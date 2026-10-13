package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int16 }
type GenS2 struct { F0 int64; F1 string; F2 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64; F1 int32; F2 int }
type GenS7 struct { F0 int64; F1 bool; F2 uint16 }
func genFn0(p0 uint8, p1 int16) uint {
    v6 := GenS1{F0: 42, F1: -1, F2: (p1 >> ((0) & 63))}
    switch -457 {
    case 2147483647:
        if true {
            var v7 uint64 = 1
            p1 -= (p1 << ((65535) & 63))
            _ = v7
        } else {
            p0 += (p0 << ((2) & 63))
            v8 := GenS1{F0: -362, F1: 9223372036854775807, F2: p1}
            p1 = 842
            _ = v8
        }
    case -449:
        switch ("hello" + "hello") {
        case "hello":
            for fi9 := 0; fi9 < 20; fi9++ {
                p1 = (-(p1))
                v10 := []int{fi9, (fi9 & 2147483647)}
                p1 -= 0
                _ = fi9
                _ = v10
            }
        case "\n\t":
            var v11 uint16 = 65535
            _ = v11
        case "a":
            v12 := []int{9223372036854775807, -422, 1}
            _ = v12
        }
    default:
        v13 := []int{9223372036854775807}
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v6
    return 273
}
func genFn1(p0 uint64) uint {
    var v14 int = 128
    var v15 bool = true
    v16 := GenS1{F0: 42, F1: (v14 ^ (v14 << ((2) & 63))), F2: 42}
    switch v14 {
    case -580:
        v15 = v15
    }
    v14 -= -959
    _ = p0
    _ = v14
    _ = v15
    _ = v16
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "hello"
    v3 := make(map[string]bool)
    var v4 uint32 = 869
    var v5 uint8 = 44
    if false {
        switch v2 {
        case "hello":
            var v17 uint32 = (v4 * v4)
            _ = v17
        case "":
            if true {
                if (255 >= 65535) {
                    v18 := GenS1{F0: -2147483648, F1: -902, F2: -32768}
                    v4 = v4
                    _ = v18
                }
                var v19 uint32 = 585
                _ = v19
            } else {
                switch v2 {
                case "":
                    switch (v2 + v2) {
                    case "a":
                        v20 := []int{}
                        _ = v20
                    case "\n\t":
                        for fi21 := 0; fi21 < 18; fi21++ {
                            if ((9223372036854775807 > -1) && true) {
                                if true {
                                    v5 += 2
                                }
                                if ((11 >= uint64(v5)) && false) {
                                    v4 -= (v4 - v4)
                                    v4 = (v4 % (((v4 / ((v4) | 1))) | 1))
                                    switch v2 {
                                    case "foo bar":
                                        var v22 uint8 = (v5 | (v5 / ((v5) | 1)))
                                        _ = v22
                                    case "\n\t":
                                        fmt.Println(v2, fi21, v5)
                                    default:
                                        if ((592 == -307) && (false && (-21 <= -102))) {
                                            v2 = v2
                                        } else {
                                            var v23 uint64 = 756
                                            var v24 int64 = 9223372036854775807
                                            switch v2 {
                                            case "foo bar":
                                                for fi25 := 0; fi25 < 2; fi25++ {
                                                    var v26 uint = 476
                                                    _ = fi25
                                                    _ = v26
                                                }
                                            default:
                                                if (-898 > fi21) {
                                                    var v27 string = v2
                                                    _ = v27
                                                } else {
                                                    var v28 uint8 = (v5 | 10)
                                                    var v29 uint16 = 2
                                                    v3[v2] = false
                                                    _ = v28
                                                    _ = v29
                                                }
                                            }
                                            _ = v23
                                            _ = v24
                                        }
                                    }
                                } else {
                                    var v30 int32 = 0
                                    _ = v30
                                }
                            }
                            var v31 uint16 = 1
                            for fi32 := 0; fi32 < 5; fi32++ {
                                var v33 uint = 978
                                switch (fi21 % ((fi32) | 1)) {
                                case 1:
                                    if true {
                                        switch (fi32 & fi21) {
                                        case 0:
                                            fmt.Println(v4, v4, fi32)
                                        case 223:
                                            if (42 < -281) {
                                                switch v2 {
                                                case "foo bar":
                                                    var v34 float32 = (-2.25 / 0.1)
                                                    _ = v34
                                                case "hello":
                                                    v31 = 568
                                                    v35 := []int{fi32}
                                                    if len(v35) > 0 {
                                                        v35[0] = fi21
                                                    }
                                                    fmt.Println(v2, v31)
                                                    _ = v35
                                                case "a":
                                                    v36 := GenS1{F0: -2147483648, F1: 4294967295, F2: -1}
                                                    _ = v36
                                                }
                                                v37 := v1[(fi21 ^ (fi32 - fi21))]
                                                if false {
                                                    fmt.Println(v33, v31)
                                                    var v38 uint32 = uint32(v5)
                                                    _ = v38
                                                } else {
                                                    v39 := GenS1{F0: 2147483647, F1: -2147483648, F2: 1000}
                                                    v37 = v4
                                                    var v40 uint32 = 540
                                                    _ = v39
                                                    _ = v40
                                                }
                                                _ = v37
                                            }
                                        default:
                                            v5 = v5
                                        }
                                        fmt.Println(fi21)
                                        v41 := GenS1{F0: -564, F1: (fi32 & (fi21 % (((fi21 / (((fi21 ^ 9223372036854775807)) | 1))) | 1))), F2: 0}
                                        switch 9223372036854775807 {
                                        case -404:
                                            v42 := v3[("\n\t" + v2)]
                                            if (!v42) {
                                                v33 = genFn0((v5 ^ (v5 / ((v5) | 1))), 1)
                                                v5 -= 34
                                            }
                                            switch (fi21 - fi21) {
                                            case 42:
                                                switch v2 {
                                                case "\n\t":
                                                    if ("a" != v2) {
                                                        v33 -= genFn0((v5 * v5), 1000)
                                                    } else {
                                                        v42 = false
                                                    }
                                                case "":
                                                    v43 := GenS1{F0: 0, F1: fi32, F2: -1}
                                                    _ = v43
                                                case "foo bar":
                                                    v4 = v4
                                                default:
                                                    var v44 int = (fi32 & (fi21 / ((1) | 1)))
                                                    _ = v44
                                                }
                                            case 32767:
                                                var v45 int32 = 2147483647
                                                _ = v45
                                            default:
                                                switch (fi21 | -9223372036854775808) {
                                                case 127:
                                                    if (1 == 2) {
                                                        var v46 int32 = -2147483648
                                                        var v47 bool = false
                                                        _ = v46
                                                        _ = v47
                                                    } else {
                                                        if ((^(v33)) >= v33) {
                                                            fmt.Println(v4, v2, fi32)
                                                        }
                                                        v2 = "\n\t"
                                                    }
                                                case 821:
                                                    v48 := GenS1{F0: -2147483648, F1: fi21, F2: -32768}
                                                    _ = v48
                                                case 4294967295:
                                                    if (!(!false)) {
                                                        if (!true) {
                                                            v49 := GenS1{F0: 379, F1: fi32, F2: 42}
                                                            v4 -= v4
                                                            _ = v49
                                                        } else {
                                                            v3[v2] = (!v42)
                                                            v50 := []int{(^(fi32))}
                                                            v51 := GenS2{F0: -1, F1: v2, F2: v42}
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                        var v52 int = int(v33)
                                                        if v42 {
                                                            var v53 uint32 = (v4 | 2147483647)
                                                            v31 -= v31
                                                            _ = v53
                                                        } else {
                                                            v3["\n\t"] = (v42 || v42)
                                                            v54 := v3[v2]
                                                            _ = v54
                                                        }
                                                        _ = v52
                                                    }
                                                }
                                            }
                                            v4 -= v4
                                            v55 := []int{fi21}
                                            var v56 uint32 = (^(v4))
                                            var v57 uint32 = (v56 % ((736) | 1))
                                            switch "hello" {
                                            case "":
                                                var v58 uint64 = 156
                                                _ = v58
                                            case "a":
                                                v57 += 2147483647
                                            default:
                                                v31 = v31
                                            }
                                            var v59 uint = v33
                                            _ = v42
                                            _ = v55
                                            _ = v56
                                            _ = v57
                                            _ = v59
                                        case -596:
                                            v5 += v5
                                        default:
                                            fmt.Println(v2, v31, v5)
                                        }
                                        v60 := GenS2{F0: 1, F1: v2, F2: ((v4 - (^(v4))) <= (v4 * v4))}
                                        var v61 uint16 = v31
                                        v31 = (v31 / ((v31) | 1))
                                        _ = v41
                                        _ = v60
                                        _ = v61
                                    } else {
                                        v62 := GenS3{F0: 9223372036854775807}
                                        if ((^(v5)) <= v5) {
                                            if false {
                                                v63 := []int{fi21, fi21}
                                                _ = v63
                                            }
                                        }
                                        v4 = (v4 - 0)
                                        var v64 uint64 = 0
                                        var v65 int16 = -32768
                                        _ = v62
                                        _ = v64
                                        _ = v65
                                    }
                                case 127:
                                    fmt.Println(v5)
                                default:
                                    var v66 uint32 = (^(v4))
                                    _ = v66
                                }
                                var v67 int8 = -48
                                _ = fi32
                                _ = v33
                                _ = v67
                            }
                            _ = fi21
                            _ = v31
                        }
                    }
                case "\n\t":
                    for fi68 := 0; fi68 < 4; fi68++ {
                        if (("a" + (("" + ("a" + ((v2 + ("hello" + "foo bar")) + ((((("hello" + ((v2 + ("a" + v2)) + v2)) + v2) + v2) + "") + "\n\t")))) + "foo bar")) == v2) {
                            v69 := v1[(fi68 << ((uint(v4)) & 63))]
                            v70 := GenS2{F0: 9223372036854775807, F1: ((v2 + v2) + v2), F2: (408 < 975)}
                            _ = v69
                            _ = v70
                        }
                        _ = fi68
                    }
                default:
                    fmt.Println(v2)
                }
            }
        case "foo bar":
            if false {
                var v71 bool = (44 <= (v5 / (((v5 >> ((genFn1(4294967295)) & 63))) | 1)))
                _ = v71
            } else {
                v72 := GenS3{F0: 4294967295}
                var v73 int64 = 42
                if false {
                    v74 := []int{128, 32768, -1}
                    var v75 int = -82
                    delete(v3, ("a" + v2))
                    _ = v74
                    _ = v75
                } else {
                    var v76 int8 = 42
                    if (!false) {
                        _ = len(v3)
                        v77 := GenS3{F0: 208}
                        v78 := GenS1{F0: -2147483648, F1: -346, F2: 0}
                        _ = v77
                        _ = v78
                    }
                    _ = v76
                }
                _ = v72
                _ = v73
            }
        }
    } else {
        switch 958 {
        case -9223372036854775808:
            var v79 uint64 = 2
            _ = v79
        }
        v80 := []int{int(v5), 101, -9223372036854775808, -2147483648}
        v81 := GenS4{F0: 65535, F1: ("\n\t" + ("\n\t" + v2))}
        _ = v80
        _ = v81
    }
    switch (v2 + "\n\t") {
    case "a":
        fmt.Println(v4)
    default:
        var v82 float32 = float32(1.5)
        _ = v82
    }
    v4 = 4294967295
    v2 = "hello"
    v83 := GenS5{F0: "a"}
    switch v2 {
    case "hello":
        for fi84 := 0; fi84 < 7; fi84++ {
            v85 := []int{(fi84 + fi84), (fi84 % (((^(fi84))) | 1)), fi84, (fi84 | (fi84 * fi84))}
            if false {
                fmt.Println(v2, fi84)
                v86 := GenS6{F0: 9223372036854775807, F1: -96, F2: int(v4)}
                _ = v86
            } else {
                fmt.Println(v5, fi84, v4)
                delete(v1, 0)
            }
            for fi87 := 0; fi87 < 18; fi87++ {
                var v88 int32 = -1
                _ = fi87
                _ = v88
            }
            _ = fi84
            _ = v85
        }
    }
    v4 += (^(v4))
    for fi89 := 0; fi89 < 1; fi89++ {
        var v90 bool = false
        if false {
            switch (fi89 | -221) {
            case 42:
                v91 := GenS5{F0: "foo bar"}
                _ = v91
            default:
                var v92 bool = v90
                _ = v92
            }
        }
        _ = fi89
        _ = v90
    }
    if (954 <= 2) {
        var v93 int = 46
        _ = v93
    } else {
        v94 := []int{127, 823}
        _ = len(v3)
        fmt.Println(v5, v4, v4)
        _ = v94
    }
    if false {
        for fi95 := 0; fi95 < 13; fi95++ {
            for fi96 := 0; fi96 < 14; fi96++ {
                _ = len(v1)
                v97 := []int{}
                _ = fi96
                _ = v97
            }
            v98 := []int{(fi95 & fi95), (fi95 << ((1) & 63)), fi95, (fi95 << ((255) & 63))}
            delete(v3, v2)
            _ = fi95
            _ = v98
        }
        for fi99 := 0; fi99 < 13; fi99++ {
            if (fi99 < fi99) {
                _ = len(v1)
                fmt.Println(v4, fi99, fi99)
                v2 = "\n\t"
            }
            _ = fi99
        }
    }
    v4 += 2
    switch "hello" {
    case "":
        v5 = v5
    case "\n\t":
        v2 = ("foo bar" + "a")
    default:
        var v100 int8 = -1
        _ = v100
    }
    v4 = v4
    if true {
        switch 32768 {
        case 4294967295:
            if (-9223372036854775808 == 9223372036854775807) {
                if (0 <= -32768) {
                    if false {
                        if false {
                            var v101 int = 967
                            var v102 uint = 992
                            _ = v101
                            _ = v102
                        } else {
                            switch 127 {
                            case 957:
                                fmt.Println(v2, v2, v2)
                            case 32768:
                                if (true || false) {
                                    switch "a" {
                                    case "\n\t":
                                        for fi103 := 0; fi103 < 13; fi103++ {
                                            var v104 float32 = 1.5
                                            _ = fi103
                                            _ = v104
                                        }
                                    }
                                    switch 32768 {
                                    case 32767:
                                        v4 += (^(v4))
                                    case -718:
                                        v2 = ""
                                    }
                                } else {
                                    for fi105 := 0; fi105 < 6; fi105++ {
                                        switch 1 {
                                        case 128:
                                            fmt.Println(v2, v4)
                                        case 42:
                                            var v106 int16 = 0
                                            _ = v106
                                        default:
                                            v2 = v2
                                        }
                                        _ = fi105
                                    }
                                }
                            case -99:
                                v5 = v5
                            }
                            v107 := []int{}
                            fmt.Println(v4, v2, v5)
                            _ = v107
                        }
                    }
                    var v108 string = ("" + v2)
                    var v109 bool = true
                    _ = v108
                    _ = v109
                }
                v110 := GenS4{F0: 0, F1: (v2 + "foo bar")}
                _ = v110
            } else {
                for fi111 := 0; fi111 < 17; fi111++ {
                    switch 9223372036854775807 {
                    case 2147483647:
                        fmt.Println(fi111, v4, v5)
                    case 42:
                        v112 := GenS6{F0: -1, F1: 0, F2: fi111}
                        _ = v112
                    default:
                        switch "\n\t" {
                        case "":
                            switch fi111 {
                            case 128:
                                fmt.Println(v5, v2)
                            case 32767:
                                v5 = (v5 >> ((1) & 63))
                            }
                        default:
                            v4 += v4
                        }
                    }
                    _ = fi111
                }
            }
        case -40:
            var v113 uint64 = 4294967295
            _ = v113
        case -734:
            if (float64(float32(0.1)) >= 3.14159) {
                for fi114 := 0; fi114 < 5; fi114++ {
                    for fi115 := 0; fi115 < 17; fi115++ {
                        if false {
                            if (v2 == "") {
                                var v116 uint64 = 950
                                delete(v3, v2)
                                _ = v116
                            }
                        }
                        var v117 int16 = -775
                        _ = fi115
                        _ = v117
                    }
                    var v118 uint64 = 4294967295
                    for fi119 := 0; fi119 < 20; fi119++ {
                        if false {
                            if false {
                                v120 := v1[fi119]
                                fmt.Println(v120, v5, fi114)
                                switch -2147483648 {
                                case -9223372036854775808:
                                    if (!((fi119 + (fi114 | (fi119 * 9223372036854775807))) < fi119)) {
                                        v2 = "\n\t"
                                        v121 := GenS1{F0: -1, F1: fi114, F2: 427}
                                        fmt.Println(v5)
                                        _ = v121
                                    } else {
                                        v122 := GenS4{F0: 32767, F1: v2}
                                        v2 = v2
                                        _ = v122
                                    }
                                default:
                                    var v123 string = "foo bar"
                                    _ = v123
                                }
                                v2 = "a"
                                fmt.Println(v118, v4)
                                _ = v120
                            }
                            v124 := GenS6{F0: -1, F1: 0, F2: -2147483648}
                            v118 = (v118 % ((v118) | 1))
                            _ = v124
                        } else {
                            var v125 uint32 = v4
                            var v126 uint8 = (^(v5))
                            var v127 uint32 = v125
                            var v128 float64 = 0.1
                            _ = v125
                            _ = v126
                            _ = v127
                            _ = v128
                        }
                        if true {
                            v129 := GenS4{F0: 301, F1: "hello"}
                            var v130 uint = 4294967295
                            if false {
                                v118 = (^(v118))
                                var v131 uint8 = (v5 ^ v5)
                                v4 = 2
                                _ = v131
                            } else {
                                var v132 float32 = 0.0
                                _ = v132
                            }
                            v133 := []int{4294967295}
                            v2 = v2
                            _ = len(v1)
                            _ = v129
                            _ = v130
                            _ = v133
                        } else {
                            v134 := GenS1{F0: 42, F1: 4294967295, F2: 42}
                            _ = v134
                        }
                        v135 := GenS7{F0: -1, F1: (genFn0(v5, 979) != 4294967295), F2: uint16(v4)}
                        _ = fi119
                        _ = v135
                    }
                    _ = fi114
                    _ = v118
                }
            }
        default:
            v1[-9223372036854775808] = (^(v4))
        }
    } else {
        fmt.Println(v5, v2)
        if (false || true) {
            v136 := []int{32767, 32767, -2147483648, 32767}
            _ = v136
        }
    }
    v5 += (v5 & (v5 % ((2) | 1)))
    v2 = (v2 + v2)
    var v137 uint = 659
    var v138 float64 = float64(-2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v137:", v137)
    fmt.Println("v138:", v138)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v137
    _ = v138
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v83
}
