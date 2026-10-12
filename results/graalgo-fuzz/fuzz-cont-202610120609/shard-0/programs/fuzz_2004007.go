package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 int32; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 uint; F2 int64 }
type GenS5 struct { F0 uint8; F1 int64 }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS8 struct { F0 string; F1 int8 }
type GenS9 struct { F0 int; F1 int; F2 uint64 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int32; F1 int16 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint16; F1 uint32 }
type GenS14 struct { F0 int16; F1 int16 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS17 struct { F0 string }
func genFn0() uint8 {
    v6 := GenS1{F0: 156}
    var v7 bool = false
    v8 := []int{}
    v7 = v7
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint32, p1 uint64, p2 bool) uint16 {
    if (true || p2) {
        var v9 uint64 = 42
        fmt.Println(p1, p2)
        _ = v9
    } else {
        v10 := GenS2{F0: 617, F1: -1, F2: 0}
        _ = v10
    }
    v11 := []int{2147483647, 468, 1}
    var v12 int64 = 9223372036854775807
    if p2 {
        for fi13 := 0; fi13 < 20; fi13++ {
            v11 = append(v11, (^(fi13)))
            v14 := GenS1{F0: genFn0()}
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = fi13
            _ = v14
        }
    } else {
        fmt.Println(p0)
        if p2 {
            fmt.Println(v12, p0, p2)
        }
        switch 0 {
        case 127:
            fmt.Println(p2, p0)
        case 9223372036854775807:
            switch 128 {
            case -9223372036854775808:
                for fi15 := 0; fi15 < 20; fi15++ {
                    switch ((((((("hello" + (("\n\t" + "a") + "\n\t")) + "") + "hello") + ("\n\t" + "hello")) + ("hello" + "")) + ("a" + "hello")) + "") {
                    case "hello":
                        var v16 bool = false
                        _ = v16
                    case "":
                        p1 += 2
                    case "foo bar":
                        for fi17 := 0; fi17 < 11; fi17++ {
                            var v18 string = ("hello" + "hello")
                            var v19 uint64 = 4294967295
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            var v20 int32 = -686
                            if (255 >= 106) {
                                v21 := GenS2{F0: 199, F1: (v20 & (v20 / ((-2147483648) | 1))), F2: (fi15 ^ -824)}
                                if p2 {
                                    if true {
                                        if (4294967295 >= 82) {
                                            p0 -= (p0 * 2147483647)
                                            v22 := GenS2{F0: (^(p1)), F1: (v20 << ((719) & 63)), F2: (fi15 << ((833) & 63))}
                                            var v23 int16 = 32767
                                            var v24 string = v18
                                            _ = v22
                                            _ = v23
                                            _ = v24
                                        }
                                        if (0 == 0) {
                                            if len(v11) > 0 {
                                                v11[0] = (fi15 - (fi15 * (fi17 << ((255) & 63))))
                                            }
                                            var v25 int = 808
                                            _ = v25
                                        }
                                        v26 := GenS3{F0: 150}
                                        _ = v26
                                    } else {
                                        fmt.Println(v20, fi15, fi17)
                                        v20 = v20
                                        v27 := GenS4{F0: 42, F1: 273, F2: v12}
                                        _ = v27
                                    }
                                    v19 = 18446744073709551615
                                }
                                fmt.Println(fi15, p2)
                                v28 := GenS3{F0: 32}
                                _ = v21
                                _ = v28
                            } else {
                                _ = len(v11)
                                v29 := GenS2{F0: v19, F1: -110, F2: 1}
                                _ = v29
                            }
                            fmt.Println(v18, v19)
                            if (0 == v12) {
                                var v30 float32 = 0.0
                                _ = v30
                            } else {
                                var v31 int16 = 0
                                _ = v31
                            }
                            v20 = (-(v20))
                            fmt.Println(v20, v18, p2)
                            v32 := GenS2{F0: v19, F1: 1, F2: fi15}
                            _ = fi17
                            _ = v18
                            _ = v19
                            _ = v20
                            _ = v32
                        }
                    }
                    v33 := GenS1{F0: genFn0()}
                    fmt.Println(v12, p1)
                    _ = fi15
                    _ = v33
                }
            case 9223372036854775807:
                for fi34 := 0; fi34 < 9; fi34++ {
                    p0 -= (p0 * p0)
                    _ = fi34
                }
            }
        }
    }
    for fi35 := 0; fi35 < 4; fi35++ {
        for fi36 := 0; fi36 < 15; fi36++ {
            var v37 bool = (!(!((p0 >> ((65535) & 63)) <= p0)))
            v11 = append(v11, (fi35 & -669))
            switch ("hello" + "a") {
            case "hello":
                v38 := GenS1{F0: 209}
                _ = v38
            case "foo bar":
                var v39 float64 = ((3.14159 - 1e10) + -2.25)
                _ = v39
            case "":
                if (1 > 931) {
                    p2 = (!false)
                } else {
                    fmt.Println(v12)
                    if true {
                        v40 := GenS1{F0: uint8(p1)}
                        _ = v40
                    }
                }
            default:
                switch "foo bar" {
                case "":
                    p0 = (^(p0))
                default:
                    v41 := GenS5{F0: 127, F1: v12}
                    _ = v41
                }
            }
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 255
}

func main() {
    var v1 int64 = 305
    var v2 uint8 = 2
    var v3 int32 = -1
    v4 := make(map[int]int)
    var v5 int32 = -97
    if ((0.1 + 3.14159) == 1.5) {
        for fi42 := 0; fi42 < 2; fi42++ {
            if (false && false) {
                switch (fi42 | fi42) {
                case 227:
                    v43 := GenS4{F0: (v2 | v2), F1: 4294967295, F2: (v1 / (((v1 & 1)) | 1))}
                    _ = v43
                default:
                    v1 -= v1
                }
                fmt.Println(fi42)
                v44 := GenS6{F0: (-(v1)), F1: v1}
                _ = v44
            }
            fmt.Println(v3, v3)
            _ = fi42
        }
    }
    fmt.Println(v2, v3, v2)
    v5 -= -6
    v4[42] = -1
    v1 = (^(v1))
    var v45 int = -9223372036854775808
    if (true && true) {
        v2 = (^(v2))
        if ((v2 & v2) >= v2) {
            for fi46 := 0; fi46 < 16; fi46++ {
                v4[v45] = (fi46 >> ((4294967295) & 63))
                fmt.Println(v5)
                _ = fi46
            }
            v2 = genFn0()
        }
        for fi47 := 0; fi47 < 7; fi47++ {
            fmt.Println(v1)
            if (746 <= -20) {
                switch "\n\t" {
                case "a":
                    for fi48 := 0; fi48 < 11; fi48++ {
                        if true {
                            if false {
                                v3 = (v5 << ((65535) & 63))
                                v49 := GenS7{F0: v2, F1: 32767, F2: v1}
                                fmt.Println(v45, v5, fi47)
                                if false {
                                    var v50 float64 = 1e10
                                    _ = v50
                                }
                                if (!true) {
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v51 uint32 = 245
                                        switch "\n\t" {
                                        case "a":
                                            var v52 int16 = -810
                                            _ = v52
                                        default:
                                            var v53 int8 = -100
                                            _ = v53
                                        }
                                        _ = v51
                                    case "hello":
                                        var v54 uint8 = 156
                                        _ = v54
                                    case "a":
                                        switch "" {
                                        case "a":
                                            v55 := GenS8{F0: "foo bar", F1: -1}
                                            _ = v55
                                        case "":
                                            var v56 uint64 = 42
                                            _ = v56
                                        case "hello":
                                            v57 := GenS3{F0: (v2 - (v2 >> ((4294967295) & 63)))}
                                            _ = v57
                                        }
                                    default:
                                        if ((0.1 >= 0.0) || (v1 != (v1 & v1))) {
                                            v1 += (v1 << ((255) & 63))
                                        } else {
                                            if true {
                                                var v58 string = "a"
                                                _ = v58
                                            } else {
                                                switch "" {
                                                case "":
                                                    var v59 int64 = (v1 * 11)
                                                    _ = v59
                                                case "hello":
                                                    switch (^(v45)) {
                                                    case 190:
                                                        var v60 uint16 = 0
                                                        _ = v60
                                                    default:
                                                        v61 := GenS5{F0: (^(v2)), F1: (v1 - v1)}
                                                        _ = v61
                                                    }
                                                case "\n\t":
                                                    var v62 int32 = 2147483647
                                                    _ = v62
                                                default:
                                                    if (!false) {
                                                        fmt.Println(v3)
                                                        var v63 int32 = v5
                                                        switch "a" {
                                                        case "hello":
                                                            v64 := GenS7{F0: genFn0(), F1: 1, F2: 962}
                                                            switch "" {
                                                            case "a":
                                                                if false {
                                                                    if false {
                                                                        fmt.Println(v3, fi47, v2)
                                                                        switch v45 {
                                                                        case 0:
                                                                            v65 := GenS8{F0: "a", F1: 2}
                                                                            _ = v65
                                                                        case -606:
                                                                            if (false || (1000 > 1000)) {
                                                                                v1 -= 42
                                                                                v66 := GenS9{F0: v45, F1: v45, F2: 1}
                                                                                _ = v66
                                                                            }
                                                                        default:
                                                                            var v67 uint8 = (v2 + v2)
                                                                            _ = v67
                                                                        }
                                                                        v68 := GenS2{F0: 4294967295, F1: -906, F2: fi48}
                                                                        _ = v68
                                                                    } else {
                                                                        fmt.Println(v2, fi47, v45)
                                                                        if true {
                                                                            var v69 uint = 4294967295
                                                                            _ = v69
                                                                        } else {
                                                                            v70 := GenS10{F0: v2}
                                                                            v3 = 2147483647
                                                                            v71 := v4[(-(fi48))]
                                                                            if true {
                                                                                v63 += v5
                                                                                fmt.Println(v3, v5, v63)
                                                                                switch (fi48 >> ((168) & 63)) {
                                                                                case -845:
                                                                                    v3 = v5
                                                                                case 9223372036854775807:
                                                                                    switch ("a" + "hello") {
                                                                                    case "\n\t":
                                                                                        var v72 uint32 = 42
                                                                                        _ = v72
                                                                                    case "hello":
                                                                                        v2 -= (^(v2))
                                                                                    default:
                                                                                        v3 += v63
                                                                                    }
                                                                                case -2147483648:
                                                                                    var v73 int32 = 1
                                                                                    _ = v73
                                                                                default:
                                                                                    fmt.Println(v2, v1)
                                                                                }
                                                                            } else {
                                                                                switch fi48 {
                                                                                case 9223372036854775807:
                                                                                    var v74 string = ""
                                                                                    _ = v74
                                                                                case 42:
                                                                                    var v75 float64 = (1.5 / 1.5)
                                                                                    _ = v75
                                                                                case 215:
                                                                                    var v76 uint32 = 2147483647
                                                                                    _ = v76
                                                                                default:
                                                                                    var v77 float64 = ((1e10 * -2.25) + 1.5)
                                                                                    _ = v77
                                                                                }
                                                                                v78 := GenS8{F0: "\n\t", F1: 0}
                                                                                switch v71 {
                                                                                case -64:
                                                                                    var v79 int32 = (^(v63))
                                                                                    _ = v79
                                                                                case 0:
                                                                                    v80 := GenS1{F0: v2}
                                                                                    _ = v80
                                                                                }
                                                                                _ = v78
                                                                            }
                                                                            if (!false) {
                                                                                v5 = v63
                                                                            } else {
                                                                                v5 = v63
                                                                                if (100 == 42) {
                                                                                    v3 = (v3 / (((^(v5))) | 1))
                                                                                    fmt.Println(fi48)
                                                                                }
                                                                            }
                                                                            _ = v70
                                                                            _ = v71
                                                                        }
                                                                        fmt.Println(v5, v45, v45)
                                                                    }
                                                                    v81 := GenS9{F0: (fi48 + fi48), F1: (v45 - (v45 & (fi47 / ((v45) | 1)))), F2: 780}
                                                                    fmt.Println(fi47, v5)
                                                                    _ = v81
                                                                }
                                                            default:
                                                                var v82 int16 = -1
                                                                _ = v82
                                                            }
                                                            _ = v64
                                                        case "a":
                                                            v83 := []int{(fi47 >> ((2) & 63)), (fi47 & (fi47 * v45)), (v45 << ((4294967295) & 63))}
                                                            _ = v83
                                                        default:
                                                            var v84 int32 = 2147483647
                                                            _ = v84
                                                        }
                                                        _ = v63
                                                    } else {
                                                        fmt.Println(v3, fi48)
                                                        var v85 uint8 = v2
                                                        if (95 <= 18446744073709551615) {
                                                            if (!false) {
                                                                switch (fi47 << ((579) & 63)) {
                                                                case 127:
                                                                    if true {
                                                                        v4[v45] = fi48
                                                                        delete(v4, fi48)
                                                                    } else {
                                                                        var v86 uint16 = 42
                                                                        _ = v86
                                                                    }
                                                                default:
                                                                    var v87 uint16 = 42
                                                                    _ = v87
                                                                }
                                                                if (!(((0.1 * 1.5) < float32((((0.0 / 1e10) / 1.0) * 0.0))) && (1e10 == 3.14159))) {
                                                                    var v88 int16 = 0
                                                                    _ = v88
                                                                } else {
                                                                    switch (fi47 ^ 42) {
                                                                    case 32767:
                                                                        switch ("" + (("a" + "foo bar") + "")) {
                                                                        case "foo bar":
                                                                            var v89 int16 = 745
                                                                            var v90 int = (v45 & fi47)
                                                                            _ = v89
                                                                            _ = v90
                                                                        case "a":
                                                                            var v91 bool = (!false)
                                                                            _ = v91
                                                                        case "hello":
                                                                            v3 = (v5 * v5)
                                                                        default:
                                                                            var v92 int16 = int16(v45)
                                                                            _ = v92
                                                                        }
                                                                    default:
                                                                        if (-2.25 > 1.5) {
                                                                            v45 -= (v45 & (fi48 & (v45 - (-(fi47)))))
                                                                            fmt.Println(fi48, v5)
                                                                            v93 := GenS11{F0: (v3 & int32(v1)), F1: 0}
                                                                            _ = v93
                                                                        }
                                                                    }
                                                                    switch (fi47 >> ((373) & 63)) {
                                                                    case 127:
                                                                        v94 := v4[v45]
                                                                        _ = v94
                                                                    case -9223372036854775808:
                                                                        delete(v4, (fi47 | fi48))
                                                                    case -2147483648:
                                                                        v5 += 148
                                                                    default:
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            var v95 int32 = 42
                                                                            _ = v95
                                                                        case "a":
                                                                            if (!((("" + "hello") + "foo bar") >= "hello")) {
                                                                                if (!true) {
                                                                                    var v96 int8 = -1
                                                                                    v2 = genFn0()
                                                                                    v97 := []int{}
                                                                                    _ = v96
                                                                                    _ = v97
                                                                                } else {
                                                                                    v98 := []int{}
                                                                                    v45 = (fi47 - v45)
                                                                                    fmt.Println(fi48, fi47)
                                                                                    _ = v98
                                                                                }
                                                                            } else {
                                                                                var v99 uint16 = 65535
                                                                                var v100 int32 = (v3 >> ((0) & 63))
                                                                                _ = v100
                                                                                _ = v99
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                switch ("hello" + (("" + "") + (((("" + "a") + "a") + "") + "foo bar"))) {
                                                                case "a":
                                                                    v101 := GenS6{F0: (v1 | (v1 + (v1 / (((v1 + (v1 / (((v1 & (v1 + -1))) | 1)))) | 1)))), F1: v1}
                                                                    _ = v101
                                                                case "hello":
                                                                    if true {
                                                                        var v102 int8 = -1
                                                                        var v103 int32 = v5
                                                                        if true {
                                                                            var v104 bool = true
                                                                            v105 := GenS6{F0: v1, F1: int64(v5)}
                                                                            _ = v104
                                                                            _ = v105
                                                                        }
                                                                        v45 -= (fi48 % ((fi48) | 1))
                                                                        _ = v102
                                                                        _ = v103
                                                                    }
                                                                case "foo bar":
                                                                    if (("hello" + "foo bar") == ("foo bar" + ("foo bar" + "a"))) {
                                                                        v3 = 2147483647
                                                                    } else {
                                                                        _ = len(v4)
                                                                        v106 := GenS12{F0: "foo bar"}
                                                                        _ = v106
                                                                    }
                                                                default:
                                                                    var v107 uint = 1
                                                                    _ = v107
                                                                }
                                                            }
                                                        } else {
                                                            v45 = 9223372036854775807
                                                            fmt.Println(v5, v1)
                                                        }
                                                        _ = v85
                                                    }
                                                }
                                                v108 := []int{}
                                                _ = v108
                                            }
                                            var v109 int16 = -183
                                            v4[0] = -2147483648
                                            _ = v109
                                        }
                                    }
                                    switch (v45 & (fi48 | (fi48 << ((42) & 63)))) {
                                    case 4294967295:
                                        v45 += (fi48 & v45)
                                    case 32767:
                                        var v110 float32 = 3.14159
                                        var v111 uint64 = 18446744073709551615
                                        _ = v110
                                        _ = v111
                                    case -378:
                                        fmt.Println(v1, v5)
                                    }
                                    switch int(v1) {
                                    case 2147483647:
                                        if true {
                                            fmt.Println(v45, fi47, v2)
                                            if (!true) {
                                                var v112 int64 = v1
                                                _ = v112
                                            }
                                        }
                                    case -2147483648:
                                        v113 := GenS8{F0: (("" + "") + "hello"), F1: -128}
                                        _ = v113
                                    case 452:
                                        v114 := []int{9223372036854775807}
                                        _ = v114
                                    }
                                    v115 := GenS13{F0: 255, F1: 48}
                                    v116 := GenS8{F0: "foo bar", F1: -1}
                                    _ = v115
                                    _ = v116
                                }
                                switch "" {
                                case "":
                                    switch "hello" {
                                    case "a":
                                        if true {
                                            v117 := GenS7{F0: v2, F1: 274, F2: v1}
                                            v118 := GenS4{F0: (v2 ^ genFn0()), F1: 42, F2: (-(v1))}
                                            _ = v117
                                            _ = v118
                                        } else {
                                            var v119 uint32 = 934
                                            _ = v119
                                        }
                                    case "\n\t":
                                        if (true || (true && true)) {
                                            fmt.Println(v45, v1, v3)
                                            if (false && false) {
                                                delete(v4, fi48)
                                            } else {
                                                var v120 bool = (genFn1(0, 173, ((true || (1 >= 271)) && (!false))) > genFn1(132, uint64(v2), true))
                                                _ = v120
                                            }
                                        } else {
                                            var v121 uint8 = v2
                                            _ = v121
                                        }
                                    }
                                case "\n\t":
                                    v122 := GenS5{F0: v2, F1: 265}
                                    fmt.Println(v45, fi48, v45)
                                    _ = v122
                                }
                                _ = v49
                            }
                        } else {
                            fmt.Println(v2, v1)
                            var v123 uint8 = 0
                            var v124 uint16 = 0
                            _ = v123
                            _ = v124
                        }
                        if true {
                            var v125 uint = 18446744073709551615
                            _ = v125
                        }
                        v2 += (^(v2))
                        _ = fi48
                    }
                case "foo bar":
                    for fi126 := 0; fi126 < 16; fi126++ {
                        var v127 float32 = -2.25
                        _ = fi126
                        _ = v127
                    }
                default:
                    v128 := GenS7{F0: v2, F1: 876, F2: v1}
                    _ = v128
                }
                if false {
                    v129 := GenS10{F0: (v2 ^ v2)}
                    v1 = 9223372036854775807
                    for fi130 := 0; fi130 < 16; fi130++ {
                        var v131 uint = 42
                        if true {
                            if (int8(v45) >= 122) {
                                if false {
                                    switch "" {
                                    case "hello":
                                        switch ("" + "") {
                                        case "hello":
                                            v132 := v4[fi47]
                                            _ = v132
                                        }
                                    case "":
                                        v133 := GenS14{F0: 0, F1: 476}
                                        _ = v133
                                    case "\n\t":
                                        fmt.Println(v45, v45, fi130)
                                    default:
                                        var v134 uint8 = genFn0()
                                        _ = v134
                                    }
                                    v135 := GenS2{F0: 2, F1: v3, F2: v45}
                                    switch 2147483647 {
                                    case 128:
                                        v136 := []int{(fi47 / ((fi130) | 1)), v45, (^(fi130)), (fi47 / ((fi130) | 1))}
                                        _ = v136
                                    default:
                                        v5 -= v5
                                    }
                                    _ = v135
                                }
                                v137 := GenS9{F0: fi130, F1: v45, F2: 0}
                                _ = v137
                            }
                            v138 := v4[v45]
                            var v139 uint16 = 65535
                            _ = v138
                            _ = v139
                        } else {
                            var v140 uint32 = 65535
                            if (!(v131 <= (v131 ^ 65535))) {
                                var v141 uint16 = 65535
                                v2 += (v2 | 255)
                                _ = v141
                            } else {
                                v140 = 42
                            }
                            _ = v140
                        }
                        _ = fi130
                        _ = v131
                    }
                    _ = v129
                }
                v45 += (fi47 >> ((417) & 63))
            }
            for fi142 := 0; fi142 < 8; fi142++ {
                v143 := GenS15{F0: 65535}
                v1 = 1
                fmt.Println(v45, v3)
                fmt.Println(fi47, fi142, v5)
                v144 := GenS16{F0: v2, F1: 488, F2: 497}
                fmt.Println(v5, v2, fi142)
                _ = fi142
                _ = v143
                _ = v144
            }
            _ = fi47
        }
    }
    v145 := []int{(v45 & 1), v45, v45, (v45 ^ v45)}
    if false {
        fmt.Println(v45, v45, v2)
        v146 := GenS17{F0: "hello"}
        _ = v146
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v145:", len(v145), v145)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v145
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}
