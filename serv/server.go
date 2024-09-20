package serv

import (
	"fmt"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("banner")
	switch r.URL.Path {
	case "/":
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	case "/ascii-art":
		//body, err := io.ReadAll(r.Body)
		// if err != nil {
		// 	log.Println("Error reading request body:", err)
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }		
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		dynamicText := text

		// Execute the template, inserting dynamicText into {{.}}
		if err := tmpl.Execute(w, dynamicText); err != nil {
			http.Error(w, err.Error(), 404)
		}
	default:
		fmt.Fprintf(w, "Sorry, no handler for path %q", r.URL.Path)
	}
}
