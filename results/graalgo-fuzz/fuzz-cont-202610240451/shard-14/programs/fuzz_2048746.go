package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int64 }
func genFn0(p0 int64, p1 string, p2 uint32) uint64 {
    p0 = (p0 << ((224) & 63))
    var v4 uint32 = p2
    v4 += 2
    var v5 int8 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 18446744073709551615
}
func genFn1(p0 uint8, p1 string, p2 uint) int64 {
    fmt.Println(p0)
    p2 = 255
    _ = p0
    _ = p1
    _ = p2
    return -604
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = true
    var v3 float32 = 0.1
    fmt.Println(v3, v3)
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(v3, v3, v3)
        delete(v1, "\n\t")
        v3 += 0.0
        _ = fi6
    }
    v7 := []int{32767, -9223372036854775808, -1}
    _ = len(v7)
    var v8 int32 = 1
    _ = len(v7)
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 int16 = 0
        _ = fi9
        _ = v10
    }
    if v2 {
        _ = len(v1)
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        if v2 {
            v2 = (v2 || (((((0.0 + 0.0) * (-2.25 + ((1e10 / 1.0) / 1e10))) + ((float64(v3) * (-2.25 + 1.5)) / (-2.25 / ((0.0 / -2.25) + 1e10)))) / 1e10) != 3.14159))
        } else {
            v2 = (!v2)
            v8 += v8
        }
        _ = fi11
    }
    var v12 uint8 = 2
    v3 = v3
    v13 := GenS1{F0: 241, F1: (v12 - 0), F2: 0}
    var v14 int32 = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
