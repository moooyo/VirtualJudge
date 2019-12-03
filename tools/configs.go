package tools

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type GlobalConfig struct {
	Address     string `yaml:"address"`
	Port        int    `yaml:"port"`
	MysqlConfig MysqlConfig
}
type MysqlConfig struct {
	Address  string `yaml:"mysqlAddress"`
	Port     string `yaml:"mysqlPort"`
	DB       string `yaml:"mysqlDbName"`
	User     string `yaml:"mysqlUser"`
	Password string `yaml:"mysqlPassword"`
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
