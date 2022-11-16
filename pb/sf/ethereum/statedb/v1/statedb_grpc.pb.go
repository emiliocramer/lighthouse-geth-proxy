// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: sf/ethereum/statedb/v1/statedb.proto

package pbstatedb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StateClient is the client API for State service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StateClient interface {
	GetBlockHeaderByHash(ctx context.Context, in *GetBlockHeaderByHashRequest, opts ...grpc.CallOption) (*GetBlockHeaderResponse, error)
	GetBlockHeaderByNumber(ctx context.Context, in *GetBlockHeaderByNumberRequest, opts ...grpc.CallOption) (*GetBlockHeaderResponse, error)
	GetLatestBlockHeader(ctx context.Context, in *GetLatestBlockHeaderRequest, opts ...grpc.CallOption) (*GetBlockHeaderResponse, error)
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	GetCode(ctx context.Context, in *GetCodeRequest, opts ...grpc.CallOption) (*GetCodeResponse, error)
	GetNonce(ctx context.Context, in *GetNonceRequest, opts ...grpc.CallOption) (*GetNonceResponse, error)
	GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*GetStateResponse, error)
}

type stateClient struct {
	cc grpc.ClientConnInterface
}

func NewStateClient(cc grpc.ClientConnInterface) StateClient {
	return &stateClient{cc}
}

func (c *stateClient) GetBlockHeaderByHash(ctx context.Context, in *GetBlockHeaderByHashRequest, opts ...grpc.CallOption) (*GetBlockHeaderResponse, error) {
	out := new(GetBlockHeaderResponse)
	err := c.cc.Invoke(ctx, "/sf.ethereum.statedb.v1.State/GetBlockHeaderByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateClient) GetBlockHeaderByNumber(ctx context.Context, in *GetBlockHeaderByNumberRequest, opts ...grpc.CallOption) (*GetBlockHeaderResponse, error) {
	out := new(GetBlockHeaderResponse)
	err := c.cc.Invoke(ctx, "/sf.ethereum.statedb.v1.State/GetBlockHeaderByNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateClient) GetLatestBlockHeader(ctx context.Context, in *GetLatestBlockHeaderRequest, opts ...grpc.CallOption) (*GetBlockHeaderResponse, error) {
	out := new(GetBlockHeaderResponse)
	err := c.cc.Invoke(ctx, "/sf.ethereum.statedb.v1.State/GetLatestBlockHeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/sf.ethereum.statedb.v1.State/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateClient) GetCode(ctx context.Context, in *GetCodeRequest, opts ...grpc.CallOption) (*GetCodeResponse, error) {
	out := new(GetCodeResponse)
	err := c.cc.Invoke(ctx, "/sf.ethereum.statedb.v1.State/GetCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateClient) GetNonce(ctx context.Context, in *GetNonceRequest, opts ...grpc.CallOption) (*GetNonceResponse, error) {
	out := new(GetNonceResponse)
	err := c.cc.Invoke(ctx, "/sf.ethereum.statedb.v1.State/GetNonce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateClient) GetState(ctx context.Context, in *GetStateRequest, opts ...grpc.CallOption) (*GetStateResponse, error) {
	out := new(GetStateResponse)
	err := c.cc.Invoke(ctx, "/sf.ethereum.statedb.v1.State/GetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StateServer is the server API for State service.
// All implementations should embed UnimplementedStateServer
// for forward compatibility
type StateServer interface {
	GetBlockHeaderByHash(context.Context, *GetBlockHeaderByHashRequest) (*GetBlockHeaderResponse, error)
	GetBlockHeaderByNumber(context.Context, *GetBlockHeaderByNumberRequest) (*GetBlockHeaderResponse, error)
	GetLatestBlockHeader(context.Context, *GetLatestBlockHeaderRequest) (*GetBlockHeaderResponse, error)
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	GetCode(context.Context, *GetCodeRequest) (*GetCodeResponse, error)
	GetNonce(context.Context, *GetNonceRequest) (*GetNonceResponse, error)
	GetState(context.Context, *GetStateRequest) (*GetStateResponse, error)
}

// UnimplementedStateServer should be embedded to have forward compatible implementations.
type UnimplementedStateServer struct {
}

func (UnimplementedStateServer) GetBlockHeaderByHash(context.Context, *GetBlockHeaderByHashRequest) (*GetBlockHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeaderByHash not implemented")
}
func (UnimplementedStateServer) GetBlockHeaderByNumber(context.Context, *GetBlockHeaderByNumberRequest) (*GetBlockHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeaderByNumber not implemented")
}
func (UnimplementedStateServer) GetLatestBlockHeader(context.Context, *GetLatestBlockHeaderRequest) (*GetBlockHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestBlockHeader not implemented")
}
func (UnimplementedStateServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedStateServer) GetCode(context.Context, *GetCodeRequest) (*GetCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCode not implemented")
}
func (UnimplementedStateServer) GetNonce(context.Context, *GetNonceRequest) (*GetNonceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNonce not implemented")
}
func (UnimplementedStateServer) GetState(context.Context, *GetStateRequest) (*GetStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}

// UnsafeStateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StateServer will
// result in compilation errors.
type UnsafeStateServer interface {
	mustEmbedUnimplementedStateServer()
}

func RegisterStateServer(s grpc.ServiceRegistrar, srv StateServer) {
	s.RegisterService(&State_ServiceDesc, srv)
}

func _State_GetBlockHeaderByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockHeaderByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateServer).GetBlockHeaderByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sf.ethereum.statedb.v1.State/GetBlockHeaderByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateServer).GetBlockHeaderByHash(ctx, req.(*GetBlockHeaderByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _State_GetBlockHeaderByNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockHeaderByNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateServer).GetBlockHeaderByNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sf.ethereum.statedb.v1.State/GetBlockHeaderByNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateServer).GetBlockHeaderByNumber(ctx, req.(*GetBlockHeaderByNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _State_GetLatestBlockHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestBlockHeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateServer).GetLatestBlockHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sf.ethereum.statedb.v1.State/GetLatestBlockHeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateServer).GetLatestBlockHeader(ctx, req.(*GetLatestBlockHeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _State_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sf.ethereum.statedb.v1.State/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _State_GetCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateServer).GetCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sf.ethereum.statedb.v1.State/GetCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateServer).GetCode(ctx, req.(*GetCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _State_GetNonce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNonceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateServer).GetNonce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sf.ethereum.statedb.v1.State/GetNonce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateServer).GetNonce(ctx, req.(*GetNonceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _State_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sf.ethereum.statedb.v1.State/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateServer).GetState(ctx, req.(*GetStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// State_ServiceDesc is the grpc.ServiceDesc for State service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var State_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sf.ethereum.statedb.v1.State",
	HandlerType: (*StateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlockHeaderByHash",
			Handler:    _State_GetBlockHeaderByHash_Handler,
		},
		{
			MethodName: "GetBlockHeaderByNumber",
			Handler:    _State_GetBlockHeaderByNumber_Handler,
		},
		{
			MethodName: "GetLatestBlockHeader",
			Handler:    _State_GetLatestBlockHeader_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _State_GetBalance_Handler,
		},
		{
			MethodName: "GetCode",
			Handler:    _State_GetCode_Handler,
		},
		{
			MethodName: "GetNonce",
			Handler:    _State_GetNonce_Handler,
		},
		{
			MethodName: "GetState",
			Handler:    _State_GetState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sf/ethereum/statedb/v1/statedb.proto",
}
