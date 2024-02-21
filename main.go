package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tverghis/emsite/handlers"
	"github.com/tverghis/emsite/util/files"
)

const Port = 8080

func main() {
	if err := files.EnsureUploadsDir(); err != nil {
		panic(err)
	}

	uploadHandler := handlers.NewUpload()
	http.HandleFunc("GET /upload", uploadHandler.GetUpload)
	http.HandleFunc("POST /upload", uploadHandler.PostUpload)

	downloadHandler := handlers.NewDownload()
	http.HandleFunc("GET /download", downloadHandler.GetDownload)

	fmt.Println("Server listening on port", Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), nil))
}
