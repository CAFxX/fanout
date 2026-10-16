package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int64 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 uint64; F2 int64 }
type GenS7 struct { F0 uint16; F1 int; F2 uint32 }
type GenS8 struct { F0 uint8; F1 int; F2 uint32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint32; F1 int16 }
type GenS14 struct { F0 int16; F1 int32 }
type GenS15 struct { F0 uint8 }

func main() {
    var v1 int64 = 1
    var v2 float64 = -2.25
    v3 := []int{0}
    var v4 int64 = v1
    var v5 int64 = (v4 * (v4 * v4))
    _ = len(v3)
    if false {
        if (false && false) {
            var v6 string = "\n\t"
            fmt.Println(v5)
            v5 = 0
            _ = v6
        } else {
            if (false || (2147483647 >= 127)) {
                var v7 int64 = 9223372036854775807
                _ = v7
            } else {
                fmt.Println(v1)
                for fi8 := 0; fi8 < 17; fi8++ {
                    v9 := GenS1{F0: 359, F1: 671, F2: v4}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi8
                    _ = v9
                }
            }
            v4 = (^(v1))
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch 0 {
    case 127:
        if (true || false) {
            v10 := GenS2{F0: (("" + "") + "a"), F1: ("a" + (((("hello" + ("\n\t" + (("hello" + (("a" + ("hello" + "a")) + "a")) + ("foo bar" + "hello")))) + "a") + "") + ""))}
            _ = v10
        } else {
            var v11 uint32 = 115
            _ = v11
        }
    case -2147483648:
        v12 := GenS3{F0: ((128 != 227) && (128 < 204))}
        _ = v12
    default:
        var v13 uint64 = 4294967295
        _ = v13
    }
    v1 -= v4
    if false {
        var v14 string = "\n\t"
        var v15 uint32 = 1
        _ = v14
        _ = v15
    } else {
        v16 := GenS2{F0: ("" + (("foo bar" + "") + "\n\t")), F1: "a"}
        _ = v16
    }
    if false {
        v2 = (v2 + v2)
        if false {
            v5 -= v1
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 uint = 835
                var v19 int16 = -592
                if ((int64(v19) >= 0) || true) {
                    fmt.Println(v1, v2, v19)
                    for fi20 := 0; fi20 < 14; fi20++ {
                        switch fi17 {
                        case 832:
                            _ = len(v3)
                        }
                        _ = fi20
                    }
                } else {
                    v2 += 0.1
                }
                _ = fi17
                _ = v18
                _ = v19
            }
            for fi21 := 0; fi21 < 3; fi21++ {
                for fi22 := 0; fi22 < 11; fi22++ {
                    v5 += v4
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v23 int8 = -2
                    v23 = v23
                    var v24 int8 = int8(fi22)
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                for fi25 := 0; fi25 < 2; fi25++ {
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v4, v2)
                    case "hello":
                        var v26 int = 32768
                        _ = v26
                    default:
                        var v27 uint32 = 0
                        _ = v27
                    }
                    switch (("" + "a") + "hello") {
                    case "foo bar":
                        v1 += 9223372036854775807
                    case "\n\t":
                        switch ("a" + "hello") {
                        case "":
                            if true {
                                v28 := GenS2{F0: "foo bar", F1: (("" + "\n\t") + (("" + "a") + (("\n\t" + "\n\t") + "\n\t")))}
                                var v29 float32 = ((-2.25 - ((0.1 / ((3.14159 * (-2.25 + 0.1)) * 0.1)) - (0.1 + (3.14159 / 1e10)))) + 1e10)
                                _ = v28
                                _ = v29
                            } else {
                                v1 -= (^(v1))
                                fmt.Println(v1, v4)
                                v30 := GenS1{F0: -9223372036854775808, F1: 32767, F2: -9223372036854775808}
                                _ = v30
                            }
                        case "a":
                            switch "foo bar" {
                            case "a":
                                fmt.Println(fi25)
                            case "foo bar":
                                if ((false || (4294967295 == 18446744073709551615)) || (fi21 < (fi25 / (((fi21 & fi25)) | 1)))) {
                                    v3 = append(v3, 127)
                                    var v31 uint16 = 32767
                                    _ = v31
                                }
                            case "hello":
                                if true {
                                    v4 += (v5 / (((v1 / ((v5) | 1))) | 1))
                                    var v32 uint8 = 128
                                    _ = v32
                                }
                            }
                        }
                    default:
                        if false {
                            switch "hello" {
                            case "a":
                                if (!false) {
                                    var v33 uint16 = 255
                                    v34 := GenS4{F0: ("" + "a"), F1: 978}
                                    fmt.Println(v5, v5, v33)
                                    switch 887 {
                                    case 32768:
                                        v35 := GenS5{F0: 1}
                                        _ = v35
                                    case -9223372036854775808:
                                        if false {
                                            if true {
                                                fmt.Println(v33, v33)
                                                var v36 uint = 0
                                                _ = v36
                                            } else {
                                                v4 += v1
                                                if true {
                                                    if len(v3) > 0 {
                                                        v3[0] = fi25
                                                    }
                                                    switch "foo bar" {
                                                    case "":
                                                        v3 = append(v3, (fi25 + fi21))
                                                        var v37 uint = 838
                                                        switch "" {
                                                        case "":
                                                            if true {
                                                                if false {
                                                                    var v38 uint = v37
                                                                    _ = v38
                                                                } else {
                                                                    v2 += v2
                                                                    var v39 uint16 = 0
                                                                    _ = v39
                                                                }
                                                                fmt.Println(v37, fi21)
                                                            }
                                                        case "foo bar":
                                                            if true {
                                                                _ = len(v3)
                                                                if true {
                                                                    v2 -= (1e10 * 1.5)
                                                                } else {
                                                                    if true {
                                                                        switch (fi25 | 128) {
                                                                        case 32768:
                                                                            fmt.Println(v1, v4, fi21)
                                                                        case 4294967295:
                                                                            if false {
                                                                                fmt.Println(v2, v5, v2)
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = (-(fi21))
                                                                                }
                                                                            }
                                                                        case 2147483647:
                                                                            if true {
                                                                                v40 := GenS6{F0: (v4 | (-(v4))), F1: 238, F2: 1}
                                                                                fmt.Println(v37, v2, v1)
                                                                                _ = v40
                                                                            }
                                                                            v41 := GenS2{F0: "foo bar", F1: ""}
                                                                            switch 519 {
                                                                            case 32768:
                                                                                v37 = v37
                                                                            }
                                                                            _ = v41
                                                                        }
                                                                        fmt.Println(v1)
                                                                    } else {
                                                                        var v42 float64 = v2
                                                                        _ = v42
                                                                    }
                                                                    switch fi25 {
                                                                    case -1:
                                                                        v43 := GenS7{F0: 748, F1: (fi21 % (((-(fi21))) | 1)), F2: 65535}
                                                                        _ = v43
                                                                    }
                                                                    v44 := GenS1{F0: (v4 ^ v5), F1: (v33 + v33), F2: (v5 ^ (v4 << ((v37) & 63)))}
                                                                    v37 += (v37 + 18446744073709551615)
                                                                    switch (fi25 << (((v37 - v37)) & 63)) {
                                                                    case 32767:
                                                                        var v45 int16 = 42
                                                                        _ = v45
                                                                    default:
                                                                        var v46 int = fi25
                                                                        _ = v46
                                                                    }
                                                                    fmt.Println(v4)
                                                                    _ = v44
                                                                }
                                                            } else {
                                                                if len(v3) > 0 {
                                                                    v3[0] = (fi21 + -683)
                                                                }
                                                                if (-65 > -1) {
                                                                    if false {
                                                                        fmt.Println(v37, v33, v1)
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            v4 = 765
                                                                        }
                                                                    }
                                                                } else {
                                                                    v2 -= (3.14159 - v2)
                                                                    var v47 uint64 = 1
                                                                    var v48 string = ("a" + "a")
                                                                    _ = v47
                                                                    _ = v48
                                                                }
                                                            }
                                                        }
                                                        v33 += (v33 - v33)
                                                        _ = v37
                                                    case "a":
                                                        var v49 int32 = -1
                                                        _ = v49
                                                    case "hello":
                                                        if (18446744073709551615 < 4294967295) {
                                                            if (!(128 != 127)) {
                                                                var v50 uint = 1
                                                                var v51 uint32 = 706
                                                                var v52 bool = false
                                                                var v53 int8 = 81
                                                                _ = v50
                                                                _ = v51
                                                                _ = v52
                                                                _ = v53
                                                            } else {
                                                                v54 := GenS1{F0: (v4 + v5), F1: v33, F2: int64(v33)}
                                                                _ = v54
                                                            }
                                                            v3 = append(v3, (fi25 * 127))
                                                            if (-51 <= -1) {
                                                                v55 := GenS8{F0: 2, F1: (fi25 & -2147483648), F2: 4294967295}
                                                                _ = v55
                                                            }
                                                        }
                                                    }
                                                }
                                                switch "a" {
                                                case "":
                                                    v33 = 699
                                                case "hello":
                                                    var v56 bool = ((v4 >= (v5 - 9223372036854775807)) || true)
                                                    _ = v56
                                                case "a":
                                                    var v57 string = "a"
                                                    _ = v57
                                                }
                                                v33 -= 255
                                            }
                                            fmt.Println(fi21)
                                            _ = len(v3)
                                        }
                                    case -1:
                                        if (!false) {
                                            _ = len(v3)
                                            if true {
                                                switch fi21 {
                                                case 0:
                                                    v33 = (v33 & (^(v33)))
                                                }
                                                var v58 int64 = (v4 ^ (v5 / ((v5) | 1)))
                                                if (42 > -1) {
                                                    if true {
                                                        var v59 uint8 = 42
                                                        var v60 float32 = float32((v2 + float64(1e10)))
                                                        _ = v59
                                                        _ = v60
                                                    } else {
                                                        v1 = v58
                                                        var v61 int = (fi25 / ((-867) | 1))
                                                        v62 := GenS7{F0: 0, F1: 424, F2: 0}
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                    if true {
                                                        v3 = append(v3, (-(fi21)))
                                                    }
                                                }
                                                fmt.Println(v1)
                                                switch 32768 {
                                                case 0:
                                                    var v63 uint64 = 148
                                                    _ = v63
                                                case -2147483648:
                                                    v64 := GenS7{F0: v33, F1: fi25, F2: 65535}
                                                    _ = v64
                                                case 42:
                                                    v5 += (v4 ^ (v1 - (^(v4))))
                                                default:
                                                    v33 -= (v33 >> ((0) & 63))
                                                }
                                                _ = v58
                                            }
                                            var v65 uint8 = 1
                                            if true {
                                                switch (fi25 ^ -233) {
                                                case 1:
                                                    v66 := GenS1{F0: 605, F1: (^(v33)), F2: (^(v4))}
                                                    _ = v66
                                                }
                                                var v67 string = "a"
                                                v1 += (-(v5))
                                                _ = v67
                                            } else {
                                                v1 = 752
                                                fmt.Println(fi25, v1)
                                            }
                                            _ = v65
                                        }
                                    default:
                                        v68 := GenS3{F0: (!(!(127 < -49)))}
                                        _ = v68
                                    }
                                    var v69 bool = (25 <= 42)
                                    _ = len(v3)
                                    _ = v33
                                    _ = v34
                                    _ = v69
                                }
                            case "foo bar":
                                if false {
                                    if (233 >= 127) {
                                        v4 = -9223372036854775808
                                    }
                                }
                            case "":
                                v1 -= (v4 >> ((0) & 63))
                            default:
                                var v70 uint32 = 65535
                                _ = v70
                            }
                        }
                    }
                    _ = fi25
                }
                if true {
                    if (false && true) {
                        for fi71 := 0; fi71 < 7; fi71++ {
                            var v72 float64 = (((3.14159 + (v2 / v2)) / v2) + v2)
                            v1 -= v5
                            _ = fi71
                            _ = v72
                        }
                        v73 := GenS5{F0: 74}
                        _ = v73
                    }
                    for fi74 := 0; fi74 < 13; fi74++ {
                        var v75 int16 = 320
                        if (65535 >= 0) {
                            var v76 uint32 = 65535
                            if (false || ((!false) || (fi74 >= (fi74 << ((255) & 63))))) {
                                v2 -= float64(1e10)
                                if len(v3) > 0 {
                                    v3[0] = 1
                                }
                                switch fi74 {
                                case 626:
                                    v2 = float64(float32(v2))
                                default:
                                    var v77 int16 = (v75 ^ -505)
                                    _ = v77
                                }
                                var v78 uint16 = 849
                                _ = v78
                            } else {
                                var v79 int32 = -254
                                var v80 bool = (true || (false || (((-2.25 - 0.1) <= v2) || ((^(fi21)) != (fi21 * -2147483648)))))
                                var v81 uint = 0
                                _ = v79
                                _ = v80
                                _ = v81
                            }
                            v75 = v75
                            _ = v76
                        }
                        if false {
                            var v82 uint16 = 65535
                            v5 -= v4
                            v2 = (((0.0 * v2) * v2) - v2)
                            if false {
                                v4 += 9223372036854775807
                                v82 -= v82
                                v5 = (v5 & (v5 - (v5 & (v5 * (v1 | v1)))))
                            }
                            _ = v82
                        }
                        if false {
                            var v83 float64 = (1e10 - (float64(3.14159) - v2))
                            v84 := GenS2{F0: "", F1: ("" + "foo bar")}
                            _ = v83
                            _ = v84
                        } else {
                            switch "hello" {
                            case "hello":
                                var v85 int = (fi74 % (((fi21 << ((255) & 63))) | 1))
                                _ = v85
                            case "\n\t":
                                v5 += -341
                                var v86 int16 = v75
                                v86 += (v75 % ((v86) | 1))
                                _ = v86
                            case "":
                                switch "\n\t" {
                                case "":
                                    v4 = (^(v1))
                                default:
                                    fmt.Println(v2)
                                }
                            }
                            var v87 string = ("\n\t" + (("\n\t" + "a") + "foo bar"))
                            fmt.Println(v75)
                            if ((0.1 - 1e10) >= (float32((v2 * (3.14159 + v2))) / -2.25)) {
                                var v88 uint32 = 42
                                fmt.Println(v5, fi74)
                                _ = len(v3)
                                switch 32768 {
                                case -2147483648:
                                    var v89 bool = (true && true)
                                    _ = v89
                                case 42:
                                    if len(v3) > 0 {
                                        v3[0] = fi21
                                    }
                                case 127:
                                    fmt.Println(v2, v2, v2)
                                }
                                _ = v88
                            } else {
                                v90 := GenS6{F0: v5, F1: 1, F2: (v4 | v4)}
                                if len(v3) > 0 {
                                    v3[0] = 226
                                }
                                _ = v90
                            }
                            var v91 int64 = (v4 / (((v5 | v1)) | 1))
                            if true {
                                var v92 uint32 = 1
                                if ((fi21 - fi21) > fi74) {
                                    var v93 bool = true
                                    _ = v93
                                }
                                v94 := GenS9{F0: "\n\t"}
                                _ = v92
                                _ = v94
                            } else {
                                v75 += 617
                                switch fi74 {
                                case -1:
                                    v1 += v4
                                case -2147483648:
                                    v95 := GenS3{F0: (0.1 != ((0.0 * 0.0) * 0.1))}
                                    _ = v95
                                case 32768:
                                    fmt.Println(fi21, v5)
                                default:
                                    switch (v87 + "hello") {
                                    case "hello":
                                        var v96 float32 = (1e10 - 1e10)
                                        _ = v96
                                    case "\n\t":
                                        if (!("\n\t" <= "\n\t")) {
                                            var v97 int8 = -50
                                            _ = v97
                                        } else {
                                            v98 := GenS1{F0: (v1 & (v91 ^ 42)), F1: 798, F2: 558}
                                            _ = v98
                                        }
                                    default:
                                        if ((v4 - v1) <= 42) {
                                            v99 := GenS7{F0: 91, F1: (fi21 & fi21), F2: 628}
                                            _ = v99
                                        }
                                    }
                                }
                            }
                            _ = v87
                            _ = v91
                        }
                        switch ((("hello" + "hello") + "foo bar") + "") {
                        case "hello":
                            fmt.Println(fi21, v4, v5)
                        case "\n\t":
                            if ((fi74 | (fi74 | (^(fi74)))) != fi21) {
                                if false {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v100 := GenS10{F0: (v75 / (((v75 / ((int16(fi74)) | 1))) | 1))}
                                    _ = v100
                                }
                                v101 := GenS7{F0: 32767, F1: fi74, F2: 572}
                                fmt.Println(v5, fi74, fi21)
                                var v102 uint = 352
                                _ = len(v3)
                                _ = v101
                                _ = v102
                            }
                        }
                        _ = fi74
                        _ = v75
                    }
                    for fi103 := 0; fi103 < 4; fi103++ {
                        var v104 float64 = float64((-2.25 / 3.14159))
                        _ = fi103
                        _ = v104
                    }
                } else {
                    v105 := GenS11{F0: 18446744073709551615}
                    _ = v105
                }
                _ = fi21
            }
        } else {
            v3 = append(v3, -557)
            v106 := GenS12{F0: 127}
            _ = v106
        }
    }
    switch (((("\n\t" + "\n\t") + ("a" + "")) + (("hello" + "hello") + (("a" + "\n\t") + "foo bar"))) + "") {
    case "\n\t":
        v2 = ((float64((float32(3.14159) + -2.25)) + (float64(0.1) / (0.0 - v2))) * v2)
    case "hello":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        switch 32768 {
        case 32768:
            v107 := GenS6{F0: v5, F1: 18446744073709551615, F2: v4}
            _ = v107
        case 32767:
            switch 0 {
            case -9223372036854775808:
                var v108 uint8 = 12
                _ = v108
            case 0:
                fmt.Println(v4, v2, v1)
            case 346:
                v5 -= v1
            }
        default:
            v2 += -2.25
        }
    }
    v5 = 1
    v1 -= (-(v4))
    for fi109 := 0; fi109 < 2; fi109++ {
        switch (fi109 & (-(fi109))) {
        case -720:
            v110 := GenS5{F0: 189}
            _ = v110
        case 946:
            _ = len(v3)
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi109
    }
    for fi111 := 0; fi111 < 9; fi111++ {
        v5 -= v1
        _ = fi111
    }
    for fi112 := 0; fi112 < 13; fi112++ {
        _ = len(v3)
        if (false || true) {
            v5 = v1
        } else {
            for fi113 := 0; fi113 < 5; fi113++ {
                if ((3.14159 - (0.0 - float32(v2))) != -2.25) {
                    var v114 bool = (42 == (v4 & v1))
                    if v114 {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if v114 {
                            if (v114 && v114) {
                                var v115 uint64 = 0
                                _ = v115
                            } else {
                                if v114 {
                                    fmt.Println(v4, v5, fi112)
                                    v2 = float64((float32((0.0 - 3.14159)) / float32(v2)))
                                }
                            }
                            var v116 float64 = float64(1e10)
                            _ = v116
                        } else {
                            switch ("foo bar" + ("" + "a")) {
                            case "\n\t":
                                var v117 uint8 = 128
                                _ = v117
                            case "":
                                switch 2147483647 {
                                case 271:
                                    var v118 float32 = 1.5
                                    _ = v118
                                case 0:
                                    _ = len(v3)
                                case 9223372036854775807:
                                    var v119 int8 = 0
                                    v120 := GenS2{F0: "\n\t", F1: "a"}
                                    _ = v119
                                    _ = v120
                                default:
                                    var v121 float32 = 0.1
                                    _ = v121
                                }
                            case "foo bar":
                                switch fi113 {
                                case 2147483647:
                                    v122 := GenS13{F0: 605, F1: -1}
                                    _ = v122
                                default:
                                    v123 := GenS8{F0: 128, F1: -845, F2: 2147483647}
                                    _ = v123
                                }
                            }
                            v4 = -431
                            v5 += v5
                        }
                        var v124 uint = 42
                        _ = v124
                    }
                    _ = v114
                } else {
                    if (!true) {
                        v5 -= v4
                        fmt.Println(v4)
                    }
                }
                var v125 int64 = (v1 % (((v5 - v4)) | 1))
                _ = fi113
                _ = v125
            }
            v3 = append(v3, (fi112 >> ((4294967295) & 63)))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi112
    }
    v5 = 375
    var v126 bool = ((-699 >= 871) || (0 != 252))
    v5 = (^(v4))
    var v127 uint64 = 42
    for fi128 := 0; fi128 < 5; fi128++ {
        switch "hello" {
        case "hello":
            v129 := GenS14{F0: 625, F1: -1}
            _ = v129
        case "foo bar":
            var v130 uint8 = 245
            _ = v130
        default:
            v1 -= v5
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi131 := 0; fi131 < 10; fi131++ {
            if (32 <= 81) {
                if (!true) {
                    switch fi128 {
                    case 304:
                        if len(v3) > 0 {
                            v3[0] = (fi128 ^ 42)
                        }
                    case -825:
                        switch -923 {
                        case 127:
                            var v132 float64 = (v2 + v2)
                            _ = v132
                        case -9223372036854775808:
                            if true {
                                fmt.Println(v127, v4, v127)
                                fmt.Println(v4, v1, fi128)
                                fmt.Println(v5)
                                v2 = v2
                            }
                        default:
                            if true {
                                v127 -= 2
                                v133 := GenS11{F0: 1}
                                _ = v133
                            } else {
                                fmt.Println(v2, v127, v126)
                            }
                        }
                    case 128:
                        v134 := GenS15{F0: 127}
                        _ = v134
                    default:
                        var v135 string = "a"
                        _ = v135
                    }
                    v3 = append(v3, 32768)
                    if (!(!(307 >= 42))) {
                        v136 := GenS6{F0: v1, F1: 42, F2: 9223372036854775807}
                        _ = v136
                    } else {
                        switch "hello" {
                        case "hello":
                            var v137 int64 = v5
                            _ = v137
                        default:
                            var v138 bool = (!(((fi131 * (-(fi128))) != fi131) || v126))
                            _ = v138
                        }
                        var v139 int16 = -117
                        _ = v139
                    }
                    if v126 {
                        var v140 int8 = 100
                        var v141 uint32 = 42
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v140
                        _ = v141
                    }
                    v5 = (v4 ^ v4)
                }
                v126 = (18446744073709551615 >= 396)
                v5 = v5
                switch ((("a" + ("\n\t" + "hello")) + ("hello" + "")) + ((("foo bar" + "") + "foo bar") + ("hello" + "\n\t"))) {
                case "a":
                    v142 := GenS4{F0: ((("foo bar" + "hello") + "\n\t") + "foo bar"), F1: 2}
                    _ = v142
                case "":
                    v5 += v5
                case "hello":
                    fmt.Println(v126, v127)
                    fmt.Println(fi128)
                    var v143 string = ("\n\t" + "foo bar")
                    _ = v143
                }
                v3 = append(v3, (fi128 ^ (-(fi131))))
            }
            _ = fi131
        }
        _ = fi128
    }
    if len(v3) > 0 {
        v3[0] = -2147483648
    }
    fmt.Println("v1:", v1)
    fmt.Println("v126:", v126)
    fmt.Println("v127:", v127)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v126
    _ = v127
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
