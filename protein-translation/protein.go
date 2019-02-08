//Package protein implements Protein Translation
package protein

import "errors"

//ErrStop is an error returned when FromCodon receives a stop-codon
var ErrStop = errors.New("Stopped too soon")

//ErrInvalidBase is an error returned when FromCodon receives an invalid codon
var ErrInvalidBase = errors.New("Invalid Base")

//FromRNA breaks a RNA sequence into multiples codons, and then using the FromCodon function, these are translated into polypeptides
//FromRNA organizes these polypeptides in sequence and returns it
func FromRNA(rna string) ([]string, error) {
	var proteins []string
	var err error

	var protein string
	var codon string
	var indexAux int
	for limite := 3; limite <= len(rna); limite += 3 {
		indexAux = limite - 3
		codon = rna[indexAux:limite]
		protein, err = FromCodon(codon)
		if err == ErrInvalidBase {
			return proteins, err
		}
		if err == ErrStop {
			return proteins, nil
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}

//FromCodon translate codons into polypeptides
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
