package brackets

const testVersion = 4

// Bracket checks input to verify if all opening brackets start then close in the right order
func Bracket(s string) (bool, error) {
	if len(s) < 1 {
		return true, nil
	}
	bracketMap := make(map[byte]byte, 3)
	bracketMap['}'] = '{'
	bracketMap[')'] = '('
	bracketMap[']'] = '['
	var b []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			b = append(b, s[i])
			continue
		}
		if len(b) > 0 && b[len(b)-1] == bracketMap[s[i]] {
			b = b[:len(b)-1]
			continue
		}
		return false, nil
	}
	if len(b) == 0 {
		return true, nil
	}
	return false, nil
}
