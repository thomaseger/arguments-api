package main

import (
  "fmt"
  "net/http"
  "log"
  "flag"
  "arguments/delegate"
)

func parseCmd() (string, int) {
  host := flag.String("host", "", "Specifies the host.");
  port := flag.Int("port", 4000, "Specifies the port to listen on.");
  flag.Parse()
  return *host, *port
}

func handle(writer http.ResponseWriter, request *http.Request) {
  log.Print("Received request on path ", request.URL.Path)
  writer.Header().Set("Content-Type", "application/json")
  writer.Header().Set("access-control-allow-origin", "*")
  json := delegate.Json(request)
  fmt.Fprint(writer, json)
}

func main() {
  host, port := parseCmd()
  address := fmt.Sprintf("%s:%d", host, port);

  log.Print("Running arguments! on ", address)

  http.HandleFunc("/", handle)
  error := http.ListenAndServe(address, nil)
  
  if error != nil {
    log.Fatal("Error while serving: ", error)
  }  
}

