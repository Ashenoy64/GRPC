// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: InverseMatrix.proto

package Inverse

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
	InverseMatrix_GetInverseMatrix_FullMethodName = "/InverseMatrix/GetInverseMatrix"
)

// InverseMatrixClient is the client API for InverseMatrix service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InverseMatrixClient interface {
	GetInverseMatrix(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[InverseMatrixRequest, InverseMatrixResponse], error)
}

type inverseMatrixClient struct {
	cc grpc.ClientConnInterface
}

func NewInverseMatrixClient(cc grpc.ClientConnInterface) InverseMatrixClient {
	return &inverseMatrixClient{cc}
}

func (c *inverseMatrixClient) GetInverseMatrix(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[InverseMatrixRequest, InverseMatrixResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &InverseMatrix_ServiceDesc.Streams[0], InverseMatrix_GetInverseMatrix_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[InverseMatrixRequest, InverseMatrixResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type InverseMatrix_GetInverseMatrixClient = grpc.BidiStreamingClient[InverseMatrixRequest, InverseMatrixResponse]

// InverseMatrixServer is the server API for InverseMatrix service.
// All implementations must embed UnimplementedInverseMatrixServer
// for forward compatibility.
type InverseMatrixServer interface {
	GetInverseMatrix(grpc.BidiStreamingServer[InverseMatrixRequest, InverseMatrixResponse]) error
	mustEmbedUnimplementedInverseMatrixServer()
}

// UnimplementedInverseMatrixServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInverseMatrixServer struct{}

func (UnimplementedInverseMatrixServer) GetInverseMatrix(grpc.BidiStreamingServer[InverseMatrixRequest, InverseMatrixResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetInverseMatrix not implemented")
}
func (UnimplementedInverseMatrixServer) mustEmbedUnimplementedInverseMatrixServer() {}
func (UnimplementedInverseMatrixServer) testEmbeddedByValue()                       {}

// UnsafeInverseMatrixServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InverseMatrixServer will
// result in compilation errors.
type UnsafeInverseMatrixServer interface {
	mustEmbedUnimplementedInverseMatrixServer()
}

func RegisterInverseMatrixServer(s grpc.ServiceRegistrar, srv InverseMatrixServer) {
	// If the following call pancis, it indicates UnimplementedInverseMatrixServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InverseMatrix_ServiceDesc, srv)
}

func _InverseMatrix_GetInverseMatrix_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(InverseMatrixServer).GetInverseMatrix(&grpc.GenericServerStream[InverseMatrixRequest, InverseMatrixResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type InverseMatrix_GetInverseMatrixServer = grpc.BidiStreamingServer[InverseMatrixRequest, InverseMatrixResponse]

// InverseMatrix_ServiceDesc is the grpc.ServiceDesc for InverseMatrix service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InverseMatrix_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "InverseMatrix",
	HandlerType: (*InverseMatrixServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetInverseMatrix",
			Handler:       _InverseMatrix_GetInverseMatrix_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "InverseMatrix.proto",
}
