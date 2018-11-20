// Package twofer is a package about sharing
package twofer

import "fmt"

// ShareWith returns a string "One for <you>, one for me". where <you> is the argument.
// If name empty, you is the string "you", else is the value in the string
// Example: "One for you, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
