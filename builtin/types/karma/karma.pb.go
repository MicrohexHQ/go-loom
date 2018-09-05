// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/karma/karma.proto

/*
Package karma is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/builtin/types/karma/karma.proto

It has these top-level messages:
	KarmaParams
	KarmaConfigValidator
	KarmaParamsValidatorNewMaxKarma
	KarmaParamsValidatorNewOracle
	KarmaParamsMutableValidator
	KarmaSourceReward
	KarmaSource
	KarmaAddressSource
	KarmaConfig
	KarmaState
	KarmaStateUser
	KarmaStateKeyUser
	KarmaInitRequest
	KarmaUserToken
	KarmaTotal
	KarmaDeploy
*/
package karma

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type KarmaParams struct {
	Config *KarmaConfig          `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	Users  []*KarmaAddressSource `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
}

func (m *KarmaParams) Reset()                    { *m = KarmaParams{} }
func (m *KarmaParams) String() string            { return proto.CompactTextString(m) }
func (*KarmaParams) ProtoMessage()               {}
func (*KarmaParams) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{0} }

func (m *KarmaParams) GetConfig() *KarmaConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *KarmaParams) GetUsers() []*KarmaAddressSource {
	if m != nil {
		return m.Users
	}
	return nil
}

type KarmaConfigValidator struct {
	Config *KarmaConfig   `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	Oracle *types.Address `protobuf:"bytes,2,opt,name=oracle" json:"oracle,omitempty"`
}

func (m *KarmaConfigValidator) Reset()                    { *m = KarmaConfigValidator{} }
func (m *KarmaConfigValidator) String() string            { return proto.CompactTextString(m) }
func (*KarmaConfigValidator) ProtoMessage()               {}
func (*KarmaConfigValidator) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{1} }

func (m *KarmaConfigValidator) GetConfig() *KarmaConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *KarmaConfigValidator) GetOracle() *types.Address {
	if m != nil {
		return m.Oracle
	}
	return nil
}

type KarmaParamsValidatorNewMaxKarma struct {
	MaxKarma int64          `protobuf:"varint,1,opt,name=max_karma,json=maxKarma,proto3" json:"max_karma,omitempty"`
	Oracle   *types.Address `protobuf:"bytes,2,opt,name=oracle" json:"oracle,omitempty"`
}

func (m *KarmaParamsValidatorNewMaxKarma) Reset()         { *m = KarmaParamsValidatorNewMaxKarma{} }
func (m *KarmaParamsValidatorNewMaxKarma) String() string { return proto.CompactTextString(m) }
func (*KarmaParamsValidatorNewMaxKarma) ProtoMessage()    {}
func (*KarmaParamsValidatorNewMaxKarma) Descriptor() ([]byte, []int) {
	return fileDescriptorKarma, []int{2}
}

func (m *KarmaParamsValidatorNewMaxKarma) GetMaxKarma() int64 {
	if m != nil {
		return m.MaxKarma
	}
	return 0
}

func (m *KarmaParamsValidatorNewMaxKarma) GetOracle() *types.Address {
	if m != nil {
		return m.Oracle
	}
	return nil
}

type KarmaParamsValidatorNewOracle struct {
	NewOracle *types.Address `protobuf:"bytes,1,opt,name=new_oracle,json=newOracle" json:"new_oracle,omitempty"`
	OldOracle *types.Address `protobuf:"bytes,2,opt,name=old_oracle,json=oldOracle" json:"old_oracle,omitempty"`
}

func (m *KarmaParamsValidatorNewOracle) Reset()         { *m = KarmaParamsValidatorNewOracle{} }
func (m *KarmaParamsValidatorNewOracle) String() string { return proto.CompactTextString(m) }
func (*KarmaParamsValidatorNewOracle) ProtoMessage()    {}
func (*KarmaParamsValidatorNewOracle) Descriptor() ([]byte, []int) {
	return fileDescriptorKarma, []int{3}
}

func (m *KarmaParamsValidatorNewOracle) GetNewOracle() *types.Address {
	if m != nil {
		return m.NewOracle
	}
	return nil
}

func (m *KarmaParamsValidatorNewOracle) GetOldOracle() *types.Address {
	if m != nil {
		return m.OldOracle
	}
	return nil
}

type KarmaParamsMutableValidator struct {
	MutableOracle bool           `protobuf:"varint,1,opt,name=mutable_oracle,json=mutableOracle,proto3" json:"mutable_oracle,omitempty"`
	Oracle        *types.Address `protobuf:"bytes,2,opt,name=oracle" json:"oracle,omitempty"`
}

func (m *KarmaParamsMutableValidator) Reset()                    { *m = KarmaParamsMutableValidator{} }
func (m *KarmaParamsMutableValidator) String() string            { return proto.CompactTextString(m) }
func (*KarmaParamsMutableValidator) ProtoMessage()               {}
func (*KarmaParamsMutableValidator) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{4} }

func (m *KarmaParamsMutableValidator) GetMutableOracle() bool {
	if m != nil {
		return m.MutableOracle
	}
	return false
}

func (m *KarmaParamsMutableValidator) GetOracle() *types.Address {
	if m != nil {
		return m.Oracle
	}
	return nil
}

type KarmaSourceReward struct {
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Reward int64  `protobuf:"varint,2,opt,name=reward,proto3" json:"reward,omitempty"`
}

func (m *KarmaSourceReward) Reset()                    { *m = KarmaSourceReward{} }
func (m *KarmaSourceReward) String() string            { return proto.CompactTextString(m) }
func (*KarmaSourceReward) ProtoMessage()               {}
func (*KarmaSourceReward) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{5} }

func (m *KarmaSourceReward) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KarmaSourceReward) GetReward() int64 {
	if m != nil {
		return m.Reward
	}
	return 0
}

type KarmaSource struct {
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Count int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *KarmaSource) Reset()                    { *m = KarmaSource{} }
func (m *KarmaSource) String() string            { return proto.CompactTextString(m) }
func (*KarmaSource) ProtoMessage()               {}
func (*KarmaSource) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{6} }

func (m *KarmaSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KarmaSource) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type KarmaAddressSource struct {
	User    *types.Address `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Sources []*KarmaSource `protobuf:"bytes,2,rep,name=sources" json:"sources,omitempty"`
}

func (m *KarmaAddressSource) Reset()                    { *m = KarmaAddressSource{} }
func (m *KarmaAddressSource) String() string            { return proto.CompactTextString(m) }
func (*KarmaAddressSource) ProtoMessage()               {}
func (*KarmaAddressSource) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{7} }

func (m *KarmaAddressSource) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *KarmaAddressSource) GetSources() []*KarmaSource {
	if m != nil {
		return m.Sources
	}
	return nil
}

type KarmaConfig struct {
	Enabled               bool                 `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	MutableOracle         bool                 `protobuf:"varint,2,opt,name=mutable_oracle,json=mutableOracle,proto3" json:"mutable_oracle,omitempty"`
	Oracle                *types.Address       `protobuf:"bytes,4,opt,name=Oracle" json:"Oracle,omitempty"`
	Sources               []*KarmaSourceReward `protobuf:"bytes,5,rep,name=sources" json:"sources,omitempty"`
	SessionMaxAccessCount int64                `protobuf:"varint,6,opt,name=session_max_access_count,json=sessionMaxAccessCount,proto3" json:"session_max_access_count,omitempty"`
	SessionDuration       int64                `protobuf:"varint,7,opt,name=session_duration,json=sessionDuration,proto3" json:"session_duration,omitempty"`
	DeployEnabled         bool                 `protobuf:"varint,8,opt,name=deploy_enabled,json=deployEnabled,proto3" json:"deploy_enabled,omitempty"`
	CallEnabled           bool                 `protobuf:"varint,9,opt,name=call_enabled,json=callEnabled,proto3" json:"call_enabled,omitempty"`
	LastUpdateTime        int64                `protobuf:"varint,10,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
}

func (m *KarmaConfig) Reset()                    { *m = KarmaConfig{} }
func (m *KarmaConfig) String() string            { return proto.CompactTextString(m) }
func (*KarmaConfig) ProtoMessage()               {}
func (*KarmaConfig) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{8} }

func (m *KarmaConfig) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *KarmaConfig) GetMutableOracle() bool {
	if m != nil {
		return m.MutableOracle
	}
	return false
}

func (m *KarmaConfig) GetOracle() *types.Address {
	if m != nil {
		return m.Oracle
	}
	return nil
}

func (m *KarmaConfig) GetSources() []*KarmaSourceReward {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *KarmaConfig) GetSessionMaxAccessCount() int64 {
	if m != nil {
		return m.SessionMaxAccessCount
	}
	return 0
}

func (m *KarmaConfig) GetSessionDuration() int64 {
	if m != nil {
		return m.SessionDuration
	}
	return 0
}

func (m *KarmaConfig) GetDeployEnabled() bool {
	if m != nil {
		return m.DeployEnabled
	}
	return false
}

func (m *KarmaConfig) GetCallEnabled() bool {
	if m != nil {
		return m.CallEnabled
	}
	return false
}

func (m *KarmaConfig) GetLastUpdateTime() int64 {
	if m != nil {
		return m.LastUpdateTime
	}
	return 0
}

type KarmaState struct {
	SourceStates   []*KarmaSource `protobuf:"bytes,1,rep,name=source_states,json=sourceStates" json:"source_states,omitempty"`
	LastUpdateTime int64          `protobuf:"varint,2,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
}

func (m *KarmaState) Reset()                    { *m = KarmaState{} }
func (m *KarmaState) String() string            { return proto.CompactTextString(m) }
func (*KarmaState) ProtoMessage()               {}
func (*KarmaState) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{9} }

func (m *KarmaState) GetSourceStates() []*KarmaSource {
	if m != nil {
		return m.SourceStates
	}
	return nil
}

func (m *KarmaState) GetLastUpdateTime() int64 {
	if m != nil {
		return m.LastUpdateTime
	}
	return 0
}

type KarmaStateUser struct {
	SourceStates []*KarmaSource `protobuf:"bytes,1,rep,name=source_states,json=sourceStates" json:"source_states,omitempty"`
	User         *types.Address `protobuf:"bytes,2,opt,name=User" json:"User,omitempty"`
	Oracle       *types.Address `protobuf:"bytes,3,opt,name=Oracle" json:"Oracle,omitempty"`
}

func (m *KarmaStateUser) Reset()                    { *m = KarmaStateUser{} }
func (m *KarmaStateUser) String() string            { return proto.CompactTextString(m) }
func (*KarmaStateUser) ProtoMessage()               {}
func (*KarmaStateUser) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{10} }

func (m *KarmaStateUser) GetSourceStates() []*KarmaSource {
	if m != nil {
		return m.SourceStates
	}
	return nil
}

func (m *KarmaStateUser) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *KarmaStateUser) GetOracle() *types.Address {
	if m != nil {
		return m.Oracle
	}
	return nil
}

type KarmaStateKeyUser struct {
	StateKeys []string       `protobuf:"bytes,1,rep,name=state_keys,json=stateKeys" json:"state_keys,omitempty"`
	User      *types.Address `protobuf:"bytes,2,opt,name=User" json:"User,omitempty"`
	Oracle    *types.Address `protobuf:"bytes,3,opt,name=Oracle" json:"Oracle,omitempty"`
}

func (m *KarmaStateKeyUser) Reset()                    { *m = KarmaStateKeyUser{} }
func (m *KarmaStateKeyUser) String() string            { return proto.CompactTextString(m) }
func (*KarmaStateKeyUser) ProtoMessage()               {}
func (*KarmaStateKeyUser) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{11} }

func (m *KarmaStateKeyUser) GetStateKeys() []string {
	if m != nil {
		return m.StateKeys
	}
	return nil
}

func (m *KarmaStateKeyUser) GetUser() *types.Address {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *KarmaStateKeyUser) GetOracle() *types.Address {
	if m != nil {
		return m.Oracle
	}
	return nil
}

type KarmaInitRequest struct {
	Params *KarmaParams `protobuf:"bytes,1,opt,name=Params" json:"Params,omitempty"`
}

func (m *KarmaInitRequest) Reset()                    { *m = KarmaInitRequest{} }
func (m *KarmaInitRequest) String() string            { return proto.CompactTextString(m) }
func (*KarmaInitRequest) ProtoMessage()               {}
func (*KarmaInitRequest) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{12} }

func (m *KarmaInitRequest) GetParams() *KarmaParams {
	if m != nil {
		return m.Params
	}
	return nil
}

type KarmaUserToken struct {
	Owner      *types.Address `protobuf:"bytes,1,opt,name=Owner" json:"Owner,omitempty"`
	TokenCount int64          `protobuf:"varint,2,opt,name=token_count,json=tokenCount,proto3" json:"token_count,omitempty"`
}

func (m *KarmaUserToken) Reset()                    { *m = KarmaUserToken{} }
func (m *KarmaUserToken) String() string            { return proto.CompactTextString(m) }
func (*KarmaUserToken) ProtoMessage()               {}
func (*KarmaUserToken) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{13} }

func (m *KarmaUserToken) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *KarmaUserToken) GetTokenCount() int64 {
	if m != nil {
		return m.TokenCount
	}
	return 0
}

type KarmaTotal struct {
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *KarmaTotal) Reset()                    { *m = KarmaTotal{} }
func (m *KarmaTotal) String() string            { return proto.CompactTextString(m) }
func (*KarmaTotal) ProtoMessage()               {}
func (*KarmaTotal) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{14} }

func (m *KarmaTotal) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type KarmaDeploy struct {
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *KarmaDeploy) Reset()                    { *m = KarmaDeploy{} }
func (m *KarmaDeploy) String() string            { return proto.CompactTextString(m) }
func (*KarmaDeploy) ProtoMessage()               {}
func (*KarmaDeploy) Descriptor() ([]byte, []int) { return fileDescriptorKarma, []int{15} }

func (m *KarmaDeploy) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*KarmaParams)(nil), "karma.KarmaParams")
	proto.RegisterType((*KarmaConfigValidator)(nil), "karma.KarmaConfigValidator")
	proto.RegisterType((*KarmaParamsValidatorNewMaxKarma)(nil), "karma.KarmaParamsValidatorNewMaxKarma")
	proto.RegisterType((*KarmaParamsValidatorNewOracle)(nil), "karma.KarmaParamsValidatorNewOracle")
	proto.RegisterType((*KarmaParamsMutableValidator)(nil), "karma.KarmaParamsMutableValidator")
	proto.RegisterType((*KarmaSourceReward)(nil), "karma.KarmaSourceReward")
	proto.RegisterType((*KarmaSource)(nil), "karma.KarmaSource")
	proto.RegisterType((*KarmaAddressSource)(nil), "karma.KarmaAddressSource")
	proto.RegisterType((*KarmaConfig)(nil), "karma.KarmaConfig")
	proto.RegisterType((*KarmaState)(nil), "karma.KarmaState")
	proto.RegisterType((*KarmaStateUser)(nil), "karma.KarmaStateUser")
	proto.RegisterType((*KarmaStateKeyUser)(nil), "karma.KarmaStateKeyUser")
	proto.RegisterType((*KarmaInitRequest)(nil), "karma.KarmaInitRequest")
	proto.RegisterType((*KarmaUserToken)(nil), "karma.KarmaUserToken")
	proto.RegisterType((*KarmaTotal)(nil), "karma.KarmaTotal")
	proto.RegisterType((*KarmaDeploy)(nil), "karma.KarmaDeploy")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/karma/karma.proto", fileDescriptorKarma)
}

var fileDescriptorKarma = []byte{
	// 705 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdf, 0x4f, 0x13, 0x41,
	0x10, 0x4e, 0x7f, 0xd2, 0x9b, 0x02, 0xe2, 0x06, 0xcd, 0x29, 0x20, 0x78, 0xc6, 0x88, 0x46, 0x5b,
	0x83, 0x0f, 0xbc, 0x61, 0x08, 0xf8, 0x60, 0x08, 0xa2, 0x27, 0xf8, 0x7a, 0x6c, 0xef, 0x16, 0x3c,
	0x7b, 0x77, 0x5b, 0x76, 0xf7, 0x52, 0xfa, 0x27, 0xf8, 0x3f, 0xfb, 0x60, 0x76, 0x76, 0x8f, 0x9e,
	0x70, 0x4d, 0xa3, 0xbe, 0x34, 0xdd, 0x99, 0x6f, 0xe6, 0x9b, 0xf9, 0xfa, 0xed, 0x16, 0xf6, 0x2e,
	0x63, 0xf5, 0x3d, 0x1f, 0xf4, 0x42, 0x9e, 0xf6, 0x13, 0xce, 0xd3, 0x8c, 0xa9, 0x31, 0x17, 0xc3,
	0xfe, 0x25, 0x7f, 0xa3, 0x8f, 0xfd, 0x41, 0x1e, 0x27, 0x2a, 0xce, 0xfa, 0x6a, 0x32, 0x62, 0xb2,
	0x3f, 0xa4, 0x22, 0xa5, 0xe6, 0xb3, 0x37, 0x12, 0x5c, 0x71, 0xd2, 0xc2, 0xc3, 0xe3, 0xb7, 0x73,
	0xda, 0x98, 0x72, 0xfc, 0x34, 0x85, 0xde, 0x0f, 0xe8, 0x1e, 0xe9, 0xd2, 0xcf, 0x54, 0xd0, 0x54,
	0x92, 0x57, 0xd0, 0x0e, 0x79, 0x76, 0x11, 0x5f, 0xba, 0xb5, 0xad, 0xda, 0x76, 0x77, 0x87, 0xf4,
	0x0c, 0x0b, 0x62, 0x0e, 0x30, 0xe3, 0x5b, 0x04, 0xe9, 0x43, 0x2b, 0x97, 0x4c, 0x48, 0xb7, 0xbe,
	0xd5, 0xd8, 0xee, 0xee, 0x3c, 0x2a, 0x43, 0xf7, 0xa3, 0x48, 0x30, 0x29, 0xbf, 0xf2, 0x5c, 0x84,
	0xcc, 0x37, 0x38, 0x2f, 0x82, 0xd5, 0x52, 0x9f, 0x6f, 0x34, 0x89, 0x23, 0xaa, 0xb8, 0xf8, 0x2b,
	0xd2, 0x2d, 0x68, 0x73, 0x41, 0xc3, 0x84, 0xb9, 0x75, 0xc4, 0x76, 0x7a, 0x96, 0xca, 0xb7, 0x71,
	0xef, 0x1c, 0x36, 0x4b, 0x1b, 0xdd, 0xb0, 0x7c, 0x62, 0xe3, 0x63, 0x7a, 0x8d, 0x19, 0xb2, 0x06,
	0x4e, 0x4a, 0xaf, 0x03, 0x64, 0x41, 0xce, 0x86, 0xdf, 0x49, 0x8b, 0xe4, 0x7c, 0x86, 0x2b, 0xd8,
	0x98, 0xc1, 0x70, 0x82, 0x00, 0xf2, 0x02, 0x20, 0x63, 0xe3, 0xc0, 0xb6, 0xa9, 0xdd, 0x6a, 0xe3,
	0x64, 0x65, 0x20, 0x4f, 0xa2, 0x60, 0x06, 0x9f, 0xc3, 0x93, 0xc8, 0x00, 0xbd, 0x0b, 0x58, 0x2b,
	0x51, 0x1e, 0xe7, 0x8a, 0x0e, 0x12, 0x36, 0x55, 0xf0, 0x39, 0x2c, 0xa7, 0x26, 0x56, 0x26, 0xed,
	0xf8, 0x4b, 0x36, 0x6a, 0xe9, 0xe6, 0xaf, 0xf6, 0x1e, 0xee, 0x23, 0x8f, 0xfd, 0xe1, 0xd8, 0x98,
	0x8a, 0x88, 0x10, 0x68, 0x66, 0x34, 0x35, 0x3d, 0x1d, 0x1f, 0xbf, 0x93, 0x87, 0xd0, 0x16, 0x98,
	0xc5, 0x56, 0x0d, 0xdf, 0x9e, 0xbc, 0x5d, 0xeb, 0x27, 0xd3, 0xa0, 0xb2, 0x74, 0x15, 0x5a, 0x21,
	0xcf, 0x33, 0x65, 0x2b, 0xcd, 0xc1, 0x3b, 0x07, 0x72, 0xd7, 0x39, 0x64, 0x1d, 0x9a, 0xda, 0x3b,
	0x77, 0x34, 0xc4, 0x28, 0x79, 0x0d, 0x0b, 0x12, 0x71, 0x85, 0x07, 0xff, 0x70, 0x8e, 0xdd, 0xa1,
	0x80, 0x78, 0xbf, 0xea, 0x76, 0x36, 0x63, 0x29, 0xe2, 0xc2, 0x02, 0xcb, 0xb4, 0x3a, 0x91, 0x55,
	0xab, 0x38, 0x56, 0xc8, 0x59, 0x9f, 0x21, 0xa7, 0xf9, 0xe6, 0x36, 0x6f, 0xcb, 0x69, 0x11, 0x3b,
	0xd3, 0x01, 0x5b, 0x38, 0xa0, 0x5b, 0x31, 0x20, 0x0a, 0x77, 0x33, 0x26, 0xd9, 0x05, 0x57, 0x32,
	0x29, 0x63, 0x9e, 0x05, 0xda, 0xa4, 0x34, 0x0c, 0x99, 0x94, 0x81, 0x51, 0xac, 0x8d, 0x8a, 0x3d,
	0xb0, 0xf9, 0x63, 0x7a, 0xbd, 0x8f, 0xd9, 0x03, 0x9d, 0x24, 0x2f, 0x61, 0xa5, 0x28, 0x8c, 0x72,
	0x41, 0x55, 0xcc, 0x33, 0x77, 0x01, 0x0b, 0xee, 0xd9, 0xf8, 0xa1, 0x0d, 0xeb, 0x05, 0x23, 0x36,
	0x4a, 0xf8, 0x24, 0x28, 0x14, 0xe8, 0x98, 0x05, 0x4d, 0xf4, 0x83, 0xd5, 0xe1, 0x29, 0x2c, 0x86,
	0x34, 0x49, 0x6e, 0x40, 0x0e, 0x82, 0xba, 0x3a, 0x56, 0x40, 0xb6, 0x61, 0x25, 0xa1, 0x52, 0x05,
	0xf9, 0x28, 0xa2, 0x8a, 0x05, 0x2a, 0x4e, 0x99, 0x0b, 0x48, 0xba, 0xac, 0xe3, 0x67, 0x18, 0x3e,
	0x8d, 0x53, 0xe6, 0x71, 0x00, 0xb3, 0xb5, 0xa2, 0x8a, 0x91, 0x5d, 0x58, 0x32, 0x0b, 0x07, 0x52,
	0x9f, 0xa5, 0x5b, 0x9b, 0xf9, 0x03, 0x2e, 0x1a, 0x20, 0xd6, 0xc9, 0x4a, 0xc2, 0x7a, 0x25, 0xe1,
	0xcf, 0x1a, 0x2c, 0x4f, 0x19, 0xcf, 0xb4, 0x61, 0xfe, 0x99, 0x75, 0x1d, 0x9a, 0xba, 0xc1, 0x9d,
	0x7b, 0x83, 0xd1, 0x92, 0x11, 0x1a, 0xd5, 0x46, 0xf0, 0x44, 0x71, 0xaf, 0x74, 0xbb, 0x23, 0x36,
	0xc1, 0xb2, 0x0d, 0x00, 0x1c, 0x23, 0x18, 0xb2, 0x89, 0x19, 0xc5, 0xf1, 0x1d, 0x69, 0x11, 0xff,
	0xcf, 0xb9, 0x07, 0x2b, 0xc8, 0xf9, 0x31, 0x8b, 0x95, 0xcf, 0xae, 0x72, 0x26, 0x95, 0x7e, 0x6a,
	0xcd, 0x13, 0x52, 0xf5, 0xd4, 0x9a, 0x8c, 0x6f, 0x11, 0xde, 0x17, 0x2b, 0x9f, 0xa6, 0x3b, 0xe5,
	0x43, 0x96, 0x91, 0x27, 0xd0, 0x3a, 0x19, 0x67, 0x15, 0xd7, 0xd1, 0x84, 0xc9, 0x26, 0x74, 0x95,
	0x06, 0x06, 0xe5, 0xfb, 0x0d, 0x18, 0x42, 0x8b, 0x7a, 0x9e, 0xf5, 0xc0, 0x29, 0x57, 0x34, 0x99,
	0x3e, 0x04, 0xb5, 0xf2, 0x43, 0xf0, 0xcc, 0xde, 0xd2, 0x43, 0xb4, 0x62, 0x35, 0x68, 0xd0, 0xc6,
	0x7f, 0xaf, 0x77, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x57, 0x15, 0xb8, 0x82, 0x38, 0x07, 0x00,
	0x00,
}
