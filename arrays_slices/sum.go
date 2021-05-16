package arrays_slices

import "fmt"

func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		fmt.Println("number: ", number)
		sum += number
	}

	return sum
}
