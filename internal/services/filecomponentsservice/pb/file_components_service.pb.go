// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.23.3
// source: file_components_service.proto

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

type FileComponent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FilePath  string `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	StartLine int32  `protobuf:"varint,3,opt,name=start_line,json=startLine,proto3" json:"start_line,omitempty"`
	EndLine   int32  `protobuf:"varint,4,opt,name=end_line,json=endLine,proto3" json:"end_line,omitempty"`
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

func (x *FileComponent) GetUserId() string {
	if x != nil {
		return x.UserId
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

type SavedFileComponent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FilePath  string `protobuf:"bytes,3,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	StartLine int32  `protobuf:"varint,4,opt,name=start_line,json=startLine,proto3" json:"start_line,omitempty"`
	EndLine   int32  `protobuf:"varint,5,opt,name=end_line,json=endLine,proto3" json:"end_line,omitempty"`
}

func (x *SavedFileComponent) Reset() {
	*x = SavedFileComponent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_components_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavedFileComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavedFileComponent) ProtoMessage() {}

func (x *SavedFileComponent) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SavedFileComponent.ProtoReflect.Descriptor instead.
func (*SavedFileComponent) Descriptor() ([]byte, []int) {
	return file_file_components_service_proto_rawDescGZIP(), []int{1}
}

func (x *SavedFileComponent) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SavedFileComponent) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SavedFileComponent) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *SavedFileComponent) GetStartLine() int32 {
	if x != nil {
		return x.StartLine
	}
	return 0
}

func (x *SavedFileComponent) GetEndLine() int32 {
	if x != nil {
		return x.EndLine
	}
	return 0
}

type SavedFileComponents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SavedFileComponents []*SavedFileComponent `protobuf:"bytes,1,rep,name=saved_file_components,json=savedFileComponents,proto3" json:"saved_file_components,omitempty"`
}

func (x *SavedFileComponents) Reset() {
	*x = SavedFileComponents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_components_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavedFileComponents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavedFileComponents) ProtoMessage() {}

func (x *SavedFileComponents) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SavedFileComponents.ProtoReflect.Descriptor instead.
func (*SavedFileComponents) Descriptor() ([]byte, []int) {
	return file_file_components_service_proto_rawDescGZIP(), []int{2}
}

func (x *SavedFileComponents) GetSavedFileComponents() []*SavedFileComponent {
	if x != nil {
		return x.SavedFileComponents
	}
	return nil
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
		mi := &file_file_components_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileComponents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileComponents) ProtoMessage() {}

func (x *FileComponents) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FileComponents.ProtoReflect.Descriptor instead.
func (*FileComponents) Descriptor() ([]byte, []int) {
	return file_file_components_service_proto_rawDescGZIP(), []int{3}
}

func (x *FileComponents) GetFileComponents() []*FileComponent {
	if x != nil {
		return x.FileComponents
	}
	return nil
}

type BatchExtractFileComponentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FilePaths []string `protobuf:"bytes,2,rep,name=file_paths,json=filePaths,proto3" json:"file_paths,omitempty"`
}

func (x *BatchExtractFileComponentsRequest) Reset() {
	*x = BatchExtractFileComponentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_components_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchExtractFileComponentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchExtractFileComponentsRequest) ProtoMessage() {}

func (x *BatchExtractFileComponentsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BatchExtractFileComponentsRequest.ProtoReflect.Descriptor instead.
func (*BatchExtractFileComponentsRequest) Descriptor() ([]byte, []int) {
	return file_file_components_service_proto_rawDescGZIP(), []int{4}
}

func (x *BatchExtractFileComponentsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BatchExtractFileComponentsRequest) GetFilePaths() []string {
	if x != nil {
		return x.FilePaths
	}
	return nil
}

type SavedFileComponentIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SavedFileComponentIds []int32 `protobuf:"varint,1,rep,packed,name=saved_file_component_ids,json=savedFileComponentIds,proto3" json:"saved_file_component_ids,omitempty"`
}

func (x *SavedFileComponentIds) Reset() {
	*x = SavedFileComponentIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_components_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavedFileComponentIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavedFileComponentIds) ProtoMessage() {}

func (x *SavedFileComponentIds) ProtoReflect() protoreflect.Message {
	mi := &file_file_components_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavedFileComponentIds.ProtoReflect.Descriptor instead.
func (*SavedFileComponentIds) Descriptor() ([]byte, []int) {
	return file_file_components_service_proto_rawDescGZIP(), []int{5}
}

func (x *SavedFileComponentIds) GetSavedFileComponentIds() []int32 {
	if x != nil {
		return x.SavedFileComponentIds
	}
	return nil
}

var File_file_components_service_proto protoreflect.FileDescriptor

var file_file_components_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7f, 0x0a, 0x0d, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6e,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65,
	0x22, 0x94, 0x01, 0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x22, 0x5e, 0x0a, 0x13, 0x53, 0x61, 0x76, 0x65, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x47,
	0x0a, 0x15, 0x73, 0x61, 0x76, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x53, 0x61, 0x76, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x52, 0x13, 0x73, 0x61, 0x76, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x49, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x0f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x52, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x5b, 0x0a, 0x21, 0x42, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x73, 0x22,
	0x50, 0x0a, 0x15, 0x53, 0x61, 0x76, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x12, 0x37, 0x0a, 0x18, 0x73, 0x61, 0x76, 0x65,
	0x64, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x15, 0x73, 0x61, 0x76, 0x65,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x73, 0x32, 0xf1, 0x01, 0x0a, 0x15, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x1a, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3d,
	0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x0f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x16, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x12, 0x46, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x1a,
	0x14, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x1b, 0x5a, 0x19, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_file_components_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_file_components_service_proto_goTypes = []interface{}{
	(*FileComponent)(nil),                     // 0: FileComponent
	(*SavedFileComponent)(nil),                // 1: SavedFileComponent
	(*SavedFileComponents)(nil),               // 2: SavedFileComponents
	(*FileComponents)(nil),                    // 3: FileComponents
	(*BatchExtractFileComponentsRequest)(nil), // 4: BatchExtractFileComponentsRequest
	(*SavedFileComponentIds)(nil),             // 5: SavedFileComponentIds
}
var file_file_components_service_proto_depIdxs = []int32{
	1, // 0: SavedFileComponents.saved_file_components:type_name -> SavedFileComponent
	0, // 1: FileComponents.file_components:type_name -> FileComponent
	4, // 2: FileComponentsService.BatchExtractFileComponents:input_type -> BatchExtractFileComponentsRequest
	3, // 3: FileComponentsService.SaveFileComponents:input_type -> FileComponents
	5, // 4: FileComponentsService.GetSavedFileComponents:input_type -> SavedFileComponentIds
	3, // 5: FileComponentsService.BatchExtractFileComponents:output_type -> FileComponents
	5, // 6: FileComponentsService.SaveFileComponents:output_type -> SavedFileComponentIds
	2, // 7: FileComponentsService.GetSavedFileComponents:output_type -> SavedFileComponents
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
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
			switch v := v.(*SavedFileComponent); i {
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
			switch v := v.(*SavedFileComponents); i {
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
		file_file_components_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchExtractFileComponentsRequest); i {
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
		file_file_components_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavedFileComponentIds); i {
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
			NumMessages:   6,
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
