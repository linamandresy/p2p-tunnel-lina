package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var CONFIG Config

type Config struct {
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
	return config, err
}
