package character

import "fmt"

func GoblinPattern(monstre *Monstre, cible *Personnage, tour int) {
	degats := monstre.Attaque

	if tour%3 == 0 {
		degats *= 2
	}
	cible.VieAct -= degats

	if cible.VieAct < 0 {
		cible.VieAct = 0
	}

	fmt.Printf("%s attaque %s et inflige %d degats \n", monstre.Nom, cible.Nom, degats)
	fmt.Printf("PV de %s :  %d/%d\n", cible.Nom, cible.VieAct, cible.VieMax)
}
