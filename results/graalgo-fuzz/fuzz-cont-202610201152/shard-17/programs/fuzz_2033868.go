package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 string }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint; F1 string; F2 int32 }
type GenS4 struct { F0 uint; F1 int16; F2 uint32 }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint; F1 uint16; F2 uint }
type GenS9 struct { F0 int32; F1 uint64 }
type GenS10 struct { F0 int64; F1 uint; F2 bool }
type GenS11 struct { F0 uint; F1 int; F2 int16 }
type GenS12 struct { F0 uint64; F1 int32 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int32; F1 int32; F2 uint8 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 uint; F1 int32; F2 int8 }
type GenS17 struct { F0 int32; F1 uint8; F2 int8 }
type GenS18 struct { F0 uint; F1 int }
type GenS19 struct { F0 int }
type GenS20 struct { F0 int16; F1 uint32; F2 bool }
type GenS21 struct { F0 int32; F1 int32; F2 uint }
type GenS22 struct { F0 int16 }
type GenS23 struct { F0 uint64; F1 bool; F2 int }
type GenS24 struct { F0 int16; F1 uint }
type GenS25 struct { F0 uint8 }
type GenS26 struct { F0 string; F1 int }
type GenS27 struct { F0 int16; F1 uint8; F2 int }
type GenS28 struct { F0 uint64; F1 string; F2 bool }
func genFn0(p0 uint64, p1 int16) uint32 {
    var v6 int16 = (p1 / (((p1 * 1000)) | 1))
    v6 -= (p1 | (p1 << ((42) & 63)))
    _ = p0
    _ = p1
    _ = v6
    return uint32(p0)
}
func genFn1(p0 int8, p1 int, p2 uint16) int {
    var v7 uint8 = 128
    v8 := GenS1{F0: 42, F1: 1, F2: ("hello" + ("a" + ""))}
    for fi9 := 0; fi9 < 14; fi9++ {
        v7 = (v7 / ((v7) | 1))
        v10 := GenS1{F0: 32768, F1: 0, F2: "\n\t"}
        if true {
            for fi11 := 0; fi11 < 12; fi11++ {
                v7 -= v7
                fmt.Println(fi9)
                if false {
                    var v12 int64 = -9223372036854775808
                    _ = v12
                } else {
                    switch (p1 + fi9) {
                    case -910:
                        v13 := []int{fi9, p1, (fi11 + (p1 % (((fi9 << ((4294967295) & 63))) | 1))), -1}
                        _ = v13
                    }
                }
                v14 := GenS1{F0: 128, F1: 1, F2: "\n\t"}
                switch p1 {
                case -447:
                    v15 := []int{(p1 - fi9), p1, (p1 + 42), 32767}
                    v7 = (v7 ^ v7)
                    var v16 int8 = (p0 | (p0 << ((148) & 63)))
                    _ = v15
                    _ = v16
                case 601:
                    if false {
                        v17 := []int{fi11}
                        if false {
                            if len(v17) > 0 {
                                v17[0] = (fi9 - 9223372036854775807)
                            }
                            var v18 uint8 = v7
                            if false {
                                p2 += p2
                                if (-2.25 > 1e10) {
                                    if false {
                                        switch (fi9 | (p1 >> ((695) & 63))) {
                                        case -374:
                                            var v19 int64 = -9223372036854775808
                                            _ = v19
                                        case 32767:
                                            var v20 int = (fi9 - (fi11 << ((192) & 63)))
                                            _ = v20
                                        case 0:
                                            var v21 uint = 977
                                            _ = v21
                                        default:
                                            var v22 int8 = 35
                                            _ = v22
                                        }
                                        v7 = v18
                                    }
                                } else {
                                    if (!(605 <= 65535)) {
                                        fmt.Println(fi11, p0, p1)
                                    } else {
                                        var v23 bool = false
                                        var v24 bool = (161 < 42)
                                        var v25 int = 42
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                    p2 += (^(p2))
                                    _ = len(v17)
                                    var v26 uint32 = 0
                                    var v27 int64 = 1
                                    _ = v26
                                    _ = v27
                                }
                                fmt.Println(p0)
                                v28 := GenS1{F0: -577, F1: 1, F2: "foo bar"}
                                _ = v28
                            } else {
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                            }
                            var v29 int = (^(fi11))
                            _ = v18
                            _ = v29
                        }
                        _ = v17
                    } else {
                        v30 := GenS1{F0: p1, F1: -935, F2: "hello"}
                        v31 := GenS1{F0: p1, F1: -406, F2: ("a" + ("a" + "hello"))}
                        _ = v30
                        _ = v31
                    }
                }
                _ = fi11
                _ = v14
            }
            fmt.Println(v7, fi9, p2)
            v32 := GenS2{F0: 1, F1: (^(p0))}
            _ = v32
        }
        _ = fi9
        _ = v10
    }
    var v33 uint = 579
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    _ = v7
    _ = v8
    return -974
}

func main() {
    var v1 int8 = 1
    v2 := make(map[int]bool)
    v3 := []uint32{4294967295, 399, 2}
    var v4 float32 = (0.1 - 1e10)
    v5 := make(map[uint32]int)
    fmt.Println(v4)
    var v34 uint = 0
    var v35 int32 = 720
    if (!false) {
        for fi36 := 0; fi36 < 14; fi36++ {
            v37 := GenS1{F0: fi36, F1: 545, F2: (("" + "foo bar") + "hello")}
            var v38 int64 = -321
            v35 += -986
            _ = fi36
            _ = v37
            _ = v38
        }
    } else {
        v1 -= (v1 & v1)
    }
    v39 := GenS3{F0: 289, F1: ("" + ""), F2: 142}
    v35 = (v35 & v35)
    fmt.Println(v4, v4, v35)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v34 = (v34 & 2)
    if (!false) {
        for fi40 := 0; fi40 < 9; fi40++ {
            for fi41 := 0; fi41 < 3; fi41++ {
                _ = len(v3)
                _ = fi41
            }
            _ = fi40
        }
        for fi42 := 0; fi42 < 16; fi42++ {
            if false {
                v35 = -2147483648
                var v43 int16 = -187
                _ = v43
            }
            var v44 uint16 = 255
            switch "hello" {
            case "foo bar":
                var v45 int = genFn1(v1, 2147483647, v44)
                _ = v45
            case "a":
                var v46 int32 = v35
                _ = v46
            case "hello":
                delete(v5, 65535)
            default:
                if (0.0 != (v4 * v4)) {
                    if (127 <= 198) {
                        if (127 > 0) {
                            if ((("hello" + "\n\t") + ("foo bar" + "\n\t")) > "foo bar") {
                                var v47 int8 = (v1 ^ v1)
                                if ((42 == 42) || true) {
                                    if len(v3) > 0 {
                                        v3[0] = 2
                                    }
                                    if false {
                                        fmt.Println(v35, v35, v44)
                                        var v48 float64 = (0.1 * 1.5)
                                        v47 += (v1 % (((-(v1))) | 1))
                                        _ = v48
                                    }
                                    var v49 int16 = -1
                                    _ = v49
                                } else {
                                    fmt.Println(v1, v34, v4)
                                }
                                v47 = (v47 % (((v47 % (((v1 - (v1 | 1))) | 1))) | 1))
                                _ = v47
                            } else {
                                if false {
                                    if (genFn0(344, -670) < 2147483648) {
                                        for fi50 := 0; fi50 < 15; fi50++ {
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v51 int32 = (^(v35))
                                                _ = v51
                                            default:
                                                if true {
                                                    var v52 uint = v34
                                                    v53 := v5[2]
                                                    v54 := GenS3{F0: v34, F1: "", F2: (-(v35))}
                                                    delete(v5, 2147483648)
                                                    _ = v52
                                                    _ = v53
                                                    _ = v54
                                                }
                                            }
                                            fmt.Println(fi42, v34, fi42)
                                            _ = fi50
                                        }
                                        v55 := GenS4{F0: (v34 & v34), F1: 1000, F2: 2147483648}
                                        _ = v55
                                    }
                                } else {
                                    v56 := GenS5{F0: 37, F1: 1}
                                    if ("\n\t" <= "") {
                                        _ = len(v5)
                                        for fi57 := 0; fi57 < 10; fi57++ {
                                            v4 -= float32(((float64(v4) - 1e10) - -2.25))
                                            v58 := GenS2{F0: (-(v35)), F1: -128}
                                            v59 := v2[int(v35)]
                                            _ = fi57
                                            _ = v58
                                            _ = v59
                                        }
                                        v60 := GenS3{F0: (v34 << (((v34 - 683)) & 63)), F1: "a", F2: (v35 >> ((v34) & 63))}
                                        _ = v60
                                    } else {
                                        if true {
                                            var v61 bool = false
                                            var v62 uint16 = 32767
                                            _ = v61
                                            _ = v62
                                        } else {
                                            fmt.Println(v44, v34, fi42)
                                        }
                                        fmt.Println(v44)
                                    }
                                    for fi63 := 0; fi63 < 13; fi63++ {
                                        v64 := GenS1{F0: (-(fi63)), F1: 9223372036854775807, F2: (("foo bar" + ("hello" + ((("\n\t" + "a") + ("hello" + ("\n\t" + ("a" + "a")))) + ("\n\t" + "\n\t")))) + "foo bar")}
                                        _ = fi63
                                        _ = v64
                                    }
                                    _ = v56
                                }
                            }
                            v65 := GenS6{F0: 1, F1: 2147483648, F2: (v1 * (v1 % ((v1) | 1)))}
                            _ = v65
                        }
                        v66 := GenS2{F0: (v35 | v35), F1: v1}
                        _ = v66
                    }
                    var v67 int8 = (v1 * (v1 % (((v1 << ((v34) & 63))) | 1)))
                    v68 := GenS7{F0: -32768}
                    _ = v67
                    _ = v68
                } else {
                    var v69 uint16 = 457
                    _ = v69
                }
            }
            _ = fi42
            _ = v44
        }
    }
    for fi70 := 0; fi70 < 19; fi70++ {
        for fi71 := 0; fi71 < 5; fi71++ {
            v72 := GenS8{F0: v34, F1: 2, F2: (v34 - v34)}
            var v73 int16 = 42
            v73 += (v73 + 1)
            v1 = (v1 ^ v1)
            _ = fi71
            _ = v72
            _ = v73
        }
        _ = fi70
    }
    v74 := GenS6{F0: 1, F1: 2, F2: v1}
    if true {
        var v75 int = 128
        _ = v75
    } else {
        v35 = (v35 + v35)
        var v76 float32 = (1e10 / 1.0)
        _ = v76
    }
    var v77 uint32 = 249
    switch ("\n\t" + "") {
    case "hello":
        switch ("foo bar" + "") {
        case "\n\t":
            fmt.Println(v1, v4, v34)
        case "hello":
            for fi78 := 0; fi78 < 13; fi78++ {
                switch "foo bar" {
                case "foo bar":
                    switch "a" {
                    case "foo bar":
                        if (-962 == -1) {
                            switch fi78 {
                            case 680:
                                if true {
                                    switch fi78 {
                                    case 447:
                                        v1 -= (v1 << ((0) & 63))
                                    default:
                                        v34 = (v34 << (((v34 >> ((v34) & 63))) & 63))
                                    }
                                    switch "" {
                                    case "\n\t":
                                        for fi79 := 0; fi79 < 7; fi79++ {
                                            var v80 uint = v34
                                            _ = len(v2)
                                            _ = fi79
                                            _ = v80
                                        }
                                    case "a":
                                        v4 = -2.25
                                    case "":
                                        fmt.Println(v34, v1)
                                    default:
                                        v77 = 2
                                    }
                                } else {
                                    if false {
                                        var v81 uint16 = 42
                                        if (((v81 & 468) != 255) && (false || true)) {
                                            var v82 uint8 = 2
                                            if (v35 <= (^(v35))) {
                                                for fi83 := 0; fi83 < 16; fi83++ {
                                                    var v84 uint32 = (v77 ^ 867)
                                                    var v85 uint8 = 128
                                                    if false {
                                                        if false {
                                                            v85 -= (v82 - (v85 >> ((v34) & 63)))
                                                            v86 := GenS2{F0: (v35 / ((2147483647) | 1)), F1: (v1 << (((v34 << (((^(v34))) & 63))) & 63))}
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v86
                                                        } else {
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                v3 = append(v3, v77)
                                                            case "":
                                                                var v87 uint32 = v77
                                                                _ = v87
                                                            }
                                                        }
                                                        if true {
                                                            v88 := GenS9{F0: v35, F1: 840}
                                                            fmt.Println(v81, v84, v1)
                                                            _ = v88
                                                        } else {
                                                            v81 = 861
                                                        }
                                                        if (true && (!(!false))) {
                                                            v89 := GenS10{F0: int64(v1), F1: (v34 * (v34 * (v34 & v34))), F2: (false && true)}
                                                            v81 -= 597
                                                            _ = v89
                                                        } else {
                                                            if false {
                                                                var v90 float32 = ((((v4 + (1e10 - v4)) * (0.1 + v4)) - v4) + -2.25)
                                                                var v91 float32 = 3.14159
                                                                _ = v90
                                                                _ = v91
                                                            }
                                                        }
                                                    }
                                                    _ = fi83
                                                    _ = v84
                                                    _ = v85
                                                }
                                                v92 := GenS5{F0: 127, F1: (v82 + v82)}
                                                if true {
                                                    switch "a" {
                                                    case "hello":
                                                        v77 = (v77 + (v77 - v77))
                                                    case "":
                                                        fmt.Println(fi78)
                                                    case "a":
                                                        for fi93 := 0; fi93 < 4; fi93++ {
                                                            v2[fi93] = ("hello" == "a")
                                                            switch "foo bar" {
                                                            case "":
                                                                var v94 int16 = 42
                                                                _ = v94
                                                            }
                                                            v95 := GenS11{F0: (v34 % ((uint(v77)) | 1)), F1: (fi93 % ((int(v77)) | 1)), F2: -32768}
                                                            _ = fi93
                                                            _ = v95
                                                        }
                                                    }
                                                    v35 -= int32(v77)
                                                    v1 = v1
                                                } else {
                                                    if ((!(((3.14159 == (((float64(v4) * 0.0) + 3.14159) / float64(v4))) && false) || false)) || false) {
                                                        for fi96 := 0; fi96 < 15; fi96++ {
                                                            if true {
                                                                var v97 int = (^(fi78))
                                                                _ = v97
                                                            } else {
                                                                var v98 int32 = (v35 + int32(v82))
                                                                if true {
                                                                    var v99 uint = 0
                                                                    fmt.Println(v77, v98, v1)
                                                                    _ = len(v3)
                                                                    _ = v99
                                                                }
                                                                fmt.Println(v4, v98, v1)
                                                                _ = v98
                                                            }
                                                            _ = fi96
                                                        }
                                                    } else {
                                                        _ = len(v3)
                                                        v1 -= v1
                                                    }
                                                    v35 = (-(v35))
                                                }
                                                _ = v92
                                            } else {
                                                var v100 uint32 = 42
                                                if true {
                                                    fmt.Println(v100, v82)
                                                    var v101 uint8 = v82
                                                    _ = v101
                                                }
                                                _ = v100
                                            }
                                            _ = v82
                                        } else {
                                            v102 := GenS1{F0: fi78, F1: 42, F2: "\n\t"}
                                            v103 := GenS12{F0: uint64(v81), F1: (v35 - (v35 << (((^(v34))) & 63)))}
                                            v104 := GenS7{F0: 32767}
                                            _ = v102
                                            _ = v103
                                            _ = v104
                                        }
                                        v105 := GenS13{F0: 848}
                                        _ = v105
                                        _ = v81
                                    }
                                    for fi106 := 0; fi106 < 14; fi106++ {
                                        if true {
                                            fmt.Println(v4, v34)
                                            v77 -= (v77 + v77)
                                            if true {
                                                var v107 int32 = -1
                                                _ = v107
                                            } else {
                                                switch (fi106 | -840) {
                                                case 42:
                                                    v2[fi78] = (870 < 49)
                                                default:
                                                    if (!false) {
                                                        if true {
                                                            v34 = (v34 & v34)
                                                        } else {
                                                            var v108 float32 = v4
                                                            v4 -= v4
                                                            _ = v108
                                                        }
                                                        fmt.Println(fi78, v35)
                                                    } else {
                                                        v109 := GenS14{F0: (-(v35)), F1: -1, F2: 199}
                                                        var v110 uint32 = 4294967295
                                                        var v111 uint64 = 449
                                                        _ = v109
                                                        _ = v110
                                                        _ = v111
                                                    }
                                                }
                                            }
                                            switch ("a" + "") {
                                            case "hello":
                                                switch fi106 {
                                                case 927:
                                                    _ = len(v3)
                                                default:
                                                    v35 = -2147483648
                                                }
                                            case "a":
                                                var v112 float32 = (v4 + float32(float64((3.14159 / 1.5))))
                                                _ = v112
                                            default:
                                                v113 := GenS11{F0: 4294967295, F1: fi106, F2: -596}
                                                _ = v113
                                            }
                                        } else {
                                            switch (^(fi106)) {
                                            case 32767:
                                                v35 = v35
                                            case 128:
                                                v114 := GenS15{F0: 9223372036854775807}
                                                _ = v114
                                            case -87:
                                                v3 = append(v3, v77)
                                            default:
                                                if false {
                                                    _ = len(v3)
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v115 uint32 = v77
                                                        _ = v115
                                                    case "":
                                                        if (0.0 <= (float64((1.5 * v4)) / (1.5 * (((float64(-2.25) * 1.5) * 3.14159) / 1e10)))) {
                                                            v3 = append(v3, (v77 / (((^(v77))) | 1)))
                                                            v1 += (v1 / ((v1) | 1))
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            v116 := GenS12{F0: 299, F1: (v35 >> ((4294967295) & 63))}
                                                            if true {
                                                                if true {
                                                                    var v117 string = ("\n\t" + "a")
                                                                    v35 = 0
                                                                    _ = v117
                                                                } else {
                                                                    _ = len(v3)
                                                                    v118 := GenS2{F0: (v35 ^ v35), F1: (v1 << ((v34) & 63))}
                                                                    _ = v118
                                                                }
                                                                switch fi106 {
                                                                case -663:
                                                                    v35 = 1
                                                                default:
                                                                    if (180 != 128) {
                                                                        var v119 string = "hello"
                                                                        switch (-(fi106)) {
                                                                        case 9223372036854775807:
                                                                            v120 := GenS16{F0: v34, F1: (v35 | (v35 ^ (v35 % ((-325) | 1)))), F2: v1}
                                                                            if (131 > 42) {
                                                                                var v121 uint8 = 128
                                                                                if true {
                                                                                    if false {
                                                                                        v121 = v121
                                                                                    }
                                                                                    if (!false) {
                                                                                        switch (fi106 % ((fi78) | 1)) {
                                                                                        case 0:
                                                                                            if (fi78 <= genFn1(42, fi78, 219)) {
                                                                                                var v122 uint8 = 128
                                                                                                if false {
                                                                                                    switch ((v119 + v119) + (v119 + "")) {
                                                                                                    case "a":
                                                                                                        fmt.Println(v4, v4)
                                                                                                    case "foo bar":
                                                                                                        fmt.Println(v4)
                                                                                                        v5[(v77 | (v77 << ((v34) & 63)))] = fi106
                                                                                                        if false {
                                                                                                            fmt.Println(v121, v1)
                                                                                                            fmt.Println(v122, v35, v77)
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        if (true && (0.1 <= 3.14159)) {
                                                                                                            v123 := GenS17{F0: -2147483648, F1: (v122 & (v122 / ((42) | 1))), F2: -73}
                                                                                                            if (255 <= v121) {
                                                                                                                var v124 int = (fi106 ^ int(v1))
                                                                                                                if len(v3) >= 2 {
                                                                                                                    v3 = v3[1:len(v3)]
                                                                                                                }
                                                                                                                switch v119 {
                                                                                                                case "foo bar":
                                                                                                                    v125 := GenS12{F0: 2, F1: -839}
                                                                                                                    _ = v125
                                                                                                                case "hello":
                                                                                                                    if len(v3) > 0 {
                                                                                                                        v3[0] = (v77 * 0)
                                                                                                                    }
                                                                                                                    switch ("\n\t" + ("hello" + v119)) {
                                                                                                                    case "foo bar":
                                                                                                                        if true {
                                                                                                                            fmt.Println(v34, fi78)
                                                                                                                            var v126 int16 = 0
                                                                                                                            if (-2.25 >= v4) {
                                                                                                                                fmt.Println(fi106)
                                                                                                                                _ = len(v5)
                                                                                                                                v126 = v126
                                                                                                                            }
                                                                                                                            _ = v126
                                                                                                                        }
                                                                                                                    case "hello":
                                                                                                                        if (v121 == v122) {
                                                                                                                            v127 := GenS13{F0: 42}
                                                                                                                            v35 = (v35 << ((2) & 63))
                                                                                                                            _ = v127
                                                                                                                        }
                                                                                                                    }
                                                                                                                    fmt.Println(v34, v119, v119)
                                                                                                                default:
                                                                                                                    v128 := GenS18{F0: 860, F1: (fi78 >> ((990) & 63))}
                                                                                                                    _ = v128
                                                                                                                }
                                                                                                                _ = v124
                                                                                                            }
                                                                                                            _ = v123
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    v121 += (v122 - v121)
                                                                                                }
                                                                                                v119 = v119
                                                                                                _ = v122
                                                                                            }
                                                                                        case 287:
                                                                                            var v129 uint8 = 255
                                                                                            _ = v129
                                                                                        case 1:
                                                                                            v130 := GenS19{F0: 32768}
                                                                                            _ = v130
                                                                                        default:
                                                                                            switch "hello" {
                                                                                            case "a":
                                                                                                v131 := GenS20{F0: 79, F1: (^(v77)), F2: false}
                                                                                                _ = v131
                                                                                            default:
                                                                                                _ = len(v5)
                                                                                            }
                                                                                        }
                                                                                        var v132 uint64 = 1
                                                                                        _ = v132
                                                                                    } else {
                                                                                        v4 = float32(1.5)
                                                                                        if true {
                                                                                            v133 := GenS21{F0: 525, F1: (-(v35)), F2: v34}
                                                                                            var v134 float32 = v4
                                                                                            var v135 bool = false
                                                                                            _ = v133
                                                                                            _ = v134
                                                                                            _ = v135
                                                                                        } else {
                                                                                            v34 += v34
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            var v136 int64 = 42
                                                                                            var v137 uint64 = 1
                                                                                            _ = v136
                                                                                            _ = v137
                                                                                        }
                                                                                    }
                                                                                    v138 := GenS22{F0: 42}
                                                                                    _ = v138
                                                                                }
                                                                                if true {
                                                                                    if (4294967295 != 18446744073709551615) {
                                                                                        v139 := GenS22{F0: 1}
                                                                                        _ = v139
                                                                                    }
                                                                                    v119 = "foo bar"
                                                                                    if (!(0 < 361)) {
                                                                                        v140 := v2[(fi106 * -2147483648)]
                                                                                        _ = v140
                                                                                    }
                                                                                }
                                                                                if (((0 >= 901) && true) || false) {
                                                                                    fmt.Println(v121, v35)
                                                                                }
                                                                                if (65 == 530) {
                                                                                    fmt.Println(fi78)
                                                                                    v141 := GenS9{F0: -2147483648, F1: 4294967295}
                                                                                    var v142 int8 = v1
                                                                                    _ = v141
                                                                                    _ = v142
                                                                                }
                                                                                v77 = v77
                                                                                if true {
                                                                                    var v143 float32 = v4
                                                                                    v144 := GenS23{F0: uint64(v121), F1: (true || false), F2: 400}
                                                                                    v145 := GenS14{F0: v35, F1: (v35 << (((v34 & 4294967295)) & 63)), F2: v121}
                                                                                    _ = v143
                                                                                    _ = v144
                                                                                    _ = v145
                                                                                }
                                                                                _ = v121
                                                                            } else {
                                                                                v119 = v119
                                                                                if false {
                                                                                    v146 := GenS12{F0: 18446744073709551615, F1: v35}
                                                                                    _ = v146
                                                                                } else {
                                                                                    fmt.Println(v4, v35, v77)
                                                                                    if true {
                                                                                        var v147 float64 = 1.5
                                                                                        _ = v147
                                                                                    }
                                                                                }
                                                                                var v148 uint = v34
                                                                                switch "" {
                                                                                case "foo bar":
                                                                                    fmt.Println(v1, v77, v119)
                                                                                case "":
                                                                                    v149 := GenS10{F0: 9223372036854775807, F1: (v34 / (((v34 ^ v34)) | 1)), F2: (!false)}
                                                                                    _ = v149
                                                                                }
                                                                                v150 := GenS14{F0: (v35 >> ((v148) & 63)), F1: (^(v35)), F2: 127}
                                                                                _ = v148
                                                                                _ = v150
                                                                            }
                                                                            _ = len(v3)
                                                                            _ = v120
                                                                        default:
                                                                            v34 = v34
                                                                        }
                                                                        _ = v119
                                                                    } else {
                                                                        v35 = v35
                                                                        v151 := GenS7{F0: 1000}
                                                                        fmt.Println(fi78)
                                                                        if (v4 <= v4) {
                                                                            var v152 uint64 = 42
                                                                            var v153 string = ""
                                                                            if true {
                                                                                v154 := GenS24{F0: 705, F1: 849}
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = (v77 ^ (^(v77)))
                                                                                }
                                                                                v155 := v5[(v77 * (v77 - (v77 + v77)))]
                                                                                v2[fi78] = (-1 <= -394)
                                                                                delete(v2, int(v35))
                                                                                _ = v154
                                                                                _ = v155
                                                                            } else {
                                                                                v156 := GenS5{F0: 224, F1: 128}
                                                                                _ = v156
                                                                            }
                                                                            _ = v152
                                                                            _ = v153
                                                                        }
                                                                        v157 := GenS25{F0: 127}
                                                                        _ = v151
                                                                        _ = v157
                                                                    }
                                                                }
                                                                if len(v3) > 0 {
                                                                    v3[0] = v77
                                                                }
                                                                v158 := GenS7{F0: int16(v77)}
                                                                v159 := GenS20{F0: 824, F1: (v77 + v77), F2: true}
                                                                var v160 int16 = 1000
                                                                v161 := v2[fi106]
                                                                var v162 float64 = (3.14159 + (float64(v4) * (float64(v4) * (float64(((v4 * float32(-2.25)) / v4)) * 1.5))))
                                                                if v161 {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    var v163 uint32 = (v77 - (^(v77)))
                                                                    if (true && v161) {
                                                                        var v164 int32 = (v35 & v35)
                                                                        var v165 bool = false
                                                                        fmt.Println(v164)
                                                                        v35 = (v164 % ((v35) | 1))
                                                                        _ = v164
                                                                        _ = v165
                                                                    }
                                                                    _ = v163
                                                                } else {
                                                                    v166 := GenS3{F0: (v34 >> (((v34 << ((v34) & 63))) & 63)), F1: (("a" + "foo bar") + ("\n\t" + "\n\t")), F2: 2147483647}
                                                                    switch (fi78 << ((2) & 63)) {
                                                                    case 9223372036854775807:
                                                                        var v167 float32 = v4
                                                                        _ = v167
                                                                    }
                                                                    fmt.Println(v4, fi78, v77)
                                                                    v168 := GenS26{F0: "", F1: int(v35)}
                                                                    _ = v166
                                                                    _ = v168
                                                                }
                                                                _ = v158
                                                                _ = v159
                                                                _ = v160
                                                                _ = v161
                                                                _ = v162
                                                            }
                                                            if (v35 < v35) {
                                                                v34 += (^(v34))
                                                                _ = len(v5)
                                                                if ("hello" <= "a") {
                                                                    v2[(fi106 & (-(fi78)))] = (!(-32768 >= -356))
                                                                    switch ("\n\t" + (("a" + "foo bar") + ("a" + ("\n\t" + "foo bar")))) {
                                                                    case "\n\t":
                                                                        v169 := GenS10{F0: 1, F1: v34, F2: (4294967295 <= v77)}
                                                                        _ = v169
                                                                    default:
                                                                        v3 = append(v3, (v77 % ((genFn0(42, -1)) | 1)))
                                                                    }
                                                                    switch "a" {
                                                                    case "":
                                                                        v77 = (v77 / ((v77) | 1))
                                                                    case "\n\t":
                                                                        v35 -= v35
                                                                    case "a":
                                                                        var v170 int = genFn1(v1, 559, 42)
                                                                        _ = v170
                                                                    default:
                                                                        v171 := GenS27{F0: 1, F1: 84, F2: int(v35)}
                                                                        _ = v171
                                                                    }
                                                                    v1 += 0
                                                                } else {
                                                                    if (true && false) {
                                                                        v34 -= 979
                                                                        fmt.Println(v4, v34, v4)
                                                                        var v172 uint16 = 2
                                                                        v77 = v77
                                                                        if (((v34 - (v34 | (^(v34)))) > v34) || (v1 <= -128)) {
                                                                            if len(v3) > 0 {
                                                                                v3[0] = (v77 ^ genFn0(1, -743))
                                                                            }
                                                                            var v173 int = -509
                                                                            _ = v173
                                                                        }
                                                                        _ = v172
                                                                    }
                                                                    v77 = (v77 >> ((v34) & 63))
                                                                    v1 = (v1 | 100)
                                                                }
                                                                v77 = 2
                                                            } else {
                                                                fmt.Println(v35)
                                                                var v174 int16 = -398
                                                                var v175 uint8 = 0
                                                                _ = v174
                                                                _ = v175
                                                            }
                                                            _ = v116
                                                        } else {
                                                            v176 := GenS28{F0: 748, F1: "hello", F2: false}
                                                            if (fi78 < -2147483648) {
                                                                v77 = (^(v77))
                                                            }
                                                            var v177 int16 = 348
                                                            _ = v176
                                                            _ = v177
                                                        }
                                                    default:
                                                        if len(v3) > 0 {
                                                            v3[0] = 65535
                                                        }
                                                    }
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        var v178 uint64 = 281
                                                        _ = v178
                                                    case "":
                                                        v4 = 1e10
                                                    default:
                                                        switch "a" {
                                                        case "a":
                                                            var v179 string = "a"
                                                            _ = v179
                                                        case "":
                                                            if ((^(fi106)) != (fi78 - (fi78 ^ (fi106 + (fi106 ^ fi106))))) {
                                                                delete(v2, (fi106 * 42))
                                                                var v180 float32 = (v4 + 1e10)
                                                                fmt.Println(fi106, v1)
                                                                _ = v180
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                        _ = fi106
                                    }
                                }
                            }
                        }
                    }
                case "a":
                    var v181 bool = true
                    _ = v181
                case "\n\t":
                    fmt.Println(v34, v77, v4)
                }
                fmt.Println(v77)
                _ = fi78
            }
        case "foo bar":
            for fi182 := 0; fi182 < 7; fi182++ {
                if (!((!true) || ((v34 + 65535) != v34))) {
                    v4 = (3.14159 * v4)
                    v3 = append(v3, (v77 - genFn0(1, 797)))
                    v5[229] = int(v77)
                }
                _ = fi182
            }
        }
    case "a":
        v4 = v4
    case "foo bar":
        switch -1 {
        case 42:
            v34 = (v34 | (v34 - uint(v77)))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v5
    _ = v74
    _ = v77
}
