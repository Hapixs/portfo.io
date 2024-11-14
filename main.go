package main

import (
	"encoding/json"
	"os"

	"github.com/Hapixs/portfolio/entities"
	"github.com/Hapixs/portfolio/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := LoadOrCreateConfig("config.json")
	if err != nil {
		println("An error occured with the config.json")
		os.Exit(1)
		return
	}

	r := gin.Default()

	r = handlers.SetupHandlers(r)
	entities.SetupMySql(config)

	r.Run(string(config.ListenIp + ":" + config.ListenPort))
}

func CreateDefaultConfig() *entities.Config {
	println("Creating default configuration file you must edit this file before continuing")
	config := &entities.Config{
		ListenPort: "8080",
		ListenIp:   "0.0.0.0",
		Database: struct {
			Host     string "json:\"host\""
			Port     int    "json:\"port\""
			User     string "json:\"user\""
			Password string "json:\"password\""
			Name     string "json:\"name\""
		}{
			Host:     "localhost",
			Port:     3307,
			User:     "user",
			Password: "password",
			Name:     "database",
		},
		DefaultApiKey: "A",
	}

	return config
}
func LoadOrCreateConfig(filename string) (*entities.Config, error) {
	println("Looking for configuration file " + filename)
	file, err := os.Open(filename)
	if err != nil {
		err = WriteDefaultConfig(filename)
		if err != nil {
			return nil, err
		}
		return LoadOrCreateConfig(filename)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &entities.Config{}
	err = decoder.Decode(config)
	return config, err
}
func WriteDefaultConfig(filename string) error {
	config := CreateDefaultConfig()
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(config)
	if err != nil {
		return err
	}

	return nil
}
