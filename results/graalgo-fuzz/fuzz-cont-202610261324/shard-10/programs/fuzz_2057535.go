package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
func genFn0(p0 int64) uint32 {
    var v6 uint = 541
    fmt.Println(v6, p0)
    v6 = 255
    _ = p0
    _ = v6
    return 2147483648
}
func genFn1(p0 bool, p1 int, p2 uint64) uint64 {
    if false {
        p2 += (^(p2))
    }
    if p0 {
        v7 := GenS1{F0: 65535, F1: 42}
        _ = v7
    } else {
        p2 -= p2
        if (-9223372036854775808 == 338) {
            if (18446744073709551615 == 1) {
                for fi8 := 0; fi8 < 4; fi8++ {
                    var v9 float32 = (1.5 - float32((0.0 * (0.0 * 1e10))))
                    for fi10 := 0; fi10 < 6; fi10++ {
                        p2 = 608
                        _ = fi10
                    }
                    fmt.Println(p1, p0, p0)
                    _ = fi8
                    _ = v9
                }
                p1 -= p1
                p1 = p1
            }
        }
        fmt.Println(p2, p1)
    }
    var v11 uint16 = 255
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return p2
}

func main() {
    var v1 uint8 = 223
    var v2 int64 = 756
    var v3 bool = (!true)
    v4 := []int{128, 773, 228, 32767}
    var v5 uint32 = 0
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 -= 127
    fmt.Println(v1)
    v3 = (!false)
    v1 = (v1 & (^(v1)))
    switch "hello" {
    case "hello":
        if false {
            for fi12 := 0; fi12 < 8; fi12++ {
                for fi13 := 0; fi13 < 12; fi13++ {
                    fmt.Println(v3, v3, fi13)
                    if v3 {
                        switch fi12 {
                        case 292:
                            if (!(!(18446744073709551615 < 117))) {
                                _ = len(v4)
                                v3 = v3
                                if v3 {
                                    if (v3 && v3) {
                                        var v14 int8 = -92
                                        _ = v14
                                    }
                                }
                            }
                        default:
                            fmt.Println(fi13, v1, v3)
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    v15 := GenS1{F0: 425, F1: -1}
                    _ = fi13
                    _ = v15
                }
                v1 -= (v1 ^ 190)
                _ = fi12
            }
        } else {
            v16 := GenS1{F0: uint(v1), F1: 1}
            var v17 uint64 = genFn1((!(!(true || (1 <= 4294967295)))), 128, 738)
            _ = v16
            _ = v17
        }
    default:
        v1 += uint8(v5)
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 int = 2147483647
        _ = fi18
        _ = v19
    }
    v4 = append(v4, -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
