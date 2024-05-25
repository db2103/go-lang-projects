package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse form error %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")
	fmt.Printf("%v", r)
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s\nAddress: %s", name, address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Herer")
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server starting at port 8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
