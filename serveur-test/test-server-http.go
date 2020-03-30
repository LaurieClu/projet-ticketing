package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", utilisateurs)
	http.Handle("/favicon.ico", http.NotFpundHandler())
	http.ListenAndServe(":8080", nil)
}

func utilisateurs{ w http.ResponseWriter, req *http.Request)
v := req.FormValue("q")
w.Header().Set("Content-Type", "text/html; charset=utf-8")
io.WriteString(w,
<form method="POST">
<input type="text" name="q">
<input type="envoyer">
</form>
<br>`+v)
}