// Sorapointa - A server software re-implementation for a certain anime game,
// and avoid sorapointa. Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: DropSubfieldType.proto

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

type DropSubfieldType int32

const (
	DropSubfieldType_DROP_SUBFIELD_TYPE_NONE                DropSubfieldType = 0
	DropSubfieldType_DROP_SUBFIELD_TYPE_ONE                 DropSubfieldType = 1
	DropSubfieldType_DROP_SUBFIELD_TYPE_Unk2700_NNGMHCEADHE DropSubfieldType = 2
	DropSubfieldType_DROP_SUBFIELD_TYPE_Unk2700_MKIJPEHKAJI DropSubfieldType = 3
	DropSubfieldType_DROP_SUBFIELD_TYPE_Unk2700_DJDNENLGIEB DropSubfieldType = 4
)

// Enum value maps for DropSubfieldType.
var (
	DropSubfieldType_name = map[int32]string{
		0: "DROP_SUBFIELD_TYPE_NONE",
		1: "DROP_SUBFIELD_TYPE_ONE",
		2: "DROP_SUBFIELD_TYPE_Unk2700_NNGMHCEADHE",
		3: "DROP_SUBFIELD_TYPE_Unk2700_MKIJPEHKAJI",
		4: "DROP_SUBFIELD_TYPE_Unk2700_DJDNENLGIEB",
	}
	DropSubfieldType_value = map[string]int32{
		"DROP_SUBFIELD_TYPE_NONE":                0,
		"DROP_SUBFIELD_TYPE_ONE":                 1,
		"DROP_SUBFIELD_TYPE_Unk2700_NNGMHCEADHE": 2,
		"DROP_SUBFIELD_TYPE_Unk2700_MKIJPEHKAJI": 3,
		"DROP_SUBFIELD_TYPE_Unk2700_DJDNENLGIEB": 4,
	}
)

func (x DropSubfieldType) Enum() *DropSubfieldType {
	p := new(DropSubfieldType)
	*p = x
	return p
}

func (x DropSubfieldType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DropSubfieldType) Descriptor() protoreflect.EnumDescriptor {
	return file_DropSubfieldType_proto_enumTypes[0].Descriptor()
}

func (DropSubfieldType) Type() protoreflect.EnumType {
	return &file_DropSubfieldType_proto_enumTypes[0]
}

func (x DropSubfieldType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DropSubfieldType.Descriptor instead.
func (DropSubfieldType) EnumDescriptor() ([]byte, []int) {
	return file_DropSubfieldType_proto_rawDescGZIP(), []int{0}
}

var File_DropSubfieldType_proto protoreflect.FileDescriptor

var file_DropSubfieldType_proto_rawDesc = []byte{
	0x0a, 0x16, 0x44, 0x72, 0x6f, 0x70, 0x53, 0x75, 0x62, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xcf, 0x01, 0x0a, 0x10, 0x44, 0x72, 0x6f,
	0x70, 0x53, 0x75, 0x62, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a,
	0x17, 0x44, 0x52, 0x4f, 0x50, 0x5f, 0x53, 0x55, 0x42, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x52,
	0x4f, 0x50, 0x5f, 0x53, 0x55, 0x42, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x2a, 0x0a, 0x26, 0x44, 0x52, 0x4f, 0x50, 0x5f, 0x53,
	0x55, 0x42, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x4e, 0x4e, 0x47, 0x4d, 0x48, 0x43, 0x45, 0x41, 0x44, 0x48, 0x45,
	0x10, 0x02, 0x12, 0x2a, 0x0a, 0x26, 0x44, 0x52, 0x4f, 0x50, 0x5f, 0x53, 0x55, 0x42, 0x46, 0x49,
	0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x4d, 0x4b, 0x49, 0x4a, 0x50, 0x45, 0x48, 0x4b, 0x41, 0x4a, 0x49, 0x10, 0x03, 0x12, 0x2a,
	0x0a, 0x26, 0x44, 0x52, 0x4f, 0x50, 0x5f, 0x53, 0x55, 0x42, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x44, 0x4a, 0x44,
	0x4e, 0x45, 0x4e, 0x4c, 0x47, 0x49, 0x45, 0x42, 0x10, 0x04, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DropSubfieldType_proto_rawDescOnce sync.Once
	file_DropSubfieldType_proto_rawDescData = file_DropSubfieldType_proto_rawDesc
)

func file_DropSubfieldType_proto_rawDescGZIP() []byte {
	file_DropSubfieldType_proto_rawDescOnce.Do(func() {
		file_DropSubfieldType_proto_rawDescData = protoimpl.X.CompressGZIP(file_DropSubfieldType_proto_rawDescData)
	})
	return file_DropSubfieldType_proto_rawDescData
}

var file_DropSubfieldType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_DropSubfieldType_proto_goTypes = []interface{}{
	(DropSubfieldType)(0), // 0: DropSubfieldType
}
var file_DropSubfieldType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DropSubfieldType_proto_init() }
func file_DropSubfieldType_proto_init() {
	if File_DropSubfieldType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_DropSubfieldType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DropSubfieldType_proto_goTypes,
		DependencyIndexes: file_DropSubfieldType_proto_depIdxs,
		EnumInfos:         file_DropSubfieldType_proto_enumTypes,
	}.Build()
	File_DropSubfieldType_proto = out.File
	file_DropSubfieldType_proto_rawDesc = nil
	file_DropSubfieldType_proto_goTypes = nil
	file_DropSubfieldType_proto_depIdxs = nil
}
