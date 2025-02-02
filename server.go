package main

import (
	"net"
	"fmt"
)

func getServerIP() (string, error) {
	// Get a list of all interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	// Iterate over the interfaces
	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			// Check if the address is an IP address
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil { // Check for IPv4
					return ipNet.IP.String(), nil
				}
			}
		}
	}
	return "", fmt.Errorf("no valid IP address found")
}