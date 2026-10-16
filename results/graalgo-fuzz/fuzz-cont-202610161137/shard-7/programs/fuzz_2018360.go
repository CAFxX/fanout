package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
func genFn0(p0 string) uint {
    v4 := []int{878, 128, 2147483647}
    v5 := GenS1{F0: 42, F1: 1}
    var v6 int64 = -966
    var v7 int = -539
    if len(v4) > 0 {
        v4[0] = int(v6)
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 2
}

func main() {
    v1 := []uint32{563, 901, 797, 4294967295}
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    var v8 uint = 18446744073709551615
    var v9 bool = (-9223372036854775808 == 0)
    var v10 uint32 = 1
    var v11 float64 = 1.5
    var v12 uint32 = (v10 / (((v10 - v10)) | 1))
    var v13 uint8 = 128
    if (!(87 < -128)) {
        v14 := GenS1{F0: 42, F1: 936}
        _ = v14
    }
    v15 := GenS1{F0: -2147483648, F1: 4294967295}
    v16 := GenS1{F0: 42, F1: uint64(v13)}
    var v17 bool = v9
    var v18 bool = (false || (((0 == uint16(v12)) || v9) || v9))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
