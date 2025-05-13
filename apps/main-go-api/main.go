package main

import (
	"fmt"
	"net/http"
)

func Hello(name string) string {
	result := "Hello " + name
	return result
}

func main() {
	fmt.Println(Hello("main-go-api"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from main-go-api")
	})

	http.ListenAndServe(":8080", nil)
}
