package main

import "fmt"

func main() {
	data := []int{1, 4, 3, 8, 5, 8, -1}

	fmt.Println(largest(data))
}

func largest(data []int) int {
	m := data[0]
	for _, j := range data {
		if m > j {
			m = j
		}
	}
	return m
}
