package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"raft/consensus/external/peer"
	consensus "raft/consensus/handler"
	"time"
)

func main() {
	handler := consensus.New()
	log.Println("Start Consensus Handler")
	handler.Start()
	rpc.Register(handler)
	rpc.HandleHTTP()

	go func() {
		time.Sleep(1000) // wait for other servers to start
		log.Println("Connecting to peers ...")
		peer.ConnectPeers()
	}()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	log.Println("Server started ...")
	log.Println(peer.PeerIPs)
	http.Serve(l, nil)

}
