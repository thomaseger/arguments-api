package core

import (
	"fmt"
	"math/rand"
)

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
	model := NewGeneratedMockModel()

	t := model.Theses

	t[0].Text = "Das Internet ist eine Kulturleistung der Menschheit von historischer Bedeutung."
	t[1].Text = "Die Menschen müssen darauf vertrauen dürfen, dass die Technologie ihnen nutzt."
	t[2].Text = "Auch im Netz kann sich Freiheit nur dann entwickeln, wenn berechtigtes Vertrauen in die Sicherheit herrscht."
	t[3].Text = "Straftaten im Internet müssen in verhältnismäßiger Form verfolgt werden. Die Anonymität des Netzes und die damit erschwerte Arbeit der Justiz wird zunehmend für kriminelle Zwecke missbraucht."
	t[4].Text = "Wirtschaft und Verwaltung haben ihnen anvertraute Daten vor Hackerangriffen zu schützen."
	t[5].Text = "Bürger dürfen die Verantwortung für ihre Sicherheit nicht auf andere abwälzen."
	t[6].Text = "Das freie und sichere Internet ist eine wichtige Triebfeder für eine Stärkung der Demokratie in aller Welt."

	return model
}

func NewGeneratedMockModel() *Model {
	var model Model
	var mock DAO
	mock = MockDAO{}
	model.SetDAO(&mock)
	
	for i := 0; i < 10; i++ {
		thesis := Thesis{
			Text: fmt.Sprintf("Thesis %d", i),
		}
		for j := 0; j < 10; j++ {
			argument := Argument { 
				Text: fmt.Sprintf("This is Argument %d in Thesis %d", j, i),
				Contra: (rand.Intn(2) % 2 == 0),
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
	Contra 			 bool
	CounterArguments []Argument
}
