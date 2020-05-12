package main

import "fmt"

func main() {
	data := []int{1, 4, 3, 8, 5, 8, -1}

	fmt.Println(largest(data))
}

func largest(data []int) int {
	valMax := 0
	indMax := 0
	for ind, val := range data {
		if val == valMax {
			data[ind] = 0
		}
		if val > valMax {
			valMax = val
			indMax = ind
		}
	}
	valMax = 0
	data[indMax] = 0
	for _, val := range data {
		if val > valMax {
			valMax = val
		}
	}
	return valMax
}
