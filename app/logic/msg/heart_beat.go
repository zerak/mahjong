package msg

import (
	"fmt"

	"mahjong/app/logic/room"
)

type HeatBeat struct {
}

func (h HeatBeat) ProcessMsg(d *room.Desk) {
	fmt.Println("process desk:", d.Id)
}
