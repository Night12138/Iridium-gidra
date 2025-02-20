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
// source: InBattleMechanicusInfo.proto

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

type InBattleMechanicusInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeftMonster           uint32                           `protobuf:"varint,5,opt,name=left_monster,json=leftMonster,proto3" json:"left_monster,omitempty"`
	WaitSeconds           uint32                           `protobuf:"varint,13,opt,name=wait_seconds,json=waitSeconds,proto3" json:"wait_seconds,omitempty"`
	EntranceList          []uint32                         `protobuf:"varint,410,rep,packed,name=entrance_list,json=entranceList,proto3" json:"entrance_list,omitempty"`
	ExitList              []uint32                         `protobuf:"varint,115,rep,packed,name=exit_list,json=exitList,proto3" json:"exit_list,omitempty"`
	HistoryCardList       []*InBattleMechanicusCardInfo    `protobuf:"bytes,11,rep,name=history_card_list,json=historyCardList,proto3" json:"history_card_list,omitempty"`
	MaxEscapeMonsterNum   uint32                           `protobuf:"varint,10,opt,name=max_escape_monster_num,json=maxEscapeMonsterNum,proto3" json:"max_escape_monster_num,omitempty"`
	BuildingStageDuration uint32                           `protobuf:"varint,4,opt,name=building_stage_duration,json=buildingStageDuration,proto3" json:"building_stage_duration,omitempty"`
	DurationMs            uint64                           `protobuf:"varint,8,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	Stage                 InBattleMechanicusStageType      `protobuf:"varint,9,opt,name=stage,proto3,enum=InBattleMechanicusStageType" json:"stage,omitempty"`
	TotalRound            uint32                           `protobuf:"varint,12,opt,name=total_round,json=totalRound,proto3" json:"total_round,omitempty"`
	MonsterList           []*InBattleMechanicusMonsterInfo `protobuf:"bytes,14,rep,name=monster_list,json=monsterList,proto3" json:"monster_list,omitempty"`
	EscapedMonsterNum     uint32                           `protobuf:"varint,6,opt,name=escaped_monster_num,json=escapedMonsterNum,proto3" json:"escaped_monster_num,omitempty"`
	Round                 uint32                           `protobuf:"varint,3,opt,name=round,proto3" json:"round,omitempty"`
	PickCardList          []*InBattleMechanicusCardInfo    `protobuf:"bytes,15,rep,name=pick_card_list,json=pickCardList,proto3" json:"pick_card_list,omitempty"`
	PlayerList            []*InBattleMechanicusPlayerInfo  `protobuf:"bytes,7,rep,name=player_list,json=playerList,proto3" json:"player_list,omitempty"`
	WaitBeginTimeUs       uint64                           `protobuf:"varint,1,opt,name=wait_begin_time_us,json=waitBeginTimeUs,proto3" json:"wait_begin_time_us,omitempty"`
	BeginTimeMs           uint64                           `protobuf:"varint,2,opt,name=begin_time_ms,json=beginTimeMs,proto3" json:"begin_time_ms,omitempty"`
}

func (x *InBattleMechanicusInfo) Reset() {
	*x = InBattleMechanicusInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_InBattleMechanicusInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InBattleMechanicusInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InBattleMechanicusInfo) ProtoMessage() {}

func (x *InBattleMechanicusInfo) ProtoReflect() protoreflect.Message {
	mi := &file_InBattleMechanicusInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InBattleMechanicusInfo.ProtoReflect.Descriptor instead.
func (*InBattleMechanicusInfo) Descriptor() ([]byte, []int) {
	return file_InBattleMechanicusInfo_proto_rawDescGZIP(), []int{0}
}

func (x *InBattleMechanicusInfo) GetLeftMonster() uint32 {
	if x != nil {
		return x.LeftMonster
	}
	return 0
}

func (x *InBattleMechanicusInfo) GetWaitSeconds() uint32 {
	if x != nil {
		return x.WaitSeconds
	}
	return 0
}

func (x *InBattleMechanicusInfo) GetEntranceList() []uint32 {
	if x != nil {
		return x.EntranceList
	}
	return nil
}

func (x *InBattleMechanicusInfo) GetExitList() []uint32 {
	if x != nil {
		return x.ExitList
	}
	return nil
}

func (x *InBattleMechanicusInfo) GetHistoryCardList() []*InBattleMechanicusCardInfo {
	if x != nil {
		return x.HistoryCardList
	}
	return nil
}

func (x *InBattleMechanicusInfo) GetMaxEscapeMonsterNum() uint32 {
	if x != nil {
		return x.MaxEscapeMonsterNum
	}
	return 0
}

func (x *InBattleMechanicusInfo) GetBuildingStageDuration() uint32 {
	if x != nil {
		return x.BuildingStageDuration
	}
	return 0
}

func (x *InBattleMechanicusInfo) GetDurationMs() uint64 {
	if x != nil {
		return x.DurationMs
	}
	return 0
}

func (x *InBattleMechanicusInfo) GetStage() InBattleMechanicusStageType {
	if x != nil {
		return x.Stage
	}
	return InBattleMechanicusStageType_IN_BATTLE_MECHANICUS_STAGE_TYPE_NONE
}

func (x *InBattleMechanicusInfo) GetTotalRound() uint32 {
	if x != nil {
		return x.TotalRound
	}
	return 0
}

func (x *InBattleMechanicusInfo) GetMonsterList() []*InBattleMechanicusMonsterInfo {
	if x != nil {
		return x.MonsterList
	}
	return nil
}

func (x *InBattleMechanicusInfo) GetEscapedMonsterNum() uint32 {
	if x != nil {
		return x.EscapedMonsterNum
	}
	return 0
}

func (x *InBattleMechanicusInfo) GetRound() uint32 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *InBattleMechanicusInfo) GetPickCardList() []*InBattleMechanicusCardInfo {
	if x != nil {
		return x.PickCardList
	}
	return nil
}

func (x *InBattleMechanicusInfo) GetPlayerList() []*InBattleMechanicusPlayerInfo {
	if x != nil {
		return x.PlayerList
	}
	return nil
}

func (x *InBattleMechanicusInfo) GetWaitBeginTimeUs() uint64 {
	if x != nil {
		return x.WaitBeginTimeUs
	}
	return 0
}

func (x *InBattleMechanicusInfo) GetBeginTimeMs() uint64 {
	if x != nil {
		return x.BeginTimeMs
	}
	return 0
}

var File_InBattleMechanicusInfo_proto protoreflect.FileDescriptor

var file_InBattleMechanicusInfo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e,
	0x69, 0x63, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63,
	0x75, 0x73, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e,
	0x69, 0x63, 0x75, 0x73, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4d,
	0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x49, 0x6e, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x06, 0x0a,
	0x16, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69,
	0x63, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x65, 0x66, 0x74, 0x5f,
	0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6c,
	0x65, 0x66, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x61,
	0x69, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x77, 0x61, 0x69, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x24, 0x0a,
	0x0d, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x9a,
	0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x73, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x47, 0x0a, 0x11, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x63, 0x61, 0x72, 0x64,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x49, 0x6e,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73,
	0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x16, 0x6d, 0x61, 0x78,
	0x5f, 0x65, 0x73, 0x63, 0x61, 0x70, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x6d, 0x61, 0x78, 0x45, 0x73,
	0x63, 0x61, 0x70, 0x65, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x36,
	0x0a, 0x17, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x15, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x67, 0x65, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x41, 0x0a, 0x0c,
	0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x63,
	0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0b, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x13, 0x65, 0x73, 0x63, 0x61, 0x70, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x65, 0x73,
	0x63, 0x61, 0x70, 0x65, 0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x41, 0x0a, 0x0e, 0x70, 0x69, 0x63, 0x6b, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63,
	0x75, 0x73, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x70, 0x69, 0x63, 0x6b,
	0x43, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x49, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63,
	0x75, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x12, 0x77, 0x61, 0x69, 0x74,
	0x5f, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x77, 0x61, 0x69, 0x74, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x55, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_InBattleMechanicusInfo_proto_rawDescOnce sync.Once
	file_InBattleMechanicusInfo_proto_rawDescData = file_InBattleMechanicusInfo_proto_rawDesc
)

func file_InBattleMechanicusInfo_proto_rawDescGZIP() []byte {
	file_InBattleMechanicusInfo_proto_rawDescOnce.Do(func() {
		file_InBattleMechanicusInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_InBattleMechanicusInfo_proto_rawDescData)
	})
	return file_InBattleMechanicusInfo_proto_rawDescData
}

var file_InBattleMechanicusInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_InBattleMechanicusInfo_proto_goTypes = []interface{}{
	(*InBattleMechanicusInfo)(nil),        // 0: InBattleMechanicusInfo
	(*InBattleMechanicusCardInfo)(nil),    // 1: InBattleMechanicusCardInfo
	(InBattleMechanicusStageType)(0),      // 2: InBattleMechanicusStageType
	(*InBattleMechanicusMonsterInfo)(nil), // 3: InBattleMechanicusMonsterInfo
	(*InBattleMechanicusPlayerInfo)(nil),  // 4: InBattleMechanicusPlayerInfo
}
var file_InBattleMechanicusInfo_proto_depIdxs = []int32{
	1, // 0: InBattleMechanicusInfo.history_card_list:type_name -> InBattleMechanicusCardInfo
	2, // 1: InBattleMechanicusInfo.stage:type_name -> InBattleMechanicusStageType
	3, // 2: InBattleMechanicusInfo.monster_list:type_name -> InBattleMechanicusMonsterInfo
	1, // 3: InBattleMechanicusInfo.pick_card_list:type_name -> InBattleMechanicusCardInfo
	4, // 4: InBattleMechanicusInfo.player_list:type_name -> InBattleMechanicusPlayerInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_InBattleMechanicusInfo_proto_init() }
func file_InBattleMechanicusInfo_proto_init() {
	if File_InBattleMechanicusInfo_proto != nil {
		return
	}
	file_InBattleMechanicusCardInfo_proto_init()
	file_InBattleMechanicusMonsterInfo_proto_init()
	file_InBattleMechanicusPlayerInfo_proto_init()
	file_InBattleMechanicusStageType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_InBattleMechanicusInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InBattleMechanicusInfo); i {
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
			RawDescriptor: file_InBattleMechanicusInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_InBattleMechanicusInfo_proto_goTypes,
		DependencyIndexes: file_InBattleMechanicusInfo_proto_depIdxs,
		MessageInfos:      file_InBattleMechanicusInfo_proto_msgTypes,
	}.Build()
	File_InBattleMechanicusInfo_proto = out.File
	file_InBattleMechanicusInfo_proto_rawDesc = nil
	file_InBattleMechanicusInfo_proto_goTypes = nil
	file_InBattleMechanicusInfo_proto_depIdxs = nil
}
