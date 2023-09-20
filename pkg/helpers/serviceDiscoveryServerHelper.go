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

type ChannelData struct {
	OnInitChan      chan bool
	OnRegisterChan  chan bool
	RetunedGuidChan chan string
	CancelChan      chan os.Signal
}
type HelperData struct {
	Connection   *grpc.ClientConn
	Conf         config.Config
	RegisterData *serviceDiscoveryProto.RegisterData
	Name         string
}

// func New(hd *HelperData) *HelperData {
// 	return &HelperData{
// 		Connection:   hd.Connection,
// 		Conf:         hd.Conf,
// 		RegisterData: hd.RegisterData}
// }

func (hd HelperData) RegisterService(ctx context.Context, cd ChannelData) {
	<-cd.OnInitChan
	initClient := serviceDiscoveryProto.NewServiceDiscoveryInitClient(hd.Connection)
	for {
		y, err := initClient.RegisterService(ctx, hd.RegisterData)
		if err != nil {

			fmt.Println(err)

		} else {

			cd.RetunedGuidChan <- y.Data
			cd.OnRegisterChan <- true
			break

		}
		time.Sleep(10 * time.Second)
	}
}

func (hd HelperData) UpdateServiceHealth(ctx context.Context, cd ChannelData) {
	<-cd.OnRegisterChan
	initClient := serviceDiscoveryProto.NewServiceDiscoveryInitClient(hd.Connection)
	for {
		hd.RegisterData.Lastupdate = timestamppb.Now()

		_, err := initClient.UpdateServiceHealth(context.Background(), hd.RegisterData)

		if err != nil {
			fmt.Println(err)
		}
		log.Println("updating " + hd.Name + " service")
		time.Sleep(10 * time.Second)

	}
}

func (hd HelperData) DeleteService(ctx context.Context, cd ChannelData) {
	initClient := serviceDiscoveryProto.NewServiceDiscoveryInitClient(hd.Connection)
	returnedguid := <-cd.RetunedGuidChan

	sig := <-cd.CancelChan
	fmt.Printf("Received signal: %v\n", sig)

	initClient.DeleteService(context.Background(), &serviceDiscoveryProto.ServiceGuid{Guid: returnedguid})
	hd.Connection.Close()
	os.Exit(1)
}
