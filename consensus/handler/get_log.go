package consensus

import (
	"log"
	"raft/consensus/state"
)

type GetLogInput struct {
}

type GetLogOutput struct {
	Log []state.Log
}

func (h *Handler) GetLog(input *GetLogInput, reply *GetLogOutput) error {
	log.Println("[GetLog] Received GetLog request")
	reply.Log = h.state.GetPersistent().GetLog()
	return nil
}
