package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int32; F2 uint32 }
type GenS4 struct { F0 int16; F1 bool }
func genFn0(p0 int32) uint64 {
    p0 -= 2147483647
    switch "foo bar" {
    case "foo bar":
        var v7 uint32 = 2147483648
        _ = v7
    default:
        var v8 int16 = 32767
        _ = v8
    }
    p0 += (^(p0))
    _ = p0
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 42
    var v3 float64 = float64((0.0 * 3.14159))
    var v4 uint64 = 183
    var v5 float64 = 1.5
    v6 := []uint32{42, 1, 702}
    var v9 int16 = 42
    v3 += 0.0
    fmt.Println(v2, v9)
    v5 += float64(0.0)
    var v10 uint = 18446744073709551615
    v3 = 0.0
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            if len(v6) > 0 {
                v6[0] = 0
            }
            if len(v6) > 0 {
                v6[0] = 48
            }
            if false {
                v13 := v1[uint32(v10)]
                _ = v13
            } else {
                v6 = append(v6, 0)
                v14 := GenS1{F0: 42, F1: (v2 | (-(v2))), F2: (v2 / ((42) | 1))}
                _ = len(v6)
                _ = v14
            }
            _ = fi12
        }
        v15 := GenS1{F0: 887, F1: (v2 | v2), F2: (v2 + v2)}
        _ = fi11
        _ = v15
    }
    var v16 uint = v10
    v17 := GenS1{F0: 2, F1: int64(v9), F2: (v2 * v2)}
    var v18 float64 = 1.5
    v19 := GenS2{F0: int32(v4)}
    v10 -= (v16 << ((v16) & 63))
    var v20 float64 = v18
    for fi21 := 0; fi21 < 1; fi21++ {
        v22 := GenS3{F0: (v2 | v2), F1: 1, F2: 813}
        v2 = 592
        _ = fi21
        _ = v22
    }
    switch 32767 {
    case 693:
        var v23 uint16 = 360
        _ = v23
    case 2147483647:
        var v24 int64 = (v2 >> ((v16) & 63))
        _ = v24
    case 32768:
        for fi25 := 0; fi25 < 18; fi25++ {
            v26 := GenS4{F0: (v9 ^ v9), F1: true}
            _ = fi25
            _ = v26
        }
    default:
        var v27 int16 = (v9 + (v9 ^ (v9 >> ((v10) & 63))))
        _ = v27
    }
    switch 9223372036854775807 {
    case 128:
        fmt.Println(v18, v5, v5)
    case 0:
        var v28 uint32 = 2
        _ = v28
    default:
        fmt.Println(v5, v4, v9)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
