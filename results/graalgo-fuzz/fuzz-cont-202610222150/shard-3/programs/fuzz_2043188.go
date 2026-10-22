package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int8 }
type GenS4 struct { F0 uint8; F1 bool }
func genFn0(p0 bool, p1 int8, p2 string) uint64 {
    v4 := []int{}
    if (838 < 191) {
        if (p0 && ((0 == 18446744073709551615) || (!p0))) {
            switch -198 {
            case -2147483648:
                fmt.Println(p2, p1)
            case 429:
                p2 = p2
            default:
                if ((0.1 - (float32((1e10 * 0.1)) + (-2.25 * 1e10))) == 1.5) {
                    if (!((-1 > 799) && p0)) {
                        v5 := GenS1{F0: (p1 + p1)}
                        if (p0 && p0) {
                            for fi6 := 0; fi6 < 11; fi6++ {
                                v7 := GenS1{F0: p1}
                                _ = len(v4)
                                for fi8 := 0; fi8 < 3; fi8++ {
                                    var v9 int = (fi6 | 32768)
                                    v4 = append(v4, fi8)
                                    _ = fi8
                                    _ = v9
                                }
                                _ = fi6
                                _ = v7
                            }
                        }
                        _ = v5
                    }
                    _ = len(v4)
                    if true {
                        var v10 float32 = 1.5
                        for fi11 := 0; fi11 < 16; fi11++ {
                            v12 := GenS2{F0: 438}
                            _ = fi11
                            _ = v12
                        }
                        var v13 int = 2147483647
                        _ = v10
                        _ = v13
                    } else {
                        p1 += p1
                        var v14 int32 = -883
                        _ = len(v4)
                        _ = v14
                    }
                } else {
                    var v15 uint64 = 2
                    var v16 uint = 65535
                    _ = v15
                    _ = v16
                }
            }
        } else {
            var v17 uint16 = 65535
            v17 = v17
            _ = v17
        }
        p2 = p2
        v18 := GenS3{F0: p1, F1: p1}
        _ = v18
    } else {
        _ = len(v4)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1(p0 uint8) int32 {
    var v19 float64 = (-2.25 / 1.0)
    v19 += v19
    v20 := []int{-2147483648, 2147483647, -571}
    p0 = 1
    _ = p0
    _ = v19
    _ = v20
    return 1
}

func main() {
    var v1 float32 = 0.0
    var v2 uint = 65535
    var v3 uint16 = 42
    v2 -= v2
    var v21 uint64 = genFn0((true || true), 0, "\n\t")
    v2 = v2
    fmt.Println(v2, v1, v1)
    fmt.Println(v21, v21)
    for fi22 := 0; fi22 < 11; fi22++ {
        fmt.Println(fi22)
        if ((42 >= 2147483647) || false) {
            v23 := GenS4{F0: 255, F1: false}
            _ = v23
        }
        _ = fi22
    }
    v24 := []int{9223372036854775807, 2147483647}
    v3 += (v3 << ((v2) & 63))
    var v25 int8 = 100
    v26 := GenS1{F0: (-(v25))}
    var v27 uint32 = 43
    v28 := GenS2{F0: v21}
    switch ("\n\t" + ("\n\t" + ("hello" + ("foo bar" + "")))) {
    case "hello":
        v25 = 127
    default:
        v21 = (v21 % ((genFn0((false && (true || true)), v25, ("" + "\n\t"))) | 1))
    }
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    var v29 bool = ((!true) || true)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
