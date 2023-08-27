package main

import (
	"fmt"
	"sort"
)

func main() {
	var charList = []string{"a", "b", "c"}

	charList = append(charList, "d")

	charList = append(charList[1:]) // slicing the slice and storing the result in the same variable
	fmt.Println("Value of charList: ", charList)

	// using the make for creating a slice
	names := make([]string, 4)
	names[0] = "John"
	names[1] = "Paul"
	names[2] = "George"
	names[3] = "Ringo"

	names = append(names, "Charles")

	fmt.Println("Unsorted: ", names)
	sort.Strings(names)
	fmt.Println("Sorted: ", names)
	fmt.Println("Checking if sorted: ", sort.StringsAreSorted(names))

	// removing a value from slice using indexing
	var courses = []string{"Go", "Python", "Ruby", "Java"}
	fmt.Println(courses)
	var index int = 1
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
