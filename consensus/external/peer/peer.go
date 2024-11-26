package peer

import (
	"log"
	"os"
	"strings"
)

// hard code these IPs for now in docker-compose.yml file
var PeerIPs = []string{"192.168.1.101", "192.168.1.102", "192.168.1.103", "192.168.1.104", "192.168.1.105"}

func init() {
	peerIps := make([]string, 0)
	// remove the IP of server itself
	log.Println("Server IP: ", os.Getenv("SERVER_IP"))

	for _, ip := range PeerIPs {
		if !strings.EqualFold(ip, os.Getenv("SERVER_IP")) { // remove the IP of server itself, it is configured in docker-compose.yml
			peerIps = append(peerIps, ip)
		}
	}

	PeerIPs = peerIps
}
