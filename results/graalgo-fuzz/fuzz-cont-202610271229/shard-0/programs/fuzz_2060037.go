package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 string; F1 int; F2 uint32 }
type GenS6 struct { F0 uint8; F1 int32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int; F1 bool }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 bool; F1 uint16; F2 int8 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int32; F1 uint16; F2 uint }
type GenS13 struct { F0 int64; F1 int64 }
func genFn0(p0 int32) uint8 {
    for fi5 := 0; fi5 < 1; fi5++ {
        if false {
            var v6 int16 = 42
            _ = v6
        } else {
            for fi7 := 0; fi7 < 11; fi7++ {
                v8 := []int{}
                var v9 int32 = 432
                switch (fi7 << ((4294967295) & 63)) {
                case 128:
                    fmt.Println(fi7)
                case -1:
                    var v10 int64 = 0
                    _ = v10
                default:
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                }
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        _ = fi5
    }
    if ((4294967295 == 714) || false) {
        p0 = p0
        for fi11 := 0; fi11 < 13; fi11++ {
            fmt.Println(p0, fi11)
            if (177 <= 127) {
                v12 := GenS1{F0: 118, F1: fi11, F2: 0}
                _ = v12
            }
            v13 := GenS2{F0: 934, F1: 111, F2: 949}
            _ = fi11
            _ = v13
        }
        v14 := []int{-806, 2147483647, -922}
        _ = v14
    } else {
        for fi15 := 0; fi15 < 14; fi15++ {
            p0 -= (p0 << ((835) & 63))
            _ = fi15
        }
        v16 := []int{}
        _ = v16
    }
    p0 += (p0 / ((p0) | 1))
    _ = p0
    return 31
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    var v3 uint16 = 65535
    v4 := []uint32{42}
    v2[423] = 1
    _ = len(v1)
    if ("" > "") {
        if len(v4) > 0 {
            v4[0] = 1
        }
    } else {
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 int64 = -648
            _ = fi17
            _ = v18
        }
    }
    if (3.14159 < 1.5) {
        for fi19 := 0; fi19 < 19; fi19++ {
            for fi20 := 0; fi20 < 16; fi20++ {
                switch 42 {
                case 127:
                    if (42 <= 0) {
                        var v21 int64 = -610
                        _ = v21
                    }
                default:
                    if false {
                        if true {
                            v22 := GenS2{F0: (v3 + (v3 >> ((0) & 63))), F1: -107, F2: 1}
                            _ = v22
                        } else {
                            var v23 int = (fi20 >> ((65535) & 63))
                            switch 362 {
                            case 32767:
                                var v24 int32 = 243
                                _ = v24
                            case -1:
                                var v25 float64 = 1.5
                                _ = v25
                            case 4294967295:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            default:
                                v23 += (fi20 >> ((65535) & 63))
                            }
                            if true {
                                switch (fi19 / (((fi20 << ((665) & 63))) | 1)) {
                                case -224:
                                    v23 = fi20
                                case -823:
                                    v26 := GenS1{F0: 1, F1: (fi19 << ((255) & 63)), F2: 167}
                                    _ = v26
                                default:
                                    var v27 uint = 42
                                    _ = len(v1)
                                    switch (fi19 * (-(v23))) {
                                    case 943:
                                        switch 4294967295 {
                                        case 728:
                                            v23 = fi20
                                        case 0:
                                            v27 = v27
                                        case 4294967295:
                                            v23 = 0
                                            var v28 uint64 = uint64(v27)
                                            var v29 int8 = 127
                                            _ = v28
                                            _ = v29
                                        }
                                    case 243:
                                        switch (fi19 >> ((666) & 63)) {
                                        case -900:
                                            if false {
                                                v27 = 871
                                                switch ("hello" + ("hello" + ("hello" + ""))) {
                                                case "foo bar":
                                                    v30 := GenS3{F0: (v23 & 0)}
                                                    _ = v30
                                                case "\n\t":
                                                    switch (((("\n\t" + "a") + "") + "") + "foo bar") {
                                                    case "\n\t":
                                                        switch (("a" + (("\n\t" + "") + ("hello" + ("hello" + "foo bar")))) + ("a" + "foo bar")) {
                                                        case "a":
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        case "\n\t":
                                                            v31 := GenS4{F0: -373, F1: "a"}
                                                            _ = v31
                                                        }
                                                    default:
                                                        var v32 uint32 = 555
                                                        _ = v32
                                                    }
                                                }
                                                v33 := GenS1{F0: -237, F1: fi19, F2: 0}
                                                v34 := GenS1{F0: -9223372036854775808, F1: 32767, F2: 18446744073709551615}
                                                v3 = 418
                                                v3 += 2
                                                if false {
                                                    var v35 uint32 = 42
                                                    _ = v35
                                                } else {
                                                    _ = len(v1)
                                                    fmt.Println(fi20, fi19, fi19)
                                                    v23 = (fi19 / (((v23 >> (((v27 % ((v27) | 1))) & 63))) | 1))
                                                }
                                                _ = v33
                                                _ = v34
                                            }
                                            switch ("a" + ("foo bar" + ("a" + "a"))) {
                                            case "hello":
                                                if (!true) {
                                                    fmt.Println(fi19, fi20, v3)
                                                    v27 = v27
                                                    v4 = append(v4, 42)
                                                } else {
                                                    if ((v27 + (v27 % (((^(v27))) | 1))) >= (^(v27))) {
                                                        var v36 float32 = (0.0 / 3.14159)
                                                        _ = v36
                                                    } else {
                                                        if (-2.25 != (0.0 - 3.14159)) {
                                                            fmt.Println(v27, v27, fi20)
                                                            v23 -= -537
                                                            var v37 bool = true
                                                            _ = v37
                                                        } else {
                                                            v23 = (v23 * (v23 / (((fi20 | 32767)) | 1)))
                                                            v27 = 2
                                                            fmt.Println(fi19)
                                                            if (23 != genFn0(-1)) {
                                                                v27 -= 2
                                                                var v38 float32 = 0.0
                                                                _ = v38
                                                            } else {
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                switch 760 {
                                                                case -612:
                                                                    v23 += (fi19 - fi20)
                                                                case 127:
                                                                    if (0 < 930) {
                                                                        if true {
                                                                            _ = len(v2)
                                                                            switch "" {
                                                                            case "":
                                                                                v3 -= (v3 << ((1) & 63))
                                                                            case "a":
                                                                                var v39 int = -2147483648
                                                                                _ = v39
                                                                            case "hello":
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    v27 = 878
                                                                                case "foo bar":
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        v40 := GenS3{F0: v23}
                                                                                        _ = v40
                                                                                    case "hello":
                                                                                        _ = len(v1)
                                                                                    case "":
                                                                                        v3 += uint16(v27)
                                                                                    default:
                                                                                        if (0 <= 42) {
                                                                                            fmt.Println(v27)
                                                                                            v23 -= 725
                                                                                        }
                                                                                        if false {
                                                                                            v41 := GenS4{F0: -729, F1: (("foo bar" + "foo bar") + ("" + "\n\t"))}
                                                                                            _ = v41
                                                                                        }
                                                                                    }
                                                                                }
                                                                            default:
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    if (32767 <= -480) {
                                                                                        switch (fi19 >> ((65535) & 63)) {
                                                                                        case -435:
                                                                                            v4 = append(v4, 4294967295)
                                                                                        }
                                                                                        if (0.0 >= 1e10) {
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = 2147483648
                                                                                            }
                                                                                            _ = len(v4)
                                                                                            v42 := GenS5{F0: "", F1: v23, F2: 2}
                                                                                            v43 := GenS1{F0: -722, F1: v23, F2: 1000}
                                                                                            v27 += (v27 % ((1) | 1))
                                                                                            if true {
                                                                                                var v44 int = (fi20 | fi20)
                                                                                                var v45 bool = (-1 >= -2147483648)
                                                                                                v45 = true
                                                                                                var v46 int8 = -66
                                                                                                _ = v44
                                                                                                _ = v45
                                                                                                _ = v46
                                                                                            }
                                                                                            var v47 uint = 2
                                                                                            if false {
                                                                                                v48 := GenS5{F0: "", F1: v23, F2: uint32(v47)}
                                                                                                v49 := GenS4{F0: 209, F1: ("foo bar" + "a")}
                                                                                                _ = v48
                                                                                                _ = v49
                                                                                            } else {
                                                                                                v3 += (^(v3))
                                                                                                v3 += 0
                                                                                                switch v23 {
                                                                                                case -814:
                                                                                                    v4 = append(v4, 4294967295)
                                                                                                default:
                                                                                                    var v50 float32 = float32(3.14159)
                                                                                                    _ = v50
                                                                                                }
                                                                                                if ((!false) || false) {
                                                                                                    if false {
                                                                                                        switch "foo bar" {
                                                                                                        case "a":
                                                                                                            var v51 string = "\n\t"
                                                                                                            _ = v51
                                                                                                        case "\n\t":
                                                                                                            if (95 != genFn0(0)) {
                                                                                                                v52 := GenS2{F0: (v3 - 1), F1: -1, F2: (v3 & (v3 | v3))}
                                                                                                                _ = v52
                                                                                                            } else {
                                                                                                                delete(v2, (fi20 + fi19))
                                                                                                                var v53 int8 = 100
                                                                                                                _ = v53
                                                                                                            }
                                                                                                        case "hello":
                                                                                                            v1[1] = (0.1 != 1e10)
                                                                                                        default:
                                                                                                            v47 = v27
                                                                                                        }
                                                                                                        var v54 uint32 = 4294967295
                                                                                                        v47 = v47
                                                                                                        v54 += v54
                                                                                                        v54 += (v54 & (v54 + (^(v54))))
                                                                                                        _ = v54
                                                                                                    }
                                                                                                    var v55 string = ""
                                                                                                    v56 := GenS6{F0: 250, F1: 908}
                                                                                                    var v57 uint64 = 453
                                                                                                    _ = v55
                                                                                                    _ = v56
                                                                                                    _ = v57
                                                                                                } else {
                                                                                                    v47 += (v27 >> (((v27 % (((v47 ^ v27)) | 1))) & 63))
                                                                                                }
                                                                                            }
                                                                                            _ = v42
                                                                                            _ = v43
                                                                                            _ = v47
                                                                                        }
                                                                                        if false {
                                                                                            fmt.Println(fi19, v3)
                                                                                            var v58 uint16 = v3
                                                                                            v59 := GenS5{F0: "a", F1: fi20, F2: 65535}
                                                                                            switch fi20 {
                                                                                            case 451:
                                                                                                v58 = 414
                                                                                            default:
                                                                                                _ = len(v1)
                                                                                            }
                                                                                            _ = v58
                                                                                            _ = v59
                                                                                        }
                                                                                        _ = len(v4)
                                                                                        _ = len(v4)
                                                                                    }
                                                                                default:
                                                                                    fmt.Println(v3, fi20, fi20)
                                                                                }
                                                                            }
                                                                            var v60 int8 = 100
                                                                            _ = v60
                                                                        }
                                                                    }
                                                                }
                                                                switch fi20 {
                                                                case 32767:
                                                                    v27 -= (v27 << (((v27 % ((v27) | 1))) & 63))
                                                                default:
                                                                    switch ("" + "a") {
                                                                    case "hello":
                                                                        v3 = v3
                                                                    case "\n\t":
                                                                        v61 := GenS6{F0: 127, F1: 42}
                                                                        _ = v61
                                                                    case "":
                                                                        v3 -= 2
                                                                    }
                                                                    fmt.Println(fi20, v3)
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        v23 -= v23
                                                                    case "\n\t":
                                                                        var v62 int = (fi20 >> ((271) & 63))
                                                                        _ = v62
                                                                    case "a":
                                                                        _ = len(v4)
                                                                    default:
                                                                        var v63 uint16 = v3
                                                                        _ = v63
                                                                    }
                                                                }
                                                            }
                                                            v3 -= v3
                                                            var v64 bool = false
                                                            _ = v64
                                                        }
                                                    }
                                                    v27 += (v27 / (((v27 / (((v27 << ((v27) & 63))) | 1))) | 1))
                                                }
                                            case "a":
                                                var v65 uint = (v27 << ((v27) & 63))
                                                _ = v65
                                            case "":
                                                v27 += (^(v27))
                                            default:
                                                var v66 bool = ((!true) && true)
                                                _ = v66
                                            }
                                        case -9223372036854775808:
                                            if ((fi20 - (v23 - 128)) <= v23) {
                                                var v67 int16 = 0
                                                _ = v67
                                            }
                                        }
                                    default:
                                        var v68 uint = (v27 / (((^(v27))) | 1))
                                        _ = v68
                                    }
                                    _ = v27
                                }
                                v4 = append(v4, 2)
                            }
                            _ = v23
                        }
                        if (0.1 < (0.1 * (-2.25 + 3.14159))) {
                            v3 += 237
                        }
                        var v69 int = (fi20 << ((42) & 63))
                        _ = v69
                    } else {
                        var v70 uint32 = 4294967295
                        v71 := GenS6{F0: 0, F1: 20}
                        _ = v70
                        _ = v71
                    }
                    var v72 bool = (false && true)
                    var v73 float32 = ((float32(0.1) * -2.25) * 3.14159)
                    _ = v72
                    _ = v73
                }
                v2[fi19] = 2147483647
                _ = fi20
            }
            var v74 int = fi19
            _ = fi19
            _ = v74
        }
        v75 := GenS7{F0: 422}
        v76 := GenS8{F0: 32768, F1: (-475 > -1)}
        _ = v75
        _ = v76
    } else {
        var v77 uint16 = v3
        _ = len(v4)
        v78 := GenS1{F0: -9223372036854775808, F1: -9223372036854775808, F2: 855}
        _ = v77
        _ = v78
    }
    v79 := GenS5{F0: "a", F1: 675, F2: 450}
    v3 -= (v3 & v3)
    if true {
        v3 += 1
        v80 := GenS8{F0: 0, F1: true}
        _ = v80
    } else {
        var v81 uint8 = 42
        _ = v81
    }
    v82 := v1[4294967295]
    if (895 > 275) {
        for fi83 := 0; fi83 < 5; fi83++ {
            if v82 {
                if (v82 || v82) {
                    fmt.Println(fi83)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    var v84 uint = 701
                    fmt.Println(v82)
                    _ = v84
                }
                if (1 >= -1) {
                    v85 := GenS9{F0: 42}
                    _ = v85
                }
                v1[583] = (!true)
            }
            fmt.Println(v82, v3, v82)
            for fi86 := 0; fi86 < 1; fi86++ {
                if true {
                    v1[128] = v82
                }
                var v87 float64 = 1e10
                _ = fi86
                _ = v87
            }
            _ = fi83
        }
        if (!v82) {
            switch -785 {
            case 9223372036854775807:
                if (1.5 >= (1.5 + ((1e10 / 1.0) + float32(float64(1e10))))) {
                    for fi88 := 0; fi88 < 19; fi88++ {
                        var v89 int8 = 1
                        v82 = true
                        _ = fi88
                        _ = v89
                    }
                } else {
                    var v90 int = -9223372036854775808
                    var v91 string = ("foo bar" + "hello")
                    for fi92 := 0; fi92 < 9; fi92++ {
                        if v82 {
                            switch v91 {
                            case "foo bar":
                                v90 = v90
                            case "\n\t":
                                v91 = ("\n\t" + "hello")
                            case "":
                                if (-792 <= -1) {
                                    v82 = v82
                                }
                            default:
                                _ = len(v4)
                            }
                            for fi93 := 0; fi93 < 1; fi93++ {
                                if (v82 && v82) {
                                    var v94 int = 128
                                    _ = v94
                                }
                                if (!(v82 || false)) {
                                    v82 = (!v82)
                                    if v82 {
                                        if (1 <= -9223372036854775808) {
                                            switch v91 {
                                            case "foo bar":
                                                var v95 int8 = 100
                                                _ = v95
                                            case "\n\t":
                                                v82 = (false || ((1 != 2) && v82))
                                            }
                                            v96 := GenS5{F0: v91, F1: fi93, F2: 2147483647}
                                            if (710 < 0) {
                                                var v97 float64 = 1e10
                                                var v98 uint64 = 297
                                                _ = len(v4)
                                                if v82 {
                                                    switch fi92 {
                                                    case 789:
                                                        v3 = v3
                                                    case 32768:
                                                        v99 := GenS7{F0: uint32(v98)}
                                                        _ = v99
                                                    }
                                                }
                                                v3 = v3
                                                _ = v97
                                                _ = v98
                                            }
                                            _ = v96
                                        }
                                        var v100 int16 = -357
                                        _ = v100
                                    } else {
                                        v91 = ("\n\t" + v91)
                                        if (false && ((!(0.1 == float32(1.5))) && (1e10 <= float64(float32(1.5))))) {
                                            var v101 int = (fi92 + (v90 % ((0) | 1)))
                                            if false {
                                                var v102 bool = (42 >= -275)
                                                _ = v102
                                            }
                                            fmt.Println(v90)
                                            _ = v101
                                        } else {
                                            v103 := GenS5{F0: v91, F1: -9223372036854775808, F2: 9}
                                            fmt.Println(fi92, v3, v82)
                                            var v104 uint32 = 86
                                            v104 = v104
                                            _ = v103
                                            _ = v104
                                        }
                                        v3 += v3
                                        if v82 {
                                            fmt.Println(v3, fi93, v82)
                                            var v105 uint32 = 2147483647
                                            _ = v105
                                        }
                                    }
                                    v91 = ""
                                } else {
                                    v90 += (v90 / (((v90 ^ fi93)) | 1))
                                }
                                v91 = ("foo bar" + "\n\t")
                                _ = fi93
                            }
                        }
                        _ = fi92
                    }
                    _ = v90
                    _ = v91
                }
            case 2147483647:
                v3 = v3
            case 32767:
                var v106 string = "foo bar"
                _ = v106
            default:
                if ((v82 || (-2147483648 != 0)) || v82) {
                    for fi107 := 0; fi107 < 11; fi107++ {
                        v82 = (v82 || false)
                        v2[32767] = 1
                        for fi108 := 0; fi108 < 1; fi108++ {
                            var v109 int64 = 511
                            if false {
                                v110 := GenS2{F0: 1, F1: 116, F2: (v3 & (v3 + v3))}
                                v111 := GenS5{F0: ("" + "\n\t"), F1: fi107, F2: 1}
                                var v112 uint8 = genFn0(-2147483648)
                                _ = v110
                                _ = v111
                                _ = v112
                            } else {
                                var v113 uint8 = genFn0(1)
                                v114 := GenS10{F0: (4294967295 <= 0), F1: (v3 | (^(v3))), F2: 1}
                                fmt.Println(v109)
                                _ = v113
                                _ = v114
                            }
                            fmt.Println(v3, fi107, v82)
                            _ = fi108
                            _ = v109
                        }
                        _ = fi107
                    }
                    fmt.Println(v82, v82)
                    v115 := GenS3{F0: -858}
                    _ = v115
                }
            }
        }
        if (-9223372036854775808 >= -226) {
            v82 = v82
            fmt.Println(v3, v3)
            var v116 float32 = 1e10
            _ = v116
        } else {
            var v117 uint8 = 42
            _ = len(v1)
            _ = v117
        }
    } else {
        for fi118 := 0; fi118 < 6; fi118++ {
            fmt.Println(fi118, v3)
            var v119 int8 = 1
            _ = fi118
            _ = v119
        }
        switch ("" + ("" + "foo bar")) {
        case "foo bar":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "hello":
            v1[32768] = false
        case "\n\t":
            if (255 == 255) {
                for fi120 := 0; fi120 < 13; fi120++ {
                    _ = len(v2)
                    for fi121 := 0; fi121 < 8; fi121++ {
                        if true {
                            fmt.Println(fi121, v3, v82)
                        }
                        v122 := GenS2{F0: (v3 % (((v3 / (((v3 ^ v3)) | 1))) | 1)), F1: -55, F2: (^(v3))}
                        _ = fi121
                        _ = v122
                    }
                    var v123 int = fi120
                    _ = fi120
                    _ = v123
                }
                if false {
                    for fi124 := 0; fi124 < 12; fi124++ {
                        var v125 int16 = 42
                        v82 = (-2.25 == (-2.25 * 0.1))
                        var v126 uint64 = 4294967295
                        _ = fi124
                        _ = v125
                        _ = v126
                    }
                }
            }
        }
    }
    for fi127 := 0; fi127 < 20; fi127++ {
        v128 := GenS5{F0: "", F1: (-(fi127)), F2: 284}
        for fi129 := 0; fi129 < 18; fi129++ {
            if v82 {
                fmt.Println(fi127, v3)
                if (42 <= 65535) {
                    var v130 int64 = 0
                    _ = v130
                } else {
                    switch "\n\t" {
                    case "a":
                        fmt.Println(fi129)
                    case "foo bar":
                        var v131 int64 = 9223372036854775807
                        _ = v131
                    default:
                        v132 := GenS5{F0: "a", F1: (fi129 % ((-87) | 1)), F2: 42}
                        _ = v132
                    }
                    if ((0.0 + 1.5) > 0.1) {
                        var v133 string = ("hello" + ("foo bar" + ("a" + "\n\t")))
                        switch fi127 {
                        case 4294967295:
                            var v134 float64 = 1.5
                            if v82 {
                                v1[(-(fi129))] = (fi127 >= fi129)
                                v135 := GenS11{F0: 103}
                                v1[fi129] = (v134 > 3.14159)
                                _ = v135
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            switch v133 {
                            case "\n\t":
                                var v136 int8 = -128
                                _ = v136
                            case "hello":
                                v4 = append(v4, 2147483648)
                            }
                            _ = v134
                        default:
                            if v82 {
                                v2[fi127] = 2147483648
                                v3 = v3
                                v1[fi127] = v82
                            } else {
                                var v137 int16 = -716
                                v133 = ("" + v133)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v137
                            }
                        }
                        v3 = (v3 / ((65535) | 1))
                        switch v133 {
                        case "foo bar":
                            if v82 {
                                v82 = (v82 && (!false))
                                switch fi129 {
                                case 1:
                                    _ = len(v4)
                                default:
                                    var v138 bool = (false && ((-473 <= 0) && (fi129 < 32767)))
                                    _ = v138
                                }
                            }
                        case "hello":
                            v4 = append(v4, 2147483648)
                        }
                        _ = v133
                    } else {
                        switch "" {
                        case "\n\t":
                            v82 = true
                        case "":
                            fmt.Println(fi129, fi127, fi129)
                        default:
                            switch "hello" {
                            case "hello":
                                var v139 uint = 18446744073709551615
                                var v140 int32 = 1
                                v141 := GenS10{F0: v82, F1: v3, F2: 94}
                                v1[(fi127 >> ((uint(v3)) & 63))] = true
                                _ = v139
                                _ = v140
                                _ = v141
                            default:
                                switch "a" {
                                case "":
                                    v3 = (v3 % ((v3) | 1))
                                case "a":
                                    v3 += (v3 / (((v3 + v3)) | 1))
                                default:
                                    var v142 uint8 = 0
                                    _ = v142
                                }
                            }
                        }
                    }
                    v143 := GenS12{F0: -369, F1: 437, F2: 2}
                    _ = v143
                }
            }
            var v144 uint16 = 430
            v144 = 1
            if (!(!(!(v82 || (!(-2147483648 < -1)))))) {
                v145 := GenS13{F0: int64(v144), F1: int64(v144)}
                if (255 >= 2) {
                    v82 = (!v82)
                    v3 += v144
                } else {
                    v82 = false
                }
                _ = v145
            } else {
                var v146 string = "hello"
                _ = v146
            }
            _ = fi129
            _ = v144
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi127
        _ = v128
    }
    delete(v1, 32768)
    for fi147 := 0; fi147 < 18; fi147++ {
        for fi148 := 0; fi148 < 18; fi148++ {
            v82 = (0.1 < ((float64(0.1) * (1e10 / (0.0 + float64(((float32(0.1) * -2.25) / 1.5))))) * ((0.1 * 1e10) / 1e10)))
            var v149 int64 = 0
            _ = fi148
            _ = v149
        }
        _ = fi147
    }
    v150 := GenS13{F0: 584, F1: 1}
    fmt.Println(v82)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v150
    _ = v2
    _ = v3
    _ = v4
    _ = v79
    _ = v82
}
