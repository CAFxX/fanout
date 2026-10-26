package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint32, p1 int16) int32 {
    p1 -= p1
    fmt.Println(p1, p0)
    if (p0 == p0) {
        var v6 uint32 = (p0 << ((0) & 63))
        _ = v6
    }
    p1 = p1
    _ = p0
    _ = p1
    return 928
}

func main() {
    var v1 string = ("foo bar" + "hello")
    var v2 int8 = 100
    var v3 float64 = 0.0
    v4 := make(map[int]bool)
    var v5 int16 = 814
    switch 42 {
    case -431:
        for fi7 := 0; fi7 < 17; fi7++ {
            if ((genFn0(4294967295, v5) <= 42) || (2 != 255)) {
                var v8 int64 = int64(v2)
                v5 = int16(v2)
                v5 = 42
                _ = v8
            }
            v9 := GenS1{F0: 841, F1: 2, F2: 403}
            v10 := []int{(fi7 - fi7), fi7, fi7, fi7}
            _ = fi7
            _ = v10
            _ = v9
        }
    default:
        var v11 int64 = 814
        _ = v11
    }
    if false {
        var v12 int64 = 9223372036854775807
        v13 := GenS1{F0: 110, F1: 470, F2: 42}
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 528, F1: 4294967295, F2: 2147483647}
    fmt.Println(v2, v5, v3)
    v15 := []int{2147483647}
    if true {
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 float64 = v3
            v18 := GenS1{F0: 810, F1: 2, F2: 42}
            _ = fi16
            _ = v17
            _ = v18
        }
        var v19 int = 2147483647
        for fi20 := 0; fi20 < 3; fi20++ {
            _ = len(v15)
            var v21 int = -410
            v19 += 4294967295
            _ = fi20
            _ = v21
        }
        _ = v19
    } else {
        v22 := GenS2{F0: 1, F1: v1}
        _ = v22
    }
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    var v23 uint8 = 128
    if false {
        if false {
            v5 -= (v5 << ((862) & 63))
        } else {
            var v24 int32 = -2147483648
            for fi25 := 0; fi25 < 17; fi25++ {
                v26 := GenS3{F0: 18446744073709551615}
                switch 0 {
                case -9223372036854775808:
                    v27 := GenS2{F0: v5, F1: ""}
                    _ = v27
                case 32767:
                    if (("foo bar" + "a") <= v1) {
                        switch fi25 {
                        case 1:
                            if (false && (430 >= -9223372036854775808)) {
                                v28 := GenS4{F0: 0, F1: -165}
                                fmt.Println(v5)
                                _ = v28
                            }
                        }
                    }
                case 1:
                    v4[(fi25 & (fi25 % ((fi25) | 1)))] = (v1 != "hello")
                default:
                    v3 -= 0.0
                }
                _ = fi25
                _ = v26
            }
            if len(v15) > 0 {
                v15[0] = 32768
            }
            _ = v24
        }
        if ((v1 > (v1 + (("" + (v1 + "a")) + (v1 + v1)))) || false) {
            v5 = 458
            for fi29 := 0; fi29 < 4; fi29++ {
                for fi30 := 0; fi30 < 11; fi30++ {
                    if false {
                        v31 := GenS3{F0: 205}
                        _ = v31
                    } else {
                        fmt.Println(v23)
                        fmt.Println(fi29)
                        _ = len(v4)
                        if (!false) {
                            var v32 int16 = v5
                            _ = v32
                        }
                        v23 = v23
                        v2 += (v2 >> ((0) & 63))
                        var v33 uint8 = (v23 & 2)
                        _ = v33
                    }
                    _ = fi30
                }
                var v34 int8 = (-(v2))
                v5 += (v5 / (((v5 << ((501) & 63))) | 1))
                _ = fi29
                _ = v34
            }
        }
        v35 := GenS3{F0: 42}
        _ = v35
    }
    v36 := GenS4{F0: -454, F1: 0}
    fmt.Println(v1, v2, v3)
    switch (v1 + v1) {
    case "a":
        if len(v15) > 0 {
            v15[0] = 32768
        }
    }
    v37 := GenS5{F0: 447}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
