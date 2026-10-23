package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int32; F1 int16 }
func genFn0(p0 uint8) int8 {
    p0 -= (^(p0))
    var v6 float32 = 0.1
    switch -746 {
    case 0:
        var v7 int64 = -464
        _ = v7
    default:
        v8 := []int{}
        _ = v8
    }
    _ = p0
    _ = v6
    return 0
}

func main() {
    var v1 int8 = -81
    var v2 uint8 = 0
    v3 := make(map[uint32]int)
    var v4 uint64 = 0
    var v5 int32 = 42
    v4 += v4
    v4 -= 4294967295
    var v9 float32 = 1.5
    if (!(364 == 828)) {
        v1 = 127
    } else {
        var v10 uint = 318
        _ = v10
    }
    var v11 float64 = 3.14159
    v12 := GenS1{F0: 2147483647, F1: 38}
    v4 = v4
    if false {
        var v13 float32 = 3.14159
        _ = v13
    }
    v5 -= v5
    switch -486 {
    case 32768:
        var v14 int = 42
        _ = v14
    default:
        if (((v1 & 42) >= (-(v1))) || true) {
            fmt.Println(v2)
            v15 := GenS1{F0: (v5 * 868), F1: 250}
            _ = v15
        }
    }
    var v16 float64 = 0.0
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 int8 = v1
        _ = fi17
        _ = v18
    }
    v2 = v2
    v19 := GenS2{F0: 1, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
