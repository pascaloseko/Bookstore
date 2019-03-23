package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// configuration details for the server and database

//RestfulEPDefault ...
var RestfulEPDefault = "localhost:8181"

//ServiceConfig ...
type ServiceConfig struct {
	RestfulEndpoint string `json:"restfulapi_endpoint"`
	DbAddress       string `json:"dbaddress"`
	DbPort          string `json:"dbport"`
	DbUser          string `json:"dbuser"`
	DbName          string `json:"dbname"`
	DbPassword      string `json:"dbpass"`
}

// Config ...
var Config ServiceConfig

//ExtractConfiguration extracts configuration details
func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{}

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}

	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}
