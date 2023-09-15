package main

import (
	"fmt"
	"strconv"

	"net"

	"github.com/ahakra/servicediscovery/config"
	pb "github.com/ahakra/servicediscovery/pkg/serviceDiscoveryProto"
	"github.com/ahakra/servicediscovery/serviceDiscoveryServer/internal/controller"
	"github.com/ahakra/servicediscovery/serviceDiscoveryServer/internal/database"
	"github.com/ahakra/servicediscovery/serviceDiscoveryServer/internal/repository"
	"google.golang.org/grpc"
)

func main() {

	// config.LoadEnv(".env")

	// conf := config.New()

	conf := config.NewFromJson("config.json")
	fmt.Println(conf)
	collection := database.Connect(
		conf.Mongodatabase.URL,
		conf.Servicediscvoreyserver.Database,
		conf.Servicediscvoreyserver.Collection)

	serverport := conf.Servicediscvoreyserver.Port

	defer database.Disconnect(collection.Database().Client()) // Close MongoDB client when the program exits

	repo := repository.NewMongoServiceRepository(collection)
	ctrl := controller.NewMongoCtrl(repo)

	server := grpc.NewServer()
	serviceDiscoveryServerinit := &ServiceDiscoveryServerInit{Ctrl: ctrl}
	serviceDicoveryServiceinfo := &ServiceDiscoveryServerInfo{Ctrl: ctrl}

	pb.RegisterServiceDiscoveryInitServer(server, serviceDiscoveryServerinit)
	pb.RegisterServiceDiscoveryInfoServer(server, serviceDicoveryServiceinfo)

	listen, err := net.Listen("tcp", ":"+strconv.Itoa(serverport)) // Specify your desired host and port
	if err != nil {
		fmt.Println("Failed to listen:", err)
		return
	}
	defer listen.Close()

	fmt.Println("Server is running on+:" + strconv.Itoa(serverport) + "..")
	if err := server.Serve(listen); err != nil {
		fmt.Println("Failed to serve:", err)
		return
	}
}
