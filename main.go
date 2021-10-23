package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jamsxd/golang-hello-world-service/pkg/application"
	"github.com/jamsxd/golang-hello-world-service/pkg/domain"
)

func main() {

	service := domain.NewBasicHelloWorld()
	handlers := application.NewBasicHelloWorldHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/hello", handlers.GetGretting).Methods("GET")

	http.ListenAndServe("0.0.0.0:8080", r)
}
