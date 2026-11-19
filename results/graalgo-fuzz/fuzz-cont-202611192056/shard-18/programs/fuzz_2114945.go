package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int32 }
type GenS2 struct { F0 int8; F1 int16; F2 int64 }
type GenS3 struct { F0 bool; F1 bool; F2 int32 }
type GenS4 struct { F0 bool; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 string; F1 string; F2 uint }
type GenS9 struct { F0 int16; F1 string }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int8; F1 int; F2 int32 }
type GenS13 struct { F0 uint64; F1 uint8 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int; F1 int64 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 int64; F1 int8 }
func genFn0() uint64 {
    v4 := []int{42}
    if false {
        var v5 uint8 = 2
        var v6 uint16 = 858
        var v7 int = int(v5)
        _ = v5
        _ = v6
        _ = v7
    }
    v4 = append(v4, 128)
    if ((0 <= 438) && true) {
        var v8 bool = (2147483647 == 128)
        for fi9 := 0; fi9 < 11; fi9++ {
            if ((1.5 >= 1e10) || v8) {
                v10 := GenS1{F0: 923, F1: -1, F2: 0}
                _ = v10
            }
            var v11 bool = (-32768 > -391)
            _ = fi9
            _ = v11
        }
        _ = v8
    } else {
        var v12 uint16 = 709
        if true {
            var v13 int16 = -32768
            _ = v13
        } else {
            if false {
                switch ((((("" + "foo bar") + "\n\t") + "foo bar") + "") + ("hello" + "a")) {
                case "":
                    if true {
                        for fi14 := 0; fi14 < 1; fi14++ {
                            v12 = 32767
                            _ = fi14
                        }
                        for fi15 := 0; fi15 < 5; fi15++ {
                            fmt.Println(v12)
                            _ = fi15
                        }
                    } else {
                        v12 = 32767
                        for fi16 := 0; fi16 < 15; fi16++ {
                            var v17 bool = true
                            v17 = v17
                            switch "\n\t" {
                            case "foo bar":
                                var v18 uint16 = 1
                                _ = v18
                            case "a":
                                if len(v4) > 0 {
                                    v4[0] = fi16
                                }
                            default:
                                var v19 uint64 = 0
                                _ = v19
                            }
                            _ = fi16
                            _ = v17
                        }
                    }
                case "foo bar":
                    v20 := GenS1{F0: 726, F1: -18, F2: 0}
                    _ = v20
                default:
                    v4 = append(v4, -9223372036854775808)
                }
                fmt.Println(v12)
                for fi21 := 0; fi21 < 11; fi21++ {
                    for fi22 := 0; fi22 < 4; fi22++ {
                        switch -480 {
                        case 4294967295:
                            var v23 int16 = 846
                            _ = v23
                        case -9223372036854775808:
                            switch fi21 {
                            case -2147483648:
                                var v24 float32 = (1.5 / (0.1 + 1e10))
                                switch (fi22 << ((2) & 63)) {
                                case -852:
                                    v25 := GenS1{F0: 0, F1: 42, F2: -612}
                                    _ = v25
                                }
                                if true {
                                    v12 = v12
                                } else {
                                    fmt.Println(fi21, v24)
                                    var v26 string = "hello"
                                    if (44 > 42) {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    } else {
                                        if false {
                                            v27 := GenS1{F0: 42, F1: 0, F2: int32(v12)}
                                            var v28 float32 = (1e10 / v24)
                                            switch fi22 {
                                            case -746:
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case -2147483648:
                                                fmt.Println(v24, v24)
                                            default:
                                                var v29 bool = false
                                                _ = v29
                                            }
                                            var v30 uint32 = 2147483647
                                            _ = v27
                                            _ = v28
                                            _ = v30
                                        } else {
                                            switch (v26 + ("a" + v26)) {
                                            case "hello":
                                                v31 := GenS2{F0: 100, F1: -32768, F2: 124}
                                                _ = v31
                                            case "a":
                                                v32 := GenS3{F0: (42 <= 42), F1: false, F2: 853}
                                                _ = v32
                                            case "\n\t":
                                                v24 -= 3.14159
                                            }
                                        }
                                    }
                                    _ = v26
                                }
                                _ = v24
                            case 32767:
                                v33 := GenS4{F0: (-1 >= -233), F1: 127, F2: (v12 / (((v12 * v12)) | 1))}
                                _ = v33
                            case -9223372036854775808:
                                if (true || false) {
                                    v12 = v12
                                }
                            default:
                                _ = len(v4)
                            }
                        }
                        if false {
                            var v34 bool = false
                            if false {
                                var v35 uint8 = 1
                                _ = len(v4)
                                v36 := GenS5{F0: (v12 % (((v12 - (v12 & v12))) | 1)), F1: v35}
                                v37 := GenS6{F0: (v12 >> ((830) & 63)), F1: 84}
                                _ = v35
                                _ = v36
                                _ = v37
                            }
                            var v38 uint = 18446744073709551615
                            v34 = true
                            switch ("" + "foo bar") {
                            case "foo bar":
                                var v39 int64 = -1
                                _ = v39
                            case "a":
                                if (0 == (v38 << (((v38 | 65535)) & 63))) {
                                    v12 += v12
                                    _ = len(v4)
                                } else {
                                    if (((3.14159 + float32(-2.25)) >= 1.5) && (3.14159 > 1.5)) {
                                        switch fi21 {
                                        case -34:
                                            var v40 bool = v34
                                            _ = v40
                                        }
                                        var v41 uint16 = v12
                                        v12 -= v41
                                        _ = v41
                                    }
                                    if false {
                                        v42 := GenS2{F0: -128, F1: 1000, F2: -1}
                                        if true {
                                            v43 := GenS6{F0: 0, F1: -247}
                                            v12 = 2
                                            _ = v43
                                        }
                                        _ = len(v4)
                                        _ = v42
                                    }
                                    v12 = (v12 * (v12 / (((v12 & (v12 * (v12 | v12)))) | 1)))
                                }
                            }
                            _ = v34
                            _ = v38
                        }
                        if false {
                            if true {
                                if (((255 <= 2) && true) && false) {
                                    v44 := GenS7{F0: -1, F1: (("hello" + "") + "")}
                                    _ = v44
                                }
                                v45 := GenS6{F0: (v12 | (v12 / ((390) | 1))), F1: -2147483648}
                                v12 = (v12 & 382)
                                _ = v45
                            }
                        }
                        _ = fi22
                    }
                    var v46 int16 = 1
                    v12 += 1
                    _ = fi21
                    _ = v46
                }
            }
            switch "\n\t" {
            case "":
                var v47 int16 = 42
                _ = v47
            case "\n\t":
                v48 := GenS5{F0: 0, F1: 1}
                _ = v48
            case "foo bar":
                if (float64(1.5) < 3.14159) {
                    fmt.Println(v12)
                }
            }
        }
        v12 -= v12
        _ = v12
    }
    _ = v4
    return 92
}
func genFn1(p0 bool, p1 int64, p2 int32) bool {
    if (169 != 172) {
        if (("foo bar" != ("a" + "\n\t")) || (2 == genFn0())) {
            var v49 float32 = 3.14159
            v50 := []int{9223372036854775807, 1}
            var v51 uint32 = 2
            _ = v49
            _ = v50
            _ = v51
        } else {
            v52 := GenS1{F0: genFn0(), F1: 77, F2: 42}
            _ = v52
        }
        var v53 bool = ((3.14159 > 0.1) || p0)
        _ = v53
    }
    var v54 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v54
    return p0
}

func main() {
    var v1 uint8 = 200
    v2 := []uint32{793, 154}
    var v3 string = "a"
    var v55 int64 = -1
    fmt.Println(v3)
    var v56 uint = 414
    v1 -= v1
    var v57 uint32 = 2
    var v58 string = ("hello" + "hello")
    switch "hello" {
    case "":
        var v59 uint8 = (^(v1))
        _ = v59
    case "\n\t":
        v57 = v57
    }
    fmt.Println(v57, v3)
    var v60 float64 = 1.5
    v57 = (v57 - (v57 << (((v56 | (v56 / (((v56 % ((65535) | 1))) | 1)))) & 63)))
    if false {
        fmt.Println(v57, v1)
        v1 = v1
    } else {
        if ((42 <= 762) && (-9223372036854775808 <= int(v1))) {
            var v61 string = "foo bar"
            for fi62 := 0; fi62 < 19; fi62++ {
                v58 = v3
                for fi63 := 0; fi63 < 8; fi63++ {
                    var v64 int32 = 949
                    _ = fi63
                    _ = v64
                }
                if true {
                    switch (v3 + "a") {
                    case "a":
                        var v65 int16 = 42
                        _ = v65
                    case "foo bar":
                        if true {
                            for fi66 := 0; fi66 < 8; fi66++ {
                                var v67 uint = (v56 % (((v56 * (v56 & v56))) | 1))
                                _ = fi66
                                _ = v67
                            }
                            if (true || true) {
                                v1 = 1
                                fmt.Println(v1)
                            } else {
                                var v68 int8 = -1
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v69 uint32 = v57
                                _ = v68
                                _ = v69
                            }
                            fmt.Println(v56, v56)
                        }
                    default:
                        for fi70 := 0; fi70 < 7; fi70++ {
                            v57 += (v57 * 234)
                            _ = fi70
                        }
                    }
                }
                _ = fi62
            }
            for fi71 := 0; fi71 < 13; fi71++ {
                for fi72 := 0; fi72 < 17; fi72++ {
                    if true {
                        v3 = v61
                        v73 := GenS3{F0: false, F1: false, F2: 1}
                        if (v57 > v57) {
                            switch "hello" {
                            case "a":
                                _ = len(v2)
                            case "":
                                v74 := GenS5{F0: 255, F1: v1}
                                var v75 int32 = -484
                                _ = v74
                                _ = v75
                            case "foo bar":
                                v76 := GenS8{F0: v61, F1: v3, F2: v56}
                                _ = v76
                            }
                            if (42 != 988) {
                                fmt.Println(v1)
                                var v77 int8 = 127
                                var v78 bool = ((v58 + ("foo bar" + "a")) >= (v3 + ((v61 + "foo bar") + "\n\t")))
                                _ = v77
                                _ = v78
                            } else {
                                var v79 uint8 = 0
                                _ = v79
                            }
                            v61 = v3
                        }
                        _ = v73
                    } else {
                        v80 := GenS6{F0: 698, F1: 42}
                        _ = v80
                    }
                    switch -113 {
                    case 42:
                        if (-448 == -552) {
                            fmt.Println(v58, v56, v60)
                            v81 := GenS9{F0: 1000, F1: "\n\t"}
                            v82 := GenS5{F0: 1, F1: v1}
                            _ = v81
                            _ = v82
                        } else {
                            v60 = float64(0.1)
                            v2 = append(v2, 4294967295)
                            v55 = -736
                        }
                    case 32768:
                        v57 -= (v57 / ((2147483648) | 1))
                    case -9223372036854775808:
                        if false {
                            v83 := GenS8{F0: v58, F1: v3, F2: 531}
                            _ = v83
                        }
                    default:
                        switch 1 {
                        case -873:
                            if ((v57 <= v57) || (uint8(v57) > (v1 & v1))) {
                                var v84 uint16 = 65535
                                _ = v84
                            } else {
                                fmt.Println(fi71)
                                v85 := GenS2{F0: 46, F1: 1000, F2: (v55 / ((v55) | 1))}
                                v86 := GenS10{F0: v1}
                                _ = v85
                                _ = v86
                            }
                        case 32767:
                            if true {
                                v87 := GenS4{F0: true, F1: 42, F2: 1}
                                v1 -= (v1 / ((v1) | 1))
                                switch (fi72 & 4294967295) {
                                case 128:
                                    var v88 int32 = -160
                                    _ = v88
                                case 590:
                                    v61 = v61
                                case 4294967295:
                                    v57 -= 1
                                default:
                                    switch fi72 {
                                    case 74:
                                        var v89 float32 = 1e10
                                        if false {
                                            v90 := GenS4{F0: true, F1: 0, F2: 2}
                                            var v91 float64 = 1e10
                                            if true {
                                                v92 := GenS11{F0: 946}
                                                var v93 uint = (v56 & 1)
                                                v2 = append(v2, v57)
                                                _ = v92
                                                _ = v93
                                            }
                                            _ = v90
                                            _ = v91
                                        }
                                        _ = v89
                                    case -2147483648:
                                        if true {
                                            v94 := GenS12{F0: -128, F1: (fi71 | (-(fi72))), F2: -905}
                                            _ = v94
                                        }
                                    }
                                }
                                _ = v87
                            } else {
                                if false {
                                    if ((!true) || (false || false)) {
                                        var v95 uint16 = 1
                                        _ = len(v2)
                                        switch (fi71 ^ fi71) {
                                        case -2147483648:
                                            switch fi72 {
                                            case -1:
                                                if ((false || ((^(fi72)) < 32767)) && (!false)) {
                                                    if false {
                                                        _ = len(v2)
                                                        v2 = append(v2, v57)
                                                        var v96 uint64 = 18446744073709551615
                                                        _ = v96
                                                    }
                                                } else {
                                                    v1 -= 11
                                                    v60 += -2.25
                                                    v60 = v60
                                                }
                                                v2 = append(v2, v57)
                                                if false {
                                                    v2 = append(v2, 4294967295)
                                                }
                                                var v97 string = "a"
                                                _ = v97
                                            case 9223372036854775807:
                                                var v98 int64 = (v55 << ((v56) & 63))
                                                _ = v98
                                            case 128:
                                                fmt.Println(v60, v95, fi71)
                                                var v99 int = fi72
                                                var v100 int8 = -128
                                                _ = v100
                                                _ = v99
                                            }
                                        case 128:
                                            if false {
                                                v101 := GenS7{F0: (v55 * v55), F1: "\n\t"}
                                                _ = v101
                                            }
                                            if false {
                                                var v102 int = fi72
                                                v103 := GenS11{F0: int(v56)}
                                                if false {
                                                    v61 = (v61 + "a")
                                                } else {
                                                    v61 = v3
                                                    var v104 float64 = (3.14159 / v60)
                                                    _ = v104
                                                }
                                                v105 := GenS13{F0: uint64(v57), F1: (v1 ^ v1)}
                                                _ = v102
                                                _ = v103
                                                _ = v105
                                            } else {
                                                if len(v2) > 0 {
                                                    v2[0] = (^(v57))
                                                }
                                                switch (fi72 & (fi72 * (fi72 / ((fi71) | 1)))) {
                                                case 4294967295:
                                                    var v106 uint32 = v57
                                                    _ = v106
                                                default:
                                                    if (!(!true)) {
                                                        fmt.Println(v55, v95, v61)
                                                        v1 = (^(v1))
                                                        var v107 float64 = 3.14159
                                                        _ = v107
                                                    }
                                                }
                                                if ((v1 > (v1 >> (((^(v56))) & 63))) || (false || false)) {
                                                    v108 := GenS14{F0: -462}
                                                    v57 -= (v57 / ((v57) | 1))
                                                    _ = v108
                                                } else {
                                                    switch fi71 {
                                                    case 0:
                                                        fmt.Println(v56, v60, v95)
                                                    case -1:
                                                        v61 = v61
                                                    default:
                                                        if (false || (-411 <= fi71)) {
                                                            var v109 int16 = 42
                                                            fmt.Println(v60, v58)
                                                            v55 += v55
                                                            _ = v109
                                                        }
                                                    }
                                                    var v110 uint = 2
                                                    if len(v2) > 0 {
                                                        v2[0] = v57
                                                    }
                                                    _ = v110
                                                }
                                            }
                                            if (false && (true && true)) {
                                                var v111 uint16 = 32767
                                                fmt.Println(v58)
                                                v61 = (v61 + v58)
                                                _ = v111
                                            } else {
                                                if true {
                                                    var v112 uint16 = v95
                                                    _ = v112
                                                } else {
                                                    fmt.Println(v3, v3, v56)
                                                    if false {
                                                        switch 128 {
                                                        case -9223372036854775808:
                                                            if false {
                                                                v113 := GenS15{F0: (fi72 - -9223372036854775808), F1: 0}
                                                                v56 = (v56 - (v56 - v56))
                                                                _ = v113
                                                            } else {
                                                                fmt.Println(fi71)
                                                                var v114 float64 = (v60 / 1.0)
                                                                fmt.Println(v95)
                                                                _ = v114
                                                            }
                                                        default:
                                                            v56 = v56
                                                        }
                                                    } else {
                                                        var v115 uint64 = 2
                                                        fmt.Println(v3)
                                                        _ = v115
                                                    }
                                                }
                                                fmt.Println(v3)
                                                v1 = 1
                                                fmt.Println(v56, v61)
                                                v116 := GenS1{F0: uint64(v57), F1: -1, F2: 2147483647}
                                                _ = v116
                                            }
                                        }
                                        _ = v95
                                    } else {
                                        if true {
                                            v60 = (v60 - (v60 / v60))
                                        }
                                        v1 = (^(v1))
                                    }
                                    var v117 uint = 1
                                    _ = v117
                                }
                                if (false || false) {
                                    var v118 int = fi72
                                    if len(v2) > 0 {
                                        v2[0] = v57
                                    }
                                    _ = v118
                                } else {
                                    var v119 bool = (false && false)
                                    if ((v55 - (-(v55))) >= (v55 % (((-(v55))) | 1))) {
                                        var v120 float32 = float32((0.0 * v60))
                                        var v121 int8 = 1
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        fmt.Println(v56, v57)
                                        _ = v120
                                        _ = v121
                                    }
                                    if (!v119) {
                                        if true {
                                            var v122 int64 = (v55 / ((-1) | 1))
                                            if len(v2) > 0 {
                                                v2[0] = v57
                                            }
                                            _ = v122
                                        } else {
                                            var v123 int64 = (v55 - 78)
                                            var v124 int32 = 42
                                            v125 := GenS16{F0: (v124 | (v124 >> ((v56) & 63)))}
                                            _ = v123
                                            _ = v124
                                            _ = v125
                                        }
                                        fmt.Println(fi72, v1, v55)
                                    } else {
                                        var v126 uint8 = (^(v1))
                                        switch v61 {
                                        case "hello":
                                            if false {
                                                fmt.Println(v55)
                                                fmt.Println(v1, fi72, v58)
                                                var v127 uint16 = 541
                                                _ = v127
                                            }
                                        case "":
                                            if (!(!(0.0 < ((0.0 + (v60 - float64(1.5))) / v60)))) {
                                                var v128 int16 = 1
                                                switch 946 {
                                                case -2147483648:
                                                    var v129 bool = true
                                                    v130 := GenS5{F0: 255, F1: (v126 / ((v1) | 1))}
                                                    switch (-(fi71)) {
                                                    case 32767:
                                                        fmt.Println(v61, v56)
                                                    }
                                                    _ = v129
                                                    _ = v130
                                                case 4294967295:
                                                    v60 += float64(-2.25)
                                                default:
                                                    if len(v2) > 0 {
                                                        v2[0] = (v57 - (v57 << ((v56) & 63)))
                                                    }
                                                }
                                                if true {
                                                    var v131 uint32 = 2
                                                    fmt.Println(v57, v3)
                                                    v57 += (^(v131))
                                                    _ = v131
                                                }
                                                _ = v128
                                            }
                                        }
                                        _ = v126
                                    }
                                    _ = v119
                                }
                                fmt.Println(v1, v57, v56)
                            }
                        }
                        if (v61 == v58) {
                            var v132 string = "\n\t"
                            var v133 uint = (^(v56))
                            _ = v132
                            _ = v133
                        }
                        v134 := GenS15{F0: 211, F1: v55}
                        var v135 int8 = 127
                        var v136 int64 = 0
                        var v137 uint32 = v57
                        v1 += (v1 - 83)
                        _ = v134
                        _ = v135
                        _ = v136
                        _ = v137
                    }
                    var v138 int64 = v55
                    _ = fi72
                    _ = v138
                }
                if false {
                    _ = len(v2)
                }
                _ = fi71
            }
            _ = v61
        } else {
            v57 = v57
            v139 := GenS13{F0: 575, F1: v1}
            for fi140 := 0; fi140 < 16; fi140++ {
                var v141 bool = (!false)
                if v141 {
                    switch v3 {
                    case "a":
                        v2 = append(v2, 4294967295)
                    }
                    fmt.Println(v57, v60, v1)
                }
                v142 := GenS14{F0: 262}
                _ = fi140
                _ = v141
                _ = v142
            }
            _ = v139
        }
    }
    switch -510 {
    case 32768:
        var v143 int32 = int32(v57)
        _ = v143
    case 4294967295:
        for fi144 := 0; fi144 < 1; fi144++ {
            var v145 int64 = (v55 % ((9223372036854775807) | 1))
            for fi146 := 0; fi146 < 15; fi146++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if (100 > 104) {
                    switch (fi146 % (((fi146 << (((v56 ^ 18446744073709551615)) & 63))) | 1)) {
                    case 42:
                        v55 += (v145 / ((51) | 1))
                    case -800:
                        v1 = v1
                    case -102:
                        var v147 float64 = float64((1.5 / -2.25))
                        v148 := GenS7{F0: v55, F1: v3}
                        _ = v147
                        _ = v148
                    }
                    _ = len(v2)
                } else {
                    _ = len(v2)
                    v58 = "a"
                }
                var v149 float32 = 1.5
                v56 -= 4294967295
                _ = fi146
                _ = v149
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi144
            _ = v145
        }
    default:
        v150 := GenS2{F0: 19, F1: -85, F2: (v55 >> ((uint(v1)) & 63))}
        _ = v150
    }
    v151 := GenS2{F0: 42, F1: 42, F2: 0}
    var v152 float64 = v60
    switch v58 {
    case "hello":
        v153 := GenS3{F0: ((v60 * ((3.14159 * float64(3.14159)) - (v60 - 3.14159))) >= -2.25), F1: false, F2: 42}
        _ = v153
    case "a":
        for fi154 := 0; fi154 < 5; fi154++ {
            for fi155 := 0; fi155 < 7; fi155++ {
                var v156 bool = false
                if v156 {
                    v1 += (v1 / ((42) | 1))
                } else {
                    if false {
                        if len(v2) > 0 {
                            v2[0] = (v57 / ((2147483648) | 1))
                        }
                    }
                    fmt.Println(fi154, v60)
                    if ((fi154 >> ((v56) & 63)) < fi155) {
                        var v157 int64 = v55
                        if true {
                            v158 := GenS7{F0: 9223372036854775807, F1: v58}
                            _ = v158
                        } else {
                            var v159 int8 = 100
                            if v156 {
                                _ = len(v2)
                                if true {
                                    if v156 {
                                        var v160 float32 = 3.14159
                                        if len(v2) > 0 {
                                            v2[0] = v57
                                        }
                                        var v161 int32 = -702
                                        _ = v160
                                        _ = v161
                                    } else {
                                        v58 = ("foo bar" + "hello")
                                        if (!v156) {
                                            switch "a" {
                                            case "foo bar":
                                                v58 = "foo bar"
                                            case "hello":
                                                if (v1 != (v1 + (v1 % (((v1 / (((v1 - (v1 * (v1 & (v1 ^ (v1 % (((v1 - v1)) | 1))))))) | 1))) | 1)))) {
                                                    v162 := GenS10{F0: uint8(v56)}
                                                    v2 = append(v2, 4294967295)
                                                    v163 := GenS10{F0: v1}
                                                    _ = v162
                                                    _ = v163
                                                }
                                            default:
                                                var v164 uint8 = (v1 - v1)
                                                _ = v164
                                            }
                                            var v165 float64 = (1e10 + v152)
                                            v166 := GenS3{F0: (0 != (v1 + v1)), F1: v156, F2: 2147483647}
                                            _ = v165
                                            _ = v166
                                        }
                                        v167 := GenS2{F0: (v159 ^ 127), F1: 0, F2: (v157 * 9223372036854775807)}
                                        _ = v167
                                    }
                                    var v168 bool = (!v156)
                                    var v169 int = (fi155 - fi154)
                                    _ = v168
                                    _ = v169
                                } else {
                                    v170 := GenS17{F0: v157, F1: (v159 * (^(v159)))}
                                    _ = len(v2)
                                    fmt.Println(v3)
                                    var v171 uint16 = 804
                                    fmt.Println(v57, v152, v156)
                                    _ = v170
                                    _ = v171
                                }
                            } else {
                                v1 = (v1 % ((v1) | 1))
                                v172 := GenS15{F0: fi155, F1: v157}
                                if len(v2) > 0 {
                                    v2[0] = v57
                                }
                                _ = v172
                            }
                            v173 := GenS3{F0: (false && (v58 >= ((v58 + "a") + "a"))), F1: (2 >= 42), F2: 1}
                            _ = v159
                            _ = v173
                        }
                        _ = v157
                    }
                    var v174 int8 = 100
                    v156 = (((v55 * -9223372036854775808) > 0) || (-2147483648 == 1))
                    _ = v174
                }
                v175 := GenS5{F0: 1, F1: (v1 << ((v56) & 63))}
                var v176 float64 = (v152 - (v60 * -2.25))
                _ = fi155
                _ = v156
                _ = v175
                _ = v176
            }
            _ = fi154
        }
    default:
        fmt.Println(v1, v55)
    }
    if true {
        if true {
            if true {
                var v177 int16 = -32768
                fmt.Println(v60, v3)
                v57 = 2147483647
                _ = v177
            }
            var v178 uint8 = 255
            _ = v178
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v152:", v152)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v151
    _ = v152
    _ = v2
    _ = v3
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v60
}
