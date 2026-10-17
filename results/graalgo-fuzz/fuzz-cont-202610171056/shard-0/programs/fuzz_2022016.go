package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint64 }
type GenS2 struct { F0 int32; F1 int64; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint8) uint64 {
    var v5 int16 = 42
    v5 = 288
    var v6 uint64 = 1
    v7 := GenS1{F0: 838, F1: "a", F2: v6}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return v6
}
func genFn1(p0 int, p1 uint, p2 int16) uint8 {
    fmt.Println(p0, p0)
    v8 := GenS2{F0: -109, F1: -648, F2: 1}
    p1 = 1
    var v9 uint8 = 2
    p1 += uint(v9)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 int = 4294967295
    v2 := make(map[int]uint32)
    v3 := []uint32{2147483647}
    v4 := []string{"", "foo bar"}
    v10 := GenS2{F0: 1, F1: -1, F2: genFn0(genFn1((^(v1)), 4294967295, 32767))}
    for fi11 := 0; fi11 < 10; fi11++ {
        v1 -= v1
        var v12 uint64 = 4294967295
        _ = fi11
        _ = v12
    }
    v1 += (v1 * v1)
    for fi13 := 0; fi13 < 10; fi13++ {
        v1 += v1
        _ = fi13
    }
    fmt.Println(v1)
    v1 += (v1 << ((1) & 63))
    if false {
        switch (v1 + 806) {
        case 2147483647:
            for fi14 := 0; fi14 < 10; fi14++ {
                switch 609 {
                case -756:
                    for fi15 := 0; fi15 < 17; fi15++ {
                        var v16 int8 = 102
                        var v17 int32 = 36
                        v1 = (fi15 & (^(v1)))
                        fmt.Println(fi14, v1, v1)
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                default:
                    var v18 uint64 = genFn0(0)
                    _ = v18
                }
                for fi19 := 0; fi19 < 14; fi19++ {
                    _ = len(v4)
                    v1 = (v1 * (^(fi19)))
                    _ = fi19
                }
                if (v1 > (-(v1))) {
                    for fi20 := 0; fi20 < 3; fi20++ {
                        var v21 int16 = -32768
                        _ = fi20
                        _ = v21
                    }
                }
                _ = fi14
            }
        default:
            for fi22 := 0; fi22 < 10; fi22++ {
                var v23 uint = 144
                _ = fi22
                _ = v23
            }
        }
        if (!false) {
            fmt.Println(v1)
            switch v1 {
            case 127:
                fmt.Println(v1)
            case -2147483648:
                v1 = v1
            }
        }
        var v24 int32 = 42
        _ = v24
    } else {
        if (!(1 > -1)) {
            var v25 int = (v1 ^ (v1 - (v1 | v1)))
            fmt.Println(v1)
            v26 := GenS1{F0: 42, F1: ("a" + ""), F2: 18446744073709551615}
            _ = v25
            _ = v26
        }
        v1 = v1
    }
    v27 := GenS1{F0: -362, F1: (("foo bar" + "") + ""), F2: genFn0(48)}
    v28 := GenS3{F0: -187, F1: 216}
    var v29 int8 = -1
    if (0.1 > 1e10) {
        if (!((131 != 227) || (4294967295 >= 2))) {
            fmt.Println(v29, v1)
        }
        if (417 == 815) {
            var v30 uint32 = 483
            v1 += (v1 ^ (v1 - (v1 >> ((1) & 63))))
            _ = v30
        }
        var v31 uint16 = 1
        _ = v31
    }
    var v32 float64 = ((1e10 * 3.14159) * -2.25)
    v29 += 52
    fmt.Println(v1, v32)
    fmt.Println(v1)
    v33 := GenS3{F0: 1, F1: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
