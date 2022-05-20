package main

import (
	"log"
	"net/http"
)

func main() {
	staticFileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", staticFileServer)

	log.Println("Starting the web application on port 8080...")
}