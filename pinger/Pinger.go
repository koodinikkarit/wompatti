package WompattiPinger

import (
	"fmt"
	"net"
	"time"

	WompattiService "github.com/koodinikkarit/wompatti/wompatti_service"
	"github.com/tatsushid/go-fastping"
)

type Pinger struct {
	pingChannels map[string]chan WompattiService.PingResponse
}

func NewPinger() *Pinger {
	return &Pinger{
		pingChannels: make(map[string]chan WompattiService.PingResponse),
	}
}

func (pinger *Pinger) Ping(ipAddress string) chan WompattiService.PingResponse {
	// var pingerChannel chan WompattiService.PingResponse
	// pingerChannel = pinger.pingChannels[ipAddress]
	// if pingerChannel != nil {

	pingerChannel := make(chan WompattiService.PingResponse)
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", ipAddress)
	if err != nil {
		fmt.Println(err)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		pingerChannel <- WompattiService.PingResponse{
			Addr:    addr.String(),
			Rtt:     rtt.String(),
			Success: true,
		}
	}
	p.OnIdle = func() {
		pingerChannel <- WompattiService.PingResponse{
			Success: false,
		}
	}
	p.RunLoop()
	if err != nil {
		fmt.Println(err)
	}
	//pinger.pingChannels[ipAddress] = pingerChannel
	//}
	return pingerChannel
}
