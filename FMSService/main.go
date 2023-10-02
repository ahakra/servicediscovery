package main

import "github.com/ahakra/servicediscovery/FMSService/internal/builder"

const rootDir = "C:\\Users\\ahmad.akra\\Desktop\\goTesting\\DB"
const rootDirLinux = "/home/amd/Desktop/ServiceDiscovery/DB"

func main() {
	for i := 0; i < 100000; i++ {
		builder := builder.FilePath{}

		builder.SetRecordType("recordTypeTest").
			SetROOTDir(rootDir).
			SetMaxDynamicRecordTypeDirs(20).
			SetMaxFilesPerFolder(20).
			SetFileName("test2.txt").
			SetData("This is the content of the file.\nHello, World!").
			Build()
		//	fmt.Println(err)
	}

}
