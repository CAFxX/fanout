package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 uint16; F2 string }
func genFn0(p0 uint64) int {
    p0 = (p0 | (p0 + p0))
    v6 := []int{-8, 435, 32768, int(p0)}
    p0 -= p0
    if len(v6) > 0 {
        v6[0] = -9223372036854775808
    }
    _ = p0
    _ = v6
    return 128
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[int]bool)
    var v3 int16 = 42
    var v4 uint32 = 689
    var v5 int8 = 1
    v7 := GenS1{F0: v5, F1: 128, F2: 3}
    v4 = (v4 / ((v4) | 1))
    fmt.Println(v1, v3, v3)
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS2{F0: 127}
        v1 -= 1
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        if (0.0 == -2.25) {
            v3 += 898
        }
        _ = fi10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        if false {
            v12 := GenS3{F0: (^(v4))}
            _ = v12
        } else {
            v13 := GenS1{F0: (v5 - v5), F1: 2, F2: 255}
            _ = v13
        }
        if false {
            v5 = v5
        } else {
            v14 := []int{fi11}
            _ = v14
        }
        _ = fi11
    }
    var v15 uint64 = v1
    v16 := []int{0, 302, genFn0(v15), 9223372036854775807}
    switch 9223372036854775807 {
    case 751:
        var v17 int64 = -9223372036854775808
        _ = v17
    case 560:
        v5 = (v5 / ((-12) | 1))
    case 32767:
        var v18 bool = true
        _ = v18
    default:
        for fi19 := 0; fi19 < 10; fi19++ {
            v20 := GenS1{F0: v5, F1: 0, F2: 255}
            fmt.Println(v15, v3)
            var v21 int = (fi19 % ((fi19) | 1))
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    v2[genFn0(352)] = true
    fmt.Println(v1, v15, v4)
    var v22 float64 = -2.25
    v23 := GenS4{F0: 2147483647, F1: 255, F2: "\n\t"}
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 int16 = v3
        var v26 float64 = float64(0.0)
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
