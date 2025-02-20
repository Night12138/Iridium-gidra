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
// source: RogueDiaryStage.proto

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

type RogueDiaryStage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StageId             uint32               `protobuf:"varint,1,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	BestRecord          *Unk2700_CMOCCENBOLJ `protobuf:"bytes,12,opt,name=best_record,json=bestRecord,proto3" json:"best_record,omitempty"`
	Unk2700_PEDCFBJLHGP bool                 `protobuf:"varint,10,opt,name=Unk2700_PEDCFBJLHGP,json=Unk2700PEDCFBJLHGP,proto3" json:"Unk2700_PEDCFBJLHGP,omitempty"`
}

func (x *RogueDiaryStage) Reset() {
	*x = RogueDiaryStage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueDiaryStage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueDiaryStage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueDiaryStage) ProtoMessage() {}

func (x *RogueDiaryStage) ProtoReflect() protoreflect.Message {
	mi := &file_RogueDiaryStage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueDiaryStage.ProtoReflect.Descriptor instead.
func (*RogueDiaryStage) Descriptor() ([]byte, []int) {
	return file_RogueDiaryStage_proto_rawDescGZIP(), []int{0}
}

func (x *RogueDiaryStage) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *RogueDiaryStage) GetBestRecord() *Unk2700_CMOCCENBOLJ {
	if x != nil {
		return x.BestRecord
	}
	return nil
}

func (x *RogueDiaryStage) GetUnk2700_PEDCFBJLHGP() bool {
	if x != nil {
		return x.Unk2700_PEDCFBJLHGP
	}
	return false
}

var File_RogueDiaryStage_proto protoreflect.FileDescriptor

var file_RogueDiaryStage_proto_rawDesc = []byte{
	0x0a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x43, 0x4d, 0x4f, 0x43, 0x43, 0x45, 0x4e, 0x42, 0x4f, 0x4c, 0x4a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x0f, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x44, 0x69, 0x61, 0x72,
	0x79, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x35, 0x0a, 0x0b, 0x62, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x43, 0x4d, 0x4f, 0x43, 0x43, 0x45, 0x4e, 0x42, 0x4f, 0x4c, 0x4a, 0x52, 0x0a, 0x62, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x5f, 0x50, 0x45, 0x44, 0x43, 0x46, 0x42, 0x4a, 0x4c, 0x48, 0x47, 0x50, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x50, 0x45,
	0x44, 0x43, 0x46, 0x42, 0x4a, 0x4c, 0x48, 0x47, 0x50, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueDiaryStage_proto_rawDescOnce sync.Once
	file_RogueDiaryStage_proto_rawDescData = file_RogueDiaryStage_proto_rawDesc
)

func file_RogueDiaryStage_proto_rawDescGZIP() []byte {
	file_RogueDiaryStage_proto_rawDescOnce.Do(func() {
		file_RogueDiaryStage_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueDiaryStage_proto_rawDescData)
	})
	return file_RogueDiaryStage_proto_rawDescData
}

var file_RogueDiaryStage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueDiaryStage_proto_goTypes = []interface{}{
	(*RogueDiaryStage)(nil),     // 0: RogueDiaryStage
	(*Unk2700_CMOCCENBOLJ)(nil), // 1: Unk2700_CMOCCENBOLJ
}
var file_RogueDiaryStage_proto_depIdxs = []int32{
	1, // 0: RogueDiaryStage.best_record:type_name -> Unk2700_CMOCCENBOLJ
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_RogueDiaryStage_proto_init() }
func file_RogueDiaryStage_proto_init() {
	if File_RogueDiaryStage_proto != nil {
		return
	}
	file_Unk2700_CMOCCENBOLJ_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueDiaryStage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueDiaryStage); i {
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
			RawDescriptor: file_RogueDiaryStage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueDiaryStage_proto_goTypes,
		DependencyIndexes: file_RogueDiaryStage_proto_depIdxs,
		MessageInfos:      file_RogueDiaryStage_proto_msgTypes,
	}.Build()
	File_RogueDiaryStage_proto = out.File
	file_RogueDiaryStage_proto_rawDesc = nil
	file_RogueDiaryStage_proto_goTypes = nil
	file_RogueDiaryStage_proto_depIdxs = nil
}
