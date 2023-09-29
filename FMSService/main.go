package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
)

const rootDir = "C:\\Users\\ahmad.akra\\Desktop\\goTesting\\DB"
const rootDirLinux = "/home/amd/Desktop/ServiceDiscovery/DB"

type FilePath struct {
	recordType               string
	fileName                 string
	rootDir                  string
	dynamicDirs              []string
	recordTypeDirs           []string
	dynamicRecordTypeDirs    []string
	fallBackDir              string
	filePath                 string
	fallBack                 bool
	maxDynamicRecordTypeDirs int
	maxFilesPerFolder        int
}

var errCannotCreateRootdir = errors.New("cann't create root dir: ")

func (builder *FilePath) SetRecordType(recordType string) *FilePath {
	builder.recordType = recordType
	return builder
}

func (builder *FilePath) SetROOTDir(rootDir string) *FilePath {
	builder.rootDir = rootDir
	return builder
}
func (builder *FilePath) SetFileName(fileName string) *FilePath {
	builder.fileName = fileName
	return builder
}

func (builder *FilePath) SetMaxDynamicRecordTypeDirs(max int) *FilePath {
	builder.maxDynamicRecordTypeDirs = max
	return builder
}

func (builder *FilePath) SetMaxFilesPerFolder(max int) *FilePath {
	builder.maxFilesPerFolder = max
	return builder
}
func (ffp *FilePath) GetOrCreateRootDIR() error {
	_, err := os.Stat(ffp.rootDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(ffp.rootDir, os.ModeDir|os.ModePerm)
		if err != nil {
			return errCannotCreateRootdir
		}
	}
	return ffp.GetOrCreateDynamicFolders()

}

func (ffp *FilePath) CreateDynamicDirFallBack() error {
	if ffp.fallBack {
		ffp.dynamicDirs = []string{}
		dynamicFolders, err := ioutil.ReadDir(ffp.rootDir)
		if err != nil {
			return err
		}

		if len(dynamicFolders) > 0 {
			for _, dB := range dynamicFolders {
				dyanmicDir := path.Join(ffp.rootDir, dB.Name())
				ffp.dynamicDirs = append(ffp.dynamicDirs, dyanmicDir)
			}

		}
		recordTypeDoesNotExists := false
		if len(ffp.dynamicDirs) > 0 {
			for _, dD := range ffp.dynamicDirs {
				recordTypePath := path.Join(dD, ffp.recordType)
				_, err := os.Stat(recordTypePath)
				if os.IsNotExist(err) {
					recordTypeDoesNotExists = true
					ffp.fallBackDir = dD
					break
				}
			}
		}
		if !recordTypeDoesNotExists {
			dyanmicDir := path.Join(ffp.rootDir, uuid.New().String())
			err = os.Mkdir(dyanmicDir, os.ModeDir|os.ModePerm)
			if err != nil {
				return err
			}
			ffp.fallBackDir = dyanmicDir
			ffp.dynamicDirs = append(ffp.dynamicDirs, dyanmicDir)

		}

	}
	return ffp.CreateRecordTypeDirFallBack()

}
func (ffp *FilePath) GetOrCreateDynamicFolders() error {

	dynamicFolders, err := ioutil.ReadDir(ffp.rootDir)
	if err != nil {
		return err
	}

	if len(dynamicFolders) > 0 {
		for _, dynamicFolder := range dynamicFolders {
			dyanmicDir := path.Join(ffp.rootDir, dynamicFolder.Name())
			ffp.dynamicDirs = append(ffp.dynamicDirs, dyanmicDir)
		}

	} else {
		dyanmicDir := path.Join(ffp.rootDir, uuid.New().String())
		err = os.Mkdir(dyanmicDir, os.ModeDir|os.ModePerm)
		if err != nil {
			return err
		}
		ffp.dynamicDirs = append(ffp.dynamicDirs, dyanmicDir)
	}
	return ffp.GetOrCreateRecordTypeDir()

}
func (ffp *FilePath) CreateRecordTypeDirFallBack() error {
	if ffp.fallBack {
		recordTypeDir := path.Join(ffp.fallBackDir, ffp.recordType)
		err := os.Mkdir(recordTypeDir, os.ModeDir|os.ModePerm)
		if err != nil {
			return err
		}
		ffp.recordTypeDirs = []string{}
		ffp.recordTypeDirs = append(ffp.recordTypeDirs, recordTypeDir)
		ffp.fallBackDir = recordTypeDir
	}
	return ffp.CreateDynamicRecordTypeDirFallBack()
}
func (ffp *FilePath) GetOrCreateRecordTypeDir() error {

	recordTypeExists := false
	for _, dynamics := range ffp.dynamicDirs {
		recordTypeDir := path.Join(dynamics, ffp.recordType)
		_, err := os.Stat(recordTypeDir)
		if !os.IsNotExist(err) {
			recordTypeExists = true
			ffp.recordTypeDirs = append(ffp.recordTypeDirs, recordTypeDir)
		}
	}

	if !recordTypeExists {
		recorTypePath := path.Join(ffp.dynamicDirs[0], ffp.recordType)
		err := os.Mkdir(recorTypePath, os.ModeDir|os.ModePerm)
		if err != nil {
			return err
		}
		ffp.recordTypeDirs = append(ffp.recordTypeDirs, recorTypePath)

	}
	return ffp.GetOrCreateDynamicRecordTypeDirs()
}
func (ffp *FilePath) CreateDynamicRecordTypeDirFallBack() error {

	if ffp.fallBack {

		dynamicRecordType := false

		for _, dynamicRecordTypeDir := range ffp.recordTypeDirs {
			index := strings.Index(dynamicRecordTypeDir, ffp.recordType)
			currentDyanmicRecordTypeDir := dynamicRecordTypeDir[:index+len(ffp.recordType)]

			entries, err := os.ReadDir(currentDyanmicRecordTypeDir)
			if err != nil {
				return err
			}
			if len(entries) < ffp.maxDynamicRecordTypeDirs {
				dynamicRecordTypePath := path.Join(currentDyanmicRecordTypeDir, uuid.NewString())

				err := os.Mkdir(dynamicRecordTypePath, os.ModeDir|os.ModePerm)
				if err != nil {
					return err
				}

				ffp.fallBackDir = dynamicRecordTypePath
				ffp.dynamicRecordTypeDirs = append(ffp.dynamicRecordTypeDirs, dynamicRecordTypePath)
				dynamicRecordType = true

			}

		}
		if dynamicRecordType {
			return ffp.GetOrCreateFilePathFallBack()

		} else {
			return ffp.CreateDynamicDirFallBack()
		}
	}
	return nil
}

func (ffp *FilePath) GetOrCreateDynamicRecordTypeDirs() error {
	dynamicRecordTypeDirExists := false
	for _, recordTypeDir := range ffp.recordTypeDirs {

		entries, err := os.ReadDir(recordTypeDir)
		if err != nil {
			return err
		}
		if len(entries) > 0 {
			for _, entry := range entries {

				dynamicRecordTypePath := path.Join(recordTypeDir, entry.Name())

				ffp.dynamicRecordTypeDirs = append(ffp.dynamicRecordTypeDirs, dynamicRecordTypePath)
				dynamicRecordTypeDirExists = true

			}

		}

	}
	if !dynamicRecordTypeDirExists {
		dynamicRecordTypePath := path.Join(ffp.recordTypeDirs[0], uuid.NewString())
		err := os.Mkdir(dynamicRecordTypePath, os.ModeDir|os.ModePerm)
		if err != nil {
			return err
		}
		ffp.dynamicRecordTypeDirs = append(ffp.dynamicRecordTypeDirs, dynamicRecordTypePath)

	}
	return ffp.GetOrCreateFilePath()

}
func (ffp *FilePath) GetOrCreateFilePathFallBack() error {
	if ffp.fallBack {
		entries, err := os.ReadDir(ffp.fallBackDir)
		if err != nil {
			return err
		}
		if len(entries) < ffp.maxFilesPerFolder {
			filePathDir := path.Join(ffp.fallBackDir, ffp.fileName)
			err := os.Mkdir(filePathDir, os.ModeDir|os.ModePerm)
			if err != nil {
				return err
			}
			ffp.filePath = filePathDir
			ffp.fallBackDir = filePathDir
			return ffp.SaveData()

		}
		return ffp.CreateDynamicRecordTypeDirFallBack()

	}
	return nil
}
func (ffp *FilePath) GetOrCreateFilePath() error {
	folderToSaveExists := false
	for _, dynamicrecodTypeDir := range ffp.dynamicRecordTypeDirs {
		entries, err := os.ReadDir(dynamicrecodTypeDir)
		if err != nil {
			return err
		}
		if len(entries) < ffp.maxFilesPerFolder && !folderToSaveExists {
			folderToSaveExists = true
			ffp.filePath = dynamicrecodTypeDir
			break
		}

	}
	if folderToSaveExists {
		filePathDir := path.Join(ffp.filePath, ffp.fileName)
		_, err := os.Stat(filePathDir)
		if !os.IsNotExist(err) {
			fileName := ffp.fileName + uuid.NewString()
			filePathDir = path.Join(ffp.filePath, fileName)
		}
		err = os.Mkdir(filePathDir, os.ModeDir|os.ModePerm)
		if err != nil {
			return err
		}
		ffp.filePath = filePathDir
		return ffp.SaveData()

	} else {
		//fmt.Println("no place to save file")
		ffp.fallBack = true

		return ffp.CreateDynamicRecordTypeDirFallBack()

	}

}

func (ffp *FilePath) SaveData() error {
	content := []byte("This is the content of the file.\nHello, World!")
	finalFile := path.Join(ffp.filePath, ffp.fileName)

	err := ioutil.WriteFile(finalFile, content, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Println("File saved successfully at:", finalFile)
	return nil
}
func main() {
	for i := 0; i < 100000; i++ {
		builder := FilePath{}
		builder.SetRecordType("recordTypeTest").
			SetROOTDir(rootDir).
			SetMaxDynamicRecordTypeDirs(1000).
			SetMaxFilesPerFolder(1000).
			SetFileName("test2.txt")

		builder.GetOrCreateRootDIR()
		//	fmt.Println(err)
	}

}
