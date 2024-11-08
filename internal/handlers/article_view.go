package handlers

import (
    "articleselection/internal/models"
    "html/template"
    "net/http"
    "strconv"
)

func ArticleViewHandler(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Path[len("/article/"):])
    if err != nil {
        http.Error(w, "ID invalide", http.StatusBadRequest)
        return
    }
    var selectedArticle *models.Article
    for _, article := range models.Articles {
        if article.ID == id {
            selectedArticle = &article
            break
        }
    }

    if selectedArticle == nil {
        http.Error(w, "Article non trouvé", http.StatusNotFound)
        return
    }

	tmpl, _ := template.ParseFiles("../web/templates/article_view.html")
    data := struct {
        Title   string
        Article models.Article
    }{
        Title:   selectedArticle.Nom,
        Article: *selectedArticle,
    }
    tmpl.Execute(w, data)
}