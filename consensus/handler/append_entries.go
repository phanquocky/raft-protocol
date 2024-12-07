package consensus

import (
	"log"
	"raft/consensus/state"
)

type LogEntry struct {
	Command string
	Term    int
}

type AppendEntriesInput struct {
	Term         int
	LeaderId     string
	PrevLogIndex int
	PrevLogTerm  int
	Entries      []LogEntry
	LeaderCommit int
}

type AppendEntriesOutput struct {
	Term    int
	Success bool
}

func (h *Handler) AppendEntries(input *AppendEntriesInput, reply *AppendEntriesOutput) error {
	// Case no log entries, this is a heartbeat
	if len(input.Entries) == 0 {
		reply.Term = h.state.GetPersistent().GetCurrentTerm()
		reply.Success = true
		// log.Println("[AppendEntries] Received heartbeat from leader")
		h.leaderHeartbeat <- struct{}{}
		return nil
	}

	// Case log entries present
	// 1. Reply false if term < currentTerm
	if input.Term < h.state.GetPersistent().GetCurrentTerm() {
		reply.Term = h.state.GetPersistent().GetCurrentTerm()
		reply.Success = false
		log.Println("[AppendEntries] Received AppendEntries from leader with lower term")
	}

	// 2. Reply false if log doesn't contain an entry at prevLogIndex whose term matches prevLogTerm
	if h.state.GetPersistent().GetLastLogIndex() < input.PrevLogIndex {
		reply.Term = h.state.GetPersistent().GetCurrentTerm()
		reply.Success = false
		log.Println("[AppendEntries] Log doesn't contain entry at prevLogIndex")
	}

	// 3. If an existing entry conflicts with a new one (same index but different terms), delete the existing entry and all that follow it
	// TODO: I don't understand this part

	// 4. Append any new entries not already in the log
	for _, entry := range input.Entries {
		h.state.GetPersistent().AppendLog(state.Log{
			Command: entry.Command,
			Index:   h.state.GetPersistent().GetLastLogIndex() + 1,
			Term:    entry.Term,
		})
	}

	// 5. If leaderCommit > commitIndex, set commitIndex = min(leaderCommit, index of last new entry)
	if input.LeaderCommit > h.state.GetVolatile().GetCommitIndex() {
		h.state.GetVolatile().SetCommitIndex(min(input.LeaderCommit, h.state.GetPersistent().GetLastLogIndex()))
	}

	return nil
}
