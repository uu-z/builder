package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":3000"
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello, World")
	})
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
