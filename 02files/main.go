package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	writeFile()
	readFile("./my1gofile.txt")

}

func errorF(err error) {
	if err != nil {
		panic(err)
	}
}

func writeFile() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file."

	file, err := os.Create("./my1gofile.txt")

	length, err := io.WriteString(file, content)
	errorF(err)

	fmt.Println("Length is ", length)
	defer file.Close()
}

func readFile(filename string) {

	databyte, err := ioutil.ReadFile(filename)
	errorF(err)

	fmt.Println("Text data inside the file is \n", databyte)
	fmt.Println("Text data inside the file is \n", string(databyte))
}
