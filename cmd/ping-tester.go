package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/tatsushid/go-fastping"
)

func main() {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", "10.70.0.30")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		fmt.Println("finish")
	}

	p.RunLoop()
	fmt.Println("End of this")
	if err != nil {
		fmt.Println(err)
	}
}
