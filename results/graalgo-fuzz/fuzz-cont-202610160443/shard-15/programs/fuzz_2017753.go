package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint16 }
func genFn0(p0 int32) bool {
    var v7 uint = 65535
    var v8 uint16 = 928
    _ = p0
    _ = v7
    _ = v8
    return (!true)
}

func main() {
    var v1 float32 = 0.0
    var v2 bool = false
    var v3 uint64 = 42
    var v4 float32 = ((-2.25 - float32(0.1)) / (v1 * v1))
    var v5 int32 = 216
    var v6 int8 = 100
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 uint8 = 42
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    var v12 float32 = v1
    v2 = v2
    v6 -= (-(v6))
    if (!(0 <= 42)) {
        fmt.Println(v3)
    }
    fmt.Println(v6)
    var v13 uint16 = 2
    v12 -= v1
    var v14 uint16 = v13
    switch 2147483647 {
    case 32768:
        v6 -= (v6 / ((v6) | 1))
    case -759:
        switch "" {
        case "a":
            v15 := GenS1{F0: 42, F1: 0, F2: v14}
            _ = v15
        }
    case 9223372036854775807:
        v16 := []int{-512, -2147483648, 32768}
        _ = v16
    default:
        v17 := GenS1{F0: 0, F1: 255, F2: v13}
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
