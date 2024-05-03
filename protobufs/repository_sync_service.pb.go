// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.23.3
// source: repository_sync_service.proto

package protobufs

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

type ReportSavedFileChunkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FilePath       string `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	ChunkIndex     int32  `protobuf:"varint,3,opt,name=chunk_index,json=chunkIndex,proto3" json:"chunk_index,omitempty"`
	NumTotalChunks int32  `protobuf:"varint,4,opt,name=num_total_chunks,json=numTotalChunks,proto3" json:"num_total_chunks,omitempty"`
}

func (x *ReportSavedFileChunkRequest) Reset() {
	*x = ReportSavedFileChunkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_sync_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportSavedFileChunkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportSavedFileChunkRequest) ProtoMessage() {}

func (x *ReportSavedFileChunkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_repository_sync_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportSavedFileChunkRequest.ProtoReflect.Descriptor instead.
func (*ReportSavedFileChunkRequest) Descriptor() ([]byte, []int) {
	return file_repository_sync_service_proto_rawDescGZIP(), []int{0}
}

func (x *ReportSavedFileChunkRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReportSavedFileChunkRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *ReportSavedFileChunkRequest) GetChunkIndex() int32 {
	if x != nil {
		return x.ChunkIndex
	}
	return 0
}

func (x *ReportSavedFileChunkRequest) GetNumTotalChunks() int32 {
	if x != nil {
		return x.NumTotalChunks
	}
	return 0
}

type ReportSavedFileChunkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsLastChunk bool `protobuf:"varint,1,opt,name=is_last_chunk,json=isLastChunk,proto3" json:"is_last_chunk,omitempty"`
}

func (x *ReportSavedFileChunkResponse) Reset() {
	*x = ReportSavedFileChunkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_sync_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportSavedFileChunkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportSavedFileChunkResponse) ProtoMessage() {}

func (x *ReportSavedFileChunkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_repository_sync_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportSavedFileChunkResponse.ProtoReflect.Descriptor instead.
func (*ReportSavedFileChunkResponse) Descriptor() ([]byte, []int) {
	return file_repository_sync_service_proto_rawDescGZIP(), []int{1}
}

func (x *ReportSavedFileChunkResponse) GetIsLastChunk() bool {
	if x != nil {
		return x.IsLastChunk
	}
	return false
}

var File_repository_sync_service_proto protoreflect.FileDescriptor

var file_repository_sync_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9e, 0x01, 0x0a, 0x1b, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73,
	0x22, 0x42, 0x0a, 0x1c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x32, 0x70, 0x0a, 0x15, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a,
	0x14, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x1c, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x61,
	0x76, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x61, 0x76, 0x65,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_repository_sync_service_proto_rawDescOnce sync.Once
	file_repository_sync_service_proto_rawDescData = file_repository_sync_service_proto_rawDesc
)

func file_repository_sync_service_proto_rawDescGZIP() []byte {
	file_repository_sync_service_proto_rawDescOnce.Do(func() {
		file_repository_sync_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_repository_sync_service_proto_rawDescData)
	})
	return file_repository_sync_service_proto_rawDescData
}

var file_repository_sync_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_repository_sync_service_proto_goTypes = []interface{}{
	(*ReportSavedFileChunkRequest)(nil),  // 0: ReportSavedFileChunkRequest
	(*ReportSavedFileChunkResponse)(nil), // 1: ReportSavedFileChunkResponse
}
var file_repository_sync_service_proto_depIdxs = []int32{
	0, // 0: RepositorySyncService.ReportSavedFileChunk:input_type -> ReportSavedFileChunkRequest
	1, // 1: RepositorySyncService.ReportSavedFileChunk:output_type -> ReportSavedFileChunkResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_repository_sync_service_proto_init() }
func file_repository_sync_service_proto_init() {
	if File_repository_sync_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_repository_sync_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportSavedFileChunkRequest); i {
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
		file_repository_sync_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportSavedFileChunkResponse); i {
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
			RawDescriptor: file_repository_sync_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_repository_sync_service_proto_goTypes,
		DependencyIndexes: file_repository_sync_service_proto_depIdxs,
		MessageInfos:      file_repository_sync_service_proto_msgTypes,
	}.Build()
	File_repository_sync_service_proto = out.File
	file_repository_sync_service_proto_rawDesc = nil
	file_repository_sync_service_proto_goTypes = nil
	file_repository_sync_service_proto_depIdxs = nil
}