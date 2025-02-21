package modules

import "fmt"

// deklarasi type struct
type User struct {
	name   string
	age    int
	gender string
}

func StructFunction() {
	// deklarasi struct
	users := []User{
		{
			name:   "Mirza",
			age:    23,
			gender: "male",
		},
		{
			name:   "Ayu",
			age:    22,
			gender: "female",
		},
	}

	fmt.Println(users)
}
