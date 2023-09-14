package config

import (
	"errors"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
)

var ErrEnvVarEmpty = errors.New("getenv: environment variable empty")

type MongoDatabase struct {
	URL string
}
type ServiceDiscvoreyServer struct {
	Address    string
	Port       int
	Database   string
	Collection string
}

type LoggerServiceReader struct {
	Address      string
	StartingPort int
	Database     string
}

type LoggerServiceWriter struct {
	Address      string
	StartingPort int
	Database     string
}

type Config struct {
	Mongodatabase          MongoDatabase
	Servicediscvoreyserver ServiceDiscvoreyServer
	Loggerservicereader    LoggerServiceReader
	Loggerservicewriter    LoggerServiceWriter
}

func LoadEnv(fileName string) {
	re := regexp.MustCompile(`^(.*` + "servicediscovery" + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/` + fileName)
	if err != nil {
		godotenv.Load()
	}
}

func New() *Config {
	serviceserverport, err := getenvInt("SERVICE_DISCOVERY_SERVER_PORT")

	if err != nil {
		return &Config{}
	}
	loggerservicereaderstartingport, err := getenvInt("LOGGER_SERVICE_READER_STARTINGPORT")

	if err != nil {
		return &Config{}
	}
	loggerservicewriterstartingport, err := getenvInt("LOGGER_SERVICE_WRITERSTARTINGPORT")

	if err != nil {
		return &Config{}
	}

	return &Config{

		Mongodatabase: MongoDatabase{
			URL: os.Getenv("MONGO_DATABASE_URL"),
		},
		Servicediscvoreyserver: ServiceDiscvoreyServer{
			Database:   os.Getenv("SERVICE_DISCOVERY_SERVER_DATABASE"),
			Collection: os.Getenv("SERVICE_DISCOVERY_SERVER_COllECTION"),
			Port:       serviceserverport,
			Address:    os.Getenv("SERVICE_DISCOVERY_SERVER_ADDRESS"),
		},
		Loggerservicereader: LoggerServiceReader{
			Address:      os.Getenv("LOGGER_DISCOVERY_READER_ADDRESS"),
			Database:     os.Getenv("LOGGER_DISCOVERY_READER_DATABASE"),
			StartingPort: loggerservicereaderstartingport,
		},
		Loggerservicewriter: LoggerServiceWriter{
			Address:      os.Getenv("LOGGER_DISCOVERY_WRITER_ADDRESS"),
			Database:     os.Getenv("LOGGER_DISCOVERY_WRITER_DATABASE"),
			StartingPort: loggerservicewriterstartingport,
		},
	}
}

func getenvInt(key string) (int, error) {
	s, err := getenvStr(key)
	if err != nil {
		return 0, err
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return v, nil
}
func getenvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, ErrEnvVarEmpty
	}
	return v, nil
}
