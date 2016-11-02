// Package hello library returns hello and name if passed in.
package hello

import "fmt"

const testVersion = 2

// HelloWorld returns a response saying hello and name, if no name submitted then hello world.
func HelloWorld(name string) string {
	output := "World"
	if len(name) > 0 {
		output = name
	}
	return fmt.Sprintf("Hello, %s!", output)
}
