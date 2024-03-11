package handlers

import (
	"database/sql"
	"fmt"
	"github.com/tverghis/emsite/models"
	"html/template"
	"net/http"
	"path"
)

type Signup struct {
	template *template.Template
	db       *sql.DB
}

func NewSignup(db *sql.DB) *Signup {
	basePath := path.Join("templates", "_base.html")
	templatePath := path.Join("templates", "signup.html")
	t := *template.Must(template.ParseFiles(templatePath, basePath))

	return &Signup{&t, db}
}

func (s *Signup) GetSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	data := struct{ UploadKey string }{fileUploadKey}

	if err := s.template.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}

func (s *Signup) PostSignup(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Username and password cannot be blank", http.StatusBadRequest)
		return
	}

	err := models.AddUser(s.db, username, password)

	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
}
