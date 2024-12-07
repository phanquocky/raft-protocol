package consensus

import (
	"log"
	"raft/consensus/external/peer"
	"raft/consensus/state"
)

type ExecuteInput struct {
	Message string
}

// Execute is the RPC handler for client to send command to the leader
func (h *Handler) Execute(args *ExecuteInput, reply *int) error {
	log.Println("Executing message: ", args.Message)
	log.Println("Current state: ", h.state.GetPersistent().GetState())

	// check if the server is the leader
	if h.state.GetPersistent().GetState() != state.Leader {
		log.Println("Not the leader, cannot execute")
		*reply = -1
		return nil
	}

	// get the last log index and term before appending the command
	prevLogIndex := h.state.GetPersistent().GetLastLogIndex()
	prevLogTerm := h.state.GetPersistent().GetLastLogTerm()

	// append the command to the log
	h.state.GetPersistent().AppendLog(state.Log{
		Command: args.Message,
		Index:   h.state.GetPersistent().GetLastLogIndex() + 1,
		Term:    h.state.GetPersistent().GetCurrentTerm(),
	})

	// issue append entries to all peers
	result := make(chan AppendEntriesOutput)
	for _, p := range peer.PeerIPs {
		go func(p string) {
			var output AppendEntriesOutput
			err := peer.GetRPC(p).Call("Handler.AppendEntries", &AppendEntriesInput{
				Term: h.state.GetPersistent().GetCurrentTerm(),
				Entries: []LogEntry{
					{
						Command: args.Message,
						Term:    h.state.GetPersistent().GetCurrentTerm(),
					},
				},
				LeaderId:     h.state.GetID(),
				PrevLogIndex: prevLogIndex,
				PrevLogTerm:  prevLogTerm,
				LeaderCommit: h.state.GetVolatile().GetCommitIndex(),
			}, &output)
			if err != nil {
				log.Printf("[sendPeriodicHeartbeats] Error sending heartbeat to %s: %v\n", p, err)
			}

			result <- output

		}(p)
	}

	// 6. Wait for append entries response
	appended := 1
	for i := 0; i < len(peer.PeerIPs); i++ {
		select {
		case output := <-result:
			if output.Success {
				appended++
			}
		}
	}

	if appended > len(peer.PeerIPs)/2 {
		// execute command
		log.Println("[Execute] command ", args.Message)

		// commit the entry
		h.state.GetVolatile().IncrementCommitIndex()
	}

	*reply = 1
	return nil
}
