package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 42
    v3 := make(map[uint32]string)
    var v4 int8 = 127
    var v5 bool = false
    v6 := GenS1{F0: 42, F1: -588}
    v2 -= -334
    var v7 uint = 4294967295
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 bool = (v7 > (v7 - (^(v7))))
        fmt.Println(v7)
        v10 := v3[1]
        _ = fi8
        _ = v10
        _ = v9
    }
    v2 = v2
    v2 -= (v2 ^ (v2 / ((v2) | 1)))
    if (!false) {
        _ = len(v3)
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        switch v2 {
        case 0:
            var v12 int64 = -1
            _ = v12
        default:
            v13 := []int{(-(fi11)), 2147483647, (fi11 * (v2 / ((fi11) | 1)))}
            _ = v13
        }
        _ = fi11
    }
    if (3.14159 != (1.5 + -2.25)) {
        v1[2] = 42
        var v14 int = v2
        _ = v14
    }
    v15 := []int{v2}
    var v16 int16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
