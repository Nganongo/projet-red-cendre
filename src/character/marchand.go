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
		fmt.Println("1. kit medical (3 Or)")
		fmt.Println("2. Serum toxique (6 Or)")
		fmt.Println("3. livre de Sort: boule de feu (25 Or)")
		fmt.Println("4. fourrure de loup (4 Or)")
		fmt.Println("5. peau de troll (7 Or)")
		fmt.Println("6. cuir de sanglier (3 Or)")
		fmt.Println("7. plume de corbeau (1 Or)")
		fmt.Println("8. potion de mana (8 Or)")
		fmt.Println("9. retour")
		fmt.Print("Votre choix : ")

		choix, _ := reader.ReadString('\n')
		choix = strings.TrimSpace(choix)

		var nomItem string
		var prix int

		switch choix {
		case "1":
			nomItem = "kit medical"
			prix = 3
		case "2":
			nomItem = "serum toxique"
			prix = 6
		case "3":
			nomItem = "livre de sort: boule de feu"
			prix = 25
		case "4":
			nomItem = "fourrure de loup"
			prix = 4
		case "5":
			nomItem = "peau de troll"
			prix = 7
		case "6":
			nomItem = "cuir de sanglier"
			prix = 3
		case "7":
			nomItem = "plume de corbeau"
			prix = 1
		case "8":
			nomItem = "potion de mana"
			prix = 8
		case "9":
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

		if nomItem == "livre de sort: boule de feu" {
			spellBook(s, "boule de feu")
		} else {
			AddInventaire(s, nomItem)
		}

		fmt.Printf(" Vous avez acheté : %s pour %d Or.\n", nomItem, prix)
	}
}
