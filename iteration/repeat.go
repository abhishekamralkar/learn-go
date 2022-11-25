package main

import (
	"fmt"
)

func Repeat(character string, repeatedCount int) string {
	if repeatedCount == 0 {
		repeatedCount = 1
	}
	var repeated string
	for i := 0; i < repeatedCount; i++ {
		repeated += character
	}
	return repeated
}

func main() {
	fmt.Println(Repeat("a", 10))

}
