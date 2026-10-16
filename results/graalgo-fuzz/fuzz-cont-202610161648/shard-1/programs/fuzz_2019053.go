package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{236, 2147483648, 384, 17}
    v3 := make(map[int]string)
    var v4 int64 = 42
    v5 := v3[620]
    v4 -= (v4 & v4)
    v6 := GenS1{F0: 0, F1: 2, F2: 32767}
    v7 := GenS1{F0: 100, F1: 42, F2: 74}
    v5 = ((((v5 + "a") + "") + v5) + (v5 + v5))
    var v8 uint8 = 0
    var v9 uint = 4294967295
    v8 = (v8 + v8)
    var v10 int64 = 54
    fmt.Println(v9, v5)
    for fi11 := 0; fi11 < 16; fi11++ {
        _ = len(v3)
        _ = fi11
    }
    if ((v4 / (((v4 << (((v9 ^ 793)) & 63))) | 1)) == 1) {
        if len(v2) > 0 {
            v2[0] = 0
        }
        v12 := v1[106]
        var v13 int64 = (v10 | (^(v10)))
        _ = v12
        _ = v13
    } else {
        v14 := v1[42]
        v5 = v5
        _ = v14
    }
    v4 = 758
    fmt.Println(v10, v5)
    v8 -= v8
    v8 = 164
    v15 := GenS1{F0: 1, F1: 42, F2: 128}
    var v16 int8 = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
