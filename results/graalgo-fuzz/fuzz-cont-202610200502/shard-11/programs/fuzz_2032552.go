package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int }
type GenS2 struct { F0 uint }
func genFn0(p0 uint32, p1 string) string {
    var v4 uint16 = 0
    p0 += (^(p0))
    switch ((p1 + (p1 + (p1 + p1))) + (p1 + p1)) {
    case "\n\t":
        fmt.Println(v4, p0)
    case "a":
        v5 := []int{4294967295, 2147483647, -185, int(v4)}
        _ = v5
    default:
        var v6 float64 = 0.0
        _ = v6
    }
    p0 = (^(p0))
    v7 := GenS1{F0: 42, F1: 1, F2: int(p0)}
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return p1
}

func main() {
    v1 := []int{127, 1}
    var v2 uint8 = 128
    var v3 uint64 = 42
    var v8 uint = 4294967295
    v1 = append(v1, 777)
    var v9 float32 = 3.14159
    var v10 uint64 = 262
    v11 := GenS2{F0: v8}
    v3 = v3
    v3 = (v10 - (v10 ^ 0))
    var v12 int64 = 0
    if false {
        for fi13 := 0; fi13 < 2; fi13++ {
            fmt.Println(v2, v9)
            _ = fi13
        }
        if (488 >= (v3 - (v3 + (v10 ^ (v10 - (v10 * (v10 << (((v8 + v8)) & 63)))))))) {
            v10 -= v3
            for fi14 := 0; fi14 < 5; fi14++ {
                v15 := GenS2{F0: 466}
                v9 = (v9 / v9)
                _ = fi14
                _ = v15
            }
        }
        v16 := GenS2{F0: 2}
        _ = v16
    }
    if len(v1) > 0 {
        v1[0] = -146
    }
    var v17 string = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
