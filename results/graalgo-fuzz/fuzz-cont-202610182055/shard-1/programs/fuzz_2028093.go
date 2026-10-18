package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int = 30
    var v2 uint8 = 2
    var v3 float64 = 1.5
    v4 := GenS1{F0: 42}
    var v5 string = "\n\t"
    v6 := []int{}
    v2 = 127
    fmt.Println(v5, v1)
    v1 = (v1 | 689)
    if len(v6) > 0 {
        v6[0] = (v1 / ((v1) | 1))
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(v5, v5)
        _ = fi7
    }
    v8 := GenS1{F0: 46}
    switch v5 {
    case "a":
        var v9 float64 = (((v3 * 1e10) - v3) + v3)
        _ = v9
    case "":
        fmt.Println(v2)
    case "\n\t":
        var v10 float64 = (float64(1e10) / v3)
        _ = v10
    default:
        v11 := GenS1{F0: int32(v2)}
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
