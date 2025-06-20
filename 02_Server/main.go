package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	// "github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	router.HandleFunc("/books", GetAllBooks).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

	g := gin.Default()
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := []string{"Atomic Habits", "The Power of Now", "The 7 Habits of Highly Effective People"}
	w.Header().Set("Content-Type", "application.json")
	bookEncoded, err := (json.Marshal(books))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bookEncoded)

}
