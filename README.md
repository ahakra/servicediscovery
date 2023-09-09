# Yet Another Project

Its implementation of service discovery using mongodb for storing and updating services

Some methods used

```
ServiceDiscoveryInitServer: 
	RegisterService(context.Context, *RegisterData) (*ReturnPayload, error)
	DeleteService(context.Context, *ServiceGuid) (*ReturnPayload, error)
	UpdateServiceHealth(context.Context, *RegisterData) (*ReturnPayload, error)
	


ServiceDiscoveryInfoServer: 
	GetAllServices(context.Context, *EmptyRequest) (*Services, error)
	GetByNameService(context.Context, *ServiceName) (*Services, error)
	


```
