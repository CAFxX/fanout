package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int16 }
type GenS2 struct { F0 bool }
func genFn0(p0 int, p1 int64) int {
    v5 := GenS1{F0: 1, F1: -98, F2: int16(p1)}
    if (!(!true)) {
        for fi6 := 0; fi6 < 7; fi6++ {
            switch "a" {
            case "\n\t":
                p1 -= p1
            case "a":
                if false {
                    if ((!true) || ((2 <= 0) && true)) {
                        p1 += -596
                        var v7 uint8 = 246
                        _ = v7
                    } else {
                        for fi8 := 0; fi8 < 20; fi8++ {
                            var v9 bool = (-2.25 < 1e10)
                            _ = fi8
                            _ = v9
                        }
                        p1 += (p1 ^ p1)
                    }
                } else {
                    v10 := []int{fi6, (p0 | p0), (p0 * (p0 * (fi6 >> ((4294967295) & 63))))}
                    p1 += p1
                    v11 := GenS2{F0: true}
                    _ = v10
                    _ = v11
                }
            }
            _ = fi6
        }
        fmt.Println(p0, p0)
    }
    v12 := []int{}
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 int32 = -2147483648
            var v16 float32 = 1e10
            _ = fi14
            _ = v15
            _ = v16
        }
        switch "" {
        case "":
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        default:
            p0 = p0
        }
        _ = len(v12)
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v5
    return 696
}
func genFn1(p0 uint32) uint32 {
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 uint64 = 780
        var v19 int32 = -982
        _ = fi17
        _ = v18
        _ = v19
    }
    p0 += 2147483647
    _ = p0
    return p0
}

func main() {
    v1 := []uint32{0, 744, 0, 765}
    var v2 int16 = -32768
    v3 := []uint32{}
    v4 := []int{4294967295, 0}
    var v20 int8 = 1
    var v21 int = genFn0(genFn0(genFn0(genFn0(-2147483648, -669), -165), 685), 1)
    for fi22 := 0; fi22 < 17; fi22++ {
        v2 -= 42
        _ = fi22
    }
    if (211 <= 65535) {
        v20 = -1
    }
    var v23 uint8 = 1
    v24 := GenS1{F0: -1, F1: v20, F2: -1}
    if false {
        var v25 bool = (739 > 2)
        for fi26 := 0; fi26 < 6; fi26++ {
            var v27 uint32 = 0
            v23 -= (^(v23))
            var v28 bool = ((true || (395 == 884)) || true)
            _ = fi26
            _ = v27
            _ = v28
        }
        v23 -= 127
        _ = v25
    }
    var v29 bool = true
    v30 := GenS2{F0: (632 > 18446744073709551615)}
    if true {
        for fi31 := 0; fi31 < 18; fi31++ {
            var v32 uint8 = (v23 + (v23 - (v23 << ((158) & 63))))
            _ = fi31
            _ = v32
        }
        for fi33 := 0; fi33 < 1; fi33++ {
            v29 = (255 <= 65535)
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
            fmt.Println(v21, v21)
            _ = fi33
        }
    }
    fmt.Println(v23)
    v23 = (v23 & (v23 & v23))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
