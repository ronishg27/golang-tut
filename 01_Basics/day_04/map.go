package day_04

import (
	"fmt"
	"maps"
)

func CheckMap() {
	person := map[string]any{
		"name": "John",
		"age":  30,
		"city": "New York",
	}

	// fmt.Println("John: ", person, " length: ", len(person))

	// for key, value := range person {
	// 	fmt.Println(key, ": ", value)

	// }

	// fmt.Println(person["name"])

	// creating map using make function
	book := make(map[string]any)
	book["title"] = "The Go Programming Language"
	book["author"] = "Alan A. A. Donovan"
	book["year"] = 2014

	// for key, value := range student {
	// 	fmt.Println(key, ": ", value)

	// }

	mergedMap := make(map[string]any)
	// two ways- using copy() and using loop
	maps.Copy(mergedMap, book)
	maps.Copy(mergedMap, person)
	for key, value := range mergedMap {
		fmt.Println(key, ": ", value)

	}
}
