package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int; F1 int32; F2 string }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 bool; F1 uint64; F2 uint }
type GenS7 struct { F0 int32; F1 uint; F2 uint }
type GenS8 struct { F0 int; F1 string }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int64; F1 int }
type GenS12 struct { F0 int8; F1 uint }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint64; F1 uint32; F2 int32 }
func genFn0(p0 uint8) int32 {
    for fi6 := 0; fi6 < 7; fi6++ {
        p0 += p0
        for fi7 := 0; fi7 < 2; fi7++ {
            switch (fi6 + (-(fi6))) {
            case 378:
                var v8 bool = true
                _ = v8
            default:
                v9 := []int{(fi6 * fi7), fi6, 0}
                _ = v9
            }
            switch 0 {
            case -9223372036854775808:
                switch fi7 {
                case 128:
                    fmt.Println(p0)
                case 42:
                    fmt.Println(fi7, p0, p0)
                default:
                    p0 -= (p0 % ((255) | 1))
                }
            case 964:
                if false {
                    if true {
                        v10 := GenS1{F0: (-1 <= 1), F1: 42}
                        var v11 int = 1
                        v11 = fi7
                        _ = v10
                        _ = v11
                    } else {
                        var v12 int = -82
                        if true {
                            v12 = 2147483647
                            if ((false && true) && (234 >= 202)) {
                                v12 -= v12
                                if true {
                                    var v13 uint8 = (^(p0))
                                    _ = v13
                                }
                                switch 42 {
                                case -871:
                                    switch "hello" {
                                    case "foo bar":
                                        v14 := GenS1{F0: ((-273 <= 0) || (!false)), F1: 0}
                                        _ = v14
                                    default:
                                        v12 += 32767
                                    }
                                case 0:
                                    v12 -= fi6
                                default:
                                    fmt.Println(v12, p0)
                                }
                            } else {
                                if false {
                                    var v15 uint = 65535
                                    var v16 int32 = 42
                                    var v17 bool = (false && ("\n\t" > "a"))
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                }
                                v18 := []int{987, fi7, v12, fi6}
                                p0 = (p0 % ((128) | 1))
                                v12 = (fi6 << ((42) & 63))
                                if (false && (true && ((!false) && (!true)))) {
                                    switch ("foo bar" + ("a" + ("foo bar" + "\n\t"))) {
                                    case "foo bar":
                                        v18 = append(v18, (^(v12)))
                                    case "hello":
                                        v12 = -392
                                    case "":
                                        p0 += 40
                                        var v19 int16 = 1
                                        _ = v19
                                    default:
                                        p0 += p0
                                        var v20 uint32 = 65535
                                        switch "hello" {
                                        case "hello":
                                            if false {
                                                v20 = v20
                                            }
                                        case "\n\t":
                                            v21 := GenS1{F0: true, F1: 0}
                                            _ = v21
                                        case "foo bar":
                                            v20 += 2
                                        default:
                                            v22 := GenS2{F0: -105, F1: -2147483648, F2: (("hello" + "\n\t") + "")}
                                            _ = v22
                                        }
                                        var v23 int32 = 2147483647
                                        _ = v20
                                        _ = v23
                                    }
                                } else {
                                    v24 := GenS3{F0: 42, F1: 103}
                                    var v25 uint16 = 1
                                    _ = v24
                                    _ = v25
                                }
                                if len(v18) > 0 {
                                    v18[0] = (fi6 << ((1) & 63))
                                }
                                if (-2147483648 != 1) {
                                    if ((fi6 ^ (^(v12))) > fi6) {
                                        var v26 int32 = 854
                                        if false {
                                            var v27 uint64 = 0
                                            _ = v27
                                        }
                                        _ = v26
                                    }
                                    if (false && false) {
                                        if (-2147483648 == -1) {
                                            v12 += (^(v12))
                                        } else {
                                            switch "hello" {
                                            case "a":
                                                v12 = fi6
                                            case "\n\t":
                                                if (false && true) {
                                                    if true {
                                                        p0 += 42
                                                        v12 = 32768
                                                    } else {
                                                        fmt.Println(fi6)
                                                        var v28 bool = true
                                                        _ = v28
                                                    }
                                                    fmt.Println(fi7, v12)
                                                    p0 = (p0 ^ 0)
                                                }
                                                v12 -= fi6
                                                var v29 uint = 1
                                                _ = v29
                                            }
                                            v30 := GenS3{F0: 42, F1: 226}
                                            var v31 float64 = 1e10
                                            _ = v30
                                            _ = v31
                                        }
                                        v12 += (v12 | 202)
                                    } else {
                                        v18 = append(v18, (fi6 | (fi6 & (v12 * fi7))))
                                        if ((1.5 * 0.0) <= 1.5) {
                                            v32 := GenS2{F0: -2147483648, F1: -312, F2: "hello"}
                                            var v33 int32 = 442
                                            var v34 float32 = 3.14159
                                            _ = v32
                                            _ = v33
                                            _ = v34
                                        } else {
                                            v35 := GenS1{F0: (0.1 >= (0.1 * -2.25)), F1: 42}
                                            _ = len(v18)
                                            _ = v35
                                        }
                                        fmt.Println(v12, fi6)
                                    }
                                    var v36 int = fi7
                                    var v37 int8 = -74
                                    p0 += (p0 - 144)
                                    p0 = p0
                                    _ = len(v18)
                                    switch (fi6 / ((127) | 1)) {
                                    case 42:
                                        if false {
                                            var v38 int8 = (-(v37))
                                            v39 := GenS3{F0: 42, F1: 915}
                                            _ = len(v18)
                                            v38 = (v37 & 1)
                                            _ = v38
                                            _ = v39
                                        } else {
                                            if true {
                                                var v40 uint32 = 2147483648
                                                v18 = append(v18, (fi6 * v36))
                                                _ = v40
                                            }
                                            v12 -= v12
                                        }
                                        if (false || false) {
                                            var v41 uint = 2
                                            _ = v41
                                        } else {
                                            var v42 uint8 = p0
                                            var v43 uint64 = 4294967295
                                            v36 = 9223372036854775807
                                            switch (fi7 ^ v36) {
                                            case -389:
                                                var v44 float32 = 3.14159
                                                _ = v44
                                            case -543:
                                                fmt.Println(fi6)
                                            }
                                            if (v43 == 2) {
                                                fmt.Println(fi6, v43, v37)
                                                if len(v18) > 0 {
                                                    v18[0] = fi7
                                                }
                                                p0 -= (p0 ^ (p0 & 128))
                                                v45 := GenS3{F0: v43, F1: 255}
                                                _ = v45
                                            }
                                            _ = v42
                                            _ = v43
                                        }
                                        p0 -= (p0 * p0)
                                    default:
                                        var v46 float64 = (0.1 - ((1e10 - 0.1) * 0.1))
                                        _ = v46
                                    }
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                    v36 -= (fi7 - fi7)
                                    _ = len(v18)
                                    var v47 int64 = -999
                                    v47 = (^(v47))
                                    if ((3.14159 > (-2.25 - (-2.25 - (3.14159 + 3.14159)))) && false) {
                                        v48 := GenS4{F0: false, F1: 255}
                                        if len(v18) > 0 {
                                            v18[0] = 2147483647
                                        }
                                        if false {
                                            v36 = fi7
                                        }
                                        switch fi7 {
                                        case 404:
                                            var v49 uint = 61
                                            _ = v49
                                        case -788:
                                            if len(v18) >= 2 {
                                                v18 = v18[1:len(v18)]
                                            }
                                        }
                                        _ = v48
                                    }
                                    v47 -= v47
                                    v36 = fi6
                                    _ = v36
                                    _ = v37
                                    _ = v47
                                } else {
                                    if false {
                                        var v50 bool = true
                                        v12 -= fi7
                                        switch -676 {
                                        case -1:
                                            if (v50 || (false || v50)) {
                                                fmt.Println(v50, v12)
                                                fmt.Println(v50, fi7, v12)
                                                v50 = (v50 || v50)
                                                v51 := GenS3{F0: 18446744073709551615, F1: 1}
                                                v12 = v12
                                                _ = v51
                                            }
                                        }
                                        _ = v50
                                    } else {
                                        if (851 == 380) {
                                            v52 := GenS5{F0: 327, F1: 0}
                                            _ = v52
                                        } else {
                                            if (32768 >= 1) {
                                                if false {
                                                    p0 -= p0
                                                } else {
                                                    v12 += (-(fi6))
                                                    v12 = 0
                                                    v12 = fi6
                                                }
                                                p0 = p0
                                            }
                                            var v53 bool = (false || false)
                                            if true {
                                                v53 = (!(!(1.5 == 1.5)))
                                                p0 -= 0
                                            } else {
                                                v18 = append(v18, (fi6 | -2147483648))
                                                v18 = append(v18, -585)
                                            }
                                            _ = v53
                                        }
                                        v54 := GenS6{F0: true, F1: 42, F2: 0}
                                        _ = v54
                                    }
                                }
                                _ = v18
                            }
                            v55 := []int{(-(fi6))}
                            _ = v55
                        }
                        _ = v12
                    }
                } else {
                    switch (fi6 | fi7) {
                    case -2147483648:
                        var v56 int64 = 42
                        _ = v56
                    case 4294967295:
                        var v57 float64 = 3.14159
                        _ = v57
                    default:
                        var v58 uint16 = 255
                        _ = v58
                    }
                }
            }
            _ = fi7
        }
        p0 -= p0
        _ = fi6
    }
    switch 128 {
    case 42:
        v59 := GenS1{F0: true, F1: 203}
        _ = v59
    default:
        if false {
            var v60 string = ("foo bar" + ((("" + "a") + "foo bar") + ("foo bar" + "\n\t")))
            var v61 int16 = -581
            v62 := GenS2{F0: 127, F1: 2147483647, F2: v60}
            _ = v60
            _ = v61
            _ = v62
        } else {
            if (255 >= 1) {
                var v63 uint = 255
                v64 := GenS2{F0: 419, F1: 992, F2: "hello"}
                switch "\n\t" {
                case "a":
                    var v65 float32 = float32((float64((float32(3.14159) - -2.25)) * (float64(1e10) / (-2.25 - (-2.25 - (float64(((((3.14159 + (0.0 - (float32(1e10) - (1e10 * (1.5 / (0.1 - (1.5 + 3.14159))))))) / (3.14159 + 0.1)) - -2.25) * -2.25)) - 0.1))))))
                    _ = v65
                case "hello":
                    if (3.14159 >= -2.25) {
                        switch "" {
                        case "a":
                            v66 := GenS1{F0: false, F1: 42}
                            _ = v66
                        case "\n\t":
                            v67 := []int{0, -481, -1}
                            _ = v67
                        default:
                            fmt.Println(p0, p0)
                        }
                        for fi68 := 0; fi68 < 1; fi68++ {
                            v63 += (v63 - (v63 >> (((^(v63))) & 63)))
                            v63 += v63
                            _ = fi68
                        }
                    }
                case "foo bar":
                    for fi69 := 0; fi69 < 17; fi69++ {
                        switch (fi69 | (^(fi69))) {
                        case 2147483647:
                            switch fi69 {
                            case -780:
                                v63 = v63
                            case 32768:
                                for fi70 := 0; fi70 < 1; fi70++ {
                                    v71 := []int{(fi70 | 2147483647), fi70}
                                    _ = fi70
                                    _ = v71
                                }
                            case -9223372036854775808:
                                var v72 uint16 = 1
                                _ = v72
                            default:
                                var v73 float32 = 1.5
                                _ = v73
                            }
                        case 32768:
                            v74 := GenS7{F0: -75, F1: (v63 + 18446744073709551615), F2: v63}
                            _ = v74
                        }
                        v75 := GenS7{F0: 1, F1: (v63 - v63), F2: 255}
                        _ = fi69
                        _ = v75
                    }
                }
                _ = v63
                _ = v64
            }
        }
    }
    _ = p0
    return -836
}
func genFn1() uint64 {
    var v76 uint64 = 91
    v77 := []int{}
    _ = v76
    _ = v77
    return v76
}

func main() {
    var v1 int8 = -92
    var v2 uint16 = 255
    v3 := make(map[uint32]string)
    var v4 uint8 = 186
    v5 := make(map[uint32]uint32)
    v4 = 42
    if (!(true && ((!(751 > 629)) && true))) {
        v2 += v2
        v3[0] = ("hello" + (("a" + "\n\t") + ("\n\t" + "foo bar")))
        if true {
            if false {
                switch "foo bar" {
                case "\n\t":
                    switch 4294967295 {
                    case 119:
                        switch ("\n\t" + "a") {
                        case "\n\t":
                            v78 := []int{4294967295, 32767, 1, 0}
                            _ = v78
                        case "a":
                            var v79 uint32 = 483
                            _ = v79
                        case "hello":
                            var v80 int16 = 0
                            _ = v80
                        default:
                            v81 := GenS8{F0: 1, F1: ("" + "\n\t")}
                            _ = v81
                        }
                    case 128:
                        _ = len(v5)
                    case 290:
                        switch "hello" {
                        case "":
                            var v82 int = 1
                            _ = v82
                        case "foo bar":
                            var v83 int64 = 9223372036854775807
                            _ = v83
                        case "hello":
                            for fi84 := 0; fi84 < 5; fi84++ {
                                v1 = 1
                                if (true || false) {
                                    fmt.Println(fi84, v2)
                                } else {
                                    switch "\n\t" {
                                    case "a":
                                        v85 := GenS1{F0: false, F1: 42}
                                        _ = v85
                                    case "hello":
                                        if false {
                                            v3[788] = ("foo bar" + "a")
                                        }
                                    }
                                    switch -331 {
                                    case 309:
                                        _ = len(v5)
                                    }
                                    fmt.Println(fi84)
                                }
                                _ = fi84
                            }
                        default:
                            if ((v4 > 127) || false) {
                                _ = len(v5)
                                switch "foo bar" {
                                case "hello":
                                    switch ("" + ("foo bar" + "a")) {
                                    case "hello":
                                        fmt.Println(v1, v4)
                                    case "":
                                        v1 -= v1
                                    }
                                case "":
                                    for fi86 := 0; fi86 < 16; fi86++ {
                                        switch (fi86 % ((-9223372036854775808) | 1)) {
                                        case 9223372036854775807:
                                            fmt.Println(v1, v2, fi86)
                                        }
                                        for fi87 := 0; fi87 < 4; fi87++ {
                                            var v88 float64 = (0.0 + (0.0 / 1.0))
                                            var v89 uint16 = 857
                                            _ = fi87
                                            _ = v88
                                            _ = v89
                                        }
                                        v3[451] = "hello"
                                        _ = fi86
                                    }
                                case "\n\t":
                                    var v90 uint32 = 940
                                    _ = v90
                                }
                            }
                        }
                    }
                default:
                    if true {
                        var v91 uint32 = 569
                        for fi92 := 0; fi92 < 1; fi92++ {
                            if true {
                                var v93 int64 = -146
                                v93 = v93
                                v93 = (v93 >> ((4294967295) & 63))
                                _ = v93
                            } else {
                                var v94 int32 = 0
                                v95 := GenS9{F0: 4294967295}
                                var v96 uint32 = v91
                                _ = v94
                                _ = v95
                                _ = v96
                            }
                            fmt.Println(v1, v1, fi92)
                            v97 := GenS10{F0: -526}
                            _ = fi92
                            _ = v97
                        }
                        v98 := []int{9223372036854775807, 435, 853}
                        _ = v91
                        _ = v98
                    } else {
                        if (42 > 1) {
                            v2 += (^(v2))
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                v2 = (v2 % ((v2) | 1))
                            case "\n\t":
                                v5[1] = uint32(v2)
                            case "hello":
                                var v99 uint32 = 2147483647
                                _ = v99
                            }
                            fmt.Println(v2, v2, v2)
                        }
                    }
                }
                v100 := GenS1{F0: false, F1: 0}
                _ = v100
            }
            var v101 float64 = ((1.5 + 1e10) + 1.5)
            for fi102 := 0; fi102 < 7; fi102++ {
                switch int(v4) {
                case 42:
                    var v103 int32 = 42
                    _ = v103
                case -1:
                    for fi104 := 0; fi104 < 12; fi104++ {
                        v105 := []int{(^(fi104)), fi102, fi102, (fi104 % ((fi104) | 1))}
                        _ = len(v5)
                        v106 := GenS9{F0: uint64(v4)}
                        _ = fi104
                        _ = v105
                        _ = v106
                    }
                case -2147483648:
                    fmt.Println(v101, v101, v101)
                }
                _ = fi102
            }
            _ = v101
        } else {
            switch 425 {
            case 32767:
                v4 = (^(v4))
            case 0:
                if ((0 <= v2) && false) {
                    var v107 float64 = (1.5 - 3.14159)
                    _ = v107
                }
            }
            v108 := GenS11{F0: 93, F1: 809}
            fmt.Println(v2, v1, v4)
            _ = v108
        }
    }
    var v109 uint64 = uint64(v2)
    var v110 float32 = float32(float64(-2.25))
    var v111 uint16 = (v2 << ((322) & 63))
    v112 := GenS12{F0: int8(v2), F1: 42}
    v4 = (v4 << ((725) & 63))
    for fi113 := 0; fi113 < 8; fi113++ {
        var v114 uint16 = v111
        _ = fi113
        _ = v114
    }
    switch int(v1) {
    case 2147483647:
        var v115 uint16 = 637
        _ = v115
    }
    v116 := GenS13{F0: "\n\t"}
    for fi117 := 0; fi117 < 20; fi117++ {
        v118 := GenS14{F0: v109, F1: 1, F2: -215}
        v2 += v2
        _ = fi117
        _ = v118
    }
    var v119 bool = (false || (int32(v2) >= -2147483648))
    switch -936 {
    case 2147483647:
        v120 := []int{}
        _ = v120
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v109:", v109)
    fmt.Println("v110:", v110)
    fmt.Println("v111:", v111)
    fmt.Println("v119:", v119)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v109
    _ = v110
    _ = v111
    _ = v112
    _ = v116
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
