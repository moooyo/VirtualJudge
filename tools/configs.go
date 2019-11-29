package tools

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type GlobalConfig struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

var configPath = "config.yaml"

var config *GlobalConfig = nil

func initConfig(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	buffer, _ := ioutil.ReadAll(f)
	config = new(GlobalConfig)
	err = yaml.Unmarshal(buffer, config)
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfig() *GlobalConfig {
	if config == nil {
		initConfig(configPath)
	}
	return config
}
