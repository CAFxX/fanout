package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint }

func main() {
    var v1 float32 = 0.0
    var v2 int8 = -1
    var v3 int8 = (^(v2))
    var v4 bool = false
    var v5 uint64 = 0
    var v6 float32 = ((0.1 * float32(-2.25)) - (v1 * v1))
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int64 = -419
        v9 := []int{(fi7 >> ((335) & 63)), 962, fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        v2 = 0
        _ = fi10
    }
    var v11 int = -1
    var v12 int32 = -2147483648
    var v13 int = v11
    switch ("hello" + "") {
    case "":
        var v14 uint64 = (v5 - (^(v5)))
        _ = v14
    case "\n\t":
        switch (-(v13)) {
        case 2147483647:
            v2 -= (v2 / ((42) | 1))
        case 4294967295:
            v15 := GenS1{F0: -1, F1: 20, F2: 255}
            _ = v15
        case 42:
            if (2 >= 2) {
                v16 := GenS2{F0: 334, F1: 127, F2: 42}
                _ = v16
            }
        default:
            v17 := []int{v11}
            _ = v17
        }
    case "a":
        var v18 uint8 = 0
        _ = v18
    default:
        fmt.Println(v4)
    }
    fmt.Println(v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
