package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint32 }

func main() {
    var v1 float64 = 1.5
    var v2 uint8 = 56
    var v3 int16 = 1
    switch 1 {
    case 2147483647:
        v1 += ((1e10 + v1) - v1)
    case -1:
        fmt.Println(v2, v2, v2)
    default:
        v4 := GenS1{F0: uint64(v2)}
        _ = v4
    }
    var v5 int32 = 2147483647
    var v6 uint64 = 635
    var v7 string = "\n\t"
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint32 = 777
        switch fi8 {
        case -735:
            v9 -= v9
        }
        v5 = (^(v5))
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: v2, F1: 1, F2: 166}
    v6 = v6
    v5 = v5
    v6 = (^(v6))
    fmt.Println(v6, v1)
    v5 = -2147483648
    var v11 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
