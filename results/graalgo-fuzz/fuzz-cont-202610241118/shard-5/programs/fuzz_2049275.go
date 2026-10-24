package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int8; F1 uint16; F2 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 int8 }
func genFn0(p0 string, p1 int, p2 uint8) int64 {
    p1 = (p1 * 32767)
    fmt.Println(p2, p0, p2)
    v6 := GenS1{F0: "", F1: (42 > 1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -9223372036854775808
}

func main() {
    var v1 uint8 = 127
    var v2 uint8 = (v1 + v1)
    v3 := []uint32{707, 1}
    v4 := []uint32{1, 4294967295}
    var v5 int64 = -1
    var v7 float32 = 1e10
    var v8 float64 = ((0.1 - 1.5) * (float64(0.1) / (3.14159 / 1e10)))
    _ = len(v4)
    for fi9 := 0; fi9 < 15; fi9++ {
        v3 = append(v3, 285)
        var v10 int16 = int16(v5)
        if false {
            v11 := GenS2{F0: -21, F1: 454, F2: (0 >= 2)}
            _ = v11
        } else {
            switch (-(fi9)) {
            case 928:
                for fi12 := 0; fi12 < 13; fi12++ {
                    var v13 int = fi9
                    v14 := GenS3{F0: 42}
                    v15 := GenS4{F0: 109, F1: (fi12 & (v13 + -133))}
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
            case 127:
                v5 = (v5 << ((42) & 63))
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
        }
        _ = fi9
        _ = v10
    }
    var v16 int8 = 42
    if len(v4) > 0 {
        v4[0] = 842
    }
    v16 = (v16 | v16)
    if false {
        var v17 float64 = 1.5
        v2 = (v2 * (v2 >> ((736) & 63)))
        v18 := GenS5{F0: v16}
        _ = v17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
