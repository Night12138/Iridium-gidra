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
// source: H5ActivityInfo.proto

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

type H5ActivityInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	H5ActivityId     uint32 `protobuf:"varint,3,opt,name=h5_activity_id,json=h5ActivityId,proto3" json:"h5_activity_id,omitempty"`
	Url              string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	IsEntranceOpen   bool   `protobuf:"varint,7,opt,name=is_entrance_open,json=isEntranceOpen,proto3" json:"is_entrance_open,omitempty"`
	H5ScheduleId     uint32 `protobuf:"varint,8,opt,name=h5_schedule_id,json=h5ScheduleId,proto3" json:"h5_schedule_id,omitempty"`
	EndTime          uint32 `protobuf:"varint,10,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	PrefabPath       string `protobuf:"bytes,11,opt,name=prefab_path,json=prefabPath,proto3" json:"prefab_path,omitempty"`
	ContentCloseTime uint32 `protobuf:"varint,2,opt,name=content_close_time,json=contentCloseTime,proto3" json:"content_close_time,omitempty"`
	BeginTime        uint32 `protobuf:"varint,13,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
}

func (x *H5ActivityInfo) Reset() {
	*x = H5ActivityInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_H5ActivityInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *H5ActivityInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*H5ActivityInfo) ProtoMessage() {}

func (x *H5ActivityInfo) ProtoReflect() protoreflect.Message {
	mi := &file_H5ActivityInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use H5ActivityInfo.ProtoReflect.Descriptor instead.
func (*H5ActivityInfo) Descriptor() ([]byte, []int) {
	return file_H5ActivityInfo_proto_rawDescGZIP(), []int{0}
}

func (x *H5ActivityInfo) GetH5ActivityId() uint32 {
	if x != nil {
		return x.H5ActivityId
	}
	return 0
}

func (x *H5ActivityInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *H5ActivityInfo) GetIsEntranceOpen() bool {
	if x != nil {
		return x.IsEntranceOpen
	}
	return false
}

func (x *H5ActivityInfo) GetH5ScheduleId() uint32 {
	if x != nil {
		return x.H5ScheduleId
	}
	return 0
}

func (x *H5ActivityInfo) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *H5ActivityInfo) GetPrefabPath() string {
	if x != nil {
		return x.PrefabPath
	}
	return ""
}

func (x *H5ActivityInfo) GetContentCloseTime() uint32 {
	if x != nil {
		return x.ContentCloseTime
	}
	return 0
}

func (x *H5ActivityInfo) GetBeginTime() uint32 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

var File_H5ActivityInfo_proto protoreflect.FileDescriptor

var file_H5ActivityInfo_proto_rawDesc = []byte{
	0x0a, 0x14, 0x48, 0x35, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x02, 0x0a, 0x0e, 0x48, 0x35, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0e, 0x68, 0x35, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x68, 0x35, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x68,
	0x35, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x68, 0x35, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x65, 0x66, 0x61, 0x62, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x66, 0x61, 0x62, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2c, 0x0a,
	0x12, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_H5ActivityInfo_proto_rawDescOnce sync.Once
	file_H5ActivityInfo_proto_rawDescData = file_H5ActivityInfo_proto_rawDesc
)

func file_H5ActivityInfo_proto_rawDescGZIP() []byte {
	file_H5ActivityInfo_proto_rawDescOnce.Do(func() {
		file_H5ActivityInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_H5ActivityInfo_proto_rawDescData)
	})
	return file_H5ActivityInfo_proto_rawDescData
}

var file_H5ActivityInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_H5ActivityInfo_proto_goTypes = []interface{}{
	(*H5ActivityInfo)(nil), // 0: H5ActivityInfo
}
var file_H5ActivityInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_H5ActivityInfo_proto_init() }
func file_H5ActivityInfo_proto_init() {
	if File_H5ActivityInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_H5ActivityInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*H5ActivityInfo); i {
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
			RawDescriptor: file_H5ActivityInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_H5ActivityInfo_proto_goTypes,
		DependencyIndexes: file_H5ActivityInfo_proto_depIdxs,
		MessageInfos:      file_H5ActivityInfo_proto_msgTypes,
	}.Build()
	File_H5ActivityInfo_proto = out.File
	file_H5ActivityInfo_proto_rawDesc = nil
	file_H5ActivityInfo_proto_goTypes = nil
	file_H5ActivityInfo_proto_depIdxs = nil
}
