package core

import "testing"

func TestDestroyIfNeeded(t *testing.T) {
	mock := new(destroyableMocker)
	DestroyIfNeeded(mock)

	if mock.destroyCalled != true {
		t.Errorf("case %d: expected: %v, got: %v", 0, true, false)
	}
}

type destroyableMocker struct {
	destroyCalled bool
}

func (t *destroyableMocker) Destroy() {
	t.destroyCalled = true
}