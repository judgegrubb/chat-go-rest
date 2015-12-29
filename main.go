package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Message struct {
	Id      int       `json:"id"`
	Message string    `json:"message"`
	Author  string    `json:"author"`
	Created time.Time `json:"created"`
}

type Messages []Message

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/messages", messages)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellow, %q", html.EscapeString(r.URL.Path))
}

func messages(w http.ResponseWriter, r *http.Request) {
	messages := Messages{
		Message{Id: 1, Message: "Hello!", Author: "fred", Created: time.Now()},
		Message{Id: 2, Message: "What's up?", Author: "tEd", Created: time.Now()},
	}

	json.NewEncoder(w).Encode(messages)
}
