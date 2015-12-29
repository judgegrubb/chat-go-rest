package main

import (
    "fmt"
    "html"
    "log"
	"net/http"
    
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hellow, %q", html.EscapeString(r.URL.Path))
}