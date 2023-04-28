package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/api/hello", handleHello)

	appengine.Main()
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the main page!")
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, from the API!")
}
