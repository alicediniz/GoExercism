// Package scrabble calculates a word's scrabble score
package scrabble

import "strings"

// Score takes a word and returns the scrabble score for that word
func Score(word string) int {
	score := 0
	word = strings.ToUpper(word)
	for _, value := range word {
		score += points(value)
	}
	return score
}

// points takes a letter and returns how many points it's worth
func points(letter rune) int {
	switch letter {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}
