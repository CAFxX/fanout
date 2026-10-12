package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int8; F2 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 uint16; F2 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint8; F1 int; F2 string }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint16; F1 string; F2 bool }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 string; F1 uint32; F2 uint }
func genFn0(p0 uint8, p1 int) uint {
    for fi7 := 0; fi7 < 18; fi7++ {
        p1 -= (p1 ^ (p1 - -1))
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS1{F0: -9223372036854775808}
            var v10 uint32 = 299
            v10 -= (^(v10))
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    var v11 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = v11
    return 0
}

func main() {
    var v1 uint8 = 0
    var v2 int32 = 2147483647
    v3 := make(map[int]uint32)
    var v4 bool = (3.14159 <= (((-2.25 - 0.1) / float32((3.14159 * ((0.0 + 3.14159) - 0.0)))) + (float32((3.14159 * 0.1)) / -2.25)))
    v5 := make(map[string]string)
    var v6 int = -443
    if v4 {
        for fi12 := 0; fi12 < 16; fi12++ {
            v6 -= (v6 % ((-109) | 1))
            v13 := []int{fi12, 2147483647}
            _ = fi12
            _ = v13
        }
    } else {
        v14 := GenS1{F0: -1}
        for fi15 := 0; fi15 < 1; fi15++ {
            if v4 {
                switch -1 {
                case 0:
                    if true {
                        fmt.Println(fi15, v2)
                        switch v6 {
                        case 1:
                            for fi16 := 0; fi16 < 12; fi16++ {
                                if true {
                                    v1 = (v1 | (^(v1)))
                                    v2 = 42
                                    var v17 float64 = float64(0.0)
                                    _ = v17
                                } else {
                                    var v18 bool = (0.1 == ((0.0 + float32(-2.25)) / ((0.1 / -2.25) / 1.5)))
                                    switch ((((("foo bar" + "\n\t") + "foo bar") + "hello") + "\n\t") + "hello") {
                                    case "":
                                        var v19 string = (("" + "a") + "foo bar")
                                        _ = v19
                                    case "\n\t":
                                        if ((737 <= 475) || v18) {
                                            if v4 {
                                                v2 = (-(v2))
                                                v3[fi15] = 529
                                                var v20 uint8 = (^(v1))
                                                _ = v20
                                            }
                                        } else {
                                            var v21 uint8 = (^(v1))
                                            _ = v21
                                        }
                                    default:
                                        v1 = v1
                                    }
                                    v22 := []int{(fi15 ^ (v6 + fi16)), (v6 + v6), v6}
                                    _ = v18
                                    _ = v22
                                }
                                var v23 bool = v4
                                v4 = v23
                                _ = fi16
                                _ = v23
                            }
                        case 32767:
                            switch "a" {
                            case "hello":
                                v24 := v5["\n\t"]
                                _ = v24
                            default:
                                switch "a" {
                                case "hello":
                                    if v4 {
                                        v25 := []int{(v6 >> ((2) & 63))}
                                        _ = v25
                                    }
                                case "\n\t":
                                    if (float32(float64(float32((1.5 - 0.1)))) < (-2.25 - ((((1.5 + (float32((1e10 * 1.5)) / (0.0 + -2.25))) - (1.5 + ((((1e10 + 1e10) / 1e10) / (float32(1e10) - (3.14159 * ((((-2.25 - -2.25) / 0.1) - 3.14159) + 0.1)))) - -2.25))) + 0.1) / 3.14159))) {
                                        var v26 int32 = v2
                                        _ = len(v3)
                                        var v27 int8 = -70
                                        _ = v26
                                        _ = v27
                                    } else {
                                        var v28 float32 = (0.0 + 3.14159)
                                        fmt.Println(fi15)
                                        var v29 uint16 = 18
                                        _ = v28
                                        _ = v29
                                    }
                                case "":
                                    var v30 bool = (!((0.1 - ((-2.25 - (((-2.25 * 0.1) + 3.14159) - (1.5 + 1.5))) * 0.0)) <= ((1.5 - (1.5 * -2.25)) + 1.5)))
                                    _ = v30
                                }
                            }
                        default:
                            for fi31 := 0; fi31 < 8; fi31++ {
                                v32 := GenS1{F0: 1}
                                if (286 == 18446744073709551615) {
                                    var v33 int64 = -17
                                    var v34 int16 = -1
                                    v35 := []int{(v6 | (fi31 + fi15)), fi15, fi31, 32767}
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                } else {
                                    v1 = (v1 / ((v1) | 1))
                                    switch v6 {
                                    case -9223372036854775808:
                                        fmt.Println(v1)
                                    }
                                }
                                if (-1 != -128) {
                                    v2 += (v2 % (((v2 / ((int32(v1)) | 1))) | 1))
                                    switch (-(fi31)) {
                                    case 32767:
                                        fmt.Println(v1, v6, v1)
                                    default:
                                        v6 += 1
                                    }
                                    v4 = (65535 <= 65535)
                                }
                                _ = fi31
                                _ = v32
                            }
                        }
                    } else {
                        for fi36 := 0; fi36 < 9; fi36++ {
                            if (42 < 255) {
                                var v37 bool = (583 <= 1000)
                                _ = v37
                            }
                            var v38 uint64 = 1
                            v39 := v5["\n\t"]
                            _ = fi36
                            _ = v38
                            _ = v39
                        }
                        switch ("foo bar" + "hello") {
                        case "\n\t":
                            if v4 {
                                v40 := GenS1{F0: 1}
                                for fi41 := 0; fi41 < 5; fi41++ {
                                    var v42 uint = 4294967295
                                    v42 += 2
                                    _ = fi41
                                    _ = v42
                                }
                                v43 := []int{}
                                _ = v40
                                _ = v43
                            } else {
                                for fi44 := 0; fi44 < 10; fi44++ {
                                    switch "foo bar" {
                                    case "hello":
                                        switch "foo bar" {
                                        case "foo bar":
                                            v2 = (v2 ^ v2)
                                        case "hello":
                                            v6 = fi44
                                            switch "foo bar" {
                                            case "\n\t":
                                                switch (("hello" + "hello") + ("a" + "\n\t")) {
                                                case "a":
                                                    var v45 uint = 18446744073709551615
                                                    _ = v45
                                                case "":
                                                    var v46 int = fi44
                                                    _ = v46
                                                default:
                                                    if false {
                                                        v47 := []int{(fi44 / (((fi44 >> ((65535) & 63))) | 1)), (fi15 & v6), (fi15 >> ((1) & 63)), (fi44 % ((2147483647) | 1))}
                                                        _ = v47
                                                    }
                                                }
                                            case "hello":
                                                var v48 uint = 42
                                                _ = v48
                                            case "foo bar":
                                                if (2 > 1) {
                                                    switch (fi15 | fi15) {
                                                    case -56:
                                                        fmt.Println(fi15, v1, fi15)
                                                    case 127:
                                                        if true {
                                                            var v49 float32 = (0.0 + (0.0 * (1e10 - -2.25)))
                                                            if v4 {
                                                                fmt.Println(v1)
                                                                var v50 int16 = 0
                                                                _ = v50
                                                            }
                                                            v51 := []int{(fi44 ^ (v6 + fi15)), (fi44 & (^(v6))), 32768}
                                                            fmt.Println(v2, v49)
                                                            _ = v49
                                                            _ = v51
                                                        } else {
                                                            v52 := GenS1{F0: 1}
                                                            _ = len(v3)
                                                            if true {
                                                                if v4 {
                                                                    var v53 string = ""
                                                                    _ = v53
                                                                }
                                                            }
                                                            v4 = ((v2 << ((genFn0((v1 ^ (^(v1))), fi15)) & 63)) <= v2)
                                                            _ = v52
                                                        }
                                                    default:
                                                        v54 := GenS1{F0: -1}
                                                        _ = v54
                                                    }
                                                    var v55 uint64 = 368
                                                    _ = v55
                                                }
                                            }
                                        default:
                                            var v56 bool = (4294967295 >= 945)
                                            _ = v56
                                        }
                                    default:
                                        fmt.Println(v1, v4)
                                    }
                                    v57 := []int{}
                                    var v58 uint8 = (v1 & 1)
                                    v59 := GenS1{F0: 9223372036854775807}
                                    _ = fi44
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                }
                                fmt.Println(v4, v6, v2)
                                v60 := GenS1{F0: -134}
                                _ = v60
                            }
                        case "":
                            v61 := GenS2{F0: v1, F1: -15, F2: 127}
                            _ = v61
                        }
                    }
                case 1:
                    v2 += v2
                case -9223372036854775808:
                    var v62 uint8 = v1
                    _ = v62
                default:
                    var v63 uint32 = 4294967295
                    _ = v63
                }
                if v4 {
                    v4 = v4
                } else {
                    var v64 uint64 = 255
                    v65 := v3[v6]
                    _ = v64
                    _ = v65
                }
                v2 = -483
            }
            v66 := []int{}
            _ = fi15
            _ = v66
        }
        fmt.Println(v1, v1, v4)
        _ = v14
    }
    var v67 string = "\n\t"
    switch ("hello" + v67) {
    case "":
        fmt.Println(v2, v67, v4)
    case "\n\t":
        v6 -= 2147483647
    }
    var v68 int64 = 42
    v69 := GenS3{F0: (65535 <= 32767)}
    v70 := GenS4{F0: genFn0((v1 / ((v1) | 1)), (v6 % ((v6) | 1))), F1: 811, F2: (v6 ^ 9223372036854775807)}
    v6 -= (v6 - int(v1))
    v71 := GenS5{F0: 0}
    for fi72 := 0; fi72 < 6; fi72++ {
        var v73 uint8 = (v1 % ((v1) | 1))
        _ = len(v5)
        v3[32767] = 42
        _ = fi72
        _ = v73
    }
    v1 = 132
    v74 := GenS5{F0: 2147483647}
    switch (v6 / ((v6) | 1)) {
    case -9223372036854775808:
        for fi75 := 0; fi75 < 2; fi75++ {
            switch 9223372036854775807 {
            case 127:
                v76 := GenS6{F0: (v1 == v1), F1: 2}
                _ = v76
            }
            if (329 == 891) {
                v67 = v67
                if (-124 < -32768) {
                    var v77 int16 = 1
                    for fi78 := 0; fi78 < 14; fi78++ {
                        fmt.Println(v77)
                        v79 := []int{fi78}
                        var v80 float64 = 1.5
                        _ = fi78
                        _ = v79
                        _ = v80
                    }
                    _ = v77
                } else {
                    v67 = ("" + v67)
                }
            }
            _ = fi75
        }
    case 42:
        v81 := GenS1{F0: (v68 << ((42) & 63))}
        _ = v81
    case -1:
        v82 := GenS2{F0: 255, F1: -128, F2: 1}
        _ = v82
    default:
        switch (v6 / ((v6) | 1)) {
        case -1:
            fmt.Println(v67)
        case 42:
            v1 = v1
        case -434:
            v83 := GenS6{F0: (!(v68 == (v68 / ((0) | 1)))), F1: 0}
            _ = v83
        }
    }
    var v84 uint64 = 1
    fmt.Println(v1, v2, v4)
    if false {
        switch v67 {
        case "":
            v85 := GenS7{F0: 302}
            _ = v85
        case "a":
            if (956 >= 723) {
                v6 = v6
            }
        case "foo bar":
            switch (v6 * v6) {
            case 781:
                var v86 uint8 = (^(v1))
                _ = v86
            default:
                var v87 string = (((v67 + v67) + v67) + "a")
                _ = v87
            }
        }
        if v4 {
            fmt.Println(v68, v68, v4)
        }
    }
    if (((v67 + (v67 + v67)) + "\n\t") == "a") {
        for fi88 := 0; fi88 < 16; fi88++ {
            if (v4 && v4) {
                v84 = v84
                var v89 float64 = 0.1
                _ = v89
            } else {
                var v90 int16 = -32768
                _ = v90
            }
            _ = fi88
        }
        if v4 {
            v1 = (v1 >> ((4294967295) & 63))
            fmt.Println(v1)
            var v91 int32 = -790
            _ = v91
        } else {
            v92 := []int{(v6 | v6), (v6 ^ (v6 * v6)), v6, -9223372036854775808}
            if (v4 || ((!v4) && (869 < 536))) {
                if v4 {
                    for fi93 := 0; fi93 < 6; fi93++ {
                        var v94 uint64 = 18446744073709551615
                        var v95 uint = 539
                        _ = fi93
                        _ = v94
                        _ = v95
                    }
                    v2 += (v2 >> ((65535) & 63))
                }
                var v96 int = (v6 | v6)
                for fi97 := 0; fi97 < 20; fi97++ {
                    for fi98 := 0; fi98 < 15; fi98++ {
                        if (int8(v2) != 127) {
                            var v99 bool = true
                            v5[(v67 + v67)] = ((((v67 + v67) + (v67 + v67)) + ((("hello" + "foo bar") + "foo bar") + v67)) + "\n\t")
                            fmt.Println(v1, fi97, v2)
                            _ = v99
                        } else {
                            fmt.Println(v84)
                            var v100 uint8 = (^(v1))
                            _ = v100
                        }
                        _ = fi98
                    }
                    _ = fi97
                }
                _ = v96
            } else {
                if v4 {
                    var v101 bool = ("" != "hello")
                    v4 = (true || ((1e10 / float64(-2.25)) >= 1e10))
                    _ = v101
                } else {
                    v67 = (((v67 + (((((v67 + "\n\t") + ("hello" + "hello")) + v67) + (v67 + v67)) + v67)) + v67) + "")
                }
                if false {
                    v6 = (v6 * (v6 + v6))
                    for fi102 := 0; fi102 < 4; fi102++ {
                        for fi103 := 0; fi103 < 9; fi103++ {
                            var v104 int64 = v68
                            fmt.Println(v68)
                            _ = fi103
                            _ = v104
                        }
                        _ = fi102
                    }
                    for fi105 := 0; fi105 < 19; fi105++ {
                        v106 := GenS4{F0: 18446744073709551615, F1: 443, F2: 0}
                        switch fi105 {
                        case 769:
                            for fi107 := 0; fi107 < 1; fi107++ {
                                delete(v5, v67)
                                var v108 bool = true
                                _ = fi107
                                _ = v108
                            }
                        default:
                            _ = len(v3)
                        }
                        switch (((("hello" + (v67 + v67)) + v67) + v67) + (v67 + "\n\t")) {
                        case "foo bar":
                            v109 := GenS4{F0: genFn0(2, (fi105 - fi105)), F1: 614, F2: (fi105 << ((42) & 63))}
                            _ = v109
                        case "hello":
                            switch v67 {
                            case "a":
                                for fi110 := 0; fi110 < 9; fi110++ {
                                    if false {
                                        if v4 {
                                            var v111 uint8 = 208
                                            fmt.Println(v4, v6)
                                            v112 := GenS1{F0: (v68 | 296)}
                                            _ = v111
                                            _ = v112
                                        }
                                        v84 += (v84 >> ((42) & 63))
                                    } else {
                                        if v4 {
                                            v6 = (-(v6))
                                        } else {
                                            if (v4 || (true || (935 < 255))) {
                                                if (127 != -1) {
                                                    fmt.Println(fi105, fi105, fi105)
                                                    fmt.Println(v4, v6, v4)
                                                } else {
                                                    var v113 uint32 = 2147483648
                                                    if ((((v67 + "foo bar") + ("foo bar" + ("" + v67))) + v67) != v67) {
                                                        var v114 int8 = 1
                                                        v115 := GenS4{F0: 497, F1: uint16(v113), F2: v6}
                                                        _ = v114
                                                        _ = v115
                                                    }
                                                    v116 := v5[(((v67 + "hello") + (v67 + ("" + "\n\t"))) + "hello")]
                                                    _ = v113
                                                    _ = v116
                                                }
                                            } else {
                                                v117 := GenS8{F0: v1, F1: (-(v6)), F2: (v67 + "")}
                                                var v118 uint32 = 2
                                                if v4 {
                                                    if false {
                                                        v119 := GenS9{F0: v68}
                                                        v4 = v4
                                                        _ = v119
                                                    }
                                                    var v120 bool = (!false)
                                                    v121 := GenS10{F0: 42, F1: v67, F2: v4}
                                                    v3[v6] = (v118 / ((v118) | 1))
                                                    var v122 float64 = 0.1
                                                    v120 = v4
                                                    _ = v120
                                                    _ = v121
                                                    _ = v122
                                                }
                                                v5["a"] = v67
                                                _ = v117
                                                _ = v118
                                            }
                                        }
                                    }
                                    _ = fi110
                                }
                            case "\n\t":
                                fmt.Println(v68, v68, v2)
                            case "foo bar":
                                var v123 uint32 = 215
                                _ = v123
                            }
                        case "":
                            if v4 {
                                v2 = (v2 ^ (^(v2)))
                                var v124 int = 733
                                _ = v124
                            } else {
                                v125 := GenS9{F0: v68}
                                var v126 int16 = -32768
                                v4 = (v4 || false)
                                _ = v125
                                _ = v126
                            }
                        default:
                            for fi127 := 0; fi127 < 19; fi127++ {
                                v128 := GenS9{F0: -441}
                                if true {
                                    v1 -= (v1 - 107)
                                    v4 = (v67 <= (v67 + "\n\t"))
                                } else {
                                    fmt.Println(fi127)
                                }
                                if v4 {
                                    var v129 bool = v4
                                    switch "" {
                                    case "a":
                                        v68 = 13
                                        if ((v2 | 519) > v2) {
                                            v130 := GenS11{F0: 296}
                                            v1 = 42
                                            var v131 int64 = (v68 ^ (v68 % (((v68 + v68)) | 1)))
                                            var v132 float32 = (3.14159 / 0.1)
                                            var v133 float32 = v132
                                            _ = v130
                                            _ = v131
                                            _ = v132
                                            _ = v133
                                        }
                                        fmt.Println(v67, v1)
                                        if ((v67 < (("foo bar" + (v67 + "foo bar")) + (v67 + (v67 + "foo bar")))) || ((v67 != "a") || (fi105 > 388))) {
                                            v3[v6] = 422
                                        }
                                        v6 = (-(fi105))
                                        v68 -= -1
                                    case "hello":
                                        fmt.Println(v68, fi105, v6)
                                    }
                                    v134 := v3[-1]
                                    _ = v129
                                    _ = v134
                                }
                                _ = fi127
                                _ = v128
                            }
                        }
                        _ = fi105
                        _ = v106
                    }
                }
                v67 = "\n\t"
            }
            _ = v92
        }
    }
    for fi135 := 0; fi135 < 19; fi135++ {
        var v136 int32 = (v2 * (v2 << ((2) & 63)))
        switch "hello" {
        case "a":
            var v137 uint = 607
            _ = v137
        case "":
            for fi138 := 0; fi138 < 10; fi138++ {
                var v139 uint16 = 42
                v140 := []int{v6, (fi135 % ((fi138) | 1)), fi135, 128}
                v141 := GenS12{F0: "", F1: 1, F2: 0}
                var v142 int8 = -1
                var v143 bool = (((127 != (^(v1))) && (!((false || (true || ((v142 ^ int8(v139)) < (v142 >> ((2) & 63))))) || (42 <= 127)))) || v4)
                _ = fi138
                _ = v139
                _ = v140
                _ = v141
                _ = v142
                _ = v143
            }
        }
        _ = fi135
        _ = v136
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
    _ = v69
    _ = v70
    _ = v71
    _ = v74
    _ = v84
}
