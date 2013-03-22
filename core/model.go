package core

import "fmt"

type Model struct {
	dao    *DAO
	Theses []Thesis
}

func (m Model) SetDAO(dao *DAO) {
	m.dao = dao
}

func (m Model) AddThesis(t Thesis) {
	m.Theses = append(m.Theses, t)
	id := (*m.dao).Create(t)
	t.Id = id
}

func NewModel() *Model {
	return NewMockModel()
}

func NewMySQLModel() *Model {
	//Use a mysqldao to access data and 
	//create initial model
	return nil
}

func NewMockModel() *Model {
	var model Model
	var mock DAO

	mock = MockDAO{}

	model.SetDAO(&mock)
	
	for i := 0; i < 10; i++ {
		thesis := Thesis{
			Text: fmt.Sprintf("Thesis %d", i),
		}
		for j := 0; j < 10; j++ {
			argument := Argument{
				Text: fmt.Sprintf("This is Argument %d in Thesis %d", j, i),
			}
			thesis.Arguments = append(thesis.Arguments, argument)
		}
		model.Theses = append(model.Theses, thesis)
	}
	return &model
}

type Thesis struct {
	Id        string
	Text      string
	Arguments []Argument
}

type Argument struct {
	Id               string
	Text             string
	Votes            int32
	CounterArguments []Argument
}
