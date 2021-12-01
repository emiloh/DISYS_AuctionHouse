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

type Ack_Response int32

const (
	Ack_FAIL      Ack_Response = 0
	Ack_SUCCES    Ack_Response = 1
	Ack_EXCEPTION Ack_Response = 2
)

// Enum value maps for Ack_Response.
var (
	Ack_Response_name = map[int32]string{
		0: "FAIL",
		1: "SUCCES",
		2: "EXCEPTION",
	}
	Ack_Response_value = map[string]int32{
		"FAIL":      0,
		"SUCCES":    1,
		"EXCEPTION": 2,
	}
)

func (x Ack_Response) Enum() *Ack_Response {
	p := new(Ack_Response)
	*p = x
	return p
}

func (x Ack_Response) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ack_Response) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_ah_proto_enumTypes[0].Descriptor()
}

func (Ack_Response) Type() protoreflect.EnumType {
	return &file_proto_ah_proto_enumTypes[0]
}

func (x Ack_Response) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ack_Response.Descriptor instead.
func (Ack_Response) EnumDescriptor() ([]byte, []int) {
	return file_proto_ah_proto_rawDescGZIP(), []int{1, 0}
}

type Offer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response Ack_Response `protobuf:"varint,1,opt,name=response,proto3,enum=Proto.Ack_Response" json:"response,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ah_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_proto_ah_proto_rawDescGZIP(), []int{1}
}

func (x *Ack) GetResponse() Ack_Response {
	if x != nil {
		return x.Response
	}
	return Ack_FAIL
}

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ah_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_proto_ah_proto_rawDescGZIP(), []int{2}
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
		mi := &file_proto_ah_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Details) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Details) ProtoMessage() {}

func (x *Details) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Details.ProtoReflect.Descriptor instead.
func (*Details) Descriptor() ([]byte, []int) {
	return file_proto_ah_proto_rawDescGZIP(), []int{3}
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

var File_proto_ah_proto protoreflect.FileDescriptor

var file_proto_ah_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x05, 0x4f, 0x66, 0x66, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x67, 0x0a, 0x03,
	0x41, 0x63, 0x6b, 0x12, 0x2f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63,
	0x6b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58, 0x43, 0x45, 0x50, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x02, 0x22, 0x18, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x75, 0x0a, 0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x65, 0x66, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0x56, 0x0a, 0x0c, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x0c, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x1a, 0x0a, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x6b, 0x12, 0x25, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x0b, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0e,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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
var file_proto_ah_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_ah_proto_goTypes = []interface{}{
	(Ack_Response)(0), // 0: Proto.Ack.Response
	(*Offer)(nil),     // 1: Proto.Offer
	(*Ack)(nil),       // 2: Proto.Ack
	(*Info)(nil),      // 3: Proto.Info
	(*Details)(nil),   // 4: Proto.Details
}
var file_proto_ah_proto_depIdxs = []int32{
	0, // 0: Proto.Ack.response:type_name -> Proto.Ack.Response
	1, // 1: Proto.AuctionHouse.Bid:input_type -> Proto.Offer
	3, // 2: Proto.AuctionHouse.Result:input_type -> Proto.Info
	2, // 3: Proto.AuctionHouse.Bid:output_type -> Proto.Ack
	4, // 4: Proto.AuctionHouse.Result:output_type -> Proto.Details
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
			switch v := v.(*Ack); i {
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
		file_proto_ah_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ah_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
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
