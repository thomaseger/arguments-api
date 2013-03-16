package service

import (
	"testing"
	"arguments/core"
)

func TestTheses(t *testing.T) {
	service := NewThesesService(core.NewModelMock())
	if service == nil {
		t.Errorf("service was nil")
	}

	if length := len(service.Theses()); length != 10 {
		t.Errorf("expected length 10 but was %d", length)	
	}
}
