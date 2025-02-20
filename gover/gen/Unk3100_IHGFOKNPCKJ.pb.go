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
// source: Unk3100_IHGFOKNPCKJ.proto

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

type Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO int32

const (
	Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO_Unk3100_CHMICKLPAKA Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO = 0
	Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO_Unk3100_GEJFGKILBLO Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO = 1
	Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO_Unk3100_HAFBECHLCIE Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO = 2
)

// Enum value maps for Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO.
var (
	Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO_name = map[int32]string{
		0: "Unk3100_DDADIDBLJGO_Unk3100_CHMICKLPAKA",
		1: "Unk3100_DDADIDBLJGO_Unk3100_GEJFGKILBLO",
		2: "Unk3100_DDADIDBLJGO_Unk3100_HAFBECHLCIE",
	}
	Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO_value = map[string]int32{
		"Unk3100_DDADIDBLJGO_Unk3100_CHMICKLPAKA": 0,
		"Unk3100_DDADIDBLJGO_Unk3100_GEJFGKILBLO": 1,
		"Unk3100_DDADIDBLJGO_Unk3100_HAFBECHLCIE": 2,
	}
)

func (x Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO) Enum() *Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO {
	p := new(Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO)
	*p = x
	return p
}

func (x Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO) Descriptor() protoreflect.EnumDescriptor {
	return file_Unk3100_IHGFOKNPCKJ_proto_enumTypes[0].Descriptor()
}

func (Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO) Type() protoreflect.EnumType {
	return &file_Unk3100_IHGFOKNPCKJ_proto_enumTypes[0]
}

func (x Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO.Descriptor instead.
func (Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO) EnumDescriptor() ([]byte, []int) {
	return file_Unk3100_IHGFOKNPCKJ_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 3160
// EnetChannelId: 0
// EnetIsReliable: true
type Unk3100_IHGFOKNPCKJ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LookPos             *Vector                                 `protobuf:"bytes,8,opt,name=look_pos,json=lookPos,proto3" json:"look_pos,omitempty"`
	TemplateId          uint32                                  `protobuf:"varint,5,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	FollowPos           *Vector                                 `protobuf:"bytes,2,opt,name=follow_pos,json=followPos,proto3" json:"follow_pos,omitempty"`
	EntityId            uint32                                  `protobuf:"varint,12,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	OtherParams         []string                                `protobuf:"bytes,13,rep,name=other_params,json=otherParams,proto3" json:"other_params,omitempty"`
	Unk3100_JHIMHLNPLGA Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO `protobuf:"varint,9,opt,name=Unk3100_JHIMHLNPLGA,json=Unk3100JHIMHLNPLGA,proto3,enum=Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO" json:"Unk3100_JHIMHLNPLGA,omitempty"`
}

func (x *Unk3100_IHGFOKNPCKJ) Reset() {
	*x = Unk3100_IHGFOKNPCKJ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk3100_IHGFOKNPCKJ_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk3100_IHGFOKNPCKJ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk3100_IHGFOKNPCKJ) ProtoMessage() {}

func (x *Unk3100_IHGFOKNPCKJ) ProtoReflect() protoreflect.Message {
	mi := &file_Unk3100_IHGFOKNPCKJ_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk3100_IHGFOKNPCKJ.ProtoReflect.Descriptor instead.
func (*Unk3100_IHGFOKNPCKJ) Descriptor() ([]byte, []int) {
	return file_Unk3100_IHGFOKNPCKJ_proto_rawDescGZIP(), []int{0}
}

func (x *Unk3100_IHGFOKNPCKJ) GetLookPos() *Vector {
	if x != nil {
		return x.LookPos
	}
	return nil
}

func (x *Unk3100_IHGFOKNPCKJ) GetTemplateId() uint32 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

func (x *Unk3100_IHGFOKNPCKJ) GetFollowPos() *Vector {
	if x != nil {
		return x.FollowPos
	}
	return nil
}

func (x *Unk3100_IHGFOKNPCKJ) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (x *Unk3100_IHGFOKNPCKJ) GetOtherParams() []string {
	if x != nil {
		return x.OtherParams
	}
	return nil
}

func (x *Unk3100_IHGFOKNPCKJ) GetUnk3100_JHIMHLNPLGA() Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO {
	if x != nil {
		return x.Unk3100_JHIMHLNPLGA
	}
	return Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO_Unk3100_CHMICKLPAKA
}

var File_Unk3100_IHGFOKNPCKJ_proto protoreflect.FileDescriptor

var file_Unk3100_IHGFOKNPCKJ_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x49, 0x48, 0x47, 0x46, 0x4f, 0x4b,
	0x4e, 0x50, 0x43, 0x4b, 0x4a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x03, 0x0a, 0x13, 0x55, 0x6e,
	0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x49, 0x48, 0x47, 0x46, 0x4f, 0x4b, 0x4e, 0x50, 0x43, 0x4b,
	0x4a, 0x12, 0x22, 0x0a, 0x08, 0x6c, 0x6f, 0x6f, 0x6b, 0x5f, 0x70, 0x6f, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x6c, 0x6f,
	0x6f, 0x6b, 0x50, 0x6f, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x70, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x6f, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f,
	0x74, 0x68, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x59,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x4a, 0x48, 0x49, 0x4d, 0x48, 0x4c,
	0x4e, 0x50, 0x4c, 0x47, 0x41, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x55, 0x6e,
	0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x49, 0x48, 0x47, 0x46, 0x4f, 0x4b, 0x4e, 0x50, 0x43, 0x4b,
	0x4a, 0x2e, 0x55, 0x6e, 0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x44, 0x44, 0x41, 0x44, 0x49, 0x44,
	0x42, 0x4c, 0x4a, 0x47, 0x4f, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x33, 0x31, 0x30, 0x30, 0x4a, 0x48,
	0x49, 0x4d, 0x48, 0x4c, 0x4e, 0x50, 0x4c, 0x47, 0x41, 0x22, 0x9c, 0x01, 0x0a, 0x13, 0x55, 0x6e,
	0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x44, 0x44, 0x41, 0x44, 0x49, 0x44, 0x42, 0x4c, 0x4a, 0x47,
	0x4f, 0x12, 0x2b, 0x0a, 0x27, 0x55, 0x6e, 0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x44, 0x44, 0x41,
	0x44, 0x49, 0x44, 0x42, 0x4c, 0x4a, 0x47, 0x4f, 0x5f, 0x55, 0x6e, 0x6b, 0x33, 0x31, 0x30, 0x30,
	0x5f, 0x43, 0x48, 0x4d, 0x49, 0x43, 0x4b, 0x4c, 0x50, 0x41, 0x4b, 0x41, 0x10, 0x00, 0x12, 0x2b,
	0x0a, 0x27, 0x55, 0x6e, 0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x44, 0x44, 0x41, 0x44, 0x49, 0x44,
	0x42, 0x4c, 0x4a, 0x47, 0x4f, 0x5f, 0x55, 0x6e, 0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x47, 0x45,
	0x4a, 0x46, 0x47, 0x4b, 0x49, 0x4c, 0x42, 0x4c, 0x4f, 0x10, 0x01, 0x12, 0x2b, 0x0a, 0x27, 0x55,
	0x6e, 0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x44, 0x44, 0x41, 0x44, 0x49, 0x44, 0x42, 0x4c, 0x4a,
	0x47, 0x4f, 0x5f, 0x55, 0x6e, 0x6b, 0x33, 0x31, 0x30, 0x30, 0x5f, 0x48, 0x41, 0x46, 0x42, 0x45,
	0x43, 0x48, 0x4c, 0x43, 0x49, 0x45, 0x10, 0x02, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk3100_IHGFOKNPCKJ_proto_rawDescOnce sync.Once
	file_Unk3100_IHGFOKNPCKJ_proto_rawDescData = file_Unk3100_IHGFOKNPCKJ_proto_rawDesc
)

func file_Unk3100_IHGFOKNPCKJ_proto_rawDescGZIP() []byte {
	file_Unk3100_IHGFOKNPCKJ_proto_rawDescOnce.Do(func() {
		file_Unk3100_IHGFOKNPCKJ_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk3100_IHGFOKNPCKJ_proto_rawDescData)
	})
	return file_Unk3100_IHGFOKNPCKJ_proto_rawDescData
}

var file_Unk3100_IHGFOKNPCKJ_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Unk3100_IHGFOKNPCKJ_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk3100_IHGFOKNPCKJ_proto_goTypes = []interface{}{
	(Unk3100_IHGFOKNPCKJ_Unk3100_DDADIDBLJGO)(0), // 0: Unk3100_IHGFOKNPCKJ.Unk3100_DDADIDBLJGO
	(*Unk3100_IHGFOKNPCKJ)(nil),                  // 1: Unk3100_IHGFOKNPCKJ
	(*Vector)(nil),                               // 2: Vector
}
var file_Unk3100_IHGFOKNPCKJ_proto_depIdxs = []int32{
	2, // 0: Unk3100_IHGFOKNPCKJ.look_pos:type_name -> Vector
	2, // 1: Unk3100_IHGFOKNPCKJ.follow_pos:type_name -> Vector
	0, // 2: Unk3100_IHGFOKNPCKJ.Unk3100_JHIMHLNPLGA:type_name -> Unk3100_IHGFOKNPCKJ.Unk3100_DDADIDBLJGO
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Unk3100_IHGFOKNPCKJ_proto_init() }
func file_Unk3100_IHGFOKNPCKJ_proto_init() {
	if File_Unk3100_IHGFOKNPCKJ_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk3100_IHGFOKNPCKJ_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk3100_IHGFOKNPCKJ); i {
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
			RawDescriptor: file_Unk3100_IHGFOKNPCKJ_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk3100_IHGFOKNPCKJ_proto_goTypes,
		DependencyIndexes: file_Unk3100_IHGFOKNPCKJ_proto_depIdxs,
		EnumInfos:         file_Unk3100_IHGFOKNPCKJ_proto_enumTypes,
		MessageInfos:      file_Unk3100_IHGFOKNPCKJ_proto_msgTypes,
	}.Build()
	File_Unk3100_IHGFOKNPCKJ_proto = out.File
	file_Unk3100_IHGFOKNPCKJ_proto_rawDesc = nil
	file_Unk3100_IHGFOKNPCKJ_proto_goTypes = nil
	file_Unk3100_IHGFOKNPCKJ_proto_depIdxs = nil
}
