package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int }
type GenS2 struct { F0 int8; F1 string; F2 int64 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS5 struct { F0 int64; F1 uint8; F2 string }
type GenS6 struct { F0 uint16; F1 int32; F2 bool }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int32; F1 uint16 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 bool; F1 string; F2 uint8 }
type GenS14 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS15 struct { F0 string; F1 int }
type GenS16 struct { F0 string }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 string; F1 int8 }
type GenS19 struct { F0 int8; F1 bool; F2 int }
type GenS20 struct { F0 int16 }
type GenS21 struct { F0 int32; F1 uint16; F2 uint }
type GenS22 struct { F0 uint16; F1 bool; F2 int8 }
type GenS23 struct { F0 int64; F1 int32; F2 uint32 }
type GenS24 struct { F0 string; F1 int }
type GenS25 struct { F0 uint64; F1 uint64 }
type GenS26 struct { F0 int8; F1 int16 }
type GenS27 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS28 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS29 struct { F0 int32 }
type GenS30 struct { F0 int8 }

func main() {
    var v1 uint16 = 42
    v2 := []uint32{4294967295, 65535, 0, 2}
    v3 := make(map[string]uint32)
    var v4 bool = false
    v5 := make(map[string]string)
    var v6 uint = 18446744073709551615
    if v4 {
        if ((v1 > v1) && true) {
            fmt.Println(v6, v6)
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := GenS1{F0: uint(v1), F1: fi7, F2: fi7}
                v4 = (1 >= v1)
                if len(v2) > 0 {
                    v2[0] = 19
                }
                _ = fi7
                _ = v8
            }
            v9 := GenS2{F0: -87, F1: "foo bar", F2: -381}
            _ = v9
        } else {
            v1 = (v1 & v1)
            var v10 int8 = 42
            _ = v10
        }
        delete(v5, "hello")
    }
    fmt.Println(v6)
    if v4 {
        v11 := GenS2{F0: int8(v6), F1: "\n\t", F2: -993}
        _ = v11
    }
    v12 := GenS1{F0: (v6 ^ v6), F1: 0, F2: 492}
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v15 := v3["hello"]
            switch fi14 {
            case 32768:
                fmt.Println(v6)
                v6 += 1
                v15 = v15
            case -2147483648:
                if v4 {
                    var v16 int16 = 863
                    var v17 string = "\n\t"
                    _ = v16
                    _ = v17
                } else {
                    v1 = v1
                    v1 = v1
                    v2 = append(v2, v15)
                }
            }
            _ = fi14
            _ = v15
        }
        switch "foo bar" {
        case "":
            if false {
                var v18 float64 = -2.25
                _ = v18
            }
        case "a":
            if (((!v4) || v4) && (true || false)) {
                for fi19 := 0; fi19 < 1; fi19++ {
                    var v20 uint32 = 42
                    _ = fi19
                    _ = v20
                }
            }
        }
        v4 = v4
        _ = fi13
    }
    if v4 {
        fmt.Println(v1)
        _ = len(v2)
        var v21 uint8 = 128
        _ = v21
    } else {
        v4 = (127 == 0)
        if v4 {
            v6 -= 65535
            var v22 uint32 = 42
            if len(v2) > 0 {
                v2[0] = (v22 - (v22 + v22))
            }
            _ = v22
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            var v24 int32 = 776
            if false {
                for fi25 := 0; fi25 < 14; fi25++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch fi25 {
                    case -1:
                        v1 += (v1 + 255)
                        v26 := GenS3{F0: 619, F1: "a"}
                        _ = v26
                    case -9223372036854775808:
                        var v27 int8 = -128
                        _ = v27
                    }
                    if v4 {
                        v28 := GenS4{F0: 624, F1: 0, F2: 128}
                        if ((42 <= v24) && v4) {
                            var v29 uint16 = v1
                            var v30 int8 = -42
                            v6 = v6
                            var v31 string = "hello"
                            fmt.Println(v30, v1)
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        _ = v28
                    } else {
                        fmt.Println(v4, fi25)
                        var v32 float32 = float32(1e10)
                        var v33 bool = true
                        _ = v32
                        _ = v33
                    }
                    _ = fi25
                }
            } else {
                v4 = (1000 <= -1)
                for fi34 := 0; fi34 < 10; fi34++ {
                    if (-9223372036854775808 <= 42) {
                        var v35 bool = v4
                        _ = v35
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            v36 := GenS4{F0: -1, F1: 225, F2: 1}
                            _ = v36
                        case "a":
                            v1 -= 42
                        default:
                            fmt.Println(v24, fi34, v4)
                        }
                        v37 := GenS5{F0: 568, F1: 128, F2: ""}
                        if v4 {
                            v24 += 42
                        }
                        _ = v37
                    }
                    fmt.Println(v1, v24, v4)
                    v38 := GenS5{F0: -1, F1: 255, F2: "\n\t"}
                    v39 := GenS2{F0: -18, F1: "a", F2: 42}
                    _ = fi34
                    _ = v38
                    _ = v39
                }
                var v40 uint8 = 77
                _ = v40
            }
            v41 := GenS6{F0: v1, F1: v24, F2: true}
            _ = fi23
            _ = v24
            _ = v41
        }
    }
    var v42 int64 = 1
    if true {
        switch "" {
        case "hello":
            v43 := GenS6{F0: v1, F1: 42, F2: v4}
            _ = v43
        case "\n\t":
            for fi44 := 0; fi44 < 7; fi44++ {
                if ((2 <= 91) || ((0 <= -686) && true)) {
                    v42 = 337
                    switch fi44 {
                    case -1:
                        var v45 float64 = 0.0
                        _ = v45
                    case 128:
                        for fi46 := 0; fi46 < 4; fi46++ {
                            if ("foo bar" >= ("hello" + "foo bar")) {
                                if len(v2) > 0 {
                                    v2[0] = 766
                                }
                                v4 = (-1 != 42)
                                var v47 int = (fi44 ^ 2147483647)
                                _ = v47
                            }
                            _ = fi46
                        }
                    case -386:
                        if (true || (!v4)) {
                            if ((((("" + "foo bar") == "a") || (68 >= 0)) || v4) && (!false)) {
                                if len(v2) > 0 {
                                    v2[0] = 245
                                }
                            } else {
                                for fi48 := 0; fi48 < 9; fi48++ {
                                    v5[("" + (("a" + "") + "a"))] = ("foo bar" + "")
                                    fmt.Println(v1, v1, v1)
                                    if true {
                                        _ = len(v5)
                                        switch ("hello" + "hello") {
                                        case "\n\t":
                                            v4 = (2147483647 < fi44)
                                        case "a":
                                            if v4 {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v5[("a" + ((("\n\t" + "hello") + "\n\t") + ""))] = "\n\t"
                                                case "":
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v49 := GenS4{F0: v42, F1: 143, F2: 127}
                                                        _ = v49
                                                    case "a":
                                                        v50 := GenS2{F0: -1, F1: "", F2: (v42 * int64(v1))}
                                                        _ = v50
                                                    case "":
                                                        if v4 {
                                                            var v51 uint = (v6 >> (((v6 / ((v6) | 1))) & 63))
                                                            _ = v51
                                                        } else {
                                                            v52 := GenS7{F0: 696}
                                                            if false {
                                                                fmt.Println(fi48)
                                                                v4 = (!(127 > 44))
                                                                fmt.Println(v6, v4, fi44)
                                                            } else {
                                                                if (int16(v42) <= -213) {
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        switch "a" {
                                                                        case "a":
                                                                            v42 = 58
                                                                        case "foo bar":
                                                                            v53 := GenS6{F0: v1, F1: 42, F2: v4}
                                                                            _ = v53
                                                                        }
                                                                    case "hello":
                                                                        v6 += (v6 + 4294967295)
                                                                    case "\n\t":
                                                                        v3["a"] = 2147483647
                                                                    }
                                                                    var v54 int8 = 42
                                                                    if v4 {
                                                                        fmt.Println(fi44)
                                                                        var v55 int = fi48
                                                                        _ = v55
                                                                    }
                                                                    _ = v54
                                                                } else {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 1
                                                                    }
                                                                    v1 -= (v1 / ((v1) | 1))
                                                                }
                                                            }
                                                            v4 = v4
                                                            _ = v52
                                                        }
                                                        switch fi48 {
                                                        case -2147483648:
                                                            _ = len(v2)
                                                            if v4 {
                                                                v1 = (v1 - (v1 << ((v6) & 63)))
                                                                v2 = append(v2, 0)
                                                                switch (fi44 * (fi44 ^ fi48)) {
                                                                case 127:
                                                                    var v56 float32 = ((-2.25 - -2.25) / ((0.0 + ((-2.25 - (0.0 * float32(-2.25))) + ((((0.1 + (-2.25 / ((1.5 + 1.5) / -2.25))) - ((0.1 - 1.5) + 1e10)) * 3.14159) * 0.0))) * 1e10))
                                                                    _ = v56
                                                                default:
                                                                    if ((1 < v1) && true) {
                                                                        v1 += uint16(v6)
                                                                        if (v1 <= (v1 & v1)) {
                                                                            v4 = ((v42 - (v42 % (((v42 ^ (v42 << ((v6) & 63)))) | 1))) < v42)
                                                                            fmt.Println(v6)
                                                                            if (1 >= 1) {
                                                                                var v57 string = ""
                                                                                _ = v57
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            } else {
                                                                switch fi44 {
                                                                case 976:
                                                                    if (v4 && v4) {
                                                                        var v58 float32 = -2.25
                                                                        switch ("foo bar" + "") {
                                                                        case "hello":
                                                                            fmt.Println(v6)
                                                                        case "":
                                                                            _ = len(v5)
                                                                            var v59 float64 = 0.0
                                                                            _ = v59
                                                                        case "foo bar":
                                                                            delete(v5, "")
                                                                        }
                                                                        if v4 {
                                                                            var v60 int8 = -128
                                                                            v60 = 66
                                                                            if false {
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    if ((0.0 * v58) == 1e10) {
                                                                                        if (false || (4294967295 <= 956)) {
                                                                                            if v4 {
                                                                                                if true {
                                                                                                    var v61 int8 = (v60 * int8(v42))
                                                                                                    v62 := GenS8{F0: 1, F1: (v1 * (v1 << (((v6 + (v6 >> (((^(v6))) & 63)))) & 63)))}
                                                                                                    v61 = (v60 + -128)
                                                                                                    _ = v61
                                                                                                    _ = v62
                                                                                                } else {
                                                                                                    var v63 string = ("hello" + "a")
                                                                                                    var v64 int16 = 32767
                                                                                                    _ = v63
                                                                                                    _ = v64
                                                                                                }
                                                                                                var v65 bool = ((2 > 0) || (!false))
                                                                                                _ = len(v2)
                                                                                                _ = v65
                                                                                            }
                                                                                            v42 -= (v42 | int64(v6))
                                                                                            var v66 uint8 = 0
                                                                                            if (false && false) {
                                                                                                var v67 int16 = -169
                                                                                                if v4 {
                                                                                                    switch -2147483648 {
                                                                                                    case -733:
                                                                                                        var v68 float64 = (0.1 + 0.1)
                                                                                                        _ = v68
                                                                                                    default:
                                                                                                        v67 = v67
                                                                                                    }
                                                                                                    v69 := GenS9{F0: 290}
                                                                                                    v70 := GenS5{F0: -9223372036854775808, F1: (^(v66)), F2: "a"}
                                                                                                    switch "" {
                                                                                                    case "a":
                                                                                                        if v4 {
                                                                                                            var v71 int = 2147483647
                                                                                                            var v72 uint8 = v66
                                                                                                            v73 := GenS3{F0: -2147483648, F1: ""}
                                                                                                            _ = v71
                                                                                                            _ = v72
                                                                                                            _ = v73
                                                                                                        } else {
                                                                                                            v67 = (v67 ^ (v67 + (-(v67))))
                                                                                                            var v74 float32 = 3.14159
                                                                                                            v75 := GenS10{F0: true}
                                                                                                            _ = v74
                                                                                                            _ = v75
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        switch "" {
                                                                                                        case "\n\t":
                                                                                                            if (v58 >= (3.14159 + 0.1)) {
                                                                                                                switch (^(fi48)) {
                                                                                                                case 570:
                                                                                                                    var v76 uint16 = v1
                                                                                                                    _ = v76
                                                                                                                case -1:
                                                                                                                    var v77 int8 = v60
                                                                                                                    _ = v77
                                                                                                                }
                                                                                                                switch "a" {
                                                                                                                case "foo bar":
                                                                                                                    v78 := GenS11{F0: (^(v66)), F1: v66, F2: (v42 ^ (v42 ^ v42))}
                                                                                                                    _ = v78
                                                                                                                case "hello":
                                                                                                                    v79 := GenS12{F0: v42}
                                                                                                                    _ = v79
                                                                                                                case "":
                                                                                                                    var v80 int16 = (v67 | (v67 >> (((^(v6))) & 63)))
                                                                                                                    _ = v80
                                                                                                                default:
                                                                                                                    v2 = append(v2, 870)
                                                                                                                }
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            v6 = 18446744073709551615
                                                                                                        default:
                                                                                                            fmt.Println(v6)
                                                                                                        }
                                                                                                    default:
                                                                                                        if (v4 && true) {
                                                                                                            if v4 {
                                                                                                                switch ("hello" + "foo bar") {
                                                                                                                case "hello":
                                                                                                                    var v81 float32 = v58
                                                                                                                    _ = v81
                                                                                                                case "":
                                                                                                                    var v82 uint16 = 32767
                                                                                                                    if (uint(v1) == (v6 * 476)) {
                                                                                                                        switch (fi48 % (((-(fi48))) | 1)) {
                                                                                                                        case 127:
                                                                                                                            var v83 float64 = 0.0
                                                                                                                            _ = v83
                                                                                                                        case -525:
                                                                                                                            fmt.Println(fi48, fi48, v66)
                                                                                                                            _ = len(v5)
                                                                                                                            v58 -= (float32(float64((1.5 * v58))) + (v58 * (1e10 / (v58 * v58))))
                                                                                                                        case 883:
                                                                                                                            v6 -= v6
                                                                                                                        }
                                                                                                                        v84 := GenS13{F0: v4, F1: "\n\t", F2: 0}
                                                                                                                        _ = v84
                                                                                                                    }
                                                                                                                    if len(v2) >= 2 {
                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                    }
                                                                                                                    _ = v82
                                                                                                                case "a":
                                                                                                                    if v4 {
                                                                                                                        v60 -= 74
                                                                                                                    } else {
                                                                                                                        v85 := GenS13{F0: v4, F1: "foo bar", F2: 127}
                                                                                                                        if v4 {
                                                                                                                            v86 := GenS14{F0: 1, F1: 18446744073709551615, F2: 393}
                                                                                                                            _ = v86
                                                                                                                        }
                                                                                                                        v87 := GenS15{F0: (("a" + ("a" + "foo bar")) + "a"), F1: (fi48 % ((0) | 1))}
                                                                                                                        _ = v85
                                                                                                                        _ = v87
                                                                                                                    }
                                                                                                                default:
                                                                                                                    var v88 int64 = (v42 - (v42 ^ -9223372036854775808))
                                                                                                                    fmt.Println(fi48, v60)
                                                                                                                    if len(v2) >= 2 {
                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                    }
                                                                                                                    _ = v88
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v89 uint = (v6 - v6)
                                                                                                                var v90 float64 = 1e10
                                                                                                                _ = v89
                                                                                                                _ = v90
                                                                                                            }
                                                                                                            var v91 int8 = int8(v1)
                                                                                                            _ = v91
                                                                                                        } else {
                                                                                                            switch (("\n\t" + "\n\t") + ("a" + "a")) {
                                                                                                            case "":
                                                                                                                if (true || (-789 >= (fi48 >> (((v6 >> ((2) & 63))) & 63)))) {
                                                                                                                    fmt.Println(v58)
                                                                                                                    v92 := GenS16{F0: "foo bar"}
                                                                                                                    v66 -= (v66 / ((v66) | 1))
                                                                                                                    if v4 {
                                                                                                                        var v93 uint16 = (v1 & 35)
                                                                                                                        if len(v2) >= 2 {
                                                                                                                            v2 = v2[1:len(v2)]
                                                                                                                        }
                                                                                                                        var v94 uint = (^(v6))
                                                                                                                        _ = v93
                                                                                                                        _ = v94
                                                                                                                    }
                                                                                                                    _ = v92
                                                                                                                }
                                                                                                            default:
                                                                                                                if (!v4) {
                                                                                                                    switch (("" + (("a" + "") + ("" + "hello"))) + "foo bar") {
                                                                                                                    case "a":
                                                                                                                        switch (fi48 % ((fi48) | 1)) {
                                                                                                                        case 0:
                                                                                                                            v58 = -2.25
                                                                                                                            _ = len(v5)
                                                                                                                        }
                                                                                                                        switch "\n\t" {
                                                                                                                        case "foo bar":
                                                                                                                            switch "a" {
                                                                                                                            case "a":
                                                                                                                                if ((fi44 / ((fi48) | 1)) == fi48) {
                                                                                                                                    v95 := GenS15{F0: ("hello" + "a"), F1: -2147483648}
                                                                                                                                    _ = v95
                                                                                                                                }
                                                                                                                            }
                                                                                                                        case "\n\t":
                                                                                                                            var v96 int32 = -2147483648
                                                                                                                            v97 := GenS3{F0: v96, F1: ("\n\t" + ("\n\t" + ("" + "")))}
                                                                                                                            _ = v96
                                                                                                                            _ = v97
                                                                                                                        case "hello":
                                                                                                                            _ = len(v3)
                                                                                                                        }
                                                                                                                        var v98 int = (fi44 << ((817) & 63))
                                                                                                                        var v99 int8 = v60
                                                                                                                        v1 = v1
                                                                                                                        if len(v2) > 0 {
                                                                                                                            v2[0] = 42
                                                                                                                        }
                                                                                                                        v58 += 3.14159
                                                                                                                        _ = v98
                                                                                                                        _ = v99
                                                                                                                    case "\n\t":
                                                                                                                        if ((((1 > -1) || v4) && (0 != 18446744073709551615)) || ((v1 % (((^(v1))) | 1)) != (^(v1)))) {
                                                                                                                            v1 = (^(v1))
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            switch fi48 {
                                                                                                            case -569:
                                                                                                                fmt.Println(fi44)
                                                                                                            }
                                                                                                            if (v4 && (!((^(v1)) < 255))) {
                                                                                                                switch "\n\t" {
                                                                                                                case "a":
                                                                                                                    fmt.Println(v6, v66)
                                                                                                                case "hello":
                                                                                                                    fmt.Println(v1, v67)
                                                                                                                case "foo bar":
                                                                                                                    v100 := GenS5{F0: v42, F1: v66, F2: ("" + "\n\t")}
                                                                                                                    v101 := GenS17{F0: 523}
                                                                                                                    _ = v100
                                                                                                                    _ = v101
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v102 uint = 2
                                                                                                                _ = v102
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    var v103 uint32 = 4294967295
                                                                                                    var v104 int8 = v60
                                                                                                    _ = v103
                                                                                                    _ = v104
                                                                                                    _ = v69
                                                                                                    _ = v70
                                                                                                } else {
                                                                                                    var v105 uint32 = 1
                                                                                                    _ = v105
                                                                                                }
                                                                                                var v106 uint16 = v1
                                                                                                var v107 int = (fi48 + -1)
                                                                                                _ = v106
                                                                                                _ = v107
                                                                                                _ = v67
                                                                                            }
                                                                                            var v108 uint = (^(v6))
                                                                                            v5[""] = "foo bar"
                                                                                            _ = v108
                                                                                            _ = v66
                                                                                        }
                                                                                        if v4 {
                                                                                            v4 = ((false && v4) && (("" + "foo bar") > "\n\t"))
                                                                                            if (46 > 127) {
                                                                                                var v109 int = fi44
                                                                                                switch "a" {
                                                                                                case "":
                                                                                                    v58 = (1.5 * v58)
                                                                                                default:
                                                                                                    fmt.Println(v4)
                                                                                                }
                                                                                                v60 = (v60 * (v60 & -8))
                                                                                                _ = v109
                                                                                            } else {
                                                                                                v58 = v58
                                                                                                fmt.Println(v6, v1, v4)
                                                                                                if (v1 > (v1 << (((v6 - uint(v1))) & 63))) {
                                                                                                    fmt.Println(v1, v60)
                                                                                                    v42 += 9223372036854775807
                                                                                                } else {
                                                                                                    fmt.Println(v60, v6)
                                                                                                    var v110 int8 = 0
                                                                                                    _ = v110
                                                                                                }
                                                                                                var v111 uint64 = 18446744073709551615
                                                                                                v60 -= 28
                                                                                                if v4 {
                                                                                                    if (!v4) {
                                                                                                        switch "\n\t" {
                                                                                                        case "":
                                                                                                            v6 -= (v6 << ((250) & 63))
                                                                                                            if (!true) {
                                                                                                                if v4 {
                                                                                                                    _ = len(v5)
                                                                                                                    v1 = (^(v1))
                                                                                                                    if false {
                                                                                                                        var v112 int32 = -599
                                                                                                                        var v113 uint8 = 2
                                                                                                                        _ = v112
                                                                                                                        _ = v113
                                                                                                                    } else {
                                                                                                                        _ = len(v2)
                                                                                                                        switch "" {
                                                                                                                        case "":
                                                                                                                            v111 = (v111 >> ((v6) & 63))
                                                                                                                        case "hello":
                                                                                                                            switch "" {
                                                                                                                            case "a":
                                                                                                                                v58 -= v58
                                                                                                                            case "\n\t":
                                                                                                                                if v4 {
                                                                                                                                    v60 = v60
                                                                                                                                    v1 = 972
                                                                                                                                }
                                                                                                                            case "":
                                                                                                                                v4 = (("hello" + "hello") != "hello")
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    _ = len(v3)
                                                                                                                }
                                                                                                            } else {
                                                                                                                fmt.Println(v42, v1, v58)
                                                                                                                switch (fi44 + (fi48 << ((v6) & 63))) {
                                                                                                                case -472:
                                                                                                                    switch "a" {
                                                                                                                    case "a":
                                                                                                                        v114 := GenS18{F0: ("a" + ""), F1: v60}
                                                                                                                        _ = v114
                                                                                                                    case "":
                                                                                                                        v3[(("" + "\n\t") + ("" + ""))] = 0
                                                                                                                    case "\n\t":
                                                                                                                        v58 += -2.25
                                                                                                                    default:
                                                                                                                        v60 += v60
                                                                                                                    }
                                                                                                                case -2147483648:
                                                                                                                    v115 := GenS5{F0: (v42 + 0), F1: 228, F2: "hello"}
                                                                                                                    _ = v115
                                                                                                                case 32768:
                                                                                                                    var v116 uint32 = 65535
                                                                                                                    _ = v116
                                                                                                                }
                                                                                                                v6 += (v6 << (((v6 & v6)) & 63))
                                                                                                                v111 = 4294967295
                                                                                                                if len(v2) > 0 {
                                                                                                                    v2[0] = 1
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        var v117 int8 = 1
                                                                                                        _ = v117
                                                                                                    }
                                                                                                    var v118 uint = 985
                                                                                                    _ = v118
                                                                                                }
                                                                                                _ = v111
                                                                                            }
                                                                                            var v119 int = -1
                                                                                            _ = v119
                                                                                        }
                                                                                        switch "\n\t" {
                                                                                        case "hello":
                                                                                            fmt.Println(v4)
                                                                                        case "\n\t":
                                                                                            switch "" {
                                                                                            case "":
                                                                                                v42 += v42
                                                                                            case "hello":
                                                                                                v1 += 0
                                                                                            default:
                                                                                                v120 := v5["foo bar"]
                                                                                                v121 := GenS19{F0: (v60 / ((v60) | 1)), F1: v4, F2: (fi48 % ((4294967295) | 1))}
                                                                                                var v122 uint32 = 91
                                                                                                if (false && ((float64(0.0) == -2.25) && (fi44 > (fi48 >> (((v6 * 65535)) & 63))))) {
                                                                                                    var v123 int64 = v42
                                                                                                    if v4 {
                                                                                                        if (1 < (v60 & v60)) {
                                                                                                            if v4 {
                                                                                                                fmt.Println(v1, v122)
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        v60 = 52
                                                                                                        fmt.Println(v1, v42)
                                                                                                        delete(v5, v120)
                                                                                                    }
                                                                                                    var v124 uint64 = 42
                                                                                                    _ = v123
                                                                                                    _ = v124
                                                                                                } else {
                                                                                                    var v125 uint = v6
                                                                                                    _ = v125
                                                                                                }
                                                                                                _ = v120
                                                                                                _ = v121
                                                                                                _ = v122
                                                                                            }
                                                                                        case "a":
                                                                                            if v4 {
                                                                                                var v126 int16 = -295
                                                                                                fmt.Println(v58, v58, v126)
                                                                                                _ = v126
                                                                                            } else {
                                                                                                var v127 string = "hello"
                                                                                                _ = v127
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        v128 := v3["\n\t"]
                                                                                        var v129 int = (fi44 + (fi44 * (fi48 - fi44)))
                                                                                        if (!((v6 >> (((v6 >> ((750) & 63))) & 63)) == v6)) {
                                                                                            v130 := GenS7{F0: v128}
                                                                                            v129 += fi48
                                                                                            _ = v130
                                                                                        } else {
                                                                                            if (!(!(!(v4 || (("hello" + (("" + (("\n\t" + "a") + "a")) + "\n\t")) <= "a"))))) {
                                                                                                switch (("" + "foo bar") + "foo bar") {
                                                                                                case "\n\t":
                                                                                                    if (42 <= v42) {
                                                                                                        fmt.Println(fi44, fi48, v1)
                                                                                                        v3[""] = (v128 / ((v128) | 1))
                                                                                                    } else {
                                                                                                        v60 = v60
                                                                                                    }
                                                                                                case "hello":
                                                                                                    v1 += (^(v1))
                                                                                                case "":
                                                                                                    var v131 float32 = (v58 * (0.0 + (float32(1.5) - v58)))
                                                                                                    _ = v131
                                                                                                default:
                                                                                                    switch fi48 {
                                                                                                    case 123:
                                                                                                        var v132 int32 = 2147483647
                                                                                                        _ = v132
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                if ((v128 >= v128) || (((v6 << (((v6 - v6)) & 63)) != 18446744073709551615) && v4)) {
                                                                                                    var v133 uint = v6
                                                                                                    _ = v133
                                                                                                } else {
                                                                                                    fmt.Println(v60, v58, v129)
                                                                                                    switch (fi48 + fi44) {
                                                                                                    case 127:
                                                                                                        var v134 int64 = v42
                                                                                                        _ = v134
                                                                                                    case 4294967295:
                                                                                                        if false {
                                                                                                            v60 += int8(v6)
                                                                                                        }
                                                                                                    case 2147483647:
                                                                                                        v2 = append(v2, (v128 % ((v128) | 1)))
                                                                                                    }
                                                                                                    if (v4 || (!false)) {
                                                                                                        if len(v2) >= 2 {
                                                                                                            v2 = v2[1:len(v2)]
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            var v135 int64 = (v42 << (((^(v6))) & 63))
                                                                                            _ = v135
                                                                                        }
                                                                                        _ = v128
                                                                                        _ = v129
                                                                                    }
                                                                                case "\n\t":
                                                                                    v136 := GenS9{F0: 65535}
                                                                                    _ = v136
                                                                                case "foo bar":
                                                                                    switch "" {
                                                                                    case "":
                                                                                        delete(v3, "\n\t")
                                                                                    default:
                                                                                        v2 = append(v2, 65535)
                                                                                    }
                                                                                default:
                                                                                    var v137 float32 = v58
                                                                                    _ = v137
                                                                                }
                                                                                v42 = (^(v42))
                                                                            } else {
                                                                                var v138 int64 = 494
                                                                                fmt.Println(fi48)
                                                                                _ = len(v3)
                                                                                _ = v138
                                                                            }
                                                                            _ = v60
                                                                        }
                                                                        if v4 {
                                                                            _ = len(v5)
                                                                            _ = len(v2)
                                                                            fmt.Println(v1, v6)
                                                                        } else {
                                                                            var v139 int8 = 115
                                                                            _ = v139
                                                                        }
                                                                        var v140 int64 = v42
                                                                        _ = v140
                                                                        _ = v58
                                                                    }
                                                                default:
                                                                    v3[("a" + ("foo bar" + ""))] = 1
                                                                }
                                                            }
                                                        case -9223372036854775808:
                                                            v5["a"] = ((("a" + "foo bar") + "\n\t") + ("" + (("" + "a") + "foo bar")))
                                                            v141 := GenS19{F0: -128, F1: (v6 != (^(v6))), F2: 2147483647}
                                                            _ = v141
                                                        case 0:
                                                            if v4 {
                                                                v42 = -9223372036854775808
                                                                var v142 uint16 = 255
                                                                v143 := GenS20{F0: 1000}
                                                                _ = v142
                                                                _ = v143
                                                            } else {
                                                                fmt.Println(v4, v1, fi48)
                                                                if (18446744073709551615 <= 308) {
                                                                    v1 -= (v1 & v1)
                                                                    if v4 {
                                                                        var v144 uint32 = 65535
                                                                        fmt.Println(v4)
                                                                        var v145 float32 = 1e10
                                                                        _ = v144
                                                                        _ = v145
                                                                    }
                                                                } else {
                                                                    v146 := v3["a"]
                                                                    _ = v146
                                                                }
                                                            }
                                                        }
                                                        fmt.Println(v6, fi44, v42)
                                                    default:
                                                        v2 = append(v2, 2)
                                                    }
                                                default:
                                                    var v147 uint16 = (v1 % ((v1) | 1))
                                                    _ = v147
                                                }
                                                v6 = (v6 & (^(v6)))
                                            } else {
                                                var v148 uint8 = 255
                                                _ = len(v2)
                                                v1 = v1
                                                v149 := GenS21{F0: 42, F1: 0, F2: 946}
                                                _ = v148
                                                _ = v149
                                            }
                                        case "":
                                            if (!v4) {
                                                fmt.Println(v42, v42)
                                            }
                                        default:
                                            var v150 float64 = (1e10 / float64(-2.25))
                                            _ = v150
                                        }
                                    } else {
                                        var v151 uint8 = 126
                                        v6 += 42
                                        _ = v151
                                    }
                                    _ = fi48
                                }
                                var v152 int = (fi44 >> ((v6) & 63))
                                var v153 uint = (^(v6))
                                _ = v152
                                _ = v153
                            }
                            for fi154 := 0; fi154 < 3; fi154++ {
                                if (v4 || v4) {
                                    var v155 string = ""
                                    switch -21 {
                                    case -9223372036854775808:
                                        if (100 >= 1) {
                                            var v156 int32 = 2147483647
                                            if v4 {
                                                var v157 uint64 = 42
                                                v157 -= 390
                                                switch fi44 {
                                                case 9223372036854775807:
                                                    var v158 bool = v4
                                                    _ = v158
                                                case 127:
                                                    fmt.Println(v1)
                                                case -706:
                                                    if v4 {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            var v159 uint8 = 107
                                                            _ = v159
                                                        default:
                                                            var v160 int16 = 1
                                                            _ = v160
                                                        }
                                                    }
                                                default:
                                                    if (255 > 186) {
                                                        delete(v3, v155)
                                                        v6 = (v6 % ((2) | 1))
                                                        switch fi44 {
                                                        case -2147483648:
                                                            var v161 string = v155
                                                            _ = v161
                                                        }
                                                        fmt.Println(fi154)
                                                    }
                                                }
                                                v162 := GenS16{F0: (v155 + "foo bar")}
                                                _ = len(v2)
                                                _ = v157
                                                _ = v162
                                            }
                                            fmt.Println(v1, fi44)
                                            switch (((v155 + (v155 + (v155 + (v155 + ((("hello" + v155) + v155) + v155))))) + ("a" + "")) + "hello") {
                                            case "foo bar":
                                                fmt.Println(fi44, v4)
                                            }
                                            _ = v156
                                        }
                                    case -1:
                                        v155 = "foo bar"
                                    default:
                                        switch "foo bar" {
                                        case "foo bar":
                                            if ((-(v42)) == v42) {
                                                if v4 {
                                                    var v163 float64 = ((-2.25 * -2.25) + 3.14159)
                                                    _ = len(v2)
                                                    _ = v163
                                                }
                                                v2 = append(v2, 1)
                                            }
                                        case "a":
                                            if (127 == -1) {
                                                v4 = v4
                                                if false {
                                                    if v4 {
                                                        if (v155 == v155) {
                                                            fmt.Println(v6)
                                                            v164 := GenS22{F0: v1, F1: true, F2: int8(v42)}
                                                            _ = v164
                                                        } else {
                                                            if (v6 > (v6 + uint(v1))) {
                                                                if (1 < 24) {
                                                                    var v165 uint8 = 83
                                                                    var v166 int16 = 435
                                                                    v165 -= (^(v165))
                                                                    _ = v165
                                                                    _ = v166
                                                                } else {
                                                                    var v167 int16 = 475
                                                                    if v4 {
                                                                        fmt.Println(v167)
                                                                        v168 := GenS6{F0: v1, F1: -2147483648, F2: (v155 > "hello")}
                                                                        _ = v168
                                                                    }
                                                                    switch 9223372036854775807 {
                                                                    case 345:
                                                                        v2 = append(v2, 609)
                                                                    case -9223372036854775808:
                                                                        if false {
                                                                            if v4 {
                                                                                switch v155 {
                                                                                case "":
                                                                                    if (0 >= 0) {
                                                                                        var v169 int = fi44
                                                                                        if v4 {
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = 2147483647
                                                                                            }
                                                                                            if v4 {
                                                                                                switch v155 {
                                                                                                case "hello":
                                                                                                    var v170 string = "foo bar"
                                                                                                    var v171 uint64 = 0
                                                                                                    if (true && v4) {
                                                                                                        if (v4 || false) {
                                                                                                            switch (v155 + v170) {
                                                                                                            case "\n\t":
                                                                                                                v172 := GenS7{F0: 2147483648}
                                                                                                                _ = v172
                                                                                                            case "":
                                                                                                                v42 -= v42
                                                                                                            default:
                                                                                                                v1 = (^(v1))
                                                                                                                _ = len(v3)
                                                                                                                if v4 {
                                                                                                                    v173 := GenS11{F0: uint8(v171), F1: 150, F2: (v42 + v42)}
                                                                                                                    v42 = (v42 & v42)
                                                                                                                    _ = v173
                                                                                                                } else {
                                                                                                                    v174 := GenS16{F0: v170}
                                                                                                                    if len(v2) > 0 {
                                                                                                                        v2[0] = 2
                                                                                                                    }
                                                                                                                    if (((v167 % ((341) | 1)) == (v167 >> ((uint(v171)) & 63))) && v4) {
                                                                                                                        v175 := GenS15{F0: v170, F1: v169}
                                                                                                                        v176 := GenS17{F0: 2147483647}
                                                                                                                        var v177 uint64 = v171
                                                                                                                        _ = v175
                                                                                                                        _ = v176
                                                                                                                        _ = v177
                                                                                                                    } else {
                                                                                                                        if ((v1 ^ (^(v1))) == (v1 << ((v6) & 63))) {
                                                                                                                            var v178 string = (v170 + (v170 + (v155 + v155)))
                                                                                                                            fmt.Println(v167, v42)
                                                                                                                            fmt.Println(fi44, v6, v169)
                                                                                                                            _ = v178
                                                                                                                        }
                                                                                                                        switch 32768 {
                                                                                                                        case -2147483648:
                                                                                                                            v167 = -32768
                                                                                                                        case -821:
                                                                                                                            v42 = -1
                                                                                                                        }
                                                                                                                        fmt.Println(v1)
                                                                                                                        if v4 {
                                                                                                                            fmt.Println(v155)
                                                                                                                            if v4 {
                                                                                                                                v179 := GenS6{F0: v1, F1: 42, F2: (101 > 108)}
                                                                                                                                switch "a" {
                                                                                                                                case "a":
                                                                                                                                    switch v155 {
                                                                                                                                    case "\n\t":
                                                                                                                                        v180 := GenS18{F0: v155, F1: 0}
                                                                                                                                        _ = v180
                                                                                                                                    case "hello":
                                                                                                                                        v42 = 0
                                                                                                                                        var v181 uint = (v6 << (((v6 ^ 18446744073709551615)) & 63))
                                                                                                                                        fmt.Println(v170, v170, v1)
                                                                                                                                        _ = v181
                                                                                                                                    }
                                                                                                                                case "\n\t":
                                                                                                                                    fmt.Println(v171, v1)
                                                                                                                                case "hello":
                                                                                                                                    v1 = v1
                                                                                                                                }
                                                                                                                                v182 := GenS14{F0: (v1 + (^(v1))), F1: uint64(v1), F2: v171}
                                                                                                                                _ = v179
                                                                                                                                _ = v182
                                                                                                                            } else {
                                                                                                                                fmt.Println(v169, v155, v170)
                                                                                                                            }
                                                                                                                            switch v155 {
                                                                                                                            case "":
                                                                                                                                if v4 {
                                                                                                                                    v183 := GenS6{F0: (v1 << ((90) & 63)), F1: -2147483648, F2: ((v1 ^ (^(v1))) == v1)}
                                                                                                                                    var v184 string = ("\n\t" + (v155 + v170))
                                                                                                                                    _ = len(v3)
                                                                                                                                    _ = v183
                                                                                                                                    _ = v184
                                                                                                                                }
                                                                                                                            case "hello":
                                                                                                                                v2 = append(v2, 331)
                                                                                                                            case "foo bar":
                                                                                                                                v2 = append(v2, 50)
                                                                                                                            }
                                                                                                                            fmt.Println(v6, v42)
                                                                                                                        }
                                                                                                                        v185 := GenS7{F0: 1}
                                                                                                                        _ = v185
                                                                                                                    }
                                                                                                                    v4 = v4
                                                                                                                    if v4 {
                                                                                                                        if (223 < 1) {
                                                                                                                            fmt.Println(v169, v1, v6)
                                                                                                                            v3[v155] = 1
                                                                                                                            v2 = append(v2, 837)
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        if len(v2) > 0 {
                                                                                                                            v2[0] = 786
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v174
                                                                                                                }
                                                                                                            }
                                                                                                            switch (v155 + v170) {
                                                                                                            case "hello":
                                                                                                                delete(v5, ((((v155 + v155) + "hello") + v155) + (v155 + v170)))
                                                                                                            case "":
                                                                                                                v186 := GenS9{F0: 2147483648}
                                                                                                                _ = v186
                                                                                                            default:
                                                                                                                fmt.Println(v4, v169)
                                                                                                            }
                                                                                                        }
                                                                                                        if ((v167 << ((v6) & 63)) > (^(v167))) {
                                                                                                            v170 = "foo bar"
                                                                                                            fmt.Println(fi44)
                                                                                                            v187 := GenS2{F0: int8(fi44), F1: "foo bar", F2: (v42 << ((v6) & 63))}
                                                                                                            _ = v187
                                                                                                        }
                                                                                                        v5[("" + (v170 + (((v155 + ((v170 + ("foo bar" + v170)) + v155)) + v170) + v155)))] = ""
                                                                                                    }
                                                                                                    switch (v169 & fi154) {
                                                                                                    case -269:
                                                                                                        v188 := GenS13{F0: v4, F1: "\n\t", F2: 255}
                                                                                                        _ = v188
                                                                                                    case 0:
                                                                                                        v189 := GenS23{F0: int64(fi44), F1: -2147483648, F2: 2147483647}
                                                                                                        _ = v189
                                                                                                    }
                                                                                                    switch v170 {
                                                                                                    case "hello":
                                                                                                        fmt.Println(v167)
                                                                                                        if v4 {
                                                                                                            fmt.Println(v6, fi44)
                                                                                                        } else {
                                                                                                            fmt.Println(v6, v155, v6)
                                                                                                        }
                                                                                                    case "\n\t":
                                                                                                        var v190 uint16 = v1
                                                                                                        _ = v190
                                                                                                    case "foo bar":
                                                                                                        delete(v3, v170)
                                                                                                    }
                                                                                                    _ = v170
                                                                                                    _ = v171
                                                                                                case "":
                                                                                                    var v191 uint32 = 811
                                                                                                    _ = v191
                                                                                                default:
                                                                                                    switch fi44 {
                                                                                                    case 127:
                                                                                                        if v4 {
                                                                                                            var v192 uint = v6
                                                                                                            var v193 string = v155
                                                                                                            var v194 uint8 = 255
                                                                                                            _ = v192
                                                                                                            _ = v193
                                                                                                            _ = v194
                                                                                                        }
                                                                                                    case 0:
                                                                                                        var v195 uint32 = 2147483648
                                                                                                        _ = v195
                                                                                                    }
                                                                                                }
                                                                                                v167 = (v167 >> (((^(v6))) & 63))
                                                                                            }
                                                                                        } else {
                                                                                            var v196 uint = v6
                                                                                            v197 := GenS13{F0: v4, F1: "", F2: 42}
                                                                                            _ = v196
                                                                                            _ = v197
                                                                                        }
                                                                                        _ = v169
                                                                                    }
                                                                                    fmt.Println(v4, v6)
                                                                                default:
                                                                                    if v4 {
                                                                                        if v4 {
                                                                                            switch (v155 + "") {
                                                                                            case "foo bar":
                                                                                                v198 := v3[""]
                                                                                                _ = v198
                                                                                            case "\n\t":
                                                                                                var v199 int8 = 42
                                                                                                _ = v199
                                                                                            case "a":
                                                                                                switch (v155 + v155) {
                                                                                                case "a":
                                                                                                    var v200 string = (((v155 + v155) + (v155 + "foo bar")) + v155)
                                                                                                    _ = v200
                                                                                                case "hello":
                                                                                                    v201 := GenS24{F0: ("a" + v155), F1: (fi154 / ((9223372036854775807) | 1))}
                                                                                                    _ = v201
                                                                                                case "":
                                                                                                    v202 := GenS12{F0: (v42 | 0)}
                                                                                                    _ = v202
                                                                                                default:
                                                                                                    delete(v5, v155)
                                                                                                }
                                                                                            default:
                                                                                                if (v4 && (0.0 < 0.1)) {
                                                                                                    var v203 int = fi154
                                                                                                    var v204 int8 = -84
                                                                                                    var v205 bool = v4
                                                                                                    _ = v203
                                                                                                    _ = v204
                                                                                                    _ = v205
                                                                                                } else {
                                                                                                    if v4 {
                                                                                                        delete(v3, v155)
                                                                                                    } else {
                                                                                                        var v206 int32 = 0
                                                                                                        v6 -= v6
                                                                                                        _ = len(v5)
                                                                                                        _ = v206
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        var v207 int64 = (-(v42))
                                                                                        _ = v207
                                                                                    } else {
                                                                                        var v208 float32 = (1e10 - -2.25)
                                                                                        if false {
                                                                                            v4 = (v6 > (v6 + 65535))
                                                                                            switch (v155 + v155) {
                                                                                            case "":
                                                                                                switch 4294967295 {
                                                                                                case 697:
                                                                                                    v5[(v155 + (v155 + (v155 + v155)))] = v155
                                                                                                default:
                                                                                                    var v209 bool = true
                                                                                                    _ = v209
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v6 -= v6
                                                                                        _ = v208
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v210 int = fi44
                                                                            var v211 float64 = float64(1.5)
                                                                            var v212 float64 = (v211 + ((0.0 / ((((((0.1 - ((1e10 * 1.5) / (v211 / v211))) * ((v211 + (v211 - float64(0.0))) * -2.25)) * v211) - v211) / v211) * float64(3.14159))) + v211))
                                                                            _ = v210
                                                                            _ = v211
                                                                            _ = v212
                                                                        } else {
                                                                            var v213 int8 = -1
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                v213 = 100
                                                                            default:
                                                                                var v214 uint8 = 255
                                                                                _ = v214
                                                                            }
                                                                            fmt.Println(v42, v213, fi44)
                                                                            _ = v213
                                                                        }
                                                                    }
                                                                    _ = v167
                                                                }
                                                            }
                                                            if (!false) {
                                                                v215 := GenS25{F0: 4294967295, F1: 1}
                                                                _ = v215
                                                            } else {
                                                                if ((("\n\t" + v155) + "foo bar") > ((("" + v155) + (v155 + (v155 + (v155 + v155)))) + (v155 + "a"))) {
                                                                    v5[(v155 + "")] = v155
                                                                    v4 = (-1 != 42)
                                                                    if (-2147483648 <= 1) {
                                                                        v4 = v4
                                                                    } else {
                                                                        switch v155 {
                                                                        case "hello":
                                                                            v216 := GenS1{F0: (v6 & 42), F1: fi44, F2: 128}
                                                                            _ = v216
                                                                        default:
                                                                            fmt.Println(fi154, v6, v1)
                                                                        }
                                                                        var v217 int64 = -786
                                                                        _ = v217
                                                                    }
                                                                }
                                                            }
                                                            v155 = (v155 + ((v155 + "hello") + ((v155 + (v155 + "\n\t")) + "a")))
                                                        }
                                                        if len(v2) > 0 {
                                                            v2[0] = 0
                                                        }
                                                    } else {
                                                        if v4 {
                                                            v6 -= v6
                                                        } else {
                                                            v6 = v6
                                                        }
                                                    }
                                                    var v218 string = v155
                                                    if (18446744073709551615 <= 368) {
                                                        switch -2147483648 {
                                                        case 9223372036854775807:
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        case 0:
                                                            var v219 int16 = 848
                                                            _ = v219
                                                        case 2147483647:
                                                            v42 -= (-(v42))
                                                        default:
                                                            fmt.Println(v6, fi44, fi154)
                                                        }
                                                    } else {
                                                        v218 = "\n\t"
                                                        if (!(!(v4 && false))) {
                                                            v220 := GenS18{F0: "a", F1: -128}
                                                            fmt.Println(v218, fi154, fi44)
                                                            _ = v220
                                                        } else {
                                                            v221 := GenS14{F0: (v1 / (((v1 >> (((^(v6))) & 63))) | 1)), F1: 0, F2: 0}
                                                            var v222 int32 = 0
                                                            _ = v221
                                                            _ = v222
                                                        }
                                                    }
                                                    _ = v218
                                                }
                                                if (v4 || v4) {
                                                    var v223 string = v155
                                                    switch fi154 {
                                                    case 982:
                                                        var v224 uint8 = 127
                                                        _ = v224
                                                    }
                                                    if (false && true) {
                                                        v1 = v1
                                                    } else {
                                                        v2 = append(v2, 4294967295)
                                                        var v225 float32 = 3.14159
                                                        var v226 uint = (v6 * (v6 % ((478) | 1)))
                                                        var v227 int16 = 32767
                                                        var v228 float32 = 3.14159
                                                        _ = v225
                                                        _ = v226
                                                        _ = v227
                                                        _ = v228
                                                    }
                                                    var v229 uint16 = 1
                                                    v223 = "foo bar"
                                                    var v230 string = "\n\t"
                                                    v231 := GenS4{F0: 354, F1: 50, F2: 0}
                                                    switch (((((("" + ((v230 + v155) + "hello")) + "a") + (v230 + (v223 + v223))) + ((v230 + "foo bar") + ((v223 + "foo bar") + "foo bar"))) + v230) + v155) {
                                                    case "\n\t":
                                                        if (-78 > -1) {
                                                            v4 = (127 < 127)
                                                            if (128 >= 10) {
                                                                v1 -= v1
                                                                _ = len(v2)
                                                            } else {
                                                                var v232 uint32 = uint32(v1)
                                                                v2 = append(v2, v232)
                                                                fmt.Println(v230)
                                                                v42 = v42
                                                                if ((!v4) && v4) {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                }
                                                                if (int(v42) <= int(v6)) {
                                                                    if ((0.1 - -2.25) > 3.14159) {
                                                                        var v233 uint = uint(v232)
                                                                        _ = v233
                                                                    }
                                                                    v4 = (4294967295 <= 0)
                                                                } else {
                                                                    v155 = v223
                                                                }
                                                                _ = v232
                                                            }
                                                        }
                                                    }
                                                    var v234 int64 = v42
                                                    _ = v223
                                                    _ = v229
                                                    _ = v230
                                                    _ = v231
                                                    _ = v234
                                                }
                                                if (726 > 856) {
                                                    var v235 int16 = 1000
                                                    _ = len(v2)
                                                    v236 := GenS26{F0: 1, F1: v235}
                                                    _ = v235
                                                    _ = v236
                                                }
                                                var v237 int32 = 42
                                                _ = v237
                                            }
                                        }
                                    }
                                    _ = v155
                                }
                                var v238 int16 = 144
                                v239 := GenS22{F0: 0, F1: v4, F2: 100}
                                delete(v3, ((("hello" + "hello") + "foo bar") + ("foo bar" + "")))
                                if (!false) {
                                    var v240 int64 = v42
                                    if v4 {
                                        if (v4 || (42 != fi44)) {
                                            var v241 bool = (-1 <= -103)
                                            var v242 float32 = 0.0
                                            _ = v241
                                            _ = v242
                                        }
                                        _ = len(v3)
                                        fmt.Println(v240, fi154)
                                        v243 := GenS27{F0: v6, F1: 610, F2: 127}
                                        _ = v243
                                    } else {
                                        v4 = true
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    _ = v240
                                } else {
                                    switch "hello" {
                                    case "foo bar":
                                        if (("foo bar" + "") != "") {
                                            switch (("foo bar" + ("a" + "hello")) + (("hello" + (("" + ("a" + ("a" + ("" + ("foo bar" + ("foo bar" + "hello")))))) + "")) + "")) {
                                            case "foo bar":
                                                v4 = (!(1e10 <= float32(0.0)))
                                            }
                                            v244 := GenS28{F0: 0, F1: 0, F2: 1}
                                            _ = v244
                                        } else {
                                            v6 += v6
                                            switch fi44 {
                                            case 32767:
                                                var v245 string = "a"
                                                _ = v245
                                            case -1:
                                                var v246 uint16 = v1
                                                _ = v246
                                            }
                                        }
                                    case "\n\t":
                                        v4 = false
                                    }
                                }
                                switch ("" + (("" + (("\n\t" + "hello") + "")) + (("" + "") + "a"))) {
                                case "hello":
                                    v4 = v4
                                case "":
                                    if v4 {
                                        var v247 int8 = 4
                                        var v248 int16 = -71
                                        _ = v247
                                        _ = v248
                                    }
                                }
                                _ = fi154
                                _ = v238
                                _ = v239
                            }
                        } else {
                            v6 = v6
                        }
                    }
                    if true {
                        if (0.0 == (0.0 - (0.0 + 0.0))) {
                            var v249 int32 = -1
                            v250 := GenS4{F0: v42, F1: 68, F2: 132}
                            var v251 uint8 = 255
                            _ = v249
                            _ = v250
                            _ = v251
                        } else {
                            v42 = v42
                        }
                    } else {
                        v42 = (v42 >> (((^(v6))) & 63))
                        switch (fi44 - (fi44 - fi44)) {
                        case 127:
                            v5["a"] = ""
                        }
                        if (65535 >= 753) {
                            var v252 int16 = 0
                            _ = v252
                        }
                    }
                }
                var v253 int8 = 29
                _ = fi44
                _ = v253
            }
        default:
            v6 = v6
        }
    }
    switch 32767 {
    case 32767:
        v1 = v1
    }
    for fi254 := 0; fi254 < 19; fi254++ {
        v42 = v42
        v6 += v6
        var v255 int32 = -2147483648
        _ = fi254
        _ = v255
    }
    delete(v3, ("" + "hello"))
    v256 := GenS29{F0: 42}
    var v257 int64 = (v42 ^ v42)
    switch "" {
    case "\n\t":
        switch 4294967295 {
        case 42:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v258 := GenS30{F0: int8(v1)}
            _ = v258
        }
    }
    switch 128 {
    case 259:
        _ = len(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v257:", v257)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v256
    _ = v257
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
