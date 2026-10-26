package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int; F1 int32; F2 int64 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint32, p1 int16, p2 uint64) uint32 {
    p1 -= (p1 - -316)
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    return 2147483648
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 0
    var v3 int8 = -103
    v4 := []uint32{}
    fmt.Println(v2)
    v5 := GenS1{F0: 42, F1: 18446744073709551615}
    _ = len(v1)
    v6 := GenS2{F0: -757, F1: 532, F2: (v2 ^ (v2 << ((249) & 63)))}
    switch 32767 {
    case 127:
        var v7 int8 = 100
        _ = v7
    case -73:
        var v8 int64 = v2
        _ = v8
    default:
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := v1[genFn0(502, 1, 659)]
            _ = len(v4)
            var v11 uint64 = 2
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    v3 = (v3 * v3)
    v2 = 42
    if (-845 >= -680) {
        var v12 uint64 = 0
        v2 = v2
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := GenS3{F0: (v12 >> ((42) & 63))}
            v15 := GenS3{F0: v12}
            var v16 int16 = 32767
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        _ = v12
    }
    switch ("foo bar" + "foo bar") {
    case "":
        if ((-2147483648 == int32(v2)) && false) {
            var v17 bool = (534 <= 32767)
            _ = v17
        }
    case "a":
        v3 -= (v3 & v3)
    default:
        if (66 > 42) {
            v3 += 100
            v2 -= v2
            v2 += v2
        } else {
            v18 := v1[4294967295]
            _ = v18
        }
    }
    var v19 uint8 = 18
    if true {
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 int16 = 415
            _ = fi20
            _ = v21
        }
        var v22 uint16 = 32767
        if false {
            var v23 float32 = (3.14159 * (0.0 - float32(-2.25)))
            _ = v23
        } else {
            if true {
                _ = len(v4)
            }
            v22 -= (^(v22))
        }
        _ = v22
    }
    var v24 int32 = 42
    v3 = (v3 >> ((uint(v19)) & 63))
    v24 = (v24 + (v24 | v24))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
