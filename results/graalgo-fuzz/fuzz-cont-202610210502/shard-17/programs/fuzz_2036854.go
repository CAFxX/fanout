package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int64, p1 uint, p2 int32) uint {
    var v6 float64 = float64(float32(3.14159))
    v7 := []int{0, 274}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p1
}

func main() {
    var v1 uint16 = 1
    var v2 int64 = -415
    v3 := make(map[string]uint32)
    v4 := []int{0, -402, 2147483647, 128}
    v5 := make(map[int]int)
    for fi8 := 0; fi8 < 20; fi8++ {
        v2 = v2
        _ = fi8
    }
    var v9 int = 11
    _ = len(v4)
    var v10 uint = uint(v1)
    v2 += v2
    if len(v4) > 0 {
        v4[0] = v9
    }
    var v11 int64 = v2
    if false {
        var v12 int8 = -106
        if (3.14159 > float64((0.0 - 3.14159))) {
            for fi13 := 0; fi13 < 3; fi13++ {
                switch v9 {
                case -741:
                    var v14 int = fi13
                    _ = v14
                case 55:
                    v1 -= (^(v1))
                }
                v15 := GenS1{F0: -1}
                _ = fi13
                _ = v15
            }
            v10 = (^(v10))
        }
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 float32 = (-2.25 * float32(0.1))
            v18 := GenS1{F0: v11}
            fmt.Println(v12)
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v12
    }
    switch (^(v9)) {
    case -848:
        var v19 float32 = (1e10 + (1.5 / float32(1e10)))
        _ = v19
    case -39:
        v5[(v9 >> ((genFn0((v2 % ((v11) | 1)), (v10 << ((v10) & 63)), 0)) & 63))] = v9
    default:
        var v20 uint8 = 211
        _ = v20
    }
    v2 += v2
    var v21 bool = (229 != -725)
    v1 -= (v1 % (((v1 >> ((4294967295) & 63))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
