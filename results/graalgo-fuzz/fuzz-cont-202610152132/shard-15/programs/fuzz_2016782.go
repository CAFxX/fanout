package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 uint; F2 int64 }
type GenS6 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS7 struct { F0 uint; F1 uint8 }
type GenS8 struct { F0 uint; F1 uint8; F2 int64 }
type GenS9 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 bool; F1 uint16 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint64; F1 int8 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 bool; F1 int64 }
type GenS16 struct { F0 int64; F1 uint64; F2 bool }
type GenS17 struct { F0 int; F1 int64 }
func genFn0(p0 uint, p1 bool, p2 int16) int {
    var v6 float32 = (0.0 - 0.0)
    if (1 == 1) {
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := []int{}
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 int8 = 42
                switch "hello" {
                case "a":
                    var v12 string = ""
                    _ = v12
                case "foo bar":
                    v13 := []int{1}
                    _ = v13
                case "hello":
                    switch "foo bar" {
                    case "foo bar":
                        if p1 {
                            switch (fi7 ^ -335) {
                            case -828:
                                p1 = (float32(-2.25) >= v6)
                            default:
                                fmt.Println(v11)
                                var v14 int64 = -9223372036854775808
                                v15 := GenS1{F0: 42, F1: 4294967295}
                                _ = v14
                                _ = v15
                            }
                        }
                    }
                }
                var v16 uint8 = 42
                _ = fi10
                _ = v11
                _ = v16
            }
            _ = fi7
        }
        if (1.5 >= (-2.25 - 1.5)) {
            var v17 float64 = ((((-2.25 * 0.0) - 1e10) + 1.5) / 1e10)
            fmt.Println(p1)
            _ = v17
        } else {
            var v18 int16 = (p2 + p2)
            _ = v18
        }
    } else {
        v19 := []int{127}
        _ = v19
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        v21 := []int{42}
        for fi22 := 0; fi22 < 8; fi22++ {
            fmt.Println(fi20)
            if true {
                var v23 bool = ((-1 == -104) && p1)
                var v24 float64 = ((3.14159 + 1e10) / (0.1 + 3.14159))
                _ = v23
                _ = v24
            } else {
                switch "\n\t" {
                case "":
                    var v25 uint32 = 1
                    _ = v25
                case "\n\t":
                    fmt.Println(p2)
                default:
                    v6 += (-2.25 - -2.25)
                }
                p0 = (p0 | 1)
            }
            v6 += 0.1
            if p1 {
                var v26 int32 = 0
                _ = v26
            }
            var v27 string = "a"
            _ = fi22
            _ = v27
        }
        var v28 int32 = int32(p2)
        _ = fi20
        _ = v21
        _ = v28
    }
    if true {
        if p1 {
            if ((p0 == p0) || ((!((-1 > -336) || (-575 < 0))) && false)) {
                fmt.Println(p1, p0, p2)
                fmt.Println(p0, v6, p2)
                if (255 != p0) {
                    v29 := []int{-2147483648, 1}
                    var v30 uint16 = 800
                    _ = v29
                    _ = v30
                }
            }
            p2 -= p2
            var v31 int32 = 0
            _ = v31
        } else {
            fmt.Println(v6)
        }
    }
    if p1 {
        p0 -= (p0 | 810)
    } else {
        var v32 float32 = float32(3.14159)
        _ = v32
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -1
}
func genFn1(p0 uint8, p1 bool, p2 int32) int {
    var v33 uint = 18446744073709551615
    v34 := []int{32767, 32768, 4294967295, 830}
    if p1 {
        if (431 > 1) {
            var v35 int8 = -66
            if p1 {
                fmt.Println(v33, p2)
            }
            var v36 int8 = v35
            _ = v35
            _ = v36
        }
        switch "" {
        case "a":
            for fi37 := 0; fi37 < 12; fi37++ {
                switch "hello" {
                case "":
                    switch "a" {
                    case "\n\t":
                        var v38 int16 = 0
                        _ = v38
                    default:
                        var v39 int32 = p2
                        _ = v39
                    }
                case "foo bar":
                    v40 := GenS2{F0: 42}
                    _ = v40
                case "a":
                    v41 := GenS3{F0: 310, F1: 200}
                    _ = v41
                }
                v42 := GenS4{F0: "foo bar"}
                _ = fi37
                _ = v42
            }
        case "\n\t":
            p1 = p1
        }
        for fi43 := 0; fi43 < 15; fi43++ {
            var v44 float32 = ((-2.25 / 1.0) * 0.0)
            _ = fi43
            _ = v44
        }
    } else {
        v45 := GenS4{F0: "a"}
        _ = v45
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    _ = v34
    return genFn0(2, p1, 585)
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 string = ""
    v3 := make(map[uint32]int)
    v4 := make(map[int]bool)
    var v5 float64 = (float64((1.5 / 1e10)) + (0.0 + 1e10))
    v46 := GenS5{F0: 4294967295, F1: 255, F2: 9223372036854775807}
    switch "a" {
    case "hello":
        v2 = (v2 + v2)
    default:
        var v47 int = -662
        _ = v47
    }
    v1 = (v1 - (v1 % ((v1) | 1)))
    v2 = v2
    for fi48 := 0; fi48 < 12; fi48++ {
        var v49 int16 = -32768
        _ = fi48
        _ = v49
    }
    fmt.Println(v5)
    var v50 int16 = -137
    switch -2147483648 {
    case -1:
        var v51 int = 201
        _ = v51
    case 127:
        if (0 <= 2) {
            fmt.Println(v50, v2, v1)
            v4[0] = true
        } else {
            var v52 uint32 = 2147483647
            for fi53 := 0; fi53 < 11; fi53++ {
                if (!(-1 != -2147483648)) {
                    for fi54 := 0; fi54 < 17; fi54++ {
                        v2 = "\n\t"
                        var v55 float64 = v5
                        _ = fi54
                        _ = v55
                    }
                    _ = len(v3)
                    for fi56 := 0; fi56 < 12; fi56++ {
                        var v57 uint64 = 4294967295
                        _ = fi56
                        _ = v57
                    }
                } else {
                    for fi58 := 0; fi58 < 8; fi58++ {
                        v5 += 1e10
                        v59 := GenS6{F0: 42, F1: 0, F2: 2}
                        if true {
                            if (false && false) {
                                if (3.14159 > float32((0.0 + (-2.25 - v5)))) {
                                    switch "a" {
                                    case "a":
                                        if (("hello" + "") >= v2) {
                                            v60 := GenS7{F0: 2, F1: 2}
                                            _ = v60
                                        }
                                    case "\n\t":
                                        var v61 int8 = -1
                                        _ = v61
                                    case "":
                                        var v62 float64 = (-2.25 / v5)
                                        _ = v62
                                    }
                                    _ = len(v4)
                                }
                                var v63 uint8 = 140
                                v64 := GenS8{F0: 2, F1: (v63 >> ((1) & 63)), F2: (-(v1))}
                                _ = v63
                                _ = v64
                            }
                            var v65 int16 = (v50 - v50)
                            _ = v65
                        }
                        _ = fi58
                        _ = v59
                    }
                }
                _ = fi53
            }
            _ = v52
        }
    default:
        v66 := GenS9{F0: -2147483648, F1: 0, F2: 255}
        _ = v66
    }
    if true {
        switch 847 {
        case 4294967295:
            if false {
                v2 = "a"
                for fi67 := 0; fi67 < 10; fi67++ {
                    for fi68 := 0; fi68 < 19; fi68++ {
                        v69 := []int{(fi68 >> ((2) & 63))}
                        switch v2 {
                        case "a":
                            var v70 int = genFn0(0, true, v50)
                            _ = v70
                        }
                        v2 = ((((v2 + (v2 + ((((v2 + "hello") + "") + (v2 + v2)) + ""))) + "") + "hello") + "hello")
                        _ = fi68
                        _ = v69
                    }
                    for fi71 := 0; fi71 < 2; fi71++ {
                        var v72 float32 = 1.5
                        v2 = v2
                        _ = fi71
                        _ = v72
                    }
                    _ = fi67
                }
            }
        case -9223372036854775808:
            if ((true || false) && true) {
                for fi73 := 0; fi73 < 8; fi73++ {
                    v5 += v5
                    _ = fi73
                }
                v74 := []int{128, -498, 42}
                v75 := GenS5{F0: 42, F1: 2, F2: v1}
                _ = v74
                _ = v75
            } else {
                switch 401 {
                case 32767:
                    switch genFn0(485, false, v50) {
                    case 127:
                        v76 := []int{-375, -834, genFn1(128, (!(v2 <= "a")), 1)}
                        _ = v76
                    case 9223372036854775807:
                        v2 = v2
                    case -412:
                        for fi77 := 0; fi77 < 13; fi77++ {
                            delete(v4, 0)
                            v78 := GenS6{F0: 32767, F1: -2147483648, F2: 2}
                            v79 := GenS10{F0: int8(fi77)}
                            _ = fi77
                            _ = v78
                            _ = v79
                        }
                    }
                }
            }
        default:
            switch -1 {
            case 2147483647:
                if true {
                    v80 := []int{}
                    var v81 float32 = (1.5 + (1e10 * float32(3.14159)))
                    _ = v80
                    _ = v81
                } else {
                    if true {
                        v4[genFn1(151, false, 1)] = true
                        fmt.Println(v50, v2)
                        if false {
                            v82 := GenS11{F0: true, F1: 1}
                            v3[65535] = 2147483647
                            _ = v82
                        } else {
                            switch 128 {
                            case 0:
                                fmt.Println(v50, v50)
                            }
                            v83 := GenS12{F0: -128}
                            _ = v83
                        }
                    }
                }
            case 32767:
                var v84 bool = ((-128 >= -1) || true)
                _ = v84
            case 0:
                v2 = (v2 + v2)
            }
        }
        if (true && true) {
            switch (v2 + ("a" + "foo bar")) {
            case "":
                v5 = v5
            case "hello":
                if true {
                    var v85 string = (((v2 + v2) + ("hello" + v2)) + v2)
                    _ = v85
                } else {
                    v86 := []int{genFn1(128, (((!(v2 > v2)) || true) && (2147483647 == 42)), 139), 192, 9223372036854775807}
                    switch -1 {
                    case -443:
                        _ = len(v86)
                    case 253:
                        for fi87 := 0; fi87 < 18; fi87++ {
                            var v88 uint16 = 42
                            switch v2 {
                            case "a":
                                var v89 int16 = -58
                                _ = v89
                            case "\n\t":
                                v88 = 328
                            case "hello":
                                v90 := GenS13{F0: 998, F1: 100}
                                _ = v90
                            }
                            fmt.Println(v5)
                            _ = fi87
                            _ = v88
                        }
                    case 9223372036854775807:
                        fmt.Println(v2, v1, v50)
                    }
                    v5 -= float64((3.14159 - 1.5))
                    _ = v86
                }
            default:
                var v91 int16 = int16(v1)
                _ = v91
            }
        } else {
            for fi92 := 0; fi92 < 12; fi92++ {
                if true {
                    v50 += v50
                    v5 = v5
                } else {
                    v2 = ("a" + "a")
                }
                if (32767 > 0) {
                    for fi93 := 0; fi93 < 20; fi93++ {
                        if false {
                            fmt.Println(v50, v2, fi93)
                            var v94 uint8 = 240
                            switch "\n\t" {
                            case "hello":
                                switch fi92 {
                                case 128:
                                    v94 -= (v94 + v94)
                                case 0:
                                    v5 -= 0.1
                                case 302:
                                    var v95 uint8 = (v94 % (((v94 % ((0) | 1))) | 1))
                                    _ = v95
                                default:
                                    var v96 int16 = v50
                                    _ = v96
                                }
                            case "a":
                                v97 := GenS14{F0: 4294967295}
                                _ = v97
                            case "":
                                var v98 float32 = (((1.5 * 0.1) - float32(0.0)) + 0.1)
                                _ = v98
                            default:
                                if (!(!false)) {
                                    v99 := GenS15{F0: (255 >= 18446744073709551615), F1: -898}
                                    var v100 float64 = float64(0.1)
                                    v101 := GenS9{F0: -1, F1: 2147483648, F2: 42}
                                    _ = v100
                                    _ = v101
                                    _ = v99
                                } else {
                                    var v102 float32 = 1.5
                                    v103 := GenS1{F0: -654, F1: 65535}
                                    _ = v102
                                    _ = v103
                                }
                            }
                            _ = v94
                        }
                        fmt.Println(v2, v2)
                        v104 := GenS11{F0: ((170 <= 42) && true), F1: 785}
                        if (false && true) {
                            v105 := GenS16{F0: v1, F1: 4294967295, F2: ((true && true) && (!true))}
                            var v106 uint32 = 4294967295
                            var v107 float64 = (v5 * float64((3.14159 / float32(v5))))
                            _ = v105
                            _ = v106
                            _ = v107
                        }
                        var v108 uint64 = 42
                        _ = fi93
                        _ = v104
                        _ = v108
                    }
                    fmt.Println(v1, fi92)
                    if (255 == 128) {
                        v2 = ((v2 + (v2 + "")) + v2)
                        v1 = (v1 / ((v1) | 1))
                    }
                }
                v50 = (v50 | v50)
                _ = fi92
            }
            fmt.Println(v2, v50, v5)
        }
        for fi109 := 0; fi109 < 9; fi109++ {
            var v110 int8 = -93
            switch fi109 {
            case 948:
                fmt.Println(fi109, v2, v110)
            case 127:
                var v111 int = (fi109 ^ 0)
                _ = v111
            case 646:
                for fi112 := 0; fi112 < 3; fi112++ {
                    fmt.Println(fi109, v110, v1)
                    v110 = 127
                    _ = fi112
                }
            }
            v50 = v50
            _ = fi109
            _ = v110
        }
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        v113 := GenS6{F0: 0, F1: 2147483647, F2: 128}
        _ = v113
    case 4294967295:
        if (1000 != v50) {
            v5 = v5
            for fi114 := 0; fi114 < 13; fi114++ {
                v115 := GenS6{F0: 255, F1: 0, F2: 2}
                fmt.Println(fi114, fi114, v1)
                v2 = v2
                _ = fi114
                _ = v115
            }
        }
    case 834:
        v116 := GenS17{F0: -1, F1: v1}
        _ = v116
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v50
}
