package clienthelper

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ahakra/servicediscovery/pkg/model"
	pb "github.com/ahakra/servicediscovery/pkg/serviceDiscoveryProto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServiceDiscoveryClient struct {
	Conn *grpc.ClientConn
}

func NewServiceDiscoveryClient(connString *model.ConnectionString) ServiceDiscoveryClient {
	serverAddr := flag.String("addr", connString.Address+":"+strconv.Itoa(connString.Port), "The server address in the format of host:port")

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		return ServiceDiscoveryClient{}
	}
	return ServiceDiscoveryClient{Conn: conn}

}

func (sdc *ServiceDiscoveryClient) RegisterClient(registerData *pb.RegisterData, port int) (string, error) {

	defer sdc.Conn.Close()
	initClient := pb.NewServiceDiscoveryInitClient(sdc.Conn)
	y, err := initClient.RegisterService(context.Background(), registerData)
	if err != nil {

		return y.Data, nil

	}
	return "", err
}

func (sdc *ServiceDiscoveryClient) UpdateClient(rd *pb.RegisterData) {
	initClient := pb.NewServiceDiscoveryInitClient(sdc.Conn)
	defer initClient.DeleteService(context.Background(), &pb.ServiceGuid{Guid: rd.Guid})
	for {
		rd.Lastupdate = timestamppb.Now()
		_, err := initClient.UpdateServiceHealth(context.Background(), rd)

		if err != nil {
			fmt.Println(err)
		}
		log.Println("updating service")
		time.Sleep(10 * time.Second)

	}

}

func (sdc *ServiceDiscoveryClient) DeleteClient(rd *pb.RegisterData) {
	initClient := pb.NewServiceDiscoveryInitClient(sdc.Conn)
	initClient.DeleteService(context.Background(), &pb.ServiceGuid{Guid: rd.Guid})
}
