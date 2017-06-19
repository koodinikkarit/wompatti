// Code generated by protoc-gen-go.
// source: ethernet_interface.proto
// DO NOT EDIT!

package WompattiService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EthernetInterface struct {
	Id    uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Ip    []byte `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Mac   []byte `protobuf:"bytes,4,opt,name=mac,proto3" json:"mac,omitempty"`
	Index uint32 `protobuf:"varint,5,opt,name=index" json:"index,omitempty"`
	Mtu   uint32 `protobuf:"varint,6,opt,name=mtu" json:"mtu,omitempty"`
	Flags uint32 `protobuf:"varint,7,opt,name=flags" json:"flags,omitempty"`
}

func (m *EthernetInterface) Reset()                    { *m = EthernetInterface{} }
func (m *EthernetInterface) String() string            { return proto.CompactTextString(m) }
func (*EthernetInterface) ProtoMessage()               {}
func (*EthernetInterface) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *EthernetInterface) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EthernetInterface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EthernetInterface) GetIp() []byte {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *EthernetInterface) GetMac() []byte {
	if m != nil {
		return m.Mac
	}
	return nil
}

func (m *EthernetInterface) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *EthernetInterface) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *EthernetInterface) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

type FetchEthernetInterfacesRequest struct {
	Offset uint32 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Limit  uint32 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *FetchEthernetInterfacesRequest) Reset()                    { *m = FetchEthernetInterfacesRequest{} }
func (m *FetchEthernetInterfacesRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchEthernetInterfacesRequest) ProtoMessage()               {}
func (*FetchEthernetInterfacesRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *FetchEthernetInterfacesRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *FetchEthernetInterfacesRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func init() {
	proto.RegisterType((*EthernetInterface)(nil), "WompattiService.EthernetInterface")
	proto.RegisterType((*FetchEthernetInterfacesRequest)(nil), "WompattiService.FetchEthernetInterfacesRequest")
}

func init() { proto.RegisterFile("ethernet_interface.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xbf, 0x4a, 0xc4, 0x40,
	0x10, 0xc6, 0xd9, 0xdc, 0x5d, 0xc4, 0xc1, 0xf3, 0xcf, 0x22, 0x32, 0x95, 0x84, 0xab, 0x52, 0xd9,
	0xf8, 0x0c, 0x0a, 0x36, 0x16, 0x6b, 0x61, 0x29, 0x6b, 0x32, 0xf1, 0x06, 0x2e, 0xd9, 0x75, 0x77,
	0x4e, 0x7c, 0x13, 0x5f, 0x57, 0x76, 0x12, 0x2b, 0xbb, 0xef, 0x37, 0xfc, 0xf8, 0xf8, 0x18, 0x40,
	0x92, 0x3d, 0xa5, 0x89, 0xe4, 0x8d, 0x27, 0xa1, 0x34, 0xf8, 0x8e, 0xee, 0x62, 0x0a, 0x12, 0xec,
	0xc5, 0x6b, 0x18, 0xa3, 0x17, 0xe1, 0x17, 0x4a, 0x5f, 0xdc, 0xd1, 0xee, 0xc7, 0xc0, 0xd5, 0xc3,
	0x62, 0x3f, 0xfd, 0xc9, 0xf6, 0x1c, 0x2a, 0xee, 0xd1, 0x34, 0xa6, 0xdd, 0xba, 0x8a, 0x7b, 0x6b,
	0x61, 0x3d, 0xf9, 0x91, 0xb0, 0x6a, 0x4c, 0x7b, 0xea, 0x34, 0xab, 0x13, 0x71, 0xd5, 0x98, 0xf6,
	0xcc, 0x55, 0x1c, 0xed, 0x25, 0xac, 0x46, 0xdf, 0xe1, 0x5a, 0x0f, 0x25, 0xda, 0x6b, 0xd8, 0xf0,
	0xd4, 0xd3, 0x37, 0x6e, 0xb4, 0x68, 0x06, 0xf5, 0xe4, 0x88, 0xb5, 0xde, 0x4a, 0x2c, 0xde, 0x70,
	0xf0, 0x1f, 0x19, 0x4f, 0x66, 0x4f, 0x61, 0xf7, 0x0c, 0xb7, 0x8f, 0x24, 0xdd, 0xfe, 0xdf, 0xba,
	0xec, 0xe8, 0xf3, 0x48, 0x59, 0xec, 0x0d, 0xd4, 0x61, 0x18, 0x32, 0xc9, 0xb2, 0x74, 0xa1, 0xd2,
	0x77, 0xe0, 0x91, 0x45, 0xe7, 0x6e, 0xdd, 0x0c, 0xef, 0xb5, 0x7e, 0xe0, 0xfe, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x52, 0x1a, 0xb0, 0x37, 0x1d, 0x01, 0x00, 0x00,
}
