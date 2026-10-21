package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 bool; F2 string }
type GenS3 struct { F0 string; F1 uint32; F2 int }
func genFn0(p0 int32) int {
    p0 = p0
    v4 := GenS1{F0: 0}
    fmt.Println(p0)
    _ = p0
    _ = v4
    return 4294967295
}
func genFn1() int64 {
    var v5 int = -560
    var v6 float64 = (0.1 * 1e10)
    if (true || (1 < 255)) {
        fmt.Println(v5, v6)
        for fi7 := 0; fi7 < 15; fi7++ {
            if true {
                v6 = ((((1.5 * (((v6 * v6) + float64(1.5)) - v6)) / 1.0) / v6) / float64(float32((v6 + (v6 + 3.14159)))))
                v8 := []int{0, fi7, 1, (fi7 / (((v5 / (((fi7 + (fi7 & -684))) | 1))) | 1))}
                if (-1 >= -492) {
                    for fi9 := 0; fi9 < 16; fi9++ {
                        var v10 float32 = 1e10
                        if true {
                            _ = len(v8)
                        } else {
                            v5 = genFn0(0)
                        }
                        v6 += (((float64(v10) * v6) - 1e10) - (((v6 + ((1.5 * (v6 + 0.0)) * v6)) - v6) / (1e10 + v6)))
                        if false {
                            fmt.Println(v5, fi9, v10)
                            var v11 uint8 = 155
                            _ = v11
                        }
                        v8 = append(v8, (fi9 | 9223372036854775807))
                        if ((0.1 + v10) > v10) {
                            v6 -= v6
                            v12 := GenS1{F0: 0}
                            switch genFn0(1) {
                            case -1:
                                v13 := GenS1{F0: 1}
                                _ = v13
                            case 352:
                                if false {
                                    if len(v8) > 0 {
                                        v8[0] = (fi9 * (v5 << ((903) & 63)))
                                    }
                                    fmt.Println(fi7, v5)
                                    var v14 int = 0
                                    if (-9223372036854775808 > 42) {
                                        v8 = append(v8, (-(v14)))
                                        switch "" {
                                        case "\n\t":
                                            v15 := GenS2{F0: 108, F1: true, F2: "hello"}
                                            var v16 uint8 = 163
                                            var v17 float32 = v10
                                            _ = v15
                                            _ = v16
                                            _ = v17
                                        default:
                                            var v18 int16 = 1000
                                            _ = v18
                                        }
                                        var v19 uint32 = 0
                                        _ = v19
                                    }
                                    if len(v8) > 0 {
                                        v8[0] = 2147483647
                                    }
                                    _ = v14
                                }
                            case 9223372036854775807:
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                            default:
                                var v20 int = genFn0(-1)
                                _ = v20
                            }
                            _ = v12
                        }
                        _ = fi9
                        _ = v10
                    }
                }
                _ = v8
            }
            _ = fi7
        }
        var v21 uint32 = 2147483648
        _ = v21
    } else {
        v22 := GenS3{F0: ("" + ("\n\t" + "")), F1: 2, F2: v5}
        _ = v22
    }
    fmt.Println(v6)
    _ = v5
    _ = v6
    return 517
}

func main() {
    var v1 string = "a"
    var v2 uint32 = 0
    v3 := []uint32{}
    var v23 int64 = 0
    v24 := GenS3{F0: v1, F1: v2, F2: 604}
    var v25 int = 9223372036854775807
    var v26 float64 = 1.5
    switch v25 {
    case -2147483648:
        v27 := GenS2{F0: (^(v23)), F1: (("" + "foo bar") < "hello"), F2: v1}
        _ = v27
    }
    fmt.Println(v26, v1)
    var v28 uint = 2
    fmt.Println(v1)
    var v29 float32 = float32(float64(0.1))
    var v30 float32 = (3.14159 + v29)
    for fi31 := 0; fi31 < 19; fi31++ {
        var v32 int8 = 42
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 6; fi33++ {
        var v34 float32 = (0.1 - float32(3.14159))
        _ = fi33
        _ = v34
    }
    v29 += v29
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
