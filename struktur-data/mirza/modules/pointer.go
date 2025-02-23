package modules

import "fmt"

func PointerFunction() {
	// deklarasi variabel dengan tipe data pointer
	number := 4
	// deklarasi variabel pointer
	pointer := &number

	// show pointer value
	fmt.Println(pointer)
	// show pointer address
	fmt.Println(&pointer)
}
