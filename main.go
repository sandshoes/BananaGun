package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloGoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([] byte("Hello net/http\n"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Tommy\n"))
}

func main() {
	fmt.Printf("Starting server at port 8080\n")

	http.HandleFunc("/", helloGoHandler)
	http.HandleFunc("about", aboutHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}


