package wompatti

// import (
// 	"fmt"
// 	"log"
// 	"net"
// 	"time"

// 	ArttuService "github.com/koodinikkarit/wompatti/arttu_service"
// 	"github.com/koodinikkarit/wompatti/db"
// 	"golang.org/x/net/context"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/grpclog"
// 	"google.golang.org/grpc/reflection"

// 	"io"

// 	gorm "github.com/jinzhu/gorm"
// )

// type ArttuServiceServer struct {
// 	db *gorm.DB
// }

// func (s *ArttuServiceServer) InsertComputerBasicInformation(ctx context.Context, in *ArttuService.InsertComputerBasicInformationRequest) (*ArttuService.InsertComputerBasicInformationResponse, error) {

// 	return &ArttuService.InsertComputerBasicInformationResponse{}, nil
// }

// func (s *ArttuServiceServer) Register(ctx context.Context, in *ArttuService.RegisterRequest) (*ArttuService.RegisterResponse, error) {
// 	arttu := wompatti.Arttu{}
// 	s.db.Create(&arttu)
// 	return &ArttuService.RegisterResponse{
// 		ArttuId: arttu.ID,
// 	}, nil
// }

// func (s *ArttuServiceServer) AddInterface(stream ArttuService.Arttu_AddInterfaceServer) error {
// 	for {
// 		addInterface, err := stream.Recv()
// 		if err == io.EOF {
// 			return stream.SendAndClose(&ArttuService.AddInterfaceResponse{})
// 		}
// 		if err != nil {
// 			return err
// 		}

// 		iface := &wompatti.Interface{
// 			ArttuID: addInterface.ArttuId,
// 			Name:    addInterface.Interface.Name,
// 			Mac:     addInterface.Interface.Mac,
// 			Index:   addInterface.Interface.Index,
// 			Mtu:     addInterface.Interface.Mtu,
// 			Flags:   addInterface.Interface.Flags,
// 		}

// 		s.db.Create(iface)
// 	}
// 	return nil
// }

// func (s *ArttuServiceServer) ShutdownComputer(ctx context.Context, in *ArttuService.ShutdownComputerRequest) (*ArttuService.ShutdownComputerResponse, error) {
// 	fmt.Println("eka")
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("Shutdown aruttu")
// 	return &ArttuService.ShutdownComputerResponse{
// 		Delay: 1,
// 	}, nil
// }

// func CreateArttuService(db *gorm.DB, port string) *grpc.Server {
// 	lis, err := net.Listen("tcp", ":"+port)
// 	if err != nil {
// 		grpclog.Fatalf("failed to listen: %v", err)
// 	}

// 	//creds, err := credentials.NewServerTLSFromFile("./ssl/server.crt", "./ssl/server.key")
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to generate credentials %v", err)
// 	// }

// 	s := grpc.NewServer()
// 	ArttuService.RegisterArttuServer(s, &ArttuServiceServer{db: db})
// 	// Register reflection service on gRPC server.
// 	reflection.Register(s)
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}

// 	return s
// }
