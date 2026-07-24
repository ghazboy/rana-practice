package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintln(w, "Only POST method is allowed here")
		return
	}
	
	var p Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Fprintln(w, "Error reading data:", err)
	return
	}

	fmt.Fprintln(w, "Received Person:", p.Name, p.Age)
}

func main() {
	http.HandleFunc("/person", createPerson)
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}