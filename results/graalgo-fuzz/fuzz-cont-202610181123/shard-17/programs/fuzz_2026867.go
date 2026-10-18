package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint64, p1 uint64, p2 bool) uint64 {
    fmt.Println(p1, p0, p0)
    p0 -= 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    return 919
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -757
    v3 := []uint32{1, 2147483647, 347}
    switch 127 {
    case 128:
        v2 += (v2 / ((v2) | 1))
    default:
        fmt.Println(v2)
    }
    v2 = v2
    var v4 int8 = -112
    var v5 float32 = 0.0
    v6 := GenS1{F0: false, F1: -121}
    v5 -= (v5 / 1.0)
    var v7 float64 = 0.0
    var v8 string = ("" + "hello")
    var v9 uint16 = 696
    v10 := GenS2{F0: false, F1: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
