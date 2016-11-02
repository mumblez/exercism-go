package etl

// Transform calculates scrabble points from old system
func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)
	for k, v := range input {
		for _, letter := range v {
			output[string(letter[0]+32)] = k
		}
	}
	return output
}
