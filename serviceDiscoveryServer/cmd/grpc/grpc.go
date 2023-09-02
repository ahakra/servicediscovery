package main

import (
	"context"

	"github.com/ahakra/servicediscovery/serviceDiscoveryServer/internal/controller"
	pb "github.com/ahakra/servicediscovery/serviceDiscoveryServer/internal/proto"
)

type ServiceDiscoveryServer struct {
	pb.UnimplementedServiceDiscoveryInitServer
	Ctrl *controller.MongoCtrl
}

func (sd *ServiceDiscoveryServer) RegisterService(ctx context.Context, in *pb.RegisterData) (*pb.ReturnPayload, error) {
	py, err := sd.Ctrl.InsertService(ctx, in)
	return py, err
}

func (sd *ServiceDiscoveryServer) DeleteService(ctx context.Context, in *pb.ServiceGuid) (*pb.ReturnPayload, error) {
	py, err := sd.Ctrl.DeleteService(ctx, in)
	return py, err
}

func (sd *ServiceDiscoveryServer) UpdateServiceHealth(ctx context.Context, in *pb.RegisterData) (*pb.ReturnPayload, error) {
	py, err := sd.Ctrl.UpdateServiceHealth(ctx, in)
	return py, err
}

func (sd *ServiceDiscoveryServer) GetAllServices(ctx context.Context, in *pb.EmptyRequest) (*pb.Services, error) {
	py, err := sd.Ctrl.GetAllServices(ctx, in)
	return py, err
}

func (sd *ServiceDiscoveryServer) GetByNameService(ctx context.Context, in *pb.ServiceName) (*pb.Services, error) {
	py, err := sd.Ctrl.GetByNameService(ctx, in)
	return py, err
}
