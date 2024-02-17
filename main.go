package main

import (
	"fmt"
	handlers "github.com/tverghis/emsite/handlers"
	"log"
	"net/http"
)

const Port = 8080

func main() {
	uploadHandler := handlers.NewUploadHandler()
	http.Handle("GET /upload", uploadHandler)

	fmt.Println("Server listening on port", Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), nil))
}
