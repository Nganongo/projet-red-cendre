package character

import (
	"bufio"
	"fmt"
)

func TrainingFight(perso *Personnage, reader *bufio.Reader) {
	monstre := InitChienCendre()
	tour := 1

	for perso.VieAct > 0 && monstre.Pvact > 0 {
		fmt.Printf("\n ------ Tour %d ------ \n", tour)

		CharcterTurn(perso, monstre, reader)
		if monstre.Pvact <= 0 {
			break
		}

		GoblinPattern(monstre, perso, tour)
		tour++

	}
	if perso.VieAct <= 0 {
		fmt.Println("Vous avez été vaincu...")
	} else {
		fmt.Println("Vous avez vaincu le", monstre.Nom, "!")
	}
}
