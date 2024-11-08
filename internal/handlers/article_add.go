package handlers

import (
	"articleselection/internal/models"
	"html/template"
	"net/http"
	"strconv"
)

func ArticleAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("../web/templates/article_add.html")
		if err != nil {
			http.Error(w, "Erreur lors du chargement de la page", http.StatusInternalServerError)
			return
		}
		data := struct {
			Title string
		}{
			Title: "Ajouter un Produit",
		}
		tmpl.Execute(w, data)
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
			Image:       "34B.webp",
			Prix:        prix,
			Stock:       stock,
			Reduction:   newReduction,
		}

		models.Articles = append(models.Articles, newArticle)
		http.Redirect(w, r, "/article/"+strconv.Itoa(newArticle.ID), http.StatusSeeOther)
	}
}
