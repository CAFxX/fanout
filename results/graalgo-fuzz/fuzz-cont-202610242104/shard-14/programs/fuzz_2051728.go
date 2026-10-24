package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int64) int32 {
    fmt.Println(p0)
    p0 = p0
    _ = p0
    return -88
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 805
    var v3 uint32 = 2147483647
    v4 := make(map[string]bool)
    var v5 uint16 = (^(v2))
    var v6 bool = (217 == 127)
    for fi7 := 0; fi7 < 6; fi7++ {
        v5 = uint16(v3)
        v8 := GenS1{F0: true, F1: v6}
        _ = fi7
        _ = v8
    }
    v2 -= (v5 * v5)
    v9 := v1[9223372036854775807]
    if (132 <= 195) {
        var v10 uint = 42
        v10 -= v10
        v3 = 936
        _ = v10
    } else {
        for fi11 := 0; fi11 < 1; fi11++ {
            var v12 int = fi11
            _ = fi11
            _ = v12
        }
        v13 := v4[((v9 + "a") + (v9 + v9))]
        if (65535 > 2) {
            v13 = (0 <= genFn0(-1))
            var v14 int32 = int32(v3)
            v2 += (v2 * (v5 << ((18446744073709551615) & 63)))
            _ = v14
        } else {
            switch 0 {
            case 32767:
                for fi15 := 0; fi15 < 13; fi15++ {
                    v16 := []int{fi15, (fi15 | fi15), -2147483648}
                    for fi17 := 0; fi17 < 6; fi17++ {
                        var v18 int64 = 0
                        v3 += (v3 | (v3 / (((v3 | 94)) | 1)))
                        _ = fi17
                        _ = v18
                    }
                    _ = fi15
                    _ = v16
                }
            case 227:
                v3 = (v3 << ((1) & 63))
            default:
                v19 := []int{-2147483648, 833}
                _ = v19
            }
        }
        _ = v13
    }
    switch -9223372036854775808 {
    case -649:
        v20 := GenS1{F0: ((v6 && v6) || false), F1: v6}
        _ = v20
    case 333:
        v5 = (v5 >> ((65535) & 63))
    case -71:
        for fi21 := 0; fi21 < 14; fi21++ {
            v22 := v1[fi21]
            for fi23 := 0; fi23 < 12; fi23++ {
                v3 = (v3 / ((2147483647) | 1))
                fmt.Println(fi23)
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
    default:
        v3 = v3
    }
    if (((1 > -40) || v6) && v6) {
        if ((v3 + v3) < (^(v3))) {
            fmt.Println(v2)
        }
        v24 := GenS1{F0: (2147483647 >= -749), F1: v6}
        _ = v24
    }
    switch v9 {
    case "\n\t":
        if (1.5 <= 1.5) {
            var v25 uint32 = 65535
            if (!(v6 && true)) {
                v26 := GenS1{F0: v6, F1: (!(true && v6))}
                for fi27 := 0; fi27 < 9; fi27++ {
                    v5 += (^(v5))
                    fmt.Println(v25, v9, v3)
                    _ = fi27
                }
                _ = v26
            }
            _ = v25
        } else {
            var v28 int8 = -128
            _ = v28
        }
    case "hello":
        var v29 uint16 = v2
        _ = v29
    case "foo bar":
        var v30 uint = uint(v3)
        _ = v30
    }
    if v6 {
        if v6 {
            if ((127 > 2) || (1000 == 1000)) {
                var v31 bool = (!(!(2 != uint64(v3))))
                var v32 uint16 = v5
                v5 = (v32 - (v5 << ((255) & 63)))
                _ = v31
                _ = v32
            }
            for fi33 := 0; fi33 < 15; fi33++ {
                v34 := GenS2{F0: v3}
                v5 -= (v5 * (v2 / (((v5 / (((v2 - v5)) | 1))) | 1)))
                _ = fi33
                _ = v34
            }
        } else {
            if (v6 || false) {
                var v35 int = -9223372036854775808
                _ = v35
            }
            v2 = 255
            var v36 int64 = 42
            _ = v36
        }
    }
    v37 := []int{9223372036854775807, -353}
    v4[(v9 + "foo bar")] = v6
    v2 += v2
    v6 = (-853 > -32768)
    var v38 int16 = 453
    fmt.Println(v9, v5, v5)
    fmt.Println(v2, v38)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
