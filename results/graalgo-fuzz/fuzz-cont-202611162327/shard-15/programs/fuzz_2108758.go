package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 string; F2 uint32 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64; F1 bool }
type GenS8 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS9 struct { F0 uint; F1 uint8 }
type GenS10 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS11 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS12 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS13 struct { F0 int; F1 int16; F2 uint16 }
type GenS14 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 bool; F1 int; F2 int64 }
type GenS17 struct { F0 bool }
func genFn0(p0 int16) int16 {
    var v5 float32 = (0.0 - 0.1)
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := []int{(fi6 & (fi6 + (fi6 / ((fi6) | 1)))), (fi6 + fi6)}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        if (false && ((42 < 2) || false)) {
            p0 -= (^(p0))
        } else {
            var v9 int32 = -2147483648
            var v10 int8 = -128
            _ = v10
            _ = v9
        }
        v11 := []int{}
        _ = fi8
        _ = v11
    }
    _ = p0
    _ = v5
    return p0
}
func genFn1(p0 int8, p1 bool, p2 int) int32 {
    fmt.Println(p2, p1)
    v12 := GenS1{F0: 255}
    if (("\n\t" <= ((("\n\t" + "hello") + "a") + "hello")) && (!p1)) {
        var v13 int = 9223372036854775807
        v14 := GenS1{F0: 661}
        if (791 <= 154) {
            switch 171 {
            case 32767:
                v15 := GenS2{F0: 758}
                _ = v15
            }
            v16 := []int{32767, v13}
            var v17 uint8 = 42
            _ = v16
            _ = v17
        }
        _ = v13
        _ = v14
    }
    if p1 {
        v18 := GenS1{F0: 255}
        _ = v18
    } else {
        if (347 >= 32767) {
            var v19 uint8 = 2
            _ = v19
        }
        switch (p2 * -1) {
        case 31:
            var v20 string = ""
            _ = v20
        case 9223372036854775807:
            v21 := GenS2{F0: 0}
            _ = v21
        case -9223372036854775808:
            for fi22 := 0; fi22 < 7; fi22++ {
                switch "\n\t" {
                case "":
                    if (42 != 455) {
                        var v23 int8 = p0
                        v24 := GenS3{F0: 1, F1: "hello", F2: 2}
                        _ = v23
                        _ = v24
                    }
                case "a":
                    var v25 uint64 = 4294967295
                    _ = v25
                case "foo bar":
                    for fi26 := 0; fi26 < 19; fi26++ {
                        p0 = (^(p0))
                        v27 := GenS4{F0: 26, F1: (p0 & p0)}
                        _ = fi26
                        _ = v27
                    }
                }
                if true {
                    v28 := []int{(p2 + fi22), (fi22 >> ((226) & 63)), fi22}
                    _ = v28
                }
                _ = fi22
            }
        }
    }
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return -564
}

func main() {
    var v1 uint64 = 0
    v2 := []int{int(v1), -2147483648, -1, -55}
    var v3 float32 = 1.5
    var v4 uint = 4294967295
    var v29 int8 = 42
    if (!(("\n\t" + "hello") <= "")) {
        var v30 bool = (838 >= 369)
        if v30 {
            switch 127 {
            case 808:
                if v30 {
                    v31 := GenS5{F0: 42, F1: (v30 || ((float32(-2.25) / v3) != (float32((0.1 + (0.1 / 1.0))) + v3)))}
                    if ((!(883 < 726)) && (-1 < 0)) {
                        v32 := GenS6{F0: -9223372036854775808}
                        switch ("hello" + "\n\t") {
                        case "foo bar":
                            v33 := GenS7{F0: (v1 % ((v1) | 1)), F1: v30}
                            _ = v33
                        }
                        _ = v32
                    } else {
                        var v34 uint64 = v1
                        v1 = 18446744073709551615
                        if ((v29 & v29) > (v29 & v29)) {
                            for fi35 := 0; fi35 < 10; fi35++ {
                                v3 += (v3 - v3)
                                var v36 uint64 = 4294967295
                                var v37 string = "a"
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                            for fi38 := 0; fi38 < 9; fi38++ {
                                switch (fi38 >> (((v4 & 255)) & 63)) {
                                case 32767:
                                    var v39 bool = v30
                                    _ = v39
                                default:
                                    if len(v2) > 0 {
                                        v2[0] = fi38
                                    }
                                }
                                _ = fi38
                            }
                        }
                        _ = v34
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v31
                }
            }
            switch -9223372036854775808 {
            case 128:
                switch 655 {
                case 0:
                    if v30 {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v40 bool = v30
                        _ = v40
                    }
                case 123:
                    if v30 {
                        var v41 int8 = (v29 >> (((v4 * v4)) & 63))
                        v1 = v1
                        _ = v41
                    } else {
                        if (((v4 >> ((528) & 63)) <= 1) && v30) {
                            v30 = (!(1 != 879))
                            switch -9223372036854775808 {
                            case 785:
                                for fi42 := 0; fi42 < 11; fi42++ {
                                    _ = len(v2)
                                    for fi43 := 0; fi43 < 9; fi43++ {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        fmt.Println(fi42)
                                        v29 = (-(v29))
                                        _ = fi43
                                    }
                                    if true {
                                        switch 0 {
                                        case 920:
                                            switch "hello" {
                                            case "a":
                                                switch "a" {
                                                case "hello":
                                                    fmt.Println(v3, v3, v1)
                                                case "\n\t":
                                                    for fi44 := 0; fi44 < 9; fi44++ {
                                                        var v45 float64 = 1.5
                                                        if (!(genFn1((v29 % (((^(v29))) | 1)), (((1 > 0) || true) && (("" <= ("a" + "hello")) || v30)), fi42) <= -1)) {
                                                            v4 = 61
                                                            var v46 int16 = genFn0(-32768)
                                                            var v47 int32 = int32(v29)
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                        switch "foo bar" {
                                                        case "":
                                                            if len(v2) > 0 {
                                                                v2[0] = (fi42 + (fi44 % ((9223372036854775807) | 1)))
                                                            }
                                                        case "foo bar":
                                                            v45 = -2.25
                                                        case "hello":
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        v29 -= (v29 - -128)
                                                        _ = fi44
                                                        _ = v45
                                                    }
                                                }
                                            case "foo bar":
                                                for fi48 := 0; fi48 < 16; fi48++ {
                                                    if len(v2) > 0 {
                                                        v2[0] = fi48
                                                    }
                                                    v49 := GenS7{F0: 669, F1: v30}
                                                    switch "\n\t" {
                                                    case "":
                                                        var v50 int = fi48
                                                        _ = v50
                                                    case "\n\t":
                                                        v30 = (!v30)
                                                    case "foo bar":
                                                        v51 := GenS8{F0: 1, F1: (v29 - (v29 - 107)), F2: 127}
                                                        _ = v51
                                                    }
                                                    var v52 bool = (76 == 42)
                                                    v53 := GenS4{F0: v29, F1: 127}
                                                    _ = fi48
                                                    _ = v49
                                                    _ = v52
                                                    _ = v53
                                                }
                                            default:
                                                var v54 uint8 = 26
                                                _ = v54
                                            }
                                        default:
                                            switch "\n\t" {
                                            case "\n\t":
                                                for fi55 := 0; fi55 < 8; fi55++ {
                                                    _ = len(v2)
                                                    var v56 int = (fi55 + (fi55 * fi42))
                                                    _ = fi55
                                                    _ = v56
                                                }
                                            }
                                        }
                                        v57 := GenS9{F0: (v4 >> ((1) & 63)), F1: 0}
                                        var v58 int64 = 9223372036854775807
                                        _ = v57
                                        _ = v58
                                    } else {
                                        fmt.Println(v29, v1)
                                        if v30 {
                                            v2 = append(v2, (fi42 + fi42))
                                            var v59 int64 = 0
                                            _ = v59
                                        } else {
                                            v3 = v3
                                            v29 += v29
                                        }
                                        var v60 uint32 = 588
                                        _ = v60
                                    }
                                    _ = fi42
                                }
                            }
                            switch 32767 {
                            case 727:
                                v4 = v4
                            case 42:
                                v1 -= v1
                            case 4294967295:
                                if false {
                                    v4 -= 255
                                    v29 = (^(v29))
                                }
                            }
                        }
                        v61 := GenS1{F0: 0}
                        switch 9223372036854775807 {
                        case 42:
                            v4 += v4
                        case 2147483647:
                            var v62 uint32 = 1
                            _ = v62
                        case -72:
                            v63 := GenS10{F0: 766, F1: 128, F2: 42}
                            _ = v63
                        }
                        _ = v61
                    }
                case -205:
                    fmt.Println(v30)
                default:
                    switch "a" {
                    case "foo bar":
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if v30 {
                                if v30 {
                                    var v64 float64 = float64(v3)
                                    v4 = v4
                                    _ = v64
                                }
                            } else {
                                switch "" {
                                case "a":
                                    if false {
                                        var v65 uint = v4
                                        v3 = (v3 / (v3 / v3))
                                        for fi66 := 0; fi66 < 1; fi66++ {
                                            switch fi66 {
                                            case -2147483648:
                                                v1 = (v1 + v1)
                                            case 9223372036854775807:
                                                v65 = v4
                                            case -593:
                                                if (v30 && false) {
                                                    if v30 {
                                                        for fi67 := 0; fi67 < 14; fi67++ {
                                                            var v68 uint = v65
                                                            _ = fi67
                                                            _ = v68
                                                        }
                                                    }
                                                    v30 = (-157 > 2147483647)
                                                } else {
                                                    var v69 uint32 = 4294967295
                                                    v65 -= v65
                                                    if (!true) {
                                                        if v30 {
                                                            fmt.Println(fi66, v69)
                                                            fmt.Println(v1, v1)
                                                        }
                                                    } else {
                                                        v65 = 4294967295
                                                        for fi70 := 0; fi70 < 2; fi70++ {
                                                            var v71 float32 = v3
                                                            _ = fi70
                                                            _ = v71
                                                        }
                                                    }
                                                    _ = v69
                                                }
                                            default:
                                                fmt.Println(v65, v4)
                                            }
                                            _ = fi66
                                        }
                                        _ = v65
                                    } else {
                                        for fi72 := 0; fi72 < 15; fi72++ {
                                            switch (fi72 - fi72) {
                                            case 641:
                                                for fi73 := 0; fi73 < 10; fi73++ {
                                                    var v74 float32 = v3
                                                    _ = fi73
                                                    _ = v74
                                                }
                                            default:
                                                switch "a" {
                                                case "":
                                                    v75 := GenS1{F0: 32767}
                                                    _ = v75
                                                }
                                            }
                                            v2 = append(v2, (fi72 * fi72))
                                            fmt.Println(v3)
                                            _ = fi72
                                        }
                                        fmt.Println(v30, v29)
                                        _ = len(v2)
                                    }
                                case "":
                                    if ((float32((float64((v3 + (float32(-2.25) + v3))) + 3.14159)) * v3) < (float32(3.14159) - ((v3 / float32(1e10)) - ((((1.5 + ((1.5 * 1.5) * v3)) / (1.5 / float32(((0.1 / 1e10) / float64((float32((0.1 + ((0.1 - 1e10) / float64((float32(1e10) + v3))))) * 1e10)))))) - v3) * v3)))) {
                                        v76 := GenS11{F0: (2 >= 2), F1: 2, F2: 1}
                                        if (7 <= 9223372036854775807) {
                                            var v77 int8 = (v29 >> ((v4) & 63))
                                            _ = v77
                                        } else {
                                            v4 += v4
                                            if ((-2.25 * 0.0) <= 0.1) {
                                                for fi78 := 0; fi78 < 6; fi78++ {
                                                    fmt.Println(v1, v4, v29)
                                                    v2 = append(v2, (fi78 - (-(fi78))))
                                                    v4 = (v4 | 329)
                                                    _ = fi78
                                                }
                                                if true {
                                                    v29 += (v29 / ((0) | 1))
                                                    fmt.Println(v1, v30, v1)
                                                }
                                                if (-1 != genFn1(v29, v30, -2147483648)) {
                                                    v79 := GenS12{F0: 272, F1: (v1 ^ v1), F2: -32768}
                                                    _ = v79
                                                }
                                            }
                                            fmt.Println(v30, v4)
                                        }
                                        _ = v76
                                    }
                                case "foo bar":
                                    v80 := GenS13{F0: -506, F1: 42, F2: 2}
                                    _ = v80
                                default:
                                    v81 := GenS12{F0: 4294967295, F1: (v1 - (v1 + (v1 << (((v4 & 18446744073709551615)) & 63)))), F2: 1}
                                    _ = v81
                                }
                            }
                            if true {
                                for fi82 := 0; fi82 < 17; fi82++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v3 = v3
                                    v1 -= 0
                                    _ = fi82
                                }
                                v83 := GenS9{F0: (v4 << ((v4) & 63)), F1: 67}
                                _ = v83
                            }
                        } else {
                            var v84 int = 9223372036854775807
                            for fi85 := 0; fi85 < 15; fi85++ {
                                if (v30 || false) {
                                    v3 = 0.0
                                    if false {
                                        var v86 uint64 = v1
                                        for fi87 := 0; fi87 < 1; fi87++ {
                                            v88 := GenS11{F0: v30, F1: uint8(v1), F2: 255}
                                            _ = fi87
                                            _ = v88
                                        }
                                        if (!false) {
                                            var v89 float64 = 0.1
                                            if v30 {
                                                if true {
                                                    switch (^(fi85)) {
                                                    case 32768:
                                                        if (v86 >= (^(v1))) {
                                                            v84 -= -2147483648
                                                            var v90 bool = (v30 || (v1 == 18446744073709551615))
                                                            _ = v90
                                                        }
                                                    case 1:
                                                        var v91 uint16 = 65535
                                                        _ = v91
                                                    }
                                                }
                                                v1 = v86
                                            }
                                            fmt.Println(fi85)
                                            _ = v89
                                        }
                                        _ = v86
                                    } else {
                                        v84 = 309
                                        var v92 uint8 = 248
                                        _ = v92
                                    }
                                } else {
                                    for fi93 := 0; fi93 < 1; fi93++ {
                                        fmt.Println(fi85)
                                        var v94 int32 = int32(v29)
                                        _ = fi93
                                        _ = v94
                                    }
                                    if (!(1 != 588)) {
                                        for fi95 := 0; fi95 < 13; fi95++ {
                                            fmt.Println(fi95, v1)
                                            if false {
                                                v30 = false
                                            }
                                            _ = fi95
                                        }
                                        fmt.Println(v3)
                                    }
                                    if false {
                                        fmt.Println(fi85, v4, v84)
                                        v96 := GenS8{F0: 566, F1: (-(v29)), F2: 139}
                                        _ = v96
                                    }
                                }
                                var v97 int16 = 745
                                if v30 {
                                    v2 = append(v2, (fi85 + fi85))
                                } else {
                                    var v98 int16 = genFn0(v97)
                                    v99 := GenS14{F0: 42, F1: ("foo bar" < (("\n\t" + "a") + "a")), F2: 127}
                                    _ = v98
                                    _ = v99
                                }
                                _ = fi85
                                _ = v97
                            }
                            for fi100 := 0; fi100 < 4; fi100++ {
                                var v101 float64 = -2.25
                                for fi102 := 0; fi102 < 16; fi102++ {
                                    var v103 bool = v30
                                    if false {
                                        if v30 {
                                            if true {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                fmt.Println(v84)
                                            } else {
                                                fmt.Println(v29, fi102)
                                            }
                                            v29 -= v29
                                            if v30 {
                                                var v104 uint8 = 128
                                                var v105 float32 = v3
                                                v30 = (((((v104 | 255) != (v104 * 181)) || v30) || v30) || ((!true) && v30))
                                                v103 = v103
                                                v103 = v30
                                                _ = v104
                                                _ = v105
                                            } else {
                                                var v106 float64 = (float64((float32((0.0 / (1e10 * 0.1))) * 0.0)) / 1.5)
                                                _ = v106
                                            }
                                        }
                                        var v107 float64 = 3.14159
                                        var v108 int8 = (v29 ^ 127)
                                        _ = v107
                                        _ = v108
                                    } else {
                                        switch ("hello" + "hello") {
                                        case "a":
                                            v103 = (!v103)
                                        case "\n\t":
                                            v109 := GenS12{F0: 42, F1: 2, F2: genFn0(473)}
                                            _ = v109
                                        }
                                        v110 := GenS2{F0: 42}
                                        if (2 >= 2) {
                                            if v30 {
                                                switch "\n\t" {
                                                case "":
                                                    var v111 uint8 = 19
                                                    _ = v111
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = (fi102 | (fi100 & fi100))
                                                }
                                                v112 := GenS10{F0: 1, F1: 2, F2: 246}
                                                _ = v112
                                            }
                                            v101 += v101
                                        }
                                        fmt.Println(fi102, v30, v84)
                                        _ = v110
                                    }
                                    if (-1 < 1000) {
                                        _ = len(v2)
                                        v113 := GenS14{F0: 32767, F1: (307 <= 1), F2: 1}
                                        fmt.Println(fi102, v103)
                                        _ = v113
                                    }
                                    _ = fi102
                                    _ = v103
                                }
                                v114 := GenS9{F0: v4, F1: 1}
                                _ = fi100
                                _ = v101
                                _ = v114
                            }
                            _ = v84
                        }
                    case "\n\t":
                        if len(v2) > 0 {
                            v2[0] = -758
                        }
                    case "":
                        if (((-2.25 / 1.0) - 3.14159) <= 1.5) {
                            var v115 string = (("foo bar" + "a") + "foo bar")
                            var v116 int = 4294967295
                            v116 = (v116 % ((v116) | 1))
                            _ = v115
                            _ = v116
                        }
                    default:
                        switch 9223372036854775807 {
                        case -410:
                            for fi117 := 0; fi117 < 19; fi117++ {
                                for fi118 := 0; fi118 < 15; fi118++ {
                                    if false {
                                        if false {
                                            v3 = (((-2.25 / v3) / v3) * float32(3.14159))
                                            v3 += ((v3 * v3) / v3)
                                            fmt.Println(v4, v1)
                                        } else {
                                            if (810 > 1) {
                                                v119 := GenS12{F0: 342, F1: v1, F2: -80}
                                                _ = v119
                                            } else {
                                                v4 = (^(v4))
                                                v4 -= (v4 | 255)
                                                fmt.Println(fi118)
                                            }
                                            _ = len(v2)
                                            var v120 uint = 4294967295
                                            _ = v120
                                        }
                                        var v121 int64 = -549
                                        v30 = (!v30)
                                        _ = v121
                                    }
                                    _ = fi118
                                }
                                var v122 int64 = 42
                                _ = fi117
                                _ = v122
                            }
                        case 9223372036854775807:
                            var v123 uint8 = 1
                            _ = v123
                        case -558:
                            var v124 uint = v4
                            _ = v124
                        default:
                            fmt.Println(v4, v4)
                        }
                    }
                }
            default:
                var v125 uint8 = 7
                _ = v125
            }
            v4 = (^(v4))
        } else {
            v2 = append(v2, 32768)
        }
        _ = v30
    } else {
        v126 := GenS13{F0: 32768, F1: genFn0(1), F2: 522}
        v1 -= v1
        v127 := GenS9{F0: (v4 & 65535), F1: 128}
        _ = v126
        _ = v127
    }
    v29 = -1
    v128 := GenS2{F0: -321}
    v29 = 42
    _ = len(v2)
    v29 = v29
    var v129 uint64 = v1
    v3 = v3
    switch 32768 {
    case 1:
        v130 := GenS12{F0: 0, F1: v129, F2: genFn0(int16(v1))}
        _ = v130
    case 9223372036854775807:
        if true {
            switch "hello" {
            case "":
                v131 := GenS9{F0: (v4 % ((v4) | 1)), F1: 2}
                _ = v131
            case "\n\t":
                fmt.Println(v4, v129, v129)
            }
            var v132 uint16 = 93
            v133 := GenS12{F0: 750, F1: (v1 & v129), F2: 42}
            _ = v132
            _ = v133
        } else {
            v4 -= v4
            v134 := GenS15{F0: -128}
            _ = v134
        }
    }
    v4 = (v4 % ((42) | 1))
    v129 = (v1 * v129)
    v135 := GenS16{F0: true, F1: 128, F2: 1}
    for fi136 := 0; fi136 < 2; fi136++ {
        var v137 uint = 637
        var v138 bool = true
        _ = fi136
        _ = v137
        _ = v138
    }
    for fi139 := 0; fi139 < 1; fi139++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v140 int64 = 1
        _ = fi139
        _ = v140
    }
    if false {
        if false {
            v141 := GenS17{F0: false}
            v29 -= (v29 % (((v29 ^ (v29 ^ v29))) | 1))
            if false {
                v129 = (^(v129))
                v1 += (^(v1))
                if ("\n\t" != ("foo bar" + "\n\t")) {
                    _ = len(v2)
                    var v142 int32 = genFn1((v29 / ((-1) | 1)), true, 32767)
                    v3 += v3
                    _ = v142
                } else {
                    for fi143 := 0; fi143 < 8; fi143++ {
                        var v144 float32 = v3
                        v1 = (v129 | (v129 % (((v1 << ((v4) & 63))) | 1)))
                        _ = fi143
                        _ = v144
                    }
                    v3 = float32((0.0 / 1.0))
                }
            }
            _ = v141
        }
    }
    if len(v2) > 0 {
        v2[0] = 42
    }
    var v145 int = -1
    fmt.Println("v1:", v1)
    fmt.Println("v129:", v129)
    fmt.Println("v145:", v145)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v128
    _ = v129
    _ = v135
    _ = v145
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
