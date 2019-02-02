//Package hamming provides a function that calculates Hamming Distance
package hamming

import "errors"

//Distance takes two strings and gives the minimum number of substitutions required to change one string into the other (Hamming Distance).
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("inputs must have the same length")
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
