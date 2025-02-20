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
// source: MultistagePlayInfo.proto

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

type MultistagePlayInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayIndex  uint32 `protobuf:"varint,13,opt,name=play_index,json=playIndex,proto3" json:"play_index,omitempty"`
	PlayType   uint32 `protobuf:"varint,11,opt,name=play_type,json=playType,proto3" json:"play_type,omitempty"`
	StageType  uint32 `protobuf:"varint,10,opt,name=stage_type,json=stageType,proto3" json:"stage_type,omitempty"`
	Duration   uint32 `protobuf:"varint,8,opt,name=duration,proto3" json:"duration,omitempty"`
	GroupId    uint32 `protobuf:"varint,12,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	BeginTime  uint32 `protobuf:"varint,9,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	StageIndex uint32 `protobuf:"varint,1,opt,name=stage_index,json=stageIndex,proto3" json:"stage_index,omitempty"`
	// Types that are assignable to Detail:
	//
	//	*MultistagePlayInfo_MechanicusInfo
	//	*MultistagePlayInfo_FleurFairInfo
	//	*MultistagePlayInfo_HideAndSeekInfo
	//	*MultistagePlayInfo_ChessInfo
	//	*MultistagePlayInfo_IrodoriChessInfo
	Detail isMultistagePlayInfo_Detail `protobuf_oneof:"detail"`
}

func (x *MultistagePlayInfo) Reset() {
	*x = MultistagePlayInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MultistagePlayInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultistagePlayInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultistagePlayInfo) ProtoMessage() {}

func (x *MultistagePlayInfo) ProtoReflect() protoreflect.Message {
	mi := &file_MultistagePlayInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultistagePlayInfo.ProtoReflect.Descriptor instead.
func (*MultistagePlayInfo) Descriptor() ([]byte, []int) {
	return file_MultistagePlayInfo_proto_rawDescGZIP(), []int{0}
}

func (x *MultistagePlayInfo) GetPlayIndex() uint32 {
	if x != nil {
		return x.PlayIndex
	}
	return 0
}

func (x *MultistagePlayInfo) GetPlayType() uint32 {
	if x != nil {
		return x.PlayType
	}
	return 0
}

func (x *MultistagePlayInfo) GetStageType() uint32 {
	if x != nil {
		return x.StageType
	}
	return 0
}

func (x *MultistagePlayInfo) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *MultistagePlayInfo) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *MultistagePlayInfo) GetBeginTime() uint32 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *MultistagePlayInfo) GetStageIndex() uint32 {
	if x != nil {
		return x.StageIndex
	}
	return 0
}

func (m *MultistagePlayInfo) GetDetail() isMultistagePlayInfo_Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (x *MultistagePlayInfo) GetMechanicusInfo() *InBattleMechanicusInfo {
	if x, ok := x.GetDetail().(*MultistagePlayInfo_MechanicusInfo); ok {
		return x.MechanicusInfo
	}
	return nil
}

func (x *MultistagePlayInfo) GetFleurFairInfo() *InBattleFleurFairInfo {
	if x, ok := x.GetDetail().(*MultistagePlayInfo_FleurFairInfo); ok {
		return x.FleurFairInfo
	}
	return nil
}

func (x *MultistagePlayInfo) GetHideAndSeekInfo() *HideAndSeekStageInfo {
	if x, ok := x.GetDetail().(*MultistagePlayInfo_HideAndSeekInfo); ok {
		return x.HideAndSeekInfo
	}
	return nil
}

func (x *MultistagePlayInfo) GetChessInfo() *InBattleChessInfo {
	if x, ok := x.GetDetail().(*MultistagePlayInfo_ChessInfo); ok {
		return x.ChessInfo
	}
	return nil
}

func (x *MultistagePlayInfo) GetIrodoriChessInfo() *IrodoriChessInfo {
	if x, ok := x.GetDetail().(*MultistagePlayInfo_IrodoriChessInfo); ok {
		return x.IrodoriChessInfo
	}
	return nil
}

type isMultistagePlayInfo_Detail interface {
	isMultistagePlayInfo_Detail()
}

type MultistagePlayInfo_MechanicusInfo struct {
	MechanicusInfo *InBattleMechanicusInfo `protobuf:"bytes,1334,opt,name=mechanicus_info,json=mechanicusInfo,proto3,oneof"`
}

type MultistagePlayInfo_FleurFairInfo struct {
	FleurFairInfo *InBattleFleurFairInfo `protobuf:"bytes,1064,opt,name=fleur_fair_info,json=fleurFairInfo,proto3,oneof"`
}

type MultistagePlayInfo_HideAndSeekInfo struct {
	HideAndSeekInfo *HideAndSeekStageInfo `protobuf:"bytes,108,opt,name=hide_and_seek_info,json=hideAndSeekInfo,proto3,oneof"`
}

type MultistagePlayInfo_ChessInfo struct {
	ChessInfo *InBattleChessInfo `protobuf:"bytes,1758,opt,name=chess_info,json=chessInfo,proto3,oneof"`
}

type MultistagePlayInfo_IrodoriChessInfo struct {
	IrodoriChessInfo *IrodoriChessInfo `protobuf:"bytes,531,opt,name=irodori_chess_info,json=irodoriChessInfo,proto3,oneof"`
}

func (*MultistagePlayInfo_MechanicusInfo) isMultistagePlayInfo_Detail() {}

func (*MultistagePlayInfo_FleurFairInfo) isMultistagePlayInfo_Detail() {}

func (*MultistagePlayInfo_HideAndSeekInfo) isMultistagePlayInfo_Detail() {}

func (*MultistagePlayInfo_ChessInfo) isMultistagePlayInfo_Detail() {}

func (*MultistagePlayInfo_IrodoriChessInfo) isMultistagePlayInfo_Detail() {}

var File_MultistagePlayInfo_proto protoreflect.FileDescriptor

var file_MultistagePlayInfo_proto_rawDesc = []byte{
	0x0a, 0x18, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x74, 0x61, 0x67, 0x65, 0x50, 0x6c, 0x61, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x48, 0x69, 0x64, 0x65,
	0x41, 0x6e, 0x64, 0x53, 0x65, 0x65, 0x6b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x43, 0x68, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46, 0x61,
	0x69, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x49, 0x6e,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x49, 0x72, 0x6f, 0x64,
	0x6f, 0x72, 0x69, 0x43, 0x68, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb8, 0x04, 0x0a, 0x12, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70,
	0x6c, 0x61, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6c, 0x61,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x43, 0x0a, 0x0f, 0x6d,
	0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0xb6,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65,
	0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00,
	0x52, 0x0e, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x41, 0x0a, 0x0f, 0x66, 0x6c, 0x65, 0x75, 0x72, 0x5f, 0x66, 0x61, 0x69, 0x72, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0xa8, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x49, 0x6e, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x46, 0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x48, 0x00, 0x52, 0x0d, 0x66, 0x6c, 0x65, 0x75, 0x72, 0x46, 0x61, 0x69, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x44, 0x0a, 0x12, 0x68, 0x69, 0x64, 0x65, 0x5f, 0x61, 0x6e, 0x64, 0x5f,
	0x73, 0x65, 0x65, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x48, 0x69, 0x64, 0x65, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x65, 0x6b, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0f, 0x68, 0x69, 0x64, 0x65, 0x41, 0x6e,
	0x64, 0x53, 0x65, 0x65, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x0a, 0x63, 0x68, 0x65,
	0x73, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0xde, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x48, 0x00, 0x52, 0x09, 0x63, 0x68, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x42, 0x0a, 0x12, 0x69, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x5f, 0x63, 0x68, 0x65, 0x73, 0x73,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x93, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x49,
	0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x43, 0x68, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x00, 0x52, 0x10, 0x69, 0x72, 0x6f, 0x64, 0x6f, 0x72, 0x69, 0x43, 0x68, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x08, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x06, 0x5a,
	0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MultistagePlayInfo_proto_rawDescOnce sync.Once
	file_MultistagePlayInfo_proto_rawDescData = file_MultistagePlayInfo_proto_rawDesc
)

func file_MultistagePlayInfo_proto_rawDescGZIP() []byte {
	file_MultistagePlayInfo_proto_rawDescOnce.Do(func() {
		file_MultistagePlayInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_MultistagePlayInfo_proto_rawDescData)
	})
	return file_MultistagePlayInfo_proto_rawDescData
}

var file_MultistagePlayInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MultistagePlayInfo_proto_goTypes = []interface{}{
	(*MultistagePlayInfo)(nil),     // 0: MultistagePlayInfo
	(*InBattleMechanicusInfo)(nil), // 1: InBattleMechanicusInfo
	(*InBattleFleurFairInfo)(nil),  // 2: InBattleFleurFairInfo
	(*HideAndSeekStageInfo)(nil),   // 3: HideAndSeekStageInfo
	(*InBattleChessInfo)(nil),      // 4: InBattleChessInfo
	(*IrodoriChessInfo)(nil),       // 5: IrodoriChessInfo
}
var file_MultistagePlayInfo_proto_depIdxs = []int32{
	1, // 0: MultistagePlayInfo.mechanicus_info:type_name -> InBattleMechanicusInfo
	2, // 1: MultistagePlayInfo.fleur_fair_info:type_name -> InBattleFleurFairInfo
	3, // 2: MultistagePlayInfo.hide_and_seek_info:type_name -> HideAndSeekStageInfo
	4, // 3: MultistagePlayInfo.chess_info:type_name -> InBattleChessInfo
	5, // 4: MultistagePlayInfo.irodori_chess_info:type_name -> IrodoriChessInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_MultistagePlayInfo_proto_init() }
func file_MultistagePlayInfo_proto_init() {
	if File_MultistagePlayInfo_proto != nil {
		return
	}
	file_HideAndSeekStageInfo_proto_init()
	file_InBattleChessInfo_proto_init()
	file_InBattleFleurFairInfo_proto_init()
	file_InBattleMechanicusInfo_proto_init()
	file_IrodoriChessInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MultistagePlayInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultistagePlayInfo); i {
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
	file_MultistagePlayInfo_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MultistagePlayInfo_MechanicusInfo)(nil),
		(*MultistagePlayInfo_FleurFairInfo)(nil),
		(*MultistagePlayInfo_HideAndSeekInfo)(nil),
		(*MultistagePlayInfo_ChessInfo)(nil),
		(*MultistagePlayInfo_IrodoriChessInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MultistagePlayInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MultistagePlayInfo_proto_goTypes,
		DependencyIndexes: file_MultistagePlayInfo_proto_depIdxs,
		MessageInfos:      file_MultistagePlayInfo_proto_msgTypes,
	}.Build()
	File_MultistagePlayInfo_proto = out.File
	file_MultistagePlayInfo_proto_rawDesc = nil
	file_MultistagePlayInfo_proto_goTypes = nil
	file_MultistagePlayInfo_proto_depIdxs = nil
}
