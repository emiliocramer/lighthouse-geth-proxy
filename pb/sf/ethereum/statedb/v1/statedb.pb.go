// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: sf/ethereum/statedb/v1/statedb.proto

package pbstatedb

import (
	v2 "github.com/streamingfast/firehose-ethereum/types/pb/sf/ethereum/type/v2"
	v1 "github.com/streamingfast/pbgo/sf/bstream/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetBlockHeaderByHashRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *GetBlockHeaderByHashRequest) Reset() {
	*x = GetBlockHeaderByHashRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockHeaderByHashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockHeaderByHashRequest) ProtoMessage() {}

func (x *GetBlockHeaderByHashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockHeaderByHashRequest.ProtoReflect.Descriptor instead.
func (*GetBlockHeaderByHashRequest) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_statedb_v1_statedb_proto_rawDescGZIP(), []int{0}
}

func (x *GetBlockHeaderByHashRequest) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type GetBlockHeaderByNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number uint64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *GetBlockHeaderByNumberRequest) Reset() {
	*x = GetBlockHeaderByNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockHeaderByNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockHeaderByNumberRequest) ProtoMessage() {}

func (x *GetBlockHeaderByNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockHeaderByNumberRequest.ProtoReflect.Descriptor instead.
func (*GetBlockHeaderByNumberRequest) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_statedb_v1_statedb_proto_rawDescGZIP(), []int{1}
}

func (x *GetBlockHeaderByNumberRequest) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type GetBlockHeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockRef *v1.BlockRef `protobuf:"bytes,1,opt,name=block_ref,json=blockRef,proto3" json:"block_ref,omitempty"`
}

func (x *GetBlockHeaderRequest) Reset() {
	*x = GetBlockHeaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockHeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockHeaderRequest) ProtoMessage() {}

func (x *GetBlockHeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockHeaderRequest.ProtoReflect.Descriptor instead.
func (*GetBlockHeaderRequest) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_statedb_v1_statedb_proto_rawDescGZIP(), []int{2}
}

func (x *GetBlockHeaderRequest) GetBlockRef() *v1.BlockRef {
	if x != nil {
		return x.BlockRef
	}
	return nil
}

type GetLatestBlockHeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLatestBlockHeaderRequest) Reset() {
	*x = GetLatestBlockHeaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestBlockHeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestBlockHeaderRequest) ProtoMessage() {}

func (x *GetLatestBlockHeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestBlockHeaderRequest.ProtoReflect.Descriptor instead.
func (*GetLatestBlockHeaderRequest) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_statedb_v1_statedb_proto_rawDescGZIP(), []int{3}
}

type GetBlockHeaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *v2.BlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *GetBlockHeaderResponse) Reset() {
	*x = GetBlockHeaderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockHeaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockHeaderResponse) ProtoMessage() {}

func (x *GetBlockHeaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockHeaderResponse.ProtoReflect.Descriptor instead.
func (*GetBlockHeaderResponse) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_statedb_v1_statedb_proto_rawDescGZIP(), []int{4}
}

func (x *GetBlockHeaderResponse) GetHeader() *v2.BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

type GetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  []byte       `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	BlockRef *v1.BlockRef `protobuf:"bytes,2,opt,name=block_ref,json=blockRef,proto3" json:"block_ref,omitempty"`
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_statedb_v1_statedb_proto_rawDescGZIP(), []int{5}
}

func (x *GetBalanceRequest) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *GetBalanceRequest) GetBlockRef() *v1.BlockRef {
	if x != nil {
		return x.BlockRef
	}
	return nil
}

type GetBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Balance represents the Wei value in hexadecimal format that can be used to
	// create a `math.BigInt` instance.
	Balance []byte `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *GetBalanceResponse) Reset() {
	*x = GetBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceResponse) ProtoMessage() {}

func (x *GetBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_statedb_v1_statedb_proto_rawDescGZIP(), []int{6}
}

func (x *GetBalanceResponse) GetBalance() []byte {
	if x != nil {
		return x.Balance
	}
	return nil
}

type GetCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  []byte       `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	BlockRef *v1.BlockRef `protobuf:"bytes,2,opt,name=block_ref,json=blockRef,proto3" json:"block_ref,omitempty"`
}

func (x *GetCodeRequest) Reset() {
	*x = GetCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCodeRequest) ProtoMessage() {}

func (x *GetCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCodeRequest.ProtoReflect.Descriptor instead.
func (*GetCodeRequest) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_statedb_v1_statedb_proto_rawDescGZIP(), []int{7}
}

func (x *GetCodeRequest) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *GetCodeRequest) GetBlockRef() *v1.BlockRef {
	if x != nil {
		return x.BlockRef
	}
	return nil
}

type GetCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Balance represents the accoint's set bytecode
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *GetCodeResponse) Reset() {
	*x = GetCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCodeResponse) ProtoMessage() {}

func (x *GetCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCodeResponse.ProtoReflect.Descriptor instead.
func (*GetCodeResponse) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_statedb_v1_statedb_proto_rawDescGZIP(), []int{8}
}

func (x *GetCodeResponse) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type GetNonceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  []byte       `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	BlockRef *v1.BlockRef `protobuf:"bytes,2,opt,name=block_ref,json=blockRef,proto3" json:"block_ref,omitempty"`
}

func (x *GetNonceRequest) Reset() {
	*x = GetNonceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNonceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNonceRequest) ProtoMessage() {}

func (x *GetNonceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNonceRequest.ProtoReflect.Descriptor instead.
func (*GetNonceRequest) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_statedb_v1_statedb_proto_rawDescGZIP(), []int{9}
}

func (x *GetNonceRequest) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *GetNonceRequest) GetBlockRef() *v1.BlockRef {
	if x != nil {
		return x.BlockRef
	}
	return nil
}

type GetNonceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce uint64 `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *GetNonceResponse) Reset() {
	*x = GetNonceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNonceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNonceResponse) ProtoMessage() {}

func (x *GetNonceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNonceResponse.ProtoReflect.Descriptor instead.
func (*GetNonceResponse) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_statedb_v1_statedb_proto_rawDescGZIP(), []int{10}
}

func (x *GetNonceResponse) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

type GetStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  []byte       `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Key      []byte       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	BlockRef *v1.BlockRef `protobuf:"bytes,3,opt,name=block_ref,json=blockRef,proto3" json:"block_ref,omitempty"`
}

func (x *GetStateRequest) Reset() {
	*x = GetStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStateRequest) ProtoMessage() {}

func (x *GetStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStateRequest.ProtoReflect.Descriptor instead.
func (*GetStateRequest) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_statedb_v1_statedb_proto_rawDescGZIP(), []int{11}
}

func (x *GetStateRequest) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *GetStateRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *GetStateRequest) GetBlockRef() *v1.BlockRef {
	if x != nil {
		return x.BlockRef
	}
	return nil
}

type GetStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetStateResponse) Reset() {
	*x = GetStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStateResponse) ProtoMessage() {}

func (x *GetStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStateResponse.ProtoReflect.Descriptor instead.
func (*GetStateResponse) Descriptor() ([]byte, []int) {
	return file_sf_ethereum_statedb_v1_statedb_proto_rawDescGZIP(), []int{12}
}

func (x *GetStateResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_sf_ethereum_statedb_v1_statedb_proto protoreflect.FileDescriptor

var file_sf_ethereum_statedb_v1_statedb_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x66, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1b,
	0x73, 0x66, 0x2f, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x73, 0x66, 0x2f,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x32,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x1b, 0x47,
	0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x79, 0x48,
	0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x37,
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x4d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x34, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x66, 0x2e, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x66, 0x52, 0x08, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x66, 0x22, 0x1d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x52, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x38, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x63, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x66,
	0x2e, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x66, 0x52, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x66, 0x22, 0x2e,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x60,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x73, 0x66, 0x2e, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x66, 0x52, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x66,
	0x22, 0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x61, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x73, 0x66, 0x2e, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x66, 0x52, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x66,
	0x22, 0x28, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x73, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x09, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73,
	0x66, 0x2e, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x66, 0x52, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x66, 0x22,
	0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x81, 0x06, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x7b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x12, 0x33, 0x2e, 0x73, 0x66, 0x2e, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x7f,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x35, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x7b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x33, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x73,
	0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x29, 0x2e, 0x73, 0x66, 0x2e,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x2e, 0x73,
	0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x66, 0x2e, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x73, 0x66, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4b, 0x5a, 0x49, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x65, 0x76, 0x6d, 0x2d, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x66, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x62, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_ethereum_statedb_v1_statedb_proto_rawDescOnce sync.Once
	file_sf_ethereum_statedb_v1_statedb_proto_rawDescData = file_sf_ethereum_statedb_v1_statedb_proto_rawDesc
)

func file_sf_ethereum_statedb_v1_statedb_proto_rawDescGZIP() []byte {
	file_sf_ethereum_statedb_v1_statedb_proto_rawDescOnce.Do(func() {
		file_sf_ethereum_statedb_v1_statedb_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_ethereum_statedb_v1_statedb_proto_rawDescData)
	})
	return file_sf_ethereum_statedb_v1_statedb_proto_rawDescData
}

var file_sf_ethereum_statedb_v1_statedb_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_sf_ethereum_statedb_v1_statedb_proto_goTypes = []interface{}{
	(*GetBlockHeaderByHashRequest)(nil),   // 0: sf.ethereum.statedb.v1.GetBlockHeaderByHashRequest
	(*GetBlockHeaderByNumberRequest)(nil), // 1: sf.ethereum.statedb.v1.GetBlockHeaderByNumberRequest
	(*GetBlockHeaderRequest)(nil),         // 2: sf.ethereum.statedb.v1.GetBlockHeaderRequest
	(*GetLatestBlockHeaderRequest)(nil),   // 3: sf.ethereum.statedb.v1.GetLatestBlockHeaderRequest
	(*GetBlockHeaderResponse)(nil),        // 4: sf.ethereum.statedb.v1.GetBlockHeaderResponse
	(*GetBalanceRequest)(nil),             // 5: sf.ethereum.statedb.v1.GetBalanceRequest
	(*GetBalanceResponse)(nil),            // 6: sf.ethereum.statedb.v1.GetBalanceResponse
	(*GetCodeRequest)(nil),                // 7: sf.ethereum.statedb.v1.GetCodeRequest
	(*GetCodeResponse)(nil),               // 8: sf.ethereum.statedb.v1.GetCodeResponse
	(*GetNonceRequest)(nil),               // 9: sf.ethereum.statedb.v1.GetNonceRequest
	(*GetNonceResponse)(nil),              // 10: sf.ethereum.statedb.v1.GetNonceResponse
	(*GetStateRequest)(nil),               // 11: sf.ethereum.statedb.v1.GetStateRequest
	(*GetStateResponse)(nil),              // 12: sf.ethereum.statedb.v1.GetStateResponse
	(*v1.BlockRef)(nil),                   // 13: sf.bstream.v1.BlockRef
	(*v2.BlockHeader)(nil),                // 14: sf.ethereum.type.v2.BlockHeader
}
var file_sf_ethereum_statedb_v1_statedb_proto_depIdxs = []int32{
	13, // 0: sf.ethereum.statedb.v1.GetBlockHeaderRequest.block_ref:type_name -> sf.bstream.v1.BlockRef
	14, // 1: sf.ethereum.statedb.v1.GetBlockHeaderResponse.header:type_name -> sf.ethereum.type.v2.BlockHeader
	13, // 2: sf.ethereum.statedb.v1.GetBalanceRequest.block_ref:type_name -> sf.bstream.v1.BlockRef
	13, // 3: sf.ethereum.statedb.v1.GetCodeRequest.block_ref:type_name -> sf.bstream.v1.BlockRef
	13, // 4: sf.ethereum.statedb.v1.GetNonceRequest.block_ref:type_name -> sf.bstream.v1.BlockRef
	13, // 5: sf.ethereum.statedb.v1.GetStateRequest.block_ref:type_name -> sf.bstream.v1.BlockRef
	0,  // 6: sf.ethereum.statedb.v1.State.GetBlockHeaderByHash:input_type -> sf.ethereum.statedb.v1.GetBlockHeaderByHashRequest
	1,  // 7: sf.ethereum.statedb.v1.State.GetBlockHeaderByNumber:input_type -> sf.ethereum.statedb.v1.GetBlockHeaderByNumberRequest
	3,  // 8: sf.ethereum.statedb.v1.State.GetLatestBlockHeader:input_type -> sf.ethereum.statedb.v1.GetLatestBlockHeaderRequest
	5,  // 9: sf.ethereum.statedb.v1.State.GetBalance:input_type -> sf.ethereum.statedb.v1.GetBalanceRequest
	7,  // 10: sf.ethereum.statedb.v1.State.GetCode:input_type -> sf.ethereum.statedb.v1.GetCodeRequest
	9,  // 11: sf.ethereum.statedb.v1.State.GetNonce:input_type -> sf.ethereum.statedb.v1.GetNonceRequest
	11, // 12: sf.ethereum.statedb.v1.State.GetState:input_type -> sf.ethereum.statedb.v1.GetStateRequest
	4,  // 13: sf.ethereum.statedb.v1.State.GetBlockHeaderByHash:output_type -> sf.ethereum.statedb.v1.GetBlockHeaderResponse
	4,  // 14: sf.ethereum.statedb.v1.State.GetBlockHeaderByNumber:output_type -> sf.ethereum.statedb.v1.GetBlockHeaderResponse
	4,  // 15: sf.ethereum.statedb.v1.State.GetLatestBlockHeader:output_type -> sf.ethereum.statedb.v1.GetBlockHeaderResponse
	6,  // 16: sf.ethereum.statedb.v1.State.GetBalance:output_type -> sf.ethereum.statedb.v1.GetBalanceResponse
	8,  // 17: sf.ethereum.statedb.v1.State.GetCode:output_type -> sf.ethereum.statedb.v1.GetCodeResponse
	10, // 18: sf.ethereum.statedb.v1.State.GetNonce:output_type -> sf.ethereum.statedb.v1.GetNonceResponse
	12, // 19: sf.ethereum.statedb.v1.State.GetState:output_type -> sf.ethereum.statedb.v1.GetStateResponse
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_sf_ethereum_statedb_v1_statedb_proto_init() }
func file_sf_ethereum_statedb_v1_statedb_proto_init() {
	if File_sf_ethereum_statedb_v1_statedb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockHeaderByHashRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockHeaderByNumberRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockHeaderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestBlockHeaderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockHeaderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCodeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNonceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNonceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sf_ethereum_statedb_v1_statedb_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sf_ethereum_statedb_v1_statedb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sf_ethereum_statedb_v1_statedb_proto_goTypes,
		DependencyIndexes: file_sf_ethereum_statedb_v1_statedb_proto_depIdxs,
		MessageInfos:      file_sf_ethereum_statedb_v1_statedb_proto_msgTypes,
	}.Build()
	File_sf_ethereum_statedb_v1_statedb_proto = out.File
	file_sf_ethereum_statedb_v1_statedb_proto_rawDesc = nil
	file_sf_ethereum_statedb_v1_statedb_proto_goTypes = nil
	file_sf_ethereum_statedb_v1_statedb_proto_depIdxs = nil
}
