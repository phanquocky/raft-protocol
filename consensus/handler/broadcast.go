package consensus

import "log"

type BroadcastInput struct {
	Message string
}

func (t *Handler) Broadcast(args *BroadcastInput, reply *int) error {
	log.Println("Broadcasting message: ", args.Message)
	log.Println("Current state: ", t.state.GetPersistent().GetState())
	*reply = 1
	return nil
}
