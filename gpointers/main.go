package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	//var ptrone *int

	//fmt.Println("Value of pointer is ", ptrone)

	myNumber := 23
	var ptrMyOne = &myNumber

	fmt.Println("Value of pointer is ", ptrMyOne)

}
