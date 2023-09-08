package repository

import (
	"context"
	"log"

	pb "github.com/ahakra/servicediscovery/serviceDiscoveryServer/internal/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepo interface {
	InsertService(ctx context.Context, in *pb.RegisterData) (*pb.ReturnPayload, error)
	DeleteService(ctx context.Context, in *pb.ServiceGuid) (*pb.ReturnPayload, error)
	UpdateServiceHealth(ctx context.Context, in *pb.RegisterData) (*pb.ReturnPayload, error)
	GetAllServices(ctx context.Context, in *pb.EmptyRequest) (*pb.Services, error)
	GetByNameService(ctx context.Context, in *pb.ServiceName) (*pb.Services, error)
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoServiceRepository(collection *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{collection: collection}
}

func (r *MongoUserRepository) InsertService(ctx context.Context, in *pb.RegisterData) (*pb.ReturnPayload, error) {

	filter := bson.M{"serviceaddress": in.Serviceaddress}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {

		return &pb.ReturnPayload{Data: "service already registered"}, nil

	}

	guid, err := r.collection.InsertOne(ctx, in)
	if err != nil {
		return nil, err
	}
	// Convert the _id field to a string
	objectID := guid.InsertedID.(primitive.ObjectID)

	return &pb.ReturnPayload{Data: objectID.Hex()}, nil
}

func (r *MongoUserRepository) DeleteService(ctx context.Context, in *pb.ServiceGuid) (*pb.ReturnPayload, error) {

	idPrimitive, err := primitive.ObjectIDFromHex(in.Guid)
	if err != nil {
		log.Printf("primitive.ObjectIDFromHex ERROR:", err)
	} else {
		filter := bson.M{"_id": idPrimitive}
		cursor, err := r.collection.DeleteOne(ctx, filter)
		if err != nil {
			return nil, err
		}
		if cursor.DeletedCount == 0 {
			return &pb.ReturnPayload{Data: "Nothing to delete"}, nil
		}
		return &pb.ReturnPayload{Data: "deleted"}, nil

	}
	return nil, err

}

func (r *MongoUserRepository) UpdateServiceHealth(ctx context.Context, in *pb.RegisterData) (*pb.ReturnPayload, error) {

	// Define the filter to find the document you want to update
	filter := bson.M{"serviceaddress": in.Serviceaddress}

	// Define the update operation (in this case, setting a field)
	update := bson.M{"$set": bson.M{"lastupdate": in.Lastupdate}}

	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	return &pb.ReturnPayload{Data: "Updated"}, nil
}

func (r *MongoUserRepository) GetAllServices(ctx context.Context, in *pb.EmptyRequest) (*pb.Services, error) {
	// Define the filter to find the document you want to update

	filter := bson.M{}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Initialize a slice to hold the services
	var services []*pb.RegisterData

	// Iterate over the cursor and decode each document into a Service object
	for cursor.Next(ctx) {
		var service pb.RegisterData
		if err := cursor.Decode(&service); err != nil {
			return nil, err
		}
		services = append(services, &service)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Create a Services response and populate it with the retrieved services
	response := &pb.Services{Services: services}
	return response, nil

}

func (r *MongoUserRepository) GetByNameService(ctx context.Context, in *pb.ServiceName) (*pb.Services, error) {
	// Define the filter to find the document you want to update

	filter := bson.M{"servicename": in.Name}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Initialize a slice to hold the services
	var services []*pb.RegisterData

	// Iterate over the cursor and decode each document into a Service object
	for cursor.Next(ctx) {
		var service pb.RegisterData
		if err := cursor.Decode(&service); err != nil {
			return nil, err
		}
		services = append(services, &service)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Create a Services response and populate it with the retrieved services
	response := &pb.Services{Services: services}
	return response, nil
}
