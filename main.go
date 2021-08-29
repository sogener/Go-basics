package main

import "fmt"

func main() {
	var n = 10
	Divisors(n)
}

func Divisors(n int) int {
	count := 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count += 1
		}
	}
	fmt.Println(count)
	return count
}
