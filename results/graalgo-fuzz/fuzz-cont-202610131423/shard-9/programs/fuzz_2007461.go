package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
func genFn0(p0 int) uint32 {
    p0 -= -771
    var v5 int64 = -703
    _ = p0
    _ = v5
    return 65535
}
func genFn1(p0 uint64) uint32 {
    p0 = p0
    p0 = 1
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 int64 = -53
        _ = fi6
        _ = v7
    }
    var v8 float64 = 0.0
    v9 := []int{9223372036854775807}
    _ = p0
    _ = v8
    _ = v9
    return genFn0(2147483647)
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    v3 := []uint32{4294967295, 0, 42}
    v4 := make(map[int]uint32)
    v3 = append(v3, 0)
    v10 := GenS1{F0: 42, F1: ""}
    var v11 float64 = (0.0 - 0.1)
    var v12 uint64 = 4294967295
    var v13 int16 = 1
    switch -412 {
    case 32767:
        v12 += 18446744073709551615
    case 127:
        var v14 int32 = 0
        _ = v14
    case 42:
        if false {
            if true {
                var v15 bool = (!false)
                _ = v15
            }
        }
    default:
        v11 = (v11 / float64(float32(v11)))
    }
    v13 = (v13 >> ((715) & 63))
    if (1e10 != 0.1) {
        var v16 bool = true
        var v17 uint8 = 128
        _ = v16
        _ = v17
    } else {
        fmt.Println(v12, v12)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
