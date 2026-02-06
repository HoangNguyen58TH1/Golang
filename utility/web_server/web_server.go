package web_server

import (
	"fmt"
	"net/http"
)

func Hanlde() {
	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Toni software system")
}
