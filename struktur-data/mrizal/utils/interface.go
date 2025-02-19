package utils

import "fmt"

func InterfaceImplement() {
	// interface kosong menerima berbagai data
	arrayInterface := []interface{}{
		"Kalimat atau string",
		30,
		4.5,
		false,
	}

	for _, value := range arrayInterface {
		fmt.Println(value)
	}
}