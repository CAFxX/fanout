package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 string }
func genFn0(p0 uint8) string {
    fmt.Println(p0)
    switch "\n\t" {
    case "\n\t":
        fmt.Println(p0)
    default:
        v5 := GenS1{F0: p0}
        _ = v5
    }
    _ = p0
    return "hello"
}
func genFn1(p0 uint64, p1 uint8, p2 int64) int32 {
    v6 := GenS1{F0: p1}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2147483647
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = -9223372036854775808
    var v3 float32 = ((0.1 - -2.25) + -2.25)
    var v4 float64 = 1.5
    var v7 string = "a"
    v7 = v7
    if false {
        var v8 bool = true
        _ = v8
    }
    var v9 uint8 = 0
    if ((-844 != 42) && true) {
        var v10 float32 = 3.14159
        _ = v10
    }
    var v11 int8 = 42
    var v12 uint = 538
    v1[-1] = 178
    var v13 float32 = ((0.0 + v3) / (((3.14159 + 3.14159) / (v3 / (1e10 * v3))) - v3))
    if true {
        v3 = (v3 + float32(float64(v3)))
        var v14 uint32 = 190
        v14 = 214
        _ = v14
    } else {
        fmt.Println(v9, v7, v4)
        v15 := []int{-9223372036854775808, 32767, 2147483647, 9223372036854775807}
        _ = v15
    }
    v16 := GenS1{F0: v9}
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := GenS1{F0: (^(v9))}
        _ = fi17
        _ = v18
    }
    if (((v12 > v12) || true) || (!((!true) || true))) {
        v19 := []int{127, 58, 561}
        v20 := GenS1{F0: 42}
        _ = v19
        _ = v20
    } else {
        var v21 int8 = (v11 * v11)
        var v22 bool = ((-(v2)) == v2)
        for fi23 := 0; fi23 < 5; fi23++ {
            fmt.Println(v21)
            v4 += ((v4 + (1e10 / (((v4 / v4) + 0.0) + -2.25))) + (1e10 - -2.25))
            _ = fi23
        }
        _ = v21
        _ = v22
    }
    var v24 int64 = (-(v2))
    var v25 int8 = (v11 - v11)
    v26 := GenS2{F0: 4294967295, F1: v7}
    fmt.Println(v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
