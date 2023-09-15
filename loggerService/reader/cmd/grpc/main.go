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

	"github.com/ahakra/servicediscovery/config"
	"github.com/ahakra/servicediscovery/loggerService/reader/internal/controller"
	"github.com/ahakra/servicediscovery/loggerService/reader/internal/database"
	"github.com/ahakra/servicediscovery/loggerService/reader/internal/handler"
	"github.com/ahakra/servicediscovery/loggerService/reader/internal/proto"
	"github.com/ahakra/servicediscovery/loggerService/reader/internal/repository"
	pb "github.com/ahakra/servicediscovery/pkg/serviceDiscoveryProto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var returnedguid string

func main() {

	conf := config.NewFromJson("config.json")

	var port = conf.Loggerservicereader.StartingPort

	//creating a channel to pass when Ctrl+C is pressed
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	db := database.Connect(conf.Mongodatabase.URL, conf.Loggerservicereader.Database)
	repo := repository.NewMongoServiceRepository(db)
	ctrl := controller.NewMongoCtrl(repo)
	grpcHandler := handler.NewLogReaderHandler(ctrl)

	//starting logreader service
	server := grpc.NewServer()
	//serviceDiscoveryServerinit := &LogReader{Ctrl: ctrl}
	proto.RegisterLogReaderServer(server, grpcHandler)

	//this is done so port will be dynamically created if port is in use starting from specific port number
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

	//Section related to registering with servicediscovery service

	//part for service disover registeration
	registerData := &pb.RegisterData{
		Servicename:    conf.Loggerservicereader.Name,
		Serviceaddress: conf.Loggerservicereader.Address + ":" + strconv.Itoa(port),
		Lastupdate:     timestamppb.Now(),
		Messages:       []string{"test", "test2"},
	}

	serverAddr := flag.String("addr", conf.Servicediscvoreyserver.Address+":"+strconv.Itoa(conf.Servicediscvoreyserver.Port), "The server address in the format of host:port")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()
	initClient := pb.NewServiceDiscoveryInitClient(conn)

	go func() {

		for {
			y, err := initClient.RegisterService(context.Background(), registerData)
			if err != nil {

				fmt.Println(err)

			} else {
				returnedguid = y.Data
				break

			}
			time.Sleep(10 * time.Second)
		}

		defer initClient.DeleteService(context.Background(), &pb.ServiceGuid{Guid: returnedguid})
		for {
			registerData := &pb.RegisterData{
				Servicename:    conf.Loggerservicereader.Name,
				Serviceaddress: conf.Loggerservicereader.Address + ":" + strconv.Itoa(port),
				Lastupdate:     timestamppb.Now(),
				Messages:       []string{"test", "test2"},
			}

			_, err := initClient.UpdateServiceHealth(context.Background(), registerData)

			if err != nil {
				fmt.Println(err)
			}
			log.Println("updating " + conf.Loggerservicereader.Name + "  service")
			time.Sleep(10 * time.Second)

		}

	}()
	go func() {
		sig := <-sigChan
		fmt.Printf("Received signal: %v\n", sig)

		initClient.DeleteService(context.Background(), &pb.ServiceGuid{Guid: returnedguid})
		listen.Close()

		os.Exit(1)
	}()

	//log reader service
	fmt.Println("Server is running on :" + strconv.Itoa(port))

	if err := server.Serve(listen); err != nil {
		fmt.Println("Failed to serve:", err)
		return
	}

}
