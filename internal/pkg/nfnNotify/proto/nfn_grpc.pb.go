// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.3
// source: internal/pkg/nfnNotify/proto/nfn.proto

package nfn

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

// NfnNotifyClient is the client API for NfnNotify service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NfnNotifyClient interface {
	Subscribe(ctx context.Context, in *SubscribeContext, opts ...grpc.CallOption) (NfnNotify_SubscribeClient, error)
}

type nfnNotifyClient struct {
	cc grpc.ClientConnInterface
}

func NewNfnNotifyClient(cc grpc.ClientConnInterface) NfnNotifyClient {
	return &nfnNotifyClient{cc}
}

func (c *nfnNotifyClient) Subscribe(ctx context.Context, in *SubscribeContext, opts ...grpc.CallOption) (NfnNotify_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &NfnNotify_ServiceDesc.Streams[0], "/nfnNotify/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &nfnNotifySubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NfnNotify_SubscribeClient interface {
	Recv() (*Notification, error)
	grpc.ClientStream
}

type nfnNotifySubscribeClient struct {
	grpc.ClientStream
}

func (x *nfnNotifySubscribeClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NfnNotifyServer is the server API for NfnNotify service.
// All implementations must embed UnimplementedNfnNotifyServer
// for forward compatibility
type NfnNotifyServer interface {
	Subscribe(*SubscribeContext, NfnNotify_SubscribeServer) error
	mustEmbedUnimplementedNfnNotifyServer()
}

// UnimplementedNfnNotifyServer must be embedded to have forward compatible implementations.
type UnimplementedNfnNotifyServer struct {
}

func (UnimplementedNfnNotifyServer) Subscribe(*SubscribeContext, NfnNotify_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedNfnNotifyServer) mustEmbedUnimplementedNfnNotifyServer() {}

// UnsafeNfnNotifyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NfnNotifyServer will
// result in compilation errors.
type UnsafeNfnNotifyServer interface {
	mustEmbedUnimplementedNfnNotifyServer()
}

func RegisterNfnNotifyServer(s grpc.ServiceRegistrar, srv NfnNotifyServer) {
	s.RegisterService(&NfnNotify_ServiceDesc, srv)
}

func _NfnNotify_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeContext)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NfnNotifyServer).Subscribe(m, &nfnNotifySubscribeServer{stream})
}

type NfnNotify_SubscribeServer interface {
	Send(*Notification) error
	grpc.ServerStream
}

type nfnNotifySubscribeServer struct {
	grpc.ServerStream
}

func (x *nfnNotifySubscribeServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

// NfnNotify_ServiceDesc is the grpc.ServiceDesc for NfnNotify service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NfnNotify_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nfnNotify",
	HandlerType: (*NfnNotifyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _NfnNotify_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/pkg/nfnNotify/proto/nfn.proto",
}