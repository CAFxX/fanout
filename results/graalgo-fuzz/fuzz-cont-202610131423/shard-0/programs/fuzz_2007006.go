package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint16, p1 int16) string {
    v7 := []int{567, 2147483647}
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int = (fi8 + (-(fi8)))
        var v10 float64 = 3.14159
        var v11 int = (fi8 * (fi8 - 127))
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 uint8 = 2
    p0 = (p0 + 368)
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    return "\n\t"
}
func genFn1(p0 int, p1 int32, p2 uint16) uint32 {
    p2 += p2
    v13 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 42
}

func main() {
    var v1 int16 = -971
    v2 := []uint32{285, 2147483647, 42}
    var v3 int8 = -119
    v4 := make(map[uint32]int)
    var v5 int8 = -128
    v6 := make(map[uint32]bool)
    v14 := GenS1{F0: 4294967295, F1: 42}
    v15 := GenS1{F0: 320, F1: v1}
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 uint = 0
        _ = len(v2)
        var v18 uint = v17
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 float64 = -2.25
    v19 = ((v19 - ((v19 - -2.25) * v19)) - (0.1 - v19))
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 uint = 65535
        v3 -= v3
        _ = fi20
        _ = v21
    }
    switch "\n\t" {
    case "":
        var v22 bool = false
        _ = v22
    case "foo bar":
        if false {
            var v23 int = 32768
            var v24 bool = false
            v6[65535] = ((3.14159 <= (0.1 / 1.0)) || ((v1 | 805) == v1))
            _ = v23
            _ = v24
        }
    default:
        v2 = append(v2, 65535)
    }
    if false {
        fmt.Println(v5, v3, v1)
        v1 = -783
        v5 = 0
    } else {
        if (-45 >= -2147483648) {
            v1 = v1
        } else {
            var v25 int64 = 974
            var v26 int = -1
            _ = v25
            _ = v26
        }
        v19 -= v19
    }
    var v27 int16 = 1
    fmt.Println(v3, v3, v27)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
