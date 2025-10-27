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
}
func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go work(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("Done")

	log.Printf("Hello %s Did you logged in?", "Divakar,")
	//log.Fatalf("Hello something went wrong")
	log.Println("Hello Divakar")
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stop Working")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}
