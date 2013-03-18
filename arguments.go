package main

import (
	"arguments/service"
	"arguments/core"
	"code.google.com/p/gorest"
	"net/http"
	"flag"
	"fmt"
	"log"
)

type Arguments struct {
	address string
	model *core.Model
}

func (a *Arguments) parseCmd() {
	host := flag.String("host", "", "Specifies the host.")
	port := flag.Int("port", 4000, "Specifies the port to listen on.")
	flag.Parse()

	a.address = fmt.Sprintf("%s:%d", *host, *port)
}

func (a *Arguments) serve() {
	a.model = core.NewModel()

	thesesService := service.NewThesesService(a.model)
	gorest.RegisterService(thesesService)	    
    
    http.Handle("/", gorest.Handle())

    log.Print("Running arguments! on ", a.address)   
    error := http.ListenAndServe(a.address, nil)
    if(error != nil) {
    	log.Fatal("Error while serving: ", error)
    }
}

func main() {
	args := new(Arguments)
	args.parseCmd()
	args.serve()
}
