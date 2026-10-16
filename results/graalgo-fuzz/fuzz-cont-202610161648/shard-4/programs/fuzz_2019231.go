package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int32; F1 int; F2 uint8 }

func main() {
    v1 := []int{970, 32767, -1}
    v2 := make(map[uint32]string)
    var v3 int32 = 42
    var v4 uint32 = 2147483648
    var v5 uint8 = 128
    v6 := GenS1{F0: 42, F1: (v4 << ((255) & 63))}
    v3 = (v3 << ((602) & 63))
    v4 = v4
    var v7 uint8 = (v5 * v5)
    v7 += (v5 & v7)
    fmt.Println(v3, v4, v7)
    var v8 uint16 = 270
    var v9 uint64 = 282
    if (1 < 65535) {
        v10 := GenS1{F0: (v8 * v8), F1: (v4 * 42)}
        if true {
            var v11 int = 373
            _ = v11
        }
        _ = v10
    }
    v5 += v7
    if (0 >= 1) {
        var v12 uint32 = 2147483648
        _ = v12
    }
    var v13 bool = (v3 < v3)
    var v14 int32 = v3
    var v15 float32 = 0.1
    v9 += (v9 & v9)
    if false {
        v7 -= uint8(v8)
        v16 := GenS2{F0: -968, F1: 127, F2: (v7 + v7)}
        _ = v16
    } else {
        var v17 bool = (v3 != 857)
        v14 -= v14
        _ = len(v2)
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
