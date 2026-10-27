package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int }

func main() {
    v1 := []int{2147483647, 128, 0}
    var v2 int64 = -623
    var v3 int = -2147483648
    var v4 int16 = 794
    v5 := make(map[uint32]uint32)
    var v6 uint = 65535
    v7 := GenS1{F0: int32(v2), F1: 2}
    for fi8 := 0; fi8 < 8; fi8++ {
        switch -1 {
        case 32767:
            fmt.Println(v6, v6)
        case 127:
            fmt.Println(fi8)
        case 790:
            v6 -= 18446744073709551615
        default:
            for fi9 := 0; fi9 < 2; fi9++ {
                v10 := v5[1]
                if (false && (0.0 >= 0.1)) {
                    v11 := GenS2{F0: v3}
                    _ = v11
                }
                _ = fi9
                _ = v10
            }
        }
        _ = fi8
    }
    v12 := GenS2{F0: (v3 ^ v3)}
    for fi13 := 0; fi13 < 11; fi13++ {
        v6 = (v6 | v6)
        var v14 bool = false
        _ = fi13
        _ = v14
    }
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = (-(v3))
    }
    var v15 float32 = 0.1
    v15 -= v15
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
