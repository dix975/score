package configuration

import (
	"dix975.com/database"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var (
	config *Configuration
)

type Configuration struct {
	SchemaFolder  string `json:"schemaFolder"`
	MongoDBConfig db.MongoServerConfig
}

func Reset(){
	config = nil
}

func Config() *Configuration{
	if config == nil {
		load()
	}

	return config
}

func load() error{

	configFile := os.Getenv("CONFIG_FILE")

	if configFile == "" {
		configFile = "configuration/configuration-local.json"

		fmt.Println("Using default config file please set env var [CONFIG_FILE=]")
	}

	fmt.Printf("Configuration file : %v\n", configFile)

	path, _ := filepath.Abs(configFile)
	file, fileErr := os.Open(path)
	if fileErr != nil {
		return fileErr
	}

	decoder := json.NewDecoder(file)

	err := decoder.Decode(&config)
	if err != nil {
		return err
	}
}
