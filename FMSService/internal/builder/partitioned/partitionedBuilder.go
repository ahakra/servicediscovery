package partitionedbuilder

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	FMSProto "github.com/ahakra/servicediscovery/FMSService/internal/FMSProto"
	"github.com/google/uuid"
)

//var wg sync.WaitGroup

// type Records struct {
// 	RecordType            string
// 	fileName              string
// 	RootDir               string
// 	data                  []string
// 	splitter              string
// 	dateTimeFieldLocation int
// 	dateTimeFormat        string
// 	dataSlice             map[string][]string
// 	unifiedDateFormat     string
// }

type PartitionedBuilder struct {
	settings *FMSProto.Records
}

func NewPartitionedBuilder() *PartitionedBuilder {
	return &PartitionedBuilder{
		settings: &FMSProto.Records{},
	}
}

func (builder *PartitionedBuilder) SetRootDir(RootDir string) *PartitionedBuilder {
	builder.settings.RootDir = RootDir
	return builder
}
func (builder *PartitionedBuilder) SetRecordType(RecordType string) *PartitionedBuilder {
	builder.settings.RecordType = RecordType
	return builder
}

func (builder *PartitionedBuilder) SetFileName(fileName string) *PartitionedBuilder {
	builder.settings.FileName = fileName
	return builder
}

func (builder *PartitionedBuilder) SetData(data []string) *PartitionedBuilder {
	builder.settings.Data = data
	return builder
}

func (builder *PartitionedBuilder) SetSplitter(splitter string) *PartitionedBuilder {
	builder.settings.Splitter = splitter
	return builder
}

func (builder *PartitionedBuilder) SetDateTimFieldLocation(location int32) *PartitionedBuilder {
	builder.settings.DateTimeFieldLocation = location
	return builder
}
func (builder *PartitionedBuilder) SetDateTimeFormat(dateTimeFormat string) *PartitionedBuilder {
	builder.settings.DateTimeFormat = dateTimeFormat
	return builder
}

func (builder *PartitionedBuilder) SetUnifiedDateFormat(unifiedDateFormat string) *PartitionedBuilder {
	builder.settings.UnifiedDateFormat = unifiedDateFormat
	return builder
}

func (builder *PartitionedBuilder) SetStoreType(storeType string) *PartitionedBuilder {
	builder.settings.StoreType = storeType
	return builder
}
func (ffp *PartitionedBuilder) Build() error {
	_, err := os.Stat(ffp.settings.RootDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(ffp.settings.RootDir, os.ModeDir|os.ModePerm)
		if err != nil {

			return fmt.Errorf("unable to init RootDir: %v", err)
		}
	}
	return ffp.GetOrCreateRecordTypeDir()

}

func (ffp *PartitionedBuilder) GetOrCreateRecordTypeDir() error {

	RecordTypeDir := path.Join(ffp.settings.RootDir, ffp.settings.RecordType)
	_, err := os.Stat(RecordTypeDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(RecordTypeDir, os.ModeDir|os.ModePerm)
		if err != nil {
			return fmt.Errorf("unable to init RootDir: %v", err)
		}
	}

	return ffp.SplitRecordsByDate()
}

func (ffp *PartitionedBuilder) SplitRecordsByDate() error {
	//result := make(map[string][]string)

	for _, record := range ffp.settings.Data {

		fields := strings.Split(record, ffp.settings.Splitter)

		if len(fields) > int(ffp.settings.DateTimeFieldLocation) {
			// Parse the date
			dateStr := fields[ffp.settings.DateTimeFieldLocation]
			parsedDate, err := time.Parse(ffp.settings.DateTimeFormat, dateStr)
			if err == nil {

				unifiedDateFormat := parsedDate.Format(ffp.settings.UnifiedDateFormat)
				keyValue := &FMSProto.KeyValue{
					Key:    unifiedDateFormat,
					Values: record,
				}

				ffp.settings.DataSlice = append(ffp.settings.DataSlice, keyValue)
			} else {
				keyValue := &FMSProto.KeyValue{
					Key:    "invalid",
					Values: record,
				}
				ffp.settings.DataSlice = append(ffp.settings.DataSlice, keyValue)
			}
		} else {
			if len(record) != 0 {
				keyValue := &FMSProto.KeyValue{
					Key:    "invalid",
					Values: record,
				}
				ffp.settings.DataSlice = append(ffp.settings.DataSlice, keyValue)
			}
		}
	}

	return ffp.SaveData()
}

func (ffp *PartitionedBuilder) SaveData() error {

	// Create a map to store data for each key
	dataByKeys := make(map[string][]string)

	// Iterate through records.DataSlice
	for _, msgValues := range ffp.settings.DataSlice {
		dataByKeys[msgValues.Key] = append(dataByKeys[msgValues.Key], msgValues.Values)
	}

	for key, value := range dataByKeys {

		fileDir := path.Join(ffp.settings.RootDir, ffp.settings.RecordType, ffp.settings.StoreType, key)

		finalFile := path.Join(fileDir, ffp.settings.FileName)

		err := os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
		}

		_, err = os.Stat(finalFile)
		if !os.IsNotExist(err) {
			finalFile = path.Join(fileDir, ffp.settings.FileName+"_"+uuid.New().String())
		}

		content := []byte(strings.Join(value, "\n"))

		err = ioutil.WriteFile(finalFile, content, 0644)
		if err != nil {

			return err
		}

		fmt.Println("File saved successfully at:", finalFile)

	}

	return nil
}
