package delegate

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"arguments/core"
)

var model core.Model

func Serve(address string) {
	initialize(address)
	serve(address)
}

func serve(address string) {
	error := http.ListenAndServe(address, nil)
	if error != nil {
		log.Fatal("Error while serving: ", error)
	}
}

func jsonify(element interface{}) string {
	response, error := json.Marshal(element)

	if error != nil {
		log.Fatal("Error while marshaling data ", error)
	}

	json := string(response)
	return json
}

func prepare(writer http.ResponseWriter, request *http.Request) {
	log.Print("Received request on path ", request.URL.Path)
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("access-control-allow-origin", "*")
}

func initialize(address string) {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/api/theses", thesesHandler)
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Go to /api")
}

func thesesHandler(writer http.ResponseWriter, request *http.Request) {
	prepare(writer, request)
	json := jsonify(model.Theses)
	fmt.Fprint(writer, json)
}