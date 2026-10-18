package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []uint32{2147483647, 791}
    var v2 float32 = (0.1 / 1e10)
    v3 := []string{"hello", "foo bar"}
    var v4 uint16 = 1
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]string)
    _ = len(v5)
    var v7 float64 = 0.0
    var v8 int32 = 42
    v9 := v5[78]
    v10 := GenS1{F0: 92, F1: -993, F2: (v8 % (((v8 >> ((753) & 63))) | 1))}
    v11 := GenS1{F0: 0, F1: -1, F2: v8}
    _ = len(v1)
    v12 := GenS1{F0: 181, F1: 718, F2: v8}
    v4 = v4
    v4 += 1
    v4 += (v4 | v4)
    v13 := GenS2{F0: true}
    v14 := GenS3{F0: -382}
    switch -6 {
    case 555:
        v15 := GenS2{F0: (2147483647 < v8)}
        _ = v15
    case 1:
        v6[2147483647] = ("" + ("" + "hello"))
    case 319:
        v7 = (v7 / float64(0.1))
    default:
        _ = len(v3)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
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
    _ = v8
    _ = v9
}
