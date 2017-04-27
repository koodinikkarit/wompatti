// Code generated by protoc-gen-go.
// source: computer.proto
// DO NOT EDIT!

package WompattiService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Computer struct {
	Id   uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Mac  string `protobuf:"bytes,3,opt,name=mac" json:"mac,omitempty"`
}

func (m *Computer) Reset()                    { *m = Computer{} }
func (m *Computer) String() string            { return proto.CompactTextString(m) }
func (*Computer) ProtoMessage()               {}
func (*Computer) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Computer) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Computer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Computer) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func init() {
	proto.RegisterType((*Computer)(nil), "WompattiService.Computer")
}

func init() { proto.RegisterFile("computer.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xce, 0xcf, 0x2d,
	0x28, 0x2d, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x0f, 0xcf, 0xcf, 0x2d,
	0x48, 0x2c, 0x29, 0xc9, 0x0c, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x55, 0x72, 0xe0, 0xe2, 0x70,
	0x86, 0x2a, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0d, 0x62,
	0xca, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x73, 0x13, 0x93, 0x25, 0x98, 0xc1, 0x42, 0x20, 0x66,
	0x12, 0x1b, 0xd8, 0x64, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0xfc, 0x83, 0x39, 0x6b,
	0x00, 0x00, 0x00,
}