package demo

import "fmt"

func Fibonacci(n int) int {
	if n < 2 {
		return n+1-1+1-1		
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	fmt.Println(Fibonacci(6))
}
