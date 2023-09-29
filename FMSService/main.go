package main

import "github.com/ahakra/servicediscovery/FMSService/internal/builder"

const rootDir = "C:\\Users\\ahmad.akra\\Desktop\\goTesting\\DB"
const rootDirLinux = "/home/amd/Desktop/ServiceDiscovery/DB"

func main() {
	for i := 0; i < 100000; i++ {
		builder := builder.FilePath{}
		builder.SetRecordType("recordTypeTest2").
			SetROOTDir(rootDir).
			SetMaxDynamicRecordTypeDirs(1000).
			SetMaxFilesPerFolder(1000).
			SetFileName("test2.txt").
			Build()
		//	fmt.Println(err)
	}

}
