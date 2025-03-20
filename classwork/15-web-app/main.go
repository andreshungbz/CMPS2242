package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func data(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"name":  "John Doe",
		"email": "john@example.com",
	}

	json, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/data", data)

	log.Print("Server Started @ :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
