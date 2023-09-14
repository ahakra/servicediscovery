package repository

import (
	"context"

	pb "github.com/ahakra/servicediscovery/loggerService/writer/internal/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MongoRepo interface {
	SaveLog(ctx context.Context, lgd *pb.LogPayload) (*emptypb.Empty, error)
}

type MongoLogRepository struct {
	database *mongo.Database
}

func NewMongoServiceRepository(database *mongo.Database) *MongoLogRepository {
	return &MongoLogRepository{database: database}
}

func (r *MongoLogRepository) SaveLog(ctx context.Context, lgd *pb.LogPayload) (*emptypb.Empty, error) {

	collection := r.database.Collection(lgd.CollectionName.Name)

	lgd.LogData.CreatedAt = timestamppb.Now()
	collection.InsertOne(ctx, lgd.LogData)
	return &emptypb.Empty{}, nil
}
