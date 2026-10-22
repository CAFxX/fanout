package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint64; F2 int64 }
type GenS3 struct { F0 bool; F1 uint16; F2 string }
type GenS4 struct { F0 uint32; F1 int8; F2 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64; F1 int; F2 int }
type GenS7 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint16; F1 string; F2 string }
type GenS10 struct { F0 int64; F1 int8 }
type GenS11 struct { F0 uint16; F1 uint16; F2 uint }
type GenS12 struct { F0 bool; F1 uint; F2 uint16 }
type GenS13 struct { F0 int32; F1 bool; F2 uint32 }
type GenS14 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS15 struct { F0 uint16 }
func genFn0(p0 uint8, p1 uint64, p2 int16) int16 {
    v5 := []int{-1, -121, 1, 960}
    switch "a" {
    case "hello":
        for fi6 := 0; fi6 < 18; fi6++ {
            p1 = 18446744073709551615
            var v7 uint8 = (p0 | p0)
            _ = fi6
            _ = v7
        }
    case "\n\t":
        if true {
            if true {
                p0 -= p0
                fmt.Println(p2)
                switch "hello" {
                case "a":
                    fmt.Println(p2, p2)
                case "\n\t":
                    var v8 uint16 = 483
                    _ = v8
                default:
                    var v9 uint64 = p1
                    _ = v9
                }
            } else {
                switch -2147483648 {
                case 810:
                    var v10 uint16 = 42
                    _ = v10
                case -1:
                    var v11 int8 = 42
                    _ = v11
                default:
                    v5 = append(v5, 1)
                }
                var v12 uint = 0
                _ = v12
            }
            if (6 >= 32767) {
                fmt.Println(p2, p2)
                var v13 uint = 401
                _ = v13
            } else {
                var v14 uint8 = p0
                if (!(65535 == 255)) {
                    var v15 int = 1
                    _ = v15
                }
                _ = v14
            }
            p2 += (p2 & (p2 << ((266) & 63)))
        } else {
            v16 := GenS1{F0: 42}
            p1 = p1
            var v17 uint = 2
            _ = v16
            _ = v17
        }
    case "foo bar":
        switch "a" {
        case "foo bar":
            var v18 int32 = 1
            _ = v18
        case "":
            var v19 int64 = 1
            _ = v19
        default:
            for fi20 := 0; fi20 < 9; fi20++ {
                v21 := GenS1{F0: 0}
                _ = fi20
                _ = v21
            }
        }
    }
    fmt.Println(p2)
    v22 := GenS2{F0: -651, F1: (p1 - 0), F2: 69}
    var v23 uint = 867
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    _ = v5
    return p2
}
func genFn1(p0 int, p1 uint16, p2 int8) int16 {
    var v24 uint64 = 0
    var v25 string = ""
    var v26 bool = true
    switch ("" + "\n\t") {
    case "":
        p1 = p1
    default:
        v27 := GenS3{F0: false, F1: 32767, F2: (v25 + "\n\t")}
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    _ = v26
    return 1
}

func main() {
    var v1 float64 = 0.0
    var v2 int32 = 2147483647
    var v3 uint = 591
    var v4 int16 = 42
    v28 := []int{}
    v2 = v2
    v2 += 0
    v29 := GenS3{F0: false, F1: 42, F2: (("a" + "hello") + (("" + ("foo bar" + ("foo bar" + ""))) + ""))}
    if false {
        var v30 int = -2147483648
        switch (v30 >> ((v3) & 63)) {
        case 9223372036854775807:
            v31 := GenS2{F0: v4, F1: 1, F2: 9223372036854775807}
            _ = v31
        default:
            fmt.Println(v1)
        }
        switch "hello" {
        case "\n\t":
            switch (v30 * -1) {
            case 127:
                v32 := GenS4{F0: 294, F1: 1, F2: v3}
                _ = v32
            case -121:
                var v33 int = (^(v30))
                _ = v33
            case 9223372036854775807:
                for fi34 := 0; fi34 < 13; fi34++ {
                    for fi35 := 0; fi35 < 16; fi35++ {
                        v1 = v1
                        if (4294967295 <= 4294967295) {
                            var v36 uint = (v3 + 18446744073709551615)
                            v1 -= (1e10 * 0.0)
                            v3 -= (v3 % ((v36) | 1))
                            _ = v36
                        }
                        fmt.Println(v1)
                        _ = fi35
                    }
                    if ("a" > "foo bar") {
                        var v37 int64 = 42
                        _ = v37
                    }
                    _ = fi34
                }
            default:
                var v38 int32 = 2147483647
                _ = v38
            }
        }
        _ = v30
    }
    switch 127 {
    case 127:
        v39 := GenS3{F0: false, F1: 2, F2: "hello"}
        _ = v39
    case -9223372036854775808:
        fmt.Println(v3, v4)
    default:
        if true {
            v4 = (v4 << (((v3 / (((^(v3))) | 1))) & 63))
            _ = len(v28)
            v40 := GenS5{F0: -1}
            _ = v40
        }
    }
    if false {
        v28 = append(v28, 4294967295)
        switch 1 {
        case -68:
            v2 += 0
        default:
            switch -2147483648 {
            case -986:
                v3 -= (v3 | (v3 | v3))
            case -103:
                v41 := GenS6{F0: 0, F1: 128, F2: -218}
                _ = v41
            case 32768:
                if ((v3 >> ((v3) & 63)) <= v3) {
                    switch "hello" {
                    case "hello":
                        switch 127 {
                        case 9223372036854775807:
                            v2 = (v2 * v2)
                        case -9223372036854775808:
                            var v42 int8 = 1
                            _ = v42
                        case 935:
                            var v43 int = 127
                            _ = v43
                        default:
                            for fi44 := 0; fi44 < 11; fi44++ {
                                var v45 float32 = 1e10
                                fmt.Println(v3, fi44)
                                var v46 float64 = ((v1 / 1.5) + 3.14159)
                                _ = fi44
                                _ = v45
                                _ = v46
                            }
                        }
                    case "":
                        v3 = (^(v3))
                    case "\n\t":
                        switch "foo bar" {
                        case "\n\t":
                            var v47 uint32 = 2
                            _ = v47
                        case "hello":
                            if true {
                                fmt.Println(v3, v1)
                                if (float32((-2.25 - 1.5)) <= (((-2.25 / 1.5) - 1e10) / 3.14159)) {
                                    var v48 float32 = float32(float64((float32((0.1 - v1)) - 0.1)))
                                    v49 := GenS5{F0: -9223372036854775808}
                                    _ = v48
                                    _ = v49
                                } else {
                                    v3 = (v3 % ((2) | 1))
                                    v50 := GenS4{F0: 0, F1: -128, F2: v3}
                                    var v51 int = 42
                                    _ = v50
                                    _ = v51
                                }
                            }
                        case "foo bar":
                            v1 = ((v1 - (0.0 * v1)) * float64(1.5))
                        }
                    default:
                        if len(v28) > 0 {
                            v28[0] = 2147483647
                        }
                    }
                    for fi52 := 0; fi52 < 11; fi52++ {
                        var v53 int8 = 1
                        for fi54 := 0; fi54 < 2; fi54++ {
                            v2 = -1
                            v4 = 32767
                            v3 -= v3
                            _ = fi54
                        }
                        _ = fi52
                        _ = v53
                    }
                } else {
                    var v55 string = "a"
                    if false {
                        fmt.Println(v2, v55, v1)
                        v56 := GenS6{F0: 42, F1: 32768, F2: 2147483647}
                        switch v55 {
                        case "":
                            if false {
                                _ = len(v28)
                                for fi57 := 0; fi57 < 16; fi57++ {
                                    if len(v28) >= 2 {
                                        v28 = v28[1:len(v28)]
                                    }
                                    _ = fi57
                                }
                                v28 = append(v28, 541)
                            } else {
                                var v58 int16 = v4
                                if len(v28) > 0 {
                                    v28[0] = -357
                                }
                                _ = v58
                            }
                        }
                        _ = v56
                    } else {
                        var v59 int64 = 9223372036854775807
                        _ = v59
                    }
                    _ = v55
                }
            }
        }
        for fi60 := 0; fi60 < 14; fi60++ {
            if (v4 != v4) {
                switch (fi60 * (fi60 >> ((v3) & 63))) {
                case -108:
                    var v61 int16 = -32768
                    _ = v61
                default:
                    v28 = append(v28, fi60)
                }
                if (!(true && true)) {
                    var v62 float32 = 0.1
                    _ = v62
                } else {
                    for fi63 := 0; fi63 < 18; fi63++ {
                        switch int(v2) {
                        case 4294967295:
                            if (!true) {
                                v28 = append(v28, fi63)
                                switch "\n\t" {
                                case "a":
                                    v64 := GenS1{F0: v2}
                                    _ = v64
                                }
                                v1 = v1
                                v4 += genFn0(160, 18446744073709551615, (v4 + v4))
                            }
                        case 32768:
                            v1 = v1
                        default:
                            var v65 int8 = -1
                            _ = v65
                        }
                        _ = fi63
                    }
                    v66 := GenS7{F0: v2, F1: 0, F2: uint32(v3)}
                    _ = v66
                }
            } else {
                if true {
                    switch (fi60 & 1) {
                    case -2147483648:
                        for fi67 := 0; fi67 < 6; fi67++ {
                            if true {
                                _ = len(v28)
                            }
                            _ = fi67
                        }
                    case 32767:
                        for fi68 := 0; fi68 < 10; fi68++ {
                            fmt.Println(v2, fi60)
                            _ = fi68
                        }
                    case 429:
                        v69 := GenS6{F0: 1, F1: (fi60 - int(v4)), F2: fi60}
                        _ = v69
                    default:
                        v4 = (v4 & v4)
                    }
                    if ((!(false || (255 >= (v3 >> ((v3) & 63))))) && (-9223372036854775808 >= 507)) {
                        if false {
                            var v70 float64 = 0.0
                            v71 := GenS8{F0: 127}
                            if ((v2 - (v2 / ((v2) | 1))) != 57) {
                                if len(v28) >= 2 {
                                    v28 = v28[1:len(v28)]
                                }
                            } else {
                                for fi72 := 0; fi72 < 8; fi72++ {
                                    var v73 uint16 = 290
                                    _ = fi72
                                    _ = v73
                                }
                                for fi74 := 0; fi74 < 10; fi74++ {
                                    fmt.Println(v3)
                                    v75 := GenS8{F0: 100}
                                    if false {
                                        fmt.Println(v1)
                                        var v76 uint16 = uint16(v3)
                                        _ = v76
                                    }
                                    if false {
                                        var v77 int8 = 100
                                        switch (("\n\t" + "a") + ("" + ("hello" + "foo bar"))) {
                                        case "foo bar":
                                            var v78 int = fi60
                                            _ = v78
                                        }
                                        _ = v77
                                    }
                                    v79 := GenS5{F0: 205}
                                    _ = fi74
                                    _ = v75
                                    _ = v79
                                }
                                v70 = -2.25
                            }
                            _ = v70
                            _ = v71
                        } else {
                            v2 = v2
                            v3 = v3
                            v4 -= (v4 & (v4 & -32768))
                        }
                        v80 := GenS9{F0: 483, F1: ("a" + "hello"), F2: "hello"}
                        if len(v28) > 0 {
                            v28[0] = -2147483648
                        }
                        _ = v80
                    }
                } else {
                    switch "hello" {
                    case "\n\t":
                        for fi81 := 0; fi81 < 11; fi81++ {
                            _ = len(v28)
                            _ = fi81
                        }
                    case "":
                        v3 += 852
                    case "hello":
                        for fi82 := 0; fi82 < 14; fi82++ {
                            v83 := GenS3{F0: false, F1: 433, F2: "hello"}
                            _ = fi82
                            _ = v83
                        }
                    }
                    v3 = v3
                    v3 = v3
                }
                for fi84 := 0; fi84 < 19; fi84++ {
                    v2 -= (v2 + v2)
                    var v85 int32 = (-(v2))
                    v86 := GenS10{F0: 9223372036854775807, F1: 42}
                    _ = fi84
                    _ = v85
                    _ = v86
                }
                switch (fi60 * fi60) {
                case -356:
                    var v87 int8 = -121
                    _ = v87
                }
            }
            if (1 <= 410) {
                if (1 != 42) {
                    v28 = append(v28, (fi60 / ((fi60) | 1)))
                    v28 = append(v28, fi60)
                    for fi88 := 0; fi88 < 3; fi88++ {
                        v2 += (v2 & -207)
                        _ = fi88
                    }
                } else {
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v3, v1)
                    default:
                        v89 := GenS9{F0: 42, F1: "a", F2: ((("hello" + "hello") + "\n\t") + "a")}
                        _ = v89
                    }
                }
            } else {
                for fi90 := 0; fi90 < 9; fi90++ {
                    var v91 uint = 2
                    switch "foo bar" {
                    case "foo bar":
                        if true {
                            switch 9223372036854775807 {
                            case 229:
                                if len(v28) >= 2 {
                                    v28 = v28[1:len(v28)]
                                }
                            case 927:
                                fmt.Println(fi60, v3, fi90)
                            case 4294967295:
                                var v92 uint8 = 71
                                _ = v92
                            default:
                                if len(v28) >= 2 {
                                    v28 = v28[1:len(v28)]
                                }
                            }
                        }
                    case "a":
                        v93 := GenS1{F0: (v2 + (-(v2)))}
                        _ = v93
                    case "":
                        v94 := GenS7{F0: v2, F1: 252, F2: 324}
                        _ = v94
                    default:
                        fmt.Println(fi90, v91)
                    }
                    _ = fi90
                    _ = v91
                }
                for fi95 := 0; fi95 < 14; fi95++ {
                    var v96 uint16 = 900
                    _ = fi95
                    _ = v96
                }
            }
            _ = fi60
        }
    } else {
        v4 -= v4
        _ = len(v28)
        fmt.Println(v4)
    }
    v97 := GenS7{F0: v2, F1: 128, F2: 848}
    for fi98 := 0; fi98 < 12; fi98++ {
        var v99 int32 = (v2 + (v2 % ((v2) | 1)))
        if false {
            if len(v28) >= 2 {
                v28 = v28[1:len(v28)]
            }
            var v100 uint8 = 128
            _ = v100
        } else {
            var v101 int64 = -76
            var v102 int8 = 42
            v103 := GenS11{F0: 42, F1: 2, F2: v3}
            _ = v101
            _ = v102
            _ = v103
        }
        if false {
            v104 := GenS12{F0: false, F1: (v3 - v3), F2: 0}
            v28 = append(v28, -812)
            if ((371 < 65535) && (v4 > int16(fi98))) {
                if false {
                    for fi105 := 0; fi105 < 17; fi105++ {
                        fmt.Println(v2)
                        v1 += (v1 * (3.14159 * 1.5))
                        if (!true) {
                            v2 -= (v2 << ((0) & 63))
                            var v106 int = fi105
                            if (123 > 1) {
                                v4 -= 154
                                switch (fi105 + v106) {
                                case 32767:
                                    if (0 > v4) {
                                        if (253 <= 153) {
                                            var v107 uint32 = 0
                                            _ = v107
                                        }
                                    } else {
                                        v99 += v99
                                        v108 := GenS10{F0: -9223372036854775808, F1: -1}
                                        _ = v108
                                    }
                                case 4294967295:
                                    var v109 int8 = 42
                                    _ = v109
                                default:
                                    var v110 string = ("foo bar" + "")
                                    _ = v110
                                }
                                var v111 uint16 = 2
                                _ = v111
                            } else {
                                var v112 uint64 = 4294967295
                                _ = v112
                            }
                            _ = v106
                        }
                        _ = fi105
                    }
                    for fi113 := 0; fi113 < 3; fi113++ {
                        _ = len(v28)
                        _ = fi113
                    }
                    v114 := GenS2{F0: 1000, F1: 1, F2: -73}
                    _ = v114
                }
                var v115 bool = ((0 <= -9223372036854775808) || (!true))
                for fi116 := 0; fi116 < 8; fi116++ {
                    fmt.Println(fi116, v1)
                    switch fi116 {
                    case 42:
                        var v117 string = ""
                        _ = v117
                    default:
                        v118 := GenS11{F0: 65535, F1: 32767, F2: (v3 + v3)}
                        _ = v118
                    }
                    var v119 uint64 = 18446744073709551615
                    _ = fi116
                    _ = v119
                }
                _ = v115
            }
            _ = v104
        }
        _ = fi98
        _ = v99
    }
    for fi120 := 0; fi120 < 11; fi120++ {
        v28 = append(v28, fi120)
        if false {
            v121 := GenS13{F0: (v2 * (v2 % (((v2 * 206)) | 1))), F1: false, F2: 2147483647}
            _ = v121
        } else {
            fmt.Println(fi120)
            fmt.Println(v2, v3)
            v4 -= (v4 - v4)
        }
        v122 := GenS14{F0: 2, F1: 18446744073709551615, F2: 1}
        _ = fi120
        _ = v122
    }
    fmt.Println(v3, v3, v3)
    var v123 float32 = (3.14159 / float32(v1))
    v124 := GenS15{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v123:", v123)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v123
    _ = v124
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v97
}
