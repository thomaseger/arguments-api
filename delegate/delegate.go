package delegate

import (
	"net/http"
	"code.google.com/p/gorest"
	"arguments/core"
	"arguments/service"
)

func Serve(address string, model *core.Model) {
	thesesService := service.NewThesesService(model)
	gorest.RegisterService(thesesService)	    
    http.Handle("/", gorest.Handle())   
    http.ListenAndServe(address, nil)
}
