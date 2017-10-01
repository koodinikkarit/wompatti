package WompattiServiceServer

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	"github.com/koodinikkarit/wompatti/context"
	"github.com/koodinikkarit/wompatti/pinger"

	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"
)

func NewWompattiService(
	newContext func() *WompattiContext.Context,
	port string,
) *grpc.Server {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// creds, err := credentials.NewServerTLSFromFile(
	// 	"./ssl/server.crt",
	// 	"./ssl/server.key",
	// )
	// if err != nil {
	// 	log.Fatalf("Failed to generate credentials %v", err)
	// }

	s := grpc.NewServer()
	WompattiService.RegisterWompattiServer(
		s,
		&Server{
			newContext: newContext,
			pinger:     WompattiPinger.NewPinger(),
		},
	)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return s
}
