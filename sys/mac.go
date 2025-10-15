package sys

import (
	"fmt"
	"net"
)

func GetMAC() (string, error) {
    inf, err := net.InterfaceByName("wlp2s0")
	
    if err != nil {
        return "nil", err
    }
	
	mac := inf.HardwareAddr
	
	return fmt.Sprint(mac), err
}