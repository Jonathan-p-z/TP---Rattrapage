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
    {ID: 1, Nom: "Produit 1", Description: "PALACE WASHED TERRY 1/4 PLACKET HOOD MOJITO", Image: "16A.webp", Prix: 89.99, Stock: 300},
    {ID: 2, Nom: "Produit 2", Description: "PALACE LONDON CAPUSHE OVERSIZE NOIR", Image: "18A.webp", Prix: 69.99, Stock: 450, Reduction: new(int)},
    {ID: 3, Nom: "Product 3", Description: "PLACAE PULL A CAPUCHE UNISEXE CHASSEUR", Image: "19A.webp", Prix: 109.99, Stock: 200},
    {ID: 4, Nom: "Produit 4", Description: "PALACE PULLCREW CAPUCHON MARINE", Image: "21A.webp", Prix: 89.99, Stock: 300},
    {ID: 4, Nom: "Produit 5", Description: "PALACE PULL CREW PASSEPORT NOIR", Image: "22A.webp", Prix: 129.99, Stock: 150, Reduction: new(int)},
    {ID: 4, Nom: "Produit 6", Description: "PALACE PANTALON CARGO GORE-TEX R-TEK NOIR", Image: "33A.webp", Prix: 169.99, Stock: 150},
    {ID: 4, Nom: "Produit 6", Description: "PALACE PANTALON BOSSY JEAN STONE", Image: "34A.webp", Prix: 159.99, Stock: 200},
    // Ajoutez d'autres articles si nécessaire
}
