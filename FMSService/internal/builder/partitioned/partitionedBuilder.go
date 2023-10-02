package partitionedbuilder

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

var wg sync.WaitGroup

type Records struct {
	recordType            string
	fileName              string
	rootDir               string
	data                  string
	splitter              string
	dateTimeFieldLocation int
	dateTimeFormat        string
	dataSlice             map[string][]string
	unifiedDateFormat     string
}

func (builder *Records) SetROOTDir(rootDir string) *Records {
	builder.rootDir = rootDir
	return builder
}
func (builder *Records) SetRecordType(recordType string) *Records {
	builder.recordType = recordType
	return builder
}

func (builder *Records) SetFileName(fileName string) *Records {
	builder.fileName = fileName
	return builder
}

func (builder *Records) SetData(data string) *Records {
	builder.data = data
	return builder
}

func (builder *Records) SetSplitter(splitter string) *Records {
	builder.splitter = splitter
	return builder
}

func (builder *Records) SetDateTimFieldLocation(location int) *Records {
	builder.dateTimeFieldLocation = location
	return builder
}
func (builder *Records) SetDateTimeFormat(dateTimeFormat string) *Records {
	builder.dateTimeFormat = dateTimeFormat
	return builder
}

func (builder *Records) SetUnifiedDateFormat(unifiedDateFormat string) *Records {
	builder.unifiedDateFormat = unifiedDateFormat
	return builder
}

func (ffp *Records) Build() error {
	_, err := os.Stat(ffp.rootDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(ffp.rootDir, os.ModeDir|os.ModePerm)
		if err != nil {

			return fmt.Errorf("unable to init rootDir: %v", err)
		}
	}
	return ffp.GetOrCreateRecordTypeDir()

}

func (ffp *Records) GetOrCreateRecordTypeDir() error {

	recordTypeDir := path.Join(ffp.rootDir, ffp.recordType)
	_, err := os.Stat(recordTypeDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(recordTypeDir, os.ModeDir|os.ModePerm)
		if err != nil {
			return fmt.Errorf("unable to init rootDir: %v", err)
		}
	}

	return ffp.SplitRecordsByDate()
}

func (ffp *Records) SplitRecordsByDate() error {
	//result := make(map[string][]string)
	ffp.dataSlice = make(map[string][]string)
	records := strings.Split(ffp.data, "\n")

	for _, record := range records {

		fields := strings.Split(record, ffp.splitter)

		if len(fields) > ffp.dateTimeFieldLocation {
			// Parse the date
			dateStr := fields[ffp.dateTimeFieldLocation]
			parsedDate, err := time.Parse(ffp.dateTimeFormat, dateStr)
			if err == nil {

				unifiedDateFormat := parsedDate.Format(ffp.unifiedDateFormat)

				ffp.dataSlice[unifiedDateFormat] = append(ffp.dataSlice[unifiedDateFormat], record)
			} else {

				ffp.dataSlice["invalid"] = append(ffp.dataSlice["invalid"], record)
			}
		} else {
			if len(record) != 0 {
				ffp.dataSlice["invalid"] = append(ffp.dataSlice["invalid"], record)
			}
		}
	}

	return ffp.SaveData()
}

func (ffp *Records) SaveData() error {

	for key, value := range ffp.dataSlice {
		fmt.Println("KEYS: ", key)
		fileDir := path.Join(ffp.rootDir, ffp.recordType, key)

		finalFile := path.Join(fileDir, ffp.fileName)

		err := os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
		} else {
			fmt.Printf("Directory created successfully: %s\n", key)
		}

		_, err = os.Stat(finalFile)
		if !os.IsNotExist(err) {
			finalFile = path.Join(fileDir, ffp.fileName+"_"+uuid.New().String())
		}

		content := []byte(strings.Join(value, "\n"))

		err = ioutil.WriteFile(finalFile, content, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return err
		}

		fmt.Println("File saved successfully at:", finalFile)

	}

	return nil
}
