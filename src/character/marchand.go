package character

import (
	"bufio"
	"fmt"
	"strings"
)

func AddInventaire(s *Personnage, item string) {
	s.Inventaire = append(s.Inventaire, item)
	fmt.Println("Vous avez obtenu :", item)
}

func Marchand(s *Personnage, reader *bufio.Reader) {
	for {
		fmt.Printf("\n====== MARCHAND ====== (Votre argent : %d Or)\n", s.Argent)
		fmt.Println("1. Kit médical (3 Or)")
		fmt.Println("2. Sérum toxique (6 Or)")
		fmt.Println("3. Livre de Sort: Boule de Feu (25 Or)")
		fmt.Println("4. Fourrure de Loup (4 Or)")
		fmt.Println("5. Peau de Troll (7 Or)")
		fmt.Println("6. Cuir de Sanglier (3 Or)")
		fmt.Println("7. Plume de Corbeau (1 Or)")
		fmt.Println("8. Retour")
		fmt.Print("Votre choix : ")

		choix, _ := reader.ReadString('\n')
		choix = strings.TrimSpace(choix)

		var nomItem string
		var prix int

		switch choix {
		case "1":
			nomItem = "Kit médical"
			prix = 3
		case "2":
			nomItem = "Sérum toxique"
			prix = 6
		case "3":
			nomItem = "Livre de Sort: Boule de Feu"
			prix = 25
		case "4":
			nomItem = "Fourrure de Loup"
			prix = 4
		case "5":
			nomItem = "Peau de Troll"
			prix = 7
		case "6":
			nomItem = "Cuir de Sanglier"
			prix = 3
		case "7":
			nomItem = "Plume de Corbeau"
			prix = 1
		case "8":
			return
		default:
			fmt.Println("Choix invalide.")
			continue
		}

		if !PeutAjouterInventaire(s) {
			fmt.Println("❌ Achat impossible : votre inventaire est plein (max 10 objets) !")
			continue
		}

		if s.Argent < prix {
			fmt.Printf("❌ Achat impossible : vous n'avez pas assez d'or (%d Or requis) !\n", prix)
			continue
		}

		s.Argent -= prix

		if nomItem == "Livre de Sort: Boule de Feu" {
			spellBook(s, "Boule de Feu")
		} else {
			AddInventaire(s, nomItem)
		}

		fmt.Printf(" Vous avez acheté : %s pour %d Or.\n", nomItem, prix)
	}
}
