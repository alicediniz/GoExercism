//Package raindrops converts an integer into a specific string, based on its factor
package raindrops

import "strconv"

//Convert receives an integer, converts it to a string and returns this string
func Convert(num int) string {
	var converted string
	if num%3 == 0 {
		converted = "Pling"
	}

	if num%5 == 0 {
		converted += "Plang"
	}

	if num%7 == 0 {
		converted += "Plong"
	}

	if len(converted) == 0 {
		converted = strconv.Itoa(num)
	}
	return converted
}
