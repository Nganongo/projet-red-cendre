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
		fmt.Println("\n=== VOTRE INVENTAIRE ===")
		if len(s.Inventaire) == 0 {
			fmt.Println("(Inventaire vide)")
		} else {
			for i, item := range s.Inventaire {
				fmt.Printf("%d. %s\n", i+1, item)
			}
		}
		fmt.Println("---------------------------------")
		fmt.Println("1. Équiper un objet depuis l'inventaire")
		fmt.Println("2. Retour")
		fmt.Print("Votre choix : ")

		choix, _ := reader.ReadString('\n')
		choix = strings.TrimSpace(choix)

		switch choix {
		case "1":
			EquiperObjet(s, reader)
		case "2":
			return
		default:
			fmt.Println("❌ Choix invalide !")
		}
	}
}

func EquiperObjet(s *Personnage, reader *bufio.Reader) {
	if len(s.Inventaire) == 0 {
		fmt.Println("❌ Votre inventaire est vide !")
		return
	}

	fmt.Println("\nObjets dans votre inventaire :")
	for i, item := range s.Inventaire {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	fmt.Print("Entrez le NUMÉRO de l'objet à équiper : ")
	choixObjet, _ := reader.ReadString('\n')
	choixObjet = strings.TrimSpace(choixObjet)

	var nomObjet string

	switch choixObjet {
	case "1":
		if len(s.Inventaire) >= 1 {
			nomObjet = s.Inventaire[0]
		}
	case "2":
		if len(s.Inventaire) >= 2 {
			nomObjet = s.Inventaire[1]
		}
	case "3":
		if len(s.Inventaire) >= 3 {
			nomObjet = s.Inventaire[2]
		}
	case "4":
		if len(s.Inventaire) >= 4 {
			nomObjet = s.Inventaire[3]
		}
	case "5":
		if len(s.Inventaire) >= 5 {
			nomObjet = s.Inventaire[4]
		}
	case "6":
		if len(s.Inventaire) >= 6 {
			nomObjet = s.Inventaire[5]
		}
	case "7":
		if len(s.Inventaire) >= 7 {
			nomObjet = s.Inventaire[6]
		}
	case "8":
		if len(s.Inventaire) >= 8 {
			nomObjet = s.Inventaire[7]
		}
	case "9":
		if len(s.Inventaire) >= 9 {
			nomObjet = s.Inventaire[8]
		}
	case "10":
		if len(s.Inventaire) >= 10 {
			nomObjet = s.Inventaire[9]
		}
	default:
		fmt.Println("❌ Numéro invalide !")
		return
	}

	if nomObjet == "" {
		fmt.Println("❌ Cet emplacement n'existe pas !")
		return
	}

	switch nomObjet {
	case "Chapeau de l'aventurier":
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
		fmt.Printf("❌ %s ne peut pas être équipé !\n", nomObjet)
	}
}
