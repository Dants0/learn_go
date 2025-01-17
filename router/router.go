package router

import (
	"fmt"
	"io"
	"net/http"

)

func GetRoute(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!")
	fmt.Println("got / routes")
}

func Server() {
	http.HandleFunc("/hello", GetRoute)
	http.ListenAndServe(":8080", nil)
}
