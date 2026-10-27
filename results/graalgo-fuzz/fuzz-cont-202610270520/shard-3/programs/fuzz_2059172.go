package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int64 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int16; F1 int32; F2 int8 }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool; F1 bool; F2 uint8 }
type GenS7 struct { F0 int64; F1 uint64 }
type GenS8 struct { F0 uint32; F1 uint8 }
type GenS9 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS10 struct { F0 int; F1 uint8 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int; F1 uint64 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 uint16; F1 bool }
type GenS16 struct { F0 uint32; F1 int32 }
type GenS17 struct { F0 int8; F1 int8; F2 uint64 }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 uint8; F1 int64 }
type GenS20 struct { F0 uint16; F1 int32 }
func genFn0(p0 bool, p1 string) uint64 {
    switch "foo bar" {
    case "a":
        v6 := []int{}
        _ = v6
    case "\n\t":
        fmt.Println(p1, p0)
    default:
        if (0.0 > 1.5) {
            for fi7 := 0; fi7 < 20; fi7++ {
                v8 := GenS1{F0: 42, F1: 42, F2: -735}
                _ = fi7
                _ = v8
            }
        }
    }
    var v9 float32 = (((1.5 / -2.25) + float32(((float64((1e10 * (1.5 / (3.14159 + -2.25)))) * ((((((((3.14159 - 1e10) + 3.14159) * 3.14159) * (0.1 / float64(3.14159))) + -2.25) + 0.1) * 3.14159) + float64(((-2.25 + -2.25) - (1.5 / -2.25))))) - 3.14159))) - float32(0.0))
    _ = p0
    _ = p1
    _ = v9
    return 1
}
func genFn1() string {
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 uint32 = 1
        var v12 uint64 = 18446744073709551615
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := GenS2{F0: 1, F1: -1}
            v15 := []int{}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    var v16 int16 = 1000
    v17 := []int{-2147483648, 9223372036854775807}
    var v18 int8 = 77
    if true {
        var v19 int32 = -610
        _ = v19
    }
    _ = v16
    _ = v17
    _ = v18
    return "\n\t"
}

func main() {
    var v1 float64 = (1e10 + 1.5)
    v2 := make(map[string]int)
    var v3 int64 = -1
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 128
    if (1 <= 1000) {
        for fi20 := 0; fi20 < 12; fi20++ {
            v21 := GenS3{F0: -111, F1: 635, F2: 42}
            for fi22 := 0; fi22 < 4; fi22++ {
                var v23 float64 = v1
                var v24 uint32 = 1
                fmt.Println(v5, v23)
                _ = fi22
                _ = v23
                _ = v24
            }
            switch fi20 {
            case 127:
                _ = len(v4)
            case 32768:
                var v25 int8 = 42
                _ = v25
            }
            _ = fi20
            _ = v21
        }
        v3 = (v3 ^ -780)
        fmt.Println(v3)
    }
    var v26 int = 760
    v26 += (v26 & (v26 * (v26 + v26)))
    if true {
        var v27 int64 = (v3 * (^(v3)))
        _ = v27
    } else {
        v26 = int(v5)
        switch "" {
        case "foo bar":
            if true {
                var v28 uint32 = 807
                if false {
                    var v29 uint8 = 0
                    var v30 uint8 = v29
                    for fi31 := 0; fi31 < 16; fi31++ {
                        fmt.Println(v3, fi31)
                        for fi32 := 0; fi32 < 15; fi32++ {
                            v33 := v2["a"]
                            if false {
                                v29 += v30
                                v34 := GenS4{F0: v28, F1: 2, F2: v28}
                                v35 := []int{(v33 % ((v26) | 1)), (-(v33)), fi31}
                                _ = v34
                                _ = v35
                            } else {
                                var v36 uint64 = 91
                                var v37 uint64 = genFn0(((v1 / (-2.25 * -2.25)) < (v1 - v1)), "foo bar")
                                if false {
                                    v36 = v36
                                    v3 = -1
                                }
                                _ = v36
                                _ = v37
                            }
                            _ = fi32
                            _ = v33
                        }
                        v4[258] = 192
                        _ = fi31
                    }
                    _ = v29
                    _ = v30
                } else {
                    fmt.Println(v3, v5)
                    switch ("" + ("hello" + "hello")) {
                    case "":
                        v5 -= 128
                    case "foo bar":
                        v38 := GenS1{F0: 2, F1: 42, F2: 9223372036854775807}
                        _ = v38
                    case "hello":
                        v3 += 42
                    default:
                        if (65535 == 65535) {
                            fmt.Println(v26)
                            v1 += v1
                            var v39 float32 = 1.5
                            _ = v39
                        } else {
                            fmt.Println(v5, v3)
                        }
                    }
                    v5 += 66
                }
                v1 -= 3.14159
                _ = v28
            }
        case "":
            v40 := GenS2{F0: 255, F1: (v26 ^ (v26 >> ((1) & 63)))}
            _ = v40
        case "hello":
            var v41 int16 = 0
            _ = v41
        default:
            for fi42 := 0; fi42 < 6; fi42++ {
                if false {
                    var v43 int = (fi42 - -9223372036854775808)
                    fmt.Println(v3, fi42)
                    v26 -= (^(v26))
                    _ = v43
                } else {
                    switch "foo bar" {
                    case "hello":
                        if true {
                            v44 := GenS2{F0: (v5 >> ((885) & 63)), F1: v26}
                            for fi45 := 0; fi45 < 18; fi45++ {
                                var v46 int32 = 0
                                _ = fi45
                                _ = v46
                            }
                            _ = v44
                        }
                    case "\n\t":
                        var v47 float64 = v1
                        _ = v47
                    case "a":
                        v26 = (v26 % ((v26) | 1))
                    default:
                        v5 -= v5
                    }
                }
                delete(v2, "\n\t")
                v3 -= v3
                _ = fi42
            }
        }
        var v48 string = ""
        _ = v48
    }
    var v49 float32 = 1e10
    var v50 int8 = 127
    v51 := GenS4{F0: 191, F1: 65535, F2: 760}
    v5 -= (v5 & v5)
    fmt.Println(v1, v3, v26)
    v52 := []int{(-(v26)), v26}
    switch -9223372036854775808 {
    case 983:
        v26 -= (v26 | (v26 << ((2) & 63)))
    case -9223372036854775808:
        fmt.Println(v49, v49)
    case 9223372036854775807:
        v3 = v3
    }
    var v53 int = v26
    var v54 int64 = int64(v5)
    if ("foo bar" > "foo bar") {
        if true {
            var v55 float32 = v49
            v56 := GenS3{F0: 42, F1: 1, F2: -128}
            _ = len(v4)
            _ = v55
            _ = v56
        }
        switch 0 {
        case 42:
            if (true && (("" + ("a" + ("a" + ""))) < "\n\t")) {
                var v57 uint32 = 504
                _ = v57
            }
        case 0:
            v58 := v4[2147483648]
            _ = v58
        case -335:
            if (true || (!true)) {
                v59 := GenS3{F0: 1000, F1: 39, F2: (-(v50))}
                _ = v59
            } else {
                var v60 int8 = v50
                _ = v60
            }
        }
    } else {
        _ = len(v52)
        if ((!true) && ((-632 == 32767) || ((32767 != 255) || ("hello" >= "a")))) {
            v61 := v4[4294967295]
            var v62 string = ""
            _ = v61
            _ = v62
        } else {
            var v63 int8 = v50
            if true {
                if (1 == 423) {
                    v64 := GenS5{F0: true}
                    v49 -= v49
                    if false {
                        for fi65 := 0; fi65 < 9; fi65++ {
                            v1 -= v1
                            if len(v52) > 0 {
                                v52[0] = (fi65 % ((v53) | 1))
                            }
                            _ = fi65
                        }
                        if ((255 == 0) && ((0 < -411) && false)) {
                            v52 = append(v52, -199)
                            switch (^(v53)) {
                            case 128:
                                if ((!(((0.1 / -2.25) - v1) <= v1)) || (false || true)) {
                                    if (v5 < (^(v5))) {
                                        for fi66 := 0; fi66 < 10; fi66++ {
                                            v67 := GenS6{F0: true, F1: (true && true), F2: v5}
                                            fmt.Println(v54, v53, v54)
                                            fmt.Println(v3, fi66)
                                            _ = fi66
                                            _ = v67
                                        }
                                        for fi68 := 0; fi68 < 10; fi68++ {
                                            if (2 >= 0) {
                                                if false {
                                                    var v69 uint32 = 42
                                                    switch 672 {
                                                    case -9223372036854775808:
                                                        fmt.Println(v63, v63)
                                                    case -2147483648:
                                                        var v70 uint64 = 882
                                                        _ = v70
                                                    }
                                                    if (42 == (fi68 / ((v53) | 1))) {
                                                        v71 := GenS7{F0: (v54 + (v54 * v54)), F1: 18446744073709551615}
                                                        _ = v71
                                                    }
                                                    _ = v69
                                                }
                                                v53 = (v26 - fi68)
                                            }
                                            v5 = v5
                                            switch (-(v26)) {
                                            case -2147483648:
                                                v52 = append(v52, -9223372036854775808)
                                            }
                                            _ = fi68
                                        }
                                    } else {
                                        v52 = append(v52, v26)
                                        if false {
                                            v72 := GenS8{F0: 65535, F1: 1}
                                            switch 127 {
                                            case 0:
                                                var v73 int16 = -32768
                                                _ = v73
                                            default:
                                                var v74 string = "\n\t"
                                                _ = v74
                                            }
                                            for fi75 := 0; fi75 < 4; fi75++ {
                                                if len(v52) >= 2 {
                                                    v52 = v52[1:len(v52)]
                                                }
                                                switch v26 {
                                                case 0:
                                                    fmt.Println(v63)
                                                case -2147483648:
                                                    _ = len(v52)
                                                default:
                                                    if ((((v5 / ((v5) | 1)) != (v5 >> ((4294967295) & 63))) && false) && (!(4294967295 == genFn0(false, "\n\t")))) {
                                                        v5 += (v5 / ((v5) | 1))
                                                        for fi76 := 0; fi76 < 8; fi76++ {
                                                            switch v26 {
                                                            case -581:
                                                                if false {
                                                                    fmt.Println(v49, v54)
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        var v77 uint32 = 752
                                                                        _ = v77
                                                                    case "\n\t":
                                                                        if (true && (!true)) {
                                                                            if false {
                                                                                v78 := GenS3{F0: -1, F1: 0, F2: (v63 >> ((42) & 63))}
                                                                                if len(v52) >= 2 {
                                                                                    v52 = v52[1:len(v52)]
                                                                                }
                                                                                _ = v78
                                                                            }
                                                                        } else {
                                                                            if false {
                                                                                _ = len(v4)
                                                                                if (false && (false || false)) {
                                                                                    fmt.Println(fi76, v53)
                                                                                    var v79 uint8 = (^(v5))
                                                                                    _ = v79
                                                                                } else {
                                                                                    v5 = 2
                                                                                }
                                                                                v5 += (v5 / ((v5) | 1))
                                                                                if false {
                                                                                    if true {
                                                                                        v4[42] = 592
                                                                                    } else {
                                                                                        v80 := GenS4{F0: 2147483648, F1: 65535, F2: 0}
                                                                                        _ = v80
                                                                                    }
                                                                                } else {
                                                                                    v81 := GenS6{F0: false, F1: false, F2: 2}
                                                                                    var v82 int = 127
                                                                                    _ = v81
                                                                                    _ = v82
                                                                                }
                                                                                if false {
                                                                                    if (false && false) {
                                                                                        fmt.Println(v53, v54)
                                                                                        if false {
                                                                                            if (!(0 <= 0)) {
                                                                                                var v83 int8 = v50
                                                                                                var v84 string = ("foo bar" + (("" + "hello") + ("foo bar" + ("" + ((("a" + "a") + ("" + "foo bar")) + "a")))))
                                                                                                _ = v83
                                                                                                _ = v84
                                                                                            } else {
                                                                                                if true {
                                                                                                    if (1 >= 591) {
                                                                                                        fmt.Println(v49, fi75, fi76)
                                                                                                        v50 -= int8(v3)
                                                                                                    }
                                                                                                    var v85 float64 = v1
                                                                                                    var v86 bool = false
                                                                                                    _ = v85
                                                                                                    _ = v86
                                                                                                }
                                                                                                var v87 uint32 = 2147483648
                                                                                                if false {
                                                                                                    fmt.Println(v63, v53, v3)
                                                                                                    if (!true) {
                                                                                                        fmt.Println(v50, v49)
                                                                                                        if false {
                                                                                                            switch fi76 {
                                                                                                            case 127:
                                                                                                                v49 -= float32(0.0)
                                                                                                            case 1:
                                                                                                                v87 = (v87 % (((v87 ^ (v87 | 4294967295))) | 1))
                                                                                                            }
                                                                                                        } else {
                                                                                                            fmt.Println(v3, fi75)
                                                                                                            _ = len(v2)
                                                                                                            if ((!false) || true) {
                                                                                                                switch "foo bar" {
                                                                                                                case "hello":
                                                                                                                    v88 := GenS8{F0: 553, F1: 2}
                                                                                                                    _ = v88
                                                                                                                default:
                                                                                                                    v89 := GenS4{F0: v87, F1: 237, F2: 2147483648}
                                                                                                                    _ = v89
                                                                                                                }
                                                                                                                switch ("hello" + "") {
                                                                                                                case "foo bar":
                                                                                                                    var v90 string = "foo bar"
                                                                                                                    _ = v90
                                                                                                                case "":
                                                                                                                    v87 = (^(v87))
                                                                                                                case "\n\t":
                                                                                                                    v91 := GenS9{F0: (v54 - 42), F1: (^(v87)), F2: (v87 % ((65535) | 1))}
                                                                                                                    _ = v91
                                                                                                                }
                                                                                                            }
                                                                                                            fmt.Println(v49, v50, v50)
                                                                                                        }
                                                                                                    }
                                                                                                    var v92 int8 = v63
                                                                                                    _ = v92
                                                                                                } else {
                                                                                                    var v93 int = v53
                                                                                                    if false {
                                                                                                        v94 := GenS9{F0: v54, F1: (v87 << ((255) & 63)), F2: (v87 & 2147483647)}
                                                                                                        var v95 string = "a"
                                                                                                        if false {
                                                                                                            var v96 uint32 = (v87 * v87)
                                                                                                            _ = v96
                                                                                                        } else {
                                                                                                            v95 = v95
                                                                                                            if (!(1 <= 1000)) {
                                                                                                                if true {
                                                                                                                    var v97 uint8 = 42
                                                                                                                    if ((0 > 65535) || (!true)) {
                                                                                                                        if len(v52) > 0 {
                                                                                                                            v52[0] = (fi76 & fi75)
                                                                                                                        }
                                                                                                                        if true {
                                                                                                                            v50 += v50
                                                                                                                            var v98 int16 = 0
                                                                                                                            v99 := GenS8{F0: (v87 << ((255) & 63)), F1: 128}
                                                                                                                            _ = v98
                                                                                                                            _ = v99
                                                                                                                        } else {
                                                                                                                            if false {
                                                                                                                                v100 := GenS10{F0: fi76, F1: (v97 - (v5 & (^(v5))))}
                                                                                                                                v101 := GenS6{F0: ((true && ((v3 >= (v3 / ((-843) | 1))) && false)) || (((v1 - v1) - (v1 - v1)) != (v1 - v1))), F1: false, F2: (v5 & (v5 - (^(v5))))}
                                                                                                                                v102 := GenS11{F0: 104}
                                                                                                                                _ = v100
                                                                                                                                _ = v101
                                                                                                                                _ = v102
                                                                                                                            } else {
                                                                                                                                var v103 uint8 = 2
                                                                                                                                delete(v4, 1)
                                                                                                                                if true {
                                                                                                                                    v104 := GenS12{F0: v93, F1: uint64(v5)}
                                                                                                                                    _ = v104
                                                                                                                                }
                                                                                                                                _ = v103
                                                                                                                            }
                                                                                                                            v105 := GenS10{F0: fi76, F1: 127}
                                                                                                                            v2[("" + v95)] = v93
                                                                                                                            _ = v105
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v97 += 1
                                                                                                                    _ = v97
                                                                                                                }
                                                                                                            }
                                                                                                            v106 := v4[0]
                                                                                                            _ = v106
                                                                                                        }
                                                                                                        if ((float64(v49) - float64(0.1)) == float64(float32(v1))) {
                                                                                                            var v107 float32 = v49
                                                                                                            _ = v107
                                                                                                        }
                                                                                                        var v108 int16 = -727
                                                                                                        _ = v108
                                                                                                        _ = v94
                                                                                                        _ = v95
                                                                                                    } else {
                                                                                                        v109 := GenS13{F0: 1000}
                                                                                                        var v110 bool = true
                                                                                                        var v111 bool = (18446744073709551615 != 520)
                                                                                                        _ = v109
                                                                                                        _ = v110
                                                                                                        _ = v111
                                                                                                    }
                                                                                                    switch "a" {
                                                                                                    case "":
                                                                                                        v5 += 0
                                                                                                    case "\n\t":
                                                                                                        switch v26 {
                                                                                                        case 32767:
                                                                                                            v5 = 127
                                                                                                            v112 := GenS10{F0: -58, F1: 42}
                                                                                                            if ((v26 * v53) != -417) {
                                                                                                                v93 = 0
                                                                                                                var v113 uint8 = (v5 << ((11) & 63))
                                                                                                                _ = v113
                                                                                                            } else {
                                                                                                                var v114 uint8 = (v5 % (((v5 ^ (v5 * (v5 % ((128) | 1))))) | 1))
                                                                                                                v52 = append(v52, v93)
                                                                                                                switch -9223372036854775808 {
                                                                                                                case 511:
                                                                                                                    switch "foo bar" {
                                                                                                                    case "hello":
                                                                                                                        var v115 uint = 0
                                                                                                                        _ = v115
                                                                                                                    case "a":
                                                                                                                        var v116 float64 = v1
                                                                                                                        _ = v116
                                                                                                                    }
                                                                                                                case 4:
                                                                                                                    var v117 string = (((("" + ("a" + "")) + ((("hello" + "") + "\n\t") + "foo bar")) + ("hello" + ((("a" + (("" + "hello") + "\n\t")) + "foo bar") + ("a" + "hello")))) + "\n\t")
                                                                                                                    _ = v117
                                                                                                                case -310:
                                                                                                                    var v118 int8 = (v50 ^ (v63 / ((v50) | 1)))
                                                                                                                    _ = v118
                                                                                                                }
                                                                                                                _ = v114
                                                                                                            }
                                                                                                            _ = v112
                                                                                                        case 139:
                                                                                                            var v119 string = ("foo bar" + "hello")
                                                                                                            _ = v119
                                                                                                        default:
                                                                                                            switch "hello" {
                                                                                                            case "foo bar":
                                                                                                                v3 -= v3
                                                                                                            case "hello":
                                                                                                                fmt.Println(v3, fi75)
                                                                                                            case "\n\t":
                                                                                                                fmt.Println(v50)
                                                                                                            default:
                                                                                                                if ((v54 | v3) != v54) {
                                                                                                                    fmt.Println(v63)
                                                                                                                    if false {
                                                                                                                        fmt.Println(v53, fi75, v87)
                                                                                                                        if true {
                                                                                                                            var v120 int64 = -1
                                                                                                                            if false {
                                                                                                                                v121 := GenS3{F0: 0, F1: 89, F2: (v63 << ((150) & 63))}
                                                                                                                                var v122 int8 = 127
                                                                                                                                var v123 int8 = 127
                                                                                                                                _ = v121
                                                                                                                                _ = v122
                                                                                                                                _ = v123
                                                                                                                            }
                                                                                                                            _ = v120
                                                                                                                        } else {
                                                                                                                            v3 -= (v3 * -9223372036854775808)
                                                                                                                            v124 := GenS14{F0: 593}
                                                                                                                            fmt.Println(fi75, v87)
                                                                                                                            _ = v124
                                                                                                                        }
                                                                                                                        delete(v2, "foo bar")
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    fmt.Println(v93)
                                                                                                                    delete(v2, "\n\t")
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        var v125 int = v53
                                                                                                        _ = v125
                                                                                                    }
                                                                                                    _ = v93
                                                                                                }
                                                                                                v1 = (float64(((((v49 * float32(v1)) + v49) / float32(((v1 - v1) / 1.0))) / float32(v1))) * v1)
                                                                                                _ = v87
                                                                                            }
                                                                                            switch "" {
                                                                                            case "foo bar":
                                                                                                var v126 uint16 = 42
                                                                                                if len(v52) >= 2 {
                                                                                                    v52 = v52[1:len(v52)]
                                                                                                }
                                                                                                _ = v126
                                                                                            case "a":
                                                                                                var v127 float32 = (v49 / 1e10)
                                                                                                _ = v127
                                                                                            default:
                                                                                                var v128 uint32 = 2147483648
                                                                                                _ = v128
                                                                                            }
                                                                                            var v129 int32 = -2147483648
                                                                                            _ = v129
                                                                                        } else {
                                                                                            v5 += (v5 >> ((255) & 63))
                                                                                            switch 32768 {
                                                                                            case -9223372036854775808:
                                                                                                var v130 int8 = (v63 | (v63 - v50))
                                                                                                _ = v130
                                                                                            case 127:
                                                                                                var v131 int = v53
                                                                                                _ = v131
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    var v132 float64 = float64((((((v49 * v49) + v49) * ((v49 + 1e10) * (v49 + ((1e10 * v49) + float32(v1))))) + v49) * float32(((v1 + (1.5 / float64((((v49 + v49) - (1e10 + 1.5)) / 1e10)))) / v1))))
                                                                                    fmt.Println(v54)
                                                                                    _ = v132
                                                                                } else {
                                                                                    var v133 bool = true
                                                                                    var v134 uint64 = genFn0(true, "hello")
                                                                                    if len(v52) >= 2 {
                                                                                        v52 = v52[1:len(v52)]
                                                                                    }
                                                                                    _ = len(v2)
                                                                                    _ = v133
                                                                                    _ = v134
                                                                                }
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    _ = len(v4)
                                                                                default:
                                                                                    if true {
                                                                                        v5 += v5
                                                                                    }
                                                                                }
                                                                                var v135 int32 = -1
                                                                                _ = v135
                                                                            }
                                                                            v63 += v63
                                                                            var v136 uint = 0
                                                                            _ = v136
                                                                        }
                                                                    default:
                                                                        v4[2147483648] = 63
                                                                    }
                                                                    v5 = (^(v5))
                                                                } else {
                                                                    v5 -= (v5 ^ 1)
                                                                    v137 := GenS15{F0: 2, F1: true}
                                                                    var v138 float64 = (v1 + (v1 - (v1 / float64((v49 + 3.14159)))))
                                                                    _ = len(v52)
                                                                    var v139 uint = 0
                                                                    var v140 float32 = (v49 - v49)
                                                                    _ = v137
                                                                    _ = v138
                                                                    _ = v139
                                                                    _ = v140
                                                                }
                                                                v141 := GenS11{F0: 1}
                                                                v52 = append(v52, v53)
                                                                _ = v141
                                                            case 9223372036854775807:
                                                                var v142 float32 = v49
                                                                _ = v142
                                                            case 557:
                                                                fmt.Println(v49, v5)
                                                            }
                                                            fmt.Println(fi76, v26)
                                                            _ = fi76
                                                        }
                                                    } else {
                                                        switch "" {
                                                        case "\n\t":
                                                            switch 1 {
                                                            case 4294967295:
                                                                v2["foo bar"] = (v53 ^ (fi75 >> ((255) & 63)))
                                                            default:
                                                                v50 = (v63 >> ((65535) & 63))
                                                            }
                                                        case "a":
                                                            if len(v52) >= 2 {
                                                                v52 = v52[1:len(v52)]
                                                            }
                                                        default:
                                                            if false {
                                                                fmt.Println(v63, v54)
                                                                v3 += (v3 | v54)
                                                                if false {
                                                                    if (!(!(922 < 0))) {
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            fmt.Println(fi75, v49)
                                                                        case "a":
                                                                            v50 += v50
                                                                        case "":
                                                                            v143 := GenS7{F0: v54, F1: 2}
                                                                            _ = v143
                                                                        }
                                                                        for fi144 := 0; fi144 < 11; fi144++ {
                                                                            var v145 uint16 = 270
                                                                            v146 := GenS16{F0: 65535, F1: -2147483648}
                                                                            if true {
                                                                                v52 = append(v52, fi144)
                                                                                _ = len(v52)
                                                                            } else {
                                                                                v147 := GenS11{F0: 623}
                                                                                v54 = v3
                                                                                v148 := GenS17{F0: 0, F1: -93, F2: 1}
                                                                                _ = v147
                                                                                _ = v148
                                                                            }
                                                                            _ = fi144
                                                                            _ = v145
                                                                            _ = v146
                                                                        }
                                                                    }
                                                                    v50 -= (v63 << ((42) & 63))
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = fi75
                                            }
                                            _ = v72
                                        } else {
                                            v49 = (((v49 - float32(3.14159)) / (1.5 + float32(v1))) * v49)
                                        }
                                    }
                                    var v149 uint8 = 94
                                    v150 := GenS18{F0: ((false || (!(!true))) && ((!false) && true))}
                                    _ = v149
                                    _ = v150
                                }
                            case 9223372036854775807:
                                v151 := GenS19{F0: 1, F1: v3}
                                _ = v151
                            case -9223372036854775808:
                                switch "" {
                                case "hello":
                                    var v152 int16 = 42
                                    _ = v152
                                default:
                                    for fi153 := 0; fi153 < 6; fi153++ {
                                        if (fi153 == (v26 % ((fi153) | 1))) {
                                            for fi154 := 0; fi154 < 14; fi154++ {
                                                v5 += v5
                                                var v155 uint16 = 42
                                                _ = fi154
                                                _ = v155
                                            }
                                        } else {
                                            v156 := GenS10{F0: -2147483648, F1: 2}
                                            if false {
                                                if false {
                                                    for fi157 := 0; fi157 < 16; fi157++ {
                                                        v50 -= v50
                                                        _ = fi157
                                                    }
                                                    var v158 uint64 = genFn0(false, "\n\t")
                                                    _ = v158
                                                } else {
                                                    switch v26 {
                                                    case 4294967295:
                                                        var v159 int8 = (v50 << ((0) & 63))
                                                        _ = v159
                                                    case 42:
                                                        var v160 int8 = -93
                                                        _ = v160
                                                    case -157:
                                                        var v161 float32 = v49
                                                        _ = v161
                                                    }
                                                    for fi162 := 0; fi162 < 5; fi162++ {
                                                        v50 -= v50
                                                        var v163 uint = 2
                                                        _ = fi162
                                                        _ = v163
                                                    }
                                                    v164 := GenS17{F0: (v50 | v50), F1: -1, F2: 4294967295}
                                                    _ = v164
                                                }
                                            } else {
                                                var v165 float32 = (float32(v1) * v49)
                                                switch 32768 {
                                                case 401:
                                                    v53 = 9223372036854775807
                                                case -186:
                                                    var v166 int32 = 1
                                                    _ = v166
                                                case 0:
                                                    if true {
                                                        delete(v2, ("hello" + ("hello" + ("foo bar" + "foo bar"))))
                                                        if ((3.14159 + v1) < -2.25) {
                                                            for fi167 := 0; fi167 < 9; fi167++ {
                                                                v63 -= (v63 >> ((42) & 63))
                                                                fmt.Println(fi167)
                                                                switch ("a" + ("" + "foo bar")) {
                                                                case "hello":
                                                                    _ = len(v2)
                                                                case "a":
                                                                    _ = len(v52)
                                                                default:
                                                                    fmt.Println(fi153)
                                                                }
                                                                _ = fi167
                                                            }
                                                        }
                                                    }
                                                default:
                                                    v54 = (v3 - (v3 * -772))
                                                }
                                                switch 0 {
                                                case 1:
                                                    for fi168 := 0; fi168 < 18; fi168++ {
                                                        var v169 int32 = -994
                                                        var v170 int16 = 32767
                                                        _ = fi168
                                                        _ = v169
                                                        _ = v170
                                                    }
                                                }
                                                _ = v165
                                            }
                                            _ = v156
                                        }
                                        fmt.Println(v1, v3)
                                        delete(v2, "hello")
                                        _ = fi153
                                    }
                                }
                            }
                            var v171 int32 = 250
                            _ = v171
                        }
                    }
                    _ = v64
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        for fi172 := 0; fi172 < 14; fi172++ {
                            var v173 float32 = (float32(v1) + (v49 / 1.0))
                            var v174 int8 = (v63 << ((0) & 63))
                            var v175 int16 = 42
                            _ = fi172
                            _ = v173
                            _ = v174
                            _ = v175
                        }
                    case "hello":
                        if false {
                            var v176 bool = (1 >= 1)
                            _ = v176
                        }
                    case "foo bar":
                        var v177 int = (v26 ^ (-(v53)))
                        _ = v177
                    }
                    switch "foo bar" {
                    case "a":
                        v178 := GenS15{F0: 42, F1: true}
                        _ = v178
                    case "\n\t":
                        for fi179 := 0; fi179 < 7; fi179++ {
                            var v180 uint8 = 128
                            delete(v2, "")
                            _ = fi179
                            _ = v180
                        }
                    }
                }
            } else {
                for fi181 := 0; fi181 < 19; fi181++ {
                    var v182 uint = 4294967295
                    v183 := GenS20{F0: 533, F1: 42}
                    var v184 int8 = (^(v63))
                    _ = fi181
                    _ = v182
                    _ = v183
                    _ = v184
                }
                v26 = (-(v26))
            }
            _ = v63
        }
        var v185 string = "\n\t"
        _ = v185
    }
    v54 = (v3 + v3)
    v186 := GenS2{F0: (v5 + 0), F1: (v53 | (^(v26)))}
    if false {
        var v187 int64 = v54
        _ = v187
    }
    v188 := GenS11{F0: uint(v5)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v186
    _ = v188
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
}
