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
// source: BattlePassBuySuccNotify.proto

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

// CmdId: 2614
// EnetChannelId: 0
// EnetIsReliable: true
type BattlePassBuySuccNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScheduleId      uint32       `protobuf:"varint,4,opt,name=schedule_id,json=scheduleId,proto3" json:"schedule_id,omitempty"`
	ProductPlayType uint32       `protobuf:"varint,11,opt,name=product_play_type,json=productPlayType,proto3" json:"product_play_type,omitempty"`
	AddPoint        uint32       `protobuf:"varint,12,opt,name=add_point,json=addPoint,proto3" json:"add_point,omitempty"`
	ItemList        []*ItemParam `protobuf:"bytes,9,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
}

func (x *BattlePassBuySuccNotify) Reset() {
	*x = BattlePassBuySuccNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BattlePassBuySuccNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattlePassBuySuccNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattlePassBuySuccNotify) ProtoMessage() {}

func (x *BattlePassBuySuccNotify) ProtoReflect() protoreflect.Message {
	mi := &file_BattlePassBuySuccNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattlePassBuySuccNotify.ProtoReflect.Descriptor instead.
func (*BattlePassBuySuccNotify) Descriptor() ([]byte, []int) {
	return file_BattlePassBuySuccNotify_proto_rawDescGZIP(), []int{0}
}

func (x *BattlePassBuySuccNotify) GetScheduleId() uint32 {
	if x != nil {
		return x.ScheduleId
	}
	return 0
}

func (x *BattlePassBuySuccNotify) GetProductPlayType() uint32 {
	if x != nil {
		return x.ProductPlayType
	}
	return 0
}

func (x *BattlePassBuySuccNotify) GetAddPoint() uint32 {
	if x != nil {
		return x.AddPoint
	}
	return 0
}

func (x *BattlePassBuySuccNotify) GetItemList() []*ItemParam {
	if x != nil {
		return x.ItemList
	}
	return nil
}

var File_BattlePassBuySuccNotify_proto protoreflect.FileDescriptor

var file_BattlePassBuySuccNotify_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x42, 0x75, 0x79, 0x53,
	0x75, 0x63, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0f, 0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xac, 0x01, 0x0a, 0x17, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x42,
	0x75, 0x79, 0x53, 0x75, 0x63, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x64, 0x64,
	0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x64,
	0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BattlePassBuySuccNotify_proto_rawDescOnce sync.Once
	file_BattlePassBuySuccNotify_proto_rawDescData = file_BattlePassBuySuccNotify_proto_rawDesc
)

func file_BattlePassBuySuccNotify_proto_rawDescGZIP() []byte {
	file_BattlePassBuySuccNotify_proto_rawDescOnce.Do(func() {
		file_BattlePassBuySuccNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_BattlePassBuySuccNotify_proto_rawDescData)
	})
	return file_BattlePassBuySuccNotify_proto_rawDescData
}

var file_BattlePassBuySuccNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BattlePassBuySuccNotify_proto_goTypes = []interface{}{
	(*BattlePassBuySuccNotify)(nil), // 0: BattlePassBuySuccNotify
	(*ItemParam)(nil),               // 1: ItemParam
}
var file_BattlePassBuySuccNotify_proto_depIdxs = []int32{
	1, // 0: BattlePassBuySuccNotify.item_list:type_name -> ItemParam
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BattlePassBuySuccNotify_proto_init() }
func file_BattlePassBuySuccNotify_proto_init() {
	if File_BattlePassBuySuccNotify_proto != nil {
		return
	}
	file_ItemParam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BattlePassBuySuccNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattlePassBuySuccNotify); i {
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
			RawDescriptor: file_BattlePassBuySuccNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BattlePassBuySuccNotify_proto_goTypes,
		DependencyIndexes: file_BattlePassBuySuccNotify_proto_depIdxs,
		MessageInfos:      file_BattlePassBuySuccNotify_proto_msgTypes,
	}.Build()
	File_BattlePassBuySuccNotify_proto = out.File
	file_BattlePassBuySuccNotify_proto_rawDesc = nil
	file_BattlePassBuySuccNotify_proto_goTypes = nil
	file_BattlePassBuySuccNotify_proto_depIdxs = nil
}
