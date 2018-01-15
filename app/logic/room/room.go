package room

import "errors"

var ErrAlreadyAdd = errors.New("already add")
var ErrNotHave = errors.New("not have desk")

type Manager struct {
	desks map[int64]*Desk
}

func (m *Manager) AddDesk(d *Desk) error {
	if _, ok := m.desks[d.Id]; ok {
		return ErrAlreadyAdd
	}
	m.desks[d.Id] = d
	return nil
}

func (m *Manager) GetDesk(id int64) (d *Desk, err error) {
	if v, ok := m.desks[id]; ok {
		return v, nil
	}
	return nil, ErrNotHave
}

func (m *Manager) DelDesk(d *Desk) {
	delete(m.desks, d.Id)
}

func (m *Manager) DelDeskById(id int64) {
	delete(m.desks, id)
}

// time since last frame
func (m *Manager) Update(tsl int64) {
	for _, v := range m.desks {
		v.Update(tsl)
	}
}

func NewManager() *Manager {
	return &Manager{desks: make(map[int64]*Desk)}
}
