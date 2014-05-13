package demo

import "testing"

func benchmarkFibonacci(number int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fibonacci(number)
	}
}

func BenchmarkFibonacci1(b *testing.B) {
	benchmarkFibonacci(1, b)
}

func BenchmarkFibonacci3(b *testing.B) {
	benchmarkFibonacci(3, b)
}

func BenchmarkFibonacci5(b *testing.B) {
	benchmarkFibonacci(5, b)
}

func BenchmarkFibonacci7(b *testing.B) {
	benchmarkFibonacci(7, b)
}

func BenchmarkFibonacci10(b *testing.B) {
	benchmarkFibonacci(10, b)
}

func BenchmarkFibonacci20(b *testing.B) {
	benchmarkFibonacci(20, b)
}

func BenchmarkFibonacci30(b *testing.B) {
	benchmarkFibonacci(30, b)
}

func BenchmarkFibonacci40(b *testing.B) {
	benchmarkFibonacci(40, b)
}
