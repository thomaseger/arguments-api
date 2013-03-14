package delegate

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "arguments/core"
)

func Json(request *http.Request) string {
  var thesis core.Thesis
  thesis.Text = "A test thesis"
  thesis.Arguments = make([]core.Argument, 3)
  
  response, error := json.Marshal(thesis)
  
  if error != nil {
    log.Fatal("Error while marshaling data ", error)
  }
  
  json := fmt.Sprintf("%s%s", string(response), "\n")
  log.Print("Prepared json: ", json)
  return json
}
