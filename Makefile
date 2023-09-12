generate:
	@read -p "Enter service dir: " service_dir; \
	mkdir ./$$service_dir ;\
	mkdir ./$$service_dir/cmd ;\
	mkdir ./$$service_dir/cmd/grpc;\
	mkdir ./$$service_dir/internal;\
	mkdir ./$$service_dir/internal/repository;\
	mkdir ./$$service_dir/internal/controller;\
	mkdir ./$$service_dir/internal/database;\
	mkdir ./$$service_dir/internal/proto

	
client:
	go run ./serviceDiscoveryClient/*.go

server:
	go run ./serviceDiscoveryServer/cmd/grpc/*.go

logwriter:
	go run ./loggerService/writer/cmd/grpc/*.go

logreader:
	go run ./loggerService/reader/cmd/grpc/*.go


generateLogWriterProto:
	
	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./loggerService/writer/internal/proto/logWriter.proto


generateLogReaderProto:
	
	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./loggerService/reader/internal/proto/logReader.proto


generateserviceDiscoveryServerproto:

	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./serviceDiscoveryServer/internal/proto/servicediscovery.proto

generatepgenerateserviceDiscoveryClientprotoroto:
	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./serviceDiscoveryClient/proto/serviceDiscovery.proto




generateauthServiceproto:

	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./authService/internal/proto/authService.proto


startmongo:
	docker start c9b138f95d89

