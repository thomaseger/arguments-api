package core

import (
	"fmt"
	"log"
	"time"
)

type MockDAO struct {
	index int
	thesisStore map[string] Thesis
	argumentStore map[string] Argument
}

func NewMockDAO() *MockDAO {
	m := MockDAO {
		thesisStore: make(map[string]Thesis),
		argumentStore: make(map[string]Argument),
	}
	return &m
}

func (m MockDAO) Create(value interface{}) string {
	m.index = m.index + 1

	switch t := value.(type) {
	case nil:
		log.Fatal("trying to create nil value in MockDAO!")
	case Thesis:
		id := newId("t", m.index)
		m.thesisStore[id] = value.(Thesis)
		log.Print("Added thesis ", id)
		return id
	case Argument:
		id := newId("a", m.index)
		m.argumentStore[id] = value.(Argument)
		log.Print("Added argument ", id)
		return id
	default:
		log.Fatalf("Can not handle type ", t)
	}
	return "unknown"
}

func (m MockDAO) Read(id string) interface{} {
	if thesis, ok := m.thesisStore[id]; ok {
		return thesis
	}
	return m.argumentStore[id]
}

func (m MockDAO) Update(id string, value interface{}) {

}

func (m MockDAO) Delete(id string) {

}

func newId(prefix string, index int) string {
	return fmt.Sprintf("%s%d-%d", prefix, index, time.Now().UTC().UnixNano())
}
