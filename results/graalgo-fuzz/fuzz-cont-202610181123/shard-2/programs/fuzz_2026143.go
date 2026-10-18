package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 uint32; F1 string }
func genFn0(p0 uint, p1 uint32) int {
    p1 += p1
    if (4294967295 == (p0 ^ p0)) {
        v7 := GenS1{F0: 127, F1: 659, F2: 519}
        _ = v7
    } else {
        p1 += p1
        v8 := GenS2{F0: 2, F1: p1}
        var v9 int64 = 9223372036854775807
        _ = v8
        _ = v9
    }
    v10 := []int{}
    _ = p0
    _ = p1
    _ = v10
    return -2147483648
}
func genFn1(p0 uint8, p1 uint64, p2 uint) string {
    p0 -= p0
    switch 32767 {
    case 9223372036854775807:
        var v11 int16 = -1
        _ = v11
    default:
        v12 := []int{-263}
        _ = v12
    }
    v13 := []int{805, 638}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return ""
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]bool)
    var v3 uint64 = 38
    var v4 uint32 = 42
    var v5 uint8 = 0
    var v6 uint8 = (v5 * 0)
    switch 2147483647 {
    case 32767:
        for fi14 := 0; fi14 < 7; fi14++ {
            if false {
                switch "foo bar" {
                case "a":
                    v5 = v5
                default:
                    fmt.Println(fi14, v3, v5)
                }
                v4 -= 728
                v6 -= (v5 << ((18446744073709551615) & 63))
            }
            v15 := GenS2{F0: (v6 ^ (^(v5))), F1: (v4 + (^(v4)))}
            _ = fi14
            _ = v15
        }
    case 397:
        var v16 int32 = 0
        _ = v16
    case 4294967295:
        v6 = (v5 & 1)
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        if len(v1) > 0 {
            v1[0] = (v4 * (v4 + v4))
        }
        _ = fi17
    }
    v6 += 2
    var v18 uint8 = 1
    switch 42 {
    case 32767:
        v19 := GenS3{F0: (((("" + ("\n\t" + ("a" + "a"))) + "a") + "a") + ("foo bar" + ("" + ("" + "foo bar"))))}
        _ = v19
    }
    if ("\n\t" > "\n\t") {
        fmt.Println(v6, v4, v6)
    } else {
        for fi20 := 0; fi20 < 17; fi20++ {
            for fi21 := 0; fi21 < 13; fi21++ {
                var v22 uint = 733
                _ = fi21
                _ = v22
            }
            var v23 uint8 = (v6 % ((128) | 1))
            _ = fi20
            _ = v23
        }
        v2[v4] = (v6 >= 0)
        v6 += (v6 / (((^(v18))) | 1))
    }
    var v24 float64 = -2.25
    switch -1 {
    case -921:
        var v25 uint16 = 32767
        _ = v25
    case 127:
        v4 = (v4 - (v4 + (v4 << ((0) & 63))))
    case 1:
        fmt.Println(v6, v3)
    }
    var v26 float64 = v24
    v5 = 127
    v27 := v2[2147483648]
    for fi28 := 0; fi28 < 2; fi28++ {
        if (2 <= 522) {
            for fi29 := 0; fi29 < 20; fi29++ {
                if v27 {
                    v24 += v24
                    var v30 uint = 139
                    _ = v30
                } else {
                    switch "" {
                    case "hello":
                        v31 := v2[v4]
                        if v31 {
                            if false {
                                var v32 int64 = 0
                                v33 := GenS4{F0: 0, F1: 0}
                                if len(v1) > 0 {
                                    v1[0] = 65535
                                }
                                _ = v32
                                _ = v33
                            }
                        } else {
                            var v34 float64 = v24
                            v35 := GenS2{F0: 1, F1: 735}
                            _ = v34
                            _ = v35
                        }
                        v36 := GenS5{F0: v4, F1: "hello"}
                        switch "\n\t" {
                        case "\n\t":
                            v37 := GenS2{F0: v6, F1: 976}
                            _ = v37
                        case "foo bar":
                            var v38 uint16 = 295
                            _ = v38
                        default:
                            _ = len(v1)
                        }
                        _ = v31
                        _ = v36
                    }
                }
                _ = fi29
            }
        } else {
            var v39 int32 = 869
            switch "hello" {
            case "foo bar":
                v40 := GenS2{F0: (v5 ^ v5), F1: v4}
                _ = v40
            case "":
                var v41 int16 = 747
                _ = v41
            }
            _ = v39
        }
        _ = fi28
    }
    v27 = (v27 || (255 == (^(v18))))
    var v42 uint32 = (^(v4))
    for fi43 := 0; fi43 < 14; fi43++ {
        var v44 float32 = (-2.25 + 0.0)
        _ = len(v2)
        _ = fi43
        _ = v44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
