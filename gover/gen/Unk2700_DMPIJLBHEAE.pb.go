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
// source: Unk2700_DMPIJLBHEAE.proto

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

type Unk2700_DMPIJLBHEAE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeType uint32 `protobuf:"varint,5,opt,name=challenge_type,json=challengeType,proto3" json:"challenge_type,omitempty"`
	IsUnlock      bool   `protobuf:"varint,12,opt,name=is_unlock,json=isUnlock,proto3" json:"is_unlock,omitempty"`
	// Types that are assignable to Unk2700_AFHAGFONBFM:
	//
	//	*Unk2700_DMPIJLBHEAE_BundleInfo
	//	*Unk2700_DMPIJLBHEAE_ScoreChallengeInfo
	//	*Unk2700_DMPIJLBHEAE_BossChallengeId
	Unk2700_AFHAGFONBFM isUnk2700_DMPIJLBHEAE_Unk2700_AFHAGFONBFM `protobuf_oneof:"Unk2700_AFHAGFONBFM"`
}

func (x *Unk2700_DMPIJLBHEAE) Reset() {
	*x = Unk2700_DMPIJLBHEAE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Unk2700_DMPIJLBHEAE_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unk2700_DMPIJLBHEAE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unk2700_DMPIJLBHEAE) ProtoMessage() {}

func (x *Unk2700_DMPIJLBHEAE) ProtoReflect() protoreflect.Message {
	mi := &file_Unk2700_DMPIJLBHEAE_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unk2700_DMPIJLBHEAE.ProtoReflect.Descriptor instead.
func (*Unk2700_DMPIJLBHEAE) Descriptor() ([]byte, []int) {
	return file_Unk2700_DMPIJLBHEAE_proto_rawDescGZIP(), []int{0}
}

func (x *Unk2700_DMPIJLBHEAE) GetChallengeType() uint32 {
	if x != nil {
		return x.ChallengeType
	}
	return 0
}

func (x *Unk2700_DMPIJLBHEAE) GetIsUnlock() bool {
	if x != nil {
		return x.IsUnlock
	}
	return false
}

func (m *Unk2700_DMPIJLBHEAE) GetUnk2700_AFHAGFONBFM() isUnk2700_DMPIJLBHEAE_Unk2700_AFHAGFONBFM {
	if m != nil {
		return m.Unk2700_AFHAGFONBFM
	}
	return nil
}

func (x *Unk2700_DMPIJLBHEAE) GetBundleInfo() *BundleInfo {
	if x, ok := x.GetUnk2700_AFHAGFONBFM().(*Unk2700_DMPIJLBHEAE_BundleInfo); ok {
		return x.BundleInfo
	}
	return nil
}

func (x *Unk2700_DMPIJLBHEAE) GetScoreChallengeInfo() *ScoreChallengeInfo {
	if x, ok := x.GetUnk2700_AFHAGFONBFM().(*Unk2700_DMPIJLBHEAE_ScoreChallengeInfo); ok {
		return x.ScoreChallengeInfo
	}
	return nil
}

func (x *Unk2700_DMPIJLBHEAE) GetBossChallengeId() uint32 {
	if x, ok := x.GetUnk2700_AFHAGFONBFM().(*Unk2700_DMPIJLBHEAE_BossChallengeId); ok {
		return x.BossChallengeId
	}
	return 0
}

type isUnk2700_DMPIJLBHEAE_Unk2700_AFHAGFONBFM interface {
	isUnk2700_DMPIJLBHEAE_Unk2700_AFHAGFONBFM()
}

type Unk2700_DMPIJLBHEAE_BundleInfo struct {
	BundleInfo *BundleInfo `protobuf:"bytes,11,opt,name=bundle_info,json=bundleInfo,proto3,oneof"`
}

type Unk2700_DMPIJLBHEAE_ScoreChallengeInfo struct {
	ScoreChallengeInfo *ScoreChallengeInfo `protobuf:"bytes,13,opt,name=score_challenge_info,json=scoreChallengeInfo,proto3,oneof"`
}

type Unk2700_DMPIJLBHEAE_BossChallengeId struct {
	BossChallengeId uint32 `protobuf:"varint,2,opt,name=boss_challenge_id,json=bossChallengeId,proto3,oneof"`
}

func (*Unk2700_DMPIJLBHEAE_BundleInfo) isUnk2700_DMPIJLBHEAE_Unk2700_AFHAGFONBFM() {}

func (*Unk2700_DMPIJLBHEAE_ScoreChallengeInfo) isUnk2700_DMPIJLBHEAE_Unk2700_AFHAGFONBFM() {}

func (*Unk2700_DMPIJLBHEAE_BossChallengeId) isUnk2700_DMPIJLBHEAE_Unk2700_AFHAGFONBFM() {}

var File_Unk2700_DMPIJLBHEAE_proto protoreflect.FileDescriptor

var file_Unk2700_DMPIJLBHEAE_proto_rawDesc = []byte{
	0x0a, 0x19, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x44, 0x4d, 0x50, 0x49, 0x4a, 0x4c,
	0x42, 0x48, 0x45, 0x41, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x42, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x5f, 0x44, 0x4d, 0x50, 0x49, 0x4a, 0x4c, 0x42, 0x48, 0x45, 0x41, 0x45, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x75, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x55, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x2e, 0x0a, 0x0b, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x42, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x0a, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x47, 0x0a, 0x14, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x12, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x43,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x11,
	0x62, 0x6f, 0x73, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0f, 0x62, 0x6f, 0x73, 0x73, 0x43,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x42, 0x15, 0x0a, 0x13, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x41, 0x46, 0x48, 0x41, 0x47, 0x46, 0x4f, 0x4e, 0x42, 0x46,
	0x4d, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_Unk2700_DMPIJLBHEAE_proto_rawDescOnce sync.Once
	file_Unk2700_DMPIJLBHEAE_proto_rawDescData = file_Unk2700_DMPIJLBHEAE_proto_rawDesc
)

func file_Unk2700_DMPIJLBHEAE_proto_rawDescGZIP() []byte {
	file_Unk2700_DMPIJLBHEAE_proto_rawDescOnce.Do(func() {
		file_Unk2700_DMPIJLBHEAE_proto_rawDescData = protoimpl.X.CompressGZIP(file_Unk2700_DMPIJLBHEAE_proto_rawDescData)
	})
	return file_Unk2700_DMPIJLBHEAE_proto_rawDescData
}

var file_Unk2700_DMPIJLBHEAE_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Unk2700_DMPIJLBHEAE_proto_goTypes = []interface{}{
	(*Unk2700_DMPIJLBHEAE)(nil), // 0: Unk2700_DMPIJLBHEAE
	(*BundleInfo)(nil),          // 1: BundleInfo
	(*ScoreChallengeInfo)(nil),  // 2: ScoreChallengeInfo
}
var file_Unk2700_DMPIJLBHEAE_proto_depIdxs = []int32{
	1, // 0: Unk2700_DMPIJLBHEAE.bundle_info:type_name -> BundleInfo
	2, // 1: Unk2700_DMPIJLBHEAE.score_challenge_info:type_name -> ScoreChallengeInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Unk2700_DMPIJLBHEAE_proto_init() }
func file_Unk2700_DMPIJLBHEAE_proto_init() {
	if File_Unk2700_DMPIJLBHEAE_proto != nil {
		return
	}
	file_BundleInfo_proto_init()
	file_ScoreChallengeInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Unk2700_DMPIJLBHEAE_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unk2700_DMPIJLBHEAE); i {
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
	file_Unk2700_DMPIJLBHEAE_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Unk2700_DMPIJLBHEAE_BundleInfo)(nil),
		(*Unk2700_DMPIJLBHEAE_ScoreChallengeInfo)(nil),
		(*Unk2700_DMPIJLBHEAE_BossChallengeId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Unk2700_DMPIJLBHEAE_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Unk2700_DMPIJLBHEAE_proto_goTypes,
		DependencyIndexes: file_Unk2700_DMPIJLBHEAE_proto_depIdxs,
		MessageInfos:      file_Unk2700_DMPIJLBHEAE_proto_msgTypes,
	}.Build()
	File_Unk2700_DMPIJLBHEAE_proto = out.File
	file_Unk2700_DMPIJLBHEAE_proto_rawDesc = nil
	file_Unk2700_DMPIJLBHEAE_proto_goTypes = nil
	file_Unk2700_DMPIJLBHEAE_proto_depIdxs = nil
}
