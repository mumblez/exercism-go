package strand

const testVersion = 3

var rnaMap = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA returns RNA complement to DNA
func ToRNA(dna string) string {
	var complement []rune
	for _, strand := range dna {
		complement = append(complement, rnaMap[strand])
	}
	return string(complement)
}
