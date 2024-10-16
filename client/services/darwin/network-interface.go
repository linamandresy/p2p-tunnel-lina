package darwin

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

const (
	KERNEL_CONTROL_PROTOCOL = 2
	AF_SYSTEM               = 32
	AF_SYS_CONTROL          = 2
	UTUN_CONTROL_NAME       = "com.apple.net.utun_control"
	CTLIOCGINFO             = 0xC0644E03
)

type ctl_info struct {
	ctl_id   uint32
	ctl_name [96]byte
}

type sockaddr_ctl struct {
	sc_len     uint8
	sc_family  uint8
	ss_sysaddr uint16
	sc_id      uint32
	sc_unit    uint32
	pad        [5]uint32
}

func StartInterface() (int, string, error) {

	fd, err := syscall.Socket(syscall.AF_SYSTEM, syscall.SOCK_DGRAM, KERNEL_CONTROL_PROTOCOL)
	if err != nil {
		return -1, "", fmt.Errorf("failed to create socket: %v", err)
	}

	var info ctl_info
	copy(info.ctl_name[:], UTUN_CONTROL_NAME)

	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), CTLIOCGINFO, uintptr(unsafe.Pointer(&info)))
	if errno != 0 {
		syscall.Close(fd)
		return -1, "", fmt.Errorf("ioctl CTLIOCGINFO failed: %v", errno)
	}

	var sc sockaddr_ctl
	sc.sc_len = uint8(unsafe.Sizeof(sc))
	sc.sc_family = AF_SYSTEM
	sc.ss_sysaddr = AF_SYS_CONTROL
	sc.sc_id = info.ctl_id
	sc.sc_unit = 0

	_, _, errno = syscall.Syscall(syscall.SYS_CONNECT, uintptr(fd), uintptr(unsafe.Pointer(&sc)), unsafe.Sizeof(sc))
	if errno != 0 {
		syscall.Close(fd)
		return -1, "", fmt.Errorf("connect syscall failed: %v", errno)
	}

	var ifName [16]byte
	optLen := uint32(len(ifName))
	_, _, errno = syscall.Syscall6(syscall.SYS_GETSOCKOPT, uintptr(fd), uintptr(KERNEL_CONTROL_PROTOCOL), uintptr(AF_SYS_CONTROL), uintptr(unsafe.Pointer(&ifName[0])), uintptr(unsafe.Pointer(&optLen)), 0)
	if errno != 0 {
		syscall.Close(fd)
		return -1, "", fmt.Errorf("failed to get interface name: %v", errno)
	}

	ifaceName := string(ifName[:])
	ifaceName = ifaceName[:len(ifaceName)-1]
	ifaceName = strings.Trim(ifaceName, "\x00")

	return fd, ifaceName, nil
}
