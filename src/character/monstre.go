package character

import "math/rand"

type Monstre struct {
	Nom        string
	Pvmax      int
	Pvact      int
	Attaque    int
	Initiative int
	Xpdonnee   int
}

func InitChienCendre() *Monstre {
	return &Monstre{
		Nom:        "chien-cendre",
		Pvmax:      40,
		Pvact:      40,
		Attaque:    5,
		Initiative: rand.Intn(20) + 1,
		Xpdonnee:   20,
	}
}
