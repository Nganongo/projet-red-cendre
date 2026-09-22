package character

import "fmt"

type Personnage struct {
	Nom        string
	Classe     string
	Niveau     int
	VieMax     int
	VieAct     int
	Inventaire []string
	Argent     int
}

func InitCharacter(nom string) *Personnage {
	return &Personnage{
		Nom:        nom,
		Classe:     "Mutant",
		Niveau:     1,
		VieMax:     100,
		VieAct:     50,
		Inventaire: []string{"Kit médical", "Kit médical", "Kit médical"},
		Argent:     100,
	}
}

func DisplayInfo(s *Personnage) {
	fmt.Println("======FICHE DU PERSONNAGE======")
	fmt.Printf("Nom:      %s\n", s.Nom)
	fmt.Printf("Classe:      %s\n", s.Classe)
	fmt.Printf("Niveau:      %d\n", s.Niveau)
	fmt.Printf("PV:      %d/%d\n", s.VieAct, s.VieMax)
	fmt.Printf("Argent:    %d\n", s.Argent)
}
