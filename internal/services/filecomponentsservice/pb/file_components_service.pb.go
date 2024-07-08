// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.23.3
// source: file_components_service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileComponent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RepositoryId string `protobuf:"bytes,2,opt,name=repository_id,json=repositoryId,proto3" json:"repository_id,omitempty"`
	FilePath     string `protobuf:"bytes,3,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	StartLine    int32  `protobuf:"varint,4,opt,name=start_line,json=startLine,proto3" json:"start_line,omitempty"`
	EndLine      int32  `protobuf:"varint,5,opt,name=end_line,json=endLine,proto3" json:"end_line,omitempty"`
	Content      string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *FileComponent) Reset() {
	*x = FileComponent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_components_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileComponent) ProtoMessage() {}

func (x *FileComponent) ProtoReflect() protoreflect.Message {
	mi := &file_file_components_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileComponent.ProtoReflect.Descriptor instead.
func (*FileComponent) Descriptor() ([]byte, []int) {
	return file_file_components_service_proto_rawDescGZIP(), []int{0}
}

func (x *FileComponent) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FileComponent) GetRepositoryId() string {
	if x != nil {
		return x.RepositoryId
	}
	return ""
}

func (x *FileComponent) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *FileComponent) GetStartLine() int32 {
	if x != nil {
		return x.StartLine
	}
	return 0
}

func (x *FileComponent) GetEndLine() int32 {
	if x != nil {
		return x.EndLine
	}
	return 0
}

func (x *FileComponent) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type FileComponents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileComponents []*FileComponent `protobuf:"bytes,1,rep,name=file_components,json=fileComponents,proto3" json:"file_components,omitempty"`
}

func (x *FileComponents) Reset() {
	*x = FileComponents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_components_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileComponents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileComponents) ProtoMessage() {}

func (x *FileComponents) ProtoReflect() protoreflect.Message {
	mi := &file_file_components_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileComponents.ProtoReflect.Descriptor instead.
func (*FileComponents) Descriptor() ([]byte, []int) {
	return file_file_components_service_proto_rawDescGZIP(), []int{1}
}

func (x *FileComponents) GetFileComponents() []*FileComponent {
	if x != nil {
		return x.FileComponents
	}
	return nil
}

type FileComponentIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileComponentIds []int32 `protobuf:"varint,1,rep,packed,name=file_component_ids,json=fileComponentIds,proto3" json:"file_component_ids,omitempty"`
}

func (x *FileComponentIds) Reset() {
	*x = FileComponentIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_components_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileComponentIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileComponentIds) ProtoMessage() {}

func (x *FileComponentIds) ProtoReflect() protoreflect.Message {
	mi := &file_file_components_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileComponentIds.ProtoReflect.Descriptor instead.
func (*FileComponentIds) Descriptor() ([]byte, []int) {
	return file_file_components_service_proto_rawDescGZIP(), []int{2}
}

func (x *FileComponentIds) GetFileComponentIds() []int32 {
	if x != nil {
		return x.FileComponentIds
	}
	return nil
}

type RepositoryFilePaths struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepositoryId string   `protobuf:"bytes,1,opt,name=repository_id,json=repositoryId,proto3" json:"repository_id,omitempty"`
	FilePaths    []string `protobuf:"bytes,2,rep,name=file_paths,json=filePaths,proto3" json:"file_paths,omitempty"`
}

func (x *RepositoryFilePaths) Reset() {
	*x = RepositoryFilePaths{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_components_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepositoryFilePaths) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepositoryFilePaths) ProtoMessage() {}

func (x *RepositoryFilePaths) ProtoReflect() protoreflect.Message {
	mi := &file_file_components_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepositoryFilePaths.ProtoReflect.Descriptor instead.
func (*RepositoryFilePaths) Descriptor() ([]byte, []int) {
	return file_file_components_service_proto_rawDescGZIP(), []int{3}
}

func (x *RepositoryFilePaths) GetRepositoryId() string {
	if x != nil {
		return x.RepositoryId
	}
	return ""
}

func (x *RepositoryFilePaths) GetFilePaths() []string {
	if x != nil {
		return x.FilePaths
	}
	return nil
}

type DeleteFileComponentsByRepositoryIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepositoryId string `protobuf:"bytes,1,opt,name=repository_id,json=repositoryId,proto3" json:"repository_id,omitempty"`
}

func (x *DeleteFileComponentsByRepositoryIdRequest) Reset() {
	*x = DeleteFileComponentsByRepositoryIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_components_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileComponentsByRepositoryIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileComponentsByRepositoryIdRequest) ProtoMessage() {}

func (x *DeleteFileComponentsByRepositoryIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_components_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileComponentsByRepositoryIdRequest.ProtoReflect.Descriptor instead.
func (*DeleteFileComponentsByRepositoryIdRequest) Descriptor() ([]byte, []int) {
	return file_file_components_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteFileComponentsByRepositoryIdRequest) GetRepositoryId() string {
	if x != nil {
		return x.RepositoryId
	}
	return ""
}

var File_file_components_service_proto protoreflect.FileDescriptor

var file_file_components_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a,
	0x0d, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0x49, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52,
	0x0e, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x40, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x10, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x73, 0x22, 0x59, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x46,
	0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x73, 0x22, 0x50, 0x0a, 0x29,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x32, 0xf9,
	0x01, 0x0a, 0x15, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x14, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x46, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x73, 0x1a, 0x0f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x1a,
	0x0f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x68, 0x0a, 0x22, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x2a, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x1b, 0x5a, 0x19, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_components_service_proto_rawDescOnce sync.Once
	file_file_components_service_proto_rawDescData = file_file_components_service_proto_rawDesc
)

func file_file_components_service_proto_rawDescGZIP() []byte {
	file_file_components_service_proto_rawDescOnce.Do(func() {
		file_file_components_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_components_service_proto_rawDescData)
	})
	return file_file_components_service_proto_rawDescData
}

var file_file_components_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_file_components_service_proto_goTypes = []interface{}{
	(*FileComponent)(nil),                             // 0: FileComponent
	(*FileComponents)(nil),                            // 1: FileComponents
	(*FileComponentIds)(nil),                          // 2: FileComponentIds
	(*RepositoryFilePaths)(nil),                       // 3: RepositoryFilePaths
	(*DeleteFileComponentsByRepositoryIdRequest)(nil), // 4: DeleteFileComponentsByRepositoryIdRequest
	(*emptypb.Empty)(nil),                             // 5: google.protobuf.Empty
}
var file_file_components_service_proto_depIdxs = []int32{
	0, // 0: FileComponents.file_components:type_name -> FileComponent
	3, // 1: FileComponentsService.CreateFileComponents:input_type -> RepositoryFilePaths
	2, // 2: FileComponentsService.GetFileComponents:input_type -> FileComponentIds
	4, // 3: FileComponentsService.DeleteFileComponentsByRepositoryId:input_type -> DeleteFileComponentsByRepositoryIdRequest
	1, // 4: FileComponentsService.CreateFileComponents:output_type -> FileComponents
	1, // 5: FileComponentsService.GetFileComponents:output_type -> FileComponents
	5, // 6: FileComponentsService.DeleteFileComponentsByRepositoryId:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_file_components_service_proto_init() }
func file_file_components_service_proto_init() {
	if File_file_components_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_components_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileComponent); i {
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
		file_file_components_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileComponents); i {
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
		file_file_components_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileComponentIds); i {
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
		file_file_components_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepositoryFilePaths); i {
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
		file_file_components_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileComponentsByRepositoryIdRequest); i {
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
			RawDescriptor: file_file_components_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_components_service_proto_goTypes,
		DependencyIndexes: file_file_components_service_proto_depIdxs,
		MessageInfos:      file_file_components_service_proto_msgTypes,
	}.Build()
	File_file_components_service_proto = out.File
	file_file_components_service_proto_rawDesc = nil
	file_file_components_service_proto_goTypes = nil
	file_file_components_service_proto_depIdxs = nil
}
