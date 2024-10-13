package config

import (
	"io/ioutil"
	"runtime"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server ServerConfig
	OSType string
}

type ServerConfig struct {
	Host string
	Port int
}

func Detect(path ...string) (Config, error) {

	if len(path) == 0 {
		path = append(path, "config.yaml")
	}
	data, err := ioutil.ReadFile(path[0])
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	DetectOs(&config)
	return config, err
}

func DetectOs(config *Config) {
	*&config.OSType = runtime.GOOS
}
