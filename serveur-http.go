package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Server Starting")
	http.HandleFunc("/index", serveFiles)
	http.ListenAndServe(":80", nil)
}

func serveFiles(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "index.html")

	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(response, "Oups...tentez votre chance ! %v", err)
		return
	}
	fmt.Fprintf(response, "Post from website! r.PostFrom = %v\n", request.PostForm)

	fmt.Println("age :", request.FormValue("Age"))
	fmt.Println("nom :", request.FormValue("nom"))
	fmt.Println("Prénom :", request.FormValue("prenom"))
	fmt.Println("user-password :", request.FormValue("user-password"))
	fmt.Println("Email:", request.FormValue("email"))
	//fmt.Println("Email:", request.Form.Get("email"))

}
