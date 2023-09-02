package main

import (
	"fmt"

	"net"

	"github.com/ahakra/servicediscovery/serviceDiscoveryServer/internal/controller"
	"github.com/ahakra/servicediscovery/serviceDiscoveryServer/internal/database"
	pb "github.com/ahakra/servicediscovery/serviceDiscoveryServer/internal/proto"
	"github.com/ahakra/servicediscovery/serviceDiscoveryServer/internal/repository"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const uri = "mongodb://root:password@localhost:27017"

func main() {

	p := pb.RegisterData{
		Servicename:    "auth_service",
		Serviceaddress: "192.168.110.242:1080",
		Lastupdate:     timestamppb.Now(),
		Messages:       []string{"test", "test2"},
	}

	collection := database.Connect(uri, "servicediscovery", "services")
	defer database.Disconnect(collection.Database().Client()) // Close MongoDB client when the program exits

	repo := repository.NewMongoServiceRepository(collection)
	ctrl := controller.NewMongoCtrl(repo)

	fmt.Println(p.Servicename)
	// var opts []grpc.DialOption

	// opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	server := grpc.NewServer()
	serviceDiscoveryServer := &ServiceDiscoveryServer{Ctrl: ctrl}
	pb.RegisterServiceDiscoveryInitServer(server, serviceDiscoveryServer)

	listen, err := net.Listen("tcp", ":1080") // Specify your desired host and port
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}
	defer listen.Close()

	fmt.Println("Server is running on :1080...")
	if err := server.Serve(listen); err != nil {
		fmt.Println("Failed to serve:", err)
		return
	}
}
