package main

import (
	"client/config"
	"fmt"
	"net"
	"os"
)

func main() {

	config, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	serverAddress := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	go receivePackets(conn)

	for {
		fmt.Println("Enter a message to send to other client : ")
		var msg string
		fmt.Scanln(&msg)

		fmt.Println("Enter the IP adress of the destination : ")
		var destIp string
		fmt.Scanln(&destIp)

		packet := fmt.Sprintf("%s:%s", destIp, msg)

		conn.Write([]byte(packet))
	}
}

func receivePackets(conn net.Conn) {
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Connection closed by server")
			os.Exit(1)
		}
		packet := buf[:n]
		fmt.Printf("%s\n", packet)
	}
}
