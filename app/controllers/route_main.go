package controllers

import (
	"net/http"
	"text/template"
)

func Top(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("github.com/soutaschool/todo-app/views/pages/index.tsx")
	t.Execute(w, nil)
}
