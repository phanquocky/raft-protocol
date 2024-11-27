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

// type Args struct {
// 	A, B int
// }

// type Quotient struct {
// 	Quo, Rem int
// }

// type Arith int

// func (t *Arith) Multiply(args *Args, reply *int) error {
// 	*reply = args.A * args.B
// 	return nil
// }

// func (t *Arith) Divide(args *Args, quo *Quotient) error {
// 	if args.B == 0 {
// 		return errors.New("divide by zero")
// 	}
// 	quo.Quo = args.A / args.B
// 	quo.Rem = args.A % args.B
// 	return nil
// }

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
