package main

import (
	"fmt"
	"log"
	"net/http"
)

const Port = 8080

func main() {
	http.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s", r.PathValue("name"))
	})

	fmt.Println("Server listening on port", Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), nil))
}
