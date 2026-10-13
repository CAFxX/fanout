package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }

func main() {
    var v1 int64 = -1
    v2 := []uint32{2, 181}
    v3 := []uint32{54, 65535}
    var v4 bool = true
    var v5 float64 = float64(1e10)
    var v6 float64 = v5
    var v7 float32 = float32(v5)
    v6 = (v5 + v6)
    v4 = (-1 >= -371)
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS1{F0: fi8, F1: 4294967295}
        _ = fi8
        _ = v9
    }
    var v10 bool = v4
    v10 = (!false)
    v1 -= (v1 & v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
