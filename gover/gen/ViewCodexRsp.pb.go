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
// source: ViewCodexRsp.proto

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

// CmdId: 4201
// EnetChannelId: 0
// EnetIsReliable: true
type ViewCodexRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode             int32            `protobuf:"varint,12,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Unk2800_IPOCJIPGNEJ []uint32         `protobuf:"varint,10,rep,packed,name=Unk2800_IPOCJIPGNEJ,json=Unk2800IPOCJIPGNEJ,proto3" json:"Unk2800_IPOCJIPGNEJ,omitempty"`
	Unk2700_DFJJHFHHIHF []uint32         `protobuf:"varint,3,rep,packed,name=Unk2700_DFJJHFHHIHF,json=Unk2700DFJJHFHHIHF,proto3" json:"Unk2700_DFJJHFHHIHF,omitempty"`
	TypeDataList        []*CodexTypeData `protobuf:"bytes,9,rep,name=type_data_list,json=typeDataList,proto3" json:"type_data_list,omitempty"`
	Unk2800_OIPJCEPGJCF []uint32         `protobuf:"varint,15,rep,packed,name=Unk2800_OIPJCEPGJCF,json=Unk2800OIPJCEPGJCF,proto3" json:"Unk2800_OIPJCEPGJCF,omitempty"`
}

func (x *ViewCodexRsp) Reset() {
	*x = ViewCodexRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ViewCodexRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewCodexRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCodexRsp) ProtoMessage() {}

func (x *ViewCodexRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ViewCodexRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewCodexRsp.ProtoReflect.Descriptor instead.
func (*ViewCodexRsp) Descriptor() ([]byte, []int) {
	return file_ViewCodexRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ViewCodexRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ViewCodexRsp) GetUnk2800_IPOCJIPGNEJ() []uint32 {
	if x != nil {
		return x.Unk2800_IPOCJIPGNEJ
	}
	return nil
}

func (x *ViewCodexRsp) GetUnk2700_DFJJHFHHIHF() []uint32 {
	if x != nil {
		return x.Unk2700_DFJJHFHHIHF
	}
	return nil
}

func (x *ViewCodexRsp) GetTypeDataList() []*CodexTypeData {
	if x != nil {
		return x.TypeDataList
	}
	return nil
}

func (x *ViewCodexRsp) GetUnk2800_OIPJCEPGJCF() []uint32 {
	if x != nil {
		return x.Unk2800_OIPJCEPGJCF
	}
	return nil
}

var File_ViewCodexRsp_proto protoreflect.FileDescriptor

var file_ViewCodexRsp_proto_rawDesc = []byte{
	0x0a, 0x12, 0x56, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x52, 0x73, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01, 0x0a, 0x0c, 0x56, 0x69,
	0x65, 0x77, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f,
	0x49, 0x50, 0x4f, 0x43, 0x4a, 0x49, 0x50, 0x47, 0x4e, 0x45, 0x4a, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x49, 0x50, 0x4f, 0x43, 0x4a, 0x49,
	0x50, 0x47, 0x4e, 0x45, 0x4a, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x44, 0x46, 0x4a, 0x4a, 0x48, 0x46, 0x48, 0x48, 0x49, 0x48, 0x46, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x44, 0x46, 0x4a, 0x4a, 0x48,
	0x46, 0x48, 0x48, 0x49, 0x48, 0x46, 0x12, 0x34, 0x0a, 0x0e, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x78, 0x54, 0x79, 0x70, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c,
	0x74, 0x79, 0x70, 0x65, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x32, 0x38, 0x30, 0x30, 0x5f, 0x4f, 0x49, 0x50, 0x4a, 0x43, 0x45, 0x50, 0x47,
	0x4a, 0x43, 0x46, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x38,
	0x30, 0x30, 0x4f, 0x49, 0x50, 0x4a, 0x43, 0x45, 0x50, 0x47, 0x4a, 0x43, 0x46, 0x42, 0x06, 0x5a,
	0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ViewCodexRsp_proto_rawDescOnce sync.Once
	file_ViewCodexRsp_proto_rawDescData = file_ViewCodexRsp_proto_rawDesc
)

func file_ViewCodexRsp_proto_rawDescGZIP() []byte {
	file_ViewCodexRsp_proto_rawDescOnce.Do(func() {
		file_ViewCodexRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ViewCodexRsp_proto_rawDescData)
	})
	return file_ViewCodexRsp_proto_rawDescData
}

var file_ViewCodexRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ViewCodexRsp_proto_goTypes = []interface{}{
	(*ViewCodexRsp)(nil),  // 0: ViewCodexRsp
	(*CodexTypeData)(nil), // 1: CodexTypeData
}
var file_ViewCodexRsp_proto_depIdxs = []int32{
	1, // 0: ViewCodexRsp.type_data_list:type_name -> CodexTypeData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ViewCodexRsp_proto_init() }
func file_ViewCodexRsp_proto_init() {
	if File_ViewCodexRsp_proto != nil {
		return
	}
	file_CodexTypeData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ViewCodexRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewCodexRsp); i {
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
			RawDescriptor: file_ViewCodexRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ViewCodexRsp_proto_goTypes,
		DependencyIndexes: file_ViewCodexRsp_proto_depIdxs,
		MessageInfos:      file_ViewCodexRsp_proto_msgTypes,
	}.Build()
	File_ViewCodexRsp_proto = out.File
	file_ViewCodexRsp_proto_rawDesc = nil
	file_ViewCodexRsp_proto_goTypes = nil
	file_ViewCodexRsp_proto_depIdxs = nil
}
