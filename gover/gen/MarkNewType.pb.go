// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: MarkNewType.proto

package gen

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

type MarkNewType int32

const (
	MarkNewType__NONE   MarkNewType = 0
	MarkNewType_COMBINE MarkNewType = 1
	MarkNewType_FORGE   MarkNewType = 2
)

// Enum value maps for MarkNewType.
var (
	MarkNewType_name = map[int32]string{
		0: "_NONE",
		1: "COMBINE",
		2: "FORGE",
	}
	MarkNewType_value = map[string]int32{
		"_NONE":   0,
		"COMBINE": 1,
		"FORGE":   2,
	}
)

func (x MarkNewType) Enum() *MarkNewType {
	p := new(MarkNewType)
	*p = x
	return p
}

func (x MarkNewType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MarkNewType) Descriptor() protoreflect.EnumDescriptor {
	return file_MarkNewType_proto_enumTypes[0].Descriptor()
}

func (MarkNewType) Type() protoreflect.EnumType {
	return &file_MarkNewType_proto_enumTypes[0]
}

func (x MarkNewType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MarkNewType.Descriptor instead.
func (MarkNewType) EnumDescriptor() ([]byte, []int) {
	return file_MarkNewType_proto_rawDescGZIP(), []int{0}
}

var File_MarkNewType_proto protoreflect.FileDescriptor

var file_MarkNewType_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4d, 0x61, 0x72, 0x6b, 0x4e, 0x65, 0x77, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x30, 0x0a, 0x0b, 0x4d, 0x61, 0x72, 0x6b, 0x4e, 0x65, 0x77, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x4f, 0x4d, 0x42, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x4f,
	0x52, 0x47, 0x45, 0x10, 0x02, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MarkNewType_proto_rawDescOnce sync.Once
	file_MarkNewType_proto_rawDescData = file_MarkNewType_proto_rawDesc
)

func file_MarkNewType_proto_rawDescGZIP() []byte {
	file_MarkNewType_proto_rawDescOnce.Do(func() {
		file_MarkNewType_proto_rawDescData = protoimpl.X.CompressGZIP(file_MarkNewType_proto_rawDescData)
	})
	return file_MarkNewType_proto_rawDescData
}

var file_MarkNewType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MarkNewType_proto_goTypes = []interface{}{
	(MarkNewType)(0), // 0: MarkNewType
}
var file_MarkNewType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MarkNewType_proto_init() }
func file_MarkNewType_proto_init() {
	if File_MarkNewType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MarkNewType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MarkNewType_proto_goTypes,
		DependencyIndexes: file_MarkNewType_proto_depIdxs,
		EnumInfos:         file_MarkNewType_proto_enumTypes,
	}.Build()
	File_MarkNewType_proto = out.File
	file_MarkNewType_proto_rawDesc = nil
	file_MarkNewType_proto_goTypes = nil
	file_MarkNewType_proto_depIdxs = nil
}
