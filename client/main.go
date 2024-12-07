package main

import (
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

	executeInput := consensus.ExecuteInput{
		Message: "Hello World",
	}

	var reply int
	err = client.Call("Handler.Execute", &executeInput, &reply)
	if err != nil {
		log.Fatal("Execute error:", err)
	}

	log.Println("Execute reply:", reply)

	// get log
	getlogInput := consensus.GetLogInput{}

	var getlogOutput consensus.GetLogOutput
	err = client.Call("Handler.GetLog", &getlogInput, &getlogOutput)
	if err != nil {
		log.Fatal("GetLog error:", err)
	}

	log.Printf("GetLog reply: %+v\n", getlogOutput)
}
