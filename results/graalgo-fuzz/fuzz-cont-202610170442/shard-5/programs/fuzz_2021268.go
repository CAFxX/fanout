package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint64 }

func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]bool)
    var v3 int64 = 494
    var v4 int8 = -57
    v5 := make(map[uint32]uint32)
    if (v1 < (v1 & v1)) {
        for fi6 := 0; fi6 < 6; fi6++ {
            var v7 uint64 = 337
            _ = fi6
            _ = v7
        }
        v4 = 42
    } else {
        v8 := GenS1{F0: 2147483647, F1: (v4 & -31), F2: 1}
        _ = v8
    }
    var v9 bool = false
    v3 += (v3 % (((v3 << ((983) & 63))) | 1))
    var v10 uint = 1
    v1 -= (-(v1))
    var v11 int64 = (-(v3))
    if false {
        var v12 int = -9223372036854775808
        _ = v12
    }
    var v13 uint16 = 0
    v1 -= v1
    v14 := []int{42, -314, 0, -1}
    v13 = (v13 | 65535)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
