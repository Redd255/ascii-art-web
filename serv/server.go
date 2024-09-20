package serv

import (
	asciiart "asciiart/src"
	"html/template"
	"net/http"
	"strings"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func AsciiWeb(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	textInput := asciiart.CheckInput(r.FormValue("text"))
	textLines := strings.Split(textInput, "\r\n")
	banner := r.FormValue("banner")
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		maps, err := asciiart.MapBanner("standard")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		asciiArt := asciiart.Draw(maps, textLines)
		err = tmpl.Execute(w, asciiArt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	maps, err := asciiart.MapBanner(banner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	asciiArt := asciiart.Draw(maps, textLines)
	err = tmpl.Execute(w, asciiArt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
