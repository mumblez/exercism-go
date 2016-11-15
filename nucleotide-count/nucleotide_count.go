package dna

import "errors"

const testVersion = 2

// DNA ...
type DNA string

// Histogram ...
type Histogram map[byte]int

// Count returns number of times nucleotide input appears in DNS strand
func (d DNA) Count(n byte) (tally int, err error) {
	switch n {
	case 'A', 'C', 'G', 'T':
		for i := 0; i < len(d); i++ {
			if d[i] == n {
				tally++
			}
		}
	default:
		err = errors.New("Invalid nucleotide")
	}
	return
}

// Counts returns a histogram count of all nucleotides
func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	seen := make(map[byte]bool)
	for i := 0; i < len(d); i++ {
		switch nuc := d[i]; nuc {
		case 'A', 'C', 'G', 'T':
			if !seen[nuc] {
				h[nuc], _ = d.Count(nuc)
				seen[nuc] = true
			}
		default:
			return nil, errors.New("Invalid nucleotide")
		}
	}
	return h, nil
}
