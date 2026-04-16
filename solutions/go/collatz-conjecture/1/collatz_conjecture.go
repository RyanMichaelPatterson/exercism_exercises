package collatzconjecture

import "errors"

func CollatzConjecture(number int) (int, error) {
	if number <= 0 {
		return 0, errors.New("Input must be a positive integer.")
	}
	iteration := 0
	for {
		if number == 1 {
			return iteration, nil
		}
		if number%2 == 0 {
			number = number / 2
		} else {
			number = (number * 3) + 1
		}
		iteration++
	}
}