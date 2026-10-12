package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 bool }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint16 }
func genFn0(p0 int64, p1 int32, p2 uint8) int16 {
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: 2, F1: 42, F2: false}
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 2
    p0 = p0
    if (!(32767 >= 128)) {
        v8 := GenS2{F0: (v7 * v7), F1: (^(p2))}
        var v9 int = -9223372036854775808
        _ = v8
        _ = v9
    } else {
        fmt.Println(v7)
        v7 = 1
        for fi10 := 0; fi10 < 20; fi10++ {
            fmt.Println(v7)
            if (100 != 42) {
                var v11 float64 = (1.5 / ((1.5 + 0.0) / -2.25))
                p2 -= (p2 * 218)
                if true {
                    p1 = 13
                    v12 := GenS1{F0: 2, F1: 0, F2: false}
                    p0 -= p0
                    _ = v12
                }
                _ = v11
            } else {
                v13 := []int{391, -2147483648, -592}
                for fi14 := 0; fi14 < 15; fi14++ {
                    v15 := GenS3{F0: 0}
                    _ = fi14
                    _ = v15
                }
                var v16 uint16 = v7
                _ = v13
                _ = v16
            }
            if (32768 <= 0) {
                v7 += v7
                for fi17 := 0; fi17 < 18; fi17++ {
                    p0 -= (p0 - p0)
                    _ = fi17
                }
            }
            _ = fi10
        }
    }
    var v18 bool = (2 == uint64(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v7
    return 663
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{}
    var v3 int = -2147483648
    var v4 uint8 = 218
    v4 += (^(v4))
    var v19 int32 = -175
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 string = ("a" + ("\n\t" + ("\n\t" + ("hello" + ("hello" + "foo bar")))))
        v22 := GenS2{F0: 42, F1: 130}
        var v23 bool = false
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    if ((0.0 - 0.0) > 1.5) {
        if true {
            v3 -= (v3 >> ((4294967295) & 63))
            switch v3 {
            case -565:
                fmt.Println(v3, v19, v3)
            case 0:
                v4 -= (v4 << ((4294967295) & 63))
            default:
                var v24 int32 = 42
                _ = v24
            }
            v25 := GenS4{F0: 255, F1: 127, F2: 846}
            _ = v25
        } else {
            switch (v3 - 0) {
            case 877:
                for fi26 := 0; fi26 < 4; fi26++ {
                    v3 -= fi26
                    _ = fi26
                }
            case -1:
                var v27 float64 = 0.1
                _ = v27
            case -294:
                var v28 uint32 = 42
                _ = v28
            default:
                v29 := GenS1{F0: 507, F1: 2, F2: (!(0.1 != 3.14159))}
                _ = v29
            }
        }
        v4 = (v4 - (v4 >> ((4294967295) & 63)))
    }
    v19 = -2147483648
    v3 = (-(v3))
    var v30 uint32 = 2
    switch "" {
    case "foo bar":
        v3 += (v3 + v3)
    default:
        var v31 float32 = ((-2.25 / 0.1) - -2.25)
        _ = v31
    }
    for fi32 := 0; fi32 < 6; fi32++ {
        fmt.Println(v19, fi32)
        _ = fi32
    }
    var v33 float32 = float32(-2.25)
    switch "foo bar" {
    case "foo bar":
        v19 = int32(v30)
    case "\n\t":
        for fi34 := 0; fi34 < 19; fi34++ {
            v35 := GenS1{F0: uint64(v4), F1: (v30 & (^(v30))), F2: true}
            _ = fi34
            _ = v35
        }
    default:
        var v36 uint8 = (v4 & v4)
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v4
}
