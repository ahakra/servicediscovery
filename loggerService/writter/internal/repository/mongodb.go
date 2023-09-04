package repository

import (
	"context"

	pb "github.com/ahakra/servicediscovery/loggerService/writter/internal/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MongoRepo interface {
	SaveLog(ctx context.Context, lgd *pb.LogData) (*emptypb.Empty, error)
}

type MongoLogRepository struct {
	database *mongo.Database
}

func NewMongoServiceRepository(database *mongo.Database) *MongoLogRepository {
	return &MongoLogRepository{database: database}
}

func (r *MongoLogRepository) SaveLog(ctx context.Context, lgd *pb.LogData) (*emptypb.Empty, error) {

}
