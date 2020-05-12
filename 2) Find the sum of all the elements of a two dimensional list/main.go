package main

import "fmt"

func main() {
	data := [][]int{
		{1, 4, 3, 8, 5, 8},
		{1, 5, 6, 1, 3, 2},
		{2, 3, 3, 1, 4, 4},
	}
	fmt.Println(sum(data))
}
func sum(data [][]int) int {
	sum := 0
	for i := range data {
		for j := range data[i] {
			sum = sum + data[i][j]
		}
	}
	return sum
}
