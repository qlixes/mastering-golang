package main

import (
	"fmt"

	"github.com/qlixes/modules/modules"
)

func main() {
	fmt.Println("== Concurrent ==")
	modules.RunCase()
	fmt.Println("== goroutine ==`")
}
