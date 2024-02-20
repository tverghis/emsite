package handlers

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/tverghis/emsite/util/files"
)

const maxFileMemBytes = 8 * 1024 * 1024
const fileUploadKey = "image"

type Upload struct {
	template template.Template
}

func NewUpload() *Upload {
	basePath := path.Join("templates", "_base.html")
	templatePath := path.Join("templates", "upload.html")
	t := *template.Must(template.ParseFiles(templatePath, basePath))

	return &Upload{t}
}

func (h *Upload) GetUpload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	data := struct{ UploadKey string }{fileUploadKey}

	if err := h.template.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}

func (h *Upload) PostUpload(w http.ResponseWriter, r *http.Request) {
	// This ensures that we control the maximum amount of form data stored in memory
	r.ParseMultipartForm(maxFileMemBytes)

	err := writeUploadedFile(r)

	if err != nil {
		fmt.Printf("%s\n", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/upload", http.StatusSeeOther)
}

func writeUploadedFile(r *http.Request) error {
	file, _, err := r.FormFile(fileUploadKey)

	if err != nil {
		return err
	}
	defer file.Close()

	if err := files.SaveUpload(file); err != nil {
		return err
	}

	return nil
}
