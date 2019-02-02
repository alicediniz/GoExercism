//Package isogram implements a function that checks whether a word or phrase has repeated letters
package isogram

import "unicode"

// IsIsogram receives a string and return if its a isogram or not
func IsIsogram(word string) bool {
	countChars := make(map[rune]int)
	for _, value := range word {
		if !unicode.IsLetter(value) {
			continue
		}
		value = unicode.ToLower(value)
		if countChars[value] >= 1 {
			return false
		}
		countChars[value]++
	}
	return true
}
