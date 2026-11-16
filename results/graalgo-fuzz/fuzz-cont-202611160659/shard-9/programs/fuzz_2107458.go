package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int64; F1 int16 }
type GenS10 struct { F0 int8; F1 uint }
type GenS11 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS12 struct { F0 uint32; F1 uint32 }
type GenS13 struct { F0 uint8; F1 int64; F2 int }
type GenS14 struct { F0 uint64; F1 bool; F2 int }
type GenS15 struct { F0 uint8; F1 int32; F2 int32 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 int64; F1 uint; F2 uint8 }
type GenS18 struct { F0 bool; F1 uint }
type GenS19 struct { F0 int; F1 uint; F2 uint32 }
type GenS20 struct { F0 uint64 }
type GenS21 struct { F0 int64; F1 int }
type GenS22 struct { F0 string; F1 int8 }
type GenS23 struct { F0 int; F1 int }
type GenS24 struct { F0 uint8 }
type GenS25 struct { F0 uint64; F1 bool; F2 uint32 }
func genFn0(p0 int64, p1 uint16, p2 bool) uint32 {
    var v7 int32 = 42
    v7 = v7
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2147483648
}
func genFn1() int {
    v8 := []int{}
    var v9 float64 = -2.25
    var v10 uint64 = 0
    if (false && (42 < 42)) {
        var v11 string = "\n\t"
        v12 := GenS1{F0: -128, F1: 42}
        _ = v11
        _ = v12
    } else {
        var v13 float32 = -2.25
        if (!(!false)) {
            var v14 uint64 = 42
            if len(v8) > 0 {
                v8[0] = -9223372036854775808
            }
            _ = v14
        } else {
            for fi15 := 0; fi15 < 17; fi15++ {
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 float32 = (0.1 * (float32(v9) + float32(float64(1e10))))
                    if len(v8) > 0 {
                        v8[0] = (fi16 * fi16)
                    }
                    if ((9223372036854775807 != -9223372036854775808) && (-32768 < 726)) {
                        fmt.Println(v17)
                        switch "hello" {
                        case "\n\t":
                            v10 = (v10 ^ v10)
                        case "":
                            var v18 int32 = 66
                            if len(v8) > 0 {
                                v8[0] = (fi16 & -9223372036854775808)
                            }
                            v10 = (v10 + v10)
                            _ = v18
                        case "foo bar":
                            v9 = (v9 * (v9 / (v9 + float64(v13))))
                        }
                        v19 := GenS2{F0: 128}
                        _ = v19
                    } else {
                        var v20 int8 = -128
                        var v21 string = "a"
                        fmt.Println(v10)
                        _ = v20
                        _ = v21
                    }
                    _ = fi16
                    _ = v17
                }
                for fi22 := 0; fi22 < 7; fi22++ {
                    var v23 uint16 = 2
                    if false {
                        v24 := GenS2{F0: 0}
                        v9 = float64(1e10)
                        v23 -= 255
                        _ = v24
                    } else {
                        v13 += (v13 / v13)
                        var v25 int64 = 0
                        _ = v25
                    }
                    _ = fi22
                    _ = v23
                }
                if (2 == 65535) {
                    v10 += (^(v10))
                    for fi26 := 0; fi26 < 14; fi26++ {
                        var v27 uint16 = 0
                        v13 += ((v13 + (((v13 - float32((float64(v13) / (v9 / 3.14159)))) * 3.14159) + float32(v9))) - v13)
                        switch "a" {
                        case "foo bar":
                            v28 := GenS1{F0: 100, F1: -32768}
                            _ = v28
                        case "hello":
                            switch "a" {
                            case "":
                                v10 = 0
                            case "hello":
                                fmt.Println(v27)
                            case "\n\t":
                                fmt.Println(v27, v10)
                            default:
                                v29 := GenS1{F0: 127, F1: 1}
                                _ = v29
                            }
                        case "\n\t":
                            var v30 int32 = 2147483647
                            _ = v30
                        default:
                            var v31 uint64 = (v10 * (v10 % (((^(v10))) | 1)))
                            _ = v31
                        }
                        _ = fi26
                        _ = v27
                    }
                    v9 -= v9
                } else {
                    var v32 bool = true
                    switch ("a" + "\n\t") {
                    case "foo bar":
                        switch "" {
                        case "":
                            if v32 {
                                v33 := GenS3{F0: ((1e10 < 3.14159) || (!(233 < 83)))}
                                if ((!(!true)) || v32) {
                                    if v32 {
                                        for fi34 := 0; fi34 < 2; fi34++ {
                                            switch "foo bar" {
                                            case "a":
                                                v10 = 4294967295
                                                v8 = append(v8, fi15)
                                                if ((v10 & (v10 << ((431) & 63))) == v10) {
                                                    if v32 {
                                                        switch (fi15 + fi15) {
                                                        case 9223372036854775807:
                                                            v10 -= v10
                                                        case -94:
                                                            if v32 {
                                                                v32 = v32
                                                            }
                                                            v10 = (v10 / (((^(v10))) | 1))
                                                        }
                                                        if (790 == 65535) {
                                                            var v35 int32 = 2147483647
                                                            switch (-(fi15)) {
                                                            case 127:
                                                                v35 += 0
                                                            }
                                                            _ = v35
                                                        }
                                                    } else {
                                                        var v36 float64 = 0.1
                                                        v36 = v36
                                                        fmt.Println(fi15, v36, v9)
                                                        var v37 uint16 = 407
                                                        v32 = (874 >= 474)
                                                        if len(v8) > 0 {
                                                            v8[0] = fi34
                                                        }
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                    var v38 string = "foo bar"
                                                    var v39 string = ("a" + v38)
                                                    if (32767 != 1000) {
                                                        if (191 <= 1) {
                                                            switch (-(fi34)) {
                                                            case 32767:
                                                                v39 = v39
                                                            default:
                                                                v13 -= v13
                                                            }
                                                            if v32 {
                                                                if ((v10 > v10) && v32) {
                                                                    fmt.Println(fi34, v10)
                                                                    fmt.Println(fi15, v38)
                                                                }
                                                                fmt.Println(v38, v32)
                                                                var v40 uint16 = 2
                                                                _ = v40
                                                            }
                                                            fmt.Println(v32, v9)
                                                        }
                                                        if len(v8) > 0 {
                                                            v8[0] = fi34
                                                        }
                                                        if (!((v9 / 1.0) < (v9 - v9))) {
                                                            v9 = v9
                                                        }
                                                    }
                                                    var v41 int64 = int64(v10)
                                                    _ = v38
                                                    _ = v39
                                                    _ = v41
                                                } else {
                                                    var v42 float64 = v9
                                                    switch -486 {
                                                    case -9223372036854775808:
                                                        fmt.Println(fi15)
                                                    case 128:
                                                        if len(v8) > 0 {
                                                            v8[0] = (fi15 >> ((277) & 63))
                                                        }
                                                    }
                                                    _ = v42
                                                }
                                                switch "a" {
                                                case "hello":
                                                    switch (fi34 | (fi15 ^ (fi15 * (fi15 / ((42) | 1))))) {
                                                    case 770:
                                                        fmt.Println(v10, fi34, fi34)
                                                    case 4294967295:
                                                        fmt.Println(v32)
                                                    case 127:
                                                        v9 -= (v9 + v9)
                                                    default:
                                                        v43 := GenS3{F0: ((^(fi15)) == -2147483648)}
                                                        _ = v43
                                                    }
                                                }
                                            }
                                            if v32 {
                                                if (42 < -127) {
                                                    v13 = 0.0
                                                    if (0 == 2) {
                                                        if v32 {
                                                            fmt.Println(v9, v10)
                                                            if true {
                                                                if (v32 || v32) {
                                                                    if len(v8) >= 2 {
                                                                        v8 = v8[1:len(v8)]
                                                                    }
                                                                    v13 -= 1.5
                                                                }
                                                            } else {
                                                                fmt.Println(v13, v10, fi15)
                                                                v44 := GenS4{F0: 2}
                                                                _ = v44
                                                            }
                                                            var v45 float64 = float64(v13)
                                                            _ = v45
                                                        } else {
                                                            if ((fi34 << ((255) & 63)) >= fi34) {
                                                                v46 := GenS1{F0: 100, F1: 387}
                                                                v10 = (^(v10))
                                                                switch fi15 {
                                                                case -9223372036854775808:
                                                                    _ = len(v8)
                                                                case 160:
                                                                    var v47 uint16 = 0
                                                                    _ = v47
                                                                case 1:
                                                                    if len(v8) >= 2 {
                                                                        v8 = v8[1:len(v8)]
                                                                    }
                                                                }
                                                                _ = v46
                                                            } else {
                                                                if false {
                                                                    _ = len(v8)
                                                                }
                                                                var v48 int = fi34
                                                                _ = v48
                                                            }
                                                            if (!v32) {
                                                                v49 := GenS5{F0: genFn0(9223372036854775807, 42, false)}
                                                                _ = v49
                                                            }
                                                            _ = len(v8)
                                                        }
                                                        var v50 uint16 = 32767
                                                        var v51 int64 = 22
                                                        var v52 int8 = 1
                                                        _ = v50
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                    if false {
                                                        switch (fi15 % (((fi34 % ((4294967295) | 1))) | 1)) {
                                                        case 667:
                                                            if v32 {
                                                                var v53 int = fi34
                                                                fmt.Println(fi15, v9)
                                                                _ = v53
                                                            } else {
                                                                var v54 uint32 = 4294967295
                                                                switch fi34 {
                                                                case -404:
                                                                    v55 := GenS1{F0: -1, F1: 42}
                                                                    _ = v55
                                                                case 9223372036854775807:
                                                                    switch (fi15 | (fi34 + (fi15 ^ 223))) {
                                                                    case 456:
                                                                        var v56 string = ""
                                                                        _ = v56
                                                                    case 128:
                                                                        v8 = append(v8, fi34)
                                                                    case 1:
                                                                        v57 := GenS6{F0: -84}
                                                                        _ = v57
                                                                    default:
                                                                        v10 += (^(v10))
                                                                    }
                                                                default:
                                                                    _ = len(v8)
                                                                }
                                                                _ = v54
                                                            }
                                                        case 2147483647:
                                                            var v58 uint64 = 1
                                                            _ = v58
                                                        }
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            if v32 {
                                                                fmt.Println(fi34)
                                                            }
                                                        case "":
                                                            v13 = v13
                                                        default:
                                                            fmt.Println(v32)
                                                        }
                                                    } else {
                                                        if false {
                                                            v10 -= v10
                                                            v13 += v13
                                                            var v59 uint32 = 65535
                                                            _ = v59
                                                        } else {
                                                            if true {
                                                                v60 := GenS4{F0: 65535}
                                                                v61 := GenS7{F0: (v32 && (!(0 <= 2))), F1: 1}
                                                                var v62 int8 = 42
                                                                _ = v60
                                                                _ = v61
                                                                _ = v62
                                                            } else {
                                                                v10 += v10
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    var v63 float64 = (v9 / (((v9 * v9) - v9) - float64(v13)))
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                    v64 := GenS8{F0: "\n\t"}
                                                    _ = v63
                                                    _ = v64
                                                }
                                                var v65 int8 = 73
                                                if (!v32) {
                                                    _ = len(v8)
                                                } else {
                                                    v9 += v9
                                                    v66 := GenS9{F0: -279, F1: 32767}
                                                    _ = v66
                                                }
                                                _ = v65
                                            } else {
                                                if (false && v32) {
                                                    if v32 {
                                                        fmt.Println(v9, fi34, fi15)
                                                        var v67 bool = v32
                                                        _ = v67
                                                    } else {
                                                        fmt.Println(fi34)
                                                        if ((0 < 128) && (!true)) {
                                                            v9 += 0.0
                                                            if len(v8) > 0 {
                                                                v8[0] = fi34
                                                            }
                                                            if (2 > 0) {
                                                                var v68 uint16 = 65535
                                                                _ = v68
                                                            } else {
                                                                if v32 {
                                                                    switch ("" + "hello") {
                                                                    case "foo bar":
                                                                        switch ("\n\t" + "") {
                                                                        case "a":
                                                                            fmt.Println(v9)
                                                                        case "":
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                if (v32 || true) {
                                                                                    if len(v8) >= 2 {
                                                                                        v8 = v8[1:len(v8)]
                                                                                    }
                                                                                } else {
                                                                                    if len(v8) > 0 {
                                                                                        v8[0] = fi34
                                                                                    }
                                                                                    if false {
                                                                                        var v69 uint16 = 32767
                                                                                        v32 = false
                                                                                        _ = v69
                                                                                    }
                                                                                    v9 -= v9
                                                                                    fmt.Println(v9, v32, v10)
                                                                                    v70 := GenS10{F0: 70, F1: 4294967295}
                                                                                    _ = v70
                                                                                }
                                                                            case "":
                                                                                v32 = ((!false) || v32)
                                                                            default:
                                                                                var v71 uint8 = 128
                                                                                v71 = v71
                                                                                if false {
                                                                                    switch ("a" + "\n\t") {
                                                                                    case "\n\t":
                                                                                        if (0 < 773) {
                                                                                            v9 = v9
                                                                                            var v72 int8 = 1
                                                                                            _ = v72
                                                                                        }
                                                                                    case "":
                                                                                        if (true && (!(v32 && v32))) {
                                                                                            var v73 uint16 = 32767
                                                                                            v74 := GenS11{F0: (v71 & (v71 / (((^(v71))) | 1))), F1: v10, F2: (^(v73))}
                                                                                            fmt.Println(fi15)
                                                                                            _ = v73
                                                                                            _ = v74
                                                                                        }
                                                                                    default:
                                                                                        var v75 int64 = 257
                                                                                        _ = v75
                                                                                    }
                                                                                    switch ("\n\t" + "foo bar") {
                                                                                    case "\n\t":
                                                                                        v76 := GenS2{F0: 2}
                                                                                        _ = v76
                                                                                    }
                                                                                    var v77 int16 = -32768
                                                                                    _ = v77
                                                                                }
                                                                                _ = v71
                                                                            }
                                                                        }
                                                                    }
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        v32 = false
                                                                        var v78 int64 = 204
                                                                        _ = v78
                                                                    case "":
                                                                        if (true || v32) {
                                                                            var v79 float64 = 1e10
                                                                            fmt.Println(v32)
                                                                            if v32 {
                                                                                if len(v8) >= 2 {
                                                                                    v8 = v8[1:len(v8)]
                                                                                }
                                                                                if v32 {
                                                                                    switch "\n\t" {
                                                                                    case "":
                                                                                        fmt.Println(v79, fi34, fi15)
                                                                                    case "\n\t":
                                                                                        v32 = (!v32)
                                                                                    case "foo bar":
                                                                                        v32 = true
                                                                                    }
                                                                                    fmt.Println(fi15)
                                                                                    v10 = 617
                                                                                }
                                                                                var v80 int16 = -1
                                                                                _ = v80
                                                                            } else {
                                                                                v81 := GenS12{F0: 1, F1: 932}
                                                                                var v82 int64 = 790
                                                                                _ = v81
                                                                                _ = v82
                                                                            }
                                                                            _ = v79
                                                                        } else {
                                                                            fmt.Println(fi15)
                                                                        }
                                                                    default:
                                                                        if len(v8) > 0 {
                                                                            v8[0] = (fi15 / (((fi34 / ((-1) | 1))) | 1))
                                                                        }
                                                                    }
                                                                    if len(v8) >= 2 {
                                                                        v8 = v8[1:len(v8)]
                                                                    }
                                                                    v32 = false
                                                                    if len(v8) > 0 {
                                                                        v8[0] = 2147483647
                                                                    }
                                                                    var v83 string = (("hello" + (("hello" + (("" + "\n\t") + ("" + "hello"))) + "")) + "")
                                                                    var v84 float32 = (0.0 + v13)
                                                                    v85 := GenS8{F0: ""}
                                                                    _ = v83
                                                                    _ = v84
                                                                    _ = v85
                                                                } else {
                                                                    if len(v8) >= 2 {
                                                                        v8 = v8[1:len(v8)]
                                                                    }
                                                                    if v32 {
                                                                        v8 = append(v8, -79)
                                                                        v13 -= (v13 + v13)
                                                                        var v86 uint = 255
                                                                        _ = v86
                                                                    }
                                                                }
                                                                if true {
                                                                    v13 -= v13
                                                                    var v87 uint8 = 252
                                                                    if false {
                                                                        v10 -= (v10 / (((^(v10))) | 1))
                                                                        v9 = 3.14159
                                                                        v88 := GenS11{F0: (^(v87)), F1: (v10 & v10), F2: 1}
                                                                        _ = v88
                                                                    }
                                                                    _ = v87
                                                                } else {
                                                                    if len(v8) >= 2 {
                                                                        v8 = v8[1:len(v8)]
                                                                    }
                                                                    if (v13 == 3.14159) {
                                                                        v13 += 3.14159
                                                                        v32 = false
                                                                    }
                                                                }
                                                                v9 = 3.14159
                                                            }
                                                        } else {
                                                            var v89 bool = v32
                                                            _ = v89
                                                        }
                                                    }
                                                }
                                                if (127 != (fi34 % (((fi15 << ((6) & 63))) | 1))) {
                                                    if (!(!true)) {
                                                        if false {
                                                            v9 -= 3.14159
                                                            v9 = v9
                                                        } else {
                                                            var v90 uint8 = 42
                                                            var v91 int = 2147483647
                                                            if v32 {
                                                                if false {
                                                                    fmt.Println(v90, v32)
                                                                    v8 = append(v8, v91)
                                                                    if false {
                                                                        fmt.Println(v90, fi34, v32)
                                                                        v9 = v9
                                                                        v92 := GenS2{F0: (v90 * uint8(v10))}
                                                                        _ = v92
                                                                    }
                                                                }
                                                                var v93 int8 = 100
                                                                var v94 string = "a"
                                                                _ = v93
                                                                _ = v94
                                                            }
                                                            v91 -= v91
                                                            _ = v90
                                                            _ = v91
                                                        }
                                                        v95 := GenS13{F0: 0, F1: -847, F2: (^(fi15))}
                                                        v32 = v32
                                                        _ = v95
                                                    }
                                                } else {
                                                    var v96 uint8 = 127
                                                    v97 := GenS14{F0: uint64(v96), F1: false, F2: 991}
                                                    switch 495 {
                                                    case -5:
                                                        v32 = v32
                                                    case 202:
                                                        v96 -= 128
                                                    case 70:
                                                        if len(v8) > 0 {
                                                            v8[0] = -2147483648
                                                        }
                                                    }
                                                    _ = v96
                                                    _ = v97
                                                }
                                            }
                                            _ = fi34
                                        }
                                        for fi98 := 0; fi98 < 17; fi98++ {
                                            var v99 float64 = ((((v9 / ((-2.25 / (v9 / (float64((-2.25 * float32(v9))) / 1e10))) - (v9 * (((0.1 * 0.1) * 3.14159) - float64(0.0))))) / float64(((v13 + 0.0) - float32(((v9 + 3.14159) + 1.5))))) + v9) + 0.1)
                                            v8 = append(v8, fi98)
                                            _ = fi98
                                            _ = v99
                                        }
                                        var v100 uint = 4294967295
                                        _ = v100
                                    } else {
                                        switch "a" {
                                        case "\n\t":
                                            v101 := GenS15{F0: 0, F1: -1, F2: 1}
                                            _ = v101
                                        }
                                        var v102 float32 = v13
                                        _ = v102
                                    }
                                    for fi103 := 0; fi103 < 17; fi103++ {
                                        if v32 {
                                            switch 42 {
                                            case 32767:
                                                v13 = 0.0
                                            case -9223372036854775808:
                                                v9 += v9
                                            case 4294967295:
                                                var v104 uint = 255
                                                _ = v104
                                            }
                                            var v105 uint16 = 552
                                            if len(v8) > 0 {
                                                v8[0] = (^(fi15))
                                            }
                                            _ = v105
                                        }
                                        v9 -= v9
                                        var v106 uint32 = 1
                                        _ = fi103
                                        _ = v106
                                    }
                                    var v107 string = "\n\t"
                                    _ = v107
                                } else {
                                    if len(v8) > 0 {
                                        v8[0] = fi15
                                    }
                                    var v108 int64 = 0
                                    if (genFn0(v108, 32767, (-151 == v108)) != 695) {
                                        fmt.Println(fi15)
                                    } else {
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v109 uint8 = 226
                                            _ = v109
                                        default:
                                            var v110 uint32 = 42
                                            _ = v110
                                        }
                                    }
                                    _ = v108
                                }
                                var v111 uint64 = 18446744073709551615
                                _ = v111
                                _ = v33
                            } else {
                                var v112 int = fi15
                                var v113 uint8 = 2
                                _ = v112
                                _ = v113
                            }
                        case "a":
                            if (v32 && (((!(4294967295 > 2)) || true) && (uint(v10) <= 1))) {
                                if (true && true) {
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    v114 := GenS16{F0: -37}
                                    _ = v114
                                } else {
                                    fmt.Println(v9, v10, v13)
                                }
                                for fi115 := 0; fi115 < 11; fi115++ {
                                    if (-479 >= 603) {
                                        var v116 int64 = -9223372036854775808
                                        var v117 int = -9223372036854775808
                                        switch ("a" + "hello") {
                                        case "hello":
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                            v9 = v9
                                        default:
                                            if (2 != 2) {
                                                v32 = (18446744073709551615 >= 42)
                                            }
                                        }
                                        _ = v116
                                        _ = v117
                                    } else {
                                        var v118 int64 = 0
                                        if v32 {
                                            if v32 {
                                                if (!v32) {
                                                    v119 := GenS17{F0: -598, F1: 0, F2: 0}
                                                    _ = v119
                                                } else {
                                                    v120 := GenS18{F0: (642 >= (fi115 ^ (^(fi15)))), F1: 18446744073709551615}
                                                    var v121 uint32 = 759
                                                    _ = v120
                                                    _ = v121
                                                }
                                            } else {
                                                v32 = false
                                                v32 = (2147483648 > 2147483648)
                                            }
                                        } else {
                                            fmt.Println(fi15)
                                            v13 = (v13 + v13)
                                            var v122 string = ""
                                            _ = v122
                                        }
                                        _ = v118
                                    }
                                    v13 += -2.25
                                    _ = fi115
                                }
                                var v123 int16 = 214
                                _ = v123
                            } else {
                                switch "\n\t" {
                                case "":
                                    var v124 int = 32768
                                    _ = v124
                                default:
                                    for fi125 := 0; fi125 < 16; fi125++ {
                                        switch (-(fi15)) {
                                        case -571:
                                            var v126 int = 32767
                                            _ = v126
                                        case 1:
                                            var v127 int64 = 0
                                            _ = v127
                                        case 48:
                                            var v128 int8 = -27
                                            _ = v128
                                        }
                                        v10 = (v10 / ((v10) | 1))
                                        _ = fi125
                                    }
                                }
                                fmt.Println(v32, fi15, v32)
                            }
                        }
                    case "\n\t":
                        for fi129 := 0; fi129 < 3; fi129++ {
                            if v32 {
                                fmt.Println(fi15)
                                var v130 uint32 = genFn0(-1, 32767, (v13 != ((float32(v9) - v13) / v13)))
                                var v131 uint = 18446744073709551615
                                var v132 float64 = float64(float32(0.0))
                                v10 += v10
                                _ = v130
                                _ = v131
                                _ = v132
                            } else {
                                v10 += v10
                                v13 = (3.14159 - float32(v9))
                                if (-719 >= 0) {
                                    v133 := GenS9{F0: 0, F1: -553}
                                    v32 = false
                                    switch "\n\t" {
                                    case "hello":
                                        v8 = append(v8, (fi129 / ((fi129) | 1)))
                                    case "foo bar":
                                        v134 := GenS15{F0: 1, F1: -1, F2: 1}
                                        _ = v134
                                    case "\n\t":
                                        v8 = append(v8, (fi15 << ((42) & 63)))
                                    }
                                    _ = v133
                                }
                                if len(v8) > 0 {
                                    v8[0] = fi129
                                }
                            }
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            if len(v8) > 0 {
                                v8[0] = fi129
                            }
                            _ = fi129
                        }
                    case "":
                        var v135 int64 = 1
                        _ = v135
                    default:
                        v13 -= (v13 - float32(((0.1 + (1e10 / ((3.14159 - 0.0) * (float64((((v13 + (float32(-2.25) + v13)) * v13) / v13)) + (v9 / (float64(float32(1.5)) - v9)))))) - ((((v9 / float64(float32(v9))) * (v9 / ((1.5 + (v9 / v9)) - (v9 / float64(v13))))) * v9) / 1e10))))
                    }
                    _ = v32
                }
                _ = fi15
            }
            fmt.Println(v13)
        }
        fmt.Println(v9)
        _ = v13
    }
    _ = v10
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    v1 := []string{"hello"}
    var v2 uint8 = 135
    var v3 string = ""
    var v4 int32 = 831
    var v5 uint32 = 65535
    var v6 uint32 = (v5 * v5)
    if (true || (true && (v3 != "\n\t"))) {
        v136 := GenS11{F0: v2, F1: 2, F2: 255}
        for fi137 := 0; fi137 < 13; fi137++ {
            for fi138 := 0; fi138 < 16; fi138++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi138
            }
            switch v3 {
            case "\n\t":
                v2 += (v2 ^ v2)
            case "":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            fmt.Println(v6, v2, v3)
            _ = fi137
        }
        v6 = 2147483648
        _ = v136
    } else {
        for fi139 := 0; fi139 < 8; fi139++ {
            v140 := GenS19{F0: fi139, F1: 0, F2: 771}
            _ = fi139
            _ = v140
        }
    }
    var v141 int64 = 149
    fmt.Println(v4, v141)
    var v142 int16 = 32767
    if true {
        var v143 string = v3
        if false {
            _ = len(v1)
        } else {
            v3 = "a"
        }
        _ = v143
    } else {
        v4 = v4
    }
    if (1.5 <= 0.1) {
        if false {
            var v144 float32 = 1e10
            fmt.Println(v5)
            switch genFn1() {
            case 924:
                for fi145 := 0; fi145 < 9; fi145++ {
                    v142 += v142
                    var v146 float64 = float64(float32((1e10 + -2.25)))
                    _ = fi145
                    _ = v146
                }
            case 781:
                fmt.Println(v2, v144, v6)
            default:
                v4 = (v4 % ((v4) | 1))
            }
            _ = v144
        } else {
            if false {
                if false {
                    v2 = v2
                    switch v3 {
                    case "foo bar":
                        v147 := GenS20{F0: 18446744073709551615}
                        _ = v147
                    case "a":
                        var v148 string = v3
                        _ = v148
                    }
                } else {
                    v142 = 1000
                }
            }
        }
    }
    if (true || ("\n\t" >= "hello")) {
        switch v3 {
        case "a":
            v141 -= v141
        case "\n\t":
            v5 = 65535
        default:
            switch genFn1() {
            case -1:
                for fi149 := 0; fi149 < 5; fi149++ {
                    if false {
                        var v150 uint8 = (v2 << ((65535) & 63))
                        if (!(uint64(v150) != 18446744073709551615)) {
                            v6 = (v5 * (v6 + genFn0((^(v141)), 42, (42 == 2))))
                            v5 = (v5 % ((2147483647) | 1))
                        }
                        switch v3 {
                        case "":
                            if len(v1) > 0 {
                                v1[0] = v3
                            }
                        }
                        _ = v150
                    } else {
                        var v151 uint8 = v2
                        _ = v151
                    }
                    v152 := GenS18{F0: ((false && true) && (!true)), F1: 486}
                    _ = fi149
                    _ = v152
                }
            case 4294967295:
                v2 += 241
            default:
                for fi153 := 0; fi153 < 16; fi153++ {
                    v154 := GenS16{F0: 127}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi153
                    _ = v154
                }
            }
        }
        v141 -= v141
        fmt.Println(v5, v5, v2)
    } else {
        var v155 int16 = (v142 / ((148) | 1))
        v156 := GenS21{F0: (v141 % (((-(v141))) | 1)), F1: genFn1()}
        _ = v155
        _ = v156
    }
    v142 += v142
    for fi157 := 0; fi157 < 7; fi157++ {
        if (((v6 + (^(v5))) <= (^(v6))) && false) {
            v141 -= (v141 & 0)
            v158 := GenS10{F0: 127, F1: 598}
            switch "\n\t" {
            case "hello":
                if (515 < 4294967295) {
                    fmt.Println(v141, fi157, v141)
                    v141 -= v141
                }
            case "foo bar":
                v4 += v4
            case "a":
                if (!false) {
                    v6 = (v6 % (((v5 << ((uint(v2)) & 63))) | 1))
                }
            default:
                v5 = (v5 & (v5 + v5))
            }
            _ = v158
        } else {
            switch (fi157 * fi157) {
            case 75:
                var v159 int8 = -128
                _ = v159
            case 735:
                v5 += v5
            default:
                switch 32767 {
                case -533:
                    fmt.Println(v6)
                case 380:
                    var v160 int = (fi157 << ((2) & 63))
                    _ = v160
                default:
                    var v161 uint64 = 4294967295
                    _ = v161
                }
            }
        }
        v162 := GenS20{F0: 2}
        _ = fi157
        _ = v162
    }
    for fi163 := 0; fi163 < 11; fi163++ {
        for fi164 := 0; fi164 < 18; fi164++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch v3 {
            case "hello":
                var v165 float32 = 0.1
                _ = v165
            case "":
                fmt.Println(v5)
            case "\n\t":
                fmt.Println(v3, v4)
            default:
                v4 = v4
            }
            _ = fi164
        }
        var v166 uint16 = 0
        _ = fi163
        _ = v166
    }
    for fi167 := 0; fi167 < 20; fi167++ {
        v4 = -2147483648
        switch fi167 {
        case -986:
            for fi168 := 0; fi168 < 17; fi168++ {
                v142 -= (v142 << ((42) & 63))
                v4 = 2147483647
                _ = fi168
            }
        }
        _ = fi167
    }
    switch genFn1() {
    case -953:
        v2 -= (v2 % (((v2 >> ((65535) & 63))) | 1))
    case -2147483648:
        if true {
            v3 = v3
        } else {
            for fi169 := 0; fi169 < 20; fi169++ {
                if (v141 <= 1) {
                    if false {
                        if true {
                            for fi170 := 0; fi170 < 5; fi170++ {
                                switch -9223372036854775808 {
                                case 0:
                                    switch v3 {
                                    case "a":
                                        var v171 bool = ((v2 << ((612) & 63)) != v2)
                                        _ = v171
                                    case "foo bar":
                                        v6 = 42
                                    default:
                                        v172 := GenS22{F0: v3, F1: -27}
                                        _ = v172
                                    }
                                    v173 := GenS23{F0: genFn1(), F1: (fi170 - 42)}
                                    _ = v173
                                case 42:
                                    switch (fi169 + (fi169 - (fi170 ^ (fi170 - (fi169 ^ (-(fi170))))))) {
                                    case 42:
                                        if (((v142 & int16(v2)) == v142) || (100 != 112)) {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v6 += (v5 << ((4294967295) & 63))
                                        } else {
                                            if (39 != 788) {
                                                fmt.Println(v2)
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = ""
                                            }
                                            switch (v3 + v3) {
                                            case "a":
                                                var v174 float64 = 1e10
                                                _ = v174
                                            default:
                                                v5 = v6
                                            }
                                        }
                                    }
                                case -2147483648:
                                    v3 = v3
                                }
                                if (false || true) {
                                    var v175 int8 = 34
                                    if len(v1) > 0 {
                                        v1[0] = v3
                                    }
                                    v176 := GenS7{F0: (true && ((18446744073709551615 != 148) || ((v6 / (((v5 - v5)) | 1)) <= 65535))), F1: v4}
                                    _ = v175
                                    _ = v176
                                }
                                if true {
                                    var v177 int32 = v4
                                    _ = v177
                                } else {
                                    if true {
                                        v142 -= (v142 % (((^(v142))) | 1))
                                    }
                                    if (!true) {
                                        var v178 int = fi169
                                        _ = v178
                                    }
                                    v4 = v4
                                }
                                _ = fi170
                            }
                            for fi179 := 0; fi179 < 20; fi179++ {
                                if true {
                                    var v180 uint16 = 42
                                    v4 = v4
                                    var v181 bool = false
                                    _ = v180
                                    _ = v181
                                }
                                var v182 int32 = v4
                                _ = fi179
                                _ = v182
                            }
                            v6 += 4294967295
                        } else {
                            fmt.Println(v141, v3)
                            var v183 uint32 = v5
                            _ = v183
                        }
                        v5 -= 65535
                    } else {
                        v184 := GenS24{F0: 0}
                        fmt.Println(v4)
                        _ = v184
                    }
                    v142 = v142
                    var v185 uint32 = (v5 / (((v6 * v5)) | 1))
                    _ = v185
                } else {
                    _ = len(v1)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v142 = (^(v142))
                }
                v186 := GenS25{F0: 2, F1: false, F2: v5}
                _ = fi169
                _ = v186
            }
            var v187 string = "hello"
            var v188 int = genFn1()
            _ = v187
            _ = v188
        }
    case 4294967295:
        var v189 uint64 = 4294967295
        _ = v189
    }
    if len(v1) > 0 {
        v1[0] = v3
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v141:", v141)
    fmt.Println("v142:", v142)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v141
    _ = v142
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
