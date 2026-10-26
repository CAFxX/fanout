package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint32 = 4294967295
    v2 := []int{671}
    var v3 uint32 = (^(v1))
    for fi4 := 0; fi4 < 13; fi4++ {
        v2 = append(v2, fi4)
        v3 = v1
        switch (fi4 >> ((381) & 63)) {
        case 683:
            v1 -= 473
        case -804:
            for fi5 := 0; fi5 < 5; fi5++ {
                var v6 uint8 = 255
                _ = fi5
                _ = v6
            }
        case -2147483648:
            var v7 uint16 = 2
            _ = v7
        default:
            v8 := GenS1{F0: "a"}
            _ = v8
        }
        _ = fi4
    }
    var v9 int16 = -1
    v1 = v3
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 float32 = 1.5
        v2 = append(v2, (fi10 / ((-243) | 1)))
        v11 = (0.1 + v11)
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        v1 -= 2147483647
        fmt.Println(v9)
        _ = fi12
    }
    v1 += (^(v3))
    v3 = (v3 & v3)
    fmt.Println(v3, v3, v1)
    v3 = v3
    switch "foo bar" {
    case "hello":
        var v13 float64 = float64((((0.1 + 0.1) - float32(0.0)) - (0.1 + 1e10)))
        _ = v13
    case "foo bar":
        v2 = append(v2, 42)
    case "\n\t":
        v14 := GenS1{F0: "foo bar"}
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := GenS1{F0: ("foo bar" + "")}
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v9
}
