package service

import (
	"testing"
	"net/http"
	"io/ioutil"
	"code.google.com/p/gorest"
	"arguments/core"
)

const (
	PortSuffix = ":4040"
	BaseUrl = "http://localhost" + PortSuffix
	ApiUrl = BaseUrl + "/api"
)

func TestInit(t *testing.T) {
	thesesService := NewThesesService(core.NewModelMock())
	gorest.RegisterService(thesesService)	    
    http.Handle("/", gorest.Handle())
	go http.ListenAndServe(PortSuffix, nil)
}

func TestTheses(t *testing.T) {	
	var target []core.Thesis
	err := get(t, "/theses", &target)

	if(err != nil) {
		t.Errorf("Error during test: ", err)
		return
	}

	expected := 10
	if length := len(target); length != expected {
		t.Errorf("Expect %d theses but was %d", expected, length)
	}
}

func get(t *testing.T, path string, target interface{}) error {
	resp, getError := http.Get(ApiUrl + path)
	if getError != nil {
		t.Errorf("Error during GET: ", getError)
		return getError
	}
	
	bytes, readError := ioutil.ReadAll(resp.Body)
	if(readError != nil) {
		t.Errorf("Error reading body: ", readError)
		return readError
	}

	marshaller := gorest.NewJSONMarshaller()
	unmarshalError := marshaller.Unmarshal(bytes, target)
	if unmarshalError != nil {
		t.Errorf("Error unmarshaling bytes: ", unmarshalError)
		return unmarshalError
	}

	return nil
}
