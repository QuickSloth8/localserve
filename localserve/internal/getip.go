package internal

import (
	"log"
	"net"
	"regexp"
	"strings"
)

// Slice defining the precedence of an interface based on its name
// This function currently supports Windows interfaces only, Linux
// support should be added in the future
var iPrecedence = map[string]int{
	"*Wi-Fi*":                 0,
	"*Local Area Connection*": 1,
}

// intName matched to regex pattern to convert it to an integer index
// (the lower the index, the higher the precedence)
func convertIntNameToPrecedence(intName string) int {
	for ptrn := range iPrecedence {
		if matched, err := regexp.MatchString(ptrn, intName); err == nil && matched {
			return iPrecedence[ptrn]
		}
	}
	return len(iPrecedence)
}

// Return and ip address based on the precedence defined in iPrecedence
func GetIp() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	addrsSlice := make([][]net.Addr, len(iPrecedence)+1)

	for _, i := range ifaces {
		iName := i.Name
		if strings.HasPrefix(iName, "Wi-Fi") || strings.HasPrefix(iName, "Local Area Connection") {
			addrs, err := i.Addrs()
			if err != nil {
				log.Fatal(err)
			}

			addrsSlice[convertIntNameToPrecedence(iName)] = addrs
		}
	}

	if len(addrsSlice) == 0 {
		return ""
	}

	for _, addrs := range addrsSlice {
		if addrs != nil {
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				// process IP address
				ipv4 := ip.To4()
				if ipv4 != nil && ipv4.IsGlobalUnicast() {
					return ipv4.String()
				}
			}
		}
	}
	return ""
}
