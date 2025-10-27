package main

import (
	"fmt"

	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Divakar! this is bracnh b3")
}
func main() {

	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
