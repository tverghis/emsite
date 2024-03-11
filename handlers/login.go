package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"path"

	"github.com/tverghis/emsite/models"
)

type Login struct {
	template *template.Template
	db       *sql.DB
}

func NewLogin(db *sql.DB) *Login {
	basePath := path.Join("templates", "_base.html")
	templatePath := path.Join("templates", "login.html")
	t := *template.Must(template.ParseFiles(templatePath, basePath))

	return &Login{&t, db}
}

func (l *Login) GetLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	data := struct{ UploadKey string }{fileUploadKey}

	if err := l.template.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}
}

func (l *Login) PostLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Username and password cannot be blank", http.StatusBadRequest)
		return
	}

	user, err := models.AuthenticateUser(l.db, username, password)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	fmt.Printf("Logged in user %s\n", user.Username)

	w.WriteHeader(200)
}
