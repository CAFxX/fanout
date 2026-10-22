package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 uint8; F2 uint16 }
type GenS4 struct { F0 bool; F1 string; F2 int16 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 string }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 int16; F1 int8 }
type GenS15 struct { F0 uint64; F1 uint8; F2 uint32 }
func genFn0(p0 uint64, p1 uint16) int16 {
    var v6 int8 = 100
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint8 = 221
        v6 = (v6 * v6)
        if true {
            var v9 uint = 238
            var v10 float32 = 3.14159
            v6 = 33
            _ = v10
            _ = v9
        } else {
            var v11 float64 = 1.5
            if (true && ((v8 ^ v8) > v8)) {
                for fi12 := 0; fi12 < 4; fi12++ {
                    v13 := GenS1{F0: v8}
                    _ = fi12
                    _ = v13
                }
            }
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(p1)
    v14 := []int{128, 128, -309, 509}
    _ = p0
    _ = p1
    _ = v14
    _ = v6
    return 42
}
func genFn1() uint16 {
    if true {
        for fi15 := 0; fi15 < 11; fi15++ {
            if (!true) {
                var v16 bool = true
                var v17 float64 = (float64(1.5) + (-2.25 - 1.5))
                fmt.Println(v16, v17, v17)
                _ = v16
                _ = v17
            } else {
                if (true || true) {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        switch "" {
                        case "hello":
                            var v19 uint = 42
                            _ = v19
                        }
                        _ = fi18
                    }
                    if true {
                        var v20 uint64 = 506
                        _ = v20
                    } else {
                        for fi21 := 0; fi21 < 1; fi21++ {
                            var v22 int64 = 0
                            _ = fi21
                            _ = v22
                        }
                        var v23 uint = 936
                        v23 += 0
                        _ = v23
                    }
                    var v24 int32 = -2147483648
                    _ = v24
                }
                fmt.Println(fi15)
                var v25 int32 = 0
                _ = v25
            }
            var v26 uint16 = 1
            _ = fi15
            _ = v26
        }
    }
    var v27 int32 = -542
    switch "\n\t" {
    case "a":
        v28 := []int{128, 128, 1}
        _ = v28
    case "foo bar":
        if (0 == 127) {
            v27 += (-(v27))
            v29 := []int{2147483647, 704, 1, 9223372036854775807}
            _ = v29
        }
    default:
        v27 -= (v27 / ((v27) | 1))
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        for fi31 := 0; fi31 < 8; fi31++ {
            fmt.Println(fi30, v27, fi31)
            v32 := []int{fi30, (fi31 >> ((4294967295) & 63))}
            v27 = 2147483647
            var v33 float32 = float32(((0.0 - 0.1) - float64(0.0)))
            var v34 uint64 = 42
            fmt.Println(v27, fi30, fi30)
            var v35 uint16 = 2
            v35 = v35
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
            _ = v35
        }
        if ((((3.14159 / (-2.25 / (-2.25 / float32((0.0 - ((3.14159 / 3.14159) - 3.14159)))))) * -2.25) != (0.1 + (0.1 * ((((3.14159 * (((-2.25 - 0.0) - 1e10) * 0.1)) - 0.0) / ((1.5 - 0.0) * float32(1.5))) + 0.0)))) || ((^(v27)) == v27)) {
            for fi36 := 0; fi36 < 12; fi36++ {
                switch fi30 {
                case 0:
                    if (fi36 < -871) {
                        fmt.Println(fi30)
                    }
                case 32767:
                    switch (fi36 | (fi36 >> ((21) & 63))) {
                    case 42:
                        switch (fi30 & 128) {
                        case 32768:
                            v37 := GenS1{F0: 124}
                            _ = v37
                        case 839:
                            v38 := GenS2{F0: 2}
                            _ = v38
                        default:
                            fmt.Println(fi36)
                        }
                    case 0:
                        var v39 float32 = (-2.25 - -2.25)
                        _ = v39
                    case 2147483647:
                        v27 = (v27 / ((v27) | 1))
                    }
                    if (((42 <= 18446744073709551615) || false) && false) {
                        v27 += 42
                        switch 0 {
                        case 127:
                            if true {
                                var v40 string = ""
                                _ = v40
                            } else {
                                if (true && false) {
                                    v27 = (v27 ^ v27)
                                }
                                fmt.Println(fi30, fi30)
                                var v41 uint64 = 1
                                _ = v41
                            }
                        }
                    } else {
                        v27 = 42
                        var v42 uint32 = 2
                        fmt.Println(v27, v27, v42)
                        fmt.Println(fi30)
                        _ = v42
                    }
                    var v43 bool = false
                    _ = v43
                case 631:
                    var v44 float32 = float32(1e10)
                    _ = v44
                }
                if ((v27 != -2147483648) && true) {
                    if true {
                        var v45 int32 = -2147483648
                        switch 32768 {
                        case 128:
                            var v46 int16 = 42
                            _ = v46
                        default:
                            if false {
                                var v47 uint64 = 2
                                if true {
                                    v27 = 1
                                    if (false || true) {
                                        v48 := GenS3{F0: ("" + "a"), F1: 255, F2: 684}
                                        v27 = v45
                                        if ((v47 != 1) && (9223372036854775807 <= -1)) {
                                            var v49 float64 = float64((3.14159 / 1.5))
                                            _ = v49
                                        }
                                        _ = v48
                                    }
                                } else {
                                    switch "foo bar" {
                                    case "hello":
                                        switch 0 {
                                        case 9223372036854775807:
                                            fmt.Println(v45)
                                            v50 := []int{(fi30 ^ fi36)}
                                            fmt.Println(fi36, fi30)
                                            _ = v50
                                        case 183:
                                            v51 := GenS2{F0: 128}
                                            var v52 int = fi30
                                            v53 := GenS1{F0: 31}
                                            _ = v51
                                            _ = v52
                                            _ = v53
                                        }
                                    case "\n\t":
                                        var v54 string = "hello"
                                        v45 = (-(v45))
                                        if (!((v45 % ((-2147483648) | 1)) == v27)) {
                                            var v55 uint = 4294967295
                                            if (65535 != 1) {
                                                var v56 uint64 = 4294967295
                                                _ = v56
                                            } else {
                                                var v57 uint32 = 147
                                                v58 := GenS4{F0: (!(v45 <= v45)), F1: "a", F2: int16(v47)}
                                                _ = v57
                                                _ = v58
                                            }
                                            _ = v55
                                        } else {
                                            var v59 uint8 = 144
                                            v54 = (("hello" + "foo bar") + ("hello" + "a"))
                                            _ = v59
                                        }
                                        _ = v54
                                    case "a":
                                        switch (fi36 >> ((18446744073709551615) & 63)) {
                                        case -2147483648:
                                            v60 := GenS1{F0: 255}
                                            _ = v60
                                        case 791:
                                            if true {
                                                v61 := []int{}
                                                _ = v61
                                            } else {
                                                v27 = 1
                                                v62 := GenS3{F0: "", F1: 255, F2: 517}
                                                _ = v62
                                            }
                                        case 127:
                                            v63 := []int{-2147483648, -79, 32768, fi36}
                                            _ = v63
                                        default:
                                            if ((-(fi36)) != -9223372036854775808) {
                                                fmt.Println(fi30, fi36, v27)
                                            } else {
                                                if ("hello" < "foo bar") {
                                                    var v64 uint8 = 255
                                                    _ = v64
                                                } else {
                                                    if false {
                                                        switch fi36 {
                                                        case -262:
                                                            v45 += v45
                                                        case 246:
                                                            switch fi30 {
                                                            case 450:
                                                                fmt.Println(v45, v47, fi36)
                                                            case 16:
                                                                switch (fi30 ^ (fi36 - (fi30 + fi30))) {
                                                                case 191:
                                                                    v27 = (v27 * v27)
                                                                case 127:
                                                                    if (42 < v45) {
                                                                        v45 += (-(v27))
                                                                        if (4294967295 > 518) {
                                                                            v47 = (^(v47))
                                                                            if (206 == 217) {
                                                                                v65 := []int{(fi30 / ((fi30) | 1)), (fi36 % ((32768) | 1)), fi30}
                                                                                switch fi30 {
                                                                                case 4294967295:
                                                                                    v65 = append(v65, fi30)
                                                                                case 127:
                                                                                    fmt.Println(v47)
                                                                                default:
                                                                                    if true {
                                                                                        v66 := GenS1{F0: 1}
                                                                                        var v67 bool = (!false)
                                                                                        v65 = append(v65, (-(fi36)))
                                                                                        if true {
                                                                                            switch "foo bar" {
                                                                                            case "":
                                                                                                v67 = v67
                                                                                            }
                                                                                        }
                                                                                        switch -1 {
                                                                                        case 127:
                                                                                            fmt.Println(fi36, v47, fi30)
                                                                                        case 1:
                                                                                            v27 -= v27
                                                                                        case 85:
                                                                                            var v68 uint = 18446744073709551615
                                                                                            _ = v68
                                                                                        default:
                                                                                            v45 += v27
                                                                                        }
                                                                                        if (((!((false && v67) || (-1 == -1))) && (v67 || v67)) || v67) {
                                                                                            switch "" {
                                                                                            case "":
                                                                                                fmt.Println(v67)
                                                                                            case "\n\t":
                                                                                                if v67 {
                                                                                                    fmt.Println(v47, fi30)
                                                                                                    v45 += (v45 * v27)
                                                                                                }
                                                                                            case "hello":
                                                                                                var v69 float32 = 1.5
                                                                                                _ = v69
                                                                                            }
                                                                                            v27 = 0
                                                                                        } else {
                                                                                            v70 := GenS3{F0: "hello", F1: 255, F2: 65535}
                                                                                            v67 = (fi36 < 9223372036854775807)
                                                                                            if v67 {
                                                                                                switch -510 {
                                                                                                case 369:
                                                                                                    var v71 uint64 = (v47 << ((2) & 63))
                                                                                                    _ = v71
                                                                                                default:
                                                                                                    fmt.Println(fi30, v47, fi30)
                                                                                                }
                                                                                                switch "a" {
                                                                                                case "hello":
                                                                                                    switch (("foo bar" + "a") + ("\n\t" + "hello")) {
                                                                                                    case "a":
                                                                                                        if (((0.1 + 1e10) <= (1.5 * ((0.0 / 1.0) + -2.25))) && false) {
                                                                                                            var v72 int = 2147483647
                                                                                                            v67 = (2 != 291)
                                                                                                            _ = v72
                                                                                                        } else {
                                                                                                            fmt.Println(fi30, fi36)
                                                                                                            var v73 int64 = -1
                                                                                                            _ = v73
                                                                                                        }
                                                                                                    case "":
                                                                                                        var v74 float32 = 1.5
                                                                                                        _ = v74
                                                                                                    case "\n\t":
                                                                                                        var v75 int16 = 42
                                                                                                        _ = v75
                                                                                                    default:
                                                                                                        var v76 int64 = -380
                                                                                                        _ = v76
                                                                                                    }
                                                                                                    if v67 {
                                                                                                        v77 := GenS2{F0: 65}
                                                                                                        var v78 uint32 = 4294967295
                                                                                                        var v79 int64 = -527
                                                                                                        _ = v77
                                                                                                        _ = v78
                                                                                                        _ = v79
                                                                                                    } else {
                                                                                                        v27 = (^(v45))
                                                                                                        var v80 int16 = -1
                                                                                                        var v81 int = 127
                                                                                                        _ = v80
                                                                                                        _ = v81
                                                                                                    }
                                                                                                    fmt.Println(v45, v45, fi30)
                                                                                                case "a":
                                                                                                    fmt.Println(fi36, fi30, fi36)
                                                                                                case "foo bar":
                                                                                                    var v82 bool = false
                                                                                                    _ = v82
                                                                                                }
                                                                                                var v83 int8 = -128
                                                                                                switch (-(fi30)) {
                                                                                                case 358:
                                                                                                    if (!(!(32767 < 32767))) {
                                                                                                        fmt.Println(v83, v67, v83)
                                                                                                        switch -151 {
                                                                                                        case 4294967295:
                                                                                                            var v84 uint8 = 0
                                                                                                            _ = v84
                                                                                                        case 32767:
                                                                                                            var v85 float64 = ((1e10 + 1.5) + 1.5)
                                                                                                            var v86 float32 = -2.25
                                                                                                            v85 = float64(0.1)
                                                                                                            switch (-(fi30)) {
                                                                                                            case 713:
                                                                                                                switch ("hello" + ("" + ("a" + ""))) {
                                                                                                                case "foo bar":
                                                                                                                    switch "a" {
                                                                                                                    case "a":
                                                                                                                        if v67 {
                                                                                                                            v45 -= -2147483648
                                                                                                                            v27 = 183
                                                                                                                            v87 := GenS4{F0: true, F1: ("a" + "a"), F2: 0}
                                                                                                                            v65 = append(v65, (fi36 ^ (fi30 % ((int(v83)) | 1))))
                                                                                                                            _ = v87
                                                                                                                        }
                                                                                                                    case "\n\t":
                                                                                                                        v47 -= v47
                                                                                                                    case "":
                                                                                                                        var v88 int64 = 1
                                                                                                                        _ = v88
                                                                                                                    default:
                                                                                                                        if v67 {
                                                                                                                            v85 = float64(3.14159)
                                                                                                                            if (1 != 42) {
                                                                                                                                var v89 int8 = v83
                                                                                                                                _ = v89
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            v90 := GenS3{F0: ("" + "foo bar"), F1: 42, F2: 32767}
                                                                                                                            var v91 float64 = 3.14159
                                                                                                                            if (65535 == 2) {
                                                                                                                                v83 += (v83 - (v83 % (((v83 / (((v83 ^ (v83 / ((0) | 1)))) | 1))) | 1)))
                                                                                                                                var v92 float64 = v85
                                                                                                                                v85 += 0.0
                                                                                                                                var v93 int16 = genFn0((v47 % ((v47) | 1)), 0)
                                                                                                                                _ = v92
                                                                                                                                _ = v93
                                                                                                                            }
                                                                                                                            _ = v90
                                                                                                                            _ = v91
                                                                                                                        }
                                                                                                                    }
                                                                                                                default:
                                                                                                                    if v67 {
                                                                                                                        v94 := GenS1{F0: 255}
                                                                                                                        _ = v94
                                                                                                                    } else {
                                                                                                                        var v95 int32 = (v27 & -1)
                                                                                                                        v65 = append(v65, -306)
                                                                                                                        _ = v95
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                v47 = 18446744073709551615
                                                                                                            }
                                                                                                            switch "" {
                                                                                                            case "hello":
                                                                                                                var v96 bool = (0 < 65535)
                                                                                                                _ = v96
                                                                                                            }
                                                                                                            _ = v85
                                                                                                            _ = v86
                                                                                                        default:
                                                                                                            v83 = (v83 | (v83 - v83))
                                                                                                        }
                                                                                                    } else {
                                                                                                        v97 := GenS3{F0: ("hello" + ""), F1: 1, F2: 2}
                                                                                                        v65 = append(v65, fi30)
                                                                                                        if v67 {
                                                                                                            switch 1 {
                                                                                                            case 9223372036854775807:
                                                                                                                v65 = append(v65, fi30)
                                                                                                            case 32767:
                                                                                                                v83 += (v83 << ((65535) & 63))
                                                                                                            default:
                                                                                                                v98 := GenS4{F0: v67, F1: ("hello" + "foo bar"), F2: 0}
                                                                                                                _ = v98
                                                                                                            }
                                                                                                            v67 = (18446744073709551615 > v47)
                                                                                                        } else {
                                                                                                            switch (-(fi30)) {
                                                                                                            case 892:
                                                                                                                v83 -= (v83 + v83)
                                                                                                            case 2147483647:
                                                                                                                if len(v65) > 0 {
                                                                                                                    v65[0] = (fi36 + fi30)
                                                                                                                }
                                                                                                            }
                                                                                                            if len(v65) >= 2 {
                                                                                                                v65 = v65[1:len(v65)]
                                                                                                            }
                                                                                                            if true {
                                                                                                                v83 += v83
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v97
                                                                                                    }
                                                                                                case 42:
                                                                                                    v27 = (-(v45))
                                                                                                }
                                                                                                v99 := GenS5{F0: -1, F1: 128}
                                                                                                _ = v83
                                                                                                _ = v99
                                                                                            } else {
                                                                                                v100 := GenS6{F0: 1}
                                                                                                var v101 int16 = 438
                                                                                                var v102 string = "a"
                                                                                                _ = v100
                                                                                                _ = v101
                                                                                                _ = v102
                                                                                            }
                                                                                            v45 -= v45
                                                                                            v67 = (1 >= -1)
                                                                                            _ = v70
                                                                                        }
                                                                                        _ = v66
                                                                                        _ = v67
                                                                                    }
                                                                                }
                                                                                v45 -= v27
                                                                                v103 := GenS6{F0: -9223372036854775808}
                                                                                switch fi30 {
                                                                                case -2147483648:
                                                                                    v104 := GenS7{F0: 0}
                                                                                    _ = v104
                                                                                case 1:
                                                                                    v65 = append(v65, fi36)
                                                                                    v45 = (-(v45))
                                                                                default:
                                                                                    v47 = 4294967295
                                                                                }
                                                                                _ = v103
                                                                                _ = v65
                                                                            }
                                                                            if (!((-179 != 1) || true)) {
                                                                                switch fi36 {
                                                                                case 32767:
                                                                                    if false {
                                                                                        if ((((((3.14159 * (0.0 / float64(float32(((0.0 + (3.14159 - ((1e10 / 1.0) + (1e10 + 1e10)))) - 1e10))))) * 0.1) - 1.5) - ((3.14159 + (((1.5 / float64(0.1)) + (0.0 * (0.0 - 3.14159))) - 3.14159)) + 1.5)) > 0.0) || false) {
                                                                                            v105 := []int{(fi30 << ((65535) & 63)), (fi30 - fi36), 4294967295}
                                                                                            _ = v105
                                                                                        } else {
                                                                                            var v106 uint64 = 2
                                                                                            fmt.Println(fi30, v45)
                                                                                            var v107 bool = false
                                                                                            _ = v106
                                                                                            _ = v107
                                                                                        }
                                                                                        v108 := []int{(fi36 + fi36), (^(fi36))}
                                                                                        _ = v108
                                                                                    } else {
                                                                                        v109 := GenS8{F0: 65535}
                                                                                        _ = v109
                                                                                    }
                                                                                }
                                                                                if (-32768 > 1) {
                                                                                    var v110 float32 = (1e10 - 3.14159)
                                                                                    v111 := []int{(fi30 >> ((2) & 63)), 1, (-(fi36))}
                                                                                    if (v110 == v110) {
                                                                                        v112 := GenS8{F0: 0}
                                                                                        _ = v112
                                                                                    }
                                                                                    _ = v110
                                                                                    _ = v111
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    if false {
                                                                        v113 := GenS8{F0: 0}
                                                                        _ = v113
                                                                    } else {
                                                                        v114 := []int{fi30}
                                                                        v115 := GenS9{F0: 0}
                                                                        _ = v114
                                                                        _ = v115
                                                                    }
                                                                }
                                                            case 127:
                                                                fmt.Println(v47, fi36, v45)
                                                            default:
                                                                v116 := GenS10{F0: 949}
                                                                _ = v116
                                                            }
                                                        case 9223372036854775807:
                                                            var v117 int32 = 0
                                                            if (0.1 >= -2.25) {
                                                                v118 := []int{int(v45), (fi30 + fi36), (fi30 - -606), (fi30 / (((fi30 - (fi36 / (((fi36 * fi30)) | 1)))) | 1))}
                                                                var v119 int8 = int8(fi30)
                                                                _ = v118
                                                                _ = v119
                                                            }
                                                            _ = v117
                                                        }
                                                        v27 -= v27
                                                    }
                                                    var v120 float32 = ((3.14159 / -2.25) + (((1e10 + 1.5) / -2.25) / 1.5))
                                                    v47 += v47
                                                    _ = v120
                                                }
                                            }
                                        }
                                    }
                                    fmt.Println(v47, v47, fi36)
                                }
                                v121 := GenS11{F0: false}
                                _ = v121
                                _ = v47
                            }
                        }
                        _ = v45
                    }
                    var v122 uint32 = 409
                    if (false && true) {
                        var v123 bool = true
                        v27 = v27
                        v124 := []int{(fi30 - 42)}
                        fmt.Println(fi30, fi36, v27)
                        _ = v123
                        _ = v124
                    }
                    _ = v122
                }
                _ = fi36
            }
            var v125 float32 = 0.0
            _ = v125
        }
        _ = fi30
    }
    if (false && (-48 > -46)) {
        v126 := GenS8{F0: 2147483648}
        _ = v126
    } else {
        var v127 string = "foo bar"
        _ = v127
    }
    _ = v27
    return 653
}

func main() {
    var v1 uint = 65535
    v2 := []int{}
    v3 := []string{(("\n\t" + (("a" + "foo bar") + ("a" + "\n\t"))) + ("" + ("" + ""))), "a", ("hello" + "foo bar"), ("hello" + ("a" + "\n\t"))}
    var v4 int64 = 0
    var v5 uint8 = 0
    switch 42 {
    case 508:
        if (3.14159 <= float32(3.14159)) {
            if false {
                fmt.Println(v5, v4)
                _ = len(v2)
            } else {
                for fi128 := 0; fi128 < 1; fi128++ {
                    var v129 int16 = 732
                    _ = fi128
                    _ = v129
                }
                v130 := GenS12{F0: (("a" + "a") + (((("\n\t" + "a") + "foo bar") + "a") + "hello"))}
                _ = v130
            }
            if false {
                fmt.Println(v1)
                if (true || true) {
                    var v131 int16 = genFn0(uint64(v1), 199)
                    var v132 uint = (v1 >> ((v1) & 63))
                    _ = v131
                    _ = v132
                } else {
                    v4 = (v4 + v4)
                    var v133 string = ""
                    v3 = append(v3, (("\n\t" + ((v133 + v133) + ("foo bar" + v133))) + v133))
                    _ = v133
                }
                v1 = (v1 - (v1 & v1))
            } else {
                if true {
                    v4 = v4
                } else {
                    var v134 float64 = 0.1
                    _ = len(v2)
                    if true {
                        v4 += v4
                    }
                    _ = v134
                }
            }
        } else {
            v135 := GenS2{F0: 127}
            if (float64((3.14159 * (((0.0 / 1.0) * 1e10) - -2.25))) <= -2.25) {
                for fi136 := 0; fi136 < 2; fi136++ {
                    if false {
                        var v137 uint = uint(v5)
                        if (42 < 42) {
                            var v138 int32 = -2147483648
                            if (!false) {
                                var v139 uint32 = 1
                                _ = v139
                            }
                            var v140 int8 = 19
                            _ = v138
                            _ = v140
                        } else {
                            var v141 int = (fi136 ^ (fi136 & (^(fi136))))
                            if false {
                                if false {
                                    v3 = append(v3, "hello")
                                    var v142 int = 127
                                    _ = v142
                                }
                            } else {
                                v137 += (v137 >> (((v137 >> ((v1) & 63))) & 63))
                                switch 344 {
                                case 4294967295:
                                    var v143 uint16 = 65535
                                    _ = v143
                                default:
                                    if false {
                                        var v144 bool = (127 <= 1)
                                        if v144 {
                                            v141 += v141
                                            v145 := GenS13{F0: v144}
                                            var v146 uint = 30
                                            _ = v145
                                            _ = v146
                                        } else {
                                            switch "hello" {
                                            case "a":
                                                var v147 int8 = -117
                                                _ = v147
                                            case "hello":
                                                if (!false) {
                                                    var v148 float64 = -2.25
                                                    _ = v148
                                                }
                                            case "":
                                                var v149 bool = (-1 >= 1)
                                                _ = v149
                                            }
                                            fmt.Println(v137, fi136)
                                            var v150 int16 = 1000
                                            _ = v150
                                        }
                                        switch (fi136 >> ((v137) & 63)) {
                                        case -1:
                                            if (1 > 42) {
                                                for fi151 := 0; fi151 < 17; fi151++ {
                                                    v4 += (v4 << (((v1 << ((2) & 63))) & 63))
                                                    _ = fi151
                                                }
                                                v5 += (v5 + (v5 | 1))
                                            } else {
                                                v152 := GenS6{F0: v4}
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v153 int64 = v4
                                                _ = v152
                                                _ = v153
                                            }
                                        }
                                        _ = v144
                                    }
                                }
                                if (65535 > (^(v1))) {
                                    if ((v4 & (v4 / ((v4) | 1))) <= (v4 & (v4 + (-(v4))))) {
                                        v3 = append(v3, (("" + ("" + ("foo bar" + ""))) + "a"))
                                        v154 := GenS14{F0: 1, F1: 0}
                                        _ = v154
                                    } else {
                                        v155 := GenS14{F0: genFn0(0, 2), F1: 0}
                                        var v156 float64 = -2.25
                                        _ = v155
                                        _ = v156
                                    }
                                } else {
                                    var v157 string = "\n\t"
                                    _ = v157
                                }
                            }
                            if true {
                                var v158 uint32 = 0
                                _ = v158
                            } else {
                                var v159 int32 = 961
                                for fi160 := 0; fi160 < 4; fi160++ {
                                    var v161 uint64 = 844
                                    fmt.Println(v159)
                                    var v162 int16 = 678
                                    _ = fi160
                                    _ = v161
                                    _ = v162
                                }
                                if len(v3) > 0 {
                                    v3[0] = "a"
                                }
                                _ = v159
                            }
                            _ = v141
                        }
                        _ = v137
                    }
                    _ = fi136
                }
                if (1e10 != -2.25) {
                    v163 := GenS13{F0: ((370 == 406) || (!true))}
                    _ = v163
                }
                switch 32767 {
                case 436:
                    var v164 uint = v1
                    _ = v164
                case 32767:
                    v165 := GenS15{F0: 142, F1: v5, F2: 1}
                    _ = v165
                }
            } else {
                if true {
                    if false {
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                    }
                    v166 := GenS6{F0: (v4 << ((v1) & 63))}
                    _ = v166
                } else {
                    var v167 int = 0
                    for fi168 := 0; fi168 < 17; fi168++ {
                        v5 = v5
                        v3 = append(v3, "")
                        _ = fi168
                    }
                    var v169 uint32 = 0
                    _ = v167
                    _ = v169
                }
                var v170 float32 = 0.1
                v5 += 128
                _ = v170
            }
            v5 -= (v5 - v5)
            _ = v135
        }
    case 9223372036854775807:
        _ = len(v3)
    }
    for fi171 := 0; fi171 < 8; fi171++ {
        fmt.Println(v5)
        _ = fi171
    }
    v172 := GenS8{F0: 2}
    var v173 float32 = 1e10
    fmt.Println(v173)
    v5 -= 0
    for fi174 := 0; fi174 < 17; fi174++ {
        if false {
            if len(v2) > 0 {
                v2[0] = fi174
            }
        }
        switch fi174 {
        case 32768:
            var v175 int64 = int64(v5)
            _ = v175
        default:
            for fi176 := 0; fi176 < 19; fi176++ {
                var v177 string = "\n\t"
                var v178 int64 = 9223372036854775807
                _ = fi176
                _ = v177
                _ = v178
            }
        }
        _ = fi174
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    var v179 uint16 = 0
    if (!false) {
        v2 = append(v2, 497)
        var v180 string = "foo bar"
        _ = v180
    } else {
        var v181 int = 128
        if false {
            switch v181 {
            case -710:
                fmt.Println(v179, v179, v173)
            }
            v182 := GenS13{F0: (!false)}
            _ = len(v2)
            _ = v182
        } else {
            v183 := GenS2{F0: 2}
            var v184 uint = 65535
            _ = v183
            _ = v184
        }
        _ = v181
    }
    _ = len(v3)
    _ = len(v2)
    fmt.Println(v179, v1)
    v185 := GenS5{F0: (v4 | v4), F1: v5}
    _ = len(v2)
    for fi186 := 0; fi186 < 19; fi186++ {
        var v187 int16 = 1000
        _ = fi186
        _ = v187
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v173:", v173)
    fmt.Println("v179:", v179)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v172
    _ = v173
    _ = v179
    _ = v185
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
