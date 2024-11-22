package peer

import (
	"log"
	"os"
	"strings"
)

var PeerIPs = []string{"192.168.0.1", "192.168.0.2", "192.168.0.3", "192.168.0.4", "192.168.0.5"}

func init() {
	peerIps := make([]string, 0)
	// remove the IP of server itself
	log.Println("Server IP: ", os.Getenv("SERVER_IP"))

	for _, ip := range PeerIPs {
		if !strings.EqualFold(ip, os.Getenv("SERVER_IP")) {
			peerIps = append(peerIps, ip)
		}
	}

	PeerIPs = peerIps
}
