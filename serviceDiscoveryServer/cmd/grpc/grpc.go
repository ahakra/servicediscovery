package main

import (
	"context"

	"github.com/ahakra/servicediscovery/serviceDiscoveryServer/internal/controller"
	//pb "github.com/ahakra/servicediscovery/serviceDiscoveryServer/internal/proto"
	pb "github.com/ahakra/servicediscovery/pkg/serviceDiscoveryProto"
)

//serverinit part
type ServiceDiscoveryServerInit struct {
	pb.UnimplementedServiceDiscoveryInitServer
	Ctrl *controller.MongoCtrl
}

func (sd *ServiceDiscoveryServerInit) RegisterService(ctx context.Context, in *pb.RegisterData) (*pb.ReturnPayload, error) {
	py, err := sd.Ctrl.InsertService(ctx, in)
	return py, err
}

func (sd *ServiceDiscoveryServerInit) DeleteService(ctx context.Context, in *pb.ServiceGuid) (*pb.ReturnPayload, error) {
	py, err := sd.Ctrl.DeleteService(ctx, in)
	return py, err
}

func (sd *ServiceDiscoveryServerInit) UpdateServiceHealth(ctx context.Context, in *pb.RegisterData) (*pb.ReturnPayload, error) {
	py, err := sd.Ctrl.UpdateServiceHealth(ctx, in)
	return py, err
}

//server info part

type ServiceDiscoveryServerInfo struct {
	pb.UnimplementedServiceDiscoveryInfoServer
	Ctrl *controller.MongoCtrl
}

func (sd *ServiceDiscoveryServerInfo) GetAllServices(ctx context.Context, in *pb.EmptyRequest) (*pb.Services, error) {
	py, err := sd.Ctrl.GetAllServices(ctx, in)
	return py, err
}

func (sd *ServiceDiscoveryServerInfo) GetServiceByName(ctx context.Context, in *pb.ServiceName) (*pb.Services, error) {
	py, err := sd.Ctrl.GetServiceByNameSer(ctx, in)
	return py, err
}
