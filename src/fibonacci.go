package demo

import "fmt"

func Fibonacci(number int) int {
	if number < 2 {
		return number
	}
	return Fibonacci(number-1) + Fibonacci(number-2)
}

func main() {
	fmt.Println(Fibonacci(6))
}
