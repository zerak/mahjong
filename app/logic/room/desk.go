package room

type Desk struct {
	Id int64
}

func (d *Desk) Update(tsl int64) {

}

func NewDesk() *Desk {
	// the id 100 for test
	return &Desk{Id: 100}
}
