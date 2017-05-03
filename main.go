package main

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/koodinikkarit/wompatti/wompatti"
	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"

	_ "github.com/jinzhu/gorm/dialects/mssql"
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

func main() {
	context := wompatti.CreateContext(
		"jaska",
		"asdf321",
		"localhost",
		"3306",
		"wompatti",
		"7856",
	)

	context.Start()
}
