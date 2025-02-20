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
// source: Unk2800_MHCFAGCKGIB.proto

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

type Unk2800_MHCFAGCKGIB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneId            uint32              `protobuf:"varint,12,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	PointId            uint32              `protobuf:"varint,6,opt,name=point_id,json=pointId,proto3" json:"point_id,omitempty"`
	DungeonEntryList   []*DungeonEntryInfo `protobuf:"bytes,1,rep,name=dungeon_entry_list,json=dungeonEntryList,proto3" json:"dungeon_entry_list,omitempty"`
	RecommendDungeonId uint32              `protobuf:"varint,8,opt,name=recommend_dungeon_id,json=recommendDungeonId,proto3" json:"recommend_dungeon_id,omitempty"`
}

func (x *Unk2800_MHCFAGCKGIB) Reset() {
	*x = Unk2800_MHCFAGCKGIB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2800_MHCFAGCKGIB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2800_MHCFAGCKGIB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2800_MHCFAGCKGIB) ProtoMessage() {}

func (x *Unk2800_MHCFAGCKGIB) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2800_MHCFAGCKGIB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2800_MHCFAGCKGIB.ProtoReflect.Descriptor instead.
func (*Unk2800_MHCFAGCKGIB) Descriptor() ([]byte, []int) {
	return file_Unk2800_MHCFAGCKGIB_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2800_MHCFAGCKGIB) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *Unk2800_MHCFAGCKGIB) GetPointId() uint32 {
	if x != nil {
		return x.PointId
	}
	return 0
}

func (x *Unk2800_MHCFAGCKGIB) GetDungeonEntryList() []*DungeonEntryInfo {
	if x != nil {
		return x.DungeonEntryList
	}
	return nil
}

func (x *Unk2800_MHCFAGCKGIB) GetRecommendDungeonId() uint32 {
	if x != nil {
		return x.RecommendDungeonId
	}
	return 0
}

var File_Unk2800_MHCFAGCKGIB_proto protoreflect.FileDescriptor

var file_Unk2800_MHCFAGCKGIB_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x4d, 0x48, 0x43, 0x46, 0x41, 0x47,
	0x43, 0x4b, 0x47, 0x49, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x44, 0x75, 0x6e,
	0x67, 0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f,
	0x4d, 0x48, 0x43, 0x46, 0x41, 0x47, 0x43, 0x4b, 0x47, 0x49, 0x42, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x3f, 0x0a, 0x12, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x10, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x5f,
	0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x12, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x44, 0x75, 0x6e, 0x67, 0x65,
	0x6f, 0x6e, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk2800_MHCFAGCKGIB_proto_rawDescOnce sync.Once
	file_Unk2800_MHCFAGCKGIB_proto_rawDescData = file_Unk2800_MHCFAGCKGIB_proto_rawDesc
)

func file_Unk2800_MHCFAGCKGIB_proto_rawDescGZIP() []byte {
	file_Unk2800_MHCFAGCKGIB_proto_rawDescOnce.Do(func() {
		file_Unk2800_MHCFAGCKGIB_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2800_MHCFAGCKGIB_proto_rawDescData)
	})
	return file_Unk2800_MHCFAGCKGIB_proto_rawDescData
}

var file_Unk2800_MHCFAGCKGIB_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk2800_MHCFAGCKGIB_proto_goTypes = []interface{}{
	(*Unk2800_MHCFAGCKGIB)(nil), // 0: Unk2800_MHCFAGCKGIB
	(*DungeonEntryInfo)(nil),    // 1: DungeonEntryInfo
}
var file_Unk2800_MHCFAGCKGIB_proto_depIdxs = []int32{
	1, // 0: Unk2800_MHCFAGCKGIB.dungeon_entry_list:type_name -> DungeonEntryInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Unk2800_MHCFAGCKGIB_proto_init() }
func file_Unk2800_MHCFAGCKGIB_proto_init() {
	if File_Unk2800_MHCFAGCKGIB_proto != nil {
		return
	}
	file_DungeonEntryInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk2800_MHCFAGCKGIB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2800_MHCFAGCKGIB); i {
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
			RawDescriptor: file_Unk2800_MHCFAGCKGIB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2800_MHCFAGCKGIB_proto_goTypes,
		DependencyIndexes: file_Unk2800_MHCFAGCKGIB_proto_depIdxs,
		MessageInfos:      file_Unk2800_MHCFAGCKGIB_proto_msgTypes,
	}.Build()
	File_Unk2800_MHCFAGCKGIB_proto = out.File
	file_Unk2800_MHCFAGCKGIB_proto_rawDesc = nil
	file_Unk2800_MHCFAGCKGIB_proto_goTypes = nil
	file_Unk2800_MHCFAGCKGIB_proto_depIdxs = nil
}
