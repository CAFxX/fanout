package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int8; F1 bool; F2 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 int16; F2 int64 }
type GenS6 struct { F0 int8; F1 bool; F2 uint }

func main() {
    v1 := []uint32{2147483647, 2}
    v2 := make(map[uint32]bool)
    var v3 int32 = 0
    v4 := []uint32{4294967295, 2147483647, 807, 4294967295}
    _ = len(v2)
    v5 := GenS1{F0: 0, F1: 0}
    v3 -= (v3 - -1)
    v3 += (v3 * 42)
    v6 := GenS1{F0: 4294967295, F1: 620}
    v7 := GenS2{F0: -81, F1: (1e10 == (0.1 - (0.0 + 1.5))), F2: -853}
    if ((float32((float64(0.0) / float64(0.1))) + (0.1 / 1e10)) <= (0.1 / 1.0)) {
        if false {
            var v8 float32 = ((1.5 - -2.25) * float32((1e10 * -2.25)))
            _ = v8
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v9 := GenS3{F0: -2}
        _ = v9
    } else {
        switch 127 {
        case 9223372036854775807:
            v10 := GenS4{F0: 42}
            _ = v10
        default:
            var v11 int64 = 0
            _ = v11
        }
    }
    fmt.Println(v3)
    v12 := GenS5{F0: v3, F1: 902, F2: 0}
    for fi13 := 0; fi13 < 10; fi13++ {
        if true {
            if len(v4) > 0 {
                v4[0] = 1
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v3, v3)
        }
        _ = fi13
    }
    v3 += (v3 >> ((212) & 63))
    v3 = v3
    v14 := GenS6{F0: 100, F1: false, F2: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
