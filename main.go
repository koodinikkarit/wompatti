package main

import (
	"github.com/koodinikkarit/wompatti/keijo"
	KeijoService "github.com/koodinikkarit/wompatti/keijo_service"

	"golang.org/x/net/context"
)

func main() {
	// context := wompatti.CreateContext(
	// 	"jaska",
	// 	"asdf321",
	// 	"localhost",
	// 	"3306",
	// 	"wompatti",
	// 	"7856",
	// )

	// context.Start()

	client := keijo.CreateKeijoClient("10.70.0.6", "12345")
	//client.TurnOn(context.Background(), &KeijoService.TurnOnRequest{Address: 0})
	client.ChangeSource(context.Background(), &KeijoService.ChangeSourceRequest{Source: KeijoService.Devices_RECORDING1, Destination: KeijoService.Devices_BROADCAST, SourceNumber: 3})
}
