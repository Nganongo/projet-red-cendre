package character

import (
	"fmt"
)

func IsDead(s *Personnage) {
	if s.VieAct <= 0 {
		fmt.Println("Vous etes mort")
		s.VieAct = s.VieMax / 2
	}
}
