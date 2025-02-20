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
// source: SceneGallerySumoInfo.proto

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

type SceneGallerySumoInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score                uint32 `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	KillNormalMosnterNum uint32 `protobuf:"varint,15,opt,name=kill_normal_mosnter_num,json=killNormalMosnterNum,proto3" json:"kill_normal_mosnter_num,omitempty"`
	KillEliteMonsterNum  uint32 `protobuf:"varint,14,opt,name=kill_elite_monster_num,json=killEliteMonsterNum,proto3" json:"kill_elite_monster_num,omitempty"`
}

func (x *SceneGallerySumoInfo) Reset() {
	*x = SceneGallerySumoInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneGallerySumoInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneGallerySumoInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneGallerySumoInfo) ProtoMessage() {}

func (x *SceneGallerySumoInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SceneGallerySumoInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneGallerySumoInfo.ProtoReflect.Descriptor instead.
func (*SceneGallerySumoInfo) Descriptor() ([]byte, []int) {
	return file_SceneGallerySumoInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneGallerySumoInfo) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *SceneGallerySumoInfo) GetKillNormalMosnterNum() uint32 {
	if x != nil {
		return x.KillNormalMosnterNum
	}
	return 0
}

func (x *SceneGallerySumoInfo) GetKillEliteMonsterNum() uint32 {
	if x != nil {
		return x.KillEliteMonsterNum
	}
	return 0
}

var File_SceneGallerySumoInfo_proto protoreflect.FileDescriptor

var file_SceneGallerySumoInfo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x75,
	0x6d, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a,
	0x14, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x53, 0x75, 0x6d,
	0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x6b,
	0x69, 0x6c, 0x6c, 0x5f, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x5f, 0x6d, 0x6f, 0x73, 0x6e, 0x74,
	0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x6b, 0x69,
	0x6c, 0x6c, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x4d, 0x6f, 0x73, 0x6e, 0x74, 0x65, 0x72, 0x4e,
	0x75, 0x6d, 0x12, 0x33, 0x0a, 0x16, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x65, 0x6c, 0x69, 0x74, 0x65,
	0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x13, 0x6b, 0x69, 0x6c, 0x6c, 0x45, 0x6c, 0x69, 0x74, 0x65, 0x4d, 0x6f, 0x6e,
	0x73, 0x74, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneGallerySumoInfo_proto_rawDescOnce sync.Once
	file_SceneGallerySumoInfo_proto_rawDescData = file_SceneGallerySumoInfo_proto_rawDesc
)

func file_SceneGallerySumoInfo_proto_rawDescGZIP() []byte {
	file_SceneGallerySumoInfo_proto_rawDescOnce.Do(func() {
		file_SceneGallerySumoInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneGallerySumoInfo_proto_rawDescData)
	})
	return file_SceneGallerySumoInfo_proto_rawDescData
}

var file_SceneGallerySumoInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneGallerySumoInfo_proto_goTypes = []interface{}{
	(*SceneGallerySumoInfo)(nil), // 0: SceneGallerySumoInfo
}
var file_SceneGallerySumoInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SceneGallerySumoInfo_proto_init() }
func file_SceneGallerySumoInfo_proto_init() {
	if File_SceneGallerySumoInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SceneGallerySumoInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneGallerySumoInfo); i {
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
			RawDescriptor: file_SceneGallerySumoInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneGallerySumoInfo_proto_goTypes,
		DependencyIndexes: file_SceneGallerySumoInfo_proto_depIdxs,
		MessageInfos:      file_SceneGallerySumoInfo_proto_msgTypes,
	}.Build()
	File_SceneGallerySumoInfo_proto = out.File
	file_SceneGallerySumoInfo_proto_rawDesc = nil
	file_SceneGallerySumoInfo_proto_goTypes = nil
	file_SceneGallerySumoInfo_proto_depIdxs = nil
}
