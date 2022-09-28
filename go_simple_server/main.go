package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error parsing form: %v", err)
	}
	fmt.Fprintf(w, "Post form successful\n")
	name := r.FormValue("name")
	surname := r.FormValue("surname")

	fmt.Fprintf(w, "Hello %s %s", name, surname)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./src"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", handleForm)

	http.HandleFunc("/hello", handleHello)

	log.Println("Listening server on port:8080...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
