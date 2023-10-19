// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: bench.proto

package benchmark

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

type BytesData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *BytesData) Reset() {
	*x = BytesData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bench_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesData) ProtoMessage() {}

func (x *BytesData) ProtoReflect() protoreflect.Message {
	mi := &file_bench_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesData.ProtoReflect.Descriptor instead.
func (*BytesData) Descriptor() ([]byte, []int) {
	return file_bench_proto_rawDescGZIP(), []int{0}
}

func (x *BytesData) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

type I32Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I32 int32 `protobuf:"varint,1,opt,name=i32,proto3" json:"i32,omitempty"`
}

func (x *I32Data) Reset() {
	*x = I32Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bench_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *I32Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*I32Data) ProtoMessage() {}

func (x *I32Data) ProtoReflect() protoreflect.Message {
	mi := &file_bench_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use I32Data.ProtoReflect.Descriptor instead.
func (*I32Data) Descriptor() ([]byte, []int) {
	return file_bench_proto_rawDescGZIP(), []int{1}
}

func (x *I32Data) GetI32() int32 {
	if x != nil {
		return x.I32
	}
	return 0
}

type U32Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	U32 uint32 `protobuf:"varint,1,opt,name=u32,proto3" json:"u32,omitempty"`
}

func (x *U32Data) Reset() {
	*x = U32Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bench_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *U32Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*U32Data) ProtoMessage() {}

func (x *U32Data) ProtoReflect() protoreflect.Message {
	mi := &file_bench_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use U32Data.ProtoReflect.Descriptor instead.
func (*U32Data) Descriptor() ([]byte, []int) {
	return file_bench_proto_rawDescGZIP(), []int{2}
}

func (x *U32Data) GetU32() uint32 {
	if x != nil {
		return x.U32
	}
	return 0
}

type I64Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I64 int64 `protobuf:"varint,1,opt,name=i64,proto3" json:"i64,omitempty"`
}

func (x *I64Data) Reset() {
	*x = I64Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bench_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *I64Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*I64Data) ProtoMessage() {}

func (x *I64Data) ProtoReflect() protoreflect.Message {
	mi := &file_bench_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use I64Data.ProtoReflect.Descriptor instead.
func (*I64Data) Descriptor() ([]byte, []int) {
	return file_bench_proto_rawDescGZIP(), []int{3}
}

func (x *I64Data) GetI64() int64 {
	if x != nil {
		return x.I64
	}
	return 0
}

type U64Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	U64 uint64 `protobuf:"varint,1,opt,name=u64,proto3" json:"u64,omitempty"`
}

func (x *U64Data) Reset() {
	*x = U64Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bench_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *U64Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*U64Data) ProtoMessage() {}

func (x *U64Data) ProtoReflect() protoreflect.Message {
	mi := &file_bench_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use U64Data.ProtoReflect.Descriptor instead.
func (*U64Data) Descriptor() ([]byte, []int) {
	return file_bench_proto_rawDescGZIP(), []int{4}
}

func (x *U64Data) GetU64() uint64 {
	if x != nil {
		return x.U64
	}
	return 0
}

var File_bench_proto protoreflect.FileDescriptor

var file_bench_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a,
	0x09, 0x42, 0x79, 0x74, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x22, 0x1b, 0x0a, 0x07, 0x49, 0x33, 0x32, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x69, 0x33, 0x32, 0x22, 0x1b, 0x0a,
	0x07, 0x55, 0x33, 0x32, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x33, 0x32, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x33, 0x32, 0x22, 0x1b, 0x0a, 0x07, 0x49, 0x36,
	0x34, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x36, 0x34, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x69, 0x36, 0x34, 0x22, 0x1b, 0x0a, 0x07, 0x55, 0x36, 0x34, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x36, 0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x75, 0x36, 0x34, 0x32, 0x0b, 0x0a, 0x09, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72,
	0x6b, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bench_proto_rawDescOnce sync.Once
	file_bench_proto_rawDescData = file_bench_proto_rawDesc
)

func file_bench_proto_rawDescGZIP() []byte {
	file_bench_proto_rawDescOnce.Do(func() {
		file_bench_proto_rawDescData = protoimpl.X.CompressGZIP(file_bench_proto_rawDescData)
	})
	return file_bench_proto_rawDescData
}

var file_bench_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bench_proto_goTypes = []interface{}{
	(*BytesData)(nil), // 0: BytesData
	(*I32Data)(nil),   // 1: I32Data
	(*U32Data)(nil),   // 2: U32Data
	(*I64Data)(nil),   // 3: I64Data
	(*U64Data)(nil),   // 4: U64Data
}
var file_bench_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bench_proto_init() }
func file_bench_proto_init() {
	if File_bench_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bench_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BytesData); i {
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
		file_bench_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*I32Data); i {
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
		file_bench_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*U32Data); i {
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
		file_bench_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*I64Data); i {
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
		file_bench_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*U64Data); i {
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
			RawDescriptor: file_bench_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bench_proto_goTypes,
		DependencyIndexes: file_bench_proto_depIdxs,
		MessageInfos:      file_bench_proto_msgTypes,
	}.Build()
	File_bench_proto = out.File
	file_bench_proto_rawDesc = nil
	file_bench_proto_goTypes = nil
	file_bench_proto_depIdxs = nil
}
