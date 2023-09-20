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

var logWriterPort int

func main() {
	channelData := helper.ChannelData{
		OnInitChan:      make(chan bool),
		OnRegisterChan:  make(chan bool),
		RetunedGuidChan: make(chan string),
		CancelChan:      make(chan os.Signal, 1),
	}

	signal.Notify(channelData.CancelChan, syscall.SIGINT, syscall.SIGTERM)

	conf := config.NewFromJson("config.json")

	db := database.Connect(conf.Mongodatabase.URL, conf.Loggerservicewriter.Database)
	repo := repository.NewMongoServiceRepository(db)
	ctrl := controller.NewMongoCtrl(repo)
	grpcHandler := handler.NewLogReaderHandler(ctrl)

	//Register Section for client
	serverAddr := flag.String("addr", conf.Servicediscvoreyserver.Address+":"+strconv.Itoa(conf.Servicediscvoreyserver.Port), "The server address in the format of host:port")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	registerData := &pb.RegisterData{
		Servicename:    conf.Loggerservicewriter.Name,
		Serviceaddress: conf.Loggerservicewriter.Address + ":" + strconv.Itoa(conf.Loggerservicewriter.StartingPort),
		Lastupdate:     timestamppb.Now(),
		Messages:       conf.Loggerservicewriter.Messages,
	}

	logWriterHelper := helper.HelperData{
		Connection:   conn,
		RegisterData: registerData,
		Conf:         *conf,
		Name:         conf.Loggerservicewriter.Name,
	}

	defer conn.Close()

	ctx := context.Background()

	go logWriterHelper.RegisterService(ctx, channelData)
	go logWriterHelper.UpdateServiceHealth(ctx, channelData)
	go logWriterHelper.DeleteService(ctx, channelData)

	//starting logwriter service
	server := grpc.NewServer()
	proto.RegisterLogwriterServer(server, grpcHandler)
	logWriterPort = conf.Loggerservicewriter.StartingPort
	listen, err := net.Listen("tcp", conf.Loggerservicewriter.Address+":"+strconv.Itoa(logWriterPort)) // Specify your desired host and port
	if err != nil {
		for {

			listen, err = net.Listen("tcp", conf.Loggerservicewriter.Address+":"+strconv.Itoa(logWriterPort)) // Specify your desired host and port
			if err != nil {
				logWriterPort = logWriterPort + 1
				log.Println(err)
				time.Sleep(1 * time.Second)

			} else {
				break
			}
		}
	}
	logWriterHelper.RegisterData.Serviceaddress = conf.Loggerservicewriter.Address + ":" + strconv.Itoa(logWriterPort)
	channelData.OnInitChan <- true

	fmt.Println(conf.Loggerservicewriter.Name + " server is running on :" + strconv.Itoa(logWriterPort))
	if err := server.Serve(listen); err != nil {
		fmt.Println("Failed to serve:", err)
		return
	}

}
