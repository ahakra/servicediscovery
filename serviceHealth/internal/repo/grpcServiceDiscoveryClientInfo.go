package repository

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/ahakra/servicediscovery/pkg/config"
	"github.com/ahakra/servicediscovery/pkg/serviceDiscoveryProto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientInfoRepo interface {
	GetAllServices() (*serviceDiscoveryProto.Services, error)
	Close() // Add a Close method
}

type ServiceDiscoveryInfoClientRepo struct {
	Client serviceDiscoveryProto.ServiceDiscoveryInfoClient
	Conn   *grpc.ClientConn // Store the gRPC connection
}

func NewClientInfo(conf *config.Config) *ServiceDiscoveryInfoClientRepo {
	fmt.Println("init new client")

	serverAddr := flag.String("addr", conf.Servicediscvoreyserver.Address+":"+strconv.Itoa(conf.Servicediscvoreyserver.Port), "The server address in the format of host:port")

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	initClientInfo := serviceDiscoveryProto.NewServiceDiscoveryInfoClient(conn)
	return &ServiceDiscoveryInfoClientRepo{
		Client: initClientInfo,
		Conn:   conn, // Store the connection for later use
	}
}

func (s *ServiceDiscoveryInfoClientRepo) GetAllServices() (*serviceDiscoveryProto.Services, error) {

	service, err := s.Client.GetAllServices(context.Background(), &serviceDiscoveryProto.EmptyRequest{})
	if err != nil {
		return &serviceDiscoveryProto.Services{}, err
	}

	return service, nil

}

func (s *ServiceDiscoveryInfoClientRepo) Close() {
	// Close the gRPC connection
	if s.Conn != nil {
		s.Conn.Close()
	}
}
