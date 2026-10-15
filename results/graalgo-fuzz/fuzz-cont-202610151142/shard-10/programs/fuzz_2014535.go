package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint8 }
func genFn0(p0 int64) uint16 {
    fmt.Println(p0)
    v6 := GenS1{F0: 283}
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 42
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]int)
    var v3 uint16 = 255
    var v4 int8 = 100
    var v5 uint64 = 2
    var v7 bool = false
    v8 := []int{962, 2147483647, 1}
    v9 := GenS2{F0: (v5 | (v5 >> ((v1) & 63))), F1: -1, F2: 128}
    delete(v2, 2)
    v2[729] = 0
    v4 += v4
    fmt.Println(v1)
    v8 = append(v8, 128)
    v5 -= (v5 >> ((v1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
