// Code generated by protoc-gen-go.
// source: severi.proto
// DO NOT EDIT!

package WompattiService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Severi struct {
	Id   uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Ip   string `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
	Port uint32 `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
}

func (m *Severi) Reset()                    { *m = Severi{} }
func (m *Severi) String() string            { return proto.CompactTextString(m) }
func (*Severi) ProtoMessage()               {}
func (*Severi) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *Severi) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Severi) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Severi) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type FetchSeverisRequest struct {
}

func (m *FetchSeverisRequest) Reset()                    { *m = FetchSeverisRequest{} }
func (m *FetchSeverisRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchSeverisRequest) ProtoMessage()               {}
func (*FetchSeverisRequest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

type FetchSeverisResponse struct {
}

func (m *FetchSeverisResponse) Reset()                    { *m = FetchSeverisResponse{} }
func (m *FetchSeverisResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchSeverisResponse) ProtoMessage()               {}
func (*FetchSeverisResponse) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

type FetchSeveriByIdRequest struct {
}

func (m *FetchSeveriByIdRequest) Reset()                    { *m = FetchSeveriByIdRequest{} }
func (m *FetchSeveriByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchSeveriByIdRequest) ProtoMessage()               {}
func (*FetchSeveriByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{3} }

type FetchSeveriByIdResponse struct {
}

func (m *FetchSeveriByIdResponse) Reset()                    { *m = FetchSeveriByIdResponse{} }
func (m *FetchSeveriByIdResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchSeveriByIdResponse) ProtoMessage()               {}
func (*FetchSeveriByIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{4} }

type CreateSeveriRequest struct {
}

func (m *CreateSeveriRequest) Reset()                    { *m = CreateSeveriRequest{} }
func (m *CreateSeveriRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateSeveriRequest) ProtoMessage()               {}
func (*CreateSeveriRequest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{5} }

type CreateSeveriResponse struct {
}

func (m *CreateSeveriResponse) Reset()                    { *m = CreateSeveriResponse{} }
func (m *CreateSeveriResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateSeveriResponse) ProtoMessage()               {}
func (*CreateSeveriResponse) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{6} }

type EditSeveriRequest struct {
}

func (m *EditSeveriRequest) Reset()                    { *m = EditSeveriRequest{} }
func (m *EditSeveriRequest) String() string            { return proto.CompactTextString(m) }
func (*EditSeveriRequest) ProtoMessage()               {}
func (*EditSeveriRequest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{7} }

type EditSeveriResponse struct {
}

func (m *EditSeveriResponse) Reset()                    { *m = EditSeveriResponse{} }
func (m *EditSeveriResponse) String() string            { return proto.CompactTextString(m) }
func (*EditSeveriResponse) ProtoMessage()               {}
func (*EditSeveriResponse) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{8} }

type RemoveSeveriRequest struct {
}

func (m *RemoveSeveriRequest) Reset()                    { *m = RemoveSeveriRequest{} }
func (m *RemoveSeveriRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveSeveriRequest) ProtoMessage()               {}
func (*RemoveSeveriRequest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{9} }

type RemoveSeveriResponse struct {
}

func (m *RemoveSeveriResponse) Reset()                    { *m = RemoveSeveriResponse{} }
func (m *RemoveSeveriResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveSeveriResponse) ProtoMessage()               {}
func (*RemoveSeveriResponse) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{10} }

func init() {
	proto.RegisterType((*Severi)(nil), "WompattiService.Severi")
	proto.RegisterType((*FetchSeverisRequest)(nil), "WompattiService.FetchSeverisRequest")
	proto.RegisterType((*FetchSeverisResponse)(nil), "WompattiService.FetchSeverisResponse")
	proto.RegisterType((*FetchSeveriByIdRequest)(nil), "WompattiService.FetchSeveriByIdRequest")
	proto.RegisterType((*FetchSeveriByIdResponse)(nil), "WompattiService.FetchSeveriByIdResponse")
	proto.RegisterType((*CreateSeveriRequest)(nil), "WompattiService.CreateSeveriRequest")
	proto.RegisterType((*CreateSeveriResponse)(nil), "WompattiService.CreateSeveriResponse")
	proto.RegisterType((*EditSeveriRequest)(nil), "WompattiService.EditSeveriRequest")
	proto.RegisterType((*EditSeveriResponse)(nil), "WompattiService.EditSeveriResponse")
	proto.RegisterType((*RemoveSeveriRequest)(nil), "WompattiService.RemoveSeveriRequest")
	proto.RegisterType((*RemoveSeveriResponse)(nil), "WompattiService.RemoveSeveriResponse")
}

func init() { proto.RegisterFile("severi.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x3b, 0x6b, 0xc3, 0x30,
	0x14, 0x85, 0xb1, 0x5a, 0x0c, 0xbd, 0xf4, 0x41, 0xe5, 0x47, 0xd5, 0xcd, 0x68, 0xf2, 0xd4, 0xa5,
	0x6b, 0xa7, 0x96, 0x16, 0xba, 0xca, 0x43, 0x67, 0xd7, 0xba, 0x10, 0x0d, 0x8e, 0x14, 0xe9, 0xc6,
	0x90, 0x7f, 0x1f, 0x90, 0x1d, 0x62, 0xe1, 0x4d, 0xfa, 0xce, 0x77, 0x8e, 0x40, 0x70, 0x1f, 0x70,
	0x42, 0x6f, 0xde, 0x9c, 0xb7, 0x64, 0xf9, 0xd3, 0x9f, 0x1d, 0x5d, 0x4f, 0x64, 0x3a, 0xf4, 0x93,
	0x19, 0x50, 0x7e, 0x40, 0xde, 0x45, 0x81, 0x3f, 0x02, 0x33, 0x5a, 0x64, 0x4d, 0xd6, 0x3e, 0x28,
	0x66, 0x74, 0xbc, 0x3b, 0xc1, 0x9a, 0xac, 0xbd, 0x53, 0xcc, 0x38, 0xce, 0xe1, 0xd6, 0x59, 0x4f,
	0xe2, 0x26, 0x1a, 0xf1, 0x2c, 0x2b, 0x28, 0x7e, 0x90, 0x86, 0xdd, 0x3c, 0x11, 0x14, 0x1e, 0x8e,
	0x18, 0x48, 0xd6, 0x50, 0xa6, 0x38, 0x38, 0xbb, 0x0f, 0x28, 0x05, 0xd4, 0x2b, 0xfe, 0x79, 0xfa,
	0xd5, 0x97, 0xc6, 0x2b, 0xbc, 0x6c, 0x92, 0xa5, 0x54, 0x41, 0xf1, 0xe5, 0xb1, 0x27, 0x9c, 0xb3,
	0xd5, 0x1b, 0x29, 0x5e, 0xf4, 0x02, 0x9e, 0xbf, 0xb5, 0xa1, 0x54, 0x2e, 0x81, 0xaf, 0xe1, 0x75,
	0x59, 0xe1, 0x68, 0xa7, 0xed, 0x72, 0x8a, 0x67, 0xfd, 0x3f, 0x8f, 0x5f, 0xf8, 0x7e, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x07, 0x5c, 0xae, 0x68, 0x52, 0x01, 0x00, 0x00,
}
