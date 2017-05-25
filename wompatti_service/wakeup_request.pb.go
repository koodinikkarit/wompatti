// Code generated by protoc-gen-go.
// source: wakeup_request.proto
// DO NOT EDIT!

package WompattiService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WakeupRequest struct {
	ComputerId uint32 `protobuf:"varint,1,opt,name=computerId" json:"computerId,omitempty"`
}

func (m *WakeupRequest) Reset()                    { *m = WakeupRequest{} }
func (m *WakeupRequest) String() string            { return proto.CompactTextString(m) }
func (*WakeupRequest) ProtoMessage()               {}
func (*WakeupRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *WakeupRequest) GetComputerId() uint32 {
	if m != nil {
		return m.ComputerId
	}
	return 0
}

func init() {
	proto.RegisterType((*WakeupRequest)(nil), "WompattiService.WakeupRequest")
}

func init() { proto.RegisterFile("wakeup_request.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 101 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x4f, 0xcc, 0x4e,
	0x2d, 0x2d, 0x88, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x0f, 0xcf, 0xcf, 0x2d, 0x48, 0x2c, 0x29, 0xc9, 0x0c, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0x55, 0xd2, 0xe7, 0xe2, 0x0d, 0x07, 0x2b, 0x0c, 0x82, 0xa8, 0x13, 0x92, 0xe3, 0xe2, 0x4a,
	0xce, 0xcf, 0x2d, 0x28, 0x2d, 0x49, 0x2d, 0xf2, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0d,
	0x42, 0x12, 0x49, 0x62, 0x03, 0x1b, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x63, 0xb3,
	0x59, 0x60, 0x00, 0x00, 0x00,
}
