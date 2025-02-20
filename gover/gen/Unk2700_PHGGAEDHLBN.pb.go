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
// source: Unk2700_PHGGAEDHLBN.proto

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

type Unk2700_PHGGAEDHLBN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_ANHJAFDEACF []uint32 `protobuf:"varint,1,rep,packed,name=Unk2700_ANHJAFDEACF,json=Unk2700ANHJAFDEACF,proto3" json:"Unk2700_ANHJAFDEACF,omitempty"`
	Unk2700_IBDCFAMBGOK bool     `protobuf:"varint,14,opt,name=Unk2700_IBDCFAMBGOK,json=Unk2700IBDCFAMBGOK,proto3" json:"Unk2700_IBDCFAMBGOK,omitempty"`
	Unk2700_KENGEGJGAEL uint32   `protobuf:"varint,6,opt,name=Unk2700_KENGEGJGAEL,json=Unk2700KENGEGJGAEL,proto3" json:"Unk2700_KENGEGJGAEL,omitempty"`
	Unk2700_DOIMMBJDALB uint32   `protobuf:"varint,4,opt,name=Unk2700_DOIMMBJDALB,json=Unk2700DOIMMBJDALB,proto3" json:"Unk2700_DOIMMBJDALB,omitempty"`
	Unk2700_FKLBCNLBBNM bool     `protobuf:"varint,3,opt,name=Unk2700_FKLBCNLBBNM,json=Unk2700FKLBCNLBBNM,proto3" json:"Unk2700_FKLBCNLBBNM,omitempty"`
	Unk2700_IFNFCNNBPIB uint32   `protobuf:"varint,10,opt,name=Unk2700_IFNFCNNBPIB,json=Unk2700IFNFCNNBPIB,proto3" json:"Unk2700_IFNFCNNBPIB,omitempty"`
	Unk2700_PBBPGFMNMNJ uint32   `protobuf:"varint,9,opt,name=Unk2700_PBBPGFMNMNJ,json=Unk2700PBBPGFMNMNJ,proto3" json:"Unk2700_PBBPGFMNMNJ,omitempty"`
}

func (x *Unk2700_PHGGAEDHLBN) Reset() {
	*x = Unk2700_PHGGAEDHLBN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2700_PHGGAEDHLBN_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2700_PHGGAEDHLBN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2700_PHGGAEDHLBN) ProtoMessage() {}

func (x *Unk2700_PHGGAEDHLBN) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2700_PHGGAEDHLBN_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2700_PHGGAEDHLBN.ProtoReflect.Descriptor instead.
func (*Unk2700_PHGGAEDHLBN) Descriptor() ([]byte, []int) {
	return file_Unk2700_PHGGAEDHLBN_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2700_PHGGAEDHLBN) GetUnk2700_ANHJAFDEACF() []uint32 {
	if x != nil {
		return x.Unk2700_ANHJAFDEACF
	}
	return nil
}

func (x *Unk2700_PHGGAEDHLBN) GetUnk2700_IBDCFAMBGOK() bool {
	if x != nil {
		return x.Unk2700_IBDCFAMBGOK
	}
	return false
}

func (x *Unk2700_PHGGAEDHLBN) GetUnk2700_KENGEGJGAEL() uint32 {
	if x != nil {
		return x.Unk2700_KENGEGJGAEL
	}
	return 0
}

func (x *Unk2700_PHGGAEDHLBN) GetUnk2700_DOIMMBJDALB() uint32 {
	if x != nil {
		return x.Unk2700_DOIMMBJDALB
	}
	return 0
}

func (x *Unk2700_PHGGAEDHLBN) GetUnk2700_FKLBCNLBBNM() bool {
	if x != nil {
		return x.Unk2700_FKLBCNLBBNM
	}
	return false
}

func (x *Unk2700_PHGGAEDHLBN) GetUnk2700_IFNFCNNBPIB() uint32 {
	if x != nil {
		return x.Unk2700_IFNFCNNBPIB
	}
	return 0
}

func (x *Unk2700_PHGGAEDHLBN) GetUnk2700_PBBPGFMNMNJ() uint32 {
	if x != nil {
		return x.Unk2700_PBBPGFMNMNJ
	}
	return 0
}

var File_Unk2700_PHGGAEDHLBN_proto protoreflect.FileDescriptor

var file_Unk2700_PHGGAEDHLBN_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x50, 0x48, 0x47, 0x47, 0x41, 0x45,
	0x44, 0x48, 0x4c, 0x42, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x02, 0x0a, 0x13,
	0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x50, 0x48, 0x47, 0x47, 0x41, 0x45, 0x44, 0x48,
	0x4c, 0x42, 0x4e, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x41,
	0x4e, 0x48, 0x4a, 0x41, 0x46, 0x44, 0x45, 0x41, 0x43, 0x46, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x41, 0x4e, 0x48, 0x4a, 0x41, 0x46, 0x44,
	0x45, 0x41, 0x43, 0x46, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x49, 0x42, 0x44, 0x43, 0x46, 0x41, 0x4d, 0x42, 0x47, 0x4f, 0x4b, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x49, 0x42, 0x44, 0x43, 0x46, 0x41,
	0x4d, 0x42, 0x47, 0x4f, 0x4b, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x4b, 0x45, 0x4e, 0x47, 0x45, 0x47, 0x4a, 0x47, 0x41, 0x45, 0x4c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4b, 0x45, 0x4e, 0x47, 0x45,
	0x47, 0x4a, 0x47, 0x41, 0x45, 0x4c, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30,
	0x30, 0x5f, 0x44, 0x4f, 0x49, 0x4d, 0x4d, 0x42, 0x4a, 0x44, 0x41, 0x4c, 0x42, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x44, 0x4f, 0x49, 0x4d,
	0x4d, 0x42, 0x4a, 0x44, 0x41, 0x4c, 0x42, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x5f, 0x46, 0x4b, 0x4c, 0x42, 0x43, 0x4e, 0x4c, 0x42, 0x42, 0x4e, 0x4d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x46, 0x4b, 0x4c,
	0x42, 0x43, 0x4e, 0x4c, 0x42, 0x42, 0x4e, 0x4d, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x5f, 0x49, 0x46, 0x4e, 0x46, 0x43, 0x4e, 0x4e, 0x42, 0x50, 0x49, 0x42, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x49, 0x46,
	0x4e, 0x46, 0x43, 0x4e, 0x4e, 0x42, 0x50, 0x49, 0x42, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x50, 0x42, 0x42, 0x50, 0x47, 0x46, 0x4d, 0x4e, 0x4d, 0x4e, 0x4a,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x50,
	0x42, 0x42, 0x50, 0x47, 0x46, 0x4d, 0x4e, 0x4d, 0x4e, 0x4a, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk2700_PHGGAEDHLBN_proto_rawDescOnce sync.Once
	file_Unk2700_PHGGAEDHLBN_proto_rawDescData = file_Unk2700_PHGGAEDHLBN_proto_rawDesc
)

func file_Unk2700_PHGGAEDHLBN_proto_rawDescGZIP() []byte {
	file_Unk2700_PHGGAEDHLBN_proto_rawDescOnce.Do(func() {
		file_Unk2700_PHGGAEDHLBN_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2700_PHGGAEDHLBN_proto_rawDescData)
	})
	return file_Unk2700_PHGGAEDHLBN_proto_rawDescData
}

var file_Unk2700_PHGGAEDHLBN_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk2700_PHGGAEDHLBN_proto_goTypes = []interface{}{
	(*Unk2700_PHGGAEDHLBN)(nil), // 0: Unk2700_PHGGAEDHLBN
}
var file_Unk2700_PHGGAEDHLBN_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Unk2700_PHGGAEDHLBN_proto_init() }
func file_Unk2700_PHGGAEDHLBN_proto_init() {
	if File_Unk2700_PHGGAEDHLBN_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Unk2700_PHGGAEDHLBN_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2700_PHGGAEDHLBN); i {
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
			RawDescriptor: file_Unk2700_PHGGAEDHLBN_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2700_PHGGAEDHLBN_proto_goTypes,
		DependencyIndexes: file_Unk2700_PHGGAEDHLBN_proto_depIdxs,
		MessageInfos:      file_Unk2700_PHGGAEDHLBN_proto_msgTypes,
	}.Build()
	File_Unk2700_PHGGAEDHLBN_proto = out.File
	file_Unk2700_PHGGAEDHLBN_proto_rawDesc = nil
	file_Unk2700_PHGGAEDHLBN_proto_goTypes = nil
	file_Unk2700_PHGGAEDHLBN_proto_depIdxs = nil
}
