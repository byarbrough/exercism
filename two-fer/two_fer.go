// Package twofer returns a two-for-one with a name
package twofer

import "fmt"

// ShareWith retruns a friendly split of items between a name and me
// Defaults to "you"
func ShareWith(name string) string {

	// Default case
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
