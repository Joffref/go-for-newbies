package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pet struct {
	Name string `json:"name"`
}

var Pet1 = Pet{"Pet1"}

func ListPets(w http.ResponseWriter, r *http.Request) {
	body, err := json.Marshal(Pet1)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(body)
	
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/pets", ListPets)
	http.ListenAndServe(":8080", nil)
}
