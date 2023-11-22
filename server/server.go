package server

import (
	"estiam-go-exos/dictionary"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/list", ListHandler).Methods("GET")
	router.HandleFunc("/add", AddHandler).Methods("POST")
	router.HandleFunc("/get", GetHandler).Methods("GET")
	router.HandleFunc("/remove", RemoveHandler).Methods("DELETE")

	http.Handle("/", router)
}

func ListenAndServe() {
	http.ListenAndServe(":8080", nil)
}

func ListHandler(writer http.ResponseWriter, request *http.Request) {
	file, openErr := os.OpenFile("dictionary.txt", os.O_RDWR|os.O_APPEND, 0777)
	if openErr != nil {
		http.Error(writer, "Error opening file", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	keys, values := dictionary.List(file)

	fmt.Fprintln(writer, "keys:", keys)
	fmt.Fprintln(writer, "values:", values)
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

func GetHandler(writer http.ResponseWriter, request *http.Request) {
	word := request.URL.Query().Get("q")

	file, openErr := os.OpenFile("dictionary.txt", os.O_RDWR|os.O_APPEND, 0777)
	if openErr != nil {
		http.Error(writer, "Error opening file", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	wordFound, translationFound, getErr := dictionary.Get(file, word)
	dictionary.Check(getErr)

	fmt.Fprintf(writer, "%s:%s", wordFound, translationFound)
}

func RemoveHandler(writer http.ResponseWriter, request *http.Request) {
	word := request.URL.Query().Get("q")

	file, openErr := os.OpenFile("dictionary.txt", os.O_RDWR|os.O_APPEND, 0777)
	if openErr != nil {
		http.Error(writer, "Error opening file", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	wordToRemove, translationToRemove, getErr := dictionary.Get(file, word)
	dictionary.Check(getErr)

	_, removeErr := dictionary.Remove(file, word)
	dictionary.Check(removeErr)

	fmt.Fprintf(writer, "remove %s:%s", wordToRemove, translationToRemove)
}
