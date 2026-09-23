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
		} else if objet != "" {
			fmt.Println("Objet inconnu ou pas encore utilisable ici.")
		}

	default:
		fmt.Println("Choix invalide.")
	}
}
