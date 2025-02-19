package modules

import "fmt"

func TaskPrimitive() {
    fmt.Println("-- Primitive --")

    // detail init variable
    var details string = "IndonesiaGelap"

    fmt.Printf("implements details variables = %s \n", details)

    simple := "IndonesiaGelap"

    fmt.Printf("implements simple variables = %s \n", simple)

    defer testDefer()

    fmt.Println("running normmal process")
}

func testDefer() {
    fmt.Println("running defer")
}
