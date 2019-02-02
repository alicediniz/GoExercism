// Package proverb generates a proverb rhyme
package proverb

import "fmt"

// Proverb receives some words and creates a proverb using them
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return nil
	}

	var proverb []string
	var strAux string
	for index := 1; index < len(rhyme); index++ {
		strAux = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[index-1], rhyme[index])
		proverb = append(proverb, strAux)
	}

	strAux = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	proverb = append(proverb, strAux)
	return proverb
}
