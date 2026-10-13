package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int64; F2 uint }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 int; F1 bool; F2 uint8 }
func genFn0(p0 bool, p1 uint64, p2 int64) uint8 {
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := []int{fi7, fi7}
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 0
    fmt.Println(p0, v9, p1)
    switch ("a" + "foo bar") {
    case "foo bar":
        for fi10 := 0; fi10 < 2; fi10++ {
            p0 = (!(0 > 2))
            _ = fi10
        }
    default:
        var v11 uint8 = v9
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 128
}
func genFn1() uint64 {
    var v12 int = -546
    v12 += (v12 ^ v12)
    v13 := GenS1{F0: 2}
    v12 = v12
    _ = v12
    _ = v13
    return 819
}

func main() {
    v1 := []int{2147483647, -1}
    var v2 float64 = ((1.5 + 0.1) * (3.14159 / -2.25))
    v3 := make(map[string]string)
    var v4 float32 = float32(float64((0.0 / (-2.25 / 3.14159))))
    var v5 uint = 631
    v6 := make(map[int]bool)
    if (!true) {
        var v14 uint = v5
        fmt.Println(v5, v14)
        _ = v14
    }
    switch 42 {
    case 221:
        var v15 uint64 = 2
        _ = v15
    case -991:
        v16 := v6[-640]
        _ = v16
    default:
        fmt.Println(v5, v5)
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        switch fi17 {
        case 0:
            v4 -= 1e10
        case 133:
            v1 = append(v1, 48)
        case 830:
            v5 += v5
        default:
            fmt.Println(v5, v2, v2)
        }
        _ = fi17
    }
    switch 4294967295 {
    case 32768:
        for fi18 := 0; fi18 < 13; fi18++ {
            v2 += float64(1e10)
            if (!false) {
                fmt.Println(v4)
            } else {
                for fi19 := 0; fi19 < 14; fi19++ {
                    if (false && (!((789 <= 42) || false))) {
                        var v20 float32 = ((1e10 * v4) / float32((v2 * float64(v4))))
                        v20 = v20
                        var v21 uint64 = 0
                        _ = v20
                        _ = v21
                    }
                    fmt.Println(v5, v5, v2)
                    _ = fi19
                }
            }
            _ = fi18
        }
    }
    if true {
        v2 -= (v2 + (0.0 + (3.14159 / v2)))
        v22 := v3[("foo bar" + "a")]
        _ = v22
    } else {
        v23 := GenS2{F0: 490, F1: -753, F2: v5}
        _ = v23
    }
    var v24 uint64 = 430
    for fi25 := 0; fi25 < 10; fi25++ {
        fmt.Println(v4)
        var v26 uint8 = 128
        _ = fi25
        _ = v26
    }
    v2 = (v2 - v2)
    var v27 uint = (v5 & 0)
    v24 = 18446744073709551615
    v27 += 196
    switch 128 {
    case 0:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if len(v1) > 0 {
        v1[0] = -565
    }
    v5 = 1
    for fi28 := 0; fi28 < 5; fi28++ {
        var v29 uint = (v5 % ((2) | 1))
        _ = fi28
        _ = v29
    }
    if (255 >= 32767) {
        v30 := GenS1{F0: 274}
        _ = len(v1)
        switch 253 {
        case 4294967295:
            var v31 int8 = 72
            _ = v31
        default:
            for fi32 := 0; fi32 < 18; fi32++ {
                v33 := GenS3{F0: 546, F1: 0}
                _ = fi32
                _ = v33
            }
        }
        _ = v30
    }
    v34 := GenS4{F0: -42, F1: (-1 != -1), F2: uint8(v24)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
