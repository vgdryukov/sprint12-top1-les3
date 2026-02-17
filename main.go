package main

import "fmt"

func main() {
	greeting := getGreeting()
	fmt.Println(greeting)
}

// Выносим логику в отдельную функцию для тестирования
func getGreeting() string {
	x := "Hello, "
	y := "word!"
	return x + y
}
