package main

import "fmt"

func main() {
	text := "Get ready" // string
	drob := 10.0        // float
	score := 11         // int
	// score /= 3
	subcribed := true // bool

	fmt.Println(text, score)
	fmt.Println(drob)
	fmt.Println(subcribed)
	fmt.Println(score % 3) // Остаток от деления

	text1 := "hello "
	text1 += "world"
	fmt.Println(text1)
	// Другой способ объявления переменных
	var test bool
	fmt.Println(test)

	testnum := 1
	testnum++ // Инкремент
	testnum-- // Декремент

	fmt.Println(testnum)
}
