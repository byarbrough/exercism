package hamming

import "fmt"

// Distance computes the Hamming Distance for two strings
func Distance(a, b string) (int, error) {
	// strings must be equal length
	if len(a) != len(b) {
		return 0, fmt.Errorf("%s and %s unequal length", a, b)
	}

	distance := 0
	// iterate over strings
	for i := range a {
		if string([]rune(a)[i]) != string([]rune(b)[i]) {
			distance++
		}
	}

	return distance, nil
}
