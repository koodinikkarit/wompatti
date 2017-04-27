

dependecies:
	go get google.golang.org/grpc
	go get github.com/jinzhu/gorm
	go get github.com/jinzhu/gorm/dialects/mysql

	
proto:
	protoc --proto_path=wompatti_service --proto_path=wompatti_service --go_out=plugins=grpc:wompatti_service wompatti_service/*.proto

all: 
	dependecies