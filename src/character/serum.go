package character

import (
	"fmt"
	"time"
)

func PoisonPot(s *Personnage) {
	index := -1
	for i, item := range s.Inventaire {
		if item == "Sérum toxique" {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Vous n'avez pas de Sérum toxique dans votre inventaire.")
		return
	}

	s.Inventaire = append(s.Inventaire[:index], s.Inventaire[index+1:]...)
	fmt.Println("⚠️ Vous consommez un Sérum toxique... Vous vous sentez mal !")

	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)

		s.VieAct -= 10
		if s.VieAct < 0 {
			s.VieAct = 0
		}

		fmt.Printf("Tick %d/3 : Dégâts de poison (-10 PV) | PV : %d/%d\n", i, s.VieAct, s.VieMax)

		if s.VieAct == 0 {
			IsDead(s)
			break
		}
	}
}
