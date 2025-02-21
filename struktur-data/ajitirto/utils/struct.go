package utils

import "fmt"

type Person struct {
	Name string
	Age  int
}

type User struct {
	Person []Person
	Role string
}

func Struct() {
	// Deklarasi struct
	var user User

	// Inisialisasi struct
	user = User{
		Person: []Person{
			{"Aji", 20},
			{"tirto", 24},
		},
		Role: "admin",
	}

	fmt.Println(user)
}