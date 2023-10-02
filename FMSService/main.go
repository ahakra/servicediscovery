package main

//import dynamicbuilder "github.com/ahakra/servicediscovery/FMSService/internal/builder/dynamic"
import (
	//"fmt"
	"sync"

	partitionedbuilder "github.com/ahakra/servicediscovery/FMSService/internal/builder/partitioned"
	"github.com/google/uuid"
)

const rootDir = "C:\\Users\\ahmad.akra\\Desktop\\goTesting\\DB"
const rootDirLinux = "/home/amd/Desktop/ServiceDiscovery/DB"
const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		// builder := dynamicbuilder.FilePath{}

		// builder.SetRecordType("recordTypeTest").
		// 	SetROOTDir(rootDir).
		// 	SetMaxDynamicRecordTypeDirs(20).
		// 	SetMaxFilesPerFolder(20).
		// 	SetFileName("test2.txt").
		// 	SetData("This is the content of the file.\nHello, World!").
		// 	Build()
		//	fmt.Println(err)
		go func() {
			defer wg.Done()

			partitionedBuilder := partitionedbuilder.Records{}
			partitionedBuilder.SetData("aa,203-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\n aa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\naa,2023-08-23 10:33,9617055\nbb,2023-08-25 12:21,9617055\n").
				SetFileName(uuid.New().String() + ".txt").
				SetRecordType("testRecord").
				SetDateTimeFormat("2006-01-02 15:04").
				SetUnifiedDateFormat("20060102").
				SetSplitter(",").
				SetROOTDir(rootDir).
				SetDateTimFieldLocation(1).Build()

		}()
	}

	wg.Wait()
}
