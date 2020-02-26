package utp

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Config struct {
	Version              uint32   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_326a99ff25d90470, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func init() {
	proto.RegisterType((*Config)(nil), "v2ray.core.transport.internet.headers.utp.Config")
}

func init() {
	proto.RegisterFile("v2ray.com/core/transport/internet/headers/utp/config.proto", fileDescriptor_326a99ff25d90470)
}

var fileDescriptor_326a99ff25d90470 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x2a, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x29, 0x4a, 0xcc, 0x2b,
	0x2e, 0xc8, 0x2f, 0x2a, 0xd1, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0x2d, 0xd1, 0xcf, 0x48,
	0x4d, 0x4c, 0x49, 0x2d, 0x2a, 0xd6, 0x2f, 0x2d, 0x29, 0xd0, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0x84, 0xe9, 0x2d, 0x4a, 0xd5, 0x83, 0xeb, 0xd3,
	0x83, 0xe9, 0xd3, 0x83, 0xea, 0xd3, 0x2b, 0x2d, 0x29, 0x50, 0x52, 0xe2, 0x62, 0x73, 0x06, 0x6b,
	0x15, 0x92, 0xe0, 0x62, 0x2f, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0x93, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0d, 0x82, 0x71, 0x9d, 0x92, 0xb8, 0x74, 0x93, 0xf3, 0x73, 0xf5, 0x88, 0x36, 0x34, 0x80,
	0x31, 0x8a, 0xb9, 0xb4, 0xa4, 0x60, 0x15, 0x93, 0x66, 0x98, 0x51, 0x50, 0x62, 0xa5, 0x9e, 0x33,
	0x48, 0x4b, 0x08, 0x5c, 0x8b, 0x27, 0x4c, 0x8b, 0x07, 0x54, 0x4b, 0x68, 0x49, 0x41, 0x12, 0x1b,
	0xd8, 0xe5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x86, 0xa8, 0x3b, 0xf7, 0x00, 0x00,
	0x00,
}
