// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/session/v1/session.proto

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

// SessionInternalAPIClient is the client API for SessionInternalAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SessionInternalAPIClient interface {
	// Save new user's session and returns auth token.
	Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error)
	// Returns user's session info by token.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Delete user's session by auth token.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type sessionInternalAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionInternalAPIClient(cc grpc.ClientConnInterface) SessionInternalAPIClient {
	return &sessionInternalAPIClient{cc}
}

func (c *sessionInternalAPIClient) Save(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, "/api.session.v1.SessionInternalAPI/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionInternalAPIClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/api.session.v1.SessionInternalAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionInternalAPIClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/api.session.v1.SessionInternalAPI/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionInternalAPIServer is the server API for SessionInternalAPI service.
// All implementations should embed UnimplementedSessionInternalAPIServer
// for forward compatibility
type SessionInternalAPIServer interface {
	// Save new user's session and returns auth token.
	Save(context.Context, *SaveRequest) (*SaveResponse, error)
	// Returns user's session info by token.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Delete user's session by auth token.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

// UnimplementedSessionInternalAPIServer should be embedded to have forward compatible implementations.
type UnimplementedSessionInternalAPIServer struct {
}

func (UnimplementedSessionInternalAPIServer) Save(context.Context, *SaveRequest) (*SaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedSessionInternalAPIServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSessionInternalAPIServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeSessionInternalAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SessionInternalAPIServer will
// result in compilation errors.
type UnsafeSessionInternalAPIServer interface {
	mustEmbedUnimplementedSessionInternalAPIServer()
}

func RegisterSessionInternalAPIServer(s grpc.ServiceRegistrar, srv SessionInternalAPIServer) {
	s.RegisterService(&SessionInternalAPI_ServiceDesc, srv)
}

func _SessionInternalAPI_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionInternalAPIServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.session.v1.SessionInternalAPI/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionInternalAPIServer).Save(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionInternalAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionInternalAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.session.v1.SessionInternalAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionInternalAPIServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionInternalAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionInternalAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.session.v1.SessionInternalAPI/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionInternalAPIServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SessionInternalAPI_ServiceDesc is the grpc.ServiceDesc for SessionInternalAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SessionInternalAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.session.v1.SessionInternalAPI",
	HandlerType: (*SessionInternalAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _SessionInternalAPI_Save_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SessionInternalAPI_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SessionInternalAPI_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/session/v1/session.proto",
}
