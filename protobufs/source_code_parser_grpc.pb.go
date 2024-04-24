// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: protobufs/source_code_parser.proto

package protobufs

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

const (
	SourceCodeParser_SyntaxParse_FullMethodName = "/protobufs.SourceCodeParser/SyntaxParse"
)

// SourceCodeParserClient is the client API for SourceCodeParser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SourceCodeParserClient interface {
	SyntaxParse(ctx context.Context, opts ...grpc.CallOption) (SourceCodeParser_SyntaxParseClient, error)
}

type sourceCodeParserClient struct {
	cc grpc.ClientConnInterface
}

func NewSourceCodeParserClient(cc grpc.ClientConnInterface) SourceCodeParserClient {
	return &sourceCodeParserClient{cc}
}

func (c *sourceCodeParserClient) SyntaxParse(ctx context.Context, opts ...grpc.CallOption) (SourceCodeParser_SyntaxParseClient, error) {
	stream, err := c.cc.NewStream(ctx, &SourceCodeParser_ServiceDesc.Streams[0], SourceCodeParser_SyntaxParse_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &sourceCodeParserSyntaxParseClient{stream}
	return x, nil
}

type SourceCodeParser_SyntaxParseClient interface {
	Send(*SyntaxParseRequest) error
	Recv() (*SyntaxParseResponse, error)
	grpc.ClientStream
}

type sourceCodeParserSyntaxParseClient struct {
	grpc.ClientStream
}

func (x *sourceCodeParserSyntaxParseClient) Send(m *SyntaxParseRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sourceCodeParserSyntaxParseClient) Recv() (*SyntaxParseResponse, error) {
	m := new(SyntaxParseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SourceCodeParserServer is the server API for SourceCodeParser service.
// All implementations must embed UnimplementedSourceCodeParserServer
// for forward compatibility
type SourceCodeParserServer interface {
	SyntaxParse(SourceCodeParser_SyntaxParseServer) error
	mustEmbedUnimplementedSourceCodeParserServer()
}

// UnimplementedSourceCodeParserServer must be embedded to have forward compatible implementations.
type UnimplementedSourceCodeParserServer struct {
}

func (UnimplementedSourceCodeParserServer) SyntaxParse(SourceCodeParser_SyntaxParseServer) error {
	return status.Errorf(codes.Unimplemented, "method SyntaxParse not implemented")
}
func (UnimplementedSourceCodeParserServer) mustEmbedUnimplementedSourceCodeParserServer() {}

// UnsafeSourceCodeParserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SourceCodeParserServer will
// result in compilation errors.
type UnsafeSourceCodeParserServer interface {
	mustEmbedUnimplementedSourceCodeParserServer()
}

func RegisterSourceCodeParserServer(s grpc.ServiceRegistrar, srv SourceCodeParserServer) {
	s.RegisterService(&SourceCodeParser_ServiceDesc, srv)
}

func _SourceCodeParser_SyntaxParse_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SourceCodeParserServer).SyntaxParse(&sourceCodeParserSyntaxParseServer{stream})
}

type SourceCodeParser_SyntaxParseServer interface {
	Send(*SyntaxParseResponse) error
	Recv() (*SyntaxParseRequest, error)
	grpc.ServerStream
}

type sourceCodeParserSyntaxParseServer struct {
	grpc.ServerStream
}

func (x *sourceCodeParserSyntaxParseServer) Send(m *SyntaxParseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sourceCodeParserSyntaxParseServer) Recv() (*SyntaxParseRequest, error) {
	m := new(SyntaxParseRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SourceCodeParser_ServiceDesc is the grpc.ServiceDesc for SourceCodeParser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SourceCodeParser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobufs.SourceCodeParser",
	HandlerType: (*SourceCodeParserServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SyntaxParse",
			Handler:       _SourceCodeParser_SyntaxParse_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protobufs/source_code_parser.proto",
}
