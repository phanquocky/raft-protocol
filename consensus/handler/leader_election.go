package consensus

import (
	"crypto/rand"
	"log"
	"math/big"
	"raft/consensus/external/peer"
	"raft/consensus/state"
	"time"
)

func (h *Handler) LeaderElection() {
	for {
		// Randomize the election timeout to be between 150ms and 300ms, but in experiment we set it to 10-20 seconds
		n, err := rand.Int(rand.Reader, big.NewInt(15))
		if err != nil {
			log.Fatal("[LeaderElection] error randomizing election timeout: ", err)
		}

		randomTimeout := time.Duration(n.Int64()+15) * time.Second

		// timer countdown for election timeout
		timer := time.NewTimer(randomTimeout)

		select {
		case <-timer.C:
			// start election
			log.Println("[LeaderElection] Election timeout reached, starting election")
			go h.startElection()
		case <-h.leaderHeartbeat:
			// reset timer
			log.Println("[LeaderElection] Received heartbeat from leader, resetting election timeout")
			timer.Stop()
		case <-h.resetElectionTimer:
			// reset timer
			log.Println("[LeaderElection] Resetting election timeout")
			timer.Stop()
		}
	}
}

func (h *Handler) startElection() {
	// 1. Set state to candidate
	h.state.GetPersistent().SetState(state.Candidate)

	// 2. Increment current term
	h.state.GetPersistent().IncrementCurrentTerm()

	// 3. Vote for self
	h.state.GetPersistent().SetVotedFor(h.state.GetID())

	// 4. Send RequestVote RPCs to all other servers

	result := make(chan RequestVoteOutput)
	for _, p := range peer.PeerIPs {
		go func(p string) {
			defer func() { result <- RequestVoteOutput{} }() // Ensure a value is always sent.
			var output RequestVoteOutput
			log.Printf("[LeaderElection] Sending RequestVote to %s, Term: %d, CandidateId: %s\n", p, h.state.GetPersistent().GetCurrentTerm(), h.state.GetID())
			err := peer.GetRPC(p).Call("Handler.RequestVote", &RequestVoteInput{
				Term:        h.state.GetPersistent().GetCurrentTerm(),
				CandidateId: h.state.GetID(),
			}, &output)
			if err != nil {
				log.Printf("[LeaderElection] Error sending RequestVote to %s: %v\n", p, err)
			}
			log.Printf("[LeaderElection] Received RequestVote from %s, Term: %d, VoteGranted: %t\n", p, output.Term, output.VoteGranted)
			result <- output
		}(p)
	}

	// // 5. Reset election timer
	// h.resetElectionTimer <- struct{}{}

	// 6. Wait for votes
	votes := 1
	for i := 0; i < len(peer.PeerIPs); i++ {
		log.Println("[LeaderElection] Waiting for votes")
		select {
		case output := <-result:
			log.Printf("[LeaderElection] Received vote: %t", output.VoteGranted)
			if output.VoteGranted {
				votes++
			}
		}
	}

	log.Println("[LeaderElection] Total votes received: ", votes)
	// 7. If votes received from majority of servers: become leader
	if votes > len(peer.PeerIPs)/2 {
		log.Println("[LeaderElection] Majority votes received, becoming leader")
		h.state.GetPersistent().SetState(state.Leader)
		h.leaderHeartbeat <- struct{}{}
	}

	// 8. If AppendEntries RPC received from new leader: convert to follower
	// 9. If election timeout elapses: start new election
}

// Leader send periodic heartbeats to followers
func (h *Handler) sendPeriodicHeartbeats() {
	for {
		time.Sleep(1000 * time.Millisecond)
		if h.state.GetPersistent().GetState() == state.Leader {
			log.Println("[sendPeriodicHeartbeats] Sending heartbeat to followers")
			// send heartbeat to all followers
			for _, p := range peer.PeerIPs {
				go func(p string) {
					output := &AppendEntriesOutput{}
					err := peer.GetRPC(p).Call("Handler.AppendEntries", &AppendEntriesInput{
						Term:     h.state.GetPersistent().GetCurrentTerm(),
						LeaderId: h.state.GetID(),
						Entries:  []int{},
					}, output)
					if err != nil {
						log.Printf("[sendPeriodicHeartbeats] Error sending heartbeat to %s: %v\n", p, err)
					}
				}(p)
			}
		}
	}
}
