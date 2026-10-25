package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint8 }

func main() {
    var v1 uint8 = 249
    var v2 float32 = (((3.14159 * 0.1) - (1e10 / 1e10)) + 0.1)
    var v3 int32 = 2147483647
    var v4 float32 = 0.0
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]uint32)
    if ((true || (v2 != (v2 * 1.5))) || true) {
        v4 += 0.1
        if true {
            v4 += 0.0
        } else {
            _ = len(v6)
            v4 = (((v2 * v4) + v4) - 1e10)
        }
    }
    var v7 int32 = -1
    v8 := GenS1{F0: 827, F1: -682, F2: v1}
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := []int{(fi9 ^ (fi9 << ((18446744073709551615) & 63))), (fi9 >> ((0) & 63)), 19}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 uint16 = 42
        _ = fi11
        _ = v12
    }
    var v13 uint64 = 0
    v14 := v5[-174]
    fmt.Println(v2, v7, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
