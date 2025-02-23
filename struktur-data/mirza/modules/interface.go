package modules

import "fmt"

// deklarasi interface
type Singer interface {
	singing() string
}

// deklarasi struct
type Person struct {
	Name string
}

// implementasi interface
func (p Person) singing() string {
	return p.Name + " is singing"
}

func InterfaceFunction() {
	person := Person{
		Name: "Mirza",
	}

	// deklarasi variabel dengan tipe data interface
	singer := Singer(person)
	fmt.Println(singer.singing())
}
