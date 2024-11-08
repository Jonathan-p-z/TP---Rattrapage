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
	{ID: 1, Nom: "PALACE WASHED TERRY 1/4 PLACKET HOOD MOJITO", Description: "Le sweat Palace Washed Terry 1/4 Placket Hood Mojito est un modèle décontracté avec une capuche, un plastron 1/4 zippé et une couleur vibrante Mojito. Confortable grâce à son tissu doux, il affiche un style streetwear frais et original.", Image: "16A.webp", Prix: 89.99, Stock: 300},
	{ID: 2, Nom: "PALACE LONDON CAPUSHE OVERSIZE NOIR", Description: "Le sweat à capuche oversize Palace London noir présente une coupe ample et décontractée. Il arbore le logo emblématique Tri-Ferg ou d’autres détails Palace sur la poitrine ou le dos. Fabriqué en coton de qualité, il offre confort et style, parfait pour un look streetwear urbain.", Image: "18A.webp", Prix: 69.99, Stock: 450, Reduction: new(int)},
	{ID: 3, Nom: "PLACAE PULL A CAPUCHE UNISEXE CHASSEUR", Description: "Le pull à capuche unisexe Palace Chasseur offre un design camo unique avec le logo Palace. Confortable et pratique, il dispose d'une capuche ajustable et d'une poche kangourou, parfait pour un look urbain.", Image: "19A.webp", Prix: 109.99, Stock: 200},
	{ID: 4, Nom: "PALACE PULLCREW CAPUCHON MARINE", Description: "Le pull Palace Crew Capuchon marine est un sweat à capuche au style streetwear épuré. Il arbore le logo Palace, souvent brodé ou imprimé sur la poitrine. Fabriqué en coton, il offre confort et chaleur avec une coupe décontractée et une capuche ajustable, idéale pour un look urbain décontracté.", Image: "21A.webp", Prix: 89.99, Stock: 300},
	{ID: 5, Nom: "PALACE PULL CREW PASSEPORT NOIR", Description: "Le pull Palace Crew Passeport noir est un modèle simple et élégant avec un design minimaliste. Il présente le logo Passeport de Palace, souvent brodé sur la poitrine, pour une touche distinctive. Conçu en coton, il offre confort et chaleur avec une coupe décontractée, idéale pour un look streetwear casual.", Image: "22A.webp", Prix: 129.99, Stock: 150, Reduction: new(int)},
	{ID: 6, Nom: "PALACE PANTALON CARGO GORE-TEX R-TEK NOIR", Description: "Le pantalon cargo Palace Gore-Tex R-Tek noir allie fonctionnalité et style streetwear. Imperméable et respirant grâce à la technologie Gore-Tex, il offre confort et chaleur. Avec une coupe ample et des poches latérales pratiques, il est à la fois technique, résistant et stylé.", Image: "33B.webp", Prix: 169.99, Stock: 150},
	{ID: 7, Nom: "PALACE PANTALON BOSSY JEAN STONE", Description: "Le pantalon Palace Bossy Jean Stone est un jean décontracté avec un délavé stone et une coupe classique. Il offre confort et style streetwear, avec des poches pratiques et un logo Palace discret.", Image: "34B.webp", Prix: 159.99, Stock: 200},
}
