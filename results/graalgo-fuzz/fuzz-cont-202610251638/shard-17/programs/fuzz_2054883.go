package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint8 }

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{87}
    v3 := make(map[uint32]string)
    var v4 int8 = 111
    var v5 bool = (true || (!false))
    var v6 float64 = -2.25
    v5 = v5
    var v7 uint8 = 242
    if (2147483647 == 2147483648) {
        if v5 {
            v6 += v6
            switch 0 {
            case -325:
                var v8 float32 = 1.5
                _ = v8
            case 2147483647:
                var v9 uint32 = 197
                _ = v9
            case 9223372036854775807:
                var v10 float64 = (-2.25 / v6)
                _ = v10
            default:
                var v11 float32 = -2.25
                _ = v11
            }
            var v12 uint32 = 1
            _ = v12
        }
        v13 := GenS1{F0: 139, F1: -9223372036854775808}
        _ = v13
    }
    var v14 int16 = 623
    v15 := GenS2{F0: v4, F1: 42, F2: v7}
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 uint8 = v7
        _ = fi16
        _ = v17
    }
    v4 -= int8(v14)
    var v18 int16 = v14
    _ = len(v1)
    v6 = (v6 + v6)
    if v5 {
        v4 = (v4 % ((v4) | 1))
        fmt.Println(v14, v14, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
