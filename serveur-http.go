package main

import (
	"fmt"
	"html/template"
	"net/http"
	"io"
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
	v := req.FormValue("q")
w.Header().Set("Content-Type", "text/html; charset=utf-8")
io.WriteString(w,
<form method="POST">
<input type="text" name="q">
<input type="envoyer">
</form>
<br>`+v)
}
