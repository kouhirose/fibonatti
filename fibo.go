package main

import "fmt"

var memo map[int]int

func fib(n int) int {
	a, b := memo[n]
	if b {
		return a
	}
	value := fib(n-1) + fib(n-2)
	memo[n] = value
	return value
}

func main() {
	memo = make(map[int]int, 10)
	memo[0] = 0
	memo[1] = 1
	const length = 40
	fmt.Println("fibo Go言語")
	for i := 0; i <= length; i++ {
		fib(i)
	}
}
