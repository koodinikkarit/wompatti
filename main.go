package main

import (
	"fmt"

	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"
	"golang.org/x/net/context"
)

type server struct{}

func (s *server) addComputer(ctx context.Context, in WompattiService.AddComputerRequest) (*WompattiService.AddComputerResponse, error) {
	return &WompattiService.AddComputerResponse{}, nil
}

func (s *server) fetchComputers(fetchComputers WompattiService.FetchComputersRequest, stream WompattiService.Wompatti_FetchComputersServer) error {
	return nil
}

func (s *server) fetchComputerByID(ctx context.Context, in WompattiService.FetchComputerByIdRequest) (*WompattiService.FetchComputerByIdResponse, error) {
	return &WompattiService.FetchComputerByIdResponse{}, nil
}

func (s *server) wakeup(ctx context.Context, in WompattiService.WakeupRequest) (*WompattiService.WakeupResponse, error) {
	return &WompattiService.WakeupResponse{}, nil
}

func main() {
	fmt.Println("Hello, 世界")
	var x float64
	x = 20.5
	fmt.Println("hello", x)
}
