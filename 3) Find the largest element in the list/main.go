package main

import "fmt"

func main() {
	data := []int{1, 4, 3, 8, 5, 8}

	fmt.Println(largest(data))
}

func largest(data []int) int {
	m := 0
	for _, j := range data {
		if m < j {
			m = j
		}
	}
	return m
}
