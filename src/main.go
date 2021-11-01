package main

import (
	"net/http"
	"web-go/src/controller"
)

func main() {
	http.HandleFunc("/", controller.Register)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/list", controller.ListDorms)
	http.ListenAndServe(":8080", nil)
}
