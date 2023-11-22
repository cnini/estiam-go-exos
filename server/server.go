package server

import (
	"estiam-go-exos/route"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/list", route.ListHandler).Methods("GET")
	router.HandleFunc("/add", route.AddHandler).Methods("POST")
	router.HandleFunc("/get", route.GetHandler).Methods("GET")
	router.HandleFunc("/remove", route.RemoveHandler).Methods("DELETE")

	http.Handle("/", router)
}

func ListenAndServe() {
	http.ListenAndServe(":8080", nil)
}
