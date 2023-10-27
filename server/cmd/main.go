package main

import (
	"QN/controller"
	"QN/utils"
	"net/http"
)

func main() {
	s := controller.New()
	utils.GetToken()
	http.ListenAndServe(":8080", s)
}
