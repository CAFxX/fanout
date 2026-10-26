package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int64 }
func genFn0(p0 int64, p1 uint64) uint32 {
    v4 := []int{}
    var v5 int16 = -32768
    v4 = append(v4, 4294967295)
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 int = -64
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        switch (v7 / (((fi6 - v7)) | 1)) {
        case -283:
            p1 -= 652
        case 326:
            switch 127 {
            case 682:
                for fi8 := 0; fi8 < 4; fi8++ {
                    v4 = append(v4, fi8)
                    _ = fi8
                }
            case -1:
                for fi9 := 0; fi9 < 10; fi9++ {
                    switch -2147483648 {
                    case -810:
                        v4 = append(v4, (v7 - 118))
                    case -9223372036854775808:
                        var v10 int = -948
                        _ = v10
                    default:
                        var v11 uint32 = 1
                        switch (-(fi9)) {
                        case -433:
                            fmt.Println(v5, v5)
                        case 4294967295:
                            if (("\n\t" <= "a") || true) {
                                v4 = append(v4, fi9)
                            }
                        }
                        var v12 uint8 = 2
                        if false {
                            p0 -= 1
                            var v13 uint64 = p1
                            if false {
                                p0 = (-(p0))
                                fmt.Println(p1, fi6, fi9)
                                if len(v4) > 0 {
                                    v4[0] = fi6
                                }
                            }
                            _ = v13
                        } else {
                            v4 = append(v4, (^(fi6)))
                        }
                        _ = v11
                        _ = v12
                    }
                    if (1 < 127) {
                        p0 -= p0
                        _ = len(v4)
                        v14 := GenS1{F0: 2, F1: 0, F2: 0}
                        _ = v14
                    } else {
                        p1 += p1
                        var v15 uint = 255
                        _ = v15
                    }
                    p1 = 477
                    _ = fi9
                }
            case 485:
                v7 = 43
            default:
                switch (fi6 | v7) {
                case 42:
                    fmt.Println(p1)
                case -2147483648:
                    fmt.Println(v7)
                case 565:
                    for fi16 := 0; fi16 < 7; fi16++ {
                        v5 -= (v5 / (((v5 | (^(v5)))) | 1))
                        _ = fi16
                    }
                default:
                    var v17 uint8 = 1
                    _ = v17
                }
            }
        }
        _ = fi6
        _ = v7
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1(p0 uint64, p1 uint16, p2 uint16) int8 {
    v18 := []int{2147483647, 32768, -9223372036854775808, 127}
    var v19 uint16 = p1
    fmt.Println(v19)
    switch -815 {
    case -2147483648:
        p2 -= p2
    case 4294967295:
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
    case 32768:
        switch "foo bar" {
        case "foo bar":
            if true {
                var v20 float64 = 0.1
                var v21 int8 = -1
                var v22 uint8 = 1
                _ = v20
                _ = v21
                _ = v22
            }
        case "a":
            var v23 string = ""
            _ = v23
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return -128
}

func main() {
    var v1 int16 = 1
    var v2 int64 = 899
    var v3 uint = 255
    var v24 int16 = (v1 >> ((v3) & 63))
    v2 = v2
    var v25 float64 = 1e10
    var v26 float64 = (v25 * v25)
    fmt.Println(v2, v1, v1)
    var v27 int32 = -1
    v28 := []int{-661}
    v24 += v1
    if (false || false) {
        for fi29 := 0; fi29 < 4; fi29++ {
            v30 := GenS1{F0: 0, F1: -531, F2: (^(v2))}
            fmt.Println(v24, v25, v2)
            _ = len(v28)
            _ = fi29
            _ = v30
        }
    }
    v2 = v2
    _ = len(v28)
    var v31 float32 = (0.0 + 1e10)
    if len(v28) > 0 {
        v28[0] = 127
    }
    v32 := GenS1{F0: 65535, F1: 127, F2: (v2 << (((v3 - (v3 / (((^(v3))) | 1)))) & 63))}
    v26 = (v26 + 1e10)
    v2 = (v2 / ((0) | 1))
    var v33 uint32 = 979
    var v34 bool = (v2 > 1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}
