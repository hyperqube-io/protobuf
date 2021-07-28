// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.15.7
// source: util.proto

package agent

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

type ActionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ActionReply) Reset() {
	*x = ActionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_util_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionReply) ProtoMessage() {}

func (x *ActionReply) ProtoReflect() protoreflect.Message {
	mi := &file_util_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionReply.ProtoReflect.Descriptor instead.
func (*ActionReply) Descriptor() ([]byte, []int) {
	return file_util_proto_rawDescGZIP(), []int{0}
}

func (x *ActionReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_util_proto protoreflect.FileDescriptor

var file_util_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x0b,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x42, 0x1d, 0x5a, 0x1b, 0x68,
	0x79, 0x70, 0x65, 0x72, 0x71, 0x75, 0x62, 0x65, 0x2d, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_util_proto_rawDescOnce sync.Once
	file_util_proto_rawDescData = file_util_proto_rawDesc
)

func file_util_proto_rawDescGZIP() []byte {
	file_util_proto_rawDescOnce.Do(func() {
		file_util_proto_rawDescData = protoimpl.X.CompressGZIP(file_util_proto_rawDescData)
	})
	return file_util_proto_rawDescData
}

var file_util_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_util_proto_goTypes = []interface{}{
	(*ActionReply)(nil), // 0: ActionReply
}
var file_util_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_util_proto_init() }
func file_util_proto_init() {
	if File_util_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_util_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionReply); i {
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
			RawDescriptor: file_util_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_util_proto_goTypes,
		DependencyIndexes: file_util_proto_depIdxs,
		MessageInfos:      file_util_proto_msgTypes,
	}.Build()
	File_util_proto = out.File
	file_util_proto_rawDesc = nil
	file_util_proto_goTypes = nil
	file_util_proto_depIdxs = nil
}
