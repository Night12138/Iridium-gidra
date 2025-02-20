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
// source: ClientTransmitReq.proto

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

// CmdId: 291
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type ClientTransmitReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneId uint32         `protobuf:"varint,2,opt,name=scene_id,json=sceneId,proto3" json:"scene_id,omitempty"`
	Reason  TransmitReason `protobuf:"varint,14,opt,name=reason,proto3,enum=TransmitReason" json:"reason,omitempty"`
	Pos     *Vector        `protobuf:"bytes,1,opt,name=pos,proto3" json:"pos,omitempty"`
	Rot     *Vector        `protobuf:"bytes,9,opt,name=rot,proto3" json:"rot,omitempty"`
}

func (x *ClientTransmitReq) Reset() {
	*x = ClientTransmitReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientTransmitReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientTransmitReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientTransmitReq) ProtoMessage() {}

func (x *ClientTransmitReq) ProtoReflect() protoreflect.Message {
	mi := &file_ClientTransmitReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientTransmitReq.ProtoReflect.Descriptor instead.
func (*ClientTransmitReq) Descriptor() ([]byte, []int) {
	return file_ClientTransmitReq_proto_rawDescGZIP(), []int{0}
}

func (x *ClientTransmitReq) GetSceneId() uint32 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *ClientTransmitReq) GetReason() TransmitReason {
	if x != nil {
		return x.Reason
	}
	return TransmitReason_TRANSMIT_REASON_NONE
}

func (x *ClientTransmitReq) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *ClientTransmitReq) GetRot() *Vector {
	if x != nil {
		return x.Rot
	}
	return nil
}

var File_ClientTransmitReq_proto protoreflect.FileDescriptor

var file_ClientTransmitReq_proto_rawDesc = []byte{
	0x0a, 0x17, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x6d, 0x69, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01,
	0x0a, 0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x27,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x70,
	0x6f, 0x73, 0x12, 0x19, 0x0a, 0x03, 0x72, 0x6f, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x72, 0x6f, 0x74, 0x42, 0x06, 0x5a,
	0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ClientTransmitReq_proto_rawDescOnce sync.Once
	file_ClientTransmitReq_proto_rawDescData = file_ClientTransmitReq_proto_rawDesc
)

func file_ClientTransmitReq_proto_rawDescGZIP() []byte {
	file_ClientTransmitReq_proto_rawDescOnce.Do(func() {
		file_ClientTransmitReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ClientTransmitReq_proto_rawDescData)
	})
	return file_ClientTransmitReq_proto_rawDescData
}

var file_ClientTransmitReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ClientTransmitReq_proto_goTypes = []interface{}{
	(*ClientTransmitReq)(nil), // 0: ClientTransmitReq
	(TransmitReason)(0),       // 1: TransmitReason
	(*Vector)(nil),            // 2: Vector
}
var file_ClientTransmitReq_proto_depIdxs = []int32{
	1, // 0: ClientTransmitReq.reason:type_name -> TransmitReason
	2, // 1: ClientTransmitReq.pos:type_name -> Vector
	2, // 2: ClientTransmitReq.rot:type_name -> Vector
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ClientTransmitReq_proto_init() }
func file_ClientTransmitReq_proto_init() {
	if File_ClientTransmitReq_proto != nil {
		return
	}
	file_TransmitReason_proto_init()
	file_Vector_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ClientTransmitReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientTransmitReq); i {
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
			RawDescriptor: file_ClientTransmitReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ClientTransmitReq_proto_goTypes,
		DependencyIndexes: file_ClientTransmitReq_proto_depIdxs,
		MessageInfos:      file_ClientTransmitReq_proto_msgTypes,
	}.Build()
	File_ClientTransmitReq_proto = out.File
	file_ClientTransmitReq_proto_rawDesc = nil
	file_ClientTransmitReq_proto_goTypes = nil
	file_ClientTransmitReq_proto_depIdxs = nil
}
