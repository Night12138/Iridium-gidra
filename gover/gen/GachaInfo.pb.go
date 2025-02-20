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
// source: GachaInfo.proto

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

type GachaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenCostItemId          uint32         `protobuf:"varint,2,opt,name=ten_cost_item_id,json=tenCostItemId,proto3" json:"ten_cost_item_id,omitempty"`
	EndTime                uint32         `protobuf:"varint,14,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	DisplayUp4ItemList     []uint32       `protobuf:"varint,1875,rep,packed,name=display_up4_item_list,json=displayUp4ItemList,proto3" json:"display_up4_item_list,omitempty"`
	Unk3100_JKILPCKLNPI    uint32         `protobuf:"varint,469,opt,name=Unk3100_JKILPCKLNPI,json=Unk3100JKILPCKLNPI,proto3" json:"Unk3100_JKILPCKLNPI,omitempty"`
	GachaUpInfoList        []*GachaUpInfo `protobuf:"bytes,1233,rep,name=gacha_up_info_list,json=gachaUpInfoList,proto3" json:"gacha_up_info_list,omitempty"`
	GachaProbUrl           string         `protobuf:"bytes,8,opt,name=gacha_prob_url,json=gachaProbUrl,proto3" json:"gacha_prob_url,omitempty"`
	GachaPrefabPath        string         `protobuf:"bytes,15,opt,name=gacha_prefab_path,json=gachaPrefabPath,proto3" json:"gacha_prefab_path,omitempty"`
	WishItemId             uint32         `protobuf:"varint,1637,opt,name=wish_item_id,json=wishItemId,proto3" json:"wish_item_id,omitempty"`
	BeginTime              uint32         `protobuf:"varint,1,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	WishMaxProgress        uint32         `protobuf:"varint,1222,opt,name=wish_max_progress,json=wishMaxProgress,proto3" json:"wish_max_progress,omitempty"`
	ScheduleId             uint32         `protobuf:"varint,10,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	GachaProbUrlOversea    string         `protobuf:"bytes,1481,opt,name=gacha_prob_url_oversea,json=gachaProbUrlOversea,proto3" json:"gacha_prob_url_oversea,omitempty"`
	GachaType              uint32         `protobuf:"varint,13,opt,name=gacha_type,json=gachaType,proto3" json:"gacha_type,omitempty"`
	LeftGachaTimes         uint32         `protobuf:"varint,5,opt,name=left_gacha_times,json=leftGachaTimes,proto3" json:"left_gacha_times,omitempty"`
	DisplayUp5ItemList     []uint32       `protobuf:"varint,2006,rep,packed,name=display_up5_item_list,json=displayUp5ItemList,proto3" json:"display_up5_item_list,omitempty"`
	GachaTimesLimit        uint32         `protobuf:"varint,11,opt,name=gacha_times_limit,json=gachaTimesLimit,proto3" json:"gacha_times_limit,omitempty"`
	CostItemNum            uint32         `protobuf:"varint,3,opt,name=cost_item_num,json=costItemNum,proto3" json:"cost_item_num,omitempty"`
	IsNewWish              bool           `protobuf:"varint,733,opt,name=is_new_wish,json=isNewWish,proto3" json:"is_new_wish,omitempty"`
	CostItemId             uint32         `protobuf:"varint,9,opt,name=cost_item_id,json=costItemId,proto3" json:"cost_item_id,omitempty"`
	TenCostItemNum         uint32         `protobuf:"varint,6,opt,name=ten_cost_item_num,json=tenCostItemNum,proto3" json:"ten_cost_item_num,omitempty"`
	GachaPreviewPrefabPath string         `protobuf:"bytes,4,opt,name=gacha_preview_prefab_path,json=gachaPreviewPrefabPath,proto3" json:"gacha_preview_prefab_path,omitempty"`
	WishProgress           uint32         `protobuf:"varint,1819,opt,name=wish_progress,json=wishProgress,proto3" json:"wish_progress,omitempty"`
	TitleTextmap           string         `protobuf:"bytes,736,opt,name=title_textmap,json=titleTextmap,proto3" json:"title_textmap,omitempty"`
	GachaRecordUrlOversea  string         `protobuf:"bytes,1854,opt,name=gacha_record_url_oversea,json=gachaRecordUrlOversea,proto3" json:"gacha_record_url_oversea,omitempty"`
	GachaSortId            uint32         `protobuf:"varint,7,opt,name=gacha_sort_id,json=gachaSortId,proto3" json:"gacha_sort_id,omitempty"`
	GachaRecordUrl         string         `protobuf:"bytes,12,opt,name=gacha_record_url,json=gachaRecordUrl,proto3" json:"gacha_record_url,omitempty"`
}

func (x *GachaInfo) Reset() {
	*x = GachaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GachaInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GachaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GachaInfo) ProtoMessage() {}

func (x *GachaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_GachaInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GachaInfo.ProtoReflect.Descriptor instead.
func (*GachaInfo) Descriptor() ([]byte, []int) {
	return file_GachaInfo_proto_rawDescGZIP(), []int{0}
}

func (x *GachaInfo) GetTenCostItemId() uint32 {
	if x != nil {
		return x.TenCostItemId
	}
	return 0
}

func (x *GachaInfo) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *GachaInfo) GetDisplayUp4ItemList() []uint32 {
	if x != nil {
		return x.DisplayUp4ItemList
	}
	return nil
}

func (x *GachaInfo) GetUnk3100_JKILPCKLNPI() uint32 {
	if x != nil {
		return x.Unk3100_JKILPCKLNPI
	}
	return 0
}

func (x *GachaInfo) GetGachaUpInfoList() []*GachaUpInfo {
	if x != nil {
		return x.GachaUpInfoList
	}
	return nil
}

func (x *GachaInfo) GetGachaProbUrl() string {
	if x != nil {
		return x.GachaProbUrl
	}
	return ""
}

func (x *GachaInfo) GetGachaPrefabPath() string {
	if x != nil {
		return x.GachaPrefabPath
	}
	return ""
}

func (x *GachaInfo) GetWishItemId() uint32 {
	if x != nil {
		return x.WishItemId
	}
	return 0
}

func (x *GachaInfo) GetBeginTime() uint32 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *GachaInfo) GetWishMaxProgress() uint32 {
	if x != nil {
		return x.WishMaxProgress
	}
	return 0
}

func (x *GachaInfo) GetScheduleId() uint32 {
	if x != nil {
		return x.ScheduleId
	}
	return 0
}

func (x *GachaInfo) GetGachaProbUrlOversea() string {
	if x != nil {
		return x.GachaProbUrlOversea
	}
	return ""
}

func (x *GachaInfo) GetGachaType() uint32 {
	if x != nil {
		return x.GachaType
	}
	return 0
}

func (x *GachaInfo) GetLeftGachaTimes() uint32 {
	if x != nil {
		return x.LeftGachaTimes
	}
	return 0
}

func (x *GachaInfo) GetDisplayUp5ItemList() []uint32 {
	if x != nil {
		return x.DisplayUp5ItemList
	}
	return nil
}

func (x *GachaInfo) GetGachaTimesLimit() uint32 {
	if x != nil {
		return x.GachaTimesLimit
	}
	return 0
}

func (x *GachaInfo) GetCostItemNum() uint32 {
	if x != nil {
		return x.CostItemNum
	}
	return 0
}

func (x *GachaInfo) GetIsNewWish() bool {
	if x != nil {
		return x.IsNewWish
	}
	return false
}

func (x *GachaInfo) GetCostItemId() uint32 {
	if x != nil {
		return x.CostItemId
	}
	return 0
}

func (x *GachaInfo) GetTenCostItemNum() uint32 {
	if x != nil {
		return x.TenCostItemNum
	}
	return 0
}

func (x *GachaInfo) GetGachaPreviewPrefabPath() string {
	if x != nil {
		return x.GachaPreviewPrefabPath
	}
	return ""
}

func (x *GachaInfo) GetWishProgress() uint32 {
	if x != nil {
		return x.WishProgress
	}
	return 0
}

func (x *GachaInfo) GetTitleTextmap() string {
	if x != nil {
		return x.TitleTextmap
	}
	return ""
}

func (x *GachaInfo) GetGachaRecordUrlOversea() string {
	if x != nil {
		return x.GachaRecordUrlOversea
	}
	return ""
}

func (x *GachaInfo) GetGachaSortId() uint32 {
	if x != nil {
		return x.GachaSortId
	}
	return 0
}

func (x *GachaInfo) GetGachaRecordUrl() string {
	if x != nil {
		return x.GachaRecordUrl
	}
	return ""
}

var File_GachaInfo_proto protoreflect.FileDescriptor

var file_GachaInfo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x47, 0x61, 0x63, 0x68, 0x61, 0x55, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x08, 0x0a, 0x09, 0x47, 0x61, 0x63, 0x68, 0x61, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x27, 0x0a, 0x10, 0x74, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x65,
	0x6e, 0x43, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x75, 0x70, 0x34, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0xd3, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x55,
	0x70, 0x34, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x55, 0x6e,
	0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x4a, 0x4b, 0x49, 0x4c, 0x50, 0x43, 0x4b, 0x4c, 0x4e, 0x50,
	0x49, 0x18, 0xd5, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x31, 0x30,
	0x30, 0x4a, 0x4b, 0x49, 0x4c, 0x50, 0x43, 0x4b, 0x4c, 0x4e, 0x50, 0x49, 0x12, 0x3a, 0x0a, 0x12,
	0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x75, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0xd1, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x61, 0x63, 0x68,
	0x61, 0x55, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x55, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x67, 0x61, 0x63, 0x68,
	0x61, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x67, 0x61, 0x63, 0x68, 0x61, 0x50, 0x72, 0x6f, 0x62, 0x55, 0x72, 0x6c, 0x12, 0x2a,
	0x0a, 0x11, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x61, 0x62, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x67, 0x61, 0x63, 0x68, 0x61,
	0x50, 0x72, 0x65, 0x66, 0x61, 0x62, 0x50, 0x61, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x69,
	0x73, 0x68, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0xe5, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x77, 0x69, 0x73, 0x68, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11,
	0x77, 0x69, 0x73, 0x68, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x18, 0xc6, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x77, 0x69, 0x73, 0x68, 0x4d, 0x61,
	0x78, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x67, 0x61,
	0x63, 0x68, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x6f, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x61, 0x18, 0xc9, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x67, 0x61, 0x63,
	0x68, 0x61, 0x50, 0x72, 0x6f, 0x62, 0x55, 0x72, 0x6c, 0x4f, 0x76, 0x65, 0x72, 0x73, 0x65, 0x61,
	0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x61, 0x63, 0x68, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x28, 0x0a, 0x10, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6c, 0x65, 0x66, 0x74, 0x47,
	0x61, 0x63, 0x68, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x70, 0x35, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0xd6, 0x0f, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x55, 0x70, 0x35, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x11, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6f, 0x73,
	0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x63, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4e, 0x75, 0x6d, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x5f, 0x77, 0x69, 0x73, 0x68, 0x18, 0xdd, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x57, 0x69, 0x73, 0x68, 0x12, 0x20,
	0x0a, 0x0c, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x12, 0x29, 0x0a, 0x11, 0x74, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74, 0x65, 0x6e,
	0x43, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4e, 0x75, 0x6d, 0x12, 0x39, 0x0a, 0x19, 0x67,
	0x61, 0x63, 0x68, 0x61, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x70, 0x72, 0x65,
	0x66, 0x61, 0x62, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16,
	0x67, 0x61, 0x63, 0x68, 0x61, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x50, 0x72, 0x65, 0x66,
	0x61, 0x62, 0x50, 0x61, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x9b, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x77, 0x69, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x6d, 0x61, 0x70, 0x18, 0xe0, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x54, 0x65, 0x78, 0x74, 0x6d,
	0x61, 0x70, 0x12, 0x38, 0x0a, 0x18, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x65, 0x61, 0x18, 0xbe,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x67, 0x61, 0x63, 0x68, 0x61, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x55, 0x72, 0x6c, 0x4f, 0x76, 0x65, 0x72, 0x73, 0x65, 0x61, 0x12, 0x22, 0x0a, 0x0d,
	0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x67, 0x61, 0x63, 0x68, 0x61, 0x53, 0x6f, 0x72, 0x74, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x10, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67, 0x61, 0x63, 0x68,
	0x61, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x72, 0x6c, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GachaInfo_proto_rawDescOnce sync.Once
	file_GachaInfo_proto_rawDescData = file_GachaInfo_proto_rawDesc
)

func file_GachaInfo_proto_rawDescGZIP() []byte {
	file_GachaInfo_proto_rawDescOnce.Do(func() {
		file_GachaInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_GachaInfo_proto_rawDescData)
	})
	return file_GachaInfo_proto_rawDescData
}

var file_GachaInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GachaInfo_proto_goTypes = []interface{}{
	(*GachaInfo)(nil),   // 0: GachaInfo
	(*GachaUpInfo)(nil), // 1: GachaUpInfo
}
var file_GachaInfo_proto_depIdxs = []int32{
	1, // 0: GachaInfo.gacha_up_info_list:type_name -> GachaUpInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GachaInfo_proto_init() }
func file_GachaInfo_proto_init() {
	if File_GachaInfo_proto != nil {
		return
	}
	file_GachaUpInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GachaInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GachaInfo); i {
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
			RawDescriptor: file_GachaInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GachaInfo_proto_goTypes,
		DependencyIndexes: file_GachaInfo_proto_depIdxs,
		MessageInfos:      file_GachaInfo_proto_msgTypes,
	}.Build()
	File_GachaInfo_proto = out.File
	file_GachaInfo_proto_rawDesc = nil
	file_GachaInfo_proto_goTypes = nil
	file_GachaInfo_proto_depIdxs = nil
}
