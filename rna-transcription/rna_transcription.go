// Package strand provides all the tools for the strand exercism task.
package strand

// ToRNA converts a dna stirng into a rna string
func ToRNA(dna string) string {
	var rna string

	for _, letter := range []rune(dna) {
		switch letter {
		case 'G':
			rna += "C"
		case 'C':
			rna += "G"
		case 'T':
			rna += "A"
		case 'A':
			rna += "U"
		}
	}

	return rna
}
