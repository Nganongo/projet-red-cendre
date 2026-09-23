package character

import (
	"fmt"
)

type Equipement struct {
	Tete  string
	Torse string
	Pieds string
}

type Personnage struct {
	Nom        string
	Classe     string
	Niveau     int
	VieMax     int
	VieAct     int
	Inventaire []string
	Argent     int
	Equipement Equipement
	Skills     []string
	Initiative int
	XpAct      int
	XpMax      int
}

func InitCharacter(nom string) *Personnage {
	return &Personnage{
		Nom:        "Paul",
		Classe:     "Mutant",
		Niveau:     1,
		VieMax:     100,
		VieAct:     50,
		Inventaire: []string{"kit medical, kit medical, kit medical"},
		Argent:     100,
		Equipement: Equipement{
			Tete:  "",
			Torse: "",
			Pieds: "",
		},
	}
}

func DisplayInfo(s *Personnage) {
	fmt.Println("======FICHE DU PERSONNAGE======")
	fmt.Printf("Nom:      %s\n", s.Nom)
	fmt.Printf("Classe:      %s\n", s.Classe)
	fmt.Printf("Niveau:      %d\n", s.Niveau)
	fmt.Printf("PV:      %d/%d\n", s.VieAct, s.VieMax)
	fmt.Printf("Argent:    %d\n", s.Argent)
	fmt.Println("=== SORTS MAÎTRISÉS ===")
	if len(s.Skills) == 0 {
		fmt.Println("Aucun sort appris.")
	} else {
		for _, skill := range s.Skills {
			fmt.Printf("- %s\n", skill)
		}
	}
}
