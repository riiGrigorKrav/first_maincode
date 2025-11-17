package main

import (
	// "errors"
	"fmt"
)

func main() {
	// result := SumDigits(-123487687565)
	// 		fmt.Println("Result SumDigits", result)
	

	mult := Multiply(1, 22)
	fmt.Println("Result Multiply", mult)
}

func SumDigits(n int) int{
	sum := 0
	if n == 0{
		return 0
	}
		if n < 0 {
		n = -n
		return SumDigits(n)
	}
		for n > 0 {
		sum += n % 10 //вычисляется остаток от деления n на 10,
		//  который представляет собой последнюю цифру числа, и добавляется к переменной sum
		n = n / 10 //выполняется целочисленное деление n на 10, что убирает последнюю цифру из числа
	}
	return sum
}

func Multiply (x, y int) int{
	result := x * y
	return result
}