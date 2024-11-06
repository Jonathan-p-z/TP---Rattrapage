package handlers

import (
	"TP---Rattrapage/internal/models"
	"html/template"
	"net/http"
)

func ArticleListHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/base.html", "web/templates/article_list.html"))
	data := struct {
		Title    string
		Articles []models.Article
	}{
		Title:    "Liste des Articles",
		Articles: models.Articles,
	}
	tmpl.ExecuteTemplate(w, "base", data)
}
