package main

import "fmt"

func main() {
	var score int = 10
	score = score + 1
	score += 1
	score++ //Инкримент
	score-- //Декремент
	text := "Hello "
	text += "World" //Конкатенация
	fmt.Println(score, text)
}
