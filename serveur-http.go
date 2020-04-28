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

	fmt.Println("Age:", request.FormValue("Age"))
	newBirthday(request.FormValue("Age"))
	fmt.Println("Surname :", request.FormValue("nom"))
	newSurname(request.FormValue("nom"))
	fmt.Println("name:", request.FormValue("prenom"))
	newUsers(request.FormValue("prenom"))
	fmt.Println("user-password :", request.FormValue("user-password"))
	newPassword(request.FormValue("user-password"))
	fmt.Println("email:", request.FormValue("email"))
	newEmail(request.FormValue("email"))
	//fmt.Println("Email:", request.Form.Get("email"))
}
