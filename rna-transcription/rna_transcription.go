// Package strand implements a function to transcribe DNA strands
package strand

import "strings"

// ToRNA takes a DNA strand and return its RNA complement
func ToRNA(dna string) string {
	rnaReplacements := strings.NewReplacer(
		"G", "C",
		"C", "G",
		"T", "A",
		"A", "U",
	)
	dna = rnaReplacements.Replace(dna)
	return dna
}
