// Code generated by protoc-gen-go.
// source: serial_interface.proto
// DO NOT EDIT!

package WompattiService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SerialInterface struct {
}

func (m *SerialInterface) Reset()                    { *m = SerialInterface{} }
func (m *SerialInterface) String() string            { return proto.CompactTextString(m) }
func (*SerialInterface) ProtoMessage()               {}
func (*SerialInterface) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func init() {
	proto.RegisterType((*SerialInterface)(nil), "WompattiService.SerialInterface")
}

func init() { proto.RegisterFile("serial_interface.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 80 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x4e, 0x2d, 0xca,
	0x4c, 0xcc, 0x89, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0x4a, 0x4b, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x0f, 0xcf, 0xcf, 0x2d, 0x48, 0x2c, 0x29, 0xc9, 0x0c, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x55, 0x12, 0xe4, 0xe2, 0x0f, 0x06, 0x2b, 0xf5, 0x84, 0xa9, 0x4c, 0x62, 0x03,
	0x2b, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x5e, 0x2c, 0x8c, 0x44, 0x00, 0x00, 0x00,
}