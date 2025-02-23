package utils

import (
	"fmt"
)

func Array() {
	var arraySimple [6]int = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arraySimple)
	fmt.Println()
}

func ArrayHello() {
	var arrayHello [2]string = [2]string{"Hello", "World"}
	fmt.Println(arrayHello)
	fmt.Println()
}

func ArrayLoop() {
	arrayLoop := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(arrayLoop); i++ {
		fmt.Println(arrayLoop[i])
	}
}

func ArrayLoop2() {
	arrayLoop := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(arrayLoop); i++ {
		fmt.Println(arrayLoop[i])
	}
}

func ArrayLoop3() {
	arrayLoop := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range arrayLoop {
		fmt.Println(i)
	}
}

func ArrayLoop4() {
	arrayLoop := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, value := range arrayLoop {
		fmt.Println(index, value)
	}
}
