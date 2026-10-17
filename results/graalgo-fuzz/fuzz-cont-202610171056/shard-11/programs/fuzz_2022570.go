package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint16 }
func genFn0(p0 int16, p1 uint8) uint32 {
    var v6 int64 = -9223372036854775808
    v7 := []int{4294967295, 0, 32767, -19}
    p1 -= p1
    switch 107 {
    case 32767:
        var v8 int32 = 42
        _ = v8
    case 2147483647:
        var v9 int16 = p0
        _ = v9
    case -1:
        v6 = -184
    default:
        v6 = (v6 * (v6 ^ v6))
    }
    v10 := GenS1{F0: 279, F1: (p1 / ((255) | 1)), F2: 608}
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint32 = 4294967295
    var v2 bool = true
    var v3 uint64 = 18446744073709551615
    v4 := make(map[uint32]int)
    var v5 uint = 2
    if (!false) {
        fmt.Println(v1)
    } else {
        v1 -= (^(v1))
    }
    switch 32767 {
    case 42:
        var v11 float64 = 3.14159
        _ = v11
    }
    var v12 float32 = ((1e10 + -2.25) * 0.0)
    var v13 int = -9223372036854775808
    fmt.Println(v2)
    var v14 uint16 = 1
    v3 += v3
    v5 = (^(v5))
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 uint32 = 746
        fmt.Println(v16, v13, v2)
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        v2 = v2
        _ = fi17
    }
    var v18 string = ("" + ("hello" + "hello"))
    v1 += (v1 | v1)
    v19 := GenS1{F0: v5, F1: 2, F2: (^(v14))}
    for fi20 := 0; fi20 < 19; fi20++ {
        v5 -= uint(v3)
        _ = fi20
    }
    v21 := GenS1{F0: (v5 * (^(v5))), F1: 1, F2: (v14 | 42)}
    switch v18 {
    case "hello":
        fmt.Println(v13)
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 int32 = -975
        _ = fi22
        _ = v23
    }
    var v24 int = (v13 / (((-(v13))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
