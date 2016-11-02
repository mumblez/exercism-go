// Package secret performs the secret handshake with relevant response
package secret

const (
	WINK uint = 1 << iota
	DOUBLEBLINK
	CLOSEYOUREYES
	JUMP
	REVERSE
	testVersion = 1
)

// Handshake returns string slice containing correct response
func Handshake(n uint) []string {
	if n < 1 {
		return nil
	}
	resp := []string{}
	if n&REVERSE > 0 {
		if n&JUMP > 0 {
			resp = append(resp, "jump")
		}
		if n&CLOSEYOUREYES > 0 {
			resp = append(resp, "close your eyes")
		}
		if n&DOUBLEBLINK > 0 {
			resp = append(resp, "double blink")
		}
		if n&WINK > 0 {
			resp = append(resp, "wink")
		}
	} else {
		if n&WINK > 0 {
			resp = append(resp, "wink")
		}
		if n&DOUBLEBLINK > 0 {
			resp = append(resp, "double blink")
		}
		if n&CLOSEYOUREYES > 0 {
			resp = append(resp, "close your eyes")
		}
		if n&JUMP > 0 {
			resp = append(resp, "jump")
		}
	}

	// if n&REVERSE > 0 {
	// 	reversed := []string{}
	// 	for i := range resp {
	// 		n := resp[len(resp)-1-i]
	// 		reversed = append(reversed, n)
	// 	}
	// 	return reversed
	// }
	return resp
}
