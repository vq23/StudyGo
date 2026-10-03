package main

import "fmt"

/*
	Многострочный комментарий
*/
// Однострочный

func main() {
	// score := 51
	// if score == 50 {
	// 	fmt.Println("Poltinnik!")
	// } else if score == 21 {
	// 	fmt.Println("Ochko")
	// } else if score == 12 {
	// 	fmt.Println("Дюжина")
	// } else {
	// 	fmt.Println("Учись")
	// }
	// fmt.Println("Спасибо за игру, ваш счет: ", score)

	subsribed := true

	if subsribed {
		fmt.Println("asd")
	} else {
		fmt.Println("else")
	}

	number := 2
	ravno1 := number == 1
	if ravno1 {
		fmt.Println("ЧИсло равно одному")
	}
	// Петушиная зона
	result := 9
	if result < 5 || result > 17 /* <-- Условное выражение*/ {
		fmt.Println("Ты урод")
	} else if result > 8 && result < 15 {
		fmt.Println("Godlike")
	}

	liked := false
	if !liked {
		fmt.Println("xt dfot ")
	}
}
