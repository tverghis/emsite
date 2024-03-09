package handlers

import (
	"html/template"
	"net/http"
	"path"
)

type Login struct {
	template *template.Template
}

func NewLogin() *Login {
	basePath := path.Join("templates", "_base.html")
	templatePath := path.Join("templates", "login.html")
	t := *template.Must(template.ParseFiles(templatePath, basePath))

	return &Login{&t}
}

func (l *Login) GetLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	data := struct{ UploadKey string }{fileUploadKey}

	if err := l.template.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}
