package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 uint16; F1 uint64 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint16; F1 uint16 }
type GenS10 struct { F0 bool; F1 bool }
type GenS11 struct { F0 uint; F1 uint16; F2 int32 }
type GenS12 struct { F0 uint8; F1 string; F2 uint16 }
func genFn0() int {
    var v7 uint32 = 2147483648
    var v8 int = 4294967295
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 string, p1 int64, p2 string) uint8 {
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := []int{(^(fi9)), (fi9 << ((697) & 63)), fi9, fi9}
        _ = fi9
        _ = v10
    }
    if ((831 >= genFn0()) || false) {
        var v11 float32 = ((0.1 + (3.14159 + -2.25)) / 1.0)
        p0 = ((p2 + (p2 + "a")) + p2)
        _ = v11
    } else {
        if true {
            v12 := []int{}
            p0 = "a"
            if (9 < 55) {
                p1 = p1
            }
            _ = v12
        } else {
            v13 := []int{}
            var v14 uint32 = 2147483647
            if len(v13) > 0 {
                v13[0] = 128
            }
            _ = v13
            _ = v14
        }
        v15 := []int{9223372036854775807, 4294967295}
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        fmt.Println(p0, fi16)
        switch (p2 + "a") {
        case "foo bar":
            var v17 float64 = (((0.0 + 0.0) + 1e10) / 1e10)
            _ = v17
        default:
            switch ("hello" + p2) {
            case "\n\t":
                switch fi16 {
                case -1:
                    v18 := GenS1{F0: (-(p1))}
                    _ = v18
                case 863:
                    for fi19 := 0; fi19 < 6; fi19++ {
                        v20 := []int{(fi16 + (fi19 >> ((0) & 63))), (-(fi19))}
                        v21 := GenS2{F0: 0, F1: false}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                case 2147483647:
                    p2 = ""
                default:
                    v22 := GenS3{F0: (p2 + p0), F1: 42}
                    _ = v22
                }
            case "foo bar":
                if true {
                    p0 = (p0 + p2)
                    for fi23 := 0; fi23 < 9; fi23++ {
                        var v24 string = ((p2 + ("\n\t" + "hello")) + "hello")
                        p1 = p1
                        var v25 uint32 = 1
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                } else {
                    if true {
                        if false {
                            switch ("foo bar" + "a") {
                            case "\n\t":
                                if false {
                                    var v26 int32 = 0
                                    _ = v26
                                }
                            default:
                                var v27 int32 = -2147483648
                                _ = v27
                            }
                        } else {
                            for fi28 := 0; fi28 < 13; fi28++ {
                                fmt.Println(p2, fi16)
                                _ = fi28
                            }
                            v29 := GenS4{F0: 225}
                            _ = v29
                        }
                        if true {
                            v30 := []int{fi16, fi16, (fi16 / (((fi16 / ((-620) | 1))) | 1)), int(p1)}
                            _ = len(v30)
                            for fi31 := 0; fi31 < 20; fi31++ {
                                var v32 uint32 = 1
                                v33 := GenS4{F0: 2}
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                            _ = v30
                        }
                    }
                    var v34 uint16 = 245
                    p2 = p0
                    _ = v34
                }
            case "hello":
                p0 = p2
            default:
                var v35 uint8 = 128
                _ = v35
            }
        }
        _ = fi16
    }
    var v36 int8 = 127
    var v37 float64 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v36
    _ = v37
    return 127
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 54
    var v3 float32 = 1e10
    var v4 int = 26
    var v5 int8 = 0
    v6 := make(map[uint32]uint32)
    v38 := GenS1{F0: -9223372036854775808}
    v39 := []int{(v4 << ((2) & 63)), v4, (-(v4)), (v4 << ((4294967295) & 63))}
    switch (v4 - 2147483647) {
    case -2147483648:
        fmt.Println(v2, v5)
    }
    var v40 uint = 1
    if (v40 >= 379) {
        v4 -= (v4 * v4)
        v5 = (v5 | (v5 << ((v40) & 63)))
    }
    for fi41 := 0; fi41 < 7; fi41++ {
        for fi42 := 0; fi42 < 12; fi42++ {
            if false {
                v2 = v2
                v1[("a" + (("foo bar" + ((((("\n\t" + "") + "foo bar") + "a") + "a") + "hello")) + ("a" + "foo bar")))] = 42
            } else {
                v39 = append(v39, (fi42 - fi42))
                v2 = genFn1("\n\t", -86, "")
                if (fi41 < -370) {
                    switch (v4 - fi41) {
                    case -2147483648:
                        var v43 int32 = -589
                        _ = v43
                    }
                    if (!(!(genFn0() != genFn0()))) {
                        if len(v39) > 0 {
                            v39[0] = 4294967295
                        }
                        fmt.Println(v3, fi41, fi41)
                        var v44 uint = (v40 | 42)
                        var v45 int16 = 0
                        _ = v44
                        _ = v45
                    } else {
                        v46 := GenS3{F0: "", F1: 18446744073709551615}
                        switch (^(fi41)) {
                        case -2147483648:
                            var v47 int32 = 42
                            _ = v47
                        case 627:
                            if (2147483647 != -1) {
                                v48 := GenS1{F0: -1}
                                if true {
                                    v40 -= 0
                                    switch 32768 {
                                    case 4294967295:
                                        if len(v39) >= 2 {
                                            v39 = v39[1:len(v39)]
                                        }
                                    default:
                                        var v49 bool = false
                                        _ = v49
                                    }
                                    var v50 int64 = -9223372036854775808
                                    _ = v50
                                }
                                _ = v48
                            } else {
                                fmt.Println(v40)
                                switch "\n\t" {
                                case "a":
                                    v39 = append(v39, genFn0())
                                    if len(v39) >= 2 {
                                        v39 = v39[1:len(v39)]
                                    }
                                case "\n\t":
                                    fmt.Println(fi42, v40)
                                case "hello":
                                    switch ("hello" + (("a" + "") + "\n\t")) {
                                    case "":
                                        delete(v6, 42)
                                    default:
                                        if true {
                                            fmt.Println(fi42)
                                        } else {
                                            if (genFn0() >= (fi42 - v4)) {
                                                if ((^(v2)) < 127) {
                                                    v40 += 42
                                                    switch 1 {
                                                    case 9223372036854775807:
                                                        v51 := GenS5{F0: -9223372036854775808, F1: -269}
                                                        _ = v51
                                                    case 4294967295:
                                                        v40 = (v40 / ((0) | 1))
                                                    case 127:
                                                        v52 := GenS6{F0: 4294967295, F1: 2147483647}
                                                        _ = v52
                                                    }
                                                } else {
                                                    v40 = (v40 >> ((v40) & 63))
                                                }
                                            }
                                            if (false && true) {
                                                var v53 uint8 = 128
                                                if false {
                                                    if (2 != 159) {
                                                        fmt.Println(v5, v4)
                                                        v40 = v40
                                                        switch (v4 | genFn0()) {
                                                        case -1:
                                                            if (1.5 == 1e10) {
                                                                var v54 int8 = (v5 % ((v5) | 1))
                                                                var v55 int16 = -959
                                                                v5 = (-(v5))
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                        case 32768:
                                                            if (int32(v5) == 42) {
                                                                if len(v39) > 0 {
                                                                    v39[0] = (fi41 / (((-(fi41))) | 1))
                                                                }
                                                                if false {
                                                                    delete(v6, 230)
                                                                    fmt.Println(v5, fi41, fi41)
                                                                }
                                                            } else {
                                                                v2 += uint8(v40)
                                                                var v56 int16 = 1
                                                                fmt.Println(fi41, v40)
                                                                v57 := GenS3{F0: (("a" + "a") + "hello"), F1: 42}
                                                                v56 = (v56 >> ((v40) & 63))
                                                                _ = v56
                                                                _ = v57
                                                            }
                                                        default:
                                                            v4 = v4
                                                        }
                                                        if false {
                                                            fmt.Println(v53, v53, v53)
                                                            if (v40 > (v40 & 863)) {
                                                                var v58 int16 = -548
                                                                var v59 int16 = v58
                                                                switch ("" + "\n\t") {
                                                                case "\n\t":
                                                                    v53 += (v2 | 223)
                                                                case "":
                                                                    if false {
                                                                        var v60 int = 0
                                                                        if (v40 <= (^(v40))) {
                                                                            _ = len(v1)
                                                                            var v61 float32 = ((v3 - -2.25) / 1e10)
                                                                            _ = v61
                                                                        } else {
                                                                            if true {
                                                                                var v62 uint64 = 0
                                                                                _ = v62
                                                                            }
                                                                            v3 -= (v3 * (v3 + v3))
                                                                        }
                                                                        _ = v60
                                                                    } else {
                                                                        var v63 string = "a"
                                                                        if true {
                                                                            var v64 int16 = 32767
                                                                            _ = v64
                                                                        } else {
                                                                            fmt.Println(v40, v5, v4)
                                                                            if (v53 < v53) {
                                                                                fmt.Println(v59, v53)
                                                                                fmt.Println(v58)
                                                                                _ = len(v6)
                                                                            } else {
                                                                                fmt.Println(v3)
                                                                                fmt.Println(v59)
                                                                            }
                                                                            switch 0 {
                                                                            case 32767:
                                                                                var v65 uint8 = v53
                                                                                _ = v65
                                                                            case -2147483648:
                                                                                _ = len(v39)
                                                                            case -103:
                                                                                v66 := v6[65535]
                                                                                _ = v66
                                                                            default:
                                                                                var v67 bool = (!(!false))
                                                                                _ = v67
                                                                            }
                                                                        }
                                                                        _ = v63
                                                                    }
                                                                default:
                                                                    v68 := GenS3{F0: "", F1: 1}
                                                                    _ = v68
                                                                }
                                                                _ = v58
                                                                _ = v59
                                                            } else {
                                                                var v69 string = "a"
                                                                _ = v69
                                                            }
                                                            v2 += 255
                                                        } else {
                                                            v2 = v2
                                                            v70 := GenS4{F0: (v53 & genFn1((("a" + "a") + "foo bar"), int64(v40), ("a" + "a")))}
                                                            _ = v70
                                                        }
                                                        if len(v39) > 0 {
                                                            v39[0] = fi42
                                                        }
                                                    }
                                                }
                                                var v71 float32 = v3
                                                _ = v53
                                                _ = v71
                                            } else {
                                                var v72 int = fi42
                                                if false {
                                                    v73 := v1[("a" + "")]
                                                    fmt.Println(fi41, v40)
                                                    var v74 uint = v40
                                                    _ = v73
                                                    _ = v74
                                                }
                                                var v75 int16 = 32767
                                                _ = v72
                                                _ = v75
                                            }
                                            var v76 bool = false
                                            _ = v76
                                        }
                                    }
                                default:
                                    var v77 int8 = 127
                                    _ = v77
                                }
                                var v78 float32 = (v3 + v3)
                                v79 := GenS3{F0: "\n\t", F1: 341}
                                if ((v5 - v5) < 74) {
                                    fmt.Println(v78, v4)
                                    if (v40 >= (v40 & v40)) {
                                        var v80 uint64 = 42
                                        v81 := GenS4{F0: v2}
                                        _ = v80
                                        _ = v81
                                    } else {
                                        switch "a" {
                                        case "a":
                                            switch "" {
                                            case "hello":
                                                var v82 float64 = 1e10
                                                _ = v82
                                            case "foo bar":
                                                if (true || (22 > 2)) {
                                                    if true {
                                                        if (!(!(((!true) && false) || true))) {
                                                            v2 = v2
                                                        }
                                                        var v83 bool = false
                                                        _ = v83
                                                    }
                                                }
                                            case "":
                                                var v84 int8 = (v5 - -128)
                                                _ = v84
                                            }
                                        case "hello":
                                            _ = len(v39)
                                        default:
                                            if false {
                                                delete(v6, 795)
                                                if (1000 >= 1) {
                                                    v4 -= (fi41 & v4)
                                                    v85 := GenS7{F0: 0, F1: 0}
                                                    fmt.Println(v4, v78)
                                                    _ = v85
                                                } else {
                                                    v6[805] = 65535
                                                    v2 -= v2
                                                    var v86 int8 = v5
                                                    if true {
                                                        var v87 int = (fi42 + fi42)
                                                        var v88 int64 = 42
                                                        _ = v87
                                                        _ = v88
                                                    } else {
                                                        var v89 bool = true
                                                        v90 := GenS8{F0: v2}
                                                        var v91 float64 = 3.14159
                                                        _ = v89
                                                        _ = v90
                                                        _ = v91
                                                    }
                                                    _ = v86
                                                }
                                            }
                                        }
                                    }
                                    v4 = genFn0()
                                } else {
                                    var v92 uint32 = 0
                                    _ = v92
                                }
                                var v93 int64 = 9223372036854775807
                                var v94 uint32 = 2
                                _ = v78
                                _ = v79
                                _ = v93
                                _ = v94
                            }
                        case 1:
                            if ("foo bar" != ("\n\t" + "foo bar")) {
                                if false {
                                    var v95 bool = (!(2 != 749))
                                    v5 -= (v5 - v5)
                                    _ = v95
                                } else {
                                    var v96 int16 = 32767
                                    _ = v96
                                }
                                if false {
                                    v40 -= v40
                                }
                            } else {
                                if (((v2 | v2) == (v2 ^ (v2 * genFn1("foo bar", 113, "foo bar")))) || (517 > 925)) {
                                    switch ("foo bar" + (("foo bar" + "") + "")) {
                                    case "foo bar":
                                        v2 = genFn1("a", 1, ("foo bar" + "\n\t"))
                                    case "":
                                        var v97 int16 = 694
                                        _ = v97
                                    case "hello":
                                        v4 = fi41
                                    default:
                                        if (-85 >= (v4 * fi41)) {
                                            v5 -= (v5 >> (((v40 + 271)) & 63))
                                            v4 = genFn0()
                                        } else {
                                            var v98 int32 = -661
                                            v99 := GenS7{F0: 2, F1: 2}
                                            _ = v98
                                            _ = v99
                                        }
                                    }
                                    if true {
                                        var v100 uint16 = 255
                                        v101 := GenS4{F0: v2}
                                        _ = v100
                                        _ = v101
                                    } else {
                                        v3 += float32(float64(v3))
                                    }
                                    var v102 float32 = (v3 * (v3 - v3))
                                    v5 += v5
                                    _ = v102
                                } else {
                                    fmt.Println(v4, v4, v40)
                                    fmt.Println(v3, v3, v4)
                                    var v103 string = "\n\t"
                                    _ = v103
                                }
                                v5 += v5
                            }
                        default:
                            v2 = genFn1(("foo bar" + "a"), 0, "a")
                        }
                        if (!true) {
                            v1["a"] = 4294967295
                        } else {
                            _ = len(v39)
                            if (43 <= 9223372036854775807) {
                                v104 := GenS8{F0: v2}
                                _ = len(v1)
                                v5 = v5
                                _ = v104
                            } else {
                                switch -9223372036854775808 {
                                case 4294967295:
                                    v3 = (v3 / (v3 / 1.0))
                                }
                            }
                            v5 -= (v5 ^ (v5 * v5))
                            switch fi41 {
                            case -1:
                                if true {
                                    v4 = (fi42 | (^(v4)))
                                    if (false && true) {
                                        v105 := GenS9{F0: 798, F1: 1}
                                        var v106 int8 = (v5 % ((v5) | 1))
                                        var v107 float32 = 0.0
                                        delete(v6, 4294967295)
                                        if (!(!false)) {
                                            v4 = (fi42 ^ fi41)
                                        }
                                        var v108 int32 = -921
                                        v106 -= v5
                                        fmt.Println(v40, v106, v2)
                                        if (!((0.1 < 3.14159) || true)) {
                                            v4 = (fi42 & (v4 / (((v4 + (v4 >> ((v40) & 63)))) | 1)))
                                            var v109 uint = v40
                                            _ = v109
                                        }
                                        fmt.Println(v40, v106)
                                        _ = v105
                                        _ = v106
                                        _ = v107
                                        _ = v108
                                    }
                                    var v110 float32 = (0.0 - v3)
                                    _ = v110
                                } else {
                                    v3 -= v3
                                    if len(v39) > 0 {
                                        v39[0] = v4
                                    }
                                    v111 := GenS6{F0: 4294967295, F1: -2147483648}
                                    if false {
                                        v5 = (v5 & 127)
                                        v112 := GenS10{F0: (2147483647 != 2147483648), F1: false}
                                        fmt.Println(v2, v40, fi41)
                                        _ = v112
                                    } else {
                                        if len(v39) >= 2 {
                                            v39 = v39[1:len(v39)]
                                        }
                                    }
                                    var v113 uint64 = 42
                                    var v114 uint = (v40 >> (((^(v40))) & 63))
                                    _ = v111
                                    _ = v113
                                    _ = v114
                                }
                            case 127:
                                switch fi41 {
                                case 32768:
                                    v2 = v2
                                case -2147483648:
                                    v115 := GenS9{F0: 33, F1: 42}
                                    _ = v115
                                case -979:
                                    var v116 int = genFn0()
                                    _ = v116
                                default:
                                    v5 += v5
                                }
                            case -2147483648:
                                v5 = (v5 & (v5 + (-(v5))))
                            default:
                                var v117 bool = true
                                _ = v117
                            }
                            var v118 float32 = (v3 * v3)
                            var v119 bool = true
                            _ = v118
                            _ = v119
                        }
                        _ = v46
                    }
                    var v120 int = (v4 + v4)
                    _ = v120
                } else {
                    if (65 > 2147483648) {
                        fmt.Println(v5, v3, fi42)
                    }
                    v4 = (^(fi41))
                    v121 := GenS1{F0: -1}
                    _ = v121
                }
            }
            var v122 uint32 = 305
            _ = fi42
            _ = v122
        }
        _ = fi41
    }
    switch "\n\t" {
    case "foo bar":
        var v123 uint8 = v2
        _ = v123
    case "":
        for fi124 := 0; fi124 < 6; fi124++ {
            v4 -= 970
            if (("\n\t" >= "\n\t") || ((v2 >> ((v40) & 63)) != (v2 >> (((v40 % ((v40) | 1))) & 63)))) {
                fmt.Println(v3)
                v5 += -116
                var v125 uint16 = 65535
                _ = v125
            } else {
                v126 := GenS5{F0: -630, F1: -1}
                fmt.Println(v2, v3)
                for fi127 := 0; fi127 < 9; fi127++ {
                    fmt.Println(fi127)
                    switch "foo bar" {
                    case "foo bar":
                        v4 += genFn0()
                    }
                    v128 := GenS11{F0: (v40 & v40), F1: 0, F2: -915}
                    _ = fi127
                    _ = v128
                }
                _ = v126
            }
            _ = fi124
        }
    case "\n\t":
        var v129 int = -9223372036854775808
        _ = v129
    }
    var v130 uint = 523
    var v131 int64 = int64(v5)
    switch "" {
    case "":
        for fi132 := 0; fi132 < 8; fi132++ {
            if (1000 == 19) {
                for fi133 := 0; fi133 < 1; fi133++ {
                    fmt.Println(v3, v40, v131)
                    v4 += 128
                    _ = fi133
                }
                if (("" < ("a" + ("foo bar" + ("" + "a")))) || ((!(!((v2 | (v2 * (v2 - (v2 * 1)))) == v2))) || (1 >= -2147483648))) {
                    var v134 float32 = (((0.1 - float32((float64(1.5) - 0.1))) * (float32(((1e10 / (3.14159 + 0.1)) * ((1.5 / 1.5) + (((0.0 + 1e10) / -2.25) * 1e10)))) * (((1.5 + 3.14159) - 3.14159) - ((1.5 / v3) / v3)))) + v3)
                    fmt.Println(v40, v130)
                    v134 -= ((0.0 + (v3 * ((1e10 + float32(((1e10 - ((1e10 * (((0.0 * 1.5) * ((0.1 + 1.5) - 0.0)) / float64((v134 - (v3 + 1.5))))) * (float64(-2.25) + 0.1))) / (1e10 + 3.14159)))) / -2.25))) * v134)
                    _ = v134
                }
                fmt.Println(v3, v40, v40)
            }
            if ((fi132 <= (v4 / ((fi132) | 1))) || (false || ((0 <= 42) || true))) {
                var v135 int32 = -285
                _ = v135
            } else {
                if true {
                    v39 = append(v39, (v4 % (((fi132 | 2147483647)) | 1)))
                } else {
                    switch "a" {
                    case "":
                        fmt.Println(v3, v4, v130)
                    case "a":
                        var v136 int8 = (-(v5))
                        _ = v136
                    case "hello":
                        fmt.Println(v131, fi132, v4)
                    }
                    var v137 uint16 = 65535
                    _ = v137
                }
            }
            _ = fi132
        }
    case "hello":
        v4 = v4
    case "foo bar":
        v138 := GenS12{F0: genFn1("hello", -9223372036854775808, "\n\t"), F1: "", F2: 255}
        _ = v138
    }
    var v139 uint16 = 265
    if false {
        for fi140 := 0; fi140 < 1; fi140++ {
            _ = len(v39)
            v131 -= (v131 | v131)
            _ = fi140
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v130:", v130)
    fmt.Println("v131:", v131)
    fmt.Println("v139:", v139)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v130
    _ = v131
    _ = v139
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
