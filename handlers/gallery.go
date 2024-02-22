package handlers

import (
	"html/template"
	"net/http"
	"path"

	"github.com/tverghis/emsite/util/files"
)

type Gallery struct {
	template template.Template
}

func NewGallery() *Gallery {
	basePath := path.Join("templates", "_base.html")
	templatePath := path.Join("templates", "gallery.html")
	t := *template.Must(template.ParseFiles(templatePath, basePath))

	return &Gallery{t}
}

func (h *Gallery) GetGallery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	filenames, err := files.GetUploadedFilePaths()

	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}

	data := struct{ Filenames []string }{filenames}

	if err := h.template.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}

func (h *Gallery) DeleteGallery(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("filename")

	if err := files.DeleteFile(name); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
