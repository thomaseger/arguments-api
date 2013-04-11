package service

import (
	"arguments/core"
	"arguments/resting"
	"bytes"
	"code.google.com/p/gorest"
	"io"
	"io/ioutil"
	"log"
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
	logTestRun("TestInit")

	thesesService := NewThesesService(core.NewMockModel())
	model = thesesService.Model()
	gorest.RegisterService(thesesService)
	http.Handle("/", gorest.Handle())
	go http.ListenAndServe(PortSuffix, nil)
}

func TestTheses(t *testing.T) {
	logTestRun("TestTheses")

	var target []core.Thesis

	get(t, ApiUrl+"/theses", &target)

	expected := 10
	if length := len(target); length != expected {
		t.Errorf("Expect %d theses but was %d", expected, length)
	}
}

func TestThesesCount(t *testing.T) {
	logTestRun("TestThesesCount")

	count := thesesCount(t)
	expected := 10
	if count != expected {
		t.Errorf("Expect %d theses but was %d", expected, count)
	}
}

func TestAddTheses(t *testing.T) {
	logTestRun("TestAddTheses")

	lengthBefore := thesesCount(t)

	thesis := core.Thesis{
		Text: "Test Thesis.",
	}

	marshaller := gorest.NewJSONMarshaller()
	target, marshalError := marshaller.Marshal(thesis)

	if marshalError != nil {
		t.Errorf("Error unmarshaling bytes: ", marshalError)
	}

	reader := bytes.NewReader(target)

	post(t, ApiUrl+"/theses", "application/json", reader, target)

	lengthAfter := thesesCount(t)

	if lengthAfter != (lengthBefore + 1) {
		t.Errorf("Add new thesis failed. Expected length is %d, actual it was %d.", lengthBefore+1, lengthAfter)
	}
}

func thesesCount(t *testing.T) int {
	var count int
	get(t, ApiUrl+"/theses/count", &count)
	return count
}

func post(t *testing.T, url string, mime string, reader io.Reader, target interface{}) {
	body := resting.PostResource(t, url, mime, reader)

	_, readError := ioutil.ReadAll(body)
	if readError != nil {
		t.Errorf("Error reading body: ", readError)
	}
}

func get(t *testing.T, url string, target interface{}) {
	body := resting.GetResource(t, url)

	bytes, readError := ioutil.ReadAll(body)
	if readError != nil {
		t.Errorf("Error reading body: ", readError)
	}

	marshaller := gorest.NewJSONMarshaller()
	unmarshalError := marshaller.Unmarshal(bytes, target)
	if unmarshalError != nil {
		t.Errorf("Error unmarshaling bytes: ", unmarshalError)
	}
}

func logTestRun(name string) {
	log.Print("Running " + name)
}
