package main

import (
	"net/http"
	"web-go/controller"
)

func main() {
	http.HandleFunc("/", controller.Register)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/list", controller.ListDorms)
	http.ListenAndServe(":10703", nil)
}