package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"projet-red/src/character"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	monper := character.CharacterCreation(reader)

	for {
		fmt.Println("\n ==== MENU ====")
		fmt.Println("1. Afficher les informations")
		fmt.Println("2. Acceder a l'inventaire")
		fmt.Println("3. Acceder au marchand")
		fmt.Println("4. Acceder au forgeron")
		fmt.Println("5. Quitter")
		fmt.Print(" Votre choix: ")

		choix, _ := reader.ReadString('\n')
		choix = strings.TrimSpace(choix)

		switch choix {
		case "1":

			character.DisplayInfo(monper)
		case "2":
			character.AccessInventory(monper)
		case "3":
			character.Marchand(monper, reader)
		case "4":
			character.Forgeron(monper, reader)
		case "5":
			fmt.Println("A Bientot !")

			return
		default:
			fmt.Println(" choix Invalide")
		}

	}
}
