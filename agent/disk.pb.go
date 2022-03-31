// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: disk.proto

package agent

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CopyFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *CopyFileRequest) Reset() {
	*x = CopyFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyFileRequest) ProtoMessage() {}

func (x *CopyFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyFileRequest.ProtoReflect.Descriptor instead.
func (*CopyFileRequest) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{0}
}

func (x *CopyFileRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *CopyFileRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type MoveFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *MoveFileRequest) Reset() {
	*x = MoveFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MoveFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MoveFileRequest) ProtoMessage() {}

func (x *MoveFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MoveFileRequest.ProtoReflect.Descriptor instead.
func (*MoveFileRequest) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{1}
}

func (x *MoveFileRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *MoveFileRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type UpdateDiskBackingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiskPath    string `protobuf:"bytes,1,opt,name=diskPath,proto3" json:"diskPath,omitempty"`
	BackingPath string `protobuf:"bytes,2,opt,name=backingPath,proto3" json:"backingPath,omitempty"`
}

func (x *UpdateDiskBackingRequest) Reset() {
	*x = UpdateDiskBackingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDiskBackingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDiskBackingRequest) ProtoMessage() {}

func (x *UpdateDiskBackingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDiskBackingRequest.ProtoReflect.Descriptor instead.
func (*UpdateDiskBackingRequest) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateDiskBackingRequest) GetDiskPath() string {
	if x != nil {
		return x.DiskPath
	}
	return ""
}

func (x *UpdateDiskBackingRequest) GetBackingPath() string {
	if x != nil {
		return x.BackingPath
	}
	return ""
}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModifiedAt  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=modifiedAt,proto3" json:"modifiedAt,omitempty"`
	IsDirectory bool                   `protobuf:"varint,2,opt,name=isDirectory,proto3" json:"isDirectory,omitempty"`
	Size        int64                  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Flags       uint32                 `protobuf:"varint,4,opt,name=flags,proto3" json:"flags,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{3}
}

func (x *FileInfo) GetModifiedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ModifiedAt
	}
	return nil
}

func (x *FileInfo) GetIsDirectory() bool {
	if x != nil {
		return x.IsDirectory
	}
	return false
}

func (x *FileInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileInfo) GetFlags() uint32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

type FileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *FileRequest) Reset() {
	*x = FileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRequest) ProtoMessage() {}

func (x *FileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRequest.ProtoReflect.Descriptor instead.
func (*FileRequest) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{4}
}

func (x *FileRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type GetFileContentsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetFileContentsReply) Reset() {
	*x = GetFileContentsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileContentsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileContentsReply) ProtoMessage() {}

func (x *GetFileContentsReply) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileContentsReply.ProtoReflect.Descriptor instead.
func (*GetFileContentsReply) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{5}
}

func (x *GetFileContentsReply) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type GetFileExistsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *GetFileExistsReply) Reset() {
	*x = GetFileExistsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileExistsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileExistsReply) ProtoMessage() {}

func (x *GetFileExistsReply) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileExistsReply.ProtoReflect.Descriptor instead.
func (*GetFileExistsReply) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{6}
}

func (x *GetFileExistsReply) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type GetFileInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *FileInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *GetFileInfoReply) Reset() {
	*x = GetFileInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileInfoReply) ProtoMessage() {}

func (x *GetFileInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileInfoReply.ProtoReflect.Descriptor instead.
func (*GetFileInfoReply) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{7}
}

func (x *GetFileInfoReply) GetInfo() *FileInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetFileSizeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *GetFileSizeReply) Reset() {
	*x = GetFileSizeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileSizeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileSizeReply) ProtoMessage() {}

func (x *GetFileSizeReply) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileSizeReply.ProtoReflect.Descriptor instead.
func (*GetFileSizeReply) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{8}
}

func (x *GetFileSizeReply) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ResizeDiskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Size int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ResizeDiskRequest) Reset() {
	*x = ResizeDiskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResizeDiskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResizeDiskRequest) ProtoMessage() {}

func (x *ResizeDiskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResizeDiskRequest.ProtoReflect.Descriptor instead.
func (*ResizeDiskRequest) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{9}
}

func (x *ResizeDiskRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ResizeDiskRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type WriteFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Url  string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *WriteFileRequest) Reset() {
	*x = WriteFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteFileRequest) ProtoMessage() {}

func (x *WriteFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteFileRequest.ProtoReflect.Descriptor instead.
func (*WriteFileRequest) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{10}
}

func (x *WriteFileRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *WriteFileRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *WriteFileRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_disk_proto protoreflect.FileDescriptor

var file_disk_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x69, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75,
	0x74, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x0f, 0x43, 0x6f, 0x70,
	0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f,
	0x22, 0x35, 0x0a, 0x0f, 0x4d, 0x6f, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x69, 0x73, 0x6b, 0x42, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x20, 0x0a, 0x0b, 0x62, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x74,
	0x68, 0x22, 0x92, 0x01, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a,
	0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x69, 0x73, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x22, 0x21, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x22, 0x31, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x3b,
	0x0a, 0x11, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x4c, 0x0a, 0x10, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0xed, 0x04, 0x0a, 0x0b, 0x44, 0x69,
	0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x43, 0x6f, 0x70,
	0x79, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x2e, 0x43, 0x6f, 0x70, 0x79, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x10, 0x2e, 0x43, 0x6f,
	0x70, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2a, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0c, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0c, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x12, 0x0c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x0c, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2a, 0x0a,
	0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x0c, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x08, 0x4d, 0x6f, 0x76,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x2e, 0x4d, 0x6f, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x69, 0x7a,
	0x65, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x12, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x44, 0x69,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x42, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x19,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x42, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x09, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x68, 0x79, 0x70,
	0x65, 0x72, 0x71, 0x75, 0x62, 0x65, 0x2d, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_disk_proto_rawDescOnce sync.Once
	file_disk_proto_rawDescData = file_disk_proto_rawDesc
)

func file_disk_proto_rawDescGZIP() []byte {
	file_disk_proto_rawDescOnce.Do(func() {
		file_disk_proto_rawDescData = protoimpl.X.CompressGZIP(file_disk_proto_rawDescData)
	})
	return file_disk_proto_rawDescData
}

var file_disk_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_disk_proto_goTypes = []interface{}{
	(*CopyFileRequest)(nil),          // 0: CopyFileRequest
	(*MoveFileRequest)(nil),          // 1: MoveFileRequest
	(*UpdateDiskBackingRequest)(nil), // 2: UpdateDiskBackingRequest
	(*FileInfo)(nil),                 // 3: FileInfo
	(*FileRequest)(nil),              // 4: FileRequest
	(*GetFileContentsReply)(nil),     // 5: GetFileContentsReply
	(*GetFileExistsReply)(nil),       // 6: GetFileExistsReply
	(*GetFileInfoReply)(nil),         // 7: GetFileInfoReply
	(*GetFileSizeReply)(nil),         // 8: GetFileSizeReply
	(*ResizeDiskRequest)(nil),        // 9: ResizeDiskRequest
	(*WriteFileRequest)(nil),         // 10: WriteFileRequest
	(*timestamppb.Timestamp)(nil),    // 11: google.protobuf.Timestamp
	(*ActionReply)(nil),              // 12: ActionReply
}
var file_disk_proto_depIdxs = []int32{
	11, // 0: FileInfo.modifiedAt:type_name -> google.protobuf.Timestamp
	3,  // 1: GetFileInfoReply.info:type_name -> FileInfo
	0,  // 2: DiskService.CopyFile:input_type -> CopyFileRequest
	0,  // 3: DiskService.CreateLinkedDisk:input_type -> CopyFileRequest
	4,  // 4: DiskService.DeleteFile:input_type -> FileRequest
	4,  // 5: DiskService.GetFileContents:input_type -> FileRequest
	4,  // 6: DiskService.GetFileExists:input_type -> FileRequest
	4,  // 7: DiskService.GetFileInfo:input_type -> FileRequest
	4,  // 8: DiskService.GetFileSize:input_type -> FileRequest
	4,  // 9: DiskService.CommitDisk:input_type -> FileRequest
	1,  // 10: DiskService.MoveFile:input_type -> MoveFileRequest
	9,  // 11: DiskService.ResizeDisk:input_type -> ResizeDiskRequest
	2,  // 12: DiskService.UpdateDiskBacking:input_type -> UpdateDiskBackingRequest
	10, // 13: DiskService.WriteFile:input_type -> WriteFileRequest
	12, // 14: DiskService.CopyFile:output_type -> ActionReply
	12, // 15: DiskService.CreateLinkedDisk:output_type -> ActionReply
	12, // 16: DiskService.DeleteFile:output_type -> ActionReply
	5,  // 17: DiskService.GetFileContents:output_type -> GetFileContentsReply
	6,  // 18: DiskService.GetFileExists:output_type -> GetFileExistsReply
	7,  // 19: DiskService.GetFileInfo:output_type -> GetFileInfoReply
	8,  // 20: DiskService.GetFileSize:output_type -> GetFileSizeReply
	12, // 21: DiskService.CommitDisk:output_type -> ActionReply
	12, // 22: DiskService.MoveFile:output_type -> ActionReply
	12, // 23: DiskService.ResizeDisk:output_type -> ActionReply
	12, // 24: DiskService.UpdateDiskBacking:output_type -> ActionReply
	12, // 25: DiskService.WriteFile:output_type -> ActionReply
	14, // [14:26] is the sub-list for method output_type
	2,  // [2:14] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_disk_proto_init() }
func file_disk_proto_init() {
	if File_disk_proto != nil {
		return
	}
	file_util_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_disk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyFileRequest); i {
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
		file_disk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MoveFileRequest); i {
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
		file_disk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDiskBackingRequest); i {
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
		file_disk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_disk_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRequest); i {
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
		file_disk_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileContentsReply); i {
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
		file_disk_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileExistsReply); i {
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
		file_disk_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileInfoReply); i {
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
		file_disk_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileSizeReply); i {
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
		file_disk_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResizeDiskRequest); i {
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
		file_disk_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteFileRequest); i {
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
			RawDescriptor: file_disk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_disk_proto_goTypes,
		DependencyIndexes: file_disk_proto_depIdxs,
		MessageInfos:      file_disk_proto_msgTypes,
	}.Build()
	File_disk_proto = out.File
	file_disk_proto_rawDesc = nil
	file_disk_proto_goTypes = nil
	file_disk_proto_depIdxs = nil
}
