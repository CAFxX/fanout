package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int8 }
type GenS2 struct { F0 bool; F1 int16 }
func genFn0(p0 int8) int {
    var v6 int8 = p0
    var v7 int64 = -1
    v8 := []int{}
    if true {
        v9 := GenS1{F0: 4294967295, F1: 1, F2: -24}
        _ = v9
    } else {
        _ = len(v8)
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return -154
}
func genFn1(p0 uint16) uint8 {
    for fi10 := 0; fi10 < 11; fi10++ {
        p0 -= p0
        v11 := GenS2{F0: true, F1: 32767}
        fmt.Println(p0, p0)
        _ = fi10
        _ = v11
    }
    var v12 uint64 = 2
    p0 = p0
    v12 = (v12 - v12)
    _ = p0
    _ = v12
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    var v3 int64 = 924
    var v4 int64 = v3
    v5 := make(map[string]uint32)
    fmt.Println(v3, v4)
    var v13 float32 = 0.0
    v14 := GenS2{F0: true, F1: 1}
    delete(v5, ("" + "hello"))
    var v15 uint8 = 127
    v16 := GenS2{F0: (false && false), F1: 0}
    switch 128 {
    case -9223372036854775808:
        var v17 float32 = v13
        _ = v17
    default:
        var v18 uint16 = 646
        _ = v18
    }
    v15 += (v15 >> ((0) & 63))
    v19 := []int{genFn0(-4), 4294967295, -157}
    v3 = v4
    v13 = (v13 * (float32((0.0 * (3.14159 / float64(v13)))) / 1.5))
    v3 = v3
    var v20 uint64 = 18446744073709551615
    if true {
        if (0 <= 1) {
            _ = len(v19)
        }
    } else {
        if (-2.25 != v13) {
            v21 := GenS1{F0: 823, F1: -134, F2: 127}
            _ = v21
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
