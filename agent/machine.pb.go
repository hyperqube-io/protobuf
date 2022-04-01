// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: machine.proto

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

type Machine_PowerState int32

const (
	Machine_UNKNOWN     Machine_PowerState = 0
	Machine_POWERED_ON  Machine_PowerState = 1
	Machine_POWERED_OFF Machine_PowerState = 2
	Machine_SUSPENDED   Machine_PowerState = 3
)

// Enum value maps for Machine_PowerState.
var (
	Machine_PowerState_name = map[int32]string{
		0: "UNKNOWN",
		1: "POWERED_ON",
		2: "POWERED_OFF",
		3: "SUSPENDED",
	}
	Machine_PowerState_value = map[string]int32{
		"UNKNOWN":     0,
		"POWERED_ON":  1,
		"POWERED_OFF": 2,
		"SUSPENDED":   3,
	}
)

func (x Machine_PowerState) Enum() *Machine_PowerState {
	p := new(Machine_PowerState)
	*p = x
	return p
}

func (x Machine_PowerState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Machine_PowerState) Descriptor() protoreflect.EnumDescriptor {
	return file_machine_proto_enumTypes[0].Descriptor()
}

func (Machine_PowerState) Type() protoreflect.EnumType {
	return &file_machine_proto_enumTypes[0]
}

func (x Machine_PowerState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Machine_PowerState.Descriptor instead.
func (Machine_PowerState) EnumDescriptor() ([]byte, []int) {
	return file_machine_proto_rawDescGZIP(), []int{1, 0}
}

type MachineDisk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MountTarget string `protobuf:"bytes,1,opt,name=mountTarget,proto3" json:"mountTarget,omitempty"`
	SourceFile  string `protobuf:"bytes,2,opt,name=sourceFile,proto3" json:"sourceFile,omitempty"`
	Size        int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *MachineDisk) Reset() {
	*x = MachineDisk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineDisk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineDisk) ProtoMessage() {}

func (x *MachineDisk) ProtoReflect() protoreflect.Message {
	mi := &file_machine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineDisk.ProtoReflect.Descriptor instead.
func (*MachineDisk) Descriptor() ([]byte, []int) {
	return file_machine_proto_rawDescGZIP(), []int{0}
}

func (x *MachineDisk) GetMountTarget() string {
	if x != nil {
		return x.MountTarget
	}
	return ""
}

func (x *MachineDisk) GetSourceFile() string {
	if x != nil {
		return x.SourceFile
	}
	return ""
}

func (x *MachineDisk) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type Machine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PowerState Machine_PowerState `protobuf:"varint,3,opt,name=powerState,proto3,enum=Machine_PowerState" json:"powerState,omitempty"`
	Disks      []*MachineDisk     `protobuf:"bytes,4,rep,name=disks,proto3" json:"disks,omitempty"`
	Networks   []*Network         `protobuf:"bytes,5,rep,name=networks,proto3" json:"networks,omitempty"`
	Xml        string             `protobuf:"bytes,6,opt,name=xml,proto3" json:"xml,omitempty"`
}

func (x *Machine) Reset() {
	*x = Machine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Machine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Machine) ProtoMessage() {}

func (x *Machine) ProtoReflect() protoreflect.Message {
	mi := &file_machine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Machine.ProtoReflect.Descriptor instead.
func (*Machine) Descriptor() ([]byte, []int) {
	return file_machine_proto_rawDescGZIP(), []int{1}
}

func (x *Machine) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Machine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Machine) GetPowerState() Machine_PowerState {
	if x != nil {
		return x.PowerState
	}
	return Machine_UNKNOWN
}

func (x *Machine) GetDisks() []*MachineDisk {
	if x != nil {
		return x.Disks
	}
	return nil
}

func (x *Machine) GetNetworks() []*Network {
	if x != nil {
		return x.Networks
	}
	return nil
}

func (x *Machine) GetXml() string {
	if x != nil {
		return x.Xml
	}
	return ""
}

type GetMachineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMachineRequest) Reset() {
	*x = GetMachineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMachineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMachineRequest) ProtoMessage() {}

func (x *GetMachineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_machine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMachineRequest.ProtoReflect.Descriptor instead.
func (*GetMachineRequest) Descriptor() ([]byte, []int) {
	return file_machine_proto_rawDescGZIP(), []int{2}
}

func (x *GetMachineRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetVncPortReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port uint32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *GetVncPortReply) Reset() {
	*x = GetVncPortReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVncPortReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVncPortReply) ProtoMessage() {}

func (x *GetVncPortReply) ProtoReflect() protoreflect.Message {
	mi := &file_machine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVncPortReply.ProtoReflect.Descriptor instead.
func (*GetVncPortReply) Descriptor() ([]byte, []int) {
	return file_machine_proto_rawDescGZIP(), []int{3}
}

func (x *GetVncPortReply) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type GetMachinesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids    []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	Fields []string `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *GetMachinesRequest) Reset() {
	*x = GetMachinesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMachinesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMachinesRequest) ProtoMessage() {}

func (x *GetMachinesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_machine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMachinesRequest.ProtoReflect.Descriptor instead.
func (*GetMachinesRequest) Descriptor() ([]byte, []int) {
	return file_machine_proto_rawDescGZIP(), []int{4}
}

func (x *GetMachinesRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GetMachinesRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

type GetMachinesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Machines []*Machine `protobuf:"bytes,1,rep,name=machines,proto3" json:"machines,omitempty"`
}

func (x *GetMachinesReply) Reset() {
	*x = GetMachinesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMachinesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMachinesReply) ProtoMessage() {}

func (x *GetMachinesReply) ProtoReflect() protoreflect.Message {
	mi := &file_machine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMachinesReply.ProtoReflect.Descriptor instead.
func (*GetMachinesReply) Descriptor() ([]byte, []int) {
	return file_machine_proto_rawDescGZIP(), []int{5}
}

func (x *GetMachinesReply) GetMachines() []*Machine {
	if x != nil {
		return x.Machines
	}
	return nil
}

type CreateMachineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CpuCount uint32 `protobuf:"varint,2,opt,name=cpuCount,proto3" json:"cpuCount,omitempty"`
	Memory   uint32 `protobuf:"varint,3,opt,name=memory,proto3" json:"memory,omitempty"`
	DiskPath string `protobuf:"bytes,4,opt,name=diskPath,proto3" json:"diskPath,omitempty"`
	Xml      string `protobuf:"bytes,5,opt,name=xml,proto3" json:"xml,omitempty"`
}

func (x *CreateMachineRequest) Reset() {
	*x = CreateMachineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMachineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMachineRequest) ProtoMessage() {}

func (x *CreateMachineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_machine_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMachineRequest.ProtoReflect.Descriptor instead.
func (*CreateMachineRequest) Descriptor() ([]byte, []int) {
	return file_machine_proto_rawDescGZIP(), []int{6}
}

func (x *CreateMachineRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateMachineRequest) GetCpuCount() uint32 {
	if x != nil {
		return x.CpuCount
	}
	return 0
}

func (x *CreateMachineRequest) GetMemory() uint32 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *CreateMachineRequest) GetDiskPath() string {
	if x != nil {
		return x.DiskPath
	}
	return ""
}

func (x *CreateMachineRequest) GetXml() string {
	if x != nil {
		return x.Xml
	}
	return ""
}

type CreateMachineReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateMachineReply) Reset() {
	*x = CreateMachineReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMachineReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMachineReply) ProtoMessage() {}

func (x *CreateMachineReply) ProtoReflect() protoreflect.Message {
	mi := &file_machine_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMachineReply.ProtoReflect.Descriptor instead.
func (*CreateMachineReply) Descriptor() ([]byte, []int) {
	return file_machine_proto_rawDescGZIP(), []int{7}
}

func (x *CreateMachineReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MachineNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineId string `protobuf:"bytes,1,opt,name=machineId,proto3" json:"machineId,omitempty"`
	NetworkId string `protobuf:"bytes,2,opt,name=networkId,proto3" json:"networkId,omitempty"`
}

func (x *MachineNetworkRequest) Reset() {
	*x = MachineNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineNetworkRequest) ProtoMessage() {}

func (x *MachineNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_machine_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineNetworkRequest.ProtoReflect.Descriptor instead.
func (*MachineNetworkRequest) Descriptor() ([]byte, []int) {
	return file_machine_proto_rawDescGZIP(), []int{8}
}

func (x *MachineNetworkRequest) GetMachineId() string {
	if x != nil {
		return x.MachineId
	}
	return ""
}

func (x *MachineNetworkRequest) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

type StopMachinesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids   []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	Force bool     `protobuf:"varint,2,opt,name=force,proto3" json:"force,omitempty"`
}

func (x *StopMachinesRequest) Reset() {
	*x = StopMachinesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopMachinesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopMachinesRequest) ProtoMessage() {}

func (x *StopMachinesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_machine_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopMachinesRequest.ProtoReflect.Descriptor instead.
func (*StopMachinesRequest) Descriptor() ([]byte, []int) {
	return file_machine_proto_rawDescGZIP(), []int{9}
}

func (x *StopMachinesRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *StopMachinesRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

var File_machine_proto protoreflect.FileDescriptor

var file_machine_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a,
	0x75, 0x74, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x0b, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22,
	0x89, 0x02, 0x0a, 0x07, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x33, 0x0a, 0x0a, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x6f,
	0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x44, 0x69, 0x73,
	0x6b, 0x52, 0x05, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x12, 0x24, 0x0a, 0x08, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x52, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x78, 0x6d, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x78, 0x6d, 0x6c,
	0x22, 0x49, 0x0a, 0x0a, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50,
	0x4f, 0x57, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x50,
	0x4f, 0x57, 0x45, 0x52, 0x45, 0x44, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09,
	0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x03, 0x22, 0x23, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x56, 0x6e, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x3e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x24, 0x0a, 0x08, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x73, 0x22, 0x8c, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x63, 0x70, 0x75, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x50, 0x61, 0x74, 0x68, 0x12, 0x10,
	0x0a, 0x03, 0x78, 0x6d, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x78, 0x6d, 0x6c,
	0x22, 0x24, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x53, 0x0a, 0x15, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x13, 0x53,
	0x74, 0x6f, 0x70, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x32, 0xae, 0x03, 0x0a, 0x0e, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x56, 0x6e, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6e, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x13,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x2d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x2c, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2c, 0x0a,
	0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0b, 0x4c,
	0x69, 0x6e, 0x6b, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x16, 0x2e, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x0d, 0x55, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x12, 0x16, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x68,
	0x79, 0x70, 0x65, 0x72, 0x71, 0x75, 0x62, 0x65, 0x2d, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_machine_proto_rawDescOnce sync.Once
	file_machine_proto_rawDescData = file_machine_proto_rawDesc
)

func file_machine_proto_rawDescGZIP() []byte {
	file_machine_proto_rawDescOnce.Do(func() {
		file_machine_proto_rawDescData = protoimpl.X.CompressGZIP(file_machine_proto_rawDescData)
	})
	return file_machine_proto_rawDescData
}

var file_machine_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_machine_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_machine_proto_goTypes = []interface{}{
	(Machine_PowerState)(0),       // 0: Machine.PowerState
	(*MachineDisk)(nil),           // 1: MachineDisk
	(*Machine)(nil),               // 2: Machine
	(*GetMachineRequest)(nil),     // 3: GetMachineRequest
	(*GetVncPortReply)(nil),       // 4: GetVncPortReply
	(*GetMachinesRequest)(nil),    // 5: GetMachinesRequest
	(*GetMachinesReply)(nil),      // 6: GetMachinesReply
	(*CreateMachineRequest)(nil),  // 7: CreateMachineRequest
	(*CreateMachineReply)(nil),    // 8: CreateMachineReply
	(*MachineNetworkRequest)(nil), // 9: MachineNetworkRequest
	(*StopMachinesRequest)(nil),   // 10: StopMachinesRequest
	(*Network)(nil),               // 11: Network
	(*ActionReply)(nil),           // 12: ActionReply
}
var file_machine_proto_depIdxs = []int32{
	0,  // 0: Machine.powerState:type_name -> Machine.PowerState
	1,  // 1: Machine.disks:type_name -> MachineDisk
	11, // 2: Machine.networks:type_name -> Network
	2,  // 3: GetMachinesReply.machines:type_name -> Machine
	3,  // 4: MachineService.GetVncPort:input_type -> GetMachineRequest
	5,  // 5: MachineService.GetMany:input_type -> GetMachinesRequest
	7,  // 6: MachineService.Create:input_type -> CreateMachineRequest
	5,  // 7: MachineService.Delete:input_type -> GetMachinesRequest
	5,  // 8: MachineService.Start:input_type -> GetMachinesRequest
	10, // 9: MachineService.Stop:input_type -> StopMachinesRequest
	9,  // 10: MachineService.LinkNetwork:input_type -> MachineNetworkRequest
	9,  // 11: MachineService.UnlinkNetwork:input_type -> MachineNetworkRequest
	4,  // 12: MachineService.GetVncPort:output_type -> GetVncPortReply
	6,  // 13: MachineService.GetMany:output_type -> GetMachinesReply
	8,  // 14: MachineService.Create:output_type -> CreateMachineReply
	12, // 15: MachineService.Delete:output_type -> ActionReply
	12, // 16: MachineService.Start:output_type -> ActionReply
	12, // 17: MachineService.Stop:output_type -> ActionReply
	12, // 18: MachineService.LinkNetwork:output_type -> ActionReply
	12, // 19: MachineService.UnlinkNetwork:output_type -> ActionReply
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_machine_proto_init() }
func file_machine_proto_init() {
	if File_machine_proto != nil {
		return
	}
	file_network_proto_init()
	file_util_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_machine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineDisk); i {
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
		file_machine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Machine); i {
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
		file_machine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMachineRequest); i {
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
		file_machine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVncPortReply); i {
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
		file_machine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMachinesRequest); i {
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
		file_machine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMachinesReply); i {
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
		file_machine_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMachineRequest); i {
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
		file_machine_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMachineReply); i {
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
		file_machine_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineNetworkRequest); i {
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
		file_machine_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopMachinesRequest); i {
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
			RawDescriptor: file_machine_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_machine_proto_goTypes,
		DependencyIndexes: file_machine_proto_depIdxs,
		EnumInfos:         file_machine_proto_enumTypes,
		MessageInfos:      file_machine_proto_msgTypes,
	}.Build()
	File_machine_proto = out.File
	file_machine_proto_rawDesc = nil
	file_machine_proto_goTypes = nil
	file_machine_proto_depIdxs = nil
}
