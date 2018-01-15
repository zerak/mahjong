package pattern

import (
	"testing"
)

func Test_Pattern(t *testing.T) {
	// pattern module
	pm := NewManager()
	pm.Register(QINGYISE, QingYiSe{})
	// register other patterns

	qysCards := "1234"
	pms := pm.GetPattern([]byte(qysCards))
	t.Log("pattern:", pms)
}
