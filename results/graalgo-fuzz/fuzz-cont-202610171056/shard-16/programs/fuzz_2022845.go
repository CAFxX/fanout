package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 string }
type GenS2 struct { F0 int16; F1 uint32 }
func genFn0(p0 int16) int64 {
    fmt.Println(p0)
    switch -1 {
    case -982:
        for fi4 := 0; fi4 < 4; fi4++ {
            fmt.Println(fi4, p0)
            _ = fi4
        }
    default:
        for fi5 := 0; fi5 < 14; fi5++ {
            fmt.Println(fi5, p0)
            p0 += (p0 / (((p0 & (p0 << ((42) & 63)))) | 1))
            _ = fi5
        }
    }
    var v6 int32 = 42
    var v7 int16 = 1000
    v7 = 1
    _ = p0
    _ = v6
    _ = v7
    return -1
}
func genFn1() int {
    if (680 < 255) {
        var v8 int64 = 9223372036854775807
        _ = v8
    }
    if (1 != 65535) {
        var v9 float64 = 0.1
        v10 := GenS1{F0: 1, F1: "a", F2: "hello"}
        _ = v10
        _ = v9
    } else {
        v11 := []int{127, -2147483648}
        var v12 int64 = 193
        fmt.Println(v12)
        _ = v11
        _ = v12
    }
    return 72
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{-143}
    var v3 float32 = float32(1.5)
    v3 = v3
    v3 = float32(0.0)
    v3 = -2.25
    delete(v1, -2147483648)
    if (179 > 2) {
        v3 = v3
    } else {
        v13 := v1[-2147483648]
        v3 -= v3
        _ = v13
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 float32 = -2.25
        switch "a" {
        case "a":
            v16 := GenS1{F0: 0, F1: (("\n\t" + "a") + ("foo bar" + ("foo bar" + "a"))), F2: ""}
            _ = v16
        case "foo bar":
            v15 += (v3 / v15)
        default:
            if (2 <= 2) {
                if (255 < 762) {
                    v15 += ((1.5 + (v3 + (0.0 - v3))) - float32(3.14159))
                    v17 := GenS2{F0: 0, F1: 552}
                    _ = v17
                } else {
                    fmt.Println(v3, fi14, fi14)
                    switch -535 {
                    case 2147483647:
                        var v18 int8 = -1
                        _ = v18
                    }
                }
                if (873 == 0) {
                    if false {
                        fmt.Println(v3, v15)
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                if (65535 != 4294967295) {
                    if (127 != -128) {
                        for fi19 := 0; fi19 < 5; fi19++ {
                            var v20 int16 = -443
                            var v21 int64 = 390
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        v3 = (v3 + (v15 * 1e10))
                    }
                } else {
                    for fi22 := 0; fi22 < 10; fi22++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi22
                    }
                    v15 = (float32(1e10) / 1.5)
                    _ = len(v2)
                }
            } else {
                v15 = v15
            }
        }
        _ = fi14
        _ = v15
    }
    v3 -= v3
    var v23 float64 = 3.14159
    fmt.Println(v23)
    fmt.Println(v3, v23)
    v23 += ((1e10 + (v23 / 1.5)) / (v23 * v23))
    switch 42 {
    case 42:
        v24 := GenS1{F0: 2, F1: "hello", F2: "a"}
        _ = v24
    case -842:
        fmt.Println(v3, v3)
    case -9223372036854775808:
        if true {
            if (!((-1 == 1) && false)) {
                v23 = 0.0
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
}
