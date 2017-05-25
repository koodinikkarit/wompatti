package wompatti

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	gorm "github.com/jinzhu/gorm"
	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"
)

type server struct {
	db *gorm.DB
}

func (s *server) AddComputer(ctx context.Context, in *WompattiService.AddComputerRequest) (*WompattiService.AddComputerResponse, error) {
	computer := Computer{Name: in.Name, Mac: in.Mac}
	s.db.Create(&computer)
	return &WompattiService.AddComputerResponse{
		Computer: &WompattiService.Computer{
			Id:   computer.ID,
			Name: computer.Name,
			Mac:  computer.Mac,
		},
	}, nil
}

func (s *server) EditComputer(ctx context.Context, in *WompattiService.EditComputerRequest) (*WompattiService.EditComputerResponse, error) {
	var computer Computer
	s.db.First(&computer, in.Id)
	if in.Name != "" {
		computer.Name = in.Name
	}
	if in.Mac != "" {
		computer.Mac = in.Mac
	}
	s.db.Save(&computer)
	return &WompattiService.EditComputerResponse{
		Computer: &WompattiService.Computer{
			Id:   computer.ID,
			Name: computer.Name,
			Mac:  computer.Mac,
		},
	}, nil
}

func (s *server) FetchComputers(fetchComputers *WompattiService.FetchComputersRequest, stream WompattiService.Wompatti_FetchComputersServer) error {
	computers := []Computer{}
	s.db.Find(&computers)

	// md, ok := metadata.FromContext(stream.Context())
	// if ok {
	// 	fmt.Println(md.Len())
	// }

	for _, computer := range computers {
		stream.Send(&WompattiService.Computer{
			Id:   computer.ID,
			Name: computer.Name,
			Mac:  computer.Mac,
		})
	}

	return nil
}

func (s *server) FetchComputerById(ctx context.Context, in *WompattiService.FetchComputerByIdRequest) (*WompattiService.FetchComputerByIdResponse, error) {
	var computer Computer
	s.db.First(&computer, in.Id)
	return &WompattiService.FetchComputerByIdResponse{
		Computer: &WompattiService.Computer{
			Id:   computer.ID,
			Name: computer.Name,
			Mac:  computer.Mac,
		},
	}, nil
}

func (s *server) Wakeup(ctx context.Context, in *WompattiService.WakeupRequest) (*WompattiService.WakeupResponse, error) {
	var computer Computer
	s.db.First(&computer, in.ComputerId)
	fmt.Println("Wakeup", computer)
	return &WompattiService.WakeupResponse{}, nil
}

func (s *server) RemoveComputer(ctx context.Context, in *WompattiService.RemoveComputerRequest) (*WompattiService.RemoveComputerResponse, error) {
	computer := &Computer{ID: in.ComputerId}
	s.db.Delete(computer)
	return &WompattiService.RemoveComputerResponse{}, nil
}

func CreateService(db *gorm.DB, port string) *grpc.Server {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("./ssl/server.crt", "./ssl/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	WompattiService.RegisterWompattiServer(s, &server{db: db})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return s
}
