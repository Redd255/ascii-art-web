package main

import (
	"asciiart/serv"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", serv.Index)
	http.HandleFunc("/ascii-art", serv.AsciiWeb)
	log.Println("http//localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
