package character

import "fmt"

func UpgradeInventorySlot(s *Personnage) {
	if s.UpgradesMax >= 3 {
		fmt.Println("❌ Vous ne pouvez pas augmenter l'inventaire plus de 3 fois !")
		s.Argent += 30 // Remboursement automatique
		return
	}

	s.InventaireMax += 10
	s.UpgradesMax++
	fmt.Printf(" Capacité d'inventaire augmentée ! Vous avez maintenant %d emplacements. (Amélioration %d/3)\n", s.InventaireMax, s.UpgradesMax)
}