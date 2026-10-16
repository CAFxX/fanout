package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int }
func genFn0(p0 string) uint {
    v5 := []int{32768}
    p0 = p0
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 408
}
func genFn1(p0 uint16, p1 int32, p2 uint32) int32 {
    v6 := GenS1{F0: 2, F1: 0, F2: 1000}
    switch "" {
    case "foo bar":
        p1 += p1
    case "":
        v7 := GenS2{F0: (p1 * p1)}
        _ = v7
    default:
        if true {
            var v8 int32 = p1
            _ = v8
        } else {
            var v9 uint32 = (^(p2))
            v9 -= p2
            fmt.Println(p2, p0)
            _ = v9
        }
    }
    switch "" {
    case "\n\t":
        switch -261 {
        case 127:
            var v10 float32 = (0.1 / float32(0.1))
            _ = v10
        case 42:
            p1 = (p1 % ((1) | 1))
        }
    case "hello":
        p2 -= 0
    default:
        if false {
            p1 += p1
            switch -856 {
            case 127:
                v11 := GenS3{F0: (("a" + "foo bar") + ("foo bar" + "a")), F1: 916}
                _ = v11
            case 79:
                v12 := []int{}
                _ = v12
            default:
                p1 = 216
            }
        } else {
            fmt.Println(p2, p1, p1)
        }
    }
    if true {
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := GenS3{F0: ("\n\t" + ""), F1: 111}
            _ = fi13
            _ = v14
        }
        switch "foo bar" {
        case "foo bar":
            p0 += (p0 - (^(p0)))
        case "a":
            var v15 uint64 = uint64(p2)
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}

func main() {
    var v1 uint = 0
    v2 := []uint32{}
    v3 := make(map[uint32]uint32)
    var v4 int16 = -32
    if len(v2) > 0 {
        v2[0] = 983
    }
    v4 = (v4 | (^(v4)))
    v4 = 42
    v1 = 65535
    if false {
        v4 += (v4 / (((v4 * (^(v4)))) | 1))
        v1 = (^(v1))
        v1 -= genFn0("\n\t")
    }
    if false {
        fmt.Println(v1, v1)
    }
    v16 := v3[519]
    var v17 int8 = 63
    v18 := GenS4{F0: true}
    v1 = 2
    var v19 float32 = 0.0
    if false {
        switch -436 {
        case 32768:
            if (0.0 < float64(v19)) {
                if (float64((v19 / v19)) < 1.5) {
                    switch -1 {
                    case 398:
                        v20 := GenS4{F0: (v4 < 1)}
                        _ = v20
                    default:
                        var v21 int8 = v17
                        _ = v21
                    }
                }
            }
        case 4294967295:
            for fi22 := 0; fi22 < 13; fi22++ {
                for fi23 := 0; fi23 < 1; fi23++ {
                    if ((2 > 587) && (v4 == (v4 / (((v4 & (v4 | 0))) | 1)))) {
                        v24 := GenS2{F0: 42}
                        _ = v24
                    } else {
                        var v25 int32 = 2147483647
                        _ = v25
                    }
                    var v26 float64 = ((1e10 - (0.0 * (((float64(v19) * 1e10) * 3.14159) - 0.0))) - -2.25)
                    _ = fi23
                    _ = v26
                }
                _ = fi22
            }
        }
        if (true && ((v4 >> ((genFn0("\n\t")) & 63)) != (v4 - (v4 >> ((v1) & 63))))) {
            v27 := GenS5{F0: -9223372036854775808, F1: 974}
            _ = v27
        }
        if true {
            var v28 uint16 = 32767
            _ = v28
        }
    } else {
        v4 = (v4 << ((271) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
