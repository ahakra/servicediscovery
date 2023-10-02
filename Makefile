client:
	go run ./serviceDiscoveryClient/*.go

server:
	go run ./serviceDiscoveryServer/cmd/grpc/*.go

logwriter:
	go run ./loggerService/writer/cmd/grpc/*.go

logreader:
	go run ./loggerService/reader/cmd/grpc/*.go

generateLogProto:
	
	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./loggerService/writer/internal/proto/logWriter.proto


generateLogReaderProto:
	
	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./loggerService/reader/internal/proto/logReader.proto


generateserviceDiscoveryServerproto:

	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./serviceDiscoveryServer/internal/proto/servicediscovery.proto

generatepgenerateserviceDiscoveryClientprotoroto:
	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./serviceDiscoveryClient/proto/serviceDiscovery.proto



protoserviceRegisterer:

	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./pkg/serviceDiscoveryProto/servicediscovery.proto
																										  

protoFMS:
	protoc --go_out=.   --go_opt=paths=source_relative ./FMSService/internal/FMSProto/records.proto

startmongo:
	docker start c9b138f95d89