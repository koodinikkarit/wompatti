

dependecies:
	go get google.golang.org/grpc
	go get github.com/golang/protobuf/proto
	go get github.com/golang/protobuf/protoc-gen-go
proto:
	protoc --proto_path=wompatti_service --go_out=pb wompatti_service/*.proto