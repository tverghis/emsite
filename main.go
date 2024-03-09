package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/tverghis/emsite/handlers"
	"github.com/tverghis/emsite/util/files"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const Port = 8080

func main() {
	if err := runDBMigrations(); err != nil {
		panic(err)
	}

	if err := files.EnsureUploadsDir(); err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("uploads"))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", fs))

	uploadHandler := handlers.NewUpload()
	http.HandleFunc("GET /upload", uploadHandler.GetUpload)
	http.HandleFunc("POST /upload", uploadHandler.PostUpload)

	downloadHandler := handlers.NewDownload()
	http.HandleFunc("GET /download", downloadHandler.GetDownload)

	galleryHandler := handlers.NewGallery()
	http.HandleFunc("GET /gallery", galleryHandler.GetGallery)
	http.HandleFunc("DELETE /gallery/{filename}", galleryHandler.DeleteGallery)

	loginHandler := handlers.NewLogin()
	http.HandleFunc("GET /login", loginHandler.GetLogin)

	fmt.Println("Server listening on port", Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), nil))
}

func runDBMigrations() error {
	m, err := migrate.New("file://db/migrations", "sqlite3://db/database.db")

	if err != nil {
		return err
	}

	if err := m.Up(); !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	fmt.Println("Successfully ran DB migrations")

	return nil
}
