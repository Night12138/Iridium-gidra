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
// source: TowerMonthlyCombatRecord.proto

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

type TowerMonthlyCombatRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MostKillAvatarPair            *TowerFightRecordPair   `protobuf:"bytes,14,opt,name=most_kill_avatar_pair,json=mostKillAvatarPair,proto3" json:"most_kill_avatar_pair,omitempty"`
	MostCastNormalSkillAvatarPair *TowerFightRecordPair   `protobuf:"bytes,8,opt,name=most_cast_normal_skill_avatar_pair,json=mostCastNormalSkillAvatarPair,proto3" json:"most_cast_normal_skill_avatar_pair,omitempty"`
	MostRevealAvatarList          []*TowerFightRecordPair `protobuf:"bytes,6,rep,name=most_reveal_avatar_list,json=mostRevealAvatarList,proto3" json:"most_reveal_avatar_list,omitempty"`
	MostCastEnergySkillAvatarPair *TowerFightRecordPair   `protobuf:"bytes,4,opt,name=most_cast_energy_skill_avatar_pair,json=mostCastEnergySkillAvatarPair,proto3" json:"most_cast_energy_skill_avatar_pair,omitempty"`
	HighestDpsAvatrPair           *TowerFightRecordPair   `protobuf:"bytes,12,opt,name=highest_dps_avatr_pair,json=highestDpsAvatrPair,proto3" json:"highest_dps_avatr_pair,omitempty"`
	MostTakeDamageAvatarPair      *TowerFightRecordPair   `protobuf:"bytes,9,opt,name=most_take_damage_avatar_pair,json=mostTakeDamageAvatarPair,proto3" json:"most_take_damage_avatar_pair,omitempty"`
}

func (x *TowerMonthlyCombatRecord) Reset() {
	*x = TowerMonthlyCombatRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TowerMonthlyCombatRecord_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TowerMonthlyCombatRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TowerMonthlyCombatRecord) ProtoMessage() {}

func (x *TowerMonthlyCombatRecord) ProtoReflect() protoreflect.Message {
	mi := &file_TowerMonthlyCombatRecord_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TowerMonthlyCombatRecord.ProtoReflect.Descriptor instead.
func (*TowerMonthlyCombatRecord) Descriptor() ([]byte, []int) {
	return file_TowerMonthlyCombatRecord_proto_rawDescGZIP(), []int{0}
}

func (x *TowerMonthlyCombatRecord) GetMostKillAvatarPair() *TowerFightRecordPair {
	if x != nil {
		return x.MostKillAvatarPair
	}
	return nil
}

func (x *TowerMonthlyCombatRecord) GetMostCastNormalSkillAvatarPair() *TowerFightRecordPair {
	if x != nil {
		return x.MostCastNormalSkillAvatarPair
	}
	return nil
}

func (x *TowerMonthlyCombatRecord) GetMostRevealAvatarList() []*TowerFightRecordPair {
	if x != nil {
		return x.MostRevealAvatarList
	}
	return nil
}

func (x *TowerMonthlyCombatRecord) GetMostCastEnergySkillAvatarPair() *TowerFightRecordPair {
	if x != nil {
		return x.MostCastEnergySkillAvatarPair
	}
	return nil
}

func (x *TowerMonthlyCombatRecord) GetHighestDpsAvatrPair() *TowerFightRecordPair {
	if x != nil {
		return x.HighestDpsAvatrPair
	}
	return nil
}

func (x *TowerMonthlyCombatRecord) GetMostTakeDamageAvatarPair() *TowerFightRecordPair {
	if x != nil {
		return x.MostTakeDamageAvatarPair
	}
	return nil
}

var File_TowerMonthlyCombatRecord_proto protoreflect.FileDescriptor

var file_TowerMonthlyCombatRecord_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x43, 0x6f,
	0x6d, 0x62, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x46, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x50, 0x61, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x04, 0x0a,
	0x18, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x43, 0x6f, 0x6d,
	0x62, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x48, 0x0a, 0x15, 0x6d, 0x6f, 0x73,
	0x74, 0x5f, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x70, 0x61,
	0x69, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x54, 0x6f, 0x77, 0x65, 0x72,
	0x46, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x50, 0x61, 0x69, 0x72, 0x52,
	0x12, 0x6d, 0x6f, 0x73, 0x74, 0x4b, 0x69, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50,
	0x61, 0x69, 0x72, 0x12, 0x60, 0x0a, 0x22, 0x6d, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x46, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x50, 0x61, 0x69, 0x72, 0x52, 0x1d, 0x6d, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x73, 0x74,
	0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x50, 0x61, 0x69, 0x72, 0x12, 0x4c, 0x0a, 0x17, 0x6d, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65,
	0x76, 0x65, 0x61, 0x6c, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x46, 0x69,
	0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x50, 0x61, 0x69, 0x72, 0x52, 0x14, 0x6d,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x76, 0x65, 0x61, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x60, 0x0a, 0x22, 0x6d, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x73, 0x74,
	0x5f, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x46, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x50, 0x61, 0x69, 0x72, 0x52, 0x1d, 0x6d, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x73, 0x74,
	0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x50, 0x61, 0x69, 0x72, 0x12, 0x4a, 0x0a, 0x16, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74,
	0x5f, 0x64, 0x70, 0x73, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x72, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x46, 0x69, 0x67,
	0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x50, 0x61, 0x69, 0x72, 0x52, 0x13, 0x68, 0x69,
	0x67, 0x68, 0x65, 0x73, 0x74, 0x44, 0x70, 0x73, 0x41, 0x76, 0x61, 0x74, 0x72, 0x50, 0x61, 0x69,
	0x72, 0x12, 0x55, 0x0a, 0x1c, 0x6d, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x64,
	0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x70, 0x61, 0x69,
	0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x54, 0x6f, 0x77, 0x65, 0x72, 0x46,
	0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x50, 0x61, 0x69, 0x72, 0x52, 0x18,
	0x6d, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x6b, 0x65, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x50, 0x61, 0x69, 0x72, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TowerMonthlyCombatRecord_proto_rawDescOnce sync.Once
	file_TowerMonthlyCombatRecord_proto_rawDescData = file_TowerMonthlyCombatRecord_proto_rawDesc
)

func file_TowerMonthlyCombatRecord_proto_rawDescGZIP() []byte {
	file_TowerMonthlyCombatRecord_proto_rawDescOnce.Do(func() {
		file_TowerMonthlyCombatRecord_proto_rawDescData = protoimpl.X.CompressGZIP(file_TowerMonthlyCombatRecord_proto_rawDescData)
	})
	return file_TowerMonthlyCombatRecord_proto_rawDescData
}

var file_TowerMonthlyCombatRecord_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TowerMonthlyCombatRecord_proto_goTypes = []interface{}{
	(*TowerMonthlyCombatRecord)(nil), // 0: TowerMonthlyCombatRecord
	(*TowerFightRecordPair)(nil),     // 1: TowerFightRecordPair
}
var file_TowerMonthlyCombatRecord_proto_depIdxs = []int32{
	1, // 0: TowerMonthlyCombatRecord.most_kill_avatar_pair:type_name -> TowerFightRecordPair
	1, // 1: TowerMonthlyCombatRecord.most_cast_normal_skill_avatar_pair:type_name -> TowerFightRecordPair
	1, // 2: TowerMonthlyCombatRecord.most_reveal_avatar_list:type_name -> TowerFightRecordPair
	1, // 3: TowerMonthlyCombatRecord.most_cast_energy_skill_avatar_pair:type_name -> TowerFightRecordPair
	1, // 4: TowerMonthlyCombatRecord.highest_dps_avatr_pair:type_name -> TowerFightRecordPair
	1, // 5: TowerMonthlyCombatRecord.most_take_damage_avatar_pair:type_name -> TowerFightRecordPair
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_TowerMonthlyCombatRecord_proto_init() }
func file_TowerMonthlyCombatRecord_proto_init() {
	if File_TowerMonthlyCombatRecord_proto != nil {
		return
	}
	file_TowerFightRecordPair_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TowerMonthlyCombatRecord_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TowerMonthlyCombatRecord); i {
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
			RawDescriptor: file_TowerMonthlyCombatRecord_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TowerMonthlyCombatRecord_proto_goTypes,
		DependencyIndexes: file_TowerMonthlyCombatRecord_proto_depIdxs,
		MessageInfos:      file_TowerMonthlyCombatRecord_proto_msgTypes,
	}.Build()
	File_TowerMonthlyCombatRecord_proto = out.File
	file_TowerMonthlyCombatRecord_proto_rawDesc = nil
	file_TowerMonthlyCombatRecord_proto_goTypes = nil
	file_TowerMonthlyCombatRecord_proto_depIdxs = nil
}
