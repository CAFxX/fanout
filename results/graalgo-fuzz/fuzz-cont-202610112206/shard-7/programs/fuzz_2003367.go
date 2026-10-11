package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int8 }
func genFn0(p0 uint64, p1 int32, p2 int64) uint32 {
    p0 -= p0
    var v4 bool = true
    v5 := []int{1, 2147483647}
    v4 = ((!(p0 != (p0 | 0))) || true)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 uint32 = 65535
    var v2 int16 = -106
    var v3 float64 = float64((1e10 - 0.0))
    var v6 uint64 = 985
    v1 = (^(v1))
    var v7 bool = false
    fmt.Println(v1, v1)
    v8 := GenS1{F0: 2, F1: v2}
    fmt.Println(v3, v1)
    for fi9 := 0; fi9 < 20; fi9++ {
        v6 -= 42
        if (42 != v6) {
            v10 := []int{(fi9 * (fi9 * (fi9 << ((42) & 63))))}
            _ = v10
        }
        fmt.Println(v1, v6, v1)
        _ = fi9
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 float32 = ((-2.25 / 1e10) + 3.14159)
        v13 := GenS2{F0: 2147483647, F1: v6}
        fmt.Println(v3, v2, v3)
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint32 = (v1 >> ((1) & 63))
    v15 := []int{}
    v16 := GenS3{F0: (v6 | 177), F1: 4294967295, F2: int8(v2)}
    v14 += (^(v1))
    v7 = v7
    if len(v15) > 0 {
        v15[0] = 42
    }
    v1 -= genFn0(v6, 1, -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
