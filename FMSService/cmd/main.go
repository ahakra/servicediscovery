package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/ahakra/servicediscovery/pkg/config"
	helper "github.com/ahakra/servicediscovery/pkg/helpers"
	pb "github.com/ahakra/servicediscovery/pkg/serviceDiscoveryProto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//import dynamicbuilder "github.com/ahakra/servicediscovery/FMSService/internal/builder/dynamic"

//"fmt"

const rootDir = "C:\\Users\\ahmad.akra\\Desktop\\goTesting\\DB"
const rootDirLinux = "/home/amd/Desktop/ServiceDiscovery/DB"

var testdata = []string{
	"aa,203-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
	"aa,2023-08-23 10:33,9617055",
	"bb,2023-08-25 12:21,9617055",
}
var wg sync.WaitGroup

func main() {

	// builder := dynamicbuilder.FilePath{}

	// builder.SetRecordType("recordTypeTest").
	// 	SetROOTDir(rootDir).
	// 	SetMaxDynamicRecordTypeDirs(20).
	// 	SetMaxFilesPerFolder(20).
	// 	SetFileName("test2.txt").
	// 	SetData("This is the content of the file.\nHello, World!").
	// 	Build()
	//	fmt.Println(err)

	//init example
	// partitionedBuilder := partitionedbuilder.NewPartitionedBuilder()
	// partitionedBuilder.SetData(testdata).
	// 	SetFileName(uuid.New().String() + ".txt").
	// 	SetRecordType("testRecord").
	// 	SetDateTimeFormat("2006-01-02 15:04").
	// 	SetUnifiedDateFormat("20060102").
	// 	SetSplitter(",").
	// 	SetRootDir(rootDirLinux).
	// 	SetStoreType("raw").
	// 	SetDateTimFieldLocation(1).Build()
	channelData := helper.ChannelData{
		OnInitChan:      make(chan bool),
		OnRegisterChan:  make(chan bool),
		RetunedGuidChan: make(chan string),
		CancelChan:      make(chan os.Signal, 1),
	}

	signal.Notify(channelData.CancelChan, syscall.SIGINT, syscall.SIGTERM)

	conf := config.NewFromJson("config.json")
	log.Println(conf)
	//Register Section for client
	serverAddr := flag.String("addr", conf.Servicediscvoreyserver.Address+":"+strconv.Itoa(conf.Servicediscvoreyserver.Port), "The server address in the format of host:port")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	log.Println(*serverAddr)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	registerData := &pb.RegisterData{
		Servicename:    conf.FMSservice.Name,
		Serviceaddress: conf.FMSservice.Address + ":" + strconv.Itoa(conf.FMSservice.StartingPort),
		Lastupdate:     timestamppb.Now(),
		Messages:       conf.FMSservice.Messages,
	}

	fmsServiceHelper := helper.HelperData{
		Connection:   conn,
		RegisterData: registerData,
		Conf:         *conf,
		Name:         conf.FMSservice.Name,
	}

	defer conn.Close()

	ctx := context.Background()

	go fmsServiceHelper.RegisterService(ctx, channelData)
	go fmsServiceHelper.UpdateServiceHealth(ctx, channelData)
	channelData.OnInitChan <- true
	fmsServiceHelper.DeleteService(ctx, channelData)

}
