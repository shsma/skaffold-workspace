package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	_, err := fmt.Fprintf(w, "Hello From APP-A!")
	if err != nil {
		http.Error(w, "Internal Server Problem.", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler) // Update this line of code

	fmt.Printf("Starting app-a at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
