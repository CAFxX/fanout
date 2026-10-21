package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 uint32 }
func genFn0(p0 int8) uint32 {
    var v7 int8 = -110
    for fi8 := 0; fi8 < 10; fi8++ {
        p0 = v7
        _ = fi8
    }
    var v9 bool = (0 != 32767)
    fmt.Println(p0, p0, p0)
    v10 := []int{0, -416, 127}
    _ = p0
    _ = v10
    _ = v7
    _ = v9
    return 545
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    var v3 float64 = 1e10
    var v4 float32 = -2.25
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]uint32)
    var v11 float64 = (0.1 / v3)
    v12 := GenS1{F0: 1}
    var v13 bool = true
    v14 := []int{550, -2147483648}
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := GenS2{F0: -1, F1: 1}
        _ = fi15
        _ = v16
    }
    v5[120] = -1
    fmt.Println(v4, v13)
    fmt.Println(v13, v13)
    var v17 uint64 = 326
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
