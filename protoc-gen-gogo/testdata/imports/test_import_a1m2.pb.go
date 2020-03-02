// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imports/test_import_a1m2.proto

package imports

import (
	fmt "fmt"
	math "math"

	proto "github.com/bamzi/protobuf/proto"
	test_a_1 "github.com/bamzi/protobuf/protoc-gen-gogo/testdata/imports/test_a_1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type A1M2 struct {
	F                    *test_a_1.M2 `protobuf:"bytes,1,opt,name=f,proto3" json:"f,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *A1M2) Reset()         { *m = A1M2{} }
func (m *A1M2) String() string { return proto.CompactTextString(m) }
func (*A1M2) ProtoMessage()    {}
func (*A1M2) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdb27b114687957d, []int{0}
}
func (m *A1M2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_A1M2.Unmarshal(m, b)
}
func (m *A1M2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_A1M2.Marshal(b, m, deterministic)
}
func (m *A1M2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_A1M2.Merge(m, src)
}
func (m *A1M2) XXX_Size() int {
	return xxx_messageInfo_A1M2.Size(m)
}
func (m *A1M2) XXX_DiscardUnknown() {
	xxx_messageInfo_A1M2.DiscardUnknown(m)
}

var xxx_messageInfo_A1M2 proto.InternalMessageInfo

func (m *A1M2) GetF() *test_a_1.M2 {
	if m != nil {
		return m.F
	}
	return nil
}

func init() {
	proto.RegisterType((*A1M2)(nil), "test.A1M2")
}

func init() { proto.RegisterFile("imports/test_import_a1m2.proto", fileDescriptor_bdb27b114687957d) }

var fileDescriptor_bdb27b114687957d = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x29, 0xd6, 0x2f, 0x49, 0x2d, 0x2e, 0x89, 0x87, 0x70, 0xe2, 0x13, 0x0d, 0x73, 0x8d,
	0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x58, 0x40, 0xe2, 0x52, 0x92, 0x28, 0xaa, 0x12, 0xe3,
	0x0d, 0xf5, 0x61, 0x0a, 0x94, 0x14, 0xb8, 0x58, 0x1c, 0x0d, 0x7d, 0x8d, 0x84, 0x24, 0xb8, 0x18,
	0xd3, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xb8, 0xf4, 0x40, 0xca, 0xf4, 0x12, 0xf5, 0x7c,
	0x8d, 0x82, 0x18, 0xd3, 0x9c, 0xac, 0xa3, 0x2c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92,
	0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0x5a, 0x93, 0x4a, 0xd3, 0x20, 0x8c, 0x64,
	0xdd, 0xf4, 0xd4, 0x3c, 0x5d, 0xb0, 0x04, 0x48, 0x63, 0x4a, 0x62, 0x49, 0xa2, 0x3e, 0xd4, 0xc2,
	0x24, 0x36, 0xb0, 0x0a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x65, 0x04, 0x17, 0xa8,
	0x00, 0x00, 0x00,
}
