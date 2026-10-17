package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
func genFn0(p0 bool, p1 uint32, p2 int64) bool {
    v6 := GenS1{F0: 100, F1: -128}
    p1 = p1
    if (4294967295 < 1) {
        switch 0 {
        case -9223372036854775808:
            p1 -= 864
        default:
            p2 = p2
        }
    }
    var v7 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return false
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 1.5
    var v3 float32 = 0.0
    var v4 float32 = ((0.0 / 0.1) * (v3 - 0.0))
    var v5 int8 = -58
    switch 32768 {
    case 32768:
        v8 := []int{128, 1, 0, 4294967295}
        _ = v8
    case 0:
        v9 := GenS1{F0: 0, F1: (v5 + v5)}
        _ = v9
    case 9223372036854775807:
        v10 := GenS1{F0: v5, F1: (-(v5))}
        _ = v10
    default:
        v5 += (v5 << ((18446744073709551615) & 63))
    }
    var v11 int8 = 91
    v12 := []int{int(v11), 612, 515}
    for fi13 := 0; fi13 < 7; fi13++ {
        v11 -= v11
        var v14 string = ""
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS1{F0: 100, F1: (v5 * v5)}
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: v5, F1: (^(v5))}
    fmt.Println(v5)
    var v18 float32 = float32(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
