package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func getValues(x int) (double int, tripe int, quad int) {
	double = x * 2
	tripe = x * 3
	quad = x * 4
	return
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(getValues(2))
}
