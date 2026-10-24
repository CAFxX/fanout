package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint32 }

func main() {
    v1 := []int{0, 42}
    v2 := make(map[int]bool)
    v3 := []int{}
    var v4 uint8 = 135
    v5 := make(map[string]int)
    v6 := GenS1{F0: 961, F1: 18446744073709551615}
    _ = len(v2)
    var v7 float32 = (-2.25 - float32(((-2.25 - (0.0 * -2.25)) / 3.14159)))
    if true {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 uint = 444
            delete(v5, "a")
            _ = fi8
            _ = v9
        }
        fmt.Println(v7, v4)
        v7 = -2.25
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v10 bool = false
    var v11 float64 = -2.25
    for fi12 := 0; fi12 < 2; fi12++ {
        v13 := v2[fi12]
        v4 += v4
        for fi14 := 0; fi14 < 15; fi14++ {
            v10 = v13
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    var v15 uint8 = (v4 + (^(v4)))
    switch 0 {
    case -1:
        var v16 uint32 = 4294967295
        _ = v16
    default:
        switch 0 {
        case -2147483648:
            fmt.Println(v4)
        case 127:
            v17 := GenS1{F0: 65535, F1: 42}
            _ = v17
        case 869:
            v18 := GenS2{F0: 2147483647, F1: 2147483647}
            _ = v18
        default:
            if (-759 <= -908) {
                var v19 int8 = 13
                var v20 uint32 = 42
                var v21 uint64 = 1
                _ = v19
                _ = v20
                _ = v21
            } else {
                v11 += 0.1
                v22 := GenS2{F0: -1, F1: 2147483648}
                _ = v22
            }
        }
    }
    if len(v1) > 0 {
        v1[0] = 32768
    }
    var v23 uint32 = 0
    var v24 bool = (100 != 1)
    var v25 uint8 = (^(v4))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
