package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 int64; F1 uint }
type GenS9 struct { F0 int16; F1 uint8 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint64; F1 uint }
type GenS14 struct { F0 string; F1 int16; F2 uint }
type GenS15 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS16 struct { F0 uint32; F1 bool }
type GenS17 struct { F0 uint64 }
func genFn0(p0 int32) uint {
    if (-20 > -891) {
        p0 += 2147483647
    } else {
        v4 := GenS1{F0: 65535}
        if true {
            v5 := []int{32768, -2147483648, 1, 993}
            v6 := GenS2{F0: 13}
            _ = v5
            _ = v6
        } else {
            for fi7 := 0; fi7 < 6; fi7++ {
                var v8 int8 = 127
                var v9 int8 = (v8 / ((v8) | 1))
                _ = fi7
                _ = v8
                _ = v9
            }
            if (true || false) {
                if (198 > 0) {
                    var v10 float64 = (3.14159 - -2.25)
                    _ = v10
                } else {
                    for fi11 := 0; fi11 < 12; fi11++ {
                        var v12 int = -748
                        _ = fi11
                        _ = v12
                    }
                }
                p0 += -385
                v13 := GenS2{F0: 214}
                _ = v13
            } else {
                fmt.Println(p0)
            }
            var v14 bool = (false && (448 < 0))
            _ = v14
        }
        v15 := GenS2{F0: 42}
        _ = v15
        _ = v4
    }
    v16 := []int{639}
    var v17 uint32 = 2
    _ = p0
    _ = v16
    _ = v17
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 1.5
    var v3 int16 = -1
    v3 = v3
    fmt.Println(v3)
    switch "" {
    case "hello":
        v18 := GenS2{F0: 214}
        _ = v18
    case "a":
        v19 := GenS2{F0: 0}
        _ = v19
    default:
        v20 := GenS3{F0: 1, F1: -1, F2: 145}
        _ = v20
    }
    if false {
        v2 -= (v2 * v2)
    } else {
        fmt.Println(v2)
    }
    v21 := GenS3{F0: 0, F1: -735, F2: 32767}
    var v22 bool = true
    for fi23 := 0; fi23 < 17; fi23++ {
        if false {
            v24 := []int{fi23, fi23, 4294967295, fi23}
            switch ((("\n\t" + "a") + ("foo bar" + "")) + "a") {
            case "":
                for fi25 := 0; fi25 < 8; fi25++ {
                    v26 := GenS2{F0: 27}
                    var v27 uint32 = 0
                    var v28 int32 = -2147483648
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                }
            case "hello":
                switch "hello" {
                case "a":
                    for fi29 := 0; fi29 < 18; fi29++ {
                        switch "\n\t" {
                        case "foo bar":
                            if (0.0 != (1e10 - 0.0)) {
                                if len(v24) >= 2 {
                                    v24 = v24[1:len(v24)]
                                }
                                switch "hello" {
                                case "a":
                                    if ("\n\t" < "foo bar") {
                                        var v30 int64 = 0
                                        switch -9223372036854775808 {
                                        case 732:
                                            var v31 uint32 = 875
                                            _ = v31
                                        case 32768:
                                            switch (fi23 * fi29) {
                                            case 2147483647:
                                                var v32 int64 = v30
                                                _ = v32
                                            default:
                                                _ = len(v24)
                                                var v33 uint32 = 4294967295
                                                _ = v33
                                            }
                                        }
                                        _ = v30
                                    }
                                case "":
                                    v34 := GenS1{F0: 494}
                                    _ = v34
                                }
                                if false {
                                    v35 := GenS4{F0: 42, F1: true}
                                    _ = v35
                                }
                            }
                        case "\n\t":
                            delete(v1, (fi23 ^ fi23))
                            _ = len(v1)
                            if (0.0 >= 0.1) {
                                v36 := GenS5{F0: 42, F1: v22}
                                v2 += (float64(0.1) + (((3.14159 - 1e10) + 1e10) - ((1e10 / (0.0 * v2)) + 3.14159)))
                                v22 = (!true)
                                _ = len(v24)
                                _ = v36
                            }
                        case "a":
                            v2 += 1.5
                        }
                        if (v22 && true) {
                            v2 = 1.5
                            v37 := v1[127]
                            v3 -= (v3 - (v3 | (^(v3))))
                            _ = v37
                        }
                        var v38 uint8 = 127
                        v2 = float64(0.0)
                        fmt.Println(v2, fi29)
                        _ = fi29
                        _ = v38
                    }
                case "hello":
                    if (!(!v22)) {
                        _ = len(v24)
                        for fi39 := 0; fi39 < 14; fi39++ {
                            v3 = v3
                            var v40 float32 = 0.1
                            _ = fi39
                            _ = v40
                        }
                    }
                case "":
                    v41 := GenS3{F0: 83, F1: 2147483647, F2: 255}
                    _ = v41
                default:
                    if (fi23 != (fi23 / ((fi23) | 1))) {
                        delete(v1, (fi23 ^ fi23))
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                        var v42 bool = false
                        _ = v42
                    }
                }
            }
            _ = v24
        }
        var v43 int = (fi23 / ((fi23) | 1))
        if (2147483647 == 42) {
            var v44 int32 = -750
            v45 := GenS1{F0: 4294967295}
            _ = v44
            _ = v45
        } else {
            if v22 {
                v1[(v43 | (fi23 | fi23))] = fi23
            } else {
                if v22 {
                    var v46 uint64 = 18446744073709551615
                    _ = v46
                } else {
                    if (true && v22) {
                        fmt.Println(v2, fi23)
                        v47 := []int{(^(v43)), (v43 % (((fi23 / ((-746) | 1))) | 1)), (fi23 - 4294967295), (fi23 / ((794) | 1))}
                        fmt.Println(v22, v3)
                        _ = v47
                    } else {
                        if v22 {
                            if (-128 < 127) {
                                if true {
                                    v48 := GenS2{F0: 2}
                                    _ = v48
                                } else {
                                    v22 = v22
                                    for fi49 := 0; fi49 < 3; fi49++ {
                                        if true {
                                            if false {
                                                if v22 {
                                                    if ((-92 >= 0) && false) {
                                                        if v22 {
                                                            v50 := []int{(^(v43)), (fi49 * fi49)}
                                                            fmt.Println(fi23)
                                                            _ = v50
                                                        } else {
                                                            switch ((("foo bar" + ("foo bar" + "hello")) + ((("\n\t" + "a") + ((("" + ("foo bar" + "foo bar")) + "\n\t") + "hello")) + ("" + "foo bar"))) + ("foo bar" + "a")) {
                                                            case "foo bar":
                                                                v22 = v22
                                                            case "hello":
                                                                if v22 {
                                                                    v2 += ((float64(0.1) * (0.1 + v2)) - v2)
                                                                    fmt.Println(fi23)
                                                                    v51 := GenS4{F0: 128, F1: true}
                                                                    var v52 int = (-(v43))
                                                                    v43 = v52
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                            case "\n\t":
                                                                var v53 uint64 = 397
                                                                _ = v53
                                                            default:
                                                                if (float32(0.0) >= float32((((float64(0.1) + 0.1) / ((((0.0 / (v2 - v2)) + (v2 * 0.1)) / float64((0.1 * 1.5))) * (v2 - (1e10 - v2)))) + (1.5 / (3.14159 + 0.0))))) {
                                                                    v54 := v1[v43]
                                                                    _ = v54
                                                                } else {
                                                                    v22 = v22
                                                                }
                                                            }
                                                        }
                                                        switch (fi49 / ((fi49) | 1)) {
                                                        case 127:
                                                            v55 := GenS6{F0: 189, F1: 939}
                                                            _ = v55
                                                        case 580:
                                                            fmt.Println(v3, fi23, v3)
                                                        case -1:
                                                            if (987 != 2) {
                                                                switch (-(fi49)) {
                                                                case 1:
                                                                    var v56 uint32 = 0
                                                                    _ = v56
                                                                case -1:
                                                                    var v57 int = 4294967295
                                                                    _ = v57
                                                                case 42:
                                                                    var v58 uint16 = 42
                                                                    _ = v58
                                                                default:
                                                                    _ = len(v1)
                                                                }
                                                                fmt.Println(v43)
                                                            }
                                                        }
                                                        v59 := []int{(fi49 >> ((2) & 63)), 32767, (^(fi23)), 4294967295}
                                                        _ = v59
                                                    } else {
                                                        if v22 {
                                                            v60 := GenS5{F0: -668, F1: v22}
                                                            if v22 {
                                                                if v22 {
                                                                    v43 -= (fi49 | int(v3))
                                                                } else {
                                                                    v61 := []int{(fi23 * fi49)}
                                                                    v62 := GenS7{F0: -306, F1: 1}
                                                                    _ = v61
                                                                    _ = v62
                                                                }
                                                                var v63 int8 = int8(v3)
                                                                _ = v63
                                                            }
                                                            fmt.Println(v3, v22)
                                                            _ = v60
                                                        } else {
                                                            var v64 uint64 = 18446744073709551615
                                                            v22 = (((129 <= -9223372036854775808) || (2147483648 >= 4)) || (!true))
                                                            _ = v64
                                                        }
                                                        v65 := GenS8{F0: 32, F1: 0}
                                                        switch (fi23 * 32768) {
                                                        case 32767:
                                                            v43 = -9223372036854775808
                                                        }
                                                        _ = v65
                                                    }
                                                    var v66 uint64 = 1
                                                    v3 += 312
                                                    _ = v66
                                                }
                                                fmt.Println(v3, v22)
                                            }
                                        }
                                        _ = fi49
                                    }
                                }
                                var v67 bool = (2 < 230)
                                switch ("foo bar" + "foo bar") {
                                case "":
                                    v67 = (false && v22)
                                case "\n\t":
                                    delete(v1, -9223372036854775808)
                                default:
                                    v68 := GenS1{F0: 2147483648}
                                    _ = v68
                                }
                                _ = v67
                            } else {
                                var v69 int64 = 42
                                _ = v69
                            }
                            v70 := GenS9{F0: -32768, F1: 225}
                            v1[(fi23 | -269)] = (v43 & fi23)
                            _ = v70
                        }
                    }
                }
            }
            if (v22 && (0 > genFn0(-1))) {
                v71 := GenS1{F0: 2}
                var v72 uint32 = 175
                var v73 uint8 = 191
                _ = v71
                _ = v72
                _ = v73
            } else {
                var v74 int32 = 0
                v43 = (v43 << ((genFn0(v74)) & 63))
                v2 -= ((-2.25 / -2.25) - float64(1.5))
                _ = v74
            }
        }
        _ = fi23
        _ = v43
    }
    switch 127 {
    case -9223372036854775808:
        if (!true) {
            for fi75 := 0; fi75 < 10; fi75++ {
                var v76 uint16 = 375
                v1[(fi75 >> ((156) & 63))] = -1
                _ = fi75
                _ = v76
            }
            v2 = v2
            v77 := GenS10{F0: 177}
            _ = v77
        } else {
            fmt.Println(v2, v2)
            if v22 {
                fmt.Println(v2, v2)
            } else {
                var v78 uint16 = 2
                v79 := []int{9223372036854775807, int(v3), -9223372036854775808, -181}
                if true {
                    v80 := GenS11{F0: v22}
                    v81 := GenS12{F0: v78}
                    _ = v80
                    _ = v81
                } else {
                    var v82 int = 32768
                    if ((!v22) || (!(false || false))) {
                        var v83 int16 = (-(v3))
                        _ = v83
                    }
                    _ = v82
                }
                _ = v78
                _ = v79
            }
        }
    }
    v84 := []int{-1, 0, 0}
    v85 := GenS13{F0: 2, F1: 180}
    v86 := v1[-9223372036854775808]
    v86 += v86
    for fi87 := 0; fi87 < 19; fi87++ {
        if len(v84) >= 2 {
            v84 = v84[1:len(v84)]
        }
        v88 := GenS14{F0: (("" + "\n\t") + ""), F1: v3, F2: 479}
        _ = fi87
        _ = v88
    }
    _ = len(v84)
    if v22 {
        v89 := GenS15{F0: 4294967295, F1: 2147483647, F2: 1}
        v84 = append(v84, v86)
        var v90 int64 = -9223372036854775808
        _ = v89
        _ = v90
    } else {
        switch 42 {
        case 32767:
            var v91 uint = 4294967295
            _ = v91
        case 32768:
            var v92 uint16 = 477
            _ = v92
        case 127:
            fmt.Println(v2, v86, v3)
        }
        v2 -= v2
    }
    if true {
        for fi93 := 0; fi93 < 9; fi93++ {
            if v22 {
                var v94 int32 = 2147483647
                var v95 float64 = v2
                _ = v94
                _ = v95
            } else {
                if true {
                    v96 := GenS3{F0: 128, F1: -1, F2: 1}
                    v86 -= 1
                    if ((((0.0 + 0.1) - 1e10) > (1.5 + -2.25)) || ((v86 % ((v86) | 1)) == fi93)) {
                        for fi97 := 0; fi97 < 18; fi97++ {
                            switch "a" {
                            case "hello":
                                if (v22 || true) {
                                    fmt.Println(fi93, v3)
                                } else {
                                    var v98 uint = 282
                                    v3 = v3
                                    v3 -= (v3 | v3)
                                    _ = v98
                                }
                            case "foo bar":
                                if true {
                                    v3 += 1000
                                    switch "a" {
                                    case "hello":
                                        fmt.Println(fi93)
                                        var v99 uint16 = 1
                                        v100 := GenS16{F0: 65535, F1: ((fi93 >> ((0) & 63)) > v86)}
                                        switch "\n\t" {
                                        case "a":
                                            v101 := GenS3{F0: 43, F1: 1, F2: 32767}
                                            _ = v101
                                        }
                                        var v102 uint32 = uint32(v99)
                                        _ = v100
                                        _ = v102
                                        _ = v99
                                    default:
                                        var v103 uint64 = 2
                                        _ = v103
                                    }
                                } else {
                                    if v22 {
                                        var v104 int16 = v3
                                        fmt.Println(fi93, fi93, v3)
                                        _ = v104
                                    } else {
                                        fmt.Println(v2, v2, fi97)
                                        var v105 int64 = -904
                                        _ = v105
                                    }
                                }
                            default:
                                var v106 bool = (!((!(1 != 103)) && v22))
                                _ = v106
                            }
                            var v107 uint8 = 174
                            _ = fi97
                            _ = v107
                        }
                        var v108 uint32 = 627
                        _ = v108
                    }
                    _ = v96
                } else {
                    delete(v1, fi93)
                    fmt.Println(v86, v3)
                    if v22 {
                        var v109 uint = genFn0(792)
                        fmt.Println(v109)
                        _ = v109
                    }
                }
            }
            if len(v84) >= 2 {
                v84 = v84[1:len(v84)]
            }
            switch v86 {
            case 1:
                v86 -= 2147483647
            }
            _ = fi93
        }
        for fi110 := 0; fi110 < 7; fi110++ {
            for fi111 := 0; fi111 < 9; fi111++ {
                switch fi111 {
                case 4294967295:
                    switch ("" + "a") {
                    case "foo bar":
                        if ((1e10 * 1.5) > -2.25) {
                            if (!v22) {
                                var v112 int32 = 42
                                var v113 uint16 = 0
                                v114 := GenS17{F0: 18446744073709551615}
                                _ = v112
                                _ = v113
                                _ = v114
                            } else {
                                var v115 int = fi110
                                if (-32768 >= (v3 >> ((4294967295) & 63))) {
                                    fmt.Println(v3, fi110)
                                }
                                switch ("a" + "") {
                                case "":
                                    var v116 bool = v22
                                    _ = v116
                                case "foo bar":
                                    var v117 float32 = float32(0.1)
                                    _ = v117
                                default:
                                    v1[v86] = fi110
                                }
                                _ = v115
                            }
                        }
                    case "\n\t":
                        v2 += 0.0
                    }
                default:
                    v118 := v1[fi110]
                    switch (fi110 | (fi110 / ((fi111) | 1))) {
                    case 1:
                        fmt.Println(v86)
                        v84 = append(v84, fi110)
                        if false {
                            switch (fi110 % ((-1) | 1)) {
                            case 9223372036854775807:
                                v119 := GenS1{F0: 65535}
                                if v22 {
                                    v2 -= (v2 - -2.25)
                                }
                                switch ("\n\t" + ("" + "foo bar")) {
                                case "hello":
                                    v118 = -793
                                }
                                _ = v119
                            case -2147483648:
                                v2 = (float64(1e10) + 1.5)
                            case 1:
                                v84 = append(v84, int(v3))
                            }
                            _ = len(v84)
                        }
                    case 128:
                        v2 = 0.0
                    case -2147483648:
                        if (!true) {
                            v120 := GenS5{F0: 762, F1: false}
                            _ = v120
                        }
                    }
                    fmt.Println(v118, fi111, v86)
                    _ = v118
                }
                var v121 float64 = (((v2 * float64(((0.0 * (3.14159 - float32(v2))) * 0.1))) * v2) - -2.25)
                v84 = append(v84, fi111)
                _ = fi111
                _ = v121
            }
            if (652 > 251) {
                fmt.Println(fi110, v86)
            }
            _ = fi110
        }
        var v122 int16 = 343
        _ = v122
    } else {
        v123 := GenS7{F0: 9223372036854775807, F1: 207}
        _ = v123
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v84:", len(v84), v84)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v84
    _ = v85
    _ = v86
}
