package atbash

// Atbash decodes the string
func Atbash(s string) (res string) {
	var counter int
	for i := 0; i < len(s); i++ {
		var letter string
		switch {
		case s[i] >= 'A' && s[i] <= 'Z':
			letter = string(122 - (s[i] + 32 - 97))
		case s[i] >= 'a' && s[i] <= 'z':
			letter = string(122 - (s[i] - 97))
		case s[i] >= '0' && s[i] <= '9':
			letter = string(s[i])
		default:
			continue
		}
		if counter == 5 {
			res += " "
			counter = 0
		}
		res += letter
		counter++
	}
	return
}
