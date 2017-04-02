package main

import (
	"fmt"
	"io"

	"github.com/golang/protobuf/proto"
	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"
	"golang.org/x/net/context"
)

type server struct{}

func (s *server) addComputer(ctx context.Context, in WompattiService.AddComputer) WompattiService.ComputerAdded {

}

func (s *server) fetchComputers(fetchComputers WompattiService.FetchComputers, stream WompattiService.FetchComputers) error {
	stream.

}

func (s *server) fetchComputerById(ctx context.Context, in WompattiService.FetchComputerById) {

}

func (s *server) wakeup(ctx context.Context, in WompattiService.Wakeup) {

}

func main() {
	fmt.Println("Hello, 世界")
	var x float64
	x = 20.5
	fmt.Println("hello", x)
}
