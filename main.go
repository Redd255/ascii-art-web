package main

import (
	"asciiart/serv"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", serv.Index)
	log.Println("http//:localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
