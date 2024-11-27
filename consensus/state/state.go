package state

import "os"

type State struct {
	id         string // using ip address like server ID
	persistent *PersistentState
	volatile   *VolatileState
}

func (s *State) GetPersistent() *PersistentState {
	if s != nil {
		return s.persistent
	}
	return nil
}

func (s *State) GetVolatile() *VolatileState {
	if s != nil {
		return s.volatile
	}
	return nil
}

func (s *State) GetID() string {
	if s != nil {
		return s.id
	}
	return ""
}

func (s *State) SetID(id string) {
	if s != nil {
		s.id = id
	}
}

func New() *State {
	return &State{
		id:         os.Getenv("SERVER_IP"), // using ip address like server ID
		persistent: newPersistentState(),
		volatile:   NewVolatileState(),
	}
}

type Role string

const (
	Follower  Role = "Follower"
	Candidate Role = "Candidate"
	Leader    Role = "Leader"
)

// PersistentState is the state that must be saved to stable storage
type PersistentState struct {
	state       Role
	currentTerm int
	votedFor    string
	log         []int
}

func newPersistentState() *PersistentState {
	return &PersistentState{
		state:       Follower,
		currentTerm: 0,
		votedFor:    "",
		log:         []int{},
	}
}

func (p *PersistentState) GetState() Role {
	if p != nil {
		return p.state
	}
	return ""
}

func (p *PersistentState) GetCurrentTerm() int {
	if p != nil {
		return p.currentTerm
	}
	return 0
}

func (p *PersistentState) SetCurrentTerm(term int) {
	if p != nil {
		p.currentTerm = term
	}
}

func (p *PersistentState) GetVotedFor() string {
	if p != nil {
		return p.votedFor
	}
	return ""
}

func (p *PersistentState) GetLog() []int {
	if p != nil {
		return p.log
	}
	return []int{}
}

func (p *PersistentState) SetState(state Role) {
	if p != nil {
		p.state = state
	}
}

func (p *PersistentState) IncrementCurrentTerm() {
	if p != nil {
		p.currentTerm++
	}
}

func (p *PersistentState) SetVotedFor(votedFor string) {
	if p != nil {
		p.votedFor = votedFor
	}
}

// VolatileState is the state that need not be saved to stable storage
type VolatileState struct {
}

func NewVolatileState() *VolatileState {
	return &VolatileState{}
}
