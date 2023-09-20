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

var logServiceReaderPort int

func main() {

	channelData := helper.ChannelData{
		OnInitChan:      make(chan bool),
		OnRegisterChan:  make(chan bool),
		RetunedGuidChan: make(chan string),
		CancelChan:      make(chan os.Signal, 1),
	}
	conf := config.NewFromJson("config.json")

	signal.Notify(channelData.CancelChan, syscall.SIGINT, syscall.SIGTERM)

	db := database.Connect(conf.Mongodatabase.URL, conf.Loggerservicereader.Database)
	repo := repository.NewMongoServiceRepository(db)
	ctrl := controller.NewMongoCtrl(repo)
	grpcHandler := handler.NewLogReaderHandler(ctrl)

	serverAddr := flag.String("addr", conf.Servicediscvoreyserver.Address+":"+strconv.Itoa(conf.Servicediscvoreyserver.Port), "The server address in the format of host:port")

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	registerData := &pb.RegisterData{
		Servicename:    conf.Loggerservicereader.Name,
		Serviceaddress: conf.Loggerservicereader.Address + ":" + strconv.Itoa(conf.Loggerservicereader.StartingPort),
		Lastupdate:     timestamppb.Now(),
		Messages:       []string{"test", "test2"},
	}

	logReaderHelper := helper.HelperData{
		Connection:   conn,
		RegisterData: registerData,
		Conf:         *conf,
		Name:         conf.Loggerservicereader.Name,
	}

	//Register Section for client

	defer conn.Close()

	ctx := context.Background()

	go logReaderHelper.RegisterService(ctx, channelData)
	go logReaderHelper.UpdateServiceHealth(ctx, channelData)
	go logReaderHelper.DeleteService(ctx, channelData)

	//starting logreader service
	server := grpc.NewServer()
	proto.RegisterLogReaderServer(server, grpcHandler)
	logServiceReaderPort = conf.Loggerservicereader.StartingPort
	listen, err := net.Listen("tcp", conf.Loggerservicereader.Address+":"+strconv.Itoa(logServiceReaderPort)) // Specify your desired host and port
	if err != nil {
		for {

			listen, err = net.Listen("tcp", conf.Loggerservicereader.Address+":"+strconv.Itoa(logServiceReaderPort)) // Specify your desired host and port

			if err != nil {
				logServiceReaderPort = logServiceReaderPort + 1
				log.Println(err)
				time.Sleep(1 * time.Second)

			} else {
				break
			}
		}
	}
	logReaderHelper.RegisterData.Serviceaddress = conf.Loggerservicereader.Address + ":" + strconv.Itoa(logServiceReaderPort)
	channelData.OnInitChan <- true

	fmt.Println("Server is running on :" + strconv.Itoa(logServiceReaderPort))
	if err := server.Serve(listen); err != nil {
		fmt.Println("Failed to serve:", err)
		return
	}

}
