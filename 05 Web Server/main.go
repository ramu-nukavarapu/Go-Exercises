package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Gophers!")
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {

	fmt.Println("Server started...")
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/greet", greet)
	http.ListenAndServe(":8080", nil)
}
