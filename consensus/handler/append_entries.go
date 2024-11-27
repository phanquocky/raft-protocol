package consensus

import "log"

type AppendEntriesInput struct {
	Term     int
	LeaderId string
	// PrevLogIndex int
	// PrevLogTerm  int
	Entries []int
	// LeaderCommit int
}

type AppendEntriesOutput struct {
	Term    int
	Success bool
}

func (h *Handler) AppendEntries(input *AppendEntriesInput, reply *AppendEntriesOutput) {
	// Case no log entries, this is a heartbeat
	if len(input.Entries) == 0 {
		reply.Term = h.state.GetPersistent().GetCurrentTerm()
		reply.Success = true
		log.Println("[AppendEntries] Received heartbeat from leader")
		h.leaderHeartbeat <- struct{}{}
		return
	}
}
