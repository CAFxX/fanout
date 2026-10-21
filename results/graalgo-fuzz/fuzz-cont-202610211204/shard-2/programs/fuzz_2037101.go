package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int }
type GenS2 struct { F0 uint8; F1 int16; F2 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool; F1 int16 }
type GenS8 struct { F0 bool; F1 int }
type GenS9 struct { F0 int64; F1 uint16 }
type GenS10 struct { F0 uint; F1 bool }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint; F1 uint; F2 uint64 }
type GenS13 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS14 struct { F0 uint8; F1 int; F2 int }
type GenS15 struct { F0 int; F1 uint64 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 uint8; F1 uint16 }
type GenS19 struct { F0 int; F1 string }
type GenS20 struct { F0 int16; F1 string }
type GenS21 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS22 struct { F0 int8 }
type GenS23 struct { F0 int; F1 uint8 }
type GenS24 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS25 struct { F0 int64 }
type GenS26 struct { F0 uint16; F1 string }
type GenS27 struct { F0 int; F1 uint64 }
type GenS28 struct { F0 int32 }
func genFn0() int16 {
    var v4 int32 = 1
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := []int{(fi5 | (fi5 * (fi5 | int(v4)))), (fi5 | fi5)}
        var v7 uint32 = 2147483648
        _ = fi5
        _ = v6
        _ = v7
    }
    v4 += v4
    _ = v4
    return 42
}
func genFn1() bool {
    if ((0.0 >= -2.25) && false) {
        var v8 string = ""
        _ = v8
    } else {
        switch "\n\t" {
        case "a":
            var v9 uint16 = 0
            _ = v9
        case "\n\t":
            v10 := GenS1{F0: 32767, F1: 42, F2: -1}
            _ = v10
        case "":
            var v11 uint16 = 2
            _ = v11
        default:
            if (((0.0 - ((1e10 / -2.25) + 0.1)) - 0.0) > -2.25) {
                v12 := GenS2{F0: 140, F1: 1, F2: 402}
                v13 := []int{}
                if true {
                    if false {
                        if false {
                            var v14 uint64 = 60
                            for fi15 := 0; fi15 < 4; fi15++ {
                                v14 = (v14 - (v14 ^ (v14 / ((v14) | 1))))
                                var v16 float32 = 1e10
                                _ = fi15
                                _ = v16
                            }
                            _ = v14
                        } else {
                            var v17 int = 9223372036854775807
                            _ = v17
                        }
                        var v18 string = ("foo bar" + "")
                        _ = v18
                    }
                    v19 := GenS2{F0: 0, F1: -1, F2: 0}
                    var v20 uint8 = 7
                    _ = v19
                    _ = v20
                }
                _ = v12
                _ = v13
            }
        }
    }
    v21 := GenS3{F0: -293}
    var v22 int64 = -1
    switch "a" {
    case "hello":
        var v23 float32 = (1e10 * 3.14159)
        _ = v23
    case "":
        v24 := []int{315}
        _ = v24
    default:
        for fi25 := 0; fi25 < 17; fi25++ {
            switch ("hello" + ("hello" + "")) {
            case "hello":
                switch "" {
                case "foo bar":
                    var v26 float64 = 0.1
                    _ = v26
                case "a":
                    fmt.Println(fi25)
                case "":
                    var v27 int = fi25
                    _ = v27
                default:
                    for fi28 := 0; fi28 < 9; fi28++ {
                        var v29 int = fi25
                        _ = fi28
                        _ = v29
                    }
                }
            case "a":
                v22 -= v22
            }
            _ = fi25
        }
    }
    var v30 int16 = 0
    _ = v21
    _ = v22
    _ = v30
    return true
}

func main() {
    var v1 uint8 = 176
    var v2 int32 = 1
    var v3 uint64 = 2
    if true {
        if (true && true) {
            for fi31 := 0; fi31 < 6; fi31++ {
                v32 := GenS1{F0: 255, F1: v3, F2: (fi31 - (fi31 >> ((213) & 63)))}
                if (true || false) {
                    v33 := GenS2{F0: v1, F1: -32768, F2: 1}
                    if (true || false) {
                        var v34 int32 = v2
                        for fi35 := 0; fi35 < 18; fi35++ {
                            v36 := []int{}
                            var v37 uint = 0
                            var v38 uint8 = v1
                            if ((2 == 4294967295) || false) {
                                var v39 int32 = v2
                                _ = v39
                            }
                            var v40 uint32 = 0
                            _ = fi35
                            _ = v36
                            _ = v37
                            _ = v38
                            _ = v40
                        }
                        var v41 int32 = (-(v34))
                        _ = v34
                        _ = v41
                    } else {
                        if true {
                            v1 += (v1 | (v1 % ((v1) | 1)))
                            var v42 uint16 = 973
                            v2 -= v2
                            _ = v42
                        }
                        switch fi31 {
                        case 9223372036854775807:
                            v3 = (^(v3))
                        case -2147483648:
                            if false {
                                v43 := []int{}
                                _ = v43
                            } else {
                                var v44 float32 = float32((-2.25 / 1.5))
                                v45 := []int{fi31, (fi31 & fi31), int(v3)}
                                for fi46 := 0; fi46 < 16; fi46++ {
                                    v3 += (v3 ^ (^(v3)))
                                    _ = fi46
                                }
                                _ = v44
                                _ = v45
                            }
                        case 4294967295:
                            v2 = v2
                        }
                        var v47 int = (fi31 ^ 1)
                        _ = v47
                    }
                    for fi48 := 0; fi48 < 4; fi48++ {
                        var v49 int16 = -103
                        v50 := []int{}
                        fmt.Println(v1, fi31, fi31)
                        if (0.0 <= -2.25) {
                            v3 = v3
                            var v51 bool = false
                            switch "hello" {
                            case "":
                                _ = len(v50)
                            default:
                                v3 += (v3 ^ v3)
                            }
                            _ = v51
                        }
                        switch (("a" + "foo bar") + "a") {
                        case "a":
                            fmt.Println(v3)
                        case "":
                            v50 = append(v50, (fi48 + (fi31 | (fi48 >> ((42) & 63)))))
                            v3 -= 467
                            v2 = (v2 ^ (v2 / (((v2 ^ (^(v2)))) | 1)))
                        case "foo bar":
                            v49 -= (v49 & v49)
                        default:
                            if true {
                                if (4294967295 != 4294967295) {
                                    v52 := GenS2{F0: v1, F1: -32768, F2: 18446744073709551615}
                                    if len(v50) > 0 {
                                        v50[0] = fi31
                                    }
                                    _ = v52
                                } else {
                                    switch fi31 {
                                    case 0:
                                        if (!(32768 != fi48)) {
                                            _ = len(v50)
                                        } else {
                                            switch "" {
                                            case "\n\t":
                                                var v53 uint = 272
                                                _ = v53
                                            }
                                            switch "\n\t" {
                                            case "":
                                                v49 -= (-(v49))
                                            case "a":
                                                var v54 string = ("hello" + "a")
                                                _ = v54
                                            }
                                        }
                                    case 128:
                                        v55 := GenS4{F0: 18446744073709551615, F1: 65535, F2: 241}
                                        _ = v55
                                    case 1:
                                        switch ("foo bar" + (("a" + ("" + "hello")) + "")) {
                                        case "":
                                            switch "hello" {
                                            case "hello":
                                                fmt.Println(v1, v2)
                                            default:
                                                var v56 bool = (!true)
                                                _ = v56
                                            }
                                        case "foo bar":
                                            fmt.Println(v2, v49, fi31)
                                        default:
                                            v1 -= (v1 & 255)
                                        }
                                    default:
                                        var v57 uint8 = 128
                                        _ = v57
                                    }
                                    if false {
                                        switch "foo bar" {
                                        case "\n\t":
                                            if false {
                                                var v58 uint16 = 42
                                                var v59 uint64 = (v3 + (v3 >> ((18446744073709551615) & 63)))
                                                switch "foo bar" {
                                                case "hello":
                                                    v60 := GenS3{F0: 635}
                                                    _ = v60
                                                case "":
                                                    v49 += (v49 & v49)
                                                    var v61 uint16 = 0
                                                    _ = v61
                                                }
                                                _ = v58
                                                _ = v59
                                            }
                                        default:
                                            switch (fi31 / ((128) | 1)) {
                                            case 32768:
                                                var v62 uint16 = 65535
                                                _ = v62
                                            }
                                            var v63 float32 = (-2.25 * 3.14159)
                                            if len(v50) > 0 {
                                                v50[0] = (fi48 | int(v3))
                                            }
                                            _ = v63
                                        }
                                    } else {
                                        switch (((("" + "") + ("" + "\n\t")) + "\n\t") + "\n\t") {
                                        case "":
                                            v50 = append(v50, 2147483647)
                                        case "foo bar":
                                            v50 = append(v50, (^(fi31)))
                                            var v64 int64 = -696
                                            if true {
                                                switch fi48 {
                                                case 32767:
                                                    v65 := GenS5{F0: true, F1: v64}
                                                    _ = v65
                                                case -2147483648:
                                                    v1 += 102
                                                case 32768:
                                                    switch "hello" {
                                                    case "\n\t":
                                                        if (v49 <= v49) {
                                                            var v66 uint = 0
                                                            _ = v66
                                                        }
                                                    case "hello":
                                                        switch ("foo bar" + "hello") {
                                                        case "\n\t":
                                                            if len(v50) >= 2 {
                                                                v50 = v50[1:len(v50)]
                                                            }
                                                        case "":
                                                            if false {
                                                                v3 += v3
                                                                var v67 uint = 166
                                                                if true {
                                                                    if (v2 >= v2) {
                                                                        fmt.Println(v67, fi31)
                                                                        _ = len(v50)
                                                                    } else {
                                                                        if false {
                                                                            v68 := GenS3{F0: -285}
                                                                            var v69 uint16 = uint16(v67)
                                                                            v49 += v49
                                                                            v49 += (v49 - genFn0())
                                                                            if true {
                                                                                v50 = append(v50, (fi31 - 127))
                                                                                fmt.Println(v2, fi48, v69)
                                                                                v67 = v67
                                                                            } else {
                                                                                v70 := GenS3{F0: v64}
                                                                                _ = len(v50)
                                                                                if false {
                                                                                    v64 = v64
                                                                                    v1 = v1
                                                                                } else {
                                                                                    if false {
                                                                                        var v71 string = "hello"
                                                                                        _ = v71
                                                                                    }
                                                                                    v64 -= -996
                                                                                    if (3.14159 > 3.14159) {
                                                                                        if len(v50) > 0 {
                                                                                            v50[0] = fi31
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v70
                                                                            }
                                                                            _ = v68
                                                                            _ = v69
                                                                        } else {
                                                                            var v72 float64 = 3.14159
                                                                            var v73 uint8 = v1
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                v50 = append(v50, int(v64))
                                                                            case "foo bar":
                                                                                fmt.Println(v73, v72)
                                                                            case "":
                                                                                if true {
                                                                                    switch fi48 {
                                                                                    case -2147483648:
                                                                                        v74 := GenS6{F0: v73}
                                                                                        _ = v74
                                                                                    case 417:
                                                                                        var v75 float64 = v72
                                                                                        var v76 string = ("foo bar" + "foo bar")
                                                                                        fmt.Println(fi48, fi48, v1)
                                                                                        switch v76 {
                                                                                        case "\n\t":
                                                                                            var v77 int8 = -127
                                                                                            _ = v77
                                                                                        default:
                                                                                            v78 := GenS5{F0: ((fi48 >> (((^(v67))) & 63)) < (fi48 % (((-(fi31))) | 1))), F1: -1}
                                                                                            _ = v78
                                                                                        }
                                                                                        v79 := GenS7{F0: false, F1: genFn0()}
                                                                                        var v80 string = ("foo bar" + v76)
                                                                                        _ = v75
                                                                                        _ = v76
                                                                                        _ = v79
                                                                                        _ = v80
                                                                                    default:
                                                                                        fmt.Println(v72)
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(fi48)
                                                                                    switch fi31 {
                                                                                    case 32768:
                                                                                        if true {
                                                                                            if len(v50) >= 2 {
                                                                                                v50 = v50[1:len(v50)]
                                                                                            }
                                                                                        } else {
                                                                                            fmt.Println(v73)
                                                                                        }
                                                                                    case -11:
                                                                                        if false {
                                                                                            if false {
                                                                                                if true {
                                                                                                    fmt.Println(v3)
                                                                                                    if (((true && false) && (!true)) && (!true)) {
                                                                                                        v81 := GenS8{F0: (1e10 != 3.14159), F1: fi31}
                                                                                                        _ = v81
                                                                                                    } else {
                                                                                                        v82 := GenS9{F0: (v64 ^ v64), F1: 843}
                                                                                                        fmt.Println(v72, v1, v64)
                                                                                                        var v83 uint8 = 255
                                                                                                        v84 := GenS5{F0: (!false), F1: 1}
                                                                                                        switch 128 {
                                                                                                        case -9223372036854775808:
                                                                                                            var v85 int16 = (v49 & (v49 % ((int16(v3)) | 1)))
                                                                                                            _ = v85
                                                                                                        case 490:
                                                                                                            v86 := GenS4{F0: uint(v3), F1: 1, F2: (v83 ^ v83)}
                                                                                                            _ = v86
                                                                                                        }
                                                                                                        _ = v82
                                                                                                        _ = v83
                                                                                                        _ = v84
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            v73 = 127
                                                                                            _ = len(v50)
                                                                                        }
                                                                                    case -2147483648:
                                                                                        if false {
                                                                                            if (4294967295 >= uint32(v1)) {
                                                                                                fmt.Println(v49, fi48)
                                                                                                v87 := GenS2{F0: 128, F1: (-(v49)), F2: (v67 >> (((v67 << ((0) & 63))) & 63))}
                                                                                                _ = v87
                                                                                            } else {
                                                                                                var v88 float32 = -2.25
                                                                                                var v89 bool = true
                                                                                                var v90 uint64 = v3
                                                                                                _ = v88
                                                                                                _ = v89
                                                                                                _ = v90
                                                                                            }
                                                                                            fmt.Println(v1, v73)
                                                                                            var v91 uint = (v67 % (((^(v67))) | 1))
                                                                                            var v92 uint64 = 425
                                                                                            var v93 bool = false
                                                                                            _ = v91
                                                                                            _ = v92
                                                                                            _ = v93
                                                                                        } else {
                                                                                            var v94 float32 = 1.5
                                                                                            var v95 float64 = v72
                                                                                            _ = v94
                                                                                            _ = v95
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            _ = v72
                                                                            _ = v73
                                                                        }
                                                                        var v96 float32 = (1.5 / ((1.5 / 3.14159) * 3.14159))
                                                                        _ = v96
                                                                    }
                                                                    v97 := GenS10{F0: v67, F1: (("" + ("foo bar" + "")) == "\n\t")}
                                                                    if true {
                                                                        v98 := GenS11{F0: -452}
                                                                        if (v1 < v1) {
                                                                            if false {
                                                                                if (!(((^(v1)) > (v1 >> (((v67 >> ((v67) & 63))) & 63))) || (((true || (2147483648 > 887)) || false) && true))) {
                                                                                    switch ("a" + (("foo bar" + "a") + "\n\t")) {
                                                                                    case "foo bar":
                                                                                        if (v3 <= 18446744073709551615) {
                                                                                            if ("\n\t" <= (("foo bar" + "\n\t") + ("hello" + (("\n\t" + "a") + "\n\t")))) {
                                                                                                v67 += v67
                                                                                                if (-111 <= -117) {
                                                                                                    v99 := GenS12{F0: v67, F1: 42, F2: v3}
                                                                                                    if len(v50) > 0 {
                                                                                                        v50[0] = fi48
                                                                                                    }
                                                                                                    v100 := GenS9{F0: (-(v64)), F1: 32767}
                                                                                                    _ = v100
                                                                                                    _ = v99
                                                                                                } else {
                                                                                                    switch (fi31 ^ (fi48 * fi48)) {
                                                                                                    case -919:
                                                                                                        if true {
                                                                                                            fmt.Println(fi31)
                                                                                                            if (65535 < 415) {
                                                                                                                v49 = (v49 % ((1000) | 1))
                                                                                                            } else {
                                                                                                                var v101 uint64 = 832
                                                                                                                v67 -= (v67 & v67)
                                                                                                                _ = v101
                                                                                                            }
                                                                                                            v49 -= v49
                                                                                                        } else {
                                                                                                            v64 = v64
                                                                                                        }
                                                                                                    case -458:
                                                                                                        var v102 int8 = 45
                                                                                                        _ = v102
                                                                                                    default:
                                                                                                        fmt.Println(v67)
                                                                                                    }
                                                                                                    var v103 float64 = 0.0
                                                                                                    switch "" {
                                                                                                    case "hello":
                                                                                                        v104 := GenS2{F0: (v1 & 2), F1: (v49 / (((v49 & (v49 | genFn0()))) | 1)), F2: 4294967295}
                                                                                                        _ = v104
                                                                                                    case "foo bar":
                                                                                                        if (v1 < v1) {
                                                                                                            v2 = int32(v67)
                                                                                                            v105 := GenS13{F0: 2147483647, F1: (v49 ^ (v49 & (v49 >> (((^(v67))) & 63)))), F2: 42}
                                                                                                            _ = v105
                                                                                                        } else {
                                                                                                            var v106 int64 = 1
                                                                                                            v107 := GenS14{F0: 64, F1: fi48, F2: (fi31 >> ((v67) & 63))}
                                                                                                            v50 = append(v50, (fi31 / ((9223372036854775807) | 1)))
                                                                                                            v108 := GenS15{F0: 127, F1: (v3 * v3)}
                                                                                                            _ = v106
                                                                                                            _ = v107
                                                                                                            _ = v108
                                                                                                        }
                                                                                                    case "a":
                                                                                                        v109 := GenS1{F0: 2, F1: (v3 & 1), F2: -1}
                                                                                                        if true {
                                                                                                            v50 = append(v50, fi48)
                                                                                                            if true {
                                                                                                                v49 -= (v49 | 32767)
                                                                                                                fmt.Println(fi31, v2, fi31)
                                                                                                                v64 = v64
                                                                                                            } else {
                                                                                                                var v110 string = ""
                                                                                                                v1 += (^(v1))
                                                                                                                _ = v110
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v109
                                                                                                    default:
                                                                                                        var v111 uint16 = 556
                                                                                                        _ = v111
                                                                                                    }
                                                                                                    _ = v103
                                                                                                }
                                                                                                v49 = genFn0()
                                                                                            } else {
                                                                                                if ((v49 >= (v49 ^ (v49 + (^(v49))))) || false) {
                                                                                                    var v112 uint8 = (v1 - (v1 | (v1 ^ 1)))
                                                                                                    _ = v112
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v3, v1)
                                                                                        }
                                                                                    }
                                                                                    var v113 uint16 = 817
                                                                                    _ = v113
                                                                                }
                                                                                v49 -= (v49 * v49)
                                                                            } else {
                                                                                var v114 float64 = float64((-2.25 * 1.5))
                                                                                v2 = v2
                                                                                var v115 int16 = (^(v49))
                                                                                var v116 uint64 = (v3 + v3)
                                                                                v117 := GenS2{F0: (v1 << ((112) & 63)), F1: (v49 + (v49 | (v115 & (v49 | v115)))), F2: uint(v116)}
                                                                                _ = v114
                                                                                _ = v115
                                                                                _ = v116
                                                                                _ = v117
                                                                            }
                                                                            if true {
                                                                                switch (fi31 / (((-(fi48))) | 1)) {
                                                                                case 128:
                                                                                    v118 := GenS16{F0: false}
                                                                                    _ = v118
                                                                                case 0:
                                                                                    if false {
                                                                                        var v119 float64 = 1e10
                                                                                        var v120 bool = (true && true)
                                                                                        _ = v119
                                                                                        _ = v120
                                                                                    }
                                                                                default:
                                                                                    var v121 bool = true
                                                                                    _ = v121
                                                                                }
                                                                                switch "hello" {
                                                                                case "a":
                                                                                    var v122 float32 = 1.5
                                                                                    _ = v122
                                                                                case "":
                                                                                    v64 = v64
                                                                                default:
                                                                                    v123 := GenS17{F0: 42}
                                                                                    _ = v123
                                                                                }
                                                                            }
                                                                            v49 = v49
                                                                            v1 = v1
                                                                            var v124 int64 = v64
                                                                            v1 -= (v1 << ((uint(v3)) & 63))
                                                                            _ = v124
                                                                        } else {
                                                                            v64 += 9223372036854775807
                                                                        }
                                                                        _ = v98
                                                                    } else {
                                                                        var v125 uint = (v67 - (v67 + v67))
                                                                        fmt.Println(v2, v3, v3)
                                                                        if (false && true) {
                                                                            v126 := GenS1{F0: 0, F1: (^(v3)), F2: fi48}
                                                                            v127 := GenS12{F0: 65535, F1: v67, F2: v3}
                                                                            v2 = v2
                                                                            _ = v126
                                                                            _ = v127
                                                                        }
                                                                        switch int(v3) {
                                                                        case 42:
                                                                            v128 := GenS2{F0: v1, F1: 1000, F2: v67}
                                                                            _ = v128
                                                                        case 4294967295:
                                                                            if (v2 != (v2 ^ (v2 * (v2 * 391)))) {
                                                                                v64 -= int64(v49)
                                                                                v129 := GenS9{F0: v64, F1: 1}
                                                                                if ((-(v2)) != (v2 | v2)) {
                                                                                    if ("" < "") {
                                                                                        if len(v50) > 0 {
                                                                                            v50[0] = -9223372036854775808
                                                                                        }
                                                                                        v130 := GenS14{F0: v1, F1: (fi48 ^ (fi48 >> ((18446744073709551615) & 63))), F2: (-(fi31))}
                                                                                        if (int64(fi48) == v64) {
                                                                                            v64 += (v64 * int64(v1))
                                                                                            v49 += genFn0()
                                                                                            if (false && true) {
                                                                                                v131 := GenS7{F0: true, F1: v49}
                                                                                                v132 := GenS18{F0: uint8(v125), F1: 32767}
                                                                                                var v133 bool = (((3.14159 - 3.14159) + 1e10) == (1e10 / 1.0))
                                                                                                var v134 uint = 18446744073709551615
                                                                                                _ = v131
                                                                                                _ = v132
                                                                                                _ = v133
                                                                                                _ = v134
                                                                                            }
                                                                                        }
                                                                                        _ = v130
                                                                                    } else {
                                                                                        switch fi48 {
                                                                                        case 32768:
                                                                                            v3 = v3
                                                                                        }
                                                                                    }
                                                                                    v2 += -538
                                                                                }
                                                                                _ = v129
                                                                            } else {
                                                                                v3 -= (v3 % (((^(v3))) | 1))
                                                                                switch "" {
                                                                                case "":
                                                                                    v135 := GenS19{F0: 915, F1: "a"}
                                                                                    _ = v135
                                                                                case "foo bar":
                                                                                    if (9223372036854775807 < (v64 / ((0) | 1))) {
                                                                                        if ((255 <= 1) || ((v64 << ((v125) & 63)) == (^(v64)))) {
                                                                                            v136 := GenS9{F0: v64, F1: 0}
                                                                                            if false {
                                                                                                fmt.Println(fi31)
                                                                                            } else {
                                                                                                v137 := GenS20{F0: 32767, F1: ""}
                                                                                                v49 = int16(v2)
                                                                                                _ = v137
                                                                                            }
                                                                                            v50 = append(v50, (fi48 >> ((uint(v3)) & 63)))
                                                                                            _ = v136
                                                                                        } else {
                                                                                            var v138 float64 = 0.0
                                                                                            if ((fi48 ^ int(v49)) == (fi31 / ((int(v64)) | 1))) {
                                                                                                v1 -= v1
                                                                                                var v139 float32 = (1.5 / -2.25)
                                                                                                _ = len(v50)
                                                                                                _ = v139
                                                                                            }
                                                                                            _ = v138
                                                                                        }
                                                                                        var v140 int16 = v49
                                                                                        v125 = 586
                                                                                        _ = v140
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    if len(v50) >= 2 {
                                                                                        v50 = v50[1:len(v50)]
                                                                                    }
                                                                                    if true {
                                                                                        if (32767 > 0) {
                                                                                            v49 = (v49 / (((^(v49))) | 1))
                                                                                            _ = len(v50)
                                                                                        } else {
                                                                                            if len(v50) >= 2 {
                                                                                                v50 = v50[1:len(v50)]
                                                                                            }
                                                                                            fmt.Println(v67)
                                                                                            v64 -= int64(v67)
                                                                                            if false {
                                                                                                switch 128 {
                                                                                                case 9223372036854775807:
                                                                                                    if true {
                                                                                                        _ = len(v50)
                                                                                                        fmt.Println(fi48, v2, fi48)
                                                                                                    }
                                                                                                default:
                                                                                                    fmt.Println(v64, v125, v49)
                                                                                                }
                                                                                                var v141 float32 = 0.0
                                                                                                v141 += ((v141 - v141) + ((v141 - (v141 + (3.14159 * 1.5))) * v141))
                                                                                                _ = v141
                                                                                            }
                                                                                        }
                                                                                        if ((0.1 > float32(((0.1 - ((-2.25 * 1e10) / -2.25)) + -2.25))) && (42 < 2147483647)) {
                                                                                            if (875 < 908) {
                                                                                                if len(v50) > 0 {
                                                                                                    v50[0] = (fi48 % (((fi31 ^ (^(fi31)))) | 1))
                                                                                                }
                                                                                                if len(v50) > 0 {
                                                                                                    v50[0] = 9223372036854775807
                                                                                                }
                                                                                                if true {
                                                                                                    if false {
                                                                                                        var v142 float32 = 3.14159
                                                                                                        _ = v142
                                                                                                    } else {
                                                                                                        _ = len(v50)
                                                                                                        var v143 int64 = (v64 | v64)
                                                                                                        v143 = v64
                                                                                                        _ = v143
                                                                                                    }
                                                                                                    if (42 >= -34) {
                                                                                                        var v144 int64 = (v64 / (((v64 ^ (v64 / ((1) | 1)))) | 1))
                                                                                                        var v145 uint32 = 1
                                                                                                        switch (fi48 % (((-(fi31))) | 1)) {
                                                                                                        case 4294967295:
                                                                                                            v67 -= v125
                                                                                                        case 32768:
                                                                                                            switch (fi31 >> (((v67 + v67)) & 63)) {
                                                                                                            case -1:
                                                                                                                v146 := GenS21{F0: v145, F1: (v145 - v145), F2: 42}
                                                                                                                _ = v146
                                                                                                            }
                                                                                                        case -1:
                                                                                                            if true {
                                                                                                                fmt.Println(v3)
                                                                                                                var v147 uint8 = (v1 | (v1 & 127))
                                                                                                                _ = v147
                                                                                                            } else {
                                                                                                                v144 += int64(v49)
                                                                                                            }
                                                                                                        default:
                                                                                                            switch "foo bar" {
                                                                                                            case "":
                                                                                                                switch "hello" {
                                                                                                                case "foo bar":
                                                                                                                    switch "foo bar" {
                                                                                                                    case "\n\t":
                                                                                                                        var v148 uint64 = uint64(v145)
                                                                                                                        _ = v148
                                                                                                                    case "a":
                                                                                                                        fmt.Println(v125)
                                                                                                                    case "":
                                                                                                                        var v149 uint16 = 42
                                                                                                                        _ = v149
                                                                                                                    default:
                                                                                                                        v1 = (^(v1))
                                                                                                                    }
                                                                                                                case "hello":
                                                                                                                    var v150 int16 = (-(v49))
                                                                                                                    _ = v150
                                                                                                                default:
                                                                                                                    switch 42 {
                                                                                                                    case 573:
                                                                                                                        _ = len(v50)
                                                                                                                    case -9223372036854775808:
                                                                                                                        switch 1 {
                                                                                                                        case -772:
                                                                                                                            v50 = append(v50, 32768)
                                                                                                                        case -843:
                                                                                                                            fmt.Println(v145, v1, v3)
                                                                                                                        default:
                                                                                                                            if len(v50) >= 2 {
                                                                                                                                v50 = v50[1:len(v50)]
                                                                                                                            }
                                                                                                                        }
                                                                                                                    case -359:
                                                                                                                        var v151 string = (("hello" + "hello") + "a")
                                                                                                                        _ = v151
                                                                                                                    default:
                                                                                                                        fmt.Println(fi48, fi48, fi48)
                                                                                                                    }
                                                                                                                }
                                                                                                            case "a":
                                                                                                                if false {
                                                                                                                    if (true || ((v49 | (v49 << ((v125) & 63))) <= v49)) {
                                                                                                                        var v152 uint = (v125 | v67)
                                                                                                                        if len(v50) >= 2 {
                                                                                                                            v50 = v50[1:len(v50)]
                                                                                                                        }
                                                                                                                        _ = v152
                                                                                                                    } else {
                                                                                                                        switch (fi48 * (fi31 + (fi48 % ((fi31) | 1)))) {
                                                                                                                        case 32767:
                                                                                                                            var v153 int32 = -1
                                                                                                                            _ = v153
                                                                                                                        }
                                                                                                                    }
                                                                                                                    switch "a" {
                                                                                                                    case "\n\t":
                                                                                                                        switch ("foo bar" + "") {
                                                                                                                        case "\n\t":
                                                                                                                            v154 := GenS22{F0: -128}
                                                                                                                            _ = v154
                                                                                                                        default:
                                                                                                                            v155 := GenS20{F0: genFn0(), F1: ("hello" + "foo bar")}
                                                                                                                            var v156 uint16 = 0
                                                                                                                            _ = v155
                                                                                                                            _ = v156
                                                                                                                        }
                                                                                                                    case "foo bar":
                                                                                                                        if len(v50) > 0 {
                                                                                                                            v50[0] = fi31
                                                                                                                        }
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    switch fi48 {
                                                                                                                    case -2147483648:
                                                                                                                        if len(v50) >= 2 {
                                                                                                                            v50 = v50[1:len(v50)]
                                                                                                                        }
                                                                                                                    case 42:
                                                                                                                        v3 += (v3 >> ((v125) & 63))
                                                                                                                    case 0:
                                                                                                                        if len(v50) >= 2 {
                                                                                                                            v50 = v50[1:len(v50)]
                                                                                                                        }
                                                                                                                        v49 += (^(v49))
                                                                                                                    }
                                                                                                                    var v157 uint64 = (v3 ^ (^(v3)))
                                                                                                                    v144 += (v144 * 42)
                                                                                                                    v144 = v144
                                                                                                                    if false {
                                                                                                                        if len(v50) > 0 {
                                                                                                                            v50[0] = (^(fi31))
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v157
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        v144 -= 9223372036854775807
                                                                                                        v67 = v67
                                                                                                        _ = v144
                                                                                                        _ = v145
                                                                                                    } else {
                                                                                                        v50 = append(v50, (fi48 | fi31))
                                                                                                        fmt.Println(fi48)
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                fmt.Println(v64, v3)
                                                                                                var v158 uint = 65535
                                                                                                v1 += uint8(v67)
                                                                                                _ = v158
                                                                                            }
                                                                                            var v159 float64 = 0.1
                                                                                            _ = v159
                                                                                        } else {
                                                                                            var v160 uint32 = 81
                                                                                            _ = v160
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v1, v67)
                                                                                    fmt.Println(v3, v3)
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v125
                                                                    }
                                                                    _ = v97
                                                                } else {
                                                                    if len(v50) >= 2 {
                                                                        v50 = v50[1:len(v50)]
                                                                    }
                                                                    if len(v50) >= 2 {
                                                                        v50 = v50[1:len(v50)]
                                                                    }
                                                                    switch (fi31 / (((fi48 >> ((v67) & 63))) | 1)) {
                                                                    case -9223372036854775808:
                                                                        _ = len(v50)
                                                                    case 42:
                                                                        var v161 int16 = v49
                                                                        _ = v161
                                                                    default:
                                                                        v49 = int16(v64)
                                                                    }
                                                                }
                                                                _ = v67
                                                            }
                                                        case "hello":
                                                            fmt.Println(v64, v3, v64)
                                                        default:
                                                            v3 += 2
                                                        }
                                                        _ = len(v50)
                                                        var v162 uint16 = 0
                                                        _ = v162
                                                    case "":
                                                        if false {
                                                            v163 := GenS23{F0: fi48, F1: v1}
                                                            if false {
                                                                if (true || true) {
                                                                    v164 := GenS24{F0: 53, F1: v1, F2: 39}
                                                                    _ = v164
                                                                }
                                                            } else {
                                                                _ = len(v50)
                                                                fmt.Println(v3, v64)
                                                            }
                                                            _ = len(v50)
                                                            _ = v163
                                                        }
                                                        var v165 int8 = 42
                                                        v165 = 0
                                                        _ = v165
                                                    default:
                                                        v64 = (v64 / ((v64) | 1))
                                                    }
                                                }
                                                switch (fi48 - (-(fi31))) {
                                                case 127:
                                                    var v166 uint64 = 0
                                                    _ = v166
                                                case 9223372036854775807:
                                                    fmt.Println(fi48, v3)
                                                default:
                                                    switch -472 {
                                                    case 42:
                                                        switch (("\n\t" + "\n\t") + "") {
                                                        case "hello":
                                                            if false {
                                                                var v167 int = (fi48 << ((794) & 63))
                                                                _ = v167
                                                            }
                                                        default:
                                                            v64 -= 1
                                                        }
                                                    case 724:
                                                        if (((v64 == v64) || ("hello" < ("\n\t" + ""))) && ((v3 - (v3 ^ v3)) >= v3)) {
                                                            v2 -= v2
                                                        }
                                                    case 853:
                                                        if false {
                                                            fmt.Println(v64, v2)
                                                        } else {
                                                            switch ("\n\t" + "") {
                                                            case "hello":
                                                                v49 += -32768
                                                            case "\n\t":
                                                                fmt.Println(v1)
                                                            case "":
                                                                switch "a" {
                                                                case "foo bar":
                                                                    var v168 string = ("foo bar" + "foo bar")
                                                                    _ = v168
                                                                case "a":
                                                                    fmt.Println(v49, v1, v2)
                                                                case "hello":
                                                                    fmt.Println(fi48, fi31, v3)
                                                                default:
                                                                    var v169 uint16 = 65535
                                                                    _ = v169
                                                                }
                                                            default:
                                                                v64 = (v64 ^ 0)
                                                            }
                                                        }
                                                    default:
                                                        v3 += (v3 & (v3 * v3))
                                                    }
                                                }
                                                v1 -= (v1 - (v1 | v1))
                                            } else {
                                                fmt.Println(v3)
                                                var v170 uint = 2
                                                if len(v50) >= 2 {
                                                    v50 = v50[1:len(v50)]
                                                }
                                                _ = v170
                                            }
                                            _ = v64
                                        case "\n\t":
                                            _ = len(v50)
                                            var v171 uint32 = 180
                                            var v172 int8 = 0
                                            _ = v171
                                            _ = v172
                                        }
                                    }
                                    fmt.Println(fi48)
                                }
                            } else {
                                v173 := GenS20{F0: -1, F1: ""}
                                _ = v173
                            }
                        }
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                    _ = v33
                } else {
                    var v174 uint64 = v3
                    var v175 int64 = -9223372036854775808
                    var v176 int32 = -494
                    _ = v174
                    _ = v175
                    _ = v176
                }
                _ = fi31
                _ = v32
            }
            switch 435 {
            case 0:
                v3 += 140
            case 42:
                switch (("" + ("" + ("\n\t" + ""))) + ("foo bar" + "foo bar")) {
                case "":
                    if false {
                        v177 := []int{}
                        _ = v177
                    }
                }
            default:
                var v178 float32 = 1e10
                _ = v178
            }
            var v179 int32 = (v2 - v2)
            _ = v179
        } else {
            var v180 int8 = 1
            var v181 uint16 = 797
            _ = v180
            _ = v181
        }
    }
    v1 += v1
    if (true || (true || (!(195 != uint(v3))))) {
        v2 -= (v2 ^ 42)
        var v182 bool = (117 > 1)
        for fi183 := 0; fi183 < 20; fi183++ {
            var v184 bool = v182
            for fi185 := 0; fi185 < 5; fi185++ {
                v186 := []int{(-(fi185)), fi185, fi185, fi183}
                _ = fi185
                _ = v186
            }
            _ = fi183
            _ = v184
        }
        _ = v182
    } else {
        for fi187 := 0; fi187 < 6; fi187++ {
            v188 := []int{fi187, -287, -855, fi187}
            _ = fi187
            _ = v188
        }
        fmt.Println(v3, v1)
        var v189 uint32 = 42
        _ = v189
    }
    if true {
        switch 4294967295 {
        case 32768:
            v190 := GenS2{F0: 1, F1: -1, F2: 42}
            _ = v190
        case 32767:
            if false {
                switch "a" {
                case "":
                    var v191 bool = false
                    _ = v191
                case "a":
                    fmt.Println(v2)
                }
            } else {
                v3 -= v3
                var v192 string = ""
                _ = v192
            }
        default:
            v193 := GenS25{F0: 0}
            _ = v193
        }
    } else {
        v194 := []int{}
        if true {
            if (!false) {
                if true {
                    var v195 uint32 = 42
                    _ = v195
                }
                var v196 uint8 = (^(v1))
                v1 = (^(v1))
                _ = v196
            }
        } else {
            var v197 uint32 = uint32(v3)
            _ = v197
        }
        _ = v194
    }
    fmt.Println(v3)
    v198 := []int{}
    v199 := GenS26{F0: 133, F1: ""}
    var v200 int8 = -128
    v3 += 478
    var v201 uint32 = 4294967295
    for fi202 := 0; fi202 < 2; fi202++ {
        if len(v198) > 0 {
            v198[0] = fi202
        }
        _ = fi202
    }
    switch -405 {
    case -573:
        if len(v198) >= 2 {
            v198 = v198[1:len(v198)]
        }
    case 32768:
        for fi203 := 0; fi203 < 4; fi203++ {
            for fi204 := 0; fi204 < 9; fi204++ {
                switch fi204 {
                case -2147483648:
                    if len(v198) >= 2 {
                        v198 = v198[1:len(v198)]
                    }
                }
                _ = fi204
            }
            _ = fi203
        }
    default:
        v1 -= v1
    }
    v201 = (v201 | (v201 / (((v201 + (v201 % ((v201) | 1)))) | 1)))
    if ("\n\t" < "\n\t") {
        v205 := GenS19{F0: -342, F1: ("\n\t" + "")}
        v206 := GenS27{F0: -9223372036854775808, F1: (v3 >> ((uint(v3)) & 63))}
        var v207 uint8 = 1
        _ = v205
        _ = v206
        _ = v207
    } else {
        for fi208 := 0; fi208 < 14; fi208++ {
            for fi209 := 0; fi209 < 13; fi209++ {
                switch "a" {
                case "":
                    var v210 uint32 = uint32(v1)
                    _ = v210
                case "a":
                    if false {
                        var v211 int32 = (v2 ^ v2)
                        v212 := GenS28{F0: (v211 % ((2147483647) | 1))}
                        _ = v211
                        _ = v212
                    }
                case "\n\t":
                    var v213 uint64 = 2
                    _ = v213
                default:
                    v214 := GenS4{F0: 4294967295, F1: 42, F2: v1}
                    _ = v214
                }
                if len(v198) > 0 {
                    v198[0] = fi209
                }
                var v215 uint32 = (v201 << ((1) & 63))
                _ = fi209
                _ = v215
            }
            v1 = (v1 << ((42) & 63))
            _ = fi208
        }
        for fi216 := 0; fi216 < 4; fi216++ {
            if len(v198) >= 2 {
                v198 = v198[1:len(v198)]
            }
            _ = fi216
        }
        var v217 float64 = -2.25
        _ = v217
    }
    for fi218 := 0; fi218 < 20; fi218++ {
        fmt.Println(v3, fi218)
        var v219 uint16 = 227
        v200 = (^(v200))
        _ = fi218
        _ = v219
    }
    v3 = (v3 - v3)
    if len(v198) >= 2 {
        v198 = v198[1:len(v198)]
    }
    for fi220 := 0; fi220 < 7; fi220++ {
        var v221 int16 = 0
        _ = fi220
        _ = v221
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v198:", len(v198), v198)
    fmt.Println("v2:", v2)
    fmt.Println("v200:", v200)
    fmt.Println("v201:", v201)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v198
    _ = v199
    _ = v2
    _ = v200
    _ = v201
    _ = v3
}
