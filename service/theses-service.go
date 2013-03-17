package service

import (
	"code.google.com/p/gorest"
	"arguments/core"
	"log"
)

type ThesesService struct {
	Testing bool

	model *core.Model

	gorest.RestService `root:"/api" consumes:"application/json" produces:"application/json"`

	//GET Endpoints
	theses gorest.EndPoint `method:"GET" path:"/theses" output:"[]Thesis"`
	thesis gorest.EndPoint `method:"GET" path:"/theses/{thesisId:int}" output:"Thesis"`
	arguments gorest.EndPoint `method:"GET" path:"/theses/{thesisId:int}/arguments" output:"[]Argument"`
	argument gorest.EndPoint `method:"GET" path:"/theses/{thesisId:int}/arguments/{argumentId:int}" output:"Argument"`
}

func NewThesesService(m *core.Model) *ThesesService {
	service := ThesesService {
		model: m,
	}
	return &service
}

func (service ThesesService) Theses() []core.Thesis {
	service.prepareResponse()
	return service.model.Theses
} 

func (service ThesesService) Thesis(thesisId int) core.Thesis {
	service.prepareResponse()
	return service.model.Theses[thesisId]
}

func (service ThesesService) Arguments(thesisId int) []core.Argument {
	service.prepareResponse()
	return service.model.Theses[thesisId].Arguments
}

func (service ThesesService) Argument(thesisId, argumentId int) core.Argument {
	service.prepareResponse()
	return service.model.Theses[thesisId].Arguments[argumentId]
}

func (service ThesesService) prepareResponse() {
	if(!service.Testing) {
		log.Printf("Received request.")
		service.RB().SetHeader("Access-Control-Allow-Origin", "*")
	}
}
