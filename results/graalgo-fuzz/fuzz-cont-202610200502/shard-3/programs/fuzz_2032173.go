package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int16 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int; F1 int16; F2 int64 }
type GenS5 struct { F0 int32; F1 string; F2 uint }
type GenS6 struct { F0 uint16; F1 uint; F2 int16 }
type GenS7 struct { F0 int; F1 uint16; F2 int64 }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 uint8; F1 int16; F2 int8 }
type GenS10 struct { F0 int; F1 uint64 }
type GenS11 struct { F0 bool; F1 uint16; F2 int16 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint32; F1 uint16 }
type GenS15 struct { F0 uint32; F1 int }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint64 }
func genFn0() int {
    var v7 bool = ("\n\t" == ("a" + "hello"))
    fmt.Println(v7)
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int32 = 0
        v10 := []int{fi8, (^(fi8)), fi8, fi8}
        v11 := GenS1{F0: (fi8 * fi8), F1: (v9 * 0), F2: 0}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    switch 2147483647 {
    case 32768:
        var v12 int32 = -2147483648
        _ = v12
    case 32767:
        v7 = ((1.5 != 0.0) && (true && v7))
    default:
        if v7 {
            v13 := GenS2{F0: v7, F1: 18446744073709551615}
            _ = v13
        }
    }
    v7 = true
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 uint8) int {
    if ((true || (2147483647 != 569)) || (-1 == 1)) {
        var v14 int = 675
        var v15 int64 = 0
        _ = v14
        _ = v15
    } else {
        switch -9223372036854775808 {
        case -470:
            v16 := []int{}
            _ = v16
        default:
            if (!(65535 < 2)) {
                var v17 int32 = 1
                _ = v17
            } else {
                p0 += p0
            }
        }
    }
    v18 := GenS3{F0: true, F1: false}
    _ = p0
    _ = v18
    return 127
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = ("" + "")
    v3 := make(map[string]string)
    var v4 int32 = 0
    var v5 float32 = (0.1 / (((1.5 - 0.1) - -2.25) / 1.0))
    var v6 uint64 = 1
    delete(v1, -2147483648)
    switch (("\n\t" + "\n\t") + v2) {
    case "foo bar":
        if true {
            v5 = float32((1.5 / 1.5))
            fmt.Println(v2)
        } else {
            for fi19 := 0; fi19 < 20; fi19++ {
                v20 := []int{fi19}
                var v21 int64 = 1
                _ = fi19
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 12; fi22++ {
                fmt.Println(v2, v5, fi22)
                if (false || true) {
                    var v23 int = int(v6)
                    _ = v23
                } else {
                    fmt.Println(v5, v2, v5)
                }
                _ = fi22
            }
            for fi24 := 0; fi24 < 19; fi24++ {
                var v25 bool = ((-(fi24)) > (fi24 % ((fi24) | 1)))
                _ = fi24
                _ = v25
            }
        }
    case "":
        var v26 bool = true
        _ = v26
    case "a":
        v27 := GenS1{F0: 0, F1: (v4 << ((0) & 63)), F2: 32767}
        _ = v27
    default:
        fmt.Println(v4, v2, v2)
    }
    var v28 int8 = 35
    if (!true) {
        if (false || true) {
            v6 = (v6 & v6)
        }
        v29 := []int{-317, 128, genFn0()}
        _ = v29
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        var v31 uint64 = 42
        var v32 string = (("" + (v2 + (v2 + v2))) + v2)
        _ = fi30
        _ = v31
        _ = v32
    }
    v6 -= v6
    delete(v3, ("" + v2))
    switch v2 {
    case "\n\t":
        _ = len(v1)
    case "foo bar":
        switch v2 {
        case "a":
            if (42 <= 255) {
                delete(v1, -2147483648)
                v33 := GenS4{F0: 268, F1: 1000, F2: 42}
                for fi34 := 0; fi34 < 18; fi34++ {
                    v4 = v4
                    for fi35 := 0; fi35 < 11; fi35++ {
                        switch (v2 + "foo bar") {
                        case "":
                            v6 = 1
                        case "foo bar":
                            v3[("\n\t" + (v2 + ((v2 + "") + (("a" + "\n\t") + "\n\t"))))] = "foo bar"
                        }
                        switch "\n\t" {
                        case "":
                            v28 += (v28 & -128)
                            v2 = (v2 + (v2 + v2))
                        case "hello":
                            var v36 int64 = -9223372036854775808
                            switch (fi35 >> ((18446744073709551615) & 63)) {
                            case 4294967295:
                                v2 = ("" + "foo bar")
                            default:
                                v5 -= 0.0
                            }
                            _ = v36
                        default:
                            switch (fi35 / (((fi35 | genFn1(42))) | 1)) {
                            case 868:
                                var v37 int8 = v28
                                _ = v37
                            case 127:
                                var v38 int64 = -1
                                _ = v38
                            case 4294967295:
                                fmt.Println(v5)
                            }
                        }
                        _ = len(v1)
                        _ = fi35
                    }
                    _ = fi34
                }
                _ = v33
            } else {
                if (false && true) {
                    fmt.Println(v28, v5, v6)
                } else {
                    switch v2 {
                    case "hello":
                        for fi39 := 0; fi39 < 20; fi39++ {
                            if false {
                                v28 = (v28 * -66)
                            } else {
                                for fi40 := 0; fi40 < 19; fi40++ {
                                    if true {
                                        fmt.Println(v6, v4, v2)
                                    }
                                    if false {
                                        v41 := []int{genFn0(), fi39, -2147483648, fi40}
                                        v6 -= v6
                                        var v42 uint32 = 2
                                        _ = v41
                                        _ = v42
                                    }
                                    if (42 > 1) {
                                        var v43 float32 = ((v5 * (0.0 - float32(0.1))) * v5)
                                        v44 := GenS2{F0: (false && (-32768 == -32768)), F1: 42}
                                        var v45 string = v2
                                        switch 0 {
                                        case -865:
                                            v2 = (v2 + "hello")
                                        case 9223372036854775807:
                                            fmt.Println(fi39, v45)
                                        default:
                                            fmt.Println(v45)
                                        }
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    }
                                    _ = fi40
                                }
                            }
                            v3[v2] = "hello"
                            _ = fi39
                        }
                    }
                    var v46 bool = (false && false)
                    for fi47 := 0; fi47 < 7; fi47++ {
                        v48 := GenS5{F0: (v4 ^ (-(v4))), F1: v2, F2: 0}
                        _ = len(v1)
                        _ = fi47
                        _ = v48
                    }
                    _ = v46
                }
            }
        }
    default:
        switch "foo bar" {
        case "a":
            switch 32767 {
            case 4294967295:
                _ = len(v3)
            case 118:
                var v49 bool = (!(true && (!false)))
                _ = v49
            default:
                fmt.Println(v28, v4)
            }
        case "hello":
            switch v2 {
            case "a":
                for fi50 := 0; fi50 < 18; fi50++ {
                    switch (-(fi50)) {
                    case 4294967295:
                        var v51 bool = true
                        _ = v51
                    }
                    _ = fi50
                }
            case "":
                var v52 uint64 = v6
                _ = v52
            default:
                v53 := GenS5{F0: v4, F1: "\n\t", F2: 0}
                _ = v53
            }
        case "":
            v54 := GenS3{F0: false, F1: (false && (1e10 < 1.5))}
            _ = v54
        }
    }
    v6 -= (v6 - 42)
    switch 1 {
    case -9223372036854775808:
        var v55 uint32 = 2147483648
        _ = v55
    case 128:
        if ((float64(1e10) + (float64(v5) - (1e10 - 1e10))) < float64(v5)) {
            if true {
                v56 := GenS3{F0: (0.0 >= (0.1 * (float64(v5) / (0.0 + -2.25)))), F1: (-9223372036854775808 != 9223372036854775807)}
                _ = v56
            } else {
                switch (v2 + v2) {
                case "":
                    if (-384 > 700) {
                        v3[(v2 + (("hello" + v2) + ("a" + v2)))] = ("\n\t" + "foo bar")
                    }
                default:
                    var v57 string = (("hello" + v2) + ("hello" + v2))
                    _ = v57
                }
            }
            v58 := []int{-2147483648, -524}
            _ = v58
        } else {
            if (-1 < 1) {
                v5 -= ((0.1 + 1.5) + v5)
                v59 := []int{}
                _ = len(v59)
                _ = v59
            }
            switch 32767 {
            case 4294967295:
                switch v2 {
                case "\n\t":
                    var v60 int8 = 127
                    _ = v60
                case "hello":
                    v28 += (v28 + v28)
                case "a":
                    if true {
                        for fi61 := 0; fi61 < 2; fi61++ {
                            for fi62 := 0; fi62 < 6; fi62++ {
                                if false {
                                    v2 = ((v2 + v2) + ("" + v2))
                                    if (0 >= 896) {
                                        var v63 int8 = -127
                                        var v64 uint8 = 0
                                        v4 = v4
                                        var v65 float64 = -2.25
                                        var v66 uint8 = v64
                                        _ = v63
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                    } else {
                                        switch v2 {
                                        case "foo bar":
                                            switch "\n\t" {
                                            case "":
                                                var v67 int = genFn0()
                                                _ = v67
                                            case "\n\t":
                                                v68 := GenS4{F0: fi62, F1: 1, F2: -583}
                                                _ = v68
                                            case "foo bar":
                                                v69 := v3[v2]
                                                _ = v69
                                            }
                                        case "hello":
                                            fmt.Println(v5)
                                        default:
                                            fmt.Println(v6)
                                            if true {
                                                v70 := []int{}
                                                if true {
                                                    if ((v28 >= v28) || true) {
                                                        v5 = float32(1.5)
                                                    } else {
                                                        fmt.Println(v4, v5, v4)
                                                        v28 = (v28 - 127)
                                                        if false {
                                                            var v71 string = "a"
                                                            v72 := GenS6{F0: 0, F1: 255, F2: 1}
                                                            v2 = "\n\t"
                                                            _ = v71
                                                            _ = v72
                                                        } else {
                                                            v28 = v28
                                                            if len(v70) >= 2 {
                                                                v70 = v70[1:len(v70)]
                                                            }
                                                        }
                                                    }
                                                }
                                                var v73 int64 = 0
                                                _ = v70
                                                _ = v73
                                            } else {
                                                v74 := GenS7{F0: (fi61 % ((204) | 1)), F1: 0, F2: 1}
                                                _ = v74
                                            }
                                        }
                                    }
                                } else {
                                    v4 += v4
                                    v75 := GenS4{F0: fi61, F1: 151, F2: 877}
                                    _ = v75
                                }
                                _ = fi62
                            }
                            _ = fi61
                        }
                        v6 = (v6 + (v6 >> ((0) & 63)))
                        for fi76 := 0; fi76 < 18; fi76++ {
                            v2 = ((v2 + v2) + ("" + v2))
                            _ = fi76
                        }
                    } else {
                        v6 -= (v6 | (^(v6)))
                        var v77 uint64 = v6
                        switch -923 {
                        case 9223372036854775807:
                            v78 := GenS8{F0: 1, F1: true}
                            _ = v78
                        case 702:
                            for fi79 := 0; fi79 < 10; fi79++ {
                                v77 -= 42
                                if (42 >= v6) {
                                    var v80 uint8 = 127
                                    fmt.Println(v80)
                                    var v81 uint8 = (v80 << ((729) & 63))
                                    _ = v80
                                    _ = v81
                                } else {
                                    switch int(v6) {
                                    case 1:
                                        v5 = ((3.14159 * 1e10) * v5)
                                    case -9223372036854775808:
                                        v4 = (v4 | v4)
                                    default:
                                        var v82 int64 = 0
                                        _ = v82
                                    }
                                }
                                if false {
                                    fmt.Println(v5, fi79, v6)
                                    switch (fi79 & (^(fi79))) {
                                    case -687:
                                        for fi83 := 0; fi83 < 18; fi83++ {
                                            v84 := []int{-501, (fi79 + (fi79 - (fi83 * (fi83 * (-(fi79)))))), (fi79 + fi83), (fi79 ^ fi83)}
                                            switch fi79 {
                                            case 128:
                                                switch fi79 {
                                                case 0:
                                                    fmt.Println(fi83)
                                                case -2147483648:
                                                    if ("\n\t" < v2) {
                                                        v6 = v77
                                                        fmt.Println(v6, v77)
                                                        var v85 string = ("\n\t" + v2)
                                                        switch v2 {
                                                        case "hello":
                                                            v28 = (v28 << ((188) & 63))
                                                        case "\n\t":
                                                            if (v5 != ((float32(-2.25) + v5) - 3.14159)) {
                                                                v5 += 0.1
                                                                v86 := GenS9{F0: 1, F1: 753, F2: v28}
                                                                _ = len(v84)
                                                                _ = v86
                                                            } else {
                                                                v87 := GenS10{F0: (fi79 % ((32768) | 1)), F1: (v6 ^ (v6 - 777))}
                                                                var v88 int8 = v28
                                                                switch (fi79 >> ((42) & 63)) {
                                                                case 101:
                                                                    v89 := GenS7{F0: 2147483647, F1: 2, F2: -494}
                                                                    v90 := GenS1{F0: genFn0(), F1: (-(v4)), F2: -1}
                                                                    v91 := GenS9{F0: 42, F1: -17, F2: (v88 / ((v28) | 1))}
                                                                    _ = v89
                                                                    _ = v90
                                                                    _ = v91
                                                                case -9223372036854775808:
                                                                    v92 := GenS11{F0: (!(2 >= 42)), F1: 451, F2: 32767}
                                                                    _ = v92
                                                                case 96:
                                                                    if (v77 == 0) {
                                                                        switch v85 {
                                                                        case "a":
                                                                            var v93 bool = true
                                                                            _ = v93
                                                                        case "":
                                                                            if (1000 == -1) {
                                                                                fmt.Println(fi79, v5)
                                                                            } else {
                                                                                v77 = 2
                                                                            }
                                                                        case "hello":
                                                                            v88 -= (v28 >> ((65535) & 63))
                                                                        default:
                                                                            if false {
                                                                                v94 := GenS10{F0: (fi79 / (((fi79 - (^(fi83)))) | 1)), F1: v77}
                                                                                v88 -= v88
                                                                                switch fi83 {
                                                                                case 743:
                                                                                    if false {
                                                                                        fmt.Println(v28, v2, v6)
                                                                                    } else {
                                                                                        v28 = (v88 ^ 1)
                                                                                        v84 = append(v84, fi79)
                                                                                        v85 = v85
                                                                                        var v95 bool = true
                                                                                        if len(v84) >= 2 {
                                                                                            v84 = v84[1:len(v84)]
                                                                                        }
                                                                                        _ = v95
                                                                                    }
                                                                                case -9223372036854775808:
                                                                                    var v96 string = v85
                                                                                    _ = v96
                                                                                case 9223372036854775807:
                                                                                    v4 -= int32(v88)
                                                                                default:
                                                                                    if len(v84) > 0 {
                                                                                        v84[0] = 463
                                                                                    }
                                                                                }
                                                                                _ = v94
                                                                            }
                                                                        }
                                                                        var v97 float64 = 0.1
                                                                        var v98 int32 = v4
                                                                        var v99 int32 = (v4 << ((39) & 63))
                                                                        fmt.Println(v6, v85, v2)
                                                                        v100 := GenS8{F0: 1, F1: (!((1000 >= -1) || (v97 < (v97 - (v97 - v97)))))}
                                                                        if (true && (!(293 > (v98 + (^(v98)))))) {
                                                                            _ = len(v3)
                                                                        }
                                                                        v101 := GenS12{F0: 42}
                                                                        if len(v84) >= 2 {
                                                                            v84 = v84[1:len(v84)]
                                                                        }
                                                                        _ = v100
                                                                        _ = v101
                                                                        _ = v97
                                                                        _ = v98
                                                                        _ = v99
                                                                    } else {
                                                                        if true {
                                                                            v102 := GenS5{F0: (^(v4)), F1: v2, F2: 0}
                                                                            _ = v102
                                                                        }
                                                                        var v103 float64 = 1.5
                                                                        if false {
                                                                            var v104 uint8 = 71
                                                                            _ = v104
                                                                        }
                                                                        if false {
                                                                            switch (fi83 & (fi79 / ((4294967295) | 1))) {
                                                                            case 245:
                                                                                v105 := GenS13{F0: 1}
                                                                                _ = v105
                                                                            }
                                                                            var v106 uint16 = 585
                                                                            var v107 string = v85
                                                                            var v108 int8 = 0
                                                                            if false {
                                                                                switch fi79 {
                                                                                case 4294967295:
                                                                                    var v109 string = v2
                                                                                    _ = v109
                                                                                }
                                                                            }
                                                                            v110 := GenS7{F0: genFn1(127), F1: 2, F2: 1}
                                                                            _ = v106
                                                                            _ = v107
                                                                            _ = v108
                                                                            _ = v110
                                                                        }
                                                                        _ = v103
                                                                    }
                                                                default:
                                                                    v111 := GenS14{F0: 2147483648, F1: 801}
                                                                    _ = v111
                                                                }
                                                                _ = v87
                                                                _ = v88
                                                            }
                                                        }
                                                        fmt.Println(v4, fi83, v28)
                                                        var v112 int64 = 237
                                                        var v113 int8 = 42
                                                        _ = v112
                                                        _ = v113
                                                        _ = v85
                                                    } else {
                                                        var v114 uint = 1
                                                        var v115 int16 = 42
                                                        var v116 int8 = -92
                                                        v5 += (v5 + v5)
                                                        _ = v114
                                                        _ = v115
                                                        _ = v116
                                                    }
                                                default:
                                                    v4 += (v4 ^ v4)
                                                }
                                            case -478:
                                                v2 = v2
                                            case 127:
                                                if (!("a" <= ((v2 + v2) + v2))) {
                                                    var v117 int32 = 251
                                                    _ = v117
                                                } else {
                                                    v6 = v77
                                                    if ((float64(3.14159) != 1.5) && false) {
                                                        v118 := v1[fi79]
                                                        var v119 int = fi83
                                                        if v118 {
                                                            switch (v2 + "hello") {
                                                            case "a":
                                                                var v120 uint8 = 128
                                                                if false {
                                                                    switch (fi83 | (fi83 - genFn0())) {
                                                                    case -807:
                                                                        v119 = (fi79 & (-(v119)))
                                                                    case -485:
                                                                        v121 := GenS15{F0: 665, F1: (fi79 - (fi79 ^ int(v120)))}
                                                                        _ = v121
                                                                    }
                                                                    v122 := GenS7{F0: -9223372036854775808, F1: 65535, F2: 1}
                                                                    v123 := GenS16{F0: 42}
                                                                    fmt.Println(v5, v5)
                                                                    if false {
                                                                        switch v2 {
                                                                        case "a":
                                                                            var v124 bool = (v118 || true)
                                                                            _ = v124
                                                                        case "hello":
                                                                            fmt.Println(v77, v28)
                                                                        }
                                                                        var v125 uint16 = 2
                                                                        _ = v125
                                                                    } else {
                                                                        if (v28 == (v28 >> ((1) & 63))) {
                                                                            var v126 uint64 = 2
                                                                            if (v119 != v119) {
                                                                                v4 = (v4 + 0)
                                                                            }
                                                                            _ = v126
                                                                        } else {
                                                                            if (213 < 2) {
                                                                                var v127 bool = false
                                                                                _ = v127
                                                                            } else {
                                                                                if (!(true || (!(0 <= 42)))) {
                                                                                    v120 = v120
                                                                                }
                                                                                v128 := GenS7{F0: 128, F1: uint16(v77), F2: 659}
                                                                                _ = v128
                                                                            }
                                                                            fmt.Println(v120, fi79, fi83)
                                                                            var v129 uint64 = 2
                                                                            _ = v129
                                                                        }
                                                                    }
                                                                    var v130 float64 = 0.1
                                                                    if (1e10 > v130) {
                                                                        switch (-(fi83)) {
                                                                        case 2147483647:
                                                                            fmt.Println(fi83, fi79, v28)
                                                                        default:
                                                                            var v131 uint64 = (v6 * 42)
                                                                            v6 += v6
                                                                            _ = v131
                                                                        }
                                                                        if v118 {
                                                                            var v132 uint32 = 883
                                                                            _ = v132
                                                                        } else {
                                                                            fmt.Println(v5, v119, fi83)
                                                                        }
                                                                        switch v119 {
                                                                        case -9223372036854775808:
                                                                            var v133 float64 = (v130 - v130)
                                                                            _ = v133
                                                                        case 474:
                                                                            _ = len(v3)
                                                                            v120 += (v120 & (v120 ^ v120))
                                                                        }
                                                                    } else {
                                                                        v3[v2] = "\n\t"
                                                                    }
                                                                    v118 = false
                                                                    _ = v122
                                                                    _ = v123
                                                                    _ = v130
                                                                }
                                                                var v134 uint8 = (v120 >> ((2) & 63))
                                                                _ = v120
                                                                _ = v134
                                                            case "\n\t":
                                                                if false {
                                                                    fmt.Println(fi79, v4, v2)
                                                                    if ((1 >= -1) || v118) {
                                                                        fmt.Println(v5, fi83)
                                                                        var v135 int = fi83
                                                                        _ = v135
                                                                    } else {
                                                                        v136 := GenS15{F0: 4294967295, F1: 804}
                                                                        if (-553 >= 681) {
                                                                            v137 := v3[v2]
                                                                            _ = v137
                                                                        }
                                                                        var v138 uint64 = (^(v77))
                                                                        v139 := GenS17{F0: 2}
                                                                        var v140 int8 = v28
                                                                        _ = v136
                                                                        _ = v138
                                                                        _ = v139
                                                                        _ = v140
                                                                    }
                                                                } else {
                                                                    var v141 float64 = (-2.25 + ((float64(0.1) + -2.25) + (3.14159 * 0.1)))
                                                                    _ = v141
                                                                }
                                                            case "hello":
                                                                fmt.Println(v118, v119)
                                                            }
                                                        }
                                                        _ = v118
                                                        _ = v119
                                                    } else {
                                                        v142 := GenS16{F0: -1}
                                                        v2 = ("foo bar" + v2)
                                                        _ = v142
                                                    }
                                                }
                                            }
                                            var v143 int8 = (v28 | 0)
                                            _ = fi83
                                            _ = v143
                                            _ = v84
                                        }
                                    }
                                    var v144 uint32 = 714
                                    _ = v144
                                }
                                _ = fi79
                            }
                        default:
                            v77 = v6
                        }
                        _ = v77
                    }
                }
            default:
                v4 = v4
            }
        }
    }
    var v145 uint8 = 168
    if (!true) {
        v146 := GenS4{F0: -532, F1: 1, F2: 42}
        _ = v146
    } else {
        v28 = v28
        fmt.Println(v5, v6, v6)
        v6 += 42
    }
    for fi147 := 0; fi147 < 13; fi147++ {
        if ((v6 <= v6) || false) {
            var v148 int = fi147
            _ = v148
        }
        v149 := GenS6{F0: uint16(v145), F1: 1, F2: 32767}
        _ = fi147
        _ = v149
    }
    if (v28 != v28) {
        var v150 float64 = 0.1
        var v151 int64 = 1
        _ = v150
        _ = v151
    }
    switch ((v2 + (v2 + "hello")) + "") {
    case "hello":
        var v152 int = 851
        _ = v152
    case "":
        var v153 int32 = 0
        _ = v153
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v145:", v145)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v145
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
