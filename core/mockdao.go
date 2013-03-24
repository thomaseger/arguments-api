package core

import "log"

type MockDAO struct {
	id int
}

func (m MockDAO) Create(value interface{}) string {
	m.id = m.id + 1
	log.Print("Added thesis ", m.id)
	return string(m.id)
}

func (m MockDAO) Read(id string) interface{} {
	return nil
}

func (m MockDAO) Update(id string, value interface{}) {

}

func (m MockDAO) Delete(id string) {

}