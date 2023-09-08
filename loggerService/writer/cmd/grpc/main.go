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
	"github.com/ahakra/servicediscovery/loggerService/writer/internal/proto"
	"github.com/ahakra/servicediscovery/loggerService/writer/internal/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const uri = "mongodb://localhost:27017"
const serviceDiscoveryPort = 1080

var returnedguid string
var port = 8111

func main() {

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 	data := struct {
	// 		name, nationality string
	// 		age               int
	// 		score             float64
	// 	}{
	// 		name:        "Anna_Hurry",
	// 		nationality: "England",
	// 		age:         21,
	// 		score:       9.5,
	// 	}
	// //
	// firstLog := proto.LogData{
	// 	CollectionName: "logs",
	// 	LogType:        proto.LogType_LOG_TYPE_INFORMATION,
	// }

	db := database.Connect(uri, "loggerservice")
	repo := repository.NewMongoServiceRepository(db)
	ctrl := controller.NewMongoCtrl(repo)

	server := grpc.NewServer()
	serviceDiscoveryServerinit := &LogWritter{Ctrl: ctrl}

	proto.RegisterLogwriterServer(server, serviceDiscoveryServerinit)

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

	//part for service disover registeration

	registerData := &proto.RegisterData{
		Servicename:    "logger_writer",
		Serviceaddress: "localhost:" + strconv.Itoa(port),
		Lastupdate:     timestamppb.Now(),
		Messages:       []string{"test", "test2"},
	}

	serverAddr := flag.String("addr", "localhost:"+strconv.Itoa(serviceDiscoveryPort), "The server address in the format of host:port")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()
	initClient := proto.NewServiceDiscoveryInitClient(conn)
	
	y, err := initClient.RegisterService(context.Background(), registerData)
	if err != nil {

		// fmt.Println(y.Data)
		returnedguid = y.Data
		fmt.Println("returned guid: " + returnedguid)

	}

	returnedguid = y.Data
	fmt.Println("returned guid: " + returnedguid)

	go func() {
		
		for {
			registerData := &proto.RegisterData{
				Servicename:    "logger_writer",
				Serviceaddress: "localhost:" + strconv.Itoa(port),
				Lastupdate:     timestamppb.Now(),
				Messages:       []string{"test", "test2"},
			}

			_, err := initClient.UpdateServiceHealth(context.Background(), registerData)

			if err != nil {
				fmt.Println(err)
			}
			log.Println("updating service")
			time.Sleep(10 * time.Second)
			
		}
		defer initClient.DeleteService(context.Background(), &proto.ServiceGuid{Guid: returnedguid})
	
	}()
	go func() {
		sig := <-sigChan
		fmt.Printf("Received signal: %v\n", sig)

		initClient.DeleteService(context.Background(), &proto.ServiceGuid{Guid: returnedguid})
		listen.Close()

		os.Exit(1)
	}()

	fmt.Println("Server is running on :" + strconv.Itoa(port))

	if err := server.Serve(listen); err != nil {
		fmt.Println("Failed to serve:", err)
		return
	}

}
