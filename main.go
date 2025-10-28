package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
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
	Roll_no int
}

func UpdateName(s *Student, newName string) {
	s.Name = newName
}
func main() {

	student := &Student{Name: "Divakar", Age: 23, Roll_no: 143}
	UpdateName(student, "Kirti")
	fmt.Println("Student Name:", student.Name)

	log.Printf("Hello %s Did you logged in?", "Divakar,")
	//log.Fatalf("Hello something went wrong")
	log.Println("Hello Divakar")
	log.Println("Hello Divakarb7")
	log.Println("Hello Divakarb7")
	log.Println("Hello Divakarb7")
	log.Println("Hello Divakar8")

	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
