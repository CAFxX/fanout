package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int32 }

func main() {
    v1 := []uint32{2147483648, 2147483647, 1}
    v2 := make(map[uint32]string)
    var v3 string = ""
    v4 := make(map[int]bool)
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]int)
    var v7 float32 = (((0.1 * 1.5) - -2.25) * 0.1)
    if len(v1) > 0 {
        v1[0] = 90
    }
    var v8 uint = 389
    v8 -= v8
    delete(v2, 0)
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 string = v3
        _ = fi9
        _ = v10
    }
    var v11 uint = (v8 & (v8 << ((2) & 63)))
    for fi12 := 0; fi12 < 13; fi12++ {
        if (-1 > -9223372036854775808) {
            var v13 float64 = ((3.14159 * (1e10 / 1e10)) / 1.5)
            _ = v13
        }
        v14 := GenS1{F0: fi12, F1: int64(fi12), F2: 42}
        _ = fi12
        _ = v14
    }
    v7 = ((float32((1e10 - ((0.1 + float64(0.0)) + 1.5))) + v7) / float32(1.5))
    v8 += 313
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
