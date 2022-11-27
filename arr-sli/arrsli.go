package main

import "fmt"

func Total(numbers [5]int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}

	fmt.Println(Total(numbers))
}
