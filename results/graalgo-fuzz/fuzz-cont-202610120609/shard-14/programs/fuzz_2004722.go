package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 bool; F1 int32; F2 uint64 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]uint32)
    v3 := []uint32{238}
    var v4 uint64 = 0
    var v5 float64 = (1.5 * 0.1)
    switch 4294967295 {
    case -794:
        var v6 uint64 = 4294967295
        _ = v6
    case 4294967295:
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint8 = 127
            _ = len(v2)
            if true {
                v9 := GenS1{F0: (^(v4)), F1: 372}
                var v10 uint32 = 42
                if (!((v1 | 0) <= 1)) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    v4 = 4294967295
                    v1 = v1
                    v11 := GenS2{F0: true, F1: -2147483648, F2: (v4 - v4)}
                    _ = v11
                }
                _ = v10
                _ = v9
            } else {
                var v12 uint16 = 32767
                v8 += uint8(v4)
                _ = v12
            }
            _ = fi7
            _ = v8
        }
    }
    v5 += (((v5 / float64(float32((-2.25 + v5)))) + 0.1) / v5)
    if true {
        for fi13 := 0; fi13 < 20; fi13++ {
            v1 += v1
            switch fi13 {
            case 32767:
                _ = len(v2)
            case -2147483648:
                fmt.Println(fi13, fi13, v4)
            default:
                fmt.Println(v5, v4, v5)
            }
            _ = fi13
        }
        v5 += -2.25
    }
    v1 += -9223372036854775808
    v1 = 9223372036854775807
    switch 6 {
    case 16:
        if (false || true) {
            var v14 int64 = (-(v1))
            var v15 int16 = 675
            _ = v14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 int = 42
        var v18 int32 = 37
        var v19 int64 = v1
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 uint64 = (v4 >> ((4294967295) & 63))
    var v21 uint32 = 73
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
