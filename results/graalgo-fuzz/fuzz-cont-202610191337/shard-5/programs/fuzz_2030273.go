package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS4 struct { F0 string; F1 bool; F2 int64 }
type GenS5 struct { F0 int32; F1 uint32; F2 uint }
type GenS6 struct { F0 uint64; F1 int; F2 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32; F1 bool; F2 uint }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint32; F1 string }
type GenS11 struct { F0 uint8; F1 uint16; F2 int }
type GenS12 struct { F0 int32; F1 string; F2 bool }
type GenS13 struct { F0 string; F1 uint32 }
type GenS14 struct { F0 int16; F1 uint }
type GenS15 struct { F0 bool; F1 int32 }
type GenS16 struct { F0 uint64; F1 bool; F2 string }
type GenS17 struct { F0 bool; F1 int8; F2 uint }
type GenS18 struct { F0 string; F1 bool; F2 int }
type GenS19 struct { F0 uint64; F1 string; F2 int16 }
type GenS20 struct { F0 uint8; F1 uint32 }
type GenS21 struct { F0 int8; F1 uint16 }
type GenS22 struct { F0 int8; F1 bool }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    var v4 int64 = -1
    var v5 uint8 = 42
    v6 := []int{1, 2147483647}
    for fi7 := 0; fi7 < 14; fi7++ {
        v4 -= 496
        var v8 float64 = ((1e10 * (1.5 - -2.25)) + 1.5)
        if false {
            var v9 bool = false
            _ = v9
        } else {
            var v10 uint16 = 255
            v11 := GenS1{F0: -21, F1: ((((("hello" + "\n\t") + ("hello" + "")) + "") + "\n\t") + "")}
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        v6 = append(v6, (fi12 >> ((0) & 63)))
        _ = fi12
    }
    switch 9223372036854775807 {
    case 4294967295:
        _ = len(v1)
    case 1:
        _ = len(v6)
    case 32767:
        for fi13 := 0; fi13 < 10; fi13++ {
            fmt.Println(fi13)
            _ = fi13
        }
    default:
        for fi14 := 0; fi14 < 10; fi14++ {
            if (false && true) {
                switch (-(fi14)) {
                case 42:
                    fmt.Println(v4)
                }
                v4 = v4
            }
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 uint8 = 235
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    }
    switch 32768 {
    case 32767:
        switch "\n\t" {
        case "a":
            var v17 int64 = -1
            _ = v17
        case "hello":
            if false {
                v4 = (v4 & 411)
            } else {
                v5 = v5
                var v18 float64 = (((float64(float32((1e10 / ((1.5 + float64(1e10)) * 0.1)))) / 1.5) / 1.5) / 1e10)
                if true {
                    delete(v1, "\n\t")
                    if len(v6) > 0 {
                        v6[0] = 944
                    }
                    v19 := GenS2{F0: 9223372036854775807, F1: 359, F2: 318}
                    _ = v19
                }
                _ = v18
            }
        case "":
            if false {
                switch -1 {
                case 4294967295:
                    var v20 int8 = 1
                    _ = v20
                case 32767:
                    var v21 uint32 = 825
                    _ = v21
                case 2147483647:
                    delete(v2, 1)
                }
                for fi22 := 0; fi22 < 16; fi22++ {
                    fmt.Println(fi22)
                    v4 = (v4 / ((1) | 1))
                    _ = fi22
                }
                if ((v5 / (((^(v5))) | 1)) <= (^(v5))) {
                    for fi23 := 0; fi23 < 17; fi23++ {
                        v4 += (v4 | v4)
                        v4 += (v4 & v4)
                        _ = fi23
                    }
                } else {
                    if false {
                        for fi24 := 0; fi24 < 10; fi24++ {
                            for fi25 := 0; fi25 < 9; fi25++ {
                                if false {
                                    var v26 float64 = 0.1
                                    v26 = float64(3.14159)
                                    var v27 int = fi24
                                    _ = v26
                                    _ = v27
                                } else {
                                    if false {
                                        var v28 bool = (!false)
                                        var v29 bool = (v28 && (773 != 0))
                                        if v29 {
                                            v29 = v28
                                            if v29 {
                                                v4 = (-(v4))
                                                v28 = v28
                                            } else {
                                                v29 = false
                                                v5 -= (v5 + v5)
                                                fmt.Println(fi24)
                                            }
                                        } else {
                                            v29 = (v28 && v28)
                                            var v30 uint16 = 0
                                            _ = v30
                                        }
                                        _ = v28
                                        _ = v29
                                    }
                                }
                                v31 := GenS1{F0: 82, F1: ("foo bar" + "hello")}
                                _ = fi25
                                _ = v31
                            }
                            _ = fi24
                        }
                    } else {
                        fmt.Println(v5, v5)
                    }
                }
            } else {
                switch "" {
                case "\n\t":
                    switch 42 {
                    case 137:
                        v5 = 128
                    case 4294967295:
                        if (((((float64(0.0) - (0.0 - 3.14159)) * float64(((1.5 - 0.1) / (float32(0.0) + 3.14159)))) - -2.25) <= 1.5) || false) {
                            for fi32 := 0; fi32 < 1; fi32++ {
                                for fi33 := 0; fi33 < 15; fi33++ {
                                    if (977 <= 42) {
                                        fmt.Println(fi32, fi32, fi33)
                                    } else {
                                        v5 -= v5
                                        switch "" {
                                        case "a":
                                            v34 := GenS3{F0: 4294967295, F1: (^(v5)), F2: 44}
                                            _ = v34
                                        case "\n\t":
                                            switch ("a" + "") {
                                            case "a":
                                                var v35 uint8 = (^(v5))
                                                _ = v35
                                            }
                                        case "":
                                            v36 := GenS4{F0: ("a" + ("a" + "hello")), F1: false, F2: 9223372036854775807}
                                            _ = v36
                                        default:
                                            v5 = v5
                                        }
                                    }
                                    v4 = 0
                                    _ = fi33
                                }
                                fmt.Println(fi32, v4)
                                v37 := GenS4{F0: "foo bar", F1: true, F2: (v4 & v4)}
                                _ = fi32
                                _ = v37
                            }
                        }
                    }
                }
                switch 935 {
                case -896:
                    for fi38 := 0; fi38 < 16; fi38++ {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        if true {
                            if true {
                                if (!true) {
                                    var v39 uint8 = (v5 + (v5 + (v5 + (v5 + v5))))
                                    for fi40 := 0; fi40 < 19; fi40++ {
                                        v4 += (-(v4))
                                        _ = fi40
                                    }
                                    _ = v39
                                }
                                var v41 string = ("\n\t" + ("hello" + ("\n\t" + ("foo bar" + "a"))))
                                _ = v41
                            } else {
                                var v42 string = ("a" + (((("foo bar" + ("" + "hello")) + ("foo bar" + ("foo bar" + "hello"))) + "hello") + "hello"))
                                var v43 int64 = (v4 / ((42) | 1))
                                if true {
                                    v5 -= 161
                                    fmt.Println(v43, v5, v5)
                                } else {
                                    if false {
                                        var v44 float32 = 0.0
                                        v45 := GenS5{F0: 2147483647, F1: 2147483647, F2: 255}
                                        if (-2.25 < (((float64(v44) * (((0.0 / (((3.14159 / (float64(-2.25) - 0.0)) * 1.5) * (float64(v44) - 0.1))) + 0.0) - float64(v44))) * (0.0 * 3.14159)) * 0.1)) {
                                            switch (fi38 - (fi38 * fi38)) {
                                            case 127:
                                                _ = len(v6)
                                            case 32768:
                                                if false {
                                                    var v46 uint8 = v5
                                                    delete(v2, (fi38 >> ((922) & 63)))
                                                    _ = v46
                                                } else {
                                                    var v47 uint32 = 2147483647
                                                    for fi48 := 0; fi48 < 10; fi48++ {
                                                        fmt.Println(v47)
                                                        if true {
                                                            v44 = float32(1.5)
                                                            switch v42 {
                                                            case "":
                                                                v44 -= v44
                                                            default:
                                                                if ((v5 - v5) < v5) {
                                                                    var v49 uint32 = (v47 & v47)
                                                                    v50 := GenS5{F0: 255, F1: v49, F2: 2}
                                                                    _ = v49
                                                                    _ = v50
                                                                } else {
                                                                    v51 := GenS6{F0: 42, F1: (fi38 - fi38), F2: 42}
                                                                    var v52 float64 = float64(float32(0.1))
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                            }
                                                        } else {
                                                            delete(v2, fi48)
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            v6 = append(v6, (-(fi38)))
                                                            switch fi38 {
                                                            case -9223372036854775808:
                                                                var v53 float32 = v44
                                                                _ = v53
                                                            }
                                                            if len(v6) > 0 {
                                                                v6[0] = (fi38 - 128)
                                                            }
                                                        }
                                                        _ = fi48
                                                    }
                                                    _ = v47
                                                }
                                            default:
                                                _ = len(v2)
                                            }
                                        } else {
                                            v54 := v2[(fi38 ^ fi38)]
                                            _ = v54
                                        }
                                        _ = v44
                                        _ = v45
                                    }
                                }
                                _ = v42
                                _ = v43
                            }
                        } else {
                            for fi55 := 0; fi55 < 18; fi55++ {
                                v4 = (v4 ^ -567)
                                var v56 int16 = 32767
                                _ = fi55
                                _ = v56
                            }
                            switch ("foo bar" + "") {
                            case "a":
                                for fi57 := 0; fi57 < 17; fi57++ {
                                    v4 = (v4 | 995)
                                    if (1 == 100) {
                                        var v58 uint16 = 255
                                        v4 = (-(v4))
                                        _ = v58
                                    } else {
                                        v59 := GenS7{F0: 0}
                                        v60 := GenS2{F0: fi57, F1: 1, F2: 0}
                                        var v61 uint16 = uint16(v5)
                                        _ = v59
                                        _ = v60
                                        _ = v61
                                    }
                                    var v62 int16 = -499
                                    _ = len(v2)
                                    _ = fi57
                                    _ = v62
                                }
                            case "\n\t":
                                fmt.Println(v5, v5, v4)
                            default:
                                v5 = 0
                            }
                        }
                        _ = fi38
                    }
                case 128:
                    var v63 int32 = -331
                    _ = v63
                default:
                    var v64 float64 = ((((-2.25 + float64(-2.25)) * -2.25) + 3.14159) * 0.1)
                    _ = v64
                }
            }
        default:
            var v65 int = -2147483648
            _ = v65
        }
    }
    for fi66 := 0; fi66 < 9; fi66++ {
        for fi67 := 0; fi67 < 11; fi67++ {
            var v68 int32 = 1
            var v69 bool = (true && false)
            _ = fi67
            _ = v68
            _ = v69
        }
        _ = fi66
    }
    if ((((329 < 32767) && (!(!true))) || (2147483647 < -1)) && false) {
        _ = len(v3)
        fmt.Println(v5, v4)
        if true {
            if (215 > v4) {
                fmt.Println(v5, v4)
                var v70 int = 2147483647
                _ = v70
            }
            var v71 int32 = -173
            _ = v71
        } else {
            v72 := GenS5{F0: 989, F1: 1, F2: 255}
            if (!(-32768 < 1)) {
                switch 0 {
                case -9223372036854775808:
                    var v73 bool = (!((!(2147483647 <= 2147483647)) && (false || (!true))))
                    _ = v73
                case 357:
                    if false {
                        var v74 uint8 = v5
                        if false {
                            switch 776 {
                            case 406:
                                for fi75 := 0; fi75 < 20; fi75++ {
                                    v76 := GenS8{F0: 0, F1: (0 > 65535), F2: 4294967295}
                                    _ = fi75
                                    _ = v76
                                }
                            case 0:
                                delete(v2, 602)
                            case 4294967295:
                                v74 = (^(v5))
                            }
                        }
                        _ = v74
                    } else {
                        if (9223372036854775807 == 0) {
                            var v77 int8 = -73
                            _ = v77
                        } else {
                            var v78 uint16 = 255
                            delete(v3, 1)
                            fmt.Println(v4, v5, v78)
                            _ = v78
                        }
                        v79 := GenS9{F0: v4}
                        _ = v79
                    }
                case -2147483648:
                    fmt.Println(v4)
                default:
                    fmt.Println(v5)
                }
                switch 9223372036854775807 {
                case 0:
                    if true {
                        fmt.Println(v5)
                        v6 = append(v6, -297)
                    } else {
                        v80 := GenS2{F0: -2147483648, F1: 361, F2: 416}
                        switch "" {
                        case "foo bar":
                            v81 := GenS10{F0: 958, F1: ("" + "")}
                            _ = v81
                        case "hello":
                            for fi82 := 0; fi82 < 5; fi82++ {
                                var v83 int32 = 2147483647
                                v84 := GenS1{F0: 0, F1: "hello"}
                                for fi85 := 0; fi85 < 14; fi85++ {
                                    v4 = (v4 << ((1) & 63))
                                    _ = fi85
                                }
                                _ = fi82
                                _ = v83
                                _ = v84
                            }
                        case "a":
                            v86 := GenS11{F0: (^(v5)), F1: 393, F2: 2147483647}
                            _ = v86
                        }
                        v4 = v4
                        _ = v80
                    }
                case 528:
                    v4 += (v4 / ((v4) | 1))
                case 42:
                    if true {
                        for fi87 := 0; fi87 < 1; fi87++ {
                            for fi88 := 0; fi88 < 15; fi88++ {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                v89 := GenS4{F0: "a", F1: false, F2: -948}
                                _ = fi88
                                _ = v89
                            }
                            if false {
                                v5 = v5
                                var v90 float64 = float64(float32(0.1))
                                var v91 uint64 = 2
                                _ = v90
                                _ = v91
                            }
                            v4 = v4
                            _ = fi87
                        }
                        fmt.Println(v5, v5)
                        for fi92 := 0; fi92 < 12; fi92++ {
                            var v93 uint64 = 4294967295
                            switch int(v93) {
                            case -465:
                                var v94 uint8 = 55
                                _ = v94
                            case -2147483648:
                                v93 = v93
                            default:
                                v5 = (^(v5))
                            }
                            var v95 float32 = -2.25
                            _ = fi92
                            _ = v93
                            _ = v95
                        }
                    } else {
                        var v96 uint32 = 4294967295
                        if len(v6) > 0 {
                            v6[0] = 32767
                        }
                        _ = v96
                    }
                }
            }
            _ = v72
        }
    }
    delete(v3, 2147483647)
    v4 -= 0
    v5 = 153
    var v97 int64 = (v4 << ((65535) & 63))
    if (!(1000 >= 0)) {
        if true {
            for fi98 := 0; fi98 < 19; fi98++ {
                var v99 int = 42
                for fi100 := 0; fi100 < 14; fi100++ {
                    if true {
                        v101 := v1[("\n\t" + ("" + ("a" + ("hello" + "hello"))))]
                        v102 := GenS12{F0: -314, F1: "hello", F2: (!(false && (1 < 255)))}
                        _ = v101
                        _ = v102
                    }
                    v103 := GenS13{F0: "\n\t", F1: 2147483648}
                    v97 += v97
                    _ = fi100
                    _ = v103
                }
                v99 += (-(fi98))
                _ = fi98
                _ = v99
            }
            if (0 < 42) {
                v97 = (v97 ^ 0)
                var v104 int8 = -107
                var v105 uint16 = 65535
                _ = v104
                _ = v105
            } else {
                v1[("\n\t" + (("foo bar" + "hello") + "hello"))] = "\n\t"
                if (false || (uint16(v5) <= 497)) {
                    switch "foo bar" {
                    case "foo bar":
                        if true {
                            v4 = -416
                            switch "" {
                            case "":
                                for fi106 := 0; fi106 < 7; fi106++ {
                                    v107 := v3[2]
                                    if false {
                                        v1["foo bar"] = "foo bar"
                                        var v108 int16 = 42
                                        _ = len(v3)
                                        _ = v108
                                    } else {
                                        _ = len(v3)
                                    }
                                    _ = fi106
                                    _ = v107
                                }
                            }
                            v109 := GenS7{F0: 32767}
                            _ = v109
                        }
                    case "":
                        switch "" {
                        case "\n\t":
                            v110 := GenS14{F0: 0, F1: 4294967295}
                            _ = v110
                        case "foo bar":
                            if len(v6) > 0 {
                                v6[0] = 128
                            }
                        }
                    }
                    fmt.Println(v97, v4, v97)
                    var v111 int16 = -1
                    _ = v111
                }
            }
            var v112 int16 = -1
            _ = v112
        } else {
            v113 := GenS15{F0: false, F1: 767}
            _ = v113
        }
        for fi114 := 0; fi114 < 17; fi114++ {
            var v115 uint8 = (v5 + (v5 - v5))
            switch 4294967295 {
            case 32768:
                if (823 > 2) {
                    var v116 uint = 18446744073709551615
                    if (!false) {
                        _ = len(v6)
                        var v117 int32 = 42
                        _ = v117
                    } else {
                        if (float32((3.14159 - float64(((float32(1e10) + 3.14159) / 1.0)))) >= ((float32((0.1 * (float64(-2.25) / -2.25))) + 0.1) / float32((0.0 + float64(1.5))))) {
                            v118 := GenS16{F0: 725, F1: (4294967295 <= 1), F2: ""}
                            v115 = (v5 / ((127) | 1))
                            _ = v118
                        }
                        for fi119 := 0; fi119 < 14; fi119++ {
                            v116 += (^(v116))
                            if false {
                                fmt.Println(v5, v4)
                            }
                            if len(v6) > 0 {
                                v6[0] = -1
                            }
                            v4 -= int64(v115)
                            _ = fi119
                        }
                    }
                    _ = v116
                }
            }
            v115 = (v115 & (v5 & (v115 * (v115 - (v115 + (v115 % (((v5 >> ((18446744073709551615) & 63))) | 1)))))))
            _ = fi114
            _ = v115
        }
        v120 := GenS16{F0: 42, F1: (!(-9223372036854775808 <= (v97 | (v4 | 0)))), F2: "hello"}
        _ = v120
    }
    for fi121 := 0; fi121 < 14; fi121++ {
        switch "\n\t" {
        case "a":
            for fi122 := 0; fi122 < 11; fi122++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                switch (^(fi122)) {
                case 954:
                    var v123 int32 = 2147483647
                    _ = v123
                }
                switch ("hello" + "a") {
                case "":
                    var v124 string = "foo bar"
                    if (!((v5 << ((2) & 63)) >= (v5 & 127))) {
                        fmt.Println(v124, v5, fi121)
                        var v125 float32 = 0.0
                        _ = v125
                    }
                    v126 := GenS17{F0: true, F1: 36, F2: 255}
                    _ = v124
                    _ = v126
                default:
                    v127 := GenS18{F0: (("\n\t" + "") + ("\n\t" + (("\n\t" + "a") + ("a" + "foo bar")))), F1: true, F2: (fi121 / (((^(fi121))) | 1))}
                    _ = v127
                }
                _ = fi122
            }
        case "hello":
            var v128 float64 = 3.14159
            _ = v128
        default:
            if true {
                switch (fi121 + fi121) {
                case 1:
                    if false {
                        v97 -= (v4 - v97)
                        v6 = append(v6, (-(fi121)))
                        var v129 float32 = 3.14159
                        _ = v129
                    } else {
                        v130 := GenS6{F0: 4294967295, F1: fi121, F2: 829}
                        _ = v130
                    }
                case 811:
                    if len(v6) > 0 {
                        v6[0] = (fi121 | fi121)
                    }
                default:
                    var v131 bool = ((!true) && false)
                    _ = v131
                }
            }
        }
        if false {
            var v132 float32 = (0.0 - (1.5 + -2.25))
            v133 := GenS11{F0: 255, F1: 988, F2: -1}
            if (45 <= 127) {
                fmt.Println(fi121, v5, v5)
                v134 := GenS19{F0: 42, F1: ((("foo bar" + ((((("\n\t" + "foo bar") + "hello") + "hello") + ("foo bar" + "hello")) + "a")) + "\n\t") + "a"), F2: 0}
                _ = v134
            } else {
                switch (fi121 - fi121) {
                case -9223372036854775808:
                    v132 = ((float32(1e10) + 1.5) * v132)
                case 128:
                    switch ("" + "") {
                    case "":
                        for fi135 := 0; fi135 < 10; fi135++ {
                            if (0 >= 65535) {
                                v136 := GenS20{F0: (^(v5)), F1: 42}
                                _ = v136
                            } else {
                                v2[fi121] = fi121
                                var v137 float64 = (0.0 * -2.25)
                                delete(v2, fi135)
                                _ = v137
                            }
                            var v138 int64 = (v4 << ((127) & 63))
                            var v139 float32 = v132
                            v138 -= v138
                            _ = fi135
                            _ = v138
                            _ = v139
                        }
                    case "a":
                        for fi140 := 0; fi140 < 16; fi140++ {
                            _ = len(v6)
                            _ = fi140
                        }
                    case "\n\t":
                        var v141 int64 = -9223372036854775808
                        _ = v141
                    default:
                        var v142 float64 = 1.5
                        _ = v142
                    }
                default:
                    v4 += (v97 - 549)
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v5 = 0
            }
            _ = v132
            _ = v133
        }
        switch -433 {
        case -9223372036854775808:
            v4 = (v97 + (v4 % (((^(v4))) | 1)))
        case 987:
            if true {
                v143 := GenS21{F0: 1, F1: 32767}
                _ = v143
            } else {
                var v144 uint8 = v5
                _ = v144
            }
        default:
            switch fi121 {
            case 32768:
                v97 = v97
            case -9223372036854775808:
                v4 -= v4
            default:
                v145 := v3[1]
                _ = v145
            }
        }
        _ = fi121
    }
    var v146 uint32 = 0
    if (!true) {
        var v147 float32 = 1e10
        _ = v147
    } else {
        if ((-2.25 * -2.25) <= 0.0) {
            var v148 uint = uint(v146)
            for fi149 := 0; fi149 < 15; fi149++ {
                v150 := GenS22{F0: 0, F1: (1 <= v97)}
                _ = fi149
                _ = v150
            }
            var v151 int8 = -45
            _ = v148
            _ = v151
        }
        switch "hello" {
        case "\n\t":
            var v152 uint64 = 0
            _ = v152
        case "hello":
            for fi153 := 0; fi153 < 20; fi153++ {
                if ((v146 | (v146 / (((v146 + 4294967295)) | 1))) > v146) {
                    var v154 float32 = (-2.25 + -2.25)
                    _ = v154
                }
                v155 := GenS12{F0: 2147483647, F1: "foo bar", F2: false}
                var v156 uint16 = 555
                _ = fi153
                _ = v155
                _ = v156
            }
        case "":
            switch "a" {
            case "hello":
                for fi157 := 0; fi157 < 2; fi157++ {
                    v6 = append(v6, (fi157 & fi157))
                    v5 = (v5 ^ v5)
                    _ = fi157
                }
            default:
                for fi158 := 0; fi158 < 13; fi158++ {
                    v97 -= -595
                    _ = fi158
                }
            }
        default:
            for fi159 := 0; fi159 < 13; fi159++ {
                v146 = (v146 & (v146 >> ((uint(v5)) & 63)))
                var v160 int8 = 0
                v146 += (^(v146))
                _ = fi159
                _ = v160
            }
        }
    }
    var v161 float64 = (((0.0 / 1.0) - (1.5 - 0.1)) * (((0.0 - 0.0) * 0.1) / (0.0 + float64(1.5))))
    fmt.Println(v5, v97)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v146:", v146)
    fmt.Println("v161:", v161)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v146
    _ = v161
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v97
}
