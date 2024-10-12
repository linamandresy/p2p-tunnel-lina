package main

import (
	"client/config"
	"fmt"
	"log"
	"os"
)

func main() {

	config, err := config.InitConfig()
	if err != nil {
		log.Fatalln(os.Stderr, err)
	}
	fmt.Printf("Serveur : %s\n", config.Server.Host)
	fmt.Printf("Port : %d\n", config.Server.Port)
}
