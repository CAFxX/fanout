package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int8; F2 uint64 }
func genFn0(p0 string, p1 bool, p2 int8) string {
    var v6 uint = 255
    var v7 uint32 = 658
    v8 := GenS1{F0: v6}
    v9 := GenS1{F0: 18446744073709551615}
    v10 := GenS2{F0: -64, F1: 127, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return ((p0 + (((p0 + (((p0 + (p0 + (p0 + p0))) + p0) + p0)) + p0) + "hello")) + p0)
}

func main() {
    v1 := []string{}
    var v2 int = -1
    v3 := make(map[uint32]int)
    v4 := []uint32{893, 2}
    v5 := []int{(v2 % (((v2 | 0)) | 1)), v2, 0}
    v2 = 128
    var v11 float32 = -2.25
    var v12 uint64 = 1
    if (562 == -228) {
        delete(v3, 366)
    }
    fmt.Println(v2)
    switch "a" {
    case "\n\t":
        if (42 <= 2) {
            v5 = append(v5, 128)
            _ = len(v3)
        }
    default:
        fmt.Println(v2)
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        v11 = (-2.25 / 1.0)
        var v14 uint32 = 2
        if false {
            v2 -= (v2 + fi13)
        } else {
            v2 = fi13
            for fi15 := 0; fi15 < 19; fi15++ {
                var v16 uint = 0
                var v17 float64 = 0.0
                _ = fi15
                _ = v16
                _ = v17
            }
            v12 = v12
        }
        _ = fi13
        _ = v14
    }
    var v18 float32 = (v11 + ((1e10 + -2.25) + (v11 / (v11 + v11))))
    fmt.Println(v11)
    v4 = append(v4, 0)
    var v19 uint32 = 65535
    v18 += v11
    var v20 int8 = 42
    for fi21 := 0; fi21 < 19; fi21++ {
        v2 = fi21
        var v22 int16 = 42
        _ = fi21
        _ = v22
    }
    var v23 float32 = (0.1 / v18)
    v2 = (-(v2))
    v20 += (v20 / (((-(v20))) | 1))
    switch v2 {
    case 969:
        _ = len(v3)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
