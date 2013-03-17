package service

import (
	"testing"
	"arguments/core"
)

func setUp(t *testing.T) *ThesesService {
	s := NewThesesService(core.NewModelMock())
	if s == nil {
		t.Errorf("s was nil")
	}
	s.Testing = true
	return s
}

func TestTheses(t *testing.T) {	
	s := setUp(t)
	expected := 10
	if length := len(s.Theses()); length != expected {
		t.Errorf("expected length %d but was %d", expected, length)	
	}
}
