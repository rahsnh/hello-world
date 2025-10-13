package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello world!")
	w.Write([]byte("Hello world!"))
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started on port 8080...")
	http.ListenAndServe(":8080", nil)
}
