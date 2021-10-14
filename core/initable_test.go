package core

import "testing"

func TestInitializeIfNeeded(t *testing.T) {
	mock := new(initableMocker)
	InitializeIfNeeded(mock)

	if mock.initCalled != true {
		t.Errorf("case %d: expected: %v, got: %v", 0, true, false)
	}
}

type initableMocker struct {
	initCalled bool
}

func (t *initableMocker) Init() {
	t.initCalled = true
}
