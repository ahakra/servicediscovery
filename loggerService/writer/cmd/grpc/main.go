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

	"github.com/ahakra/servicediscovery/loggerService/writer/internal/controller"
	"github.com/ahakra/servicediscovery/loggerService/writer/internal/database"
	"github.com/ahakra/servicediscovery/loggerService/writer/internal/handler"
	"github.com/ahakra/servicediscovery/loggerService/writer/internal/proto"

	"github.com/ahakra/servicediscovery/loggerService/writer/internal/repository"
	"github.com/ahakra/servicediscovery/pkg/config"
	helper "github.com/ahakra/servicediscovery/pkg/helpers"
	pb "github.com/ahakra/servicediscovery/pkg/serviceDiscoveryProto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {

	sigChan := make(chan os.Signal, 1)
	returnedGuidChan := make(chan string, 1)
	onRegisterChan := make(chan bool, 1)
	onInitServerChan := make(chan bool, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	conf := config.NewFromJson("config.json")
	var serviceDiscoveryPort = conf.Servicediscvoreyserver.Port
	var port = conf.Loggerservicewriter.StartingPort

	db := database.Connect(conf.Mongodatabase.URL, conf.Loggerservicewriter.Database)
	repo := repository.NewMongoServiceRepository(db)
	ctrl := controller.NewMongoCtrl(repo)
	grpcHandler := handler.NewLogReaderHandler(ctrl)

	//Register Section for client
	serverAddr := flag.String("addr", conf.Servicediscvoreyserver.Address+":"+strconv.Itoa(serviceDiscoveryPort), "The server address in the format of host:port")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	registerData := &pb.RegisterData{
		Servicename:    conf.Loggerservicewriter.Name,
		Serviceaddress: conf.Loggerservicewriter.Address + ":" + strconv.Itoa(port),
		Lastupdate:     timestamppb.Now(),
		Messages:       []string{"test", "test2"},
	}

	helper := helper.HelperData{
		Connection:   conn,
		RegisterData: registerData,
		Conf:         *conf,
		Name:         conf.Loggerservicewriter.Name,
	}
	ctx := context.Background()

	go helper.RegisterService(ctx, returnedGuidChan, onInitServerChan, onRegisterChan)
	go helper.UpdateServiceHealth(ctx, onRegisterChan)
	go helper.DeleteService(ctx, returnedGuidChan, sigChan)

	//starting logwriter service
	server := grpc.NewServer()
	proto.RegisterLogwriterServer(server, grpcHandler)
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
	helper.RegisterData.Serviceaddress = conf.Loggerservicewriter.Address + ":" + strconv.Itoa(port)
	onInitServerChan <- true

	fmt.Println("Server is running on :" + strconv.Itoa(port))
	if err := server.Serve(listen); err != nil {
		fmt.Println("Failed to serve:", err)
		return
	}

}
