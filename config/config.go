package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	config *Config
)

type Config struct {
	AppName  string
	MongoURL string
	Port     string
}

//Init Initializes the configuration struct
func Init() {
	//Read in configuration from a file
	configFile, err := ioutil.ReadFile("../config.json")
	if err != nil {
		log.Panic("There was an error loading the config file")
	}

	if err := json.Unmarshal(configFile, config); err != nil {
		log.Panic("Error decoding into struct. Please check the config file")
	}
}

//Get Returns a config file
func Get() *Config {
	return config
}
