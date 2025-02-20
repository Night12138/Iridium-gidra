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
// source: AbilityInvokeArgument.proto

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

type AbilityInvokeArgument int32

const (
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_NONE                                 AbilityInvokeArgument = 0
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_MODIFIER_CHANGE                 AbilityInvokeArgument = 1
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_COMMAND_MODIFIER_CHANGE_REQUEST AbilityInvokeArgument = 2
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_SPECIAL_FLOAT_ARGUMENT          AbilityInvokeArgument = 3
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_OVERRIDE_PARAM                  AbilityInvokeArgument = 4
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_CLEAR_OVERRIDE_PARAM            AbilityInvokeArgument = 5
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_REINIT_OVERRIDEMAP              AbilityInvokeArgument = 6
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_GLOBAL_FLOAT_VALUE              AbilityInvokeArgument = 7
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_CLEAR_GLOBAL_FLOAT_VALUE        AbilityInvokeArgument = 8
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_ABILITY_ELEMENT_STRENGTH        AbilityInvokeArgument = 9
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_ADD_OR_GET_ABILITY_AND_TRIGGER  AbilityInvokeArgument = 10
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_SET_KILLED_STATE                AbilityInvokeArgument = 11
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_SET_ABILITY_TRIGGER             AbilityInvokeArgument = 12
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_ADD_NEW_ABILITY                 AbilityInvokeArgument = 13
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_REMOVE_ABILITY                  AbilityInvokeArgument = 14
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_SET_MODIFIER_APPLY_ENTITY       AbilityInvokeArgument = 15
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_MODIFIER_DURABILITY_CHANGE      AbilityInvokeArgument = 16
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_ELEMENT_REACTION_VISUAL         AbilityInvokeArgument = 17
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_SET_POSE_PARAMETER              AbilityInvokeArgument = 18
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_UPDATE_BASE_REACTION_DAMAGE     AbilityInvokeArgument = 19
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_TRIGGER_ELEMENT_REACTION        AbilityInvokeArgument = 20
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_META_LOSE_HP                         AbilityInvokeArgument = 21
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_Unk2700_JDDDLJELBLJ                  AbilityInvokeArgument = 22
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_ACTION_TRIGGER_ABILITY               AbilityInvokeArgument = 50
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_ACTION_SET_CRASH_DAMAGE              AbilityInvokeArgument = 51
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_ACTION_EFFECT                        AbilityInvokeArgument = 52
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_ACTION_SUMMON                        AbilityInvokeArgument = 53
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_ACTION_BLINK                         AbilityInvokeArgument = 54
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_ACTION_CREATE_GADGET                 AbilityInvokeArgument = 55
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_ACTION_APPLY_LEVEL_MODIFIER          AbilityInvokeArgument = 56
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_ACTION_GENERATE_ELEM_BALL            AbilityInvokeArgument = 57
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_ACTION_SET_RANDOM_OVERRIDE_MAP_VALUE AbilityInvokeArgument = 58
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_ACTION_SERVER_MONSTER_LOG            AbilityInvokeArgument = 59
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_ACTION_CREATE_TILE                   AbilityInvokeArgument = 60
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_ACTION_DESTROY_TILE                  AbilityInvokeArgument = 61
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_ACTION_FIRE_AFTER_IMAGE              AbilityInvokeArgument = 62
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_Unk2700_FNANDDPDLOL                  AbilityInvokeArgument = 63
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_Unk3000_EEANHJONEEP                  AbilityInvokeArgument = 64
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_Unk3000_ADEHJMKKBJD                  AbilityInvokeArgument = 65
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_MIXIN_AVATAR_STEER_BY_CAMERA         AbilityInvokeArgument = 100
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_MIXIN_MONSTER_DEFEND                 AbilityInvokeArgument = 101
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_MIXIN_WIND_ZONE                      AbilityInvokeArgument = 102
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_MIXIN_COST_STAMINA                   AbilityInvokeArgument = 103
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_MIXIN_ELITE_SHIELD                   AbilityInvokeArgument = 104
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_MIXIN_ELEMENT_SHIELD                 AbilityInvokeArgument = 105
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_MIXIN_GLOBAL_SHIELD                  AbilityInvokeArgument = 106
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_MIXIN_SHIELD_BAR                     AbilityInvokeArgument = 107
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_MIXIN_WIND_SEED_SPAWNER              AbilityInvokeArgument = 108
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_MIXIN_DO_ACTION_BY_ELEMENT_REACTION  AbilityInvokeArgument = 109
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_MIXIN_FIELD_ENTITY_COUNT_CHANGE      AbilityInvokeArgument = 110
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_MIXIN_SCENE_PROP_SYNC                AbilityInvokeArgument = 111
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_MIXIN_WIDGET_MP_SUPPORT              AbilityInvokeArgument = 112
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_Unk2700_NJHBFADEOON                  AbilityInvokeArgument = 113
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_Unk2700_EGCIFFFLLBG                  AbilityInvokeArgument = 114
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_Unk2700_OFDGFACOLDI                  AbilityInvokeArgument = 115
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_Unk2700_KDPKJGJNGFB                  AbilityInvokeArgument = 116
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_Unk3000_BNECPACGKHP                  AbilityInvokeArgument = 117
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_Unk3000_LGIPOCBHKAL                  AbilityInvokeArgument = 118
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_Unk3000_EFJIGCEGHJG                  AbilityInvokeArgument = 119
	AbilityInvokeArgument_ABILITY_INVOKE_ARGUMENT_Unk3100_NLIPBBOINEO                  AbilityInvokeArgument = 120
)

// Enum value maps for AbilityInvokeArgument.
var (
	AbilityInvokeArgument_name = map[int32]string{
		0:   "ABILITY_INVOKE_ARGUMENT_NONE",
		1:   "ABILITY_INVOKE_ARGUMENT_META_MODIFIER_CHANGE",
		2:   "ABILITY_INVOKE_ARGUMENT_META_COMMAND_MODIFIER_CHANGE_REQUEST",
		3:   "ABILITY_INVOKE_ARGUMENT_META_SPECIAL_FLOAT_ARGUMENT",
		4:   "ABILITY_INVOKE_ARGUMENT_META_OVERRIDE_PARAM",
		5:   "ABILITY_INVOKE_ARGUMENT_META_CLEAR_OVERRIDE_PARAM",
		6:   "ABILITY_INVOKE_ARGUMENT_META_REINIT_OVERRIDEMAP",
		7:   "ABILITY_INVOKE_ARGUMENT_META_GLOBAL_FLOAT_VALUE",
		8:   "ABILITY_INVOKE_ARGUMENT_META_CLEAR_GLOBAL_FLOAT_VALUE",
		9:   "ABILITY_INVOKE_ARGUMENT_META_ABILITY_ELEMENT_STRENGTH",
		10:  "ABILITY_INVOKE_ARGUMENT_META_ADD_OR_GET_ABILITY_AND_TRIGGER",
		11:  "ABILITY_INVOKE_ARGUMENT_META_SET_KILLED_STATE",
		12:  "ABILITY_INVOKE_ARGUMENT_META_SET_ABILITY_TRIGGER",
		13:  "ABILITY_INVOKE_ARGUMENT_META_ADD_NEW_ABILITY",
		14:  "ABILITY_INVOKE_ARGUMENT_META_REMOVE_ABILITY",
		15:  "ABILITY_INVOKE_ARGUMENT_META_SET_MODIFIER_APPLY_ENTITY",
		16:  "ABILITY_INVOKE_ARGUMENT_META_MODIFIER_DURABILITY_CHANGE",
		17:  "ABILITY_INVOKE_ARGUMENT_META_ELEMENT_REACTION_VISUAL",
		18:  "ABILITY_INVOKE_ARGUMENT_META_SET_POSE_PARAMETER",
		19:  "ABILITY_INVOKE_ARGUMENT_META_UPDATE_BASE_REACTION_DAMAGE",
		20:  "ABILITY_INVOKE_ARGUMENT_META_TRIGGER_ELEMENT_REACTION",
		21:  "ABILITY_INVOKE_ARGUMENT_META_LOSE_HP",
		22:  "ABILITY_INVOKE_ARGUMENT_Unk2700_JDDDLJELBLJ",
		50:  "ABILITY_INVOKE_ARGUMENT_ACTION_TRIGGER_ABILITY",
		51:  "ABILITY_INVOKE_ARGUMENT_ACTION_SET_CRASH_DAMAGE",
		52:  "ABILITY_INVOKE_ARGUMENT_ACTION_EFFECT",
		53:  "ABILITY_INVOKE_ARGUMENT_ACTION_SUMMON",
		54:  "ABILITY_INVOKE_ARGUMENT_ACTION_BLINK",
		55:  "ABILITY_INVOKE_ARGUMENT_ACTION_CREATE_GADGET",
		56:  "ABILITY_INVOKE_ARGUMENT_ACTION_APPLY_LEVEL_MODIFIER",
		57:  "ABILITY_INVOKE_ARGUMENT_ACTION_GENERATE_ELEM_BALL",
		58:  "ABILITY_INVOKE_ARGUMENT_ACTION_SET_RANDOM_OVERRIDE_MAP_VALUE",
		59:  "ABILITY_INVOKE_ARGUMENT_ACTION_SERVER_MONSTER_LOG",
		60:  "ABILITY_INVOKE_ARGUMENT_ACTION_CREATE_TILE",
		61:  "ABILITY_INVOKE_ARGUMENT_ACTION_DESTROY_TILE",
		62:  "ABILITY_INVOKE_ARGUMENT_ACTION_FIRE_AFTER_IMAGE",
		63:  "ABILITY_INVOKE_ARGUMENT_Unk2700_FNANDDPDLOL",
		64:  "ABILITY_INVOKE_ARGUMENT_Unk3000_EEANHJONEEP",
		65:  "ABILITY_INVOKE_ARGUMENT_Unk3000_ADEHJMKKBJD",
		100: "ABILITY_INVOKE_ARGUMENT_MIXIN_AVATAR_STEER_BY_CAMERA",
		101: "ABILITY_INVOKE_ARGUMENT_MIXIN_MONSTER_DEFEND",
		102: "ABILITY_INVOKE_ARGUMENT_MIXIN_WIND_ZONE",
		103: "ABILITY_INVOKE_ARGUMENT_MIXIN_COST_STAMINA",
		104: "ABILITY_INVOKE_ARGUMENT_MIXIN_ELITE_SHIELD",
		105: "ABILITY_INVOKE_ARGUMENT_MIXIN_ELEMENT_SHIELD",
		106: "ABILITY_INVOKE_ARGUMENT_MIXIN_GLOBAL_SHIELD",
		107: "ABILITY_INVOKE_ARGUMENT_MIXIN_SHIELD_BAR",
		108: "ABILITY_INVOKE_ARGUMENT_MIXIN_WIND_SEED_SPAWNER",
		109: "ABILITY_INVOKE_ARGUMENT_MIXIN_DO_ACTION_BY_ELEMENT_REACTION",
		110: "ABILITY_INVOKE_ARGUMENT_MIXIN_FIELD_ENTITY_COUNT_CHANGE",
		111: "ABILITY_INVOKE_ARGUMENT_MIXIN_SCENE_PROP_SYNC",
		112: "ABILITY_INVOKE_ARGUMENT_MIXIN_WIDGET_MP_SUPPORT",
		113: "ABILITY_INVOKE_ARGUMENT_Unk2700_NJHBFADEOON",
		114: "ABILITY_INVOKE_ARGUMENT_Unk2700_EGCIFFFLLBG",
		115: "ABILITY_INVOKE_ARGUMENT_Unk2700_OFDGFACOLDI",
		116: "ABILITY_INVOKE_ARGUMENT_Unk2700_KDPKJGJNGFB",
		117: "ABILITY_INVOKE_ARGUMENT_Unk3000_BNECPACGKHP",
		118: "ABILITY_INVOKE_ARGUMENT_Unk3000_LGIPOCBHKAL",
		119: "ABILITY_INVOKE_ARGUMENT_Unk3000_EFJIGCEGHJG",
		120: "ABILITY_INVOKE_ARGUMENT_Unk3100_NLIPBBOINEO",
	}
	AbilityInvokeArgument_value = map[string]int32{
		"ABILITY_INVOKE_ARGUMENT_NONE":                                 0,
		"ABILITY_INVOKE_ARGUMENT_META_MODIFIER_CHANGE":                 1,
		"ABILITY_INVOKE_ARGUMENT_META_COMMAND_MODIFIER_CHANGE_REQUEST": 2,
		"ABILITY_INVOKE_ARGUMENT_META_SPECIAL_FLOAT_ARGUMENT":          3,
		"ABILITY_INVOKE_ARGUMENT_META_OVERRIDE_PARAM":                  4,
		"ABILITY_INVOKE_ARGUMENT_META_CLEAR_OVERRIDE_PARAM":            5,
		"ABILITY_INVOKE_ARGUMENT_META_REINIT_OVERRIDEMAP":              6,
		"ABILITY_INVOKE_ARGUMENT_META_GLOBAL_FLOAT_VALUE":              7,
		"ABILITY_INVOKE_ARGUMENT_META_CLEAR_GLOBAL_FLOAT_VALUE":        8,
		"ABILITY_INVOKE_ARGUMENT_META_ABILITY_ELEMENT_STRENGTH":        9,
		"ABILITY_INVOKE_ARGUMENT_META_ADD_OR_GET_ABILITY_AND_TRIGGER":  10,
		"ABILITY_INVOKE_ARGUMENT_META_SET_KILLED_STATE":                11,
		"ABILITY_INVOKE_ARGUMENT_META_SET_ABILITY_TRIGGER":             12,
		"ABILITY_INVOKE_ARGUMENT_META_ADD_NEW_ABILITY":                 13,
		"ABILITY_INVOKE_ARGUMENT_META_REMOVE_ABILITY":                  14,
		"ABILITY_INVOKE_ARGUMENT_META_SET_MODIFIER_APPLY_ENTITY":       15,
		"ABILITY_INVOKE_ARGUMENT_META_MODIFIER_DURABILITY_CHANGE":      16,
		"ABILITY_INVOKE_ARGUMENT_META_ELEMENT_REACTION_VISUAL":         17,
		"ABILITY_INVOKE_ARGUMENT_META_SET_POSE_PARAMETER":              18,
		"ABILITY_INVOKE_ARGUMENT_META_UPDATE_BASE_REACTION_DAMAGE":     19,
		"ABILITY_INVOKE_ARGUMENT_META_TRIGGER_ELEMENT_REACTION":        20,
		"ABILITY_INVOKE_ARGUMENT_META_LOSE_HP":                         21,
		"ABILITY_INVOKE_ARGUMENT_Unk2700_JDDDLJELBLJ":                  22,
		"ABILITY_INVOKE_ARGUMENT_ACTION_TRIGGER_ABILITY":               50,
		"ABILITY_INVOKE_ARGUMENT_ACTION_SET_CRASH_DAMAGE":              51,
		"ABILITY_INVOKE_ARGUMENT_ACTION_EFFECT":                        52,
		"ABILITY_INVOKE_ARGUMENT_ACTION_SUMMON":                        53,
		"ABILITY_INVOKE_ARGUMENT_ACTION_BLINK":                         54,
		"ABILITY_INVOKE_ARGUMENT_ACTION_CREATE_GADGET":                 55,
		"ABILITY_INVOKE_ARGUMENT_ACTION_APPLY_LEVEL_MODIFIER":          56,
		"ABILITY_INVOKE_ARGUMENT_ACTION_GENERATE_ELEM_BALL":            57,
		"ABILITY_INVOKE_ARGUMENT_ACTION_SET_RANDOM_OVERRIDE_MAP_VALUE": 58,
		"ABILITY_INVOKE_ARGUMENT_ACTION_SERVER_MONSTER_LOG":            59,
		"ABILITY_INVOKE_ARGUMENT_ACTION_CREATE_TILE":                   60,
		"ABILITY_INVOKE_ARGUMENT_ACTION_DESTROY_TILE":                  61,
		"ABILITY_INVOKE_ARGUMENT_ACTION_FIRE_AFTER_IMAGE":              62,
		"ABILITY_INVOKE_ARGUMENT_Unk2700_FNANDDPDLOL":                  63,
		"ABILITY_INVOKE_ARGUMENT_Unk3000_EEANHJONEEP":                  64,
		"ABILITY_INVOKE_ARGUMENT_Unk3000_ADEHJMKKBJD":                  65,
		"ABILITY_INVOKE_ARGUMENT_MIXIN_AVATAR_STEER_BY_CAMERA":         100,
		"ABILITY_INVOKE_ARGUMENT_MIXIN_MONSTER_DEFEND":                 101,
		"ABILITY_INVOKE_ARGUMENT_MIXIN_WIND_ZONE":                      102,
		"ABILITY_INVOKE_ARGUMENT_MIXIN_COST_STAMINA":                   103,
		"ABILITY_INVOKE_ARGUMENT_MIXIN_ELITE_SHIELD":                   104,
		"ABILITY_INVOKE_ARGUMENT_MIXIN_ELEMENT_SHIELD":                 105,
		"ABILITY_INVOKE_ARGUMENT_MIXIN_GLOBAL_SHIELD":                  106,
		"ABILITY_INVOKE_ARGUMENT_MIXIN_SHIELD_BAR":                     107,
		"ABILITY_INVOKE_ARGUMENT_MIXIN_WIND_SEED_SPAWNER":              108,
		"ABILITY_INVOKE_ARGUMENT_MIXIN_DO_ACTION_BY_ELEMENT_REACTION":  109,
		"ABILITY_INVOKE_ARGUMENT_MIXIN_FIELD_ENTITY_COUNT_CHANGE":      110,
		"ABILITY_INVOKE_ARGUMENT_MIXIN_SCENE_PROP_SYNC":                111,
		"ABILITY_INVOKE_ARGUMENT_MIXIN_WIDGET_MP_SUPPORT":              112,
		"ABILITY_INVOKE_ARGUMENT_Unk2700_NJHBFADEOON":                  113,
		"ABILITY_INVOKE_ARGUMENT_Unk2700_EGCIFFFLLBG":                  114,
		"ABILITY_INVOKE_ARGUMENT_Unk2700_OFDGFACOLDI":                  115,
		"ABILITY_INVOKE_ARGUMENT_Unk2700_KDPKJGJNGFB":                  116,
		"ABILITY_INVOKE_ARGUMENT_Unk3000_BNECPACGKHP":                  117,
		"ABILITY_INVOKE_ARGUMENT_Unk3000_LGIPOCBHKAL":                  118,
		"ABILITY_INVOKE_ARGUMENT_Unk3000_EFJIGCEGHJG":                  119,
		"ABILITY_INVOKE_ARGUMENT_Unk3100_NLIPBBOINEO":                  120,
	}
)

func (x AbilityInvokeArgument) Enum() *AbilityInvokeArgument {
	p := new(AbilityInvokeArgument)
	*p = x
	return p
}

func (x AbilityInvokeArgument) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AbilityInvokeArgument) Descriptor() protoreflect.EnumDescriptor {
	return file_AbilityInvokeArgument_proto_enumTypes[0].Descriptor()
}

func (AbilityInvokeArgument) Type() protoreflect.EnumType {
	return &file_AbilityInvokeArgument_proto_enumTypes[0]
}

func (x AbilityInvokeArgument) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AbilityInvokeArgument.Descriptor instead.
func (AbilityInvokeArgument) EnumDescriptor() ([]byte, []int) {
	return file_AbilityInvokeArgument_proto_rawDescGZIP(), []int{0}
}

var File_AbilityInvokeArgument_proto protoreflect.FileDescriptor

var file_AbilityInvokeArgument_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x41,
	0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xd1, 0x18,
	0x0a, 0x15, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x41,
	0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x42, 0x49, 0x4c, 0x49,
	0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x30, 0x0a, 0x2c, 0x41, 0x42, 0x49,
	0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49,
	0x45, 0x52, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x01, 0x12, 0x40, 0x0a, 0x3c, 0x41,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52,
	0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x5f, 0x43, 0x4f, 0x4d, 0x4d,
	0x41, 0x4e, 0x44, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x02, 0x12, 0x37, 0x0a,
	0x33, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f,
	0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x5f, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x5f, 0x41, 0x52, 0x47, 0x55,
	0x4d, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x2f, 0x0a, 0x2b, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x52, 0x49, 0x44, 0x45, 0x5f,
	0x50, 0x41, 0x52, 0x41, 0x4d, 0x10, 0x04, 0x12, 0x35, 0x0a, 0x31, 0x41, 0x42, 0x49, 0x4c, 0x49,
	0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x5f, 0x43, 0x4c, 0x45, 0x41, 0x52, 0x5f, 0x4f, 0x56,
	0x45, 0x52, 0x52, 0x49, 0x44, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x10, 0x05, 0x12, 0x33,
	0x0a, 0x2f, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45,
	0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x5f, 0x52,
	0x45, 0x49, 0x4e, 0x49, 0x54, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x52, 0x49, 0x44, 0x45, 0x4d, 0x41,
	0x50, 0x10, 0x06, 0x12, 0x33, 0x0a, 0x2f, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49,
	0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d,
	0x45, 0x54, 0x41, 0x5f, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54,
	0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x07, 0x12, 0x39, 0x0a, 0x35, 0x41, 0x42, 0x49, 0x4c,
	0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x5f, 0x43, 0x4c, 0x45, 0x41, 0x52, 0x5f, 0x47,
	0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x5f, 0x56, 0x41, 0x4c, 0x55,
	0x45, 0x10, 0x08, 0x12, 0x39, 0x0a, 0x35, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49,
	0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d,
	0x45, 0x54, 0x41, 0x5f, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x45, 0x4c, 0x45, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x4e, 0x47, 0x54, 0x48, 0x10, 0x09, 0x12, 0x3f,
	0x0a, 0x3b, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45,
	0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x5f, 0x41,
	0x44, 0x44, 0x5f, 0x4f, 0x52, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x10, 0x0a, 0x12,
	0x31, 0x0a, 0x2d, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b,
	0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x5f,
	0x53, 0x45, 0x54, 0x5f, 0x4b, 0x49, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x10, 0x0b, 0x12, 0x34, 0x0a, 0x30, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e,
	0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45,
	0x54, 0x41, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x54,
	0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x10, 0x0c, 0x12, 0x30, 0x0a, 0x2c, 0x41, 0x42, 0x49, 0x4c,
	0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x4e, 0x45, 0x57,
	0x5f, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x0d, 0x12, 0x2f, 0x0a, 0x2b, 0x41, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47,
	0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56,
	0x45, 0x5f, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x0e, 0x12, 0x3a, 0x0a, 0x36, 0x41,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52,
	0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x5f, 0x53, 0x45, 0x54, 0x5f,
	0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x45,
	0x4e, 0x54, 0x49, 0x54, 0x59, 0x10, 0x0f, 0x12, 0x3b, 0x0a, 0x37, 0x41, 0x42, 0x49, 0x4c, 0x49,
	0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x45, 0x52,
	0x5f, 0x44, 0x55, 0x52, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x48, 0x41, 0x4e,
	0x47, 0x45, 0x10, 0x10, 0x12, 0x38, 0x0a, 0x34, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f,
	0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x4d, 0x45, 0x54, 0x41, 0x5f, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x49, 0x53, 0x55, 0x41, 0x4c, 0x10, 0x11, 0x12, 0x33,
	0x0a, 0x2f, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45,
	0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x5f, 0x53,
	0x45, 0x54, 0x5f, 0x50, 0x4f, 0x53, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45,
	0x52, 0x10, 0x12, 0x12, 0x3c, 0x0a, 0x38, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49,
	0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d,
	0x45, 0x54, 0x41, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f,
	0x52, 0x45, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x4d, 0x41, 0x47, 0x45, 0x10,
	0x13, 0x12, 0x39, 0x0a, 0x35, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56,
	0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54,
	0x41, 0x5f, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x5f, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x14, 0x12, 0x28, 0x0a, 0x24,
	0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41,
	0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x5f, 0x4c, 0x4f, 0x53,
	0x45, 0x5f, 0x48, 0x50, 0x10, 0x15, 0x12, 0x2f, 0x0a, 0x2b, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4a, 0x44, 0x44, 0x44, 0x4c, 0x4a,
	0x45, 0x4c, 0x42, 0x4c, 0x4a, 0x10, 0x16, 0x12, 0x32, 0x0a, 0x2e, 0x41, 0x42, 0x49, 0x4c, 0x49,
	0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45,
	0x52, 0x5f, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x32, 0x12, 0x33, 0x0a, 0x2f, 0x41,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52,
	0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45,
	0x54, 0x5f, 0x43, 0x52, 0x41, 0x53, 0x48, 0x5f, 0x44, 0x41, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x33,
	0x12, 0x29, 0x0a, 0x25, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f,
	0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x45, 0x46, 0x46, 0x45, 0x43, 0x54, 0x10, 0x34, 0x12, 0x29, 0x0a, 0x25, 0x41,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52,
	0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x55,
	0x4d, 0x4d, 0x4f, 0x4e, 0x10, 0x35, 0x12, 0x28, 0x0a, 0x24, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x36,
	0x12, 0x30, 0x0a, 0x2c, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f,
	0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x47, 0x41, 0x44, 0x47, 0x45, 0x54,
	0x10, 0x37, 0x12, 0x37, 0x0a, 0x33, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e,
	0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x59, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c,
	0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x45, 0x52, 0x10, 0x38, 0x12, 0x35, 0x0a, 0x31, 0x41,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52,
	0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x47, 0x45,
	0x4e, 0x45, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x4c, 0x45, 0x4d, 0x5f, 0x42, 0x41, 0x4c, 0x4c,
	0x10, 0x39, 0x12, 0x40, 0x0a, 0x3c, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e,
	0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x5f,
	0x4f, 0x56, 0x45, 0x52, 0x52, 0x49, 0x44, 0x45, 0x5f, 0x4d, 0x41, 0x50, 0x5f, 0x56, 0x41, 0x4c,
	0x55, 0x45, 0x10, 0x3a, 0x12, 0x35, 0x0a, 0x31, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f,
	0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x4d, 0x4f,
	0x4e, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x4c, 0x4f, 0x47, 0x10, 0x3b, 0x12, 0x2e, 0x0a, 0x2a, 0x41,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52,
	0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x49, 0x4c, 0x45, 0x10, 0x3c, 0x12, 0x2f, 0x0a, 0x2b, 0x41,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52,
	0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45,
	0x53, 0x54, 0x52, 0x4f, 0x59, 0x5f, 0x54, 0x49, 0x4c, 0x45, 0x10, 0x3d, 0x12, 0x33, 0x0a, 0x2f,
	0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41,
	0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46,
	0x49, 0x52, 0x45, 0x5f, 0x41, 0x46, 0x54, 0x45, 0x52, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10,
	0x3e, 0x12, 0x2f, 0x0a, 0x2b, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56,
	0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x6e, 0x6b,
	0x32, 0x37, 0x30, 0x30, 0x5f, 0x46, 0x4e, 0x41, 0x4e, 0x44, 0x44, 0x50, 0x44, 0x4c, 0x4f, 0x4c,
	0x10, 0x3f, 0x12, 0x2f, 0x0a, 0x2b, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e,
	0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x6e,
	0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x45, 0x45, 0x41, 0x4e, 0x48, 0x4a, 0x4f, 0x4e, 0x45, 0x45,
	0x50, 0x10, 0x40, 0x12, 0x2f, 0x0a, 0x2b, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49,
	0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55,
	0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30, 0x5f, 0x41, 0x44, 0x45, 0x48, 0x4a, 0x4d, 0x4b, 0x4b, 0x42,
	0x4a, 0x44, 0x10, 0x41, 0x12, 0x38, 0x0a, 0x34, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f,
	0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x4d, 0x49, 0x58, 0x49, 0x4e, 0x5f, 0x41, 0x56, 0x41, 0x54, 0x41, 0x52, 0x5f, 0x53, 0x54, 0x45,
	0x45, 0x52, 0x5f, 0x42, 0x59, 0x5f, 0x43, 0x41, 0x4d, 0x45, 0x52, 0x41, 0x10, 0x64, 0x12, 0x30,
	0x0a, 0x2c, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45,
	0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x58, 0x49, 0x4e, 0x5f,
	0x4d, 0x4f, 0x4e, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x46, 0x45, 0x4e, 0x44, 0x10, 0x65,
	0x12, 0x2b, 0x0a, 0x27, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f,
	0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x58, 0x49,
	0x4e, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x5f, 0x5a, 0x4f, 0x4e, 0x45, 0x10, 0x66, 0x12, 0x2e, 0x0a,
	0x2a, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f,
	0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x58, 0x49, 0x4e, 0x5f, 0x43,
	0x4f, 0x53, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x4d, 0x49, 0x4e, 0x41, 0x10, 0x67, 0x12, 0x2e, 0x0a,
	0x2a, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f,
	0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x58, 0x49, 0x4e, 0x5f, 0x45,
	0x4c, 0x49, 0x54, 0x45, 0x5f, 0x53, 0x48, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x68, 0x12, 0x30, 0x0a,
	0x2c, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f,
	0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x58, 0x49, 0x4e, 0x5f, 0x45,
	0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x48, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x69, 0x12,
	0x2f, 0x0a, 0x2b, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b,
	0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x58, 0x49, 0x4e,
	0x5f, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x5f, 0x53, 0x48, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x6a,
	0x12, 0x2c, 0x0a, 0x28, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f,
	0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x58, 0x49,
	0x4e, 0x5f, 0x53, 0x48, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x42, 0x41, 0x52, 0x10, 0x6b, 0x12, 0x33,
	0x0a, 0x2f, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45,
	0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x58, 0x49, 0x4e, 0x5f,
	0x57, 0x49, 0x4e, 0x44, 0x5f, 0x53, 0x45, 0x45, 0x44, 0x5f, 0x53, 0x50, 0x41, 0x57, 0x4e, 0x45,
	0x52, 0x10, 0x6c, 0x12, 0x3f, 0x0a, 0x3b, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49,
	0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d,
	0x49, 0x58, 0x49, 0x4e, 0x5f, 0x44, 0x4f, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x42,
	0x59, 0x5f, 0x45, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x6d, 0x12, 0x3b, 0x0a, 0x37, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f,
	0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x4d, 0x49, 0x58, 0x49, 0x4e, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x45, 0x4e, 0x54, 0x49,
	0x54, 0x59, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10,
	0x6e, 0x12, 0x31, 0x0a, 0x2d, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56,
	0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x58,
	0x49, 0x4e, 0x5f, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x5f, 0x53, 0x59,
	0x4e, 0x43, 0x10, 0x6f, 0x12, 0x33, 0x0a, 0x2f, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f,
	0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x4d, 0x49, 0x58, 0x49, 0x4e, 0x5f, 0x57, 0x49, 0x44, 0x47, 0x45, 0x54, 0x5f, 0x4d, 0x50, 0x5f,
	0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x70, 0x12, 0x2f, 0x0a, 0x2b, 0x41, 0x42, 0x49,
	0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4e, 0x4a, 0x48,
	0x42, 0x46, 0x41, 0x44, 0x45, 0x4f, 0x4f, 0x4e, 0x10, 0x71, 0x12, 0x2f, 0x0a, 0x2b, 0x41, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47,
	0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x45, 0x47,
	0x43, 0x49, 0x46, 0x46, 0x46, 0x4c, 0x4c, 0x42, 0x47, 0x10, 0x72, 0x12, 0x2f, 0x0a, 0x2b, 0x41,
	0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41, 0x52,
	0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f, 0x4f,
	0x46, 0x44, 0x47, 0x46, 0x41, 0x43, 0x4f, 0x4c, 0x44, 0x49, 0x10, 0x73, 0x12, 0x2f, 0x0a, 0x2b,
	0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x41,
	0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x6e, 0x6b, 0x32, 0x37, 0x30, 0x30, 0x5f,
	0x4b, 0x44, 0x50, 0x4b, 0x4a, 0x47, 0x4a, 0x4e, 0x47, 0x46, 0x42, 0x10, 0x74, 0x12, 0x2f, 0x0a,
	0x2b, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f,
	0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30, 0x30,
	0x5f, 0x42, 0x4e, 0x45, 0x43, 0x50, 0x41, 0x43, 0x47, 0x4b, 0x48, 0x50, 0x10, 0x75, 0x12, 0x2f,
	0x0a, 0x2b, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45,
	0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x6e, 0x6b, 0x33, 0x30, 0x30,
	0x30, 0x5f, 0x4c, 0x47, 0x49, 0x50, 0x4f, 0x43, 0x42, 0x48, 0x4b, 0x41, 0x4c, 0x10, 0x76, 0x12,
	0x2f, 0x0a, 0x2b, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f, 0x4b,
	0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x6e, 0x6b, 0x33, 0x30,
	0x30, 0x30, 0x5f, 0x45, 0x46, 0x4a, 0x49, 0x47, 0x43, 0x45, 0x47, 0x48, 0x4a, 0x47, 0x10, 0x77,
	0x12, 0x2f, 0x0a, 0x2b, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x4f,
	0x4b, 0x45, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x6e, 0x6b, 0x33,
	0x31, 0x30, 0x30, 0x5f, 0x4e, 0x4c, 0x49, 0x50, 0x42, 0x42, 0x4f, 0x49, 0x4e, 0x45, 0x4f, 0x10,
	0x78, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_AbilityInvokeArgument_proto_rawDescOnce sync.Once
	file_AbilityInvokeArgument_proto_rawDescData = file_AbilityInvokeArgument_proto_rawDesc
)

func file_AbilityInvokeArgument_proto_rawDescGZIP() []byte {
	file_AbilityInvokeArgument_proto_rawDescOnce.Do(func() {
		file_AbilityInvokeArgument_proto_rawDescData = protoimpl.X.CompressGZIP(file_AbilityInvokeArgument_proto_rawDescData)
	})
	return file_AbilityInvokeArgument_proto_rawDescData
}

var file_AbilityInvokeArgument_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AbilityInvokeArgument_proto_goTypes = []interface{}{
	(AbilityInvokeArgument)(0), // 0: AbilityInvokeArgument
}
var file_AbilityInvokeArgument_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AbilityInvokeArgument_proto_init() }
func file_AbilityInvokeArgument_proto_init() {
	if File_AbilityInvokeArgument_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_AbilityInvokeArgument_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AbilityInvokeArgument_proto_goTypes,
		DependencyIndexes: file_AbilityInvokeArgument_proto_depIdxs,
		EnumInfos:         file_AbilityInvokeArgument_proto_enumTypes,
	}.Build()
	File_AbilityInvokeArgument_proto = out.File
	file_AbilityInvokeArgument_proto_rawDesc = nil
	file_AbilityInvokeArgument_proto_goTypes = nil
	file_AbilityInvokeArgument_proto_depIdxs = nil
}
