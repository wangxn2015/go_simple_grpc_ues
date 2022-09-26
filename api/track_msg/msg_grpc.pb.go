// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: msg.proto

package track_msg

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

// UEsClient is the client API for UEs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UEsClient interface {
	// rpc CreateUser (users.CreateUserRequest) returns (users.CreateUserResult);
	GetUEs(ctx context.Context, in *UERequest, opts ...grpc.CallOption) (UEs_GetUEsClient, error)
}

type uEsClient struct {
	cc grpc.ClientConnInterface
}

func NewUEsClient(cc grpc.ClientConnInterface) UEsClient {
	return &uEsClient{cc}
}

func (c *uEsClient) GetUEs(ctx context.Context, in *UERequest, opts ...grpc.CallOption) (UEs_GetUEsClient, error) {
	stream, err := c.cc.NewStream(ctx, &UEs_ServiceDesc.Streams[0], "/msg.UEs/GetUEs", opts...)
	if err != nil {
		return nil, err
	}
	x := &uEsGetUEsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UEs_GetUEsClient interface {
	Recv() (*UEInfo, error)
	grpc.ClientStream
}

type uEsGetUEsClient struct {
	grpc.ClientStream
}

func (x *uEsGetUEsClient) Recv() (*UEInfo, error) {
	m := new(UEInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UEsServer is the server API for UEs service.
// All implementations must embed UnimplementedUEsServer
// for forward compatibility
type UEsServer interface {
	// rpc CreateUser (users.CreateUserRequest) returns (users.CreateUserResult);
	GetUEs(*UERequest, UEs_GetUEsServer) error
	mustEmbedUnimplementedUEsServer()
}

// UnimplementedUEsServer must be embedded to have forward compatible implementations.
type UnimplementedUEsServer struct {
}

func (UnimplementedUEsServer) GetUEs(*UERequest, UEs_GetUEsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUEs not implemented")
}
func (UnimplementedUEsServer) mustEmbedUnimplementedUEsServer() {}

// UnsafeUEsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UEsServer will
// result in compilation errors.
type UnsafeUEsServer interface {
	mustEmbedUnimplementedUEsServer()
}

func RegisterUEsServer(s grpc.ServiceRegistrar, srv UEsServer) {
	s.RegisterService(&UEs_ServiceDesc, srv)
}

func _UEs_GetUEs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UERequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UEsServer).GetUEs(m, &uEsGetUEsServer{stream})
}

type UEs_GetUEsServer interface {
	Send(*UEInfo) error
	grpc.ServerStream
}

type uEsGetUEsServer struct {
	grpc.ServerStream
}

func (x *uEsGetUEsServer) Send(m *UEInfo) error {
	return x.ServerStream.SendMsg(m)
}

// UEs_ServiceDesc is the grpc.ServiceDesc for UEs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UEs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msg.UEs",
	HandlerType: (*UEsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUEs",
			Handler:       _UEs_GetUEs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "msg.proto",
}
