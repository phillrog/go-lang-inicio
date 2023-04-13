package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome string
	Msg  string
}

func main() {
	templates = template.Must(template.ParseGlob("C:/_dev/repo/Go/src/go-lang-inicio/22 - HTML/index.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"Mundo", "Feito em GO!"}
		templates.ExecuteTemplate(w, "index.html", u)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
