package character

import (
	"bufio"
	"fmt"
	"strings"
)

func CompterItem(s *Personnage, nomItem string) int {
	compteur := 0
	for _, item := range s.Inventaire {
		if item == nomItem {
			compteur++
		}
	}
	return compteur
}

func RetirerItem(s *Personnage, nomItem string, quantite int) {
	retires := 0
	i := 0
	for i < len(s.Inventaire) && retires < quantite {
		if s.Inventaire[i] == nomItem {
			s.Inventaire = append(s.Inventaire[:i], s.Inventaire[i+1:]...)
			retires++
		} else {
			i++
		}
	}
}

func Forgeron(s *Personnage, reader *bufio.Reader) {
	for {
		fmt.Println("\n====== FORGERON ======")
		fmt.Println("1. Chapeau de l'aventurier (Recette : 1 Plume de Corbeau + 1 Cuir de Sanglier)")
		fmt.Println("2. Tunique de cuir (Recette : 2 Fourrures de Loup + 1 Peau de Troll)")
		fmt.Println("3. Bottes renforcées (Recette : 2 Cuirs de Sanglier)")
		fmt.Println("4. Retour")
		fmt.Print("Votre choix : ")

		choix, _ := reader.ReadString('\n')
		choix = strings.TrimSpace(choix)

		switch choix {
		case "1":
			fabriquerEquipement(s, "Chapeau de l'aventurier", map[string]int{
				"Plume de Corbeau": 1,
				"cuir de sanglier": 1,
			})
		case "2":
			fabriquerEquipement(s, "Tunique de cuir", map[string]int{
				"fourrure de loup": 2,
				"peau de troll":    1,
			})
		case "3":
			fabriquerEquipement(s, "bottes renforcées", map[string]int{
				"cuir de sanglier": 2,
			})
		case "4":
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func fabriquerEquipement(s *Personnage, nomResultat string, recette map[string]int) {
	if !PeutAjouterInventaire(s) {
		fmt.Println("❌ Fabrication impossible : inventaire plein !")
		return
	}

	for composant, quantiteRequise := range recette {
		if CompterItem(s, composant) < quantiteRequise {
			fmt.Printf("❌ Composants insuffisants ! Il vous faut : %d x %s.\n", quantiteRequise, composant)
			return
		}
	}

	for composant, quantiteRequise := range recette {
		RetirerItem(s, composant, quantiteRequise)
	}

	AddInventaire(s, nomResultat)
	fmt.Printf("🔨 Succès ! Vous avez fabriqué : %s !\n", nomResultat)
}
