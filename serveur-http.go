package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Nouveaux struct {
	Id   int
	Name string
}

func indexHTMLTemplateVariableHandler(w http.ResponseWriter, r *http.Request) {
	tmplt := template.New("index-templated.html")
	p := Nouveaux{Id: 1, Name: "chavdif"}

	tmplt.Execute(w, p)

}

func main() {
	http.HandleFunc("/", Inscription)
	fmt.Println("Server Starting")
	http.ListenAndServe(":80", nil)
	http.HandleFunc("/index-templated.html", indexHTMLTemplateVariableHandler)
}

func Inscription(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Inscription, %s!", r.URL.Path[1:])
}
