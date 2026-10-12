package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint64; F1 int16; F2 bool }
type GenS8 struct { F0 int8; F1 int64 }
type GenS9 struct { F0 int8; F1 int8; F2 uint64 }
type GenS10 struct { F0 int32; F1 int32; F2 int32 }
type GenS11 struct { F0 bool; F1 int64 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int64; F1 uint }
type GenS14 struct { F0 bool; F1 int }
type GenS15 struct { F0 int16; F1 bool; F2 string }
type GenS16 struct { F0 int; F1 uint64 }
type GenS17 struct { F0 int8; F1 int16; F2 string }
type GenS18 struct { F0 int64 }
func genFn0() int {
    if (-1 < -2147483648) {
        if true {
            if false {
                v6 := []int{32768, 0}
                var v7 string = "foo bar"
                fmt.Println(v7)
                _ = v6
                _ = v7
            } else {
                v8 := GenS1{F0: 0}
                var v9 bool = false
                if (!((1 >= -1) && v9)) {
                    for fi10 := 0; fi10 < 14; fi10++ {
                        v9 = v9
                        switch (fi10 & fi10) {
                        case 42:
                            fmt.Println(v9, v9)
                        default:
                            var v11 uint16 = 0
                            _ = v11
                        }
                        var v12 uint64 = 42
                        _ = fi10
                        _ = v12
                    }
                    if (!v9) {
                        if (1 >= 2147483647) {
                            var v13 bool = v9
                            _ = v13
                        } else {
                            v14 := []int{42, -169, 42}
                            _ = v14
                        }
                    } else {
                        v9 = ((9223372036854775807 > 9223372036854775807) && (!false))
                        fmt.Println(v9)
                    }
                } else {
                    if (v9 || (127 > 42)) {
                        var v15 uint16 = 65535
                        for fi16 := 0; fi16 < 19; fi16++ {
                            if true {
                                v9 = v9
                                var v17 int16 = 1000
                                _ = v17
                            }
                            _ = fi16
                        }
                        _ = v15
                    } else {
                        v18 := []int{128, 128, 510}
                        v18 = append(v18, -2147483648)
                        _ = v18
                    }
                }
                _ = v8
                _ = v9
            }
        }
    } else {
        for fi19 := 0; fi19 < 15; fi19++ {
            for fi20 := 0; fi20 < 13; fi20++ {
                v21 := []int{9223372036854775807}
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := []int{(fi22 & (fi22 * -688)), fi22}
            fmt.Println(fi22)
            _ = fi22
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        var v25 float32 = 1.5
        _ = fi24
        _ = v25
    }
    return -2147483648
}

func main() {
    var v1 float64 = ((0.1 / 0.1) * 1e10)
    var v2 bool = (!(-9223372036854775808 < -325))
    var v3 string = "\n\t"
    v4 := []uint32{65535, 118}
    var v5 float32 = (0.0 * 3.14159)
    switch v3 {
    case "hello":
        v26 := GenS1{F0: 1}
        _ = v26
    case "foo bar":
        for fi27 := 0; fi27 < 8; fi27++ {
            if (2 < 1) {
                switch (v3 + v3) {
                case "a":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                v28 := GenS1{F0: 18446744073709551615}
                _ = v28
            }
            _ = fi27
        }
    }
    if len(v4) > 0 {
        v4[0] = 1
    }
    if (v2 || true) {
        v29 := GenS2{F0: 42}
        v30 := GenS2{F0: 0}
        _ = v29
        _ = v30
    } else {
        v4 = append(v4, 2)
        var v31 uint64 = 18446744073709551615
        v31 += (v31 + v31)
        _ = v31
    }
    var v32 float64 = 1.5
    if len(v4) > 0 {
        v4[0] = 1
    }
    v33 := GenS3{F0: -106, F1: 1}
    if false {
        v34 := GenS4{F0: v2}
        for fi35 := 0; fi35 < 11; fi35++ {
            fmt.Println(v5, v1)
            _ = fi35
        }
        var v36 int32 = -2147483648
        _ = v34
        _ = v36
    }
    v37 := GenS4{F0: ((-1 < 0) || true)}
    var v38 int8 = 127
    if false {
        for fi39 := 0; fi39 < 7; fi39++ {
            switch v3 {
            case "hello":
                var v40 int8 = v38
                _ = v40
            case "a":
                for fi41 := 0; fi41 < 10; fi41++ {
                    v32 -= v1
                    var v42 int16 = -551
                    if (409 < 18446744073709551615) {
                        if v2 {
                            if (2147483647 >= 42) {
                                if (208 < 1) {
                                    switch (v3 + (v3 + (v3 + "foo bar"))) {
                                    case "hello":
                                        var v43 float32 = ((3.14159 - v5) - ((v5 + 1e10) - 1.5))
                                        _ = v43
                                    case "foo bar":
                                        fmt.Println(v38)
                                    case "":
                                        if v2 {
                                            v3 = ("foo bar" + "a")
                                            var v44 int64 = -1
                                            if v2 {
                                                var v45 uint64 = 42
                                                if (v2 && false) {
                                                    if len(v4) > 0 {
                                                        v4[0] = 2147483648
                                                    }
                                                    v3 = (v3 + v3)
                                                    v46 := GenS4{F0: v2}
                                                    var v47 float32 = v5
                                                    v48 := GenS3{F0: v38, F1: 2147483647}
                                                    _ = v46
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    switch (fi41 << ((2) & 63)) {
                                                    case 954:
                                                        var v49 bool = (true && v2)
                                                        _ = v49
                                                    default:
                                                        var v50 uint = 2
                                                        _ = v50
                                                    }
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                                _ = v45
                                            }
                                            _ = v44
                                        }
                                    default:
                                        var v51 float64 = float64(((0.1 + v5) * v5))
                                        var v52 uint16 = 2
                                        fmt.Println(v5)
                                        _ = v51
                                        _ = v52
                                    }
                                }
                            }
                            switch (v3 + "\n\t") {
                            case "a":
                                var v53 uint8 = 108
                                _ = v53
                            case "\n\t":
                                switch v3 {
                                case "":
                                    if (0 != 1) {
                                        v54 := GenS3{F0: v38, F1: 2147483647}
                                        v1 = v1
                                        _ = v54
                                    }
                                }
                            case "":
                                v5 = float32((v1 + v1))
                            default:
                                var v55 uint16 = 1
                                _ = v55
                            }
                            fmt.Println(fi41, fi39, fi41)
                        }
                        v56 := GenS4{F0: (-32768 >= v42)}
                        v4 = append(v4, 2)
                        _ = v56
                    }
                    v42 -= 225
                    if v2 {
                        switch v3 {
                        case "":
                            if (v2 && ((453 >= 65535) || (((1 == 2) || (9223372036854775807 == -346)) || (v42 <= v42)))) {
                                v57 := GenS4{F0: true}
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v57
                            } else {
                                var v58 int = fi39
                                v42 -= 1000
                                var v59 int32 = 42
                                var v60 int8 = v38
                                _ = v58
                                _ = v59
                                _ = v60
                            }
                        case "foo bar":
                            switch (fi41 % ((genFn0()) | 1)) {
                            case 9223372036854775807:
                                if (v2 && (2147483647 != -1)) {
                                    v61 := GenS5{F0: 2, F1: 302}
                                    var v62 uint32 = 2
                                    _ = v61
                                    _ = v62
                                } else {
                                    if false {
                                        var v63 uint8 = 0
                                        _ = v63
                                    } else {
                                        v64 := GenS6{F0: 2}
                                        v3 = "hello"
                                        v2 = (v1 >= 1.5)
                                        _ = v64
                                    }
                                    fmt.Println(v1, v5)
                                    v32 -= v32
                                    v65 := GenS1{F0: 4294967295}
                                    _ = v65
                                }
                            }
                        case "hello":
                            if v2 {
                                if (v2 && false) {
                                    var v66 int32 = -2147483648
                                    v67 := GenS7{F0: 42, F1: 673, F2: (!((1.5 * -2.25) != 1.5))}
                                    v4 = append(v4, 773)
                                    _ = v66
                                    _ = v67
                                } else {
                                    if ((!(((float64(((float32((v1 / 1e10)) / 3.14159) + v5)) / v1) + float64(v5)) <= 0.0)) || (fi39 < (fi41 & fi41))) {
                                        _ = len(v4)
                                        switch 42 {
                                        case 812:
                                            var v68 string = (v3 + v3)
                                            _ = v68
                                        case 32767:
                                            if (1.5 != v32) {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                if v2 {
                                                    _ = len(v4)
                                                } else {
                                                    switch fi39 {
                                                    case -707:
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    default:
                                                        v2 = v2
                                                    }
                                                }
                                                switch fi39 {
                                                case 2147483647:
                                                    v69 := GenS8{F0: (^(v38)), F1: int64(v42)}
                                                    _ = v69
                                                case -119:
                                                    v70 := GenS7{F0: 81, F1: 1000, F2: false}
                                                    _ = v70
                                                }
                                            }
                                        case -2147483648:
                                            if false {
                                                var v71 bool = v2
                                                _ = v71
                                            } else {
                                                fmt.Println(v1, fi39)
                                                var v72 uint32 = 2147483648
                                                v32 = (v1 * -2.25)
                                                _ = v72
                                            }
                                        default:
                                            v32 = 0.0
                                        }
                                    } else {
                                        v5 -= v5
                                    }
                                    if v2 {
                                        fmt.Println(fi41, fi39)
                                        v2 = v2
                                        if (!true) {
                                            v3 = v3
                                            if (((2 <= 1) && (2147483647 > -1)) && (!(!v2))) {
                                                var v73 uint8 = 1
                                                var v74 uint8 = (v73 ^ v73)
                                                _ = v73
                                                _ = v74
                                            } else {
                                                fmt.Println(v42, v3, v42)
                                            }
                                            var v75 int8 = (v38 << ((65535) & 63))
                                            _ = v75
                                        }
                                    } else {
                                        v32 += ((float64(1.5) / 1e10) + float64(v5))
                                        var v76 float32 = 1e10
                                        if true {
                                            var v77 uint64 = 2
                                            fmt.Println(v76)
                                            v78 := GenS2{F0: v38}
                                            _ = v77
                                            _ = v78
                                        }
                                        var v79 int64 = -903
                                        _ = v76
                                        _ = v79
                                    }
                                    switch -2147483648 {
                                    case 1:
                                        v2 = v2
                                    }
                                }
                                switch fi39 {
                                case 42:
                                    var v80 float32 = v5
                                    _ = v80
                                case 316:
                                    var v81 float64 = v32
                                    _ = v81
                                case 0:
                                    if v2 {
                                        var v82 float64 = v32
                                        switch ("hello" + (("hello" + v3) + (("hello" + ((("a" + v3) + (v3 + v3)) + (((v3 + v3) + "hello") + v3))) + ((v3 + v3) + v3)))) {
                                        case "":
                                            if v2 {
                                                var v83 int32 = 0
                                                var v84 int32 = (v83 >> ((65535) & 63))
                                                if len(v4) > 0 {
                                                    v4[0] = 217
                                                }
                                                if (fi41 > (-(fi41))) {
                                                    var v85 int = (fi39 | (fi39 << ((0) & 63)))
                                                    var v86 string = ((v3 + (v3 + v3)) + "")
                                                    _ = len(v4)
                                                    var v87 float64 = (v1 * v32)
                                                    _ = v85
                                                    _ = v86
                                                    _ = v87
                                                } else {
                                                    switch (fi41 << ((42) & 63)) {
                                                    case 1:
                                                        var v88 string = v3
                                                        _ = v88
                                                    }
                                                }
                                                _ = v83
                                                _ = v84
                                            } else {
                                                var v89 float32 = ((((0.1 * v5) / 1.0) - v5) - 0.1)
                                                if (2 < 42) {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v4 = append(v4, 765)
                                                    switch v3 {
                                                    case "foo bar":
                                                        v4 = append(v4, 830)
                                                    case "hello":
                                                        v1 += 3.14159
                                                    default:
                                                        v1 -= (1.5 * (v32 - 3.14159))
                                                    }
                                                }
                                                var v90 string = "foo bar"
                                                _ = v89
                                                _ = v90
                                            }
                                        default:
                                            if (255 < 465) {
                                                var v91 float32 = 1e10
                                                _ = v91
                                            }
                                        }
                                        v92 := GenS3{F0: v38, F1: 2147483647}
                                        _ = v82
                                        _ = v92
                                    }
                                default:
                                    if v2 {
                                        v3 = ((v3 + v3) + "hello")
                                        if (!v2) {
                                            if (v5 > 0.0) {
                                                fmt.Println(v2, v32)
                                            } else {
                                                var v93 int32 = 2147483647
                                                _ = v93
                                            }
                                            var v94 float64 = v32
                                            var v95 float64 = v94
                                            _ = v94
                                            _ = v95
                                        } else {
                                            switch v3 {
                                            case "foo bar":
                                                fmt.Println(v5)
                                            case "hello":
                                                var v96 float64 = (v1 + ((float64(1e10) * v32) - ((v32 + 1.5) - v1)))
                                                _ = v96
                                            default:
                                                var v97 uint16 = 424
                                                _ = v97
                                            }
                                            switch fi41 {
                                            case -2147483648:
                                                if (18446744073709551615 < 2) {
                                                    var v98 uint16 = 734
                                                    if ((v3 + (v3 + ((("foo bar" + v3) + "hello") + ("\n\t" + (v3 + v3))))) == (((v3 + "foo bar") + ("" + v3)) + v3)) {
                                                        v5 += v5
                                                        var v99 uint8 = 2
                                                        _ = v99
                                                    } else {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        v100 := GenS1{F0: 4294967295}
                                                        v4 = append(v4, 2)
                                                        _ = v100
                                                    }
                                                    _ = v98
                                                } else {
                                                    fmt.Println(v38, v32, v32)
                                                }
                                            }
                                        }
                                    }
                                }
                                var v101 uint = 18446744073709551615
                                _ = v101
                            } else {
                                if v2 {
                                    v3 = ((v3 + v3) + v3)
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v102 := GenS7{F0: 42, F1: -1, F2: true}
                                    _ = v102
                                }
                                switch v3 {
                                case "a":
                                    fmt.Println(v38, v1)
                                case "foo bar":
                                    switch (fi39 << ((255) & 63)) {
                                    case 4294967295:
                                        v2 = (!v2)
                                    case -9223372036854775808:
                                        var v103 float32 = float32((-2.25 + v1))
                                        _ = v103
                                    case 0:
                                        v104 := GenS9{F0: 127, F1: (v38 & (v38 + (v38 - (-(v38))))), F2: 625}
                                        _ = v104
                                    }
                                default:
                                    var v105 bool = v2
                                    _ = v105
                                }
                            }
                            v32 += ((v1 * 1.5) - v1)
                        default:
                            v106 := GenS6{F0: 42}
                            _ = v106
                        }
                        fmt.Println(fi41, v32, v3)
                    } else {
                        v107 := GenS7{F0: 42, F1: (v42 << ((815) & 63)), F2: true}
                        v3 = "a"
                        _ = v107
                    }
                    var v108 int64 = 1
                    _ = fi41
                    _ = v108
                    _ = v42
                }
            case "":
                v32 = v32
            }
            for fi109 := 0; fi109 < 4; fi109++ {
                _ = len(v4)
                _ = fi109
            }
            for fi110 := 0; fi110 < 8; fi110++ {
                switch fi110 {
                case 61:
                    var v111 int32 = 761
                    _ = v111
                case 204:
                    if true {
                        if (v38 == v38) {
                            if (!((v5 >= v5) && true)) {
                                var v112 float32 = v5
                                if (2 < 1) {
                                    v113 := GenS5{F0: 42, F1: 0}
                                    fmt.Println(v5, v3, v2)
                                    var v114 uint = 18446744073709551615
                                    if (v3 > v3) {
                                        if ((v2 && false) && (false && (255 < uint8(v114)))) {
                                            var v115 uint = v114
                                            v2 = v2
                                            v116 := GenS9{F0: v38, F1: (-(v38)), F2: 2}
                                            _ = v115
                                            _ = v116
                                        } else {
                                            var v117 string = v3
                                            _ = v117
                                        }
                                        v114 -= 788
                                    } else {
                                        if v2 {
                                            var v118 int8 = v38
                                            var v119 float32 = 0.0
                                            _ = v118
                                            _ = v119
                                        }
                                        var v120 uint16 = 91
                                        if false {
                                            v121 := GenS10{F0: -417, F1: -2147483648, F2: -1}
                                            v5 = (0.1 / v112)
                                            v2 = v2
                                            _ = v121
                                        } else {
                                            v38 = 42
                                            v122 := GenS11{F0: true, F1: 1}
                                            _ = len(v4)
                                            _ = v122
                                        }
                                        _ = v120
                                    }
                                    _ = v113
                                    _ = v114
                                } else {
                                    v1 = -2.25
                                    var v123 float64 = 1.5
                                    _ = v123
                                }
                                if (((("\n\t" + "foo bar") + (((v3 + v3) + (v3 + v3)) + (v3 + (v3 + ("foo bar" + v3))))) + v3) >= "hello") {
                                    if false {
                                        var v124 uint32 = 0
                                        v125 := GenS12{F0: (v124 << ((2) & 63))}
                                        var v126 uint64 = 42
                                        _ = v124
                                        _ = v125
                                        _ = v126
                                    }
                                    v127 := GenS11{F0: true, F1: -195}
                                    _ = v127
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if (!(!v2)) {
                                        if v2 {
                                            var v128 uint8 = 255
                                            _ = v128
                                        }
                                        v129 := GenS13{F0: -9223372036854775808, F1: 18446744073709551615}
                                        if v2 {
                                            if (!(true && (0.1 <= -2.25))) {
                                                v4 = append(v4, 42)
                                                v38 += 42
                                            }
                                            if false {
                                                if true {
                                                    switch ("\n\t" + "hello") {
                                                    case "hello":
                                                        switch fi39 {
                                                        case 4294967295:
                                                            switch "\n\t" {
                                                            case "a":
                                                                v112 = v112
                                                            default:
                                                                v3 = v3
                                                            }
                                                        }
                                                    case "":
                                                        v3 = ("a" + ((("hello" + (v3 + v3)) + v3) + v3))
                                                    case "a":
                                                        var v130 int16 = -655
                                                        if (v2 || (v2 && (147 >= 168))) {
                                                            v131 := GenS8{F0: (v38 - 25), F1: int64(v38)}
                                                            v132 := GenS1{F0: 880}
                                                            v130 += v130
                                                            _ = v131
                                                            _ = v132
                                                        }
                                                        _ = v130
                                                    default:
                                                        fmt.Println(v1)
                                                    }
                                                    v32 += 0.0
                                                }
                                                var v133 bool = (163 >= 951)
                                                var v134 int8 = v38
                                                _ = v133
                                                _ = v134
                                            } else {
                                                switch (v3 + v3) {
                                                case "foo bar":
                                                    fmt.Println(v3)
                                                default:
                                                    v1 -= (v1 + float64(v112))
                                                }
                                                var v135 int = fi39
                                                switch genFn0() {
                                                case 32767:
                                                    if v2 {
                                                        if (!((v38 << ((630) & 63)) != -128)) {
                                                            fmt.Println(fi110, v2, v135)
                                                            var v136 int8 = v38
                                                            _ = v136
                                                        }
                                                    } else {
                                                        var v137 int8 = v38
                                                        if v2 {
                                                            var v138 uint32 = 117
                                                            var v139 bool = (genFn0() < (fi110 >> ((4294967295) & 63)))
                                                            if (v2 && v139) {
                                                                v140 := GenS14{F0: (255 > 255), F1: (fi39 % ((32767) | 1))}
                                                                fmt.Println(v5, v135, fi110)
                                                                switch "a" {
                                                                case "foo bar":
                                                                    if v139 {
                                                                        if (128 <= 127) {
                                                                            var v141 int64 = -709
                                                                            _ = v141
                                                                        }
                                                                        if v2 {
                                                                            fmt.Println(v32)
                                                                            v142 := GenS9{F0: (v137 << ((974) & 63)), F1: (v38 | -1), F2: 857}
                                                                            fmt.Println(v112, v5, v5)
                                                                            _ = v142
                                                                        }
                                                                        if (477 == 661) {
                                                                            switch (fi110 / ((-304) | 1)) {
                                                                            case -1:
                                                                                fmt.Println(v135)
                                                                            case 1:
                                                                                switch v3 {
                                                                                case "hello":
                                                                                    if (18446744073709551615 == 1) {
                                                                                        var v143 int = v135
                                                                                        _ = v143
                                                                                    } else {
                                                                                        var v144 bool = v2
                                                                                        _ = v144
                                                                                    }
                                                                                default:
                                                                                    var v145 float32 = (((1.5 / float32(v1)) * ((((v112 * (0.0 / v5)) / 1e10) * 0.0) - v5)) + (v5 - float32(((v32 * v32) - v32))))
                                                                                    _ = v145
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v2 = v139
                                                                            v135 = (^(fi39))
                                                                        }
                                                                        v146 := GenS15{F0: 1, F1: ((v112 > float32((1.5 * v32))) && (1 <= 1)), F2: ("\n\t" + "")}
                                                                        _ = v146
                                                                    }
                                                                case "\n\t":
                                                                    switch v135 {
                                                                    case -1:
                                                                        var v147 int16 = -683
                                                                        _ = v147
                                                                    default:
                                                                        fmt.Println(v5, fi39)
                                                                    }
                                                                default:
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                }
                                                                _ = v140
                                                            }
                                                            switch int(v138) {
                                                            case 2147483647:
                                                                v148 := GenS5{F0: 900, F1: 2}
                                                                _ = v148
                                                            case 9223372036854775807:
                                                                fmt.Println(v2, v1, v135)
                                                            case 1:
                                                                v4 = append(v4, (v138 & 0))
                                                            default:
                                                                var v149 bool = v139
                                                                _ = v149
                                                            }
                                                            v32 += 3.14159
                                                            _ = v138
                                                            _ = v139
                                                        }
                                                        v2 = v2
                                                        v112 -= v5
                                                        v150 := GenS13{F0: -253, F1: 4294967295}
                                                        _ = v137
                                                        _ = v150
                                                    }
                                                case 32768:
                                                    var v151 uint64 = 1
                                                    _ = v151
                                                }
                                                var v152 float64 = v32
                                                _ = v135
                                                _ = v152
                                            }
                                            if true {
                                                var v153 uint32 = 65535
                                                switch ("\n\t" + "") {
                                                case "\n\t":
                                                    v38 -= (-(v38))
                                                case "hello":
                                                    v154 := GenS13{F0: -577, F1: 42}
                                                    _ = v154
                                                }
                                                v155 := GenS7{F0: 1, F1: -786, F2: (!false)}
                                                _ = v153
                                                _ = v155
                                            }
                                        }
                                        var v156 float32 = v5
                                        if ((42 >= 910) && (!v2)) {
                                            var v157 int16 = -32768
                                            v156 = 3.14159
                                            var v158 string = "a"
                                            _ = v157
                                            _ = v158
                                        }
                                        _ = v129
                                        _ = v156
                                    } else {
                                        var v159 float64 = 1.5
                                        v159 += ((v159 / (1.5 + (3.14159 / v159))) / v1)
                                        _ = v159
                                    }
                                }
                                _ = v112
                            }
                            v32 -= (v1 * v1)
                        }
                        if v2 {
                            var v160 uint64 = 18446744073709551615
                            v160 = (^(v160))
                            if v2 {
                                v2 = v2
                            } else {
                                v161 := GenS5{F0: 255, F1: 255}
                                v160 = v160
                                var v162 int16 = 32767
                                _ = v161
                                _ = v162
                            }
                            _ = v160
                        } else {
                            v38 = v38
                            v163 := GenS15{F0: 32767, F1: v2, F2: v3}
                            _ = v163
                        }
                        fmt.Println(v5, v1, v5)
                    } else {
                        var v164 int32 = 42
                        v165 := GenS14{F0: v2, F1: fi110}
                        _ = v164
                        _ = v165
                    }
                    if (!v2) {
                        switch "hello" {
                        case "a":
                            v166 := GenS16{F0: -1, F1: 0}
                            _ = v166
                        case "\n\t":
                            var v167 string = "\n\t"
                            _ = v167
                        }
                        switch -1 {
                        case 32768:
                            var v168 bool = (-280 > -9223372036854775808)
                            _ = v168
                        case 42:
                            v3 = v3
                        }
                        v3 = (v3 + v3)
                    }
                case 32767:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                switch fi110 {
                case 127:
                    v4 = append(v4, 0)
                case 42:
                    v4 = append(v4, 42)
                }
                _ = fi110
            }
            _ = fi39
        }
        for fi169 := 0; fi169 < 13; fi169++ {
            v2 = true
            v170 := GenS13{F0: 0, F1: 255}
            var v171 int16 = 1000
            _ = fi169
            _ = v170
            _ = v171
        }
    } else {
        for fi172 := 0; fi172 < 12; fi172++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v173 int16 = -32768
            _ = fi172
            _ = v173
        }
        switch 4294967295 {
        case 127:
            v38 -= (v38 << ((1) & 63))
        case -2147483648:
            v174 := GenS1{F0: 255}
            _ = v174
        case 128:
            v175 := GenS17{F0: 127, F1: 32767, F2: "a"}
            _ = v175
        default:
            var v176 bool = v2
            _ = v176
        }
    }
    if (v5 >= float32(v32)) {
        switch -38 {
        case 32767:
            switch genFn0() {
            case 4294967295:
                v38 = 84
            }
        default:
            v38 -= (v38 * 42)
        }
        if true {
            for fi177 := 0; fi177 < 20; fi177++ {
                var v178 int32 = -1
                switch (-(fi177)) {
                case -234:
                    if (("a" < v3) && v2) {
                        var v179 int16 = 1
                        _ = v179
                    } else {
                        if false {
                            v2 = v2
                            var v180 int16 = int16(v38)
                            _ = v180
                        } else {
                            var v181 float64 = 1e10
                            _ = v181
                        }
                    }
                case 9223372036854775807:
                    if false {
                        if (42 != 42) {
                            if len(v4) > 0 {
                                v4[0] = 1
                            }
                        }
                        v1 -= (v32 + (((1.5 + v32) / float64(v5)) - ((v32 / (v1 + 3.14159)) + 3.14159)))
                        switch v3 {
                        case "hello":
                            var v182 int = 82
                            _ = v182
                        case "\n\t":
                            var v183 uint32 = 2147483648
                            _ = v183
                        case "a":
                            v2 = false
                        default:
                            v32 = float64(v5)
                        }
                    }
                case 32767:
                    for fi184 := 0; fi184 < 9; fi184++ {
                        switch "" {
                        case "\n\t":
                            var v185 string = (v3 + ((v3 + v3) + v3))
                            _ = v185
                        case "foo bar":
                            if ((!(127 >= 0)) || v2) {
                                var v186 uint = 961
                                if v2 {
                                    fmt.Println(v38, v5)
                                    v3 = v3
                                    var v187 int16 = 42
                                    v2 = false
                                    fmt.Println(fi177, v32, v3)
                                    fmt.Println(v186)
                                    _ = v187
                                }
                                _ = v186
                            } else {
                                v3 = v3
                                v188 := GenS15{F0: 42, F1: (0 != 1), F2: ((v3 + v3) + "\n\t")}
                                fmt.Println(v2, v178, v3)
                                _ = v188
                            }
                        case "":
                            v5 = v5
                        }
                        v189 := GenS8{F0: v38, F1: -566}
                        _ = fi184
                        _ = v189
                    }
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi177
                _ = v178
            }
        } else {
            for fi190 := 0; fi190 < 4; fi190++ {
                if ((v38 % ((v38) | 1)) == (-(v38))) {
                    v32 += v1
                }
                if len(v4) > 0 {
                    v4[0] = 588
                }
                v38 = v38
                _ = fi190
            }
            fmt.Println(v32, v2, v2)
            fmt.Println(v3, v1)
        }
        v191 := GenS4{F0: v2}
        _ = v191
    } else {
        for fi192 := 0; fi192 < 7; fi192++ {
            v193 := GenS18{F0: 9223372036854775807}
            v4 = append(v4, 42)
            _ = fi192
            _ = v193
        }
        fmt.Println(v3, v5)
    }
    var v194 uint = 65535
    var v195 float64 = (v32 / 0.1)
    if len(v4) > 0 {
        v4[0] = 2
    }
    if (2147483647 == 187) {
        switch "foo bar" {
        case "hello":
            var v196 int64 = -1
            _ = v196
        case "a":
            if v2 {
                var v197 bool = (148 >= 1)
                v198 := GenS2{F0: v38}
                _ = v197
                _ = v198
            } else {
                var v199 uint8 = 48
                _ = v199
            }
        case "foo bar":
            v200 := GenS13{F0: -537, F1: (v194 << (((^(v194))) & 63))}
            _ = v200
        default:
            var v201 uint32 = 0
            _ = v201
        }
        if (!true) {
            v4 = append(v4, 2147483647)
            if (v38 > -1) {
                if v2 {
                    var v202 uint16 = 32767
                    _ = v202
                }
                fmt.Println(v3)
                v5 -= v5
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v194:", v194)
    fmt.Println("v195:", v195)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v194
    _ = v195
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
