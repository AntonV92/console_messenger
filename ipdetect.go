package main

import (
	"log"
	"net"
)

const LocalNetMask = "ffffff00"

func getLocalAddr() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	var result string

	for _, interf := range interfaces {
		addrs, err := interf.Addrs()
		if err != nil {
			log.Fatal(err)
		}

		for _, add := range addrs {
			if ip, ok := add.(*net.IPNet); ok {

				if ip.Mask.String() == LocalNetMask {
					result = ip.IP.String()
				}
			}
		}
	}

	return result
}
