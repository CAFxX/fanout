package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }

func main() {
    var v1 uint64 = 248
    var v2 uint16 = 1
    var v3 int = -346
    var v4 float64 = float64((0.0 - 0.0))
    for fi5 := 0; fi5 < 15; fi5++ {
        switch "" {
        case "hello":
            v6 := []int{(v3 % ((fi5) | 1)), (fi5 << ((0) & 63)), (v3 - fi5), (fi5 - (-(v3)))}
            _ = v6
        case "":
            var v7 uint16 = v2
            _ = v7
        case "\n\t":
            v8 := GenS1{F0: (fi5 - (fi5 + (fi5 * (fi5 * fi5)))), F1: v1}
            _ = v8
        default:
            v9 := []int{-1, v3}
            _ = v9
        }
        v10 := []int{}
        _ = fi5
        _ = v10
    }
    var v11 uint8 = 127
    v12 := []int{128}
    if ((v3 * 32767) == v3) {
        v2 = v2
    }
    fmt.Println(v1, v2)
    var v13 int16 = 577
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 int16 = v13
        _ = fi14
        _ = v15
    }
    v3 = v3
    v16 := GenS1{F0: v3, F1: (v1 & v1)}
    v3 = (^(v3))
    fmt.Println(v2, v13, v3)
    v17 := GenS1{F0: -370, F1: v1}
    v3 += (v3 ^ v3)
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 uint8 = v11
        v20 := GenS1{F0: 4294967295, F1: v1}
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
