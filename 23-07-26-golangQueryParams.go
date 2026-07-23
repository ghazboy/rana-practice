package main

import (
	"fmt"
	"net/http"
)

func greetUser(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprintln(w, "Hello, stranger! Try adding ?name=YourName to the URL")
	} else {
		fmt.Fprintln(w, "Hello, "+name+"!")
	}
}

func main() {
	http.HandleFunc("/greet", greetUser)
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}