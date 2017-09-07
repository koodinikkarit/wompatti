// Code generated by protoc-gen-go.
// source: command.proto
// DO NOT EDIT!

package WompattiService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateCommandResponse_State int32

const (
	CreateCommandResponse_SUCCESS   CreateCommandResponse_State = 0
	CreateCommandResponse_NOT_FOUND CreateCommandResponse_State = 1
)

var CreateCommandResponse_State_name = map[int32]string{
	0: "SUCCESS",
	1: "NOT_FOUND",
}
var CreateCommandResponse_State_value = map[string]int32{
	"SUCCESS":   0,
	"NOT_FOUND": 1,
}

func (x CreateCommandResponse_State) String() string {
	return proto.EnumName(CreateCommandResponse_State_name, int32(x))
}
func (CreateCommandResponse_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{8, 0}
}

type EditCommandReponse_State int32

const (
	EditCommandReponse_SUCCESS   EditCommandReponse_State = 0
	EditCommandReponse_NOT_FOUND EditCommandReponse_State = 1
)

var EditCommandReponse_State_name = map[int32]string{
	0: "SUCCESS",
	1: "NOT_FOUND",
}
var EditCommandReponse_State_value = map[string]int32{
	"SUCCESS":   0,
	"NOT_FOUND": 1,
}

func (x EditCommandReponse_State) String() string {
	return proto.EnumName(EditCommandReponse_State_name, int32(x))
}
func (EditCommandReponse_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{10, 0} }

type RemoveCommandResponse_State int32

const (
	RemoveCommandResponse_SUCCESS   RemoveCommandResponse_State = 0
	RemoveCommandResponse_NOT_FOUND RemoveCommandResponse_State = 1
)

var RemoveCommandResponse_State_name = map[int32]string{
	0: "SUCCESS",
	1: "NOT_FOUND",
}
var RemoveCommandResponse_State_value = map[string]int32{
	"SUCCESS":   0,
	"NOT_FOUND": 1,
}

func (x RemoveCommandResponse_State) String() string {
	return proto.EnumName(RemoveCommandResponse_State_name, int32(x))
}
func (RemoveCommandResponse_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{12, 0}
}

type Command struct {
	Id           uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	DeviceTypeId uint32 `protobuf:"varint,2,opt,name=deviceTypeId" json:"deviceTypeId,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Value        string `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Command) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Command) GetDeviceTypeId() uint32 {
	if m != nil {
		return m.DeviceTypeId
	}
	return 0
}

func (m *Command) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Command) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type FetchCommandsRequest struct {
}

func (m *FetchCommandsRequest) Reset()                    { *m = FetchCommandsRequest{} }
func (m *FetchCommandsRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchCommandsRequest) ProtoMessage()               {}
func (*FetchCommandsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type FetchCommandsResponse struct {
}

func (m *FetchCommandsResponse) Reset()                    { *m = FetchCommandsResponse{} }
func (m *FetchCommandsResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchCommandsResponse) ProtoMessage()               {}
func (*FetchCommandsResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type FetchCommandsByDeviceTypeIdRequest struct {
}

func (m *FetchCommandsByDeviceTypeIdRequest) Reset()         { *m = FetchCommandsByDeviceTypeIdRequest{} }
func (m *FetchCommandsByDeviceTypeIdRequest) String() string { return proto.CompactTextString(m) }
func (*FetchCommandsByDeviceTypeIdRequest) ProtoMessage()    {}
func (*FetchCommandsByDeviceTypeIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{3}
}

type FetchCommandsByDeviceTypeIdResponse struct {
}

func (m *FetchCommandsByDeviceTypeIdResponse) Reset()         { *m = FetchCommandsByDeviceTypeIdResponse{} }
func (m *FetchCommandsByDeviceTypeIdResponse) String() string { return proto.CompactTextString(m) }
func (*FetchCommandsByDeviceTypeIdResponse) ProtoMessage()    {}
func (*FetchCommandsByDeviceTypeIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{4}
}

type FetchCommandByIdRequest struct {
}

func (m *FetchCommandByIdRequest) Reset()                    { *m = FetchCommandByIdRequest{} }
func (m *FetchCommandByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchCommandByIdRequest) ProtoMessage()               {}
func (*FetchCommandByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type FetchCommandByIdResponse struct {
}

func (m *FetchCommandByIdResponse) Reset()                    { *m = FetchCommandByIdResponse{} }
func (m *FetchCommandByIdResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchCommandByIdResponse) ProtoMessage()               {}
func (*FetchCommandByIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

type CreateCommandRequest struct {
	DeviceTypeId uint32 `protobuf:"varint,1,opt,name=deviceTypeId" json:"deviceTypeId,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Value        string `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
}

func (m *CreateCommandRequest) Reset()                    { *m = CreateCommandRequest{} }
func (m *CreateCommandRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateCommandRequest) ProtoMessage()               {}
func (*CreateCommandRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *CreateCommandRequest) GetDeviceTypeId() uint32 {
	if m != nil {
		return m.DeviceTypeId
	}
	return 0
}

func (m *CreateCommandRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateCommandRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CreateCommandResponse struct {
	State   CreateCommandResponse_State `protobuf:"varint,1,opt,name=state,enum=WompattiService.CreateCommandResponse_State" json:"state,omitempty"`
	Command *Command                    `protobuf:"bytes,2,opt,name=command" json:"command,omitempty"`
}

func (m *CreateCommandResponse) Reset()                    { *m = CreateCommandResponse{} }
func (m *CreateCommandResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateCommandResponse) ProtoMessage()               {}
func (*CreateCommandResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *CreateCommandResponse) GetState() CreateCommandResponse_State {
	if m != nil {
		return m.State
	}
	return CreateCommandResponse_SUCCESS
}

func (m *CreateCommandResponse) GetCommand() *Command {
	if m != nil {
		return m.Command
	}
	return nil
}

type EditCommandRequest struct {
	CommandId uint32 `protobuf:"varint,1,opt,name=commandId" json:"commandId,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Value     string `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
}

func (m *EditCommandRequest) Reset()                    { *m = EditCommandRequest{} }
func (m *EditCommandRequest) String() string            { return proto.CompactTextString(m) }
func (*EditCommandRequest) ProtoMessage()               {}
func (*EditCommandRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *EditCommandRequest) GetCommandId() uint32 {
	if m != nil {
		return m.CommandId
	}
	return 0
}

func (m *EditCommandRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EditCommandRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type EditCommandReponse struct {
	State   EditCommandReponse_State `protobuf:"varint,1,opt,name=state,enum=WompattiService.EditCommandReponse_State" json:"state,omitempty"`
	Command *Command                 `protobuf:"bytes,2,opt,name=command" json:"command,omitempty"`
}

func (m *EditCommandReponse) Reset()                    { *m = EditCommandReponse{} }
func (m *EditCommandReponse) String() string            { return proto.CompactTextString(m) }
func (*EditCommandReponse) ProtoMessage()               {}
func (*EditCommandReponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *EditCommandReponse) GetState() EditCommandReponse_State {
	if m != nil {
		return m.State
	}
	return EditCommandReponse_SUCCESS
}

func (m *EditCommandReponse) GetCommand() *Command {
	if m != nil {
		return m.Command
	}
	return nil
}

type RemoveCommandRequest struct {
	CommandId uint32 `protobuf:"varint,1,opt,name=commandId" json:"commandId,omitempty"`
}

func (m *RemoveCommandRequest) Reset()                    { *m = RemoveCommandRequest{} }
func (m *RemoveCommandRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveCommandRequest) ProtoMessage()               {}
func (*RemoveCommandRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *RemoveCommandRequest) GetCommandId() uint32 {
	if m != nil {
		return m.CommandId
	}
	return 0
}

type RemoveCommandResponse struct {
	State RemoveCommandResponse_State `protobuf:"varint,1,opt,name=state,enum=WompattiService.RemoveCommandResponse_State" json:"state,omitempty"`
}

func (m *RemoveCommandResponse) Reset()                    { *m = RemoveCommandResponse{} }
func (m *RemoveCommandResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveCommandResponse) ProtoMessage()               {}
func (*RemoveCommandResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *RemoveCommandResponse) GetState() RemoveCommandResponse_State {
	if m != nil {
		return m.State
	}
	return RemoveCommandResponse_SUCCESS
}

func init() {
	proto.RegisterType((*Command)(nil), "WompattiService.Command")
	proto.RegisterType((*FetchCommandsRequest)(nil), "WompattiService.FetchCommandsRequest")
	proto.RegisterType((*FetchCommandsResponse)(nil), "WompattiService.FetchCommandsResponse")
	proto.RegisterType((*FetchCommandsByDeviceTypeIdRequest)(nil), "WompattiService.FetchCommandsByDeviceTypeIdRequest")
	proto.RegisterType((*FetchCommandsByDeviceTypeIdResponse)(nil), "WompattiService.FetchCommandsByDeviceTypeIdResponse")
	proto.RegisterType((*FetchCommandByIdRequest)(nil), "WompattiService.FetchCommandByIdRequest")
	proto.RegisterType((*FetchCommandByIdResponse)(nil), "WompattiService.FetchCommandByIdResponse")
	proto.RegisterType((*CreateCommandRequest)(nil), "WompattiService.CreateCommandRequest")
	proto.RegisterType((*CreateCommandResponse)(nil), "WompattiService.CreateCommandResponse")
	proto.RegisterType((*EditCommandRequest)(nil), "WompattiService.EditCommandRequest")
	proto.RegisterType((*EditCommandReponse)(nil), "WompattiService.EditCommandReponse")
	proto.RegisterType((*RemoveCommandRequest)(nil), "WompattiService.RemoveCommandRequest")
	proto.RegisterType((*RemoveCommandResponse)(nil), "WompattiService.RemoveCommandResponse")
	proto.RegisterEnum("WompattiService.CreateCommandResponse_State", CreateCommandResponse_State_name, CreateCommandResponse_State_value)
	proto.RegisterEnum("WompattiService.EditCommandReponse_State", EditCommandReponse_State_name, EditCommandReponse_State_value)
	proto.RegisterEnum("WompattiService.RemoveCommandResponse_State", RemoveCommandResponse_State_name, RemoveCommandResponse_State_value)
}

func init() { proto.RegisterFile("command.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xd1, 0x4b, 0xfa, 0x50,
	0x14, 0xc7, 0x7f, 0xd7, 0x9f, 0x26, 0x1e, 0xd3, 0xe4, 0x32, 0x73, 0x45, 0x0f, 0x72, 0x2d, 0x30,
	0x88, 0x3d, 0x58, 0xef, 0x81, 0x53, 0xa1, 0x17, 0x85, 0x4d, 0xe9, 0x25, 0x88, 0xe5, 0x3d, 0xd4,
	0xa0, 0xb9, 0xe5, 0xae, 0x82, 0x6f, 0xfd, 0x45, 0x3d, 0xf5, 0x07, 0x86, 0x77, 0xd7, 0x74, 0x9b,
	0x84, 0x12, 0xf4, 0xb6, 0x9d, 0x73, 0x3e, 0xdf, 0xef, 0xee, 0xf7, 0x8c, 0x0b, 0xa5, 0xb1, 0xef,
	0x79, 0xce, 0x84, 0x1b, 0xc1, 0xd4, 0x17, 0x3e, 0x3d, 0xba, 0xf7, 0xbd, 0xc0, 0x11, 0xc2, 0xb5,
	0x71, 0x3a, 0x77, 0xc7, 0xc8, 0x9e, 0x21, 0x6f, 0x46, 0x13, 0xb4, 0x0c, 0x19, 0x97, 0xeb, 0xa4,
	0x4e, 0x9a, 0x25, 0x2b, 0xe3, 0x72, 0xca, 0xe0, 0x90, 0xe3, 0x72, 0x68, 0xb8, 0x08, 0xf0, 0x8e,
	0xeb, 0x19, 0xd9, 0x89, 0xd5, 0x28, 0x85, 0xec, 0xc4, 0xf1, 0x50, 0xff, 0x5f, 0x27, 0xcd, 0x82,
	0x25, 0x9f, 0xa9, 0x06, 0xb9, 0xb9, 0xf3, 0x3a, 0x43, 0x3d, 0x2b, 0x8b, 0xd1, 0x0b, 0x3b, 0x06,
	0xad, 0x87, 0x62, 0xfc, 0xa2, 0xdc, 0x42, 0x0b, 0xdf, 0x66, 0x18, 0x0a, 0x56, 0x83, 0x6a, 0xa2,
	0x1e, 0x06, 0xfe, 0x24, 0x44, 0x76, 0x0e, 0x2c, 0xd6, 0x68, 0x2f, 0x3a, 0x1b, 0xce, 0x2b, 0xfc,
	0x02, 0x1a, 0x3f, 0x4e, 0x29, 0xb1, 0x13, 0xa8, 0x6d, 0x8e, 0xb5, 0x17, 0x6b, 0x85, 0x53, 0xd0,
	0xd3, 0x2d, 0x85, 0x71, 0xd0, 0xcc, 0x29, 0x3a, 0x02, 0x55, 0x53, 0x31, 0xa9, 0x68, 0xc8, 0xaf,
	0xa2, 0xf9, 0x24, 0x50, 0x4d, 0xd8, 0x44, 0xfe, 0xb4, 0x0d, 0xb9, 0x50, 0x38, 0x02, 0xa5, 0x41,
	0xb9, 0x75, 0x65, 0x24, 0xd6, 0x67, 0x6c, 0xc5, 0x0c, 0x7b, 0xc9, 0x58, 0x11, 0x4a, 0x5b, 0x90,
	0x57, 0xff, 0x80, 0xdc, 0x60, 0xb1, 0xa5, 0xa7, 0x55, 0x14, 0xbf, 0x1a, 0x64, 0x0d, 0xc8, 0x49,
	0x0d, 0x5a, 0x84, 0xbc, 0x3d, 0x32, 0xcd, 0xae, 0x6d, 0x57, 0xfe, 0xd1, 0x12, 0x14, 0xfa, 0x83,
	0xe1, 0x63, 0x6f, 0x30, 0xea, 0x77, 0x2a, 0x84, 0x3d, 0x00, 0xed, 0x72, 0x57, 0x24, 0xa2, 0x39,
	0x83, 0x82, 0x52, 0xf9, 0xce, 0x65, 0x5d, 0xd8, 0x23, 0x94, 0x0f, 0x92, 0x90, 0x8f, 0x12, 0xb9,
	0x8d, 0x27, 0x72, 0x99, 0x3a, 0x4b, 0x9a, 0xf9, 0xa3, 0x38, 0x6e, 0x40, 0xb3, 0xd0, 0xf3, 0xe7,
	0xb8, 0x4f, 0x20, 0xec, 0x9d, 0x40, 0x35, 0x81, 0xed, 0xba, 0xfb, 0xad, 0x58, 0xec, 0xb0, 0x3b,
	0x7d, 0xf8, 0xd3, 0x81, 0xbc, 0x1a, 0xae, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x07, 0x55, 0x9a,
	0x0d, 0x2b, 0x04, 0x00, 0x00,
}
