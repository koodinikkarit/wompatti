package wompatti

import (
	"net"

	"golang.org/x/net/context"

	wompatti "github.com/koodinikkarit/wompatti/db"
	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"
)

func (s *WompattiServiceServer) FetchComputers(ctx context.Context, in *WompattiService.FetchComputersRequest) (*WompattiService.ComputersConnection, error) {
	res := &WompattiService.ComputersConnection{}

	computers := []wompatti.Computer{}

	s.db.Find(&computers).Count(&res.TotalCount)

	q, startCursor := wompatti.ConnectionsQuery(s.db, in.After, in.Before, in.First, in.Last)

	q = q.Find(&computers)

	currentCursor := startCursor + 1

	for _, computer := range computers {
		cc := &WompattiService.Computer{
			Id:             computer.ID,
			ArttuId:        computer.ArttuID,
			DeviceInfoId:   computer.DeviceInfoID,
			WolInterfaceId: computer.WolInterfaceID,
			Name:           computer.Name,
		}
		res.Edges = append(res.Edges, &WompattiService.ComputersEdge{
			Cursor: currentCursor,
			Node:   cc,
		})
		currentCursor++
	}

	return res, nil
}

func (s *WompattiServiceServer) FetchComputerById(ctx context.Context, in *WompattiService.FetchComputerByIdRequest) (*WompattiService.FetchComputerByIdResponse, error) {
	res := &WompattiService.FetchComputerByIdResponse{}
	computers := []wompatti.Computer{}
	s.db.Where("id in (?)", in.ComputerIdt).Find(&computers)
	for _, computerId := range in.ComputerIdt {
		var found bool
		for _, computer := range computers {
			if computer.ID == computerId {
				res.Computers = append(res.Computers, &WompattiService.Computer{
					Id:             computer.ID,
					Name:           computer.Name,
					ArttuId:        computer.ArttuID,
					DeviceInfoId:   computer.DeviceInfoID,
					WolInterfaceId: computer.WolInterfaceID,
				})
				found = true
				break
			}
		}
		if found == false {
			res.Computers = append(res.Computers, &WompattiService.Computer{
				Id: 0,
			})
		}
	}
	return res, nil
}

func (s *WompattiServiceServer) FetchDevices(ctx context.Context, in *WompattiService.FetchDevicesRequest) (*WompattiService.DevicesConnection, error) {
	res := &WompattiService.DevicesConnection{}

	var devices = []wompatti.Device{}

	s.db.Find(&devices).Count(&res.TotalCount)

	q, startCursor := wompatti.ConnectionsQuery(s.db, in.After, in.Before, in.First, in.Last)

	q.Find(&devices)

	currentCursor := startCursor + 1

	for _, device := range devices {
		res.Edges = append(res.Edges, &WompattiService.DevicesEdge{
			Node: &WompattiService.Device{
				Id:           device.ID,
				Name:         device.Name,
				DeviceInfoId: device.DeviceInfoID,
				DeviceTypeId: device.DeviceTypeID,
			},
			Cursor: currentCursor,
		})
		currentCursor++
	}

	return res, nil
}

func (s *WompattiServiceServer) FetchDeviceById(ctx context.Context, in *WompattiService.FetchDeviceByIdRequest) (*WompattiService.FetchDeviceByIdResponse, error) {
	res := &WompattiService.FetchDeviceByIdResponse{}
	devices := []wompatti.Device{}
	s.db.Where("id in (?)", in.DeviceIdt).Find(&devices)
	for _, deviceId := range in.DeviceIdt {
		var found bool
		for _, device := range devices {
			if device.ID == deviceId {
				res.Devices = append(res.Devices, &WompattiService.Device{
					Id:           device.ID,
					Name:         device.Name,
					DeviceInfoId: uint32(device.DeviceInfoID),
					DeviceTypeId: device.DeviceTypeID,
				})
				found = true
				break
			}
		}
		if found == false {
			res.Devices = append(res.Devices, &WompattiService.Device{
				Id: 0,
			})
		}
	}

	return nil, nil
}

func (s *WompattiServiceServer) FetchDeviceInfoById(ctx context.Context, in *WompattiService.FetchDeviceInfoByIdRequest) (*WompattiService.FetchDeviceInfoByIdResponse, error) {
	res := &WompattiService.FetchDeviceInfoByIdResponse{}
	deviceInfos := []wompatti.DeviceInfo{}
	s.db.Where("id in (?)", in.DeviceInfoIdt).Find(&deviceInfos)
	for _, deviceInfoId := range in.DeviceInfoIdt {
		var found bool
		for _, deviceInfo := range deviceInfos {
			if deviceInfoId == uint32(deviceInfo.ID) {
				res.DeviceInfos = append(res.DeviceInfos, &WompattiService.DeviceInfo{
					Id: uint32(deviceInfo.ID),
				})
				found = true
				break
			}
		}
		if found == false {
			res.DeviceInfos = append(res.DeviceInfos, &WompattiService.DeviceInfo{
				Id: 0,
			})
		}
	}

	return res, nil
}

func (s *WompattiServiceServer) FetchKeyValuesByDeviceInfoId(ctx context.Context, in *WompattiService.FetchKeyValuesByDeviceInfoIdRequest) (*WompattiService.FetchKeyValuesByDeviceInfoIdResponse, error) {
	res := &WompattiService.FetchKeyValuesByDeviceInfoIdResponse{}

	keyValues := []wompatti.KeyValue{}
	s.db.Where("device_info_id IN (?)", in.DeviceInfoIdt).Find(&keyValues)
	for _, deviceInfoId := range in.DeviceInfoIdt {
		deviceInfoKeyValues := &WompattiService.DeviceInfoKeyValues{
			DeviceInfoId: deviceInfoId,
		}

		for _, keyValue := range keyValues {
			if deviceInfoId == keyValue.DeviceInfoID {
				deviceInfoKeyValues.KeyValues = append(deviceInfoKeyValues.KeyValues, &WompattiService.KeyValue{
					Id:           keyValue.ID,
					DeviceInfoId: keyValue.DeviceInfoID,
					Key:          keyValue.Key,
					Value:        keyValue.Value,
				})
			}
		}
	}

	return res, nil
}

func (s *WompattiServiceServer) FetchEthernetInterfaces(ctx context.Context, in *WompattiService.FetchEthernetInterfacesRequest) (*WompattiService.EthernetInterfacesConnection, error) {
	res := &WompattiService.EthernetInterfacesConnection{}

	ifaces, _ := net.Interfaces()

	// s.db.Find(&devices).Count(&res.TotalCount)

	// q, startCursor := wompatti.ConnectionsQuery(s.db, in.After, in.Before, in.First, in.Last)

	// q.Find(&devices)

	//currentCursor := startCursor + 1

	for _, iface := range ifaces {
		res.Edges = append(res.Edges, &WompattiService.EthernetInterfaceEdge{
			Node: &WompattiService.EthernetInterface{
				Id:    uint32(iface.Index),
				Name:  iface.Name,
				Mac:   iface.HardwareAddr,
				Index: uint32(iface.Index),
				Mtu:   uint32(iface.MTU),
				Flags: uint32(iface.Flags),
			},
			//Cursor: currentCursor,
		})
		//currentCursor++
	}
	return nil, nil
}

func (s *WompattiServiceServer) FetchDeviceTypes(ctx context.Context, in *WompattiService.FetchDeviceTypesRequest) (*WompattiService.DeviceTypesConnection, error) {
	res := &WompattiService.DeviceTypesConnection{}

	deviceTypes := []wompatti.DeviceType{}

	s.db.Find(&deviceTypes).Count(&res.TotalCount)

	q, startCursor := wompatti.ConnectionsQuery(s.db, in.After, in.Before, in.First, in.Last)

	q.Find(&deviceTypes)

	currentCursor := startCursor + 1

	for _, deviceType := range deviceTypes {
		res.Edges = append(res.Edges, &WompattiService.DeviceTypesEdge{
			Node: &WompattiService.DeviceType{
				Id:   deviceType.ID,
				Name: deviceType.Name,
			},
			Cursor: currentCursor,
		})
		currentCursor++
	}

	return res, nil
}

func (s *WompattiServiceServer) FetchDeviceTypeById(ctx context.Context, in *WompattiService.FetchDeviceTypeByIdRequest) (*WompattiService.FetchDeviceTypeByIdResponse, error) {
	res := &WompattiService.FetchDeviceTypeByIdResponse{}

	deviceTypes := []wompatti.DeviceType{}

	s.db.Where("id in (?)", in.DeviceTypesIdt).Find(&deviceTypes)

	for _, deviceTypeId := range in.DeviceTypesIdt {
		var found bool
		for _, deviceType := range deviceTypes {
			if deviceType.ID == deviceTypeId {
				res.DeviceTypes = append(res.DeviceTypes, &WompattiService.DeviceType{
					Id:   deviceType.ID,
					Name: deviceType.Name,
				})
				found = true
				break
			}
		}
		if found != false {
			res.DeviceTypes = append(res.DeviceTypes, &WompattiService.DeviceType{
				Id: 0,
			})
		}
	}

	return res, nil
}

func (s *WompattiServiceServer) FetchCommandsByDeviceTypeId(ctx context.Context, in *WompattiService.FetchCommandsByDeviceTypeIdRequest) (*WompattiService.FetchCommandsByDeviceTypeIdResponse, error) {
	res := &WompattiService.FetchCommandsByDeviceTypeIdResponse{}

	commands := []wompatti.Command{}

	s.db.Where("device_type_id in (?)", in.DeviceTypeIdt).Find(&commands)

	for _, deviceTypeId := range in.DeviceTypeIdt {
		deviceTypeCommands := &WompattiService.DeviceTypeCommands{
			DeviceTypeId: deviceTypeId,
		}

		for _, command := range commands {
			if command.DeviceTypeID == deviceTypeId {
				deviceTypeCommands.Commands = append(deviceTypeCommands.Commands, &WompattiService.Command{
					Id:           command.ID,
					DeviceTypeId: command.DeviceTypeID,
					Name:         command.Name,
					Value:        command.Value,
				})
			}
		}
		res.DeviceTypeCommands = append(res.DeviceTypeCommands, deviceTypeCommands)
	}

	return nil, nil
}

func (s *WompattiServiceServer) FetchWolInterfaceById(ctx context.Context, in *WompattiService.FetchWolInterfaceByIdRequest) (*WompattiService.FetchWolInterfaceByIdResponse, error) {
	res := &WompattiService.FetchWolInterfaceByIdResponse{}

	wolInterfaces := []wompatti.WolInterface{}

	s.db.Where("id in (?)", in.WolInterfaceIdt).Find(&wolInterfaces)

	for _, wolInterfaceId := range in.WolInterfaceIdt {
		var found bool
		for _, wolInterface := range wolInterfaces {
			if wolInterface.ID == wolInterfaceId {
				res.WolInterfaces = append(res.WolInterfaces, &WompattiService.WolInterface{
					Id:                  wolInterface.ID,
					EthernetInterfaceId: wolInterface.EthernetInterfaceId,
					Mac:                 wolInterface.Mac,
				})
				found = true
				break
			}
		}
		if found != false {
			res.WolInterfaces = append(res.WolInterfaces, &WompattiService.WolInterface{
				Id: 0,
			})
		}
	}

	return res, nil
}

func (s *WompattiServiceServer) FetchTelnetInterfaces(ctx context.Context, in *WompattiService.FetchTelnetInterfacesRequest) (*WompattiService.TelnetInterfacesConnection, error) {
	res := &WompattiService.TelnetInterfacesConnection{}

	telnetInterfaces := []wompatti.TelnetInterface{}

	s.db.Find(&telnetInterfaces).Count(&res.TotalCount)

	q, startCursor := wompatti.ConnectionsQuery(s.db, in.After, in.Before, in.First, in.Last)

	q.Find(&telnetInterfaces)

	currentCursor := startCursor + 1

	for _, telnetInterface := range telnetInterfaces {
		res.Edges = append(res.Edges, &WompattiService.TelnetInterfacesEdge{
			Node: &WompattiService.TelnetInterface{
				Id:   telnetInterface.ID,
				Ip:   telnetInterface.Ip,
				Port: telnetInterface.Port,
			},
			Cursor: currentCursor,
		})
		currentCursor++
	}

	return res, nil
}

func (s *WompattiServiceServer) FetchTelnetInterfaceById(ctx context.Context, in *WompattiService.FetchTelnetInterfaceByIdRequest) (*WompattiService.FetchTelnetInterfaceByIdResponse, error) {
	res := &WompattiService.FetchTelnetInterfaceByIdResponse{}

	telnetInterfaces := []wompatti.TelnetInterface{}

	s.db.Where("id in (?)", in.TelnetInterfaceId).Find(&telnetInterfaces)

	for _, telnetInterfaceId := range in.TelnetInterfaceId {
		var found bool
		for _, telnetInterface := range telnetInterfaces {
			if telnetInterfaceId == telnetInterface.ID {
				res.TelnetInterfaces = append(res.TelnetInterfaces, &WompattiService.TelnetInterface{
					Id:   telnetInterface.ID,
					Ip:   telnetInterface.Ip,
					Port: telnetInterface.Port,
				})
				found = true
			}
		}
		if found == false {
			res.TelnetInterfaces = append(res.TelnetInterfaces, &WompattiService.TelnetInterface{
				Id: 0,
			})
		}
	}

	return res, nil
}

func (s *WompattiServiceServer) FetchSerialInterfaces(ctx context.Context, in *WompattiService.FetchSerialInterfacesRequest) (*WompattiService.SerialInterfacesConnection, error) {
	res := &WompattiService.SerialInterfacesConnection{}

	serialInterfaces := []wompatti.SerialInterface{}

	s.db.Find(&serialInterfaces).Count(&res.TotalCount)

	q, startCursor := wompatti.ConnectionsQuery(s.db, in.After, in.Before, in.First, in.Last)

	q.Find(&serialInterfaces)

	currentCursor := startCursor + 1

	for _, serialInterface := range serialInterfaces {
		res.Edges = append(res.Edges, &WompattiService.SerialInterfacesEdge{
			Node: &WompattiService.SerialInterface{
				Id:           serialInterface.ID,
				SerialPortId: serialInterface.SerialPortID,
			},
			Cursor: currentCursor,
		})
		currentCursor++
	}

	return res, nil
}

func (s *WompattiServiceServer) FetchSerialInterfaceById(ctx context.Context, in *WompattiService.FetchSerialInterfaceByIdRequest) (*WompattiService.FetchSerialInterfaceByIdResponse, error) {
	res := &WompattiService.FetchSerialInterfaceByIdResponse{}

	serialInterfaces := []wompatti.SerialInterface{}

	s.db.Where("id in (?)", in.SerialInterfaceIdt).Find(&serialInterfaces)

	for _, serialInterfaceId := range in.SerialInterfaceIdt {
		var found bool
		for _, serialInterface := range serialInterfaces {
			if serialInterface.ID == serialInterfaceId {
				res.SerialInterfaces = append(res.SerialInterfaces, &WompattiService.SerialInterface{
					Id:           serialInterface.ID,
					SerialPortId: serialInterfaceId,
				})
				found = true
			}
		}
		if found == false {
			res.SerialInterfaces = append(res.SerialInterfaces, &WompattiService.SerialInterface{
				Id: 0,
			})
		}
	}

	return res, nil
}

// func (s *WompattiServiceServer) Ping(in *WompattiService.PingRequest, stream WompattiService.Wompatti_PingServer) error {
// 	c <- s.pinger.Ping(in.IpAddress)

// 	for pingResponse <- c {
// 		if err := stream.Send(pingResponse); err != nil {
// 			return err;
// 		}
// 	}
// 	return nil;
// }

func (s *WompattiServiceServer) Ping(
	ctx context.Context, 
	in *WompattiService.PingRequest,
) (
	*WompattiService.PingResponse,
	error
) {
	return c <- s.pinger.Ping(in.IpAddress)
}
