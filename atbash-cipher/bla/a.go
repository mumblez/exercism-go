package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	a := "Aye"
	// a = 97, z = 122
	// + 32 if capital
	// - 97 to get 0
	// cipher = 26 - (letter - 97)
	// numbers = 48 (1) - 57 (9)
	b := 122 - (a[0] - 97)
	fmt.Printf("a[0] = %+v\n", string(b))
	fmt.Printf("a[0] + 'a' = %+v\n", a[0]+32)
}
