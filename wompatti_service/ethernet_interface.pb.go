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
func (*EthernetInterface) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

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
	After  uint32 `protobuf:"varint,1,opt,name=after" json:"after,omitempty"`
	Before uint32 `protobuf:"varint,2,opt,name=before" json:"before,omitempty"`
	First  uint32 `protobuf:"varint,3,opt,name=first" json:"first,omitempty"`
	Last   uint32 `protobuf:"varint,4,opt,name=last" json:"last,omitempty"`
}

func (m *FetchEthernetInterfacesRequest) Reset()                    { *m = FetchEthernetInterfacesRequest{} }
func (m *FetchEthernetInterfacesRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchEthernetInterfacesRequest) ProtoMessage()               {}
func (*FetchEthernetInterfacesRequest) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *FetchEthernetInterfacesRequest) GetAfter() uint32 {
	if m != nil {
		return m.After
	}
	return 0
}

func (m *FetchEthernetInterfacesRequest) GetBefore() uint32 {
	if m != nil {
		return m.Before
	}
	return 0
}

func (m *FetchEthernetInterfacesRequest) GetFirst() uint32 {
	if m != nil {
		return m.First
	}
	return 0
}

func (m *FetchEthernetInterfacesRequest) GetLast() uint32 {
	if m != nil {
		return m.Last
	}
	return 0
}

type EthernetInterfacesConnection struct {
	PageInfo   *PageInfo                `protobuf:"bytes,1,opt,name=pageInfo" json:"pageInfo,omitempty"`
	Edges      []*EthernetInterfaceEdge `protobuf:"bytes,2,rep,name=edges" json:"edges,omitempty"`
	TotalCount uint32                   `protobuf:"varint,3,opt,name=totalCount" json:"totalCount,omitempty"`
}

func (m *EthernetInterfacesConnection) Reset()                    { *m = EthernetInterfacesConnection{} }
func (m *EthernetInterfacesConnection) String() string            { return proto.CompactTextString(m) }
func (*EthernetInterfacesConnection) ProtoMessage()               {}
func (*EthernetInterfacesConnection) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *EthernetInterfacesConnection) GetPageInfo() *PageInfo {
	if m != nil {
		return m.PageInfo
	}
	return nil
}

func (m *EthernetInterfacesConnection) GetEdges() []*EthernetInterfaceEdge {
	if m != nil {
		return m.Edges
	}
	return nil
}

func (m *EthernetInterfacesConnection) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

type EthernetInterfaceEdge struct {
	Node   *EthernetInterface `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	Cursor uint32             `protobuf:"varint,2,opt,name=cursor" json:"cursor,omitempty"`
}

func (m *EthernetInterfaceEdge) Reset()                    { *m = EthernetInterfaceEdge{} }
func (m *EthernetInterfaceEdge) String() string            { return proto.CompactTextString(m) }
func (*EthernetInterfaceEdge) ProtoMessage()               {}
func (*EthernetInterfaceEdge) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *EthernetInterfaceEdge) GetNode() *EthernetInterface {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *EthernetInterfaceEdge) GetCursor() uint32 {
	if m != nil {
		return m.Cursor
	}
	return 0
}

func init() {
	proto.RegisterType((*EthernetInterface)(nil), "WompattiService.EthernetInterface")
	proto.RegisterType((*FetchEthernetInterfacesRequest)(nil), "WompattiService.FetchEthernetInterfacesRequest")
	proto.RegisterType((*EthernetInterfacesConnection)(nil), "WompattiService.EthernetInterfacesConnection")
	proto.RegisterType((*EthernetInterfaceEdge)(nil), "WompattiService.EthernetInterfaceEdge")
}

func init() { proto.RegisterFile("ethernet_interface.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4f, 0xe3, 0x30,
	0x10, 0xc5, 0x95, 0xf4, 0xcf, 0xee, 0x4e, 0xb7, 0xdb, 0x5d, 0x6b, 0x41, 0x06, 0xa1, 0x2a, 0xca,
	0x01, 0xe5, 0xd4, 0x43, 0x11, 0x9c, 0xb8, 0x55, 0x45, 0xea, 0x0d, 0x99, 0x03, 0xc7, 0xca, 0x4d,
	0x26, 0xa9, 0xa5, 0xd6, 0x0e, 0xce, 0x04, 0xf5, 0x9b, 0xf0, 0x41, 0xf8, 0x82, 0xc8, 0x4e, 0x8a,
	0x10, 0x41, 0xe2, 0x36, 0x6f, 0xf4, 0x7b, 0x79, 0xcf, 0xa3, 0x00, 0x47, 0xda, 0xa2, 0xd5, 0x48,
	0x6b, 0xa5, 0x09, 0x6d, 0x2e, 0x53, 0x9c, 0x95, 0xd6, 0x90, 0x61, 0x93, 0x47, 0xb3, 0x2f, 0x25,
	0x91, 0x7a, 0x40, 0xfb, 0xac, 0x52, 0x3c, 0x9f, 0x94, 0xb2, 0xc0, 0xb5, 0xd2, 0xb9, 0x69, 0x88,
	0xf8, 0x25, 0x80, 0x7f, 0xcb, 0xd6, 0xbe, 0x3a, 0xba, 0xd9, 0x1f, 0x08, 0x55, 0xc6, 0x83, 0x28,
	0x48, 0xc6, 0x22, 0x54, 0x19, 0x63, 0xd0, 0xd7, 0x72, 0x8f, 0x3c, 0x8c, 0x82, 0xe4, 0x97, 0xf0,
	0xb3, 0x67, 0x4a, 0xde, 0x8b, 0x82, 0xe4, 0xb7, 0x08, 0x55, 0xc9, 0xfe, 0x42, 0x6f, 0x2f, 0x53,
	0xde, 0xf7, 0x0b, 0x37, 0xb2, 0xff, 0x30, 0x50, 0x3a, 0xc3, 0x03, 0x1f, 0xf8, 0x0f, 0x35, 0xc2,
	0x73, 0x54, 0xf3, 0xa1, 0xdf, 0xb9, 0xd1, 0x71, 0xf9, 0x4e, 0x16, 0x15, 0xff, 0xd1, 0x70, 0x5e,
	0xc4, 0x07, 0x98, 0xde, 0x21, 0xa5, 0xdb, 0x4e, 0xbb, 0x4a, 0xe0, 0x53, 0x8d, 0x15, 0x39, 0x9f,
	0xcc, 0x09, 0x6d, 0x5b, 0xb4, 0x11, 0xec, 0x14, 0x86, 0x1b, 0xcc, 0x8d, 0x6d, 0xda, 0x8e, 0x45,
	0xab, 0x7c, 0x8a, 0xb2, 0x15, 0xf9, 0xca, 0x2e, 0xc5, 0x09, 0xf7, 0xb2, 0x9d, 0xac, 0xc8, 0xd7,
	0x1e, 0x0b, 0x3f, 0xc7, 0xaf, 0x01, 0x5c, 0x74, 0x53, 0x17, 0x46, 0x6b, 0x4c, 0x49, 0x19, 0xcd,
	0xae, 0xe1, 0xa7, 0xbb, 0xe3, 0x4a, 0xe7, 0xc6, 0x67, 0x8f, 0xe6, 0x67, 0xb3, 0x4f, 0x97, 0x9e,
	0xdd, 0xb7, 0x80, 0x78, 0x47, 0xd9, 0x2d, 0x0c, 0x30, 0x2b, 0xb0, 0xe2, 0x61, 0xd4, 0x4b, 0x46,
	0xf3, 0xcb, 0x8e, 0xa7, 0x13, 0xba, 0xcc, 0x0a, 0x14, 0x8d, 0x89, 0x4d, 0x01, 0xc8, 0x90, 0xdc,
	0x2d, 0x4c, 0xad, 0x8f, 0x8f, 0xf8, 0xb0, 0x89, 0x0b, 0x38, 0xf9, 0xd2, 0xcf, 0x6e, 0xa0, 0xaf,
	0x4d, 0x86, 0x6d, 0xd3, 0xf8, 0xfb, 0x54, 0xe1, 0x79, 0x77, 0xc8, 0xb4, 0xb6, 0x95, 0xb1, 0xc7,
	0x43, 0x36, 0x6a, 0x33, 0xf4, 0x7f, 0xce, 0xd5, 0x5b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xe2,
	0x5d, 0xc1, 0x77, 0x02, 0x00, 0x00,
}
