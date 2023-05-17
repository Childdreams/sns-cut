// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sns_rpc

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

// CuterClient is the client API for Cuter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CuterClient interface {
	Cut(ctx context.Context, opts ...grpc.CallOption) (Cuter_CutClient, error)
}

type cuterClient struct {
	cc grpc.ClientConnInterface
}

func NewCuterClient(cc grpc.ClientConnInterface) CuterClient {
	return &cuterClient{cc}
}

func (c *cuterClient) Cut(ctx context.Context, opts ...grpc.CallOption) (Cuter_CutClient, error) {
	stream, err := c.cc.NewStream(ctx, &Cuter_ServiceDesc.Streams[0], "/sns_rpc.Cuter/Cut", opts...)
	if err != nil {
		return nil, err
	}
	x := &cuterCutClient{stream}
	return x, nil
}

type Cuter_CutClient interface {
	Send(*SnsRawContentReq) error
	Recv() (*SnsReply, error)
	grpc.ClientStream
}

type cuterCutClient struct {
	grpc.ClientStream
}

func (x *cuterCutClient) Send(m *SnsRawContentReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cuterCutClient) Recv() (*SnsReply, error) {
	m := new(SnsReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CuterServer is the server API for Cuter service.
// All implementations must embed UnimplementedCuterServer
// for forward compatibility
type CuterServer interface {
	Cut(Cuter_CutServer) error
	mustEmbedUnimplementedCuterServer()
}

// UnimplementedCuterServer must be embedded to have forward compatible implementations.
type UnimplementedCuterServer struct {
}

func (UnimplementedCuterServer) Cut(Cuter_CutServer) error {
	return status.Errorf(codes.Unimplemented, "method Cut not implemented")
}
func (UnimplementedCuterServer) mustEmbedUnimplementedCuterServer() {}

// UnsafeCuterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CuterServer will
// result in compilation errors.
type UnsafeCuterServer interface {
	mustEmbedUnimplementedCuterServer()
}

func RegisterCuterServer(s grpc.ServiceRegistrar, srv CuterServer) {
	s.RegisterService(&Cuter_ServiceDesc, srv)
}

func _Cuter_Cut_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CuterServer).Cut(&cuterCutServer{stream})
}

type Cuter_CutServer interface {
	Send(*SnsReply) error
	Recv() (*SnsRawContentReq, error)
	grpc.ServerStream
}

type cuterCutServer struct {
	grpc.ServerStream
}

func (x *cuterCutServer) Send(m *SnsReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cuterCutServer) Recv() (*SnsRawContentReq, error) {
	m := new(SnsRawContentReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Cuter_ServiceDesc is the grpc.ServiceDesc for Cuter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cuter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sns_rpc.Cuter",
	HandlerType: (*CuterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Cut",
			Handler:       _Cuter_Cut_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/sns_rpc/sns.proto",
}