package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 uint8; F1 int16; F2 int32 }
type GenS8 struct { F0 int8; F1 int32 }
type GenS9 struct { F0 int; F1 int32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 bool; F1 uint16; F2 int16 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int32; F1 int; F2 uint }
type GenS14 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS15 struct { F0 string; F1 uint }
type GenS16 struct { F0 uint8; F1 uint8 }
type GenS17 struct { F0 uint16; F1 uint8; F2 int }
type GenS18 struct { F0 uint; F1 int16 }
type GenS19 struct { F0 uint32 }
type GenS20 struct { F0 int32 }
type GenS21 struct { F0 int32; F1 bool; F2 int8 }
type GenS22 struct { F0 int; F1 uint8; F2 uint64 }
type GenS23 struct { F0 uint64; F1 int32 }
type GenS24 struct { F0 uint16; F1 uint64 }
type GenS25 struct { F0 uint8 }
type GenS26 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS27 struct { F0 uint }
type GenS28 struct { F0 bool; F1 int8; F2 int }
type GenS29 struct { F0 uint64; F1 int16; F2 string }
type GenS30 struct { F0 uint32; F1 int8; F2 int8 }
type GenS31 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS32 struct { F0 int8; F1 int8 }

func main() {
    var v1 int8 = 127
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v4 bool = true
    var v5 int64 = 571
    var v6 int16 = 583
    if (42 <= 100) {
        switch "foo bar" {
        case "a":
            var v7 float32 = 1e10
            _ = v7
        case "foo bar":
            v8 := GenS1{F0: 65535, F1: (4294967295 < 18446744073709551615)}
            _ = v8
        default:
            v4 = (3.14159 != 1e10)
        }
        v4 = true
        if (("a" + ("hello" + ("foo bar" + "foo bar"))) <= "hello") {
            v4 = (!false)
        } else {
            if v4 {
                if v4 {
                    v9 := GenS2{F0: 345, F1: (v1 % ((1) | 1)), F2: 1}
                    for fi10 := 0; fi10 < 4; fi10++ {
                        var v11 uint8 = 0
                        v12 := v3[2]
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    var v13 uint64 = 1
                    _ = v13
                    _ = v9
                }
                var v14 int8 = 0
                v15 := []int{950}
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 11; fi16++ {
                switch (("hello" + "hello") + "hello") {
                case "a":
                    v5 = v5
                case "":
                    v6 += (^(v6))
                default:
                    v17 := GenS2{F0: 1, F1: int8(v6), F2: 18446744073709551615}
                    _ = v17
                }
                switch "foo bar" {
                case "a":
                    switch ((("hello" + ("hello" + "\n\t")) + "") + "\n\t") {
                    case "":
                        v18 := GenS2{F0: 18446744073709551615, F1: v1, F2: 178}
                        _ = v18
                    case "\n\t":
                        var v19 bool = (((v6 & 1000) > (v6 & int16(v1))) && (0 < 178))
                        _ = v19
                    }
                default:
                    var v20 uint64 = 132
                    _ = v20
                }
                _ = fi16
            }
        }
    } else {
        switch 127 {
        case 1:
            for fi21 := 0; fi21 < 13; fi21++ {
                v22 := v2[42]
                switch fi21 {
                case 0:
                    var v23 uint64 = 18446744073709551615
                    _ = v23
                case 128:
                    if (18446744073709551615 >= 4294967295) {
                        var v24 int16 = v6
                        v3[42] = (v4 || v4)
                        if (856 > 1) {
                            var v25 int = (v22 - 95)
                            var v26 uint8 = 32
                            v27 := []int{v22, v22, v22, (-(v25))}
                            _ = v25
                            _ = v26
                            _ = v27
                        } else {
                            v5 -= (v5 & (v5 - v5))
                        }
                        _ = v24
                    } else {
                        switch "\n\t" {
                        case "":
                            v4 = (894 == 480)
                        }
                        var v28 uint32 = 2147483648
                        v29 := GenS3{F0: (v28 | 0), F1: 42}
                        _ = v28
                        _ = v29
                    }
                }
                for fi30 := 0; fi30 < 1; fi30++ {
                    v31 := []int{-9223372036854775808}
                    var v32 int = (fi30 & (-(v22)))
                    switch v32 {
                    case 42:
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                    }
                    v33 := GenS2{F0: 18446744073709551615, F1: v1, F2: 0}
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
                _ = fi21
                _ = v22
            }
        case 2147483647:
            for fi34 := 0; fi34 < 10; fi34++ {
                var v35 int = (fi34 | fi34)
                v36 := v2[2147483647]
                var v37 int8 = v1
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
            }
        case 138:
            fmt.Println(v4)
        default:
            v38 := []int{}
            _ = v38
        }
    }
    if v4 {
        v39 := GenS1{F0: 2147483648, F1: (9223372036854775807 <= 127)}
        _ = v39
    }
    v1 = v1
    for fi40 := 0; fi40 < 18; fi40++ {
        delete(v2, 963)
        v6 += v6
        var v41 int = fi40
        _ = fi40
        _ = v41
    }
    v5 += v5
    v42 := GenS3{F0: 4294967295, F1: 4294967295}
    v43 := GenS4{F0: 32768}
    if v4 {
        var v44 bool = v4
        var v45 float64 = (((float64(1.5) + ((1e10 - 0.0) - 1.5)) + (float64((0.1 - (((3.14159 * 1e10) / 1e10) * (1.5 * (1.5 * 1.5))))) * 1e10)) / -2.25)
        _ = v44
        _ = v45
    }
    v5 = 0
    if (v4 && v4) {
        v6 = (^(v6))
        if ("foo bar" == "hello") {
            if v4 {
                fmt.Println(v4, v4)
            } else {
                var v46 uint16 = 344
                fmt.Println(v1, v1, v5)
                if v4 {
                    var v47 uint32 = 65535
                    delete(v3, (v47 + (v47 % ((v47) | 1))))
                    switch ("" + ("foo bar" + "hello")) {
                    case "foo bar":
                        delete(v2, (v47 & (^(v47))))
                    case "hello":
                        var v48 uint = 0
                        _ = v48
                    default:
                        if v4 {
                            v49 := []int{2147483647}
                            for fi50 := 0; fi50 < 6; fi50++ {
                                if len(v49) > 0 {
                                    v49[0] = 4294967295
                                }
                                var v51 int8 = v1
                                var v52 uint64 = 2
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                            v46 = v46
                            _ = v49
                        }
                    }
                    _ = v47
                } else {
                    var v53 int32 = 2147483647
                    var v54 uint = 490
                    _ = v53
                    _ = v54
                }
                _ = v46
            }
            var v55 string = ("" + ("foo bar" + ("hello" + (("hello" + "\n\t") + ""))))
            _ = v55
        } else {
            for fi56 := 0; fi56 < 4; fi56++ {
                var v57 int64 = 42
                fmt.Println(v4, fi56, fi56)
                var v58 uint8 = 42
                _ = fi56
                _ = v57
                _ = v58
            }
            switch "" {
            case "":
                v5 += v5
            case "hello":
                v59 := GenS4{F0: -492}
                _ = v59
            case "\n\t":
                switch 686 {
                case -856:
                    switch 2147483647 {
                    case 32767:
                        switch -113 {
                        case -284:
                            if v4 {
                                v60 := GenS3{F0: 2147483648, F1: -940}
                                v61 := GenS5{F0: 9223372036854775807}
                                v62 := []int{}
                                _ = v60
                                _ = v61
                                _ = v62
                            }
                        case 2147483647:
                            v2[774] = -900
                        case 9223372036854775807:
                            if v4 {
                                for fi63 := 0; fi63 < 14; fi63++ {
                                    switch "a" {
                                    case "hello":
                                        v64 := GenS3{F0: 2147483647, F1: (fi63 >> ((4294967295) & 63))}
                                        _ = v64
                                    }
                                    v4 = true
                                    v5 = v5
                                    _ = fi63
                                }
                            }
                        default:
                            fmt.Println(v4, v5)
                        }
                    case 128:
                        v4 = true
                    case 990:
                        if (v6 <= (^(v6))) {
                            var v65 int8 = v1
                            v66 := []int{32767}
                            _ = v65
                            _ = v66
                        } else {
                            for fi67 := 0; fi67 < 5; fi67++ {
                                var v68 uint8 = 0
                                v69 := GenS6{F0: (v68 - (v68 / (((v68 - 128)) | 1))), F1: 2}
                                v4 = ((!(v1 != (v1 * (v1 ^ (v1 ^ (-(v1))))))) || ((!(((!(!(-9223372036854775808 > -9223372036854775808))) && true) && false)) && ((fi67 - (^(fi67))) == -796)))
                                _ = fi67
                                _ = v68
                                _ = v69
                            }
                            v70 := GenS4{F0: 127}
                            _ = v70
                        }
                    default:
                        var v71 bool = v4
                        _ = v71
                    }
                case 0:
                    for fi72 := 0; fi72 < 14; fi72++ {
                        var v73 int = (-(fi72))
                        v1 = (v1 >> ((65535) & 63))
                        _ = fi72
                        _ = v73
                    }
                }
            }
        }
    } else {
        v3[4294967295] = ((v6 - (v6 % ((v6) | 1))) < v6)
        var v74 uint32 = 2147483647
        _ = v74
    }
    for fi75 := 0; fi75 < 15; fi75++ {
        var v76 uint64 = 18446744073709551615
        if ((0 >= 9223372036854775807) || (v4 || v4)) {
            for fi77 := 0; fi77 < 16; fi77++ {
                v1 = v1
                _ = fi77
            }
            delete(v3, 828)
        }
        _ = fi75
        _ = v76
    }
    if ((0.0 - (1e10 + ((float32((float64(((1e10 * 1e10) / 0.1)) + 0.0)) + -2.25) * (-2.25 * float32(float64((((((3.14159 / ((1.5 - 0.0) - (3.14159 + 0.0))) * (0.0 / 1.0)) / 0.1) + (0.0 - 3.14159)) / 3.14159))))))) < ((1.5 * float32((0.0 * -2.25))) - (3.14159 - 1.5))) {
        switch "\n\t" {
        case "\n\t":
            for fi78 := 0; fi78 < 11; fi78++ {
                for fi79 := 0; fi79 < 3; fi79++ {
                    switch fi79 {
                    case 9223372036854775807:
                        if (v4 || v4) {
                            fmt.Println(v1, v6, v4)
                            v1 = v1
                        }
                    }
                    var v80 int64 = 1
                    v81 := GenS7{F0: 255, F1: (v6 << ((18446744073709551615) & 63)), F2: 2147483647}
                    fmt.Println(fi78, v4, v80)
                    var v82 uint = 255
                    switch fi79 {
                    case 0:
                        if ((3.14159 / (3.14159 / 1.5)) < 1e10) {
                            switch (fi79 % ((0) | 1)) {
                            case 128:
                                fmt.Println(v4)
                            case -132:
                                switch 0 {
                                case 128:
                                    var v83 int16 = v6
                                    var v84 int16 = v6
                                    _ = v83
                                    _ = v84
                                case 2147483647:
                                    if false {
                                        switch fi78 {
                                        case 127:
                                            switch 127 {
                                            case 32768:
                                                v85 := GenS8{F0: 27, F1: -60}
                                                _ = v85
                                            case 603:
                                                if (!(-2.25 >= 0.0)) {
                                                    if false {
                                                        var v86 uint32 = 662
                                                        _ = v86
                                                    } else {
                                                        v80 = v80
                                                        var v87 string = ("hello" + "hello")
                                                        _ = v87
                                                    }
                                                } else {
                                                    fmt.Println(v82)
                                                    switch "a" {
                                                    case "foo bar":
                                                        fmt.Println(v1, v82)
                                                    }
                                                }
                                            default:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v88 := GenS9{F0: fi79, F1: 2147483647}
                                                    _ = v88
                                                default:
                                                    switch 0 {
                                                    case 1:
                                                        switch ("a" + "hello") {
                                                        case "\n\t":
                                                            v82 += (v82 >> (((^(v82))) & 63))
                                                        default:
                                                            switch ("hello" + ((("" + ("" + "")) + "a") + "a")) {
                                                            case "\n\t":
                                                                switch "hello" {
                                                                case "":
                                                                    v89 := []int{fi78, fi79}
                                                                    _ = v89
                                                                case "\n\t":
                                                                    if false {
                                                                        v1 += -128
                                                                        var v90 bool = (0 >= v1)
                                                                        _ = v90
                                                                    }
                                                                default:
                                                                    v91 := GenS10{F0: (v82 * (v82 % ((71) | 1)))}
                                                                    _ = v91
                                                                }
                                                            case "":
                                                                if (131 <= 0) {
                                                                    if ((-394 <= -2147483648) || ((v1 & 42) != -128)) {
                                                                        if v4 {
                                                                            if ((v80 >> (((v82 + (v82 - v82))) & 63)) <= (-(v5))) {
                                                                                v6 += 1000
                                                                                v5 += (v80 * 9223372036854775807)
                                                                            } else {
                                                                                v92 := []int{fi79, fi79, 128, 976}
                                                                                _ = v92
                                                                            }
                                                                            switch fi79 {
                                                                            case 9223372036854775807:
                                                                                v93 := GenS3{F0: 404, F1: 155}
                                                                                _ = v93
                                                                            case 128:
                                                                                fmt.Println(fi79)
                                                                            default:
                                                                                switch (-(fi78)) {
                                                                                case -88:
                                                                                    var v94 uint8 = 42
                                                                                    _ = v94
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    v5 = v80
                                                                    v95 := []int{fi78}
                                                                    var v96 string = "a"
                                                                    _ = v95
                                                                    _ = v96
                                                                } else {
                                                                    _ = len(v3)
                                                                    var v97 int16 = v6
                                                                    _ = v97
                                                                }
                                                            }
                                                        }
                                                    case -1:
                                                        fmt.Println(fi78, v4)
                                                    case 32768:
                                                        if (!(false || v4)) {
                                                            switch (fi79 / ((fi79) | 1)) {
                                                            case -313:
                                                                var v98 uint64 = 18446744073709551615
                                                                _ = v98
                                                            default:
                                                                fmt.Println(v82, v5, v5)
                                                            }
                                                            v99 := GenS2{F0: 0, F1: (v1 | int8(v5)), F2: 392}
                                                            v5 -= v80
                                                            _ = v99
                                                        } else {
                                                            v3[627] = ((v82 * 0) > v82)
                                                        }
                                                    }
                                                }
                                            }
                                        case 73:
                                            delete(v3, 1)
                                        default:
                                            if v4 {
                                                var v100 float32 = float32(3.14159)
                                                var v101 uint16 = 42
                                                _ = v100
                                                _ = v101
                                            } else {
                                                fmt.Println(v5)
                                                v1 += (v1 & (^(v1)))
                                                var v102 string = ""
                                                if ((1 >= 255) || (!(!(!(2147483647 >= int32(fi78)))))) {
                                                    switch "foo bar" {
                                                    case "a":
                                                        if (3.14159 < (float32((0.0 / 1.0)) + (((-2.25 + 1.5) * -2.25) / ((3.14159 - (-2.25 - float32(0.0))) + 0.1)))) {
                                                            delete(v2, 42)
                                                            switch v102 {
                                                            case "hello":
                                                                v103 := GenS11{F0: true, F1: uint16(v82), F2: (^(v6))}
                                                                _ = v103
                                                            case "a":
                                                                fmt.Println(v5, v80, v80)
                                                            case "":
                                                                if v4 {
                                                                    v104 := GenS3{F0: 2, F1: (fi78 * fi78)}
                                                                    _ = v104
                                                                } else {
                                                                    fmt.Println(v4, fi78)
                                                                    var v105 string = v102
                                                                    _ = v105
                                                                }
                                                            default:
                                                                var v106 int8 = (v1 >> ((v82) & 63))
                                                                _ = v106
                                                            }
                                                            if (false && v4) {
                                                                switch fi79 {
                                                                case -2147483648:
                                                                    switch (v102 + "a") {
                                                                    case "a":
                                                                        v107 := GenS12{F0: v1}
                                                                        _ = v107
                                                                    }
                                                                case -567:
                                                                    v5 = (v80 % ((int64(v6)) | 1))
                                                                case -828:
                                                                    fmt.Println(v6, v82)
                                                                }
                                                                v108 := GenS13{F0: 2147483647, F1: fi78, F2: (v82 / (((^(v82))) | 1))}
                                                                var v109 uint64 = 42
                                                                _ = v108
                                                                _ = v109
                                                            } else {
                                                                v110 := v2[4294967295]
                                                                if (v6 == -32768) {
                                                                    v111 := GenS12{F0: int8(fi79)}
                                                                    fmt.Println(v110, v82, v82)
                                                                    fmt.Println(v102, fi79, v4)
                                                                    _ = v111
                                                                } else {
                                                                    var v112 int32 = -1
                                                                    v113 := GenS14{F0: 65535, F1: 0, F2: 2}
                                                                    _ = v112
                                                                    _ = v113
                                                                }
                                                                _ = v110
                                                            }
                                                        } else {
                                                            if ((fi79 % (((fi78 * -744)) | 1)) < (fi79 >> ((18446744073709551615) & 63))) {
                                                                if (!(!(((1.5 + 0.0) * 3.14159) >= 0.0))) {
                                                                    v114 := GenS15{F0: "", F1: v82}
                                                                    delete(v2, 381)
                                                                    _ = v114
                                                                } else {
                                                                    v80 = (v80 % ((1) | 1))
                                                                    v3[2147483648] = v4
                                                                }
                                                            } else {
                                                                fmt.Println(v4, fi78)
                                                            }
                                                        }
                                                    }
                                                    var v115 int = (fi78 >> ((402) & 63))
                                                    v82 -= (^(v82))
                                                    _ = v115
                                                } else {
                                                    var v116 uint8 = 163
                                                    v116 = (v116 / (((v116 >> (((v82 | (v82 & v82))) & 63))) | 1))
                                                    if v4 {
                                                        fmt.Println(fi78, v102, fi79)
                                                    } else {
                                                        v117 := GenS16{F0: (v116 << ((uint(v116)) & 63)), F1: v116}
                                                        fmt.Println(v116, v82)
                                                        _ = v117
                                                    }
                                                    if (v4 && (-2.25 > -2.25)) {
                                                        var v118 uint32 = 919
                                                        fmt.Println(v116, v1, v116)
                                                        var v119 int32 = 85
                                                        var v120 int8 = 0
                                                        v119 = (v119 & v119)
                                                        if v4 {
                                                            switch (fi79 + int(v82)) {
                                                            case -2147483648:
                                                                v4 = (!(v80 <= (v80 ^ v80)))
                                                            default:
                                                                var v121 bool = true
                                                                _ = v121
                                                            }
                                                            var v122 int = fi79
                                                            _ = v122
                                                        } else {
                                                            v5 = (v5 - (v5 / (((v80 ^ (v5 % ((v80) | 1)))) | 1)))
                                                            v118 += 316
                                                            fmt.Println(v80, fi78)
                                                        }
                                                        v123 := GenS17{F0: 766, F1: v116, F2: fi78}
                                                        _ = v118
                                                        _ = v119
                                                        _ = v120
                                                        _ = v123
                                                    } else {
                                                        var v124 int16 = (v6 % ((v6) | 1))
                                                        v125 := GenS18{F0: 4294967295, F1: v6}
                                                        _ = v124
                                                        _ = v125
                                                    }
                                                    var v126 int16 = v6
                                                    _ = v116
                                                    _ = v126
                                                }
                                                _ = v102
                                            }
                                            switch "" {
                                            case "foo bar":
                                                var v127 uint8 = 127
                                                _ = v127
                                            case "\n\t":
                                                if ((v4 && v4) && (-77 <= 78)) {
                                                    v128 := []int{fi79, 9223372036854775807, int(v6)}
                                                    var v129 int32 = -141
                                                    _ = v128
                                                    _ = v129
                                                } else {
                                                    var v130 int32 = 42
                                                    var v131 int8 = (^(v1))
                                                    switch "" {
                                                    case "":
                                                        var v132 int32 = v130
                                                        _ = v132
                                                    case "\n\t":
                                                        v5 -= (-(v5))
                                                    }
                                                    v133 := GenS19{F0: 65535}
                                                    _ = v130
                                                    _ = v131
                                                    _ = v133
                                                }
                                            default:
                                                v2[854] = -855
                                            }
                                        }
                                        if ((-(fi79)) >= 2147483647) {
                                            if (!(-2147483648 > -2147483648)) {
                                                switch fi78 {
                                                case 0:
                                                    delete(v2, 2147483648)
                                                    var v134 string = "hello"
                                                    v134 = ("\n\t" + (v134 + v134))
                                                    _ = v134
                                                case 42:
                                                    if ((977 > 0) && (4294967295 <= 0)) {
                                                        if (42 >= (v1 * -64)) {
                                                            v135 := GenS2{F0: 4294967295, F1: (v1 ^ (-(v1))), F2: 148}
                                                            v136 := GenS8{F0: 100, F1: -1}
                                                            _ = v135
                                                            _ = v136
                                                        }
                                                    }
                                                    fmt.Println(fi78, v5)
                                                case 32768:
                                                    if (1 <= 1) {
                                                        v4 = ((v82 - (v82 + v82)) >= (v82 >> ((4294967295) & 63)))
                                                        fmt.Println(v1, v6, fi78)
                                                        var v137 int = (fi79 / ((fi79) | 1))
                                                        var v138 float64 = (1.5 / 3.14159)
                                                        fmt.Println(v1, v1)
                                                        _ = v137
                                                        _ = v138
                                                    }
                                                }
                                            }
                                            v139 := []int{(fi78 + (fi79 + 128))}
                                            _ = v139
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            if (255 > v82) {
                                                fmt.Println(v80, v5, v82)
                                                if ((v82 * (^(v82))) > v82) {
                                                    if false {
                                                        switch (fi78 >> ((381) & 63)) {
                                                        case 280:
                                                            v1 += (-(v1))
                                                        case 32767:
                                                            v80 += (v5 >> ((v82) & 63))
                                                        case 127:
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                v140 := GenS20{F0: -1}
                                                                if true {
                                                                    v141 := GenS17{F0: 65535, F1: 127, F2: 32767}
                                                                    v142 := GenS18{F0: v82, F1: (v6 * 42)}
                                                                    _ = v141
                                                                    _ = v142
                                                                }
                                                                v143 := []int{fi78, fi78, -856}
                                                                _ = v140
                                                                _ = v143
                                                            case "":
                                                                if v4 {
                                                                    var v144 uint8 = 0
                                                                    if (v4 && (818 < 660)) {
                                                                        var v145 int8 = -109
                                                                        if (0 != 255) {
                                                                            if v4 {
                                                                                var v146 int16 = -55
                                                                                _ = v146
                                                                            } else {
                                                                                var v147 int8 = v1
                                                                                _ = v147
                                                                            }
                                                                            if (!true) {
                                                                                switch "" {
                                                                                case "a":
                                                                                    fmt.Println(v6, v144)
                                                                                case "foo bar":
                                                                                    var v148 bool = v4
                                                                                    _ = v148
                                                                                }
                                                                                switch ((("\n\t" + "\n\t") + "\n\t") + "a") {
                                                                                case "a":
                                                                                    v82 -= (v82 % (((v82 / (((v82 + v82)) | 1))) | 1))
                                                                                case "\n\t":
                                                                                    v149 := GenS17{F0: 685, F1: (^(v144)), F2: int(v6)}
                                                                                    v5 = (v80 + (v80 + v5))
                                                                                    v3[0] = (("\n\t" + "\n\t") <= "\n\t")
                                                                                    _ = v149
                                                                                default:
                                                                                    fmt.Println(v82, v1)
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v145 = v145
                                                                            var v150 string = "\n\t"
                                                                            var v151 string = v150
                                                                            _ = v150
                                                                            _ = v151
                                                                        }
                                                                        v152 := []int{(fi79 >> ((v82) & 63)), fi78, 128}
                                                                        _ = v145
                                                                        _ = v152
                                                                    } else {
                                                                        v153 := []int{int(v80), (fi78 + (fi79 ^ (fi79 & (fi78 / (((fi78 - (fi78 >> ((v82) & 63)))) | 1))))), (fi79 + fi78), fi78}
                                                                        var v154 int32 = 0
                                                                        if len(v153) > 0 {
                                                                            v153[0] = (fi79 * 42)
                                                                        }
                                                                        _ = v153
                                                                        _ = v154
                                                                    }
                                                                    v5 -= v5
                                                                    _ = v144
                                                                } else {
                                                                    v155 := []int{}
                                                                    v4 = v4
                                                                    _ = v155
                                                                }
                                                                v82 = v82
                                                                v80 += v80
                                                            case "hello":
                                                                if (783 == 0) {
                                                                    var v156 float32 = float32(3.14159)
                                                                    v156 = v156
                                                                    _ = v156
                                                                }
                                                            default:
                                                                var v157 uint8 = 127
                                                                v158 := GenS21{F0: -711, F1: ((((fi79 + (fi78 | fi79)) < (fi79 % ((fi79) | 1))) || (!v4)) || (1 >= 1)), F2: v1}
                                                                v159 := []int{-9223372036854775808, fi79, (fi79 + (fi79 >> (((v82 >> ((v82) & 63))) & 63))), fi79}
                                                                _ = v157
                                                                _ = v158
                                                                _ = v159
                                                            }
                                                        default:
                                                            switch "a" {
                                                            case "":
                                                                v6 = v6
                                                            case "hello":
                                                                if false {
                                                                    if true {
                                                                        var v160 uint32 = 0
                                                                        v160 += v160
                                                                        v161 := []int{(fi79 & 2147483647), (fi79 * 225), (fi79 + fi78), fi79}
                                                                        _ = v160
                                                                        _ = v161
                                                                    } else {
                                                                        v162 := GenS22{F0: 9223372036854775807, F1: 221, F2: 18446744073709551615}
                                                                        v163 := GenS23{F0: 42, F1: 0}
                                                                        switch ("a" + "a") {
                                                                        case "a":
                                                                            var v164 int32 = -1
                                                                            _ = v164
                                                                        case "hello":
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                if ((42 >= v6) || v4) {
                                                                                    fmt.Println(v4, v80)
                                                                                    v165 := GenS17{F0: 397, F1: 127, F2: (^(fi78))}
                                                                                    v166 := []int{(-(fi79)), (fi79 | fi78), fi79}
                                                                                    switch "foo bar" {
                                                                                    case "a":
                                                                                        if false {
                                                                                            v167 := GenS24{F0: 65535, F1: 42}
                                                                                            _ = len(v2)
                                                                                            _ = v167
                                                                                        } else {
                                                                                            v5 = v80
                                                                                        }
                                                                                    case "":
                                                                                        fmt.Println(v4, fi78, v80)
                                                                                    case "\n\t":
                                                                                        v80 = int64(fi79)
                                                                                    }
                                                                                    _ = v165
                                                                                    _ = v166
                                                                                }
                                                                            default:
                                                                                v4 = (!(!v4))
                                                                            }
                                                                        case "\n\t":
                                                                            if v4 {
                                                                                v168 := GenS9{F0: (fi79 / (((fi79 / ((fi78) | 1))) | 1)), F1: 972}
                                                                                v169 := []int{fi78, (fi78 & fi78), fi79, 1}
                                                                                v5 = 1
                                                                                _ = v168
                                                                                _ = v169
                                                                            } else {
                                                                                if true {
                                                                                    var v170 uint16 = 1
                                                                                    _ = v170
                                                                                }
                                                                                var v171 int16 = v6
                                                                                fmt.Println(v82, v5)
                                                                                _ = v171
                                                                            }
                                                                        }
                                                                        _ = v162
                                                                        _ = v163
                                                                    }
                                                                    switch -149 {
                                                                    case -2147483648:
                                                                        v6 = (v6 ^ v6)
                                                                    case 1:
                                                                        v172 := GenS6{F0: 0, F1: 255}
                                                                        _ = v172
                                                                    }
                                                                    _ = len(v2)
                                                                } else {
                                                                    if false {
                                                                        v82 += v82
                                                                    } else {
                                                                        fmt.Println(v4, fi78, v82)
                                                                        v173 := []int{42, fi78, 128}
                                                                        v174 := GenS9{F0: fi79, F1: 0}
                                                                        _ = v173
                                                                        _ = v174
                                                                    }
                                                                    switch 32767 {
                                                                    case -9223372036854775808:
                                                                        v175 := GenS18{F0: v82, F1: v6}
                                                                        _ = v175
                                                                    case 275:
                                                                        var v176 uint32 = 176
                                                                        _ = v176
                                                                    }
                                                                }
                                                                if v4 {
                                                                    v177 := v2[1]
                                                                    _ = v177
                                                                }
                                                            case "foo bar":
                                                                var v178 int16 = v6
                                                                _ = v178
                                                            default:
                                                                v179 := GenS8{F0: v1, F1: 574}
                                                                if false {
                                                                    v4 = false
                                                                    v180 := GenS25{F0: 42}
                                                                    v181 := GenS23{F0: 441, F1: -2147483648}
                                                                    if v4 {
                                                                        var v182 uint16 = 65535
                                                                        var v183 float64 = -2.25
                                                                        _ = v182
                                                                        _ = v183
                                                                    } else {
                                                                        var v184 int8 = int8(fi79)
                                                                        if v4 {
                                                                            v185 := GenS16{F0: 27, F1: 128}
                                                                            if (-1 == 1) {
                                                                                var v186 int32 = 879
                                                                                if (!(!v4)) {
                                                                                    v1 = v1
                                                                                    if v4 {
                                                                                        var v187 int8 = v184
                                                                                        _ = len(v2)
                                                                                        _ = v187
                                                                                    } else {
                                                                                        var v188 uint16 = 65535
                                                                                        var v189 uint64 = 2
                                                                                        _ = v188
                                                                                        _ = v189
                                                                                    }
                                                                                } else {
                                                                                    var v190 uint16 = 2
                                                                                    if (963 >= 221) {
                                                                                        v82 = v82
                                                                                        v191 := []int{42, fi78}
                                                                                        switch fi79 {
                                                                                        case 32768:
                                                                                            v190 = (v190 | 65535)
                                                                                        case -1:
                                                                                            v191 = append(v191, (fi78 & 354))
                                                                                        default:
                                                                                            if v4 {
                                                                                                var v192 int = fi79
                                                                                                _ = v192
                                                                                            } else {
                                                                                                if true {
                                                                                                    var v193 uint = (v82 / ((v82) | 1))
                                                                                                    v4 = (!(((^(v186)) == v186) || (18446744073709551615 < 793)))
                                                                                                    _ = v193
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        switch fi78 {
                                                                                        case -2147483648:
                                                                                            v80 += -1
                                                                                        default:
                                                                                            var v194 int64 = v80
                                                                                            _ = v194
                                                                                        }
                                                                                        if true {
                                                                                            v186 -= 2147483647
                                                                                            v4 = (v4 || true)
                                                                                            v4 = v4
                                                                                        }
                                                                                        v82 -= (v82 % ((v82) | 1))
                                                                                        _ = v191
                                                                                    } else {
                                                                                        var v195 int8 = v1
                                                                                        fmt.Println(v5, v186, v4)
                                                                                        v1 = (v1 * 127)
                                                                                        _ = v195
                                                                                    }
                                                                                    fmt.Println(v184, fi79)
                                                                                    _ = v190
                                                                                }
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    v196 := []int{fi79, (fi78 % ((fi79) | 1)), fi79, (fi79 ^ fi79)}
                                                                                    _ = v196
                                                                                case "foo bar":
                                                                                    v197 := GenS24{F0: 1, F1: 0}
                                                                                    _ = v197
                                                                                default:
                                                                                    if (v4 || ((1e10 + float32(-2.25)) >= 1e10)) {
                                                                                        switch ("" + ((("a" + "\n\t") + "foo bar") + "hello")) {
                                                                                        case "":
                                                                                            v4 = true
                                                                                        case "a":
                                                                                            v198 := GenS5{F0: (fi79 >> ((v82) & 63))}
                                                                                            _ = v198
                                                                                        case "\n\t":
                                                                                            fmt.Println(v1, v5)
                                                                                        }
                                                                                        var v199 uint = (v82 / (((v82 / (((v82 >> ((1) & 63))) | 1))) | 1))
                                                                                        v200 := GenS23{F0: 120, F1: v186}
                                                                                        _ = v199
                                                                                        _ = v200
                                                                                    } else {
                                                                                        switch int(v82) {
                                                                                        case -9223372036854775808:
                                                                                            v6 -= v6
                                                                                        }
                                                                                        v2[2147483647] = (fi79 << ((v82) & 63))
                                                                                    }
                                                                                }
                                                                                var v201 bool = (((fi79 << (((v82 - v82)) & 63)) >= (fi79 * fi78)) || (v82 >= (^(v82))))
                                                                                v186 += (v186 & (v186 - 2147483647))
                                                                                _ = v186
                                                                                _ = v201
                                                                            }
                                                                            var v202 uint32 = 2147483647
                                                                            _ = v185
                                                                            _ = v202
                                                                        }
                                                                        v184 += int8(fi79)
                                                                        _ = v184
                                                                    }
                                                                    switch (^(fi78)) {
                                                                    case 0:
                                                                        v203 := []int{}
                                                                        _ = v203
                                                                    case 42:
                                                                        v80 = 42
                                                                    default:
                                                                        fmt.Println(v4)
                                                                    }
                                                                    switch ("a" + ("foo bar" + "")) {
                                                                    case "hello":
                                                                        if (2 == 0) {
                                                                            var v204 float64 = (float64(3.14159) * 1.5)
                                                                            v205 := GenS8{F0: v1, F1: 1}
                                                                            fmt.Println(v4, v6, v1)
                                                                            _ = v204
                                                                            _ = v205
                                                                        } else {
                                                                            var v206 int64 = v5
                                                                            _ = v206
                                                                        }
                                                                    case "foo bar":
                                                                        v207 := []int{}
                                                                        _ = v207
                                                                    case "":
                                                                        var v208 int32 = -745
                                                                        _ = v208
                                                                    default:
                                                                        var v209 uint64 = 18446744073709551615
                                                                        _ = v209
                                                                    }
                                                                    var v210 bool = true
                                                                    v211 := []int{381, (-(fi78))}
                                                                    var v212 string = "\n\t"
                                                                    _ = v180
                                                                    _ = v181
                                                                    _ = v210
                                                                    _ = v211
                                                                    _ = v212
                                                                }
                                                                _ = v179
                                                            }
                                                        }
                                                        _ = len(v2)
                                                    } else {
                                                        v213 := GenS26{F0: uint8(v82), F1: v1, F2: 0}
                                                        var v214 int64 = (-(v5))
                                                        if v4 {
                                                            v215 := []int{(fi78 - fi79), 189}
                                                            switch fi78 {
                                                            case 743:
                                                                var v216 uint = (v82 | v82)
                                                                _ = v216
                                                            case 128:
                                                                v214 -= (v5 / (((v214 << ((v82) & 63))) | 1))
                                                            case 0:
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    var v217 int16 = -32768
                                                                    _ = v217
                                                                case "foo bar":
                                                                    v218 := GenS26{F0: 1, F1: (v1 ^ v1), F2: 2147483647}
                                                                    _ = v218
                                                                case "a":
                                                                    if (!((-(v6)) == (^(v6)))) {
                                                                        var v219 bool = (0 <= 127)
                                                                        var v220 uint32 = 510
                                                                        if (true || (0 < 549)) {
                                                                            var v221 string = ("\n\t" + ((("a" + "a") + ("" + ("a" + ""))) + "foo bar"))
                                                                            _ = v221
                                                                        } else {
                                                                            var v222 string = "hello"
                                                                            var v223 int32 = 416
                                                                            _ = v222
                                                                            _ = v223
                                                                        }
                                                                        _ = v219
                                                                        _ = v220
                                                                    } else {
                                                                        var v224 uint64 = 1
                                                                        var v225 int32 = 42
                                                                        _ = v224
                                                                        _ = v225
                                                                    }
                                                                default:
                                                                    var v226 bool = (226 >= 42)
                                                                    _ = v226
                                                                }
                                                            default:
                                                                if false {
                                                                    var v227 int = (fi79 << ((v82) & 63))
                                                                    var v228 int64 = 1
                                                                    var v229 int8 = 100
                                                                    _ = v227
                                                                    _ = v228
                                                                    _ = v229
                                                                } else {
                                                                    var v230 uint8 = 42
                                                                    _ = v230
                                                                }
                                                            }
                                                            v5 -= -9223372036854775808
                                                            if ((v82 * (v82 >> ((4294967295) & 63))) > v82) {
                                                                var v231 uint64 = 0
                                                                switch "hello" {
                                                                case "":
                                                                    v3[2] = (v4 || ((2 == 188) && (!false)))
                                                                    if v4 {
                                                                        v1 += (v1 - v1)
                                                                        v1 -= v1
                                                                        v5 = v214
                                                                    }
                                                                    if v4 {
                                                                        fmt.Println(v4, fi79, v231)
                                                                        if ((1.5 - 3.14159) == 0.0) {
                                                                            var v232 uint64 = v231
                                                                            var v233 uint32 = 178
                                                                            var v234 uint16 = uint16(v82)
                                                                            _ = v232
                                                                            _ = v233
                                                                            _ = v234
                                                                        }
                                                                        var v235 int32 = 95
                                                                        _ = v235
                                                                    }
                                                                case "foo bar":
                                                                    var v236 uint = v82
                                                                    _ = v236
                                                                default:
                                                                    v237 := GenS15{F0: ("hello" + ("" + "")), F1: v82}
                                                                    _ = v237
                                                                }
                                                                if v4 {
                                                                    if (!v4) {
                                                                        switch fi78 {
                                                                        case 1:
                                                                            if (21 > 1) {
                                                                                v6 -= v6
                                                                            } else {
                                                                                fmt.Println(fi78)
                                                                                v238 := GenS25{F0: 255}
                                                                                _ = v238
                                                                            }
                                                                        case 32768:
                                                                            if len(v215) > 0 {
                                                                                v215[0] = fi79
                                                                            }
                                                                        case -9223372036854775808:
                                                                            if len(v215) > 0 {
                                                                                v215[0] = -1
                                                                            }
                                                                        }
                                                                        v239 := GenS27{F0: 65535}
                                                                        if (v6 >= (^(v6))) {
                                                                            var v240 int32 = -1
                                                                            v214 -= v5
                                                                            if ((v6 << ((v82) & 63)) <= int16(fi78)) {
                                                                                var v241 string = "a"
                                                                                var v242 int = (fi78 ^ (fi79 % (((fi79 >> ((65535) & 63))) | 1)))
                                                                                if (true && true) {
                                                                                    fmt.Println(v5, v241, fi78)
                                                                                    v82 = v82
                                                                                    v6 -= int16(v240)
                                                                                } else {
                                                                                    var v243 uint32 = 688
                                                                                    v244 := GenS8{F0: 127, F1: (^(v240))}
                                                                                    _ = v243
                                                                                    _ = v244
                                                                                }
                                                                                _ = v241
                                                                                _ = v242
                                                                            }
                                                                            _ = v240
                                                                        } else {
                                                                            var v245 int32 = -2147483648
                                                                            if ((0 < fi79) && (!((-(v6)) <= int16(fi79)))) {
                                                                                if false {
                                                                                    var v246 uint8 = 2
                                                                                    if ((45 == 65535) || (!v4)) {
                                                                                        v247 := GenS28{F0: v4, F1: v1, F2: fi78}
                                                                                        fmt.Println(v231, v231)
                                                                                        v214 -= int64(v6)
                                                                                        _ = v247
                                                                                    }
                                                                                    v248 := GenS16{F0: (^(v246)), F1: 2}
                                                                                    _ = v246
                                                                                    _ = v248
                                                                                } else {
                                                                                    v249 := GenS23{F0: 2, F1: (v245 << ((2) & 63))}
                                                                                    _ = v249
                                                                                }
                                                                                if len(v215) >= 2 {
                                                                                    v215 = v215[1:len(v215)]
                                                                                }
                                                                            } else {
                                                                                var v250 uint8 = 2
                                                                                if v4 {
                                                                                    if (false && false) {
                                                                                        fmt.Println(fi79, v250, v231)
                                                                                        v251 := GenS17{F0: uint16(v250), F1: (v250 >> (((v82 << ((0) & 63))) & 63)), F2: fi78}
                                                                                        if v4 {
                                                                                            v252 := GenS29{F0: 0, F1: (v6 & (v6 - -1)), F2: "hello"}
                                                                                            _ = v252
                                                                                        }
                                                                                        _ = v251
                                                                                    } else {
                                                                                        var v253 uint8 = (v250 / ((v250) | 1))
                                                                                        fmt.Println(v1)
                                                                                        v231 += v231
                                                                                        v254 := GenS30{F0: 1, F1: -95, F2: (v1 ^ (v1 - (v1 / (((v1 >> ((v82) & 63))) | 1))))}
                                                                                        v255 := GenS2{F0: 1, F1: v1, F2: (^(v231))}
                                                                                        _ = v253
                                                                                        _ = v254
                                                                                        _ = v255
                                                                                    }
                                                                                    var v256 int16 = (v6 << (((^(v82))) & 63))
                                                                                    v6 = v6
                                                                                    _ = v256
                                                                                }
                                                                                _ = v250
                                                                            }
                                                                            v80 = (-(v214))
                                                                            var v257 float32 = (0.0 / 1.0)
                                                                            v82 -= v82
                                                                            _ = v245
                                                                            _ = v257
                                                                        }
                                                                        v258 := GenS14{F0: 42, F1: (v231 + 4294967295), F2: 193}
                                                                        if len(v215) > 0 {
                                                                            v215[0] = 1
                                                                        }
                                                                        _ = v239
                                                                        _ = v258
                                                                    }
                                                                    var v259 float64 = -2.25
                                                                    v1 -= v1
                                                                    if len(v215) > 0 {
                                                                        v215[0] = (fi79 >> (((^(v82))) & 63))
                                                                    }
                                                                    switch (fi78 - fi79) {
                                                                    case 2147483647:
                                                                        fmt.Println(v214)
                                                                    default:
                                                                        var v260 int = fi79
                                                                        _ = v260
                                                                    }
                                                                    _ = v259
                                                                } else {
                                                                    var v261 float32 = 0.0
                                                                    _ = v261
                                                                }
                                                                switch ("hello" + ("foo bar" + "foo bar")) {
                                                                case "foo bar":
                                                                    if false {
                                                                        if (!false) {
                                                                            v262 := GenS12{F0: 1}
                                                                            if (-1 >= (^(v80))) {
                                                                                v214 = v214
                                                                                v263 := GenS31{F0: (^(v231)), F1: v82, F2: 1}
                                                                                fmt.Println(v6, v1)
                                                                                _ = v263
                                                                            } else {
                                                                                if (!("" >= "hello")) {
                                                                                    var v264 float64 = 3.14159
                                                                                    if ("foo bar" == "hello") {
                                                                                        v215 = append(v215, fi79)
                                                                                        fmt.Println(v5)
                                                                                        var v265 uint8 = 1
                                                                                        _ = v265
                                                                                    }
                                                                                    v82 = v82
                                                                                    _ = v264
                                                                                } else {
                                                                                    fmt.Println(v231, v82)
                                                                                    var v266 uint16 = 274
                                                                                    v267 := GenS20{F0: 1}
                                                                                    _ = v266
                                                                                    _ = v267
                                                                                }
                                                                            }
                                                                            if (0 == 65535) {
                                                                                var v268 float64 = float64(0.0)
                                                                                v215 = append(v215, 9223372036854775807)
                                                                                _ = v268
                                                                            } else {
                                                                                switch (("hello" + "") + "") {
                                                                                case "hello":
                                                                                    var v269 string = "a"
                                                                                    _ = v269
                                                                                }
                                                                                var v270 int = -9223372036854775808
                                                                                var v271 float64 = 1.5
                                                                                _ = v270
                                                                                _ = v271
                                                                            }
                                                                            _ = v262
                                                                        } else {
                                                                            var v272 bool = (-9223372036854775808 >= (v5 | 42))
                                                                            if v4 {
                                                                                var v273 int64 = v5
                                                                                v1 -= v1
                                                                                var v274 int = (^(fi78))
                                                                                _ = v273
                                                                                _ = v274
                                                                            }
                                                                            var v275 uint32 = 2147483647
                                                                            _ = v272
                                                                            _ = v275
                                                                        }
                                                                    }
                                                                default:
                                                                    if false {
                                                                        v276 := v3[589]
                                                                        _ = v276
                                                                    } else {
                                                                        if (fi79 > (fi79 * fi78)) {
                                                                            v82 -= v82
                                                                            v1 -= (v1 ^ int8(fi79))
                                                                        }
                                                                        v277 := GenS21{F0: -579, F1: v4, F2: v1}
                                                                        _ = v277
                                                                    }
                                                                }
                                                                _ = v231
                                                            } else {
                                                                fmt.Println(v82, v6)
                                                                v80 = (v5 | 1)
                                                                var v278 int16 = (v6 & v6)
                                                                _ = v278
                                                            }
                                                            v279 := v3[2147483648]
                                                            _ = v215
                                                            _ = v279
                                                        }
                                                        _ = v213
                                                        _ = v214
                                                    }
                                                    v80 += (v5 & v80)
                                                }
                                                switch "" {
                                                case "\n\t":
                                                    if v4 {
                                                        var v280 int64 = v5
                                                        v1 = (-(v1))
                                                        var v281 uint16 = 65535
                                                        _ = v280
                                                        _ = v281
                                                    } else {
                                                        switch ("" + ("a" + "hello")) {
                                                        case "":
                                                            switch (fi78 * fi78) {
                                                            case 0:
                                                                if (v4 || ((v4 && false) || v4)) {
                                                                    if (2 < 38) {
                                                                        var v282 uint32 = 961
                                                                        _ = v282
                                                                    }
                                                                } else {
                                                                    v283 := []int{}
                                                                    _ = v283
                                                                }
                                                            }
                                                        }
                                                        var v284 string = "a"
                                                        v285 := GenS32{F0: -81, F1: (v1 / (((v1 ^ 0)) | 1))}
                                                        _ = v284
                                                        _ = v285
                                                    }
                                                case "hello":
                                                    v4 = ((1 <= 1) || true)
                                                default:
                                                    var v286 uint = (v82 - v82)
                                                    _ = v286
                                                }
                                            }
                                        default:
                                            v1 = int8(v82)
                                        }
                                    }
                                }
                            case 854:
                                var v287 int8 = (v1 >> ((60) & 63))
                                _ = v287
                            default:
                                var v288 int = (fi78 * (-(fi78)))
                                _ = v288
                            }
                            var v289 uint8 = 127
                            _ = v289
                        }
                    }
                    _ = fi79
                    _ = v80
                    _ = v81
                    _ = v82
                }
                _ = fi78
            }
        case "foo bar":
            v1 = (^(v1))
        default:
            for fi290 := 0; fi290 < 16; fi290++ {
                switch "foo bar" {
                case "a":
                    v291 := []int{32767, (fi290 / (((fi290 / (((fi290 >> ((711) & 63))) | 1))) | 1)), 877, fi290}
                    _ = v291
                case "\n\t":
                    v3[42] = true
                case "":
                    delete(v3, 2147483648)
                default:
                    v292 := []int{fi290}
                    _ = v292
                }
                if false {
                    v1 += (v1 ^ v1)
                    if true {
                        fmt.Println(v4, v6)
                        var v293 uint64 = 18446744073709551615
                        var v294 uint16 = 32767
                        _ = v293
                        _ = v294
                    }
                    v295 := v2[21]
                    _ = v295
                }
                _ = fi290
            }
        }
        v4 = ((!false) || false)
    } else {
        v6 = -1
        var v296 int32 = 42
        _ = v296
    }
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
