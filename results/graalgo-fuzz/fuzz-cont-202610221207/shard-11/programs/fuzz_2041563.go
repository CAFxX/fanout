package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int16 }
type GenS2 struct { F0 string; F1 bool }
func genFn0(p0 bool, p1 uint32, p2 string) uint16 {
    var v5 int64 = 0
    var v6 int16 = 485
    v7 := []int{0, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 854
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    var v3 string = ""
    v4 := make(map[uint32]int)
    if (245 < 65535) {
        if (65535 < 2) {
            v8 := []int{}
            _ = v8
        } else {
            v3 = (v3 + (v3 + (v3 + v3)))
        }
        var v9 uint64 = 18446744073709551615
        _ = v9
    } else {
        fmt.Println(v3)
        v3 = ""
        v10 := GenS1{F0: "\n\t", F1: 0, F2: 32767}
        _ = v10
    }
    if (false && ((2 < 1) || false)) {
        var v11 uint64 = 0
        _ = v11
    }
    var v12 uint32 = 2147483647
    for fi13 := 0; fi13 < 2; fi13++ {
        fmt.Println(fi13)
        _ = fi13
    }
    var v14 float64 = (0.0 * 1e10)
    var v15 uint32 = v12
    v3 = "\n\t"
    if false {
        delete(v2, 65535)
        v16 := GenS2{F0: (v3 + v3), F1: true}
        _ = v16
    }
    v12 = 105
    if (42 <= 225) {
        var v17 uint64 = 0
        var v18 int = -866
        _ = v17
        _ = v18
    }
    delete(v2, (v12 | 42))
    v3 = v3
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
