package server

import (
	"estiam-go-exos/dictionary"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/add", AddHandler).Methods("POST")

	http.Handle("/", router)
}

func ListenAndServe() {
	http.ListenAndServe(":8080", nil)
}

func AddHandler(writer http.ResponseWriter, request *http.Request) {
	word := request.URL.Query().Get("key")
	translation := request.URL.Query().Get("value")

	file, openErr := os.OpenFile("dictionary.txt", os.O_RDWR|os.O_APPEND, 0777)
	if openErr != nil {
		http.Error(writer, "Error opening file", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	dictionary.Add(file, word, translation)
}
