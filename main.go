package main

import (
	"context"
	"fmt"

	"github.com/koodinikkarit/wompatti/db"
	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"
)

type server struct{}

func (s *server) AddComputer(ctx context.Context, in *WompattiService.AddComputerRequest) (*WompattiService.AddComputerResponse, error) {
	fmt.Println("Addcomputer", in.Name)
	return &WompattiService.AddComputerResponse{}, nil
}

func (s *server) EditComputer(ctx context.Context, in *WompattiService.EditComputerRequest) (*WompattiService.EditComputerResponse, error) {

	return &WompattiService.EditComputerResponse{}, nil
}

func (s *server) FetchComputers(fetchComputers *WompattiService.FetchComputersRequest, stream WompattiService.Wompatti_FetchComputersServer) error {
	fmt.Println("fetchComputers")
	stream.Send(&WompattiService.Computer{Name: "tykki"})
	return nil
}

func (s *server) FetchComputerById(ctx context.Context, in *WompattiService.FetchComputerByIdRequest) (*WompattiService.FetchComputerByIdResponse, error) {
	return &WompattiService.FetchComputerByIdResponse{}, nil
}

func (s *server) Wakeup(ctx context.Context, in *WompattiService.WakeupRequest) (*WompattiService.WakeupResponse, error) {
	fmt.Println("Wakeup", in.ComputerId)
	return &WompattiService.WakeupResponse{}, nil
}

type computer struct {
	id   int
	name string
	mac  string
}

func main() {
	fmt.Println("Hello, 世界")
	var x float64
	x = 20.5
	fmt.Println("hello", x)

	context := db.New()

	context.Add()

	// l := list.New()

	// l.PushBack(&Book{ad: 5})

	// for i := l.Front(); i != nil; i = i.Next() {
	// 	fmt.Println(i.Value.(*Book).ad)
	// }

	/*
		lis, err := net.Listen("tcp", ":7856")
		if err != nil {
			grpclog.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		WompattiService.RegisterWompattiServer(s, &server{})

		// Register reflection service on gRPC server.
		reflection.Register(s)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}*/
}
