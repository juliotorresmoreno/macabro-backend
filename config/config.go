package config

import (
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

// Config s
type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database struct {
		Driver string `json:"driver"`
		DSN    string `json:"dsn"`
	} `json:"database"`
}

var conf Config

// Init s
func Init(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.NewDecoder(f).Decode(&conf)
	if err != nil {
		log.Fatal(err)
	}
}

// GetConfig s
func GetConfig() Config {
	return conf
}
