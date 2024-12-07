package state

import "os"

type State struct {
	id         string // using ip address like server ID
	persistent *PersistentState
	volatile   *VolatileState
}

type Log struct {
	Command string
	Index   int
	Term    int
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
	log         []Log
}

func newPersistentState() *PersistentState {
	return &PersistentState{
		state:       Follower,
		currentTerm: 0,
		votedFor:    "",
		log:         []Log{{Command: "", Index: 0, Term: 0}}, // dummy log entry at index 0, using 1-based indexing
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

func (p *PersistentState) GetLog() []Log {
	if p != nil {
		return p.log
	}
	return []Log{}
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

func (p *PersistentState) AppendLog(log Log) {
	if p != nil {
		p.log = append(p.log, log)
	}
}

func (p *PersistentState) GetLastLogIndex() int {
	if p != nil {
		return len(p.log) - 1
	}
	return 0
}

func (p *PersistentState) GetLastLogTerm() int {
	if p != nil {
		return p.log[p.GetLastLogIndex()].Term
	}
	return 0
}

// VolatileState is the state that need not be saved to stable storage
type VolatileState struct {
	commitIndex int
	// lastApplied int
}

// func (v *VolatileState) IncrementLastApplied() {
// 	if v != nil {
// 		v.lastApplied++
// 	}
// }

// func (v *VolatileState) GetLastApplied() int {
// 	if v != nil {
// 		return v.lastApplied
// 	}
// 	return 0
// }

func (v *VolatileState) GetCommitIndex() int {
	if v != nil {
		return v.commitIndex
	}
	return 0
}

func (v *VolatileState) IncrementCommitIndex() {
	if v != nil {
		v.commitIndex++
	}
}

func (v *VolatileState) SetCommitIndex(commitIndex int) {
	if v != nil {
		v.commitIndex = commitIndex
	}
}

func NewVolatileState() *VolatileState {
	return &VolatileState{
		commitIndex: 0,
		// lastApplied: 0,
	}
}
