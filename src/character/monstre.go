package character

type Monstre struct {
	Nom     string
	Pvmax   int
	Pvact   int
	Attaque int
}

func InitChienCendre() *Monstre {
	return &Monstre{
		Nom:     "chien-cendre",
		Pvmax:   40,
		Pvact:   40,
		Attaque: 5,
	}
}
