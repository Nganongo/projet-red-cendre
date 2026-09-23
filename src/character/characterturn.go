package character

import (
	"bufio"
	"fmt"
	"strings"
)

func CharcterTurn(perso *Personnage, monstre *Monstre, reader *bufio.Reader) {
	fmt.Println("\n==== Tour de", perso.Nom, "====")
	fmt.Println("1. Attaque")
	fmt.Println("2. Inventaire")
	fmt.Println("3. Sorts")
	fmt.Print("Votre choix : ")

	choix, _ := reader.ReadString('\n')
	choix = strings.TrimSpace(choix)

	switch choix {
	case "1":
		degats := 5
		monstre.Pvact -= degats
		if monstre.Pvact < 0 {
			monstre.Pvact = 0
		}
		fmt.Printf("%s utilise Attaque basique et inflige %d dégâts à %s !\n", perso.Nom, degats, monstre.Nom)
		fmt.Printf("PV de %s :  %d/%d\n", monstre.Nom, monstre.Pvact, monstre.Pvmax)

	case "2":
		AccessInventory(perso)
		fmt.Print("Utiliser quel objet ? (nom exacte ou vide pour annuler) : ")
		objet, _ := reader.ReadString('\n')
		objet = strings.TrimSpace(objet)

		if objet == "kit medical" {
			TakePot(perso)
		} else if objet == "potion de mana" {
			TakeManaPot(perso)
		} else if objet != "" {
			fmt.Println("Objet inconnu ou pas encore utilisable ici.")
		}

	case "3":
		if len(perso.Skills) == 0 {
			fmt.Println("Vous ne connaissez aucun sort.")
			break

		}
		fmt.Println("=== Vos sorts ===")
		for i, sort := range perso.Skills {
			fmt.Printf("%d. %s\n", i+1, sort)
		}

		fmt.Print("Quel sort utiliser ? (nom exact) : ")
		sort, _ := reader.ReadString('\n')
		sort = strings.TrimSpace(sort)

		degats := 0
		cout := 0

		if sort == "coup de poing" {
			degats = 8
			cout = 3
		} else if sort == "boule de feu" {
			degats = 18
			cout = 10
		} else {
			fmt.Println("Sort inconnu.")
			break
		}
		if perso.Mana < cout {
			fmt.Println("Pas assez de mana pour lancer ce sort.")
			break
		}

		perso.Mana -= cout
		monstre.Pvact -= degats

		if monstre.Pvact < 0 {
			monstre.Pvact = 0
		}
		fmt.Printf("%s lance %s et inflige %d dégâts à %s ! (Mana : %d/%d)\n", perso.Nom, sort, degats, monstre.Nom, perso.Mana, perso.ManaMax)
		fmt.Printf("PV de %s : %d/%d\n", monstre.Nom, monstre.Pvact, monstre.Pvmax)

	default:
		fmt.Println("Choix invalide.")
	}
}
