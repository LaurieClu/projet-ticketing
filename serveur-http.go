package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", Inscription)
    http.ListenAndServe(":80", nil)
}

func Inscription(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Inscription, %s!", r.URL.Path[1:])
}