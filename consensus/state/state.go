package state

type State struct {
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

func New() *State {
	return &State{
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
	votedFor    int
	log         []int
}

func newPersistentState() *PersistentState {
	return &PersistentState{
		state:       Follower,
		currentTerm: 0,
		votedFor:    -1,
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

func (p *PersistentState) GetVotedFor() int {
	if p != nil {
		return p.votedFor
	}
	return -1
}

func (p *PersistentState) GetLog() []int {
	if p != nil {
		return p.log
	}
	return []int{}
}

// VolatileState is the state that need not be saved to stable storage
type VolatileState struct {
}

func NewVolatileState() *VolatileState {
	return &VolatileState{}
}
