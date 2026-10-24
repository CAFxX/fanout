package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = -2.25
    v3 := make(map[uint32]uint32)
    var v4 int32 = 1
    var v5 int64 = 1
    v3[0] = 42
    v2 = float32((-2.25 - 1.5))
    var v6 float32 = float32((-2.25 + 0.0))
    var v7 uint8 = 81
    var v8 float64 = 0.1
    v9 := []int{42}
    v10 := v1[65535]
    var v11 uint = 255
    var v12 uint8 = v7
    v13 := GenS1{F0: (v10 % ((0) | 1)), F1: 78, F2: 2}
    var v14 uint8 = (v12 / ((v7) | 1))
    var v15 float64 = v8
    fmt.Println(v11, v5)
    var v16 uint = (v11 | v11)
    switch "\n\t" {
    case "":
        v16 = 574
    case "hello":
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    default:
        var v17 uint64 = 940
        _ = v17
    }
    v18 := GenS2{F0: 42, F1: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
