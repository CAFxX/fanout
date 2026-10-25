package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int; F1 int32; F2 uint }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint64; F1 int32; F2 string }
type GenS9 struct { F0 int64; F1 int32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int; F1 uint32 }
type GenS13 struct { F0 uint64; F1 int32 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int64; F1 uint16; F2 int16 }
type GenS16 struct { F0 uint16; F1 uint32 }
type GenS17 struct { F0 int32 }
type GenS18 struct { F0 string; F1 int16; F2 int32 }
type GenS19 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS20 struct { F0 int16; F1 int8 }
type GenS21 struct { F0 uint64 }
type GenS22 struct { F0 uint }
type GenS23 struct { F0 int; F1 int64; F2 uint64 }
type GenS24 struct { F0 uint8 }
type GenS25 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS26 struct { F0 int32 }
type GenS27 struct { F0 int8; F1 int; F2 bool }
type GenS28 struct { F0 uint16; F1 int16; F2 int16 }
type GenS29 struct { F0 uint32 }
type GenS30 struct { F0 int8; F1 bool; F2 uint8 }
type GenS31 struct { F0 int; F1 uint8; F2 int16 }
type GenS32 struct { F0 int16; F1 uint8; F2 int16 }
type GenS33 struct { F0 int; F1 int32; F2 int }
type GenS34 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS35 struct { F0 int16; F1 uint16; F2 int8 }
type GenS36 struct { F0 uint16; F1 uint32 }
type GenS37 struct { F0 uint32; F1 int }
type GenS38 struct { F0 uint32; F1 uint16 }
type GenS39 struct { F0 bool; F1 string }
type GenS40 struct { F0 bool }
type GenS41 struct { F0 uint64 }
type GenS42 struct { F0 uint16; F1 uint }
type GenS43 struct { F0 int64 }
type GenS44 struct { F0 uint16 }
type GenS45 struct { F0 int8; F1 uint32; F2 int32 }
type GenS46 struct { F0 uint8; F1 int16; F2 uint32 }
func genFn0(p0 int8, p1 bool, p2 int32) uint32 {
    if p1 {
        var v7 int64 = 615
        p0 -= (p0 >> ((369) & 63))
        fmt.Println(v7)
        _ = v7
    }
    if true {
        if false {
            p2 += (p2 & int32(p0))
            p2 = p2
            v8 := GenS1{F0: 1}
            _ = v8
        } else {
            var v9 float64 = (0.1 / (((-2.25 / 3.14159) - 3.14159) / 1.0))
            _ = v9
        }
        if p1 {
            v10 := GenS1{F0: 0}
            _ = v10
        }
        fmt.Println(p2, p0, p1)
    } else {
        if false {
            p0 -= p0
            if p1 {
                v11 := []int{2147483647, 1, 4294967295}
                _ = v11
            }
        } else {
            v12 := []int{}
            var v13 int32 = p2
            _ = v12
            _ = v13
        }
        var v14 uint8 = 227
        for fi15 := 0; fi15 < 6; fi15++ {
            if (("foo bar" + (("\n\t" + (("\n\t" + "\n\t") + "")) + "")) <= "\n\t") {
                var v16 uint32 = 4294967295
                _ = v16
            } else {
                var v17 float32 = -2.25
                fmt.Println(v17, p2, fi15)
                _ = v17
            }
            _ = fi15
        }
        _ = v14
    }
    p1 = false
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 uint8, p1 uint) uint16 {
    p0 = 0
    v18 := []int{32768}
    for fi19 := 0; fi19 < 7; fi19++ {
        for fi20 := 0; fi20 < 2; fi20++ {
            fmt.Println(p0, p0, fi19)
            v18 = append(v18, (fi20 | (-(fi20))))
            var v21 int64 = 0
            v22 := GenS2{F0: -534, F1: false}
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    v23 := GenS3{F0: 759}
    _ = p0
    _ = p1
    _ = v18
    _ = v23
    return uint16(p0)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 854
    var v3 uint16 = 0
    v4 := make(map[uint32]bool)
    var v5 bool = false
    v6 := make(map[string]bool)
    if false {
        var v24 int = -2147483648
        v25 := GenS3{F0: 65535}
        if (!(!(1 <= 160))) {
            var v26 int16 = 71
            _ = v26
        }
        _ = v24
        _ = v25
    } else {
        for fi27 := 0; fi27 < 8; fi27++ {
            v28 := []int{(fi27 ^ fi27), (fi27 % ((-206) | 1)), (-(fi27)), 9223372036854775807}
            for fi29 := 0; fi29 < 17; fi29++ {
                v30 := v4[(v2 << ((18446744073709551615) & 63))]
                _ = fi29
                _ = v30
            }
            _ = fi27
            _ = v28
        }
        v31 := GenS4{F0: ((v3 >> ((611) & 63)) <= (^(v3))), F1: -702}
        _ = v31
    }
    v32 := GenS2{F0: 9223372036854775807, F1: false}
    v33 := GenS5{F0: 2147483647, F1: 0, F2: 61}
    v3 -= v3
    fmt.Println(v5, v2, v2)
    var v34 uint = 1
    if v5 {
        fmt.Println(v2, v5, v3)
        switch "" {
        case "hello":
            v35 := GenS5{F0: -9223372036854775808, F1: 42, F2: v34}
            _ = v35
        case "\n\t":
            v36 := GenS6{F0: v3, F1: ("a" + ("foo bar" + ""))}
            _ = v36
        case "":
            for fi37 := 0; fi37 < 2; fi37++ {
                switch int(v3) {
                case -1:
                    switch "foo bar" {
                    case "foo bar":
                        if (!false) {
                            v38 := GenS2{F0: -1, F1: (!false)}
                            _ = v38
                        } else {
                            v39 := GenS3{F0: (v3 | v3)}
                            _ = v39
                        }
                    case "":
                        var v40 int = fi37
                        _ = v40
                    }
                default:
                    v34 += (v34 % (((v34 >> ((490) & 63))) | 1))
                }
                v6[("foo bar" + "foo bar")] = v5
                _ = fi37
            }
        default:
            fmt.Println(v3)
        }
        fmt.Println(v3, v3, v3)
    } else {
        if (1.5 > 1e10) {
            var v41 int32 = 1
            var v42 string = ("\n\t" + "\n\t")
            v42 = "a"
            _ = v41
            _ = v42
        }
    }
    if true {
        fmt.Println(v2)
        if v5 {
            _ = len(v4)
            if (264 != 42) {
                var v43 float64 = 1e10
                _ = v43
            } else {
                delete(v4, v2)
                var v44 string = "foo bar"
                fmt.Println(v2, v34, v44)
                _ = v44
            }
        } else {
            v5 = (((^(v2)) > (v2 ^ 2)) && true)
            v45 := GenS6{F0: (v3 & (v3 - (v3 << (((^(v34))) & 63)))), F1: ("a" + "\n\t")}
            var v46 uint32 = 2147483648
            _ = v45
            _ = v46
        }
        for fi47 := 0; fi47 < 5; fi47++ {
            var v48 uint64 = 131
            for fi49 := 0; fi49 < 19; fi49++ {
                if false {
                    var v50 int64 = -1
                    v51 := GenS6{F0: v3, F1: ""}
                    _ = v50
                    _ = v51
                } else {
                    if v5 {
                        v52 := []int{(fi49 / (((fi49 * -903)) | 1)), fi49, (fi49 ^ (fi49 + fi47)), (fi47 / (((fi49 & -637)) | 1))}
                        var v53 int64 = -625
                        _ = v52
                        _ = v53
                    }
                    v2 = (v2 << ((v34) & 63))
                }
                var v54 uint8 = uint8(v48)
                switch (fi49 + fi49) {
                case -9223372036854775808:
                    v55 := GenS5{F0: (^(fi47)), F1: 0, F2: 1}
                    _ = v55
                case 0:
                    v48 += v48
                default:
                    v54 = (^(v54))
                }
                _ = fi49
                _ = v54
            }
            _ = fi47
            _ = v48
        }
    }
    v3 = (v3 % ((v3) | 1))
    switch "a" {
    case "a":
        fmt.Println(v2, v2)
    default:
        v56 := GenS7{F0: 719}
        _ = v56
    }
    fmt.Println(v5, v2)
    var v57 uint = v34
    for fi58 := 0; fi58 < 3; fi58++ {
        v59 := GenS3{F0: genFn1(127, (v34 & v57))}
        switch fi58 {
        case 128:
            for fi60 := 0; fi60 < 2; fi60++ {
                v57 += (v34 - 2)
                fmt.Println(v34, v5, v2)
                if v5 {
                    if false {
                        v5 = (v3 <= (v3 + (v3 / ((genFn1(173, (^(v57)))) | 1))))
                        if v5 {
                            if false {
                                v2 = genFn0(127, false, 0)
                            }
                        }
                        var v61 int32 = 246
                        var v62 float64 = 0.1
                        var v63 int16 = 1
                        _ = v61
                        _ = v62
                        _ = v63
                    } else {
                        if v5 {
                            var v64 uint8 = 0
                            _ = v64
                        }
                        v34 += (v57 - 657)
                    }
                    fmt.Println(v5, fi58)
                } else {
                    if v5 {
                        v65 := GenS7{F0: -714}
                        _ = len(v4)
                        fmt.Println(v57)
                        v66 := GenS8{F0: 244, F1: 1, F2: ("\n\t" + (("\n\t" + "\n\t") + "foo bar"))}
                        _ = v65
                        _ = v66
                    } else {
                        v3 += (^(v3))
                    }
                    var v67 uint8 = 229
                    v68 := GenS4{F0: (930 == uint64(v2)), F1: -1}
                    v69 := GenS9{F0: -9223372036854775808, F1: -40}
                    _ = v67
                    _ = v68
                    _ = v69
                }
                switch fi60 {
                case -543:
                    _ = len(v1)
                    switch ("hello" + "foo bar") {
                    case "":
                        fmt.Println(v34, fi60)
                    case "\n\t":
                        v57 = (v57 | v34)
                        switch (-(fi60)) {
                        case 499:
                            var v70 uint = (v57 / ((v34) | 1))
                            _ = v70
                        case 681:
                            if (v5 || false) {
                                switch ("" + "a") {
                                case "a":
                                    v71 := GenS10{F0: v5}
                                    _ = v71
                                }
                            }
                        }
                        if (127 < -1) {
                            v34 -= (v57 & v34)
                            v72 := GenS5{F0: (fi60 << (((v34 * 42)) & 63)), F1: 570, F2: (v57 * 2)}
                            v73 := GenS11{F0: (("foo bar" + "hello") + "\n\t")}
                            _ = v72
                            _ = v73
                        }
                    case "hello":
                        if (true && ((-2147483648 <= 42) && (!false))) {
                            v57 -= (v34 % ((v57) | 1))
                            v74 := GenS7{F0: 1}
                            _ = v74
                        }
                    }
                    v5 = (-40 != -832)
                case 9223372036854775807:
                    var v75 uint = v57
                    _ = v75
                case 0:
                    switch "hello" {
                    case "":
                        _ = len(v4)
                    default:
                        switch fi58 {
                        case 127:
                            v76 := GenS12{F0: fi58, F1: 1}
                            _ = v76
                        case 4294967295:
                            var v77 int8 = 42
                            _ = v77
                        }
                        if ((v2 | 4294967295) >= 4294967295) {
                            switch (^(fi58)) {
                            case 2147483647:
                                v5 = ((!v5) && true)
                                v34 = (v57 * v57)
                                v5 = v5
                            case 32768:
                                v78 := []int{fi60, 757, fi58, (fi58 + (fi60 ^ fi60))}
                                var v79 uint32 = v2
                                var v80 string = "\n\t"
                                _ = v78
                                _ = v79
                                _ = v80
                            case 1:
                                fmt.Println(v2, v3)
                            }
                            v34 -= 837
                            v81 := GenS13{F0: 561, F1: 1}
                            _ = v81
                        } else {
                            v34 += 333
                            if v5 {
                                v2 += v2
                                v34 += (v34 & 18446744073709551615)
                            }
                            if (0 == 1) {
                                if (v5 && (4 == 42)) {
                                    v5 = true
                                    if true {
                                        var v82 int = fi58
                                        if ((!(-1 >= -9223372036854775808)) || (-1 <= -1)) {
                                            v83 := GenS12{F0: -2147483648, F1: (v2 << (((v34 | v57)) & 63))}
                                            if true {
                                                if ("a" >= "a") {
                                                    v82 = fi58
                                                    v84 := GenS11{F0: ("a" + "\n\t")}
                                                    if ("\n\t" < "foo bar") {
                                                        var v85 uint32 = v2
                                                        _ = v85
                                                    } else {
                                                        if (42 != -1) {
                                                            var v86 string = (("hello" + ((("hello" + ("hello" + "hello")) + "") + "a")) + "foo bar")
                                                            _ = v86
                                                        } else {
                                                            var v87 uint = 18446744073709551615
                                                            var v88 int32 = 690
                                                            var v89 uint8 = 2
                                                            _ = v87
                                                            _ = v88
                                                            _ = v89
                                                        }
                                                        v90 := v4[(^(v2))]
                                                        _ = v90
                                                    }
                                                    v91 := []int{-9223372036854775808}
                                                    _ = v84
                                                    _ = v91
                                                }
                                                var v92 uint16 = genFn1(181, (^(v57)))
                                                _ = v92
                                            }
                                            var v93 int32 = -243
                                            _ = v83
                                            _ = v93
                                        } else {
                                            if ((v34 / ((v34) | 1)) <= v57) {
                                                fmt.Println(v82, v34, v82)
                                                switch (fi60 & (-(fi60))) {
                                                case 42:
                                                    var v94 int8 = -103
                                                    var v95 uint16 = uint16(v34)
                                                    v3 = 811
                                                    _ = v94
                                                    _ = v95
                                                case -663:
                                                    v96 := []int{}
                                                    _ = v96
                                                }
                                                if ((-1 == 42) || (!(true && (-521 <= 0)))) {
                                                    var v97 uint = (v34 % ((uint(v2)) | 1))
                                                    _ = v97
                                                }
                                                v5 = ((true && (!v5)) && false)
                                            }
                                            v82 -= (fi58 + (-(fi60)))
                                        }
                                        var v98 int32 = 1
                                        _ = v82
                                        _ = v98
                                    }
                                } else {
                                    v99 := v4[(v2 >> (((^(v57))) & 63))]
                                    switch (("foo bar" + (("" + "foo bar") + "a")) + "") {
                                    case "":
                                        if (471 > v2) {
                                            var v100 int = 42
                                            var v101 uint8 = 14
                                            _ = v100
                                            _ = v101
                                        } else {
                                            v34 -= v34
                                            var v102 int32 = -101
                                            if true {
                                                if v5 {
                                                    delete(v1, 2147483647)
                                                    var v103 float64 = float64(((float32(-2.25) + float32(((0.1 - 1e10) + -2.25))) / (((-2.25 + 0.0) * 1.5) / (-2.25 / 1e10))))
                                                    _ = len(v6)
                                                    fmt.Println(v102, v34, fi60)
                                                    v2 += v2
                                                    _ = v103
                                                } else {
                                                    v102 = (v102 << (((v57 | (v57 & v57))) & 63))
                                                    var v104 int32 = v102
                                                    v6["hello"] = v99
                                                    _ = v104
                                                }
                                            } else {
                                                v105 := GenS1{F0: v2}
                                                v106 := GenS14{F0: uint32(v3)}
                                                _ = v105
                                                _ = v106
                                            }
                                            _ = v102
                                        }
                                    case "\n\t":
                                        v3 -= v3
                                    case "foo bar":
                                        fmt.Println(fi58, v99, v2)
                                    }
                                    _ = v99
                                }
                                fmt.Println(fi60, v2)
                            } else {
                                fmt.Println(v57)
                            }
                        }
                        v107 := GenS8{F0: 18446744073709551615, F1: -2147483648, F2: "hello"}
                        switch (^(fi60)) {
                        case -2147483648:
                            if (!v5) {
                                v5 = v5
                                v108 := GenS3{F0: genFn1(42, v57)}
                                _ = v108
                            }
                        case 548:
                            v109 := GenS15{F0: 600, F1: v3, F2: 975}
                            var v110 uint8 = 0
                            var v111 bool = v5
                            fmt.Println(fi60, v3)
                            if v111 {
                                var v112 float64 = 0.0
                                v113 := GenS16{F0: (v3 ^ v3), F1: (^(v2))}
                                v114 := GenS13{F0: 0, F1: 229}
                                _ = v112
                                _ = v113
                                _ = v114
                            } else {
                                var v115 uint16 = (v3 % ((700) | 1))
                                switch "a" {
                                case "":
                                    if v5 {
                                        var v116 uint = (v57 + 73)
                                        if (255 < v3) {
                                            switch fi58 {
                                            case -1:
                                                var v117 uint8 = uint8(v116)
                                                _ = v117
                                            case 42:
                                                if v111 {
                                                    var v118 uint = (v34 - (v34 - v57))
                                                    v115 += (v115 | (v115 / ((v3) | 1)))
                                                    _ = v118
                                                }
                                            }
                                            v119 := v4[(v2 - 217)]
                                            _ = v119
                                        }
                                        switch "hello" {
                                        case "":
                                            v111 = v5
                                        case "hello":
                                            if true {
                                                v57 += 65535
                                                var v120 float32 = 0.0
                                                switch fi58 {
                                                case 128:
                                                    v5 = (235 != 18446744073709551615)
                                                case -39:
                                                    if v111 {
                                                        var v121 string = (("\n\t" + "hello") + "a")
                                                        _ = v121
                                                    } else {
                                                        var v122 uint = (v116 + 255)
                                                        var v123 int16 = 948
                                                        _ = v122
                                                        _ = v123
                                                    }
                                                default:
                                                    if (0 >= -435) {
                                                        v124 := GenS17{F0: -1}
                                                        v120 -= (1e10 * (1e10 - v120))
                                                        if ((fi58 + (fi58 - (fi60 * fi58))) != (fi60 % ((42) | 1))) {
                                                            _ = len(v6)
                                                        } else {
                                                            v125 := v4[uint32(v57)]
                                                            v126 := []int{fi58}
                                                            fmt.Println(v34, fi60)
                                                            _ = v125
                                                            _ = v126
                                                        }
                                                        _ = v124
                                                    } else {
                                                        if (v110 == v110) {
                                                            fmt.Println(fi60, v34)
                                                            var v127 float32 = 1.5
                                                            v57 -= v34
                                                            _ = v127
                                                        } else {
                                                            if (false && v5) {
                                                                var v128 int16 = 498
                                                                delete(v1, (^(v2)))
                                                                v1[uint32(v115)] = (fi60 & (fi58 << ((v34) & 63)))
                                                                _ = v128
                                                            } else {
                                                                v129 := GenS1{F0: (v2 / ((2147483648) | 1))}
                                                                v110 += uint8(v57)
                                                                v130 := GenS14{F0: genFn0(-111, v111, -667)}
                                                                _ = v129
                                                                _ = v130
                                                            }
                                                        }
                                                    }
                                                }
                                                var v131 string = ("foo bar" + "hello")
                                                _ = v120
                                                _ = v131
                                            }
                                        }
                                        _ = v116
                                    }
                                    v3 -= v115
                                    v115 = 2
                                    v34 += 2
                                }
                                _ = v115
                            }
                            switch "hello" {
                            case "\n\t":
                                switch ("hello" + "foo bar") {
                                case "":
                                    var v132 int16 = 42
                                    _ = v132
                                case "hello":
                                    if (v5 || (3.14159 >= 0.0)) {
                                        if (-9223372036854775808 != int64(v2)) {
                                            if false {
                                                if false {
                                                    v133 := []int{(^(fi58)), fi58, (fi60 & fi58), -798}
                                                    v34 = v34
                                                    _ = len(v4)
                                                    v57 += (^(v57))
                                                    _ = v133
                                                } else {
                                                    v134 := []int{128}
                                                    fmt.Println(v110, fi58, v34)
                                                    _ = v134
                                                }
                                                var v135 uint64 = 717
                                                if false {
                                                    switch "hello" {
                                                    case "a":
                                                        if (uint8(v34) == v110) {
                                                            fmt.Println(fi58, v135)
                                                            switch "\n\t" {
                                                            case "":
                                                                v34 = v57
                                                            case "\n\t":
                                                                if v111 {
                                                                    v2 = (v2 / ((v2) | 1))
                                                                    if false {
                                                                        fmt.Println(v57, v2)
                                                                    } else {
                                                                        var v136 uint16 = (^(v3))
                                                                        v137 := []int{}
                                                                        var v138 int64 = 0
                                                                        _ = v136
                                                                        _ = v137
                                                                        _ = v138
                                                                    }
                                                                } else {
                                                                    var v139 float32 = 0.1
                                                                    v140 := GenS18{F0: "\n\t", F1: 32767, F2: -2147483648}
                                                                    _ = v139
                                                                    _ = v140
                                                                }
                                                            case "foo bar":
                                                                v135 = (v135 / ((uint64(v57)) | 1))
                                                            }
                                                        }
                                                    case "":
                                                        fmt.Println(v111, v135)
                                                    case "hello":
                                                        v110 = (v110 % ((128) | 1))
                                                    }
                                                } else {
                                                    v5 = true
                                                    var v141 int8 = 55
                                                    v34 -= (v34 >> (((v34 << ((18446744073709551615) & 63))) & 63))
                                                    _ = v141
                                                }
                                                _ = v135
                                            }
                                            fmt.Println(v5, fi60)
                                        } else {
                                            v2 -= 4294967295
                                            var v142 int16 = 32767
                                            var v143 int8 = -29
                                            _ = v142
                                            _ = v143
                                        }
                                    } else {
                                        v144 := v1[(v2 - 2147483648)]
                                        if (false && v5) {
                                            switch 4294967295 {
                                            case 42:
                                                if v111 {
                                                    switch (("hello" + ("" + ("" + ""))) + "") {
                                                    case "hello":
                                                        v57 = 1
                                                    case "\n\t":
                                                        fmt.Println(v34, fi60, v3)
                                                    }
                                                    if v5 {
                                                        v2 = 1
                                                    } else {
                                                        v111 = (fi60 == (fi58 ^ (v144 - fi58)))
                                                        if v5 {
                                                            v145 := GenS17{F0: 2147483647}
                                                            var v146 int64 = int64(v2)
                                                            v147 := []int{(fi60 * v144), fi60, -2147483648, fi60}
                                                            _ = v145
                                                            _ = v146
                                                            _ = v147
                                                        } else {
                                                            if true {
                                                                v148 := GenS19{F0: (v110 | 42), F1: -32768, F2: 2}
                                                                switch (v144 << (((v34 | (v34 / (((v34 / ((1) | 1))) | 1)))) & 63)) {
                                                                case 1:
                                                                    if (0 == 0) {
                                                                        v1[v2] = v144
                                                                        v149 := GenS17{F0: -40}
                                                                        var v150 uint8 = uint8(v57)
                                                                        _ = v149
                                                                        _ = v150
                                                                    }
                                                                case -917:
                                                                    var v151 uint8 = (v110 ^ v110)
                                                                    _ = v151
                                                                default:
                                                                    fmt.Println(v57, v144)
                                                                }
                                                                var v152 float64 = ((float64(0.1) / ((0.1 + ((0.0 / (3.14159 + 0.0)) - -2.25)) + 0.1)) * 1e10)
                                                                fmt.Println(v34, v34)
                                                                v152 = ((v152 * 3.14159) * 1.5)
                                                                var v153 int = (v144 & v144)
                                                                _ = v148
                                                                _ = v152
                                                                _ = v153
                                                            } else {
                                                                v110 = (v110 - v110)
                                                                v154 := []int{127}
                                                                if v5 {
                                                                    if v5 {
                                                                        var v155 float32 = float32(-2.25)
                                                                        _ = v155
                                                                    }
                                                                }
                                                                _ = v154
                                                            }
                                                        }
                                                    }
                                                    v111 = v5
                                                    switch (fi60 - -369) {
                                                    case 446:
                                                        v156 := []int{v144}
                                                        _ = v156
                                                    case -958:
                                                        if false {
                                                            if ((^(v110)) > (v110 + v110)) {
                                                                fmt.Println(v2, v110, v144)
                                                                fmt.Println(v110)
                                                                if ((((((v34 % ((2) | 1)) > v57) || false) || true) || (!(false || false))) || false) {
                                                                    v157 := GenS16{F0: 0, F1: genFn0(int8(v3), true, 2147483647)}
                                                                    v111 = (-2147483648 > (v144 >> ((v34) & 63)))
                                                                    _ = v157
                                                                } else {
                                                                    v158 := GenS12{F0: int(v34), F1: genFn0(127, (-2.25 != -2.25), 1)}
                                                                    v144 += (fi60 << ((18446744073709551615) & 63))
                                                                    v159 := GenS14{F0: (v2 >> ((uint(v2)) & 63))}
                                                                    _ = v158
                                                                    _ = v159
                                                                }
                                                                delete(v1, genFn0(1, v111, 0))
                                                                switch -883 {
                                                                case -210:
                                                                    var v160 int32 = 2147483647
                                                                    _ = v160
                                                                case 0:
                                                                    var v161 uint8 = (v110 % (((v110 / (((v110 | v110)) | 1))) | 1))
                                                                    _ = v161
                                                                default:
                                                                    if v111 {
                                                                        v6[""] = (!(uint(v3) <= (^(v34))))
                                                                        v34 -= (v34 ^ (v57 >> (((v57 - uint(v3))) & 63)))
                                                                    } else {
                                                                        var v162 int = fi60
                                                                        if v111 {
                                                                            v3 = genFn1(v110, (^(v57)))
                                                                        }
                                                                        _ = v162
                                                                    }
                                                                }
                                                                switch ("\n\t" + (("\n\t" + "a") + (("foo bar" + "a") + "\n\t"))) {
                                                                case "":
                                                                    switch (v144 + v144) {
                                                                    case 42:
                                                                        v163 := GenS11{F0: "a"}
                                                                        _ = v163
                                                                    case 178:
                                                                        if v5 {
                                                                            fmt.Println(fi58, v110, fi58)
                                                                        } else {
                                                                            fmt.Println(v5)
                                                                        }
                                                                    case -330:
                                                                        v164 := GenS20{F0: 32767, F1: -128}
                                                                        _ = v164
                                                                    }
                                                                }
                                                                switch "a" {
                                                                case "hello":
                                                                    switch v144 {
                                                                    case -2147483648:
                                                                        if v5 {
                                                                            var v165 uint16 = (v3 - (v3 - v3))
                                                                            _ = v165
                                                                        }
                                                                    default:
                                                                        if (!(v3 > uint16(v110))) {
                                                                            v166 := GenS11{F0: (("a" + "") + ("foo bar" + (("a" + "a") + "foo bar")))}
                                                                            v167 := GenS7{F0: int32(v2)}
                                                                            _ = v166
                                                                            _ = v167
                                                                        } else {
                                                                            var v168 uint = v57
                                                                            v169 := GenS21{F0: 42}
                                                                            var v170 int8 = 111
                                                                            _ = v168
                                                                            _ = v169
                                                                            _ = v170
                                                                        }
                                                                    }
                                                                default:
                                                                    v171 := GenS18{F0: ("a" + ""), F1: 1, F2: -773}
                                                                    _ = v171
                                                                }
                                                                var v172 uint = 18446744073709551615
                                                                var v173 int8 = -64
                                                                var v174 int64 = -490
                                                                _ = v172
                                                                _ = v173
                                                                _ = v174
                                                            } else {
                                                                v57 -= 42
                                                                if (v111 || (!((v3 - (v3 - (v3 ^ (v3 >> ((v57) & 63))))) >= 65535))) {
                                                                    _ = len(v1)
                                                                    var v175 float32 = 1.5
                                                                    _ = v175
                                                                }
                                                                fmt.Println(v2, fi60)
                                                                v5 = (!(-706 <= -1))
                                                            }
                                                            var v176 uint8 = (v110 * v110)
                                                            _ = v176
                                                        } else {
                                                            v177 := GenS22{F0: (v34 | (v57 * 837))}
                                                            _ = v177
                                                        }
                                                    case 32767:
                                                        v178 := GenS13{F0: 42, F1: -2147483648}
                                                        _ = v178
                                                    default:
                                                        v34 += (^(v57))
                                                    }
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        switch -11 {
                                                        case 4294967295:
                                                            v179 := GenS8{F0: 0, F1: 0, F2: ("hello" + (("" + "foo bar") + ("\n\t" + "foo bar")))}
                                                            _ = v179
                                                        case 1:
                                                            v5 = true
                                                        case 9223372036854775807:
                                                            if (127 <= 126) {
                                                                var v180 float32 = (((-2.25 / (-2.25 + 0.0)) * 3.14159) / -2.25)
                                                                v34 = 1
                                                                _ = v180
                                                            }
                                                        default:
                                                            var v181 int32 = 42
                                                            _ = v181
                                                        }
                                                        v57 -= v57
                                                        var v182 uint16 = (v3 ^ genFn1((v110 ^ (v110 - uint8(v3))), 65535))
                                                        _ = v182
                                                    case "a":
                                                        if v111 {
                                                            var v183 float64 = (0.1 + 3.14159)
                                                            var v184 uint64 = 4294967295
                                                            var v185 uint64 = v184
                                                            _ = v183
                                                            _ = v184
                                                            _ = v185
                                                        }
                                                    case "":
                                                        if (0 <= 42) {
                                                            var v186 uint64 = 42
                                                            v111 = (!((-1 == 1) && (v110 <= 2)))
                                                            _ = v186
                                                        }
                                                    }
                                                } else {
                                                    v110 -= uint8(v3)
                                                }
                                            case -1:
                                                var v187 int64 = -9223372036854775808
                                                _ = v187
                                            case 4294967295:
                                                v5 = (((v2 >> ((v57) & 63)) <= v2) || false)
                                            default:
                                                delete(v1, genFn0(127, (v110 <= (v110 >> ((v34) & 63))), 2147483647))
                                            }
                                        }
                                        _ = v144
                                    }
                                case "a":
                                    var v188 int64 = -342
                                    _ = v188
                                default:
                                    var v189 uint8 = 42
                                    _ = v189
                                }
                            case "a":
                                var v190 int8 = 1
                                _ = v190
                            default:
                                var v191 int64 = -1
                                _ = v191
                            }
                            _ = v109
                            _ = v110
                            _ = v111
                        }
                        _ = v107
                    }
                }
                _ = fi60
            }
        case -9223372036854775808:
            for fi192 := 0; fi192 < 7; fi192++ {
                var v193 uint8 = 0
                switch ("a" + "\n\t") {
                case "hello":
                    v194 := []int{}
                    switch fi192 {
                    case 9223372036854775807:
                        var v195 float64 = ((0.0 - -2.25) * float64(((1.5 - 1.5) * 1.5)))
                        _ = v195
                    default:
                        switch 32768 {
                        case 9223372036854775807:
                            fmt.Println(v5, v2)
                        case 128:
                            if ((-1 < int16(v57)) || (false && (v5 && false))) {
                                v196 := GenS18{F0: "a", F1: 42, F2: 664}
                                _ = v196
                            } else {
                                v197 := GenS5{F0: int(v2), F1: 376, F2: v34}
                                _ = v197
                            }
                            switch (fi58 * fi58) {
                            case -120:
                                _ = len(v1)
                            default:
                                var v198 int8 = -128
                                _ = v198
                            }
                        case 0:
                            if false {
                                var v199 uint64 = 0
                                if (!((v5 || true) && (3.14159 > (1.5 + ((float64(0.1) / 0.1) / 1.5))))) {
                                    v34 += 2
                                    _ = len(v6)
                                    switch fi58 {
                                    case 42:
                                        if ((v57 + v34) == uint(v193)) {
                                            v199 = (^(v199))
                                            delete(v4, uint32(v3))
                                            if v5 {
                                                if ((v34 == (v57 ^ v34)) && (!((((v3 << ((v34) & 63)) <= 0) && ((!(((v193 >> ((1) & 63)) > (v193 * (v193 & v193))) || v5)) && false)) || (fi58 > (fi192 / ((fi58) | 1)))))) {
                                                    v199 += (^(v199))
                                                    v2 = v2
                                                    fmt.Println(fi58, v34)
                                                }
                                                v193 = 28
                                            } else {
                                                if ((v2 << (((v34 % ((18446744073709551615) | 1))) & 63)) <= (v2 << ((v34) & 63))) {
                                                    fmt.Println(v5, v199, v5)
                                                    fmt.Println(v34)
                                                    switch "foo bar" {
                                                    case "a":
                                                        switch "hello" {
                                                        case "":
                                                            fmt.Println(v34)
                                                        case "hello":
                                                            var v200 int8 = 42
                                                            _ = v200
                                                        case "foo bar":
                                                            v201 := GenS13{F0: 2, F1: 42}
                                                            _ = v201
                                                        default:
                                                            var v202 int32 = 42
                                                            _ = v202
                                                        }
                                                    case "hello":
                                                        if true {
                                                            v203 := GenS23{F0: 61, F1: int64(v2), F2: v199}
                                                            if (-423 > 952) {
                                                                v204 := v4[(v2 - v2)]
                                                                _ = v204
                                                            } else {
                                                                var v205 float32 = (float32((0.0 / float64(0.1))) / 3.14159)
                                                                if (0 <= 9223372036854775807) {
                                                                    if (("" + "a") <= "\n\t") {
                                                                        var v206 uint8 = v193
                                                                        v199 = (^(v199))
                                                                        v5 = ("foo bar" <= "foo bar")
                                                                        var v207 int16 = 42
                                                                        v206 = uint8(v3)
                                                                        _ = v206
                                                                        _ = v207
                                                                    }
                                                                    fmt.Println(v205, v3, v205)
                                                                    v208 := GenS24{F0: 127}
                                                                    v34 = (v57 >> ((0) & 63))
                                                                    var v209 uint16 = v3
                                                                    _ = v208
                                                                    _ = v209
                                                                } else {
                                                                    v199 = uint64(v57)
                                                                    switch (-(fi58)) {
                                                                    case -31:
                                                                        if (v57 != (v34 ^ v34)) {
                                                                            v193 -= 42
                                                                            var v210 float32 = v205
                                                                            _ = v210
                                                                        } else {
                                                                            switch "hello" {
                                                                            case "foo bar":
                                                                                fmt.Println(v193, v57, v3)
                                                                            case "":
                                                                                v57 += (v34 * (v57 ^ uint(v2)))
                                                                            case "hello":
                                                                                var v211 uint = (v57 | v34)
                                                                                _ = v211
                                                                            }
                                                                        }
                                                                    case 1:
                                                                        _ = len(v194)
                                                                    case 874:
                                                                        v5 = v5
                                                                    default:
                                                                        v212 := GenS7{F0: 1}
                                                                        _ = v212
                                                                    }
                                                                    v205 = ((v205 - (float32(0.0) / 1.0)) / 1.0)
                                                                }
                                                                _ = v205
                                                            }
                                                            var v213 float64 = (3.14159 + ((-2.25 / -2.25) * float64((float32((0.1 / 1.0)) - -2.25))))
                                                            _ = v203
                                                            _ = v213
                                                        } else {
                                                            switch (fi58 + fi58) {
                                                            case 2147483647:
                                                                switch ("hello" + "\n\t") {
                                                                case "":
                                                                    v6[""] = v5
                                                                }
                                                            case 1:
                                                                var v214 int64 = 284
                                                                _ = v214
                                                            case -9223372036854775808:
                                                                v4[(^(v2))] = true
                                                            }
                                                            switch "a" {
                                                            case "hello":
                                                                if false {
                                                                    switch fi192 {
                                                                    case 127:
                                                                        var v215 string = ("a" + "a")
                                                                        _ = v215
                                                                    default:
                                                                        if v5 {
                                                                            v34 += (v57 << (((v57 + (v34 + (v57 >> ((v34) & 63))))) & 63))
                                                                            if (false && ((v193 % ((v193) | 1)) != (v193 / (((v193 * 127)) | 1)))) {
                                                                                v216 := GenS25{F0: 85, F1: v193, F2: uint8(v57)}
                                                                                v5 = (true || (0.0 == 3.14159))
                                                                                fmt.Println(v199, v199, v199)
                                                                                _ = v216
                                                                            }
                                                                            if len(v194) >= 2 {
                                                                                v194 = v194[1:len(v194)]
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    if (!(!v5)) {
                                                                        v199 += uint64(v2)
                                                                        v217 := GenS22{F0: v34}
                                                                        var v218 bool = v5
                                                                        _ = v217
                                                                        _ = v218
                                                                    } else {
                                                                        fmt.Println(v34)
                                                                    }
                                                                    v219 := GenS19{F0: v193, F1: -589, F2: 4294967295}
                                                                    _ = v219
                                                                }
                                                            case "\n\t":
                                                                v57 += v34
                                                            }
                                                            v5 = v5
                                                        }
                                                    }
                                                }
                                                fmt.Println(v199)
                                            }
                                        } else {
                                            if true {
                                                var v220 uint64 = 109
                                                _ = v220
                                            }
                                        }
                                        if (((fi58 % (((-(fi58))) | 1)) > int(v193)) || v5) {
                                            fmt.Println(v5)
                                        } else {
                                            v34 += uint(v3)
                                            _ = len(v4)
                                            var v221 int64 = 1
                                            _ = v221
                                        }
                                        if false {
                                            var v222 string = ("foo bar" + ("a" + ("\n\t" + ("" + ""))))
                                            _ = v222
                                        }
                                        v194 = append(v194, fi192)
                                        if (0 != (v57 | (^(v57)))) {
                                            v199 = v199
                                        } else {
                                            v223 := GenS23{F0: (^(fi58)), F1: 42, F2: v199}
                                            _ = v223
                                        }
                                        v34 -= v34
                                    case 889:
                                        if v5 {
                                            fmt.Println(fi58)
                                        } else {
                                            fmt.Println(v5, fi192)
                                            v224 := GenS26{F0: -320}
                                            var v225 uint16 = (v3 * uint16(v193))
                                            v225 = (v3 << ((v57) & 63))
                                            _ = v224
                                            _ = v225
                                        }
                                    case 2147483647:
                                        v226 := GenS7{F0: 0}
                                        _ = v226
                                    default:
                                        delete(v4, (^(v2)))
                                    }
                                } else {
                                    if (!(1e10 < 1e10)) {
                                        var v227 uint64 = v199
                                        v228 := GenS1{F0: (v2 >> ((42) & 63))}
                                        v229 := GenS27{F0: 75, F1: fi58, F2: (v2 >= v2)}
                                        if ((true || v5) && (!(fi192 != (fi58 | -159)))) {
                                            if v5 {
                                                v3 += (^(v3))
                                                if (!false) {
                                                    var v230 bool = (("" + "hello") <= ("" + "foo bar"))
                                                    _ = v230
                                                }
                                            }
                                            if ((!true) || (fi58 != (fi58 - fi58))) {
                                                v231 := GenS8{F0: 18446744073709551615, F1: int32(v34), F2: ""}
                                                _ = v231
                                            } else {
                                                fmt.Println(v199, v3)
                                                v232 := GenS28{F0: v3, F1: 866, F2: -1}
                                                _ = v232
                                            }
                                        } else {
                                            switch (^(fi192)) {
                                            case 128:
                                                if v5 {
                                                    v233 := GenS29{F0: uint32(v34)}
                                                    v234 := GenS30{F0: 42, F1: true, F2: v193}
                                                    _ = v233
                                                    _ = v234
                                                } else {
                                                    if true {
                                                        v5 = v5
                                                        var v235 float64 = 0.0
                                                        v34 = v57
                                                        _ = v235
                                                    }
                                                    switch 32768 {
                                                    case 976:
                                                        switch "\n\t" {
                                                        case "a":
                                                            var v236 uint16 = genFn1((v193 + (v193 * uint8(v34))), v34)
                                                            if (v57 > v34) {
                                                                if len(v194) > 0 {
                                                                    v194[0] = -662
                                                                }
                                                                v237 := GenS4{F0: true, F1: 0}
                                                                _ = v237
                                                            }
                                                            if (!(0 <= -491)) {
                                                                if false {
                                                                    v238 := GenS3{F0: 65535}
                                                                    var v239 float32 = (float32(float64(1.5)) + float32((float64(-2.25) * (0.1 / (((-2.25 / (float64(float32(float64(3.14159))) - 0.1)) * (3.14159 * 1e10)) / 1e10)))))
                                                                    v240 := GenS31{F0: fi192, F1: 0, F2: 32767}
                                                                    _ = v238
                                                                    _ = v239
                                                                    _ = v240
                                                                } else {
                                                                    var v241 bool = (!((!v5) && v5))
                                                                    v3 = 32767
                                                                    _ = v241
                                                                }
                                                                if (v5 || false) {
                                                                    if v5 {
                                                                        v5 = true
                                                                        if (((4294967295 > int(v236)) && (2147483647 != -317)) && v5) {
                                                                            switch fi58 {
                                                                            case 42:
                                                                                v194 = append(v194, (fi58 * fi192))
                                                                                if (409 > -1) {
                                                                                    switch fi58 {
                                                                                    case -2147483648:
                                                                                        var v242 bool = false
                                                                                        _ = v242
                                                                                    case 128:
                                                                                        if v5 {
                                                                                            switch ("\n\t" + "") {
                                                                                            case "hello":
                                                                                                v6[("hello" + "a")] = v5
                                                                                            case "foo bar":
                                                                                                v243 := GenS32{F0: -1, F1: 255, F2: -1}
                                                                                                _ = v243
                                                                                            default:
                                                                                                var v244 uint16 = 0
                                                                                                _ = v244
                                                                                            }
                                                                                        } else {
                                                                                            var v245 int64 = -355
                                                                                            if false {
                                                                                                switch "foo bar" {
                                                                                                case "":
                                                                                                    v246 := GenS33{F0: -9223372036854775808, F1: 42, F2: fi58}
                                                                                                    _ = v246
                                                                                                case "a":
                                                                                                    switch (fi58 + fi192) {
                                                                                                    case -890:
                                                                                                        var v247 uint32 = 205
                                                                                                        fmt.Println(v247, v199)
                                                                                                        delete(v6, ("a" + ""))
                                                                                                        _ = v247
                                                                                                    default:
                                                                                                        if true {
                                                                                                            var v248 uint64 = v227
                                                                                                            fmt.Println(v199, v245, v3)
                                                                                                            if (v5 && (!v5)) {
                                                                                                                switch (fi192 * fi58) {
                                                                                                                case -9223372036854775808:
                                                                                                                    var v249 int16 = -32768
                                                                                                                    _ = v249
                                                                                                                case -349:
                                                                                                                    if (1 > 42) {
                                                                                                                        v250 := GenS23{F0: 2147483647, F1: (v245 * v245), F2: v248}
                                                                                                                        var v251 uint32 = (v2 % (((^(v2))) | 1))
                                                                                                                        if (42 <= uint16(v34)) {
                                                                                                                            switch ("foo bar" + "") {
                                                                                                                            case "\n\t":
                                                                                                                                v199 = v248
                                                                                                                            default:
                                                                                                                                v252 := GenS30{F0: -49, F1: (v2 != (v251 % (((^(v251))) | 1))), F2: (v193 << (((v57 | 0)) & 63))}
                                                                                                                                _ = v252
                                                                                                                            }
                                                                                                                            if ((((((v193 != (v193 >> (((v57 - 308)) & 63))) && v5) && true) && true) && (785 > 1000)) && v5) {
                                                                                                                                if v5 {
                                                                                                                                    var v253 bool = true
                                                                                                                                    _ = v253
                                                                                                                                } else {
                                                                                                                                    var v254 float64 = ((float64(-2.25) - float64((-2.25 / (1e10 - (0.0 / 1.0))))) - float64(3.14159))
                                                                                                                                    _ = v254
                                                                                                                                }
                                                                                                                                var v255 uint = v34
                                                                                                                                var v256 string = "a"
                                                                                                                                _ = v255
                                                                                                                                _ = v256
                                                                                                                            } else {
                                                                                                                                var v257 uint32 = (^(v251))
                                                                                                                                var v258 int64 = v245
                                                                                                                                v227 = 2
                                                                                                                                _ = v257
                                                                                                                                _ = v258
                                                                                                                            }
                                                                                                                            v2 = v251
                                                                                                                        }
                                                                                                                        _ = v250
                                                                                                                        _ = v251
                                                                                                                    }
                                                                                                                }
                                                                                                                var v259 uint = (v34 + uint(v227))
                                                                                                                _ = v259
                                                                                                            }
                                                                                                            v227 -= uint64(v2)
                                                                                                            _ = v248
                                                                                                        } else {
                                                                                                            var v260 float32 = 1e10
                                                                                                            switch (("" + "hello") + "hello") {
                                                                                                            case "a":
                                                                                                                switch fi192 {
                                                                                                                case 42:
                                                                                                                    fmt.Println(fi192)
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                v199 += v199
                                                                                                            }
                                                                                                            v3 -= (v3 | (v236 << (((^(v34))) & 63)))
                                                                                                            _ = v260
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    var v261 int = -965
                                                                                                    _ = v261
                                                                                                }
                                                                                                var v262 float32 = 1.5
                                                                                                v193 -= uint8(v236)
                                                                                                fmt.Println(v199, fi58, v34)
                                                                                                _ = v262
                                                                                            }
                                                                                            _ = v245
                                                                                        }
                                                                                    }
                                                                                    v193 = v193
                                                                                    if (!v5) {
                                                                                        v263 := GenS34{F0: genFn1((v193 % ((188) | 1)), 0), F1: v199, F2: (^(v236))}
                                                                                        _ = v263
                                                                                    }
                                                                                }
                                                                            }
                                                                            v57 += (v57 * 768)
                                                                        }
                                                                        fmt.Println(v57, fi192)
                                                                    } else {
                                                                        if (!((!v5) || false)) {
                                                                            v264 := GenS27{F0: 100, F1: int(v227), F2: (1000 == 0)}
                                                                            v193 = v193
                                                                            switch 0 {
                                                                            case -2147483648:
                                                                                v2 += 4294967295
                                                                            case -313:
                                                                                switch 42 {
                                                                                case 2147483647:
                                                                                    var v265 uint = uint(v2)
                                                                                    _ = v265
                                                                                case -1:
                                                                                    v193 = (v193 & v193)
                                                                                case 4294967295:
                                                                                    _ = len(v6)
                                                                                }
                                                                            }
                                                                            _ = v264
                                                                        } else {
                                                                            v3 -= genFn1(uint8(v2), 65535)
                                                                            var v266 bool = v5
                                                                            switch (fi192 << ((v34) & 63)) {
                                                                            case 42:
                                                                                if v266 {
                                                                                    v34 -= v57
                                                                                    var v267 uint64 = 1
                                                                                    if (1 == 2147483647) {
                                                                                        if (127 > -101) {
                                                                                            fmt.Println(v236)
                                                                                            v268 := GenS35{F0: 0, F1: (v236 / (((v3 >> ((v57) & 63))) | 1)), F2: 100}
                                                                                            if (fi192 <= (fi58 << ((v34) & 63))) {
                                                                                                var v269 float32 = 0.1
                                                                                                v4[728] = false
                                                                                                _ = v269
                                                                                            }
                                                                                            _ = v268
                                                                                        } else {
                                                                                            var v270 string = ""
                                                                                            var v271 uint = uint(v3)
                                                                                            fmt.Println(fi192)
                                                                                            _ = v270
                                                                                            _ = v271
                                                                                        }
                                                                                        v236 = (v3 >> ((v57) & 63))
                                                                                        var v272 float64 = (3.14159 / (-2.25 * (((0.1 + -2.25) * 3.14159) + ((0.1 - float64(1e10)) * (1e10 * 1e10)))))
                                                                                        _ = v272
                                                                                    } else {
                                                                                        if (-115 <= 0) {
                                                                                            v273 := GenS3{F0: (v3 + (^(v3)))}
                                                                                            v274 := GenS36{F0: 923, F1: uint32(v267)}
                                                                                            fmt.Println(v227)
                                                                                            switch "hello" {
                                                                                            case "hello":
                                                                                                fmt.Println(v3, v2)
                                                                                            case "a":
                                                                                                var v275 bool = ((0.0 - 0.0) <= 1e10)
                                                                                                _ = v275
                                                                                            }
                                                                                            var v276 uint = 255
                                                                                            _ = v273
                                                                                            _ = v274
                                                                                            _ = v276
                                                                                        } else {
                                                                                            v194 = append(v194, (fi192 % ((-495) | 1)))
                                                                                            if ((v2 & (v2 / ((2147483648) | 1))) >= (v2 % (((v2 | 1)) | 1))) {
                                                                                                var v277 float64 = 1e10
                                                                                                _ = v277
                                                                                            }
                                                                                        }
                                                                                        v236 += 65535
                                                                                    }
                                                                                    _ = v267
                                                                                }
                                                                            }
                                                                            _ = v266
                                                                        }
                                                                    }
                                                                    if false {
                                                                        if (v5 || false) {
                                                                            switch ("hello" + "hello") {
                                                                            case "hello":
                                                                                var v278 int32 = 43
                                                                                var v279 float64 = float64(3.14159)
                                                                                _ = v278
                                                                                _ = v279
                                                                            case "":
                                                                                v236 = 2
                                                                            }
                                                                            v280 := GenS5{F0: fi58, F1: 0, F2: v34}
                                                                            _ = v280
                                                                        } else {
                                                                            if len(v194) > 0 {
                                                                                v194[0] = fi58
                                                                            }
                                                                            v199 += v199
                                                                            if ((!(true && (-103 >= 1))) && false) {
                                                                                if false {
                                                                                    v2 = v2
                                                                                    var v281 float64 = 1e10
                                                                                    _ = v281
                                                                                }
                                                                            } else {
                                                                                v282 := GenS11{F0: "a"}
                                                                                if ((61 <= (^(v236))) || v5) {
                                                                                    var v283 int32 = int32(v193)
                                                                                    _ = v283
                                                                                } else {
                                                                                    var v284 uint16 = 255
                                                                                    if v5 {
                                                                                        v284 -= (v284 - (v236 & (v3 + (v3 % ((genFn1((v193 * 89), v57)) | 1)))))
                                                                                        if v5 {
                                                                                            if len(v194) >= 2 {
                                                                                                v194 = v194[1:len(v194)]
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        if v5 {
                                                                                            delete(v1, 0)
                                                                                            var v285 int64 = -654
                                                                                            _ = v285
                                                                                        } else {
                                                                                            v193 = 1
                                                                                            var v286 int8 = int8(fi192)
                                                                                            v236 -= 1
                                                                                            if (32767 <= genFn1((v193 << ((2) & 63)), v34)) {
                                                                                                fmt.Println(v34)
                                                                                                v227 -= v199
                                                                                            }
                                                                                            v284 = (^(v3))
                                                                                            _ = v286
                                                                                        }
                                                                                        if len(v194) > 0 {
                                                                                            v194[0] = (-(fi58))
                                                                                        }
                                                                                        fmt.Println(v236, v5)
                                                                                    }
                                                                                    var v287 string = "foo bar"
                                                                                    _ = v284
                                                                                    _ = v287
                                                                                }
                                                                                switch ("hello" + "\n\t") {
                                                                                case "foo bar":
                                                                                    switch "" {
                                                                                    case "foo bar":
                                                                                        v288 := GenS37{F0: (v2 - (^(v2))), F1: (fi192 >> (((v57 - v34)) & 63))}
                                                                                        _ = v288
                                                                                    case "hello":
                                                                                        if (-995 >= 1) {
                                                                                            switch "a" {
                                                                                            case "\n\t":
                                                                                                v5 = (genFn1(v193, (v57 >> (((v34 & v57)) & 63))) > (v3 / ((v3) | 1)))
                                                                                            case "":
                                                                                                v194 = append(v194, fi58)
                                                                                            default:
                                                                                                switch fi192 {
                                                                                                case 0:
                                                                                                    switch fi58 {
                                                                                                    case 32767:
                                                                                                        var v289 uint = 0
                                                                                                        _ = v289
                                                                                                    case 0:
                                                                                                        v193 = v193
                                                                                                    }
                                                                                                case -9:
                                                                                                    switch "hello" {
                                                                                                    case "\n\t":
                                                                                                        if false {
                                                                                                            var v290 uint8 = (v193 << (((v57 ^ uint(v2))) & 63))
                                                                                                            if true {
                                                                                                                var v291 int16 = -979
                                                                                                                v292 := v1[(^(v2))]
                                                                                                                fmt.Println(v2)
                                                                                                                _ = v291
                                                                                                                _ = v292
                                                                                                            }
                                                                                                            _ = v290
                                                                                                        } else {
                                                                                                            v5 = v5
                                                                                                        }
                                                                                                    case "foo bar":
                                                                                                        v5 = (1 < 0)
                                                                                                    case "":
                                                                                                        fmt.Println(v227, v3, v57)
                                                                                                    }
                                                                                                default:
                                                                                                    if (0.1 != -2.25) {
                                                                                                        switch ("foo bar" + "foo bar") {
                                                                                                        case "hello":
                                                                                                            if v5 {
                                                                                                                var v293 int = 32768
                                                                                                                v227 = v199
                                                                                                                if true {
                                                                                                                    v193 -= 1
                                                                                                                    v294 := GenS38{F0: (v2 ^ 128), F1: 42}
                                                                                                                    var v295 uint16 = 255
                                                                                                                    _ = v294
                                                                                                                    _ = v295
                                                                                                                } else {
                                                                                                                    if ((v193 >> (((v57 - v34)) & 63)) == (v193 - v193)) {
                                                                                                                        v236 = 42
                                                                                                                        v34 += (v34 * (^(v34)))
                                                                                                                        var v296 int64 = 1
                                                                                                                        _ = v296
                                                                                                                    }
                                                                                                                    v297 := v6[("foo bar" + ("hello" + ("hello" + "hello")))]
                                                                                                                    var v298 int64 = -9223372036854775808
                                                                                                                    _ = v297
                                                                                                                    _ = v298
                                                                                                                }
                                                                                                                _ = v293
                                                                                                            }
                                                                                                        default:
                                                                                                            fmt.Println(v227, v57, v236)
                                                                                                        }
                                                                                                        var v299 int64 = int64(v193)
                                                                                                        v300 := GenS39{F0: v5, F1: "\n\t"}
                                                                                                        var v301 int8 = 42
                                                                                                        v236 += v236
                                                                                                        _ = v299
                                                                                                        _ = v300
                                                                                                        _ = v301
                                                                                                    }
                                                                                                    fmt.Println(fi192, fi192)
                                                                                                    if v5 {
                                                                                                        v302 := GenS18{F0: ("" + ""), F1: -32768, F2: 0}
                                                                                                        v193 -= (v193 << ((v34) & 63))
                                                                                                        if (!true) {
                                                                                                            fmt.Println(v227, v57, v199)
                                                                                                        }
                                                                                                        _ = v302
                                                                                                    } else {
                                                                                                        v303 := GenS34{F0: (v236 << ((uint(v2)) & 63)), F1: 200, F2: v236}
                                                                                                        if v5 {
                                                                                                            var v304 float64 = (0.0 + 0.1)
                                                                                                            _ = v304
                                                                                                        } else {
                                                                                                            fmt.Println(v3)
                                                                                                        }
                                                                                                        fmt.Println(v3, v236, v3)
                                                                                                        _ = v303
                                                                                                    }
                                                                                                    _ = len(v194)
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    case "":
                                                                                        v5 = false
                                                                                    default:
                                                                                        var v305 float64 = 1e10
                                                                                        _ = v305
                                                                                    }
                                                                                case "hello":
                                                                                    switch ("foo bar" + "") {
                                                                                    case "foo bar":
                                                                                        fmt.Println(v3, fi192, fi58)
                                                                                    }
                                                                                }
                                                                                _ = v282
                                                                            }
                                                                        }
                                                                        var v306 string = ("\n\t" + "a")
                                                                        _ = v306
                                                                    }
                                                                }
                                                            } else {
                                                                if (v57 != uint(v199)) {
                                                                    var v307 uint16 = (v3 & uint16(v2))
                                                                    var v308 int16 = 1000
                                                                    _ = v307
                                                                    _ = v308
                                                                }
                                                                var v309 int64 = -744
                                                                v310 := GenS19{F0: uint8(v34), F1: 42, F2: 42}
                                                                v311 := GenS40{F0: false}
                                                                v3 = v236
                                                                _ = v309
                                                                _ = v310
                                                                _ = v311
                                                            }
                                                            _ = v236
                                                        case "\n\t":
                                                            var v312 float32 = 0.0
                                                            _ = v312
                                                        case "foo bar":
                                                            v313 := v6["hello"]
                                                            _ = v313
                                                        }
                                                    case -2147483648:
                                                        v227 = uint64(v3)
                                                    case 42:
                                                        var v314 bool = (!true)
                                                        _ = v314
                                                    }
                                                    switch 32768 {
                                                    case 127:
                                                        var v315 uint32 = 486
                                                        _ = v315
                                                    case 2147483647:
                                                        if (1 == 884) {
                                                            v2 = (v2 / ((genFn0(-1, (v193 != uint8(v3)), 0)) | 1))
                                                        } else {
                                                            if len(v194) > 0 {
                                                                v194[0] = fi58
                                                            }
                                                            if (!(v5 || (1 > -32768))) {
                                                                var v316 uint = (v57 | (v57 % (((v34 >> ((v57) & 63))) | 1)))
                                                                var v317 string = "\n\t"
                                                                v318 := GenS37{F0: v2, F1: (-(fi58))}
                                                                v319 := GenS32{F0: 32767, F1: 127, F2: 42}
                                                                fmt.Println(v5)
                                                                if false {
                                                                    switch v317 {
                                                                    case "hello":
                                                                        switch v317 {
                                                                        case "foo bar":
                                                                            var v320 int32 = 42
                                                                            if v5 {
                                                                                v227 = v227
                                                                                v227 -= 888
                                                                            }
                                                                            _ = v320
                                                                        case "\n\t":
                                                                            v34 -= (v57 >> ((v316) & 63))
                                                                        default:
                                                                            var v321 int8 = 1
                                                                            _ = v321
                                                                        }
                                                                    case "\n\t":
                                                                        switch v317 {
                                                                        case "a":
                                                                            var v322 uint32 = (v2 | genFn0(42, true, -2147483648))
                                                                            _ = v322
                                                                        case "":
                                                                            if len(v194) >= 2 {
                                                                                v194 = v194[1:len(v194)]
                                                                            }
                                                                        case "foo bar":
                                                                            var v323 int32 = -1
                                                                            _ = v323
                                                                        default:
                                                                            v4[2] = v5
                                                                        }
                                                                    }
                                                                    _ = len(v4)
                                                                } else {
                                                                    if v5 {
                                                                        v324 := GenS20{F0: 1, F1: 125}
                                                                        _ = v324
                                                                    }
                                                                }
                                                                v3 += 32767
                                                                var v325 bool = v5
                                                                fmt.Println(v199, v2, v325)
                                                                _ = v316
                                                                _ = v317
                                                                _ = v318
                                                                _ = v319
                                                                _ = v325
                                                            }
                                                        }
                                                    case 128:
                                                        var v326 string = ""
                                                        _ = v326
                                                    }
                                                    v57 -= uint(v3)
                                                    fmt.Println(v199)
                                                }
                                            case -4:
                                                fmt.Println(v227, fi58)
                                            }
                                            v34 += v57
                                            fmt.Println(v34)
                                        }
                                        v5 = true
                                        fmt.Println(fi192)
                                        _ = v227
                                        _ = v228
                                        _ = v229
                                    }
                                    v2 = (v2 / ((v2) | 1))
                                }
                                var v327 int = -1
                                _ = v199
                                _ = v327
                            } else {
                                var v328 float32 = 1.5
                                switch "hello" {
                                case "foo bar":
                                    v329 := GenS41{F0: 889}
                                    _ = v329
                                case "hello":
                                    v2 -= (v2 / ((314) | 1))
                                case "":
                                    v330 := v1[genFn0(42, true, 306)]
                                    _ = v330
                                }
                                var v331 uint32 = genFn0(-128, (!v5), 42)
                                v3 += v3
                                _ = v328
                                _ = v331
                            }
                        default:
                            v332 := GenS13{F0: 4294967295, F1: -1}
                            _ = v332
                        }
                    }
                    _ = v194
                case "foo bar":
                    fmt.Println(v193, v57)
                case "":
                    if (-1 >= 9223372036854775807) {
                        fmt.Println(fi192)
                        if v5 {
                            fmt.Println(v2, v57)
                        } else {
                            v1[genFn0(116, ((0 == 599) && v5), -1)] = (fi58 / ((2147483647) | 1))
                            _ = len(v6)
                        }
                    }
                }
                _ = fi192
                _ = v193
            }
        case 1:
            var v333 float32 = -2.25
            _ = v333
        default:
            delete(v1, 65535)
        }
        for fi334 := 0; fi334 < 13; fi334++ {
            switch fi334 {
            case 2147483647:
                var v335 uint16 = (v3 << ((4294967295) & 63))
                _ = v335
            case 32767:
                v336 := []int{fi334, fi58, int(v3)}
                v2 = genFn0(-58, (32767 <= -504), 1)
                if v5 {
                    v34 = 4294967295
                    v5 = true
                } else {
                    var v337 bool = (false || v5)
                    _ = v337
                }
                _ = v336
            case 172:
                v338 := []int{(fi58 % (((-(fi334))) | 1)), fi334, (^(fi58))}
                _ = v338
            default:
                v339 := GenS8{F0: 0, F1: 0, F2: (("a" + "hello") + ("a" + "\n\t"))}
                _ = v339
            }
            switch "" {
            case "\n\t":
                v2 -= 65535
                fmt.Println(fi58, fi334)
                if v5 {
                    v340 := GenS42{F0: (v3 % ((v3) | 1)), F1: v34}
                    v2 = genFn0(-128, true, 42)
                    v3 += v3
                    _ = v340
                } else {
                    if (3.14159 <= (0.1 / 1.5)) {
                        if (float32(0.1) > (0.1 - (((0.0 - 0.1) - 0.1) + 1e10))) {
                            var v341 uint = (v34 + v34)
                            v34 -= (v341 % ((639) | 1))
                            v342 := []int{}
                            v343 := GenS43{F0: -1}
                            fmt.Println(v5, v341)
                            _ = v341
                            _ = v342
                            _ = v343
                        } else {
                            if ("foo bar" == "foo bar") {
                                var v344 float32 = (1e10 - 1e10)
                                v345 := GenS44{F0: (v3 & v3)}
                                switch "foo bar" {
                                case "a":
                                    v346 := GenS14{F0: v2}
                                    _ = v346
                                case "\n\t":
                                    if true {
                                        if true {
                                            switch ("foo bar" + ("hello" + "foo bar")) {
                                            case "\n\t":
                                                if ((638 > 18446744073709551615) && (-9223372036854775808 < 9223372036854775807)) {
                                                    switch (fi58 * fi334) {
                                                    case 4294967295:
                                                        v34 = 217
                                                    case -725:
                                                        _ = len(v1)
                                                    default:
                                                        v5 = (((1.5 + (3.14159 * (-2.25 / -2.25))) + (0.0 * 1.5)) == float64(3.14159))
                                                    }
                                                    var v347 string = ("hello" + (("foo bar" + "foo bar") + "hello"))
                                                    _ = v347
                                                }
                                            case "foo bar":
                                                v3 += 32767
                                            case "":
                                                v3 += 525
                                            }
                                            if v5 {
                                                v348 := []int{(fi58 << ((0) & 63)), -9223372036854775808}
                                                _ = v348
                                            } else {
                                                var v349 float64 = -2.25
                                                _ = v349
                                            }
                                            v350 := []int{fi334}
                                            _ = v350
                                        }
                                        v351 := GenS30{F0: 82, F1: v5, F2: 127}
                                        _ = v351
                                    }
                                }
                                _ = v344
                                _ = v345
                            } else {
                                fmt.Println(v5, v5, v5)
                            }
                            var v352 uint = (^(v57))
                            _ = v352
                        }
                        var v353 int64 = -1
                        _ = v353
                    } else {
                        v3 = v3
                        v34 += (^(v34))
                        v3 = (v3 & (v3 << ((v57) & 63)))
                    }
                    switch fi58 {
                    case 439:
                        fmt.Println(fi334, v5)
                    case -369:
                        v354 := []int{(fi334 | 1), fi58}
                        _ = v354
                    }
                    if v5 {
                        if v5 {
                            v355 := GenS45{F0: -22, F1: v2, F2: 603}
                            if ((v2 + v2) == 0) {
                                var v356 uint32 = 969
                                if (v57 >= (v57 * v57)) {
                                    var v357 float64 = (3.14159 / 1.0)
                                    _ = v357
                                }
                                _ = v356
                            } else {
                                var v358 float32 = 1e10
                                _ = v358
                            }
                            _ = v355
                        }
                        v5 = (v5 || true)
                    } else {
                        fmt.Println(v57)
                    }
                }
                fmt.Println(fi334, fi58)
                v2 += uint32(v57)
            case "a":
                v5 = v5
            case "foo bar":
                v34 -= v34
            }
            v359 := v6[("a" + ((("\n\t" + ("a" + "")) + "a") + "hello"))]
            _ = fi334
            _ = v359
        }
        _ = fi58
        _ = v59
    }
    v360 := []int{9223372036854775807, 128, 32768, 32767}
    var v361 int16 = -1
    if false {
        _ = len(v4)
        if (v5 && v5) {
            v362 := v6["\n\t"]
            switch 1 {
            case 4294967295:
                var v363 string = ((("foo bar" + ((("" + "foo bar") + "hello") + "foo bar")) + "hello") + "")
                _ = v363
            case -539:
                v364 := v1[(^(v2))]
                _ = v364
            case 2147483647:
                _ = len(v4)
            }
            _ = v362
        }
        if false {
            for fi365 := 0; fi365 < 8; fi365++ {
                v57 = 4294967295
                var v366 int64 = -90
                if (-7 >= fi365) {
                    switch "" {
                    case "\n\t":
                        switch -9223372036854775808 {
                        case -1:
                            v4[(v2 ^ 194)] = (uint8(v3) < 42)
                        case -9223372036854775808:
                            var v367 float32 = 0.1
                            _ = v367
                        }
                    }
                    switch "a" {
                    case "":
                        for fi368 := 0; fi368 < 10; fi368++ {
                            v361 -= v361
                            v5 = true
                            _ = fi368
                        }
                    case "foo bar":
                        var v369 int = (^(fi365))
                        _ = v369
                    case "hello":
                        switch int(v3) {
                        case 9223372036854775807:
                            for fi370 := 0; fi370 < 16; fi370++ {
                                v371 := GenS5{F0: 32767, F1: 0, F2: (v57 << ((v34) & 63))}
                                var v372 int16 = v361
                                _ = fi370
                                _ = v371
                                _ = v372
                            }
                        case 127:
                            fmt.Println(v3, v5)
                        case 2147483647:
                            for fi373 := 0; fi373 < 13; fi373++ {
                                v374 := GenS35{F0: 1, F1: v3, F2: -1}
                                fmt.Println(v5, fi365)
                                v375 := GenS22{F0: v34}
                                _ = fi373
                                _ = v374
                                _ = v375
                            }
                        default:
                            switch fi365 {
                            case -51:
                                var v376 uint8 = 255
                                _ = v376
                            }
                        }
                    }
                    switch "" {
                    case "a":
                        if true {
                            v377 := v6["foo bar"]
                            _ = v377
                        }
                    case "":
                        delete(v6, ("hello" + ("\n\t" + "a")))
                    case "foo bar":
                        if (!(("\n\t" + "") < "hello")) {
                            fmt.Println(v2, v366, v57)
                            var v378 float64 = ((3.14159 - (0.1 + 3.14159)) / 1e10)
                            _ = v378
                        } else {
                            v379 := GenS24{F0: 1}
                            delete(v4, v2)
                            for fi380 := 0; fi380 < 18; fi380++ {
                                v381 := GenS46{F0: 73, F1: v361, F2: v2}
                                _ = fi380
                                _ = v381
                            }
                            _ = v379
                        }
                    default:
                        v34 += v34
                    }
                }
                _ = fi365
                _ = v366
            }
            _ = len(v360)
        } else {
            v360 = append(v360, -857)
            v34 += (^(v57))
            switch "hello" {
            case "a":
                fmt.Println(v5, v5, v3)
            case "foo bar":
                switch ("" + "\n\t") {
                case "foo bar":
                    _ = len(v1)
                }
            case "\n\t":
                v360 = append(v360, 32768)
            }
        }
    } else {
        var v382 uint8 = 70
        v383 := GenS18{F0: ("hello" + "\n\t"), F1: (v361 % (((v361 << ((v34) & 63))) | 1)), F2: -883}
        _ = v382
        _ = v383
    }
    if (!false) {
        switch "foo bar" {
        case "hello":
            var v384 uint32 = (v2 - v2)
            _ = v384
        case "a":
            var v385 float64 = 1.5
            _ = v385
        case "foo bar":
            fmt.Println(v34, v57, v34)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v360:", len(v360), v360)
    fmt.Println("v361:", v361)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v360
    _ = v361
    _ = v4
    _ = v5
    _ = v57
    _ = v6
}
