package modules

import "fmt"

func SliceFunction() {
	// deklarasi array yang mau dijadikan slice
	fruits := []string{"apple", "grape", "banana", "melon"}

	// slice fruits to only first 3 elements
	first3Fruits := fruits[:3]
	// slice fruits to only the last element
	lastFruit := fruits[len(fruits)-1]

	// show slice result
	fmt.Println(first3Fruits)
	fmt.Println(lastFruit)
}
