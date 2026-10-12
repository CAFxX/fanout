package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int64; F1 uint16 }
type GenS10 struct { F0 int32; F1 int64 }
type GenS11 struct { F0 string; F1 uint }
type GenS12 struct { F0 int64; F1 int16 }
func genFn0(p0 uint8, p1 int32, p2 int) int {
    v6 := []int{-690, p2, p2, 4294967295}
    var v7 int64 = -9223372036854775808
    var v8 float32 = (3.14159 - float32(1e10))
    p2 += p2
    var v9 int64 = int64(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return p2
}
func genFn1(p0 bool, p1 uint) int32 {
    switch (("hello" + "hello") + (("a" + "hello") + "a")) {
    case "hello":
        for fi10 := 0; fi10 < 12; fi10++ {
            for fi11 := 0; fi11 < 4; fi11++ {
                if false {
                    switch "hello" {
                    case "":
                        if p0 {
                            v12 := []int{(fi11 | fi10), 0, (fi10 >> ((p1) & 63)), fi11}
                            switch "" {
                            case "a":
                                if (127 < 75) {
                                    var v13 string = "hello"
                                    switch v13 {
                                    case "a":
                                        v14 := GenS1{F0: 928, F1: -32768, F2: 111}
                                        _ = v14
                                    case "":
                                        var v15 int16 = 1000
                                        _ = v15
                                    default:
                                        v12 = append(v12, genFn0(127, -2147483648, genFn0(1, 46, fi11)))
                                    }
                                    _ = v13
                                } else {
                                    if true {
                                        v16 := GenS2{F0: -65, F1: 65535}
                                        if p0 {
                                            if len(v12) > 0 {
                                                v12[0] = (fi10 & (fi10 & (fi11 ^ fi10)))
                                            }
                                            p0 = (fi11 != -733)
                                        } else {
                                            fmt.Println(p1)
                                            p0 = (18446744073709551615 != p1)
                                            switch "foo bar" {
                                            case "\n\t":
                                                v17 := GenS2{F0: 65, F1: 1}
                                                _ = v17
                                            case "hello":
                                                p1 -= (p1 + p1)
                                            default:
                                                fmt.Println(p1, p1)
                                            }
                                        }
                                        var v18 bool = (255 >= 42)
                                        _ = v16
                                        _ = v18
                                    }
                                    if ((fi11 % (((fi11 & fi10)) | 1)) <= (fi10 + 2147483647)) {
                                        v19 := GenS2{F0: 1, F1: 1}
                                        _ = v19
                                    } else {
                                        v20 := GenS2{F0: 17, F1: 42}
                                        if (48 == 322) {
                                            var v21 string = "foo bar"
                                            p1 = p1
                                            switch "hello" {
                                            case "\n\t":
                                                if len(v12) >= 2 {
                                                    v12 = v12[1:len(v12)]
                                                }
                                            case "foo bar":
                                                var v22 string = ("" + "")
                                                _ = v22
                                            case "a":
                                                p0 = false
                                            default:
                                                _ = len(v12)
                                            }
                                            v23 := GenS3{F0: 9223372036854775807}
                                            if ((p0 || ((113 <= 42) || (p0 && (fi10 != (fi11 % ((genFn0(128, 2147483647, (fi11 / ((fi11) | 1)))) | 1)))))) || (-1 < 42)) {
                                                var v24 bool = p0
                                                _ = v24
                                            } else {
                                                p0 = (2 != 65535)
                                                p1 += (p1 | 42)
                                            }
                                            v25 := GenS4{F0: 1}
                                            fmt.Println(fi11, p0)
                                            _ = v21
                                            _ = v23
                                            _ = v25
                                        }
                                        if (!((165 < 18446744073709551615) && p0)) {
                                            p0 = p0
                                            v26 := GenS5{F0: 550, F1: 13}
                                            if (("hello" <= "") && p0) {
                                                p1 -= (^(p1))
                                                var v27 int = (^(fi10))
                                                p1 -= 343
                                                _ = v27
                                            }
                                            _ = v26
                                        } else {
                                            var v28 float64 = (((0.0 - 0.1) / 1.0) - -2.25)
                                            _ = v28
                                        }
                                        if ((-119 >= 100) || p0) {
                                            if len(v12) > 0 {
                                                v12[0] = (fi11 * (fi10 << (((^(p1))) & 63)))
                                            }
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            if p0 {
                                                var v29 uint16 = 32767
                                                _ = v29
                                            }
                                            if (65535 < 142) {
                                                _ = len(v12)
                                                v30 := GenS5{F0: (p1 * 2), F1: 1}
                                                fmt.Println(p1)
                                                if len(v12) >= 2 {
                                                    v12 = v12[1:len(v12)]
                                                }
                                                _ = v30
                                            } else {
                                                switch "foo bar" {
                                                case "hello":
                                                    if p0 {
                                                        var v31 bool = p0
                                                        p1 = (p1 & 18446744073709551615)
                                                        _ = v31
                                                    }
                                                case "a":
                                                    var v32 float64 = (0.0 - 1.5)
                                                    _ = v32
                                                case "":
                                                    var v33 int8 = -1
                                                    _ = v33
                                                }
                                                switch "\n\t" {
                                                case "a":
                                                    var v34 int16 = 42
                                                    _ = v34
                                                case "hello":
                                                    switch "a" {
                                                    case "foo bar":
                                                        if p0 {
                                                            v35 := GenS6{F0: -32768, F1: 4294967295}
                                                            p0 = (!(0.1 <= float64((float32((1.5 + (-2.25 * (((0.0 * float64((-2.25 * 1e10))) + (1.5 - 1e10)) * 1.5)))) * 0.0))))
                                                            _ = v35
                                                        }
                                                    case "a":
                                                        if true {
                                                            if len(v12) >= 2 {
                                                                v12 = v12[1:len(v12)]
                                                            }
                                                            v36 := GenS6{F0: -1, F1: 933}
                                                            p1 -= 18446744073709551615
                                                            _ = v36
                                                        }
                                                        if len(v12) >= 2 {
                                                            v12 = v12[1:len(v12)]
                                                        }
                                                    case "\n\t":
                                                        fmt.Println(p1, p1, fi11)
                                                    default:
                                                        if len(v12) >= 2 {
                                                            v12 = v12[1:len(v12)]
                                                        }
                                                        if len(v12) >= 2 {
                                                            v12 = v12[1:len(v12)]
                                                        }
                                                    }
                                                }
                                            }
                                            v37 := GenS7{F0: -9223372036854775808, F1: -29}
                                            _ = v37
                                        } else {
                                            fmt.Println(fi11, p1)
                                            v38 := GenS8{F0: 65535}
                                            v39 := GenS8{F0: 2147483648}
                                            _ = v38
                                            _ = v39
                                        }
                                        _ = v20
                                    }
                                    var v40 uint64 = 18446744073709551615
                                    _ = v40
                                }
                            case "foo bar":
                                var v41 int8 = -1
                                _ = v41
                            case "hello":
                                p0 = (1e10 <= 0.1)
                            }
                            _ = v12
                        } else {
                            v42 := []int{(fi11 >> (((p1 % (((^(p1))) | 1))) & 63)), fi10, -619, (fi10 - genFn0(72, 2147483647, genFn0(216, 0, fi11)))}
                            _ = v42
                        }
                        if (p0 && ((p1 << ((p1) & 63)) != 2)) {
                            var v43 int64 = 1
                            var v44 int64 = v43
                            p0 = (true && (1000 <= 1000))
                            _ = v43
                            _ = v44
                        }
                        var v45 int8 = -128
                        if ((float64(1e10) == 3.14159) && p0) {
                            p0 = (p0 && (!p0))
                            var v46 string = "foo bar"
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(p1, p0)
                            case "a":
                                v47 := []int{}
                                _ = v47
                            default:
                                v48 := GenS6{F0: 232, F1: 4294967295}
                                var v49 int32 = 874
                                _ = v48
                                _ = v49
                            }
                            _ = v46
                        }
                        _ = v45
                    case "hello":
                        v50 := GenS9{F0: 0, F1: 32767}
                        _ = v50
                    case "a":
                        switch (("\n\t" + ("hello" + "hello")) + ("hello" + "a")) {
                        case "hello":
                            p1 = (p1 & p1)
                        case "":
                            fmt.Println(fi10, p1, p0)
                            var v51 uint8 = 197
                            v52 := []int{9223372036854775807}
                            var v53 uint16 = 255
                            if len(v52) >= 2 {
                                v52 = v52[1:len(v52)]
                            }
                            var v54 uint = (p1 >> (((p1 & (p1 << ((p1) & 63)))) & 63))
                            fmt.Println(p0, v54, fi11)
                            _ = v51
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                    }
                    switch ("a" + "") {
                    case "":
                        var v55 uint8 = 127
                        _ = v55
                    case "\n\t":
                        fmt.Println(p1, p0)
                    case "hello":
                        p1 -= p1
                    default:
                        v56 := []int{}
                        _ = v56
                    }
                    var v57 bool = (false || true)
                    var v58 int64 = 42
                    v59 := GenS4{F0: 243}
                    var v60 float64 = 0.0
                    _ = v57
                    _ = v58
                    _ = v59
                    _ = v60
                } else {
                    var v61 float64 = ((float64(1e10) / (((0.0 + 0.1) / 1.0) / 1.0)) * (0.1 - 0.0))
                    switch "foo bar" {
                    case "hello":
                        var v62 uint64 = 1
                        _ = v62
                    case "foo bar":
                        if (42 != 97) {
                            var v63 int64 = -9223372036854775808
                            var v64 int16 = -32768
                            if (0 != 230) {
                                if (p0 && (583 != uint64(p1))) {
                                    switch "hello" {
                                    case "":
                                        p1 = p1
                                    case "foo bar":
                                        fmt.Println(p0, v64, v64)
                                    case "hello":
                                        v65 := []int{}
                                        _ = v65
                                    }
                                    p1 = (^(p1))
                                } else {
                                    p0 = false
                                    fmt.Println(p0)
                                    var v66 int32 = 0
                                    _ = v66
                                }
                                v67 := GenS3{F0: v63}
                                _ = v67
                            } else {
                                v68 := []int{fi11, (fi11 ^ (fi10 >> (((p1 << ((947) & 63))) & 63))), genFn0(128, 0, genFn0(1, -226, 360)), fi10}
                                _ = v68
                            }
                            fmt.Println(v64, p1, p1)
                            v69 := []int{}
                            _ = v63
                            _ = v64
                            _ = v69
                        } else {
                            var v70 uint8 = 137
                            if (p0 || ((p0 || true) || p0)) {
                                if p0 {
                                    p1 = 861
                                    if (p0 && (-1 != 0)) {
                                        p1 += p1
                                    }
                                }
                                v70 = v70
                                v71 := []int{(fi11 | fi10), fi11}
                                _ = v71
                            } else {
                                v72 := []int{fi11, 128, 32768, fi11}
                                var v73 uint16 = 380
                                var v74 float32 = (((1e10 + (0.1 * 0.1)) / 1.5) + 0.1)
                                _ = v72
                                _ = v73
                                _ = v74
                            }
                            v70 -= 255
                            v75 := []int{(fi11 - (fi10 << ((496) & 63)))}
                            _ = v70
                            _ = v75
                        }
                    case "\n\t":
                        var v76 uint64 = 0
                        _ = v76
                    default:
                        var v77 uint32 = 1
                        _ = v77
                    }
                    var v78 uint64 = 18446744073709551615
                    _ = v61
                    _ = v78
                }
                p0 = (0 > 18446744073709551615)
                var v79 uint = p1
                _ = fi11
                _ = v79
            }
            _ = fi10
        }
    case "\n\t":
        var v80 float64 = -2.25
        _ = v80
    }
    if (4294967295 > 2) {
        var v81 uint64 = 1
        if false {
            for fi82 := 0; fi82 < 6; fi82++ {
                switch "" {
                case "":
                    var v83 uint8 = 195
                    _ = v83
                case "hello":
                    var v84 int = fi82
                    _ = v84
                case "\n\t":
                    fmt.Println(v81, v81, fi82)
                default:
                    var v85 string = "a"
                    _ = v85
                }
                var v86 int16 = -32768
                var v87 int16 = v86
                _ = fi82
                _ = v86
                _ = v87
            }
            if false {
                if p0 {
                    if p0 {
                        for fi88 := 0; fi88 < 12; fi88++ {
                            v89 := []int{(fi88 - fi88), fi88, (fi88 + genFn0(159, 1, fi88)), genFn0(110, 166, (fi88 * genFn0(128, 2147483647, (-(fi88)))))}
                            for fi90 := 0; fi90 < 3; fi90++ {
                                switch (("foo bar" + (("hello" + "foo bar") + ("a" + ""))) + "foo bar") {
                                case "":
                                    v81 = (^(v81))
                                case "hello":
                                    v89 = append(v89, fi88)
                                default:
                                    p1 += (^(p1))
                                }
                                if true {
                                    if p0 {
                                        p1 += 18446744073709551615
                                        if len(v89) > 0 {
                                            v89[0] = genFn0(148, 42, fi88)
                                        }
                                        v91 := GenS8{F0: 583}
                                        _ = v91
                                    }
                                    v92 := GenS10{F0: 1, F1: 0}
                                    if ((p0 || (65535 != 65535)) && (94 <= 73)) {
                                        var v93 uint32 = 0
                                        var v94 uint64 = (v81 - v81)
                                        fmt.Println(fi90)
                                        var v95 bool = p0
                                        _ = v93
                                        _ = v94
                                        _ = v95
                                    }
                                    _ = v92
                                } else {
                                    v81 -= v81
                                    v89 = append(v89, (fi88 / ((fi88) | 1)))
                                }
                                _ = fi90
                            }
                            _ = fi88
                            _ = v89
                        }
                        for fi96 := 0; fi96 < 9; fi96++ {
                            v81 = v81
                            _ = fi96
                        }
                        fmt.Println(p0, p0, v81)
                    }
                    v97 := GenS4{F0: 174}
                    _ = v97
                }
                switch -1 {
                case -1:
                    var v98 int32 = 653
                    _ = v98
                }
            } else {
                fmt.Println(v81)
                v99 := GenS11{F0: "\n\t", F1: p1}
                _ = v99
            }
            var v100 string = "a"
            _ = v100
        }
        _ = v81
    } else {
        var v101 uint = (p1 % (((p1 << (((^(p1))) & 63))) | 1))
        v102 := GenS12{F0: 9223372036854775807, F1: 1}
        if (1 >= 0) {
            fmt.Println(v101, p0)
            var v103 int16 = -1
            if p0 {
                p1 = (p1 / ((v101) | 1))
                v104 := []int{127, -1, -122, 127}
                switch 4294967295 {
                case -543:
                    if len(v104) >= 2 {
                        v104 = v104[1:len(v104)]
                    }
                }
                _ = v104
            } else {
                fmt.Println(p1, p1)
                v105 := []int{}
                for fi106 := 0; fi106 < 11; fi106++ {
                    if false {
                        switch "foo bar" {
                        case "\n\t":
                            var v107 int16 = (-(v103))
                            _ = v107
                        }
                        v108 := GenS8{F0: 42}
                        _ = len(v105)
                        _ = v108
                    }
                    if len(v105) >= 2 {
                        v105 = v105[1:len(v105)]
                    }
                    _ = fi106
                }
                _ = v105
            }
            _ = v103
        }
        _ = v101
        _ = v102
    }
    v109 := []int{-825, 127, -2147483648}
    _ = p0
    _ = p1
    _ = v109
    return 1
}

func main() {
    var v1 int16 = 1
    var v2 uint16 = 657
    v3 := []string{"\n\t", ""}
    v4 := make(map[int]int)
    var v5 float64 = 1e10
    switch "a" {
    case "foo bar":
        switch "\n\t" {
        case "":
            for fi110 := 0; fi110 < 16; fi110++ {
                for fi111 := 0; fi111 < 9; fi111++ {
                    var v112 uint64 = 0
                    var v113 string = "a"
                    _ = fi111
                    _ = v112
                    _ = v113
                }
                for fi114 := 0; fi114 < 2; fi114++ {
                    fmt.Println(v2, v5)
                    var v115 int64 = -9223372036854775808
                    if true {
                        v116 := GenS3{F0: (v115 + (v115 & v115))}
                        var v117 float32 = (0.0 + 3.14159)
                        _ = v116
                        _ = v117
                    }
                    fmt.Println(v1, fi110, v2)
                    _ = fi114
                    _ = v115
                }
                v118 := GenS10{F0: -2147483648, F1: 816}
                _ = fi110
                _ = v118
            }
        }
    case "":
        if (((("\n\t" + "hello") + "hello") + ("" + "foo bar")) == "a") {
            v5 -= v5
            var v119 float64 = (1.5 * (v5 + (v5 / float64(0.1))))
            v119 -= -2.25
            _ = v119
        }
    case "hello":
        for fi120 := 0; fi120 < 10; fi120++ {
            switch "" {
            case "":
                v121 := GenS5{F0: 0, F1: 0}
                _ = v121
            case "foo bar":
                if (!(4294967295 >= 677)) {
                    var v122 bool = false
                    _ = v122
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        v123 := GenS9{F0: 42, F1: v2}
                        _ = v123
                    case "\n\t":
                        if (0 != 422) {
                            v5 += v5
                            v4[genFn0(uint8(v2), 2147483647, (-(fi120)))] = fi120
                            if true {
                                v124 := GenS4{F0: 1}
                                for fi125 := 0; fi125 < 9; fi125++ {
                                    if len(v3) > 0 {
                                        v3[0] = "\n\t"
                                    }
                                    v4[0] = 32767
                                    if (698 <= 2147483648) {
                                        v2 += v2
                                    } else {
                                        switch (fi120 + fi120) {
                                        case -894:
                                            fmt.Println(fi125)
                                        case 0:
                                            _ = len(v4)
                                        default:
                                            v2 += (v2 ^ (v2 & v2))
                                        }
                                    }
                                    _ = fi125
                                }
                                var v126 uint32 = 4294967295
                                _ = v124
                                _ = v126
                            }
                        }
                    default:
                        fmt.Println(v5)
                    }
                }
            default:
                v5 -= v5
            }
            _ = fi120
        }
    }
    v127 := GenS8{F0: 65535}
    if false {
        v4[genFn0(255, 42, 1)] = 42
        v128 := GenS2{F0: 127, F1: 2}
        v1 -= -424
        _ = v128
    }
    switch "foo bar" {
    case "\n\t":
        _ = len(v4)
    case "foo bar":
        switch "hello" {
        case "hello":
            v2 = v2
        case "":
            v5 += (((v5 / (v5 / v5)) - (0.1 + v5)) + ((0.0 + v5) / (v5 - 0.1)))
        case "\n\t":
            var v129 float64 = v5
            _ = v129
        }
    case "a":
        for fi130 := 0; fi130 < 4; fi130++ {
            var v131 uint8 = 2
            var v132 int = fi130
            _ = fi130
            _ = v131
            _ = v132
        }
    }
    for fi133 := 0; fi133 < 9; fi133++ {
        v1 -= (v1 - -32768)
        if (718 != genFn0(118, -1, -9223372036854775808)) {
            fmt.Println(v5, fi133, fi133)
            v134 := GenS3{F0: -1}
            fmt.Println(v2)
            _ = v134
        }
        var v135 int8 = -128
        _ = fi133
        _ = v135
    }
    v2 -= 1
    var v136 int16 = (v1 % ((v1) | 1))
    fmt.Println(v1, v5, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v136:", v136)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v127
    _ = v136
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
