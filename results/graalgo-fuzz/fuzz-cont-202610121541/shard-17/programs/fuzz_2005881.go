package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 string; F1 int64; F2 bool }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 uint8; F1 int8; F2 string }
type GenS7 struct { F0 int16; F1 uint8 }

func main() {
    v1 := []int{9223372036854775807, -658, 1}
    var v2 int = 270
    var v3 uint64 = 990
    var v4 uint16 = 42
    v5 := []string{}
    v6 := GenS1{F0: 527, F1: 647}
    v4 = 255
    for fi7 := 0; fi7 < 19; fi7++ {
        switch ("a" + "\n\t") {
        case "":
            if true {
                for fi8 := 0; fi8 < 3; fi8++ {
                    v9 := GenS1{F0: 542, F1: 87}
                    _ = fi8
                    _ = v9
                }
                fmt.Println(v3, fi7, v3)
                var v10 float64 = 0.1
                _ = v10
            } else {
                if (true && (true || (!(1 > 2147483647)))) {
                    if false {
                        if ((0.1 * 0.0) == 3.14159) {
                            if (-6 <= -128) {
                                for fi11 := 0; fi11 < 7; fi11++ {
                                    v12 := GenS1{F0: 2147483647, F1: 42}
                                    _ = fi11
                                    _ = v12
                                }
                                v4 = (v4 * 42)
                            }
                            v13 := GenS1{F0: -2147483648, F1: 669}
                            var v14 uint64 = (^(v3))
                            _ = v13
                            _ = v14
                        } else {
                            fmt.Println(v3, v3, v4)
                            v4 += 898
                            v15 := GenS1{F0: int32(v3), F1: 2}
                            _ = v15
                        }
                        var v16 float32 = 1.5
                        if len(v1) > 0 {
                            v1[0] = -2147483648
                        }
                        _ = v16
                    }
                    switch "hello" {
                    case "\n\t":
                        if len(v1) > 0 {
                            v1[0] = v2
                        }
                    case "a":
                        var v17 uint64 = (v3 + (^(v3)))
                        _ = v17
                    }
                }
                var v18 float32 = (-2.25 * (-2.25 + 1.5))
                var v19 int64 = 0
                _ = v18
                _ = v19
            }
        default:
            if false {
                v4 += (v4 % (((v4 + (v4 + (v4 << ((42) & 63))))) | 1))
                fmt.Println(v4, v3, fi7)
                fmt.Println(v2, v3, v3)
            } else {
                v4 = (v4 % (((v4 & (v4 * v4))) | 1))
            }
        }
        for fi20 := 0; fi20 < 20; fi20++ {
            v21 := GenS1{F0: 0, F1: 2}
            _ = fi20
            _ = v21
        }
        v3 -= v3
        _ = fi7
    }
    v22 := GenS2{F0: 0, F1: v3}
    var v23 float64 = (float64((0.1 - 1e10)) + (((-2.25 + ((1e10 / 1.5) / 1.0)) - (float64(0.0) / 1.5)) - 0.1))
    v4 += v4
    v2 += (-(v2))
    v23 = (v23 / v23)
    if false {
        fmt.Println(v4, v23, v2)
        if len(v1) > 0 {
            v1[0] = v2
        }
        switch "foo bar" {
        case "hello":
            if ((3.14159 + (1e10 + (float64(-2.25) + v23))) <= float64((((1e10 * ((((1.5 / 1e10) / 0.1) * 1.5) - (3.14159 * (float32(v23) * 0.1)))) - float32((float64(-2.25) + v23))) / 1.0))) {
                if true {
                    switch v2 {
                    case 1:
                        var v24 int8 = 127
                        _ = v24
                    }
                } else {
                    fmt.Println(v23)
                    fmt.Println(v2, v23)
                    var v25 float32 = (1.5 / 1.0)
                    _ = v25
                }
                for fi26 := 0; fi26 < 7; fi26++ {
                    var v27 uint16 = v4
                    _ = fi26
                    _ = v27
                }
                v28 := GenS1{F0: -1, F1: 211}
                _ = v28
            }
        case "foo bar":
            var v29 string = ""
            _ = v29
        case "a":
            var v30 int32 = 42
            _ = v30
        }
    } else {
        var v31 uint8 = 42
        for fi32 := 0; fi32 < 8; fi32++ {
            if false {
                for fi33 := 0; fi33 < 19; fi33++ {
                    v34 := GenS3{F0: "\n\t", F1: 9223372036854775807, F2: (1 > 2)}
                    v31 += v31
                    var v35 uint64 = 0
                    fmt.Println(fi32, v35, v35)
                    var v36 int32 = 945
                    var v37 uint64 = (v3 | 904)
                    _ = fi33
                    _ = v34
                    _ = v35
                    _ = v36
                    _ = v37
                }
                fmt.Println(fi32)
            }
            var v38 uint64 = (v3 & v3)
            _ = fi32
            _ = v38
        }
        switch 1 {
        case 575:
            v23 = (1e10 * (-2.25 * v23))
        default:
            fmt.Println(v31, v2)
        }
        _ = v31
    }
    if false {
        switch "" {
        case "":
            var v39 uint8 = 191
            _ = v39
        case "\n\t":
            v3 = (v3 << ((65535) & 63))
        }
        var v40 uint32 = 1
        _ = v40
    } else {
        v23 -= -2.25
    }
    switch ("" + "foo bar") {
    case "\n\t":
        for fi41 := 0; fi41 < 17; fi41++ {
            var v42 int32 = -1
            var v43 bool = ((0.0 - 1e10) > 0.1)
            _ = fi41
            _ = v42
            _ = v43
        }
    case "":
        v4 -= (v4 % (((v4 - 0)) | 1))
    case "hello":
        v44 := GenS1{F0: 0, F1: 2}
        _ = v44
    default:
        v3 += v3
    }
    v45 := GenS4{F0: 704, F1: "hello"}
    for fi46 := 0; fi46 < 19; fi46++ {
        var v47 uint32 = 42
        v4 += (v4 >> ((uint(v4)) & 63))
        for fi48 := 0; fi48 < 11; fi48++ {
            v3 = 293
            if false {
                switch "" {
                case "":
                    var v49 bool = (0.1 < 1.5)
                    _ = v49
                default:
                    var v50 uint64 = 18446744073709551615
                    _ = v50
                }
                _ = len(v5)
                v3 += (^(v3))
                var v51 int32 = -2147483648
                _ = v51
            } else {
                v52 := GenS5{F0: 104, F1: 1}
                if true {
                    v2 -= (fi48 + fi46)
                } else {
                    var v53 uint32 = uint32(v3)
                    v54 := GenS5{F0: 127, F1: -2}
                    _ = v53
                    _ = v54
                }
                switch 2147483647 {
                case 42:
                    fmt.Println(fi48)
                default:
                    v23 = (v23 * (v23 - 0.0))
                }
                _ = v52
            }
            _ = fi48
        }
        _ = fi46
        _ = v47
    }
    if (!true) {
        if (0.1 >= (((((3.14159 * 1.5) - 0.1) - (float32(1e10) * 0.0)) / (1e10 * 1.5)) - 0.1)) {
            v55 := GenS6{F0: 42, F1: 127, F2: "foo bar"}
            v56 := GenS4{F0: 0, F1: "\n\t"}
            for fi57 := 0; fi57 < 12; fi57++ {
                v58 := GenS7{F0: -644, F1: 127}
                v59 := GenS7{F0: -9, F1: 217}
                v60 := GenS4{F0: -1, F1: "hello"}
                _ = fi57
                _ = v58
                _ = v59
                _ = v60
            }
            _ = v55
            _ = v56
        } else {
            for fi61 := 0; fi61 < 8; fi61++ {
                v2 -= int(v4)
                v1 = append(v1, 32767)
                for fi62 := 0; fi62 < 12; fi62++ {
                    v4 = 552
                    if len(v1) > 0 {
                        v1[0] = v2
                    }
                    v2 -= fi62
                    _ = fi62
                }
                _ = fi61
            }
        }
        var v63 uint = 408
        switch "hello" {
        case "foo bar":
            if true {
                v64 := GenS4{F0: 0, F1: ("foo bar" + "foo bar")}
                _ = v64
            }
        case "":
            v3 -= (^(v3))
        default:
            v65 := GenS7{F0: 1000, F1: 42}
            _ = v65
        }
        _ = v63
    } else {
        var v66 uint64 = 42
        _ = v66
    }
    if false {
        v3 = (v3 - v3)
        for fi67 := 0; fi67 < 14; fi67++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi67
        }
        switch -10 {
        case 381:
            v2 = (v2 ^ v2)
        }
    }
    if false {
        if (-128 <= -128) {
            var v68 int64 = 9223372036854775807
            v69 := GenS1{F0: 1, F1: 2147483648}
            _ = v68
            _ = v69
        }
        fmt.Println(v23)
        fmt.Println(v4, v3)
    }
    var v70 string = "a"
    v5 = append(v5, v70)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v70
}
