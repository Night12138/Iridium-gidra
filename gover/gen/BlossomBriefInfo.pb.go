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
// source: BlossomBriefInfo.proto

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

type BlossomBriefInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshId     uint32  `protobuf:"varint,13,opt,name=refresh_id,json=refreshId,proto3" json:"refresh_id,omitempty"`
	RewardId      uint32  `protobuf:"varint,5,opt,name=reward_id,json=rewardId,proto3" json:"reward_id,omitempty"`
	CityId        uint32  `protobuf:"varint,10,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
	Resin         uint32  `protobuf:"varint,11,opt,name=resin,proto3" json:"resin,omitempty"`
	State         uint32  `protobuf:"varint,7,opt,name=state,proto3" json:"state,omitempty"`
	IsGuideOpened bool    `protobuf:"varint,1,opt,name=is_guide_opened,json=isGuideOpened,proto3" json:"is_guide_opened,omitempty"`
	MonsterLevel  uint32  `protobuf:"varint,8,opt,name=monster_level,json=monsterLevel,proto3" json:"monster_level,omitempty"`
	CircleCampId  uint32  `protobuf:"varint,15,opt,name=circle_camp_id,json=circleCampId,proto3" json:"circle_camp_id,omitempty"`
	Pos           *Vector `protobuf:"bytes,12,opt,name=pos,proto3" json:"pos,omitempty"`
	SceneId       uint32  `protobuf:"varint,9,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
}

func (x *BlossomBriefInfo) Reset() {
	*x = BlossomBriefInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BlossomBriefInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlossomBriefInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlossomBriefInfo) ProtoMessage() {}

func (x *BlossomBriefInfo) ProtoReflect() protoreflect.Message {
	mi := &file_BlossomBriefInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlossomBriefInfo.ProtoReflect.Descriptor instead.
func (*BlossomBriefInfo) Descriptor() ([]byte, []int) {
	return file_BlossomBriefInfo_proto_rawDescGZIP(), []int{0}
}

func (x *BlossomBriefInfo) GetRefreshId() uint32 {
	if x != nil {
		return x.RefreshId
	}
	return 0
}

func (x *BlossomBriefInfo) GetRewardId() uint32 {
	if x != nil {
		return x.RewardId
	}
	return 0
}

func (x *BlossomBriefInfo) GetCityId() uint32 {
	if x != nil {
		return x.CityId
	}
	return 0
}

func (x *BlossomBriefInfo) GetResin() uint32 {
	if x != nil {
		return x.Resin
	}
	return 0
}

func (x *BlossomBriefInfo) GetState() uint32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *BlossomBriefInfo) GetIsGuideOpened() bool {
	if x != nil {
		return x.IsGuideOpened
	}
	return false
}

func (x *BlossomBriefInfo) GetMonsterLevel() uint32 {
	if x != nil {
		return x.MonsterLevel
	}
	return 0
}

func (x *BlossomBriefInfo) GetCircleCampId() uint32 {
	if x != nil {
		return x.CircleCampId
	}
	return 0
}

func (x *BlossomBriefInfo) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *BlossomBriefInfo) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

var File_BlossomBriefInfo_proto protoreflect.FileDescriptor

var file_BlossomBriefInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x42, 0x6c, 0x6f, 0x73, 0x73, 0x6f, 0x6d, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x02, 0x0a, 0x10, 0x42, 0x6c, 0x6f, 0x73, 0x73,
	0x6f, 0x6d, 0x42, 0x72, 0x69, 0x65, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x73, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x72, 0x65, 0x73, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0f,
	0x69, 0x73, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x47, 0x75, 0x69, 0x64, 0x65, 0x4f, 0x70,
	0x65, 0x6e, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x6f, 0x6e,
	0x73, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x69, 0x72,
	0x63, 0x6c, 0x65, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x63, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x6d, 0x70, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x49, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BlossomBriefInfo_proto_rawDescOnce sync.Once
	file_BlossomBriefInfo_proto_rawDescData = file_BlossomBriefInfo_proto_rawDesc
)

func file_BlossomBriefInfo_proto_rawDescGZIP() []byte {
	file_BlossomBriefInfo_proto_rawDescOnce.Do(func() {
		file_BlossomBriefInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_BlossomBriefInfo_proto_rawDescData)
	})
	return file_BlossomBriefInfo_proto_rawDescData
}

var file_BlossomBriefInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BlossomBriefInfo_proto_goTypes = []interface{}{
	(*BlossomBriefInfo)(nil), // 0: BlossomBriefInfo
	(*Vector)(nil),           // 1: Vector
}
var file_BlossomBriefInfo_proto_depIdxs = []int32{
	1, // 0: BlossomBriefInfo.pos:type_name -> Vector
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BlossomBriefInfo_proto_init() }
func file_BlossomBriefInfo_proto_init() {
	if File_BlossomBriefInfo_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BlossomBriefInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlossomBriefInfo); i {
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
			RawDescriptor: file_BlossomBriefInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BlossomBriefInfo_proto_goTypes,
		DependencyIndexes: file_BlossomBriefInfo_proto_depIdxs,
		MessageInfos:      file_BlossomBriefInfo_proto_msgTypes,
	}.Build()
	File_BlossomBriefInfo_proto = out.File
	file_BlossomBriefInfo_proto_rawDesc = nil
	file_BlossomBriefInfo_proto_goTypes = nil
	file_BlossomBriefInfo_proto_depIdxs = nil
}
