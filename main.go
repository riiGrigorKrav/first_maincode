// Калькулятор Идекса массы тела (ИМТ = вес/рост*рост)
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("___Калькулятор индекса массы тела___")
	userWidthKg, userHeight  := getUserInput()
	IMT := calculateIMT(userWidthKg, userHeight)
	outputResult(IMT)
	}
// fmt.Print("Введите свой рост в сантиметрах: ")
	// fmt.Scan(&userHeight)
	// fmt.Print("Введите свой вес в кг: ")
	// fmt.Scan(&userWidthKg)
	// result := fmt.Sprintf("Ваш индекс массы тела: %.1f\n", IMT)
	// fmt.Print(result, "Введенные данные:", userHeight, userWidthKg)

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWidthKg float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan(&userWidthKg)
	return userWidthKg, userHeight
}

func calculateIMT(userWidthKg float64, userHeight float64) float64 {
	const IMTPower = 2
	IM := userWidthKg / math.Pow(userHeight/100, IMTPower)
	return IM
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.1f\n", imt)
	fmt.Print(result)
}

// fmt.Println(Massage, "Я скоро стану GOLANG NINJA")float64
// massage := "Я тоже скоро стану GOLANG NINJA"
// fmt.Println(Massage, massage)
// massage = "Я уже становлюсь GOLANG NINJA"
// fmt.Println(massage)

// 	Msg := &massage
// 	fmt.Println(Msg)
// 	fmt.Println(*Msg)
// 	*Msg = "VO kak"
// 	fmt.Println(*Msg)

// 	var Fil int = 122
// 	fmt.Println(Fil)
// 	Fil, Fil2 := 65, "Hi you"
// 	fmt.Println(Fil, string(Fil), float64(Fil), Fil2)
// 	//так не работает var Fil string = "Сто двадцать три"
// 	fmt.Printf("%T %T \n", Fil, Fil2)

// 	a, b, c := 1, 3, 5
// 	fmt.Println(a, b, c)
// 	a, b = b, a
// 	fmt.Println(a, b, c)
// 	a, _, c = 12, 23, 34
// 	fmt.Println(a, b, c)
// 	a, b, _ = 10, 30, 50
// 	fmt.Println(a, b, c)
// }
