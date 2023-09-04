client:
	go run ./serviceDiscoveryClient/*.go

server:
	go run ./serviceDiscoveryServer/cmd/grpc/*.go


generateLogProto:
	
	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./loggerService/writter/internal/proto/logWritter.proto