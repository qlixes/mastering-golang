package utils

import "fmt"

func Map() {
	word := map[string]string {
		"Tech": "Technology",
		"JS": "JavaScript",
		"PHP": "Bahasa PHP",
		"GO": "Golang",
	}

	for i, v := range word {
		fmt.Println("index : ", i,", value: ", v)
	}
}