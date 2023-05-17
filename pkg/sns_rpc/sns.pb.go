// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: pkg/sns_rpc/sns.proto

package sns_rpc

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

type SnsRawContentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []string `protobuf:"bytes,1,rep,name=content,proto3" json:"content,omitempty"`
}

func (x *SnsRawContentReq) Reset() {
	*x = SnsRawContentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sns_rpc_sns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnsRawContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnsRawContentReq) ProtoMessage() {}

func (x *SnsRawContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sns_rpc_sns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnsRawContentReq.ProtoReflect.Descriptor instead.
func (*SnsRawContentReq) Descriptor() ([]byte, []int) {
	return file_pkg_sns_rpc_sns_proto_rawDescGZIP(), []int{0}
}

func (x *SnsRawContentReq) GetContent() []string {
	if x != nil {
		return x.Content
	}
	return nil
}

type Cut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Word  string  `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
	Count int32   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Freq  float32 `protobuf:"fixed32,3,opt,name=freq,proto3" json:"freq,omitempty"`
	Poly  float32 `protobuf:"fixed32,4,opt,name=poly,proto3" json:"poly,omitempty"`
	Flex  float32 `protobuf:"fixed32,5,opt,name=flex,proto3" json:"flex,omitempty"`
}

func (x *Cut) Reset() {
	*x = Cut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sns_rpc_sns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cut) ProtoMessage() {}

func (x *Cut) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sns_rpc_sns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cut.ProtoReflect.Descriptor instead.
func (*Cut) Descriptor() ([]byte, []int) {
	return file_pkg_sns_rpc_sns_proto_rawDescGZIP(), []int{1}
}

func (x *Cut) GetWord() string {
	if x != nil {
		return x.Word
	}
	return ""
}

func (x *Cut) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Cut) GetFreq() float32 {
	if x != nil {
		return x.Freq
	}
	return 0
}

func (x *Cut) GetPoly() float32 {
	if x != nil {
		return x.Poly
	}
	return 0
}

func (x *Cut) GetFlex() float32 {
	if x != nil {
		return x.Flex
	}
	return 0
}

type SnsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cuts []*Cut `protobuf:"bytes,1,rep,name=cuts,proto3" json:"cuts,omitempty"`
}

func (x *SnsReply) Reset() {
	*x = SnsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sns_rpc_sns_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnsReply) ProtoMessage() {}

func (x *SnsReply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sns_rpc_sns_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnsReply.ProtoReflect.Descriptor instead.
func (*SnsReply) Descriptor() ([]byte, []int) {
	return file_pkg_sns_rpc_sns_proto_rawDescGZIP(), []int{2}
}

func (x *SnsReply) GetCuts() []*Cut {
	if x != nil {
		return x.Cuts
	}
	return nil
}

var File_pkg_sns_rpc_sns_proto protoreflect.FileDescriptor

var file_pkg_sns_rpc_sns_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x6e, 0x73, 0x5f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x6e, 0x73, 0x5f, 0x72, 0x70, 0x63,
	0x22, 0x2c, 0x0a, 0x10, 0x53, 0x6e, 0x73, 0x52, 0x61, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x6b,
	0x0a, 0x03, 0x43, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x66,
	0x72, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x04, 0x70, 0x6f, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x65, 0x78, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x66, 0x6c, 0x65, 0x78, 0x22, 0x2c, 0x0a, 0x08, 0x53,
	0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x04, 0x63, 0x75, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x6e, 0x73, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x75, 0x74, 0x52, 0x04, 0x63, 0x75, 0x74, 0x73, 0x32, 0x42, 0x0a, 0x05, 0x43, 0x75, 0x74,
	0x65, 0x72, 0x12, 0x39, 0x0a, 0x03, 0x43, 0x75, 0x74, 0x12, 0x19, 0x2e, 0x73, 0x6e, 0x73, 0x5f,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x6e, 0x73, 0x52, 0x61, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x73, 0x6e, 0x73, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x31, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x66, 0x78, 0x74, 0x2e, 0x63, 0x6e, 0x2f, 0x7a,
	0x68, 0x61, 0x6e, 0x67, 0x78, 0x69, 0x61, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x6e, 0x73,
	0x2d, 0x63, 0x75, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x6e, 0x73, 0x5f, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_sns_rpc_sns_proto_rawDescOnce sync.Once
	file_pkg_sns_rpc_sns_proto_rawDescData = file_pkg_sns_rpc_sns_proto_rawDesc
)

func file_pkg_sns_rpc_sns_proto_rawDescGZIP() []byte {
	file_pkg_sns_rpc_sns_proto_rawDescOnce.Do(func() {
		file_pkg_sns_rpc_sns_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_sns_rpc_sns_proto_rawDescData)
	})
	return file_pkg_sns_rpc_sns_proto_rawDescData
}

var file_pkg_sns_rpc_sns_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_sns_rpc_sns_proto_goTypes = []interface{}{
	(*SnsRawContentReq)(nil), // 0: sns_rpc.SnsRawContentReq
	(*Cut)(nil),              // 1: sns_rpc.Cut
	(*SnsReply)(nil),         // 2: sns_rpc.SnsReply
}
var file_pkg_sns_rpc_sns_proto_depIdxs = []int32{
	1, // 0: sns_rpc.SnsReply.cuts:type_name -> sns_rpc.Cut
	0, // 1: sns_rpc.Cuter.Cut:input_type -> sns_rpc.SnsRawContentReq
	2, // 2: sns_rpc.Cuter.Cut:output_type -> sns_rpc.SnsReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_sns_rpc_sns_proto_init() }
func file_pkg_sns_rpc_sns_proto_init() {
	if File_pkg_sns_rpc_sns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_sns_rpc_sns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnsRawContentReq); i {
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
		file_pkg_sns_rpc_sns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cut); i {
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
		file_pkg_sns_rpc_sns_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnsReply); i {
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
			RawDescriptor: file_pkg_sns_rpc_sns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_sns_rpc_sns_proto_goTypes,
		DependencyIndexes: file_pkg_sns_rpc_sns_proto_depIdxs,
		MessageInfos:      file_pkg_sns_rpc_sns_proto_msgTypes,
	}.Build()
	File_pkg_sns_rpc_sns_proto = out.File
	file_pkg_sns_rpc_sns_proto_rawDesc = nil
	file_pkg_sns_rpc_sns_proto_goTypes = nil
	file_pkg_sns_rpc_sns_proto_depIdxs = nil
}