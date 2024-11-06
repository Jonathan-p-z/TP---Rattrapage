package main

import (
    "net/http"
    "TP---Rattrapage/internal/handlers"
)

func main() {
    http.HandleFunc("/", handlers.ArticleListHandler)
    http.HandleFunc("/article/", handlers.ArticleViewHandler)
    http.HandleFunc("/ajouter", handlers.ArticleAddHandler)
    http.Handle("/web/static/", http.StripPrefix("/web/static/", http.FileServer(http.Dir("web/static"))))

    http.ListenAndServe(":8080", nil)
}
