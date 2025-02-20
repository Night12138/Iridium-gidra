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
// source: Unk2700_FGJFFMPOJON.proto

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

type Unk2700_FGJFFMPOJON struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname            string            `protobuf:"bytes,7,opt,name=nickname,proto3" json:"nickname,omitempty"`
	RemarkName          string            `protobuf:"bytes,3,opt,name=remark_name,json=remarkName,proto3" json:"remark_name,omitempty"`
	ProfilePicture      *ProfilePicture   `protobuf:"bytes,11,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty"`
	Unk2700_IFCNGIPPOAE map[uint32]uint32 `protobuf:"bytes,9,rep,name=Unk2700_IFCNGIPPOAE,json=Unk2700IFCNGIPPOAE,proto3" json:"Unk2700_IFCNGIPPOAE,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Uid                 uint32            `protobuf:"varint,8,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *Unk2700_FGJFFMPOJON) Reset() {
	*x = Unk2700_FGJFFMPOJON{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2700_FGJFFMPOJON_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2700_FGJFFMPOJON) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2700_FGJFFMPOJON) ProtoMessage() {}

func (x *Unk2700_FGJFFMPOJON) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2700_FGJFFMPOJON_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2700_FGJFFMPOJON.ProtoReflect.Descriptor instead.
func (*Unk2700_FGJFFMPOJON) Descriptor() ([]byte, []int) {
	return file_Unk2700_FGJFFMPOJON_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2700_FGJFFMPOJON) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Unk2700_FGJFFMPOJON) GetRemarkName() string {
	if x != nil {
		return x.RemarkName
	}
	return ""
}

func (x *Unk2700_FGJFFMPOJON) GetProfilePicture() *ProfilePicture {
	if x != nil {
		return x.ProfilePicture
	}
	return nil
}

func (x *Unk2700_FGJFFMPOJON) GetUnk2700_IFCNGIPPOAE() map[uint32]uint32 {
	if x != nil {
		return x.Unk2700_IFCNGIPPOAE
	}
	return nil
}

func (x *Unk2700_FGJFFMPOJON) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

var File_Unk2700_FGJFFMPOJON_proto protoreflect.FileDescriptor

var file_Unk2700_FGJFFMPOJON_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x46, 0x47, 0x4a, 0x46, 0x46, 0x4d,
	0x50, 0x4f, 0x4a, 0x4f, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc4, 0x02, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x46, 0x47,
	0x4a, 0x46, 0x46, 0x4d, 0x50, 0x4f, 0x4a, 0x4f, 0x4e, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x5d, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x49, 0x46, 0x43, 0x4e,
	0x47, 0x49, 0x50, 0x50, 0x4f, 0x41, 0x45, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x46, 0x47, 0x4a, 0x46, 0x46, 0x4d, 0x50, 0x4f,
	0x4a, 0x4f, 0x4e, 0x2e, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x49, 0x46, 0x43, 0x4e, 0x47,
	0x49, 0x50, 0x50, 0x4f, 0x41, 0x45, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x49, 0x46, 0x43, 0x4e, 0x47, 0x49, 0x50, 0x50, 0x4f, 0x41, 0x45, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x1a, 0x45, 0x0a, 0x17, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x49, 0x46, 0x43, 0x4e,
	0x47, 0x49, 0x50, 0x50, 0x4f, 0x41, 0x45, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Unk2700_FGJFFMPOJON_proto_rawDescOnce sync.Once
	file_Unk2700_FGJFFMPOJON_proto_rawDescData = file_Unk2700_FGJFFMPOJON_proto_rawDesc
)

func file_Unk2700_FGJFFMPOJON_proto_rawDescGZIP() []byte {
	file_Unk2700_FGJFFMPOJON_proto_rawDescOnce.Do(func() {
		file_Unk2700_FGJFFMPOJON_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2700_FGJFFMPOJON_proto_rawDescData)
	})
	return file_Unk2700_FGJFFMPOJON_proto_rawDescData
}

var file_Unk2700_FGJFFMPOJON_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Unk2700_FGJFFMPOJON_proto_goTypes = []interface{}{
	(*Unk2700_FGJFFMPOJON)(nil), // 0: Unk2700_FGJFFMPOJON
	nil,                         // 1: Unk2700_FGJFFMPOJON.Unk2700IFCNGIPPOAEEntry
	(*ProfilePicture)(nil),      // 2: ProfilePicture
}
var file_Unk2700_FGJFFMPOJON_proto_depIdxs = []int32{
	2, // 0: Unk2700_FGJFFMPOJON.profile_picture:type_name -> ProfilePicture
	1, // 1: Unk2700_FGJFFMPOJON.Unk2700_IFCNGIPPOAE:type_name -> Unk2700_FGJFFMPOJON.Unk2700IFCNGIPPOAEEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Unk2700_FGJFFMPOJON_proto_init() }
func file_Unk2700_FGJFFMPOJON_proto_init() {
	if File_Unk2700_FGJFFMPOJON_proto != nil {
		return
	}
	file_ProfilePicture_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk2700_FGJFFMPOJON_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2700_FGJFFMPOJON); i {
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
			RawDescriptor: file_Unk2700_FGJFFMPOJON_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2700_FGJFFMPOJON_proto_goTypes,
		DependencyIndexes: file_Unk2700_FGJFFMPOJON_proto_depIdxs,
		MessageInfos:      file_Unk2700_FGJFFMPOJON_proto_msgTypes,
	}.Build()
	File_Unk2700_FGJFFMPOJON_proto = out.File
	file_Unk2700_FGJFFMPOJON_proto_rawDesc = nil
	file_Unk2700_FGJFFMPOJON_proto_goTypes = nil
	file_Unk2700_FGJFFMPOJON_proto_depIdxs = nil
}
