package main

import (
	"articleselection/internal/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.ArticleListHandler)
	http.HandleFunc("/article/", handlers.ArticleViewHandler)
	http.HandleFunc("/ajouter", handlers.ArticleAddHandler)
	http.Handle("/web/static/", http.StripPrefix("/web/static/", http.FileServer(http.Dir("web/static"))))

	http.ListenAndServe("Localhost:8080", nil)
}
