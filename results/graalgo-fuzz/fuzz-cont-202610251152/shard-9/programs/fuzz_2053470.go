package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 int16; F2 uint }
type GenS5 struct { F0 uint64; F1 bool }
func genFn0(p0 int64, p1 int64, p2 int8) uint16 {
    v7 := GenS1{F0: 1000, F1: (32767 == 0)}
    if (42 <= 1000) {
        var v8 uint16 = 322
        v8 += 255
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 int16 = 42
    var v2 int16 = v1
    var v3 int16 = v2
    v4 := []int{-1, 127, 1, -211}
    v5 := make(map[uint32]uint32)
    var v6 uint8 = 127
    if (!((!false) && (1e10 != 0.1))) {
        if true {
            v9 := GenS1{F0: (v1 - -578), F1: (false || false)}
            var v10 string = "hello"
            _ = v10
            _ = v9
        }
    }
    v1 += 200
    switch -9223372036854775808 {
    case 221:
        var v11 uint8 = 128
        _ = v11
    case -2147483648:
        switch -337 {
        case -740:
            v12 := GenS2{F0: 0, F1: 2, F2: 2147483647}
            _ = v12
        case 2147483647:
            for fi13 := 0; fi13 < 11; fi13++ {
                fmt.Println(v3, v6)
                _ = fi13
            }
        case 0:
            fmt.Println(v6)
        default:
            var v14 uint = 4294967295
            _ = v14
        }
    }
    v3 = (v1 | (v3 >> ((457) & 63)))
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 int32 = 2147483647
        if true {
            for fi17 := 0; fi17 < 7; fi17++ {
                fmt.Println(v2)
                v5[uint32(v6)] = 1
                if (((fi15 >> ((18446744073709551615) & 63)) == fi15) || true) {
                    v18 := GenS3{F0: 2}
                    _ = len(v4)
                    fmt.Println(fi17, v6, v2)
                    v6 = 42
                    _ = v18
                }
                _ = fi17
            }
            v4 = append(v4, (^(fi15)))
        } else {
            var v19 int32 = -2147483648
            _ = len(v4)
            _ = v19
        }
        _ = fi15
        _ = v16
    }
    v1 -= int16(v6)
    fmt.Println(v6)
    v2 = (v1 | (v3 - (-(v1))))
    var v20 int = -885
    v20 = v20
    v20 = (^(v20))
    for fi21 := 0; fi21 < 18; fi21++ {
        var v22 int32 = -1
        _ = fi21
        _ = v22
    }
    fmt.Println(v1, v20, v20)
    var v23 bool = (3.14159 != 1.5)
    if v23 {
        for fi24 := 0; fi24 < 8; fi24++ {
            v1 += int16(v20)
            _ = fi24
        }
        v2 = (v2 * v2)
    } else {
        v25 := GenS4{F0: ("foo bar" + "a"), F1: -593, F2: 18446744073709551615}
        delete(v5, 2147483648)
        var v26 float64 = 0.1
        _ = v25
        _ = v26
    }
    v27 := GenS5{F0: 4294967295, F1: (true || v23)}
    v20 -= -216
    v3 = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
