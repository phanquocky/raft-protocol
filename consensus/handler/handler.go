package consensus

import "raft/consensus/state"

// Only methods that satisfy these criteria will be made available for remote access; other methods will be ignored:

// the method's type is exported.
// the method is exported.
// the method has two arguments, both exported (or builtin) types.
// the method's second argument is a pointer.
// the method has return type error.

type Handler struct {
	state              *state.State
	leaderHeartbeat    chan struct{}
	resetElectionTimer chan struct{}
}

func New() *Handler {
	handler := &Handler{
		state:              state.New(),
		leaderHeartbeat:    make(chan struct{}),
		resetElectionTimer: make(chan struct{}),
	}

	return handler
}

func (h *Handler) Start() {
	go h.LeaderElection()

	// case Leader will start sending heartbeat periodically
	go h.sendPeriodicHeartbeats()
}
