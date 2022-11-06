package exercises

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Length of both DNA strings should be the same")
	}

	distance := 0

	for idx, letter := range a {
		if b[idx] != byte(letter) {
			distance += 1
		}
	}
	return distance, nil
}
