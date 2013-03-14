package main

import (
	"arguments/delegate"
	"flag"
	"fmt"
	"log"
)

func parseCmd() (string, int) {
	host := flag.String("host", "", "Specifies the host.")
	port := flag.Int("port", 4000, "Specifies the port to listen on.")
	flag.Parse()
	return *host, *port
}

func main() {
	host, port := parseCmd()
	address := fmt.Sprintf("%s:%d", host, port)

	log.Print("Running arguments! on ", address)
	delegate.Serve(address)
}
