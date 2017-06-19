package wompatti

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	gorm "github.com/jinzhu/gorm"
	wompatti "github.com/koodinikkarit/wompatti/db"
	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"
)

type WompattiServiceServer struct {
	db *gorm.DB
}

// func (s *WompattiServiceServer) AddComputer(ctx context.Context, in *WompattiService.AddComputerRequest) (*WompattiService.AddComputerResponse, error) {
// 	computer := wompatti.Computer{Name: in.Name, Mac: in.Mac}
// 	s.db.Create(&computer)
// 	return &WompattiService.AddComputerResponse{
// 		Computer: &WompattiService.Computer{
// 			Id:   computer.ID,
// 			Name: computer.Name,
// 			Mac:  computer.Mac,
// 		},
// 	}, nil
// }

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

func (s *WompattiServiceServer) FetchComputers(fetchComputers *WompattiService.FetchComputersRequest, stream WompattiService.Wompatti_FetchComputersServer) error {
	computers := []wompatti.Computer{}
	s.db.Find(&computers)

	// md, ok := metadata.FromContext(stream.Context())
	// if ok {
	// 	fmt.Println(md.Len())
	// }

	for _, computer := range computers {
		stream.Send(&WompattiService.Computer{
			Id:           computer.ID,
			Name:         computer.Name,
			ArttuId:      computer.ArttuID,
			DeviceInfoId: computer.DeviceInfoID,
		})
	}

	return nil
}

func (s *WompattiServiceServer) FetchComputerById(in *WompattiService.FetchComputerByIdRequest, stream WompattiService.Wompatti_FetchComputerByIdServer) error {
	computers := []wompatti.Computer{}
	s.db.Where("id in (?)", in.ComputerIdt).Find(&computers)
	for _, computerId := range in.ComputerIdt {
		var found bool
		for _, computer := range computers {
			if computer.ID == computerId {
				stream.Send(&WompattiService.FetchComputerByIdResponse{
					State: WompattiService.FetchComputerByIdResponse_SUCCESS,
					Computer: &WompattiService.Computer{
						Id:           computer.ID,
						Name:         computer.Name,
						ArttuId:      computer.ArttuID,
						DeviceInfoId: computer.DeviceInfoID,
					},
				})
				found = true
				break
			}
		}
		if found == false {
			stream.Send(&WompattiService.FetchComputerByIdResponse{
				State:      WompattiService.FetchComputerByIdResponse_NOT_FOUND,
				ComputerId: computerId,
			})
		}
	}
	return nil
}

func (s *WompattiServiceServer) FetchDeviceInfoById(in *WompattiService.FetchDeviceInfoByIdRequest, stream WompattiService.Wompatti_FetchDeviceInfoByIdServer) error {
	deviceInfos := []wompatti.DeviceInfo{}
	s.db.Where("id in (?)", in.DeviceInfoIdt).Find(&deviceInfos)
	for _, deviceInfoId := range in.DeviceInfoIdt {
		var found bool
		for _, deviceInfo := range deviceInfos {
			if deviceInfoId == uint32(deviceInfo.ID) {
				stream.Send(&WompattiService.FetchDeviceInfoByIdResponse{
					State: WompattiService.FetchDeviceInfoByIdResponse_SUCCESS,
					DeviceInfo: &WompattiService.DeviceInfo{
						Id: uint32(deviceInfo.ID),
					},
				})
				found = true
				break
			}
		}
		if found == false {
			stream.Send(&WompattiService.FetchDeviceInfoByIdResponse{
				State:        WompattiService.FetchDeviceInfoByIdResponse_NOT_FOUND,
				DeviceInfoId: uint32(deviceInfoId),
			})
		}
	}

	return nil
}

func (s *WompattiServiceServer) FetchKeyValuesByDeviceInfoId(in *WompattiService.FetchKeyValuesByDeviceInfoIdRequest, stream WompattiService.Wompatti_FetchKeyValuesByDeviceInfoIdServer) error {
	keyValues := []wompatti.KeyValue{}
	s.db.Where("device_info_id IN (?)", in.DeviceInfoIdt).Find(&keyValues)
	for _, deviceInfoId := range in.DeviceInfoIdt {
		var fetchKeyValuesByDeviceInfoIdResponse WompattiService.FetchKeyValuesByDeviceInfoIdResponse
		fetchKeyValuesByDeviceInfoIdResponse.DeviceInfoId = deviceInfoId
		for _, keyValue := range keyValues {
			if deviceInfoId == keyValue.DeviceInfoID {
				fetchKeyValuesByDeviceInfoIdResponse.KeyValues = append(fetchKeyValuesByDeviceInfoIdResponse.KeyValues, &WompattiService.KeyValue{
					Id:           keyValue.ID,
					DeviceInfoId: keyValue.DeviceInfoID,
					Key:          keyValue.Key,
					Value:        keyValue.Value,
				})
			}
		}
		stream.Send(&fetchKeyValuesByDeviceInfoIdResponse)
	}

	return nil
}

// func (s *WompattiServiceServer) EditComputer(ctx context.Context, in *WompattiService.EditComputerRequest) (*WompattiService.EditComputerResponse, error) {
// 	var computer wompatti.Computer
// 	s.db.First(&computer, in.Id)
// 	if in.Name != "" {
// 		computer.Name = in.Name
// 	}
// 	if in.Mac != "" {
// 		computer.Mac = in.Mac
// 	}
// 	s.db.Save(&computer)
// 	return &WompattiService.EditComputerResponse{
// 		Computer: &WompattiService.Computer{
// 			Id:   computer.ID,
// 			Name: computer.Name,
// 			Mac:  computer.Mac,
// 		},
// 	}, nil
// }

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

	response := &WompattiService.EditKeyValueResponse{}

	if in.Key != "" && in.Key != keyValue.Key {
		keyValue.Key = in.Key
	}
	if in.Value != "" && in.Value != keyValue.Value {
		keyValue.Value = in.Value
	}
	s.db.Save(keyValue)
	response.State = WompattiService.EditKeyValueResponse_SUCCESS
	response.KeyValue = &WompattiService.KeyValue{
		Id:           keyValue.ID,
		DeviceInfoId: keyValue.DeviceInfoID,
		Key:          keyValue.Key,
		Value:        keyValue.Value,
	}

	return response, nil
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

func (s *WompattiServiceServer) FetchDevices(in *WompattiService.FetchDevicesRequest, stream WompattiService.Device) error {
	return nil
}

func (s *WompattiServiceServer) FetchDeviceById(in *WompattiService.FetchDeviceByIdRequest, stream WompattiService.Wompatti_FetchDeviceByIdServer) error {
	return nil
}

func (s *WompattiServiceServer) FetchEthernetInterfaces(in *WompattiService.FetchEthernetInterfacesRequest, stream WompattiService.Wompatti_FetchEthernetInterfacesServer) error {
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		stream.Send(&WompattiService.EthernetInterface{
			Name:  iface.Name,
			Mac:   iface.HardwareAddr,
			Index: uint32(iface.Index),
			Mtu:   uint32(iface.MTU),
			Flags: uint32(iface.Flags),
		})
	}
	return nil
}

func (s *WompattiServiceServer) CreateWolInterface(ctx context.Context, in *WompattiService.CreateWolInterfaceRequest) (*WompattiService.CreateWolInterfaceResponse, error) {
	return &WompattiService.CreateWolInterfaceResponse{}, nil
}

func CreateWompattiService(db *gorm.DB, port string) *grpc.Server {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("./ssl/server.crt", "./ssl/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	WompattiService.RegisterWompattiServer(s, &WompattiServiceServer{db: db})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return s
}

// func (s *WompattiServiceServer) AddComputer(ctx context.Context, in *WompattiService.AddComputerRequest) (*WompattiService.AddComputerResponse, error) {
// 	computer := wompatti.Computer{Name: in.Name, Mac: in.Mac}
// 	s.db.Create(&computer)
// 	return &WompattiService.AddComputerResponse{
// 		Computer: &WompattiService.Computer{
// 			Id:   computer.ID,
// 			Name: computer.Name,
// 			Mac:  computer.Mac,
// 		},
// 	}, nil
// }

// func (s *WompattiServiceServer) EditComputer(ctx context.Context, in *WompattiService.EditComputerRequest) (*WompattiService.EditComputerResponse, error) {
// 	var computer wompatti.Computer
// 	s.db.First(&computer, in.Id)
// 	if in.Name != "" {
// 		computer.Name = in.Name
// 	}
// 	if in.Mac != "" {
// 		computer.Mac = in.Mac
// 	}
// 	s.db.Save(&computer)
// 	return &WompattiService.EditComputerResponse{
// 		Computer: &WompattiService.Computer{
// 			Id:   computer.ID,
// 			Name: computer.Name,
// 			Mac:  computer.Mac,
// 		},
// 	}, nil
// }

// func (s *WompattiServiceServer) FetchComputerById(ctx context.Context, in *WompattiService.FetchComputerByIdRequest) (*WompattiService.FetchComputerByIdResponse, error) {
// 	var computer wompatti.Computer
// 	s.db.First(&computer, in.Id)
// 	return &WompattiService.FetchComputerByIdResponse{
// 		Computer: &WompattiService.Computer{
// 			Id:   computer.ID,
// 			Name: computer.Name,
// 			Mac:  computer.Mac,
// 		},
// 	}, nil
// }

// func (s *WompattiServiceServer) Wakeup(ctx context.Context, in *WompattiService.WakeupRequest) (*WompattiService.WakeupResponse, error) {
// 	var computer wompatti.Computer
// 	s.db.First(&computer, in.ComputerId)
// 	fmt.Println("Wakeup", computer)
// 	return &WompattiService.WakeupResponse{}, nil
// }

// func (s *WompattiServiceServer) RemoveComputer(ctx context.Context, in *WompattiService.RemoveComputerRequest) (*WompattiService.RemoveComputerResponse, error) {
// 	computer := &wompatti.Computer{ID: in.ComputerId}
// 	s.db.Delete(computer)
// 	return &WompattiService.RemoveComputerResponse{}, nil
// }

// func (s *WompattiServiceServer) AddKeijo(ctx context.Context, in *WompattiService.AddKeijoRequest) (*WompattiService.AddKeijoResponse, error) {
// 	keijo := Keijo{Name: in.Name, Ip: in.Ip, Port: in.Port}
// 	s.db.Create(&keijo)
// 	return &WompattiService.AddKeijoResponse{
// 		Keijo: &WompattiService.Keijo{
// 			Id:   keijo.ID,
// 			Name: keijo.Name,
// 			Ip:   keijo.Ip,
// 			Port: keijo.Port,
// 		},
// 	}, nil
// }

// func (s *WompattiServiceServer) EditKeijo(ctx context.Context, in *WompattiService.EditKeijoRequest) (*WompattiService.EditKeijoResponse, error) {
// 	var keijo Keijo
// 	s.db.Find(&keijo, in.KeijoId)
// 	if in.Name != "" {
// 		keijo.Name = in.Name
// 	}
// 	if in.Ip != "" {
// 		keijo.Ip = in.Ip
// 	}
// 	if in.Port != "" {
// 		keijo.Port = in.Port
// 	}
// 	s.db.Save(&keijo)
// 	return &WompattiService.EditKeijoResponse{
// 		Keijo: &WompattiService.Keijo{
// 			Id:   keijo.ID,
// 			Name: keijo.Name,
// 			Ip:   keijo.Ip,
// 			Port: keijo.Port,
// 		},
// 	}, nil
// }

// func (s *WompattiServiceServer) RemoveKeijo(ctx context.Context, in *WompattiService.RemoveKeijoRequest) (*WompattiService.RemoveKeijoResponse, error) {
// 	keijo := &Keijo{ID: in.KeijoId}
// 	s.db.Delete(keijo)
// 	return &WompattiService.RemoveKeijoResponse{}, nil
// }

// func (s *WompattiServiceServer) FetchKeijos(in *WompattiService.FetchKeijosRequest, stream WompattiService.Wompatti_FetchKeijosServer) error {
// 	keijos := []Keijo{}
// 	s.db.Find(&keijos)
// 	for _, keijo := range keijos {
// 		stream.Send(&WompattiService.Keijo{
// 			Id:   keijo.ID,
// 			Name: keijo.Name,
// 			Ip:   keijo.Ip,
// 			Port: keijo.Port,
// 		})
// 	}

// 	return nil
// }

// func (s *WompattiServiceServer) FetchKeijoById(ctx context.Context, in *WompattiService.FetchKeijoByIdRequest) (*WompattiService.FetchKeijoByIdResponse, error) {
// 	var keijo Keijo
// 	s.db.First(&keijo, in.KeijoId)
// 	return &WompattiService.FetchKeijoByIdResponse{
// 		Keijo: &WompattiService.Keijo{
// 			Id:   keijo.ID,
// 			Name: keijo.Name,
// 			Ip:   keijo.Ip,
// 			Port: keijo.Port,
// 		},
// 	}, nil
// }

// func (s *WompattiServiceServer) TurnOnCecDevice(ctx context.Context, in *WompattiService.TurnOnCecDeviceRequest) (*WompattiService.TurnOnCecDeviceResponse, error) {
// 	var keijo Keijo
// 	s.db.First(&keijo, in.KeijoId)

// 	client := CreateKeijoClient(keijo.Ip, keijo.Port)
// 	client.TurnOn(context.Background(), &KeijoService.TurnOnRequest{Address: in.Address})

// 	return &WompattiService.TurnOnCecDeviceResponse{}, nil
// }

// func (s *WompattiServiceServer) ChangeKeijoSource(ctx context.Context, in *WompattiService.ChangeKeijoSourceRequest) (*WompattiService.ChangeKeijoSourceResponse, error) {
// 	var keijo Keijo
// 	s.db.First(&keijo, in.KeijoId)

// 	client := CreateKeijoClient(keijo.Ip, keijo.Port)
// 	client.ChangeSource(context.Background(), &KeijoService.ChangeSourceRequest{Source: KeijoService.Devices_RECORDING1, Destination: KeijoService.Devices_BROADCAST, SourceNumber: in.SourceNumber})

// 	return &WompattiService.ChangeKeijoSourceResponse{}, nil
// }

// func (s *WompattiServiceServer) ShutDownCecDevice(ctx context.Context, in *WompattiService.ShutDownCecDeviceRequest) (*WompattiService.ShutDownCecDeviceResponse, error) {
// 	return &WompattiService.ShutDownCecDeviceResponse{}, nil
// }

// func (s *WompattiServiceServer) FetchCecDevicePowerStatusByKeijoId(ctx context.Context, in *WompattiService.FetchCecDevicePowerStatusByKeijoIdRequest) (*WompattiService.FetchCecDevicePowerStatusByKeijoIdResponse, error) {
// 	return &WompattiService.FetchCecDevicePowerStatusByKeijoIdResponse{}, nil
// }

// func (s *WompattiServiceServer) FetchCecTvDeviceSourceByKeijoId(ctx context.Context, in *WompattiService.FetchCecTvDeviceSourceByKeijoIdRequest) (*WompattiService.FetchCecTvDeviceSourceByKeijoIdResponse, error) {
// 	return &WompattiService.FetchCecTvDeviceSourceByKeijoIdResponse{}, nil
// }
