package darwin

import (
	"client/config"
	"fmt"
	"git.zx2c4.com/wireguard-go"
	"github.com/google/gopacket/pcap"
	"log"
	"strings"
)

func StartService(config config.Config) {
	var deviceName string = GetDeviceName()

}

func GetDeviceName() string {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	var no int = 0
	for _, device := range devices {
		if strings.HasPrefix(device.Name, "utun") {
			no++
		}
	}
	return fmt.Sprintf("utun%d", no)
}
