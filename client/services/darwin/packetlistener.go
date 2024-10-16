package darwin

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func StartListener(deviceName string) {
	// devices, err := pcap.FindAllDevs()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Available network devices :")
	// for _, device := range devices {
	// 	fmt.Println(device.Name)
	// 	for _, address := range device.Addresses {
	// 		fmt.Printf("IP :%s, Subnet : %s \n", address.IP.String(), address.Netmask.String())
	// 	}
	// 	fmt.Println()
	// }

	handle, err := pcap.OpenLive(deviceName, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	var filter = "tcp"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listening on %s [filter : %s]\n", deviceName, filter)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer != nil {
			ip, _ := ipLayer.(*layers.IPv4)
			srcIP := ip.SrcIP
			destIP := ip.DstIP

			fmt.Printf("Source IP : %s, Destination IP : %s\nValeur : %v", srcIP, destIP, packet)
		}

	}
}
