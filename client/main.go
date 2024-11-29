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

	requestInput := &consensus.RequestVoteInput{
		Term:        1,
		CandidateId: "1",
	}
	var requestOutput consensus.RequestVoteOutput
	err = client.Call("Handler.RequestVote", requestInput, &requestOutput)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	log.Printf("RequestVote... Success %v", requestOutput)

	appendEntriesInput := &consensus.AppendEntriesInput{
		Term:     1,
		LeaderId: "1",
		Entries:  []int{},
	}

	var appendEntriesOutput consensus.AppendEntriesOutput
	err = client.Call("Handler.AppendEntries", appendEntriesInput, &appendEntriesOutput)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	log.Printf("AppendEntries... Success %v", appendEntriesOutput)

}
