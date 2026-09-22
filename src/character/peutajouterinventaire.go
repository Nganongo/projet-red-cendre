package character

const MaxInventaire = 10

func PeutAjouterInventaire(s *Personnage) bool {
	return len(s.Inventaire) < MaxInventaire
}
