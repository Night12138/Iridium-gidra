// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: LunchBoxSlotType.proto

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

type LunchBoxSlotType int32

const (
	LunchBoxSlotType_LUNCH_BOX_SLOT_NONE   LunchBoxSlotType = 0
	LunchBoxSlotType_LUNCH_BOX_SLOT_REVIVE LunchBoxSlotType = 1
	LunchBoxSlotType_LUNCH_BOX_SLOT_HEAL   LunchBoxSlotType = 2
)

// Enum value maps for LunchBoxSlotType.
var (
	LunchBoxSlotType_name = map[int32]string{
		0: "LUNCH_BOX_SLOT_NONE",
		1: "LUNCH_BOX_SLOT_REVIVE",
		2: "LUNCH_BOX_SLOT_HEAL",
	}
	LunchBoxSlotType_value = map[string]int32{
		"LUNCH_BOX_SLOT_NONE":   0,
		"LUNCH_BOX_SLOT_REVIVE": 1,
		"LUNCH_BOX_SLOT_HEAL":   2,
	}
)

func (x LunchBoxSlotType) Enum() *LunchBoxSlotType {
	p := new(LunchBoxSlotType)
	*p = x
	return p
}

func (x LunchBoxSlotType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LunchBoxSlotType) Descriptor() protoreflect.EnumDescriptor {
	return file_LunchBoxSlotType_proto_enumTypes[0].Descriptor()
}

func (LunchBoxSlotType) Type() protoreflect.EnumType {
	return &file_LunchBoxSlotType_proto_enumTypes[0]
}

func (x LunchBoxSlotType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LunchBoxSlotType.Descriptor instead.
func (LunchBoxSlotType) EnumDescriptor() ([]byte, []int) {
	return file_LunchBoxSlotType_proto_rawDescGZIP(), []int{0}
}

var File_LunchBoxSlotType_proto protoreflect.FileDescriptor

var file_LunchBoxSlotType_proto_rawDesc = []byte{
	0x0a, 0x16, 0x4c, 0x75, 0x6e, 0x63, 0x68, 0x42, 0x6f, 0x78, 0x53, 0x6c, 0x6f, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x5f, 0x0a, 0x10, 0x4c, 0x75, 0x6e, 0x63,
	0x68, 0x42, 0x6f, 0x78, 0x53, 0x6c, 0x6f, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13,
	0x4c, 0x55, 0x4e, 0x43, 0x48, 0x5f, 0x42, 0x4f, 0x58, 0x5f, 0x53, 0x4c, 0x4f, 0x54, 0x5f, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x55, 0x4e, 0x43, 0x48, 0x5f, 0x42,
	0x4f, 0x58, 0x5f, 0x53, 0x4c, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x56, 0x45, 0x10, 0x01,
	0x12, 0x17, 0x0a, 0x13, 0x4c, 0x55, 0x4e, 0x43, 0x48, 0x5f, 0x42, 0x4f, 0x58, 0x5f, 0x53, 0x4c,
	0x4f, 0x54, 0x5f, 0x48, 0x45, 0x41, 0x4c, 0x10, 0x02, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LunchBoxSlotType_proto_rawDescOnce sync.Once
	file_LunchBoxSlotType_proto_rawDescData = file_LunchBoxSlotType_proto_rawDesc
)

func file_LunchBoxSlotType_proto_rawDescGZIP() []byte {
	file_LunchBoxSlotType_proto_rawDescOnce.Do(func() {
		file_LunchBoxSlotType_proto_rawDescData = protoimpl.X.CompressGZIP(file_LunchBoxSlotType_proto_rawDescData)
	})
	return file_LunchBoxSlotType_proto_rawDescData
}

var file_LunchBoxSlotType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_LunchBoxSlotType_proto_goTypes = []interface{}{
	(LunchBoxSlotType)(0), // 0: LunchBoxSlotType
}
var file_LunchBoxSlotType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_LunchBoxSlotType_proto_init() }
func file_LunchBoxSlotType_proto_init() {
	if File_LunchBoxSlotType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_LunchBoxSlotType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LunchBoxSlotType_proto_goTypes,
		DependencyIndexes: file_LunchBoxSlotType_proto_depIdxs,
		EnumInfos:         file_LunchBoxSlotType_proto_enumTypes,
	}.Build()
	File_LunchBoxSlotType_proto = out.File
	file_LunchBoxSlotType_proto_rawDesc = nil
	file_LunchBoxSlotType_proto_goTypes = nil
	file_LunchBoxSlotType_proto_depIdxs = nil
}
