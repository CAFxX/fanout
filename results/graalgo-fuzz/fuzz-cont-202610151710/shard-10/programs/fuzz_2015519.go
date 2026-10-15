package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int8, p1 int8, p2 string) int16 {
    p1 = p1
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := []int{fi5, 9223372036854775807, fi5}
        v7 := GenS1{F0: fi5, F1: -32768}
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return -239
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    var v3 float32 = 1e10
    v4 := make(map[string]uint32)
    v8 := []int{4294967295}
    v9 := GenS1{F0: 128, F1: 42}
    var v10 float64 = (1e10 - 3.14159)
    var v11 uint64 = 42
    v3 -= v3
    v11 = (v11 & v11)
    fmt.Println(v10)
    v12 := GenS2{F0: 605, F1: uint16(v11)}
    var v13 uint32 = 1
    var v14 int64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
