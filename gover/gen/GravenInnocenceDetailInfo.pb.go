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
// source: GravenInnocenceDetailInfo.proto

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

type GravenInnocenceDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsContentClosed     bool                 `protobuf:"varint,8,opt,name=is_content_closed,json=isContentClosed,proto3" json:"is_content_closed,omitempty"`
	Unk3000_JGJKABIPGLK *Unk3000_OFMFFECMKLE `protobuf:"bytes,10,opt,name=Unk3000_JGJKABIPGLK,json=Unk3000JGJKABIPGLK,proto3" json:"Unk3000_JGJKABIPGLK,omitempty"`
	Unk3000_CDDIFHNEDOO *Unk3000_ILLNKBDNGKP `protobuf:"bytes,7,opt,name=Unk3000_CDDIFHNEDOO,json=Unk3000CDDIFHNEDOO,proto3" json:"Unk3000_CDDIFHNEDOO,omitempty"`
	Unk3000_BDFIOPBIOEB *Unk3000_ALPEACOMIPG `protobuf:"bytes,13,opt,name=Unk3000_BDFIOPBIOEB,json=Unk3000BDFIOPBIOEB,proto3" json:"Unk3000_BDFIOPBIOEB,omitempty"`
	Unk3000_KDPJGGENAJM *Unk3000_FFOBEKMOHOI `protobuf:"bytes,12,opt,name=Unk3000_KDPJGGENAJM,json=Unk3000KDPJGGENAJM,proto3" json:"Unk3000_KDPJGGENAJM,omitempty"`
}

func (x *GravenInnocenceDetailInfo) Reset() {
	*x = GravenInnocenceDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GravenInnocenceDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GravenInnocenceDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GravenInnocenceDetailInfo) ProtoMessage() {}

func (x *GravenInnocenceDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_GravenInnocenceDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GravenInnocenceDetailInfo.ProtoReflect.Descriptor instead.
func (*GravenInnocenceDetailInfo) Descriptor() ([]byte, []int) {
	return file_GravenInnocenceDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *GravenInnocenceDetailInfo) GetIsContentClosed() bool {
	if x != nil {
		return x.IsContentClosed
	}
	return false
}

func (x *GravenInnocenceDetailInfo) GetUnk3000_JGJKABIPGLK() *Unk3000_OFMFFECMKLE {
	if x != nil {
		return x.Unk3000_JGJKABIPGLK
	}
	return nil
}

func (x *GravenInnocenceDetailInfo) GetUnk3000_CDDIFHNEDOO() *Unk3000_ILLNKBDNGKP {
	if x != nil {
		return x.Unk3000_CDDIFHNEDOO
	}
	return nil
}

func (x *GravenInnocenceDetailInfo) GetUnk3000_BDFIOPBIOEB() *Unk3000_ALPEACOMIPG {
	if x != nil {
		return x.Unk3000_BDFIOPBIOEB
	}
	return nil
}

func (x *GravenInnocenceDetailInfo) GetUnk3000_KDPJGGENAJM() *Unk3000_FFOBEKMOHOI {
	if x != nil {
		return x.Unk3000_KDPJGGENAJM
	}
	return nil
}

var File_GravenInnocenceDetailInfo_proto protoreflect.FileDescriptor

var file_GravenInnocenceDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x47, 0x72, 0x61, 0x76, 0x65, 0x6e, 0x49, 0x6e, 0x6e, 0x6f, 0x63, 0x65, 0x6e, 0x63,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x41, 0x4c, 0x50, 0x45, 0x41,
	0x43, 0x4f, 0x4d, 0x49, 0x50, 0x47, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e,
	0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x46, 0x46, 0x4f, 0x42, 0x45, 0x4b, 0x4d, 0x4f, 0x48, 0x4f,
	0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30,
	0x5f, 0x49, 0x4c, 0x4c, 0x4e, 0x4b, 0x42, 0x44, 0x4e, 0x47, 0x4b, 0x50, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x4f, 0x46, 0x4d, 0x46,
	0x46, 0x45, 0x43, 0x4d, 0x4b, 0x4c, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x02,
	0x0a, 0x19, 0x47, 0x72, 0x61, 0x76, 0x65, 0x6e, 0x49, 0x6e, 0x6e, 0x6f, 0x63, 0x65, 0x6e, 0x63,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x11, 0x69,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30,
	0x30, 0x30, 0x5f, 0x4a, 0x47, 0x4a, 0x4b, 0x41, 0x42, 0x49, 0x50, 0x47, 0x4c, 0x4b, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x4f,
	0x46, 0x4d, 0x46, 0x46, 0x45, 0x43, 0x4d, 0x4b, 0x4c, 0x45, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33,
	0x30, 0x30, 0x30, 0x4a, 0x47, 0x4a, 0x4b, 0x41, 0x42, 0x49, 0x50, 0x47, 0x4c, 0x4b, 0x12, 0x45,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x43, 0x44, 0x44, 0x49, 0x46, 0x48,
	0x4e, 0x45, 0x44, 0x4f, 0x4f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e,
	0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x49, 0x4c, 0x4c, 0x4e, 0x4b, 0x42, 0x44, 0x4e, 0x47, 0x4b,
	0x50, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x43, 0x44, 0x44, 0x49, 0x46, 0x48,
	0x4e, 0x45, 0x44, 0x4f, 0x4f, 0x12, 0x45, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30,
	0x5f, 0x42, 0x44, 0x46, 0x49, 0x4f, 0x50, 0x42, 0x49, 0x4f, 0x45, 0x42, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x41, 0x4c, 0x50,
	0x45, 0x41, 0x43, 0x4f, 0x4d, 0x49, 0x50, 0x47, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30,
	0x30, 0x42, 0x44, 0x46, 0x49, 0x4f, 0x50, 0x42, 0x49, 0x4f, 0x45, 0x42, 0x12, 0x45, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x4b, 0x44, 0x50, 0x4a, 0x47, 0x47, 0x45, 0x4e,
	0x41, 0x4a, 0x4d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x33,
	0x30, 0x30, 0x30, 0x5f, 0x46, 0x46, 0x4f, 0x42, 0x45, 0x4b, 0x4d, 0x4f, 0x48, 0x4f, 0x49, 0x52,
	0x12, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x4b, 0x44, 0x50, 0x4a, 0x47, 0x47, 0x45, 0x4e,
	0x41, 0x4a, 0x4d, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_GravenInnocenceDetailInfo_proto_rawDescOnce sync.Once
	file_GravenInnocenceDetailInfo_proto_rawDescData = file_GravenInnocenceDetailInfo_proto_rawDesc
)

func file_GravenInnocenceDetailInfo_proto_rawDescGZIP() []byte {
	file_GravenInnocenceDetailInfo_proto_rawDescOnce.Do(func() {
		file_GravenInnocenceDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_GravenInnocenceDetailInfo_proto_rawDescData)
	})
	return file_GravenInnocenceDetailInfo_proto_rawDescData
}

var file_GravenInnocenceDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GravenInnocenceDetailInfo_proto_goTypes = []interface{}{
	(*GravenInnocenceDetailInfo)(nil), // 0: GravenInnocenceDetailInfo
	(*Unk3000_OFMFFECMKLE)(nil),       // 1: Unk3000_OFMFFECMKLE
	(*Unk3000_ILLNKBDNGKP)(nil),       // 2: Unk3000_ILLNKBDNGKP
	(*Unk3000_ALPEACOMIPG)(nil),       // 3: Unk3000_ALPEACOMIPG
	(*Unk3000_FFOBEKMOHOI)(nil),       // 4: Unk3000_FFOBEKMOHOI
}
var file_GravenInnocenceDetailInfo_proto_depIdxs = []int32{
	1, // 0: GravenInnocenceDetailInfo.Unk3000_JGJKABIPGLK:type_name -> Unk3000_OFMFFECMKLE
	2, // 1: GravenInnocenceDetailInfo.Unk3000_CDDIFHNEDOO:type_name -> Unk3000_ILLNKBDNGKP
	3, // 2: GravenInnocenceDetailInfo.Unk3000_BDFIOPBIOEB:type_name -> Unk3000_ALPEACOMIPG
	4, // 3: GravenInnocenceDetailInfo.Unk3000_KDPJGGENAJM:type_name -> Unk3000_FFOBEKMOHOI
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_GravenInnocenceDetailInfo_proto_init() }
func file_GravenInnocenceDetailInfo_proto_init() {
	if File_GravenInnocenceDetailInfo_proto != nil {
		return
	}
	file_Unk3000_ALPEACOMIPG_proto_init()
	file_Unk3000_FFOBEKMOHOI_proto_init()
	file_Unk3000_ILLNKBDNGKP_proto_init()
	file_Unk3000_OFMFFECMKLE_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GravenInnocenceDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GravenInnocenceDetailInfo); i {
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
			RawDescriptor: file_GravenInnocenceDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GravenInnocenceDetailInfo_proto_goTypes,
		DependencyIndexes: file_GravenInnocenceDetailInfo_proto_depIdxs,
		MessageInfos:      file_GravenInnocenceDetailInfo_proto_msgTypes,
	}.Build()
	File_GravenInnocenceDetailInfo_proto = out.File
	file_GravenInnocenceDetailInfo_proto_rawDesc = nil
	file_GravenInnocenceDetailInfo_proto_goTypes = nil
	file_GravenInnocenceDetailInfo_proto_depIdxs = nil
}
