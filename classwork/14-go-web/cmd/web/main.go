// Filename: cmd/web/main.go

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from newsletter"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About newsletter"))
}

func messages(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"test":  "hello",
		"const": "apple",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	// mux is our router (multiplexer)
	mux := http.NewServeMux()

	// the route patter/endpoint/URL path
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/messages", messages)

	// start a local web server
	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
