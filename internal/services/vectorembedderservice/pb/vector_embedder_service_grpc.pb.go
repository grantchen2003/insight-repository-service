// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: vector_embedder_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	VectorEmbedderService_GetSimilarFileComponentIds_FullMethodName                                           = "/VectorEmbedderService/GetSimilarFileComponentIds"
	VectorEmbedderService_CreateFileComponentVectorEmbeddings_FullMethodName                                  = "/VectorEmbedderService/CreateFileComponentVectorEmbeddings"
	VectorEmbedderService_DeleteFileComponentVectorEmbeddingsByRepositoryId_FullMethodName                    = "/VectorEmbedderService/DeleteFileComponentVectorEmbeddingsByRepositoryId"
	VectorEmbedderService_DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIds_FullMethodName = "/VectorEmbedderService/DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIds"
)

// VectorEmbedderServiceClient is the client API for VectorEmbedderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VectorEmbedderServiceClient interface {
	GetSimilarFileComponentIds(ctx context.Context, in *GetSimilarFileComponentIdsRequest, opts ...grpc.CallOption) (*GetSimilarFileComponentIdsResponse, error)
	CreateFileComponentVectorEmbeddings(ctx context.Context, in *CreateFileComponentVectorEmbeddingsRequest, opts ...grpc.CallOption) (*CreateFileComponentVectorEmbeddingsResponse, error)
	DeleteFileComponentVectorEmbeddingsByRepositoryId(ctx context.Context, in *DeleteFileComponentVectorEmbeddingsByRepositoryIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIds(ctx context.Context, in *DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIdsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type vectorEmbedderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVectorEmbedderServiceClient(cc grpc.ClientConnInterface) VectorEmbedderServiceClient {
	return &vectorEmbedderServiceClient{cc}
}

func (c *vectorEmbedderServiceClient) GetSimilarFileComponentIds(ctx context.Context, in *GetSimilarFileComponentIdsRequest, opts ...grpc.CallOption) (*GetSimilarFileComponentIdsResponse, error) {
	out := new(GetSimilarFileComponentIdsResponse)
	err := c.cc.Invoke(ctx, VectorEmbedderService_GetSimilarFileComponentIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vectorEmbedderServiceClient) CreateFileComponentVectorEmbeddings(ctx context.Context, in *CreateFileComponentVectorEmbeddingsRequest, opts ...grpc.CallOption) (*CreateFileComponentVectorEmbeddingsResponse, error) {
	out := new(CreateFileComponentVectorEmbeddingsResponse)
	err := c.cc.Invoke(ctx, VectorEmbedderService_CreateFileComponentVectorEmbeddings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vectorEmbedderServiceClient) DeleteFileComponentVectorEmbeddingsByRepositoryId(ctx context.Context, in *DeleteFileComponentVectorEmbeddingsByRepositoryIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, VectorEmbedderService_DeleteFileComponentVectorEmbeddingsByRepositoryId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vectorEmbedderServiceClient) DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIds(ctx context.Context, in *DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIdsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, VectorEmbedderService_DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VectorEmbedderServiceServer is the server API for VectorEmbedderService service.
// All implementations must embed UnimplementedVectorEmbedderServiceServer
// for forward compatibility
type VectorEmbedderServiceServer interface {
	GetSimilarFileComponentIds(context.Context, *GetSimilarFileComponentIdsRequest) (*GetSimilarFileComponentIdsResponse, error)
	CreateFileComponentVectorEmbeddings(context.Context, *CreateFileComponentVectorEmbeddingsRequest) (*CreateFileComponentVectorEmbeddingsResponse, error)
	DeleteFileComponentVectorEmbeddingsByRepositoryId(context.Context, *DeleteFileComponentVectorEmbeddingsByRepositoryIdRequest) (*emptypb.Empty, error)
	DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIds(context.Context, *DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIdsRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedVectorEmbedderServiceServer()
}

// UnimplementedVectorEmbedderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVectorEmbedderServiceServer struct {
}

func (UnimplementedVectorEmbedderServiceServer) GetSimilarFileComponentIds(context.Context, *GetSimilarFileComponentIdsRequest) (*GetSimilarFileComponentIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSimilarFileComponentIds not implemented")
}
func (UnimplementedVectorEmbedderServiceServer) CreateFileComponentVectorEmbeddings(context.Context, *CreateFileComponentVectorEmbeddingsRequest) (*CreateFileComponentVectorEmbeddingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFileComponentVectorEmbeddings not implemented")
}
func (UnimplementedVectorEmbedderServiceServer) DeleteFileComponentVectorEmbeddingsByRepositoryId(context.Context, *DeleteFileComponentVectorEmbeddingsByRepositoryIdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFileComponentVectorEmbeddingsByRepositoryId not implemented")
}
func (UnimplementedVectorEmbedderServiceServer) DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIds(context.Context, *DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIdsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIds not implemented")
}
func (UnimplementedVectorEmbedderServiceServer) mustEmbedUnimplementedVectorEmbedderServiceServer() {}

// UnsafeVectorEmbedderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VectorEmbedderServiceServer will
// result in compilation errors.
type UnsafeVectorEmbedderServiceServer interface {
	mustEmbedUnimplementedVectorEmbedderServiceServer()
}

func RegisterVectorEmbedderServiceServer(s grpc.ServiceRegistrar, srv VectorEmbedderServiceServer) {
	s.RegisterService(&VectorEmbedderService_ServiceDesc, srv)
}

func _VectorEmbedderService_GetSimilarFileComponentIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSimilarFileComponentIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VectorEmbedderServiceServer).GetSimilarFileComponentIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VectorEmbedderService_GetSimilarFileComponentIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VectorEmbedderServiceServer).GetSimilarFileComponentIds(ctx, req.(*GetSimilarFileComponentIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VectorEmbedderService_CreateFileComponentVectorEmbeddings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFileComponentVectorEmbeddingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VectorEmbedderServiceServer).CreateFileComponentVectorEmbeddings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VectorEmbedderService_CreateFileComponentVectorEmbeddings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VectorEmbedderServiceServer).CreateFileComponentVectorEmbeddings(ctx, req.(*CreateFileComponentVectorEmbeddingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VectorEmbedderService_DeleteFileComponentVectorEmbeddingsByRepositoryId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileComponentVectorEmbeddingsByRepositoryIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VectorEmbedderServiceServer).DeleteFileComponentVectorEmbeddingsByRepositoryId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VectorEmbedderService_DeleteFileComponentVectorEmbeddingsByRepositoryId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VectorEmbedderServiceServer).DeleteFileComponentVectorEmbeddingsByRepositoryId(ctx, req.(*DeleteFileComponentVectorEmbeddingsByRepositoryIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VectorEmbedderService_DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VectorEmbedderServiceServer).DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VectorEmbedderService_DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VectorEmbedderServiceServer).DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIds(ctx, req.(*DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VectorEmbedderService_ServiceDesc is the grpc.ServiceDesc for VectorEmbedderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VectorEmbedderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "VectorEmbedderService",
	HandlerType: (*VectorEmbedderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSimilarFileComponentIds",
			Handler:    _VectorEmbedderService_GetSimilarFileComponentIds_Handler,
		},
		{
			MethodName: "CreateFileComponentVectorEmbeddings",
			Handler:    _VectorEmbedderService_CreateFileComponentVectorEmbeddings_Handler,
		},
		{
			MethodName: "DeleteFileComponentVectorEmbeddingsByRepositoryId",
			Handler:    _VectorEmbedderService_DeleteFileComponentVectorEmbeddingsByRepositoryId_Handler,
		},
		{
			MethodName: "DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIds",
			Handler:    _VectorEmbedderService_DeleteFileComponentVectorEmbeddingsByRepositoryIdAndFileComponentIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vector_embedder_service.proto",
}
