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
// source: SpiceActivityDetailInfo.proto

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

type SpiceActivityDetailInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_IGMHNDNGNPG uint32        `protobuf:"varint,15,opt,name=Unk2700_IGMHNDNGNPG,json=Unk2700IGMHNDNGNPG,proto3" json:"Unk2700_IGMHNDNGNPG,omitempty"`
	SpiceStageList      []*SpiceStage `protobuf:"bytes,7,rep,name=spice_stage_list,json=spiceStageList,proto3" json:"spice_stage_list,omitempty"`
	Unk2700_KIAHJKGOLGO uint32        `protobuf:"varint,13,opt,name=Unk2700_KIAHJKGOLGO,json=Unk2700KIAHJKGOLGO,proto3" json:"Unk2700_KIAHJKGOLGO,omitempty"`
}

func (x *SpiceActivityDetailInfo) Reset() {
	*x = SpiceActivityDetailInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SpiceActivityDetailInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpiceActivityDetailInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpiceActivityDetailInfo) ProtoMessage() {}

func (x *SpiceActivityDetailInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SpiceActivityDetailInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpiceActivityDetailInfo.ProtoReflect.Descriptor instead.
func (*SpiceActivityDetailInfo) Descriptor() ([]byte, []int) {
	return file_SpiceActivityDetailInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SpiceActivityDetailInfo) GetUnk2700_IGMHNDNGNPG() uint32 {
	if x != nil {
		return x.Unk2700_IGMHNDNGNPG
	}
	return 0
}

func (x *SpiceActivityDetailInfo) GetSpiceStageList() []*SpiceStage {
	if x != nil {
		return x.SpiceStageList
	}
	return nil
}

func (x *SpiceActivityDetailInfo) GetUnk2700_KIAHJKGOLGO() uint32 {
	if x != nil {
		return x.Unk2700_KIAHJKGOLGO
	}
	return 0
}

var File_SpiceActivityDetailInfo_proto protoreflect.FileDescriptor

var file_SpiceActivityDetailInfo_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x53, 0x70, 0x69, 0x63, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x10, 0x53, 0x70, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x17, 0x53, 0x70, 0x69, 0x63, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a,
	0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x49, 0x47, 0x4d, 0x48, 0x4e, 0x44, 0x4e,
	0x47, 0x4e, 0x50, 0x47, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x49, 0x47, 0x4d, 0x48, 0x4e, 0x44, 0x4e, 0x47, 0x4e, 0x50, 0x47, 0x12, 0x35,
	0x0a, 0x10, 0x73, 0x70, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x70, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x73, 0x70, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x4b, 0x49, 0x41, 0x48, 0x4a, 0x4b, 0x47, 0x4f, 0x4c, 0x47, 0x4f, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4b, 0x49, 0x41, 0x48, 0x4a,
	0x4b, 0x47, 0x4f, 0x4c, 0x47, 0x4f, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SpiceActivityDetailInfo_proto_rawDescOnce sync.Once
	file_SpiceActivityDetailInfo_proto_rawDescData = file_SpiceActivityDetailInfo_proto_rawDesc
)

func file_SpiceActivityDetailInfo_proto_rawDescGZIP() []byte {
	file_SpiceActivityDetailInfo_proto_rawDescOnce.Do(func() {
		file_SpiceActivityDetailInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SpiceActivityDetailInfo_proto_rawDescData)
	})
	return file_SpiceActivityDetailInfo_proto_rawDescData
}

var file_SpiceActivityDetailInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SpiceActivityDetailInfo_proto_goTypes = []interface{}{
	(*SpiceActivityDetailInfo)(nil), // 0: SpiceActivityDetailInfo
	(*SpiceStage)(nil),              // 1: SpiceStage
}
var file_SpiceActivityDetailInfo_proto_depIdxs = []int32{
	1, // 0: SpiceActivityDetailInfo.spice_stage_list:type_name -> SpiceStage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SpiceActivityDetailInfo_proto_init() }
func file_SpiceActivityDetailInfo_proto_init() {
	if File_SpiceActivityDetailInfo_proto != nil {
		return
	}
	file_SpiceStage_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SpiceActivityDetailInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpiceActivityDetailInfo); i {
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
			RawDescriptor: file_SpiceActivityDetailInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SpiceActivityDetailInfo_proto_goTypes,
		DependencyIndexes: file_SpiceActivityDetailInfo_proto_depIdxs,
		MessageInfos:      file_SpiceActivityDetailInfo_proto_msgTypes,
	}.Build()
	File_SpiceActivityDetailInfo_proto = out.File
	file_SpiceActivityDetailInfo_proto_rawDesc = nil
	file_SpiceActivityDetailInfo_proto_goTypes = nil
	file_SpiceActivityDetailInfo_proto_depIdxs = nil
}
