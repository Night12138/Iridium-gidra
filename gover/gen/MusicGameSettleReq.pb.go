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
// source: MusicGameSettleReq.proto

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

// CmdId: 8892
// EnetChannelId: 0
// EnetIsReliable: true
// IsAllowClient: true
type MusicGameSettleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unk2700_GDPKOANEDEB []uint32 `protobuf:"varint,384,rep,packed,name=Unk2700_GDPKOANEDEB,json=Unk2700GDPKOANEDEB,proto3" json:"Unk2700_GDPKOANEDEB,omitempty"`
	Unk2700_NMHGADLANMM uint32   `protobuf:"varint,795,opt,name=Unk2700_NMHGADLANMM,json=Unk2700NMHGADLANMM,proto3" json:"Unk2700_NMHGADLANMM,omitempty"`
	Unk2700_NNHGOCJLKFH []uint32 `protobuf:"varint,4,rep,packed,name=Unk2700_NNHGOCJLKFH,json=Unk2700NNHGOCJLKFH,proto3" json:"Unk2700_NNHGOCJLKFH,omitempty"`
	Unk2700_NCHHEJNFECG uint32   `protobuf:"varint,15,opt,name=Unk2700_NCHHEJNFECG,json=Unk2700NCHHEJNFECG,proto3" json:"Unk2700_NCHHEJNFECG,omitempty"`
	Score               uint32   `protobuf:"varint,9,opt,name=score,proto3" json:"score,omitempty"`
	Unk2700_CEPGMKAHHCD uint64   `protobuf:"varint,6,opt,name=Unk2700_CEPGMKAHHCD,json=Unk2700CEPGMKAHHCD,proto3" json:"Unk2700_CEPGMKAHHCD,omitempty"`
	Unk2700_MMHHGALFHGA uint32   `protobuf:"varint,13,opt,name=Unk2700_MMHHGALFHGA,json=Unk2700MMHHGALFHGA,proto3" json:"Unk2700_MMHHGALFHGA,omitempty"`
	Unk2700_CBLIJHDFKHA bool     `protobuf:"varint,422,opt,name=Unk2700_CBLIJHDFKHA,json=Unk2700CBLIJHDFKHA,proto3" json:"Unk2700_CBLIJHDFKHA,omitempty"`
	MaxCombo            uint32   `protobuf:"varint,5,opt,name=max_combo,json=maxCombo,proto3" json:"max_combo,omitempty"`
	Unk2700_FBLBGPHMLAL uint32   `protobuf:"varint,1058,opt,name=Unk2700_FBLBGPHMLAL,json=Unk2700FBLBGPHMLAL,proto3" json:"Unk2700_FBLBGPHMLAL,omitempty"`
	Speed               float32  `protobuf:"fixed32,409,opt,name=speed,proto3" json:"speed,omitempty"`
	Unk2700_IOKPIKJDEHG bool     `protobuf:"varint,3,opt,name=Unk2700_IOKPIKJDEHG,json=Unk2700IOKPIKJDEHG,proto3" json:"Unk2700_IOKPIKJDEHG,omitempty"`
	Combo               uint32   `protobuf:"varint,1,opt,name=combo,proto3" json:"combo,omitempty"`
	MusicBasicId        uint32   `protobuf:"varint,7,opt,name=music_basic_id,json=musicBasicId,proto3" json:"music_basic_id,omitempty"`
	Unk2700_DIMBABOGNEM uint32   `protobuf:"varint,2,opt,name=Unk2700_DIMBABOGNEM,json=Unk2700DIMBABOGNEM,proto3" json:"Unk2700_DIMBABOGNEM,omitempty"`
	Unk2700_IOMOHEKJJLJ uint32   `protobuf:"varint,1953,opt,name=Unk2700_IOMOHEKJJLJ,json=Unk2700IOMOHEKJJLJ,proto3" json:"Unk2700_IOMOHEKJJLJ,omitempty"`
	CorrectHit          uint32   `protobuf:"varint,14,opt,name=correct_hit,json=correctHit,proto3" json:"correct_hit,omitempty"`
	Unk2700_LKJNLDJAGGL bool     `protobuf:"varint,1285,opt,name=Unk2700_LKJNLDJAGGL,json=Unk2700LKJNLDJAGGL,proto3" json:"Unk2700_LKJNLDJAGGL,omitempty"`
}

func (x *MusicGameSettleReq) Reset() {
	*x = MusicGameSettleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MusicGameSettleReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MusicGameSettleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MusicGameSettleReq) ProtoMessage() {}

func (x *MusicGameSettleReq) ProtoReflect() protoreflect.Message {
	mi := &file_MusicGameSettleReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MusicGameSettleReq.ProtoReflect.Descriptor instead.
func (*MusicGameSettleReq) Descriptor() ([]byte, []int) {
	return file_MusicGameSettleReq_proto_rawDescGZIP(), []int{0}
}

func (x *MusicGameSettleReq) GetUnk2700_GDPKOANEDEB() []uint32 {
	if x != nil {
		return x.Unk2700_GDPKOANEDEB
	}
	return nil
}

func (x *MusicGameSettleReq) GetUnk2700_NMHGADLANMM() uint32 {
	if x != nil {
		return x.Unk2700_NMHGADLANMM
	}
	return 0
}

func (x *MusicGameSettleReq) GetUnk2700_NNHGOCJLKFH() []uint32 {
	if x != nil {
		return x.Unk2700_NNHGOCJLKFH
	}
	return nil
}

func (x *MusicGameSettleReq) GetUnk2700_NCHHEJNFECG() uint32 {
	if x != nil {
		return x.Unk2700_NCHHEJNFECG
	}
	return 0
}

func (x *MusicGameSettleReq) GetScore() uint32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *MusicGameSettleReq) GetUnk2700_CEPGMKAHHCD() uint64 {
	if x != nil {
		return x.Unk2700_CEPGMKAHHCD
	}
	return 0
}

func (x *MusicGameSettleReq) GetUnk2700_MMHHGALFHGA() uint32 {
	if x != nil {
		return x.Unk2700_MMHHGALFHGA
	}
	return 0
}

func (x *MusicGameSettleReq) GetUnk2700_CBLIJHDFKHA() bool {
	if x != nil {
		return x.Unk2700_CBLIJHDFKHA
	}
	return false
}

func (x *MusicGameSettleReq) GetMaxCombo() uint32 {
	if x != nil {
		return x.MaxCombo
	}
	return 0
}

func (x *MusicGameSettleReq) GetUnk2700_FBLBGPHMLAL() uint32 {
	if x != nil {
		return x.Unk2700_FBLBGPHMLAL
	}
	return 0
}

func (x *MusicGameSettleReq) GetSpeed() float32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *MusicGameSettleReq) GetUnk2700_IOKPIKJDEHG() bool {
	if x != nil {
		return x.Unk2700_IOKPIKJDEHG
	}
	return false
}

func (x *MusicGameSettleReq) GetCombo() uint32 {
	if x != nil {
		return x.Combo
	}
	return 0
}

func (x *MusicGameSettleReq) GetMusicBasicId() uint32 {
	if x != nil {
		return x.MusicBasicId
	}
	return 0
}

func (x *MusicGameSettleReq) GetUnk2700_DIMBABOGNEM() uint32 {
	if x != nil {
		return x.Unk2700_DIMBABOGNEM
	}
	return 0
}

func (x *MusicGameSettleReq) GetUnk2700_IOMOHEKJJLJ() uint32 {
	if x != nil {
		return x.Unk2700_IOMOHEKJJLJ
	}
	return 0
}

func (x *MusicGameSettleReq) GetCorrectHit() uint32 {
	if x != nil {
		return x.CorrectHit
	}
	return 0
}

func (x *MusicGameSettleReq) GetUnk2700_LKJNLDJAGGL() bool {
	if x != nil {
		return x.Unk2700_LKJNLDJAGGL
	}
	return false
}

var File_MusicGameSettleReq_proto protoreflect.FileDescriptor

var file_MusicGameSettleReq_proto_rawDesc = []byte{
	0x0a, 0x18, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x06, 0x0a, 0x12, 0x4d,
	0x75, 0x73, 0x69, 0x63, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x30, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x47, 0x44, 0x50,
	0x4b, 0x4f, 0x41, 0x4e, 0x45, 0x44, 0x45, 0x42, 0x18, 0x80, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x47, 0x44, 0x50, 0x4b, 0x4f, 0x41, 0x4e, 0x45,
	0x44, 0x45, 0x42, 0x12, 0x30, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4e,
	0x4d, 0x48, 0x47, 0x41, 0x44, 0x4c, 0x41, 0x4e, 0x4d, 0x4d, 0x18, 0x9b, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4e, 0x4d, 0x48, 0x47, 0x41, 0x44,
	0x4c, 0x41, 0x4e, 0x4d, 0x4d, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x4e, 0x4e, 0x48, 0x47, 0x4f, 0x43, 0x4a, 0x4c, 0x4b, 0x46, 0x48, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4e, 0x4e, 0x48, 0x47, 0x4f,
	0x43, 0x4a, 0x4c, 0x4b, 0x46, 0x48, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30,
	0x30, 0x5f, 0x4e, 0x43, 0x48, 0x48, 0x45, 0x4a, 0x4e, 0x46, 0x45, 0x43, 0x47, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4e, 0x43, 0x48, 0x48,
	0x45, 0x4a, 0x4e, 0x46, 0x45, 0x43, 0x47, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2f, 0x0a,
	0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x43, 0x45, 0x50, 0x47, 0x4d, 0x4b, 0x41,
	0x48, 0x48, 0x43, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x43, 0x45, 0x50, 0x47, 0x4d, 0x4b, 0x41, 0x48, 0x48, 0x43, 0x44, 0x12, 0x2f,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4d, 0x4d, 0x48, 0x48, 0x47, 0x41,
	0x4c, 0x46, 0x48, 0x47, 0x41, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x4d, 0x4d, 0x48, 0x48, 0x47, 0x41, 0x4c, 0x46, 0x48, 0x47, 0x41, 0x12,
	0x30, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x43, 0x42, 0x4c, 0x49, 0x4a,
	0x48, 0x44, 0x46, 0x4b, 0x48, 0x41, 0x18, 0xa6, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x55,
	0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x43, 0x42, 0x4c, 0x49, 0x4a, 0x48, 0x44, 0x46, 0x4b, 0x48,
	0x41, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6d, 0x62, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6d, 0x62, 0x6f, 0x12, 0x30,
	0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x46, 0x42, 0x4c, 0x42, 0x47, 0x50,
	0x48, 0x4d, 0x4c, 0x41, 0x4c, 0x18, 0xa2, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e,
	0x6b, 0x32, 0x37, 0x30, 0x30, 0x46, 0x42, 0x4c, 0x42, 0x47, 0x50, 0x48, 0x4d, 0x4c, 0x41, 0x4c,
	0x12, 0x15, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x99, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37,
	0x30, 0x30, 0x5f, 0x49, 0x4f, 0x4b, 0x50, 0x49, 0x4b, 0x4a, 0x44, 0x45, 0x48, 0x47, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x49, 0x4f, 0x4b,
	0x50, 0x49, 0x4b, 0x4a, 0x44, 0x45, 0x48, 0x47, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6d, 0x62,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x6d, 0x62, 0x6f, 0x12, 0x24,
	0x0a, 0x0e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x5f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x44, 0x49, 0x4d, 0x42, 0x41, 0x42, 0x4f, 0x47, 0x4e, 0x45, 0x4d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x44, 0x49, 0x4d, 0x42, 0x41, 0x42,
	0x4f, 0x47, 0x4e, 0x45, 0x4d, 0x12, 0x30, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30,
	0x5f, 0x49, 0x4f, 0x4d, 0x4f, 0x48, 0x45, 0x4b, 0x4a, 0x4a, 0x4c, 0x4a, 0x18, 0xa1, 0x0f, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x49, 0x4f, 0x4d, 0x4f,
	0x48, 0x45, 0x4b, 0x4a, 0x4a, 0x4c, 0x4a, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x63, 0x74, 0x5f, 0x68, 0x69, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x63, 0x74, 0x48, 0x69, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x32,
	0x37, 0x30, 0x30, 0x5f, 0x4c, 0x4b, 0x4a, 0x4e, 0x4c, 0x44, 0x4a, 0x41, 0x47, 0x47, 0x4c, 0x18,
	0x85, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x4c,
	0x4b, 0x4a, 0x4e, 0x4c, 0x44, 0x4a, 0x41, 0x47, 0x47, 0x4c, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MusicGameSettleReq_proto_rawDescOnce sync.Once
	file_MusicGameSettleReq_proto_rawDescData = file_MusicGameSettleReq_proto_rawDesc
)

func file_MusicGameSettleReq_proto_rawDescGZIP() []byte {
	file_MusicGameSettleReq_proto_rawDescOnce.Do(func() {
		file_MusicGameSettleReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_MusicGameSettleReq_proto_rawDescData)
	})
	return file_MusicGameSettleReq_proto_rawDescData
}

var file_MusicGameSettleReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MusicGameSettleReq_proto_goTypes = []interface{}{
	(*MusicGameSettleReq)(nil), // 0: MusicGameSettleReq
}
var file_MusicGameSettleReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MusicGameSettleReq_proto_init() }
func file_MusicGameSettleReq_proto_init() {
	if File_MusicGameSettleReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MusicGameSettleReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MusicGameSettleReq); i {
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
			RawDescriptor: file_MusicGameSettleReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MusicGameSettleReq_proto_goTypes,
		DependencyIndexes: file_MusicGameSettleReq_proto_depIdxs,
		MessageInfos:      file_MusicGameSettleReq_proto_msgTypes,
	}.Build()
	File_MusicGameSettleReq_proto = out.File
	file_MusicGameSettleReq_proto_rawDesc = nil
	file_MusicGameSettleReq_proto_goTypes = nil
	file_MusicGameSettleReq_proto_depIdxs = nil
}
