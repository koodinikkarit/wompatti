// Code generated by protoc-gen-go.
// source: fetch_computer_by_id_response.proto
// DO NOT EDIT!

package WompattiService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FetchComputerByIdResponse struct {
}

func (m *FetchComputerByIdResponse) Reset()                    { *m = FetchComputerByIdResponse{} }
func (m *FetchComputerByIdResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchComputerByIdResponse) ProtoMessage()               {}
func (*FetchComputerByIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func init() {
	proto.RegisterType((*FetchComputerByIdResponse)(nil), "WompattiService.FetchComputerByIdResponse")
}

func init() { proto.RegisterFile("fetch_computer_by_id_response.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 100 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x4b, 0x2d, 0x49,
	0xce, 0x88, 0x4f, 0xce, 0xcf, 0x2d, 0x28, 0x2d, 0x49, 0x2d, 0x8a, 0x4f, 0xaa, 0x8c, 0xcf, 0x4c,
	0x89, 0x2f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x0f, 0xcf, 0xcf, 0x2d, 0x48, 0x2c, 0x29, 0xc9, 0x0c, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0x55, 0x92, 0xe6, 0x92, 0x74, 0x03, 0xe9, 0x73, 0x86, 0x6a, 0x73, 0xaa, 0xf4, 0x4c, 0x09, 0x82,
	0xea, 0x49, 0x62, 0x03, 0x6b, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x46, 0x79, 0x06, 0xb2,
	0x5b, 0x00, 0x00, 0x00,
}
