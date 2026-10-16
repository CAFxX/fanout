package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int32; F1 int8; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 uint32; F1 int64 }
type GenS8 struct { F0 int8; F1 int32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int64; F1 bool }
type GenS11 struct { F0 int16; F1 uint32 }
type GenS12 struct { F0 uint32; F1 string }
type GenS13 struct { F0 int8; F1 int16 }
type GenS14 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int16; F1 uint8; F2 bool }
type GenS17 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS18 struct { F0 uint16; F1 int64; F2 int }
type GenS19 struct { F0 int16; F1 bool }
type GenS20 struct { F0 uint; F1 int }
type GenS21 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS22 struct { F0 uint16; F1 int16; F2 int16 }
type GenS23 struct { F0 string }
type GenS24 struct { F0 bool }
type GenS25 struct { F0 int8 }
type GenS26 struct { F0 uint16; F1 uint }
type GenS27 struct { F0 int; F1 uint64; F2 int64 }
type GenS28 struct { F0 bool; F1 int }
type GenS29 struct { F0 uint8; F1 uint }
func genFn0(p0 int64, p1 uint64) uint {
    p1 += p1
    var v7 int32 = 1
    v8 := GenS1{F0: 1, F1: 255}
    v7 = (v7 & -432)
    var v9 uint32 = 65535
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 255
}
func genFn1(p0 int8, p1 string, p2 int32) uint8 {
    if true {
        switch ("\n\t" + p1) {
        case "foo bar":
            var v10 int32 = -490
            _ = v10
        case "hello":
            switch p1 {
            case "a":
                v11 := []int{0, 0, -575, 4294967295}
                _ = v11
            case "\n\t":
                v12 := GenS2{F0: (-(p2)), F1: p0, F2: 791}
                _ = v12
            }
        default:
            for fi13 := 0; fi13 < 11; fi13++ {
                if (42 >= 1) {
                    for fi14 := 0; fi14 < 10; fi14++ {
                        switch ("a" + p1) {
                        case "a":
                            var v15 uint32 = 1
                            _ = v15
                        case "\n\t":
                            p2 -= 350
                        case "foo bar":
                            var v16 uint16 = 42
                            _ = v16
                        }
                        if false {
                            var v17 int8 = p0
                            switch 4294967295 {
                            case 32767:
                                switch fi14 {
                                case 2147483647:
                                    if true {
                                        v17 = 56
                                    } else {
                                        switch "a" {
                                        case "hello":
                                            switch p1 {
                                            case "":
                                                var v18 float64 = float64(float32(((float64(0.0) + 1.5) / 1.0)))
                                                _ = v18
                                            }
                                        default:
                                            v19 := GenS2{F0: p2, F1: 100, F2: 691}
                                            _ = v19
                                        }
                                        p1 = (p1 + p1)
                                        p2 -= p2
                                    }
                                default:
                                    p0 -= v17
                                }
                            }
                            if (!false) {
                                v20 := []int{fi14}
                                switch p1 {
                                case "a":
                                    fmt.Println(p2, v17, p2)
                                case "hello":
                                    if (true && true) {
                                        fmt.Println(p1)
                                        if true {
                                            p2 -= (^(p2))
                                            p2 = -2147483648
                                            var v21 int8 = p0
                                            _ = v21
                                        }
                                        var v22 int64 = 1
                                        _ = v22
                                    }
                                case "\n\t":
                                    if (p0 <= -1) {
                                        if true {
                                            v23 := GenS2{F0: (^(p2)), F1: 42, F2: 2147483647}
                                            _ = v23
                                        }
                                        var v24 uint16 = 255
                                        var v25 uint = 4294967295
                                        fmt.Println(p2, v25, p1)
                                        var v26 float64 = 0.1
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    }
                                default:
                                    v17 = v17
                                }
                                _ = v20
                            } else {
                                switch fi14 {
                                case 128:
                                    if (((float64((-2.25 + 0.1)) + 3.14159) + 0.0) <= 1e10) {
                                        p0 = p0
                                        v27 := []int{-2147483648, 9223372036854775807}
                                        if (false && false) {
                                            if true {
                                                if true {
                                                    _ = len(v27)
                                                    var v28 uint = 4294967295
                                                    _ = v28
                                                } else {
                                                    fmt.Println(v17, p2)
                                                }
                                            } else {
                                                if ((128 >= (^(fi14))) && (-1 >= (p0 >> ((279) & 63)))) {
                                                    var v29 bool = (("" + p1) == p1)
                                                    switch (fi14 ^ fi13) {
                                                    case 965:
                                                        if (!v29) {
                                                            var v30 uint16 = 32767
                                                            if (!(("\n\t" != (("a" + ("hello" + p1)) + (p1 + ("\n\t" + "\n\t")))) || (v29 || (v30 <= (v30 % ((v30) | 1)))))) {
                                                                var v31 uint16 = v30
                                                                p0 += 100
                                                                if v29 {
                                                                    if (v30 > 807) {
                                                                        var v32 uint32 = 1
                                                                        _ = v32
                                                                    }
                                                                    v33 := GenS3{F0: true}
                                                                    _ = v33
                                                                } else {
                                                                    var v34 uint = 674
                                                                    v35 := GenS4{F0: p2}
                                                                    _ = v34
                                                                    _ = v35
                                                                }
                                                                p1 = ""
                                                                _ = v31
                                                            } else {
                                                                switch (((((("a" + (("a" + p1) + p1)) + p1) + "a") + p1) + "hello") + p1) {
                                                                case "hello":
                                                                    if true {
                                                                        var v36 int8 = v17
                                                                        v30 += v30
                                                                        _ = v36
                                                                    } else {
                                                                        fmt.Println(v17)
                                                                        fmt.Println(v30, fi14)
                                                                        v17 = (p0 << ((genFn0(9223372036854775807, 4294967295)) & 63))
                                                                        p0 += v17
                                                                    }
                                                                case "\n\t":
                                                                    p2 -= 145
                                                                    fmt.Println(p2)
                                                                case "a":
                                                                    switch p1 {
                                                                    case "a":
                                                                        fmt.Println(fi14)
                                                                    default:
                                                                        v37 := GenS5{F0: 9223372036854775807, F1: 0}
                                                                        _ = v37
                                                                    }
                                                                default:
                                                                    var v38 float64 = (1e10 - (0.1 - 1e10))
                                                                    _ = v38
                                                                }
                                                            }
                                                            v27 = append(v27, (fi14 / (((fi14 / (((-(fi13))) | 1))) | 1)))
                                                            _ = v30
                                                        } else {
                                                            if ((fi14 | fi13) > 128) {
                                                                v29 = v29
                                                                v39 := GenS6{F0: -862, F1: p1}
                                                                _ = v39
                                                            } else {
                                                                if len(v27) > 0 {
                                                                    v27[0] = fi13
                                                                }
                                                            }
                                                        }
                                                    case 2147483647:
                                                        switch "foo bar" {
                                                        case "hello":
                                                            if len(v27) >= 2 {
                                                                v27 = v27[1:len(v27)]
                                                            }
                                                        case "\n\t":
                                                            v29 = ((v17 >= 1) || ((^(p2)) >= p2))
                                                            var v40 uint8 = 42
                                                            var v41 int64 = -1
                                                            if (v40 >= v40) {
                                                                if len(v27) >= 2 {
                                                                    v27 = v27[1:len(v27)]
                                                                }
                                                            }
                                                            _ = v40
                                                            _ = v41
                                                        case "foo bar":
                                                            p2 -= p2
                                                        }
                                                    }
                                                    fmt.Println(p2, fi14, fi13)
                                                    _ = v29
                                                }
                                            }
                                        } else {
                                            var v42 uint64 = 944
                                            if (fi14 <= fi13) {
                                                fmt.Println(p2)
                                                var v43 int64 = 71
                                                var v44 uint8 = 255
                                                _ = v43
                                                _ = v44
                                            }
                                            if true {
                                                if (v17 < (-(v17))) {
                                                    v45 := GenS7{F0: uint32(v42), F1: -1}
                                                    fmt.Println(v17, v17)
                                                    if len(v27) >= 2 {
                                                        v27 = v27[1:len(v27)]
                                                    }
                                                    _ = v45
                                                }
                                                switch p1 {
                                                case "a":
                                                    var v46 int16 = -482
                                                    _ = v46
                                                case "hello":
                                                    if (542 >= 502) {
                                                        p0 -= 42
                                                        if true {
                                                            fmt.Println(fi13, v17)
                                                            v27 = append(v27, 529)
                                                            v47 := GenS4{F0: (p2 | (p2 & (-(p2))))}
                                                            _ = v47
                                                        }
                                                    }
                                                default:
                                                    v48 := GenS7{F0: 654, F1: 42}
                                                    _ = v48
                                                }
                                            }
                                            v17 -= (^(p0))
                                            if len(v27) > 0 {
                                                v27[0] = (fi14 | (fi13 - fi13))
                                            }
                                            _ = v42
                                        }
                                        _ = v27
                                    } else {
                                        switch (fi14 + int(v17)) {
                                        case 128:
                                            if (32767 >= 255) {
                                                if (fi14 > (fi14 + (fi13 ^ (fi14 / ((0) | 1))))) {
                                                    var v49 int16 = 32767
                                                    if true {
                                                        p0 += (p0 + -128)
                                                    } else {
                                                        v50 := GenS4{F0: p2}
                                                        _ = v50
                                                    }
                                                    p2 += p2
                                                    if false {
                                                        switch p1 {
                                                        case "":
                                                            if (147 > 401) {
                                                                p0 += (v17 & v17)
                                                                switch p1 {
                                                                case "":
                                                                    v51 := []int{fi13, fi14}
                                                                    _ = v51
                                                                case "a":
                                                                    var v52 bool = (p2 >= 1)
                                                                    _ = v52
                                                                default:
                                                                    switch (p1 + "foo bar") {
                                                                    case "\n\t":
                                                                        fmt.Println(fi14, v49, p0)
                                                                    case "hello":
                                                                        v49 += 821
                                                                    default:
                                                                        if ((fi13 << ((1) & 63)) <= 0) {
                                                                            v53 := []int{(fi13 & fi14), fi13, fi14, (^(fi13))}
                                                                            v53 = append(v53, fi13)
                                                                            _ = len(v53)
                                                                            var v54 int32 = p2
                                                                            fmt.Println(p1, fi14, p0)
                                                                            var v55 float32 = 0.1
                                                                            _ = v53
                                                                            _ = v54
                                                                            _ = v55
                                                                        } else {
                                                                            switch "a" {
                                                                            case "hello":
                                                                                if (!(!true)) {
                                                                                    var v56 int8 = (p0 & -128)
                                                                                    p2 += (p2 / (((p2 | (p2 ^ p2))) | 1))
                                                                                    _ = v56
                                                                                } else {
                                                                                    v57 := []int{(fi13 ^ 32768), -316}
                                                                                    p0 += 100
                                                                                    var v58 int32 = 1
                                                                                    _ = v57
                                                                                    _ = v58
                                                                                }
                                                                            case "a":
                                                                                if (!(((p1 + "\n\t") + "") == "hello")) {
                                                                                    var v59 float64 = 1.5
                                                                                    _ = v59
                                                                                }
                                                                            case "foo bar":
                                                                                switch fi13 {
                                                                                case 1:
                                                                                    v60 := []int{fi13, (^(fi13)), fi14}
                                                                                    _ = v60
                                                                                }
                                                                            default:
                                                                                if (!true) {
                                                                                    var v61 uint64 = 4294967295
                                                                                    var v62 uint8 = 127
                                                                                    _ = v61
                                                                                    _ = v62
                                                                                } else {
                                                                                    var v63 uint32 = 2147483647
                                                                                    _ = v63
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                p1 = "a"
                                                            } else {
                                                                var v64 float64 = 3.14159
                                                                _ = v64
                                                            }
                                                        case "hello":
                                                            fmt.Println(fi14, fi13)
                                                        case "a":
                                                            if true {
                                                                p0 = v17
                                                            }
                                                        }
                                                        if ((!true) && false) {
                                                            fmt.Println(p2, fi13)
                                                            if true {
                                                                p0 = v17
                                                            } else {
                                                                fmt.Println(p0, v17, v17)
                                                                v65 := GenS8{F0: v17, F1: p2}
                                                                if (!true) {
                                                                    v66 := GenS9{F0: ((p1 + (((p1 + "foo bar") + (((p1 + "\n\t") + ((p1 + p1) + ((p1 + p1) + "a"))) + p1)) + p1)) + p1)}
                                                                    _ = v66
                                                                }
                                                                _ = v65
                                                            }
                                                            var v67 int8 = (-(v17))
                                                            switch p1 {
                                                            case "a":
                                                                var v68 uint8 = 42
                                                                _ = v68
                                                            case "":
                                                                p2 = (p2 >> ((42) & 63))
                                                            default:
                                                                var v69 uint16 = 1
                                                                _ = v69
                                                            }
                                                            _ = v67
                                                        } else {
                                                            switch (fi13 % (((fi13 >> ((255) & 63))) | 1)) {
                                                            case 30:
                                                                var v70 int64 = -9223372036854775808
                                                                _ = v70
                                                            case -9223372036854775808:
                                                                v71 := []int{}
                                                                _ = v71
                                                            default:
                                                                v72 := []int{int(v49)}
                                                                _ = v72
                                                            }
                                                            v49 += v49
                                                            v73 := GenS8{F0: (v17 | -1), F1: (-(p2))}
                                                            _ = v73
                                                        }
                                                        v74 := []int{(fi13 | fi14), (fi13 | (fi14 * (-(fi14)))), fi14}
                                                        _ = v74
                                                    }
                                                    _ = v49
                                                }
                                                switch p1 {
                                                case "a":
                                                    var v75 int32 = (^(p2))
                                                    _ = v75
                                                }
                                            }
                                        case 32768:
                                            switch ((p1 + "a") + "hello") {
                                            case "hello":
                                                var v76 uint16 = 255
                                                _ = v76
                                            }
                                        case 42:
                                            v77 := GenS6{F0: -9223372036854775808, F1: ((p1 + p1) + ((p1 + ("hello" + (p1 + p1))) + "hello"))}
                                            _ = v77
                                        default:
                                            v78 := GenS10{F0: -209, F1: true}
                                            _ = v78
                                        }
                                        v79 := GenS11{F0: 1, F1: 169}
                                        _ = v79
                                    }
                                    var v80 uint64 = 1
                                    _ = v80
                                case -9223372036854775808:
                                    var v81 uint32 = 911
                                    switch p1 {
                                    case "a":
                                        switch 1 {
                                        case 128:
                                            var v82 uint32 = v81
                                            _ = v82
                                        case -1:
                                            var v83 uint8 = 255
                                            _ = v83
                                        case 4294967295:
                                            var v84 float32 = 1.5
                                            _ = v84
                                        default:
                                            v85 := GenS12{F0: v81, F1: (p1 + "\n\t")}
                                            _ = v85
                                        }
                                        v86 := []int{(fi13 % ((fi14) | 1))}
                                        fmt.Println(p0, p2)
                                        _ = v86
                                    case "foo bar":
                                        fmt.Println(v17, p2, v81)
                                    }
                                    _ = v81
                                case 42:
                                    p0 -= -1
                                default:
                                    fmt.Println(fi13, p2)
                                }
                                if (false && true) {
                                    var v87 int = (fi14 - 42)
                                    _ = v87
                                } else {
                                    v88 := []int{9223372036854775807, int(p2), (fi13 / (((fi14 + (fi14 + fi13))) | 1)), fi14}
                                    v17 = -128
                                    p0 -= v17
                                    var v89 uint8 = 0
                                    var v90 bool = ((!(!false)) || (1 != fi13))
                                    _ = v88
                                    _ = v89
                                    _ = v90
                                }
                                fmt.Println(p0, p2, v17)
                            }
                            if (!(18446744073709551615 <= 4294967295)) {
                                if ("" > "foo bar") {
                                    v91 := []int{fi14, fi14}
                                    _ = v91
                                }
                                var v92 uint16 = 300
                                p0 = p0
                                p2 -= p2
                                _ = v92
                            }
                            v17 = v17
                            _ = v17
                        }
                        fmt.Println(p1, fi13, fi13)
                        var v93 string = p1
                        _ = fi14
                        _ = v93
                    }
                    v94 := []int{fi13, (^(fi13)), (fi13 - 9223372036854775807)}
                    _ = v94
                } else {
                    p1 = ((p1 + p1) + "foo bar")
                    v95 := []int{528}
                    var v96 int64 = -1
                    _ = v95
                    _ = v96
                }
                p2 -= (p2 * 0)
                _ = fi13
            }
        }
        for fi97 := 0; fi97 < 4; fi97++ {
            v98 := []int{-87, (-(fi97)), 127, fi97}
            _ = fi97
            _ = v98
        }
        p2 -= p2
    }
    if (!(float32(1e10) >= -2.25)) {
        p2 = -342
    }
    v99 := []int{42}
    _ = p0
    _ = p1
    _ = p2
    _ = v99
    return 127
}

func main() {
    v1 := []string{"hello"}
    var v2 uint32 = 0
    var v3 int = 2147483647
    var v4 int = v3
    var v5 uint = 4294967295
    v6 := make(map[int]uint32)
    for fi100 := 0; fi100 < 6; fi100++ {
        v1 = append(v1, ("foo bar" + ""))
        fmt.Println(v3, v4, fi100)
        for fi101 := 0; fi101 < 13; fi101++ {
            v4 += fi101
            _ = fi101
        }
        _ = fi100
    }
    if true {
        v1 = append(v1, "\n\t")
        if false {
            if len(v1) > 0 {
                v1[0] = ("" + "")
            }
            var v102 string = "hello"
            var v103 uint = v5
            _ = v102
            _ = v103
        } else {
            switch "" {
            case "\n\t":
                var v104 int64 = -911
                _ = v104
            case "foo bar":
                for fi105 := 0; fi105 < 7; fi105++ {
                    fmt.Println(v2)
                    v2 += (v2 ^ 65535)
                    _ = fi105
                }
            }
            v4 += (^(v4))
            var v106 int32 = -1
            _ = v106
        }
        if false {
            if true {
                v107 := GenS10{F0: -9223372036854775808, F1: (!(!(("\n\t" + "foo bar") >= "")))}
                _ = v107
            } else {
                for fi108 := 0; fi108 < 4; fi108++ {
                    v2 -= 0
                    switch "hello" {
                    case "":
                        if false {
                            v3 -= 128
                        }
                    case "\n\t":
                        if false {
                            fmt.Println(v3, v3, v5)
                            switch v4 {
                            case -2147483648:
                                if (v2 <= (v2 ^ (v2 >> (((v5 << ((v5) & 63))) & 63)))) {
                                    var v109 uint64 = 42
                                    if true {
                                        if (v109 < v109) {
                                            for fi110 := 0; fi110 < 1; fi110++ {
                                                v111 := v6[fi110]
                                                switch "\n\t" {
                                                case "\n\t":
                                                    v5 = (v5 ^ (v5 >> ((v5) & 63)))
                                                }
                                                v4 = 127
                                                var v112 int = fi110
                                                var v113 int = (v112 >> (((v5 + (^(v5)))) & 63))
                                                _ = fi110
                                                _ = v111
                                                _ = v112
                                                _ = v113
                                            }
                                        }
                                        for fi114 := 0; fi114 < 14; fi114++ {
                                            if false {
                                                switch (v4 >> (((v5 >> ((v5) & 63))) & 63)) {
                                                case 9223372036854775807:
                                                    _ = len(v6)
                                                case 128:
                                                    v2 = v2
                                                }
                                            }
                                            var v115 uint64 = (v109 | (v109 & 1))
                                            _ = fi114
                                            _ = v115
                                        }
                                        v116 := GenS13{F0: 105, F1: -1}
                                        _ = v116
                                    } else {
                                        if (100 == -1) {
                                            _ = len(v1)
                                        } else {
                                            switch "a" {
                                            case "a":
                                                v117 := GenS1{F0: 842, F1: 255}
                                                _ = v117
                                            case "foo bar":
                                                v4 += 4294967295
                                            case "":
                                                if (false || true) {
                                                    v118 := GenS14{F0: 586, F1: v2, F2: (v109 + v109)}
                                                    _ = v118
                                                }
                                            }
                                            v4 = (-(fi108))
                                            for fi119 := 0; fi119 < 5; fi119++ {
                                                var v120 uint32 = uint32(v109)
                                                var v121 int16 = -750
                                                var v122 uint8 = genFn1(100, "hello", 0)
                                                _ = fi119
                                                _ = v120
                                                _ = v121
                                                _ = v122
                                            }
                                        }
                                        v123 := GenS15{F0: v109}
                                        var v124 string = ("foo bar" + "a")
                                        _ = v123
                                        _ = v124
                                    }
                                    var v125 float32 = 1.5
                                    _ = v109
                                    _ = v125
                                } else {
                                    v126 := GenS14{F0: 0, F1: v2, F2: 42}
                                    var v127 int16 = -1
                                    _ = v126
                                    _ = v127
                                }
                            case -147:
                                if (true || (!true)) {
                                    v2 = 65535
                                    fmt.Println(v3, v4, fi108)
                                }
                            case 0:
                                if (!true) {
                                    v6[(-(fi108))] = v2
                                    v3 = (^(v3))
                                    switch -403 {
                                    case 32768:
                                        switch "" {
                                        case "\n\t":
                                            v2 -= v2
                                        }
                                    case -54:
                                        for fi128 := 0; fi128 < 7; fi128++ {
                                            fmt.Println(v5, v5, fi108)
                                            v1 = append(v1, (("foo bar" + "foo bar") + "foo bar"))
                                            v1 = append(v1, "hello")
                                            _ = fi128
                                        }
                                    }
                                } else {
                                    v129 := v6[fi108]
                                    var v130 uint8 = genFn1(-1, (("\n\t" + ("" + ("\n\t" + "\n\t"))) + (("\n\t" + "a") + "a")), 1)
                                    var v131 float32 = -2.25
                                    _ = v129
                                    _ = v130
                                    _ = v131
                                }
                            }
                        } else {
                            var v132 float64 = 3.14159
                            _ = v132
                        }
                    default:
                        for fi133 := 0; fi133 < 11; fi133++ {
                            v5 = 18446744073709551615
                            if false {
                                v5 = genFn0(42, 0)
                            } else {
                                switch "\n\t" {
                                case "hello":
                                    if true {
                                        var v134 int16 = 208
                                        fmt.Println(v134)
                                        fmt.Println(v3, v3)
                                        var v135 uint8 = 239
                                        v2 = (v2 + 2)
                                        _ = v134
                                        _ = v135
                                    }
                                case "foo bar":
                                    v2 += 2147483648
                                case "a":
                                    switch "" {
                                    case "\n\t":
                                        switch v4 {
                                        case 128:
                                            v5 += v5
                                        case 1:
                                            v4 = (v4 & (fi108 | (-(v3))))
                                        case 0:
                                            v136 := GenS16{F0: 354, F1: 1, F2: ((1.5 >= (1e10 * ((((float64(1e10) * -2.25) / 0.1) + 3.14159) / 0.1))) || true)}
                                            _ = v136
                                        default:
                                            var v137 float32 = -2.25
                                            _ = v137
                                        }
                                    }
                                default:
                                    var v138 int16 = 652
                                    _ = v138
                                }
                                if ((!(255 < 0)) && (-1 < 32767)) {
                                    if (42 == 126) {
                                        v139 := GenS2{F0: 0, F1: -128, F2: v2}
                                        fmt.Println(fi133, v2)
                                        v5 = genFn0(-370, 42)
                                        _ = v139
                                    } else {
                                        switch (fi133 & v4) {
                                        case 9223372036854775807:
                                            v140 := GenS17{F0: (v2 << ((v5) & 63)), F1: (false || ((v5 | (v5 << ((v5) & 63))) <= 235)), F2: 1}
                                            _ = v140
                                        case -1:
                                            var v141 bool = false
                                            _ = v141
                                        }
                                        v142 := GenS17{F0: (^(v2)), F1: (!true), F2: 1}
                                        switch 466 {
                                        case -9223372036854775808:
                                            var v143 uint16 = 1
                                            _ = v143
                                        case 96:
                                            v5 -= (v5 % (((v5 << ((v5) & 63))) | 1))
                                        case -1:
                                            if true {
                                                var v144 int8 = -5
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v144
                                            }
                                            v5 += genFn0(0, 18446744073709551615)
                                            _ = len(v6)
                                        }
                                        _ = v142
                                    }
                                    v4 = (v3 ^ 289)
                                    if len(v1) > 0 {
                                        v1[0] = ("\n\t" + "\n\t")
                                    }
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch -1 {
                                    case 4294967295:
                                        v145 := GenS1{F0: 42, F1: 42}
                                        _ = v145
                                    case 32768:
                                        var v146 int8 = 1
                                        _ = v146
                                    case -1:
                                        switch (fi133 / (((v3 & (v3 * 9223372036854775807))) | 1)) {
                                        case 9223372036854775807:
                                            var v147 uint8 = genFn1(1, ((("" + "") + ("foo bar" + (("\n\t" + ("foo bar" + "foo bar")) + ("" + "hello")))) + ""), -1)
                                            _ = v147
                                        case -2147483648:
                                            fmt.Println(v5, fi133)
                                        case -1:
                                            switch ("\n\t" + "a") {
                                            case "\n\t":
                                                if false {
                                                    switch "a" {
                                                    case "":
                                                        switch "a" {
                                                        case "hello":
                                                            v2 = (v2 | v2)
                                                        }
                                                    default:
                                                        v148 := GenS16{F0: 42, F1: 128, F2: false}
                                                        _ = v148
                                                    }
                                                    var v149 int16 = -32768
                                                    if (881 != 0) {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                    v4 += fi133
                                                    _ = v149
                                                } else {
                                                    var v150 bool = (!(!true))
                                                    var v151 uint64 = 201
                                                    if (!v150) {
                                                        if (true && (v150 || (42 > 370))) {
                                                            var v152 string = ""
                                                            if v150 {
                                                                if false {
                                                                    var v153 int16 = 0
                                                                    fmt.Println(v4, v3, v152)
                                                                    _ = v153
                                                                } else {
                                                                    v152 = v152
                                                                    var v154 int8 = -54
                                                                    if v150 {
                                                                        v155 := GenS18{F0: 65535, F1: -439, F2: 2147483647}
                                                                        _ = v155
                                                                    } else {
                                                                        if v150 {
                                                                            var v156 uint32 = v2
                                                                            _ = v156
                                                                        } else {
                                                                            if ((!v150) || (!false)) {
                                                                                var v157 uint8 = 1
                                                                                if (!true) {
                                                                                    var v158 uint = (v5 ^ uint(v2))
                                                                                    fmt.Println(v2, v152, fi133)
                                                                                    _ = v158
                                                                                } else {
                                                                                    fmt.Println(v152)
                                                                                }
                                                                                v159 := GenS19{F0: -32768, F1: (v150 || (true || v150))}
                                                                                _ = v157
                                                                                _ = v159
                                                                            }
                                                                            var v160 float64 = float64(((3.14159 - 1e10) + -2.25))
                                                                            _ = v160
                                                                        }
                                                                        var v161 uint = genFn0(-899, (v151 / ((18446744073709551615) | 1)))
                                                                        var v162 int16 = 1000
                                                                        if len(v1) > 0 {
                                                                            v1[0] = ((v152 + "") + v152)
                                                                        }
                                                                        var v163 uint = v5
                                                                        var v164 float32 = (0.1 - 3.14159)
                                                                        _ = v161
                                                                        _ = v162
                                                                        _ = v163
                                                                        _ = v164
                                                                    }
                                                                    var v165 int8 = 100
                                                                    _ = v154
                                                                    _ = v165
                                                                }
                                                                if v150 {
                                                                    v166 := v6[42]
                                                                    fmt.Println(v3, v152)
                                                                    v152 = v152
                                                                    v167 := GenS20{F0: v5, F1: v4}
                                                                    _ = v166
                                                                    _ = v167
                                                                }
                                                                if v150 {
                                                                    if (false || ((!(v150 && v150)) && ((fi108 & fi133) < 98))) {
                                                                        switch 4294967295 {
                                                                        case 209:
                                                                            var v168 uint16 = 2
                                                                            _ = v168
                                                                        default:
                                                                            if v150 {
                                                                                var v169 uint32 = 42
                                                                                _ = v169
                                                                            }
                                                                        }
                                                                        if (v4 != fi133) {
                                                                            delete(v6, 128)
                                                                        } else {
                                                                            v3 += 127
                                                                            var v170 uint8 = 232
                                                                            _ = v170
                                                                        }
                                                                        v171 := GenS14{F0: 0, F1: v2, F2: v151}
                                                                        _ = v171
                                                                    }
                                                                    var v172 int32 = 54
                                                                    _ = v172
                                                                }
                                                            }
                                                            fmt.Println(v5, v5, fi133)
                                                            var v173 uint64 = v151
                                                            if v150 {
                                                                v151 = (v151 << ((815) & 63))
                                                                _ = len(v1)
                                                            } else {
                                                                v174 := GenS18{F0: 2, F1: 9223372036854775807, F2: (v4 + v3)}
                                                                _ = v174
                                                            }
                                                            if ((v2 - 1) < (v2 * v2)) {
                                                                switch "hello" {
                                                                case "hello":
                                                                    v4 -= 127
                                                                case "foo bar":
                                                                    switch int(v5) {
                                                                    case -2147483648:
                                                                        _ = len(v6)
                                                                        v175 := GenS21{F0: (v173 ^ v151), F1: v151, F2: 0}
                                                                        fmt.Println(v4)
                                                                        _ = v175
                                                                    default:
                                                                        var v176 uint32 = 2147483647
                                                                        _ = v176
                                                                    }
                                                                }
                                                                var v177 int8 = 100
                                                                fmt.Println(v5)
                                                                _ = v177
                                                            } else {
                                                                switch fi108 {
                                                                case 4294967295:
                                                                    v4 += fi108
                                                                    v178 := GenS1{F0: 32767, F1: 129}
                                                                    _ = v178
                                                                }
                                                            }
                                                            if v150 {
                                                                switch -2147483648 {
                                                                case 4294967295:
                                                                    v179 := v6[fi108]
                                                                    _ = v179
                                                                }
                                                                v180 := GenS20{F0: v5, F1: fi133}
                                                                var v181 string = v152
                                                                _ = v180
                                                                _ = v181
                                                            }
                                                            _ = v152
                                                            _ = v173
                                                        }
                                                        var v182 int = v3
                                                        _ = len(v1)
                                                        _ = v182
                                                    }
                                                    _ = v150
                                                    _ = v151
                                                }
                                            case "a":
                                                v2 = v2
                                            case "":
                                                if (!((2 > 2) && (0 > 59))) {
                                                    if len(v1) > 0 {
                                                        v1[0] = "\n\t"
                                                    }
                                                    v2 = (v2 | (v2 ^ v2))
                                                }
                                            default:
                                                if (0 >= int(v5)) {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                } else {
                                                    if (((1.5 - float64(-2.25)) / 1.5) <= float64(3.14159)) {
                                                        v183 := GenS5{F0: 9223372036854775807, F1: 476}
                                                        var v184 uint64 = 213
                                                        _ = v183
                                                        _ = v184
                                                    } else {
                                                        if false {
                                                            if (false && false) {
                                                                if len(v1) > 0 {
                                                                    v1[0] = "\n\t"
                                                                }
                                                                var v185 uint8 = 1
                                                                var v186 uint64 = 4294967295
                                                                _ = v185
                                                                _ = v186
                                                            }
                                                        } else {
                                                            if false {
                                                                var v187 uint = genFn0(721, 42)
                                                                v188 := GenS1{F0: 32767, F1: 128}
                                                                _ = v187
                                                                _ = v188
                                                            } else {
                                                                if (2 != 114) {
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        var v189 uint = genFn0(1, 297)
                                                                        _ = v189
                                                                    case "a":
                                                                        v190 := v6[128]
                                                                        _ = v190
                                                                    case "\n\t":
                                                                        v191 := GenS6{F0: 42, F1: ("" + "foo bar")}
                                                                        _ = v191
                                                                    default:
                                                                        if false {
                                                                            _ = len(v6)
                                                                            var v192 int16 = 0
                                                                            _ = v192
                                                                        } else {
                                                                            var v193 string = ""
                                                                            _ = v193
                                                                        }
                                                                    }
                                                                    if false {
                                                                        var v194 uint64 = 18446744073709551615
                                                                        v194 = (^(v194))
                                                                        var v195 int16 = 562
                                                                        _ = v194
                                                                        _ = v195
                                                                    } else {
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            if len(v1) > 0 {
                                                                                v1[0] = ""
                                                                            }
                                                                        case "":
                                                                            v4 -= (v3 << ((2) & 63))
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                delete(v6, fi133)
                                                                            case "a":
                                                                                var v196 float64 = (float64(1e10) * ((((((-2.25 / 0.1) + (((-2.25 * 0.0) + (3.14159 + 0.1)) - 0.0)) / 1e10) * 1e10) / 3.14159) / 3.14159))
                                                                                _ = v196
                                                                            case "foo bar":
                                                                                v197 := GenS22{F0: 0, F1: -648, F2: 654}
                                                                                _ = v197
                                                                            }
                                                                        }
                                                                        v1 = append(v1, ("foo bar" + "a"))
                                                                        v3 = 816
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        if (!false) {
                                                                            v2 -= (v2 + 42)
                                                                            if ((3.14159 * 0.1) != ((0.0 + ((1e10 * 3.14159) + (3.14159 + ((0.0 * (1e10 - 3.14159)) / 3.14159)))) - (0.0 / 0.1))) {
                                                                                var v198 uint16 = 65535
                                                                                fmt.Println(v198, v198, fi133)
                                                                                v4 = v4
                                                                                _ = v198
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                v5 = (v5 & genFn0(0, 434))
                                                            }
                                                            var v199 int = fi133
                                                            _ = v199
                                                        }
                                                        if (!false) {
                                                            v200 := GenS20{F0: (v5 | genFn0(1, 77)), F1: fi108}
                                                            _ = len(v1)
                                                            _ = v200
                                                        }
                                                        v201 := GenS9{F0: ""}
                                                        switch fi108 {
                                                        case 628:
                                                            v4 = v3
                                                        case 0:
                                                            v202 := GenS15{F0: 0}
                                                            _ = v202
                                                        case -9223372036854775808:
                                                            fmt.Println(v3)
                                                            fmt.Println(v5)
                                                            var v203 uint64 = 18446744073709551615
                                                            _ = v203
                                                        default:
                                                            v3 = 1
                                                        }
                                                        _ = v201
                                                    }
                                                    fmt.Println(fi108, v4, v4)
                                                }
                                            }
                                        }
                                    }
                                    v204 := GenS5{F0: -1, F1: 0}
                                    _ = v204
                                }
                            }
                            _ = fi133
                        }
                    }
                    if false {
                        v205 := GenS23{F0: ("foo bar" + "a")}
                        _ = v205
                    } else {
                        v2 -= 1
                        v5 -= v5
                        if ((!false) && false) {
                            var v206 uint64 = 2
                            _ = v206
                        }
                    }
                    _ = fi108
                }
            }
            switch 4294967295 {
            case -523:
                _ = len(v6)
            case 990:
                if true {
                    for fi207 := 0; fi207 < 19; fi207++ {
                        var v208 bool = true
                        v2 -= v2
                        _ = fi207
                        _ = v208
                    }
                    v209 := GenS24{F0: false}
                    _ = v209
                }
            }
            v2 = v2
        } else {
            _ = len(v6)
            switch v3 {
            case 32768:
                v210 := GenS21{F0: 2, F1: 4294967295, F2: -1}
                _ = v210
            case -494:
                var v211 uint16 = 32767
                _ = v211
            default:
                if (((!(false || (1 <= 4294967295))) && false) || (2 < 581)) {
                    var v212 int64 = 247
                    _ = v212
                } else {
                    var v213 uint8 = genFn1(100, (("" + "") + ""), -2147483648)
                    _ = v213
                }
            }
        }
    }
    for fi214 := 0; fi214 < 6; fi214++ {
        v215 := v6[fi214]
        if (!true) {
            v216 := GenS25{F0: 118}
            v2 = 153
            switch (("\n\t" + "hello") + (("a" + "hello") + "\n\t")) {
            case "\n\t":
                if true {
                    var v217 float64 = float64(0.0)
                    v218 := GenS26{F0: 42, F1: (^(v5))}
                    v219 := GenS26{F0: 0, F1: 2}
                    _ = v217
                    _ = v218
                    _ = v219
                } else {
                    _ = len(v6)
                    fmt.Println(v2, fi214)
                    v215 = (v215 ^ v2)
                }
            case "foo bar":
                for fi220 := 0; fi220 < 20; fi220++ {
                    v221 := GenS7{F0: v215, F1: -1}
                    _ = fi220
                    _ = v221
                }
            }
            _ = v216
        }
        _ = fi214
        _ = v215
    }
    switch v4 {
    case -280:
        v4 = v4
    case -1:
        switch (("foo bar" + "foo bar") + "a") {
        case "":
            fmt.Println(v2, v4)
        case "hello":
            if ((767 > 42) && true) {
                v2 = v2
                _ = len(v6)
                fmt.Println(v2, v5)
            } else {
                var v222 float64 = float64(1.5)
                _ = v222
            }
        case "foo bar":
            v2 -= v2
        default:
            for fi223 := 0; fi223 < 16; fi223++ {
                v5 = 0
                for fi224 := 0; fi224 < 7; fi224++ {
                    var v225 int16 = 0
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                    _ = fi224
                    _ = v225
                }
                v226 := GenS27{F0: fi223, F1: 703, F2: -1}
                _ = fi223
                _ = v226
            }
        }
    case 128:
        for fi227 := 0; fi227 < 1; fi227++ {
            switch "" {
            case "foo bar":
                v6[v3] = (v2 | 1)
            case "":
                v5 -= (^(v5))
            case "\n\t":
                v228 := GenS25{F0: -128}
                _ = v228
            default:
                v3 += v4
            }
            if (255 <= uint8(v5)) {
                if false {
                    fmt.Println(v3, v4)
                } else {
                    v4 = (v3 % (((fi227 + v4)) | 1))
                    for fi229 := 0; fi229 < 19; fi229++ {
                        if (!true) {
                            v230 := GenS5{F0: 0, F1: 1}
                            fmt.Println(v2, v2)
                            v1 = append(v1, "")
                            var v231 float32 = ((3.14159 + (3.14159 / (1.5 + 1e10))) / (((1e10 - 0.1) / ((1.5 / 0.1) / 3.14159)) - float32(3.14159)))
                            _ = v230
                            _ = v231
                        }
                        _ = fi229
                    }
                    var v232 uint8 = 0
                    _ = v232
                }
                v233 := GenS7{F0: (v2 - (v2 << ((v5) & 63))), F1: -306}
                fmt.Println(v3, fi227, v2)
                _ = v233
            } else {
                for fi234 := 0; fi234 < 12; fi234++ {
                    v235 := GenS9{F0: "a"}
                    fmt.Println(v4)
                    _ = fi234
                    _ = v235
                }
            }
            _ = fi227
        }
    default:
        for fi236 := 0; fi236 < 2; fi236++ {
            var v237 uint8 = 127
            _ = fi236
            _ = v237
        }
    }
    _ = len(v6)
    _ = len(v1)
    for fi238 := 0; fi238 < 15; fi238++ {
        v2 = (v2 * 2)
        _ = fi238
    }
    v3 = (v4 + (v4 * v4))
    var v239 float32 = 1.5
    v240 := GenS28{F0: true, F1: v4}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v241 int64 = 0
    v242 := GenS28{F0: (!false), F1: 128}
    v243 := GenS29{F0: 42, F1: (v5 + (v5 & genFn0(v241, 984)))}
    _ = len(v1)
    for fi244 := 0; fi244 < 10; fi244++ {
        v3 = v3
        _ = fi244
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v239:", v239)
    fmt.Println("v241:", v241)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v239
    _ = v240
    _ = v241
    _ = v242
    _ = v243
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
