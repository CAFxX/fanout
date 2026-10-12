package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int; F2 int64 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint64; F1 uint }
type GenS9 struct { F0 int64; F1 int; F2 uint32 }
type GenS10 struct { F0 int32; F1 bool; F2 int }
type GenS11 struct { F0 int16; F1 int }
type GenS12 struct { F0 bool; F1 string }
type GenS13 struct { F0 uint64; F1 uint32 }
type GenS14 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS15 struct { F0 uint8; F1 int8 }
type GenS16 struct { F0 int8; F1 uint; F2 uint8 }

func main() {
    var v1 int8 = 0
    var v2 bool = false
    var v3 float64 = 0.1
    v4 := []int{9223372036854775807, -30, 995}
    var v5 uint8 = 127
    v5 += 197
    var v6 float32 = (1e10 * 0.0)
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(v2, v5, v1)
        if (int(v5) == fi7) {
            for fi8 := 0; fi8 < 20; fi8++ {
                if true {
                    switch ("" + "a") {
                    case "foo bar":
                        var v9 int64 = -9223372036854775808
                        _ = v9
                    default:
                        v6 += v6
                    }
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        var v10 string = "a"
                        _ = v10
                    default:
                        v11 := GenS1{F0: (v1 / ((v1) | 1))}
                        _ = v11
                    }
                    var v12 uint8 = v5
                    _ = v12
                }
                var v13 string = "hello"
                _ = fi8
                _ = v13
            }
            switch "" {
            case "\n\t":
                var v14 bool = (434 != 0)
                _ = v14
            case "":
                for fi15 := 0; fi15 < 19; fi15++ {
                    _ = len(v4)
                    switch "hello" {
                    case "hello":
                        if (!true) {
                            switch fi7 {
                            case 0:
                                var v16 int8 = -128
                                _ = v16
                            case 9223372036854775807:
                                v5 += v5
                            case 128:
                                if false {
                                    v1 -= 0
                                    fmt.Println(fi15, v5)
                                } else {
                                    if (false || (455 > 42)) {
                                        v6 -= ((-2.25 - (3.14159 / v6)) - ((1e10 / 1.5) / v6))
                                        switch (fi15 * 899) {
                                        case 1:
                                            var v17 int64 = -356
                                            _ = v17
                                        }
                                    } else {
                                        var v18 uint32 = 2147483648
                                        var v19 bool = (845 != 18446744073709551615)
                                        _ = v18
                                        _ = v19
                                    }
                                    var v20 float32 = v6
                                    v4 = append(v4, fi7)
                                    _ = v20
                                }
                            }
                        }
                    default:
                        switch ("foo bar" + "a") {
                        case "foo bar":
                            if v2 {
                                if v2 {
                                    switch fi7 {
                                    case -512:
                                        v3 = v3
                                    case -1:
                                        if (!(v2 && (!(!(false || (!false)))))) {
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v21 int64 = 83
                                                _ = v21
                                            default:
                                                fmt.Println(fi15)
                                            }
                                            v22 := GenS2{F0: (!false)}
                                            v23 := GenS1{F0: v1}
                                            _ = v22
                                            _ = v23
                                        }
                                    case 128:
                                        if (v6 < (v6 + 0.1)) {
                                            v2 = v2
                                            if ((v5 * (v5 / ((v5) | 1))) < (v5 << ((2) & 63))) {
                                                fmt.Println(v5)
                                            } else {
                                                if ((-361 != 505) || v2) {
                                                    switch (("a" + "hello") + "foo bar") {
                                                    case "":
                                                        var v24 float32 = (v6 / 1e10)
                                                        _ = v24
                                                    default:
                                                        var v25 float32 = 3.14159
                                                        _ = v25
                                                    }
                                                    switch 128 {
                                                    case 776:
                                                        if (v5 <= v5) {
                                                            fmt.Println(fi15, v1, fi15)
                                                            v26 := GenS2{F0: (!(0 <= 42))}
                                                            if (1 == uint(v5)) {
                                                                switch (fi7 << ((65535) & 63)) {
                                                                case 32767:
                                                                    var v27 int64 = 607
                                                                    _ = v27
                                                                case -277:
                                                                    v1 += v1
                                                                case -736:
                                                                    v5 += v5
                                                                default:
                                                                    var v28 int64 = 9223372036854775807
                                                                    _ = v28
                                                                }
                                                                v1 = (v1 | 42)
                                                                if true {
                                                                    if true {
                                                                        if v2 {
                                                                            var v29 int64 = -759
                                                                            switch ("foo bar" + "\n\t") {
                                                                            case "":
                                                                                switch "\n\t" {
                                                                                case "a":
                                                                                    v30 := GenS2{F0: (("\n\t" + "a") != "")}
                                                                                    var v31 uint32 = 65535
                                                                                    _ = v30
                                                                                    _ = v31
                                                                                case "":
                                                                                    var v32 int = 9223372036854775807
                                                                                    _ = v32
                                                                                }
                                                                            case "foo bar":
                                                                                if false {
                                                                                    switch "foo bar" {
                                                                                    case "hello":
                                                                                        var v33 uint8 = 33
                                                                                        _ = v33
                                                                                    case "":
                                                                                        if (32767 == 261) {
                                                                                            v2 = true
                                                                                            if false {
                                                                                                var v34 int64 = -545
                                                                                                if (true && (v2 || (!v2))) {
                                                                                                    v3 -= 0.1
                                                                                                }
                                                                                                _ = v34
                                                                                            }
                                                                                            v2 = (42 != 411)
                                                                                            var v35 uint64 = 2
                                                                                            _ = v35
                                                                                        }
                                                                                    default:
                                                                                        var v36 uint8 = 86
                                                                                        _ = v36
                                                                                    }
                                                                                    v29 = v29
                                                                                }
                                                                                if true {
                                                                                    var v37 string = ("hello" + "foo bar")
                                                                                    v38 := GenS2{F0: ((^(v5)) >= v5)}
                                                                                    _ = len(v4)
                                                                                    var v39 string = ""
                                                                                    v39 = "a"
                                                                                    _ = v37
                                                                                    _ = v38
                                                                                    _ = v39
                                                                                }
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            case "\n\t":
                                                                                var v40 float32 = float32((3.14159 / float64(v6)))
                                                                                _ = v40
                                                                            }
                                                                            _ = v29
                                                                        } else {
                                                                            fmt.Println(v5, v6, v2)
                                                                            fmt.Println(v5, fi15)
                                                                        }
                                                                        v2 = (-248 != 536)
                                                                        fmt.Println(v1, v5, v3)
                                                                    } else {
                                                                        if (2 > 2) {
                                                                            var v41 bool = v2
                                                                            v5 -= 162
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            _ = v41
                                                                        }
                                                                        if (!v2) {
                                                                            if v2 {
                                                                                fmt.Println(v2, v2, v3)
                                                                                var v42 uint8 = (v5 * (v5 ^ (v5 / ((v5) | 1))))
                                                                                var v43 int16 = 1000
                                                                                if (!(v6 == v6)) {
                                                                                    fmt.Println(v5, v5, v6)
                                                                                } else {
                                                                                    if true {
                                                                                        var v44 string = "\n\t"
                                                                                        var v45 int = fi7
                                                                                        _ = v44
                                                                                        _ = v45
                                                                                    } else {
                                                                                        var v46 uint8 = (v42 * 165)
                                                                                        fmt.Println(v6)
                                                                                        _ = v46
                                                                                    }
                                                                                    if true {
                                                                                        var v47 int16 = 741
                                                                                        switch "hello" {
                                                                                        case "":
                                                                                            v42 -= (v5 << ((65535) & 63))
                                                                                        case "a":
                                                                                            switch (fi7 >> ((1) & 63)) {
                                                                                            case 32767:
                                                                                                v48 := GenS3{F0: 0, F1: 790, F2: 1}
                                                                                                _ = v48
                                                                                            }
                                                                                        case "foo bar":
                                                                                            v49 := GenS2{F0: (!v2)}
                                                                                            _ = v49
                                                                                        default:
                                                                                            var v50 int16 = v47
                                                                                            fmt.Println(v6, v43, v47)
                                                                                            v1 += (-(v1))
                                                                                            _ = len(v4)
                                                                                            v3 -= v3
                                                                                            _ = v50
                                                                                        }
                                                                                        v2 = v2
                                                                                        _ = v47
                                                                                    } else {
                                                                                        var v51 float32 = v6
                                                                                        _ = v51
                                                                                    }
                                                                                    var v52 uint16 = 42
                                                                                    v53 := GenS4{F0: v2, F1: 0}
                                                                                    _ = v52
                                                                                    _ = v53
                                                                                }
                                                                                _ = v42
                                                                                _ = v43
                                                                            } else {
                                                                                v4 = append(v4, (fi7 >> ((65535) & 63)))
                                                                                var v54 int = (fi7 * (^(fi7)))
                                                                                v1 = (v1 * v1)
                                                                                _ = v54
                                                                            }
                                                                            var v55 float64 = ((((v3 - (3.14159 + -2.25)) + (v3 - (3.14159 / v3))) - -2.25) - v3)
                                                                            var v56 float32 = (v6 / (v6 * v6))
                                                                            _ = v55
                                                                            _ = v56
                                                                        }
                                                                        v2 = (fi15 > (^(fi15)))
                                                                    }
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    var v57 uint = 4294967295
                                                                    _ = v57
                                                                } else {
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    fmt.Println(v3, v2, v6)
                                                                    fmt.Println(fi15, v2, fi7)
                                                                }
                                                            }
                                                            var v58 int8 = 0
                                                            _ = v26
                                                            _ = v58
                                                        }
                                                    case -9223372036854775808:
                                                        var v59 float64 = v3
                                                        if len(v4) > 0 {
                                                            v4[0] = 127
                                                        }
                                                        _ = v59
                                                    default:
                                                        v60 := GenS4{F0: v2, F1: 42}
                                                        _ = v60
                                                    }
                                                } else {
                                                    v1 = -1
                                                    var v61 uint8 = v5
                                                    _ = v61
                                                }
                                                var v62 uint8 = 28
                                                _ = v62
                                            }
                                            var v63 uint16 = 0
                                            v64 := GenS5{F0: (fi7 >> ((18446744073709551615) & 63)), F1: -32768}
                                            var v65 uint8 = v5
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                        } else {
                                            v5 -= (^(v5))
                                        }
                                        var v66 uint = 0
                                        _ = v66
                                    }
                                }
                                v2 = (true || v2)
                            } else {
                                v67 := GenS1{F0: v1}
                                _ = v67
                            }
                        default:
                            var v68 uint16 = 65535
                            _ = v68
                        }
                    }
                    _ = fi15
                }
            case "hello":
                v1 = (-(v1))
            }
        }
        if v2 {
            v5 = 2
        } else {
            fmt.Println(v6, v3)
            v6 = (v6 + (float32(v3) + -2.25))
            fmt.Println(v3, v6, v5)
        }
        _ = fi7
    }
    v5 -= v5
    v5 -= v5
    var v69 float64 = v3
    var v70 uint32 = 2
    var v71 uint32 = 0
    var v72 bool = (!false)
    var v73 bool = true
    fmt.Println(v71, v5)
    switch ("a" + "\n\t") {
    case "foo bar":
        var v74 uint32 = 291
        _ = v74
    case "hello":
        for fi75 := 0; fi75 < 2; fi75++ {
            switch 113 {
            case 289:
                var v76 int32 = 1
                _ = v76
            default:
                v5 += v5
            }
            v5 = (v5 % ((255) | 1))
            for fi77 := 0; fi77 < 15; fi77++ {
                if v73 {
                    v78 := GenS4{F0: (fi77 < 604), F1: 2147483647}
                    v79 := GenS2{F0: false}
                    var v80 uint16 = 464
                    _ = v78
                    _ = v79
                    _ = v80
                } else {
                    fmt.Println(v5)
                }
                v81 := GenS6{F0: fi75}
                _ = fi77
                _ = v81
            }
            _ = fi75
        }
    }
    v69 -= v3
    switch 32768 {
    case 4294967295:
        v82 := GenS7{F0: 229}
        _ = v82
    case 127:
        var v83 int16 = 32767
        _ = v83
    case 92:
        if (((0.1 / 0.1) / float32(1.5)) <= ((0.1 - v6) / float32((-2.25 + ((0.1 + ((v69 + v69) + 0.1)) / v69))))) {
            v72 = (v72 || true)
            for fi84 := 0; fi84 < 1; fi84++ {
                v85 := GenS8{F0: 18446744073709551615, F1: 935}
                _ = fi84
                _ = v85
            }
            var v86 int8 = 100
            _ = v86
        }
    default:
        if false {
            for fi87 := 0; fi87 < 5; fi87++ {
                if v73 {
                    for fi88 := 0; fi88 < 11; fi88++ {
                        if (-1 < -8) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v89 uint16 = 1
                            _ = v89
                        } else {
                            fmt.Println(v71, v1)
                            if (((fi88 << ((0) & 63)) > fi87) || v2) {
                                v71 = v70
                                if len(v4) > 0 {
                                    v4[0] = (-(fi88))
                                }
                                var v90 int = fi88
                                _ = v90
                            } else {
                                fmt.Println(v73, v6, v1)
                                var v91 uint16 = 571
                                _ = v91
                            }
                        }
                        _ = fi88
                    }
                } else {
                    if false {
                        if (!true) {
                            v69 = 0.1
                            var v92 uint64 = 93
                            var v93 int64 = 42
                            _ = v92
                            _ = v93
                        } else {
                            if (v5 <= (v5 ^ 84)) {
                                _ = len(v4)
                                switch (^(fi87)) {
                                case 68:
                                    _ = len(v4)
                                default:
                                    if (0 < v70) {
                                        fmt.Println(v70, v72)
                                        for fi94 := 0; fi94 < 16; fi94++ {
                                            switch (fi87 - fi87) {
                                            case 311:
                                                if true {
                                                    fmt.Println(v2, v3)
                                                } else {
                                                    v95 := GenS2{F0: true}
                                                    var v96 float32 = v6
                                                    var v97 uint = 255
                                                    _ = v95
                                                    _ = v96
                                                    _ = v97
                                                }
                                            case 2147483647:
                                                v6 += (((v6 + v6) + ((0.1 + v6) * v6)) / v6)
                                            case 0:
                                                v98 := GenS5{F0: fi87, F1: -509}
                                                _ = v98
                                            default:
                                                v3 += v69
                                            }
                                            var v99 float32 = 0.0
                                            _ = fi94
                                            _ = v99
                                        }
                                        v71 = 2
                                    }
                                }
                                if (!v72) {
                                    v100 := GenS1{F0: v1}
                                    v101 := GenS7{F0: fi87}
                                    switch (-(fi87)) {
                                    case 127:
                                        for fi102 := 0; fi102 < 14; fi102++ {
                                            v69 = v69
                                            if (((928 == -9223372036854775808) && (-1000 > int64(v1))) || (!((v5 | (v5 * (v5 | 1))) < v5))) {
                                                if false {
                                                    var v103 int = -419
                                                    if false {
                                                        v70 += (v70 + v71)
                                                        if (!(0 != 995)) {
                                                            if v2 {
                                                                v103 -= fi87
                                                                if v2 {
                                                                    v4 = append(v4, v103)
                                                                    var v104 uint32 = (^(v71))
                                                                    v103 += 319
                                                                    _ = v104
                                                                } else {
                                                                    if (v2 && ((^(v1)) > int8(fi102))) {
                                                                        var v105 float32 = v6
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            if len(v4) > 0 {
                                                                                v4[0] = (fi87 >> ((255) & 63))
                                                                            }
                                                                            var v106 int8 = v1
                                                                            v4 = append(v4, fi87)
                                                                            _ = v106
                                                                        case "\n\t":
                                                                            switch fi87 {
                                                                            case 32768:
                                                                                v107 := GenS2{F0: (1e10 <= v105)}
                                                                                _ = v107
                                                                            case -9223372036854775808:
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = v103
                                                                                }
                                                                            default:
                                                                                var v108 float32 = (v6 * 1.5)
                                                                                _ = v108
                                                                            }
                                                                            v5 -= 42
                                                                            fmt.Println(v103, v2)
                                                                        default:
                                                                            var v109 float64 = 1.5
                                                                            _ = v109
                                                                        }
                                                                        _ = v105
                                                                    }
                                                                    v110 := GenS1{F0: (v1 % ((105) | 1))}
                                                                    v69 = (v69 / 1e10)
                                                                    v103 = (-(fi87))
                                                                    v72 = (v73 && true)
                                                                    _ = v110
                                                                }
                                                                var v111 int8 = v1
                                                                v71 = v71
                                                                v103 = (^(fi102))
                                                                v5 = (^(v5))
                                                                if v73 {
                                                                    v4 = append(v4, -1)
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    if v2 {
                                                                        var v112 int16 = 506
                                                                        switch fi87 {
                                                                        case 42:
                                                                            var v113 bool = false
                                                                            _ = v113
                                                                        default:
                                                                            switch 1 {
                                                                            case -439:
                                                                                fmt.Println(v71)
                                                                            case -1:
                                                                                if v2 {
                                                                                    var v114 int = (v103 % ((fi102) | 1))
                                                                                    v114 = fi87
                                                                                    _ = v114
                                                                                }
                                                                            }
                                                                        }
                                                                        if v72 {
                                                                            if v72 {
                                                                                v69 = 3.14159
                                                                                if v72 {
                                                                                    v103 += fi102
                                                                                    v4 = append(v4, -1)
                                                                                }
                                                                                v73 = false
                                                                            }
                                                                            v115 := GenS4{F0: v72, F1: -860}
                                                                            _ = v115
                                                                        }
                                                                        _ = v112
                                                                    } else {
                                                                        v4 = append(v4, (fi87 >> ((42) & 63)))
                                                                        var v116 int32 = -326
                                                                        _ = v116
                                                                    }
                                                                    var v117 float32 = v6
                                                                    v6 = v117
                                                                    if (v5 >= v5) {
                                                                        v73 = (-2147483648 > -2147483648)
                                                                    } else {
                                                                        v69 -= v3
                                                                        fmt.Println(v69, v117, v117)
                                                                        var v118 bool = true
                                                                        _ = v118
                                                                    }
                                                                    _ = v117
                                                                } else {
                                                                    fmt.Println(v3, v5)
                                                                    v71 = v71
                                                                }
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                v111 = -3
                                                                var v119 uint32 = 0
                                                                v2 = v72
                                                                v3 -= (-2.25 / 0.1)
                                                                if v2 {
                                                                    var v120 int8 = (v1 % ((v111) | 1))
                                                                    if (v1 != v1) {
                                                                        if (!(!v73)) {
                                                                            switch 42 {
                                                                            case -566:
                                                                                v2 = v73
                                                                            default:
                                                                                v1 = int8(v71)
                                                                            }
                                                                            var v121 uint64 = 42
                                                                            switch (fi87 >> ((877) & 63)) {
                                                                            case -816:
                                                                                v2 = v2
                                                                            }
                                                                            _ = v121
                                                                        } else {
                                                                            if ((!((!(!v2)) && ((^(v111)) == v1))) || v2) {
                                                                                if ((float64((v6 + 0.1)) + 1e10) != float64((v6 - ((v6 + ((1.5 - float32(v3)) / v6)) - 3.14159)))) {
                                                                                    var v122 uint64 = 4294967295
                                                                                    switch "a" {
                                                                                    case "foo bar":
                                                                                        if ((0.1 + ((v6 * -2.25) * ((0.0 - 1.5) / float32(v69)))) == v6) {
                                                                                            v73 = v2
                                                                                            var v123 string = "\n\t"
                                                                                            var v124 int = v103
                                                                                            _ = v123
                                                                                            _ = v124
                                                                                        }
                                                                                        var v125 uint = uint(v119)
                                                                                        v2 = ((v69 - v69) <= ((v3 / (v69 + v3)) + 1e10))
                                                                                        _ = v125
                                                                                    default:
                                                                                        v4 = append(v4, fi87)
                                                                                    }
                                                                                    var v126 string = "hello"
                                                                                    v73 = v72
                                                                                    if v73 {
                                                                                        var v127 bool = false
                                                                                        var v128 uint32 = 511
                                                                                        v103 -= int(v122)
                                                                                        _ = v127
                                                                                        _ = v128
                                                                                    }
                                                                                    _ = v122
                                                                                    _ = v126
                                                                                }
                                                                            } else {
                                                                                if ((128 < 127) && false) {
                                                                                    v111 = v120
                                                                                    var v129 uint64 = 18446744073709551615
                                                                                    _ = v129
                                                                                }
                                                                            }
                                                                            var v130 float32 = 1e10
                                                                            fmt.Println(v1)
                                                                            _ = v130
                                                                        }
                                                                        switch (-(v103)) {
                                                                        case 861:
                                                                            var v131 int8 = 74
                                                                            _ = v131
                                                                        default:
                                                                            v132 := GenS6{F0: 4294967295}
                                                                            _ = v132
                                                                        }
                                                                    }
                                                                    fmt.Println(v2)
                                                                    _ = v120
                                                                }
                                                                if v72 {
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                }
                                                                _ = v111
                                                                _ = v119
                                                            } else {
                                                                v133 := GenS9{F0: 42, F1: 2147483647, F2: 2147483647}
                                                                if false {
                                                                    fmt.Println(v5, fi102, v6)
                                                                    var v134 uint = 255
                                                                    fmt.Println(v72)
                                                                    _ = v134
                                                                }
                                                                if len(v4) > 0 {
                                                                    v4[0] = (v103 % ((fi87) | 1))
                                                                }
                                                                _ = v133
                                                            }
                                                            var v135 float64 = v3
                                                            var v136 string = ""
                                                            _ = v135
                                                            _ = v136
                                                        }
                                                        v69 -= v69
                                                    }
                                                    _ = v103
                                                }
                                                switch (fi102 - fi102) {
                                                case 886:
                                                    var v137 float32 = (0.0 + v6)
                                                    _ = v137
                                                case 4294967295:
                                                    var v138 string = "hello"
                                                    _ = v138
                                                case -858:
                                                    if len(v4) > 0 {
                                                        v4[0] = fi102
                                                    }
                                                }
                                            } else {
                                                v139 := GenS1{F0: v1}
                                                _ = v139
                                            }
                                            _ = fi102
                                        }
                                    default:
                                        if len(v4) > 0 {
                                            v4[0] = (fi87 << ((uint(v71)) & 63))
                                        }
                                    }
                                    _ = v100
                                    _ = v101
                                }
                            } else {
                                for fi140 := 0; fi140 < 16; fi140++ {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v141 := GenS4{F0: (!true), F1: int32(fi140)}
                                    v1 += (v1 >> ((uint(v5)) & 63))
                                    _ = fi140
                                    _ = v141
                                }
                                switch 9223372036854775807 {
                                case 1:
                                    v71 -= v70
                                case -35:
                                    switch (fi87 - (fi87 - -559)) {
                                    case 42:
                                        if ((1 > 513) || v72) {
                                            v142 := GenS10{F0: 42, F1: v72, F2: (fi87 / (((-(fi87))) | 1))}
                                            for fi143 := 0; fi143 < 13; fi143++ {
                                                var v144 int = (fi87 << ((255) & 63))
                                                _ = fi143
                                                _ = v144
                                            }
                                            _ = v142
                                        }
                                    }
                                case 974:
                                    if ((-311 == 0) || (v3 == v69)) {
                                        var v145 int32 = 2147483647
                                        for fi146 := 0; fi146 < 19; fi146++ {
                                            v145 += 0
                                            v4 = append(v4, 42)
                                            switch (^(fi146)) {
                                            case 795:
                                                var v147 uint64 = 876
                                                _ = v147
                                            case 9223372036854775807:
                                                if true {
                                                    var v148 int64 = 477
                                                    fmt.Println(v1)
                                                    v149 := GenS5{F0: (fi87 << ((255) & 63)), F1: -1}
                                                    if (float64(v6) > 1e10) {
                                                        v150 := GenS11{F0: -32768, F1: (fi146 / ((fi146) | 1))}
                                                        v151 := GenS12{F0: (v73 && true), F1: "a"}
                                                        var v152 int = -718
                                                        _ = v150
                                                        _ = v151
                                                        _ = v152
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = (fi146 & (-(fi87)))
                                                    }
                                                    fmt.Println(v6, v3, v148)
                                                    _ = v148
                                                    _ = v149
                                                } else {
                                                    v71 -= 65535
                                                }
                                            default:
                                                switch fi87 {
                                                case 830:
                                                    fmt.Println(v70, fi146)
                                                case 1:
                                                    fmt.Println(v6)
                                                case 128:
                                                    v2 = (v73 || (!((-31 > -32768) || v72)))
                                                default:
                                                    v153 := GenS13{F0: 69, F1: v70}
                                                    _ = v153
                                                }
                                            }
                                            _ = fi146
                                        }
                                        if (fi87 != (fi87 | (-(fi87)))) {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        } else {
                                            for fi154 := 0; fi154 < 16; fi154++ {
                                                fmt.Println(v2)
                                                _ = fi154
                                            }
                                            fmt.Println(v72, v73, v2)
                                        }
                                        _ = v145
                                    }
                                default:
                                    switch fi87 {
                                    case 127:
                                        var v155 int = fi87
                                        _ = v155
                                    case -2147483648:
                                        for fi156 := 0; fi156 < 15; fi156++ {
                                            fmt.Println(v2, v3)
                                            if v72 {
                                                if (("\n\t" + "a") > "foo bar") {
                                                    var v157 float64 = (v3 + ((((v69 - float64(v6)) * float64(1e10)) * 1e10) + v69))
                                                    _ = v157
                                                } else {
                                                    v158 := GenS14{F0: 657, F1: (v71 * (v70 | (v71 & (v70 / ((86) | 1))))), F2: uint64(v71)}
                                                    var v159 uint8 = 192
                                                    _ = v158
                                                    _ = v159
                                                }
                                                switch 4294967295 {
                                                case -1:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v160 uint64 = 2
                                                    if v2 {
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            fmt.Println(fi156, v6)
                                                        case "hello":
                                                            switch "" {
                                                            case "hello":
                                                                v70 = 0
                                                            case "":
                                                                var v161 bool = true
                                                                _ = v161
                                                            default:
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    switch "" {
                                                                    case "a":
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            var v162 uint32 = v71
                                                                            v2 = (v6 <= 1.5)
                                                                            _ = v162
                                                                        case "hello":
                                                                            var v163 int16 = -152
                                                                            _ = v163
                                                                        }
                                                                    case "":
                                                                        if v73 {
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                var v164 int8 = (v1 | (v1 / ((v1) | 1)))
                                                                                _ = v164
                                                                            case "\n\t":
                                                                                v6 = v6
                                                                            }
                                                                            if true {
                                                                                v5 = (v5 ^ (v5 - 255))
                                                                                var v165 float32 = 0.1
                                                                                switch "" {
                                                                                case "a":
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                default:
                                                                                    v166 := GenS15{F0: (v5 | v5), F1: -1}
                                                                                    _ = v166
                                                                                }
                                                                                _ = v165
                                                                            } else {
                                                                                var v167 uint8 = 127
                                                                                v69 = -2.25
                                                                                fmt.Println(fi87, v73)
                                                                                _ = v167
                                                                            }
                                                                            v70 = (v70 * (v71 | 273))
                                                                        } else {
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                if false {
                                                                                    if v72 {
                                                                                        var v168 uint8 = 42
                                                                                        v169 := GenS10{F0: 42, F1: v2, F2: (fi87 / ((9223372036854775807) | 1))}
                                                                                        _ = v168
                                                                                        _ = v169
                                                                                    } else {
                                                                                        v6 = 0.0
                                                                                        v170 := GenS6{F0: fi87}
                                                                                        fmt.Println(v160)
                                                                                        _ = v170
                                                                                    }
                                                                                }
                                                                            case "hello":
                                                                                var v171 int32 = 0
                                                                                _ = v171
                                                                            default:
                                                                                v2 = true
                                                                            }
                                                                            v172 := GenS16{F0: (^(v1)), F1: 42, F2: v5}
                                                                            var v173 string = "a"
                                                                            _ = v172
                                                                            _ = v173
                                                                        }
                                                                    case "\n\t":
                                                                        v3 = (3.14159 - v3)
                                                                    }
                                                                default:
                                                                    v160 = 2
                                                                }
                                                            }
                                                        default:
                                                            v72 = (0 <= -9223372036854775808)
                                                        }
                                                        var v174 int64 = 1
                                                        v71 -= v70
                                                        _ = v174
                                                    } else {
                                                        var v175 uint8 = v5
                                                        _ = len(v4)
                                                        _ = v175
                                                    }
                                                    _ = v160
                                                case -2147483648:
                                                    var v176 int = int(v1)
                                                    _ = v176
                                                }
                                                var v177 string = "hello"
                                                _ = v177
                                            }
                                            _ = fi156
                                        }
                                    case 128:
                                        v178 := GenS10{F0: -2147483648, F1: (fi87 > fi87), F2: (^(fi87))}
                                        _ = v178
                                    }
                                }
                                for fi179 := 0; fi179 < 1; fi179++ {
                                    if len(v4) > 0 {
                                        v4[0] = (fi87 & -547)
                                    }
                                    _ = fi179
                                }
                            }
                        }
                        var v180 uint = 4294967295
                        var v181 uint8 = v5
                        _ = v180
                        _ = v181
                    }
                    for fi182 := 0; fi182 < 2; fi182++ {
                        v183 := GenS12{F0: (v73 && (((^(v5)) >= v5) || (v2 || (2 != 638)))), F1: (("" + "\n\t") + ("a" + ("hello" + ("\n\t" + ("" + "foo bar")))))}
                        _ = fi182
                        _ = v183
                    }
                }
                v2 = ((^(v1)) >= (v1 + 84))
                _ = fi87
            }
            var v184 int16 = 0
            switch 691 {
            case 2147483647:
                var v185 float64 = 0.1
                _ = v185
            case 32767:
                v2 = (!true)
            case 9223372036854775807:
                v186 := GenS1{F0: (-(v1))}
                _ = v186
            }
            _ = v184
        } else {
            for fi187 := 0; fi187 < 8; fi187++ {
                for fi188 := 0; fi188 < 19; fi188++ {
                    if (-9223372036854775808 == 42) {
                        v71 += 2
                        if len(v4) > 0 {
                            v4[0] = fi187
                        }
                        v189 := GenS7{F0: (-(fi187))}
                        v6 = v6
                        _ = v189
                    } else {
                        switch "hello" {
                        case "":
                            var v190 bool = (!(!(-761 >= 42)))
                            _ = v190
                        default:
                            var v191 float64 = v3
                            _ = v191
                        }
                    }
                    _ = fi188
                }
                v6 -= v6
                _ = fi187
            }
            v2 = (!v2)
            v69 = 1.5
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v70
    _ = v71
    _ = v72
    _ = v73
}
