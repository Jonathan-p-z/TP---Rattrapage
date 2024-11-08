package handlers

import (
	"articleselection/internal/models"
	"html/template"
	"log"
	"net/http"
)

type test struct {
	Title    string
	Articles []models.Article
}

func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../web/templates/base.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	data := test{
		Title:    "Liste des Articles",
		Articles: models.Articles,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}
