// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: idgen.proto

package gen

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

// IdGenClient is the client API for IdGen service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdGenClient interface {
	Gen(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*IdResponse, error)
}

type idGenClient struct {
	cc grpc.ClientConnInterface
}

func NewIdGenClient(cc grpc.ClientConnInterface) IdGenClient {
	return &idGenClient{cc}
}

func (c *idGenClient) Gen(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*IdResponse, error) {
	out := new(IdResponse)
	err := c.cc.Invoke(ctx, "/idgen.IdGen/Gen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdGenServer is the server API for IdGen service.
// All implementations must embed UnimplementedIdGenServer
// for forward compatibility
type IdGenServer interface {
	Gen(context.Context, *IdRequest) (*IdResponse, error)
	mustEmbedUnimplementedIdGenServer()
}

// UnimplementedIdGenServer must be embedded to have forward compatible implementations.
type UnimplementedIdGenServer struct {
}

func (UnimplementedIdGenServer) Gen(context.Context, *IdRequest) (*IdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gen not implemented")
}
func (UnimplementedIdGenServer) mustEmbedUnimplementedIdGenServer() {}

// UnsafeIdGenServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdGenServer will
// result in compilation errors.
type UnsafeIdGenServer interface {
	mustEmbedUnimplementedIdGenServer()
}

func RegisterIdGenServer(s grpc.ServiceRegistrar, srv IdGenServer) {
	s.RegisterService(&IdGen_ServiceDesc, srv)
}

func _IdGen_Gen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdGenServer).Gen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idgen.IdGen/Gen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdGenServer).Gen(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IdGen_ServiceDesc is the grpc.ServiceDesc for IdGen service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdGen_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "idgen.IdGen",
	HandlerType: (*IdGenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Gen",
			Handler:    _IdGen_Gen_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idgen.proto",
}
