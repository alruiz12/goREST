package config

import (
	"net"
	"strings"
)

func GetMyIP(givenInterface string )string{
	var MyIP string
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, iface := range ifaces {
		if iface.Flags & net.FlagUp == 0 {
			continue //interface down
		}
		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}
		if strings.Compare(iface.Name, givenInterface) == 0 {
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type){
				case *net.IPNet: ip = v.IP
				case *net.IPAddr: ip = v.IP
				}
				ip = ip.To4()
				MyIP = ip.String()
				break

			}
			break
		}

	}
	return MyIP
}