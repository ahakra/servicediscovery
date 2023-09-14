package controller

import (
	"context"

	pb "github.com/ahakra/servicediscovery/loggerService/writer/internal/proto"
	"github.com/ahakra/servicediscovery/loggerService/writer/internal/repository"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MongoCtrl struct {
	MongoRepo repository.MongoRepo
}

func NewMongoCtrl(mr repository.MongoRepo) *MongoCtrl {
	return &MongoCtrl{MongoRepo: mr}
}

func (mc *MongoCtrl) SaveLog(ctx context.Context, in *pb.LogPayload) (*emptypb.Empty, error) {
	_, err := mc.MongoRepo.SaveLog(ctx, in)
	return &emptypb.Empty{}, err
}
