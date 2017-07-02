package wompatti

import (
	wompatti "github.com/koodinikkarit/wompatti/db"
	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"
)

func GenerateServiceDeviceType(model *wompatti.DeviceType) *WompattiService.DeviceType {
	return &WompattiService.DeviceType{
		Id:   model.ID,
		Name: model.Name,
	}
}

func GenerateServiceCommand(model *wompatti.Command) *WompattiService.Command {
	return &WompattiService.Command{
		Id:           model.ID,
		DeviceTypeId: model.DeviceTypeID,
		Name:         model.Name,
		Value:        model.Value,
	}
}

func GenerateServiceTelnetInterface(model *wompatti.TelnetInterface) *WompattiService.TelnetInterface {
	return &WompattiService.TelnetInterface{
		Id:   model.ID,
		Ip:   model.Ip,
		Port: model.Port,
	}
}

func GenerateServiceSerialInterface(model *wompatti.SerialInterface) *WompattiService.SerialInterface {
	return &WompattiService.SerialInterface{
		Id:           model.ID,
		SerialPortId: model.SerialPortID,
	}
}
