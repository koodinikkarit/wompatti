// Code generated by protoc-gen-go.
// source: arttu.proto
// DO NOT EDIT!

/*
Package WompattiService is a generated protocol buffer package.

It is generated from these files:
	arttu.proto
	computer.proto
	device.proto
	device_info.proto
	device_type.proto
	ethernet_interface.proto
	page_info.proto
	serial_interface.proto
	telnet_interface.proto
	wol_interface.proto
	wompatti_service.proto

It has these top-level messages:
	Arttu
	FetchArttuByComputerIdRequest
	FetchArttuByComputerIdResponse
	FetchArttuByIdRequest
	FetchArttuByIdResponse
	Computer
	ComputersEdge
	ComputersConnection
	CreateComputerRequest
	CreateComputerResponse
	EditComputerRequest
	EditComputerResponse
	RemoveComputerRequest
	RemoveComputerResponse
	FetchComputerByIdRequest
	FetchComputerByIdResponse
	FetchManyComputerByIdRequest
	FetchManyComputerByIdResponse
	FetchComputersRequest
	Device
	DevicesConnection
	DevicesEdge
	CreateDeviceRequest
	CreateDeviceResponse
	EditDeviceRequest
	EditDeviceResponse
	RemoveDeviceRequest
	RemoveDeviceResponse
	FetchDevicesRequest
	FetchDeviceByIdRequest
	FetchDeviceByIdResponse
	DeviceInfo
	KeyValue
	FetchDeviceInfoByIdRequest
	FetchDeviceInfoByIdResponse
	FetchKeyValuesByDeviceInfoIdRequest
	DeviceInfoKeyValues
	FetchKeyValuesByDeviceInfoIdResponse
	CreateKeyValueRequest
	CreateKeyValueResponse
	EditKeyValueRequest
	EditKeyValueResponse
	RemoveKeyValueRequest
	RemoveKeyValueResponse
	DeviceType
	Command
	DeviceTypesConnection
	DeviceTypesEdge
	CreateDeviceTypeRequest
	CreateDeviceTypeResponse
	EditDeviceTypeRequest
	EditDeviceTypeResponse
	RemoveDeviceTypeRequest
	RemoveDeviceTypeResponse
	FetchDeviceTypesRequest
	FetchDeviceTypeByIdRequest
	FetchDeviceTypeByIdResponse
	CreateCommandRequest
	CreateCommandResponse
	EditCommandRequest
	EditCommandReponse
	RemoveCommandRequest
	RemoveCommandResponse
	FetchCommandsByDeviceTypeIdRequest
	DeviceTypeCommands
	FetchCommandsByDeviceTypeIdResponse
	EthernetInterface
	FetchEthernetInterfacesRequest
	EthernetInterfacesConnection
	EthernetInterfaceEdge
	PageInfo
	SerialInterface
	SerialInterfacesEdge
	SerialInterfacesConnection
	FetchSerialInterfacesRequest
	FetchSerialInterfaceByIdRequest
	FetchSerialInterfaceByIdResponse
	CreateSerialInterfaceRequest
	CreateSerialInterfaceResponse
	EditSerialInterfaceRequest
	EditSerialInterfaceResponse
	RemoveSerialInterfaceRequest
	RemoveSerialInterfaceResponse
	TelnetInterface
	TelnetInterfacesConnection
	TelnetInterfacesEdge
	FetchTelnetInterfacesRequest
	FetchTelnetInterfaceByIdRequest
	FetchTelnetInterfaceByIdResponse
	CreateTelnetInterfaceRequest
	CreateTelnetInterfaceResponse
	EditTelnetInterfaceRequest
	EditTelnetInterfaceResponse
	RemoveTelnetInterfaceRequest
	RemoveTelnetInterfaceResponse
	WolInterface
	CreateWolInterfaceRequest
	CreateWolInterfaceResponse
	EditWolInterfaceRequest
	EditWolInterfaceResponse
	RemoveWolInterfaceRequest
	RemoveWolInterfaceResponse
	ExecuteWolInterfaceRequest
	ExecuteWolInterfaceResponse
	FetchWolInterfaceByIdRequest
	FetchWolInterfaceByIdResponse
*/
package WompattiService

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

type FetchArttuByComputerIdResponse_State int32

const (
	FetchArttuByComputerIdResponse_SUCCESS   FetchArttuByComputerIdResponse_State = 0
	FetchArttuByComputerIdResponse_NOT_FOUND FetchArttuByComputerIdResponse_State = 1
)

var FetchArttuByComputerIdResponse_State_name = map[int32]string{
	0: "SUCCESS",
	1: "NOT_FOUND",
}
var FetchArttuByComputerIdResponse_State_value = map[string]int32{
	"SUCCESS":   0,
	"NOT_FOUND": 1,
}

func (x FetchArttuByComputerIdResponse_State) String() string {
	return proto.EnumName(FetchArttuByComputerIdResponse_State_name, int32(x))
}
func (FetchArttuByComputerIdResponse_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0}
}

type FetchArttuByIdResponse_State int32

const (
	FetchArttuByIdResponse_SUCCESS   FetchArttuByIdResponse_State = 0
	FetchArttuByIdResponse_NOT_FOUND FetchArttuByIdResponse_State = 1
)

var FetchArttuByIdResponse_State_name = map[int32]string{
	0: "SUCCESS",
	1: "NOT_FOUND",
}
var FetchArttuByIdResponse_State_value = map[string]int32{
	"SUCCESS":   0,
	"NOT_FOUND": 1,
}

func (x FetchArttuByIdResponse_State) String() string {
	return proto.EnumName(FetchArttuByIdResponse_State_name, int32(x))
}
func (FetchArttuByIdResponse_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

type Arttu struct {
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *Arttu) Reset()                    { *m = Arttu{} }
func (m *Arttu) String() string            { return proto.CompactTextString(m) }
func (*Arttu) ProtoMessage()               {}
func (*Arttu) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Arttu) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FetchArttuByComputerIdRequest struct {
	ComputerId uint32 `protobuf:"varint,1,opt,name=computerId" json:"computerId,omitempty"`
}

func (m *FetchArttuByComputerIdRequest) Reset()                    { *m = FetchArttuByComputerIdRequest{} }
func (m *FetchArttuByComputerIdRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchArttuByComputerIdRequest) ProtoMessage()               {}
func (*FetchArttuByComputerIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FetchArttuByComputerIdRequest) GetComputerId() uint32 {
	if m != nil {
		return m.ComputerId
	}
	return 0
}

type FetchArttuByComputerIdResponse struct {
	Arttu *Arttu                               `protobuf:"bytes,1,opt,name=arttu" json:"arttu,omitempty"`
	State FetchArttuByComputerIdResponse_State `protobuf:"varint,2,opt,name=state,enum=WompattiService.FetchArttuByComputerIdResponse_State" json:"state,omitempty"`
}

func (m *FetchArttuByComputerIdResponse) Reset()                    { *m = FetchArttuByComputerIdResponse{} }
func (m *FetchArttuByComputerIdResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchArttuByComputerIdResponse) ProtoMessage()               {}
func (*FetchArttuByComputerIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FetchArttuByComputerIdResponse) GetArttu() *Arttu {
	if m != nil {
		return m.Arttu
	}
	return nil
}

func (m *FetchArttuByComputerIdResponse) GetState() FetchArttuByComputerIdResponse_State {
	if m != nil {
		return m.State
	}
	return FetchArttuByComputerIdResponse_SUCCESS
}

type FetchArttuByIdRequest struct {
	ArttuId uint32 `protobuf:"varint,1,opt,name=arttuId" json:"arttuId,omitempty"`
}

func (m *FetchArttuByIdRequest) Reset()                    { *m = FetchArttuByIdRequest{} }
func (m *FetchArttuByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchArttuByIdRequest) ProtoMessage()               {}
func (*FetchArttuByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FetchArttuByIdRequest) GetArttuId() uint32 {
	if m != nil {
		return m.ArttuId
	}
	return 0
}

type FetchArttuByIdResponse struct {
	Arttu *Arttu                       `protobuf:"bytes,1,opt,name=arttu" json:"arttu,omitempty"`
	State FetchArttuByIdResponse_State `protobuf:"varint,2,opt,name=state,enum=WompattiService.FetchArttuByIdResponse_State" json:"state,omitempty"`
}

func (m *FetchArttuByIdResponse) Reset()                    { *m = FetchArttuByIdResponse{} }
func (m *FetchArttuByIdResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchArttuByIdResponse) ProtoMessage()               {}
func (*FetchArttuByIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FetchArttuByIdResponse) GetArttu() *Arttu {
	if m != nil {
		return m.Arttu
	}
	return nil
}

func (m *FetchArttuByIdResponse) GetState() FetchArttuByIdResponse_State {
	if m != nil {
		return m.State
	}
	return FetchArttuByIdResponse_SUCCESS
}

func init() {
	proto.RegisterType((*Arttu)(nil), "WompattiService.Arttu")
	proto.RegisterType((*FetchArttuByComputerIdRequest)(nil), "WompattiService.FetchArttuByComputerIdRequest")
	proto.RegisterType((*FetchArttuByComputerIdResponse)(nil), "WompattiService.FetchArttuByComputerIdResponse")
	proto.RegisterType((*FetchArttuByIdRequest)(nil), "WompattiService.FetchArttuByIdRequest")
	proto.RegisterType((*FetchArttuByIdResponse)(nil), "WompattiService.FetchArttuByIdResponse")
	proto.RegisterEnum("WompattiService.FetchArttuByComputerIdResponse_State", FetchArttuByComputerIdResponse_State_name, FetchArttuByComputerIdResponse_State_value)
	proto.RegisterEnum("WompattiService.FetchArttuByIdResponse_State", FetchArttuByIdResponse_State_name, FetchArttuByIdResponse_State_value)
}

func init() { proto.RegisterFile("arttu.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2c, 0x2a, 0x29,
	0x29, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x0f, 0xcf, 0xcf, 0x2d, 0x48, 0x2c, 0x29,
	0xc9, 0x0c, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x55, 0x12, 0xe7, 0x62, 0x75, 0x04, 0xc9, 0x0b,
	0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x06, 0x31, 0x65, 0xa6, 0x28,
	0xd9, 0x73, 0xc9, 0xba, 0xa5, 0x96, 0x24, 0x67, 0x80, 0x65, 0x9d, 0x2a, 0x9d, 0xf3, 0x73, 0x0b,
	0x4a, 0x4b, 0x52, 0x8b, 0x3c, 0x53, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xe4, 0xb8,
	0xb8, 0x92, 0xe1, 0x82, 0x50, 0x8d, 0x48, 0x22, 0x4a, 0x07, 0x18, 0xb9, 0xe4, 0x70, 0x99, 0x50,
	0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0xa4, 0xc3, 0xc5, 0x0a, 0x76, 0x1c, 0x58, 0x37, 0xb7, 0x91,
	0x98, 0x1e, 0x9a, 0xeb, 0xf4, 0xc0, 0x5a, 0x83, 0x20, 0x8a, 0x84, 0xbc, 0xb9, 0x58, 0x8b, 0x4b,
	0x12, 0x4b, 0x52, 0x25, 0x98, 0x14, 0x18, 0x35, 0xf8, 0x8c, 0x4c, 0x31, 0x54, 0xe3, 0xb7, 0x4d,
	0x2f, 0x18, 0xa4, 0x39, 0x08, 0x62, 0x86, 0x92, 0x32, 0x17, 0x2b, 0x98, 0x2f, 0xc4, 0xcd, 0xc5,
	0x1e, 0x1c, 0xea, 0xec, 0xec, 0x1a, 0x1c, 0x2c, 0xc0, 0x20, 0xc4, 0xcb, 0xc5, 0xe9, 0xe7, 0x1f,
	0x12, 0xef, 0xe6, 0x1f, 0xea, 0xe7, 0x22, 0xc0, 0xa8, 0x64, 0xc8, 0x25, 0x8a, 0x6c, 0x26, 0xc2,
	0xef, 0x12, 0x5c, 0xec, 0x60, 0x37, 0xc1, 0x3d, 0x0e, 0xe3, 0x2a, 0x6d, 0x60, 0xe4, 0x12, 0x43,
	0xd7, 0x43, 0x96, 0x6f, 0x9d, 0x51, 0x7d, 0xab, 0x8b, 0xd7, 0xb7, 0x94, 0xf8, 0x32, 0x89, 0x0d,
	0x9c, 0x34, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x80, 0x2d, 0x73, 0x6a, 0x29, 0x02, 0x00,
	0x00,
}
