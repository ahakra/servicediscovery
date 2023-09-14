package main

import (
	"fmt"

	"net"

	"github.com/ahakra/servicediscovery/config"
	pb "github.com/ahakra/servicediscovery/pkg/serviceDiscoveryProto"
	"github.com/ahakra/servicediscovery/serviceDiscoveryServer/internal/controller"
	"github.com/ahakra/servicediscovery/serviceDiscoveryServer/internal/database"
	"github.com/ahakra/servicediscovery/serviceDiscoveryServer/internal/repository"
	"google.golang.org/grpc"
)

func main() {

	config.LoadEnv(".env")

	conf := config.New()

	collection := database.Connect(conf.Database.URL, conf.ServiceDiscoveryDatabase, conf.ServiceDiscoveryCollection)
	defer database.Disconnect(collection.Database().Client()) // Close MongoDB client when the program exits

	repo := repository.NewMongoServiceRepository(collection)
	ctrl := controller.NewMongoCtrl(repo)

	server := grpc.NewServer()
	serviceDiscoveryServerinit := &ServiceDiscoveryServerInit{Ctrl: ctrl}
	serviceDicoveryServiceinfo := &ServiceDiscoveryServerInfo{Ctrl: ctrl}

	pb.RegisterServiceDiscoveryInitServer(server, serviceDiscoveryServerinit)
	pb.RegisterServiceDiscoveryInfoServer(server, serviceDicoveryServiceinfo)

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
