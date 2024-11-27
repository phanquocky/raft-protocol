package peer

import (
	"log"
	"net/rpc"
	"os"
	"strings"
	"time"
)

type BroadcastInput struct {
	Message string
}

// hard code these IPs for now in docker-compose.yml file
var PeerIPs = []string{"192.168.1.101", "192.168.1.102", "192.168.1.103", "192.168.1.104", "192.168.1.105"}

var mapConnections = make(map[string]*rpc.Client)

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

func ConnectPeers() {
	for _, ip := range PeerIPs {

		client, err := rpc.DialHTTP("tcp", ip+":1234")
		if err != nil {
			log.Println("Error connecting to peer: ", ip)
			log.Println("Retrying in 1 second ...")
			for err != nil {
				time.Sleep(1 * time.Second)
				client, err = rpc.Dial("tcp", ip+":1234")
			}
		}
		// log.Println("Ping to peer: ", ip)
		// // Synchronous call
		// input := &BroadcastInput{
		// 	Message: "Hello",
		// }
		// var reply int
		// err = client.Call("Handler.Broadcast", input, &reply)
		// if err != nil {
		// 	log.Fatal("arith error:", err)
		// }
		// fmt.Printf("Broadcast... Success %d", reply)

		log.Println("Connect successfully to peer: ", ip)
		mapConnections[ip] = client
	}
}

func GetRPC(ip string) *rpc.Client {
	return mapConnections[ip]
}
