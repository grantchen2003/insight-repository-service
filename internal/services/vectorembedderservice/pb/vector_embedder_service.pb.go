// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.23.3
// source: vector_embedder_service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateFileComponentVectorEmbeddingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileComponentIds []int32 `protobuf:"varint,1,rep,packed,name=file_component_ids,json=fileComponentIds,proto3" json:"file_component_ids,omitempty"`
}

func (x *CreateFileComponentVectorEmbeddingsRequest) Reset() {
	*x = CreateFileComponentVectorEmbeddingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vector_embedder_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileComponentVectorEmbeddingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileComponentVectorEmbeddingsRequest) ProtoMessage() {}

func (x *CreateFileComponentVectorEmbeddingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vector_embedder_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileComponentVectorEmbeddingsRequest.ProtoReflect.Descriptor instead.
func (*CreateFileComponentVectorEmbeddingsRequest) Descriptor() ([]byte, []int) {
	return file_vector_embedder_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFileComponentVectorEmbeddingsRequest) GetFileComponentIds() []int32 {
	if x != nil {
		return x.FileComponentIds
	}
	return nil
}

type CreateFileComponentVectorEmbeddingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileComponentVectorEmbeddingIds []int32 `protobuf:"varint,1,rep,packed,name=file_component_vector_embedding_ids,json=fileComponentVectorEmbeddingIds,proto3" json:"file_component_vector_embedding_ids,omitempty"`
}

func (x *CreateFileComponentVectorEmbeddingsResponse) Reset() {
	*x = CreateFileComponentVectorEmbeddingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vector_embedder_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileComponentVectorEmbeddingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileComponentVectorEmbeddingsResponse) ProtoMessage() {}

func (x *CreateFileComponentVectorEmbeddingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vector_embedder_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileComponentVectorEmbeddingsResponse.ProtoReflect.Descriptor instead.
func (*CreateFileComponentVectorEmbeddingsResponse) Descriptor() ([]byte, []int) {
	return file_vector_embedder_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFileComponentVectorEmbeddingsResponse) GetFileComponentVectorEmbeddingIds() []int32 {
	if x != nil {
		return x.FileComponentVectorEmbeddingIds
	}
	return nil
}

var File_vector_embedder_service_proto protoreflect.FileDescriptor

var file_vector_embedder_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5a, 0x0a, 0x2a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6d, 0x62, 0x65,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x12, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x22, 0x7b, 0x0a, 0x2b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x23, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x5f, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x1f, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6d, 0x62, 0x65,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x73, 0x32, 0x9a, 0x01, 0x0a, 0x15, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x80, 0x01, 0x0a, 0x23, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2b, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x2f, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vector_embedder_service_proto_rawDescOnce sync.Once
	file_vector_embedder_service_proto_rawDescData = file_vector_embedder_service_proto_rawDesc
)

func file_vector_embedder_service_proto_rawDescGZIP() []byte {
	file_vector_embedder_service_proto_rawDescOnce.Do(func() {
		file_vector_embedder_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_vector_embedder_service_proto_rawDescData)
	})
	return file_vector_embedder_service_proto_rawDescData
}

var file_vector_embedder_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vector_embedder_service_proto_goTypes = []interface{}{
	(*CreateFileComponentVectorEmbeddingsRequest)(nil),  // 0: CreateFileComponentVectorEmbeddingsRequest
	(*CreateFileComponentVectorEmbeddingsResponse)(nil), // 1: CreateFileComponentVectorEmbeddingsResponse
}
var file_vector_embedder_service_proto_depIdxs = []int32{
	0, // 0: VectorEmbedderService.CreateFileComponentVectorEmbeddings:input_type -> CreateFileComponentVectorEmbeddingsRequest
	1, // 1: VectorEmbedderService.CreateFileComponentVectorEmbeddings:output_type -> CreateFileComponentVectorEmbeddingsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vector_embedder_service_proto_init() }
func file_vector_embedder_service_proto_init() {
	if File_vector_embedder_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vector_embedder_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileComponentVectorEmbeddingsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vector_embedder_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileComponentVectorEmbeddingsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vector_embedder_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vector_embedder_service_proto_goTypes,
		DependencyIndexes: file_vector_embedder_service_proto_depIdxs,
		MessageInfos:      file_vector_embedder_service_proto_msgTypes,
	}.Build()
	File_vector_embedder_service_proto = out.File
	file_vector_embedder_service_proto_rawDesc = nil
	file_vector_embedder_service_proto_goTypes = nil
	file_vector_embedder_service_proto_depIdxs = nil
}
