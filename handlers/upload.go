package handlers

import (
	"html/template"
	"net/http"
	"path"
)

type UploadHandler struct {
	template template.Template
}

func NewUploadHandler() *UploadHandler {
	templatePath := path.Join("templates", "upload.html")
	t := *template.Must(template.ParseFiles(templatePath))

	return &UploadHandler{t}
}

func (h *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.template.Execute(w, nil); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}
