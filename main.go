package main

import (
	"log"
	"net/http"
)

func main() {
	staticFileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", staticFileServer)

	log.Println("Starting the web application on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
