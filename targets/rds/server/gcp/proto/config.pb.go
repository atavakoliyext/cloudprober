// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/yext/cloudprober/targets/rds/server/gcp/proto/config.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GCEInstances struct {
	// Zone filter, e.g. zone_filter: "us-east1-*" for all zones in us-east1
	// region.
	ZoneFilter *string `protobuf:"bytes,1,opt,name=zone_filter,json=zoneFilter" json:"zone_filter,omitempty"`
	// How often resources should be evaluated/expanded.
	ReEvalSec            *int32   `protobuf:"varint,98,opt,name=re_eval_sec,json=reEvalSec,def=300" json:"re_eval_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GCEInstances) Reset()         { *m = GCEInstances{} }
func (m *GCEInstances) String() string { return proto.CompactTextString(m) }
func (*GCEInstances) ProtoMessage()    {}
func (*GCEInstances) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_78533f985346d64e, []int{0}
}
func (m *GCEInstances) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GCEInstances.Unmarshal(m, b)
}
func (m *GCEInstances) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GCEInstances.Marshal(b, m, deterministic)
}
func (dst *GCEInstances) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GCEInstances.Merge(dst, src)
}
func (m *GCEInstances) XXX_Size() int {
	return xxx_messageInfo_GCEInstances.Size(m)
}
func (m *GCEInstances) XXX_DiscardUnknown() {
	xxx_messageInfo_GCEInstances.DiscardUnknown(m)
}

var xxx_messageInfo_GCEInstances proto.InternalMessageInfo

const Default_GCEInstances_ReEvalSec int32 = 300

func (m *GCEInstances) GetZoneFilter() string {
	if m != nil && m.ZoneFilter != nil {
		return *m.ZoneFilter
	}
	return ""
}

func (m *GCEInstances) GetReEvalSec() int32 {
	if m != nil && m.ReEvalSec != nil {
		return *m.ReEvalSec
	}
	return Default_GCEInstances_ReEvalSec
}

type RegionalForwardingRules struct {
	// Region filter, e.g. "region_filter:europe-*"
	RegionFilter         *string  `protobuf:"bytes,1,opt,name=region_filter,json=regionFilter" json:"region_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegionalForwardingRules) Reset()         { *m = RegionalForwardingRules{} }
func (m *RegionalForwardingRules) String() string { return proto.CompactTextString(m) }
func (*RegionalForwardingRules) ProtoMessage()    {}
func (*RegionalForwardingRules) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_78533f985346d64e, []int{1}
}
func (m *RegionalForwardingRules) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegionalForwardingRules.Unmarshal(m, b)
}
func (m *RegionalForwardingRules) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegionalForwardingRules.Marshal(b, m, deterministic)
}
func (dst *RegionalForwardingRules) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegionalForwardingRules.Merge(dst, src)
}
func (m *RegionalForwardingRules) XXX_Size() int {
	return xxx_messageInfo_RegionalForwardingRules.Size(m)
}
func (m *RegionalForwardingRules) XXX_DiscardUnknown() {
	xxx_messageInfo_RegionalForwardingRules.DiscardUnknown(m)
}

var xxx_messageInfo_RegionalForwardingRules proto.InternalMessageInfo

func (m *RegionalForwardingRules) GetRegionFilter() string {
	if m != nil && m.RegionFilter != nil {
		return *m.RegionFilter
	}
	return ""
}

// Runtime configurator variables.
type RTCVariables struct {
	RtcConfig            []*RTCVariables_RTCConfig `protobuf:"bytes,1,rep,name=rtc_config,json=rtcConfig" json:"rtc_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RTCVariables) Reset()         { *m = RTCVariables{} }
func (m *RTCVariables) String() string { return proto.CompactTextString(m) }
func (*RTCVariables) ProtoMessage()    {}
func (*RTCVariables) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_78533f985346d64e, []int{2}
}
func (m *RTCVariables) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RTCVariables.Unmarshal(m, b)
}
func (m *RTCVariables) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RTCVariables.Marshal(b, m, deterministic)
}
func (dst *RTCVariables) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RTCVariables.Merge(dst, src)
}
func (m *RTCVariables) XXX_Size() int {
	return xxx_messageInfo_RTCVariables.Size(m)
}
func (m *RTCVariables) XXX_DiscardUnknown() {
	xxx_messageInfo_RTCVariables.DiscardUnknown(m)
}

var xxx_messageInfo_RTCVariables proto.InternalMessageInfo

func (m *RTCVariables) GetRtcConfig() []*RTCVariables_RTCConfig {
	if m != nil {
		return m.RtcConfig
	}
	return nil
}

type RTCVariables_RTCConfig struct {
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// How often RTC variables should be evaluated/expanded.
	ReEvalSec            *int32   `protobuf:"varint,2,opt,name=re_eval_sec,json=reEvalSec,def=10" json:"re_eval_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RTCVariables_RTCConfig) Reset()         { *m = RTCVariables_RTCConfig{} }
func (m *RTCVariables_RTCConfig) String() string { return proto.CompactTextString(m) }
func (*RTCVariables_RTCConfig) ProtoMessage()    {}
func (*RTCVariables_RTCConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_78533f985346d64e, []int{2, 0}
}
func (m *RTCVariables_RTCConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RTCVariables_RTCConfig.Unmarshal(m, b)
}
func (m *RTCVariables_RTCConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RTCVariables_RTCConfig.Marshal(b, m, deterministic)
}
func (dst *RTCVariables_RTCConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RTCVariables_RTCConfig.Merge(dst, src)
}
func (m *RTCVariables_RTCConfig) XXX_Size() int {
	return xxx_messageInfo_RTCVariables_RTCConfig.Size(m)
}
func (m *RTCVariables_RTCConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RTCVariables_RTCConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RTCVariables_RTCConfig proto.InternalMessageInfo

const Default_RTCVariables_RTCConfig_ReEvalSec int32 = 10

func (m *RTCVariables_RTCConfig) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *RTCVariables_RTCConfig) GetReEvalSec() int32 {
	if m != nil && m.ReEvalSec != nil {
		return *m.ReEvalSec
	}
	return Default_RTCVariables_RTCConfig_ReEvalSec
}

// GCP provider config.
type ProviderConfig struct {
	// GCP projects. If running on GCE, it defaults to the local project.
	Project []string `protobuf:"bytes,1,rep,name=project" json:"project,omitempty"`
	// GCE instances discovery options. This field should be declared for the GCE
	// instances discovery to be enabled.
	GceInstances *GCEInstances `protobuf:"bytes,2,opt,name=gce_instances,json=gceInstances" json:"gce_instances,omitempty"`
	// Regional forwarding rules discovery options. This field should be declared
	// for the regional forwarding rules discovery to be enabled.
	// NOTE: This is not supported yet.
	RegionalForwardingRules *RegionalForwardingRules `protobuf:"bytes,3,opt,name=regional_forwarding_rules,json=regionalForwardingRules" json:"regional_forwarding_rules,omitempty"`
	// RTC variables discovery options.
	RtcVariables *RTCVariables `protobuf:"bytes,4,opt,name=rtc_variables,json=rtcVariables" json:"rtc_variables,omitempty"`
	// Compute API version.
	ApiVersion           *string  `protobuf:"bytes,99,opt,name=api_version,json=apiVersion,def=v1" json:"api_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProviderConfig) Reset()         { *m = ProviderConfig{} }
func (m *ProviderConfig) String() string { return proto.CompactTextString(m) }
func (*ProviderConfig) ProtoMessage()    {}
func (*ProviderConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_78533f985346d64e, []int{3}
}
func (m *ProviderConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderConfig.Unmarshal(m, b)
}
func (m *ProviderConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderConfig.Marshal(b, m, deterministic)
}
func (dst *ProviderConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderConfig.Merge(dst, src)
}
func (m *ProviderConfig) XXX_Size() int {
	return xxx_messageInfo_ProviderConfig.Size(m)
}
func (m *ProviderConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderConfig proto.InternalMessageInfo

const Default_ProviderConfig_ApiVersion string = "v1"

func (m *ProviderConfig) GetProject() []string {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *ProviderConfig) GetGceInstances() *GCEInstances {
	if m != nil {
		return m.GceInstances
	}
	return nil
}

func (m *ProviderConfig) GetRegionalForwardingRules() *RegionalForwardingRules {
	if m != nil {
		return m.RegionalForwardingRules
	}
	return nil
}

func (m *ProviderConfig) GetRtcVariables() *RTCVariables {
	if m != nil {
		return m.RtcVariables
	}
	return nil
}

func (m *ProviderConfig) GetApiVersion() string {
	if m != nil && m.ApiVersion != nil {
		return *m.ApiVersion
	}
	return Default_ProviderConfig_ApiVersion
}

func init() {
	proto.RegisterType((*GCEInstances)(nil), "cloudprober.targets.rds.gcp.GCEInstances")
	proto.RegisterType((*RegionalForwardingRules)(nil), "cloudprober.targets.rds.gcp.RegionalForwardingRules")
	proto.RegisterType((*RTCVariables)(nil), "cloudprober.targets.rds.gcp.RTCVariables")
	proto.RegisterType((*RTCVariables_RTCConfig)(nil), "cloudprober.targets.rds.gcp.RTCVariables.RTCConfig")
	proto.RegisterType((*ProviderConfig)(nil), "cloudprober.targets.rds.gcp.ProviderConfig")
}

func init() {
	proto.RegisterFile("github.com/yext/cloudprober/targets/rds/server/gcp/proto/config.proto", fileDescriptor_config_78533f985346d64e)
}

var fileDescriptor_config_78533f985346d64e = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3f, 0x6f, 0xdb, 0x30,
	0x10, 0xc5, 0x21, 0x39, 0x45, 0xa1, 0x93, 0xdc, 0x41, 0x4b, 0xd4, 0x76, 0xa8, 0x61, 0x2f, 0xee,
	0x22, 0xe5, 0x4f, 0xa7, 0x0c, 0x5d, 0x8c, 0x24, 0xe8, 0x12, 0x14, 0x6c, 0x90, 0x55, 0xa0, 0xa9,
	0x33, 0xcb, 0x42, 0x21, 0x89, 0x23, 0xad, 0x02, 0xfd, 0x42, 0xfd, 0x5c, 0xfd, 0x26, 0x85, 0x28,
	0x3b, 0x91, 0x81, 0xd6, 0xc8, 0xc6, 0x7b, 0xc2, 0x3b, 0xea, 0xbd, 0x1f, 0xe1, 0x56, 0x2a, 0xff,
	0x7d, 0xbb, 0x2e, 0x85, 0x79, 0xac, 0xa4, 0x31, 0xb2, 0xc5, 0x4a, 0xb4, 0x66, 0xdb, 0x58, 0x32,
	0x6b, 0xa4, 0xca, 0x73, 0x92, 0xe8, 0x5d, 0x45, 0x8d, 0xab, 0x1c, 0x52, 0x87, 0x54, 0x49, 0x61,
	0x2b, 0x4b, 0xc6, 0x9b, 0x4a, 0x18, 0xbd, 0x51, 0xb2, 0x0c, 0x43, 0xfe, 0x7e, 0x64, 0x2b, 0x77,
	0xb6, 0x92, 0x1a, 0x57, 0x4a, 0x61, 0xe7, 0xf7, 0x90, 0xdd, 0xae, 0xae, 0xbf, 0x68, 0xe7, 0xb9,
	0x16, 0xe8, 0xf2, 0x0f, 0x90, 0xfe, 0x32, 0x1a, 0xeb, 0x8d, 0x6a, 0x3d, 0x52, 0x11, 0xcd, 0xa2,
	0x65, 0xc2, 0xa0, 0x97, 0x6e, 0x82, 0x92, 0x2f, 0x20, 0x25, 0xac, 0xb1, 0xe3, 0x6d, 0xed, 0x50,
	0x14, 0xeb, 0x59, 0xb4, 0x7c, 0x75, 0x35, 0xb9, 0x3c, 0x3b, 0x63, 0x09, 0xe1, 0x75, 0xc7, 0xdb,
	0x6f, 0x28, 0xe6, 0x9f, 0xe1, 0x94, 0xa1, 0x54, 0x46, 0xf3, 0xf6, 0xc6, 0xd0, 0x4f, 0x4e, 0x8d,
	0xd2, 0x92, 0x6d, 0x5b, 0x74, 0xf9, 0x02, 0xa6, 0x14, 0x3e, 0x1d, 0x5e, 0x91, 0x0d, 0xe2, 0x70,
	0xc9, 0xfc, 0x77, 0x04, 0x19, 0xbb, 0x5f, 0x3d, 0x70, 0x52, 0x7c, 0xdd, 0xbb, 0x18, 0x00, 0x79,
	0x51, 0x0f, 0xb9, 0x8a, 0x68, 0x36, 0x59, 0xa6, 0x17, 0x97, 0xe5, 0x91, 0x60, 0xe5, 0xd8, 0xde,
	0x0f, 0xab, 0x60, 0x65, 0x09, 0x79, 0x31, 0x1c, 0xdf, 0xad, 0x20, 0x79, 0xd2, 0xf3, 0x1c, 0x4e,
	0x34, 0x7f, 0xc4, 0x22, 0x9a, 0xc5, 0xcb, 0x84, 0x85, 0x73, 0x3e, 0x3f, 0x8c, 0x1a, 0x87, 0xa8,
	0xf1, 0xf9, 0x41, 0xd2, 0x3f, 0x31, 0xbc, 0xf9, 0x4a, 0xa6, 0x53, 0x0d, 0xd2, 0x6e, 0x55, 0x01,
	0xaf, 0x2d, 0x99, 0x1f, 0x28, 0x7c, 0xf8, 0xd1, 0x84, 0xed, 0xc7, 0xfc, 0x0e, 0xa6, 0x52, 0x60,
	0xad, 0xf6, 0x6d, 0x87, 0x95, 0xe9, 0xc5, 0xc7, 0xa3, 0x41, 0xc6, 0x78, 0x58, 0x26, 0x05, 0x3e,
	0xc3, 0xb2, 0xf0, 0x96, 0x76, 0x35, 0xd7, 0x9b, 0xa7, 0x9e, 0x6b, 0xea, 0x8b, 0x2e, 0x26, 0x61,
	0xf7, 0xa7, 0xe3, 0x25, 0xfd, 0x1b, 0x12, 0x3b, 0xa5, 0xff, 0xd0, 0xbb, 0x83, 0x69, 0xcf, 0xa1,
	0xdb, 0x37, 0x5b, 0x9c, 0xbc, 0x20, 0xc1, 0x18, 0x05, 0xcb, 0xc8, 0x8b, 0x67, 0xae, 0x0b, 0x48,
	0xb9, 0x55, 0x75, 0x87, 0xe4, 0x94, 0xd1, 0x85, 0xe8, 0xdf, 0xc2, 0x55, 0xdc, 0x9d, 0x33, 0xe0,
	0x56, 0x3d, 0x0c, 0xea, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x2f, 0xea, 0xe4, 0x0a, 0x03,
	0x00, 0x00,
}
