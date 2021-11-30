// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: proto/ah.proto

package Proto

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

type Acknowledgement_Status int32

const (
	Acknowledgement_FAIL      Acknowledgement_Status = 0
	Acknowledgement_SUCCES    Acknowledgement_Status = 1
	Acknowledgement_EXCEPTION Acknowledgement_Status = 2
)

// Enum value maps for Acknowledgement_Status.
var (
	Acknowledgement_Status_name = map[int32]string{
		0: "FAIL",
		1: "SUCCES",
		2: "EXCEPTION",
	}
	Acknowledgement_Status_value = map[string]int32{
		"FAIL":      0,
		"SUCCES":    1,
		"EXCEPTION": 2,
	}
)

func (x Acknowledgement_Status) Enum() *Acknowledgement_Status {
	p := new(Acknowledgement_Status)
	*p = x
	return p
}

func (x Acknowledgement_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Acknowledgement_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_ah_proto_enumTypes[0].Descriptor()
}

func (Acknowledgement_Status) Type() protoreflect.EnumType {
	return &file_proto_ah_proto_enumTypes[0]
}

func (x Acknowledgement_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Acknowledgement_Status.Descriptor instead.
func (Acknowledgement_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_ah_proto_rawDescGZIP(), []int{4, 0}
}

type Offer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount int64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	User   string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Offer) Reset() {
	*x = Offer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ah_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Offer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Offer) ProtoMessage() {}

func (x *Offer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ah_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Offer.ProtoReflect.Descriptor instead.
func (*Offer) Descriptor() ([]byte, []int) {
	return file_proto_ah_proto_rawDescGZIP(), []int{0}
}

func (x *Offer) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Offer) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Offer) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details         *Details         `protobuf:"bytes,1,opt,name=details,proto3,oneof" json:"details,omitempty"`
	Acknowledgement *Acknowledgement `protobuf:"bytes,2,opt,name=acknowledgement,proto3,oneof" json:"acknowledgement,omitempty"`
	DetailList      *DetailsList     `protobuf:"bytes,3,opt,name=detailList,proto3,oneof" json:"detailList,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ah_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ah_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_ah_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetDetails() *Details {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Response) GetAcknowledgement() *Acknowledgement {
	if x != nil {
		return x.Acknowledgement
	}
	return nil
}

func (x *Response) GetDetailList() *DetailsList {
	if x != nil {
		return x.DetailList
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ah_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ah_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_ah_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ah_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ah_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_proto_ah_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Acknowledgement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Acknowledgement_Status `protobuf:"varint,1,opt,name=status,proto3,enum=Proto.Acknowledgement_Status" json:"status,omitempty"`
}

func (x *Acknowledgement) Reset() {
	*x = Acknowledgement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ah_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Acknowledgement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Acknowledgement) ProtoMessage() {}

func (x *Acknowledgement) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ah_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Acknowledgement.ProtoReflect.Descriptor instead.
func (*Acknowledgement) Descriptor() ([]byte, []int) {
	return file_proto_ah_proto_rawDescGZIP(), []int{4}
}

func (x *Acknowledgement) GetStatus() Acknowledgement_Status {
	if x != nil {
		return x.Status
	}
	return Acknowledgement_FAIL
}

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ah_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ah_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_proto_ah_proto_rawDescGZIP(), []int{5}
}

func (x *Info) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Info) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type Details struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Timeleft int64  `protobuf:"varint,3,opt,name=timeleft,proto3" json:"timeleft,omitempty"`
	Amount   int64  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	User     string `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Details) Reset() {
	*x = Details{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ah_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Details) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Details) ProtoMessage() {}

func (x *Details) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ah_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Details.ProtoReflect.Descriptor instead.
func (*Details) Descriptor() ([]byte, []int) {
	return file_proto_ah_proto_rawDescGZIP(), []int{6}
}

func (x *Details) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Details) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Details) GetTimeleft() int64 {
	if x != nil {
		return x.Timeleft
	}
	return 0
}

func (x *Details) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Details) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type DetailsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       []int64  `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
	Name     []string `protobuf:"bytes,2,rep,name=name,proto3" json:"name,omitempty"`
	Timeleft []int64  `protobuf:"varint,3,rep,packed,name=timeleft,proto3" json:"timeleft,omitempty"`
	Amount   []int64  `protobuf:"varint,4,rep,packed,name=amount,proto3" json:"amount,omitempty"`
	User     []string `protobuf:"bytes,5,rep,name=user,proto3" json:"user,omitempty"`
}

func (x *DetailsList) Reset() {
	*x = DetailsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ah_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailsList) ProtoMessage() {}

func (x *DetailsList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ah_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailsList.ProtoReflect.Descriptor instead.
func (*DetailsList) Descriptor() ([]byte, []int) {
	return file_proto_ah_proto_rawDescGZIP(), []int{7}
}

func (x *DetailsList) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *DetailsList) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *DetailsList) GetTimeleft() []int64 {
	if x != nil {
		return x.Timeleft
	}
	return nil
}

func (x *DetailsList) GetAmount() []int64 {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *DetailsList) GetUser() []string {
	if x != nil {
		return x.User
	}
	return nil
}

var File_proto_ah_proto protoreflect.FileDescriptor

var file_proto_ah_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x05, 0x4f, 0x66, 0x66, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xe8, 0x01, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x48, 0x00, 0x52, 0x07, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x0f, 0x61, 0x63, 0x6b, 0x6e,
	0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x01, 0x52, 0x0f, 0x61, 0x63, 0x6b,
	0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x37, 0x0a, 0x0a, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x02, 0x52, 0x0a, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x4c, 0x69, 0x73, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x21, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x77, 0x0a, 0x0f, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63,
	0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2d, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x45, 0x58, 0x43, 0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x22, 0x28, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x75, 0x0a, 0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x65, 0x66, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x65, 0x66, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x79, 0x0a, 0x0b,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x65, 0x66, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xc2, 0x01, 0x0a, 0x0c, 0x41, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12,
	0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x1a, 0x16, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x0b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0e, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x27, 0x0a, 0x04,
	0x56, 0x69, 0x65, 0x77, 0x12, 0x0b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x12, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x16, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ah_proto_rawDescOnce sync.Once
	file_proto_ah_proto_rawDescData = file_proto_ah_proto_rawDesc
)

func file_proto_ah_proto_rawDescGZIP() []byte {
	file_proto_ah_proto_rawDescOnce.Do(func() {
		file_proto_ah_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ah_proto_rawDescData)
	})
	return file_proto_ah_proto_rawDescData
}

var file_proto_ah_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_ah_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_ah_proto_goTypes = []interface{}{
	(Acknowledgement_Status)(0), // 0: Proto.Acknowledgement.Status
	(*Offer)(nil),               // 1: Proto.Offer
	(*Response)(nil),            // 2: Proto.Response
	(*User)(nil),                // 3: Proto.User
	(*RegisterRequest)(nil),     // 4: Proto.RegisterRequest
	(*Acknowledgement)(nil),     // 5: Proto.Acknowledgement
	(*Info)(nil),                // 6: Proto.Info
	(*Details)(nil),             // 7: Proto.Details
	(*DetailsList)(nil),         // 8: Proto.DetailsList
}
var file_proto_ah_proto_depIdxs = []int32{
	7, // 0: Proto.Response.details:type_name -> Proto.Details
	5, // 1: Proto.Response.acknowledgement:type_name -> Proto.Acknowledgement
	8, // 2: Proto.Response.detailList:type_name -> Proto.DetailsList
	0, // 3: Proto.Acknowledgement.status:type_name -> Proto.Acknowledgement.Status
	1, // 4: Proto.AuctionHouse.Bid:input_type -> Proto.Offer
	6, // 5: Proto.AuctionHouse.Result:input_type -> Proto.Info
	3, // 6: Proto.AuctionHouse.View:input_type -> Proto.User
	4, // 7: Proto.AuctionHouse.Register:input_type -> Proto.RegisterRequest
	5, // 8: Proto.AuctionHouse.Bid:output_type -> Proto.Acknowledgement
	7, // 9: Proto.AuctionHouse.Result:output_type -> Proto.Details
	8, // 10: Proto.AuctionHouse.View:output_type -> Proto.DetailsList
	2, // 11: Proto.AuctionHouse.Register:output_type -> Proto.Response
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_ah_proto_init() }
func file_proto_ah_proto_init() {
	if File_proto_ah_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ah_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Offer); i {
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
		file_proto_ah_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_ah_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_ah_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_proto_ah_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Acknowledgement); i {
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
		file_proto_ah_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
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
		file_proto_ah_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Details); i {
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
		file_proto_ah_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailsList); i {
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
	file_proto_ah_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ah_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ah_proto_goTypes,
		DependencyIndexes: file_proto_ah_proto_depIdxs,
		EnumInfos:         file_proto_ah_proto_enumTypes,
		MessageInfos:      file_proto_ah_proto_msgTypes,
	}.Build()
	File_proto_ah_proto = out.File
	file_proto_ah_proto_rawDesc = nil
	file_proto_ah_proto_goTypes = nil
	file_proto_ah_proto_depIdxs = nil
}
