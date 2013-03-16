package core

import "fmt"

type Model struct {
	Theses []Thesis
}

type Thesis struct {
	Text      string
	Arguments []Argument
}

type Argument struct {
	Text             string
	Votes            int32
	CounterArguments []Argument
}

func NewModel() *Model {
	return NewModelMock()
}

func NewModelMock() *Model {
	var model Model
	for i := 0; i < 10; i++ {
		thesis := Thesis{
			Text: fmt.Sprintf("Thesis %d", i),
		}
		for j := 0; j < 10; j++ {
			argument := Argument {
				Text: fmt.Sprintf("This is Argument %d in Thesis %d", j, i),
			}
			thesis.Arguments = append(thesis.Arguments, argument)
		}
		model.Theses = append(model.Theses, thesis)
	}
	return &model
}
