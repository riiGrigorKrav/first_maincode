package mymath

import "fmt"

func Average(numbers []float64) float64 {
	var num float64

	if len(numbers) == 0 {
		fmt.Println("Need to enter numbers")
		return 0
	}
	for _, n := range numbers {
		num += n
	}
	res := num / float64(len(numbers))
	return res
}
