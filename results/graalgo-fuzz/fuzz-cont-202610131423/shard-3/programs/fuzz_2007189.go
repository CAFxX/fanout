package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int }

func main() {
    var v1 float64 = 1.5
    v2 := make(map[string]string)
    v3 := make(map[int]string)
    var v4 string = "\n\t"
    var v5 bool = false
    v5 = v5
    v6 := GenS1{F0: -32768}
    v7 := GenS2{F0: 286, F1: 9223372036854775807}
    v8 := GenS2{F0: 4294967295, F1: 9223372036854775807}
    switch v4 {
    case "a":
        if v5 {
            fmt.Println(v4)
        }
    case "":
        switch v4 {
        case "a":
            v5 = (-490 < -431)
        case "":
            v4 = v4
        }
    default:
        v1 += 0.0
    }
    if v5 {
        v5 = v5
    } else {
        for fi9 := 0; fi9 < 6; fi9++ {
            v5 = (!((65535 < 4294967295) && v5))
            fmt.Println(v1)
            _ = fi9
        }
    }
    var v10 int16 = -32768
    fmt.Println(v4, v4, v10)
    for fi11 := 0; fi11 < 1; fi11++ {
        v10 -= v10
        _ = fi11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
