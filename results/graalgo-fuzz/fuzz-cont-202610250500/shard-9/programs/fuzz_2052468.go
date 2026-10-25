package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 int8; F1 string; F2 int32 }
type GenS5 struct { F0 bool; F1 int8; F2 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint; F1 uint }
type GenS8 struct { F0 int; F1 int64 }
type GenS9 struct { F0 uint8; F1 int64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 bool; F1 int; F2 uint64 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint32; F1 uint16 }
type GenS14 struct { F0 int; F1 string; F2 int16 }
type GenS15 struct { F0 uint16; F1 string; F2 uint }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int; F1 uint8 }
type GenS18 struct { F0 int64; F1 uint16 }
func genFn0(p0 int16) int8 {
    for fi6 := 0; fi6 < 3; fi6++ {
        if true {
            for fi7 := 0; fi7 < 16; fi7++ {
                var v8 string = "hello"
                v9 := []int{540, 42, (fi7 + fi7)}
                _ = fi7
                _ = v8
                _ = v9
            }
        } else {
            switch fi6 {
            case 0:
                var v10 uint = 1
                _ = v10
            case 9223372036854775807:
                v11 := []int{fi6, (-(fi6))}
                _ = v11
            }
            p0 -= p0
            p0 += 42
        }
        v12 := []int{}
        _ = fi6
        _ = v12
    }
    p0 = p0
    v13 := []int{0, 1, -2147483648, 128}
    v14 := GenS1{F0: 146}
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 int16 = 32767
            if true {
                p0 -= -1
            } else {
                v17 -= v17
                switch 484 {
                case 0:
                    if (!false) {
                        v18 := GenS1{F0: 128}
                        var v19 uint16 = 32767
                        _ = v18
                        _ = v19
                    } else {
                        v20 := GenS2{F0: v17, F1: 709, F2: 0}
                        _ = v20
                    }
                case 32767:
                    if true {
                        switch (fi15 + 245) {
                        case 42:
                            fmt.Println(fi15, fi16, v17)
                        default:
                            var v21 uint64 = 0
                            _ = v21
                        }
                        var v22 uint64 = 1
                        _ = len(v13)
                        _ = v22
                    }
                case -587:
                    v23 := GenS3{F0: (fi15 / ((fi15) | 1)), F1: 533}
                    v24 := GenS3{F0: fi15, F1: 786}
                    v13 = append(v13, fi15)
                    var v25 uint8 = 165
                    _ = v23
                    _ = v24
                    _ = v25
                }
            }
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    _ = p0
    _ = v13
    _ = v14
    return -60
}
func genFn1(p0 uint16, p1 uint) string {
    v26 := []int{1, 32767}
    for fi27 := 0; fi27 < 15; fi27++ {
        if len(v26) > 0 {
            v26[0] = (fi27 >> (((^(p1))) & 63))
        }
        p1 = 1
        var v28 string = "hello"
        _ = fi27
        _ = v28
    }
    _ = p0
    _ = p1
    _ = v26
    return (("a" + ("hello" + "foo bar")) + "foo bar")
}

func main() {
    var v1 int32 = 1
    var v2 float32 = 0.0
    var v3 uint32 = 42
    v4 := make(map[uint32]bool)
    var v5 uint64 = 18446744073709551615
    var v29 uint8 = 255
    v1 -= (-(v1))
    if false {
        if true {
            v3 -= (^(v3))
            v30 := []int{214}
            _ = v30
        } else {
            _ = len(v4)
        }
        switch -488 {
        case 4294967295:
            if (!(!true)) {
                fmt.Println(v3, v5)
                switch 4294967295 {
                case -1:
                    for fi31 := 0; fi31 < 17; fi31++ {
                        v5 = v5
                        v2 += v2
                        _ = fi31
                    }
                default:
                    for fi32 := 0; fi32 < 8; fi32++ {
                        v3 += (v3 / ((v3) | 1))
                        if false {
                            for fi33 := 0; fi33 < 2; fi33++ {
                                fmt.Println(fi32, fi32)
                                v34 := []int{}
                                _ = fi33
                                _ = v34
                            }
                        }
                        var v35 uint8 = v29
                        _ = fi32
                        _ = v35
                    }
                }
            }
        default:
            if true {
                v29 = v29
                if (604 != 524) {
                    v36 := GenS4{F0: 42, F1: "", F2: (v1 >> ((uint(v29)) & 63))}
                    _ = v36
                }
            } else {
                fmt.Println(v2, v2, v3)
                if ((v3 | (v3 % ((v3) | 1))) <= v3) {
                    var v37 uint8 = v29
                    fmt.Println(v37, v29)
                    _ = v37
                } else {
                    v29 = uint8(v5)
                    if false {
                        v5 = v5
                        v5 = (^(v5))
                        v38 := GenS5{F0: (!false), F1: genFn0(0), F2: 65535}
                        _ = v38
                    } else {
                        for fi39 := 0; fi39 < 6; fi39++ {
                            fmt.Println(v1, v3, v5)
                            _ = len(v4)
                            _ = fi39
                        }
                        v40 := []int{834, 32768, 81}
                        _ = v40
                    }
                }
            }
        }
        if ((v2 / float32(1.5)) == (v2 * 0.1)) {
            if false {
                for fi41 := 0; fi41 < 16; fi41++ {
                    if false {
                        var v42 uint8 = v29
                        v3 = v3
                        _ = v42
                    }
                    fmt.Println(v3)
                    _ = fi41
                }
            } else {
                switch 128 {
                case 225:
                    v2 -= float32(float64((float32(0.1) + v2)))
                case 415:
                    if true {
                        fmt.Println(v2, v3, v1)
                        v29 += uint8(v5)
                    } else {
                        v29 = v29
                        fmt.Println(v5)
                        var v43 uint8 = 119
                        _ = v43
                    }
                }
                var v44 string = "foo bar"
                fmt.Println(v1, v3, v44)
                _ = v44
            }
            for fi45 := 0; fi45 < 12; fi45++ {
                _ = len(v4)
                _ = fi45
            }
            v2 = float32((0.0 * 0.0))
        } else {
            v46 := GenS1{F0: (v29 & (v29 ^ 127))}
            _ = v46
        }
    } else {
        for fi47 := 0; fi47 < 13; fi47++ {
            v29 -= (v29 % (((v29 * v29)) | 1))
            switch (fi47 & -9223372036854775808) {
            case 0:
                if false {
                    for fi48 := 0; fi48 < 15; fi48++ {
                        var v49 int8 = 1
                        v50 := v4[v3]
                        v51 := GenS6{F0: v5}
                        v52 := GenS7{F0: 584, F1: 18446744073709551615}
                        v53 := GenS6{F0: (v5 - (v5 >> ((311) & 63)))}
                        _ = fi48
                        _ = v49
                        _ = v50
                        _ = v51
                        _ = v52
                        _ = v53
                    }
                    var v54 int = (fi47 + fi47)
                    _ = v54
                } else {
                    if false {
                        switch "\n\t" {
                        case "hello":
                            switch fi47 {
                            case 967:
                                if ((!false) && (true || true)) {
                                    var v55 int64 = -1
                                    _ = v55
                                } else {
                                    for fi56 := 0; fi56 < 19; fi56++ {
                                        var v57 uint = 255
                                        if (0 != 1) {
                                            switch int(v5) {
                                            case 533:
                                                v57 = v57
                                            case 527:
                                                fmt.Println(fi56, v1, v57)
                                            case -1:
                                                v1 = v1
                                                v58 := GenS5{F0: (false && ((1000 < -1) || true)), F1: 42, F2: v57}
                                                if ("foo bar" < "") {
                                                    var v59 string = ""
                                                    _ = v59
                                                } else {
                                                    if true {
                                                        v60 := GenS5{F0: (true || (!((true && (-632 < -1)) && (v2 != -2.25)))), F1: -1, F2: 4294967295}
                                                        v3 = v3
                                                        _ = v60
                                                    }
                                                }
                                                if false {
                                                    var v61 uint64 = uint64(v3)
                                                    _ = v61
                                                }
                                                v29 = v29
                                                var v62 uint64 = v5
                                                var v63 int = fi56
                                                var v64 uint32 = (v3 >> ((2) & 63))
                                                v63 = fi47
                                                var v65 int8 = genFn0(42)
                                                _ = v58
                                                _ = v62
                                                _ = v63
                                                _ = v64
                                                _ = v65
                                            }
                                            v29 += v29
                                            var v66 float32 = (v2 * float32(1e10))
                                            _ = v66
                                        } else {
                                            var v67 uint64 = (^(v5))
                                            var v68 uint = v57
                                            v57 = 2
                                            var v69 uint64 = (^(v67))
                                            _ = v67
                                            _ = v68
                                            _ = v69
                                        }
                                        var v70 uint16 = 549
                                        _ = fi56
                                        _ = v57
                                        _ = v70
                                    }
                                    var v71 bool = true
                                    _ = v71
                                }
                            case -272:
                                switch -9223372036854775808 {
                                case 32768:
                                    var v72 uint = 933
                                    _ = v72
                                case 127:
                                    v29 = (v29 | v29)
                                case 42:
                                    v1 += -1
                                default:
                                    var v73 float64 = (-2.25 - (float64(v2) / 1.0))
                                    _ = v73
                                }
                            case 9223372036854775807:
                                v74 := GenS5{F0: (fi47 >= (fi47 >> ((18446744073709551615) & 63))), F1: -1, F2: 42}
                                _ = v74
                            }
                        default:
                            v29 = (v29 * (v29 & v29))
                        }
                    }
                    var v75 float32 = (((-2.25 / 3.14159) - (v2 / ((3.14159 * v2) + (((v2 * (float32(3.14159) - 1e10)) / v2) + ((1.5 - (v2 + (float32(1.5) * (v2 - float32(((0.1 + -2.25) - (0.1 - ((0.1 * 1.5) / float64((v2 * 3.14159)))))))))) - (v2 / v2)))))) + v2)
                    switch 2147483647 {
                    case -2147483648:
                        fmt.Println(v2, v2, fi47)
                    }
                    _ = v75
                }
            case -1:
                for fi76 := 0; fi76 < 14; fi76++ {
                    var v77 int = (fi47 >> ((1) & 63))
                    switch (fi76 * -406) {
                    case 32768:
                        if true {
                            v78 := GenS8{F0: -793, F1: 42}
                            _ = v78
                        } else {
                            fmt.Println(v2, fi47)
                            var v79 string = ""
                            var v80 uint = 65535
                            _ = v79
                            _ = v80
                        }
                    case 2147483647:
                        if ((0.1 - float64(v2)) > 0.0) {
                            v81 := GenS9{F0: v29, F1: -652}
                            switch -807 {
                            case 481:
                                var v82 float32 = v2
                                _ = v82
                            case 127:
                                switch "" {
                                case "":
                                    switch 4294967295 {
                                    case 127:
                                        v29 -= uint8(v3)
                                        switch (v77 & fi47) {
                                        case -288:
                                            var v83 uint = 65535
                                            _ = v83
                                        case 32767:
                                            v2 -= 0.1
                                        case 0:
                                            var v84 int8 = -85
                                            _ = v84
                                        default:
                                            v2 -= v2
                                        }
                                    default:
                                        switch fi76 {
                                        case 32767:
                                            switch v77 {
                                            case 402:
                                                v29 -= v29
                                            default:
                                                v1 += v1
                                            }
                                        case 9223372036854775807:
                                            switch -668 {
                                            case -2147483648:
                                                fmt.Println(v2)
                                            case 9223372036854775807:
                                                v2 = 1e10
                                            }
                                        case 0:
                                            switch (-(v77)) {
                                            case 42:
                                                delete(v4, uint32(v29))
                                            case 796:
                                                switch "a" {
                                                case "\n\t":
                                                    v85 := GenS4{F0: 100, F1: "", F2: int32(v5)}
                                                    _ = v85
                                                case "a":
                                                    switch -639 {
                                                    case -2147483648:
                                                        var v86 float32 = 3.14159
                                                        var v87 int8 = 42
                                                        var v88 uint = 592
                                                        _ = v86
                                                        _ = v87
                                                        _ = v88
                                                    case -47:
                                                        v77 -= (v77 >> ((2) & 63))
                                                        v5 += (v5 | (v5 - (v5 / ((v5) | 1))))
                                                        v89 := GenS10{F0: (^(v77))}
                                                        fmt.Println(fi76, fi76)
                                                        _ = v89
                                                    case 0:
                                                        var v90 float64 = float64((v2 / 0.1))
                                                        _ = v90
                                                    default:
                                                        v91 := []int{fi47, 168, 9223372036854775807, 2147483647}
                                                        if true {
                                                            fmt.Println(v1)
                                                            if false {
                                                                v2 -= ((v2 / v2) + 0.0)
                                                                var v92 uint8 = (v29 - (^(v29)))
                                                                _ = len(v91)
                                                                _ = v92
                                                            } else {
                                                                var v93 int16 = 32767
                                                                var v94 uint = 909
                                                                if false {
                                                                    if (-1 >= -9223372036854775808) {
                                                                        v1 += (^(v1))
                                                                        v91 = append(v91, (fi47 >> ((v94) & 63)))
                                                                        if ((v1 << (((v94 >> ((uint(v3)) & 63))) & 63)) <= (-(v1))) {
                                                                            switch fi76 {
                                                                            case 127:
                                                                                v95 := GenS4{F0: 42, F1: "foo bar", F2: -936}
                                                                                _ = v95
                                                                            default:
                                                                                v29 += v29
                                                                                var v96 bool = (true || false)
                                                                                var v97 bool = ((v1 % ((v1) | 1)) != -532)
                                                                                _ = v96
                                                                                _ = v97
                                                                            }
                                                                            if true {
                                                                                fmt.Println(v2)
                                                                                var v98 uint32 = uint32(v5)
                                                                                _ = v98
                                                                            }
                                                                            if (540 == 42) {
                                                                                v77 += (^(v77))
                                                                                v99 := GenS11{F0: true, F1: (fi76 >> ((v94) & 63)), F2: (^(v5))}
                                                                                _ = v99
                                                                            } else {
                                                                                fmt.Println(fi47, fi76, v77)
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v100 uint8 = v29
                                                                        v100 += v29
                                                                        _ = v100
                                                                    }
                                                                    if (!(false || ((^(v77)) <= fi76))) {
                                                                        if ((v94 + (v94 / ((v94) | 1))) <= (v94 * (v94 / ((v94) | 1)))) {
                                                                            v77 -= fi76
                                                                        } else {
                                                                            fmt.Println(v1)
                                                                        }
                                                                    }
                                                                    if true {
                                                                        v93 -= (v93 >> ((v94) & 63))
                                                                        v2 = float32(1e10)
                                                                        if (((v1 - (v1 + (v1 ^ int32(v5)))) > (v1 - 1)) || (-1 != 100)) {
                                                                            var v101 int16 = -32768
                                                                            v1 -= -551
                                                                            switch (fi47 % ((int(v29)) | 1)) {
                                                                            case 42:
                                                                                if ("a" < (("a" + "") + "foo bar")) {
                                                                                    switch (fi76 << ((v94) & 63)) {
                                                                                    case 256:
                                                                                        v4[v3] = (!false)
                                                                                    case 9223372036854775807:
                                                                                        var v102 uint16 = 851
                                                                                        _ = v102
                                                                                    }
                                                                                }
                                                                            case -1:
                                                                                if (42 < -9223372036854775808) {
                                                                                    if (v93 == v93) {
                                                                                        var v103 float32 = v2
                                                                                        switch "a" {
                                                                                        case "hello":
                                                                                            var v104 string = "a"
                                                                                            _ = v104
                                                                                        case "":
                                                                                            v105 := GenS12{F0: v77}
                                                                                            _ = v105
                                                                                        case "a":
                                                                                            var v106 int16 = (v101 & -32768)
                                                                                            _ = v106
                                                                                        default:
                                                                                            var v107 int64 = 1
                                                                                            _ = v107
                                                                                        }
                                                                                        _ = v103
                                                                                    } else {
                                                                                        v3 = v3
                                                                                    }
                                                                                    v3 = (v3 * v3)
                                                                                    var v108 uint8 = uint8(v3)
                                                                                    _ = v108
                                                                                }
                                                                            case 44:
                                                                                v109 := GenS8{F0: 127, F1: 0}
                                                                                _ = v109
                                                                            }
                                                                            if (v1 != (v1 & 42)) {
                                                                                if (-1 > 42) {
                                                                                    if ((!((^(v5)) < v5)) && true) {
                                                                                        _ = len(v91)
                                                                                    }
                                                                                    var v110 uint64 = v5
                                                                                    if len(v91) > 0 {
                                                                                        v91[0] = v77
                                                                                    }
                                                                                    var v111 int = fi47
                                                                                    _ = v110
                                                                                    _ = v111
                                                                                }
                                                                                delete(v4, v3)
                                                                                switch (fi76 - 4294967295) {
                                                                                case -2147483648:
                                                                                    var v112 uint64 = 164
                                                                                    _ = v112
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v77, v29, v29)
                                                                                var v113 uint64 = (v5 + v5)
                                                                                _ = v113
                                                                            }
                                                                            v77 = fi47
                                                                            var v114 uint16 = 876
                                                                            var v115 float32 = (v2 + (1.5 + v2))
                                                                            v116 := GenS10{F0: fi47}
                                                                            _ = v101
                                                                            _ = v114
                                                                            _ = v115
                                                                            _ = v116
                                                                        } else {
                                                                            if true {
                                                                                var v117 int = 407
                                                                                var v118 int16 = -1
                                                                                _ = v117
                                                                                _ = v118
                                                                            } else {
                                                                                if (42 >= 42) {
                                                                                    v94 += uint(v3)
                                                                                    v119 := GenS4{F0: 42, F1: ("\n\t" + (("\n\t" + "\n\t") + "\n\t")), F2: int32(v29)}
                                                                                    var v120 uint8 = 127
                                                                                    _ = v119
                                                                                    _ = v120
                                                                                }
                                                                                fmt.Println(v3, v1)
                                                                                var v121 int = (fi47 / (((fi47 / ((-9223372036854775808) | 1))) | 1))
                                                                                _ = v121
                                                                            }
                                                                            if (!((v1 >> ((42) & 63)) <= v1)) {
                                                                                _ = len(v4)
                                                                                v93 += v93
                                                                            } else {
                                                                                switch fi76 {
                                                                                case -1:
                                                                                    v122 := GenS13{F0: 7, F1: 2}
                                                                                    _ = v122
                                                                                default:
                                                                                    v93 -= 1
                                                                                    _ = len(v91)
                                                                                    fmt.Println(v94, v93, v1)
                                                                                }
                                                                                v123 := GenS7{F0: v94, F1: uint(v29)}
                                                                                fmt.Println(v2, v5, v1)
                                                                                _ = v123
                                                                            }
                                                                        }
                                                                        v124 := GenS3{F0: (^(fi47)), F1: (v3 << (((v94 / (((v94 * 1)) | 1))) & 63))}
                                                                        var v125 float32 = -2.25
                                                                        switch (fi76 * fi76) {
                                                                        case -2147483648:
                                                                            if (1 == (^(v5))) {
                                                                                fmt.Println(v125)
                                                                                delete(v4, 486)
                                                                            }
                                                                        case 343:
                                                                            var v126 int32 = v1
                                                                            _ = v126
                                                                        default:
                                                                            fmt.Println(v3, v3, v5)
                                                                        }
                                                                        v1 += v1
                                                                        _ = v124
                                                                        _ = v125
                                                                    }
                                                                } else {
                                                                    var v127 uint32 = (^(v3))
                                                                    var v128 int32 = (v1 & (v1 - (-(v1))))
                                                                    v3 -= v127
                                                                    _ = v127
                                                                    _ = v128
                                                                }
                                                                _ = v93
                                                                _ = v94
                                                            }
                                                            var v129 uint8 = v29
                                                            _ = v129
                                                        } else {
                                                            v4[(v3 % (((v3 << ((2) & 63))) | 1))] = (true || (((("" + "") + "hello") + "foo bar") == ("" + "foo bar")))
                                                            v130 := GenS7{F0: 65535, F1: 709}
                                                            _ = v130
                                                        }
                                                        v4[uint32(v5)] = false
                                                        _ = v91
                                                    }
                                                case "":
                                                    var v131 float64 = 0.1
                                                    _ = v131
                                                default:
                                                    v132 := GenS14{F0: -2147483648, F1: "", F2: 32767}
                                                    _ = v132
                                                }
                                            case -1:
                                                _ = len(v4)
                                            }
                                        }
                                    }
                                }
                            }
                            _ = v81
                        }
                    }
                    v29 = (v29 * v29)
                    _ = fi76
                    _ = v77
                }
            default:
                if ((v5 ^ v5) <= v5) {
                    for fi133 := 0; fi133 < 14; fi133++ {
                        delete(v4, (^(v3)))
                        v134 := []int{-664, 2147483647}
                        fmt.Println(v1, fi133)
                        _ = fi133
                        _ = v134
                    }
                } else {
                    switch "hello" {
                    case "a":
                        for fi135 := 0; fi135 < 18; fi135++ {
                            v2 -= (v2 * (float32(((((0.1 / (1e10 + (1e10 + -2.25))) / (-2.25 + 0.0)) * 0.0) * 3.14159)) * ((float32((1.5 + 3.14159)) * v2) / float32((-2.25 * 1e10)))))
                            _ = fi135
                        }
                    case "foo bar":
                        v136 := GenS15{F0: 255, F1: (((("a" + "foo bar") + "\n\t") + "") + ("" + "a")), F2: 255}
                        _ = v136
                    case "hello":
                        if ((fi47 | fi47) >= 543) {
                            var v137 bool = (v29 <= (v29 / ((42) | 1)))
                            _ = v137
                        } else {
                            var v138 uint64 = 42
                            _ = v138
                        }
                    }
                    v3 = 65535
                }
            }
            switch (fi47 ^ (fi47 >> ((255) & 63))) {
            case -38:
                if (1 == 1000) {
                    switch "a" {
                    case "foo bar":
                        v5 = (v5 - (^(v5)))
                    case "":
                        v2 = v2
                    case "a":
                        var v139 uint32 = 2147483648
                        _ = v139
                    }
                    v140 := []int{(fi47 % ((fi47) | 1)), fi47, (fi47 + (^(fi47)))}
                    v1 = (v1 / ((v1) | 1))
                    _ = v140
                } else {
                    for fi141 := 0; fi141 < 5; fi141++ {
                        var v142 uint64 = (v5 << ((12) & 63))
                        v5 -= (v5 & (v142 % (((^(v142))) | 1)))
                        _ = fi141
                        _ = v142
                    }
                    var v143 bool = true
                    for fi144 := 0; fi144 < 5; fi144++ {
                        fmt.Println(fi144)
                        var v145 int8 = 100
                        _ = fi144
                        _ = v145
                    }
                    _ = v143
                }
            case 128:
                v146 := GenS16{F0: ((!(((true || false) || false) && (!true))) || (!((^(v3)) >= v3)))}
                _ = v146
            default:
                for fi147 := 0; fi147 < 15; fi147++ {
                    var v148 int64 = 42
                    switch fi47 {
                    case 0:
                        v5 = v5
                    case 32768:
                        fmt.Println(fi147, v2)
                    case -1:
                        var v149 string = "hello"
                        _ = v149
                    }
                    if (1e10 <= (float64(v2) - float64(v2))) {
                        if (!true) {
                            v1 += -1
                            v150 := GenS5{F0: false, F1: genFn0(991), F2: 0}
                            v2 += ((v2 / (3.14159 / (v2 / v2))) + float32(0.1))
                            v1 += v1
                            fmt.Println(v3)
                            v151 := v4[42]
                            var v152 int16 = int16(v3)
                            v153 := GenS12{F0: (fi47 / (((fi147 % (((fi147 & fi47)) | 1))) | 1))}
                            if v151 {
                                v154 := GenS10{F0: -246}
                                _ = v154
                            }
                            v5 -= v5
                            _ = v150
                            _ = v151
                            _ = v152
                            _ = v153
                        } else {
                            _ = len(v4)
                            v3 += v3
                            v155 := GenS17{F0: fi147, F1: v29}
                            _ = v155
                        }
                        if false {
                            fmt.Println(v3, v29, fi47)
                        } else {
                            if ((!true) && true) {
                                var v156 uint16 = 32767
                                _ = v156
                            }
                            _ = len(v4)
                        }
                        var v157 string = ""
                        _ = v157
                    } else {
                        if false {
                            fmt.Println(v2, v29, v29)
                        } else {
                            fmt.Println(v29, v2)
                            var v158 uint = 0
                            var v159 uint16 = 1
                            _ = v158
                            _ = v159
                        }
                        var v160 uint32 = (v3 >> ((18446744073709551615) & 63))
                        v2 += (3.14159 + (1.5 * v2))
                        v29 = (v29 & v29)
                        v161 := []int{(fi47 & fi147), fi47}
                        _ = v160
                        _ = v161
                    }
                    _ = fi147
                    _ = v148
                }
            }
            _ = fi47
        }
        v2 += 3.14159
    }
    v162 := GenS10{F0: -1}
    fmt.Println(v1)
    v163 := []int{-158, 42}
    var v164 uint32 = v3
    v5 += v5
    for fi165 := 0; fi165 < 3; fi165++ {
        v3 += 1
        _ = fi165
    }
    v166 := GenS8{F0: 1, F1: 159}
    v2 = float32(((-2.25 / (-2.25 / 1e10)) - (-2.25 * 0.1)))
    switch (("" + "hello") + ((((("\n\t" + "\n\t") + "\n\t") + "") + "\n\t") + "")) {
    case "a":
        delete(v4, v164)
    case "\n\t":
        var v167 uint = 632
        _ = v167
    case "hello":
        v168 := GenS7{F0: 0, F1: 119}
        _ = v168
    default:
        v3 += v3
    }
    if false {
        var v169 string = "hello"
        var v170 uint32 = (^(v164))
        for fi171 := 0; fi171 < 10; fi171++ {
            fmt.Println(v5, v3, v170)
            var v172 int32 = v1
            v173 := GenS18{F0: 42, F1: 199}
            _ = fi171
            _ = v172
            _ = v173
        }
        _ = v169
        _ = v170
    } else {
        fmt.Println(v1, v5)
        switch ("" + ((((("\n\t" + "hello") + "\n\t") + "") + "hello") + ("a" + ("foo bar" + "hello")))) {
        case "\n\t":
            var v174 uint = 42
            _ = v174
        default:
            var v175 uint = 840
            _ = v175
        }
        fmt.Println(v3)
    }
    var v176 uint8 = 127
    fmt.Println(v176, v176, v1)
    v176 -= 255
    var v177 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v163:", len(v163), v163)
    fmt.Println("v164:", v164)
    fmt.Println("v176:", v176)
    fmt.Println("v177:", v177)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v162
    _ = v163
    _ = v164
    _ = v166
    _ = v176
    _ = v177
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
