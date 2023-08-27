package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Lorem Ipsum is simply dummy text of the printing and typesetting industry"

	file, err := os.Create("./demoFile.txt") // creating a file

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)
	fmt.Printf("%v bytes written successfully", length)
	file.Close()
	readFile("./demoFile.txt")
}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)

	checkNilError(err)

	fmt.Println("Text-data inside the file is \n", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
