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
// source: SceneSurfaceMaterial.proto

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

type SceneSurfaceMaterial int32

const (
	SceneSurfaceMaterial_SCENE_SURFACE_MATERIAL_INVALID SceneSurfaceMaterial = 0
	SceneSurfaceMaterial_SCENE_SURFACE_MATERIAL_GRASS   SceneSurfaceMaterial = 1
	SceneSurfaceMaterial_SCENE_SURFACE_MATERIAL_DIRT    SceneSurfaceMaterial = 2
	SceneSurfaceMaterial_SCENE_SURFACE_MATERIAL_ROCK    SceneSurfaceMaterial = 3
	SceneSurfaceMaterial_SCENE_SURFACE_MATERIAL_SNOW    SceneSurfaceMaterial = 4
	SceneSurfaceMaterial_SCENE_SURFACE_MATERIAL_WATER   SceneSurfaceMaterial = 5
	SceneSurfaceMaterial_SCENE_SURFACE_MATERIAL_TILE    SceneSurfaceMaterial = 6
	SceneSurfaceMaterial_SCENE_SURFACE_MATERIAL_SAND    SceneSurfaceMaterial = 7
)

// Enum value maps for SceneSurfaceMaterial.
var (
	SceneSurfaceMaterial_name = map[int32]string{
		0: "SCENE_SURFACE_MATERIAL_INVALID",
		1: "SCENE_SURFACE_MATERIAL_GRASS",
		2: "SCENE_SURFACE_MATERIAL_DIRT",
		3: "SCENE_SURFACE_MATERIAL_ROCK",
		4: "SCENE_SURFACE_MATERIAL_SNOW",
		5: "SCENE_SURFACE_MATERIAL_WATER",
		6: "SCENE_SURFACE_MATERIAL_TILE",
		7: "SCENE_SURFACE_MATERIAL_SAND",
	}
	SceneSurfaceMaterial_value = map[string]int32{
		"SCENE_SURFACE_MATERIAL_INVALID": 0,
		"SCENE_SURFACE_MATERIAL_GRASS":   1,
		"SCENE_SURFACE_MATERIAL_DIRT":    2,
		"SCENE_SURFACE_MATERIAL_ROCK":    3,
		"SCENE_SURFACE_MATERIAL_SNOW":    4,
		"SCENE_SURFACE_MATERIAL_WATER":   5,
		"SCENE_SURFACE_MATERIAL_TILE":    6,
		"SCENE_SURFACE_MATERIAL_SAND":    7,
	}
)

func (x SceneSurfaceMaterial) Enum() *SceneSurfaceMaterial {
	p := new(SceneSurfaceMaterial)
	*p = x
	return p
}

func (x SceneSurfaceMaterial) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SceneSurfaceMaterial) Descriptor() protoreflect.EnumDescriptor {
	return file_SceneSurfaceMaterial_proto_enumTypes[0].Descriptor()
}

func (SceneSurfaceMaterial) Type() protoreflect.EnumType {
	return &file_SceneSurfaceMaterial_proto_enumTypes[0]
}

func (x SceneSurfaceMaterial) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SceneSurfaceMaterial.Descriptor instead.
func (SceneSurfaceMaterial) EnumDescriptor() ([]byte, []int) {
	return file_SceneSurfaceMaterial_proto_rawDescGZIP(), []int{0}
}

var File_SceneSurfaceMaterial_proto protoreflect.FileDescriptor

var file_SceneSurfaceMaterial_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x53, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4d, 0x61,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xa3, 0x02, 0x0a,
	0x14, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x53, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4d, 0x61, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x5f, 0x53,
	0x55, 0x52, 0x46, 0x41, 0x43, 0x45, 0x5f, 0x4d, 0x41, 0x54, 0x45, 0x52, 0x49, 0x41, 0x4c, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x43, 0x45,
	0x4e, 0x45, 0x5f, 0x53, 0x55, 0x52, 0x46, 0x41, 0x43, 0x45, 0x5f, 0x4d, 0x41, 0x54, 0x45, 0x52,
	0x49, 0x41, 0x4c, 0x5f, 0x47, 0x52, 0x41, 0x53, 0x53, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x53,
	0x43, 0x45, 0x4e, 0x45, 0x5f, 0x53, 0x55, 0x52, 0x46, 0x41, 0x43, 0x45, 0x5f, 0x4d, 0x41, 0x54,
	0x45, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x44, 0x49, 0x52, 0x54, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b,
	0x53, 0x43, 0x45, 0x4e, 0x45, 0x5f, 0x53, 0x55, 0x52, 0x46, 0x41, 0x43, 0x45, 0x5f, 0x4d, 0x41,
	0x54, 0x45, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x52, 0x4f, 0x43, 0x4b, 0x10, 0x03, 0x12, 0x1f, 0x0a,
	0x1b, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x5f, 0x53, 0x55, 0x52, 0x46, 0x41, 0x43, 0x45, 0x5f, 0x4d,
	0x41, 0x54, 0x45, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x53, 0x4e, 0x4f, 0x57, 0x10, 0x04, 0x12, 0x20,
	0x0a, 0x1c, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x5f, 0x53, 0x55, 0x52, 0x46, 0x41, 0x43, 0x45, 0x5f,
	0x4d, 0x41, 0x54, 0x45, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x57, 0x41, 0x54, 0x45, 0x52, 0x10, 0x05,
	0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x5f, 0x53, 0x55, 0x52, 0x46, 0x41, 0x43,
	0x45, 0x5f, 0x4d, 0x41, 0x54, 0x45, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x54, 0x49, 0x4c, 0x45, 0x10,
	0x06, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x5f, 0x53, 0x55, 0x52, 0x46, 0x41,
	0x43, 0x45, 0x5f, 0x4d, 0x41, 0x54, 0x45, 0x52, 0x49, 0x41, 0x4c, 0x5f, 0x53, 0x41, 0x4e, 0x44,
	0x10, 0x07, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_SceneSurfaceMaterial_proto_rawDescOnce sync.Once
	file_SceneSurfaceMaterial_proto_rawDescData = file_SceneSurfaceMaterial_proto_rawDesc
)

func file_SceneSurfaceMaterial_proto_rawDescGZIP() []byte {
	file_SceneSurfaceMaterial_proto_rawDescOnce.Do(func() {
		file_SceneSurfaceMaterial_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneSurfaceMaterial_proto_rawDescData)
	})
	return file_SceneSurfaceMaterial_proto_rawDescData
}

var file_SceneSurfaceMaterial_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SceneSurfaceMaterial_proto_goTypes = []interface{}{
	(SceneSurfaceMaterial)(0), // 0: SceneSurfaceMaterial
}
var file_SceneSurfaceMaterial_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SceneSurfaceMaterial_proto_init() }
func file_SceneSurfaceMaterial_proto_init() {
	if File_SceneSurfaceMaterial_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SceneSurfaceMaterial_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneSurfaceMaterial_proto_goTypes,
		DependencyIndexes: file_SceneSurfaceMaterial_proto_depIdxs,
		EnumInfos:         file_SceneSurfaceMaterial_proto_enumTypes,
	}.Build()
	File_SceneSurfaceMaterial_proto = out.File
	file_SceneSurfaceMaterial_proto_rawDesc = nil
	file_SceneSurfaceMaterial_proto_goTypes = nil
	file_SceneSurfaceMaterial_proto_depIdxs = nil
}
