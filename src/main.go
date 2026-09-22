package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"projet-red/src/character"
)

func main() {
	monper := character.InitCharacter("jordan")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n ==== MENU ====")
		fmt.Println("1. Afficher les informations")
		fmt.Println("2. Acceder a l'inventaire")
		fmt.Println("3. Quitter")
		fmt.Print(" Votre choix: ")

		choix, _ := reader.ReadString('\n')
		choix = strings.TrimSpace(choix)

		switch choix {
		case "1":
			character.DisplayInfo(monper)
		case "2":
			character.AccessInventory(monper)
		case "3":
			fmt.Println("A Bientot !")
			return
		default:
			fmt.Println(" choix Invalide")
		}

	}
}
