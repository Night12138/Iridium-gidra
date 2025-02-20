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
// source: GachaWishRsp.proto

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

// CmdId: 1521
// EnetChannelId: 0
// EnetIsReliable: true
type GachaWishRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GachaType       uint32 `protobuf:"varint,8,opt,name=gacha_type,json=gachaType,proto3" json:"gacha_type,omitempty"`
	GachaScheduleId uint32 `protobuf:"varint,7,opt,name=gacha_schedule_id,json=gachaScheduleId,proto3" json:"gacha_schedule_id,omitempty"`
	WishMaxProgress uint32 `protobuf:"varint,2,opt,name=wish_max_progress,json=wishMaxProgress,proto3" json:"wish_max_progress,omitempty"`
	WishProgress    uint32 `protobuf:"varint,5,opt,name=wish_progress,json=wishProgress,proto3" json:"wish_progress,omitempty"`
	WishItemId      uint32 `protobuf:"varint,3,opt,name=wish_item_id,json=wishItemId,proto3" json:"wish_item_id,omitempty"`
	Retcode         int32  `protobuf:"varint,14,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GachaWishRsp) Reset() {
	*x = GachaWishRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GachaWishRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GachaWishRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GachaWishRsp) ProtoMessage() {}

func (x *GachaWishRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GachaWishRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GachaWishRsp.ProtoReflect.Descriptor instead.
func (*GachaWishRsp) Descriptor() ([]byte, []int) {
	return file_GachaWishRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GachaWishRsp) GetGachaType() uint32 {
	if x != nil {
		return x.GachaType
	}
	return 0
}

func (x *GachaWishRsp) GetGachaScheduleId() uint32 {
	if x != nil {
		return x.GachaScheduleId
	}
	return 0
}

func (x *GachaWishRsp) GetWishMaxProgress() uint32 {
	if x != nil {
		return x.WishMaxProgress
	}
	return 0
}

func (x *GachaWishRsp) GetWishProgress() uint32 {
	if x != nil {
		return x.WishProgress
	}
	return 0
}

func (x *GachaWishRsp) GetWishItemId() uint32 {
	if x != nil {
		return x.WishItemId
	}
	return 0
}

func (x *GachaWishRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_GachaWishRsp_proto protoreflect.FileDescriptor

var file_GachaWishRsp_proto_rawDesc = []byte{
	0x0a, 0x12, 0x47, 0x61, 0x63, 0x68, 0x61, 0x57, 0x69, 0x73, 0x68, 0x52, 0x73, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x01, 0x0a, 0x0c, 0x47, 0x61, 0x63, 0x68, 0x61, 0x57, 0x69,
	0x73, 0x68, 0x52, 0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x67, 0x61, 0x63, 0x68, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0f, 0x67, 0x61, 0x63, 0x68, 0x61, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64,
	0x12, 0x2a, 0x0a, 0x11, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x77, 0x69, 0x73,
	0x68, 0x4d, 0x61, 0x78, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x77, 0x69, 0x73, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x77, 0x69, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x20, 0x0a, 0x0c, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x69, 0x73, 0x68, 0x49, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x06, 0x5a,
	0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GachaWishRsp_proto_rawDescOnce sync.Once
	file_GachaWishRsp_proto_rawDescData = file_GachaWishRsp_proto_rawDesc
)

func file_GachaWishRsp_proto_rawDescGZIP() []byte {
	file_GachaWishRsp_proto_rawDescOnce.Do(func() {
		file_GachaWishRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GachaWishRsp_proto_rawDescData)
	})
	return file_GachaWishRsp_proto_rawDescData
}

var file_GachaWishRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GachaWishRsp_proto_goTypes = []interface{}{
	(*GachaWishRsp)(nil), // 0: GachaWishRsp
}
var file_GachaWishRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GachaWishRsp_proto_init() }
func file_GachaWishRsp_proto_init() {
	if File_GachaWishRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GachaWishRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GachaWishRsp); i {
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
			RawDescriptor: file_GachaWishRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GachaWishRsp_proto_goTypes,
		DependencyIndexes: file_GachaWishRsp_proto_depIdxs,
		MessageInfos:      file_GachaWishRsp_proto_msgTypes,
	}.Build()
	File_GachaWishRsp_proto = out.File
	file_GachaWishRsp_proto_rawDesc = nil
	file_GachaWishRsp_proto_goTypes = nil
	file_GachaWishRsp_proto_depIdxs = nil
}
