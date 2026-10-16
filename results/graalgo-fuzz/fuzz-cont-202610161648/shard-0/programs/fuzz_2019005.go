package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int8 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint32; F1 int64; F2 bool }
type GenS4 struct { F0 int8; F1 int32 }

func main() {
    var v1 uint8 = 42
    v2 := []uint32{}
    var v3 float64 = 1e10
    v4 := make(map[int]bool)
    if len(v2) > 0 {
        v2[0] = 619
    }
    if false {
        v4[2147483647] = false
        v5 := GenS1{F0: 4294967295, F1: -9223372036854775808, F2: 25}
        _ = v5
    } else {
        var v6 int8 = -104
        v3 = (v3 * v3)
        _ = v6
    }
    switch "hello" {
    case "hello":
        v3 -= (1.5 / float64(1e10))
    case "\n\t":
        for fi7 := 0; fi7 < 19; fi7++ {
            fmt.Println(v1, v1)
            var v8 uint8 = (^(v1))
            _ = fi7
            _ = v8
        }
    }
    if len(v2) > 0 {
        v2[0] = 2
    }
    switch "hello" {
    case "hello":
        if (!(527 == 18446744073709551615)) {
            v9 := GenS1{F0: 2, F1: -9223372036854775808, F2: 127}
            _ = v9
        } else {
            v1 = (v1 * (v1 & 66))
            switch 681 {
            case 128:
                switch "hello" {
                case "":
                    fmt.Println(v3)
                case "hello":
                    if (358 < 282) {
                        fmt.Println(v3, v3)
                        switch 422 {
                        case -2147483648:
                            v4[127] = (18446744073709551615 >= 687)
                        case 32768:
                            v3 -= (v3 + 1e10)
                        case 9223372036854775807:
                            if len(v2) > 0 {
                                v2[0] = 65535
                            }
                        default:
                            var v10 uint = 42
                            _ = v10
                        }
                        v1 = (v1 * (v1 + (v1 >> ((18446744073709551615) & 63))))
                    }
                case "a":
                    fmt.Println(v3)
                default:
                    if (true && ((4294967295 == 344) || true)) {
                        var v11 int = -734
                        var v12 float32 = 1e10
                        _ = v11
                        _ = v12
                    } else {
                        if false {
                            v1 = v1
                            v13 := GenS1{F0: 422, F1: -9223372036854775808, F2: 65}
                            _ = v13
                        }
                        var v14 uint8 = 0
                        if (311 == 707) {
                            v2 = append(v2, 1)
                            if false {
                                v3 = 1e10
                            } else {
                                var v15 uint8 = v14
                                _ = v15
                            }
                        } else {
                            v16 := GenS2{F0: 100, F1: 42}
                            v14 = v14
                            _ = v16
                        }
                        _ = v14
                    }
                }
            default:
                var v17 float64 = v3
                _ = v17
            }
            v1 += (v1 + (v1 + v1))
        }
    case "":
        v3 -= v3
    }
    if true {
        fmt.Println(v3, v3)
    }
    switch 1 {
    case -9223372036854775808:
        var v18 int32 = -1
        _ = v18
    case 385:
        v19 := GenS1{F0: 990, F1: -9223372036854775808, F2: 0}
        _ = v19
    case 4294967295:
        v20 := GenS1{F0: 58, F1: 0, F2: -95}
        _ = v20
    default:
        fmt.Println(v3, v1)
    }
    switch "hello" {
    case "hello":
        var v21 bool = ((2 > 2) || (uint(v1) <= 255))
        _ = v21
    default:
        switch 266 {
        case 747:
            var v22 uint64 = 487
            _ = v22
        }
    }
    var v23 int32 = -1
    switch 32768 {
    case 0:
        v24 := v4[32767]
        _ = v24
    default:
        _ = len(v2)
    }
    var v25 string = "hello"
    if (!(0 >= -24)) {
        for fi26 := 0; fi26 < 1; fi26++ {
            if true {
                v1 -= (^(v1))
                v27 := GenS3{F0: 278, F1: -274, F2: true}
                _ = v27
            }
            _ = fi26
        }
        if (!(!false)) {
            for fi28 := 0; fi28 < 1; fi28++ {
                v23 += (v23 * (v23 << ((42) & 63)))
                _ = fi28
            }
        }
    }
    if true {
        v1 += v1
        var v29 int16 = 0
        _ = v29
    } else {
        switch 4294967295 {
        case -528:
            v1 = v1
        case -633:
            if len(v2) > 0 {
                v2[0] = 1
            }
        default:
            v30 := GenS4{F0: -1, F1: (v23 * v23)}
            _ = v30
        }
    }
    var v31 int32 = v23
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v31
    _ = v4
}
