package config

import (
	"errors"
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
)

var ErrEnvVarEmpty = errors.New("getenv: environment variable empty")

type Mongodatabase struct {
	URL string
}

type Config struct {
	Database                   Mongodatabase
	serviceDiscvoeryServerPort int
	ServiceDiscoveryDatabase   string
	ServiceDiscoveryCollection string
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
	serviceport, err := getenvInt("SERVICE_DISCOVERY_SERVER_PORT")
	if err != nil {
		return &Config{}
	}
	return &Config{
		Database: Mongodatabase{
			URL: os.Getenv("MONGO_DATABASE_URL"),
		},

		serviceDiscvoeryServerPort: serviceport,
		ServiceDiscoveryDatabase:   os.Getenv("SERVICE_DISCOVERY_DATABASE"),
		ServiceDiscoveryCollection: os.Getenv("services"),
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
