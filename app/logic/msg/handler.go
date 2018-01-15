package msg

import (
	"mahjong/app/logic/room"
)

type Handler interface {
	ProcessMsg(d *room.Desk)
}

type HandleManager struct {
	handlers map[int]Handler
}

func (hm *HandleManager) Register(msg int, handle Handler) {
	if _, ok := hm.handlers[msg]; !ok {
		hm.handlers[msg] = handle
	}
}

func NewManager() *HandleManager {
	return &HandleManager{handlers: make(map[int]Handler)}
}
