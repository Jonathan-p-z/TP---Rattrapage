package models

type Article struct {
    ID          int
    Nom         string
    Description string
    Image       string
    Prix        float64
    Stock       int
    Reduction   *int // Pointeur pour gérer l'optionnel
}

var Articles = []Article{
    {ID: 1, Nom: "Produit 1", Description: "Description du produit 1", Image: "product1.jpg", Prix: 19.99, Stock: 10},
    {ID: 2, Nom: "Produit 2", Description: "Description du produit 2", Image: "product2.jpg", Prix: 29.99, Stock: 5, Reduction: new(int)},
    // Ajoutez d'autres articles si nécessaire
}
