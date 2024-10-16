package darwin

import (
	"fmt"
	"net"
	"os/exec"
	"p2p-tunnel-lina/client/config"
	"strings"
)

func DialToServer(ifaceName string) (net.Conn, error) {
	con, err := net.Dial("tcp", config.GetServerURL(config.CONFIG))
	if err != nil {
		return nil, err
	}

	var data []byte = make([]byte, 4096)

	n, err := con.Read(data)
	if err != nil {
		return nil, err
	}

	ipAddress := string(data[:n])
	// _ = ipAddress
	ConfigureDHCP(ifaceName, ipAddress)
	return con, nil
}

func ConfigureDHCP(ifaceName, ipAddress string) {

	err := exec.Command("ifconfig", ifaceName, ipAddress, ipAddress, "up").Run()
	if err != nil {
		panic(err)
	}

	destIP := fmt.Sprintf("%s.0.0.0/8", strings.Split(ipAddress, ".")[0])
	err = exec.Command("route", "-n", "add", "-net", destIP, "-interface", ifaceName).Run()
	if err != nil {
		panic(err)
	}
}
