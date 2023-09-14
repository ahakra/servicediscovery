package controller

import (
	"context"
	"log"

	pb "github.com/ahakra/servicediscovery/loggerService/reader/internal/proto"
	"github.com/ahakra/servicediscovery/loggerService/reader/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type IMongoCtrl interface {
	ReadLog(ctx context.Context, in *pb.LogFilter) (*pb.ReturnedData, error)
}

type MongoCtrl struct {
	MongoRepo repository.MongoRepo
}

func NewMongoCtrl(mr repository.MongoRepo) *MongoCtrl {
	return &MongoCtrl{MongoRepo: mr}
}

func (mc *MongoCtrl) ReadLog(ctx context.Context, in *pb.LogFilter) (*pb.ReturnedData, error) {
	filter := bson.M{}
	log.Println(in.LogType)
	log.Println(pb.LogType_LOG_TYPE_NIL.Enum())
	if in.LogType != pb.LogType_LOG_TYPE_NIL && in.LogType.Enum() != nil {
		filter["logtype"] = in.LogType
		log.Println("inside enum cond")
	}
	if in.DateFrom != nil {
		filter["createdat.seconds"] = bson.M{
			"$gte": in.DateFrom.Seconds,
		}
	}
	if in.DateTo != nil {
		filter["createdat.seconds"] = bson.M{
			"$lte": in.DateTo.Seconds,
		}
	}

	retuneddata, err := mc.MongoRepo.ReadLog(ctx, in.CollectionName, &filter)
	if err != nil {
		return &pb.ReturnedData{}, nil
	}
	return retuneddata, nil

}
