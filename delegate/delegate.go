package delegate

import (
  "log"
  "net/http"
  "encoding/json"
  "arguments/core"
)

func Json(request *http.Request) string {
  response, error := json.Marshal(core.NewModel())
  
  if error != nil {
    log.Fatal("Error while marshaling data ", error)
  }
  
  json := string(response)
  return json
}
