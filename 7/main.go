package main

import "fmt"

func main() {
	N := 5
	fmt.Println(sum(N))
}
func sum(N int) int {
	sum := 0
	for i := 1; i < N+1; i++ {
		sum = sum + i
	}
	return sum
}
