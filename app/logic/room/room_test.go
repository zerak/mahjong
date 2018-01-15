package room

import (
	"testing"
)

func Test_Room(t *testing.T) {
	// pattern module
	rm := NewManager()
	rm.AddDesk(NewDesk())
	// add other desks

	rm.Update(0)
	rm.DelDeskById(100)
	v, err := rm.GetDesk(100)
	if err != nil {
		t.Log("there is no desk:", err)
	} else {
		t.Error("must be no desk id:", v.Id)
	}
}
