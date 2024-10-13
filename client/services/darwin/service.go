package darwin

import (
	"client/config"
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

func StartService(config config.Config) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Available network devices :")
	for _, device := range devices {
		fmt.Println(device.Name)
		for _, address := range device.Addresses {
			fmt.Printf("IP : %v , Subnet : %v \n", address.IP, address.Netmask)
		}
		fmt.Println("================")
	}
}
