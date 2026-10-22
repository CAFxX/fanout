package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int32; F2 int }
type GenS3 struct { F0 int; F1 uint8 }
func genFn0(p0 uint, p1 int, p2 int8) uint {
    v4 := GenS1{F0: 1}
    var v5 int8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (p0 % (((p0 + (p0 | p0))) | 1))
}

func main() {
    v1 := []int{-2147483648, -1, 697, 9223372036854775807}
    var v2 uint8 = 128
    var v3 int64 = -622
    if false {
        if false {
            if false {
                switch "\n\t" {
                case "a":
                    fmt.Println(v3, v3)
                case "\n\t":
                    if false {
                        if true {
                            fmt.Println(v2)
                        } else {
                            var v6 float32 = 0.1
                            v1 = append(v1, 127)
                            var v7 uint16 = 260
                            _ = v6
                            _ = v7
                        }
                        fmt.Println(v2, v3)
                    }
                case "foo bar":
                    switch "\n\t" {
                    case "":
                        for fi8 := 0; fi8 < 13; fi8++ {
                            fmt.Println(v2, fi8, v2)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v9 := GenS1{F0: 449}
                            _ = fi8
                            _ = v9
                        }
                    case "\n\t":
                        var v10 int64 = v3
                        _ = v10
                    case "foo bar":
                        var v11 int32 = -381
                        _ = v11
                    default:
                        v3 -= (-(v3))
                    }
                default:
                    if true {
                        v2 = (v2 - v2)
                        v2 = v2
                    } else {
                        if ((3.14159 / 0.1) <= (((float64(3.14159) + 1e10) + 0.0) - -2.25)) {
                            for fi12 := 0; fi12 < 16; fi12++ {
                                if (65535 <= 782) {
                                    var v13 string = "a"
                                    var v14 int32 = -687
                                    _ = v13
                                    _ = v14
                                } else {
                                    v3 -= (v3 * -1)
                                }
                                _ = fi12
                            }
                            v15 := GenS2{F0: (!false), F1: 1, F2: 418}
                            _ = v15
                        } else {
                            v3 += (v3 % ((v3) | 1))
                            fmt.Println(v2, v2)
                        }
                        fmt.Println(v2, v2)
                    }
                }
                for fi16 := 0; fi16 < 8; fi16++ {
                    var v17 uint64 = 0
                    for fi18 := 0; fi18 < 9; fi18++ {
                        var v19 uint64 = v17
                        _ = fi18
                        _ = v19
                    }
                    var v20 int8 = 110
                    _ = fi16
                    _ = v17
                    _ = v20
                }
            }
            v3 = v3
            v2 = 112
        }
        fmt.Println(v3, v3)
        v2 = v2
    } else {
        var v21 uint16 = 42
        _ = v21
    }
    v3 = 9223372036854775807
    if (!((v2 - v2) <= 42)) {
        if len(v1) > 0 {
            v1[0] = -9223372036854775808
        }
        var v22 int8 = -26
        _ = v22
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        if (65535 >= 1) {
            var v24 uint8 = (v2 - (v2 << ((42) & 63)))
            _ = v24
        }
        v1 = append(v1, fi23)
        _ = fi23
    }
    fmt.Println(v3, v2)
    for fi25 := 0; fi25 < 19; fi25++ {
        fmt.Println(v3, v2)
        var v26 float64 = 0.0
        _ = fi25
        _ = v26
    }
    var v27 float64 = (((((-2.25 - 1.5) * (-2.25 + -2.25)) * float64(-2.25)) - 1.5) * 3.14159)
    _ = len(v1)
    v28 := GenS3{F0: 192, F1: (v2 | v2)}
    _ = len(v1)
    v1 = append(v1, 588)
    v27 = v27
    var v29 uint8 = v2
    var v30 uint8 = 42
    var v31 int = 1
    for fi32 := 0; fi32 < 13; fi32++ {
        _ = len(v1)
        switch fi32 {
        case -247:
            switch "a" {
            case "\n\t":
                v1 = append(v1, 606)
            default:
                v33 := GenS2{F0: false, F1: 1, F2: (^(fi32))}
                _ = v33
            }
        }
        v30 = v30
        _ = fi32
    }
    if true {
        if (!true) {
            fmt.Println(v27, v31, v30)
            fmt.Println(v30)
            fmt.Println(v30)
        }
        if ((-1 <= -32768) || (genFn0(486, v31, 68) <= uint(v29))) {
            var v34 uint32 = 4294967295
            _ = len(v1)
            _ = v34
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
