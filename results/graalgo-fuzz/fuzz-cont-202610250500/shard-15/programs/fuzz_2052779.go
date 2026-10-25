package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int32 }
type GenS2 struct { F0 bool; F1 int; F2 uint16 }
func genFn0(p0 int8, p1 uint8) uint64 {
    var v7 float32 = 1e10
    if (42 > 170) {
        v8 := []int{2147483647, -770}
        v7 += 1.5
        var v9 uint64 = 42
        _ = v8
        _ = v9
    } else {
        var v10 int16 = -1
        fmt.Println(p1, p1, p1)
        var v11 uint32 = 65535
        _ = v10
        _ = v11
    }
    switch "" {
    case "foo bar":
        fmt.Println(v7, p1, p1)
    default:
        p0 = 100
    }
    var v12 float32 = (v7 / (v7 / 1.0))
    p0 = p0
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    return 835
}

func main() {
    var v1 float32 = float32(float64(float32(-2.25)))
    v2 := []int{196}
    var v3 uint = 999
    v4 := []string{"a"}
    v5 := []string{""}
    v6 := make(map[int]int)
    for fi13 := 0; fi13 < 15; fi13++ {
        delete(v6, fi13)
        var v14 uint = v3
        _ = fi13
        _ = v14
    }
    var v15 float64 = 0.1
    switch 2147483647 {
    case 470:
        if false {
            for fi16 := 0; fi16 < 8; fi16++ {
                v6[(fi16 << ((100) & 63))] = 4294967295
                _ = fi16
            }
            v17 := GenS1{F0: 916, F1: 428, F2: 0}
            if false {
                var v18 uint64 = 1
                _ = v18
            } else {
                var v19 float32 = (v1 + 0.0)
                var v20 uint32 = 42
                _ = v19
                _ = v20
            }
            _ = v17
        }
    case -847:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v2 = append(v2, 42)
    var v21 float64 = 1e10
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 bool = (79 <= 42)
        _ = fi22
        _ = v23
    }
    switch "hello" {
    case "hello":
        fmt.Println(v3, v15, v21)
    }
    v24 := GenS2{F0: (-689 > 634), F1: int(v3), F2: 261}
    v25 := GenS2{F0: (0 != 2), F1: 32768, F2: 1}
    v15 = ((float64(((v1 * 3.14159) / v1)) + ((float64(v1) * v21) / float64(float32((v15 * 1.5))))) * ((0.0 * v15) + ((v21 - 0.0) - v15)))
    v26 := GenS2{F0: true, F1: 4294967295, F2: 255}
    var v27 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
