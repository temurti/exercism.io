package hamming

import "errors"

// Distance - calculate Hamming distance between two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("DNA strands have different lengths")
	}
	var dist int // Hamming distance integer variable
	for pos := range a {
		if a[pos] != b[pos] {
			dist++
		}
	}
	return dist, nil
}
