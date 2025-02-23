package main

import (
	"fmt"
	"mirza/modules"
)

func main() {
	fmt.Println("Array Implementation")
	modules.ArrayFunction()

	fmt.Println("\nSlice Implementation")
	modules.SliceFunction()

	fmt.Println("\nMap Implementation")
	modules.MapFunction()

	fmt.Println("\nStruct Implementation")
	modules.StructFunction()

	fmt.Println("\nPointer Implementation")
	modules.PointerFunction()

	fmt.Println("\nInterface Implementation")
	modules.InterfaceFunction()

	fmt.Println("\nChannel Implementation")
	modules.ChannelFunction()
}
