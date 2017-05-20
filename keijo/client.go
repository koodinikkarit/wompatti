package keijo

import (
	"flag"
	"fmt"

	"google.golang.org/grpc"

	KeijoService "github.com/koodinikkarit/wompatti/keijo_service"
)

type KeijoClient struct {
}

// CreateKeijoClient creates new client
func CreateKeijoClient(ip string, port string) KeijoService.KeijoClient {
	var (
		serverAddr = flag.String("server_addr", ip+":"+port, "Serverin ip ja portti")
	)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		fmt.Println("Virhe yhteydessa")
	}

	client := KeijoService.NewKeijoClient(conn)
	return client
}
