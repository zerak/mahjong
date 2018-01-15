// pattern model
// hand of cards that fit
// 3n2 or some special
package pattern

type Module interface {
	// Check check hand of cards hu or not
	Check([]byte) bool
}

type Manager struct {
	pattern map[int]Module
}

func (m *Manager) Check(id int, hands []byte) bool {
	if v, ok := m.pattern[id]; ok {
		return v.Check(hands)
	}
	return false
}

// Register register a handler of a custom pattern
func (m *Manager) Register(id int, handler Module) {
	if _, ok := m.pattern[id]; !ok {
		m.pattern[id] = handler
	}
}

// GetPattern get all patterns that the hand of cards fit
func (m *Manager) GetPattern(hands []byte) []int {
	var p = []int{}
	for k, v := range m.pattern {
		if v.Check(hands) {
			p = append(p, k)
		}
	}
	return p
}

func NewManager() *Manager {
	return &Manager{
		pattern: make(map[int]Module),
	}
}
