package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint64, p1 uint64) uint8 {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int64 = 9223372036854775807
        _ = fi7
        _ = v8
    }
    var v9 int8 = 100
    fmt.Println(v9, p0)
    fmt.Println(p0, p0)
    var v10 float32 = 0.0
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 174
}

func main() {
    v1 := []uint32{65535, 609, 4294967295}
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 786
    var v4 float64 = (0.0 / (-2.25 / 3.14159))
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]int)
    v3 -= (v3 / (((v3 | (v3 & (v3 * (^(v3)))))) | 1))
    delete(v2, 223)
    var v11 int64 = 9223372036854775807
    fmt.Println(v4)
    v12 := v2[4294967295]
    v13 := GenS1{F0: 128}
    switch 128 {
    case 861:
        fmt.Println(v3, v4)
    default:
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := GenS1{F0: uint8(v12)}
            _ = fi14
            _ = v15
        }
    }
    if false {
        var v16 float32 = 1e10
        _ = v16
    } else {
        fmt.Println(v12)
        v12 = v12
    }
    var v17 float32 = (0.1 / float32(v4))
    var v18 int8 = 100
    fmt.Println(v12, v3)
    if (false && false) {
        var v19 int16 = 1000
        v20 := v6[65535]
        var v21 uint = 298
        _ = v19
        _ = v20
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
