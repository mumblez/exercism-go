package main

import "fmt"

func main() {
	var b []byte
	a := "{}()[]"
	b = append(b, a[0])
	b = append(b, a[1])
	b = append(b, a[2])
	b = append(b, a[3])
	b = append(b, a[4])
	b = append(b, a[5])
	fmt.Printf("b = %+v\n", b)
	b = b[:(len(b) - 1)]
	fmt.Printf("b = %+v\n", b)
	fmt.Printf("b[0] = a[0] : %+v\n", b[0] == a[0])
	fmt.Printf("'a' = %+v\n", 'a')
	fmt.Printf("a[0] == '{' = %+v\n", b[1] == '{')
}
