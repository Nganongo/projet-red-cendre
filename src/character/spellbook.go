package character

import "fmt"

func spellBook(c *Personnage, spell string) {
	for _, s := range c.Skills {
		if s == spell {
			fmt.Printf("\nVous maîtrisez déjà le sort %s !\n", spell)
			return
		}
	}
	c.Skills = append(c.Skills, spell)
	fmt.Printf("\nFélicitations ! Vous avez appris le sort : %s !\n", spell)
}
