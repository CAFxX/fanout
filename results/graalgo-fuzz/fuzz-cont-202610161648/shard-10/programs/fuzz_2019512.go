package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint16, p1 uint8, p2 int) uint64 {
    p1 = p1
    v6 := []int{}
    v7 := GenS1{F0: 2147483647}
    for fi8 := 0; fi8 < 15; fi8++ {
        p1 = 32
        v6 = append(v6, 32768)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1() bool {
    v9 := GenS2{F0: 900}
    var v10 bool = false
    v11 := GenS3{F0: false}
    _ = v10
    _ = v11
    _ = v9
    return true
}

func main() {
    var v1 int = -916
    var v2 int32 = 42
    var v3 float64 = (1.5 - (0.1 - 1e10))
    var v4 int32 = (-(v2))
    v5 := make(map[string]uint32)
    var v12 int64 = 1
    fmt.Println(v4, v4)
    for fi13 := 0; fi13 < 3; fi13++ {
        fmt.Println(v12, v2, v3)
        _ = fi13
    }
    v14 := GenS2{F0: 122}
    var v15 int8 = 71
    var v16 int64 = (-(v12))
    v17 := []int{}
    v15 = (v15 << ((504) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
