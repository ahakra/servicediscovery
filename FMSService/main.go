package main

//import dynamicbuilder "github.com/ahakra/servicediscovery/FMSService/internal/builder/dynamic"
import (
	//"fmt"

	partitionedbuilder "github.com/ahakra/servicediscovery/FMSService/internal/builder/partitioned"
	"github.com/google/uuid"
)

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

	partitionedBuilder := partitionedbuilder.NewPartitionedBuilder()
	partitionedBuilder.SetData(testdata).
		SetFileName(uuid.New().String() + ".txt").
		SetRecordType("testRecord").
		SetDateTimeFormat("2006-01-02 15:04").
		SetUnifiedDateFormat("20060102").
		SetSplitter(",").
		SetRootDir(rootDirLinux).
		SetStoreType("raw").
		SetDateTimFieldLocation(1).Build()

}
