package modules

import "fmt"

func MapFunction() {
	// deklarasi map
	student := map[string]string{
		"name":   "Mirza",
		"age":    "23",
		"gender": "male",
	}

	// accessing map value
	fmt.Println(student["name"])
	fmt.Println(student["age"])
}
