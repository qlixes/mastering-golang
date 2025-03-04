package main

import (
	"fmt"

	"github.com/qlixes/flowcamp/modules"
)

func main() {
	fmt.Println("== Primitive ==")
	modules.TaskPrimitive()
	fmt.Println("== Array ==")
	modules.TaskArray()
	fmt.Println("== Map ==")
	modules.TaskMap()
	fmt.Println("== Slice ==")
	modules.TaskSlice()
	fmt.Println("== Struct ==")
	modules.TaskStruct()
	fmt.Println("== application ==")
	modules.Exchange(255600, 300000)
}
