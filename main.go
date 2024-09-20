package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	case "/ascii-art":
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading request body:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Fprint(w, string(body))
	default:
		fmt.Fprintf(w, "Sorry, no handler for path %q", r.URL.Path)
	}
}
func main() {
	http.HandleFunc("/", homePage)
	log.Println("Server running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
