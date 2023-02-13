// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package spawn

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

// SpawnClient is the client API for Spawn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpawnClient interface {
	Generate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringResponse, error)
}

type spawnClient struct {
	cc grpc.ClientConnInterface
}

func NewSpawnClient(cc grpc.ClientConnInterface) SpawnClient {
	return &spawnClient{cc}
}

func (c *spawnClient) Generate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StringResponse, error) {
	out := new(StringResponse)
	err := c.cc.Invoke(ctx, "/spawn.Spawn/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpawnServer is the server API for Spawn service.
// All implementations must embed UnimplementedSpawnServer
// for forward compatibility
type SpawnServer interface {
	Generate(context.Context, *Empty) (*StringResponse, error)
	mustEmbedUnimplementedSpawnServer()
}

// UnimplementedSpawnServer must be embedded to have forward compatible implementations.
type UnimplementedSpawnServer struct {
}

func (UnimplementedSpawnServer) Generate(context.Context, *Empty) (*StringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedSpawnServer) mustEmbedUnimplementedSpawnServer() {}

// UnsafeSpawnServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpawnServer will
// result in compilation errors.
type UnsafeSpawnServer interface {
	mustEmbedUnimplementedSpawnServer()
}

func RegisterSpawnServer(s grpc.ServiceRegistrar, srv SpawnServer) {
	s.RegisterService(&Spawn_ServiceDesc, srv)
}

func _Spawn_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spawn.Spawn/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnServer).Generate(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Spawn_ServiceDesc is the grpc.ServiceDesc for Spawn service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Spawn_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spawn.Spawn",
	HandlerType: (*SpawnServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Spawn_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/spawn.proto",
}