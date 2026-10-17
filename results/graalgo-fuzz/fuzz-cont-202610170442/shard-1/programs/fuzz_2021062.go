package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 string }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32; F1 string; F2 int16 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int }

func main() {
    var v1 uint32 = 121
    var v2 uint16 = 1
    v3 := []uint32{v1, (v1 + v1), v1, v1}
    var v4 int16 = -32768
    var v5 uint16 = v2
    var v6 float32 = -2.25
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            switch "" {
            case "hello":
                v4 = -350
            case "foo bar":
                var v9 int16 = int16(v1)
                _ = v9
            case "":
                var v10 uint16 = v5
                _ = v10
            }
            fmt.Println(v4, v5, v1)
            fmt.Println(v6, v5, v6)
            _ = fi8
        }
        v11 := GenS1{F0: 18446744073709551615, F1: (v2 == v5), F2: ("hello" + ("foo bar" + ("\n\t" + "foo bar")))}
        _ = fi7
        _ = v11
    }
    v1 -= (^(v1))
    v12 := GenS2{F0: int8(v1), F1: 943}
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS2{F0: -6, F1: 0}
        fmt.Println(fi13, v2)
        _ = fi13
        _ = v14
    }
    var v15 int = 32767
    var v16 int = v15
    switch -2147483648 {
    case -12:
        if true {
            switch "a" {
            case "foo bar":
                for fi17 := 0; fi17 < 2; fi17++ {
                    fmt.Println(fi17, v2, v1)
                    _ = fi17
                }
            }
            v18 := GenS3{F0: (("hello" + "") + "hello")}
            if ((("hello" == ("\n\t" + "a")) || (!false)) || true) {
                v16 = v16
                if len(v3) > 0 {
                    v3[0] = v1
                }
            }
            _ = v18
        }
    case 32767:
        _ = len(v3)
    default:
        v19 := GenS1{F0: 42, F1: (true && ((((float64((0.0 * ((1e10 - v6) + v6))) / (0.0 + -2.25)) * float64(3.14159)) + float64(1.5)) < (0.1 * (-2.25 - -2.25)))), F2: "a"}
        _ = v19
    }
    if len(v3) > 0 {
        v3[0] = (v1 >> ((18446744073709551615) & 63))
    }
    fmt.Println(v5)
    for fi20 := 0; fi20 < 14; fi20++ {
        v16 = (^(v15))
        _ = fi20
    }
    fmt.Println(v5)
    if true {
        v6 = ((v6 * 3.14159) / v6)
        if true {
            for fi21 := 0; fi21 < 17; fi21++ {
                if false {
                    var v22 uint32 = v1
                    v23 := GenS3{F0: ("" + "hello")}
                    for fi24 := 0; fi24 < 12; fi24++ {
                        var v25 uint16 = uint16(v22)
                        _ = fi24
                        _ = v25
                    }
                    _ = v22
                    _ = v23
                } else {
                    var v26 int32 = 0
                    if false {
                        v27 := GenS3{F0: "a"}
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 3; fi28++ {
                        var v29 uint64 = 0
                        v5 -= (v5 >> ((2) & 63))
                        _ = len(v3)
                        _ = fi28
                        _ = v29
                    }
                    _ = v26
                }
                if ((0.0 * (1e10 / 1e10)) >= ((1.5 - (1.5 - (0.0 * 3.14159))) - float64(v6))) {
                    fmt.Println(v2, v16)
                    switch "foo bar" {
                    case "a":
                        _ = len(v3)
                    default:
                        v15 = (v15 >> ((667) & 63))
                    }
                    v4 = v4
                } else {
                    fmt.Println(fi21, v4, v6)
                    if len(v3) > 0 {
                        v3[0] = (v1 - 277)
                    }
                    for fi30 := 0; fi30 < 17; fi30++ {
                        v2 -= v5
                        switch (-(v16)) {
                        case 1:
                            switch (fi21 >> ((241) & 63)) {
                            case -906:
                                _ = len(v3)
                            }
                        }
                        _ = fi30
                    }
                }
                _ = fi21
            }
        }
    }
    var v31 string = (("foo bar" + "hello") + "foo bar")
    for fi32 := 0; fi32 < 15; fi32++ {
        var v33 int8 = 42
        v16 += fi32
        if ((!(v6 >= v6)) && false) {
            for fi34 := 0; fi34 < 4; fi34++ {
                switch fi32 {
                case -2147483648:
                    v3 = append(v3, v1)
                case -547:
                    if (!false) {
                        v5 = (v5 * v5)
                    }
                case 128:
                    var v35 float32 = ((v6 * (((v6 - 3.14159) / v6) + float32(0.0))) - (1e10 - (v6 * 0.1)))
                    _ = v35
                default:
                    v36 := GenS2{F0: -61, F1: 751}
                    _ = v36
                }
                if false {
                    if (((0.1 + 0.1) - float64(v6)) <= 3.14159) {
                        v4 += 623
                        var v37 int = -9223372036854775808
                        if (true || (uint8(v2) <= 127)) {
                            v1 = v1
                            switch 42 {
                            case 127:
                                switch -9223372036854775808 {
                                case 0:
                                    v38 := GenS4{F0: "hello"}
                                    _ = v38
                                default:
                                    v39 := GenS3{F0: v31}
                                    _ = v39
                                }
                            case -9223372036854775808:
                                v2 -= (v5 << ((4294967295) & 63))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if (24 <= 127) {
                                    if (25 > -1) {
                                        v15 = int(v33)
                                        _ = len(v3)
                                        if ((^(v1)) == v1) {
                                            if (("a" + ("a" + v31)) != (("foo bar" + (v31 + ((("\n\t" + v31) + ((((v31 + v31) + v31) + v31) + ("a" + v31))) + "foo bar"))) + (v31 + ""))) {
                                                switch (fi32 | v37) {
                                                case 2147483647:
                                                    if len(v3) > 0 {
                                                        v3[0] = v1
                                                    }
                                                case 536:
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v40 := GenS2{F0: (v33 << ((2) & 63)), F1: -1}
                                                        _ = v40
                                                    case "":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                }
                                                var v41 float64 = (((0.0 / 1e10) * 0.1) / 1e10)
                                                _ = v41
                                            }
                                        }
                                    }
                                    fmt.Println(fi32)
                                    switch v37 {
                                    case -578:
                                        if true {
                                            fmt.Println(v15)
                                            _ = len(v3)
                                        } else {
                                            var v42 uint32 = (^(v1))
                                            var v43 float64 = (1e10 * float64(v6))
                                            _ = v42
                                            _ = v43
                                        }
                                    default:
                                        if false {
                                            var v44 float32 = v6
                                            _ = v44
                                        } else {
                                            fmt.Println(v2, v33, v37)
                                            if false {
                                                v45 := GenS4{F0: ("a" + "\n\t")}
                                                var v46 uint8 = 127
                                                _ = len(v3)
                                                _ = v45
                                                _ = v46
                                            } else {
                                                fmt.Println(v37, v6, v6)
                                                var v47 bool = (!(-84 >= 1))
                                                _ = v47
                                            }
                                        }
                                    }
                                    v48 := GenS1{F0: 0, F1: (4294967295 < 2), F2: ("foo bar" + (("\n\t" + ("\n\t" + "hello")) + "foo bar"))}
                                    _ = len(v3)
                                    _ = v48
                                }
                            case 436:
                                fmt.Println(fi32, v4, v16)
                            }
                            if true {
                                if len(v3) > 0 {
                                    v3[0] = (v1 | (^(v1)))
                                }
                                v49 := GenS3{F0: "foo bar"}
                                _ = v49
                            } else {
                                var v50 uint64 = 42
                                v51 := GenS5{F0: (("\n\t" + (v31 + (v31 + v31))) + "")}
                                v52 := GenS4{F0: "a"}
                                if ((!(((v4 - v4) != (v4 - v4)) && false)) || false) {
                                    _ = len(v3)
                                    var v53 string = v31
                                    switch v37 {
                                    case 4294967295:
                                        switch -409 {
                                        case -399:
                                            v54 := GenS6{F0: -1, F1: v53, F2: -1}
                                            v16 = v16
                                            _ = len(v3)
                                            _ = v54
                                        default:
                                            v2 -= 854
                                        }
                                    case -987:
                                        if (!(v50 <= 0)) {
                                            v55 := GenS2{F0: (v33 | (v33 << ((18446744073709551615) & 63))), F1: -291}
                                            switch 306 {
                                            case 32767:
                                                switch ("hello" + v31) {
                                                case "a":
                                                    v56 := GenS3{F0: ((v31 + "a") + v31)}
                                                    _ = v56
                                                case "":
                                                    if (!(((v1 >> ((65535) & 63)) < 42) || true)) {
                                                        var v57 uint64 = (v50 ^ v50)
                                                        _ = v57
                                                    } else {
                                                        v33 += v33
                                                        v58 := GenS1{F0: (v50 / (((^(v50))) | 1)), F1: false, F2: ""}
                                                        if len(v3) > 0 {
                                                            v3[0] = 2147483647
                                                        }
                                                        _ = v58
                                                    }
                                                }
                                            case 32768:
                                                var v59 int = 32768
                                                _ = v59
                                            }
                                            _ = v55
                                        }
                                    case -2147483648:
                                        v33 = v33
                                    }
                                    var v60 int64 = -1
                                    var v61 uint64 = (v50 ^ 2)
                                    _ = v53
                                    _ = v60
                                    _ = v61
                                }
                                _ = v50
                                _ = v51
                                _ = v52
                            }
                        }
                        _ = v37
                    }
                    var v62 float64 = 0.1
                    var v63 string = v31
                    _ = v62
                    _ = v63
                } else {
                    v31 = v31
                    var v64 string = "a"
                    _ = v64
                }
                switch (v31 + (v31 + "hello")) {
                case "hello":
                    fmt.Println(fi32, v31, fi34)
                case "\n\t":
                    v3 = append(v3, 2147483648)
                case "a":
                    if true {
                        if (!true) {
                            var v65 uint8 = 145
                            _ = len(v3)
                            if ((v33 + (v33 % (((^(v33))) | 1))) >= v33) {
                                v66 := GenS3{F0: v31}
                                switch 529 {
                                case -1:
                                    v67 := GenS1{F0: 161, F1: true, F2: v31}
                                    _ = v67
                                default:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                fmt.Println(v65, v65, v6)
                                _ = v66
                            } else {
                                var v68 int8 = (v33 >> ((797) & 63))
                                v6 -= 0.1
                                v6 += float32((0.0 * 0.0))
                                _ = v68
                            }
                            _ = v65
                        }
                    } else {
                        var v69 uint = 18446744073709551615
                        _ = v69
                    }
                default:
                    v4 += v4
                }
                v70 := GenS4{F0: "a"}
                fmt.Println(v15, v5)
                _ = fi34
                _ = v70
            }
        } else {
            if true {
                v1 += v1
            } else {
                if true {
                    var v71 uint64 = 42
                    v16 += (v15 + v16)
                    for fi72 := 0; fi72 < 1; fi72++ {
                        v3 = append(v3, (v1 % (((^(v1))) | 1)))
                        v15 += (v16 - -1)
                        if (true && (false && ((v5 % ((v5) | 1)) >= v2))) {
                            switch (v31 + v31) {
                            case "hello":
                                v4 = (^(v4))
                            case "foo bar":
                                if len(v3) > 0 {
                                    v3[0] = v1
                                }
                            case "a":
                                switch v31 {
                                case "\n\t":
                                    var v73 float32 = 0.0
                                    _ = v73
                                }
                                var v74 uint64 = (v71 / ((v71) | 1))
                                var v75 uint16 = (v5 / ((v5) | 1))
                                _ = v74
                                _ = v75
                            }
                            v6 = (v6 + float32(-2.25))
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                var v76 uint = 1
                                var v77 uint16 = v2
                                if (false || ((v33 != v33) && false)) {
                                    if (false && false) {
                                        var v78 int64 = int64(v76)
                                        fmt.Println(v15, v1, v15)
                                        _ = v78
                                    } else {
                                        if (((v1 >> ((v76) & 63)) == 0) && false) {
                                            var v79 int16 = v4
                                            v16 = v15
                                            if (true || (!((v15 % ((4294967295) | 1)) > 116))) {
                                                switch fi72 {
                                                case 9223372036854775807:
                                                    fmt.Println(v1, v1)
                                                }
                                            } else {
                                                var v80 uint8 = 0
                                                var v81 string = v31
                                                _ = v80
                                                _ = v81
                                            }
                                            _ = v79
                                        }
                                    }
                                    if (!(0.1 > v6)) {
                                        if (!true) {
                                            v71 = (^(v71))
                                            if false {
                                                var v82 int8 = v33
                                                switch 0 {
                                                case 4294967295:
                                                    var v83 uint16 = 0
                                                    _ = v83
                                                }
                                                var v84 int32 = 0
                                                if true {
                                                    if (true || false) {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        fmt.Println(v16, fi32, v82)
                                                        if (!((v6 != ((((3.14159 - v6) - float32(((-2.25 - 1e10) * 0.1))) * v6) + 0.0)) && false)) {
                                                            var v85 int = (v15 / (((v15 - v16)) | 1))
                                                            v16 = int(v76)
                                                            _ = v85
                                                        } else {
                                                            var v86 string = v31
                                                            v87 := GenS7{F0: v84, F1: "\n\t"}
                                                            _ = v86
                                                            _ = v87
                                                        }
                                                    } else {
                                                        var v88 float32 = (v6 + 3.14159)
                                                        var v89 int8 = (v82 + -100)
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        _ = v88
                                                        _ = v89
                                                    }
                                                    var v90 int = v16
                                                    _ = v90
                                                }
                                                _ = v82
                                                _ = v84
                                            }
                                            switch (fi72 * int(v33)) {
                                            case 32768:
                                                var v91 string = v31
                                                _ = v91
                                            case -880:
                                                var v92 uint = (^(v76))
                                                _ = v92
                                            case 32767:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v93 float32 = ((((float32(3.14159) * v6) + (0.1 * v6)) * v6) / 1.0)
                                                _ = v93
                                            }
                                            var v94 int32 = 0
                                            v95 := GenS2{F0: (^(v33)), F1: v94}
                                            _ = v94
                                            _ = v95
                                        } else {
                                            v5 += v2
                                        }
                                        var v96 bool = true
                                        switch (v16 << (((v76 / ((255) | 1))) & 63)) {
                                        case 384:
                                            v4 = (v4 << ((v76) & 63))
                                        case -2147483648:
                                            v3 = append(v3, v1)
                                        default:
                                            fmt.Println(v76, v2)
                                        }
                                        _ = v96
                                    }
                                }
                                fmt.Println(fi32)
                                if (-1 >= 42) {
                                    var v97 uint8 = 128
                                    v76 += 4294967295
                                    _ = v97
                                } else {
                                    v98 := GenS1{F0: 0, F1: true, F2: "hello"}
                                    _ = v98
                                }
                                _ = v76
                                _ = v77
                            }
                            v3 = append(v3, v1)
                        }
                        switch (fi32 % ((v15) | 1)) {
                        case -2147483648:
                            switch v31 {
                            case "foo bar":
                                var v99 uint32 = v1
                                _ = v99
                            case "":
                                switch ("\n\t" + v31) {
                                case "":
                                    switch (v16 | fi32) {
                                    case 2147483647:
                                        if len(v3) > 0 {
                                            v3[0] = 387
                                        }
                                    default:
                                        v100 := GenS7{F0: 2147483647, F1: v31}
                                        _ = v100
                                    }
                                case "foo bar":
                                    v3 = append(v3, (^(v1)))
                                    if (((^(v1)) < (v1 + v1)) && (v31 <= "")) {
                                        var v101 string = v31
                                        v3 = append(v3, v1)
                                        v2 -= (v2 & v2)
                                        v102 := GenS5{F0: ("\n\t" + (("a" + "hello") + (v101 + "hello")))}
                                        _ = v101
                                        _ = v102
                                    } else {
                                        v5 -= v5
                                        switch 0 {
                                        case 671:
                                            switch v31 {
                                            case "\n\t":
                                                if ((v1 ^ (v1 >> ((255) & 63))) != v1) {
                                                    v33 += (v33 - v33)
                                                    var v103 float32 = (v6 - v6)
                                                    var v104 float64 = (((0.1 + (1.5 - (-2.25 / ((1e10 - (((((float64(((0.0 + v6) / v6)) - -2.25) + 1.5) * 0.1) + (1e10 * ((-2.25 - 1.5) / 1.5))) + 0.0)) + 0.0)))) * 3.14159) - float64(v103))
                                                    _ = v103
                                                    _ = v104
                                                } else {
                                                    v3 = append(v3, (v1 * (v1 + (v1 / ((v1) | 1)))))
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v4 = int16(fi72)
                                                }
                                            case "a":
                                                var v105 bool = false
                                                _ = v105
                                            case "foo bar":
                                                v15 = 860
                                            }
                                        case -1:
                                            if (false || (!(false && true))) {
                                                var v106 uint32 = 1
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                v107 := GenS3{F0: v31}
                                                _ = v106
                                                _ = v107
                                            }
                                        case -678:
                                            switch v31 {
                                            case "a":
                                                v1 = 116
                                            }
                                        }
                                        v6 = v6
                                    }
                                    if true {
                                        v31 = v31
                                    }
                                case "a":
                                    var v108 string = ((v31 + ("" + ("\n\t" + v31))) + v31)
                                    _ = v108
                                }
                            default:
                                v16 = v16
                                v33 += v33
                                v109 := GenS1{F0: v71, F1: true, F2: v31}
                                _ = v109
                            }
                        case 33:
                            switch fi32 {
                            case -463:
                                _ = len(v3)
                            case 32767:
                                if true {
                                    if ((0.0 + -2.25) < (1e10 / 0.1)) {
                                        if false {
                                            var v110 int32 = 42
                                            v2 = uint16(v1)
                                            fmt.Println(v4)
                                            _ = v110
                                        } else {
                                            if true {
                                                fmt.Println(v1)
                                                if len(v3) > 0 {
                                                    v3[0] = v1
                                                }
                                                v111 := GenS6{F0: int32(v1), F1: "a", F2: -262}
                                                _ = v111
                                            } else {
                                                if (!(-2147483648 < 2147483647)) {
                                                    switch (-(fi72)) {
                                                    case 885:
                                                        var v112 float32 = v6
                                                        _ = v112
                                                    case 9223372036854775807:
                                                        v113 := GenS1{F0: (v71 | v71), F1: (!(false && ((-2.25 - -2.25) != (0.0 - (-2.25 / (-2.25 - 3.14159)))))), F2: "\n\t"}
                                                        if ((((v6 - float32(0.0)) / 3.14159) + 3.14159) > v6) {
                                                            v114 := GenS2{F0: v33, F1: -2147483648}
                                                            var v115 uint8 = 255
                                                            v3 = append(v3, 820)
                                                            _ = v114
                                                            _ = v115
                                                        } else {
                                                            v2 += (v5 >> ((0) & 63))
                                                        }
                                                        _ = v113
                                                    }
                                                }
                                                var v116 int8 = (v33 % (((-(v33))) | 1))
                                                var v117 uint16 = 255
                                                var v118 uint8 = 42
                                                v71 = v71
                                                _ = v116
                                                _ = v117
                                                _ = v118
                                            }
                                            var v119 uint = 0
                                            _ = v119
                                        }
                                        var v120 uint8 = 1
                                        if len(v3) > 0 {
                                            v3[0] = v1
                                        }
                                        _ = v120
                                    } else {
                                        var v121 int8 = v33
                                        _ = v121
                                    }
                                    var v122 uint64 = (v71 * uint64(v5))
                                    _ = v122
                                }
                            case -577:
                                v123 := GenS2{F0: (v33 | 42), F1: 2147483647}
                                _ = v123
                            }
                            v124 := GenS3{F0: v31}
                            _ = v124
                        case 2147483647:
                            var v125 int32 = -2147483648
                            _ = v125
                        default:
                            if len(v3) > 0 {
                                v3[0] = v1
                            }
                        }
                        switch -2147483648 {
                        case -928:
                            var v126 int16 = (-(v4))
                            _ = v126
                        case -1:
                            v3 = append(v3, v1)
                        }
                        _ = fi72
                    }
                    _ = v71
                }
                v127 := GenS7{F0: -436, F1: ((v31 + v31) + v31)}
                for fi128 := 0; fi128 < 16; fi128++ {
                    v129 := GenS8{F0: v4}
                    fmt.Println(v2, fi32, v31)
                    _ = fi128
                    _ = v129
                }
                _ = v127
            }
            if true {
                v4 -= -32768
                var v130 uint64 = 411
                for fi131 := 0; fi131 < 14; fi131++ {
                    if true {
                        switch (fi131 ^ v15) {
                        case 4294967295:
                            v3 = append(v3, v1)
                        case 0:
                            if ((v6 + ((v6 / (float32(((1.5 - float64(v6)) - 0.0)) - v6)) * v6)) <= v6) {
                                var v132 float64 = (-2.25 / 0.1)
                                v33 -= v33
                                _ = v132
                            } else {
                                v4 += (-(v4))
                            }
                            if len(v3) > 0 {
                                v3[0] = v1
                            }
                        case -673:
                            var v133 bool = (!false)
                            _ = v133
                        default:
                            var v134 int64 = 9223372036854775807
                            _ = v134
                        }
                        v33 += (v33 - v33)
                        v16 = int(v33)
                    } else {
                        if (!false) {
                            var v135 int8 = (v33 | (^(v33)))
                            _ = v135
                        } else {
                            var v136 uint16 = (v5 * v2)
                            v5 = v2
                            _ = v136
                        }
                        v16 -= fi32
                    }
                    fmt.Println(fi32, v16, v31)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi131
                }
                _ = v130
            } else {
                var v137 int32 = 1
                _ = len(v3)
                v138 := GenS9{F0: (fi32 << ((64) & 63))}
                _ = v137
                _ = v138
            }
        }
        _ = fi32
        _ = v33
    }
    v139 := GenS1{F0: 2, F1: false, F2: (v31 + ("\n\t" + "foo bar"))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v139
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
