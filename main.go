package main

import (
	"os"

	"github.com/koodinikkarit/wompatti/context"
	"github.com/koodinikkarit/wompatti/wompatti_service_server"
)

func main() {
	contextGenerator := WompattiContext.NewContextGenerator(
		os.Getenv("WOMPATTI_DB_USERNAME"),
		os.Getenv("WOMPATTI_DB_PASSWORD"),
		os.Getenv("WOMPATTI_DB_IP"),
		os.Getenv("WOMPATTI_DB_PORT"),
		os.Getenv("WOMPATTI_DB_NAME"),
		true,
	)

	WompattiServiceServer.NewWompattiService(
		contextGenerator.NewContext,
		os.Getenv("WOMPATTI_PORT"),
	)
}
