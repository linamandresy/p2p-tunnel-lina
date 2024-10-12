package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig `yaml:server`
}

type ServerConfig struct {
	Host string
	Port int
}

func LoadData(path string) (Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, nil
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	return config, err
}

func InitConfig() (Config, error) {
	return LoadData("config.yaml")

}
