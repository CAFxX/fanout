package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint32; F1 bool; F2 int }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint }
type GenS6 struct { F0 string; F1 uint8; F2 uint }
type GenS7 struct { F0 int32; F1 int32; F2 uint64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int32; F1 uint8 }
type GenS10 struct { F0 bool; F1 uint64 }
type GenS11 struct { F0 int64; F1 int32; F2 bool }
type GenS12 struct { F0 int16 }
func genFn0(p0 uint64, p1 bool) int8 {
    p0 += (p0 - (p0 * (p0 - 2)))
    p0 = (p0 - p0)
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 int64 = 42
    v2 := make(map[int]bool)
    v3 := []int{-1}
    var v4 int16 = 1
    var v5 int = 42
    var v6 uint64 = 430
    if true {
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 int8 = 42
            v9 := GenS1{F0: 2, F1: (v5 << ((187) & 63))}
            v4 += (v4 - v4)
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := GenS2{F0: "\n\t"}
        _ = v10
    } else {
        if false {
            var v11 int8 = -1
            _ = v11
        }
        if (!false) {
            fmt.Println(v1, v6, v6)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi12 := 0; fi12 < 17; fi12++ {
                fmt.Println(v4)
                v13 := GenS3{F0: 32767, F1: v4}
                v14 := GenS3{F0: 2, F1: 32767}
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        var v15 float32 = float32(1e10)
        _ = v15
    }
    switch v5 {
    case 9223372036854775807:
        v16 := GenS4{F0: 2, F1: ((1.5 + 0.1) <= 1e10), F2: (v5 ^ 269)}
        _ = v16
    }
    var v17 float32 = 0.0
    if (false && true) {
        var v18 uint64 = v6
        v19 := GenS1{F0: 1, F1: 1}
        v20 := GenS3{F0: 1, F1: (v4 ^ int16(v6))}
        _ = v18
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 4; fi21++ {
            for fi22 := 0; fi22 < 19; fi22++ {
                fmt.Println(v4, v4, fi21)
                var v23 int = -9223372036854775808
                _ = fi22
                _ = v23
            }
            v17 = v17
            _ = fi21
        }
        for fi24 := 0; fi24 < 11; fi24++ {
            fmt.Println(v6)
            v5 = v5
            v25 := GenS5{F0: 514, F1: 1, F2: 42}
            _ = fi24
            _ = v25
        }
        fmt.Println(v17)
    }
    v2[v5] = (-1 > (^(v1)))
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 int8 = 127
        _ = fi26
        _ = v27
    }
    v5 += v5
    if false {
        v5 = 4294967295
        for fi28 := 0; fi28 < 3; fi28++ {
            var v29 int64 = (v1 - v1)
            for fi30 := 0; fi30 < 12; fi30++ {
                var v31 int16 = -32768
                if false {
                    v29 = (v29 - v29)
                }
                switch v5 {
                case 9223372036854775807:
                    var v32 float64 = 3.14159
                    _ = v32
                case -9223372036854775808:
                    fmt.Println(v17)
                default:
                    if ((1 >= 255) && true) {
                        v33 := GenS6{F0: "\n\t", F1: 3, F2: 488}
                        v34 := v2[fi30]
                        if (-2147483648 >= -2147483648) {
                            _ = len(v3)
                        }
                        v5 = v5
                        var v35 float64 = (float64(v17) + (-2.25 - (1e10 / 1e10)))
                        var v36 string = "hello"
                        if (true || (v34 && v34)) {
                            _ = len(v2)
                        } else {
                            var v37 int16 = (v4 ^ v4)
                            v38 := GenS5{F0: 908, F1: 781, F2: 289}
                            _ = v37
                            _ = v38
                        }
                        _ = v33
                        _ = v34
                        _ = v35
                        _ = v36
                    } else {
                        switch (("\n\t" + ("" + ("a" + "a"))) + "foo bar") {
                        case "foo bar":
                            var v39 uint8 = 128
                            _ = v39
                        }
                        var v40 int8 = -1
                        _ = v40
                    }
                }
                _ = fi30
                _ = v31
            }
            switch "foo bar" {
            case "":
                if false {
                    for fi41 := 0; fi41 < 11; fi41++ {
                        v42 := GenS1{F0: 2, F1: fi28}
                        v5 = (-(fi28))
                        _ = fi41
                        _ = v42
                    }
                    if (!false) {
                        var v43 int16 = (v4 & v4)
                        v1 = (v29 + (v1 & (v29 | (v29 ^ 9223372036854775807))))
                        switch (fi28 - (v5 - (fi28 | (fi28 ^ -2147483648)))) {
                        case -58:
                            for fi44 := 0; fi44 < 19; fi44++ {
                                var v45 int8 = 100
                                if (v4 > (v4 + v43)) {
                                    var v46 int32 = 2147483647
                                    _ = v46
                                }
                                var v47 int32 = -988
                                _ = fi44
                                _ = v45
                                _ = v47
                            }
                        case 9223372036854775807:
                            if (-122 == -98) {
                                if false {
                                    v48 := GenS3{F0: 42, F1: v43}
                                    var v49 uint8 = 127
                                    _ = v48
                                    _ = v49
                                }
                                for fi50 := 0; fi50 < 20; fi50++ {
                                    var v51 int64 = (v1 - v1)
                                    var v52 uint8 = 23
                                    if len(v3) > 0 {
                                        v3[0] = fi28
                                    }
                                    _ = fi50
                                    _ = v51
                                    _ = v52
                                }
                            } else {
                                v1 = (^(v29))
                                if len(v3) > 0 {
                                    v3[0] = v5
                                }
                                for fi53 := 0; fi53 < 19; fi53++ {
                                    if ((!false) || ((0.1 != 1e10) || false)) {
                                        if false {
                                            v54 := GenS4{F0: 934, F1: true, F2: 2147483647}
                                            _ = v54
                                        } else {
                                            if true {
                                                switch "hello" {
                                                case "":
                                                    if (807 == 4294967295) {
                                                        if false {
                                                            v4 -= v43
                                                            if (143 > 0) {
                                                                var v55 uint64 = v6
                                                                _ = v55
                                                            }
                                                            var v56 int16 = v4
                                                            _ = v56
                                                        }
                                                        if (true || ((v17 + 0.0) > v17)) {
                                                            v6 = v6
                                                        }
                                                        switch (fi28 / ((-307) | 1)) {
                                                        case -9223372036854775808:
                                                            if len(v3) > 0 {
                                                                v3[0] = 32767
                                                            }
                                                        }
                                                        var v57 float64 = 1.5
                                                        v6 = 42
                                                        if false {
                                                            switch "foo bar" {
                                                            case "":
                                                                var v58 uint = 18446744073709551615
                                                                _ = v58
                                                            }
                                                            v2[-9223372036854775808] = false
                                                            var v59 uint32 = 42
                                                            _ = v59
                                                        }
                                                        fmt.Println(v5, fi28, v17)
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        _ = v57
                                                    }
                                                default:
                                                    switch (-(fi53)) {
                                                    case 568:
                                                        var v60 uint = 18446744073709551615
                                                        _ = v60
                                                    case -9223372036854775808:
                                                        var v61 int = fi28
                                                        _ = v61
                                                    default:
                                                        v62 := GenS7{F0: -977, F1: 42, F2: (v6 * (^(v6)))}
                                                        _ = v62
                                                    }
                                                }
                                                var v63 bool = true
                                                var v64 uint64 = (^(v6))
                                                _ = v63
                                                _ = v64
                                            } else {
                                                v17 = (v17 / 3.14159)
                                                _ = len(v3)
                                                var v65 float64 = -2.25
                                                _ = v65
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = fi53
                                            }
                                            var v66 uint8 = 232
                                            _ = v66
                                        }
                                        var v67 uint16 = 0
                                        _ = v67
                                    } else {
                                        v68 := GenS2{F0: ("\n\t" + ((("" + "") + "a") + "\n\t"))}
                                        _ = v68
                                    }
                                    var v69 int32 = 1
                                    if true {
                                        v70 := GenS1{F0: 772, F1: fi53}
                                        v6 += 129
                                        v43 = int16(fi53)
                                        _ = v70
                                    }
                                    _ = fi53
                                    _ = v69
                                }
                            }
                        default:
                            if true {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                for fi71 := 0; fi71 < 9; fi71++ {
                                    var v72 uint8 = 128
                                    v73 := GenS1{F0: 32767, F1: (v5 - v5)}
                                    var v74 uint8 = uint8(v6)
                                    _ = fi71
                                    _ = v72
                                    _ = v73
                                    _ = v74
                                }
                            }
                        }
                        _ = v43
                    } else {
                        var v75 float32 = v17
                        delete(v2, (^(v5)))
                        v75 = 0.1
                        _ = v75
                    }
                    if false {
                        var v76 int64 = (v1 * v1)
                        v6 -= 42
                        _ = v76
                    } else {
                        if false {
                            v1 = v29
                            var v77 uint16 = 2
                            v6 -= v6
                            _ = v77
                        } else {
                            for fi78 := 0; fi78 < 16; fi78++ {
                                var v79 int = (v5 << ((18446744073709551615) & 63))
                                v80 := GenS1{F0: uint16(v6), F1: (v5 % (((fi28 * 0)) | 1))}
                                _ = fi78
                                _ = v79
                                _ = v80
                            }
                        }
                    }
                } else {
                    fmt.Println(v6, v4)
                }
            case "foo bar":
                fmt.Println(v4, v6)
            }
            _ = fi28
            _ = v29
        }
        switch v5 {
        case 9223372036854775807:
            var v81 uint64 = v6
            _ = v81
        case 128:
            v2[(v5 + -695)] = ((!false) || (711 >= 1))
        }
    } else {
        v5 -= (v5 & (v5 / (((^(v5))) | 1)))
        v2[(v5 >> ((399) & 63))] = (!(("\n\t" + (("foo bar" + ("a" + ("a" + ("hello" + ((("foo bar" + "") + (("" + "") + "\n\t")) + "a"))))) + "hello")) != "\n\t"))
    }
    v5 += v5
    if (1 >= 0) {
        if len(v3) > 0 {
            v3[0] = 128
        }
        switch "a" {
        case "hello":
            var v82 uint32 = 2147483648
            _ = v82
        default:
            if true {
                fmt.Println(v6, v5, v4)
            } else {
                for fi83 := 0; fi83 < 13; fi83++ {
                    v4 += int16(v1)
                    var v84 int64 = (v1 ^ (^(v1)))
                    _ = fi83
                    _ = v84
                }
                fmt.Println(v4, v4, v4)
                if len(v3) > 0 {
                    v3[0] = v5
                }
            }
        }
    } else {
        if true {
            for fi85 := 0; fi85 < 9; fi85++ {
                var v86 uint8 = 255
                var v87 string = "foo bar"
                v87 = "foo bar"
                _ = fi85
                _ = v86
                _ = v87
            }
        }
    }
    switch (v5 | (v5 >> ((688) & 63))) {
    case -167:
        v4 = v4
    case -9223372036854775808:
        for fi88 := 0; fi88 < 7; fi88++ {
            var v89 float64 = float64(((v17 / -2.25) * 3.14159))
            switch "" {
            case "\n\t":
                fmt.Println(v5, v6)
            default:
                for fi90 := 0; fi90 < 7; fi90++ {
                    v91 := GenS5{F0: 2147483648, F1: 106, F2: 42}
                    _ = fi90
                    _ = v91
                }
            }
            switch "a" {
            case "foo bar":
                for fi92 := 0; fi92 < 20; fi92++ {
                    switch ((("hello" + "foo bar") + "a") + "hello") {
                    case "a":
                        v4 = v4
                    }
                    v6 += (v6 ^ v6)
                    if (v89 >= (3.14159 - v89)) {
                        switch "hello" {
                        case "":
                            v5 = fi92
                        case "hello":
                            delete(v2, fi88)
                        case "\n\t":
                            var v93 float64 = v89
                            v89 += 1e10
                            _ = v93
                        default:
                            if len(v3) > 0 {
                                v3[0] = (fi92 << ((uint(v6)) & 63))
                            }
                        }
                        v1 = -1
                    }
                    _ = fi92
                }
            case "a":
                var v94 int16 = (-(v4))
                _ = v94
            case "":
                for fi95 := 0; fi95 < 14; fi95++ {
                    if ((2147483647 != 4294967295) && true) {
                        if (0 != genFn0(v6, true)) {
                            if true {
                                v96 := GenS8{F0: -1}
                                _ = v96
                            }
                            var v97 int16 = (v4 - v4)
                            _ = v97
                        }
                    } else {
                        if len(v3) > 0 {
                            v3[0] = (v5 ^ v5)
                        }
                        fmt.Println(fi88, v6, fi95)
                    }
                    if false {
                        v1 -= v1
                    } else {
                        var v98 float64 = v89
                        var v99 int64 = 9223372036854775807
                        if (0.1 > float32((v89 * v89))) {
                            var v100 int8 = -5
                            if ("\n\t" >= "a") {
                                _ = len(v2)
                            } else {
                                v100 = v100
                            }
                            _ = v100
                        }
                        _ = v98
                        _ = v99
                    }
                    switch ("a" + "") {
                    case "":
                        v101 := v2[fi88]
                        _ = v101
                    case "foo bar":
                        v2[464] = (false && (!true))
                    case "\n\t":
                        if (32767 != 1) {
                            v102 := GenS7{F0: 42, F1: 2147483647, F2: (^(v6))}
                            v4 = (v4 - 1)
                            if true {
                                v103 := GenS9{F0: 42, F1: 0}
                                v104 := GenS10{F0: (fi88 >= 176), F1: (v6 >> ((0) & 63))}
                                fmt.Println(v17, fi88)
                                _ = v103
                                _ = v104
                            }
                            _ = v102
                        }
                    default:
                        var v105 float64 = 0.0
                        _ = v105
                    }
                    v106 := GenS7{F0: 87, F1: -903, F2: v6}
                    _ = fi95
                    _ = v106
                }
            default:
                if (v17 <= v17) {
                    v89 = v89
                } else {
                    if (9223372036854775807 != (v1 - 0)) {
                        if true {
                            for fi107 := 0; fi107 < 8; fi107++ {
                                var v108 int16 = v4
                                _ = fi107
                                _ = v108
                            }
                        } else {
                            v4 += 1000
                        }
                        var v109 int = (fi88 % (((fi88 * v5)) | 1))
                        var v110 int64 = (v1 & (v1 | v1))
                        _ = v109
                        _ = v110
                    } else {
                        v111 := GenS11{F0: v1, F1: -1, F2: ((true || true) && true)}
                        _ = v111
                    }
                }
            }
            _ = fi88
            _ = v89
        }
    case 0:
        v1 = (v1 + (v1 << ((459) & 63)))
    default:
        for fi112 := 0; fi112 < 18; fi112++ {
            if ((v6 >> ((2) & 63)) > (^(v6))) {
                fmt.Println(v1, v4)
                if true {
                    var v113 float64 = float64(v17)
                    for fi114 := 0; fi114 < 12; fi114++ {
                        if (0 == 43) {
                            fmt.Println(v1, fi114, v113)
                            var v115 uint8 = 1
                            _ = v115
                        }
                        fmt.Println(v5, v113, v1)
                        _ = fi114
                    }
                    if true {
                        v5 -= 2147483647
                    } else {
                        v113 = (v113 * (v113 + v113))
                        v5 += (^(v5))
                    }
                    _ = v113
                }
                var v116 uint32 = 0
                _ = v116
            } else {
                for fi117 := 0; fi117 < 13; fi117++ {
                    v118 := GenS12{F0: 0}
                    _ = fi117
                    _ = v118
                }
                v5 = v5
            }
            _ = fi112
        }
    }
    var v119 float64 = 0.1
    fmt.Println(v119, v1, v119)
    v5 -= (v5 ^ (v5 >> ((798) & 63)))
    var v120 uint16 = 2
    switch "a" {
    case "foo bar":
        for fi121 := 0; fi121 < 12; fi121++ {
            v120 = 42
            _ = fi121
        }
    case "hello":
        v4 = (v4 / (((v4 & (v4 | -1))) | 1))
    }
    v122 := GenS11{F0: (v1 / ((v1) | 1)), F1: 329, F2: (!(-118 > -6))}
    fmt.Println("v1:", v1)
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v119
    _ = v120
    _ = v122
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
