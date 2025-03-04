package modules

import (
    "fmt"
)

func TaskMap() {
    fmt.Println("-- Map --")

    // init map
    map1 := map[string]interface{}{
        "go": "Go Lang",
        "php": "Preprocessor Hypertext Programming",
        "JS": "Java Script",
        "TS": "Typescript",
        "RS": "Rust",
        "ID": 19.45,
    }

    fmt.Println(map1)

    // add map
    map1["Gibran"] = 2024

    fmt.Println(map1)

    // load data
    fmt.Println(map1["ID"])
}
