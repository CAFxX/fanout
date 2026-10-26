package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 uint; F2 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 uint64; F2 string }
type GenS7 struct { F0 uint32; F1 int16; F2 int8 }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 uint8; F1 int32; F2 int }
type GenS10 struct { F0 uint; F1 uint8 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 string; F1 int8 }
type GenS13 struct { F0 uint8; F1 int32 }
type GenS14 struct { F0 uint8; F1 int64; F2 int }
func genFn0(p0 bool) uint32 {
    v7 := []int{-516, 1, 32767, -136}
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 int = (fi9 | -1)
            var v11 float32 = (((3.14159 + 3.14159) + 1e10) - 1.5)
            if false {
                var v12 float32 = float32(float64(v11))
                v13 := GenS1{F0: "a", F1: 4294967295, F2: 1}
                v14 := GenS2{F0: 1000}
                _ = v12
                _ = v13
                _ = v14
            } else {
                var v15 int = fi9
                if true {
                    switch "\n\t" {
                    case "a":
                        v16 := GenS1{F0: "foo bar", F1: 4294967295, F2: 125}
                        _ = v16
                    case "hello":
                        v17 := GenS1{F0: "hello", F1: 2, F2: 255}
                        _ = v17
                    default:
                        var v18 float64 = -2.25
                        _ = v18
                    }
                    var v19 uint16 = 65535
                    _ = v19
                }
                if (-226 > 9223372036854775807) {
                    var v20 uint8 = 42
                    switch (("a" + ((((("foo bar" + "") + "a") + "\n\t") + "\n\t") + "a")) + ("foo bar" + ("" + "\n\t"))) {
                    case "foo bar":
                        if (!(("hello" + "\n\t") < "a")) {
                            var v21 uint32 = 423
                            v11 += v11
                            _ = v21
                        }
                    case "\n\t":
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(v10)
                        case "foo bar":
                            v10 = (fi9 << ((255) & 63))
                        default:
                            switch (v15 & (v10 & -528)) {
                            case 4294967295:
                                if ("" == "hello") {
                                    v7 = append(v7, fi9)
                                    v22 := GenS1{F0: "\n\t", F1: 4294967295, F2: 255}
                                    fmt.Println(p0)
                                    _ = v22
                                }
                            case 1:
                                if (!true) {
                                    var v23 uint64 = 1
                                    _ = v23
                                }
                            case -9223372036854775808:
                                var v24 int8 = -80
                                _ = v24
                            }
                        }
                    case "hello":
                        var v25 float32 = -2.25
                        v26 := GenS2{F0: -64}
                        _ = v25
                        _ = v26
                    }
                    _ = v20
                }
                var v27 uint16 = 413
                _ = v15
                _ = v27
            }
            fmt.Println(v11, v11)
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v28 int = (fi9 * fi9)
            _ = fi9
            _ = v10
            _ = v11
            _ = v28
        }
        _ = fi8
    }
    if p0 {
        switch -859 {
        case 4294967295:
            p0 = (!(p0 && ("hello" == "")))
        case 0:
            var v29 uint64 = 2
            _ = v29
        default:
            p0 = p0
        }
    }
    if p0 {
        switch 127 {
        case 703:
            v30 := GenS1{F0: "foo bar", F1: 0, F2: 255}
            _ = v30
        case 4294967295:
            var v31 string = (("foo bar" + "hello") + (("" + "\n\t") + ("hello" + ("a" + ("hello" + ("foo bar" + "foo bar"))))))
            _ = v31
        case 1:
            fmt.Println(p0)
        default:
            v32 := GenS3{F0: 42}
            _ = v32
        }
        if p0 {
            p0 = p0
            switch ("" + "a") {
            case "":
                v33 := GenS3{F0: 127}
                _ = v33
            case "foo bar":
                v7 = append(v7, 677)
            case "\n\t":
                v34 := GenS3{F0: 120}
                _ = v34
            }
        } else {
            var v35 uint32 = 0
            for fi36 := 0; fi36 < 1; fi36++ {
                fmt.Println(v35, fi36)
                switch "foo bar" {
                case "\n\t":
                    _ = len(v7)
                case "":
                    var v37 uint16 = 0
                    _ = v37
                case "foo bar":
                    v35 += v35
                }
                _ = fi36
            }
            switch ((("\n\t" + "") + "hello") + "foo bar") {
            case "":
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            case "hello":
                for fi38 := 0; fi38 < 1; fi38++ {
                    v35 += (v35 * 42)
                    _ = fi38
                }
            case "foo bar":
                if (false && p0) {
                    if ("foo bar" < "hello") {
                        if len(v7) > 0 {
                            v7[0] = -9223372036854775808
                        }
                    } else {
                        switch 127 {
                        case 4294967295:
                            _ = len(v7)
                        case 32767:
                            for fi39 := 0; fi39 < 7; fi39++ {
                                v35 += 2147483647
                                _ = fi39
                            }
                        case 609:
                            if p0 {
                                var v40 int = 0
                                switch v40 {
                                case 32768:
                                    for fi41 := 0; fi41 < 20; fi41++ {
                                        p0 = (p0 && false)
                                        v42 := GenS4{F0: (fi41 / ((v40) | 1)), F1: 121, F2: -1}
                                        _ = fi41
                                        _ = v42
                                    }
                                case -9223372036854775808:
                                    for fi43 := 0; fi43 < 1; fi43++ {
                                        p0 = p0
                                        _ = fi43
                                    }
                                case 4294967295:
                                    _ = len(v7)
                                default:
                                    if p0 {
                                        var v44 uint64 = 468
                                        for fi45 := 0; fi45 < 7; fi45++ {
                                            if p0 {
                                                v40 += 0
                                                fmt.Println(fi45, v35, v40)
                                            } else {
                                                var v46 float32 = (3.14159 / -2.25)
                                                _ = v46
                                            }
                                            _ = len(v7)
                                            var v47 float64 = 3.14159
                                            _ = fi45
                                            _ = v47
                                        }
                                        var v48 int64 = 0
                                        _ = v44
                                        _ = v48
                                    } else {
                                        v35 = 2147483647
                                    }
                                }
                                _ = v40
                            }
                        default:
                            v49 := GenS5{F0: 1}
                            _ = v49
                        }
                        var v50 int = 32767
                        var v51 float32 = 3.14159
                        _ = v50
                        _ = v51
                    }
                    var v52 int32 = 2147483647
                    _ = v52
                }
            default:
                var v53 bool = (p0 && p0)
                _ = v53
            }
            _ = v35
        }
    } else {
        v54 := GenS6{F0: 9223372036854775807, F1: 1, F2: "foo bar"}
        if ("" != ("\n\t" + (("hello" + (((((("a" + (("a" + "\n\t") + (((("foo bar" + ("a" + ("a" + "foo bar"))) + "foo bar") + "hello") + "\n\t"))) + "hello") + ("a" + "\n\t")) + ("hello" + (("\n\t" + "a") + "hello"))) + ("\n\t" + "")) + "foo bar")) + ("foo bar" + ("a" + "hello"))))) {
            p0 = (!p0)
            switch -9223372036854775808 {
            case -2147483648:
                p0 = (p0 && p0)
            }
        } else {
            if p0 {
                fmt.Println(p0)
                for fi55 := 0; fi55 < 6; fi55++ {
                    var v56 float32 = (3.14159 / 1.0)
                    var v57 int64 = -135
                    v58 := GenS3{F0: 96}
                    _ = fi55
                    _ = v56
                    _ = v57
                    _ = v58
                }
                v59 := GenS2{F0: 1000}
                _ = v59
            } else {
                var v60 uint32 = 2147483647
                if (!p0) {
                    for fi61 := 0; fi61 < 5; fi61++ {
                        p0 = ((^(v60)) == v60)
                        _ = fi61
                    }
                }
                var v62 float32 = (1e10 - (((((((3.14159 * 1.5) / 0.1) * ((1e10 * -2.25) - 3.14159)) / 0.1) / 1e10) / 3.14159) / -2.25))
                _ = v60
                _ = v62
            }
            switch ("hello" + ("" + ("a" + "foo bar"))) {
            case "hello":
                fmt.Println(p0)
            case "":
                for fi63 := 0; fi63 < 16; fi63++ {
                    var v64 bool = (false && p0)
                    switch (fi63 + (fi63 | 42)) {
                    case 198:
                        if len(v7) > 0 {
                            v7[0] = 4294967295
                        }
                    }
                    p0 = p0
                    _ = fi63
                    _ = v64
                }
            case "a":
                if (0.0 >= float32((3.14159 - 0.1))) {
                    for fi65 := 0; fi65 < 12; fi65++ {
                        v7 = append(v7, fi65)
                        fmt.Println(p0, p0)
                        v66 := GenS3{F0: 127}
                        _ = fi65
                        _ = v66
                    }
                }
            default:
                for fi67 := 0; fi67 < 7; fi67++ {
                    if p0 {
                        p0 = (0 <= 42)
                    } else {
                        for fi68 := 0; fi68 < 5; fi68++ {
                            if (-1 > -1) {
                                p0 = ((60 > 1) && (!p0))
                                var v69 uint = 204
                                _ = v69
                            } else {
                                if true {
                                    var v70 int64 = -288
                                    v71 := GenS7{F0: 2147483648, F1: -119, F2: -1}
                                    if (!((!((42 == -87) || (-551 == (v70 / (((v70 | (v70 & (^(v70))))) | 1))))) && (true || ((127 == 54) && ((2147483648 > 42) || (65 == -96)))))) {
                                        _ = len(v7)
                                        v72 := GenS3{F0: 138}
                                        _ = v72
                                    }
                                    _ = v70
                                    _ = v71
                                }
                                v73 := GenS4{F0: fi67, F1: 255, F2: 42}
                                p0 = p0
                                var v74 string = ("" + "a")
                                if (!(0 >= 561)) {
                                    if false {
                                        switch (fi67 * fi67) {
                                        case -2147483648:
                                            var v75 uint64 = 2
                                            _ = v75
                                        case -571:
                                            if (-690 > -2147483648) {
                                                v76 := GenS8{F0: (p0 || true), F1: 682}
                                                v77 := GenS9{F0: 2, F1: 2147483647, F2: 1}
                                                _ = v76
                                                _ = v77
                                            } else {
                                                p0 = p0
                                                v78 := GenS6{F0: 42, F1: 1, F2: v74}
                                                v79 := GenS3{F0: 127}
                                                if true {
                                                    var v80 int32 = 2147483647
                                                    _ = v80
                                                }
                                                if len(v7) >= 2 {
                                                    v7 = v7[1:len(v7)]
                                                }
                                                switch v74 {
                                                case "\n\t":
                                                    v74 = v74
                                                case "":
                                                    var v81 uint8 = 42
                                                    _ = v81
                                                case "a":
                                                    v7 = append(v7, (fi68 - fi67))
                                                default:
                                                    if (!p0) {
                                                        if p0 {
                                                            p0 = false
                                                            if len(v7) > 0 {
                                                                v7[0] = -211
                                                            }
                                                            v74 = v74
                                                            var v82 float64 = (float64(3.14159) * (0.1 - 0.1))
                                                            _ = v82
                                                        }
                                                        var v83 bool = (!(1 <= 540))
                                                        if p0 {
                                                            v84 := GenS4{F0: fi68, F1: 42, F2: 13}
                                                            v83 = true
                                                            switch (fi68 + 283) {
                                                            case -33:
                                                                v74 = ((v74 + v74) + v74)
                                                            case 61:
                                                                if (0.0 >= 0.0) {
                                                                    if len(v7) > 0 {
                                                                        v7[0] = (-(fi67))
                                                                    }
                                                                    _ = len(v7)
                                                                    if true {
                                                                        var v85 string = v74
                                                                        _ = v85
                                                                    } else {
                                                                        switch fi68 {
                                                                        case 32767:
                                                                            var v86 int16 = 0
                                                                            _ = v86
                                                                        case 1:
                                                                            if len(v7) >= 2 {
                                                                                v7 = v7[1:len(v7)]
                                                                            }
                                                                        case -9223372036854775808:
                                                                            var v87 float64 = 0.0
                                                                            _ = v87
                                                                        }
                                                                        v74 = v74
                                                                    }
                                                                } else {
                                                                    p0 = (186 <= 18446744073709551615)
                                                                }
                                                            }
                                                            _ = v84
                                                        } else {
                                                            v83 = (!(p0 || ((2147483647 != 940) && (-1 > 13))))
                                                            if len(v7) > 0 {
                                                                v7[0] = (fi67 | -1)
                                                            }
                                                        }
                                                        _ = v83
                                                    } else {
                                                        if (false && ((fi67 < (-(fi68))) || p0)) {
                                                            if true {
                                                                v88 := GenS10{F0: 4294967295, F1: 42}
                                                                v89 := GenS11{F0: -374}
                                                                var v90 int = (-(fi68))
                                                                _ = v88
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                            switch (fi67 >> ((235) & 63)) {
                                                            case 9223372036854775807:
                                                                p0 = false
                                                            }
                                                        } else {
                                                            var v91 uint16 = 0
                                                            var v92 uint32 = 4294967295
                                                            fmt.Println(v74, v91)
                                                            _ = v91
                                                            _ = v92
                                                        }
                                                    }
                                                }
                                                var v93 int16 = 32767
                                                var v94 int64 = 1
                                                _ = v78
                                                _ = v79
                                                _ = v93
                                                _ = v94
                                            }
                                        }
                                    } else {
                                        p0 = p0
                                    }
                                } else {
                                    v7 = append(v7, fi68)
                                    p0 = (255 < 0)
                                }
                                var v95 uint8 = 255
                                _ = v73
                                _ = v74
                                _ = v95
                            }
                            if len(v7) > 0 {
                                v7[0] = fi68
                            }
                            _ = fi68
                        }
                        p0 = (42 > -121)
                    }
                    v7 = append(v7, fi67)
                    _ = fi67
                }
            }
            for fi96 := 0; fi96 < 17; fi96++ {
                p0 = p0
                v7 = append(v7, 1)
                _ = fi96
            }
        }
        _ = v54
    }
    _ = p0
    _ = v7
    return 452
}

func main() {
    var v1 int8 = 42
    var v2 int16 = 1000
    var v3 int8 = v1
    v4 := []string{"a", ("hello" + (("foo bar" + ("" + (("hello" + "a") + (("\n\t" + ("\n\t" + "\n\t")) + "\n\t")))) + ("" + ("\n\t" + ""))))}
    var v5 int32 = -445
    var v6 bool = ((v1 % ((-128) | 1)) <= v1)
    for fi97 := 0; fi97 < 3; fi97++ {
        for fi98 := 0; fi98 < 2; fi98++ {
            v5 += 0
            _ = fi98
        }
        _ = fi97
    }
    if v6 {
        for fi99 := 0; fi99 < 20; fi99++ {
            fmt.Println(v3, v3, v1)
            var v100 string = ""
            _ = fi99
            _ = v100
        }
        if v6 {
            for fi101 := 0; fi101 < 1; fi101++ {
                var v102 float32 = (((-2.25 * 0.1) * 1.5) - 0.1)
                _ = fi101
                _ = v102
            }
            var v103 int = 42
            _ = v103
        } else {
            var v104 int16 = v2
            fmt.Println(v104, v3)
            v104 += 1
            _ = v104
        }
        if (-32768 < (v2 % (((v2 * (v2 + (v2 / (((v2 + (v2 + -357))) | 1))))) | 1))) {
            v105 := GenS12{F0: "a", F1: 0}
            if v6 {
                switch -19 {
                case -9223372036854775808:
                    if ((false || v6) && (!true)) {
                        var v106 int = 4294967295
                        _ = v106
                    } else {
                        v107 := GenS7{F0: 42, F1: (v2 - v2), F2: (v3 - v1)}
                        switch -625 {
                        case -376:
                            if ((v2 * (v2 ^ 42)) != (v2 | 0)) {
                                if (2 >= 0) {
                                    for fi108 := 0; fi108 < 17; fi108++ {
                                        if v6 {
                                            _ = len(v4)
                                            var v109 int8 = -128
                                            _ = v109
                                        }
                                        _ = fi108
                                    }
                                    v2 = v2
                                    var v110 int = 453
                                    _ = v110
                                } else {
                                    for fi111 := 0; fi111 < 2; fi111++ {
                                        fmt.Println(fi111)
                                        var v112 int32 = -2147483648
                                        _ = fi111
                                        _ = v112
                                    }
                                }
                                v113 := GenS6{F0: 1, F1: 4294967295, F2: "a"}
                                _ = v113
                            }
                        case -744:
                            var v114 bool = (42 == 42)
                            _ = v114
                        default:
                            var v115 uint = 1
                            _ = v115
                        }
                        v116 := GenS13{F0: 97, F1: (^(v5))}
                        _ = v107
                        _ = v116
                    }
                case 1:
                    switch -9223372036854775808 {
                    case -2147483648:
                        fmt.Println(v2, v6)
                    case 9223372036854775807:
                        for fi117 := 0; fi117 < 3; fi117++ {
                            var v118 uint16 = 0
                            _ = fi117
                            _ = v118
                        }
                    case -9223372036854775808:
                        switch 0 {
                        case -2147483648:
                            var v119 uint = 578
                            _ = v119
                        case 519:
                            var v120 uint8 = 242
                            _ = v120
                        default:
                            v121 := GenS11{F0: -9223372036854775808}
                            _ = v121
                        }
                    default:
                        fmt.Println(v1, v1, v6)
                    }
                }
            } else {
                if true {
                    v5 = (v5 - v5)
                }
                v3 -= 42
                for fi122 := 0; fi122 < 3; fi122++ {
                    if v6 {
                        var v123 int8 = (v3 << ((18446744073709551615) & 63))
                        for fi124 := 0; fi124 < 18; fi124++ {
                            if (!false) {
                                var v125 int64 = -726
                                var v126 uint = 255
                                _ = v125
                                _ = v126
                            } else {
                                if v6 {
                                    var v127 bool = true
                                    _ = v127
                                }
                            }
                            _ = fi124
                        }
                        v2 = v2
                        _ = v123
                    } else {
                        v128 := GenS14{F0: 203, F1: 9223372036854775807, F2: (fi122 - -1)}
                        _ = v128
                    }
                    switch fi122 {
                    case 127:
                        switch fi122 {
                        case 2147483647:
                            v4 = append(v4, "")
                        }
                    }
                    _ = fi122
                }
            }
            _ = v105
        } else {
            if len(v4) > 0 {
                v4[0] = (("hello" + ("foo bar" + "a")) + ("foo bar" + "a"))
            }
            var v129 int16 = 1
            var v130 uint = 398
            _ = v129
            _ = v130
        }
    } else {
        _ = len(v4)
        v3 = v3
    }
    fmt.Println(v3)
    if v6 {
        for fi131 := 0; fi131 < 8; fi131++ {
            v1 = (v1 << ((18446744073709551615) & 63))
            if (1 < 0) {
                fmt.Println(v6, v2)
            } else {
                fmt.Println(v2, fi131, v5)
                v2 = (v2 << ((4294967295) & 63))
            }
            _ = fi131
        }
        v1 -= (-(v1))
    }
    if ((v5 + 217) < (v5 ^ -123)) {
        v5 -= (v5 % (((v5 & v5)) | 1))
    }
    fmt.Println(v5, v2, v1)
    var v132 int8 = v3
    _ = len(v4)
    var v133 uint = 65535
    var v134 int = -431
    var v135 uint16 = 0
    var v136 int16 = int16(v134)
    var v137 uint32 = 764
    v2 -= (v136 & (v2 / ((int16(v135)) | 1)))
    for fi138 := 0; fi138 < 9; fi138++ {
        if ((int(v133) > fi138) || (false || (5 < 127))) {
            v133 -= v133
        } else {
            v3 = v1
            var v139 int16 = v2
            _ = v139
        }
        v140 := GenS9{F0: 34, F1: (v5 - (v5 << ((65535) & 63))), F2: (v134 % ((fi138) | 1))}
        v132 = (v132 >> (((v133 << ((18446744073709551615) & 63))) & 63))
        _ = fi138
        _ = v140
    }
    var v141 uint32 = (v137 ^ (v137 ^ (v137 >> ((uint(v135)) & 63))))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v132:", v132)
    fmt.Println("v133:", v133)
    fmt.Println("v134:", v134)
    fmt.Println("v135:", v135)
    fmt.Println("v136:", v136)
    fmt.Println("v137:", v137)
    fmt.Println("v141:", v141)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v132
    _ = v133
    _ = v134
    _ = v135
    _ = v136
    _ = v137
    _ = v141
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
