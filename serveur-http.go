package main

import (
	"fmt"
	"net/http"
)

func serveFiles(w http.ResponseWriter, r *http.Request) {
	p := "." + (r.URL.Path)
	if p == "./" {
		p = "./static.index.html"
	}
	http.ServeFile(w, r, p)
}
func main() {
	http.HandleFunc("/", Inscription)
	fmt.Println("Server Starting")
	http.ListenAndServe(":80", nil)
	http.HandleFunc("/index.html", serveFiles)
}

func Inscription(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Inscription, %s!", r.URL.Path[1:])
}
