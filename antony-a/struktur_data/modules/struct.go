package modules

import (
    "fmt"
)

type Negara struct {
    TahunMerdeka int
    FullName string
    KepalaNegara string
    Ibukota string
}

func TaskStruct() {
    fmt.Println("-- Struct --")

    // init simple
    indonesia := Negara{
        TahunMerdeka: 1945,
        FullName: "Indonesia",
        KepalaNegara: "Presiden",
        Ibukota: "IKN (WIP)",
    }

    fmt.Println(indonesia)

    // init details
    var Arab Negara
    Arab.KepalaNegara = "Raja"
    Arab.Ibukota = "Mekkah"
    Arab.TahunMerdeka = 1890
    Arab.FullName = "United Emirat Arab"

    fmt.Println(Arab)

    // collection struct
    OutOfTopic := [2]Negara{
        Negara{
            TahunMerdeka: 1945,
            FullName: "Afganistan",
            KepalaNegara: "Raja",
            Ibukota: "Afganistan",
        },
        {
            TahunMerdeka: 1900,
            FullName: "Palestina",
            KepalaNegara: "Hamas",
            Ibukota: "Gaza",
        },
    }

    fmt.Println(OutOfTopic)
}
