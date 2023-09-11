package repository

import (
	"context"

	"github.com/ahakra/servicediscovery/loggerService/reader/internal/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepo interface {
	ReadLog(ctx context.Context, col string, filter *bson.M) (*proto.ReturnedData, error)
}

type MongoLogRepository struct {
	database *mongo.Database
}

func NewMongoServiceRepository(database *mongo.Database) *MongoLogRepository {
	return &MongoLogRepository{database: database}
}

func (r *MongoLogRepository) ReadLog(ctx context.Context, col string, filter *bson.M) (*proto.ReturnedData, error) {

	collection := r.database.Collection(col)
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var services []*proto.LogData

	for cursor.Next(ctx) {
		var service proto.LogData
		if err := cursor.Decode(&service); err != nil {
			return nil, err
		}
		services = append(services, &service)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	response := &proto.ReturnedData{Data: services}
	return response, nil

}
