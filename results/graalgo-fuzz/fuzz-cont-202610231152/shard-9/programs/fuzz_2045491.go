package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int64 }
func genFn0(p0 int8, p1 int8) uint8 {
    p0 = (p0 / ((p0) | 1))
    var v5 int64 = 42
    _ = p0
    _ = p1
    _ = v5
    return 127
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = 2147483647
    v3 := make(map[int]int)
    var v4 uint = 255
    v4 = (v4 + (v4 + v4))
    v6 := GenS1{F0: (v2 ^ (v2 >> ((v4) & 63))), F1: 127, F2: 865}
    var v7 uint64 = 595
    var v8 uint = (v4 >> (((v4 << (((v4 << (((v4 | (v4 / (((v4 % ((18446744073709551615) | 1))) | 1)))) & 63))) & 63))) & 63))
    v7 = 2
    v9 := GenS1{F0: v2, F1: 32767, F2: 42}
    var v10 uint64 = v7
    fmt.Println(v4, v2)
    v7 = (v10 * v10)
    for fi11 := 0; fi11 < 15; fi11++ {
        fmt.Println(v2, v7)
        var v12 uint8 = 1
        var v13 uint = (^(v8))
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
