// Code generated by protoc-gen-go.
// source: page_info.proto
// DO NOT EDIT!

package WompattiService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PageInfo struct {
	HasNextPage     bool   `protobuf:"varint,1,opt,name=hasNextPage" json:"hasNextPage,omitempty"`
	HasPreviousPage bool   `protobuf:"varint,2,opt,name=hasPreviousPage" json:"hasPreviousPage,omitempty"`
	StartCursor     uint32 `protobuf:"varint,3,opt,name=startCursor" json:"startCursor,omitempty"`
	EndCursor       uint32 `protobuf:"varint,4,opt,name=endCursor" json:"endCursor,omitempty"`
}

func (m *PageInfo) Reset()                    { *m = PageInfo{} }
func (m *PageInfo) String() string            { return proto.CompactTextString(m) }
func (*PageInfo) ProtoMessage()               {}
func (*PageInfo) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *PageInfo) GetHasNextPage() bool {
	if m != nil {
		return m.HasNextPage
	}
	return false
}

func (m *PageInfo) GetHasPreviousPage() bool {
	if m != nil {
		return m.HasPreviousPage
	}
	return false
}

func (m *PageInfo) GetStartCursor() uint32 {
	if m != nil {
		return m.StartCursor
	}
	return 0
}

func (m *PageInfo) GetEndCursor() uint32 {
	if m != nil {
		return m.EndCursor
	}
	return 0
}

func init() {
	proto.RegisterType((*PageInfo)(nil), "WompattiService.PageInfo")
}

func init() { proto.RegisterFile("page_info.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x48, 0x4c, 0x4f,
	0x8d, 0xcf, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x0f, 0xcf, 0xcf,
	0x2d, 0x48, 0x2c, 0x29, 0xc9, 0x0c, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x55, 0x9a, 0xc6, 0xc8,
	0xc5, 0x11, 0x90, 0x98, 0x9e, 0xea, 0x99, 0x97, 0x96, 0x2f, 0xa4, 0xc0, 0xc5, 0x9d, 0x91, 0x58,
	0xec, 0x97, 0x5a, 0x51, 0x02, 0x12, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x08, 0x42, 0x16, 0x12,
	0xd2, 0xe0, 0xe2, 0xcf, 0x48, 0x2c, 0x0e, 0x28, 0x4a, 0x2d, 0xcb, 0xcc, 0x2f, 0x2d, 0x06, 0xab,
	0x62, 0x02, 0xab, 0x42, 0x17, 0x06, 0x99, 0x55, 0x5c, 0x92, 0x58, 0x54, 0xe2, 0x5c, 0x5a, 0x54,
	0x9c, 0x5f, 0x24, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x1b, 0x84, 0x2c, 0x24, 0x24, 0xc3, 0xc5, 0x99,
	0x9a, 0x97, 0x02, 0x95, 0x67, 0x01, 0xcb, 0x23, 0x04, 0x92, 0xd8, 0xc0, 0x0e, 0x36, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xfa, 0x8b, 0x5d, 0xfa, 0xc3, 0x00, 0x00, 0x00,
}
