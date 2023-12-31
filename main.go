package main

import (
	"fmt"
	"log"
	"net/http"
)

func newsHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm err: %v", err)
		return
	}
	fmt.Fprintf(w, "Added successfully \n")
	name := r.FormValue("name")
	address := r.FormValue("email")
	fmt.Fprintf(w, "Name= %s\n", name)
	fmt.Fprintf(w, "Email= %s\n", address)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/newsletter", newsHandler)
	http.HandleFunc("/about", aboutHandler)
	fmt.Print("Starting server at port 8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
