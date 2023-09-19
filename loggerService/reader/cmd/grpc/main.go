package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/ahakra/servicediscovery/loggerService/reader/internal/controller"
	"github.com/ahakra/servicediscovery/loggerService/reader/internal/database"
	"github.com/ahakra/servicediscovery/loggerService/reader/internal/handler"
	"github.com/ahakra/servicediscovery/loggerService/reader/internal/proto"
	"github.com/ahakra/servicediscovery/loggerService/reader/internal/repository"
	"github.com/ahakra/servicediscovery/pkg/config"
	helper "github.com/ahakra/servicediscovery/pkg/helpers"
	pb "github.com/ahakra/servicediscovery/pkg/serviceDiscoveryProto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {

	conf := config.NewFromJson("config.json")
	var port = conf.Loggerservicereader.StartingPort

	sigChan := make(chan os.Signal, 1)
	returnedGuidChan := make(chan string, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	db := database.Connect(conf.Mongodatabase.URL, conf.Loggerservicereader.Database)
	repo := repository.NewMongoServiceRepository(db)
	ctrl := controller.NewMongoCtrl(repo)
	grpcHandler := handler.NewLogReaderHandler(ctrl)

	//Register client
	//Register Section for client
	serverAddr := flag.String("addr", conf.Servicediscvoreyserver.Address+":"+strconv.Itoa(conf.Servicediscvoreyserver.Port), "The server address in the format of host:port")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Println("starting")
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	registerData := &pb.RegisterData{
		Servicename:    conf.Loggerservicereader.Name,
		Serviceaddress: conf.Loggerservicereader.Address + ":" + strconv.Itoa(port),
		Lastupdate:     timestamppb.Now(),
		Messages:       []string{"test", "test2"},
	}

	helper := helper.HelperData{
		Connection:   conn,
		RegisterData: registerData,
		Conf:         *conf,
		Name:         conf.Loggerservicereader.Name,
	}
	ctx := context.Background()

	go helper.RegisterService(ctx, returnedGuidChan)
	go helper.UpdateServiceHealth(ctx)
	go helper.DeleteService(ctx, returnedGuidChan, sigChan)

	fmt.Println("Server is running on :" + strconv.Itoa(port))

	//starting logreader service
	server := grpc.NewServer()
	proto.RegisterLogReaderServer(server, grpcHandler)
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(port)) // Specify your desired host and port
	if err != nil {
		for {

			listen, err = net.Listen("tcp", ":"+strconv.Itoa(port)) // Specify your desired host and port
			if err != nil {
				port = port + 1
				log.Println(err)
				time.Sleep(1 * time.Second)
				log.Println(port)
			} else {
				break
			}
		}
	}

	if err := server.Serve(listen); err != nil {
		fmt.Println("Failed to serve:", err)
		return
	}

}
