package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int8, p1 uint64) int64 {
    var v5 uint32 = 2
    v5 = v5
    v5 = (^(v5))
    v5 = 545
    var v6 uint16 = 2
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 int64) uint16 {
    v7 := []int{}
    var v8 int = -2147483648
    var v9 int32 = -813
    p0 = -9223372036854775808
    if ((!false) && (2147483647 > 2147483647)) {
        switch -228 {
        case 4294967295:
            if len(v7) > 0 {
                v7[0] = v8
            }
        }
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 595
}

func main() {
    var v1 uint16 = 2
    var v2 float64 = (0.1 / (-2.25 - 1.5))
    var v3 float64 = ((v2 / float64(-2.25)) - 1e10)
    v4 := []uint32{785, 0}
    switch "foo bar" {
    case "":
        v10 := GenS1{F0: 128, F1: -743, F2: 12}
        _ = v10
    case "\n\t":
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
    case "hello":
        var v11 int32 = 877
        _ = v11
    }
    var v12 uint32 = 1
    if len(v4) > 0 {
        v4[0] = v12
    }
    var v13 bool = (false || (((!true) || true) && false))
    var v14 string = ("foo bar" + "")
    _ = len(v4)
    v14 = "hello"
    var v15 uint64 = 4294967295
    fmt.Println(v1, v1)
    v12 += (v12 | v12)
    fmt.Println(v14, v15)
    if (v14 <= v14) {
        switch (v14 + (v14 + ("a" + v14))) {
        case "":
            for fi16 := 0; fi16 < 19; fi16++ {
                v12 += (v12 % ((2) | 1))
                if ((^(v15)) != v15) {
                    switch fi16 {
                    case 226:
                        var v17 int32 = -1
                        _ = v17
                    case 660:
                        v18 := GenS2{F0: (^(v15))}
                        _ = v18
                    default:
                        switch (v14 + v14) {
                        case "":
                            v12 = (v12 ^ v12)
                        default:
                            switch 2147483647 {
                            case 32768:
                                v14 = v14
                            case -2147483648:
                                var v19 uint8 = 11
                                _ = v19
                            case 402:
                                var v20 uint8 = uint8(v15)
                                _ = v20
                            default:
                                for fi21 := 0; fi21 < 9; fi21++ {
                                    v15 -= v15
                                    _ = len(v4)
                                    var v22 string = v14
                                    _ = fi21
                                    _ = v22
                                }
                            }
                        }
                    }
                    _ = len(v4)
                } else {
                    var v23 uint16 = (^(v1))
                    if v13 {
                        var v24 int64 = 42
                        fmt.Println(v2, v14, v13)
                        for fi25 := 0; fi25 < 18; fi25++ {
                            v4 = append(v4, (v12 ^ (^(v12))))
                            _ = len(v4)
                            _ = fi25
                        }
                        _ = v24
                    }
                    if (0.1 != (0.0 - 3.14159)) {
                        v26 := GenS2{F0: v15}
                        _ = v26
                    }
                    _ = v23
                }
                _ = fi16
            }
        case "\n\t":
            if ((v1 ^ 0) != uint16(v15)) {
                if ((655 >= 0) || false) {
                    v2 = v3
                }
                if len(v4) > 0 {
                    v4[0] = 65535
                }
            }
        case "foo bar":
            var v27 int64 = -107
            _ = v27
        default:
            fmt.Println(v3, v3)
        }
        if (((610 < int32(v12)) && v13) || (-827 != int16(v15))) {
            if false {
                v28 := GenS2{F0: (v15 * (v15 % ((v15) | 1)))}
                v13 = v13
                _ = v28
            }
            if (!v13) {
                fmt.Println(v14, v1, v14)
                for fi29 := 0; fi29 < 1; fi29++ {
                    var v30 int64 = 0
                    for fi31 := 0; fi31 < 17; fi31++ {
                        v32 := GenS1{F0: 248, F1: -1, F2: (^(v15))}
                        _ = fi31
                        _ = v32
                    }
                    _ = fi29
                    _ = v30
                }
                v14 = "foo bar"
            }
        }
        v33 := GenS2{F0: (v15 ^ v15)}
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
