package core

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
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
}

func (m Model) FindThesis(id string) Thesis {
	index, _ := strconv.Atoi(id)
	return m.Theses[index]
}

func (m Model) FindArgument(thesisId, argumentId string) Argument {
	thesisIndex, _ := strconv.Atoi(thesisId)
	argumentIndex, _ := strconv.Atoi(argumentId)
	return m.Theses[thesisIndex].Arguments[argumentIndex]
}

func NewModel() *Model {
	return NewMockModel()
}

func NewMockModel() *Model {
	model := NewGeneratedMockModel()

	t := model.Theses

	t[9].Text = "Das Internet ist eine Kulturleistung der Menschheit von historischer Bedeutung."
	t[1].Text = "Die Menschen müssen darauf vertrauen dürfen, dass die Technologie ihnen nutzt."
	t[2].Text = "Auch im Netz kann sich Freiheit nur dann entwickeln, wenn berechtigtes Vertrauen in die Sicherheit herrscht."
	t[7].Text = "Straftaten im Internet müssen in verhältnismäßiger Form verfolgt werden. Die Anonymität des Netzes und die damit erschwerte Arbeit der Justiz wird zunehmend für kriminelle Zwecke missbraucht."
	t[4].Text = "Wirtschaft und Verwaltung haben ihnen anvertraute Daten vor Hackerangriffen zu schützen."
	t[8].Text = "Bürger dürfen die Verantwortung für ihre Sicherheit nicht auf andere abwälzen."
	t[6].Text = "Das freie und sichere Internet ist eine wichtige Triebfeder für eine Stärkung der Demokratie in aller Welt."
	t[3].Text = "Der Himmel ist blau."
	t[5].Text = "Äpfel schmecken besser als Birnen."
	t[0].Text = "Max ist cooler als Moritz."
	return model
}

func NewGeneratedMockModel() *Model {
	var model Model
	var mock DAO
	mock = NewMockDAO()
	model.SetDAO(&mock)

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 10; i++ {
		id := strconv.Itoa(i)
		thesis := Thesis{
			Id:   id,
			Text: fmt.Sprintf("Thesis %d", i),
		}
		for j := 0; j < 1+rand.Intn(9); j++ {
			id := strconv.Itoa(j)
			argument := Argument{
				Id:     id,
				Text:   randomString(100 + rand.Intn(900)),
				Votes:  rand.Intn(1234),
				Contra: (rand.Intn(2)%2 == 0),
			}
			mock.Create(argument)
			thesis.Arguments = append(thesis.Arguments, argument)
		}
		mock.Create(thesis)
		model.Theses = append(model.Theses, thesis)
	}
	return &model
}

func randomString(l int) string {
	text := "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Lorem ipsum dolor sit amet, consectetuer adipiscing elit, sed diam nonummy nibh euismod tincidunt ut laoreet dolore magna aliquam erat volutpat. Ut wisi enim ad minim veniam, quis nostrud exerci tation ullamcorper suscipit lobortis nisl ut aliquip ex ea commodo consequat. Duis autem vel eum iriure dolor in hendrerit in vulputate velit esse molestie consequat, vel illum dolore eu feugiat nulla facilisis at vero eros et accumsan et iusto odio dignissim qui blandit praesent luptatum zzril delenit augue duis dolore te feugait nulla facilisi. Nam liber tempor cum soluta nobis eleifend option congue nihil imperdiet doming id quod mazim placerat facer."
	start := rand.Intn(500)
	return text[start : start+l]
}

type Thesis struct {
	Id        string
	Text      string
	Arguments []Argument
}

type Argument struct {
	Id               string
	Text             string
	Votes            int
	Contra           bool
	CounterArguments []Argument
}
