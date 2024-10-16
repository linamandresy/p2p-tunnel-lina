package main

import (
	"p2p-tunnel-lina/client/config"
	"p2p-tunnel-lina/client/services/darwin"
)

func main() {
	var err error
	config.CONFIG, err = config.Detect()
	if err != nil {
		panic(err)
	}

	StartServices(config.CONFIG)
}

func StartServices(config config.Config) {
	switch config.OSType {
	case "darwin":
		darwin.StartService(config)
	default:
		panic("unsupported operating system")
	}
}
