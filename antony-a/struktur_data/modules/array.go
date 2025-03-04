package modules

import "fmt"

func TaskArray() {
    fmt.Println("-- Array Task --")
    // init array static
    array1 := [5]string{"go","js","php","py","rs"}

    // implements get value using index
    fmt.Printf("call array with index 3 : %s \n", array1[2]) 

    // init unlimeted array
    array2 := [...]string{"Indonesia","Gelap", "2045"}

    // implements unlimited array
    fmt.Println(array2)

    // array multi dimension
    array3 := [2][5]int{
        {1,2,3,4,5},
        {6,7,8,9,10},
    }

    // select dynamic array
    fmt.Printf("get pos(1,2) = %d \n", array3[1][2])

    // array multiple data type
    array4 := []interface{}{"Indonesia", 2045, true, 0.100}

    fmt.Println(array4)

    // using alias of interface{}
    array5 := []any{"Indonesia", 2045, 0.1}

    fmt.Println(array5)
}
