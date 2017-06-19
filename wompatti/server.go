package wompatti

import (
	"github.com/jinzhu/gorm"
	"github.com/koodinikkarit/wompatti/db"
	"google.golang.org/grpc"
)

type WompattiServer struct {
	db              *gorm.DB
	wompattiService *grpc.Server
	arttuService    *grpc.Server
}

func CreateWompattiServer(
	dbUser string,
	dbPass string,
	dbIp string,
	dbPort string,
	dbName string,
	wompattiServicePort string,
	arttuServicePort string) {

	db := wompatti.CreateDb(dbUser, dbPass, dbIp, dbPort, dbName)

	CreateWompattiService(db, wompattiServicePort)
	//CreateArttuService(db, arttuServicePort)

	// return &WompattiServer{
	// 	db: db,
	// }
}
