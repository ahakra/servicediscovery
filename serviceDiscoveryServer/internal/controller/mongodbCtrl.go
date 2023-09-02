package controller

import (
	"context"

	pb "github.com/ahakra/servicediscovery/serviceDiscoveryServer/internal/proto"
	"github.com/ahakra/servicediscovery/serviceDiscoveryServer/internal/repository"
)

type MongoCtrl struct {
	MongoRepo repository.MongoRepo
}

func NewMongoCtrl(mr repository.MongoRepo) *MongoCtrl {
	return &MongoCtrl{MongoRepo: mr}
}

func (mc *MongoCtrl) InsertService(ctx context.Context, in *pb.RegisterData) (*pb.ReturnPayload, error) {
	py, err := mc.MongoRepo.InsertService(ctx, in)
	return py, err
}
func (mc *MongoCtrl) DeleteService(ctx context.Context, in *pb.ServiceGuid) (*pb.ReturnPayload, error) {

	py, err := mc.MongoRepo.DeleteService(ctx, in)
	return py, err
}
func (mc *MongoCtrl) UpdateServiceHealth(ctx context.Context, in *pb.RegisterData) (*pb.ReturnPayload, error) {
	py, err := mc.MongoRepo.UpdateServiceHealth(ctx, in)
	return py, err
}

func (mc *MongoCtrl) GetAllServices(ctx context.Context, in *pb.EmptyRequest) (*pb.Services, error) {

	py, err := mc.MongoRepo.GetAllServices(ctx, in)
	return py, err
}
func (mc *MongoCtrl) GetByNameService(ctx context.Context, in *pb.ServiceName) (*pb.Services, error) {
	py, err := mc.MongoRepo.GetByNameService(ctx, in)
	return py, err
}
