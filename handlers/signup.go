package handlers

import (
	"html/template"
	"net/http"
	"path"
)

type Signup struct {
	template *template.Template
}

func NewSignup() *Signup {
	basePath := path.Join("templates", "_base.html")
	templatePath := path.Join("templates", "signup.html")
	t := *template.Must(template.ParseFiles(templatePath, basePath))

	return &Signup{&t}
}

func (l *Signup) GetSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	data := struct{ UploadKey string }{fileUploadKey}

	if err := l.template.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}
