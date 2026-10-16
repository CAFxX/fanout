package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int16 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []uint32{4294967295, 4294967295}
    var v3 uint64 = 362
    var v4 int16 = -1
    v5 := []string{}
    v6 := []uint32{1, 234, 908}
    var v7 uint16 = 832
    var v8 uint32 = 210
    fmt.Println(v3, v8, v4)
    var v9 uint8 = 42
    var v10 float32 = -2.25
    v9 -= (v9 & v9)
    fmt.Println(v1, v10)
    v11 := GenS1{F0: 18446744073709551615, F1: 0, F2: v4}
    var v12 bool = (false || true)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
