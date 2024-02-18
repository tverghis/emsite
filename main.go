package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	handlers "github.com/tverghis/emsite/handlers"
)

const Port = 8080

func main() {
	if err := ensureUploadsDir(); err != nil && !errors.Is(err, fs.ErrExist) {
		panic(err)
	}

	uploadHandler := handlers.NewUpload()
	http.HandleFunc("GET /upload", uploadHandler.GetUpload)
	http.HandleFunc("POST /upload", uploadHandler.PostUpload)

	fmt.Println("Server listening on port", Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), nil))
}

func ensureUploadsDir() error {
	return os.Mkdir("uploads", 0755)
}
