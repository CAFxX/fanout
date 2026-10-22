package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 uint; F2 int }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint8; F1 int64; F2 bool }
type GenS7 struct { F0 uint64; F1 uint; F2 int8 }
type GenS8 struct { F0 bool; F1 int16; F2 uint }
type GenS9 struct { F0 uint16; F1 uint8 }
type GenS10 struct { F0 uint8; F1 int }
type GenS11 struct { F0 int32; F1 int64; F2 string }
type GenS12 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS13 struct { F0 uint; F1 int8; F2 bool }
type GenS14 struct { F0 uint; F1 int; F2 int8 }
type GenS15 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint64) uint64 {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            p0 += p0
            _ = fi7
        }
        _ = fi6
    }
    switch "\n\t" {
    case "foo bar":
        if true {
            fmt.Println(p0)
            var v8 uint64 = (p0 & p0)
            v9 := []int{}
            _ = v8
            _ = v9
        } else {
            p0 = 4294967295
            var v10 int = 4294967295
            for fi11 := 0; fi11 < 15; fi11++ {
                for fi12 := 0; fi12 < 15; fi12++ {
                    p0 -= 0
                    switch fi11 {
                    case 127:
                        var v13 uint64 = (p0 >> ((1) & 63))
                        _ = v13
                    case 32768:
                        fmt.Println(fi11)
                    case -154:
                        var v14 uint16 = 1
                        switch -1 {
                        case 128:
                            v15 := GenS1{F0: true}
                            _ = v15
                        }
                        v16 := []int{v10, (fi11 << ((0) & 63)), fi12}
                        _ = v14
                        _ = v16
                    }
                    var v17 uint = 0
                    _ = fi12
                    _ = v17
                }
                fmt.Println(v10)
                _ = fi11
            }
            _ = v10
        }
    default:
        if (255 < 0) {
            p0 -= p0
        } else {
            for fi18 := 0; fi18 < 13; fi18++ {
                if (true && true) {
                    var v19 uint8 = 127
                    _ = v19
                }
                for fi20 := 0; fi20 < 16; fi20++ {
                    switch fi18 {
                    case -2147483648:
                        if (!(0.0 != 3.14159)) {
                            if false {
                                v21 := []int{fi20, -74, (^(fi18)), (fi20 >> ((1) & 63))}
                                if true {
                                    switch (fi20 * (fi18 % ((fi18) | 1))) {
                                    case -698:
                                        var v22 int16 = 32767
                                        v22 += (v22 - (v22 >> ((733) & 63)))
                                        var v23 int = fi18
                                        _ = v22
                                        _ = v23
                                    default:
                                        fmt.Println(p0, fi18)
                                    }
                                    v24 := GenS2{F0: p0, F1: 2, F2: 65535}
                                    p0 = p0
                                    _ = v24
                                } else {
                                    p0 = p0
                                }
                                _ = v21
                            }
                            switch "a" {
                            case "hello":
                                p0 = (p0 << ((18446744073709551615) & 63))
                            case "foo bar":
                                p0 = 138
                            }
                            var v25 string = ""
                            v26 := []int{fi18}
                            _ = v25
                            _ = v26
                        } else {
                            if false {
                                p0 = (^(p0))
                            }
                            v27 := GenS2{F0: 454, F1: 0, F2: 430}
                            switch fi20 {
                            case -1:
                                v28 := GenS1{F0: false}
                                _ = v28
                            }
                            _ = v27
                        }
                    case 1:
                        v29 := GenS2{F0: p0, F1: 2147483648, F2: 2}
                        _ = v29
                    case -9223372036854775808:
                        v30 := []int{(fi18 * 2147483647), (fi20 << ((972) & 63))}
                        _ = v30
                    }
                    p0 -= p0
                    if (true || (false || false)) {
                        var v31 uint8 = 184
                        var v32 bool = true
                        fmt.Println(v32, v31)
                        _ = v31
                        _ = v32
                    } else {
                        fmt.Println(p0)
                        fmt.Println(fi18)
                        p0 -= (^(p0))
                    }
                    _ = fi20
                }
                switch "" {
                case "foo bar":
                    p0 += (p0 >> ((911) & 63))
                case "a":
                    for fi33 := 0; fi33 < 2; fi33++ {
                        v34 := GenS3{F0: 66}
                        var v35 int8 = 0
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                case "\n\t":
                    v36 := GenS2{F0: p0, F1: 875, F2: 65535}
                    _ = v36
                default:
                    v37 := GenS3{F0: 77}
                    _ = v37
                }
                _ = fi18
            }
            fmt.Println(p0)
        }
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        v39 := GenS1{F0: false}
        var v40 bool = (((!true) || true) && ((2 <= 4294967295) && (!false)))
        v41 := GenS2{F0: (p0 / ((0) | 1)), F1: 2, F2: 4294967295}
        _ = fi38
        _ = v39
        _ = v40
        _ = v41
    }
    _ = p0
    return (p0 & 851)
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 348
    var v3 uint64 = 21
    var v4 int32 = 0
    var v5 string = ("hello" + "a")
    var v42 float32 = 0.0
    v43 := []int{-176, -9223372036854775808, -1}
    if false {
        v5 = v5
    }
    v4 += v4
    v3 += v3
    fmt.Println(v3, v2)
    if true {
        v3 = genFn0((v3 * uint64(v2)))
    } else {
        v44 := GenS3{F0: 1}
        if ("\n\t" <= ((v5 + ("hello" + v5)) + ("" + (v5 + v5)))) {
            var v45 uint64 = 311
            v45 -= (v3 % (((v3 << (((v2 << ((2) & 63))) & 63))) | 1))
            _ = v45
        } else {
            v46 := GenS3{F0: -90}
            var v47 uint8 = 128
            _ = v46
            _ = v47
        }
        _ = v44
    }
    for fi48 := 0; fi48 < 14; fi48++ {
        if true {
            v4 = 0
            for fi49 := 0; fi49 < 15; fi49++ {
                switch "a" {
                case "hello":
                    switch fi49 {
                    case 127:
                        v50 := GenS2{F0: 969, F1: 2147483648, F2: uint(v3)}
                        _ = v50
                    case 1:
                        switch fi49 {
                        case 4294967295:
                            fmt.Println(fi49, v5)
                        default:
                            v51 := GenS1{F0: false}
                            if false {
                                var v52 uint64 = (^(v3))
                                _ = v52
                            } else {
                                switch fi48 {
                                case 604:
                                    var v53 bool = true
                                    _ = v53
                                default:
                                    if false {
                                        _ = len(v43)
                                        v5 = (v5 + "")
                                    } else {
                                        v54 := GenS2{F0: v3, F1: 2147483648, F2: v2}
                                        v2 = 255
                                        _ = v54
                                    }
                                }
                                if ((((1.5 / 3.14159) / 1e10) - (-2.25 + 0.1)) <= (3.14159 + (1e10 - float64(float32((3.14159 - 3.14159)))))) {
                                    fmt.Println(v5)
                                } else {
                                    var v55 int16 = -481
                                    v56 := v1[v5]
                                    var v57 uint64 = v3
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                }
                            }
                            v58 := GenS4{F0: 98, F1: v2, F2: fi49}
                            _ = v51
                            _ = v58
                        }
                    case 9223372036854775807:
                        var v59 bool = true
                        _ = v59
                    }
                case "\n\t":
                    switch (((((v5 + "") + (("a" + v5) + v5)) + v5) + ((v5 + (("foo bar" + "") + (v5 + v5))) + (v5 + v5))) + "a") {
                    case "a":
                        v60 := GenS5{F0: (v4 | (v4 & -2147483648)), F1: (-399 > -65)}
                        _ = v60
                    case "\n\t":
                        var v61 uint64 = v3
                        _ = v61
                    }
                case "a":
                    if false {
                        if true {
                            var v62 int = (fi49 + (-(fi49)))
                            if false {
                                if (100 != int8(v4)) {
                                    fmt.Println(v62)
                                } else {
                                    fmt.Println(v3)
                                    v3 += (v3 | (v3 / ((genFn0(genFn0(v3))) | 1)))
                                    var v63 int8 = 0
                                    switch (fi48 * v62) {
                                    case 127:
                                        v64 := GenS5{F0: (^(v4)), F1: ("foo bar" < "foo bar")}
                                        _ = v64
                                    case -9223372036854775808:
                                        fmt.Println(fi48, v42)
                                    }
                                    if true {
                                        fmt.Println(v62, v3)
                                        fmt.Println(v62, v2)
                                    }
                                    _ = v63
                                }
                                var v65 uint64 = (v3 / ((v3) | 1))
                                if ((fi49 - fi48) != fi49) {
                                    if false {
                                        v62 -= (fi49 | 127)
                                    } else {
                                        v66 := GenS5{F0: (v4 % ((-20) | 1)), F1: (!false)}
                                        _ = v66
                                    }
                                    v67 := GenS4{F0: (v65 / (((v65 / ((v65) | 1))) | 1)), F1: (^(v2)), F2: 32768}
                                    var v68 int = fi48
                                    _ = v67
                                    _ = v68
                                } else {
                                    fmt.Println(v5)
                                    var v69 int64 = -1
                                    var v70 uint32 = 2
                                    _ = v69
                                    _ = v70
                                }
                                _ = v65
                            }
                            if (0.1 <= -2.25) {
                                var v71 int64 = -1
                                switch fi48 {
                                case 1:
                                    var v72 string = "foo bar"
                                    _ = v72
                                case -582:
                                    v73 := GenS6{F0: 38, F1: (^(v71)), F2: (uint16(v3) < 255)}
                                    _ = v73
                                case 406:
                                    var v74 int64 = v71
                                    _ = v74
                                default:
                                    switch "\n\t" {
                                    case "foo bar":
                                        v71 -= v71
                                    case "\n\t":
                                        if (0 != genFn0((v3 - (v3 / (((v3 | v3)) | 1))))) {
                                            switch -2147483648 {
                                            case 32768:
                                                v42 += v42
                                            case -9223372036854775808:
                                                _ = len(v1)
                                            default:
                                                if false {
                                                    var v75 uint32 = 1
                                                    fmt.Println(v75, v4, v62)
                                                    v76 := GenS7{F0: (v3 / (((^(v3))) | 1)), F1: (v2 + (v2 / (((v2 ^ (v2 - (v2 ^ v2)))) | 1))), F2: 125}
                                                    _ = v75
                                                    _ = v76
                                                }
                                            }
                                            _ = len(v1)
                                        } else {
                                            _ = len(v43)
                                            fmt.Println(v4, v71)
                                            switch (v5 + v5) {
                                            case "a":
                                                v77 := GenS7{F0: genFn0((v3 | 202)), F1: v2, F2: 100}
                                                _ = v77
                                            case "hello":
                                                var v78 int64 = (-(v71))
                                                _ = v78
                                            default:
                                                v3 = 1
                                            }
                                        }
                                    case "a":
                                        v5 = v5
                                        if (0.0 > 3.14159) {
                                            switch "\n\t" {
                                            case "hello":
                                                v4 = v4
                                            case "foo bar":
                                                if len(v43) >= 2 {
                                                    v43 = v43[1:len(v43)]
                                                }
                                            case "\n\t":
                                                var v79 float32 = 0.1
                                                _ = v79
                                            }
                                            v1["hello"] = (-(fi48))
                                        } else {
                                            v2 = 0
                                            v62 = fi48
                                            v80 := GenS4{F0: v3, F1: v2, F2: (fi49 / ((fi49) | 1))}
                                            _ = v80
                                        }
                                        var v81 uint = v2
                                        _ = v81
                                    default:
                                        var v82 int32 = v4
                                        _ = v82
                                    }
                                }
                                var v83 uint16 = 283
                                _ = v71
                                _ = v83
                            }
                            _ = v62
                        } else {
                            switch "\n\t" {
                            case "":
                                var v84 float32 = (0.1 / -2.25)
                                _ = v84
                            }
                            if ((fi49 | fi48) < fi49) {
                                if (!(2 == (v2 ^ (^(v2))))) {
                                    if (("foo bar" + ((v5 + v5) + v5)) != v5) {
                                        var v85 uint32 = 1
                                        fmt.Println(v5)
                                        fmt.Println(fi49, v3)
                                        fmt.Println(v4, v4)
                                        v1["\n\t"] = (^(fi48))
                                        _ = v85
                                    }
                                    if (487 > 42) {
                                        v86 := v1["foo bar"]
                                        v87 := GenS8{F0: false, F1: -32768, F2: (v2 * (v2 + v2))}
                                        _ = v86
                                        _ = v87
                                    } else {
                                        var v88 uint8 = 3
                                        _ = v88
                                    }
                                } else {
                                    if false {
                                        if (true || (9223372036854775807 != 845)) {
                                            if (255 >= 127) {
                                                if true {
                                                    if false {
                                                        var v89 uint16 = 1
                                                        if (0.0 < v42) {
                                                            v4 = (-(v4))
                                                        } else {
                                                            fmt.Println(v2, v3)
                                                        }
                                                        if false {
                                                            var v90 string = ("a" + (("\n\t" + "foo bar") + v5))
                                                            v89 = (v89 >> ((v2) & 63))
                                                            _ = v90
                                                        } else {
                                                            v89 += 65535
                                                            v5 = v5
                                                        }
                                                        _ = v89
                                                    }
                                                    if false {
                                                        v91 := GenS9{F0: 234, F1: 122}
                                                        v92 := GenS8{F0: (0 > 511), F1: 1, F2: 4294967295}
                                                        switch "" {
                                                        case "foo bar":
                                                            var v93 int32 = v4
                                                            var v94 uint8 = 203
                                                            var v95 uint8 = (v94 - v94)
                                                            _ = v93
                                                            _ = v94
                                                            _ = v95
                                                        }
                                                        _ = v91
                                                        _ = v92
                                                    } else {
                                                        v3 = (v3 - v3)
                                                        if false {
                                                            if (((-114 != 127) || (v3 < (v3 % (((^(v3))) | 1)))) || (genFn0(uint64(v2)) >= genFn0(uint64(v2)))) {
                                                                v4 += (v4 * v4)
                                                                v96 := GenS10{F0: 255, F1: 4294967295}
                                                                if (127 >= 127) {
                                                                    switch fi49 {
                                                                    case 128:
                                                                        v2 += v2
                                                                    case 0:
                                                                        if false {
                                                                            var v97 uint16 = 255
                                                                            _ = v97
                                                                        }
                                                                        if true {
                                                                            var v98 uint8 = 128
                                                                            _ = v98
                                                                        }
                                                                    case 42:
                                                                        fmt.Println(fi48, v4)
                                                                    }
                                                                    var v99 bool = false
                                                                    v4 -= v4
                                                                    _ = v99
                                                                }
                                                                var v100 uint = v2
                                                                _ = v100
                                                                _ = v96
                                                            }
                                                            var v101 bool = (240 > 42)
                                                            _ = v101
                                                        }
                                                        v42 -= ((float32((3.14159 - -2.25)) / v42) + (v42 - 0.1))
                                                        v102 := GenS7{F0: genFn0((v3 | 4294967295)), F1: (v2 >> ((v2) & 63)), F2: 0}
                                                        v5 = ""
                                                        if len(v43) >= 2 {
                                                            v43 = v43[1:len(v43)]
                                                        }
                                                        v103 := GenS9{F0: 796, F1: 121}
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                }
                                            } else {
                                                var v104 float64 = 0.0
                                                fmt.Println(v4)
                                                if false {
                                                    var v105 int16 = -824
                                                    _ = v105
                                                } else {
                                                    if false {
                                                        v106 := GenS11{F0: -2147483648, F1: -9223372036854775808, F2: ("a" + (v5 + ((v5 + (v5 + "a")) + "hello")))}
                                                        _ = v106
                                                    }
                                                    v2 -= 254
                                                }
                                                _ = v104
                                            }
                                            v4 -= v4
                                        }
                                        var v107 int = fi49
                                        var v108 int = 128
                                        _ = v107
                                        _ = v108
                                    } else {
                                        var v109 uint64 = v3
                                        fmt.Println(fi49, fi48)
                                        _ = v109
                                    }
                                    switch fi49 {
                                    case -2147483648:
                                        switch v5 {
                                        case "hello":
                                            _ = len(v43)
                                        case "a":
                                            v4 = -2147483648
                                        }
                                        if false {
                                            v110 := GenS12{F0: 145, F1: 65535, F2: 65535}
                                            _ = v110
                                        }
                                        var v111 int16 = -376
                                        _ = v111
                                    case -9223372036854775808:
                                        fmt.Println(fi48)
                                    }
                                    if len(v43) > 0 {
                                        v43[0] = 127
                                    }
                                }
                                fmt.Println(fi49, v5, v4)
                                v5 = v5
                            } else {
                                if true {
                                    switch "hello" {
                                    case "\n\t":
                                        v43 = append(v43, fi48)
                                    case "":
                                        v112 := GenS13{F0: (v2 << (((v2 % (((v2 << ((v2) & 63))) | 1))) & 63)), F1: int8(fi48), F2: (false && (254 == 2))}
                                        _ = v112
                                    }
                                } else {
                                    if (42 != 23) {
                                        if true {
                                            v2 = 0
                                            v113 := v1[("hello" + (v5 + v5))]
                                            _ = v113
                                        } else {
                                            var v114 int32 = (-(v4))
                                            fmt.Println(fi48, v2)
                                            _ = v114
                                        }
                                        var v115 int64 = -1
                                        _ = v115
                                    } else {
                                        var v116 int32 = (-(v4))
                                        v4 += v116
                                        var v117 uint = v2
                                        _ = v116
                                        _ = v117
                                    }
                                    switch ((v5 + "") + (v5 + v5)) {
                                    case "foo bar":
                                        switch (fi48 >> ((v2) & 63)) {
                                        case -2147483648:
                                            if len(v43) > 0 {
                                                v43[0] = fi49
                                            }
                                        }
                                    case "\n\t":
                                        var v118 uint64 = (v3 + (v3 & v3))
                                        _ = v118
                                    case "":
                                        v3 = v3
                                    }
                                    var v119 bool = true
                                    _ = v119
                                }
                            }
                        }
                        v120 := GenS7{F0: (v3 & v3), F1: (v2 - (v2 ^ (v2 - 1))), F2: 1}
                        var v121 uint = v2
                        _ = v120
                        _ = v121
                    } else {
                        var v122 bool = (307 != (v3 & (v3 >> ((v2) & 63))))
                        v5 = ("a" + v5)
                        _ = v122
                    }
                default:
                    v123 := GenS14{F0: 42, F1: fi48, F2: 22}
                    _ = v123
                }
                v2 = v2
                var v124 int8 = 1
                fmt.Println(v42)
                _ = fi49
                _ = v124
            }
            var v125 uint64 = v3
            _ = v125
        }
        for fi126 := 0; fi126 < 13; fi126++ {
            var v127 uint = 89
            v43 = append(v43, (fi126 >> ((v127) & 63)))
            _ = fi126
            _ = v127
        }
        _ = fi48
    }
    if false {
        var v128 float32 = v42
        v129 := GenS15{F0: 0, F1: 42}
        _ = v128
        _ = v129
    }
    if len(v43) >= 2 {
        v43 = v43[1:len(v43)]
    }
    if (!((v3 * v3) < v3)) {
        v1[("a" + v5)] = -408
    } else {
        fmt.Println(v42)
    }
    var v130 bool = (((("" + v5) + (v5 + ("foo bar" + ("foo bar" + v5)))) + (v5 + v5)) <= v5)
    for fi131 := 0; fi131 < 19; fi131++ {
        v5 = (v5 + v5)
        if true {
            var v132 uint64 = 4294967295
            _ = v132
        }
        _ = fi131
    }
    v2 = v2
    if v130 {
        if false {
            for fi133 := 0; fi133 < 2; fi133++ {
                fmt.Println(fi133, v5, fi133)
                v134 := GenS9{F0: 423, F1: 6}
                var v135 float64 = (-2.25 / 3.14159)
                _ = fi133
                _ = v134
                _ = v135
            }
            switch -956 {
            case -646:
                v42 += (float32((1.5 * 0.0)) * 3.14159)
            case 9223372036854775807:
                var v136 int64 = 473
                _ = v136
            case -2147483648:
                var v137 float64 = 0.1
                _ = v137
            }
        }
        switch 127 {
        case -2147483648:
            if (true && true) {
                delete(v1, ("hello" + ("\n\t" + (("\n\t" + v5) + "foo bar"))))
            } else {
                for fi138 := 0; fi138 < 1; fi138++ {
                    if len(v43) >= 2 {
                        v43 = v43[1:len(v43)]
                    }
                    _ = fi138
                }
                var v139 int32 = v4
                var v140 int = -1
                _ = v139
                _ = v140
            }
        case 42:
            var v141 int32 = 208
            _ = v141
        case 326:
            var v142 uint32 = 0
            _ = v142
        }
    }
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v130:", v130)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v130
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
