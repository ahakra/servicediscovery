package main

import (
	"context"
	"flag"
	"log"
	"time"

	// Import the generated client code
	"fmt"

	pb "github.com/ahakra/servicediscovery/serviceDiscoveryClient/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	//tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	//caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr = flag.String("addr", "localhost:1080", "The server address in the format of host:port")
	//serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	//opts = append(opts, flag.Bool("tls", false, "in"))
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()
	initClient := pb.NewServiceDiscoveryInitClient(conn)
	infoClient := pb.NewServiceDiscoveryInfoClient(conn)

	registerData := &pb.RegisterData{
		Servicename:    "logger_service",
		Serviceaddress: "192.168.110.242:1083",
		Lastupdate:     timestamppb.Now(),
		Messages:       []string{"test", "test2"},
	}

	for {

		state := conn.GetState()
		fmt.Println("Connection State:", state.String())
		time.Sleep(1 * time.Second)
		if state == connectivity.Ready {
			break
		}

	}
	// Print the connection state

	y, err := initClient.RegisterService(context.Background(), registerData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(y)
	x, err := initClient.UpdateServiceHealth(context.Background(), registerData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x)

	z, err := infoClient.GetAllServices(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		fmt.Println(err)
	}
	for _, value := range z.Services {
		fmt.Println(value.Servicename + " : " + value.Serviceaddress)
	}

	zz, err := infoClient.GetByNameService(context.Background(), &pb.ServiceName{Name: "auth_service"})
	if err != nil {
		fmt.Println(err)
	}
	for _, value := range zz.Services {
		fmt.Println(value.Servicename + " : " + value.Serviceaddress)
	}

	zzz, err := initClient.DeleteService(context.Background(), &pb.ServiceGuid{Guid: "64f307dc2a79b39b7ab0ad5e"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(zzz)
}
