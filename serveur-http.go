package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Server Starting")
	//http.HandleFunc("/", Inscription)
	http.HandleFunc("/index", serveFiles)
	http.ListenAndServe(":80", nil)

}

func serveFiles(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "index.html")
}

//func serveFile(w http.ResponseWriter, r *http.Request) {
//	http.ServeFile(w, r, p)

//func Inscription(w http.ResponseWriter, r *http.Request) {
