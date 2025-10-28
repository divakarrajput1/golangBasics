package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Divakar! this is bracnh b4")
	fmt.Fprintf(w, "Hello, Divakar! this is bracnh b4")
	fmt.Fprintf(w, "Hello, Divakar! this is bracnh b4")
	fmt.Fprintf(w, "Hello, Divakar! this is bracnh b4")
	fmt.Fprintf(w, "Hello, Divakar! this is bracnh b4")
	fmt.Fprintf(w, "Hello, Divakar! this is bracnh b6")
	fmt.Fprintf(w, "Hello, Divakar! this is bracnh b6")
	fmt.Fprintf(w, "Hello, Divakar! this is bracnhb5")
	fmt.Fprintf(w, "Hello, Divakar! this is bracnhb5")
	fmt.Fprintf(w, "Hello, Divakar! this is bracnhb5")
	fmt.Fprintf(w, "Hello, Divakar! this is bracnhb5")
	fmt.Fprintf(w, "Hello, Divakar! this is bracnhb8")
	fmt.Fprintf(w, "Hello, Divakar! this is bracnhb8")
}

type Student struct {
	Name    string
	Age     int
	Roll_No int
}

func UpdateName(s *Student, newName string) {
	s.Name = newName
}
func main() {

	fmt.Println("Done")

	log.Printf("Hello %s Did you logged in?", "Divakar,")

	log.Println("Hello Divakar")
	log.Println("Hello Divakarb7")
	log.Println("Hello Divakarb7")
	log.Println("Hello Divakarb7")
	log.Println("Hello Divakar8")

	student := &Student{Name: "Kirti", Age: 21, Roll_No: 143}
	UpdateName(student, "Divakar")
	fmt.Println("Student Name:", student.Name)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
