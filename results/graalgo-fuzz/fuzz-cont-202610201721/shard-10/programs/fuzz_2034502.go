package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint16 }

func main() {
    var v1 uint16 = 2
    var v2 uint = 842
    v3 := []int{-10}
    var v4 float32 = (float32((0.0 - -2.25)) + (1.5 * -2.25))
    var v5 uint16 = v1
    var v6 uint = 65535
    var v7 float32 = v4
    fmt.Println(v4, v2, v2)
    for fi8 := 0; fi8 < 20; fi8++ {
        if ("" == ("\n\t" + "")) {
            v3 = append(v3, fi8)
        }
        if (499 >= 1) {
            v6 = (v2 / ((v2) | 1))
            fmt.Println(v6, v1)
        }
        _ = fi8
    }
    var v9 uint = (v6 + (^(v6)))
    v10 := GenS1{F0: 58, F1: 1, F2: 4294967295}
    var v11 int64 = -825
    v5 = 1
    var v12 float32 = v4
    v6 += v9
    v13 := GenS1{F0: v11, F1: 1, F2: 0}
    var v14 int8 = 0
    switch 128 {
    case 127:
        var v15 int64 = (v11 ^ v11)
        _ = v15
    }
    v2 -= v9
    fmt.Println(v7, v7, v7)
    if false {
        v16 := GenS2{F0: 65535, F1: 0, F2: 835}
        _ = v16
    } else {
        var v17 bool = (((v7 * (-2.25 - v12)) - v12) >= v4)
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
