package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int32; F1 uint; F2 int64 }
func genFn0(p0 int8, p1 int8, p2 uint8) uint8 {
    var v4 int8 = (p0 ^ int8(p2))
    p0 = (p1 << ((4294967295) & 63))
    v5 := []int{}
    var v6 int32 = -462
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return p2
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    var v3 bool = true
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int16 = 42
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS1{F0: (fi9 ^ (fi9 * fi7)), F1: 42}
            v11 := []int{}
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    v12 := GenS2{F0: 0, F1: 30, F2: -9223372036854775808}
    for fi13 := 0; fi13 < 10; fi13++ {
        if v3 {
            fmt.Println(v3)
            v14 := GenS1{F0: (fi13 & (fi13 * fi13)), F1: 0}
            _ = v14
        }
        _ = fi13
    }
    v3 = v3
    v3 = v3
    fmt.Println(v3)
    fmt.Println(v3)
    v3 = (2147483647 > 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
}
