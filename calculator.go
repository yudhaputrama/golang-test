package main

import "strconv"

// SumString is function to sum all converted string value to integer
func SumString(values []string) int {
	total := 0
	for _, val := range values {
		n, err := strconv.Atoi(val)
		if err == nil {
			total += n
		}
	}
	return total
}
