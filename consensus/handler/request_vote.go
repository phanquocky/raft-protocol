package consensus

import (
	"log"
	"raft/consensus/state"
)

type RequestVoteInput struct {
	Term        int
	CandidateId string
	// LastLogIndex int
	// LastLogTerm int
}

type RequestVoteOutput struct {
	Term        int
	VoteGranted bool // true means candidate received vote
}

func (h *Handler) RequestVote(input *RequestVoteInput, reply *RequestVoteOutput) error {
	log.Printf("[RequestVote] Received RequestVote from %s, Term: %d\n", input.CandidateId, input.Term)
	// 2. If votedFor is null or candidateId, and candidate's log is at least as up-to-date as receiver's log, grant vote
	// 3. If RPC request or response contains term T > currentTerm: set currentTerm = T, convert to follower
	// 4. If candidateId is not in own peer list, reject vote
	// 5. If votedFor is not null, reject vote
	// 6. If candidate's log is not at least as up-to-date as receiver's log, reject vote
	// 7. If vote granted, reset election timer
	// 8. If election timeout elapses: start new election

	// 1. Reply false if term < currentTerm
	if input.Term < h.state.GetPersistent().GetCurrentTerm() {
		reply.Term = h.state.GetPersistent().GetCurrentTerm()
		reply.VoteGranted = false
		return nil
	}

	// 3. If RPC request or response contains term T > currentTerm: set currentTerm = T, convert to follower
	if input.Term > h.state.GetPersistent().GetCurrentTerm() {
		h.state.GetPersistent().SetCurrentTerm(input.Term)
		h.state.GetPersistent().SetState(state.Follower)
		h.state.GetPersistent().SetVotedFor("")
	}

	// 2. If votedFor is null or candidateId, and candidate's log is at least as up-to-date as receiver's log, grant vote
	if h.state.GetPersistent().GetVotedFor() == "" || h.state.GetPersistent().GetVotedFor() == h.state.GetID() {
		reply.Term = input.Term
		reply.VoteGranted = true
		h.state.GetPersistent().SetVotedFor(input.CandidateId)
		return nil
	}

	// 4. If candidateId is not in own peer list, reject vote
	// 5. If votedFor is not null, reject vote
	// 6. If candidate's log is not at least as up-to-date as receiver's log, reject vote
	reply.Term = h.state.GetPersistent().GetCurrentTerm()
	reply.VoteGranted = false

	// 7. If vote granted, reset election timer
	h.resetElectionTimer <- struct{}{}

	// 8. If election timeout elapses: start new election
	return nil
}
