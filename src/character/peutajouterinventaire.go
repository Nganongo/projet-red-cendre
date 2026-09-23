package character

func PeutAjouterInventaire(s *Personnage) bool {
	return len(s.Inventaire) < s.InventaireMax
}
