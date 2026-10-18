package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int64; F2 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 int16; F1 uint32; F2 int32 }
type GenS8 struct { F0 uint16; F1 bool }
type GenS9 struct { F0 uint8; F1 int16 }
type GenS10 struct { F0 bool; F1 uint; F2 bool }
type GenS11 struct { F0 uint16; F1 int16; F2 int64 }
type GenS12 struct { F0 int8; F1 uint8 }
type GenS13 struct { F0 uint64; F1 uint; F2 int32 }
type GenS14 struct { F0 uint; F1 int8; F2 int16 }
func genFn0(p0 int32, p1 uint, p2 int32) uint {
    var v4 float64 = 3.14159
    var v5 uint8 = 128
    switch "hello" {
    case "a":
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 float64 = v4
            _ = fi6
            _ = v7
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p1
}

func main() {
    var v1 int = -219
    var v2 int = v1
    var v3 bool = (false && ((true || false) || true))
    var v8 string = "\n\t"
    fmt.Println(v1, v2)
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := GenS1{F0: 127, F1: 511}
            v2 += (fi10 << ((4294967295) & 63))
            _ = fi10
            _ = v11
        }
        v12 := []int{fi9, (v2 & 42), 127}
        _ = fi9
        _ = v12
    }
    switch ((("hello" + ("\n\t" + (v8 + "hello"))) + v8) + v8) {
    case "a":
        for fi13 := 0; fi13 < 11; fi13++ {
            if v3 {
                switch (v2 ^ 159) {
                case -9223372036854775808:
                    fmt.Println(v1, v8)
                case 2147483647:
                    switch "" {
                    case "hello":
                        v2 = (fi13 % (((^(fi13))) | 1))
                    case "":
                        for fi14 := 0; fi14 < 1; fi14++ {
                            var v15 bool = v3
                            _ = fi14
                            _ = v15
                        }
                    }
                default:
                    v1 = (v1 % ((0) | 1))
                }
            } else {
                if (42 < 0) {
                    if v3 {
                        v2 = (fi13 + fi13)
                    }
                    for fi16 := 0; fi16 < 15; fi16++ {
                        switch ("hello" + (v8 + v8)) {
                        case "hello":
                            v3 = ((fi16 < v1) || (v3 && (int8(v1) >= -1)))
                        case "\n\t":
                            if true {
                                v17 := GenS2{F0: 244, F1: -9223372036854775808, F2: v8}
                                v18 := []int{42, (v2 / (((fi13 >> ((18446744073709551615) & 63))) | 1))}
                                _ = v17
                                _ = v18
                            }
                        case "":
                            var v19 int32 = -996
                            _ = v19
                        default:
                            v3 = false
                        }
                        _ = fi16
                    }
                } else {
                    fmt.Println(v3)
                    var v20 float32 = -2.25
                    v21 := []int{}
                    _ = v20
                    _ = v21
                }
                switch v8 {
                case "foo bar":
                    v22 := GenS1{F0: -128, F1: 105}
                    _ = v22
                case "hello":
                    var v23 int = v1
                    _ = v23
                }
            }
            v3 = (129 == 986)
            var v24 float32 = ((float32(float64(((1.5 + 1.5) * (float32(3.14159) / 1.5)))) - (1.5 * (1.5 + -2.25))) - float32((((0.0 + float64(1e10)) + -2.25) / 3.14159)))
            _ = fi13
            _ = v24
        }
    case "":
        for fi25 := 0; fi25 < 10; fi25++ {
            v1 -= 32768
            var v26 float32 = (1e10 * 3.14159)
            switch (fi25 / (((v2 >> ((4294967295) & 63))) | 1)) {
            case 128:
                switch v8 {
                case "a":
                    v27 := GenS3{F0: 0}
                    _ = v27
                case "hello":
                    v1 += (fi25 - (fi25 >> ((814) & 63)))
                default:
                    if v3 {
                        for fi28 := 0; fi28 < 17; fi28++ {
                            v29 := GenS1{F0: -24, F1: 42}
                            _ = fi28
                            _ = v29
                        }
                        v30 := []int{(v2 - fi25)}
                        _ = v30
                    } else {
                        var v31 uint8 = 42
                        if (((1e10 - (((0.1 * float64(((float32(-2.25) * (float32((1e10 / 1e10)) / v26)) * 1.5))) / ((0.1 / -2.25) * (0.1 - 1.5))) + 1e10)) * ((float64(((0.1 * v26) + v26)) + 0.0) / 1.0)) >= 0.0) {
                            v32 := GenS2{F0: 42, F1: -961, F2: v8}
                            _ = v32
                        } else {
                            if v3 {
                                for fi33 := 0; fi33 < 3; fi33++ {
                                    fmt.Println(v31, v31)
                                    switch (("hello" + v8) + "hello") {
                                    case "a":
                                        var v34 uint = 255
                                        _ = v34
                                    case "":
                                        var v35 int8 = -60
                                        _ = v35
                                    default:
                                        v36 := []int{128, 748, (v2 * (fi33 ^ 32768))}
                                        _ = v36
                                    }
                                    _ = fi33
                                }
                                var v37 int = 1
                                if true {
                                    if ((!(((((1.5 - v26) - (v26 / (-2.25 + (float32((((-2.25 / 0.1) * 1e10) - float64(v26))) - v26)))) + v26) - -2.25) >= 3.14159)) && (255 != 0)) {
                                        if v3 {
                                            fmt.Println(v8)
                                            v8 = v8
                                        } else {
                                            if (!v3) {
                                                var v38 float64 = 1.5
                                                _ = v38
                                            } else {
                                                fmt.Println(v8, fi25)
                                                v26 = 1.5
                                            }
                                            for fi39 := 0; fi39 < 5; fi39++ {
                                                fmt.Println(v8)
                                                var v40 int32 = 42
                                                if (1.5 != 0.1) {
                                                    if (v3 || true) {
                                                        var v41 int = fi39
                                                        var v42 int32 = v40
                                                        v3 = v3
                                                        v43 := GenS4{F0: v41}
                                                        var v44 uint = 18446744073709551615
                                                        v1 -= -9223372036854775808
                                                        var v45 float32 = float32(0.0)
                                                        _ = v41
                                                        _ = v42
                                                        _ = v43
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        var v46 uint = 18446744073709551615
                                                        if v3 {
                                                            v47 := GenS5{F0: v3, F1: (v31 + v31)}
                                                            var v48 uint32 = 2147483648
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                        var v49 float32 = v26
                                                        switch v2 {
                                                        case -1:
                                                            if v3 {
                                                                var v50 bool = v3
                                                                if ((1e10 - float32(0.1)) >= float32(float64(float32(3.14159)))) {
                                                                    v51 := GenS6{F0: -2147483648, F1: 50}
                                                                    _ = v51
                                                                }
                                                                switch (v8 + "a") {
                                                                case "":
                                                                    var v52 int64 = -1
                                                                    _ = v52
                                                                case "a":
                                                                    v53 := []int{}
                                                                    v54 := GenS6{F0: v40, F1: int8(v31)}
                                                                    _ = v53
                                                                    _ = v54
                                                                default:
                                                                    switch (^(v2)) {
                                                                    case 32767:
                                                                        v26 -= float32(1e10)
                                                                    default:
                                                                        if (!((!(650 != 42)) && true)) {
                                                                            v8 = ("hello" + v8)
                                                                            fmt.Println(v2, v3, v40)
                                                                        } else {
                                                                            v55 := []int{}
                                                                            var v56 float32 = 1e10
                                                                            v55 = append(v55, 32767)
                                                                            _ = v55
                                                                            _ = v56
                                                                        }
                                                                    }
                                                                }
                                                                _ = v50
                                                            }
                                                        case 1:
                                                            var v57 uint8 = (v31 - (v31 & 203))
                                                            _ = v57
                                                        case -2147483648:
                                                            fmt.Println(v37)
                                                        default:
                                                            if v3 {
                                                                v3 = (42 == 703)
                                                            }
                                                        }
                                                        fmt.Println(fi25)
                                                        _ = v46
                                                        _ = v49
                                                    }
                                                } else {
                                                    v31 += 128
                                                    v58 := GenS6{F0: (v40 >> ((182) & 63)), F1: 0}
                                                    switch -9223372036854775808 {
                                                    case 32767:
                                                        v59 := []int{}
                                                        _ = v59
                                                    case 42:
                                                        var v60 uint16 = 65535
                                                        _ = v60
                                                    }
                                                    _ = v58
                                                }
                                                _ = fi39
                                                _ = v40
                                            }
                                            for fi61 := 0; fi61 < 2; fi61++ {
                                                var v62 int32 = -264
                                                if (!true) {
                                                    var v63 string = ("a" + ("a" + "\n\t"))
                                                    var v64 uint = 2
                                                    var v65 string = ("foo bar" + (v8 + v8))
                                                    _ = v63
                                                    _ = v64
                                                    _ = v65
                                                } else {
                                                    switch (((v8 + v8) + v8) + "a") {
                                                    case "\n\t":
                                                        fmt.Println(v26)
                                                    case "":
                                                        switch v8 {
                                                        case "hello":
                                                            fmt.Println(v37, v62)
                                                        case "":
                                                            var v66 int8 = -56
                                                            _ = v66
                                                        case "foo bar":
                                                            if (2 != 1) {
                                                                v3 = v3
                                                            } else {
                                                                v1 = (v37 & 2147483647)
                                                            }
                                                        }
                                                    }
                                                }
                                                _ = fi61
                                                _ = v62
                                            }
                                        }
                                    }
                                    for fi67 := 0; fi67 < 15; fi67++ {
                                        var v68 int64 = 547
                                        v69 := GenS7{F0: 1, F1: 933, F2: 2147483647}
                                        if (0 <= -2147483648) {
                                            v70 := []int{}
                                            switch (("foo bar" + "hello") + v8) {
                                            case "\n\t":
                                                switch "hello" {
                                                case "hello":
                                                    v37 += fi67
                                                case "\n\t":
                                                    _ = len(v70)
                                                default:
                                                    v1 -= (^(fi67))
                                                }
                                                v26 += v26
                                                v70 = append(v70, (fi67 >> ((0) & 63)))
                                                fmt.Println(v37, v37)
                                            case "foo bar":
                                                var v71 uint16 = 837
                                                var v72 int64 = (^(v68))
                                                v1 = fi67
                                                v31 -= (v31 - v31)
                                                if len(v70) > 0 {
                                                    v70[0] = -1
                                                }
                                                _ = v71
                                                _ = v72
                                            }
                                            _ = v70
                                        }
                                        _ = fi67
                                        _ = v68
                                        _ = v69
                                    }
                                } else {
                                    var v73 string = v8
                                    _ = v73
                                }
                                _ = v37
                            } else {
                                v74 := []int{fi25}
                                var v75 float32 = 3.14159
                                v75 -= v26
                                _ = v74
                                _ = v75
                            }
                        }
                        _ = v31
                    }
                }
            case 9223372036854775807:
                if ("foo bar" > "foo bar") {
                    var v76 uint = 591
                    v77 := GenS8{F0: 441, F1: ((42 <= -403) || (0 <= 0))}
                    if (false && (int16(v1) < 579)) {
                        for fi78 := 0; fi78 < 10; fi78++ {
                            v79 := []int{}
                            if false {
                                v1 += v1
                                var v80 int16 = -1
                                var v81 uint32 = 222
                                _ = v80
                                _ = v81
                            }
                            if (-452 != -1) {
                                if len(v79) >= 2 {
                                    v79 = v79[1:len(v79)]
                                }
                                var v82 uint32 = 42
                                v79 = append(v79, (fi25 | int(v82)))
                                _ = v82
                            }
                            v3 = (!false)
                            if len(v79) > 0 {
                                v79[0] = v1
                            }
                            v1 += (-(v2))
                            _ = fi78
                            _ = v79
                        }
                        if v3 {
                            if v3 {
                                v83 := []int{}
                                for fi84 := 0; fi84 < 7; fi84++ {
                                    v2 += (-(fi25))
                                    _ = fi84
                                }
                                v83 = append(v83, 896)
                                _ = v83
                            } else {
                                if (1000 <= -1) {
                                    fmt.Println(v2)
                                    v85 := GenS9{F0: 1, F1: -798}
                                    v86 := []int{}
                                    _ = v85
                                    _ = v86
                                }
                                v26 += (1e10 + 1e10)
                                if (100 > 92) {
                                    v87 := GenS6{F0: -1, F1: 88}
                                    for fi88 := 0; fi88 < 10; fi88++ {
                                        if v3 {
                                            v89 := GenS10{F0: (255 <= 1), F1: 934, F2: false}
                                            switch "" {
                                            case "hello":
                                                v90 := []int{(fi88 + 805), 32767, v1}
                                                _ = v90
                                            case "a":
                                                fmt.Println(v1, v3, v8)
                                            }
                                            _ = v89
                                        }
                                        fmt.Println(fi88, v2)
                                        _ = fi88
                                    }
                                    fmt.Println(v1, fi25, v2)
                                    _ = v87
                                } else {
                                    v91 := []int{v1, (fi25 ^ (v2 - 0)), (^(v1))}
                                    for fi92 := 0; fi92 < 13; fi92++ {
                                        v3 = (!((v8 + v8) == ("foo bar" + v8)))
                                        _ = fi92
                                    }
                                    _ = v91
                                }
                            }
                        }
                        v93 := GenS11{F0: 1, F1: -27, F2: -1}
                        _ = v93
                    }
                    _ = v76
                    _ = v77
                }
            }
            _ = fi25
            _ = v26
        }
    }
    for fi94 := 0; fi94 < 2; fi94++ {
        v95 := GenS7{F0: 0, F1: 4294967295, F2: 42}
        v96 := []int{}
        v97 := GenS6{F0: 0, F1: 42}
        _ = fi94
        _ = v95
        _ = v96
        _ = v97
    }
    var v98 int32 = 42
    fmt.Println(v98)
    if v3 {
        switch "\n\t" {
        case "\n\t":
            v99 := GenS12{F0: 0, F1: 76}
            _ = v99
        }
        v1 -= v2
    } else {
        var v100 int64 = -1
        v98 = v98
        if false {
            v98 = (-(v98))
        }
        _ = v100
    }
    for fi101 := 0; fi101 < 10; fi101++ {
        v102 := GenS5{F0: (0.0 <= (float32((0.0 - 1.5)) + float32((float64(1e10) * 0.0)))), F1: 255}
        _ = fi101
        _ = v102
    }
    v103 := GenS1{F0: -118, F1: 1}
    v104 := []int{}
    for fi105 := 0; fi105 < 6; fi105++ {
        v2 = v2
        for fi106 := 0; fi106 < 3; fi106++ {
            v2 = 32768
            _ = fi106
        }
        _ = fi105
    }
    if false {
        if false {
            var v107 int64 = 0
            var v108 float32 = (0.0 / -2.25)
            var v109 string = "a"
            _ = v107
            _ = v108
            _ = v109
        }
        if len(v104) >= 2 {
            v104 = v104[1:len(v104)]
        }
        for fi110 := 0; fi110 < 18; fi110++ {
            v3 = (100 < 127)
            _ = fi110
        }
    } else {
        for fi111 := 0; fi111 < 1; fi111++ {
            for fi112 := 0; fi112 < 7; fi112++ {
                var v113 bool = true
                v114 := GenS2{F0: 0, F1: 910, F2: v8}
                v115 := GenS13{F0: 900, F1: 2, F2: v98}
                _ = fi112
                _ = v113
                _ = v114
                _ = v115
            }
            _ = fi111
        }
        var v116 int = (v1 >> ((1) & 63))
        for fi117 := 0; fi117 < 12; fi117++ {
            if (0 < 4294967295) {
                var v118 float32 = 0.0
                var v119 int64 = -387
                _ = v118
                _ = v119
            }
            _ = len(v104)
            _ = fi117
        }
        _ = v116
    }
    for fi120 := 0; fi120 < 11; fi120++ {
        var v121 float32 = -2.25
        switch v8 {
        case "foo bar":
            fmt.Println(v1, v98, v8)
        case "":
            var v122 uint16 = 4
            _ = v122
        case "a":
            fmt.Println(v121, fi120)
        }
        _ = fi120
        _ = v121
    }
    if v3 {
        v123 := GenS9{F0: 42, F1: -25}
        v124 := GenS1{F0: 100, F1: 42}
        if (127 == 42) {
            v125 := GenS14{F0: 2, F1: -100, F2: -982}
            _ = v125
        }
        _ = v123
        _ = v124
    }
    var v126 bool = false
    v1 += (-(v1))
    switch "hello" {
    case "foo bar":
        v98 -= v98
    case "\n\t":
        v104 = append(v104, (v1 - v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v104:", len(v104), v104)
    fmt.Println("v126:", v126)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v103
    _ = v104
    _ = v126
    _ = v2
    _ = v3
    _ = v8
    _ = v98
}
