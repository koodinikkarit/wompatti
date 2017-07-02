package wompatti

import (
	"golang.org/x/net/context"

	wompatti "github.com/koodinikkarit/wompatti/db"
	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"
)

func (s *WompattiServiceServer) CreateComputer(ctx context.Context, in *WompattiService.CreateComputerRequest) (*WompattiService.CreateComputerResponse, error) {
	deviceInfo := &wompatti.DeviceInfo{}
	s.db.Create(deviceInfo)

	computer := &wompatti.Computer{
		Name:         in.Name,
		DeviceInfoID: uint32(deviceInfo.ID),
	}
	s.db.Create(computer)

	return &WompattiService.CreateComputerResponse{
		Computer: &WompattiService.Computer{
			Id:   computer.ID,
			Name: computer.Name,
		},
		DeviceInfo: &WompattiService.DeviceInfo{
			Id: uint32(deviceInfo.ID),
		},
	}, nil
}

func (s *WompattiServiceServer) EditComputer(ctx context.Context, in *WompattiService.EditComputerRequest) (*WompattiService.EditComputerResponse, error) {
	var computer wompatti.Computer
	s.db.First(&computer, in.ComputerId)
	if in.Name != "" {
		computer.Name = in.Name
	}
	if in.ArttuId != 0 {
		computer.ArttuID = in.ArttuId
	}
	if in.DeviceInfoId != 0 {
		computer.DeviceInfoID = in.DeviceInfoId
	}

	s.db.Save(&computer)
	return &WompattiService.EditComputerResponse{
		State: WompattiService.EditComputerResponse_SUCCESS,
		Computer: &WompattiService.Computer{
			Id:           computer.ID,
			Name:         computer.Name,
			ArttuId:      computer.ArttuID,
			DeviceInfoId: computer.DeviceInfoID,
		},
	}, nil
}

func (s *WompattiServiceServer) RemoveComputer(ctx context.Context, in *WompattiService.RemoveComputerRequest) (*WompattiService.RemoveComputerResponse, error) {
	response := &WompattiService.RemoveComputerResponse{}
	computer := &wompatti.Computer{ID: in.ComputerId}

	s.db.First(computer, in.ComputerId)
	if computer.ID != 0 {
		response.State = WompattiService.RemoveComputerResponse_SUCCESS
		s.db.Delete(computer)
	} else {
		response.State = WompattiService.RemoveComputerResponse_NOT_FOUND
	}

	return response, nil
}

func (s *WompattiServiceServer) CreateDevice(ctx context.Context, in *WompattiService.CreateDeviceRequest) (*WompattiService.CreateDeviceResponse, error) {
	deviceInfo := &wompatti.DeviceInfo{}

	s.db.Create(deviceInfo)

	device := &wompatti.Device{
		DeviceInfoID: uint32(deviceInfo.ID),
		Name:         in.Name,
	}

	s.db.Create(device)

	return &WompattiService.CreateDeviceResponse{
		Device: &WompattiService.Device{
			Id:           device.ID,
			DeviceInfoId: device.DeviceInfoID,
			Name:         device.Name,
		},
		DeviceInfo: &WompattiService.DeviceInfo{
			Id: uint32(deviceInfo.ID),
		},
	}, nil
}

func (s *WompattiServiceServer) EditDevice(ctx context.Context, in *WompattiService.EditDeviceRequest) (*WompattiService.EditDeviceResponse, error) {
	res := &WompattiService.EditDeviceResponse{}

	device := &wompatti.Device{}

	s.db.Find(&device, in.DeviceId)

	if device.ID > 0 {
		res.State = WompattiService.EditDeviceResponse_SUCCESS

		if in.Name != "" {
			device.Name = in.Name
		}

		if in.DeviceTypeId != 0 {
			device.DeviceTypeID = in.DeviceTypeId
		}

		if in.SerialInterfaceId != 0 {
			device.SerialInterfaceID = in.SerialInterfaceId
		}

		if in.TelnetInterfaceId != 0 {
			device.TelnetInterfaceID = in.TelnetInterfaceId
		}

		if in.CecInterfaceId != 0 {
			device.CecInterfaceID = in.CecInterfaceId
		}

		s.db.Save(&device)

		res.Device = &WompattiService.Device{
			Name:              device.Name,
			DeviceTypeId:      device.DeviceTypeID,
			SerialInterfaceId: device.SerialInterfaceID,
			TelnetInterfaceId: device.TelnetInterfaceID,
			CecInterfaceId:    device.CecInterfaceID,
		}
	} else {
		res.State = WompattiService.EditDeviceResponse_NOT_FOUND
	}

	return res, nil
}

func (s *WompattiServiceServer) RemoveDevice(ctx context.Context, in *WompattiService.RemoveDeviceRequest) (*WompattiService.RemoveDeviceResponse, error) {
	response := &WompattiService.RemoveDeviceResponse{}
	device := &wompatti.Device{}
	s.db.First(device, in.DeviceId)
	if device.ID != 0 {
		response.State = WompattiService.RemoveDeviceResponse_SUCCESS
		s.db.Delete(device)
	} else {
		response.State = WompattiService.RemoveDeviceResponse_NOT_FOUND
	}

	return response, nil
}

func (s *WompattiServiceServer) CreateDeviceType(ctx context.Context, in *WompattiService.CreateDeviceTypeRequest) (*WompattiService.CreateDeviceTypeResponse, error) {
	res := &WompattiService.CreateDeviceTypeResponse{}

	deviceType := wompatti.DeviceType{
		Name: in.Name,
	}

	s.db.Create(&deviceType)

	res.DeviceType = GenerateServiceDeviceType(&deviceType)

	return res, nil
}

func (s *WompattiServiceServer) EditDeviceType(ctx context.Context, in *WompattiService.EditDeviceTypeRequest) (*WompattiService.EditDeviceTypeResponse, error) {
	res := &WompattiService.EditDeviceTypeResponse{}

	deviceType := &wompatti.DeviceType{}

	s.db.First(&deviceType, in.DeviceTypeId)

	if deviceType.ID > 0 {

		if in.Name != "" {
			deviceType.Name = in.Name
		}

		s.db.Save(&deviceType)

		res.State = WompattiService.EditDeviceTypeResponse_SUCCESS
		res.DeviceType = GenerateServiceDeviceType(deviceType)
	} else {
		res.State = WompattiService.EditDeviceTypeResponse_NOT_FOUND
	}

	return res, nil
}

func (s *WompattiServiceServer) RemoveDeviceType(ctx context.Context, in *WompattiService.RemoveDeviceTypeRequest) (*WompattiService.RemoveDeviceTypeResponse, error) {
	res := &WompattiService.RemoveDeviceTypeResponse{}

	deviceType := &wompatti.DeviceType{}

	s.db.First(&deviceType, in.DeviceTypeId)

	if deviceType.ID > 0 {
		res.State = WompattiService.RemoveDeviceTypeResponse_SUCCESS
		s.db.Delete(&deviceType)
	} else {
		res.State = WompattiService.RemoveDeviceTypeResponse_NOT_FOUND
	}

	return nil, nil
}

func (s *WompattiServiceServer) CreateCommand(ctx context.Context, in *WompattiService.CreateCommandRequest) (*WompattiService.CreateCommandResponse, error) {
	res := &WompattiService.CreateCommandResponse{}

	command := &wompatti.Command{
		DeviceTypeID: in.DeviceTypeId,
		Name:         in.Name,
		Value:        in.Value,
	}

	s.db.Create(&command)

	res.Command = GenerateServiceCommand(command)

	return res, nil
}

func (s *WompattiServiceServer) EditCommand(ctx context.Context, in *WompattiService.EditCommandRequest) (*WompattiService.EditCommandReponse, error) {
	res := &WompattiService.EditCommandReponse{}

	command := &wompatti.Command{}

	s.db.First(&command, in.CommandId)

	if command.ID > 0 {
		if in.Name != "" {
			command.Name = in.Name
		}

		if in.Value != "" {
			command.Value = in.Value
		}

		s.db.Save(&command)

		res.State = WompattiService.EditCommandReponse_SUCCESS
		res.Command = GenerateServiceCommand(command)
	} else {
		res.State = WompattiService.EditCommandReponse_NOT_FOUND
	}

	return res, nil
}

func (s *WompattiServiceServer) RemoveCommand(ctx context.Context, in *WompattiService.RemoveCommandRequest) (*WompattiService.RemoveCommandResponse, error) {
	res := &WompattiService.RemoveCommandResponse{}

	command := &wompatti.Command{}

	s.db.First(&command, in.CommandId)

	if command.ID > 0 {
		res.State = WompattiService.RemoveCommandResponse_SUCCESS
		s.db.Delete(&command)
	} else {
		res.State = WompattiService.RemoveCommandResponse_NOT_FOUND
	}

	return res, nil
}

func (s *WompattiServiceServer) CreateKeyValue(ctx context.Context, in *WompattiService.CreateKeyValueRequest) (*WompattiService.CreateKeyValueResponse, error) {
	keyValue := &wompatti.KeyValue{
		DeviceInfoID: in.DeviceInfoId,
		Key:          in.Key,
		Value:        in.Value,
	}

	s.db.Create(keyValue)

	return &WompattiService.CreateKeyValueResponse{
		KeyValue: &WompattiService.KeyValue{
			Id:           keyValue.ID,
			DeviceInfoId: keyValue.DeviceInfoID,
			Key:          keyValue.Key,
			Value:        keyValue.Value,
		},
	}, nil
}

func (s *WompattiServiceServer) EditKeyValue(ctx context.Context, in *WompattiService.EditKeyValueRequest) (*WompattiService.EditKeyValueResponse, error) {
	var keyValue wompatti.KeyValue
	s.db.First(&keyValue, in.KeyValueId)

	res := &WompattiService.EditKeyValueResponse{}

	if keyValue.ID > 0 {
		res.State = WompattiService.EditKeyValueResponse_SUCCESS

		if in.Key != "" && in.Key != keyValue.Key {
			keyValue.Key = in.Key
		}
		if in.Value != "" && in.Value != keyValue.Value {
			keyValue.Value = in.Value
		}
		s.db.Save(keyValue)

		res.KeyValue = &WompattiService.KeyValue{
			Id:           keyValue.ID,
			DeviceInfoId: keyValue.DeviceInfoID,
			Key:          keyValue.Key,
			Value:        keyValue.Value,
		}
	} else {
		res.State = WompattiService.EditKeyValueResponse_NOT_FOUND
	}

	return res, nil
}

func (s *WompattiServiceServer) RemoveKeyValue(ctx context.Context, in *WompattiService.RemoveKeyValueRequest) (*WompattiService.RemoveKeyValueResponse, error) {
	response := &WompattiService.RemoveKeyValueResponse{}
	keyValue := &wompatti.KeyValue{}
	s.db.First(&keyValue, in.KeyValueId)

	if keyValue.ID != 0 {
		response.State = WompattiService.RemoveKeyValueResponse_SUCCESS
		s.db.Delete(&keyValue)
	} else {
		response.State = WompattiService.RemoveKeyValueResponse_NOT_FOUND
	}

	return response, nil
}

func (s *WompattiServiceServer) CreateWolInterface(ctx context.Context, in *WompattiService.CreateWolInterfaceRequest) (*WompattiService.CreateWolInterfaceResponse, error) {
	wolInterface := &wompatti.WolInterface{
		EthernetInterfaceId: in.EthernetInterfaceId,
		Mac:                 in.Mac,
	}

	s.db.Create(wolInterface)

	return &WompattiService.CreateWolInterfaceResponse{
		WolInterface: &WompattiService.WolInterface{
			Id:                  wolInterface.ID,
			EthernetInterfaceId: wolInterface.EthernetInterfaceId,
			Mac:                 wolInterface.Mac,
		},
	}, nil
}

func (s *WompattiServiceServer) EditWolInterface(ctx context.Context, in *WompattiService.EditWolInterfaceRequest) (*WompattiService.EditWolInterfaceResponse, error) {
	res := &WompattiService.EditWolInterfaceResponse{}

	wolInterface := &wompatti.WolInterface{}

	s.db.First(&wolInterface, in.WolInterfaceId)

	if wolInterface.ID > 0 {
		res.State = WompattiService.EditWolInterfaceResponse_SUCCESS

		if in.EthernetInterfaceId != 0 {
			wolInterface.EthernetInterfaceId = in.EthernetInterfaceId
		}

		if in.Mac != nil {
			wolInterface.Mac = in.Mac
		}

		s.db.Save(&wolInterface)

		res.WolInterface = &WompattiService.WolInterface{
			Id:                  wolInterface.ID,
			EthernetInterfaceId: wolInterface.EthernetInterfaceId,
			Mac:                 wolInterface.Mac,
		}
	} else {
		res.State = WompattiService.EditWolInterfaceResponse_NOT_FOUND
	}

	return nil, nil
}

func (s *WompattiServiceServer) RemoveWolInterface(ctx context.Context, in *WompattiService.RemoveWolInterfaceRequest) (*WompattiService.RemoveWolInterfaceResponse, error) {
	res := &WompattiService.RemoveWolInterfaceResponse{}
	wolInterface := &wompatti.WolInterface{}

	s.db.First(&wolInterface, in.WolInterfaceId)

	if wolInterface.ID > 0 {
		res.State = WompattiService.RemoveWolInterfaceResponse_SUCCESS
		s.db.Delete(&wolInterface)
	} else {
		res.State = WompattiService.RemoveWolInterfaceResponse_NOT_FOUND
	}

	return res, nil
}

func (s *WompattiServiceServer) ExecuteWolInterface(ctx context.Context, in *WompattiService.ExecuteWolInterfaceRequest) (*WompattiService.ExecuteWolInterfaceResponse, error) {
	return nil, nil
}

func (s *WompattiServiceServer) CreateTelnetInterface(ctx context.Context, in *WompattiService.CreateTelnetInterfaceRequest) (*WompattiService.CreateTelnetInterfaceResponse, error) {
	res := &WompattiService.CreateTelnetInterfaceResponse{}

	telnetInterface := &wompatti.TelnetInterface{
		Ip:   in.Ip,
		Port: in.Port,
	}

	s.db.Save(&telnetInterface)

	res.TelnetInterface = GenerateServiceTelnetInterface(telnetInterface)

	return res, nil
}

func (s *WompattiServiceServer) EditTelnetInterface(ctx context.Context, in *WompattiService.EditTelnetInterfaceRequest) (*WompattiService.EditTelnetInterfaceResponse, error) {
	res := &WompattiService.EditTelnetInterfaceResponse{}

	telnetInterface := &wompatti.TelnetInterface{}

	s.db.First(&telnetInterface, in.Id)

	if telnetInterface.ID > 0 {
		if in.Ip != "" {
			telnetInterface.Ip = in.Ip
		}
		if in.Port != "" {
			telnetInterface.Port = in.Port
		}

		res.State = WompattiService.EditTelnetInterfaceResponse_SUCCESS
		res.TelnetInterface = GenerateServiceTelnetInterface(telnetInterface)
	} else {
		res.State = WompattiService.EditTelnetInterfaceResponse_NOT_FOUND
	}

	return res, nil
}

func (s *WompattiServiceServer) RemoveTelnetInterface(ctx context.Context, in *WompattiService.RemoveTelnetInterfaceRequest) (*WompattiService.RemoveTelnetInterfaceResponse, error) {
	res := &WompattiService.RemoveTelnetInterfaceResponse{}

	telnetInterface := &wompatti.TelnetInterface{}

	s.db.First(&telnetInterface, in.TelnetInterfaceId)

	if telnetInterface.ID > 0 {
		res.State = WompattiService.RemoveTelnetInterfaceResponse_SUCCESS
		s.db.Delete(&telnetInterface)
	} else {
		res.State = WompattiService.RemoveTelnetInterfaceResponse_NOT_FOUND
	}

	return res, nil
}

func (s *WompattiServiceServer) CreateSerialInterface(ctx context.Context, in *WompattiService.CreateSerialInterfaceRequest) (*WompattiService.CreateSerialInterfaceResponse, error) {
	res := &WompattiService.CreateSerialInterfaceResponse{}

	serialInterface := &wompatti.SerialInterface{
		SerialPortID: in.SerialPortId,
	}

	s.db.Create(&serialInterface)

	res.SerialInterface = GenerateServiceSerialInterface(serialInterface)

	return res, nil
}

func (s *WompattiServiceServer) EditSerialInterface(ctx context.Context, in *WompattiService.EditSerialInterfaceRequest) (*WompattiService.EditSerialInterfaceResponse, error) {
	res := &WompattiService.EditSerialInterfaceResponse{}

	serialInterface := &wompatti.SerialInterface{}

	s.db.First(&serialInterface, in.SerialInterfaceId)

	if serialInterface.ID > 0 {
		if in.SerialPortId != 0 {
			serialInterface.SerialPortID = in.SerialInterfaceId
		}

		s.db.Save(&serialInterface)

		res.State = WompattiService.EditSerialInterfaceResponse_SUCCESS
		res.SerialInterface = GenerateServiceSerialInterface(serialInterface)
	} else {
		res.State = WompattiService.EditSerialInterfaceResponse_NOT_FOUND
	}

	return res, nil
}

func (s *WompattiServiceServer) RemoveSerialInterface(ctx context.Context, in *WompattiService.RemoveSerialInterfaceRequest) (*WompattiService.RemoveSerialInterfaceResponse, error) {
	res := &WompattiService.RemoveSerialInterfaceResponse{}

	serialInterface := &wompatti.SerialInterface{}

	s.db.First(&serialInterface, in.SerialInterfaceId)

	if serialInterface.ID > 0 {
		res.State = WompattiService.RemoveSerialInterfaceResponse_SUCCESS
		s.db.Delete(&serialInterface)
	} else {
		res.State = WompattiService.RemoveSerialInterfaceResponse_NOT_FOUND
	}

	return res, nil
}
