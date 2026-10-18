package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint64; F1 int16 }
type GenS10 struct { F0 int32; F1 uint32; F2 int16 }

func main() {
    v1 := []int{-9223372036854775808, 2147483647, -708}
    var v2 int64 = -1
    v3 := make(map[string]uint32)
    var v4 uint8 = 255
    var v5 int16 = -32768
    var v6 int64 = -1
    v6 = (v6 - 9223372036854775807)
    fmt.Println(v2, v5, v5)
    for fi7 := 0; fi7 < 6; fi7++ {
        _ = len(v3)
        fmt.Println(v6, v2)
        _ = fi7
    }
    v8 := GenS1{F0: v4, F1: ((!(!(127 == -128))) && (!(0 > 127))), F2: 255}
    if false {
        v4 -= (v4 | v4)
        fmt.Println(v2, v5, v6)
        var v9 uint = 255
        _ = v9
    }
    v10 := GenS1{F0: v4, F1: true, F2: 32767}
    for fi11 := 0; fi11 < 12; fi11++ {
        v1 = append(v1, fi11)
        v12 := v3[""]
        if (42 <= 127) {
            v13 := GenS1{F0: 128, F1: false, F2: 255}
            var v14 int32 = -231
            var v15 uint16 = 242
            _ = v13
            _ = v14
            _ = v15
        } else {
            var v16 int = (fi11 % (((fi11 >> ((375) & 63))) | 1))
            for fi17 := 0; fi17 < 18; fi17++ {
                if true {
                    v18 := GenS2{F0: (v12 - (v12 | (v12 - (v12 | v12))))}
                    _ = v18
                } else {
                    v3[""] = (v12 << ((1) & 63))
                    v2 -= v6
                }
                _ = fi17
            }
            if false {
                switch (fi11 & (v16 & -1)) {
                case -9223372036854775808:
                    if (!false) {
                        var v19 uint16 = 1
                        var v20 float32 = 1e10
                        if false {
                            delete(v3, "a")
                        }
                        _ = v19
                        _ = v20
                    } else {
                        for fi21 := 0; fi21 < 18; fi21++ {
                            if false {
                                v12 -= 65535
                                v12 -= (^(v12))
                                var v22 uint16 = 947
                                _ = v22
                            } else {
                                fmt.Println(v16, v5)
                                if ((true || (0 >= 255)) && (18446744073709551615 != 0)) {
                                    switch "hello" {
                                    case "a":
                                        if (42 > 42) {
                                            switch "a" {
                                            case "\n\t":
                                                switch (fi11 | -9223372036854775808) {
                                                case -9223372036854775808:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                switch (("\n\t" + (((("a" + (("" + "") + "\n\t")) + "hello") + "") + "\n\t")) + "") {
                                                case "a":
                                                    var v23 uint16 = 32767
                                                    _ = v23
                                                case "foo bar":
                                                    v24 := GenS3{F0: v5}
                                                    _ = v24
                                                case "":
                                                    var v25 string = ""
                                                    _ = v25
                                                }
                                            case "a":
                                                v2 = v6
                                            default:
                                                if len(v1) > 0 {
                                                    v1[0] = fi21
                                                }
                                            }
                                            v1 = append(v1, -2147483648)
                                        } else {
                                            v5 = (-(v5))
                                            v26 := GenS4{F0: (v5 % (((v5 / ((-575) | 1))) | 1))}
                                            _ = v26
                                        }
                                    }
                                    v4 = v4
                                    var v27 float32 = 1e10
                                    switch (v16 ^ v16) {
                                    case -9223372036854775808:
                                        v28 := GenS1{F0: v4, F1: true, F2: uint16(v4)}
                                        _ = v28
                                    case 1:
                                        var v29 int32 = -31
                                        _ = v29
                                    }
                                    v6 = (-(v6))
                                    if (0.0 <= (0.0 + (1.5 / 1e10))) {
                                        var v30 uint8 = v4
                                        _ = v30
                                    }
                                    _ = v27
                                } else {
                                    v1 = append(v1, (-(v16)))
                                    v31 := GenS5{F0: 772}
                                    fmt.Println(v2)
                                    var v32 string = "foo bar"
                                    var v33 uint64 = 474
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            fmt.Println(v2, v6)
                            v34 := GenS2{F0: v12}
                            _ = fi21
                            _ = v34
                        }
                        fmt.Println(fi11)
                    }
                case -855:
                    if (!((-1 <= 2147483647) || ((1 > 255) || false))) {
                        var v35 int32 = 2147483647
                        var v36 bool = true
                        _ = v35
                        _ = v36
                    }
                default:
                    v6 += 0
                }
                for fi37 := 0; fi37 < 1; fi37++ {
                    v4 -= v4
                    _ = fi37
                }
                var v38 int16 = v5
                _ = v38
            } else {
                if (!(false && (-1 == (v5 >> ((42) & 63))))) {
                    if true {
                        var v39 int8 = -128
                        if true {
                            switch (("\n\t" + "\n\t") + "a") {
                            case "":
                                if false {
                                    v16 -= int(v2)
                                    v5 = 1
                                    switch "" {
                                    case "a":
                                        v40 := GenS6{F0: (v39 - (v39 / (((v39 % ((v39) | 1))) | 1))), F1: v12}
                                        _ = v40
                                    case "\n\t":
                                        v41 := GenS7{F0: 0, F1: false}
                                        _ = v41
                                    default:
                                        v6 = v2
                                    }
                                } else {
                                    for fi42 := 0; fi42 < 2; fi42++ {
                                        if (!false) {
                                            v43 := GenS8{F0: v39}
                                            _ = v43
                                        } else {
                                            v12 = (v12 >> ((uint(v4)) & 63))
                                        }
                                        v6 = v2
                                        var v44 uint32 = 612
                                        if true {
                                            var v45 uint16 = 975
                                            var v46 float64 = 0.0
                                            if ((v16 | (-(fi42))) != 545) {
                                                v47 := GenS9{F0: 511, F1: v5}
                                                _ = v47
                                            }
                                            _ = v45
                                            _ = v46
                                        } else {
                                            var v48 int8 = 100
                                            switch ("hello" + "foo bar") {
                                            case "":
                                                v49 := GenS2{F0: v12}
                                                _ = v49
                                            case "foo bar":
                                                fmt.Println(v5)
                                            default:
                                                var v50 uint64 = 1
                                                _ = v50
                                            }
                                            v48 = (-(v48))
                                            _ = v48
                                        }
                                        var v51 uint8 = v4
                                        _ = fi42
                                        _ = v44
                                        _ = v51
                                    }
                                    var v52 uint8 = (v4 / ((230) | 1))
                                    if (!false) {
                                        for fi53 := 0; fi53 < 13; fi53++ {
                                            delete(v3, "a")
                                            _ = fi53
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    } else {
                                        if false {
                                            for fi54 := 0; fi54 < 20; fi54++ {
                                                if false {
                                                    v5 = (v5 & v5)
                                                }
                                                _ = len(v3)
                                                fmt.Println(v39, v39, v2)
                                                _ = fi54
                                            }
                                        } else {
                                            v55 := GenS10{F0: -493, F1: v12, F2: v5}
                                            if false {
                                                for fi56 := 0; fi56 < 9; fi56++ {
                                                    if len(v1) > 0 {
                                                        v1[0] = v16
                                                    }
                                                    var v57 uint16 = 2
                                                    _ = fi56
                                                    _ = v57
                                                }
                                                switch "a" {
                                                case "":
                                                    v58 := GenS6{F0: 42, F1: 42}
                                                    _ = v58
                                                case "a":
                                                    fmt.Println(v6)
                                                default:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                var v59 int = 716
                                                _ = v59
                                            }
                                            _ = v55
                                        }
                                        var v60 uint32 = (v12 / (((v12 % (((v12 ^ v12)) | 1))) | 1))
                                        delete(v3, "\n\t")
                                        _ = v60
                                    }
                                    _ = v52
                                }
                            }
                        }
                        _ = v39
                    } else {
                        if ((false && false) || (false || (((42 < 1) || true) || false))) {
                            if (42 >= (v5 + v5)) {
                                v2 -= v6
                                v12 = (v12 / ((uint32(v4)) | 1))
                            } else {
                                v61 := v3[("\n\t" + "a")]
                                _ = v61
                            }
                            for fi62 := 0; fi62 < 7; fi62++ {
                                v5 += v5
                                v63 := GenS8{F0: 60}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = fi62
                                _ = v63
                            }
                            v64 := GenS8{F0: -128}
                            _ = v64
                        } else {
                            fmt.Println(v16, fi11)
                            switch ("" + ("foo bar" + "")) {
                            case "\n\t":
                                if len(v1) > 0 {
                                    v1[0] = -2147483648
                                }
                            }
                            for fi65 := 0; fi65 < 10; fi65++ {
                                var v66 uint8 = v4
                                _ = fi65
                                _ = v66
                            }
                        }
                        v12 = 4294967295
                    }
                } else {
                    switch ("hello" + "\n\t") {
                    case "hello":
                        for fi67 := 0; fi67 < 8; fi67++ {
                            var v68 int32 = -1
                            switch "foo bar" {
                            case "hello":
                                var v69 float32 = (-2.25 * (1e10 * 3.14159))
                                switch "\n\t" {
                                case "foo bar":
                                    var v70 uint16 = 255
                                    _ = v70
                                case "":
                                    fmt.Println(v2, v69)
                                }
                                _ = v69
                            }
                            _ = fi67
                            _ = v68
                        }
                    case "":
                        var v71 int = (fi11 / (((-(fi11))) | 1))
                        _ = v71
                    case "a":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        var v72 int16 = 32767
                        _ = v72
                    }
                }
            }
            _ = v16
        }
        _ = fi11
        _ = v12
    }
    var v73 int16 = (-(v5))
    for fi74 := 0; fi74 < 17; fi74++ {
        for fi75 := 0; fi75 < 14; fi75++ {
            if (!true) {
                v6 = v6
                v76 := GenS10{F0: -1, F1: 1, F2: v5}
                _ = v76
            }
            var v77 uint8 = 188
            v6 -= (v2 * -9223372036854775808)
            _ = fi75
            _ = v77
        }
        fmt.Println(v73)
        _ = fi74
    }
    if (1 < -128) {
        switch "\n\t" {
        case "hello":
            if false {
                if (0 <= 623) {
                    var v78 uint64 = 0
                    for fi79 := 0; fi79 < 11; fi79++ {
                        if false {
                            fmt.Println(fi79, v2)
                        }
                        _ = fi79
                    }
                    _ = v78
                }
                fmt.Println(v5, v2, v6)
            }
        case "a":
            for fi80 := 0; fi80 < 14; fi80++ {
                var v81 int = 42
                if (2147483648 != 0) {
                    var v82 int32 = 42
                    _ = v82
                }
                if (-1 < 42) {
                    v83 := GenS2{F0: 42}
                    _ = v83
                } else {
                    v5 = v5
                }
                _ = fi80
                _ = v81
            }
        case "foo bar":
            v84 := GenS9{F0: 0, F1: (v73 >> ((4294967295) & 63))}
            _ = v84
        }
        var v85 uint8 = (v4 * (v4 / ((0) | 1)))
        var v86 int8 = -128
        _ = v85
        _ = v86
    }
    v2 = v2
    _ = len(v1)
    fmt.Println(v5, v4)
    for fi87 := 0; fi87 < 15; fi87++ {
        v88 := GenS5{F0: v6}
        _ = fi87
        _ = v88
    }
    switch "foo bar" {
    case "hello":
        v5 = v5
    default:
        var v89 uint8 = v4
        _ = v89
    }
    var v90 float64 = float64(1e10)
    v91 := v3[("\n\t" + "a")]
    for fi92 := 0; fi92 < 20; fi92++ {
        fmt.Println(v73)
        v90 = 0.1
        _ = fi92
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v8
    _ = v90
    _ = v91
}
