package models

type Article struct {
    ID          int
    Nom         string
    Description string
    Image       string
    Prix        float64
    Stock       int
    Reduction   *int
}

var Articles = []Article{
    {ID: 1, Nom: "PALACE WASHED TERRY 1/4 PLACKET HOOD MOJITO", Description: "", Image: "16A.webp", Prix: 89.99, Stock: 300},
    {ID: 2, Nom: "PALACE LONDON CAPUSHE OVERSIZE NOIR", Description: "", Image: "18A.webp", Prix: 69.99, Stock: 450, Reduction: new(int)},
    {ID: 3, Nom: "PLACAE PULL A CAPUCHE UNISEXE CHASSEUR", Description: "", Image: "19A.webp", Prix: 109.99, Stock: 200},
    {ID: 4, Nom: "PALACE PULLCREW CAPUCHON MARINE", Description: "", Image: "21A.webp", Prix: 89.99, Stock: 300},
    {ID: 5, Nom: "PALACE PULL CREW PASSEPORT NOIR", Description: "", Image: "22A.webp", Prix: 129.99, Stock: 150, Reduction: new(int)},
    {ID: 6, Nom: "PALACE PANTALON CARGO GORE-TEX R-TEK NOIR", Description: "", Image: "33B.webp", Prix: 169.99, Stock: 150},
    {ID: 7, Nom: "PALACE PANTALON BOSSY JEAN STONE", Description: "", Image: "34B.webp", Prix: 159.99, Stock: 200},
}
