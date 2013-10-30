package service

import (
	"github.com/thomaseger/arguments-api/core"
	"code.google.com/p/gorest"
	"log"
)

type ThesesService struct {
	model *core.Model

	gorest.RestService `root:"/api" consumes:"application/json" produces:"application/json"`

	theses      gorest.EndPoint `method:"GET" path:"/theses" output:"[]Thesis"`
	thesis      gorest.EndPoint `method:"GET" path:"/theses/{thesisId:string}" output:"Thesis"`
	thesesCount gorest.EndPoint `method:"GET" path:"/theses/count" output:"int"`
	addThesis   gorest.EndPoint `method:"POST" path:"/theses" postdata:"Thesis"`
	arguments   gorest.EndPoint `method:"GET" path:"/theses/{thesisId:string}/arguments" output:"[]Argument"`
	argument    gorest.EndPoint `method:"GET" path:"/theses/{thesisId:string}/arguments/{argumentId:string}" output:"Argument"`
}

func NewThesesService(m *core.Model) *ThesesService {
	service := ThesesService{
		model: m,
	}
	return &service
}

func (service ThesesService) Theses() []core.Thesis {
	service.prepareResponse()
	return service.model.Theses
}

func (service ThesesService) ThesesCount() int {
	service.prepareResponse()
	return len(service.model.Theses)
}

func (service ThesesService) Thesis(thesisId string) core.Thesis {
	service.prepareResponse()
	return service.model.FindThesis(thesisId)
}

func (service ThesesService) Arguments(thesisId string) []core.Argument {
	service.prepareResponse()
	return service.model.FindThesis(thesisId).Arguments
}

func (service ThesesService) Argument(thesisId, argumentId string) core.Argument {
	service.prepareResponse()
	return service.model.FindArgument(thesisId, argumentId)
}

func (s ThesesService) AddThesis(t core.Thesis) {
	s.prepareResponse()
	s.model.AddThesis(t)
}

func (service ThesesService) prepareResponse() {
	log.Printf("Received request.")
	service.RB().SetHeader("Access-Control-Allow-Origin", "*")
}

func (s ThesesService) Model() *core.Model {
	return s.model
}
