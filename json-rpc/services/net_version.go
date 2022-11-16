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

package services

import (
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/streamingfast/logging"
)

func (e *NetService) Version(r *http.Request, args *VersionArgs, reply *NetworkID) error {
	ctx := r.Context()
	zlogger := logging.Logger(ctx, zlog)
	zlogger.Debug("net version")

	*reply = e.networkID
	return nil
}

type VersionArgs struct {
}

func (a *VersionArgs) Validate(requestInfo *rpc.RequestInfo) error {
	return nil
}