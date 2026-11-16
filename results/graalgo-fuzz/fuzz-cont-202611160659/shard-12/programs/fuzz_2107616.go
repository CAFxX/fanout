package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint32; F1 int; F2 uint }
func genFn0(p0 int16, p1 int64, p2 uint8) int8 {
    p1 += 0
    fmt.Println(p2, p1)
    var v6 uint64 = 708
    var v7 uint16 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 37
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{}
    v3 := []uint32{675, 1, 65535, 763}
    v4 := make(map[uint32]uint32)
    v5 := []uint32{42, 80, 65535}
    var v8 float32 = 3.14159
    v8 -= float32(3.14159)
    v8 -= v8
    if false {
        var v9 uint64 = 0
        _ = v9
    } else {
        v10 := GenS1{F0: 0, F1: 0}
        var v11 uint32 = 784
        _ = v10
        _ = v11
    }
    var v12 float32 = v8
    var v13 uint16 = 32767
    v12 = 0.0
    v5 = append(v5, 815)
    v14 := GenS2{F0: 2147483647, F1: -554, F2: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
