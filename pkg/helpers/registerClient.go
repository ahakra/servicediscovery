package helper

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ahakra/servicediscovery/pkg/config"
	"github.com/ahakra/servicediscovery/pkg/serviceDiscoveryProto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type HelperData struct {
	Connection   *grpc.ClientConn
	Conf         config.Config
	RegisterData *serviceDiscoveryProto.RegisterData
}

func New(hd *HelperData) *HelperData {
	return &HelperData{
		Connection:   hd.Connection,
		Conf:         hd.Conf,
		RegisterData: hd.RegisterData}
}

func (hd *HelperData) RegisterService(ctx context.Context, GuidChan chan string) {

	initClient := serviceDiscoveryProto.NewServiceDiscoveryInitClient(hd.Connection)
	for {
		y, err := initClient.RegisterService(ctx, hd.RegisterData)
		if err != nil {

			fmt.Println(err)

		} else {
			GuidChan <- y.Data
			break

		}
		time.Sleep(10 * time.Second)
	}
}

func (hd *HelperData) UpdateServiceHealth(ctx context.Context, GuidChan chan string) {
	retunedGuid := <-GuidChan
	GuidChan <- retunedGuid
	initClient := serviceDiscoveryProto.NewServiceDiscoveryInitClient(hd.Connection)
	for {
		hd.RegisterData.Lastupdate = timestamppb.Now()

		_, err := initClient.UpdateServiceHealth(context.Background(), hd.RegisterData)

		if err != nil {
			fmt.Println(err)
		}
		log.Println("updating " + hd.Conf.Loggerservicewriter.Name + " service")
		time.Sleep(10 * time.Second)

	}
}

func (hd *HelperData) DeleteService(ctx context.Context, GuidChan chan string, sigChan chan os.Signal) {
	initClient := serviceDiscoveryProto.NewServiceDiscoveryInitClient(hd.Connection)
	returnedguid := <-GuidChan
	sig := <-sigChan
	fmt.Printf("Received signal: %v\n", sig)

	initClient.DeleteService(context.Background(), &serviceDiscoveryProto.ServiceGuid{Guid: returnedguid})

	os.Exit(1)
}
