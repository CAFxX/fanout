package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int16 }
type GenS2 struct { F0 bool; F1 bool; F2 int16 }
type GenS3 struct { F0 int8; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 int32; F1 uint16; F2 bool }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint32; F1 int32 }
type GenS10 struct { F0 uint; F1 int8; F2 int8 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int8; F1 int }
type GenS13 struct { F0 int64; F1 int16; F2 int }
type GenS14 struct { F0 uint16; F1 string }
type GenS15 struct { F0 uint; F1 int; F2 uint32 }
func genFn0(p0 uint, p1 uint, p2 uint64) int16 {
    if false {
        v6 := GenS1{F0: 0, F1: 9223372036854775807, F2: 0}
        var v7 float32 = -2.25
        _ = v6
        _ = v7
    } else {
        p0 = (p1 << ((144) & 63))
        if (42 >= -145) {
            switch 32767 {
            case -9223372036854775808:
                v8 := GenS1{F0: 65535, F1: 0, F2: -1}
                _ = v8
            case -2147483648:
                fmt.Println(p2)
            }
        } else {
            switch -2147483648 {
            case 127:
                var v9 uint16 = 600
                _ = v9
            case 401:
                var v10 uint16 = 255
                _ = v10
            default:
                p1 += p0
            }
            fmt.Println(p2)
            var v11 bool = (!(-128 < -53))
            _ = v11
        }
    }
    v12 := []int{9223372036854775807, -660, -107, 127}
    var v13 float64 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 int8 = 42
    var v2 int8 = 38
    var v3 uint32 = 65535
    var v4 uint = 0
    v5 := []uint32{v3, (v3 | v3)}
    if (251 >= 1) {
        v3 += v3
        switch -496 {
        case 93:
            v4 = v4
        case -843:
            fmt.Println(v1, v4, v2)
        case 51:
            switch -9223372036854775808 {
            case 387:
                if (((-2.25 - 1.5) / ((1e10 - 0.0) - 0.0)) < 0.0) {
                    switch 93 {
                    case 248:
                        var v14 int16 = 866
                        _ = v14
                    case 127:
                        v3 -= 42
                    case 0:
                        v3 += (v3 - 0)
                    default:
                        fmt.Println(v4)
                    }
                    v15 := GenS1{F0: (v3 << ((65535) & 63)), F1: 128, F2: -32768}
                    _ = v15
                } else {
                    v16 := GenS2{F0: (!("" < "hello")), F1: true, F2: genFn0((^(v4)), (v4 + (v4 >> ((884) & 63))), 2)}
                    if true {
                        v4 -= (v4 & 65535)
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v16
                }
            case -57:
                if len(v5) > 0 {
                    v5[0] = v3
                }
            }
        default:
            v4 -= v4
        }
    } else {
        v17 := GenS2{F0: false, F1: (uint8(v4) < 255), F2: 32767}
        v3 = (v3 >> (((v4 >> (((v4 & (v4 - (^(v4))))) & 63))) & 63))
        _ = v17
    }
    switch "" {
    case "\n\t":
        var v18 uint8 = 42
        _ = v18
    case "":
        var v19 uint16 = 65535
        _ = v19
    }
    if false {
        v2 = v1
        _ = len(v5)
        if (!(-9223372036854775808 <= 42)) {
            var v20 uint16 = 2
            _ = v20
        } else {
            if (true || (false || false)) {
                var v21 bool = true
                if v21 {
                    fmt.Println(v3, v21, v4)
                    fmt.Println(v3, v3, v3)
                    if v21 {
                        var v22 uint16 = 321
                        _ = v22
                    } else {
                        var v23 bool = (0 == 9223372036854775807)
                        _ = v23
                    }
                } else {
                    if false {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v24 := GenS3{F0: v2, F1: (v2 >> (((v4 >> ((v4) & 63))) & 63)), F2: 69}
                        _ = v24
                    } else {
                        _ = len(v5)
                        v21 = true
                    }
                    v21 = (v1 > v1)
                }
                v3 = 2147483647
                _ = v21
            }
            var v25 uint = 543
            for fi26 := 0; fi26 < 4; fi26++ {
                var v27 int16 = genFn0((v4 ^ v25), (^(v25)), 0)
                fmt.Println(fi26, fi26)
                var v28 uint64 = 18446744073709551615
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = v25
        }
    }
    switch ("" + "hello") {
    case "foo bar":
        v29 := GenS2{F0: (2 == 2), F1: (!false), F2: 1}
        _ = v29
    default:
        switch ("a" + "foo bar") {
        case "a":
            v30 := GenS4{F0: 128}
            _ = v30
        default:
            var v31 int64 = 0
            _ = v31
        }
    }
    fmt.Println(v1, v3, v1)
    v32 := GenS5{F0: "a", F1: 1}
    if (!(!true)) {
        var v33 float64 = float64(3.14159)
        _ = v33
    }
    var v34 int16 = 32767
    for fi35 := 0; fi35 < 8; fi35++ {
        if (42 <= 255) {
            if len(v5) > 0 {
                v5[0] = 2147483648
            }
            var v36 uint8 = 127
            _ = v36
        }
        if true {
            var v37 uint8 = 255
            v37 += v37
            _ = v37
        }
        var v38 int = fi35
        _ = fi35
        _ = v38
    }
    if (728 != 640) {
        for fi39 := 0; fi39 < 3; fi39++ {
            v2 += (v2 ^ (v2 * 105))
            _ = fi39
        }
    }
    var v40 uint8 = 128
    switch "a" {
    case "\n\t":
        v41 := GenS6{F0: -2147483648, F1: 782, F2: (false && false)}
        _ = v41
    case "a":
        v42 := GenS7{F0: (^(v3))}
        _ = v42
    default:
        if ((65535 == 537) || true) {
            v40 = 158
            v43 := GenS3{F0: -1, F1: (v2 % ((-68) | 1)), F2: v40}
            for fi44 := 0; fi44 < 3; fi44++ {
                fmt.Println(v1, v34, v34)
                switch ("" + (("\n\t" + "") + "hello")) {
                case "hello":
                    var v45 int8 = 1
                    _ = v45
                }
                _ = fi44
            }
            _ = v43
        }
    }
    for fi46 := 0; fi46 < 14; fi46++ {
        v3 += (v3 | (^(v3)))
        if true {
            switch "" {
            case "a":
                if (true && (661 < 65535)) {
                    v1 = 127
                    switch "" {
                    case "hello":
                        if false {
                            for fi47 := 0; fi47 < 7; fi47++ {
                                var v48 uint = 255
                                v5 = append(v5, (v3 + 4294967295))
                                _ = fi47
                                _ = v48
                            }
                            v49 := GenS8{F0: 427}
                            var v50 int64 = 9223372036854775807
                            _ = v49
                            _ = v50
                        }
                    case "foo bar":
                        v51 := GenS9{F0: 451, F1: -156}
                        _ = v51
                    case "a":
                        v5 = append(v5, v3)
                    }
                } else {
                    for fi52 := 0; fi52 < 15; fi52++ {
                        var v53 int = -170
                        var v54 float32 = 1.5
                        _ = fi52
                        _ = v53
                        _ = v54
                    }
                    v34 += genFn0((^(v4)), v4, 0)
                    if (true && (true || false)) {
                        v2 += (v1 >> ((v4) & 63))
                        var v55 uint = v4
                        if true {
                            v5 = append(v5, v3)
                            v2 = (-(v1))
                            switch fi46 {
                            case 32767:
                                switch "" {
                                case "hello":
                                    var v56 uint32 = 2
                                    _ = v56
                                case "\n\t":
                                    fmt.Println(v2, v1, v4)
                                }
                            case 0:
                                var v57 string = "hello"
                                _ = v57
                            default:
                                v58 := GenS10{F0: (^(v55)), F1: v2, F2: (v2 * (v1 - v2))}
                                _ = v58
                            }
                        } else {
                            _ = len(v5)
                            var v59 bool = false
                            fmt.Println(v3, v4, v3)
                            _ = v59
                        }
                        _ = v55
                    } else {
                        switch fi46 {
                        case 127:
                            if true {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                fmt.Println(v2)
                                var v60 uint16 = 2
                                _ = v60
                            } else {
                                switch ("\n\t" + (("foo bar" + "") + ("foo bar" + ""))) {
                                case "foo bar":
                                    var v61 string = "foo bar"
                                    _ = v61
                                default:
                                    fmt.Println(v34, v40)
                                }
                                for fi62 := 0; fi62 < 13; fi62++ {
                                    _ = len(v5)
                                    var v63 uint = (v4 % (((^(v4))) | 1))
                                    if (-1 >= 0) {
                                        _ = len(v5)
                                        v3 -= 1
                                    }
                                    _ = fi62
                                    _ = v63
                                }
                                v64 := GenS11{F0: 500}
                                _ = v64
                            }
                        default:
                            switch int(v3) {
                            case 128:
                                var v65 uint64 = 2
                                _ = v65
                            default:
                                var v66 int64 = -1
                                _ = v66
                            }
                        }
                        for fi67 := 0; fi67 < 11; fi67++ {
                            fmt.Println(fi67, v3, v4)
                            _ = fi67
                        }
                        if false {
                            var v68 uint8 = v40
                            if false {
                                for fi69 := 0; fi69 < 16; fi69++ {
                                    if true {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        var v70 uint8 = v68
                                        if (-1 > fi46) {
                                            v4 = v4
                                            var v71 uint64 = 1
                                            _ = v71
                                        }
                                        _ = v70
                                    } else {
                                        if ((1 > int64(v68)) && (fi69 >= -1)) {
                                            fmt.Println(fi46, v1, fi46)
                                        }
                                    }
                                    if (true || false) {
                                        if false {
                                            var v72 uint16 = 1
                                            var v73 uint16 = v72
                                            var v74 uint16 = v72
                                            if (443 < 42) {
                                                switch (("" + "") + "") {
                                                case "":
                                                    var v75 int16 = genFn0(v4, (v4 & (v4 - 721)), 42)
                                                    _ = v75
                                                case "hello":
                                                    if len(v5) > 0 {
                                                        v5[0] = v3
                                                    }
                                                case "foo bar":
                                                    v73 = (v74 / ((v73) | 1))
                                                }
                                                if (((!(("foo bar" + ((("hello" + "foo bar") + "a") + "a")) > ("\n\t" + ""))) && true) || true) {
                                                    _ = len(v5)
                                                    var v76 string = ("foo bar" + "hello")
                                                    var v77 uint16 = (v74 ^ v73)
                                                    _ = v76
                                                    _ = v77
                                                }
                                            }
                                            _ = v72
                                            _ = v73
                                            _ = v74
                                        }
                                    }
                                    fmt.Println(v68, fi46)
                                    _ = fi69
                                }
                                v78 := GenS1{F0: (v3 >> ((2) & 63)), F1: int(v1), F2: (v34 * v34)}
                                _ = v78
                            }
                            if true {
                                if false {
                                    var v79 uint32 = (^(v3))
                                    _ = v79
                                }
                                if false {
                                    _ = len(v5)
                                    v80 := GenS3{F0: (v2 + v1), F1: (v1 & (v1 / (((v1 - (v2 + int8(v34)))) | 1))), F2: (v40 << (((v4 & v4)) & 63))}
                                    var v81 string = (("\n\t" + "hello") + "\n\t")
                                    _ = v80
                                    _ = v81
                                } else {
                                    v5 = append(v5, (v3 >> ((uint(v3)) & 63)))
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    v5 = append(v5, 42)
                                }
                                if (-1 < 839) {
                                    for fi82 := 0; fi82 < 6; fi82++ {
                                        if ((1.5 >= 1.5) && ((0.1 <= 1.5) && ((137 < 4294967295) || true))) {
                                            var v83 uint32 = 2147483647
                                            if (v4 < 0) {
                                                var v84 string = "foo bar"
                                                if len(v5) > 0 {
                                                    v5[0] = (v3 >> ((uint(v68)) & 63))
                                                }
                                                _ = v84
                                            } else {
                                                var v85 float64 = 0.0
                                                var v86 uint = 551
                                                _ = v85
                                                _ = v86
                                            }
                                            v40 = v40
                                            fmt.Println(v4, v40)
                                            _ = v83
                                        }
                                        _ = fi82
                                    }
                                    if true {
                                        v4 += (v4 / (((^(v4))) | 1))
                                    } else {
                                        fmt.Println(fi46, v34)
                                    }
                                    if false {
                                        for fi87 := 0; fi87 < 12; fi87++ {
                                            if len(v5) > 0 {
                                                v5[0] = 65535
                                            }
                                            v5 = append(v5, 4294967295)
                                            v34 = (v34 << (((v4 - v4)) & 63))
                                            _ = fi87
                                        }
                                        if (false && true) {
                                            v88 := GenS12{F0: (v2 - v2), F1: 32768}
                                            fmt.Println(v4, v40, v68)
                                            _ = v88
                                        } else {
                                            if len(v5) > 0 {
                                                v5[0] = (v3 ^ 785)
                                            }
                                        }
                                    } else {
                                        switch fi46 {
                                        case 0:
                                            _ = len(v5)
                                        default:
                                            var v89 int64 = 11
                                            _ = v89
                                        }
                                        v90 := GenS3{F0: v2, F1: (v2 - (v2 - v1)), F2: 23}
                                        fmt.Println(v1)
                                        _ = v90
                                    }
                                } else {
                                    fmt.Println(v3)
                                    var v91 uint8 = (v68 >> ((46) & 63))
                                    v91 = v40
                                    _ = v91
                                }
                            } else {
                                var v92 uint16 = 1
                                var v93 int32 = 846
                                var v94 float32 = (((0.0 - float32((1e10 + (-2.25 * 0.0)))) * 0.0) - (1.5 + (3.14159 + 1e10)))
                                _ = v92
                                _ = v93
                                _ = v94
                            }
                            _ = v68
                        }
                    }
                }
            }
            for fi95 := 0; fi95 < 13; fi95++ {
                switch 32767 {
                case 32768:
                    switch -9223372036854775808 {
                    case 32768:
                        v3 = 932
                    case -9223372036854775808:
                        if true {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            fmt.Println(v34, v4)
                            var v96 bool = true
                            v34 -= (v34 ^ 1000)
                            if ((-(fi95)) >= fi95) {
                                var v97 int32 = 222
                                _ = len(v5)
                                v98 := GenS13{F0: 723, F1: genFn0(v4, 397, 18446744073709551615), F2: fi46}
                                _ = v97
                                _ = v98
                            } else {
                                v40 += 181
                                var v99 uint32 = v3
                                _ = v99
                            }
                            v2 = v1
                            _ = v96
                        } else {
                            switch -539 {
                            case 2147483647:
                                var v100 int32 = 0
                                _ = v100
                            case 9223372036854775807:
                                if (!false) {
                                    _ = len(v5)
                                } else {
                                    if (v3 != v3) {
                                        if false {
                                            v40 -= v40
                                        } else {
                                            v4 -= v4
                                            var v101 int16 = v34
                                            fmt.Println(fi46)
                                            _ = v101
                                        }
                                    }
                                    if ((!true) || (42 != 956)) {
                                        v102 := GenS14{F0: 2, F1: "\n\t"}
                                        _ = v102
                                    } else {
                                        v34 += genFn0((v4 % (((v4 + v4)) | 1)), v4, 42)
                                        v103 := GenS7{F0: v3}
                                        _ = v103
                                    }
                                    v1 = (v1 + (v2 - (-(v1))))
                                }
                                v40 -= 2
                                switch fi95 {
                                case 33:
                                    switch (fi95 << ((970) & 63)) {
                                    case 2147483647:
                                        var v104 int64 = 0
                                        _ = v104
                                    case -2147483648:
                                        var v105 uint32 = v3
                                        _ = v105
                                    case -9223372036854775808:
                                        switch (fi95 / (((^(fi46))) | 1)) {
                                        case 32767:
                                            var v106 float64 = 3.14159
                                            _ = v106
                                        }
                                    default:
                                        var v107 uint64 = 1
                                        _ = v107
                                    }
                                }
                                v40 += (v40 / ((v40) | 1))
                            case 598:
                                if (0.0 >= (float32(1e10) * (0.0 + 1.5))) {
                                    v108 := GenS10{F0: v4, F1: v1, F2: v2}
                                    v109 := GenS15{F0: v4, F1: fi95, F2: (v3 / ((v3) | 1))}
                                    v4 += 18446744073709551615
                                    _ = v108
                                    _ = v109
                                }
                            default:
                                v110 := GenS9{F0: v3, F1: -374}
                                _ = v110
                            }
                            switch (fi95 >> (((v4 | v4)) & 63)) {
                            case 32768:
                                var v111 int64 = 9223372036854775807
                                _ = v111
                            case 9223372036854775807:
                                var v112 int32 = 0
                                _ = v112
                            default:
                                switch "\n\t" {
                                case "a":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                case "foo bar":
                                    v113 := GenS13{F0: 9223372036854775807, F1: 1, F2: 763}
                                    _ = v113
                                case "\n\t":
                                    var v114 int32 = 42
                                    _ = v114
                                }
                            }
                            if ((!(2 >= 700)) && false) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        }
                    case 959:
                        v4 = (^(v4))
                    default:
                        if true {
                            switch fi46 {
                            case 4294967295:
                                v115 := GenS13{F0: 42, F1: 349, F2: fi46}
                                _ = v115
                            case 2147483647:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            var v116 bool = (!((v2 << ((676) & 63)) > 1))
                            var v117 uint16 = 65535
                            if (fi95 > fi95) {
                                var v118 int = fi46
                                _ = v118
                            } else {
                                v119 := GenS12{F0: v2, F1: 42}
                                var v120 uint = uint(v3)
                                if (uint16(v120) >= (v117 % (((^(v117))) | 1))) {
                                    switch (-(fi46)) {
                                    case 4294967295:
                                        var v121 uint8 = v40
                                        _ = v121
                                    case 32768:
                                        if v116 {
                                            if v116 {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            } else {
                                                v117 = (v117 | (v117 % (((v117 * (v117 ^ 65535))) | 1)))
                                            }
                                        } else {
                                            v40 = uint8(v120)
                                            var v122 int = -2147483648
                                            _ = v122
                                        }
                                    case -77:
                                        if true {
                                            v2 += -1
                                            v117 = (v117 / ((v117) | 1))
                                            var v123 int32 = 42
                                            _ = v123
                                        } else {
                                            v2 = v2
                                            var v124 uint64 = uint64(v120)
                                            fmt.Println(v4, v2)
                                            _ = v124
                                        }
                                    }
                                    v125 := GenS5{F0: "a", F1: 42}
                                    _ = v125
                                } else {
                                    var v126 int8 = (v2 - v2)
                                    var v127 int32 = -700
                                    var v128 uint32 = v3
                                    _ = v126
                                    _ = v127
                                    _ = v128
                                }
                                _ = v119
                                _ = v120
                            }
                            _ = v116
                            _ = v117
                        } else {
                            v34 += (v34 % ((877) | 1))
                        }
                    }
                }
                if true {
                    switch (fi46 | 32767) {
                    case 42:
                        switch fi46 {
                        case 208:
                            switch (("foo bar" + ((("" + "foo bar") + ("hello" + ("foo bar" + (("\n\t" + ("a" + "\n\t")) + "foo bar")))) + "foo bar")) + ("a" + "\n\t")) {
                            case "foo bar":
                                v1 = v1
                            case "hello":
                                if (((-(v2)) != int8(v4)) || true) {
                                    var v129 int8 = (v2 + v1)
                                    if true {
                                        fmt.Println(v2, v34, v3)
                                        if (!true) {
                                            v130 := GenS7{F0: (^(v3))}
                                            v5 = append(v5, v3)
                                            _ = v130
                                        } else {
                                            var v131 int8 = v129
                                            v40 -= 42
                                            _ = v131
                                        }
                                        var v132 bool = true
                                        v5 = append(v5, v3)
                                        v3 += (v3 * v3)
                                        _ = v132
                                    }
                                    v133 := GenS15{F0: (^(v4)), F1: (-(fi46)), F2: 4294967295}
                                    v134 := GenS13{F0: -950, F1: v34, F2: 32768}
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v129
                                    _ = v133
                                    _ = v134
                                }
                            default:
                                _ = len(v5)
                            }
                        }
                    case 9223372036854775807:
                        fmt.Println(v3, v40)
                    case 127:
                        if (!true) {
                            var v135 int = -336
                            fmt.Println(v34)
                            _ = v135
                        }
                    default:
                        var v136 int16 = v34
                        _ = v136
                    }
                    var v137 float64 = (float64(1.5) / (3.14159 + (3.14159 / 1.5)))
                    _ = v137
                } else {
                    if true {
                        var v138 uint = (v4 << (((v4 & v4)) & 63))
                        _ = v138
                    } else {
                        if len(v5) > 0 {
                            v5[0] = (^(v3))
                        }
                        if false {
                            var v139 uint8 = v40
                            _ = len(v5)
                            if ((^(v1)) > (-(v1))) {
                                var v140 int32 = 42
                                _ = v140
                            }
                            _ = v139
                        }
                        if ((false || (((v3 >> ((v4) & 63)) > v3) && (false && true))) || true) {
                            v141 := GenS7{F0: (v3 | v3)}
                            _ = v141
                        } else {
                            var v142 uint32 = v3
                            v142 = v142
                            _ = v142
                        }
                        v4 = 1
                    }
                }
                _ = fi95
            }
            fmt.Println(fi46, v2, v4)
        } else {
            fmt.Println(v34, v40)
            v1 += (v1 ^ 1)
            v2 = v2
        }
        v2 = (v1 & (v1 | 57))
        _ = fi46
    }
    switch "hello" {
    case "a":
        if (87 < (v1 * v1)) {
            var v143 int32 = 36
            var v144 uint = (v4 / (((v4 % ((v4) | 1))) | 1))
            _ = v143
            _ = v144
        }
    default:
        v34 -= 732
    }
    for fi145 := 0; fi145 < 9; fi145++ {
        v3 -= v3
        _ = fi145
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v4
    _ = v40
    _ = v5
}
