package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", &myHandler{})
	http.ListenAndServe(":8080", nil)
}

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
