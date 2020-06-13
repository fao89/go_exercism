// Package twofer plays two-fer challenge
package twofer

import "fmt"

// ShareWith tells you how to share with the given name
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
