package core

import "fmt"

type Model struct {
	dao    *DAO
	Theses []Thesis
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

func NewModel() *Model {
	return NewModelMock()
}

func NewMySQLModel() *Model {
	//Use a mysqldao to access data and 
	//create initial model
	return nil
}

func NewModelMock() *Model {
	var model Model
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
