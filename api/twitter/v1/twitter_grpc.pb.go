// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/twitter/v1/twitter.proto

package pb

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

// StartClient is the client API for Start service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StartClient interface {
	TwitByAuthors(ctx context.Context, in *TwitByAuthorRequest, opts ...grpc.CallOption) (*TwitByAuthorResponse, error)
	UpdateTwit(ctx context.Context, in *UpdateTwitRequest, opts ...grpc.CallOption) (*UpdateTwitResponse, error)
	DeleteTwit(ctx context.Context, in *DeleteTwitRequest, opts ...grpc.CallOption) (*DeleteTwitResponse, error)
	CreateTwit(ctx context.Context, in *CreateTwitRequest, opts ...grpc.CallOption) (*CreateTwitResponse, error)
}

type startClient struct {
	cc grpc.ClientConnInterface
}

func NewStartClient(cc grpc.ClientConnInterface) StartClient {
	return &startClient{cc}
}

func (c *startClient) TwitByAuthors(ctx context.Context, in *TwitByAuthorRequest, opts ...grpc.CallOption) (*TwitByAuthorResponse, error) {
	out := new(TwitByAuthorResponse)
	err := c.cc.Invoke(ctx, "/api.twitter.v1.Start/TwitByAuthors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startClient) UpdateTwit(ctx context.Context, in *UpdateTwitRequest, opts ...grpc.CallOption) (*UpdateTwitResponse, error) {
	out := new(UpdateTwitResponse)
	err := c.cc.Invoke(ctx, "/api.twitter.v1.Start/UpdateTwit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startClient) DeleteTwit(ctx context.Context, in *DeleteTwitRequest, opts ...grpc.CallOption) (*DeleteTwitResponse, error) {
	out := new(DeleteTwitResponse)
	err := c.cc.Invoke(ctx, "/api.twitter.v1.Start/DeleteTwit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *startClient) CreateTwit(ctx context.Context, in *CreateTwitRequest, opts ...grpc.CallOption) (*CreateTwitResponse, error) {
	out := new(CreateTwitResponse)
	err := c.cc.Invoke(ctx, "/api.twitter.v1.Start/CreateTwit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StartServer is the server API for Start service.
// All implementations should embed UnimplementedStartServer
// for forward compatibility
type StartServer interface {
	TwitByAuthors(context.Context, *TwitByAuthorRequest) (*TwitByAuthorResponse, error)
	UpdateTwit(context.Context, *UpdateTwitRequest) (*UpdateTwitResponse, error)
	DeleteTwit(context.Context, *DeleteTwitRequest) (*DeleteTwitResponse, error)
	CreateTwit(context.Context, *CreateTwitRequest) (*CreateTwitResponse, error)
}

// UnimplementedStartServer should be embedded to have forward compatible implementations.
type UnimplementedStartServer struct {
}

func (UnimplementedStartServer) TwitByAuthors(context.Context, *TwitByAuthorRequest) (*TwitByAuthorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TwitByAuthors not implemented")
}
func (UnimplementedStartServer) UpdateTwit(context.Context, *UpdateTwitRequest) (*UpdateTwitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTwit not implemented")
}
func (UnimplementedStartServer) DeleteTwit(context.Context, *DeleteTwitRequest) (*DeleteTwitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTwit not implemented")
}
func (UnimplementedStartServer) CreateTwit(context.Context, *CreateTwitRequest) (*CreateTwitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTwit not implemented")
}

// UnsafeStartServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StartServer will
// result in compilation errors.
type UnsafeStartServer interface {
	mustEmbedUnimplementedStartServer()
}

func RegisterStartServer(s grpc.ServiceRegistrar, srv StartServer) {
	s.RegisterService(&Start_ServiceDesc, srv)
}

func _Start_TwitByAuthors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwitByAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartServer).TwitByAuthors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.twitter.v1.Start/TwitByAuthors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartServer).TwitByAuthors(ctx, req.(*TwitByAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Start_UpdateTwit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTwitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartServer).UpdateTwit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.twitter.v1.Start/UpdateTwit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartServer).UpdateTwit(ctx, req.(*UpdateTwitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Start_DeleteTwit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTwitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartServer).DeleteTwit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.twitter.v1.Start/DeleteTwit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartServer).DeleteTwit(ctx, req.(*DeleteTwitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Start_CreateTwit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTwitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartServer).CreateTwit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.twitter.v1.Start/CreateTwit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartServer).CreateTwit(ctx, req.(*CreateTwitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Start_ServiceDesc is the grpc.ServiceDesc for Start service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Start_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.twitter.v1.Start",
	HandlerType: (*StartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TwitByAuthors",
			Handler:    _Start_TwitByAuthors_Handler,
		},
		{
			MethodName: "UpdateTwit",
			Handler:    _Start_UpdateTwit_Handler,
		},
		{
			MethodName: "DeleteTwit",
			Handler:    _Start_DeleteTwit_Handler,
		},
		{
			MethodName: "CreateTwit",
			Handler:    _Start_CreateTwit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/twitter/v1/twitter.proto",
}