package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var ErrEnvVarEmpty = errors.New("getenv: environment variable empty")

type MongoDatabase struct {
	URL string `json:"URL"`
}

type ServiceDiscoveryServer struct {
	Port       int    `json:"Port"`
	Address    string `json:"Address"`
	Database   string `json:"Database"`
	Collection string `json:"Collection"`
}

type LoggerServiceReader struct {
	Name         string   `json:"Name"`
	Address      string   `json:"Address"`
	StartingPort int      `json:"StartingPort"`
	Database     string   `json:"Database"`
	Messages     []string `json:"Messages"`
}

type LoggerServiceWriter struct {
	Name         string   `json:"Name"`
	Address      string   `json:"Address"`
	StartingPort int      `json:"StartingPort"`
	Database     string   `json:"Database"`
	Messages     []string `json:"Messages"`
}

type Config struct {
	Mongodatabase          MongoDatabase          `json:"MongoDatabase"`
	Servicediscvoreyserver ServiceDiscoveryServer `json:"ServiceDiscoveryServer"`
	Loggerservicereader    LoggerServiceReader    `json:"LoggerServiceReader"`
	Loggerservicewriter    LoggerServiceWriter    `json:"LoggerServiceWriter"`
}

// type MongoDatabase struct {
// 	URL string
// }
// type ServiceDiscvoreyServer struct {
// 	Address    string
// 	Port       int
// 	Database   string
// 	Collection string
// }

// type LoggerServiceReader struct {
// 	Name         string
// 	Address      string
// 	StartingPort int
// 	Database     string
// }

// type LoggerServiceWriter struct {
// 	Name         string
// 	Address      string
// 	StartingPort int
// 	Database     string
// }

// func LoadEnv(fileName string) {
// 	re := regexp.MustCompile(`^(.*` + "servicediscovery" + `)`)
// 	cwd, _ := os.Getwd()
// 	rootPath := re.Find([]byte(cwd))

// 	err := godotenv.Load(string(rootPath) + `/` + fileName)
// 	if err != nil {
// 		godotenv.Load()
// 	}
// }

func NewFromJson(fileName string) *Config {

	rootPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

	}

	fmt.Printf("Working Directory: %s\n", rootPath)

	jsonData, err := ioutil.ReadFile(string(rootPath) + `/` + fileName)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Create a Config struct to unmarshal JSON into
	var config Config

	// Unmarshal JSON data into the struct
	if err := json.Unmarshal(jsonData, &config); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}
	return &config
}

// func New() *Config {
// 	serviceserverport, err := getenvInt("SERVICE_DISCOVERY_SERVER_PORT")

// 	if err != nil {
// 		return &Config{}
// 	}
// 	loggerservicereaderstartingport, err := getenvInt("LOGGER_SERVICE_READER_STARTINGPORT")

// 	if err != nil {
// 		return &Config{}
// 	}
// 	loggerservicewriterstartingport, err := getenvInt("LOGGER_SERVICE_WRITERSTARTINGPORT")

// 	if err != nil {
// 		return &Config{}
// 	}

// 	return &Config{

// 		Mongodatabase: MongoDatabase{
// 			URL: os.Getenv("MONGO_DATABASE_URL"),
// 		},
// 		Servicediscvoreyserver: ServiceDiscoveryServer{
// 			Database:   os.Getenv("SERVICE_DISCOVERY_SERVER_DATABASE"),
// 			Collection: os.Getenv("SERVICE_DISCOVERY_SERVER_COllECTION"),
// 			Port:       serviceserverport,
// 			Address:    os.Getenv("SERVICE_DISCOVERY_SERVER_ADDRESS"),
// 		},
// 		Loggerservicereader: LoggerServiceReader{
// 			Name:         os.Getenv("LOGGER_DISCOVERY_READER_NAME"),
// 			Address:      os.Getenv("LOGGER_DISCOVERY_READER_ADDRESS"),
// 			Database:     os.Getenv("LOGGER_DISCOVERY_READER_DATABASE"),
// 			StartingPort: loggerservicereaderstartingport,
// 		},
// 		Loggerservicewriter: LoggerServiceWriter{
// 			Name:         os.Getenv("LOGGER_DISCOVERY_WRITER_NAME"),
// 			Address:      os.Getenv("LOGGER_DISCOVERY_WRITER_ADDRESS"),
// 			Database:     os.Getenv("LOGGER_DISCOVERY_WRITER_DATABASE"),
// 			StartingPort: loggerservicewriterstartingport,
// 		},
// 	}
// }

// func getenvInt(key string) (int, error) {
// 	s, err := getenvStr(key)
// 	if err != nil {
// 		return 0, err
// 	}
// 	v, err := strconv.Atoi(s)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return v, nil
// }
// func getenvStr(key string) (string, error) {
// 	v := os.Getenv(key)
// 	if v == "" {
// 		return v, ErrEnvVarEmpty
// 	}
// 	return v, nil
// }
