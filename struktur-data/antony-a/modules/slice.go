package modules

import (
    "fmt"
)

func TaskSlice() {
    fmt.Println("-- Slice --")

    // init
    slice1 := [5]string{"Indonesia", "Gelap", "Efishery", "Fraud", "Winter"}

    // get collection betwen index 1 until 3
    fmt.Println(slice1[3:4])

    // get all collection start from index 4
    fmt.Println(slice1[4:])

    // get all collection until index 3
    fmt.Println(slice1[:3])
}
