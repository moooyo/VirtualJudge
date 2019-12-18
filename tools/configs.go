package tools

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type GlobalConfig struct {
	WebConfig   Web   `yaml:"web"`
	MysqlConfig Mysql `yaml:"mysql"`
}

type Web struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type Mysql struct {
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"database"`
	User     string `yaml:"username"`
	Password string `yaml:"password"`
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
