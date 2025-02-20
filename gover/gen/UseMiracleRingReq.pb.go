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
// source: UseMiracleRingReq.proto

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

type UseMiracleRingReq_MiracleRingOpType int32

const (
	UseMiracleRingReq_MIRACLE_RING_OP_TYPE_NONE    UseMiracleRingReq_MiracleRingOpType = 0
	UseMiracleRingReq_MIRACLE_RING_OP_TYPE_PLACE   UseMiracleRingReq_MiracleRingOpType = 1
	UseMiracleRingReq_MIRACLE_RING_OP_TYPE_RETRACT UseMiracleRingReq_MiracleRingOpType = 2
)

// Enum value maps for UseMiracleRingReq_MiracleRingOpType.
var (
	UseMiracleRingReq_MiracleRingOpType_name = map[int32]string{
		0: "MIRACLE_RING_OP_TYPE_NONE",
		1: "MIRACLE_RING_OP_TYPE_PLACE",
		2: "MIRACLE_RING_OP_TYPE_RETRACT",
	}
	UseMiracleRingReq_MiracleRingOpType_value = map[string]int32{
		"MIRACLE_RING_OP_TYPE_NONE":    0,
		"MIRACLE_RING_OP_TYPE_PLACE":   1,
		"MIRACLE_RING_OP_TYPE_RETRACT": 2,
	}
)

func (x UseMiracleRingReq_MiracleRingOpType) Enum() *UseMiracleRingReq_MiracleRingOpType {
	p := new(UseMiracleRingReq_MiracleRingOpType)
	*p = x
	return p
}

func (x UseMiracleRingReq_MiracleRingOpType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UseMiracleRingReq_MiracleRingOpType) Descriptor() protoreflect.EnumDescriptor {
	return file_UseMiracleRingReq_proto_enumTypes[0].Descriptor()
}

func (UseMiracleRingReq_MiracleRingOpType) Type() protoreflect.EnumType {
	return &file_UseMiracleRingReq_proto_enumTypes[0]
}

func (x UseMiracleRingReq_MiracleRingOpType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UseMiracleRingReq_MiracleRingOpType.Descriptor instead.
func (UseMiracleRingReq_MiracleRingOpType) EnumDescriptor() ([]byte, []int) {
	return file_UseMiracleRingReq_proto_rawDescGZIP(), []int{0, 0}
}

// CmdId: 5226
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type UseMiracleRingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MiracleRingOpType uint32  `protobuf:"varint,13,opt,name=miracle_ring_op_type,json=miracleRingOpType,proto3" json:"miracle_ring_op_type,omitempty"`
	Pos               *Vector `protobuf:"bytes,8,opt,name=pos,proto3" json:"pos,omitempty"`
	Rot               *Vector `protobuf:"bytes,7,opt,name=rot,proto3" json:"rot,omitempty"`
}

func (x *UseMiracleRingReq) Reset() {
	*x = UseMiracleRingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UseMiracleRingReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UseMiracleRingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UseMiracleRingReq) ProtoMessage() {}

func (x *UseMiracleRingReq) ProtoReflect() protoreflect.Message {
	mi := &file_UseMiracleRingReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UseMiracleRingReq.ProtoReflect.Descriptor instead.
func (*UseMiracleRingReq) Descriptor() ([]byte, []int) {
	return file_UseMiracleRingReq_proto_rawDescGZIP(), []int{0}
}

func (x *UseMiracleRingReq) GetMiracleRingOpType() uint32 {
	if x != nil {
		return x.MiracleRingOpType
	}
	return 0
}

func (x *UseMiracleRingReq) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *UseMiracleRingReq) GetRot() *Vector {
	if x != nil {
		return x.Rot
	}
	return nil
}

var File_UseMiracleRingReq_proto protoreflect.FileDescriptor

var file_UseMiracleRingReq_proto_rawDesc = []byte{
	0x0a, 0x17, 0x55, 0x73, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x01, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x4d,
	0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x2f, 0x0a,
	0x14, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x6d, 0x69, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12, 0x19, 0x0a, 0x03, 0x72, 0x6f, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x03, 0x72, 0x6f, 0x74, 0x22, 0x74, 0x0a, 0x11, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52,
	0x69, 0x6e, 0x67, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x49, 0x52,
	0x41, 0x43, 0x4c, 0x45, 0x5f, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x50, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x49, 0x52, 0x41,
	0x43, 0x4c, 0x45, 0x5f, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x49, 0x52, 0x41,
	0x43, 0x4c, 0x45, 0x5f, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x52, 0x45, 0x54, 0x52, 0x41, 0x43, 0x54, 0x10, 0x02, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_UseMiracleRingReq_proto_rawDescOnce sync.Once
	file_UseMiracleRingReq_proto_rawDescData = file_UseMiracleRingReq_proto_rawDesc
)

func file_UseMiracleRingReq_proto_rawDescGZIP() []byte {
	file_UseMiracleRingReq_proto_rawDescOnce.Do(func() {
		file_UseMiracleRingReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_UseMiracleRingReq_proto_rawDescData)
	})
	return file_UseMiracleRingReq_proto_rawDescData
}

var file_UseMiracleRingReq_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_UseMiracleRingReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_UseMiracleRingReq_proto_goTypes = []interface{}{
	(UseMiracleRingReq_MiracleRingOpType)(0), // 0: UseMiracleRingReq.MiracleRingOpType
	(*UseMiracleRingReq)(nil),                // 1: UseMiracleRingReq
	(*Vector)(nil),                           // 2: Vector
}
var file_UseMiracleRingReq_proto_depIdxs = []int32{
	2, // 0: UseMiracleRingReq.pos:type_name -> Vector
	2, // 1: UseMiracleRingReq.rot:type_name -> Vector
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_UseMiracleRingReq_proto_init() }
func file_UseMiracleRingReq_proto_init() {
	if File_UseMiracleRingReq_proto != nil {
		return
	}
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_UseMiracleRingReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UseMiracleRingReq); i {
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
			RawDescriptor: file_UseMiracleRingReq_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UseMiracleRingReq_proto_goTypes,
		DependencyIndexes: file_UseMiracleRingReq_proto_depIdxs,
		EnumInfos:         file_UseMiracleRingReq_proto_enumTypes,
		MessageInfos:      file_UseMiracleRingReq_proto_msgTypes,
	}.Build()
	File_UseMiracleRingReq_proto = out.File
	file_UseMiracleRingReq_proto_rawDesc = nil
	file_UseMiracleRingReq_proto_goTypes = nil
	file_UseMiracleRingReq_proto_depIdxs = nil
}
