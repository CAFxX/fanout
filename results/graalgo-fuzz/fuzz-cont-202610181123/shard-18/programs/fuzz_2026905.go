package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 int16; F1 string; F2 int32 }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 int64; F1 int }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS10 struct { F0 uint32; F1 string }
type GenS11 struct { F0 int32; F1 uint16 }
type GenS12 struct { F0 int16; F1 string; F2 uint }
type GenS13 struct { F0 int8; F1 int64; F2 int16 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int32; F1 uint32 }
type GenS17 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]bool)
    var v3 int32 = 193
    var v4 uint = 391
    v5 := make(map[int]int)
    var v6 uint8 = 169
    v7 := GenS1{F0: -1, F1: 2147483648, F2: false}
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS1{F0: 32767, F1: 4294967295, F2: (-1 < 9223372036854775807)}
            v4 += (v4 >> ((2) & 63))
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 13; fi11++ {
            switch "" {
            case "foo bar":
                var v12 int64 = -1
                _ = v12
            default:
                var v13 int16 = 256
                v14 := GenS2{F0: (v4 * 350)}
                if (32767 <= 256) {
                    if false {
                        switch (fi8 * fi11) {
                        case 156:
                            if false {
                                v15 := GenS1{F0: v13, F1: 42, F2: true}
                                if (true || (true || ((v6 + 255) >= v6))) {
                                    if true {
                                        fmt.Println(v3, fi11)
                                    } else {
                                        var v16 bool = true
                                        if (587 <= 127) {
                                            var v17 float64 = (0.1 + -2.25)
                                            var v18 float64 = (1e10 - v17)
                                            v13 = v13
                                            _ = v17
                                            _ = v18
                                        } else {
                                            var v19 uint8 = 51
                                            v20 := []int{int(v13), fi11, fi8, int(v4)}
                                            var v21 uint8 = v6
                                            _ = v19
                                            _ = v20
                                            _ = v21
                                        }
                                        v22 := GenS2{F0: (^(v4))}
                                        _ = v16
                                        _ = v22
                                    }
                                } else {
                                    if false {
                                        v6 += 0
                                        v23 := []int{fi8, -9223372036854775808, fi11}
                                        _ = v23
                                    } else {
                                        fmt.Println(v4, v13)
                                        var v24 int8 = 100
                                        switch ((("\n\t" + ("foo bar" + "foo bar")) + "\n\t") + ("a" + "a")) {
                                        case "foo bar":
                                            switch ("foo bar" + ("hello" + "hello")) {
                                            case "\n\t":
                                                v3 += v3
                                            case "foo bar":
                                                v25 := GenS1{F0: -37, F1: 4294967295, F2: (false && false)}
                                                _ = v25
                                            case "hello":
                                                if ((!((1e10 / 1.0) != 1e10)) && (281 > 65535)) {
                                                    if false {
                                                        if (3.14159 == -2.25) {
                                                            v6 = (^(v6))
                                                        } else {
                                                            if ((v24 + v24) >= v24) {
                                                                v26 := v2[fi8]
                                                                v24 -= v24
                                                                _ = v26
                                                            }
                                                            v27 := GenS3{F0: (v13 << ((v4) & 63))}
                                                            _ = v27
                                                        }
                                                        if (!((fi8 >> ((v4) & 63)) >= fi8)) {
                                                            v28 := []int{}
                                                            v4 = (v4 << (((v4 % (((v4 % ((65535) | 1))) | 1))) & 63))
                                                            if false {
                                                                var v29 int = (fi8 * (fi11 | fi11))
                                                                var v30 int64 = 580
                                                                v31 := GenS1{F0: (v13 * (v13 + -440)), F1: 665, F2: (float32(((0.1 / (0.0 + (1e10 / 0.1))) + (((3.14159 - float64(0.0)) + -2.25) / -2.25))) <= 0.0)}
                                                                _ = v29
                                                                _ = v30
                                                                _ = v31
                                                            }
                                                            _ = v28
                                                        } else {
                                                            if ((!false) || (13 > 1)) {
                                                                v24 -= v24
                                                                v32 := GenS4{F0: "", F1: (-9223372036854775808 <= 9223372036854775807)}
                                                                if (!(9223372036854775807 <= 1)) {
                                                                    var v33 uint64 = 0
                                                                    _ = v33
                                                                } else {
                                                                    fmt.Println(v3, fi8)
                                                                    var v34 int32 = (v3 ^ v3)
                                                                    if false {
                                                                        var v35 bool = false
                                                                        if v35 {
                                                                            var v36 int64 = int64(v4)
                                                                            _ = v36
                                                                        } else {
                                                                            v3 -= (v3 + (v34 >> ((523) & 63)))
                                                                            var v37 float32 = 3.14159
                                                                            _ = v37
                                                                        }
                                                                        var v38 float64 = 1e10
                                                                        _ = v35
                                                                        _ = v38
                                                                    }
                                                                    _ = v34
                                                                }
                                                                _ = v32
                                                            } else {
                                                                v39 := GenS5{F0: v13, F1: ("foo bar" + "a"), F2: v3}
                                                                _ = v39
                                                            }
                                                            if false {
                                                                fmt.Println(v6, v13, v4)
                                                            } else {
                                                                v24 += v24
                                                                v4 = v4
                                                                fmt.Println(v24, v13)
                                                            }
                                                            v13 += (v13 & v13)
                                                            v40 := []int{(fi8 * -1)}
                                                            _ = v40
                                                        }
                                                    } else {
                                                        var v41 int8 = 127
                                                        var v42 uint = 0
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    if (553 == 764) {
                                                        if (true && true) {
                                                            v43 := []int{fi11, (fi11 ^ (-(fi8))), fi11}
                                                            _ = v43
                                                        } else {
                                                            delete(v1, 42)
                                                        }
                                                        v3 = (^(v3))
                                                        if (false && true) {
                                                            var v44 uint32 = 382
                                                            _ = v44
                                                        }
                                                    } else {
                                                        if true {
                                                            var v45 bool = false
                                                            if (v45 && (v45 && false)) {
                                                                if (85 <= int8(v4)) {
                                                                    var v46 int = -9223372036854775808
                                                                    v13 = (v13 % ((v13) | 1))
                                                                    var v47 uint32 = 363
                                                                    v3 -= v3
                                                                    switch (((("hello" + "a") + "foo bar") + "hello") + ("\n\t" + "a")) {
                                                                    case "":
                                                                        v24 += v24
                                                                    }
                                                                    _ = v46
                                                                    _ = v47
                                                                }
                                                            }
                                                            if false {
                                                                var v48 float32 = 3.14159
                                                                v49 := GenS6{F0: 0, F1: 1}
                                                                _ = v48
                                                                _ = v49
                                                            } else {
                                                                v3 = (v3 ^ (v3 | (-(v3))))
                                                            }
                                                            _ = v45
                                                        }
                                                        fmt.Println(fi11, v6)
                                                        v24 += v24
                                                    }
                                                    var v50 float32 = (float32(1e10) * 0.0)
                                                    _ = v50
                                                }
                                            }
                                        case "\n\t":
                                            v6 += v6
                                        case "":
                                            var v51 uint8 = v6
                                            if (!(!true)) {
                                                if (!true) {
                                                    v52 := GenS6{F0: 67, F1: 0}
                                                    _ = v52
                                                }
                                            }
                                            v53 := GenS6{F0: 11, F1: 506}
                                            _ = v51
                                            _ = v53
                                        }
                                        fmt.Println(fi11)
                                        _ = v24
                                    }
                                }
                                var v54 float32 = float32(((((0.0 * 0.1) * (float64(1.5) * 0.1)) + ((3.14159 / (((0.1 - 3.14159) * -2.25) / float64((0.1 / 0.1)))) + 0.1)) + float64(float32((3.14159 / -2.25)))))
                                var v55 int16 = v13
                                switch (fi11 | (^(fi8))) {
                                case -902:
                                    var v56 uint16 = 65535
                                    _ = v56
                                case 42:
                                    v54 = 1e10
                                }
                                v57 := GenS7{F0: -126, F1: (fi8 / (((^(fi11))) | 1))}
                                if false {
                                    v4 = uint(v6)
                                } else {
                                    if true {
                                        v4 = v4
                                        if (v13 >= 1000) {
                                            if false {
                                                var v58 uint16 = 310
                                                _ = v58
                                            } else {
                                                if false {
                                                    switch (fi11 | int(v4)) {
                                                    case 127:
                                                        var v59 int64 = 365
                                                        _ = v59
                                                    }
                                                } else {
                                                    var v60 uint = v4
                                                    v61 := GenS5{F0: (v13 * (v55 / ((v55) | 1))), F1: "", F2: v3}
                                                    if (true || (uint(v6) > (v60 / ((v4) | 1)))) {
                                                        var v62 int64 = 42
                                                        switch 32768 {
                                                        case 893:
                                                            fmt.Println(v54, v62)
                                                        case 32768:
                                                            fmt.Println(v6, v60, v4)
                                                        case 230:
                                                            var v63 bool = false
                                                            _ = v63
                                                        default:
                                                            var v64 int64 = -9223372036854775808
                                                            _ = v64
                                                        }
                                                        var v65 float64 = ((-2.25 - 0.0) + 3.14159)
                                                        var v66 uint8 = (v6 + 203)
                                                        var v67 int64 = (v62 % ((v62) | 1))
                                                        _ = v62
                                                        _ = v65
                                                        _ = v66
                                                        _ = v67
                                                    } else {
                                                        _ = len(v1)
                                                        switch "" {
                                                        case "foo bar":
                                                            v68 := GenS2{F0: v4}
                                                            _ = v68
                                                        case "hello":
                                                            var v69 int64 = 625
                                                            _ = v69
                                                        case "\n\t":
                                                            var v70 uint8 = uint8(v4)
                                                            _ = v70
                                                        default:
                                                            var v71 int32 = (v3 % ((559) | 1))
                                                            _ = v71
                                                        }
                                                    }
                                                    _ = v60
                                                    _ = v61
                                                }
                                            }
                                        }
                                    } else {
                                        fmt.Println(fi8, v4)
                                        if (-9223372036854775808 < -674) {
                                            switch fi8 {
                                            case 128:
                                                switch (fi11 + -138) {
                                                case 9223372036854775807:
                                                    var v72 uint = v4
                                                    _ = v72
                                                case 128:
                                                    var v73 uint32 = 1
                                                    _ = v73
                                                case 2147483647:
                                                    v74 := []int{(fi8 | -2147483648), fi8, fi8, fi8}
                                                    _ = v74
                                                default:
                                                    v75 := GenS1{F0: (v55 % ((-1) | 1)), F1: 58, F2: true}
                                                    _ = v75
                                                }
                                            case -1:
                                                v54 -= (v54 * ((v54 / (v54 - (1e10 + ((((v54 - 1e10) + v54) * v54) + (v54 / (((v54 / v54) - float32((float64(0.1) * 0.1))) / v54)))))) / float32(1e10)))
                                            default:
                                                var v76 int64 = 0
                                                _ = v76
                                            }
                                            if (true || ((fi8 * (fi8 ^ fi8)) == (fi11 & (fi8 ^ (fi8 & fi11))))) {
                                                var v77 int32 = (v3 / ((42) | 1))
                                                switch (fi11 & 9223372036854775807) {
                                                case 32767:
                                                    if (975 == 18446744073709551615) {
                                                        switch (fi11 | 2147483647) {
                                                        case 2147483647:
                                                            v6 += 1
                                                        case 360:
                                                            var v78 uint16 = 801
                                                            _ = v78
                                                        case 533:
                                                            if ((v77 / (((v3 + v3)) | 1)) > 187) {
                                                                if true {
                                                                    v6 -= v6
                                                                    switch "hello" {
                                                                    case "":
                                                                        if (0 >= 0) {
                                                                            var v79 int = -968
                                                                            _ = v79
                                                                        }
                                                                    }
                                                                    v13 = (v55 >> ((127) & 63))
                                                                }
                                                                fmt.Println(v3)
                                                                v80 := GenS3{F0: v55}
                                                                _ = v80
                                                            } else {
                                                                if (!false) {
                                                                    var v81 string = "a"
                                                                    v4 = (v4 / ((255) | 1))
                                                                    _ = v81
                                                                }
                                                            }
                                                        default:
                                                            var v82 uint64 = 181
                                                            _ = v82
                                                        }
                                                        var v83 int16 = v13
                                                        if (!((fi8 & fi8) <= (-(fi8)))) {
                                                            v54 = float32(1e10)
                                                            v55 -= (v13 + v55)
                                                            switch fi8 {
                                                            case 42:
                                                                v13 = v55
                                                            case -428:
                                                                switch fi11 {
                                                                case -2147483648:
                                                                    v6 += 27
                                                                default:
                                                                    v84 := GenS4{F0: "", F1: (2147483648 <= 0)}
                                                                    _ = v84
                                                                }
                                                            default:
                                                                v85 := GenS3{F0: (v13 << (((v4 << ((2) & 63))) & 63))}
                                                                _ = v85
                                                            }
                                                            var v86 string = ("" + "hello")
                                                            _ = v86
                                                        }
                                                        if false {
                                                            var v87 int = fi8
                                                            v2[32767] = (1 <= 871)
                                                            _ = v87
                                                        }
                                                        v6 -= (v6 ^ 29)
                                                        _ = v83
                                                    } else {
                                                        var v88 int = (fi8 / (((fi8 - fi11)) | 1))
                                                        switch ("a" + ("" + "foo bar")) {
                                                        case "a":
                                                            v89 := GenS6{F0: 42, F1: 0}
                                                            switch 254 {
                                                            case 32768:
                                                                if (!(100 <= int8(v3))) {
                                                                    if true {
                                                                        if false {
                                                                            var v90 uint = (v4 >> ((255) & 63))
                                                                            v54 -= (v54 * v54)
                                                                            _ = v90
                                                                        }
                                                                    } else {
                                                                        var v91 int8 = 42
                                                                        if true {
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                v92 := []int{}
                                                                                _ = v92
                                                                            default:
                                                                                delete(v2, (fi11 % ((fi11) | 1)))
                                                                            }
                                                                            if true {
                                                                                v91 = v91
                                                                                if (1 > 2147483647) {
                                                                                    if ((v6 & v6) > 128) {
                                                                                        v13 = (v55 << ((255) & 63))
                                                                                    }
                                                                                    v91 = 42
                                                                                } else {
                                                                                    if (false || false) {
                                                                                        if true {
                                                                                            switch fi8 {
                                                                                            case 4294967295:
                                                                                                _ = len(v5)
                                                                                            case 2147483647:
                                                                                                switch (("\n\t" + "") + "foo bar") {
                                                                                                case "a":
                                                                                                    v88 += (fi11 << (((^(v4))) & 63))
                                                                                                case "foo bar":
                                                                                                    v6 = v6
                                                                                                case "\n\t":
                                                                                                    v6 = 214
                                                                                                }
                                                                                            case 609:
                                                                                                if true {
                                                                                                    if ((!(v91 <= (v91 - (v91 - v91)))) && false) {
                                                                                                        v88 = 128
                                                                                                    }
                                                                                                } else {
                                                                                                    v93 := GenS3{F0: -643}
                                                                                                    if false {
                                                                                                        v91 -= v91
                                                                                                    }
                                                                                                    _ = v93
                                                                                                }
                                                                                            default:
                                                                                                v94 := v1[2147483648]
                                                                                                _ = v94
                                                                                            }
                                                                                            var v95 float32 = v54
                                                                                            _ = v95
                                                                                        }
                                                                                        var v96 uint64 = 42
                                                                                        v97 := []int{int(v4), (v88 | (fi8 - 42)), (fi8 * (fi11 >> ((449) & 63)))}
                                                                                        _ = v96
                                                                                        _ = v97
                                                                                    } else {
                                                                                        var v98 float64 = float64(-2.25)
                                                                                        if (!true) {
                                                                                            v99 := GenS8{F0: 111}
                                                                                            v54 = v54
                                                                                            if true {
                                                                                                if (!false) {
                                                                                                    v6 = 123
                                                                                                } else {
                                                                                                    v54 += float32((v98 - (1e10 + (v98 + (v98 + -2.25)))))
                                                                                                    var v100 int16 = int16(v4)
                                                                                                    v3 += (v77 | v77)
                                                                                                    v54 = v54
                                                                                                    _ = v100
                                                                                                }
                                                                                                fmt.Println(v88, v13)
                                                                                                switch "hello" {
                                                                                                case "a":
                                                                                                    switch (^(fi11)) {
                                                                                                    case 1:
                                                                                                        v101 := GenS9{F0: 2, F1: 0, F2: 2147483647}
                                                                                                        _ = v101
                                                                                                    case 127:
                                                                                                        var v102 string = "foo bar"
                                                                                                        _ = v102
                                                                                                    }
                                                                                                case "":
                                                                                                    v91 = v91
                                                                                                    v103 := []int{fi8, v88, (fi11 | 127)}
                                                                                                    _ = v103
                                                                                                default:
                                                                                                    switch "foo bar" {
                                                                                                    case "foo bar":
                                                                                                        fmt.Println(fi8, fi8, v88)
                                                                                                    case "hello":
                                                                                                        v6 += (^(v6))
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                v104 := GenS6{F0: 772, F1: 4294967295}
                                                                                                var v105 bool = true
                                                                                                _ = v104
                                                                                                _ = v105
                                                                                            }
                                                                                            v106 := GenS4{F0: "a", F1: (((!false) && true) && true)}
                                                                                            _ = v106
                                                                                            _ = v99
                                                                                        }
                                                                                        _ = v98
                                                                                    }
                                                                                    if false {
                                                                                        v107 := GenS10{F0: 0, F1: ("" + "foo bar")}
                                                                                        var v108 uint32 = 606
                                                                                        if (-1 == 0) {
                                                                                            switch ("hello" + ("foo bar" + "hello")) {
                                                                                            case "a":
                                                                                                if (uint16(v108) == 108) {
                                                                                                    if (418 >= 4294967295) {
                                                                                                        v4 += 4294967295
                                                                                                    } else {
                                                                                                        v13 = (v55 * v55)
                                                                                                        var v109 uint64 = uint64(v108)
                                                                                                        v55 = (v13 >> ((v4) & 63))
                                                                                                        _ = v109
                                                                                                    }
                                                                                                } else {
                                                                                                    if false {
                                                                                                        v110 := GenS5{F0: -859, F1: ("\n\t" + ""), F2: v77}
                                                                                                        var v111 int8 = -1
                                                                                                        fmt.Println(v3, fi8)
                                                                                                        _ = v110
                                                                                                        _ = v111
                                                                                                    } else {
                                                                                                        var v112 uint = v4
                                                                                                        _ = v112
                                                                                                    }
                                                                                                    if true {
                                                                                                        switch fi11 {
                                                                                                        case 32768:
                                                                                                            switch "" {
                                                                                                            case "foo bar":
                                                                                                                fmt.Println(v108)
                                                                                                            case "a":
                                                                                                                var v113 uint = v4
                                                                                                                _ = v113
                                                                                                            default:
                                                                                                                v6 += (v6 % ((v6) | 1))
                                                                                                            }
                                                                                                            fmt.Println(v54, v77, fi11)
                                                                                                            v1[v108] = (v108 << (((^(v4))) & 63))
                                                                                                        case 80:
                                                                                                            if true {
                                                                                                                v114 := []int{(fi11 % ((1) | 1)), fi8}
                                                                                                                fmt.Println(fi11, v108)
                                                                                                                _ = v114
                                                                                                            }
                                                                                                        case 2147483647:
                                                                                                            var v115 int16 = (v55 << ((18446744073709551615) & 63))
                                                                                                            _ = v115
                                                                                                        default:
                                                                                                            fmt.Println(v55, v6, fi8)
                                                                                                        }
                                                                                                        if ((^(v4)) >= uint(v108)) {
                                                                                                            fmt.Println(v54, v77)
                                                                                                            v4 = v4
                                                                                                            v54 = (1e10 / 0.1)
                                                                                                        } else {
                                                                                                            v116 := GenS2{F0: (v4 | (v4 | (v4 + (^(v4)))))}
                                                                                                            _ = v116
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v117 float64 = (1e10 / 0.1)
                                                                                                        var v118 uint16 = 32767
                                                                                                        var v119 string = ("\n\t" + "hello")
                                                                                                        var v120 bool = true
                                                                                                        fmt.Println(v108, v6, v55)
                                                                                                        switch v88 {
                                                                                                        case -2147483648:
                                                                                                            v121 := []int{fi11, (v88 & 410), fi11, fi11}
                                                                                                            _ = v121
                                                                                                        case -1:
                                                                                                            v77 += 1
                                                                                                        case 2147483647:
                                                                                                            v122 := GenS11{F0: (^(v77)), F1: 65535}
                                                                                                            _ = v122
                                                                                                        default:
                                                                                                            var v123 int32 = int32(v118)
                                                                                                            _ = v123
                                                                                                        }
                                                                                                        _ = v117
                                                                                                        _ = v118
                                                                                                        _ = v119
                                                                                                        _ = v120
                                                                                                    }
                                                                                                    var v124 int8 = (v91 * int8(v88))
                                                                                                    _ = v124
                                                                                                }
                                                                                            case "hello":
                                                                                                v125 := []int{(v88 * fi8), (fi11 % ((fi11) | 1))}
                                                                                                _ = v125
                                                                                            case "foo bar":
                                                                                                var v126 float32 = v54
                                                                                                _ = v126
                                                                                            default:
                                                                                                v6 += (^(v6))
                                                                                                if (1 <= 892) {
                                                                                                    if false {
                                                                                                        v6 += 42
                                                                                                    } else {
                                                                                                        var v127 uint = (v4 - (v4 / ((v4) | 1)))
                                                                                                        if (((!false) || ((!false) || (18446744073709551615 == 2))) && false) {
                                                                                                            if true {
                                                                                                                fmt.Println(v13)
                                                                                                            }
                                                                                                            v128 := GenS10{F0: (v108 - (v108 | (v108 << (((v127 & (^(v4)))) & 63)))), F1: "\n\t"}
                                                                                                            _ = v128
                                                                                                        } else {
                                                                                                            var v129 uint16 = 1
                                                                                                            var v130 int32 = v3
                                                                                                            v3 += (v130 % ((v3) | 1))
                                                                                                            _ = v129
                                                                                                            _ = v130
                                                                                                        }
                                                                                                        _ = v127
                                                                                                    }
                                                                                                    if true {
                                                                                                        v4 = 648
                                                                                                        var v131 string = (("\n\t" + ("\n\t" + ("foo bar" + "foo bar"))) + "hello")
                                                                                                        var v132 uint16 = uint16(v4)
                                                                                                        _ = v131
                                                                                                        _ = v132
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            fmt.Println(v77, v4, v3)
                                                                                        } else {
                                                                                            var v133 float64 = float64(-2.25)
                                                                                            v134 := GenS3{F0: (v13 & (v13 << ((v4) & 63)))}
                                                                                            v3 += v77
                                                                                            _ = v133
                                                                                            _ = v134
                                                                                        }
                                                                                        _ = v107
                                                                                        _ = v108
                                                                                    }
                                                                                    v13 -= (v13 % (((v13 | v55)) | 1))
                                                                                }
                                                                                fmt.Println(v3)
                                                                            } else {
                                                                                var v135 int8 = 68
                                                                                var v136 uint8 = v6
                                                                                _ = v135
                                                                                _ = v136
                                                                            }
                                                                        }
                                                                        switch (fi8 & (v88 ^ 127)) {
                                                                        case 42:
                                                                            v88 += v88
                                                                        case 0:
                                                                            fmt.Println(v54, v3, fi8)
                                                                        }
                                                                        var v137 uint16 = 255
                                                                        _ = v137
                                                                        _ = v91
                                                                    }
                                                                }
                                                            case 128:
                                                                if true {
                                                                    if (true && (!((^(v13)) > (v55 % ((v55) | 1))))) {
                                                                        v88 = 4294967295
                                                                    }
                                                                    fmt.Println(fi11)
                                                                }
                                                            case 127:
                                                                v77 = v3
                                                            default:
                                                                v138 := GenS12{F0: (v13 + (v13 | v13)), F1: "", F2: (v4 + (v4 / (((v4 / (((v4 * (v4 - v4))) | 1))) | 1)))}
                                                                v139 := GenS13{F0: 1, F1: 114, F2: v55}
                                                                v140 := []int{42}
                                                                _ = v138
                                                                _ = v139
                                                                _ = v140
                                                            }
                                                            v141 := []int{127, (fi11 | (^(fi11)))}
                                                            _ = v141
                                                            _ = v89
                                                        case "\n\t":
                                                            var v142 float32 = 1.5
                                                            _ = v142
                                                        case "foo bar":
                                                            v143 := v2[fi8]
                                                            _ = v143
                                                        default:
                                                            v54 -= v54
                                                        }
                                                        _ = v88
                                                    }
                                                }
                                                _ = v77
                                            } else {
                                                v144 := []int{32767, (fi8 ^ int(v3)), 42}
                                                _ = v144
                                            }
                                        }
                                        switch "hello" {
                                        case "\n\t":
                                            v145 := []int{}
                                            _ = v145
                                        case "a":
                                            v13 = v13
                                            switch (fi11 ^ (fi11 + fi8)) {
                                            case 268:
                                                if false {
                                                    v13 += (v13 % (((v13 << ((4294967295) & 63))) | 1))
                                                }
                                            }
                                            if (v54 == 0.1) {
                                                v146 := GenS1{F0: v55, F1: 4294967295, F2: (4294967295 == 2147483648)}
                                                var v147 float32 = (v54 / 3.14159)
                                                _ = v146
                                                _ = v147
                                            } else {
                                                v6 -= (v6 + (^(v6)))
                                                var v148 int64 = int64(v13)
                                                var v149 bool = false
                                                _ = v148
                                                _ = v149
                                            }
                                        case "":
                                            var v150 float32 = (v54 + ((v54 - float32(-2.25)) * 1e10))
                                            _ = v150
                                        }
                                        switch "" {
                                        case "\n\t":
                                            var v151 uint16 = 1
                                            _ = v151
                                        case "foo bar":
                                            v152 := []int{fi8, fi11, fi11}
                                            _ = v152
                                        }
                                    }
                                    var v153 int16 = (v13 ^ v13)
                                    var v154 float64 = (-2.25 + (3.14159 / (float64(v54) / float64(1.5))))
                                    _ = v153
                                    _ = v154
                                }
                                var v155 string = ("foo bar" + ((("\n\t" + "") + (("\n\t" + "foo bar") + "")) + "foo bar"))
                                _ = v15
                                _ = v155
                                _ = v54
                                _ = v55
                                _ = v57
                            } else {
                                v156 := GenS1{F0: v13, F1: 626, F2: false}
                                if (int64(v3) > 0) {
                                    switch (fi11 / (((fi8 << (((v4 ^ v4)) & 63))) | 1)) {
                                    case -9223372036854775808:
                                        v6 = (v6 * 255)
                                    default:
                                        delete(v1, 96)
                                    }
                                    if true {
                                        v13 = -32768
                                        v6 = 128
                                        fmt.Println(v3, fi8, v3)
                                        _ = len(v2)
                                    } else {
                                        v157 := GenS14{F0: 587}
                                        _ = v157
                                    }
                                    fmt.Println(v13)
                                } else {
                                    v6 = (v6 | v6)
                                    v1[2147483647] = 1
                                }
                                v158 := GenS13{F0: 0, F1: -1, F2: v13}
                                delete(v1, 42)
                                var v159 int16 = v13
                                _ = v156
                                _ = v158
                                _ = v159
                            }
                        }
                    }
                    if (!((1 >= 127) || false)) {
                        if (!(uint32(v6) >= 65535)) {
                            var v160 uint16 = 255
                            if false {
                                switch (fi8 >> ((917) & 63)) {
                                case -2147483648:
                                    v161 := []int{(fi11 % ((9223372036854775807) | 1)), fi11, (fi8 * fi8), 241}
                                    _ = v161
                                case 261:
                                    v162 := GenS11{F0: (v3 | (v3 & v3)), F1: (v160 % ((42) | 1))}
                                    _ = v162
                                case 4294967295:
                                    v13 += (-(v13))
                                }
                                v160 = (v160 + v160)
                                var v163 int16 = v13
                                _ = v163
                            }
                            _ = v160
                        } else {
                            v164 := []int{}
                            v5[fi8] = fi8
                            _ = len(v2)
                            _ = v164
                        }
                        if (!(((0.0 / 1.0) * float64(3.14159)) != 3.14159)) {
                            switch ("hello" + "foo bar") {
                            case "foo bar":
                                v165 := GenS14{F0: 198}
                                v3 -= v3
                                v6 -= (v6 - v6)
                                _ = v165
                            }
                            if (false || false) {
                                v3 += (v3 & 42)
                            } else {
                                v3 = (^(v3))
                            }
                        } else {
                            var v166 uint = 0
                            v167 := GenS8{F0: 2147483647}
                            _ = v166
                            _ = v167
                        }
                        switch "hello" {
                        case "hello":
                            v168 := GenS13{F0: 100, F1: 0, F2: v13}
                            _ = len(v1)
                            _ = v168
                        case "":
                            switch 920 {
                            case 1:
                                var v169 uint64 = 631
                                _ = v169
                            }
                        default:
                            switch (fi11 - (fi8 ^ 42)) {
                            case -779:
                                switch "hello" {
                                case "hello":
                                    var v170 bool = true
                                    _ = v170
                                }
                            case 125:
                                fmt.Println(v4, v13, v13)
                            }
                        }
                        v171 := GenS15{F0: -150}
                        v13 -= v13
                        v3 = 1
                        _ = v171
                    } else {
                        v13 += (v13 * 648)
                        var v172 int32 = 42
                        v173 := GenS16{F0: v172, F1: 65535}
                        _ = v172
                        _ = v173
                    }
                    var v174 float32 = 1e10
                    _ = v174
                }
                _ = v13
                _ = v14
            }
            fmt.Println(fi8)
            v6 = uint8(v4)
            _ = fi11
        }
        _ = fi8
    }
    v6 -= (v6 / (((v6 >> ((18446744073709551615) & 63))) | 1))
    v3 = -734
    v3 += v3
    var v175 int8 = -6
    for fi176 := 0; fi176 < 4; fi176++ {
        switch "\n\t" {
        case "a":
            fmt.Println(v6, v3, v4)
        }
        _ = fi176
    }
    for fi177 := 0; fi177 < 16; fi177++ {
        for fi178 := 0; fi178 < 11; fi178++ {
            v4 = v4
            switch (("foo bar" + "") + "hello") {
            case "\n\t":
                fmt.Println(v175, fi178)
                var v179 uint = v4
                switch "foo bar" {
                case "":
                    if true {
                        var v180 int = 42
                        v5[(v180 / ((fi178) | 1))] = fi177
                        var v181 int16 = -1
                        _ = v180
                        _ = v181
                    } else {
                        fmt.Println(v175, v6)
                        switch -713 {
                        case 9223372036854775807:
                            v182 := []int{int(v3)}
                            _ = v182
                        default:
                            v183 := GenS15{F0: 1}
                            _ = v183
                        }
                        v184 := []int{-670, (fi178 | (fi178 << ((v4) & 63))), (fi178 & (fi178 / ((fi178) | 1)))}
                        _ = v184
                    }
                case "hello":
                    v6 = (v6 / ((v6) | 1))
                }
                _ = v179
            case "a":
                if false {
                    v1[2] = 586
                }
            }
            _ = fi178
        }
        v6 = (v6 / (((v6 * v6)) | 1))
        _ = fi177
    }
    v4 = 18446744073709551615
    v185 := GenS17{F0: 932}
    if (!true) {
        var v186 uint32 = 52
        var v187 uint32 = (v186 % ((0) | 1))
        for fi188 := 0; fi188 < 20; fi188++ {
            var v189 uint16 = 65535
            _ = fi188
            _ = v189
        }
        _ = v186
        _ = v187
    }
    var v190 uint8 = v6
    var v191 uint32 = 65535
    v4 -= 42
    for fi192 := 0; fi192 < 15; fi192++ {
        var v193 uint64 = 0
        for fi194 := 0; fi194 < 11; fi194++ {
            var v195 float64 = 0.1
            if ((v193 / ((v193) | 1)) > 4294967295) {
                fmt.Println(v6, fi192)
            } else {
                fmt.Println(fi192, v195, v190)
                v195 += float64(((((float32((float64((float32(v195) / 1e10)) - (v195 / 0.1))) - (1e10 - 0.1)) * 0.0) / (1.5 - (1.5 + -2.25))) / (float32(v195) / float32((float64(3.14159) - 1.5)))))
            }
            v6 = 128
            _ = fi194
            _ = v195
        }
        _ = fi192
        _ = v193
    }
    var v196 int64 = 9223372036854775807
    v190 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v175:", v175)
    fmt.Println("v190:", v190)
    fmt.Println("v191:", v191)
    fmt.Println("v196:", v196)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v175
    _ = v185
    _ = v190
    _ = v191
    _ = v196
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
