package wompatti

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	gorm "github.com/jinzhu/gorm"
	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"
)

type WompattiServiceServer struct {
	db *gorm.DB
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
