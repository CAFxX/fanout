package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint8; F1 uint8; F2 int }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64; F1 int; F2 uint32 }
type GenS9 struct { F0 bool; F1 int64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 string; F1 uint32; F2 uint }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 int32; F1 int; F2 uint }
func genFn0(p0 uint32, p1 int64) int32 {
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(fi7, fi7)
            fmt.Println(fi8, p0, p1)
            if (-1 != -1) {
                var v9 float32 = float32(-2.25)
                switch fi8 {
                case 42:
                    v10 := []int{(fi8 - (fi7 + (fi7 % (((fi8 - fi8)) | 1)))), fi8, (-(fi8))}
                    _ = v10
                case 32768:
                    p1 -= (p1 & p1)
                }
                _ = v9
            } else {
                p0 -= p0
            }
            _ = fi8
        }
        v11 := []int{}
        _ = fi7
        _ = v11
    }
    switch 654 {
    case -899:
        p1 = 0
    default:
        p0 = (p0 - p0)
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint16 = 2
        _ = fi12
        _ = v13
    }
    var v14 float32 = 1.5
    _ = p0
    _ = p1
    _ = v14
    return 879
}

func main() {
    var v1 bool = false
    v2 := make(map[string]uint32)
    var v3 uint8 = 1
    var v4 uint16 = 55
    var v5 int16 = -1
    var v6 int8 = -92
    switch 699 {
    case 0:
        v5 -= (v5 % (((v5 * 0)) | 1))
    case 820:
        if v1 {
            switch 63 {
            case -9223372036854775808:
                v15 := GenS1{F0: 65535, F1: 4294967295}
                _ = v15
            default:
                var v16 int64 = 9223372036854775807
                _ = v16
            }
            v17 := GenS1{F0: 18446744073709551615, F1: 918}
            _ = v17
        }
    case 2147483647:
        if ((-(v5)) != -32768) {
            v18 := GenS1{F0: 564, F1: 2147483647}
            _ = v18
        } else {
            fmt.Println(v1, v5)
            switch (((("hello" + "") + "\n\t") + "hello") + ("a" + ("\n\t" + ""))) {
            case "a":
                v19 := GenS1{F0: 42, F1: 128}
                _ = v19
            case "hello":
                v6 += v6
            case "\n\t":
                if (v5 != v5) {
                    v6 = v6
                }
            }
        }
    }
    var v20 float64 = ((0.1 + ((float64(float32(1.5)) * (1e10 / 3.14159)) + -2.25)) / (3.14159 * 1.5))
    var v21 uint64 = 927
    if (-98 > 1) {
        var v22 int32 = 42
        if false {
            v6 -= (v6 % ((v6) | 1))
            if (42 <= 42) {
                if false {
                    switch int(v4) {
                    case 127:
                        if (v5 < (v5 ^ v5)) {
                            var v23 float64 = ((v20 + ((((1e10 + 1e10) - 1.5) + v20) - (v20 * -2.25))) / 3.14159)
                            fmt.Println(v23, v4)
                            _ = v23
                        } else {
                            var v24 bool = (!v1)
                            v25 := GenS2{F0: 0, F1: 4294967295}
                            for fi26 := 0; fi26 < 5; fi26++ {
                                for fi27 := 0; fi27 < 1; fi27++ {
                                    if (v1 || (v20 < v20)) {
                                        if false {
                                            v4 = (v4 + 425)
                                        }
                                        var v28 int32 = 203
                                        if v1 {
                                            v29 := GenS3{F0: 127, F1: (v3 - v3), F2: (fi27 * (fi26 - fi27))}
                                            fmt.Println(v22, v4)
                                            v30 := GenS4{F0: v21, F1: 127}
                                            _ = v29
                                            _ = v30
                                        }
                                        _ = v28
                                    }
                                    _ = fi27
                                }
                                v31 := GenS5{F0: (true || false)}
                                v32 := GenS5{F0: (353 > 595)}
                                _ = fi26
                                _ = v31
                                _ = v32
                            }
                            _ = v24
                            _ = v25
                        }
                    case 1:
                        var v33 int8 = (-(v6))
                        _ = v33
                    }
                    v2["hello"] = 2147483647
                    for fi34 := 0; fi34 < 3; fi34++ {
                        delete(v2, "hello")
                        v2["foo bar"] = 65535
                        var v35 int64 = 9223372036854775807
                        _ = fi34
                        _ = v35
                    }
                } else {
                    if ("hello" < (("foo bar" + "") + "hello")) {
                        switch 127 {
                        case 459:
                            v21 -= (v21 >> ((2) & 63))
                        }
                        v36 := GenS6{F0: "\n\t", F1: 1}
                        switch (("a" + ("" + ("hello" + ("foo bar" + "hello")))) + "a") {
                        case "\n\t":
                            var v37 int = -1
                            _ = v37
                        case "":
                            v1 = v1
                        }
                        _ = v36
                    } else {
                        var v38 string = ("hello" + ("\n\t" + ""))
                        _ = v38
                    }
                }
                v39 := GenS6{F0: (("" + ("a" + "hello")) + ""), F1: -1}
                if v1 {
                    v40 := v2[((("a" + "foo bar") + ("hello" + "")) + "\n\t")]
                    if (true || v1) {
                        if v1 {
                            if v1 {
                                if (508 > (v21 | (v21 & v21))) {
                                    for fi41 := 0; fi41 < 14; fi41++ {
                                        for fi42 := 0; fi42 < 10; fi42++ {
                                            v43 := GenS1{F0: 1, F1: 42}
                                            if true {
                                                if v1 {
                                                    if true {
                                                        if v1 {
                                                            if v1 {
                                                                if ((v22 / (((v22 & v22)) | 1)) == (^(v22))) {
                                                                    v6 = 1
                                                                    v44 := []int{(fi41 ^ 42)}
                                                                    if len(v44) >= 2 {
                                                                        v44 = v44[1:len(v44)]
                                                                    }
                                                                    _ = v44
                                                                }
                                                                var v45 int16 = -32768
                                                                _ = v45
                                                            }
                                                        }
                                                        var v46 int = fi41
                                                        v2["a"] = (v40 + v40)
                                                        _ = v46
                                                    } else {
                                                        if false {
                                                            v47 := GenS1{F0: 814, F1: (fi41 ^ -1)}
                                                            _ = v47
                                                        } else {
                                                            v22 += 0
                                                            v2["hello"] = 881
                                                        }
                                                        v48 := []int{(^(fi42)), fi41}
                                                        fmt.Println(v22, v6, fi42)
                                                        _ = v48
                                                    }
                                                }
                                                if v1 {
                                                    v49 := v2[""]
                                                    if (genFn0((v40 & (v49 ^ 0)), int64(v3)) == v22) {
                                                        var v50 uint64 = (v21 << ((18446744073709551615) & 63))
                                                        switch "a" {
                                                        case "foo bar":
                                                            switch 32768 {
                                                            case 42:
                                                                switch "" {
                                                                case "foo bar":
                                                                    if (v50 <= 4294967295) {
                                                                        var v51 int8 = 22
                                                                        v52 := GenS7{F0: 191}
                                                                        v40 = 760
                                                                        _ = v51
                                                                        _ = v52
                                                                    }
                                                                }
                                                                var v53 uint = 0
                                                                fmt.Println(v49)
                                                                v5 += (v5 + (v5 >> ((65535) & 63)))
                                                                v2[("" + "hello")] = 4
                                                                _ = v53
                                                            case 1:
                                                                var v54 float32 = 3.14159
                                                                _ = v54
                                                            case 2147483647:
                                                                switch ((("hello" + ("a" + "a")) + ("hello" + ("foo bar" + ""))) + "") {
                                                                case "":
                                                                    if (v1 || true) {
                                                                        _ = len(v2)
                                                                        var v55 int64 = 566
                                                                        _ = v55
                                                                    }
                                                                }
                                                            }
                                                        case "hello":
                                                            v20 = 0.0
                                                        case "":
                                                            v56 := GenS1{F0: 4294967295, F1: fi42}
                                                            _ = v56
                                                        default:
                                                            var v57 uint32 = 161
                                                            _ = v57
                                                        }
                                                        switch fi41 {
                                                        case 127:
                                                            var v58 bool = true
                                                            _ = v58
                                                        }
                                                        _ = v50
                                                    } else {
                                                        v59 := GenS8{F0: -1, F1: 127, F2: (v40 << ((42) & 63))}
                                                        if ((float64(3.14159) - v20) != 1.5) {
                                                            fmt.Println(v22, v20, v5)
                                                            fmt.Println(fi42)
                                                            _ = len(v2)
                                                        } else {
                                                            switch int(v6) {
                                                            case 32768:
                                                                v6 -= v6
                                                            }
                                                            v1 = false
                                                            var v60 int16 = v5
                                                            switch "hello" {
                                                            case "":
                                                                var v61 int64 = 42
                                                                _ = v61
                                                            case "hello":
                                                                var v62 float64 = v20
                                                                _ = v62
                                                            case "\n\t":
                                                                if ((196 == v4) || v1) {
                                                                    v5 += (v60 << ((0) & 63))
                                                                    v2["foo bar"] = (v40 ^ 2)
                                                                }
                                                            default:
                                                                var v63 uint32 = (v49 * uint32(v21))
                                                                _ = v63
                                                            }
                                                            _ = v60
                                                        }
                                                        var v64 uint16 = 125
                                                        _ = v59
                                                        _ = v64
                                                    }
                                                    _ = v49
                                                } else {
                                                    v40 += (^(v40))
                                                    fmt.Println(v3, v5, fi41)
                                                }
                                            } else {
                                                v65 := []int{}
                                                v66 := GenS2{F0: 1, F1: fi42}
                                                if (!(v1 && v1)) {
                                                    v1 = (3.14159 != -2.25)
                                                }
                                                _ = v65
                                                _ = v66
                                            }
                                            _ = fi42
                                            _ = v43
                                        }
                                        v1 = (!v1)
                                        _ = fi41
                                    }
                                    v67 := []int{32768, -9223372036854775808, 1, 9223372036854775807}
                                    if v1 {
                                        v67 = append(v67, int(v22))
                                    } else {
                                        switch 981 {
                                        case -424:
                                            var v68 bool = false
                                            _ = v68
                                        case 9223372036854775807:
                                            if false {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v69 int = 4294967295
                                                    _ = v69
                                                default:
                                                    for fi70 := 0; fi70 < 19; fi70++ {
                                                        v6 = (v6 | v6)
                                                        v3 = (v3 ^ (v3 - v3))
                                                        _ = fi70
                                                    }
                                                }
                                            } else {
                                                for fi71 := 0; fi71 < 12; fi71++ {
                                                    for fi72 := 0; fi72 < 17; fi72++ {
                                                        var v73 uint16 = v4
                                                        if v1 {
                                                            if (!v1) {
                                                                var v74 uint64 = 1
                                                                var v75 float32 = ((1.5 + 3.14159) - 0.1)
                                                                _ = v74
                                                                _ = v75
                                                            }
                                                            if v1 {
                                                                if (v1 || (-119 != int64(v3))) {
                                                                    switch "" {
                                                                    case "hello":
                                                                        var v76 uint16 = 0
                                                                        _ = v76
                                                                    case "foo bar":
                                                                        if true {
                                                                            if len(v67) >= 2 {
                                                                                v67 = v67[1:len(v67)]
                                                                            }
                                                                        }
                                                                    case "":
                                                                        var v77 uint8 = (^(v3))
                                                                        _ = v77
                                                                    default:
                                                                        var v78 int16 = -201
                                                                        _ = v78
                                                                    }
                                                                    v22 = int32(v3)
                                                                } else {
                                                                    var v79 int16 = 1
                                                                    if v1 {
                                                                        fmt.Println(v40)
                                                                        var v80 bool = true
                                                                        _ = v80
                                                                    }
                                                                    _ = v79
                                                                }
                                                                v4 -= v73
                                                            }
                                                            fmt.Println(v6)
                                                        }
                                                        var v81 uint16 = v73
                                                        fmt.Println(v5, v5)
                                                        _ = fi72
                                                        _ = v73
                                                        _ = v81
                                                    }
                                                    _ = fi71
                                                }
                                                v4 -= (v4 / ((42) | 1))
                                                v20 = (v20 / (v20 * 1e10))
                                            }
                                        }
                                        v2[("" + "foo bar")] = uint32(v3)
                                    }
                                    _ = v67
                                } else {
                                    fmt.Println(v20, v5)
                                    var v82 uint16 = (v4 | uint16(v3))
                                    fmt.Println(v5, v20)
                                    _ = v82
                                }
                            } else {
                                v21 = v21
                                v40 -= 420
                                if (v1 || v1) {
                                    delete(v2, "")
                                }
                            }
                        }
                        v83 := v2["hello"]
                        _ = v83
                    } else {
                        fmt.Println(v3, v22)
                    }
                    var v84 int16 = (v5 + 429)
                    _ = v40
                    _ = v84
                } else {
                    for fi85 := 0; fi85 < 9; fi85++ {
                        for fi86 := 0; fi86 < 11; fi86++ {
                            var v87 string = ""
                            delete(v2, "\n\t")
                            var v88 uint = 288
                            _ = fi86
                            _ = v87
                            _ = v88
                        }
                        var v89 int = -239
                        v90 := GenS2{F0: v4, F1: 128}
                        _ = fi85
                        _ = v89
                        _ = v90
                    }
                    v4 += v4
                    switch 666 {
                    case 0:
                        var v91 int8 = -49
                        _ = v91
                    case 32768:
                        if v1 {
                            v92 := []int{int(v6), -846, 1, 32767}
                            _ = v92
                        } else {
                            var v93 bool = (!v1)
                            v6 -= (v6 * (v6 >> ((0) & 63)))
                            _ = v93
                        }
                    case 2147483647:
                        v3 -= v3
                    default:
                        v3 -= (^(v3))
                    }
                }
                _ = v39
            }
            v94 := GenS9{F0: true, F1: 1}
            _ = v94
        }
        for fi95 := 0; fi95 < 7; fi95++ {
            v96 := GenS4{F0: 291, F1: v6}
            v21 = (v21 ^ (v21 * v21))
            _ = fi95
            _ = v96
        }
        _ = v22
    }
    if ((v5 + v5) >= (v5 >> ((uint(v3)) & 63))) {
        fmt.Println(v1, v1, v1)
    }
    v2[""] = 934
    fmt.Println(v4)
    v97 := GenS9{F0: v1, F1: 42}
    var v98 int8 = -101
    v1 = v1
    var v99 uint8 = (v3 << ((2) & 63))
    if (305 != -439) {
        switch 161 {
        case 294:
            for fi100 := 0; fi100 < 20; fi100++ {
                for fi101 := 0; fi101 < 2; fi101++ {
                    switch "foo bar" {
                    case "\n\t":
                        var v102 float32 = ((float32((-2.25 - 0.0)) / float32(v20)) - 1e10)
                        _ = v102
                    case "a":
                        switch int(v4) {
                        case 32767:
                            var v103 uint64 = v21
                            _ = v103
                        case 825:
                            var v104 int64 = 0
                            _ = v104
                        case 42:
                            v2["foo bar"] = 983
                        }
                    case "foo bar":
                        var v105 string = "hello"
                        _ = v105
                    default:
                        v106 := []int{-124, fi100, 32768}
                        _ = v106
                    }
                    fmt.Println(v20)
                    _ = fi101
                }
                _ = fi100
            }
        }
    } else {
        for fi107 := 0; fi107 < 3; fi107++ {
            switch 32768 {
            case -661:
                v108 := GenS10{F0: uint32(v3)}
                _ = v108
            case 0:
                fmt.Println(v1, fi107)
            }
            for fi109 := 0; fi109 < 1; fi109++ {
                v110 := GenS2{F0: v4, F1: fi107}
                _ = fi109
                _ = v110
            }
            if (v20 == 0.1) {
                var v111 string = "hello"
                v112 := GenS11{F0: v111, F1: 65535, F2: 255}
                _ = v111
                _ = v112
            }
            _ = fi107
        }
        v98 -= v6
        v113 := v2[""]
        _ = v113
    }
    v3 = (v99 << ((463) & 63))
    v5 = int16(v4)
    fmt.Println(v20)
    v20 = v20
    if true {
        for fi114 := 0; fi114 < 7; fi114++ {
            fmt.Println(v20)
            _ = fi114
        }
        v115 := GenS12{F0: (v4 >> ((255) & 63))}
        v98 += 42
        _ = v115
    }
    if (!true) {
        for fi116 := 0; fi116 < 11; fi116++ {
            if (false && v1) {
                v4 -= 0
                delete(v2, ("hello" + "\n\t"))
            } else {
                for fi117 := 0; fi117 < 3; fi117++ {
                    v2[(("foo bar" + ("hello" + "a")) + "hello")] = 527
                    fmt.Println(v98, v3, v1)
                    _ = fi117
                }
                switch 2147483647 {
                case 1:
                    fmt.Println(v98, v5)
                case 32767:
                    var v118 float32 = 0.0
                    _ = v118
                default:
                    v20 = (((v20 * -2.25) * v20) * v20)
                }
                fmt.Println(v1, v6, v20)
            }
            v119 := []int{(^(fi116)), (fi116 >> ((0) & 63)), 776}
            if v1 {
                for fi120 := 0; fi120 < 7; fi120++ {
                    v3 = v3
                    if len(v119) >= 2 {
                        v119 = v119[1:len(v119)]
                    }
                    if ("foo bar" == "") {
                        fmt.Println(v5, v6, v5)
                    }
                    switch fi116 {
                    case 127:
                        v4 -= uint16(v21)
                    case 481:
                        v1 = (0 == 42)
                    case -737:
                        switch (("\n\t" + "hello") + ("\n\t" + "")) {
                        case "hello":
                            v3 += v3
                        default:
                            if (0 >= genFn0(879, -144)) {
                                var v121 uint = 0
                                v122 := GenS5{F0: ((0.1 < (0.0 - float32(v20))) && ((-(v5)) < v5))}
                                _ = v121
                                _ = v122
                            }
                        }
                    default:
                        v5 = v5
                    }
                    _ = fi120
                }
                for fi123 := 0; fi123 < 12; fi123++ {
                    var v124 bool = (true && (!false))
                    if v1 {
                        if false {
                            v21 = 4294967295
                            var v125 int8 = v6
                            _ = v125
                        }
                        switch fi123 {
                        case -9223372036854775808:
                            if v124 {
                                v4 += (v4 / (((v4 + 42)) | 1))
                                var v126 int16 = v5
                                _ = v126
                            } else {
                                var v127 bool = v1
                                v128 := GenS2{F0: (v4 + v4), F1: 32767}
                                var v129 int8 = v6
                                fmt.Println(v124, v20)
                                var v130 bool = ((v124 && v1) || ((-(fi116)) > 42))
                                _ = v127
                                _ = v128
                                _ = v129
                                _ = v130
                            }
                        case 32767:
                            if false {
                                if (v20 == -2.25) {
                                    switch "hello" {
                                    case "foo bar":
                                        _ = len(v119)
                                    default:
                                        v1 = false
                                    }
                                    var v131 uint64 = (v21 * v21)
                                    _ = v131
                                }
                            }
                        case 0:
                            var v132 uint = 65535
                            _ = v132
                        default:
                            fmt.Println(v1, v98)
                        }
                    }
                    _ = fi123
                    _ = v124
                }
            }
            _ = fi116
            _ = v119
        }
    } else {
        for fi133 := 0; fi133 < 17; fi133++ {
            var v134 uint8 = v99
            _ = fi133
            _ = v134
        }
        switch 956 {
        case 165:
            if v1 {
                if v1 {
                    v1 = false
                    v5 = v5
                }
                delete(v2, (("" + (("hello" + "hello") + "hello")) + "\n\t"))
            } else {
                if (false && (float32(-2.25) < 1.5)) {
                    if true {
                        if (v1 && v1) {
                            fmt.Println(v5, v20, v99)
                            switch "" {
                            case "a":
                                fmt.Println(v4)
                            }
                            v135 := []int{42, -1, 32768}
                            _ = v135
                        }
                    } else {
                        v136 := GenS10{F0: 4}
                        _ = v136
                    }
                }
                for fi137 := 0; fi137 < 1; fi137++ {
                    _ = len(v2)
                    var v138 uint = uint(v21)
                    _ = fi137
                    _ = v138
                }
                for fi139 := 0; fi139 < 14; fi139++ {
                    _ = len(v2)
                    v140 := GenS13{F0: 198, F1: (^(fi139)), F2: 65535}
                    for fi141 := 0; fi141 < 11; fi141++ {
                        fmt.Println(v3)
                        switch fi139 {
                        case 745:
                            var v142 int16 = (^(v5))
                            _ = v142
                        case -9223372036854775808:
                            fmt.Println(v6)
                            var v143 int64 = -9223372036854775808
                            _ = v143
                        }
                        _ = fi141
                    }
                    _ = fi139
                    _ = v140
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v97
    _ = v98
    _ = v99
}
