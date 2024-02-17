package main

import (
	"fmt"
	handlers "github.com/tverghis/emsite/handlers"
	"log"
	"net/http"
)

const Port = 8080

func main() {
	uploadHandler := handlers.NewUpload()
	http.HandleFunc("GET /upload", uploadHandler.GetUpload)
	http.HandleFunc("POST /upload", uploadHandler.PostUpload)

	fmt.Println("Server listening on port", Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), nil))
}
