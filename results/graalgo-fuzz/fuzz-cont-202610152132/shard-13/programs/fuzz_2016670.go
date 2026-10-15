package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
func genFn0(p0 uint64, p1 uint, p2 int64) int16 {
    p2 = 0
    v7 := GenS1{F0: 42, F1: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 768
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 42
    var v3 uint32 = 227
    v4 := []uint32{(v3 >> ((456) & 63)), v3, 4294967295}
    var v5 uint64 = 403
    var v6 uint64 = v5
    var v8 uint32 = 2
    v9 := GenS1{F0: -1, F1: 1}
    v4 = append(v4, (v3 * v3))
    v5 -= (v6 ^ v5)
    if false {
        var v10 int64 = 42
        _ = len(v1)
        _ = v10
    }
    if (0 != 732) {
        var v11 float32 = float32((0.1 + 0.1))
        v2 = (v2 >> ((18446744073709551615) & 63))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v11
    }
    v2 -= (v2 % (((v2 | (^(v2)))) | 1))
    var v12 uint8 = 127
    v12 -= 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
