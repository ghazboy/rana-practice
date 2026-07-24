package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Rana!")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page :)")
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", about)
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}