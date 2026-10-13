package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 int8; F1 uint16 }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int; F1 int16; F2 uint8 }
type GenS12 struct { F0 uint32; F1 uint }
type GenS13 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS14 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS15 struct { F0 bool; F1 int32; F2 bool }
func genFn0(p0 int) int {
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 uint = 479
        var v8 uint8 = 0
        var v9 uint16 = uint16(v8)
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := []int{127, (p0 ^ p0)}
    p0 += p0
    var v11 string = ("\n\t" + "")
    if len(v10) > 0 {
        v10[0] = p0
    }
    _ = p0
    _ = v10
    _ = v11
    return p0
}
func genFn1() uint32 {
    var v12 int32 = 1
    if true {
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 uint64 = 42
            v14 = (v14 * 358)
            _ = fi13
            _ = v14
        }
    } else {
        fmt.Println(v12)
        fmt.Println(v12)
    }
    switch genFn0(-1) {
    case 127:
        v12 -= v12
    case 414:
        v15 := GenS1{F0: 65535, F1: (v12 - (v12 >> ((1) & 63))), F2: 42}
        _ = v15
    case -1:
        v16 := []int{128, 0, -2147483648, 982}
        _ = v16
    default:
        for fi17 := 0; fi17 < 9; fi17++ {
            v12 = v12
            fmt.Println(v12)
            _ = fi17
        }
    }
    _ = v12
    return 2
}

func main() {
    var v1 uint = 42
    var v2 int64 = 1
    var v3 string = "foo bar"
    var v4 uint16 = 2
    v5 := make(map[string]int)
    if (32767 == v4) {
        v5[v3] = -1
    }
    if false {
        var v18 int32 = -509
        for fi19 := 0; fi19 < 13; fi19++ {
            v1 += v1
            _ = fi19
        }
        v20 := GenS1{F0: 0, F1: (v18 * (v18 | v18)), F2: 1}
        _ = v18
        _ = v20
    } else {
        var v21 uint = (v1 >> ((v1) & 63))
        if false {
            v1 = v21
            if (!((v1 ^ v1) != v21)) {
                if (v4 > 1) {
                    var v22 uint64 = 184
                    _ = v22
                }
                var v23 uint16 = v4
                fmt.Println(v23)
                _ = v23
            }
            if ("hello" >= v3) {
                _ = len(v5)
                if ((true && false) && (-1 == -1)) {
                    v24 := []int{}
                    var v25 int32 = 407
                    v3 = "\n\t"
                    _ = v24
                    _ = v25
                }
                for fi26 := 0; fi26 < 18; fi26++ {
                    switch 725 {
                    case -1:
                        if (0 > 42) {
                            v27 := GenS1{F0: 65535, F1: -594, F2: 1}
                            if true {
                                v2 += v2
                                v2 -= -518
                            }
                            _ = v27
                        }
                    case 9223372036854775807:
                        switch ("foo bar" + (v3 + ("\n\t" + v3))) {
                        case "":
                            if false {
                                if false {
                                    if (!false) {
                                        var v28 uint32 = genFn1()
                                        if (v2 > -9223372036854775808) {
                                            var v29 int64 = 1
                                            _ = v29
                                        }
                                        switch (fi26 + genFn0(fi26)) {
                                        case 0:
                                            var v30 float64 = -2.25
                                            _ = v30
                                        }
                                        _ = v28
                                    }
                                    v31 := GenS2{F0: 116}
                                    _ = v31
                                }
                                if (v3 >= ((v3 + (("" + v3) + ("a" + v3))) + v3)) {
                                    fmt.Println(fi26, fi26)
                                    var v32 uint8 = 255
                                    v21 = 383
                                    _ = v32
                                } else {
                                    for fi33 := 0; fi33 < 2; fi33++ {
                                        switch "hello" {
                                        case "foo bar":
                                            v34 := GenS3{F0: 2}
                                            _ = v34
                                        }
                                        var v35 float32 = 0.1
                                        var v36 float64 = float64(-2.25)
                                        _ = fi33
                                        _ = v35
                                        _ = v36
                                    }
                                    var v37 uint32 = 42
                                    v38 := []int{fi26, fi26, int(v21)}
                                    _ = v37
                                    _ = v38
                                }
                                for fi39 := 0; fi39 < 14; fi39++ {
                                    var v40 int8 = 55
                                    var v41 uint = (^(v21))
                                    var v42 int32 = -941
                                    _ = fi39
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                }
                            }
                        case "foo bar":
                            var v43 string = "\n\t"
                            _ = v43
                        default:
                            v44 := []int{0, fi26}
                            _ = v44
                        }
                    case -9223372036854775808:
                        fmt.Println(v1, v3, v2)
                    default:
                        fmt.Println(v3, v1)
                    }
                    _ = fi26
                }
            } else {
                fmt.Println(v2, v21, v2)
                var v45 int32 = 42
                var v46 int = 32768
                _ = v45
                _ = v46
            }
        } else {
            if true {
                v47 := GenS2{F0: 2}
                v2 -= v2
                if (!(313 > 65535)) {
                    if false {
                        v2 = v2
                    } else {
                        v1 = v21
                    }
                    var v48 int16 = -32768
                    v49 := GenS4{F0: (("a" + "\n\t") + v3), F1: 11, F2: 1}
                    _ = v48
                    _ = v49
                } else {
                    var v50 float32 = (((0.0 - 3.14159) + float32((1.5 / 1e10))) * 1e10)
                    for fi51 := 0; fi51 < 14; fi51++ {
                        v52 := GenS5{F0: (v21 ^ v21), F1: (!((v50 == v50) && true))}
                        var v53 uint64 = 275
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                    v1 = (v21 << (((v1 ^ (v1 & v1))) & 63))
                    _ = v50
                }
                _ = v47
            } else {
                v54 := GenS2{F0: 124}
                _ = v54
            }
            v3 = "\n\t"
        }
        _ = v21
    }
    v5[("foo bar" + ("a" + v3))] = 127
    if false {
        switch (v3 + "foo bar") {
        case "a":
            var v55 uint32 = genFn1()
            _ = v55
        case "foo bar":
            for fi56 := 0; fi56 < 1; fi56++ {
                for fi57 := 0; fi57 < 10; fi57++ {
                    var v58 bool = (false && true)
                    if v58 {
                        v59 := GenS6{F0: 586, F1: fi56}
                        if v58 {
                            if ((v4 | (v4 / (((v4 * v4)) | 1))) != 852) {
                                var v60 uint = v1
                                var v61 uint64 = 2
                                _ = v60
                                _ = v61
                            } else {
                                v3 = (v3 + v3)
                                if ((^(fi57)) == 42) {
                                    var v62 int8 = -17
                                    v58 = (1 == 42)
                                    var v63 uint16 = v4
                                    v4 -= v4
                                    _ = v62
                                    _ = v63
                                }
                                var v64 uint = 18446744073709551615
                                switch "" {
                                case "":
                                    v65 := v5[(("\n\t" + v3) + v3)]
                                    _ = v65
                                case "foo bar":
                                    delete(v5, (((v3 + "hello") + (v3 + (v3 + v3))) + ""))
                                case "hello":
                                    var v66 int16 = -1
                                    _ = v66
                                default:
                                    v58 = ((v4 == v4) || false)
                                }
                                if (v58 || (-355 != 42)) {
                                    if (2147483647 != 42) {
                                        if (!(-46 == 1)) {
                                            v67 := GenS7{F0: 127, F1: v4}
                                            _ = v67
                                        } else {
                                            fmt.Println(v4, v4, v58)
                                        }
                                        fmt.Println(fi56, v2)
                                        v1 += 415
                                    }
                                    if (v2 >= (v2 / ((v2) | 1))) {
                                        v3 = ("" + "foo bar")
                                        var v68 int = (^(fi57))
                                        switch "a" {
                                        case "a":
                                            switch (v68 ^ 0) {
                                            case 890:
                                                switch v3 {
                                                case "\n\t":
                                                    v69 := []int{-2147483648, fi56, (fi57 / ((genFn0((fi56 - (v68 | (fi57 + 919))))) | 1))}
                                                    _ = v69
                                                }
                                            case 42:
                                                switch v3 {
                                                case "\n\t":
                                                    fmt.Println(v68, v68, v58)
                                                case "":
                                                    if true {
                                                        v64 += v1
                                                        fmt.Println(fi56)
                                                    } else {
                                                        if v58 {
                                                            var v70 uint16 = 65535
                                                            var v71 int32 = -2147483648
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                        v72 := GenS3{F0: 2}
                                                        _ = v72
                                                    }
                                                case "hello":
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        var v73 float64 = -2.25
                                                        _ = v73
                                                    case "foo bar":
                                                        var v74 uint64 = 104
                                                        _ = v74
                                                    case "a":
                                                        if (42 >= 0) {
                                                            var v75 float32 = 1e10
                                                            _ = v75
                                                        } else {
                                                            if v58 {
                                                                if v58 {
                                                                    fmt.Println(v68, v2, v64)
                                                                    var v76 float32 = ((1.5 + (1.5 / 3.14159)) + 1e10)
                                                                    _ = v76
                                                                }
                                                                fmt.Println(fi56)
                                                                v3 = ((v3 + ((v3 + (((v3 + v3) + v3) + "\n\t")) + v3)) + (v3 + ((v3 + "a") + v3)))
                                                            }
                                                        }
                                                        switch v3 {
                                                        case "a":
                                                            v77 := GenS8{F0: genFn0((fi57 | fi56)), F1: 128}
                                                            _ = v77
                                                        case "foo bar":
                                                            v78 := GenS8{F0: v68, F1: 0}
                                                            _ = v78
                                                        case "\n\t":
                                                            fmt.Println(v4, fi56)
                                                        default:
                                                            var v79 float32 = 1.5
                                                            _ = v79
                                                        }
                                                        if true {
                                                            v58 = (0 >= 764)
                                                            v64 += v64
                                                            var v80 bool = ((("" + ("" + "\n\t")) + v3) > (v3 + (("hello" + v3) + v3)))
                                                            if (!(v58 || (901 < 18446744073709551615))) {
                                                                if ((1e10 - 1e10) != 1.5) {
                                                                    v2 = v2
                                                                }
                                                                v81 := []int{128, (v68 / (((fi56 >> ((18446744073709551615) & 63))) | 1))}
                                                                _ = v81
                                                            }
                                                            if v80 {
                                                                var v82 float32 = 0.0
                                                                _ = v82
                                                            }
                                                            _ = v80
                                                        } else {
                                                            v83 := GenS4{F0: "\n\t", F1: 124, F2: 2}
                                                            _ = v83
                                                        }
                                                    default:
                                                        v58 = (v58 && v58)
                                                        switch "hello" {
                                                        case "a":
                                                            if (("" + v3) >= v3) {
                                                                if (!(1 <= -162)) {
                                                                    v84 := GenS9{F0: v58, F1: 2, F2: 150}
                                                                    _ = v84
                                                                }
                                                            } else {
                                                                var v85 int16 = -32768
                                                                v3 = (v3 + ((v3 + v3) + v3))
                                                                _ = v85
                                                            }
                                                        case "foo bar":
                                                            var v86 uint32 = 0
                                                            _ = v86
                                                        case "hello":
                                                            if true {
                                                                if v58 {
                                                                    fmt.Println(fi56, v64)
                                                                    v68 += 32768
                                                                }
                                                                v87 := GenS10{F0: -357}
                                                                var v88 int8 = 1
                                                                if v58 {
                                                                    switch v3 {
                                                                    case "\n\t":
                                                                        var v89 int16 = 42
                                                                        _ = v89
                                                                    case "foo bar":
                                                                        v64 -= (v64 * 255)
                                                                    }
                                                                    if v58 {
                                                                        v3 = (("hello" + (("\n\t" + "") + ((v3 + "a") + v3))) + v3)
                                                                        v2 = (v2 << ((255) & 63))
                                                                        var v90 int = fi56
                                                                        _ = v90
                                                                    }
                                                                    var v91 int8 = v88
                                                                    if v58 {
                                                                        var v92 uint16 = v4
                                                                        _ = v92
                                                                    } else {
                                                                        v64 -= 217
                                                                        var v93 float32 = 1.5
                                                                        var v94 float64 = 0.0
                                                                        _ = v93
                                                                        _ = v94
                                                                    }
                                                                    _ = v91
                                                                }
                                                                _ = v87
                                                                _ = v88
                                                            }
                                                        default:
                                                            var v95 uint16 = (v4 ^ 42)
                                                            switch 42 {
                                                            case -1:
                                                                switch v3 {
                                                                case "a":
                                                                    var v96 uint32 = 2147483647
                                                                    _ = v96
                                                                default:
                                                                    if (-128 != 1) {
                                                                        v4 += v95
                                                                        v2 += v2
                                                                        switch v3 {
                                                                        case "hello":
                                                                            var v97 uint = 1
                                                                            _ = v97
                                                                        }
                                                                    }
                                                                }
                                                            case 42:
                                                                if false {
                                                                    v98 := GenS11{F0: (fi56 ^ fi56), F1: 1, F2: 0}
                                                                    var v99 float32 = (0.1 * 1.5)
                                                                    if ((v58 && (!(v58 || true))) && ("\n\t" > (v3 + (v3 + v3)))) {
                                                                        if (v58 || (v3 <= (v3 + ""))) {
                                                                            switch int(v4) {
                                                                            case -1:
                                                                                v68 -= (fi57 >> ((v1) & 63))
                                                                            }
                                                                            v100 := GenS12{F0: 42, F1: (v64 * v1)}
                                                                            if ((^(v4)) < v95) {
                                                                                var v101 int = int(v64)
                                                                                var v102 float64 = 0.1
                                                                                if v58 {
                                                                                    if ((((v3 + (v3 + (v3 + ("foo bar" + "a")))) + "\n\t") <= (("\n\t" + v3) + "foo bar")) || ((!v58) || (797 >= 183))) {
                                                                                        fmt.Println(v99)
                                                                                        var v103 int16 = 1000
                                                                                        var v104 int64 = int64(v64)
                                                                                        _ = v103
                                                                                        _ = v104
                                                                                    } else {
                                                                                        switch int(v2) {
                                                                                        case 110:
                                                                                            v105 := []int{(fi56 ^ fi56), 843, 2147483647, 454}
                                                                                            v106 := GenS13{F0: 456, F1: v4, F2: (v2 * (v2 | -41))}
                                                                                            _ = v105
                                                                                            _ = v106
                                                                                        }
                                                                                        v107 := GenS12{F0: 2, F1: 65535}
                                                                                        var v108 int64 = v2
                                                                                        _ = v107
                                                                                        _ = v108
                                                                                    }
                                                                                    v58 = false
                                                                                }
                                                                                _ = v101
                                                                                _ = v102
                                                                            } else {
                                                                                v109 := GenS14{F0: -1, F1: genFn1(), F2: genFn1()}
                                                                                _ = v109
                                                                            }
                                                                            _ = v100
                                                                        } else {
                                                                            var v110 uint8 = 0
                                                                            _ = v110
                                                                        }
                                                                        var v111 uint8 = 1
                                                                        delete(v5, ("" + v3))
                                                                        _ = v111
                                                                    }
                                                                    var v112 uint64 = uint64(v1)
                                                                    _ = v112
                                                                    _ = v98
                                                                    _ = v99
                                                                }
                                                            case -755:
                                                                var v113 uint64 = 1
                                                                var v114 int32 = 2147483647
                                                                _ = v113
                                                                _ = v114
                                                            default:
                                                                v115 := []int{(-(v68)), (fi57 & fi56), (fi57 & -9223372036854775808)}
                                                                _ = v115
                                                            }
                                                            v58 = true
                                                            fmt.Println(v4)
                                                            v4 += (v4 % ((v4) | 1))
                                                            _ = v95
                                                        }
                                                        switch v3 {
                                                        case "hello":
                                                            v116 := []int{(-(v68)), (fi57 << (((v1 >> ((v64) & 63))) & 63)), (fi57 % ((fi57) | 1)), (fi56 % (((fi57 | (^(v68)))) | 1))}
                                                            v117 := GenS9{F0: v58, F1: 51, F2: 4294967295}
                                                            v3 = v3
                                                            _ = v116
                                                            _ = v117
                                                        case "":
                                                            v64 += 18446744073709551615
                                                        }
                                                        v118 := GenS13{F0: 626, F1: 2, F2: -334}
                                                        var v119 uint16 = 282
                                                        v120 := []int{}
                                                        _ = v118
                                                        _ = v119
                                                        _ = v120
                                                    }
                                                    v121 := []int{(fi56 / ((fi56) | 1)), v68, genFn0(fi57), (fi57 | v68)}
                                                    _ = v121
                                                default:
                                                    var v122 bool = (1 <= 1)
                                                    _ = v122
                                                }
                                            default:
                                                v2 += (v2 & v2)
                                            }
                                        case "":
                                            v4 -= v4
                                        }
                                        switch (v3 + "hello") {
                                        case "foo bar":
                                            var v123 uint = 652
                                            _ = v123
                                        case "":
                                            fmt.Println(v4)
                                            var v124 uint16 = 0
                                            v4 += (^(v124))
                                            _ = v124
                                        default:
                                            var v125 uint64 = 42
                                            _ = v125
                                        }
                                        _ = v68
                                    } else {
                                        v1 += 1
                                    }
                                    v126 := []int{(fi56 >> ((425) & 63))}
                                    _ = v126
                                } else {
                                    var v127 uint64 = 42
                                    var v128 uint = (v1 | v1)
                                    if v58 {
                                        switch (fi56 + 42) {
                                        case -9223372036854775808:
                                            v3 = "a"
                                        case 1:
                                            if (true || (49 > 128)) {
                                                v3 = v3
                                            }
                                        case -1:
                                            v129 := GenS10{F0: 0}
                                            _ = v129
                                        default:
                                            delete(v5, v3)
                                            switch genFn0((fi57 * -9223372036854775808)) {
                                            case 9223372036854775807:
                                                switch genFn0(fi57) {
                                                case 1:
                                                    switch (fi57 * (fi57 >> (((^(v64))) & 63))) {
                                                    case 190:
                                                        v58 = (v58 || false)
                                                    }
                                                }
                                            case -423:
                                                v130 := []int{genFn0(fi56), (fi56 + (-(fi57)))}
                                                _ = v130
                                            default:
                                                switch "a" {
                                                case "foo bar":
                                                    var v131 uint16 = v4
                                                    _ = v131
                                                }
                                            }
                                            switch (fi56 + int(v4)) {
                                            case 2147483647:
                                                v128 += v1
                                            case -1:
                                                v132 := GenS15{F0: v58, F1: 0, F2: v58}
                                                _ = v132
                                            }
                                        }
                                        v4 += v4
                                        fmt.Println(v1, v3, fi57)
                                    } else {
                                        fmt.Println(v1)
                                        v133 := []int{(fi57 % ((1) | 1)), fi57, fi56, genFn0(127)}
                                        _ = v133
                                    }
                                    _ = v127
                                    _ = v128
                                }
                                var v134 uint64 = 1
                                var v135 bool = (((float64((1e10 * 3.14159)) - (float64(((1.5 * 0.1) - float32((0.1 - 0.1)))) + 0.1)) > -2.25) || (100 < 47))
                                v136 := GenS4{F0: v3, F1: 42, F2: 237}
                                _ = v134
                                _ = v135
                                _ = v136
                                _ = v64
                            }
                            var v137 uint16 = v4
                            _ = v137
                        }
                        var v138 string = "a"
                        _ = v138
                        _ = v59
                    } else {
                        if v58 {
                            var v139 string = ""
                            v2 = int64(v4)
                            fmt.Println(fi57, fi56, v58)
                            fmt.Println(v139)
                            var v140 int32 = -1
                            _ = v139
                            _ = v140
                        } else {
                            v141 := GenS1{F0: 2147483648, F1: 88, F2: uint64(v4)}
                            v4 = v4
                            v1 = v1
                            _ = v141
                        }
                    }
                    if true {
                        v142 := GenS6{F0: 0, F1: 438}
                        _ = v142
                    }
                    _ = fi57
                    _ = v58
                }
                if true {
                    var v143 float64 = (3.14159 + (0.0 * 0.0))
                    v144 := []int{}
                    fmt.Println(v143, fi56, v2)
                    _ = v143
                    _ = v144
                } else {
                    var v145 uint8 = 1
                    if ((v145 - (^(v145))) > (^(v145))) {
                        switch genFn0((fi56 >> (((v1 << ((v1) & 63))) & 63))) {
                        case 14:
                            fmt.Println(v4)
                        }
                        var v146 uint = v1
                        _ = v146
                    } else {
                        fmt.Println(fi56)
                        var v147 int8 = 37
                        _ = v147
                    }
                    _ = v145
                }
                _ = fi56
            }
        }
        switch 65 {
        case -507:
            var v148 float32 = (((1e10 / 1e10) * 1e10) + 1e10)
            _ = v148
        case -225:
            switch v3 {
            case "":
                v4 -= 42
            case "\n\t":
                if (false && true) {
                    for fi149 := 0; fi149 < 4; fi149++ {
                        if (2 > 18446744073709551615) {
                            v3 = v3
                            var v150 uint16 = 643
                            _ = v150
                        }
                        _ = fi149
                    }
                } else {
                    v151 := []int{127, 4294967295, 42, 520}
                    if false {
                        var v152 float32 = 1.5
                        v4 = (^(v4))
                        v4 = uint16(v1)
                        _ = v152
                    }
                    _ = v151
                }
            case "foo bar":
                v1 = v1
            default:
                v2 = v2
            }
        case -1:
            _ = len(v5)
        }
    } else {
        _ = len(v5)
    }
    _ = len(v5)
    v153 := []int{}
    var v154 int = -1
    var v155 uint32 = 65535
    var v156 bool = (genFn0(int(v155)) == -847)
    v1 = v1
    v3 = "a"
    var v157 int64 = (v2 & (v2 >> (((v1 * (v1 / ((v1) | 1)))) & 63)))
    v4 = v4
    v158 := GenS8{F0: -336, F1: 128}
    var v159 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v153:", len(v153), v153)
    fmt.Println("v154:", v154)
    fmt.Println("v155:", v155)
    fmt.Println("v156:", v156)
    fmt.Println("v157:", v157)
    fmt.Println("v159:", v159)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v153
    _ = v154
    _ = v155
    _ = v156
    _ = v157
    _ = v158
    _ = v159
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
