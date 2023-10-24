package main

import (
	"fmt"
	"net/http"
)

const PORT = 8080

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	fmt.Fprintf(w, "Hello %s", name)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello")

}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting at %d", PORT)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Print("Yo error")
	}

}
