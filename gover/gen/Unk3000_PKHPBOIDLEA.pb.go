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
// source: Unk3000_PKHPBOIDLEA.proto

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

type Unk3000_PKHPBOIDLEA int32

const (
	Unk3000_PKHPBOIDLEA_Unk3000_PKHPBOIDLEA_Unk3000_KANMGBLJEHC Unk3000_PKHPBOIDLEA = 0
	Unk3000_PKHPBOIDLEA_Unk3000_PKHPBOIDLEA_Unk3000_ICFILKDKFNL Unk3000_PKHPBOIDLEA = 1
	Unk3000_PKHPBOIDLEA_Unk3000_PKHPBOIDLEA_Unk3000_FBFKPBGLMAD Unk3000_PKHPBOIDLEA = 2
	Unk3000_PKHPBOIDLEA_Unk3000_PKHPBOIDLEA_Unk3000_KEOIEIKLFDN Unk3000_PKHPBOIDLEA = 3
)

// Enum value maps for Unk3000_PKHPBOIDLEA.
var (
	Unk3000_PKHPBOIDLEA_name = map[int32]string{
		0: "Unk3000_PKHPBOIDLEA_Unk3000_KANMGBLJEHC",
		1: "Unk3000_PKHPBOIDLEA_Unk3000_ICFILKDKFNL",
		2: "Unk3000_PKHPBOIDLEA_Unk3000_FBFKPBGLMAD",
		3: "Unk3000_PKHPBOIDLEA_Unk3000_KEOIEIKLFDN",
	}
	Unk3000_PKHPBOIDLEA_value = map[string]int32{
		"Unk3000_PKHPBOIDLEA_Unk3000_KANMGBLJEHC": 0,
		"Unk3000_PKHPBOIDLEA_Unk3000_ICFILKDKFNL": 1,
		"Unk3000_PKHPBOIDLEA_Unk3000_FBFKPBGLMAD": 2,
		"Unk3000_PKHPBOIDLEA_Unk3000_KEOIEIKLFDN": 3,
	}
)

func (x Unk3000_PKHPBOIDLEA) Enum() *Unk3000_PKHPBOIDLEA {
	p := new(Unk3000_PKHPBOIDLEA)
	*p = x
	return p
}

func (x Unk3000_PKHPBOIDLEA) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Unk3000_PKHPBOIDLEA) Descriptor() protoreflect.EnumDescriptor {
	return file_Unk3000_PKHPBOIDLEA_proto_enumTypes[0].Descriptor()
}

func (Unk3000_PKHPBOIDLEA) Type() protoreflect.EnumType {
	return &file_Unk3000_PKHPBOIDLEA_proto_enumTypes[0]
}

func (x Unk3000_PKHPBOIDLEA) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unk3000_PKHPBOIDLEA.Descriptor instead.
func (Unk3000_PKHPBOIDLEA) EnumDescriptor() ([]byte, []int) {
	return file_Unk3000_PKHPBOIDLEA_proto_rawDescGZIP(), []int{0}
}

var File_Unk3000_PKHPBOIDLEA_proto protoreflect.FileDescriptor

var file_Unk3000_PKHPBOIDLEA_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x50, 0x4b, 0x48, 0x50, 0x42, 0x4f,
	0x49, 0x44, 0x4c, 0x45, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xc9, 0x01, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x50, 0x4b, 0x48, 0x50, 0x42, 0x4f, 0x49, 0x44,
	0x4c, 0x45, 0x41, 0x12, 0x2b, 0x0a, 0x27, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x50,
	0x4b, 0x48, 0x50, 0x42, 0x4f, 0x49, 0x44, 0x4c, 0x45, 0x41, 0x5f, 0x55, 0x6e, 0x6b, 0x33, 0x30,
	0x30, 0x30, 0x5f, 0x4b, 0x41, 0x4e, 0x4d, 0x47, 0x42, 0x4c, 0x4a, 0x45, 0x48, 0x43, 0x10, 0x00,
	0x12, 0x2b, 0x0a, 0x27, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x50, 0x4b, 0x48, 0x50,
	0x42, 0x4f, 0x49, 0x44, 0x4c, 0x45, 0x41, 0x5f, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f,
	0x49, 0x43, 0x46, 0x49, 0x4c, 0x4b, 0x44, 0x4b, 0x46, 0x4e, 0x4c, 0x10, 0x01, 0x12, 0x2b, 0x0a,
	0x27, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x50, 0x4b, 0x48, 0x50, 0x42, 0x4f, 0x49,
	0x44, 0x4c, 0x45, 0x41, 0x5f, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x46, 0x42, 0x46,
	0x4b, 0x50, 0x42, 0x47, 0x4c, 0x4d, 0x41, 0x44, 0x10, 0x02, 0x12, 0x2b, 0x0a, 0x27, 0x55, 0x6e,
	0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x50, 0x4b, 0x48, 0x50, 0x42, 0x4f, 0x49, 0x44, 0x4c, 0x45,
	0x41, 0x5f, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x4b, 0x45, 0x4f, 0x49, 0x45, 0x49,
	0x4b, 0x4c, 0x46, 0x44, 0x4e, 0x10, 0x03, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk3000_PKHPBOIDLEA_proto_rawDescOnce sync.Once
	file_Unk3000_PKHPBOIDLEA_proto_rawDescData = file_Unk3000_PKHPBOIDLEA_proto_rawDesc
)

func file_Unk3000_PKHPBOIDLEA_proto_rawDescGZIP() []byte {
	file_Unk3000_PKHPBOIDLEA_proto_rawDescOnce.Do(func() {
		file_Unk3000_PKHPBOIDLEA_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk3000_PKHPBOIDLEA_proto_rawDescData)
	})
	return file_Unk3000_PKHPBOIDLEA_proto_rawDescData
}

var file_Unk3000_PKHPBOIDLEA_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Unk3000_PKHPBOIDLEA_proto_goTypes = []interface{}{
	(Unk3000_PKHPBOIDLEA)(0), // 0: Unk3000_PKHPBOIDLEA
}
var file_Unk3000_PKHPBOIDLEA_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Unk3000_PKHPBOIDLEA_proto_init() }
func file_Unk3000_PKHPBOIDLEA_proto_init() {
	if File_Unk3000_PKHPBOIDLEA_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Unk3000_PKHPBOIDLEA_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk3000_PKHPBOIDLEA_proto_goTypes,
		DependencyIndexes: file_Unk3000_PKHPBOIDLEA_proto_depIdxs,
		EnumInfos:         file_Unk3000_PKHPBOIDLEA_proto_enumTypes,
	}.Build()
	File_Unk3000_PKHPBOIDLEA_proto = out.File
	file_Unk3000_PKHPBOIDLEA_proto_rawDesc = nil
	file_Unk3000_PKHPBOIDLEA_proto_goTypes = nil
	file_Unk3000_PKHPBOIDLEA_proto_depIdxs = nil
}
