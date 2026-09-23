package character

import (
	"bufio"
	"fmt"
	"strings"
)

func AccessEquipement(s *Personnage, reader *bufio.Reader) {
	for {
		fmt.Println("\n====== ÉQUIPEMENT ACTUEL ======")
		fmt.Printf("Tête  : %s\n", s.Equipement.Tete)
		fmt.Printf("Torse : %s\n", s.Equipement.Torse)
		fmt.Printf("Pieds : %s\n", s.Equipement.Pieds)
		fmt.Println("--------------------------------")
		fmt.Println("1. Équiper un objet depuis l'inventaire")
		fmt.Println("2. Retour")
		fmt.Print("Votre choix : ")

		choix, _ := reader.ReadString('\n')
		choix = strings.TrimSpace(choix)

		if choix == "1" {
			EquiperObjet(s, reader)
		} else if choix == "2" {
			return
		} else {
			fmt.Println("Choix invalide.")
		}
	}
}

func EquiperObjet(s *Personnage, reader *bufio.Reader) {
	if len(s.Inventaire) == 0 {
		fmt.Println("❌ Votre inventaire est vide !")
		return
	}

	fmt.Println("\nObjets équipables dans votre inventaire :")
	for i, item := range s.Inventaire {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Print("Entrez le nom EXACT de l'objet à équiper : ")

	nomObjet, _ := reader.ReadString('\n')
	nomObjet = strings.TrimSpace(nomObjet)

	// Vérification si l'objet est bien dans l'inventaire
	trouve := false
	for _, item := range s.Inventaire {
		if item == nomObjet {
			trouve = true
			break
		}
	}

	if !trouve {
		fmt.Println("❌ Vous ne possédez pas cet objet.")
		return
	}

	// Logique d'équipement selon l'objet
	switch nomObjet {
	case "Chapeau de l'aventurier":
		// Si un objet était déjà équipé à la Tête, on le réintègre dans l'inventaire
		if s.Equipement.Tete != "" {
			AddInventaire(s, s.Equipement.Tete)
		}
		s.Equipement.Tete = nomObjet
		s.VieMax += 10
		s.VieAct += 10
		RetirerItem(s, nomObjet, 1)
		fmt.Println("🎩 Chapeau de l'aventurier équipé ! (+10 PV Max)")

	case "Tunique de cuir":
		if s.Equipement.Torse != "" {
			AddInventaire(s, s.Equipement.Torse)
		}
		s.Equipement.Torse = nomObjet
		s.VieMax += 25
		s.VieAct += 25
		RetirerItem(s, nomObjet, 1)
		fmt.Println("🥋 Tunique de cuir équipée ! (+25 PV Max)")

	case "Bottes renforcées":
		if s.Equipement.Pieds != "" {
			AddInventaire(s, s.Equipement.Pieds)
		}
		s.Equipement.Pieds = nomObjet
		s.VieMax += 15
		s.VieAct += 15
		RetirerItem(s, nomObjet, 1)
		fmt.Println("🥾 Bottes renforcées équipées ! (+15 PV Max)")

	default:
		fmt.Println("❌ Cet objet ne peut pas être équipé.")
	}
}
