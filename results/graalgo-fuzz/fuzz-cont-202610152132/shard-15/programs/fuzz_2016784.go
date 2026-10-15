package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint64; F2 int8 }
type GenS3 struct { F0 int32; F1 bool; F2 bool }
type GenS4 struct { F0 bool; F1 string; F2 uint }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 int16; F1 string; F2 uint16 }
type GenS7 struct { F0 uint32; F1 int; F2 uint }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint32; F1 int }
type GenS12 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 string; F1 int64; F2 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{32768, 4294967295, 1, -1}
    v3 := []uint32{2147483648, 65535, 42}
    var v4 bool = false
    var v5 uint8 = 249
    var v6 string = "hello"
    v7 := GenS1{F0: v6}
    v4 = true
    if (100 > -118) {
        v6 = "foo bar"
        var v8 string = v6
        for fi9 := 0; fi9 < 17; fi9++ {
            for fi10 := 0; fi10 < 17; fi10++ {
                var v11 uint8 = v5
                v8 = ("hello" + ("foo bar" + v6))
                _ = fi10
                _ = v11
            }
            var v12 uint = 495
            v13 := GenS1{F0: v8}
            _ = fi9
            _ = v12
            _ = v13
        }
        _ = v8
    } else {
        v14 := GenS1{F0: "foo bar"}
        fmt.Println(v5, v5)
        if ((v6 + "\n\t") < v6) {
            if (8 != 42) {
                var v15 float32 = float32(3.14159)
                _ = v15
            } else {
                v16 := GenS1{F0: (v6 + v6)}
                for fi17 := 0; fi17 < 18; fi17++ {
                    fmt.Println(v5)
                    _ = fi17
                }
                fmt.Println(v4, v5, v4)
                _ = v16
            }
            fmt.Println(v6, v5)
        }
        _ = v14
    }
    var v18 int32 = -2147483648
    switch 4294967295 {
    case 133:
        v19 := GenS1{F0: "\n\t"}
        _ = v19
    case 2147483647:
        var v20 int32 = v18
        _ = v20
    case -1:
        for fi21 := 0; fi21 < 4; fi21++ {
            var v22 float64 = 1e10
            var v23 uint16 = 335
            _ = fi21
            _ = v22
            _ = v23
        }
    default:
        v5 = v5
    }
    fmt.Println(v6)
    v24 := GenS2{F0: "\n\t", F1: 887, F2: -128}
    v5 += 255
    _ = len(v3)
    var v25 int64 = 896
    switch 32768 {
    case 361:
        if v4 {
            for fi26 := 0; fi26 < 15; fi26++ {
                var v27 uint16 = 0
                var v28 int32 = 0
                _ = fi26
                _ = v27
                _ = v28
            }
            v4 = (-1 == 1)
            if ((0.0 / float64((1e10 - 0.0))) == 0.0) {
                fmt.Println(v6, v25, v25)
                v6 = ""
                if (v4 && (777 >= 2147483648)) {
                    switch v6 {
                    case "foo bar":
                        switch 9223372036854775807 {
                        case -2147483648:
                            switch -1 {
                            case 42:
                                var v29 float32 = 1.5
                                _ = v29
                            case -1:
                                v30 := v1[65535]
                                _ = v30
                            default:
                                if false {
                                    if v4 {
                                        if (!v4) {
                                            v18 = -61
                                            switch "foo bar" {
                                            case "\n\t":
                                                v31 := GenS2{F0: "\n\t", F1: 0, F2: -43}
                                                _ = v31
                                            default:
                                                var v32 uint = 978
                                                _ = v32
                                            }
                                        } else {
                                            _ = len(v3)
                                        }
                                        if v4 {
                                            if true {
                                                v33 := GenS3{F0: v18, F1: v4, F2: v4}
                                                v5 += (v5 - v5)
                                                v5 += v5
                                                _ = v33
                                            }
                                            for fi34 := 0; fi34 < 13; fi34++ {
                                                var v35 string = ((v6 + "hello") + ((v6 + "hello") + v6))
                                                fmt.Println(v18)
                                                _ = fi34
                                                _ = v35
                                            }
                                        }
                                    } else {
                                        var v36 uint64 = 18446744073709551615
                                        v37 := GenS4{F0: ((!(42 > 42)) && (1e10 > 0.1)), F1: ((v6 + (v6 + "hello")) + v6), F2: 614}
                                        if false {
                                            for fi38 := 0; fi38 < 4; fi38++ {
                                                fmt.Println(fi38)
                                                if ((v6 + v6) == (v6 + "")) {
                                                    fmt.Println(v25, v25)
                                                    v4 = v4
                                                    switch 42 {
                                                    case -1:
                                                        for fi39 := 0; fi39 < 13; fi39++ {
                                                            v40 := GenS1{F0: (("foo bar" + v6) + ("" + v6))}
                                                            var v41 uint32 = 2147483648
                                                            _ = fi39
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                    }
                                                }
                                                _ = fi38
                                            }
                                            var v42 uint32 = 65535
                                            _ = v42
                                        } else {
                                            switch -1 {
                                            case 127:
                                                var v43 float32 = 0.0
                                                _ = v43
                                            default:
                                                if true {
                                                    v44 := GenS5{F0: 695, F1: 18446744073709551615}
                                                    v45 := GenS2{F0: "", F1: v36, F2: 78}
                                                    _ = v44
                                                    _ = v45
                                                }
                                            }
                                            v18 = -2147483648
                                        }
                                        _ = v36
                                        _ = v37
                                    }
                                    var v46 int16 = 387
                                    _ = v46
                                }
                            }
                        case 550:
                            fmt.Println(v18, v6, v18)
                        }
                    case "a":
                        v47 := GenS3{F0: v18, F1: true, F2: true}
                        _ = v47
                    }
                    for fi48 := 0; fi48 < 17; fi48++ {
                        v49 := GenS6{F0: 0, F1: v6, F2: 42}
                        for fi50 := 0; fi50 < 5; fi50++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if v4 {
                                switch (fi50 | 4294967295) {
                                case 127:
                                    switch "a" {
                                    case "\n\t":
                                        fmt.Println(fi50, v25)
                                    case "hello":
                                        var v51 uint = 65535
                                        _ = v51
                                    default:
                                        switch (fi48 - (-(fi48))) {
                                        case 127:
                                            var v52 float32 = 0.0
                                            _ = v52
                                        default:
                                            if (v25 < (^(v25))) {
                                                var v53 float64 = 1.5
                                                v54 := GenS7{F0: 264, F1: fi48, F2: 2}
                                                _ = v53
                                                _ = v54
                                            }
                                        }
                                    }
                                case 0:
                                    switch (((v6 + v6) + "a") + v6) {
                                    case "hello":
                                        var v55 uint64 = 82
                                        _ = v55
                                    case "a":
                                        if (((((-1 >= 127) || (18446744073709551615 == 515)) || v4) || (false && (v18 <= -853))) || v4) {
                                            v25 += int64(v5)
                                        } else {
                                            switch "hello" {
                                            case "\n\t":
                                                delete(v1, 77)
                                            case "hello":
                                                var v56 uint64 = 514
                                                _ = v56
                                            default:
                                                v6 = v6
                                            }
                                            v3 = append(v3, 0)
                                        }
                                    case "foo bar":
                                        v57 := GenS4{F0: (v4 && v4), F1: v6, F2: 4294967295}
                                        _ = v57
                                    default:
                                        switch "foo bar" {
                                        case "hello":
                                            v5 = (v5 | v5)
                                        case "\n\t":
                                            if (95 > 18446744073709551615) {
                                                if true {
                                                    if (v6 != v6) {
                                                        fmt.Println(v6, v5, fi50)
                                                    } else {
                                                        if v4 {
                                                            switch v6 {
                                                            case "foo bar":
                                                                var v58 uint16 = 322
                                                                _ = v58
                                                            case "hello":
                                                                if v4 {
                                                                    v59 := v1[643]
                                                                    if (!true) {
                                                                        var v60 uint = 2
                                                                        var v61 bool = false
                                                                        _ = v60
                                                                        _ = v61
                                                                    } else {
                                                                        fmt.Println(fi48)
                                                                        v4 = (2 <= 0)
                                                                        v4 = (float64(float32((1e10 + 0.0))) == 3.14159)
                                                                    }
                                                                    _ = v59
                                                                } else {
                                                                    v25 -= 1
                                                                    var v62 uint64 = 18446744073709551615
                                                                    if (!v4) {
                                                                        v62 += (^(v62))
                                                                        var v63 bool = (v25 <= v25)
                                                                        v2 = append(v2, (^(fi50)))
                                                                        if v63 {
                                                                            v4 = (3.14159 >= 1e10)
                                                                            var v64 uint32 = 65535
                                                                            _ = v64
                                                                        } else {
                                                                            v5 += 42
                                                                            v4 = (!true)
                                                                        }
                                                                        _ = v63
                                                                    } else {
                                                                        var v65 bool = v4
                                                                        _ = v65
                                                                    }
                                                                    _ = v62
                                                                }
                                                            }
                                                            if v4 {
                                                                switch "hello" {
                                                                case "hello":
                                                                    if (!(!(false && true))) {
                                                                        var v66 int = fi50
                                                                        var v67 float64 = (1.5 * float64(3.14159))
                                                                        v67 = ((v67 - v67) / (((v67 + 1.5) + ((((v67 + 1e10) / v67) + float64(((0.0 + 1.5) * 3.14159))) - v67)) / v67))
                                                                        _ = v66
                                                                        _ = v67
                                                                    } else {
                                                                        if (127 > -1) {
                                                                            v68 := GenS8{F0: 42}
                                                                            var v69 uint16 = 255
                                                                            _ = v68
                                                                            _ = v69
                                                                        }
                                                                        if v4 {
                                                                            var v70 uint16 = 431
                                                                            _ = v70
                                                                        } else {
                                                                            v4 = v4
                                                                        }
                                                                    }
                                                                case "\n\t":
                                                                    switch (^(fi50)) {
                                                                    case -922:
                                                                        v25 -= 42
                                                                    case 128:
                                                                        if (2147483648 > 4294967295) {
                                                                            var v71 int8 = -49
                                                                            v5 += v5
                                                                            var v72 int32 = v18
                                                                            v5 = 88
                                                                            if (42 == 487) {
                                                                                var v73 uint8 = 127
                                                                                _ = v73
                                                                            }
                                                                            _ = v71
                                                                            _ = v72
                                                                        } else {
                                                                            switch 0 {
                                                                            case -1:
                                                                                v6 = ("" + v6)
                                                                            }
                                                                            var v74 float32 = (0.0 - float32(0.1))
                                                                            v75 := GenS9{F0: 604}
                                                                            _ = v74
                                                                            _ = v75
                                                                        }
                                                                    case 32767:
                                                                        var v76 int64 = -1
                                                                        _ = v76
                                                                    }
                                                                default:
                                                                    v4 = v4
                                                                }
                                                                if (((1.5 / 1.5) == 0.1) || v4) {
                                                                    v5 -= v5
                                                                    fmt.Println(v6, v4, fi48)
                                                                    fmt.Println(v18, fi50, v4)
                                                                }
                                                            } else {
                                                                if ((v18 | v18) == (v18 + v18)) {
                                                                    var v77 uint16 = 255
                                                                    var v78 int32 = v18
                                                                    v78 = v18
                                                                    if v4 {
                                                                        var v79 bool = true
                                                                        v6 = ((((v6 + ((((("a" + "a") + v6) + (v6 + (("a" + v6) + v6))) + (("\n\t" + (v6 + "hello")) + (v6 + v6))) + "hello")) + "foo bar") + ("hello" + (v6 + ((v6 + "hello") + "hello")))) + v6)
                                                                        if false {
                                                                            v25 = v25
                                                                        } else {
                                                                            if (true || v79) {
                                                                                v80 := GenS5{F0: 402, F1: 1}
                                                                                var v81 int32 = (v78 - (v78 - 42))
                                                                                var v82 uint64 = 42
                                                                                _ = v80
                                                                                _ = v81
                                                                                _ = v82
                                                                            }
                                                                            var v83 int64 = (v25 / (((v25 >> ((0) & 63))) | 1))
                                                                            var v84 string = v6
                                                                            _ = v83
                                                                            _ = v84
                                                                        }
                                                                        _ = v79
                                                                    }
                                                                    var v85 uint16 = (v77 | (v77 * 32767))
                                                                    v86 := GenS10{F0: 520}
                                                                    _ = v77
                                                                    _ = v78
                                                                    _ = v85
                                                                    _ = v86
                                                                }
                                                                var v87 int16 = 902
                                                                v88 := GenS7{F0: 0, F1: (-(fi48)), F2: 65535}
                                                                _ = v87
                                                                _ = v88
                                                            }
                                                        }
                                                        switch fi50 {
                                                        case 32767:
                                                            v89 := GenS8{F0: 65535}
                                                            _ = v89
                                                        case -716:
                                                            var v90 float32 = (0.1 * 1e10)
                                                            _ = v90
                                                        default:
                                                            fmt.Println(fi50, v5)
                                                        }
                                                        delete(v1, 609)
                                                        var v91 int8 = -122
                                                        _ = v91
                                                    }
                                                    if (284 >= (v25 & (-(v25)))) {
                                                        if (v5 != 66) {
                                                            v4 = v4
                                                            v6 = "\n\t"
                                                            var v92 float64 = 0.1
                                                            _ = v92
                                                        }
                                                        _ = len(v3)
                                                        delete(v1, 4294967295)
                                                    } else {
                                                        v93 := GenS11{F0: 2, F1: (fi50 ^ -9223372036854775808)}
                                                        var v94 float32 = ((float32((3.14159 / -2.25)) + (-2.25 - -2.25)) - 3.14159)
                                                        _ = v93
                                                        _ = v94
                                                    }
                                                    if (!(-1 < 42)) {
                                                        var v95 string = ((v6 + v6) + v6)
                                                        _ = v95
                                                    } else {
                                                        if true {
                                                            v25 += v25
                                                            v5 += v5
                                                        } else {
                                                            var v96 string = v6
                                                            var v97 uint8 = v5
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                        if (v4 && v4) {
                                                            var v98 float32 = 3.14159
                                                            v99 := GenS6{F0: 32767, F1: "foo bar", F2: 2}
                                                            v5 -= 42
                                                            _ = v98
                                                            _ = v99
                                                        } else {
                                                            v25 = -1
                                                        }
                                                        var v100 int64 = (^(v25))
                                                        if (((fi50 > fi48) && (!v4)) && (v4 || true)) {
                                                            switch ((v6 + (v6 + v6)) + (v6 + v6)) {
                                                            case "foo bar":
                                                                var v101 int16 = 32767
                                                                _ = v101
                                                            case "":
                                                                if len(v3) > 0 {
                                                                    v3[0] = 4294967295
                                                                }
                                                            case "\n\t":
                                                                _ = len(v3)
                                                            default:
                                                                if false {
                                                                    switch (-(fi48)) {
                                                                    case 9223372036854775807:
                                                                        fmt.Println(v100, fi48)
                                                                    case 4294967295:
                                                                        v5 = v5
                                                                    case -295:
                                                                        v102 := GenS12{F0: 2, F1: v5, F2: 0}
                                                                        _ = v102
                                                                    }
                                                                    if true {
                                                                        var v103 uint = 42
                                                                        v104 := GenS13{F0: 42}
                                                                        v105 := GenS13{F0: 580}
                                                                        _ = v103
                                                                        _ = v104
                                                                        _ = v105
                                                                    }
                                                                    v106 := GenS12{F0: 0, F1: 128, F2: 42}
                                                                    _ = v106
                                                                }
                                                            }
                                                            if (65535 <= 2) {
                                                                v18 = 2147483647
                                                                if ((v5 - v5) > (v5 & (v5 ^ (v5 * 153)))) {
                                                                    if v4 {
                                                                        if (0 != 0) {
                                                                            v5 += (^(v5))
                                                                            if len(v2) > 0 {
                                                                                v2[0] = fi48
                                                                            }
                                                                            v18 = (v18 + -2147483648)
                                                                        } else {
                                                                            if (1e10 <= (((0.1 / 1.5) / -2.25) / 0.1)) {
                                                                                switch (("hello" + v6) + (v6 + v6)) {
                                                                                case "\n\t":
                                                                                    v100 -= 0
                                                                                case "hello":
                                                                                    var v107 int8 = 0
                                                                                    _ = v107
                                                                                case "foo bar":
                                                                                    var v108 uint = 412
                                                                                    _ = v108
                                                                                default:
                                                                                    switch "\n\t" {
                                                                                    case "hello":
                                                                                        v18 = v18
                                                                                    default:
                                                                                        switch 9223372036854775807 {
                                                                                        case 128:
                                                                                            _ = len(v3)
                                                                                        case -9223372036854775808:
                                                                                            fmt.Println(fi50, fi50, v5)
                                                                                        default:
                                                                                            var v109 int32 = (v18 ^ v18)
                                                                                            v4 = ((342 != -1) || v4)
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            _ = v109
                                                                                        }
                                                                                    }
                                                                                }
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    v1[1] = 856
                                                                                default:
                                                                                    var v110 float32 = ((0.0 / float32((1e10 + 0.0))) + 3.14159)
                                                                                    _ = v110
                                                                                }
                                                                            } else {
                                                                                v4 = ((-(v18)) != 298)
                                                                            }
                                                                        }
                                                                        v1[2147483647] = 65535
                                                                    } else {
                                                                        var v111 uint8 = 1
                                                                        _ = v111
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v100
                                                    }
                                                }
                                                fmt.Println(v6, fi48)
                                                if v4 {
                                                    fmt.Println(fi48, fi50)
                                                    if false {
                                                        fmt.Println(v4, v6, v25)
                                                        fmt.Println(v18)
                                                    }
                                                } else {
                                                    if v4 {
                                                        v2 = append(v2, fi50)
                                                        if (-1 == (-(v18))) {
                                                            v25 = 0
                                                            var v112 uint64 = 1
                                                            _ = v112
                                                        }
                                                        v25 = 626
                                                    } else {
                                                        var v113 uint8 = (v5 / ((v5) | 1))
                                                        var v114 uint = 978
                                                        _ = v113
                                                        _ = v114
                                                    }
                                                }
                                            } else {
                                                var v115 uint32 = 736
                                                if v4 {
                                                    var v116 int16 = -1
                                                    var v117 int16 = v116
                                                    fmt.Println(v117, fi48)
                                                    v118 := GenS11{F0: 2, F1: (-(fi50))}
                                                    var v119 float32 = (0.0 * float32(3.14159))
                                                    _ = v116
                                                    _ = v117
                                                    _ = v118
                                                    _ = v119
                                                } else {
                                                    if (!false) {
                                                        var v120 bool = v4
                                                        _ = v120
                                                    } else {
                                                        if v4 {
                                                            v121 := GenS13{F0: (v115 * v115)}
                                                            _ = v121
                                                        }
                                                        var v122 uint16 = 0
                                                        if false {
                                                            var v123 int16 = -755
                                                            _ = v123
                                                        }
                                                        _ = v122
                                                    }
                                                }
                                                _ = v115
                                            }
                                        default:
                                            v18 += v18
                                        }
                                    }
                                }
                                var v124 uint8 = (^(v5))
                                _ = v124
                            } else {
                                var v125 uint16 = 65535
                                _ = v125
                            }
                            v6 = v6
                            var v126 uint = 190
                            var v127 uint64 = uint64(v126)
                            _ = fi50
                            _ = v126
                            _ = v127
                        }
                        v18 = (v18 + (v18 ^ v18))
                        _ = fi48
                        _ = v49
                    }
                } else {
                    var v128 bool = (v25 >= v25)
                    _ = v128
                }
            }
        } else {
            var v129 uint64 = 260
            var v130 uint32 = 42
            _ = len(v3)
            _ = v129
            _ = v130
        }
    case 2147483647:
        fmt.Println(v18)
    default:
        if (!(v6 > v6)) {
            switch 9223372036854775807 {
            case 0:
                v5 += 159
            case -1:
                fmt.Println(v4, v6, v6)
            }
        } else {
            v18 -= 1
        }
    }
    for fi131 := 0; fi131 < 3; fi131++ {
        v25 += v25
        if v4 {
            for fi132 := 0; fi132 < 7; fi132++ {
                v133 := GenS14{F0: v6, F1: (v25 & v25), F2: 32767}
                if v4 {
                    switch fi131 {
                    case -137:
                        _ = len(v3)
                    }
                    if (!true) {
                        fmt.Println(v5, v5, fi131)
                        v6 = "foo bar"
                    }
                }
                _ = fi132
                _ = v133
            }
        }
        _ = fi131
    }
    var v134 float64 = 3.14159
    v25 = v25
    v25 = v25
    var v135 uint16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v134:", v134)
    fmt.Println("v135:", v135)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v134
    _ = v135
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
