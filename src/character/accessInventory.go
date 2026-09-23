package character

import "fmt"

func AccessInventory(s *Personnage) {
	if len(s.Inventaire) == 0 {
		fmt.Println("Votre inventaire est vide")
		return
	}

	fmt.Println("======INVENTAIRE======")
	for i, item := range s.Inventaire {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}

func TakePot(s *Personnage) {
	index := -1
	for i, item := range s.Inventaire {
		if item == "kit medical" {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("vous n'avez pas de kit medical")
		return
	}

	s.Inventaire = append(s.Inventaire[:index], s.Inventaire[index+1:]...)

	s.VieAct += 50
	if s.VieAct > s.VieMax {
		s.VieAct = s.VieMax
	}

	fmt.Println("Vous utiliser un kit medical")
	fmt.Printf("PV %d/%d", s.VieAct, s.VieMax)
}
