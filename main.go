package main

import (
	"github.com/koodinikkarit/wompatti/context"
	"github.com/koodinikkarit/wompatti/wompatti_service_server"
)

func main() {
	contextGenerator := WompattiContext.NewContextGenerator(
		"jaska",
		"asdf321",
		"localhost",
		"3306",
		"wompatti",
		true,
	)

	WompattiServiceServer.NewWompattiService(
		contextGenerator.NewContext,
		"3112",
	)
}
