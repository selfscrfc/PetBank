// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: Accounts/AccountsService.proto

package Accounts

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

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	GetDetail(ctx context.Context, in *GetDetailsRequest, opts ...grpc.CallOption) (*GetDetailsResponse, error)
	GetTransactionsHistory(ctx context.Context, in *GetTransactionsHistoryRequest, opts ...grpc.CallOption) (*GetTransactionsHistoryResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Block(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockResponse, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) GetDetail(ctx context.Context, in *GetDetailsRequest, opts ...grpc.CallOption) (*GetDetailsResponse, error) {
	out := new(GetDetailsResponse)
	err := c.cc.Invoke(ctx, "/accounts.AccountService/GetDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetTransactionsHistory(ctx context.Context, in *GetTransactionsHistoryRequest, opts ...grpc.CallOption) (*GetTransactionsHistoryResponse, error) {
	out := new(GetTransactionsHistoryResponse)
	err := c.cc.Invoke(ctx, "/accounts.AccountService/GetTransactionsHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/accounts.AccountService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/accounts.AccountService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Block(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockResponse, error) {
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, "/accounts.AccountService/Block", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations must embed UnimplementedAccountServiceServer
// for forward compatibility
type AccountServiceServer interface {
	GetDetail(context.Context, *GetDetailsRequest) (*GetDetailsResponse, error)
	GetTransactionsHistory(context.Context, *GetTransactionsHistoryRequest) (*GetTransactionsHistoryResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Block(context.Context, *BlockRequest) (*BlockResponse, error)
	mustEmbedUnimplementedAccountServiceServer()
}

// UnimplementedAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) GetDetail(context.Context, *GetDetailsRequest) (*GetDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetail not implemented")
}
func (UnimplementedAccountServiceServer) GetTransactionsHistory(context.Context, *GetTransactionsHistoryRequest) (*GetTransactionsHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionsHistory not implemented")
}
func (UnimplementedAccountServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAccountServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAccountServiceServer) Block(context.Context, *BlockRequest) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Block not implemented")
}
func (UnimplementedAccountServiceServer) mustEmbedUnimplementedAccountServiceServer() {}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_GetDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.AccountService/GetDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetDetail(ctx, req.(*GetDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetTransactionsHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetTransactionsHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.AccountService/GetTransactionsHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetTransactionsHistory(ctx, req.(*GetTransactionsHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.AccountService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.AccountService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Block_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Block(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.AccountService/Block",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Block(ctx, req.(*BlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accounts.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDetail",
			Handler:    _AccountService_GetDetail_Handler,
		},
		{
			MethodName: "GetTransactionsHistory",
			Handler:    _AccountService_GetTransactionsHistory_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AccountService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AccountService_Delete_Handler,
		},
		{
			MethodName: "Block",
			Handler:    _AccountService_Block_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Accounts/AccountsService.proto",
}
