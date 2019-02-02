// Package dna implements a function that counts how many times each nucleotide appears in a DNA strand
package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a Histogram of valid nucleotides in the given DNA.
// Returns an error if the given DNA contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	for _, v := range d {
		_, ok := h[v]
		if !ok {
			return nil, fmt.Errorf("%v is not a nucleotide", v)
		}
		h[v]++
	}
	return h, nil
}
