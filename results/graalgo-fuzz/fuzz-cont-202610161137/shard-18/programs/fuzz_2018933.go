package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int32; F1 uint; F2 int64 }
type GenS4 struct { F0 int16; F1 int32; F2 uint64 }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 int32; F1 int; F2 uint8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int; F1 uint64 }
type GenS10 struct { F0 uint; F1 string }
type GenS11 struct { F0 uint8; F1 int16; F2 uint32 }
func genFn0() int16 {
    v6 := []int{0, -778, 9223372036854775807, -98}
    var v7 uint64 = 2
    for fi8 := 0; fi8 < 17; fi8++ {
        switch fi8 {
        case -1:
            var v9 int16 = 555
            _ = v9
        case 32767:
            _ = len(v6)
        default:
            v10 := GenS1{F0: 2}
            _ = v10
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            if true {
                v7 = (v7 % (((v7 >> ((4294967295) & 63))) | 1))
                switch (^(fi8)) {
                case 32768:
                    if false {
                        v7 = (v7 | v7)
                        var v12 int64 = 0
                        fmt.Println(v7)
                        _ = v12
                    }
                case -9223372036854775808:
                    v13 := GenS2{F0: 0, F1: false}
                    _ = v13
                case 127:
                    if true {
                        v7 += v7
                    } else {
                        switch (^(fi11)) {
                        case 128:
                            if true {
                                if true {
                                    v7 -= 0
                                } else {
                                    v14 := GenS1{F0: 65535}
                                    _ = v14
                                }
                                var v15 int16 = 726
                                _ = v15
                            } else {
                                switch "foo bar" {
                                case "\n\t":
                                    if (false || true) {
                                        v16 := GenS3{F0: -1, F1: 42, F2: 0}
                                        if len(v6) > 0 {
                                            v6[0] = (fi11 & (fi8 >> ((42) & 63)))
                                        }
                                        _ = v16
                                    }
                                case "a":
                                    if true {
                                        var v17 int32 = -2147483648
                                        _ = v17
                                    } else {
                                        if false {
                                            v7 = (v7 % ((4294967295) | 1))
                                            v7 = v7
                                        }
                                    }
                                    if (true || (!(53 <= 42))) {
                                        if (!false) {
                                            if (!true) {
                                                switch ("hello" + "a") {
                                                case "a":
                                                    switch "foo bar" {
                                                    case "a":
                                                        var v18 string = ""
                                                        _ = v18
                                                    }
                                                case "hello":
                                                    v7 -= v7
                                                case "foo bar":
                                                    if len(v6) > 0 {
                                                        v6[0] = 318
                                                    }
                                                default:
                                                    v7 = 2
                                                }
                                                var v19 int8 = 42
                                                _ = v19
                                            } else {
                                                var v20 int64 = 164
                                                v21 := GenS4{F0: 1000, F1: 2147483647, F2: v7}
                                                var v22 uint32 = 769
                                                if (467 >= -408) {
                                                    if true {
                                                        var v23 int = fi8
                                                        v20 -= v20
                                                        v23 += 128
                                                        var v24 uint32 = v22
                                                        _ = v23
                                                        _ = v24
                                                    }
                                                }
                                                v7 += 907
                                                if (true || (false || (v7 != (^(v7))))) {
                                                    var v25 uint16 = 2
                                                    v20 += v20
                                                    if (("" + "a") > "\n\t") {
                                                        if false {
                                                            v25 = (v25 * v25)
                                                            var v26 uint64 = (v7 | 0)
                                                            _ = v26
                                                        } else {
                                                            var v27 bool = true
                                                            v28 := GenS5{F0: (((v25 << ((987) & 63)) != (v25 << ((393) & 63))) && false), F1: int32(v20)}
                                                            if len(v6) > 0 {
                                                                v6[0] = 972
                                                            }
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            _ = v27
                                                            _ = v28
                                                        }
                                                        fmt.Println(v25, v20, v20)
                                                    } else {
                                                        var v29 string = ("a" + "a")
                                                        var v30 uint8 = 128
                                                        _ = v29
                                                        _ = v30
                                                    }
                                                    _ = v25
                                                } else {
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    v22 -= v22
                                                }
                                                var v31 int = 1
                                                var v32 uint64 = (v7 % ((v7) | 1))
                                                _ = v20
                                                _ = v21
                                                _ = v22
                                                _ = v31
                                                _ = v32
                                            }
                                            v33 := GenS2{F0: 2, F1: true}
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            _ = v33
                                        } else {
                                            v7 -= (v7 / ((42) | 1))
                                        }
                                        var v34 int = (fi11 << ((286) & 63))
                                        fmt.Println(v34, v7, fi11)
                                        _ = v34
                                    } else {
                                        if (true && ("foo bar" <= "hello")) {
                                            var v35 float64 = (0.1 - 0.1)
                                            switch (fi8 * -100) {
                                            case -9223372036854775808:
                                                fmt.Println(fi8, fi8)
                                            case 4294967295:
                                                var v36 int = -9223372036854775808
                                                _ = v36
                                            case 9223372036854775807:
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            default:
                                                if ((0 == 42) || (1e10 > v35)) {
                                                    var v37 uint32 = 65535
                                                    v37 = (v37 + v37)
                                                    _ = v37
                                                } else {
                                                    fmt.Println(fi11, v7, v35)
                                                }
                                            }
                                            _ = v35
                                        }
                                        if true {
                                            var v38 uint16 = 2
                                            _ = v38
                                        } else {
                                            var v39 int16 = -32768
                                            _ = v39
                                        }
                                        var v40 uint16 = 598
                                        _ = v40
                                    }
                                    _ = len(v6)
                                default:
                                    switch fi8 {
                                    case 456:
                                        switch "a" {
                                        case "a":
                                            if false {
                                                if (255 <= 128) {
                                                    if true {
                                                        v41 := GenS5{F0: true, F1: -25}
                                                        _ = v41
                                                    }
                                                }
                                                v42 := GenS6{F0: -1, F1: (fi11 & -311), F2: 2}
                                                _ = v42
                                            } else {
                                                if true {
                                                    v7 = (v7 | (v7 & 42))
                                                } else {
                                                    if (("\n\t" + "a") > "a") {
                                                        v7 = 0
                                                        fmt.Println(fi11)
                                                    }
                                                    if false {
                                                        fmt.Println(fi11)
                                                        if (42 == 2147483647) {
                                                            v6 = append(v6, fi11)
                                                            v43 := GenS2{F0: 157, F1: (v7 <= (v7 << ((42) & 63)))}
                                                            _ = v43
                                                        }
                                                        var v44 float32 = 1e10
                                                        var v45 string = "\n\t"
                                                        var v46 uint8 = 1
                                                        _ = v44
                                                        _ = v45
                                                        _ = v46
                                                    } else {
                                                        var v47 int8 = -114
                                                        _ = v47
                                                    }
                                                    v7 = v7
                                                }
                                                var v48 uint16 = 551
                                                _ = v48
                                            }
                                        case "hello":
                                            var v49 string = "\n\t"
                                            _ = v49
                                        }
                                    case 32768:
                                        if (false || false) {
                                            var v50 int16 = 786
                                            _ = v50
                                        }
                                    case 1:
                                        switch ("a" + "hello") {
                                        case "foo bar":
                                            var v51 int8 = 127
                                            _ = v51
                                        }
                                    default:
                                        var v52 uint = 584
                                        _ = v52
                                    }
                                }
                            }
                        case -581:
                            var v53 uint32 = 2147483647
                            _ = v53
                        case 32768:
                            v7 += 18446744073709551615
                        }
                    }
                default:
                    v7 += (v7 - (v7 >> ((65535) & 63)))
                }
            } else {
                v7 = v7
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                if ((v7 ^ 4294967295) >= 18446744073709551615) {
                    v7 += 979
                }
            }
            switch "foo bar" {
            case "foo bar":
                switch "a" {
                case "a":
                    var v54 uint16 = 42
                    _ = v54
                default:
                    v55 := GenS2{F0: 6, F1: false}
                    _ = v55
                }
            }
            var v56 int16 = 263
            if true {
                var v57 uint8 = 127
                if false {
                    v56 = v56
                    if true {
                        v58 := GenS1{F0: 65535}
                        _ = v58
                    } else {
                        fmt.Println(fi8, fi8, v57)
                    }
                }
                if (-1 <= 42) {
                    v56 -= v56
                    v7 = v7
                    var v59 int32 = 0
                    _ = v59
                } else {
                    var v60 int8 = -125
                    v61 := GenS7{F0: 0}
                    var v62 uint = 2
                    _ = v60
                    _ = v61
                    _ = v62
                }
                if true {
                    if true {
                        v57 = v57
                        v57 = (v57 ^ (v57 ^ (v57 & 127)))
                        var v63 uint64 = v7
                        _ = v63
                    }
                    v6 = append(v6, -1)
                    fmt.Println(fi11)
                } else {
                    v64 := GenS8{F0: "\n\t"}
                    if true {
                        var v65 int32 = -2147483648
                        var v66 uint = 65535
                        v6 = append(v6, fi8)
                        _ = v65
                        _ = v66
                    }
                    if ((((85 > 704) || (127 == 100)) || true) && false) {
                        switch "hello" {
                        case "":
                            switch fi8 {
                            case -349:
                                if false {
                                    var v67 float64 = 3.14159
                                    v56 = (v56 % ((v56) | 1))
                                    switch fi11 {
                                    case 1:
                                        switch "\n\t" {
                                        case "\n\t":
                                            v7 = 2
                                        }
                                    }
                                    _ = v67
                                }
                            case 127:
                                v7 += v7
                            case 128:
                                var v68 uint16 = uint16(v7)
                                _ = v68
                            default:
                                var v69 float64 = 0.1
                                _ = v69
                            }
                        case "foo bar":
                            var v70 uint = 1
                            _ = v70
                        default:
                            var v71 uint8 = 128
                            _ = v71
                        }
                        v57 -= (v57 << ((0) & 63))
                    } else {
                        var v72 uint16 = 65535
                        _ = v72
                    }
                    _ = v64
                }
                var v73 bool = true
                switch (-(fi8)) {
                case -9223372036854775808:
                    var v74 uint = uint(v7)
                    _ = v74
                case -571:
                    if (!(2 == 192)) {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        v56 -= 42
                        var v75 uint16 = 1
                        _ = v75
                    }
                case 953:
                    var v76 int32 = 363
                    _ = v76
                }
                v77 := GenS9{F0: 697, F1: v7}
                _ = v57
                _ = v73
                _ = v77
            }
            if true {
                v78 := GenS2{F0: 469, F1: (v7 == (^(v7)))}
                v7 -= v7
                _ = v78
            } else {
                if false {
                    if len(v6) > 0 {
                        v6[0] = 128
                    }
                    v79 := GenS4{F0: (-(v56)), F1: -2147483648, F2: v7}
                    fmt.Println(fi11, v7)
                    _ = v79
                }
                if len(v6) > 0 {
                    v6[0] = -227
                }
            }
            var v80 float32 = 0.1
            v81 := GenS4{F0: (v56 + v56), F1: -1, F2: v7}
            _ = fi11
            _ = v56
            _ = v80
            _ = v81
        }
        _ = fi8
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v7 = (v7 ^ 20)
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint8, p1 uint) int8 {
    var v82 int16 = -32768
    var v83 int = -2147483648
    for fi84 := 0; fi84 < 18; fi84++ {
        v85 := GenS8{F0: ("foo bar" + "")}
        _ = fi84
        _ = v85
    }
    v86 := []int{v83}
    fmt.Println(p1, v82)
    _ = p0
    _ = p1
    _ = v82
    _ = v83
    _ = v86
    return 0
}

func main() {
    var v1 float64 = 0.1
    v2 := []int{1, 128, 32767, 42}
    var v3 string = ""
    var v4 string = ""
    v5 := []int{-286}
    switch "hello" {
    case "\n\t":
        if (564 > 575) {
            switch 0 {
            case 32768:
                v4 = v4
            case -9223372036854775808:
                switch "\n\t" {
                case "hello":
                    switch -1 {
                    case 565:
                        fmt.Println(v4)
                    default:
                        if true {
                            _ = len(v2)
                            v2 = append(v2, 0)
                            v4 = ((v4 + v4) + "a")
                        } else {
                            for fi87 := 0; fi87 < 4; fi87++ {
                                var v88 uint32 = 0
                                var v89 uint64 = 1
                                _ = fi87
                                _ = v88
                                _ = v89
                            }
                            var v90 uint8 = 128
                            _ = v90
                        }
                    }
                case "a":
                    if (-2147483648 != -97) {
                        for fi91 := 0; fi91 < 18; fi91++ {
                            if (((!(true || false)) || (0 <= 2147483647)) && true) {
                                var v92 float64 = (1e10 / v1)
                                _ = v92
                            } else {
                                switch 127 {
                                case 535:
                                    var v93 uint64 = 861
                                    _ = v93
                                case 42:
                                    switch "\n\t" {
                                    case "a":
                                        v94 := GenS3{F0: 128, F1: 65535, F2: -9223372036854775808}
                                        _ = v94
                                    default:
                                        for fi95 := 0; fi95 < 17; fi95++ {
                                            v96 := GenS2{F0: 0, F1: (0 > 2)}
                                            switch (-(fi95)) {
                                            case 32767:
                                                v4 = v3
                                            }
                                            _ = fi95
                                            _ = v96
                                        }
                                    }
                                default:
                                    var v97 float64 = (float64(0.0) * v1)
                                    _ = v97
                                }
                                v4 = v3
                            }
                            for fi98 := 0; fi98 < 19; fi98++ {
                                switch v4 {
                                case "a":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                var v99 int64 = -1
                                _ = fi98
                                _ = v99
                            }
                            for fi100 := 0; fi100 < 19; fi100++ {
                                switch v3 {
                                case "foo bar":
                                    if false {
                                        fmt.Println(v4)
                                    } else {
                                        v4 = ""
                                    }
                                default:
                                    v1 = v1
                                }
                                var v101 int = 32767
                                _ = fi100
                                _ = v101
                            }
                            _ = fi91
                        }
                    }
                }
            case 899:
                v102 := GenS5{F0: false, F1: -355}
                _ = v102
            }
            v4 = ""
            fmt.Println(v3)
        } else {
            var v103 int = -9223372036854775808
            _ = v103
        }
    case "a":
        fmt.Println(v4, v1, v3)
    }
    v4 = (v3 + v4)
    if false {
        var v104 int32 = 2147483647
        _ = len(v2)
        _ = v104
    }
    _ = len(v5)
    v2 = append(v2, 127)
    var v105 uint8 = 2
    var v106 uint = 4294967295
    var v107 int64 = 42
    var v108 int32 = 42
    for fi109 := 0; fi109 < 11; fi109++ {
        v107 -= v107
        switch -9223372036854775808 {
        case -748:
            for fi110 := 0; fi110 < 13; fi110++ {
                switch 127 {
                case 4294967295:
                    var v111 float32 = (1e10 / 1.5)
                    _ = v111
                case 32768:
                    v4 = (v4 + (("foo bar" + v3) + ("" + v4)))
                case 2147483647:
                    if false {
                        v112 := GenS2{F0: 232, F1: ((!true) && false)}
                        _ = v112
                    } else {
                        if len(v2) > 0 {
                            v2[0] = fi110
                        }
                        if false {
                            var v113 int8 = 54
                            _ = v113
                        } else {
                            fmt.Println(v108, fi109, fi109)
                            if false {
                                var v114 float64 = float64(0.0)
                                switch (v3 + (v3 + "foo bar")) {
                                case "a":
                                    var v115 int32 = v108
                                    _ = v115
                                }
                                _ = v114
                            }
                            if (fi110 > fi110) {
                                if len(v5) > 0 {
                                    v5[0] = (fi109 / ((fi109) | 1))
                                }
                            } else {
                                var v116 uint = (v106 % (((v106 - 674)) | 1))
                                v108 = int32(v116)
                                _ = v116
                            }
                        }
                        v4 = v4
                        switch fi109 {
                        case 639:
                            switch (v3 + (v4 + v3)) {
                            case "":
                                if (false && (false && (1 <= 42))) {
                                    var v117 uint = 0
                                    _ = v117
                                } else {
                                    v106 -= 936
                                    var v118 float64 = v1
                                    _ = v118
                                }
                            case "hello":
                                var v119 uint8 = 1
                                _ = v119
                            default:
                                v3 = v4
                            }
                        case -9223372036854775808:
                            if false {
                                var v120 uint = (v106 | v106)
                                if false {
                                    v3 = v4
                                    v5 = append(v5, fi110)
                                    fmt.Println(v1, v106)
                                    v121 := GenS8{F0: (v3 + (("hello" + (v3 + v4)) + v3))}
                                    var v122 uint = v120
                                    _ = v121
                                    _ = v122
                                } else {
                                    var v123 uint64 = 18446744073709551615
                                    v1 -= v1
                                    v124 := GenS8{F0: ""}
                                    if (!true) {
                                        var v125 uint64 = v123
                                        _ = v125
                                    } else {
                                        var v126 float64 = (-2.25 / v1)
                                        v127 := GenS10{F0: v120, F1: v3}
                                        _ = v126
                                        _ = v127
                                    }
                                    var v128 uint32 = 2
                                    _ = v123
                                    _ = v124
                                    _ = v128
                                }
                                _ = v120
                            } else {
                                if (v1 == (v1 + 1.5)) {
                                    v129 := GenS11{F0: v105, F1: genFn0(), F2: 2}
                                    _ = v129
                                }
                                v106 = 105
                            }
                        case -198:
                            var v130 uint16 = 0
                            _ = v130
                        default:
                            v1 -= v1
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                default:
                    switch (v4 + v3) {
                    case "":
                        v108 = (v108 ^ 797)
                    default:
                        var v131 bool = false
                        _ = v131
                    }
                }
                var v132 int32 = -1
                _ = fi110
                _ = v132
            }
        }
        _ = fi109
    }
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("v107:", v107)
    fmt.Println("v108:", v108)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v105
    _ = v106
    _ = v107
    _ = v108
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
