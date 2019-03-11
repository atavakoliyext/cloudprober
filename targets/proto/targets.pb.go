// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/yext/cloudprober/targets/proto/targets.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "github.com/yext/cloudprober/targets/gce/proto"
import proto4 "github.com/yext/cloudprober/targets/lameduck/proto"
import proto2 "github.com/yext/cloudprober/targets/rds/client/proto"
import proto3 "github.com/yext/cloudprober/targets/rtc/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TargetsDef struct {
	// Types that are valid to be assigned to Type:
	//	*TargetsDef_HostNames
	//	*TargetsDef_GceTargets
	//	*TargetsDef_RdsTargets
	//	*TargetsDef_RtcTargets
	//	*TargetsDef_DummyTargets
	Type isTargetsDef_Type `protobuf_oneof:"type"`
	// Regex to apply on the targets.
	Regex *string `protobuf:"bytes,21,opt,name=regex" json:"regex,omitempty"`
	// Exclude lameducks. Lameduck targets can be set through RTC (realtime
	// configurator) service. This functionality works only if lame_duck_options
	// are specified.
	ExcludeLameducks *bool `protobuf:"varint,22,opt,name=exclude_lameducks,json=excludeLameducks,def=1" json:"exclude_lameducks,omitempty"`
	// How often targets should be evaluated. Any number less than or equal to 0
	// will result in no target caching (targets will be reevaluated on demand).
	// Note that individual target types may have their own caches implemented
	// (specifically GCE instances/forwarding rules). This does not impact those
	// caches.
	ReEvalSec                    *int32   `protobuf:"varint,23,opt,name=re_eval_sec,json=reEvalSec,def=0" json:"re_eval_sec,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *TargetsDef) Reset()         { *m = TargetsDef{} }
func (m *TargetsDef) String() string { return proto.CompactTextString(m) }
func (*TargetsDef) ProtoMessage()    {}
func (*TargetsDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_targets_d8fa4415f08434d8, []int{0}
}

var extRange_TargetsDef = []proto.ExtensionRange{
	{Start: 200, End: 536870911},
}

func (*TargetsDef) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_TargetsDef
}
func (m *TargetsDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetsDef.Unmarshal(m, b)
}
func (m *TargetsDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetsDef.Marshal(b, m, deterministic)
}
func (dst *TargetsDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetsDef.Merge(dst, src)
}
func (m *TargetsDef) XXX_Size() int {
	return xxx_messageInfo_TargetsDef.Size(m)
}
func (m *TargetsDef) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetsDef.DiscardUnknown(m)
}

var xxx_messageInfo_TargetsDef proto.InternalMessageInfo

const Default_TargetsDef_ExcludeLameducks bool = true
const Default_TargetsDef_ReEvalSec int32 = 0

type isTargetsDef_Type interface {
	isTargetsDef_Type()
}

type TargetsDef_HostNames struct {
	HostNames string `protobuf:"bytes,1,opt,name=host_names,json=hostNames,oneof"`
}

type TargetsDef_GceTargets struct {
	GceTargets *proto1.TargetsConf `protobuf:"bytes,2,opt,name=gce_targets,json=gceTargets,oneof"`
}

type TargetsDef_RdsTargets struct {
	RdsTargets *proto2.ClientConf `protobuf:"bytes,3,opt,name=rds_targets,json=rdsTargets,oneof"`
}

type TargetsDef_RtcTargets struct {
	RtcTargets *proto3.TargetsConf `protobuf:"bytes,4,opt,name=rtc_targets,json=rtcTargets,oneof"`
}

type TargetsDef_DummyTargets struct {
	DummyTargets *DummyTargets `protobuf:"bytes,20,opt,name=dummy_targets,json=dummyTargets,oneof"`
}

func (*TargetsDef_HostNames) isTargetsDef_Type() {}

func (*TargetsDef_GceTargets) isTargetsDef_Type() {}

func (*TargetsDef_RdsTargets) isTargetsDef_Type() {}

func (*TargetsDef_RtcTargets) isTargetsDef_Type() {}

func (*TargetsDef_DummyTargets) isTargetsDef_Type() {}

func (m *TargetsDef) GetType() isTargetsDef_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *TargetsDef) GetHostNames() string {
	if x, ok := m.GetType().(*TargetsDef_HostNames); ok {
		return x.HostNames
	}
	return ""
}

func (m *TargetsDef) GetGceTargets() *proto1.TargetsConf {
	if x, ok := m.GetType().(*TargetsDef_GceTargets); ok {
		return x.GceTargets
	}
	return nil
}

func (m *TargetsDef) GetRdsTargets() *proto2.ClientConf {
	if x, ok := m.GetType().(*TargetsDef_RdsTargets); ok {
		return x.RdsTargets
	}
	return nil
}

func (m *TargetsDef) GetRtcTargets() *proto3.TargetsConf {
	if x, ok := m.GetType().(*TargetsDef_RtcTargets); ok {
		return x.RtcTargets
	}
	return nil
}

func (m *TargetsDef) GetDummyTargets() *DummyTargets {
	if x, ok := m.GetType().(*TargetsDef_DummyTargets); ok {
		return x.DummyTargets
	}
	return nil
}

func (m *TargetsDef) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *TargetsDef) GetExcludeLameducks() bool {
	if m != nil && m.ExcludeLameducks != nil {
		return *m.ExcludeLameducks
	}
	return Default_TargetsDef_ExcludeLameducks
}

func (m *TargetsDef) GetReEvalSec() int32 {
	if m != nil && m.ReEvalSec != nil {
		return *m.ReEvalSec
	}
	return Default_TargetsDef_ReEvalSec
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TargetsDef) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TargetsDef_OneofMarshaler, _TargetsDef_OneofUnmarshaler, _TargetsDef_OneofSizer, []interface{}{
		(*TargetsDef_HostNames)(nil),
		(*TargetsDef_GceTargets)(nil),
		(*TargetsDef_RdsTargets)(nil),
		(*TargetsDef_RtcTargets)(nil),
		(*TargetsDef_DummyTargets)(nil),
	}
}

func _TargetsDef_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TargetsDef)
	// type
	switch x := m.Type.(type) {
	case *TargetsDef_HostNames:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.HostNames)
	case *TargetsDef_GceTargets:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GceTargets); err != nil {
			return err
		}
	case *TargetsDef_RdsTargets:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RdsTargets); err != nil {
			return err
		}
	case *TargetsDef_RtcTargets:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RtcTargets); err != nil {
			return err
		}
	case *TargetsDef_DummyTargets:
		b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DummyTargets); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TargetsDef.Type has unexpected type %T", x)
	}
	return nil
}

func _TargetsDef_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TargetsDef)
	switch tag {
	case 1: // type.host_names
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Type = &TargetsDef_HostNames{x}
		return true, err
	case 2: // type.gce_targets
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(proto1.TargetsConf)
		err := b.DecodeMessage(msg)
		m.Type = &TargetsDef_GceTargets{msg}
		return true, err
	case 3: // type.rds_targets
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(proto2.ClientConf)
		err := b.DecodeMessage(msg)
		m.Type = &TargetsDef_RdsTargets{msg}
		return true, err
	case 4: // type.rtc_targets
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(proto3.TargetsConf)
		err := b.DecodeMessage(msg)
		m.Type = &TargetsDef_RtcTargets{msg}
		return true, err
	case 20: // type.dummy_targets
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DummyTargets)
		err := b.DecodeMessage(msg)
		m.Type = &TargetsDef_DummyTargets{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TargetsDef_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TargetsDef)
	// type
	switch x := m.Type.(type) {
	case *TargetsDef_HostNames:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.HostNames)))
		n += len(x.HostNames)
	case *TargetsDef_GceTargets:
		s := proto.Size(x.GceTargets)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TargetsDef_RdsTargets:
		s := proto.Size(x.RdsTargets)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TargetsDef_RtcTargets:
		s := proto.Size(x.RtcTargets)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TargetsDef_DummyTargets:
		s := proto.Size(x.DummyTargets)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// DummyTargets represent empty targets, which are useful for external
// probes that do not have any "proper" targets.  Such as ilbprober.
type DummyTargets struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DummyTargets) Reset()         { *m = DummyTargets{} }
func (m *DummyTargets) String() string { return proto.CompactTextString(m) }
func (*DummyTargets) ProtoMessage()    {}
func (*DummyTargets) Descriptor() ([]byte, []int) {
	return fileDescriptor_targets_d8fa4415f08434d8, []int{1}
}
func (m *DummyTargets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DummyTargets.Unmarshal(m, b)
}
func (m *DummyTargets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DummyTargets.Marshal(b, m, deterministic)
}
func (dst *DummyTargets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DummyTargets.Merge(dst, src)
}
func (m *DummyTargets) XXX_Size() int {
	return xxx_messageInfo_DummyTargets.Size(m)
}
func (m *DummyTargets) XXX_DiscardUnknown() {
	xxx_messageInfo_DummyTargets.DiscardUnknown(m)
}

var xxx_messageInfo_DummyTargets proto.InternalMessageInfo

// Global targets options. These options are independent of the per-probe
// targets which are defined by the "Targets" type above.
//
// Currently these options are used only for GCE targets to control things like
// how often to re-evaluate the targets and whether to check for lame ducks or
// not.
type GlobalTargetsOptions struct {
	// GCE targets options.
	GlobalGceTargetsOptions *proto1.GlobalOptions `protobuf:"bytes,1,opt,name=global_gce_targets_options,json=globalGceTargetsOptions" json:"global_gce_targets_options,omitempty"`
	// Lame duck options. If provided, targets module checks for the lame duck
	// targets and removes them from the targets list.
	LameDuckOptions      *proto4.Options `protobuf:"bytes,2,opt,name=lame_duck_options,json=lameDuckOptions" json:"lame_duck_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GlobalTargetsOptions) Reset()         { *m = GlobalTargetsOptions{} }
func (m *GlobalTargetsOptions) String() string { return proto.CompactTextString(m) }
func (*GlobalTargetsOptions) ProtoMessage()    {}
func (*GlobalTargetsOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_targets_d8fa4415f08434d8, []int{2}
}
func (m *GlobalTargetsOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalTargetsOptions.Unmarshal(m, b)
}
func (m *GlobalTargetsOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalTargetsOptions.Marshal(b, m, deterministic)
}
func (dst *GlobalTargetsOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalTargetsOptions.Merge(dst, src)
}
func (m *GlobalTargetsOptions) XXX_Size() int {
	return xxx_messageInfo_GlobalTargetsOptions.Size(m)
}
func (m *GlobalTargetsOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalTargetsOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalTargetsOptions proto.InternalMessageInfo

func (m *GlobalTargetsOptions) GetGlobalGceTargetsOptions() *proto1.GlobalOptions {
	if m != nil {
		return m.GlobalGceTargetsOptions
	}
	return nil
}

func (m *GlobalTargetsOptions) GetLameDuckOptions() *proto4.Options {
	if m != nil {
		return m.LameDuckOptions
	}
	return nil
}

func init() {
	proto.RegisterType((*TargetsDef)(nil), "cloudprober.targets.TargetsDef")
	proto.RegisterType((*DummyTargets)(nil), "cloudprober.targets.DummyTargets")
	proto.RegisterType((*GlobalTargetsOptions)(nil), "cloudprober.targets.GlobalTargetsOptions")
}

func init() {
	proto.RegisterFile("github.com/yext/cloudprober/targets/proto/targets.proto", fileDescriptor_targets_d8fa4415f08434d8)
}

var fileDescriptor_targets_d8fa4415f08434d8 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe3, 0x36, 0x45, 0xcd, 0xb8, 0x40, 0x6b, 0x02, 0xb5, 0x72, 0x21, 0x0d, 0x7f, 0x14,
	0x71, 0xb0, 0xa1, 0x37, 0x2a, 0x2e, 0x90, 0x40, 0x72, 0x40, 0x20, 0x0c, 0x77, 0xcb, 0x19, 0x4f,
	0xb6, 0x51, 0xd7, 0xde, 0x68, 0x77, 0x5d, 0xb5, 0xb7, 0x3c, 0x1e, 0x27, 0xc4, 0x23, 0x21, 0xdb,
	0xbb, 0x56, 0x84, 0x0c, 0x8a, 0xd4, 0xe3, 0x7c, 0x33, 0xf3, 0xdb, 0x99, 0xf9, 0x16, 0xde, 0xb2,
	0x95, 0xbe, 0x2c, 0x16, 0x01, 0x8a, 0x2c, 0x64, 0x42, 0x30, 0x4e, 0x21, 0x72, 0x51, 0xa4, 0x6b,
	0x29, 0x16, 0x24, 0x43, 0x9d, 0x48, 0x46, 0x5a, 0x85, 0x6b, 0x29, 0xb4, 0xb0, 0x51, 0x50, 0x45,
	0xde, 0xa3, 0xad, 0xc2, 0xc0, 0xa4, 0x06, 0xef, 0x77, 0xe3, 0xf1, 0x24, 0xa3, 0xb4, 0xc0, 0x2b,
	0x03, 0x46, 0x91, 0x2f, 0x57, 0xac, 0xe6, 0x0e, 0xde, 0xed, 0x86, 0x60, 0x48, 0x6d, 0xdd, 0x93,
	0xdd, 0xba, 0x65, 0xaa, 0x42, 0xe4, 0x2b, 0xca, 0xf5, 0x1d, 0x46, 0x90, 0x1a, 0x5b, 0xba, 0x47,
	0xbf, 0xf7, 0x01, 0x7e, 0xd4, 0x25, 0x53, 0x5a, 0x7a, 0x4f, 0x01, 0x2e, 0x85, 0xd2, 0x71, 0x9e,
	0x64, 0xa4, 0x7c, 0x67, 0xe8, 0x8c, 0x7b, 0xf3, 0x4e, 0xd4, 0x2b, 0xb5, 0x2f, 0xa5, 0xe4, 0xcd,
	0xc0, 0x65, 0x48, 0xb1, 0xa1, 0xfa, 0x7b, 0x43, 0x67, 0xec, 0x9e, 0x3f, 0x0f, 0x5a, 0xce, 0x1b,
	0x30, 0xa4, 0xc0, 0xa0, 0x27, 0x22, 0x5f, 0xce, 0x3b, 0x11, 0x30, 0x24, 0xa3, 0x78, 0x9f, 0xc0,
	0x95, 0xa9, 0x6a, 0x40, 0xfb, 0x15, 0xe8, 0x59, 0x2b, 0x48, 0xa6, 0x2a, 0x98, 0x54, 0xfb, 0x5b,
	0x8e, 0x4c, 0x95, 0xe5, 0xcc, 0xc0, 0x95, 0x1a, 0x1b, 0x4e, 0xf7, 0x3f, 0x03, 0x49, 0x8d, 0x7f,
	0x0f, 0x24, 0x35, 0x5a, 0xd0, 0x1c, 0xee, 0xa7, 0x45, 0x96, 0xdd, 0x36, 0xa8, 0x7e, 0x85, 0x3a,
	0x6b, 0x45, 0x4d, 0xcb, 0x4a, 0xd3, 0x39, 0xef, 0x44, 0x47, 0xe9, 0x56, 0xec, 0xf5, 0xe1, 0x40,
	0x12, 0xa3, 0x1b, 0xff, 0x71, 0x79, 0xbf, 0xa8, 0x0e, 0xbc, 0x37, 0x70, 0x42, 0x37, 0xc8, 0x8b,
	0x94, 0x62, 0xfb, 0xa3, 0x94, 0xff, 0x64, 0xe8, 0x8c, 0x0f, 0x2f, 0xba, 0x5a, 0x16, 0x14, 0x1d,
	0x9b, 0xf4, 0x67, 0x9b, 0xf5, 0xce, 0xc0, 0x95, 0x14, 0xd3, 0x75, 0xc2, 0x63, 0x45, 0xe8, 0x9f,
	0x0e, 0x9d, 0xf1, 0xc1, 0x85, 0xf3, 0x3a, 0xea, 0x49, 0xfa, 0x78, 0x9d, 0xf0, 0xef, 0x84, 0xaf,
	0x7a, 0x87, 0x3f, 0x9d, 0xe3, 0xcd, 0x66, 0xb3, 0xd9, 0xfb, 0x70, 0x0f, 0xba, 0xfa, 0x76, 0x4d,
	0xa3, 0x07, 0x70, 0xb4, 0x3d, 0xde, 0xe8, 0x97, 0x03, 0xfd, 0x19, 0x17, 0x8b, 0x84, 0x1b, 0xe5,
	0xeb, 0x5a, 0xaf, 0x44, 0xae, 0x3c, 0x84, 0x01, 0xab, 0xf4, 0x78, 0xcb, 0xd2, 0x58, 0xd4, 0xd9,
	0xca, 0x7c, 0xf7, 0xfc, 0xe5, 0x3f, 0xad, 0xad, 0x91, 0x86, 0x15, 0x9d, 0xd6, 0xa4, 0x59, 0x63,
	0xb0, 0x7d, 0xe4, 0x1b, 0x9c, 0x94, 0xeb, 0xc6, 0xe5, 0x46, 0x0d, 0xbb, 0xfe, 0x36, 0x2f, 0x5a,
	0xd9, 0xf6, 0x38, 0x81, 0x45, 0x3f, 0x2c, 0x95, 0x69, 0x81, 0x57, 0x46, 0xf8, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x8e, 0xc2, 0xbc, 0x78, 0x08, 0x04, 0x00, 0x00,
}
