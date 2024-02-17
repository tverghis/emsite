package handlers

import (
	"html/template"
	"net/http"
	"path"
)

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
	if err := h.template.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}

func (h *Upload) PostUpload(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/upload", http.StatusSeeOther)
}
