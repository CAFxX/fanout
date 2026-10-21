package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 bool; F2 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 bool; F2 uint }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 uint64; F1 int64; F2 int8 }

func main() {
    v1 := []uint32{65535, 0}
    var v2 int16 = -1
    var v3 float32 = 1e10
    switch "foo bar" {
    case "hello":
        if (1 <= 0) {
            if false {
                var v4 uint16 = 928
                v4 = 767
                _ = v4
            } else {
                var v5 int8 = -1
                var v6 int8 = -34
                v7 := GenS1{F0: 1, F1: 115}
                _ = v5
                _ = v6
                _ = v7
            }
            for fi8 := 0; fi8 < 11; fi8++ {
                var v9 int64 = -167
                v3 = (v3 + v3)
                if true {
                    switch -1 {
                    case 0:
                        v10 := GenS2{F0: 0}
                        _ = v10
                    case 338:
                        switch "foo bar" {
                        case "a":
                            for fi11 := 0; fi11 < 16; fi11++ {
                                if (("" + "a") == "hello") {
                                    var v12 int16 = (v2 / ((32767) | 1))
                                    if (42 < 1) {
                                        v2 += (v2 * v2)
                                        var v13 uint32 = 532
                                        _ = v13
                                    } else {
                                        fmt.Println(fi11, v2, v12)
                                    }
                                    fmt.Println(v9)
                                    var v14 uint8 = 42
                                    v15 := GenS2{F0: 687}
                                    _ = v12
                                    _ = v14
                                    _ = v15
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = len(v1)
                                }
                                _ = fi11
                            }
                        case "foo bar":
                            for fi16 := 0; fi16 < 6; fi16++ {
                                if true {
                                    var v17 int16 = v2
                                    v3 -= ((1.5 - ((0.0 + float32(1e10)) - 1e10)) - -2.25)
                                    _ = v17
                                }
                                _ = fi16
                            }
                        default:
                            var v18 string = "foo bar"
                            _ = v18
                        }
                    }
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if (v3 < -2.25) {
                        fmt.Println(v3, v9)
                        var v19 string = ""
                        v3 = (v3 * v3)
                        _ = v19
                    } else {
                        switch "foo bar" {
                        case "a":
                            if (!true) {
                                v3 = v3
                                var v20 float64 = 0.1
                                var v21 uint8 = 2
                                _ = v20
                                _ = v21
                            } else {
                                for fi22 := 0; fi22 < 4; fi22++ {
                                    fmt.Println(v9)
                                    _ = fi22
                                }
                                var v23 float64 = (0.0 + 1e10)
                                _ = v23
                            }
                        case "":
                            switch (fi8 | fi8) {
                            case -486:
                                fmt.Println(v2, v3)
                            case -857:
                                v9 -= v9
                            }
                        default:
                            for fi24 := 0; fi24 < 15; fi24++ {
                                v25 := GenS1{F0: 2, F1: 128}
                                var v26 float32 = (v3 - v3)
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                        }
                        v27 := GenS1{F0: 4294967295, F1: 0}
                        _ = v27
                    }
                }
                _ = fi8
                _ = v9
            }
        }
    case "a":
        for fi28 := 0; fi28 < 1; fi28++ {
            v3 += v3
            _ = fi28
        }
    default:
        if false {
            _ = len(v1)
            switch 127 {
            case 9223372036854775807:
                fmt.Println(v2)
            }
        } else {
            if (2 < 255) {
                v3 = (v3 + v3)
                v29 := GenS3{F0: 0, F1: true, F2: 0}
                _ = v29
            }
        }
    }
    v30 := GenS4{F0: 1}
    var v31 uint32 = 4294967295
    switch 0 {
    case 124:
        if false {
            v3 += (1e10 * ((float32((-2.25 / 1e10)) + 0.1) / v3))
        }
    case 189:
        for fi32 := 0; fi32 < 9; fi32++ {
            for fi33 := 0; fi33 < 4; fi33++ {
                switch ((("\n\t" + ("a" + "foo bar")) + "") + "hello") {
                case "foo bar":
                    v2 -= (v2 % ((-1) | 1))
                }
                _ = fi33
            }
            _ = fi32
        }
    case 9223372036854775807:
        var v34 int8 = 2
        _ = v34
    default:
        v31 = 0
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        for fi35 := 0; fi35 < 7; fi35++ {
            _ = len(v1)
            v3 -= 3.14159
            v2 += (v2 & v2)
            _ = fi35
        }
    case 42:
        fmt.Println(v2)
    default:
        v3 -= (float32((0.1 / 1e10)) / v3)
    }
    fmt.Println(v31, v3, v31)
    v3 = ((1e10 / v3) + v3)
    var v36 float32 = (1e10 + v3)
    for fi37 := 0; fi37 < 5; fi37++ {
        switch 42 {
        case 347:
            var v38 float32 = (float32(0.0) * 0.0)
            _ = v38
        case 44:
            v3 = -2.25
        case 938:
            var v39 float32 = v3
            _ = v39
        }
        _ = fi37
    }
    for fi40 := 0; fi40 < 20; fi40++ {
        if (false || (false && true)) {
            if true {
                v41 := GenS5{F0: 1, F1: (!true), F2: 0}
                fmt.Println(fi40, v2, fi40)
                _ = v41
            } else {
                v3 = v3
                _ = len(v1)
            }
        }
        _ = fi40
    }
    switch 127 {
    case 2147483647:
        if (int(v31) < 4294967295) {
            for fi42 := 0; fi42 < 3; fi42++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if len(v1) > 0 {
                    v1[0] = (v31 - v31)
                }
                v1 = append(v1, 0)
                _ = fi42
            }
        } else {
            v31 = 2147483647
            var v43 float64 = (float64(-2.25) - ((((float64(v36) * 0.1) / (1e10 - 0.1)) * (float64(0.0) / 1.5)) * 0.1))
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 2; fi44++ {
        if (!(4294967295 >= 1)) {
            var v45 float64 = 1e10
            if false {
                if (true && (9223372036854775807 == (fi44 * (fi44 >> ((2) & 63))))) {
                    v36 = -2.25
                }
                var v46 float64 = (v45 / ((0.1 / v45) * (v45 * v45)))
                _ = v46
            } else {
                v47 := GenS2{F0: 0}
                fmt.Println(fi44)
                v2 -= -1
                _ = v47
            }
            var v48 uint32 = (v31 - v31)
            _ = v45
            _ = v48
        }
        var v49 string = (("foo bar" + ("hello" + "a")) + "\n\t")
        _ = fi44
        _ = v49
    }
    v50 := GenS6{F0: -1, F1: 747}
    var v51 int32 = -2147483648
    v1 = append(v1, (^(v31)))
    for fi52 := 0; fi52 < 19; fi52++ {
        _ = len(v1)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch (fi52 ^ -779) {
        case -910:
            v53 := GenS2{F0: 4294967295}
            _ = v53
        case 285:
            var v54 bool = true
            _ = v54
        case 4294967295:
            switch "a" {
            case "\n\t":
                if (4294967295 < (^(v31))) {
                    v55 := GenS4{F0: v2}
                    if true {
                        var v56 uint32 = 479
                        var v57 int = 0
                        v3 -= -2.25
                        _ = v56
                        _ = v57
                    }
                    v2 -= (v2 - v2)
                    _ = v55
                }
            case "a":
                var v58 int16 = v2
                _ = v58
            case "":
                v2 -= (-(v2))
            }
        }
        _ = fi52
    }
    if true {
        if true {
            for fi59 := 0; fi59 < 16; fi59++ {
                v51 -= (v51 - (-(v51)))
                for fi60 := 0; fi60 < 7; fi60++ {
                    var v61 string = "foo bar"
                    v62 := GenS5{F0: 252, F1: (!true), F2: 4294967295}
                    _ = fi60
                    _ = v61
                    _ = v62
                }
                _ = fi59
            }
        } else {
            v63 := GenS1{F0: 18446744073709551615, F1: 1}
            _ = v63
        }
        if false {
            var v64 int16 = (v2 - (-(v2)))
            var v65 int64 = 65
            var v66 uint = 4294967295
            _ = v64
            _ = v65
            _ = v66
        } else {
            for fi67 := 0; fi67 < 1; fi67++ {
                for fi68 := 0; fi68 < 7; fi68++ {
                    v69 := GenS1{F0: 18446744073709551615, F1: 127}
                    _ = fi68
                    _ = v69
                }
                v51 = (v51 & 0)
                if false {
                    v3 -= v36
                    if len(v1) > 0 {
                        v1[0] = v31
                    }
                }
                _ = fi67
            }
            if (!(("a" > "hello") || true)) {
                v70 := GenS7{F0: 42, F1: 0, F2: 0}
                _ = v70
            }
        }
        if (v2 == int16(v51)) {
            for fi71 := 0; fi71 < 17; fi71++ {
                var v72 bool = (!(fi71 == fi71))
                _ = fi71
                _ = v72
            }
            if (!true) {
                switch int(v31) {
                case 2147483647:
                    var v73 uint16 = 65535
                    _ = v73
                case -2147483648:
                    for fi74 := 0; fi74 < 12; fi74++ {
                        v75 := GenS3{F0: (^(v51)), F1: (v31 < 2), F2: -1}
                        v76 := GenS1{F0: 18446744073709551615, F1: uint8(v31)}
                        _ = fi74
                        _ = v75
                        _ = v76
                    }
                }
                v51 += (v51 & (-(v51)))
            } else {
                switch (("\n\t" + ((((("" + "foo bar") + "") + "hello") + "a") + "hello")) + ("a" + "hello")) {
                case "foo bar":
                    for fi77 := 0; fi77 < 6; fi77++ {
                        _ = len(v1)
                        _ = fi77
                    }
                case "hello":
                    v78 := GenS6{F0: 9223372036854775807, F1: 923}
                    _ = v78
                case "":
                    switch "hello" {
                    case "hello":
                        if false {
                            v31 -= (v31 % ((v31) | 1))
                            var v79 uint64 = 42
                            if len(v1) > 0 {
                                v1[0] = (v31 | (v31 >> ((691) & 63)))
                            }
                            _ = v79
                        }
                    default:
                        fmt.Println(v36, v3)
                    }
                }
                if len(v1) > 0 {
                    v1[0] = 2
                }
                v3 += ((1e10 * v3) / 3.14159)
            }
            if ((("a" < "a") || (float32((((0.0 + ((-2.25 / 1e10) - float64(v3))) / 0.1) - -2.25)) != float32(0.1))) && true) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                switch 128 {
                case 169:
                    v80 := GenS1{F0: 18446744073709551615, F1: 1}
                    _ = v80
                case 619:
                    for fi81 := 0; fi81 < 3; fi81++ {
                        var v82 int8 = 127
                        _ = fi81
                        _ = v82
                    }
                default:
                    v51 += v51
                }
            } else {
                var v83 uint = 2
                var v84 uint8 = 0
                v36 = v36
                _ = v83
                _ = v84
            }
        }
    } else {
        var v85 uint64 = 910
        var v86 uint = 2
        v85 = 310
        _ = v85
        _ = v86
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v50
    _ = v51
}
