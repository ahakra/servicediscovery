# First Golang Project

Its implementation of service discovery using mongodb for storing and updating services

Some methods used

```
	InsertService(ctx context.Context, in *pb.RegisterData) (*pb.ReturnPayload, error)
	DeleteService(ctx context.Context, in *pb.ServiceGuid) (*pb.ReturnPayload, error)
	UpdateServiceHealth(ctx context.Context, in *pb.RegisterData) (*pb.ReturnPayload, error)
	GetAllServices(ctx context.Context, in *pb.EmptyRequest) (*pb.Services, error)
	GetByNameService(ctx context.Context, in *pb.ServiceName) (*pb.Services, error)


```
