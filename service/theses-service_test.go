package service

import (
	"arguments/core"
	"arguments/resting"
	"code.google.com/p/gorest"
	"io/ioutil"
	"net/http"
	"testing"
)

const (
	PortSuffix = ":4040"
	BaseUrl    = "http://localhost" + PortSuffix
	ApiUrl     = BaseUrl + "/api"
)

var model *core.Model

func TestInit(t *testing.T) {
	thesesService := NewThesesService(core.NewMockModel())
	model = thesesService.Model()
	gorest.RegisterService(thesesService)
	http.Handle("/", gorest.Handle())
	go http.ListenAndServe(PortSuffix, nil)
}

func TestTheses(t *testing.T) {
	var target []core.Thesis
	err := get(t, ApiUrl + "/theses", &target)

	if err != nil {
		t.Errorf("Error during test: ", err)
		return
	}

	expected := 10
	if length := len(target); length != expected {
		t.Errorf("Expect %d theses but was %d", expected, length)
	}
}

func TestAddTheses(t *testing.T) {
	lengthBefore := len(model.Theses)
	post(t, ApiUrl + "/theses/add/ThisIsATest.")
	lengthAfter := len(model.Theses)

	if lengthAfter != (lengthBefore + 1) {
		t.Errorf("Add new thesis failed. Expected length is %d, actual it was %d.", lengthBefore + 1, lengthAfter)
	}
}

func post(t *testing.T, path string) error {
	_, postError := http.Post(ApiUrl + path, "", nil)
	if postError != nil {
		t.Errorf("Error during POST: ", postError)
		return postError
	}

	return nil
}

func get(t *testing.T, path string, target interface{}) error {
	body := resting.GetResource(t, path)

	bytes, readError := ioutil.ReadAll(body)
	if readError != nil {
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
