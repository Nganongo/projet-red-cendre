package character

import (
	"bufio"
	"fmt"
)

func TrainingFight(perso *Personnage, reader *bufio.Reader) {
	monstre := InitChienCendre()
	fmt.Print(ChienCendreArt)
	tour := 1

	fmt.Printf("Initiative : %s = %d, %s = %d\n", perso.Nom, perso.Initiative, monstre.Nom, monstre.Initiative)

	if monstre.Initiative > perso.Initiative {
		fmt.Println(monstre.Nom, "est plus rapide et attaque en premier !")

		GoblinPattern(monstre, perso, tour)
		tour++
	} else {
		fmt.Println(perso.Nom, "est plus rapide et commence !")
	}

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
		fmt.Print(DefaiteArt)
		fmt.Println("Vous avez été vaincu...")
	} else {
		fmt.Print(VictoireArt)
		fmt.Println("Vous avez vaincu le", monstre.Nom, "!")
		perso.XpAct += monstre.Xpdonnee
		fmt.Printf("Vous gagnez %d XP ! (%d/%d)\n", monstre.Xpdonnee, perso.XpAct, perso.XpMax)

		if perso.XpAct >= perso.XpMax {
			excedent := perso.XpAct - perso.XpMax
			perso.Niveau++
			perso.VieMax += 20
			perso.XpMax += 20
			perso.XpAct = excedent
			fmt.Printf("Niveau supérieur ! %s passe niveau %d (PV max : %d)\n", perso.Nom, perso.Niveau, perso.VieMax)
		}
	}
}
