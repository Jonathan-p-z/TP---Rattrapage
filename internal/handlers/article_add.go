package handlers

import (
	"TP---Rattrapage/internal/models"
	"html/template"
	"net/http"
	"strconv"
)

func ArticleAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("web/templates/base.html", "web/templates/article_add.html"))
		data := struct {
			Title string
		}{
			Title: "Ajouter un Produit",
		}
		tmpl.ExecuteTemplate(w, "base", data)
	} else if r.Method == http.MethodPost {
		nom := r.FormValue("nom")
		description := r.FormValue("description")
		prix, _ := strconv.ParseFloat(r.FormValue("prix"), 64)
		stock, _ := strconv.Atoi(r.FormValue("stock"))
		reduction, _ := strconv.Atoi(r.FormValue("reduction"))

		var newReduction *int
		if reduction > 0 {
			newReduction = &reduction
		}

		newArticle := models.Article{
			ID:          len(models.Articles) + 1,
			Nom:         nom,
			Description: description,
			Image:       "default.jpg",
			Prix:        prix,
			Stock:       stock,
			Reduction:   newReduction,
		}

		models.Articles = append(models.Articles, newArticle)
		http.Redirect(w, r, "/article/"+strconv.Itoa(newArticle.ID), http.StatusSeeOther)
	}
}
