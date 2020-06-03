package main

import (
	"fmt"
	"net/http"
)

func main() {
	printAllIDAllName()
	fmt.Println("Server Starting")
	http.HandleFunc("/index", serveFiles)
	http.ListenAndServe(":80", nil)
}

func serveFiles(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "index.html")

	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(response, "Oups...retentez votre chance ! %v", err)
		return
	}
	fmt.Fprintf(response, "Post from website! r.PostFrom = %v\n", request.PostForm)

	fmt.Printf("%v\n%v\n%v\n%v\n", request.FormValue("email"), request.FormValue("password"), request.FormValue("surname"), request.FormValue("name"))
	newUsers(request.FormValue("email"), request.FormValue("password"), request.FormValue("surname"), request.FormValue("name"))
}
