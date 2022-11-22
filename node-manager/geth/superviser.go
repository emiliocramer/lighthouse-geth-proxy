// Copyright 2021 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package geth

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ShinyTrinkets/overseer"
	nodemanager "github.com/streamingfast/firehose-ethereum/node-manager"
	nodeManager "github.com/streamingfast/node-manager"
	logplugin "github.com/streamingfast/node-manager/log_plugin"
	"github.com/streamingfast/node-manager/superviser"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var enodeRegexp = regexp.MustCompile(`enode://([a-f0-9]*)@.*$`)

type Superviser struct {
	*nodemanager.Superviser

	backupMutex         sync.Mutex
	infoMutex           sync.Mutex
	binary              string
	arguments           []string
	dataDir             string
	ipcFilePath         string
	lastBlockSeen       uint64
	enodeStr            string
	connectedPeers      []string
	headBlockUpdateFunc nodeManager.HeadBlockUpdater
}

func (s *Superviser) GetName() string {
	return "geth"
}

func NewGethSuperviser(
	binary string,
	dataDir string,
	nodeIPCPath string,
	arguments []string,
	debugDeepMind bool,
	headBlockUpdateFunc nodeManager.HeadBlockUpdater,
	enforcePeersStr string,
	logToZap bool,
	appLogger *zap.Logger,
	nodelogger *zap.Logger,
) (*Superviser, error) {
	// Ensure process manager line buffer is large enough (50 MiB) for our Deep Mind instrumentation outputting lot's of text.
	overseer.DEFAULT_LINE_BUFFER_SIZE = 50 * 1024 * 1024

	gethSuperviser := &Superviser{
		Superviser: &nodemanager.Superviser{
			Superviser: superviser.New(appLogger, binary, arguments),
			Logger:     appLogger,
		},
		binary:              binary,
		arguments:           arguments,
		dataDir:             dataDir,
		ipcFilePath:         nodeIPCPath,
		connectedPeers:      []string{},
		headBlockUpdateFunc: headBlockUpdateFunc,
	}

	gethSuperviser.RegisterLogPlugin(logplugin.LogPluginFunc(gethSuperviser.lastBlockSeenLogPlugin))

	if logToZap {
		gethSuperviser.RegisterLogPlugin(nodemanager.NewGethToZapLogPlugin(debugDeepMind, nodelogger))
	} else {
		gethSuperviser.RegisterLogPlugin(logplugin.NewToConsoleLogPlugin(debugDeepMind))
	}

	if enforcePeersStr != "" {
		enforcedPeers := strings.Split(enforcePeersStr, ",")
		appLogger.Info("enforcing peers by dns", zap.Strings("peers", enforcedPeers))
		go gethSuperviser.EnsurePeersByDNS(enforcedPeers)
	}

	appLogger.Info("created geth superviser", zap.Object("superviser", gethSuperviser))
	return gethSuperviser, nil
}

func (s *Superviser) GetCommand() string {
	return s.binary + " " + strings.Join(s.arguments, " ")
}

func (s *Superviser) LastSeenBlockNum() uint64 {
	return s.lastBlockSeen
}

func (s *Superviser) ServerID() (string, error) {
	id := s.enodeStr
	if id != "" {
		return id, nil
	}
	return "", fmt.Errorf("enode not fetched yet")
}

func (s *Superviser) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("binary", s.binary)
	enc.AddArray("arguments", nodemanager.StringArray(s.arguments))
	enc.AddString("data_dir", s.dataDir)
	enc.AddString("ipc_file_path", s.ipcFilePath)
	enc.AddUint64("last_block_seen", s.lastBlockSeen)
	enc.AddString("enode_str", s.enodeStr)

	return nil
}

func (s *Superviser) lastBlockSeenLogPlugin(line string) {
	switch {
	case strings.HasPrefix(line, "DMLOG FINALIZE_BLOCK"):
		line = strings.TrimSpace(strings.TrimPrefix(line, "DMLOG FINALIZE_BLOCK"))
	case strings.HasPrefix(line, "FIRE FINALIZE_BLOCK"):
		line = strings.TrimSpace(strings.TrimPrefix(line, "FIRE FINALIZE_BLOCK"))
	default:
		return
	}

	blockNum, err := strconv.ParseUint(line, 10, 64)
	if err != nil {
		s.Logger.Error("unable to extract last block num", zap.String("line", line), zap.Error(err))
		return
	}

	//metrics.SetHeadBlockNumber(blockNum)
	s.lastBlockSeen = blockNum
}

// AddPeer sends a command through IPC socket to connect geth to the given peer
func (s *Superviser) AddPeer(peer string) error {

	for _, peerPrefix := range append(s.connectedPeers, s.enodeStr[0:19]) { //connected or ourself
		if strings.Contains(peer, peerPrefix) {
			return nil
		}
	}

	resp, err := s.sendGethCommand(fmt.Sprintf(`{"jsonrpc":"2.0","method":"admin_addPeer","params":["%s"],"id":1}`, peer))
	if err != nil {
		return err
	}
	if !gjson.Get(resp, "result").Bool() {
		return fmt.Errorf("result not true")
	}
	return nil
}

func (s *Superviser) sendGethCommand(cmd string) (string, error) {
	c, err := net.Dial("unix", s.ipcFilePath)
	if err != nil {
		return "", err
	}
	defer c.Close()

	_, err = c.Write([]byte(cmd))
	if err != nil {
		return "", err
	}

	resp, err := readString(c)
	return resp, err
}

func (s *Superviser) setEnodeStr(enodeStr string) error {
	ipAddr := getIPAddress()
	if ipAddr == "" {
		return fmt.Errorf("cannot find local IP address")
	}

	s.infoMutex.Lock()
	defer s.infoMutex.Unlock()
	fixedEnodeStr := enodeRegexp.ReplaceAllString(enodeStr, fmt.Sprintf(`enode://${1}@%s:30303`, ipAddr))
	if fixedEnodeStr != "" && s.enodeStr != fixedEnodeStr {
		s.enodeStr = fixedEnodeStr
	}
	return nil
}

// EnsurePeersByDNS periodically checks IP addresses on the given FQDNs,
// calls /v1/server_id on port 8080 (or other if specified in hostname) and adds them as peers
// wantedPeersHostnames can point to the headless service name in k8s
func (s *Superviser) EnsurePeersByDNS(wantedPeersHostnames []string) {
	for {
		time.Sleep(10 * time.Second)
		if !s.IsRunning() {
			s.Logger.Info("supervisor not running, will try to add peers later")
			continue
		}
		if len(s.enodeStr) < 20 {
			s.Logger.Info("wrong enode string will retry", zap.String("enode", s.enodeStr))
			continue
		}

		var allEnodes []string
		for _, hostname := range wantedPeersHostnames {
			enodes := s.getEnodesFromPeers(hostname)
			s.Logger.Debug("got enode", zap.String("hostname", hostname), zap.Strings("enodes", enodes))
			allEnodes = append(allEnodes, enodes...)
		}

		for _, enode := range allEnodes {
			if err := s.AddPeer(enode); err != nil {
				s.Logger.Warn("cannot add peer", zap.String("enode", enode))
			}
		}
	}
}

func (s *Superviser) getEnodesFromPeers(hostname string) []string {
	port := "8080"
	if splitted := strings.Split(hostname, ":"); len(splitted) == 2 {
		port = splitted[1]
		hostname = splitted[0]
	}
	ips, err := net.LookupIP(hostname)
	if err != nil {
		s.Logger.Warn("cannot get IP for hostname", zap.Error(err), zap.String("hostname", hostname))
		return nil
	}
	var enodes []string
	for _, ip := range ips {
		enodeAddr, err := httpGet(fmt.Sprintf("http://%s:%s/v1/server_id", ip, port))
		if err != nil {
			s.Logger.Warn("error getting enode string from IP", zap.Stringer("ip", ip))
			continue
		}
		if !strings.HasPrefix(enodeAddr, "enode://") {
			s.Logger.Warn("got invalid enode string from IP", zap.Stringer("ip", ip), zap.String("enode", enodeAddr))
			continue
		}
		enodes = append(enodes, enodeAddr)
	}
	return enodes
}

// cannot use ReadAll on an IPC socket
func readString(r io.Reader) (string, error) {
	br := bufio.NewReader(r)
	var line string
	for {
		l, err := br.ReadString('\n')
		if len(l) > 0 {
			line += l
		}
		switch err {
		case bufio.ErrBufferFull:
			continue
		case io.EOF, nil:
			return line, nil
		default:
			return "", err
		}
	}
}

func hex2int(hexStr string) int64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseInt(cleaned, 16, 64)
	return result
}

func hex2uint(hexStr string) uint64 {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hexStr, "0x", "", -1)

	// base 16 for hexadecimal
	result, _ := strconv.ParseUint(cleaned, 16, 64)
	return result
}

func hex2string(hexStr string) string {
	// remove 0x suffix if found in the input string
	return strings.Replace(hexStr, "0x", "", -1)
}

func getIPAddress() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip.IsGlobalUnicast() {
				return ip.String()
			}
		}
	}
	return ""
}

func httpGet(addr string) (string, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
