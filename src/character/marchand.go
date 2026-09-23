package character

import (
	"bufio"
	"fmt"
	"strings"
)

func AddInventaire(s *Personnage, item string) {
	s.Inventaire = append(s.Inventaire, item)
	fmt.Println("Vous avez obtenu :", item)
}

func Marchand(s *Personnage, reader *bufio.Reader) {
	fmt.Println("\n===== MARCHAND =====")
	fmt.Println("1. Kit medical (gratuit)")
	fmt.Println("2. Retour")
	fmt.Print("Votre choix : ")

	choix, _ := reader.ReadString('\n')
	choix = strings.TrimSpace(choix)

	switch choix {
	case "1":
		AddInventaire(s, "kit medical")
	case "2":
		return
	default:
		fmt.Println("Choix invalide")
	}
}
