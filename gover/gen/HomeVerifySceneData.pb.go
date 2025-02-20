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
// source: HomeVerifySceneData.proto

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

type HomeVerifySceneData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks   []*HomeVerifyBlockData `protobuf:"bytes,6,rep,name=blocks,proto3" json:"blocks,omitempty"`
	ModuleId uint32                 `protobuf:"varint,11,opt,name=module_id,json=moduleId,proto3" json:"module_id,omitempty"`
	SceneId  uint32                 `protobuf:"varint,4,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	Version  uint32                 `protobuf:"varint,14,opt,name=version,proto3" json:"version,omitempty"`
	IsRoom   uint32                 `protobuf:"varint,2,opt,name=is_room,json=isRoom,proto3" json:"is_room,omitempty"`
}

func (x *HomeVerifySceneData) Reset() {
	*x = HomeVerifySceneData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HomeVerifySceneData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeVerifySceneData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeVerifySceneData) ProtoMessage() {}

func (x *HomeVerifySceneData) ProtoReflect() protoreflect.Message {
	mi := &file_HomeVerifySceneData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeVerifySceneData.ProtoReflect.Descriptor instead.
func (*HomeVerifySceneData) Descriptor() ([]byte, []int) {
	return file_HomeVerifySceneData_proto_rawDescGZIP(), []int{0}
}

func (x *HomeVerifySceneData) GetBlocks() []*HomeVerifyBlockData {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *HomeVerifySceneData) GetModuleId() uint32 {
	if x != nil {
		return x.ModuleId
	}
	return 0
}

func (x *HomeVerifySceneData) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *HomeVerifySceneData) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *HomeVerifySceneData) GetIsRoom() uint32 {
	if x != nil {
		return x.IsRoom
	}
	return 0
}

var File_HomeVerifySceneData_proto protoreflect.FileDescriptor

var file_HomeVerifySceneData_proto_rawDesc = []byte{
	0x0a, 0x19, 0x48, 0x6f, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x48, 0x6f, 0x6d,
	0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x13, 0x48, 0x6f, 0x6d, 0x65, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c,
	0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65,
	0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65,
	0x6e, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17,
	0x0a, 0x07, 0x69, 0x73, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x69, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HomeVerifySceneData_proto_rawDescOnce sync.Once
	file_HomeVerifySceneData_proto_rawDescData = file_HomeVerifySceneData_proto_rawDesc
)

func file_HomeVerifySceneData_proto_rawDescGZIP() []byte {
	file_HomeVerifySceneData_proto_rawDescOnce.Do(func() {
		file_HomeVerifySceneData_proto_rawDescData = protoimpl.X.CompressGZIP(file_HomeVerifySceneData_proto_rawDescData)
	})
	return file_HomeVerifySceneData_proto_rawDescData
}

var file_HomeVerifySceneData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HomeVerifySceneData_proto_goTypes = []interface{}{
	(*HomeVerifySceneData)(nil), // 0: HomeVerifySceneData
	(*HomeVerifyBlockData)(nil), // 1: HomeVerifyBlockData
}
var file_HomeVerifySceneData_proto_depIdxs = []int32{
	1, // 0: HomeVerifySceneData.blocks:type_name -> HomeVerifyBlockData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HomeVerifySceneData_proto_init() }
func file_HomeVerifySceneData_proto_init() {
	if File_HomeVerifySceneData_proto != nil {
		return
	}
	file_HomeVerifyBlockData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HomeVerifySceneData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeVerifySceneData); i {
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
			RawDescriptor: file_HomeVerifySceneData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HomeVerifySceneData_proto_goTypes,
		DependencyIndexes: file_HomeVerifySceneData_proto_depIdxs,
		MessageInfos:      file_HomeVerifySceneData_proto_msgTypes,
	}.Build()
	File_HomeVerifySceneData_proto = out.File
	file_HomeVerifySceneData_proto_rawDesc = nil
	file_HomeVerifySceneData_proto_goTypes = nil
	file_HomeVerifySceneData_proto_depIdxs = nil
}
