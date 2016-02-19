package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting Server")

	http.Handle("/", http.FileServer(http.Dir("./public/")))

	log.Println("Listening on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
