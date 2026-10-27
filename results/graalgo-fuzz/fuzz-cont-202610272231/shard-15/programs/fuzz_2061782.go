package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int32; F1 uint8; F2 string }
type GenS4 struct { F0 bool; F1 int8; F2 int }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 int16; F1 uint64; F2 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint; F1 uint8; F2 string }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 string; F1 int }
type GenS13 struct { F0 string; F1 uint32 }
type GenS14 struct { F0 string; F1 uint64; F2 string }
type GenS15 struct { F0 uint8; F1 int16; F2 int32 }
type GenS16 struct { F0 string; F1 uint32 }
type GenS17 struct { F0 uint16; F1 int32 }
type GenS18 struct { F0 int32; F1 bool; F2 int16 }

func main() {
    var v1 uint32 = 217
    var v2 int64 = -696
    v3 := make(map[int]int)
    var v4 uint8 = 149
    var v5 uint64 = 0
    fmt.Println(v2, v2)
    v6 := GenS1{F0: v5, F1: 4294967295}
    if false {
        v7 := GenS2{F0: 255, F1: -1}
        for fi8 := 0; fi8 < 2; fi8++ {
            for fi9 := 0; fi9 < 9; fi9++ {
                fmt.Println(fi8)
                v10 := GenS1{F0: 18446744073709551615, F1: 612}
                var v11 uint16 = 42
                _ = fi9
                _ = v10
                _ = v11
            }
            v4 -= (v4 << ((4294967295) & 63))
            if false {
                for fi12 := 0; fi12 < 20; fi12++ {
                    v13 := GenS3{F0: -1, F1: 127, F2: ((("foo bar" + ("foo bar" + ("a" + "\n\t"))) + ("hello" + "\n\t")) + "\n\t")}
                    _ = fi12
                    _ = v13
                }
                v14 := []int{-714, (fi8 / ((fi8) | 1))}
                var v15 string = ("foo bar" + "\n\t")
                _ = v14
                _ = v15
            } else {
                for fi16 := 0; fi16 < 5; fi16++ {
                    if (!false) {
                        v17 := GenS4{F0: true, F1: 0, F2: (fi16 >> ((4294967295) & 63))}
                        v2 -= (v2 >> ((237) & 63))
                        v5 = 18446744073709551615
                        _ = v17
                    }
                    _ = fi16
                }
                var v18 float64 = (3.14159 * 0.1)
                var v19 int32 = 422
                _ = v18
                _ = v19
            }
            _ = fi8
        }
        _ = v7
    } else {
        var v20 uint8 = uint8(v5)
        if false {
            switch 127 {
            case 389:
                var v21 uint = 684
                _ = v21
            case -656:
                for fi22 := 0; fi22 < 19; fi22++ {
                    switch fi22 {
                    case 1:
                        v23 := []int{2147483647, (fi22 << ((922) & 63)), (fi22 % ((int(v1)) | 1)), (fi22 / (((-(fi22))) | 1))}
                        _ = v23
                    case -1:
                        v24 := GenS3{F0: -619, F1: (^(v20)), F2: "\n\t"}
                        _ = v24
                    case 42:
                        v20 = 2
                    default:
                        fmt.Println(fi22, v1)
                    }
                    switch 113 {
                    case 35:
                        v25 := []int{(fi22 | fi22), -148, (fi22 + fi22)}
                        _ = v25
                    }
                    _ = fi22
                }
            case 1:
                for fi26 := 0; fi26 < 18; fi26++ {
                    v20 += uint8(v1)
                    var v27 string = ("" + "\n\t")
                    _ = len(v3)
                    _ = fi26
                    _ = v27
                }
            }
        } else {
            switch ((("hello" + "foo bar") + (("\n\t" + "foo bar") + "foo bar")) + ("" + "\n\t")) {
            case "a":
                fmt.Println(v4)
            }
            if true {
                v2 = (v2 & -787)
            } else {
                for fi28 := 0; fi28 < 2; fi28++ {
                    v20 -= v20
                    _ = fi28
                }
                for fi29 := 0; fi29 < 4; fi29++ {
                    for fi30 := 0; fi30 < 11; fi30++ {
                        v31 := GenS4{F0: ("foo bar" != "\n\t"), F1: 32, F2: (fi29 + 0)}
                        switch fi30 {
                        case 9223372036854775807:
                            v1 -= 1
                        case -317:
                            switch ("foo bar" + "") {
                            case "foo bar":
                                v32 := v3[1]
                                _ = v32
                            case "\n\t":
                                switch "hello" {
                                case "":
                                    if false {
                                        v1 += 763
                                        var v33 uint8 = v20
                                        v34 := []int{(-(fi30)), 528, fi29, (-(fi30))}
                                        _ = v33
                                        _ = v34
                                    }
                                case "hello":
                                    var v35 string = ("foo bar" + ("\n\t" + "hello"))
                                    _ = v35
                                case "a":
                                    var v36 int8 = 127
                                    var v37 float64 = 0.0
                                    _ = v36
                                    _ = v37
                                default:
                                    v38 := []int{32768}
                                    _ = v38
                                }
                            }
                        case -836:
                            switch (^(fi29)) {
                            case 343:
                                v4 += (v20 | 0)
                            case -687:
                                if (0 < -1) {
                                    if false {
                                        if (32767 != fi30) {
                                            if true {
                                                v39 := GenS4{F0: (int16(v1) != 115), F1: 127, F2: fi29}
                                                var v40 uint8 = v20
                                                if (!true) {
                                                    v41 := GenS5{F0: "\n\t", F1: (fi30 << ((33) & 63))}
                                                    v42 := GenS1{F0: (v5 + (v5 ^ (v5 * v5))), F1: 18446744073709551615}
                                                    var v43 int = (^(fi29))
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                }
                                                _ = v39
                                                _ = v40
                                            }
                                        } else {
                                            v44 := GenS6{F0: 1, F1: v5, F2: 2}
                                            var v45 int8 = 100
                                            var v46 uint64 = (v5 & 1)
                                            if ((v5 - v5) != v46) {
                                                if false {
                                                    v47 := GenS7{F0: (65535 >= 246)}
                                                    var v48 float32 = 0.0
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    if (-750 != -1) {
                                                        fmt.Println(v1)
                                                        if ((v45 ^ (v45 % ((v45) | 1))) >= v45) {
                                                            if true {
                                                                var v49 int = fi30
                                                                _ = v49
                                                            } else {
                                                                var v50 int64 = 42
                                                                _ = v50
                                                            }
                                                            v5 -= (v46 % (((v5 % ((v5) | 1))) | 1))
                                                            fmt.Println(v4, v46)
                                                            switch int(v2) {
                                                            case 2147483647:
                                                                v51 := GenS1{F0: v46, F1: 0}
                                                                _ = v51
                                                            default:
                                                                if ((v1 ^ 65535) < v1) {
                                                                    v1 = v1
                                                                } else {
                                                                    v52 := GenS8{F0: 1000}
                                                                    var v53 float64 = ((0.1 + 0.0) / -2.25)
                                                                    v54 := GenS7{F0: ((v45 * v45) == (v45 << ((uint(v1)) & 63)))}
                                                                    _ = v52
                                                                    _ = v53
                                                                    _ = v54
                                                                }
                                                            }
                                                            if true {
                                                                v1 = (^(v1))
                                                            } else {
                                                                if ((!false) && false) {
                                                                    v45 -= v45
                                                                } else {
                                                                    switch fi30 {
                                                                    case -1:
                                                                        v4 -= (v20 / ((v4) | 1))
                                                                    }
                                                                    v55 := []int{fi29, 128}
                                                                    if false {
                                                                        v5 = 76
                                                                    }
                                                                    _ = v55
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        switch (fi30 | 128) {
                                                        case 649:
                                                            v5 -= v46
                                                        case 127:
                                                            var v56 float32 = 0.0
                                                            _ = v56
                                                        case -574:
                                                            if false {
                                                                if true {
                                                                    delete(v3, 127)
                                                                } else {
                                                                    var v57 float64 = ((0.1 * float64(3.14159)) - -2.25)
                                                                    _ = v57
                                                                }
                                                                fmt.Println(v4, v20, v4)
                                                            } else {
                                                                switch "a" {
                                                                case "":
                                                                    if false {
                                                                        var v58 uint16 = 859
                                                                        switch -785 {
                                                                        case -342:
                                                                            if ((v20 >> ((1) & 63)) < (v20 + (^(v4)))) {
                                                                                v2 = 42
                                                                            } else {
                                                                                v2 = (v2 >> ((2) & 63))
                                                                            }
                                                                        case -383:
                                                                            switch (fi30 - int(v2)) {
                                                                            case 267:
                                                                                switch (-(fi30)) {
                                                                                case 941:
                                                                                    var v59 uint32 = (v1 % ((4294967295) | 1))
                                                                                    _ = v59
                                                                                case 89:
                                                                                    v45 -= -1
                                                                                }
                                                                            case 0:
                                                                                v60 := GenS2{F0: v58, F1: (v45 + v45)}
                                                                                _ = v60
                                                                            case 700:
                                                                                fmt.Println(v2)
                                                                            }
                                                                        case 32768:
                                                                            fmt.Println(v46, v46)
                                                                        }
                                                                        if false {
                                                                            _ = len(v3)
                                                                            if false {
                                                                                v4 = v20
                                                                                var v61 int16 = -1
                                                                                _ = v61
                                                                            }
                                                                            v62 := GenS1{F0: v5, F1: 1}
                                                                            _ = v62
                                                                        }
                                                                        fmt.Println(fi30, v58, fi30)
                                                                        var v63 uint64 = v5
                                                                        _ = v58
                                                                        _ = v63
                                                                    }
                                                                }
                                                                if (true || (uint16(v46) != 744)) {
                                                                    v45 -= (v45 << ((816) & 63))
                                                                    v64 := GenS2{F0: 752, F1: v45}
                                                                    if (-2.25 < 0.0) {
                                                                        if true {
                                                                            var v65 float64 = 3.14159
                                                                            v66 := v3[-321]
                                                                            v67 := GenS3{F0: 504, F1: v20, F2: "a"}
                                                                            var v68 int = (^(fi29))
                                                                            _ = v65
                                                                            _ = v66
                                                                            _ = v67
                                                                            _ = v68
                                                                        } else {
                                                                            v1 = (v1 ^ v1)
                                                                        }
                                                                        switch "" {
                                                                        case "hello":
                                                                            var v69 uint32 = v1
                                                                            _ = v69
                                                                        case "foo bar":
                                                                            v70 := GenS9{F0: 65535}
                                                                            _ = v70
                                                                        default:
                                                                            if false {
                                                                                v71 := GenS7{F0: false}
                                                                                if (false || false) {
                                                                                    v5 = v5
                                                                                    v72 := GenS4{F0: (!false), F1: (v45 ^ 100), F2: int(v4)}
                                                                                    _ = v72
                                                                                } else {
                                                                                    var v73 int16 = -32768
                                                                                    _ = v73
                                                                                }
                                                                                var v74 uint = 378
                                                                                _ = v71
                                                                                _ = v74
                                                                            }
                                                                        }
                                                                        var v75 int64 = (v2 + (v2 >> ((42) & 63)))
                                                                        _ = v75
                                                                    } else {
                                                                        v20 = 1
                                                                        fmt.Println(v45)
                                                                        v4 = (v4 | v20)
                                                                    }
                                                                    _ = v64
                                                                } else {
                                                                    delete(v3, 4294967295)
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                                v76 := GenS1{F0: (v5 >> ((255) & 63)), F1: 1000}
                                                _ = v76
                                            }
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                        }
                                        if ((3.14159 / -2.25) == -2.25) {
                                            v1 -= v1
                                            if true {
                                                var v77 uint = 785
                                                if false {
                                                    var v78 bool = false
                                                    v79 := GenS5{F0: "foo bar", F1: int(v1)}
                                                    v80 := []int{(^(fi29)), fi30, (^(fi29)), (fi29 ^ fi30)}
                                                    _ = v78
                                                    _ = v79
                                                    _ = v80
                                                }
                                                v81 := GenS8{F0: 42}
                                                var v82 bool = false
                                                _ = v77
                                                _ = v81
                                                _ = v82
                                            }
                                            v4 = 2
                                        } else {
                                            v1 += (v1 / (((v1 - 2)) | 1))
                                            if true {
                                                fmt.Println(v20)
                                                v83 := GenS5{F0: "a", F1: 127}
                                                v1 += v1
                                                _ = v83
                                            }
                                            switch "" {
                                            case "":
                                                v1 = (v1 % ((v1) | 1))
                                            default:
                                                v84 := GenS10{F0: 1, F1: 78, F2: ("" + "foo bar")}
                                                _ = v84
                                            }
                                        }
                                    } else {
                                        var v85 int8 = 57
                                        delete(v3, (fi29 | (fi30 >> ((1) & 63))))
                                        _ = v85
                                    }
                                    switch -1 {
                                    case -632:
                                        var v86 int = -1
                                        _ = v86
                                    case -2147483648:
                                        fmt.Println(v20, v20)
                                    }
                                } else {
                                    v87 := GenS11{F0: 32767}
                                    _ = v87
                                }
                                v3[fi29] = (fi29 & (fi30 & -258))
                            }
                        default:
                            v20 -= uint8(v5)
                        }
                        var v88 string = "a"
                        _ = fi30
                        _ = v31
                        _ = v88
                    }
                    _ = fi29
                }
                for fi89 := 0; fi89 < 18; fi89++ {
                    v90 := v3[(fi89 * (fi89 % (((fi89 & fi89)) | 1)))]
                    switch fi89 {
                    case 0:
                        fmt.Println(v1, v90)
                    case 9223372036854775807:
                        v2 = -9223372036854775808
                    case -598:
                        var v91 uint16 = uint16(v1)
                        _ = v91
                    }
                    _ = fi89
                    _ = v90
                }
            }
        }
        _ = v20
    }
    var v92 int8 = 100
    if false {
        fmt.Println(v4)
        for fi93 := 0; fi93 < 1; fi93++ {
            var v94 string = "\n\t"
            for fi95 := 0; fi95 < 3; fi95++ {
                if false {
                    if ((true && true) || false) {
                        v1 = (^(v1))
                        v96 := GenS3{F0: 2147483647, F1: (v4 << ((65535) & 63)), F2: v94}
                        _ = v96
                    }
                } else {
                    v1 += (v1 * 42)
                    fmt.Println(v1)
                }
                switch (^(fi95)) {
                case 1:
                    switch v94 {
                    case "hello":
                        v4 = 0
                    case "":
                        v97 := GenS12{F0: ("a" + ("a" + v94)), F1: -1}
                        _ = v97
                    case "foo bar":
                        var v98 float64 = 1e10
                        switch fi93 {
                        case 1:
                            var v99 uint16 = 255
                            _ = v99
                        }
                        _ = v98
                    }
                    var v100 bool = (!(("" + v94) >= v94))
                    v101 := GenS8{F0: -96}
                    _ = v100
                    _ = v101
                case 4294967295:
                    if false {
                        var v102 int = (fi95 - (fi93 << ((288) & 63)))
                        _ = v102
                    } else {
                        delete(v3, fi93)
                        fmt.Println(v92, v92, v1)
                    }
                default:
                    if (false && ((!((v4 << ((18446744073709551615) & 63)) >= (v4 % ((v4) | 1)))) || false)) {
                        v1 = (v1 >> ((42) & 63))
                    } else {
                        var v103 int32 = 42
                        v2 = v2
                        _ = v103
                    }
                }
                if false {
                    if true {
                        v4 = v4
                    }
                } else {
                    v104 := GenS4{F0: ((fi95 + int(v92)) <= fi95), F1: (v92 ^ v92), F2: (fi95 & fi93)}
                    fmt.Println(v4, v5, v1)
                    v3[fi95] = (fi95 ^ 32767)
                    _ = v104
                }
                v105 := GenS13{F0: v94, F1: (v1 ^ 2147483647)}
                _ = fi95
                _ = v105
            }
            _ = fi93
            _ = v94
        }
        for fi106 := 0; fi106 < 13; fi106++ {
            if false {
                for fi107 := 0; fi107 < 10; fi107++ {
                    if ((-(v2)) != (v2 * (v2 ^ v2))) {
                        var v108 uint64 = (v5 * v5)
                        if false {
                            switch "a" {
                            case "foo bar":
                                v5 = (v5 * v108)
                            }
                            delete(v3, -140)
                        }
                        _ = v108
                    }
                    if (255 > 408) {
                        switch (fi106 + -650) {
                        case 32768:
                            switch "\n\t" {
                            case "\n\t":
                                if (v5 == v5) {
                                    switch "" {
                                    case "foo bar":
                                        var v109 int8 = v92
                                        _ = v109
                                    case "\n\t":
                                        delete(v3, (fi106 + (-(fi106))))
                                    default:
                                        v110 := []int{-2}
                                        _ = v110
                                    }
                                    switch "" {
                                    case "\n\t":
                                        v3[fi107] = (fi106 / ((fi106) | 1))
                                    }
                                    var v111 int8 = v92
                                    _ = v111
                                } else {
                                    _ = len(v3)
                                }
                            default:
                                switch "hello" {
                                case "\n\t":
                                    if (!true) {
                                        var v112 float32 = (0.0 * 0.1)
                                        _ = v112
                                    }
                                default:
                                    if (0.1 < -2.25) {
                                        v113 := []int{fi106, fi107}
                                        _ = v113
                                    } else {
                                        v114 := GenS6{F0: 0, F1: (v5 | 282), F2: 2}
                                        _ = v114
                                    }
                                }
                            }
                        }
                    }
                    delete(v3, -2147483648)
                    v115 := GenS14{F0: "hello", F1: 0, F2: ((("" + "a") + (("" + "") + "")) + "")}
                    var v116 int32 = -2147483648
                    _ = fi107
                    _ = v115
                    _ = v116
                }
                fmt.Println(v1)
                v117 := GenS8{F0: 42}
                _ = v117
            } else {
                var v118 uint32 = v1
                switch ("foo bar" + "") {
                case "foo bar":
                    v4 += 71
                case "hello":
                    var v119 uint8 = 128
                    _ = v119
                default:
                    if true {
                        if (-32768 > 1) {
                            var v120 uint32 = (^(v1))
                            _ = v120
                        } else {
                            fmt.Println(v118)
                            for fi121 := 0; fi121 < 4; fi121++ {
                                fmt.Println(v92)
                                v1 = (v118 % ((2147483647) | 1))
                                _ = fi121
                            }
                            var v122 float32 = 0.0
                            _ = v122
                        }
                    } else {
                        v2 = (v2 * (v2 * 9223372036854775807))
                    }
                }
                if false {
                    for fi123 := 0; fi123 < 10; fi123++ {
                        if (v4 >= 42) {
                            v124 := GenS15{F0: 127, F1: -33, F2: int32(v5)}
                            if true {
                                v118 = (^(v1))
                                var v125 uint8 = 42
                                _ = v125
                            }
                            _ = v124
                        } else {
                            switch (^(fi123)) {
                            case 128:
                                switch ("foo bar" + "") {
                                case "hello":
                                    fmt.Println(v4)
                                case "\n\t":
                                    var v126 uint64 = (v5 * 381)
                                    _ = v126
                                default:
                                    switch fi106 {
                                    case 0:
                                        if (false || (2 != 275)) {
                                            v2 += (v2 / ((-56) | 1))
                                            var v127 string = "a"
                                            _ = v127
                                        }
                                    default:
                                        fmt.Println(v118)
                                    }
                                }
                            case 350:
                                if ((^(fi123)) > 127) {
                                    var v128 float32 = 3.14159
                                    _ = v128
                                }
                            case -10:
                                if true {
                                    if (v2 == v2) {
                                        delete(v3, (fi106 >> ((2) & 63)))
                                        _ = len(v3)
                                        fmt.Println(v4, v1)
                                    } else {
                                        var v129 uint64 = (v5 % ((0) | 1))
                                        v130 := []int{}
                                        _ = v129
                                        _ = v130
                                    }
                                } else {
                                    switch "foo bar" {
                                    case "":
                                        if true {
                                            switch (^(fi123)) {
                                            case 564:
                                                v1 -= (v118 - uint32(v4))
                                            default:
                                                v131 := []int{}
                                                _ = v131
                                            }
                                            v132 := GenS8{F0: -1}
                                            _ = v132
                                        } else {
                                            fmt.Println(v118)
                                            var v133 float64 = 3.14159
                                            _ = v133
                                        }
                                    case "hello":
                                        switch (fi106 << ((42) & 63)) {
                                        case 79:
                                            v134 := GenS16{F0: "", F1: (v118 - v1)}
                                            _ = v134
                                        case 648:
                                            var v135 uint32 = 124
                                            _ = v135
                                        default:
                                            v92 = (v92 + (v92 & v92))
                                        }
                                    case "\n\t":
                                        delete(v3, fi123)
                                    default:
                                        var v136 bool = ((true || false) && false)
                                        _ = v136
                                    }
                                    if (v118 == v118) {
                                        var v137 uint = 114
                                        var v138 uint8 = v4
                                        var v139 uint32 = v118
                                        _ = v137
                                        _ = v138
                                        _ = v139
                                    } else {
                                        var v140 uint32 = v1
                                        v141 := []int{(^(fi123)), (fi106 / ((fi123) | 1)), fi106, fi123}
                                        if true {
                                            var v142 uint8 = v4
                                            var v143 float64 = 3.14159
                                            if ((v5 << ((0) & 63)) > v5) {
                                                var v144 int8 = v92
                                                _ = v144
                                            }
                                            var v145 uint8 = v142
                                            _ = v142
                                            _ = v143
                                            _ = v145
                                        }
                                        if (-713 >= 574) {
                                            var v146 int16 = 367
                                            _ = v146
                                        }
                                        _ = v140
                                        _ = v141
                                    }
                                    var v147 uint32 = 127
                                    _ = v147
                                }
                            }
                            v92 += 127
                            _ = len(v3)
                            if ((false || (true || (((^(v2)) >= (v2 >> ((2) & 63))) && true))) || false) {
                                v4 = (v4 * uint8(v5))
                                var v148 float32 = 1.5
                                _ = v148
                            }
                            switch ("hello" + "foo bar") {
                            case "":
                                var v149 float32 = 1.5
                                _ = v149
                            default:
                                fmt.Println(v4)
                            }
                        }
                        v2 -= 9223372036854775807
                        _ = fi123
                    }
                    if (!(!((!(false || (v1 > uint32(v5)))) && (42 <= 1)))) {
                        if ((((v118 ^ (^(v118))) != v118) && (!false)) || true) {
                            v118 += v118
                        } else {
                            if (!(100 != (v92 * 42))) {
                                for fi150 := 0; fi150 < 18; fi150++ {
                                    var v151 bool = (!(v5 > v5))
                                    _ = fi150
                                    _ = v151
                                }
                                v5 = (^(v5))
                            }
                        }
                    } else {
                        switch fi106 {
                        case -696:
                            var v152 int64 = v2
                            _ = v152
                        case 2147483647:
                            v118 = 4294967295
                        }
                        var v153 uint16 = 65535
                        _ = v153
                    }
                }
                _ = v118
            }
            _ = fi106
        }
    }
    v1 = v1
    switch "\n\t" {
    case "":
        var v154 uint = 286
        _ = v154
    default:
        v155 := GenS7{F0: ((uint32(v5) >= (^(v1))) && false)}
        _ = v155
    }
    if (false || (128 > 42)) {
        v2 -= (v2 ^ v2)
        for fi156 := 0; fi156 < 20; fi156++ {
            var v157 string = ""
            v158 := GenS7{F0: (!true)}
            v157 = (v157 + "hello")
            _ = fi156
            _ = v157
            _ = v158
        }
    } else {
        v159 := GenS17{F0: 1, F1: 0}
        var v160 uint16 = 1
        v161 := GenS18{F0: 42, F1: (true && (("foo bar" + "\n\t") == ("\n\t" + ""))), F2: -32768}
        _ = v159
        _ = v160
        _ = v161
    }
    var v162 string = "\n\t"
    var v163 int64 = v2
    v1 -= (v1 % ((v1) | 1))
    if false {
        v92 = 42
        for fi164 := 0; fi164 < 1; fi164++ {
            for fi165 := 0; fi165 < 14; fi165++ {
                v5 = 199
                var v166 uint32 = (v1 & v1)
                _ = fi165
                _ = v166
            }
            _ = fi164
        }
    }
    fmt.Println(v163)
    fmt.Println(v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v162:", v162)
    fmt.Println("v163:", v163)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v162
    _ = v163
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v92
}
