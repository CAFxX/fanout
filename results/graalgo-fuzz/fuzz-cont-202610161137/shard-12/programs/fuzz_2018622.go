package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int8 }
type GenS2 struct { F0 uint64; F1 int16; F2 string }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool; F1 string; F2 uint16 }
type GenS9 struct { F0 string; F1 string }
type GenS10 struct { F0 uint32; F1 bool }
type GenS11 struct { F0 uint64; F1 uint }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int; F1 uint32 }
type GenS14 struct { F0 uint64; F1 int64; F2 uint }
type GenS15 struct { F0 uint64; F1 uint; F2 int16 }
func genFn0(p0 int, p1 int16) bool {
    switch ("foo bar" + ("foo bar" + (("" + (("\n\t" + "hello") + "foo bar")) + (("foo bar" + "a") + (("a" + "a") + (("\n\t" + "foo bar") + "a")))))) {
    case "":
        var v7 int64 = 42
        _ = v7
    case "\n\t":
        switch (-(p0)) {
        case -846:
            var v8 bool = (!true)
            _ = v8
        case 2147483647:
            v9 := GenS1{F0: "\n\t", F1: 0, F2: 100}
            _ = v9
        case 42:
            if ((!("foo bar" < "a")) || true) {
                v10 := GenS1{F0: ("a" + "a"), F1: (^(p0)), F2: 0}
                v11 := GenS2{F0: 114, F1: (p1 >> ((0) & 63)), F2: ""}
                switch "" {
                case "a":
                    if true {
                        if true {
                            v12 := GenS1{F0: "", F1: p0, F2: 42}
                            p0 -= p0
                            _ = v12
                        }
                    }
                case "hello":
                    v13 := GenS3{F0: (p0 % (((p0 - -1)) | 1)), F1: -860}
                    _ = v13
                case "\n\t":
                    v14 := GenS1{F0: "foo bar", F1: 32768, F2: -128}
                    _ = v14
                }
                _ = v10
                _ = v11
            }
        }
    case "hello":
        v15 := GenS3{F0: p0, F1: -324}
        _ = v15
    }
    var v16 int16 = (^(p1))
    var v17 bool = true
    switch 0 {
    case -2147483648:
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 int16 = v16
            if (v17 || v17) {
                switch (-(p0)) {
                case 2147483647:
                    v19 += (^(v19))
                case 9223372036854775807:
                    fmt.Println(p1, p0, p0)
                case 831:
                    switch "foo bar" {
                    case "foo bar":
                        for fi20 := 0; fi20 < 18; fi20++ {
                            if true {
                                if (!true) {
                                    v19 = (v19 + (p1 >> ((593) & 63)))
                                } else {
                                    var v21 int32 = 42
                                    _ = v21
                                }
                                var v22 int32 = -152
                                _ = v22
                            }
                            _ = fi20
                        }
                    case "hello":
                        p1 = (p1 / (((-(p1))) | 1))
                    case "\n\t":
                        v23 := []int{(fi18 & (-(p0))), (fi18 >> ((2) & 63)), (fi18 + (fi18 & (fi18 % ((-9223372036854775808) | 1))))}
                        _ = v23
                    default:
                        v24 := GenS2{F0: 2, F1: (v19 << ((180) & 63)), F2: "foo bar"}
                        _ = v24
                    }
                default:
                    v25 := []int{fi18}
                    _ = v25
                }
            }
            var v26 int16 = (v19 ^ 632)
            _ = fi18
            _ = v19
            _ = v26
        }
    case 32768:
        v17 = (!false)
    default:
        for fi27 := 0; fi27 < 3; fi27++ {
            var v28 int64 = 1
            v17 = (368 > 591)
            for fi29 := 0; fi29 < 3; fi29++ {
                if v17 {
                    p1 = p1
                    v30 := []int{p0, fi29, fi29, fi29}
                    if len(v30) > 0 {
                        v30[0] = -356
                    }
                    _ = v30
                }
                _ = fi29
            }
            _ = fi27
            _ = v28
        }
    }
    fmt.Println(v17)
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return ((((("a" + "foo bar") + "foo bar") + (("hello" + ("\n\t" + ("hello" + "hello"))) + ("hello" + "\n\t"))) + ("foo bar" + ("foo bar" + (("\n\t" + ("a" + "hello")) + "foo bar")))) <= "a")
}
func genFn1() int32 {
    var v31 uint64 = 11
    v32 := []int{127, 554, 2147483647, -2147483648}
    var v33 uint = 2
    for fi34 := 0; fi34 < 6; fi34++ {
        v35 := GenS3{F0: (fi34 >> ((v33) & 63)), F1: -2147483648}
        _ = fi34
        _ = v35
    }
    _ = v31
    _ = v32
    _ = v33
    return -1
}

func main() {
    v1 := []int{127, -845, 1, -9223372036854775808}
    v2 := make(map[uint32]string)
    var v3 int16 = -1
    var v4 string = "hello"
    v5 := make(map[string]int)
    var v6 int32 = 182
    v4 = (("a" + v4) + "")
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        var v37 uint16 = 255
        switch fi36 {
        case -393:
            v38 := GenS4{F0: 255}
            _ = v38
        }
        switch 9223372036854775807 {
        case 32767:
            switch -906 {
            case 4294967295:
                var v39 int16 = (v3 % ((v3) | 1))
                _ = v39
            case 801:
                _ = len(v5)
            case 416:
                for fi40 := 0; fi40 < 9; fi40++ {
                    v41 := GenS4{F0: 255}
                    switch (fi40 + (fi36 % ((fi36) | 1))) {
                    case 32768:
                        fmt.Println(v37, v6, v37)
                    case 0:
                        if true {
                            v6 += (v6 - genFn1())
                        } else {
                            if true {
                                if false {
                                    v42 := GenS5{F0: 2}
                                    _ = v42
                                }
                                v3 -= (v3 * v3)
                                v4 = "hello"
                            } else {
                                var v43 bool = false
                                _ = len(v1)
                                fmt.Println(fi40, v3)
                                _ = v43
                            }
                        }
                    case 128:
                        var v44 float64 = 0.1
                        fmt.Println(v37, v4)
                        _ = v44
                    default:
                        if ((v37 >> ((0) & 63)) > v37) {
                            var v45 uint = 42
                            var v46 uint8 = 243
                            _ = v45
                            _ = v46
                        } else {
                            if ((v37 ^ (v37 ^ 42)) == v37) {
                                switch "hello" {
                                case "a":
                                    fmt.Println(v3, fi40, fi40)
                                case "\n\t":
                                    if false {
                                        v37 += 855
                                        if false {
                                            switch (fi36 - -89) {
                                            case 1:
                                                v2[556] = ("foo bar" + v4)
                                            case 32767:
                                                v6 = v6
                                            case 9223372036854775807:
                                                if false {
                                                    v47 := GenS6{F0: v4}
                                                    var v48 uint64 = 42
                                                    v49 := GenS3{F0: (fi36 - fi36), F1: (v6 / ((v6) | 1))}
                                                    _ = v47
                                                    _ = v48
                                                    _ = v49
                                                } else {
                                                    var v50 uint64 = 4294967295
                                                    if false {
                                                        var v51 uint8 = 2
                                                        v4 = v4
                                                        v51 = (^(v51))
                                                        _ = v51
                                                    }
                                                    v52 := GenS7{F0: 892}
                                                    _ = v50
                                                    _ = v52
                                                }
                                            default:
                                                v3 = (v3 ^ (v3 * 587))
                                            }
                                        }
                                    } else {
                                        var v53 int32 = genFn1()
                                        fmt.Println(fi36, v3)
                                        _ = v53
                                    }
                                case "":
                                    var v54 uint16 = v37
                                    _ = v54
                                default:
                                    v55 := GenS2{F0: 42, F1: (v3 / (((v3 & (-(v3)))) | 1)), F2: (("foo bar" + "a") + (v4 + "a"))}
                                    _ = v55
                                }
                            }
                            var v56 uint16 = v37
                            v1 = append(v1, (^(fi36)))
                            _ = v56
                        }
                    }
                    v57 := GenS2{F0: 2, F1: (^(v3)), F2: "hello"}
                    _ = fi40
                    _ = v41
                    _ = v57
                }
            default:
                v4 = ""
            }
        case -668:
            for fi58 := 0; fi58 < 7; fi58++ {
                if (true && true) {
                    if false {
                        v37 = v37
                        var v59 uint16 = v37
                        var v60 int32 = (v6 / ((-2147483648) | 1))
                        _ = v59
                        _ = v60
                    }
                    var v61 uint8 = 0
                    fmt.Println(fi36, v4)
                    _ = v61
                }
                if true {
                    v3 = v3
                    if (1e10 == float32(((0.0 - (float64((((1e10 * float32(1e10)) / 0.1) + (float32(float64((-2.25 / 1e10))) + (-2.25 - (1e10 * (float32(-2.25) / 1.0)))))) - 1.5)) + 3.14159))) {
                        var v62 int32 = -2147483648
                        v4 = (v4 + v4)
                        v63 := GenS3{F0: fi58, F1: (^(v62))}
                        v6 -= v6
                        _ = v62
                        _ = v63
                    }
                    var v64 uint64 = 0
                    fmt.Println(v64, v64)
                    fmt.Println(v4)
                    _ = v64
                } else {
                    var v65 int8 = -1
                    v3 = (-(v3))
                    _ = v65
                }
                var v66 int = -153
                _ = fi58
                _ = v66
            }
        case -1:
            var v67 uint16 = v37
            _ = v67
        }
        _ = fi36
        _ = v37
    }
    switch 2147483647 {
    case -2147483648:
        v68 := GenS1{F0: ("" + (("foo bar" + v4) + "foo bar")), F1: -1, F2: -128}
        _ = v68
    case 32768:
        switch 0 {
        case -9223372036854775808:
            if false {
                switch 0 {
                case 1:
                    for fi69 := 0; fi69 < 14; fi69++ {
                        v70 := GenS2{F0: 4294967295, F1: v3, F2: (v4 + v4)}
                        if true {
                            switch "hello" {
                            case "a":
                                var v71 float64 = (((0.0 + -2.25) - (0.1 - -2.25)) * 0.1)
                                _ = v71
                            case "hello":
                                v72 := GenS4{F0: 728}
                                _ = v72
                            }
                            v4 = (v4 + v4)
                            var v73 uint64 = 18446744073709551615
                            _ = v73
                        } else {
                            v6 = (-(v6))
                        }
                        _ = fi69
                        _ = v70
                    }
                default:
                    switch 0 {
                    case 32768:
                        var v74 uint32 = 4294967295
                        _ = v74
                    }
                }
                switch 1 {
                case 32768:
                    for fi75 := 0; fi75 < 4; fi75++ {
                        switch v4 {
                        case "":
                            _ = len(v1)
                        case "hello":
                            for fi76 := 0; fi76 < 4; fi76++ {
                                v6 += 1
                                v77 := v2[1]
                                var v78 int8 = -76
                                v2[0] = "\n\t"
                                _ = fi76
                                _ = v77
                                _ = v78
                            }
                        case "foo bar":
                            switch v4 {
                            case "hello":
                                v3 += v3
                            case "\n\t":
                                v79 := GenS5{F0: 2}
                                _ = v79
                            }
                        default:
                            _ = len(v5)
                        }
                        if true {
                            v4 = v4
                        }
                        if (!(255 <= 2)) {
                            if true {
                                if true {
                                    if false {
                                        var v80 uint = 617
                                        switch (v4 + v4) {
                                        case "\n\t":
                                            v81 := GenS8{F0: true, F1: ("foo bar" + ""), F2: 2}
                                            _ = v81
                                        case "a":
                                            v82 := GenS9{F0: v4, F1: v4}
                                            _ = v82
                                        default:
                                            for fi83 := 0; fi83 < 17; fi83++ {
                                                var v84 float32 = 0.0
                                                var v85 string = ("" + v4)
                                                _ = fi83
                                                _ = v84
                                                _ = v85
                                            }
                                        }
                                        var v86 uint16 = 0
                                        _ = v80
                                        _ = v86
                                    } else {
                                        for fi87 := 0; fi87 < 14; fi87++ {
                                            fmt.Println(v3, fi75, v3)
                                            v3 -= (v3 | v3)
                                            if len(v1) > 0 {
                                                v1[0] = fi75
                                            }
                                            _ = fi87
                                        }
                                        v3 -= v3
                                        v5[v4] = fi75
                                    }
                                    if (true && ((fi75 % (((fi75 * (fi75 ^ (fi75 << ((1) & 63))))) | 1)) >= fi75)) {
                                        v6 += v6
                                    }
                                    if (true && true) {
                                        for fi88 := 0; fi88 < 19; fi88++ {
                                            var v89 float32 = (0.0 - 1e10)
                                            _ = fi88
                                            _ = v89
                                        }
                                    }
                                }
                                _ = len(v5)
                            } else {
                                v90 := GenS2{F0: 1, F1: v3, F2: v4}
                                _ = v90
                            }
                            v3 -= v3
                        }
                        _ = fi75
                    }
                case 128:
                    v3 -= v3
                }
                v91 := GenS6{F0: (v4 + v4)}
                _ = v91
            }
        case 378:
            switch 2147483647 {
            case 9223372036854775807:
                if (127 <= 100) {
                    for fi92 := 0; fi92 < 8; fi92++ {
                        v93 := GenS4{F0: 18446744073709551615}
                        _ = fi92
                        _ = v93
                    }
                    if true {
                        fmt.Println(v3, v4, v4)
                        if len(v1) > 0 {
                            v1[0] = -1
                        }
                        if (v4 >= v4) {
                            v94 := GenS10{F0: 1, F1: ((927 <= 42) || false)}
                            _ = v94
                        }
                    } else {
                        var v95 uint16 = 1
                        v96 := GenS8{F0: (("\n\t" + v4) > (v4 + "")), F1: "foo bar", F2: (v95 ^ (v95 * 115))}
                        if (-1 < -533) {
                            switch -9223372036854775808 {
                            case 4294967295:
                                if (!(42 < 0)) {
                                    var v97 bool = false
                                    if v97 {
                                        for fi98 := 0; fi98 < 10; fi98++ {
                                            v6 = v6
                                            if (18446744073709551615 < 4294967295) {
                                                var v99 uint8 = 42
                                                var v100 uint = 531
                                                for fi101 := 0; fi101 < 20; fi101++ {
                                                    v102 := GenS11{F0: 0, F1: uint(v99)}
                                                    v103 := GenS12{F0: 42}
                                                    _ = fi101
                                                    _ = v102
                                                    _ = v103
                                                }
                                                _ = v100
                                                _ = v99
                                            }
                                            switch ("" + "\n\t") {
                                            case "hello":
                                                v104 := v5[v4]
                                                _ = v104
                                            case "foo bar":
                                                v3 -= v3
                                            case "":
                                                v3 += (v3 >> ((0) & 63))
                                            }
                                            _ = fi98
                                        }
                                        if ((-2.25 - 1.5) > -2.25) {
                                            v2[1] = v4
                                        } else {
                                            v3 = v3
                                            v6 = (v6 & v6)
                                        }
                                        v95 += v95
                                    } else {
                                        var v105 uint8 = 42
                                        _ = v105
                                    }
                                    var v106 bool = v97
                                    _ = v106
                                    _ = v97
                                }
                            default:
                                v107 := v5[v4]
                                _ = v107
                            }
                            v108 := GenS12{F0: 4294967295}
                            _ = v108
                        } else {
                            for fi109 := 0; fi109 < 12; fi109++ {
                                var v110 int64 = 493
                                if true {
                                    var v111 uint64 = 927
                                    v4 = ((v4 + "hello") + "foo bar")
                                    var v112 int = fi109
                                    _ = v111
                                    _ = v112
                                }
                                _ = fi109
                                _ = v110
                            }
                        }
                        _ = v95
                        _ = v96
                    }
                }
            case -356:
                v5["foo bar"] = -1
            case 57:
                if true {
                    v4 = (((v4 + v4) + "") + "hello")
                    for fi113 := 0; fi113 < 1; fi113++ {
                        if len(v1) > 0 {
                            v1[0] = fi113
                        }
                        v4 = "\n\t"
                        _ = fi113
                    }
                } else {
                    var v114 int8 = 0
                    fmt.Println(v6, v6, v3)
                    var v115 float64 = 1e10
                    _ = v114
                    _ = v115
                }
            }
        default:
            _ = len(v2)
        }
    }
    v3 += (v3 ^ (v3 ^ 32767))
    v116 := GenS5{F0: 18446744073709551615}
    if len(v1) > 0 {
        v1[0] = 0
    }
    v4 = ((v4 + v4) + "\n\t")
    switch -2147483648 {
    case 32767:
        v6 = v6
    default:
        v4 = v4
    }
    v117 := GenS7{F0: 2}
    var v118 uint32 = 606
    switch int(v3) {
    case 281:
        v118 += 0
    }
    if (4294967295 < 0) {
        switch 128 {
        case 32768:
            var v119 uint = 255
            _ = v119
        }
        v6 -= -1
    } else {
        v3 = (v3 - (-(v3)))
        if (65535 >= uint16(v118)) {
            v3 += int16(v118)
            var v120 int64 = 42
            _ = v120
        } else {
            if false {
                for fi121 := 0; fi121 < 15; fi121++ {
                    if (!true) {
                        var v122 int16 = v3
                        fmt.Println(v4)
                        fmt.Println(v118)
                        _ = v122
                    }
                    for fi123 := 0; fi123 < 17; fi123++ {
                        _ = len(v2)
                        v124 := GenS7{F0: (v118 >> ((65535) & 63))}
                        _ = fi123
                        _ = v124
                    }
                    v125 := GenS9{F0: (("\n\t" + v4) + v4), F1: v4}
                    _ = fi121
                    _ = v125
                }
            } else {
                if true {
                    v3 = v3
                    _ = len(v1)
                    for fi126 := 0; fi126 < 12; fi126++ {
                        v127 := GenS13{F0: fi126, F1: (v118 | v118)}
                        var v128 float32 = 1e10
                        _ = fi126
                        _ = v127
                        _ = v128
                    }
                }
                v129 := GenS14{F0: 458, F1: 0, F2: 65535}
                _ = v129
            }
            v118 += v118
            for fi130 := 0; fi130 < 20; fi130++ {
                v4 = (((("a" + (v4 + "a")) + v4) + v4) + v4)
                _ = fi130
            }
        }
        switch v4 {
        case "":
            v131 := GenS15{F0: 321, F1: 2, F2: v3}
            _ = v131
        }
    }
    _ = len(v2)
    switch v4 {
    case "a":
        if (47 > 237) {
            v4 = (v4 + v4)
            v5[(((("foo bar" + v4) + v4) + (((v4 + "hello") + v4) + v4)) + v4)] = 9223372036854775807
        } else {
            var v132 int = 389
            _ = v132
        }
    case "\n\t":
        var v133 int16 = (v3 + 1000)
        _ = v133
    }
    v1 = append(v1, 32767)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v118:", v118)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v116
    _ = v117
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
