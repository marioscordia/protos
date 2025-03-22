// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: updates/updates.proto

package updatesv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Updates_GetUpdates_FullMethodName = "/updates.Updates/GetUpdates"
)

// UpdatesClient is the client API for Updates service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpdatesClient interface {
	GetUpdates(ctx context.Context, in *GetUpdatesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetUpdatesResponse], error)
}

type updatesClient struct {
	cc grpc.ClientConnInterface
}

func NewUpdatesClient(cc grpc.ClientConnInterface) UpdatesClient {
	return &updatesClient{cc}
}

func (c *updatesClient) GetUpdates(ctx context.Context, in *GetUpdatesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetUpdatesResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Updates_ServiceDesc.Streams[0], Updates_GetUpdates_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetUpdatesRequest, GetUpdatesResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Updates_GetUpdatesClient = grpc.ServerStreamingClient[GetUpdatesResponse]

// UpdatesServer is the server API for Updates service.
// All implementations must embed UnimplementedUpdatesServer
// for forward compatibility.
type UpdatesServer interface {
	GetUpdates(*GetUpdatesRequest, grpc.ServerStreamingServer[GetUpdatesResponse]) error
	mustEmbedUnimplementedUpdatesServer()
}

// UnimplementedUpdatesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUpdatesServer struct{}

func (UnimplementedUpdatesServer) GetUpdates(*GetUpdatesRequest, grpc.ServerStreamingServer[GetUpdatesResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetUpdates not implemented")
}
func (UnimplementedUpdatesServer) mustEmbedUnimplementedUpdatesServer() {}
func (UnimplementedUpdatesServer) testEmbeddedByValue()                 {}

// UnsafeUpdatesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpdatesServer will
// result in compilation errors.
type UnsafeUpdatesServer interface {
	mustEmbedUnimplementedUpdatesServer()
}

func RegisterUpdatesServer(s grpc.ServiceRegistrar, srv UpdatesServer) {
	// If the following call pancis, it indicates UnimplementedUpdatesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Updates_ServiceDesc, srv)
}

func _Updates_GetUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetUpdatesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UpdatesServer).GetUpdates(m, &grpc.GenericServerStream[GetUpdatesRequest, GetUpdatesResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Updates_GetUpdatesServer = grpc.ServerStreamingServer[GetUpdatesResponse]

// Updates_ServiceDesc is the grpc.ServiceDesc for Updates service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Updates_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "updates.Updates",
	HandlerType: (*UpdatesServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUpdates",
			Handler:       _Updates_GetUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "updates/updates.proto",
}
