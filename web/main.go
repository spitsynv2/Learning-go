package main

import (
	"github.com/spitsynv2/web/http_server"
	"log"
	"net/http"
)

// main.go
func main() {
	server := web.NewPlayerServer(&web.InMemoryPlayerStore{})
	log.Fatal(http.ListenAndServe(":8080", server))
}
