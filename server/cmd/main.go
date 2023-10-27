package main

import (
	"QN/controller"
	"net/http"
)

type ThisHandler struct{}

func (m *ThisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ThisHandler's ServeHTTP"))
}

func main() {
	s := controller.New()
	http.ListenAndServe(":8080", s)
}
