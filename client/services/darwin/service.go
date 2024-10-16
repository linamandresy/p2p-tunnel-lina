package darwin

import (
	"fmt"
	"p2p-tunnel-lina/client/config"
	"p2p-tunnel-lina/client/network"
	"syscall"
)

func StartService(config config.Config) {
	fd, ifaceName, err := StartInterface()
	if err != nil {
		panic(err)
	}
	defer syscall.Close(fd)

	network.NETWORK_CONN, err = DialToServer(ifaceName)
	if err != nil {
		panic(err)
	}
	defer network.NETWORK_CONN.Close()
	// StartListener(ifaceName)

	// StartListener("en0")

	var buf []byte = make([]byte, 4600)
	for {
		n, err := syscall.Read(fd, buf)
		if err != nil {
			panic(err)
		}
		fmt.Println(buf[:n])
		fmt.Println("==============")
	}
}
