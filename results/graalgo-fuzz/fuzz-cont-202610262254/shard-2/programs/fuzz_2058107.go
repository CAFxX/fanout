package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 bool; F2 uint }
type GenS3 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int; F1 int16; F2 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string; F1 uint; F2 uint }
type GenS7 struct { F0 int16; F1 uint32 }
type GenS8 struct { F0 int16; F1 uint }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int32; F1 int64 }
type GenS12 struct { F0 uint16; F1 uint64 }
type GenS13 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 uint8 }
func genFn0(p0 int16, p1 uint) uint32 {
    if true {
        v6 := []int{42, -995, 32767, -580}
        var v7 uint = (p1 >> ((p1) & 63))
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: p0}
    switch 9223372036854775807 {
    case -312:
        for fi9 := 0; fi9 < 14; fi9++ {
            switch "a" {
            case "":
                var v10 int16 = -832
                _ = v10
            case "foo bar":
                v11 := []int{}
                _ = v11
            case "a":
                var v12 int8 = 1
                _ = v12
            }
            _ = fi9
        }
    case 32768:
        switch -134 {
        case 958:
            p0 -= (p0 | (^(p0)))
        case -493:
            var v13 int = 42
            _ = v13
        default:
            p0 = p0
        }
    case 2147483647:
        switch -1 {
        case 472:
            var v14 uint8 = 42
            _ = v14
        case 0:
            for fi15 := 0; fi15 < 2; fi15++ {
                for fi16 := 0; fi16 < 2; fi16++ {
                    v17 := []int{}
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 5; fi18++ {
                    switch "\n\t" {
                    case "":
                        v19 := []int{0, (-(fi18)), -809}
                        p1 = p1
                        var v20 int64 = int64(p0)
                        _ = v19
                        _ = v20
                    case "\n\t":
                        var v21 uint = 574
                        _ = v21
                    case "foo bar":
                        var v22 string = ""
                        _ = v22
                    }
                    if true {
                        v23 := []int{fi15, 42, (fi18 | (fi18 - fi15))}
                        var v24 float32 = 1.5
                        _ = v23
                        _ = v24
                    } else {
                        var v25 float32 = 1.5
                        _ = v25
                    }
                    v26 := GenS2{F0: (-(p0)), F1: true, F2: p1}
                    _ = fi18
                    _ = v26
                }
                p1 -= p1
                _ = fi15
            }
        case -2147483648:
            for fi27 := 0; fi27 < 15; fi27++ {
                if true {
                    for fi28 := 0; fi28 < 18; fi28++ {
                        v29 := []int{fi28, (fi27 & fi27), 32767, -785}
                        if false {
                            p1 -= 65535
                            p1 = p1
                        } else {
                            if true {
                                fmt.Println(p0, fi27, fi27)
                                v30 := GenS3{F0: (p1 + p1), F1: 1, F2: 42}
                                if true {
                                    p1 += 65535
                                    p0 -= (-(p0))
                                    p1 += p1
                                } else {
                                    v31 := GenS4{F0: (fi28 & 0), F1: (p0 * int16(fi28)), F2: 847}
                                    p0 = p0
                                    _ = v31
                                }
                                var v32 bool = (42 == uint32(p1))
                                _ = v30
                                _ = v32
                            } else {
                                if len(v29) > 0 {
                                    v29[0] = -863
                                }
                                var v33 uint32 = 287
                                v33 = v33
                                _ = v33
                            }
                        }
                        v34 := GenS3{F0: (p1 - 358), F1: 89, F2: 4294967295}
                        fmt.Println(fi28, p1, fi28)
                        var v35 int16 = p0
                        _ = fi28
                        _ = v29
                        _ = v34
                        _ = v35
                    }
                    if (float64(0.1) <= (3.14159 - (-2.25 / (((-2.25 / 1.5) / (1e10 + 1.5)) - 1e10)))) {
                        p0 = 32767
                    } else {
                        v36 := []int{42}
                        v37 := GenS5{F0: 781}
                        v36 = append(v36, fi27)
                        _ = v36
                        _ = v37
                    }
                } else {
                    v38 := GenS5{F0: 664}
                    p0 = int16(p1)
                    _ = v38
                }
                for fi39 := 0; fi39 < 4; fi39++ {
                    p0 = p0
                    _ = fi39
                }
                for fi40 := 0; fi40 < 7; fi40++ {
                    v41 := []int{(fi40 * (^(fi27)))}
                    _ = fi40
                    _ = v41
                }
                _ = fi27
            }
        default:
            p1 = 857
        }
    }
    v42 := []int{-2147483648, 4294967295, int(p0)}
    _ = p0
    _ = p1
    _ = v42
    _ = v8
    return 540
}
func genFn1(p0 int16) uint {
    fmt.Println(p0)
    v43 := []int{128}
    _ = p0
    _ = v43
    return 0
}

func main() {
    var v1 float32 = float32(3.14159)
    var v2 int16 = 1
    var v3 int64 = 42
    v4 := make(map[string]uint32)
    var v5 uint = 4294967295
    switch ("\n\t" + ("foo bar" + ("" + ""))) {
    case "":
        v5 -= 1
    case "foo bar":
        for fi44 := 0; fi44 < 14; fi44++ {
            if true {
                var v45 uint32 = 65535
                var v46 float64 = 0.0
                var v47 float64 = v46
                _ = v45
                _ = v46
                _ = v47
            } else {
                var v48 uint64 = 4
                if (false && false) {
                    switch ("foo bar" + "hello") {
                    case "\n\t":
                        var v49 float32 = v1
                        _ = v49
                    }
                }
                var v50 int = -9223372036854775808
                _ = v48
                _ = v50
            }
            _ = fi44
        }
    default:
        var v51 uint32 = 1
        _ = v51
    }
    var v52 uint16 = 2
    v5 -= uint(v52)
    switch ("" + ("\n\t" + "")) {
    case "":
        v53 := GenS4{F0: 127, F1: 1, F2: 2}
        _ = v53
    }
    if ("\n\t" > "\n\t") {
        v2 += v2
        if ((v5 % (((v5 + v5)) | 1)) <= (^(v5))) {
            switch 894 {
            case 32767:
                v2 = (v2 ^ v2)
            }
            switch 655 {
            case 9223372036854775807:
                v5 -= v5
            case 0:
                for fi54 := 0; fi54 < 18; fi54++ {
                    if true {
                        if false {
                            if (2147483648 >= 2147483647) {
                                v3 += (^(v3))
                                v55 := GenS6{F0: "foo bar", F1: (v5 + (v5 % (((v5 - genFn1(int16(v52)))) | 1))), F2: genFn1((v2 << ((genFn1((v2 << ((v5) & 63)))) & 63)))}
                                switch 127 {
                                case -9223372036854775808:
                                    v56 := []int{fi54, 32768, (^(fi54)), fi54}
                                    _ = v56
                                }
                                _ = v55
                            } else {
                                var v57 string = "foo bar"
                                switch "foo bar" {
                                case "hello":
                                    v1 = 0.1
                                default:
                                    if false {
                                        v2 += (^(v2))
                                        var v58 int16 = v2
                                        _ = v58
                                    } else {
                                        for fi59 := 0; fi59 < 20; fi59++ {
                                            if false {
                                                v60 := GenS5{F0: 0}
                                                switch (v57 + "\n\t") {
                                                case "hello":
                                                    var v61 bool = false
                                                    _ = v61
                                                case "":
                                                    v62 := GenS7{F0: v2, F1: 285}
                                                    _ = v62
                                                case "a":
                                                    var v63 uint64 = 4294967295
                                                    _ = v63
                                                default:
                                                    var v64 uint = (v5 * v5)
                                                    _ = v64
                                                }
                                                _ = v60
                                            } else {
                                                v1 += -2.25
                                                fmt.Println(v5, v52, v52)
                                            }
                                            var v65 float32 = v1
                                            v2 = 352
                                            _ = fi59
                                            _ = v65
                                        }
                                    }
                                }
                                for fi66 := 0; fi66 < 5; fi66++ {
                                    v52 = (^(v52))
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v67 uint64 = 18446744073709551615
                                        _ = v67
                                    case "":
                                        var v68 int8 = 83
                                        _ = v68
                                    case "a":
                                        var v69 int16 = 838
                                        _ = v69
                                    }
                                    var v70 float64 = (-2.25 - 3.14159)
                                    _ = fi66
                                    _ = v70
                                }
                                _ = v57
                            }
                        }
                        v52 = v52
                    }
                    _ = fi54
                }
            case 128:
                switch ("hello" + "hello") {
                case "a":
                    v52 += v52
                case "foo bar":
                    if (false && false) {
                        for fi71 := 0; fi71 < 5; fi71++ {
                            for fi72 := 0; fi72 < 19; fi72++ {
                                fmt.Println(v2)
                                switch (fi71 & -172) {
                                case 1:
                                    switch "" {
                                    case "a":
                                        v73 := GenS1{F0: (v2 / ((-32768) | 1))}
                                        _ = v73
                                    case "hello":
                                        var v74 float32 = v1
                                        _ = v74
                                    case "\n\t":
                                        fmt.Println(v5, v1, fi71)
                                    default:
                                        var v75 uint32 = 1
                                        _ = v75
                                    }
                                case 42:
                                    fmt.Println(v52, v52)
                                case -1:
                                    if (!(uint64(v52) <= 2)) {
                                        v1 += v1
                                        if false {
                                            if false {
                                                if true {
                                                    v4["a"] = 31
                                                    switch fi72 {
                                                    case -1:
                                                        if ((v5 <= v5) && (v2 < (v2 / (((-(v2))) | 1)))) {
                                                            fmt.Println(v5, fi71)
                                                        }
                                                    case 139:
                                                        if (false && ((v2 + (v2 + -1)) == v2)) {
                                                            fmt.Println(fi72)
                                                            v52 = 929
                                                        }
                                                    default:
                                                        v76 := []int{(fi72 << ((genFn1((^(v2)))) & 63)), (fi71 + 9223372036854775807), (fi71 * fi71), 9223372036854775807}
                                                        _ = v76
                                                    }
                                                    fmt.Println(v3)
                                                } else {
                                                    fmt.Println(v2)
                                                    var v77 float32 = (v1 / -2.25)
                                                    _ = v77
                                                }
                                                v78 := GenS8{F0: 73, F1: (^(v5))}
                                                var v79 uint = (v5 & v5)
                                                v2 -= v2
                                                var v80 uint64 = 2
                                                var v81 int64 = (-(v3))
                                                _ = v78
                                                _ = v79
                                                _ = v80
                                                _ = v81
                                            }
                                            v52 -= v52
                                            var v82 uint64 = 4294967295
                                            if ((^(v3)) != (v3 / ((v3) | 1))) {
                                                v83 := []int{374, fi71, fi71, (fi72 >> ((v5) & 63))}
                                                var v84 uint8 = 170
                                                var v85 float64 = 1e10
                                                _ = len(v4)
                                                var v86 int16 = v2
                                                _ = v83
                                                _ = v84
                                                _ = v85
                                                _ = v86
                                            } else {
                                                var v87 int32 = 42
                                                if true {
                                                    if ((v1 * v1) <= 0.0) {
                                                        v88 := GenS8{F0: (v2 - v2), F1: 612}
                                                        if (v5 <= (v5 / (((v5 + v5)) | 1))) {
                                                            if (!(float32(0.0) != 1.5)) {
                                                                if (1 != 42) {
                                                                    fmt.Println(v3)
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        v89 := GenS9{F0: ((v5 == v5) || (-54 < 0))}
                                                                        _ = v89
                                                                    case "":
                                                                        if false {
                                                                            if false {
                                                                                v90 := GenS6{F0: ("" + "\n\t"), F1: (v5 - 0), F2: v5}
                                                                                var v91 uint8 = 246
                                                                                var v92 int8 = -128
                                                                                _ = v90
                                                                                _ = v91
                                                                                _ = v92
                                                                            }
                                                                            v93 := GenS10{F0: v2}
                                                                            _ = v93
                                                                        }
                                                                    case "\n\t":
                                                                        v94 := []int{(fi72 + fi72)}
                                                                        _ = v94
                                                                    default:
                                                                        v5 = (v5 & (v5 * (v5 | (v5 >> ((genFn1(0)) & 63)))))
                                                                    }
                                                                    v5 = genFn1((v2 - v2))
                                                                    var v95 uint64 = v82
                                                                    v96 := []int{}
                                                                    _ = v95
                                                                    _ = v96
                                                                }
                                                            }
                                                        }
                                                        v82 = v82
                                                        _ = v88
                                                    }
                                                } else {
                                                    v82 += 18446744073709551615
                                                    v97 := GenS4{F0: (fi71 ^ 477), F1: 953, F2: 2147483648}
                                                    v98 := []int{}
                                                    _ = v97
                                                    _ = v98
                                                }
                                                _ = v87
                                            }
                                            _ = v82
                                        }
                                        v99 := []int{fi72}
                                        _ = v99
                                    }
                                default:
                                    if (((v3 & int64(v2)) < (^(v3))) || ((!(!(((132 > v2) || true) && (!false)))) || (!(970 == 339)))) {
                                        if (!false) {
                                            var v100 uint8 = 128
                                            v101 := []int{(fi72 ^ fi72), (fi72 / ((-377) | 1))}
                                            _ = v100
                                            _ = v101
                                        }
                                        fmt.Println(fi71, v5)
                                    } else {
                                        v102 := []int{}
                                        if (42 > 4294967295) {
                                            fmt.Println(v3)
                                            v1 = float32((0.0 - 1.5))
                                        }
                                        v103 := GenS11{F0: -1, F1: (v3 * v3)}
                                        _ = v102
                                        _ = v103
                                    }
                                }
                                _ = fi72
                            }
                            if true {
                                v1 += -2.25
                                v5 = (v5 << (((v5 - 42)) & 63))
                                v2 -= v2
                            }
                            var v104 string = "\n\t"
                            _ = fi71
                            _ = v104
                        }
                        for fi105 := 0; fi105 < 15; fi105++ {
                            v3 = (v3 % (((v3 - 183)) | 1))
                            var v106 uint8 = uint8(v5)
                            fmt.Println(v3, fi105)
                            _ = fi105
                            _ = v106
                        }
                        v107 := GenS12{F0: (v52 + v52), F1: 4294967295}
                        _ = v107
                    }
                }
            }
        } else {
            v1 = 1.5
            if true {
                v2 = (v2 | (v2 * v2))
            } else {
                if (false || true) {
                    var v108 uint64 = 18446744073709551615
                    for fi109 := 0; fi109 < 17; fi109++ {
                        for fi110 := 0; fi110 < 1; fi110++ {
                            v2 = -32768
                            v111 := GenS13{F0: 2147483648, F1: (v108 << ((genFn1((v2 | v2))) & 63)), F2: 2}
                            v112 := GenS10{F0: v2}
                            _ = fi110
                            _ = v111
                            _ = v112
                        }
                        _ = fi109
                    }
                    _ = v108
                } else {
                    v3 -= 9223372036854775807
                    var v113 int32 = 1
                    _ = v113
                }
                v52 += v52
                var v114 int = 42
                _ = v114
            }
        }
    } else {
        switch "\n\t" {
        case "a":
            v4["a"] = 46
        case "\n\t":
            v2 -= v2
        default:
            if ((0.1 < (-2.25 * 1.5)) && false) {
                for fi115 := 0; fi115 < 7; fi115++ {
                    if (0 < 0) {
                        if (93 >= 203) {
                            for fi116 := 0; fi116 < 20; fi116++ {
                                v52 = (v52 - v52)
                                _ = fi116
                            }
                            v1 -= (v1 * 3.14159)
                        } else {
                            var v117 uint32 = 0
                            _ = len(v4)
                            v118 := []int{-9223372036854775808, 42, 4294967295, fi115}
                            _ = v117
                            _ = v118
                        }
                        fmt.Println(v5, v1, v2)
                    } else {
                        if false {
                            v119 := GenS14{F0: 42}
                            var v120 uint8 = uint8(v5)
                            for fi121 := 0; fi121 < 15; fi121++ {
                                v1 = (0.1 * (v1 + v1))
                                var v122 int8 = -128
                                var v123 uint16 = (^(v52))
                                _ = fi121
                                _ = v122
                                _ = v123
                            }
                            _ = v119
                            _ = v120
                        }
                    }
                    for fi124 := 0; fi124 < 14; fi124++ {
                        v125 := GenS7{F0: (-(v2)), F1: 0}
                        v126 := GenS15{F0: ("foo bar" + "")}
                        v1 -= (v1 - 0.0)
                        var v127 uint16 = v52
                        _ = fi124
                        _ = v125
                        _ = v126
                        _ = v127
                    }
                    v128 := GenS8{F0: -32768, F1: (v5 & (^(v5)))}
                    _ = fi115
                    _ = v128
                }
                var v129 int16 = (v2 ^ (v2 / ((v2) | 1)))
                var v130 uint64 = 164
                _ = v129
                _ = v130
            } else {
                for fi131 := 0; fi131 < 18; fi131++ {
                    switch "\n\t" {
                    case "":
                        fmt.Println(v52, v5, v2)
                    case "hello":
                        v1 += (3.14159 / v1)
                    case "a":
                        for fi132 := 0; fi132 < 4; fi132++ {
                            switch (fi132 % ((fi131) | 1)) {
                            case -1:
                                var v133 int64 = v3
                                _ = v133
                            default:
                                fmt.Println(v52)
                            }
                            var v134 int64 = v3
                            var v135 float32 = v1
                            v136 := GenS16{F0: 2}
                            v5 = v5
                            _ = fi132
                            _ = v134
                            _ = v135
                            _ = v136
                        }
                    }
                    v137 := []int{(fi131 | (fi131 + 32768)), fi131}
                    _ = fi131
                    _ = v137
                }
            }
        }
        var v138 uint32 = 42
        v139 := v4["a"]
        _ = v138
        _ = v139
    }
    var v140 string = "foo bar"
    for fi141 := 0; fi141 < 14; fi141++ {
        fmt.Println(v1, v140)
        _ = fi141
    }
    fmt.Println(v140, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v140:", v140)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v140
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
}
