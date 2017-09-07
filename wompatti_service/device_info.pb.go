// Code generated by protoc-gen-go.
// source: device_info.proto
// DO NOT EDIT!

package WompattiService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DeviceInfo struct {
	Id uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeviceInfo) Reset()                    { *m = DeviceInfo{} }
func (m *DeviceInfo) String() string            { return proto.CompactTextString(m) }
func (*DeviceInfo) ProtoMessage()               {}
func (*DeviceInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *DeviceInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FetchDeviceInfoByIdRequest struct {
	DeviceInfoIdt []uint32 `protobuf:"varint,1,rep,packed,name=deviceInfoIdt" json:"deviceInfoIdt,omitempty"`
}

func (m *FetchDeviceInfoByIdRequest) Reset()                    { *m = FetchDeviceInfoByIdRequest{} }
func (m *FetchDeviceInfoByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchDeviceInfoByIdRequest) ProtoMessage()               {}
func (*FetchDeviceInfoByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *FetchDeviceInfoByIdRequest) GetDeviceInfoIdt() []uint32 {
	if m != nil {
		return m.DeviceInfoIdt
	}
	return nil
}

type FetchDeviceInfoByIdResponse struct {
	DeviceInfos []*DeviceInfo `protobuf:"bytes,1,rep,name=deviceInfos" json:"deviceInfos,omitempty"`
}

func (m *FetchDeviceInfoByIdResponse) Reset()                    { *m = FetchDeviceInfoByIdResponse{} }
func (m *FetchDeviceInfoByIdResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchDeviceInfoByIdResponse) ProtoMessage()               {}
func (*FetchDeviceInfoByIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *FetchDeviceInfoByIdResponse) GetDeviceInfos() []*DeviceInfo {
	if m != nil {
		return m.DeviceInfos
	}
	return nil
}

func init() {
	proto.RegisterType((*DeviceInfo)(nil), "WompattiService.DeviceInfo")
	proto.RegisterType((*FetchDeviceInfoByIdRequest)(nil), "WompattiService.FetchDeviceInfoByIdRequest")
	proto.RegisterType((*FetchDeviceInfoByIdResponse)(nil), "WompattiService.FetchDeviceInfoByIdResponse")
}

func init() { proto.RegisterFile("device_info.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x49, 0x2d, 0xcb,
	0x4c, 0x4e, 0x8d, 0xcf, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x0f,
	0xcf, 0xcf, 0x2d, 0x48, 0x2c, 0x29, 0xc9, 0x0c, 0x4e, 0x2d, 0x02, 0xc9, 0x29, 0xc9, 0x70, 0x71,
	0xb9, 0x80, 0x55, 0x79, 0xe6, 0xa5, 0xe5, 0x0b, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0xf0, 0x06, 0x31, 0x65, 0xa6, 0x28, 0x39, 0x71, 0x49, 0xb9, 0xa5, 0x96, 0x24, 0x67,
	0x20, 0x94, 0x38, 0x55, 0x7a, 0xa6, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xa9, 0x70,
	0xf1, 0xa6, 0xc0, 0x25, 0x3c, 0x53, 0x4a, 0x24, 0x18, 0x15, 0x98, 0x35, 0x78, 0x83, 0x50, 0x05,
	0x95, 0x62, 0xb8, 0xa4, 0xb1, 0x9a, 0x51, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x64, 0xcb, 0xc5,
	0x8d, 0x50, 0x5f, 0x0c, 0x36, 0x82, 0xdb, 0x48, 0x5a, 0x0f, 0xcd, 0x9d, 0x7a, 0x08, 0xdd, 0x41,
	0xc8, 0xea, 0x93, 0xd8, 0xc0, 0xfe, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x01, 0xeb, 0x21,
	0x67, 0xec, 0x00, 0x00, 0x00,
}
