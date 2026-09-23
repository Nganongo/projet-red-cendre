package character

import (
	"bufio"
	"fmt"
	"math/rand"
	"strings"
	"unicode"
)

func CharacterCreation(reader *bufio.Reader) *Personnage {
	var nom string
	for {
		fmt.Println("Quel est votre nom ?")
		nom, _ = reader.ReadString('\n')
		nom = strings.TrimSpace(nom)

		valide := nom != ""

		for _, car := range nom {
			if !unicode.IsLetter(car) {
				valide = false
				break
			}
		}
		if valide {
			break
		}
		fmt.Println("Nom invalide : uniquement des lettres, réessayez.")

	}
	nom = strings.ToLower(nom)
	nom = strings.ToUpper(nom[:1]) + nom[1:]

	fmt.Printf("Nom choisi : %s\n", nom)

	fmt.Println("Choisisser votre classe:")
	fmt.Println("1. Survivant (100 PV)")
	fmt.Println("2. mutant (80 PV)")
	fmt.Println("3. Blindé (120 PV)")

	choixClasse, _ := reader.ReadString('\n')
	choixClasse = strings.TrimSpace(choixClasse)

	var classe string
	var viemax int

	switch choixClasse {
	case "1":
		classe = "Humain"
		viemax = 100

	case "2":
		classe = "Mutant"
		viemax = 80

	case "3":
		classe = "Blindé"
		viemax = 120
	default:
		fmt.Println("Choix invalide, Survivant par défaut.")
		classe = "Survivant"
		viemax = 100

	}

	return &Personnage{
		Nom:        nom,
		Classe:     classe,
		Niveau:     1,
		VieMax:     viemax,
		VieAct:     viemax / 2,
		Inventaire: []string{},
		Argent:     100,
		Skills:     []string{"coup de poing"},
		Initiative: rand.Intn(20) + 1,
		XpAct:      0,
		XpMax:      50,
		Mana:       20,
		ManaMax:    20,
	}
}
