package main

import (
	"fmt"
	"net/http"
)

func HttpFileHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "index.html")
}
func main() {
	http.HandleFunc("/", Inscription)
	fmt.Println("Server Starting")
	http.ListenAndServe(":80", nil)
	http.HandleFunc("/index.html", HttpFileHandler)
}

func Inscription(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Inscription, %s!", r.URL.Path[1:])
}
