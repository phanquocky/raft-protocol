package main

import (
	"fmt"
	"log"
	"net/rpc"
	consensus "raft/consensus/handler"
)

// type BroadcastInput struct {
// 	Message string
// }

func main() {

	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	input := &consensus.BroadcastInput{
		Message: "Hello",
	}
	var reply int
	err = client.Call("Handler.Broadcast", input, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Broadcast... Success %d", reply)
}
