package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -2147483648
    v3 := []int{9223372036854775807, 671}
    var v4 uint16 = 42
    v5 := make(map[uint32]int)
    var v6 uint16 = 1
    var v7 uint = 924
    if false {
        var v8 uint16 = v4
        v9 := GenS1{F0: 9223372036854775807, F1: 2147483647}
        _ = v8
        _ = v9
    }
    var v10 uint16 = (v6 % ((v4) | 1))
    switch -270 {
    case 4294967295:
        switch 2147483647 {
        case 32768:
            var v11 uint8 = 42
            _ = v11
        case 550:
            v2 += 0
        default:
            v4 += 349
        }
    case -425:
        v2 -= (v2 * 2147483647)
    case 9223372036854775807:
        v12 := GenS1{F0: 460, F1: 42}
        _ = v12
    default:
        var v13 uint64 = 918
        _ = v13
    }
    fmt.Println(v10, v6, v7)
    var v14 float64 = ((((1e10 * 3.14159) * ((((1e10 - (1.5 * 0.0)) - 0.0) + 0.1) * 1e10)) * 0.1) - float64(-2.25))
    if ((((0.1 + ((0.0 * ((0.0 / float32(1.5)) + -2.25)) * -2.25)) * 3.14159) - 0.1) > 3.14159) {
        v6 = v10
        v7 += v7
        v6 = (v4 % ((0) | 1))
    }
    var v15 int8 = -128
    fmt.Println(v10, v7, v7)
    v16 := GenS1{F0: 0, F1: v2}
    fmt.Println(v10, v15, v15)
    fmt.Println(v2, v14)
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 int = -2147483648
        _ = fi17
        _ = v18
    }
    v2 -= (v2 * (^(v2)))
    var v19 uint8 = 117
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
