package main

import (
	"client/config"
	"client/services/darwin"
)

func main() {

	config, err := config.Detect()
	if err != nil {
		panic(err)
	}

	StartServices(config)
}

func StartServices(config config.Config) {
	switch config.OSType {
	case "darwin":
		darwin.StartService(config)
	default:
		panic("unsupported operating system")
	}
}
