package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int16, p1 uint8, p2 int8) uint64 {
    switch "foo bar" {
    case "\n\t":
        v5 := []int{42, -1, -778}
        _ = v5
    default:
        p2 = (p2 - (p2 * (p2 / ((43) | 1))))
    }
    if (128 < p1) {
        var v6 uint64 = 0
        p1 -= (p1 >> ((830) & 63))
        _ = v6
    }
    p0 = p0
    p1 = 0
    for fi7 := 0; fi7 < 4; fi7++ {
        if true {
            p1 += 0
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    return 735
}
func genFn1() int16 {
    var v8 uint8 = 62
    if ((1.5 - 0.0) <= (float32(1e10) + (0.0 / ((((1.5 - (-2.25 - 1e10)) * 3.14159) * 1.5) / float32(((3.14159 / 1e10) + -2.25)))))) {
        if ((255 == 110) && true) {
            v8 += v8
            v8 = v8
            var v9 uint16 = uint16(v8)
            _ = v9
        } else {
            if (!true) {
                v8 = v8
                v10 := GenS1{F0: 234, F1: -431}
                v8 = 198
                _ = v10
            }
        }
        if (false || ((-2.25 + 1.5) == 1e10)) {
            v11 := []int{-2147483648, int(v8)}
            v8 = (^(v8))
            _ = v11
        }
        var v12 bool = false
        _ = v12
    } else {
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 int32 = -2147483648
            _ = fi13
            _ = v14
        }
    }
    _ = v8
    return -496
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]uint32)
    v3 := []uint32{2147483648, 1, 2147483647, 2}
    v4 := make(map[int]string)
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 float32 = ((-2.25 / (((-2.25 / 1.0) - (1.5 + 1.5)) - 3.14159)) - float32(-2.25))
        v17 := GenS1{F0: 65535, F1: 1}
        var v18 uint16 = 995
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 uint32 = 42
    var v20 int16 = 0
    fmt.Println(v20, v20)
    var v21 int = -2147483648
    v22 := GenS1{F0: 42, F1: 0}
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    var v23 uint8 = 2
    switch v21 {
    case -208:
        var v24 int8 = 0
        _ = v24
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        var v26 uint16 = uint16(v23)
        var v27 int = (^(v21))
        v19 = v19
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        v29 := GenS1{F0: 209, F1: 42}
        v4[(fi28 * (v21 * (fi28 ^ fi28)))] = "\n\t"
        var v30 uint16 = 65535
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 uint16 = 883
    var v32 int8 = 100
    var v33 float64 = (3.14159 + ((-2.25 - (0.0 + -2.25)) / 0.1))
    if (42 >= 42) {
        for fi34 := 0; fi34 < 8; fi34++ {
            var v35 float64 = v33
            v36 := GenS2{F0: v23}
            v19 += v19
            _ = fi34
            _ = v35
            _ = v36
        }
    } else {
        if true {
            fmt.Println(v20, v33)
        } else {
            var v37 int = v21
            switch "a" {
            case "foo bar":
                v37 += v21
            default:
                for fi38 := 0; fi38 < 12; fi38++ {
                    v33 += v33
                    _ = fi38
                }
            }
            _ = v37
        }
        fmt.Println(v33, v19)
        var v39 uint64 = uint64(v19)
        _ = v39
    }
    v3 = append(v3, uint32(v23))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
