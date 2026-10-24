package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint64; F1 int64; F2 string }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint32; F1 uint8 }
func genFn0() uint {
    var v6 uint64 = 1
    v6 = v6
    var v7 uint = 65535
    _ = v6
    _ = v7
    return (v7 + (v7 * (v7 & v7)))
}
func genFn1(p0 int16, p1 int, p2 uint32) bool {
    p2 += (p2 % ((p2) | 1))
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 int16 = p0
        var v10 float32 = 1.5
        p1 = -502
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return (!(64 == 128))
}

func main() {
    var v1 uint32 = 136
    var v2 int16 = -705
    v3 := []int{-122, 32767, 0}
    v4 := make(map[string]string)
    v5 := make(map[int]string)
    if true {
        fmt.Println(v2, v1)
        if (true || false) {
            _ = len(v3)
            var v11 int8 = -128
            _ = v11
        } else {
            v2 -= v2
            v12 := GenS1{F0: 112, F1: 697}
            _ = v12
        }
        switch -9223372036854775808 {
        case 279:
            v13 := GenS1{F0: 42, F1: 2147483647}
            _ = v13
        case -2147483648:
            if ((v1 | 272) == (^(v1))) {
                var v14 float64 = 3.14159
                var v15 bool = true
                v16 := GenS2{F0: v15, F1: (v2 >> ((2) & 63))}
                _ = v14
                _ = v15
                _ = v16
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v17 := GenS1{F0: 42, F1: 0}
                v18 := v5[4294967295]
                _ = v17
                _ = v18
            }
        case 231:
            if (32767 == (v2 & (v2 + (v2 * (v2 | v2))))) {
                if (-2147483648 == 958) {
                    v2 += (v2 | v2)
                    var v19 int = 2147483647
                    for fi20 := 0; fi20 < 1; fi20++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi20
                    }
                    _ = v19
                } else {
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                    if (true || ((-9223372036854775808 > 2147483647) && (402 != 42))) {
                        v2 -= v2
                        v1 -= v1
                    }
                }
                if true {
                    fmt.Println(v2, v1)
                    fmt.Println(v1, v1)
                }
            }
        }
    }
    v5[-515] = "foo bar"
    var v21 int64 = -1
    v1 = v1
    if (-1 >= -40) {
        v22 := GenS3{F0: 2, F1: (v21 / (((^(v21))) | 1)), F2: (("" + "\n\t") + "foo bar")}
        v1 += (v1 ^ v1)
        _ = v22
    } else {
        v1 += 2147483647
        var v23 int8 = 101
        for fi24 := 0; fi24 < 6; fi24++ {
            v25 := GenS1{F0: 65535, F1: 2147483647}
            v26 := GenS2{F0: (!(1.5 < (1e10 + float64(-2.25)))), F1: v2}
            if true {
                v27 := GenS4{F0: (v21 + (v21 << ((genFn0()) & 63))), F1: 2}
                var v28 int64 = -9223372036854775808
                v21 -= (v28 / ((377) | 1))
                _ = v27
                _ = v28
            } else {
                if ((0 >= 18446744073709551615) || (false && (26 == (v23 << ((genFn0()) & 63))))) {
                    v21 -= (v21 >> ((512) & 63))
                    v21 -= v21
                }
            }
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = v23
    }
    var v29 int32 = -1
    if (!(!(2 != 255))) {
        v30 := GenS4{F0: -1, F1: uint16(v1)}
        _ = v30
    } else {
        v31 := GenS5{F0: v1, F1: 1}
        v5[2147483647] = ("" + "")
        _ = v31
    }
    fmt.Println(v21, v2, v1)
    if (32767 > 2) {
        var v32 uint8 = 42
        delete(v5, 9223372036854775807)
        _ = v32
    } else {
        var v33 int16 = v2
        v33 = 0
        _ = v33
    }
    var v34 string = "hello"
    var v35 int32 = (v29 & v29)
    fmt.Println(v34)
    v36 := GenS4{F0: v21, F1: 699}
    var v37 uint = 1
    var v38 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
