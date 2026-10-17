package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 uint; F2 uint8 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint64; F1 int; F2 uint32 }
type GenS5 struct { F0 int; F1 uint; F2 uint8 }
type GenS6 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS7 struct { F0 uint8; F1 string }
type GenS8 struct { F0 int64; F1 int16; F2 int8 }
type GenS9 struct { F0 uint64; F1 string }
type GenS10 struct { F0 int16; F1 int16 }
type GenS11 struct { F0 int8; F1 uint8; F2 int64 }
type GenS12 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS13 struct { F0 int16; F1 uint }
type GenS14 struct { F0 uint32; F1 uint8; F2 uint }
type GenS15 struct { F0 uint8; F1 int; F2 int8 }
type GenS16 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS17 struct { F0 uint16; F1 uint; F2 int16 }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 uint64; F1 int; F2 uint64 }
type GenS20 struct { F0 int; F1 uint32 }
type GenS21 struct { F0 int32; F1 uint8; F2 int16 }
type GenS22 struct { F0 uint64; F1 int; F2 string }
func genFn0() uint64 {
    var v7 float32 = 3.14159
    if true {
        v7 += float32(0.1)
        var v8 float64 = (float64((v7 * (1e10 * v7))) + 1e10)
        _ = v8
    } else {
        v9 := GenS1{F0: 1}
        var v10 bool = false
        if v10 {
            if (1000 > -1) {
                if ((4294967295 > 84) && v10) {
                    switch 9223372036854775807 {
                    case 32767:
                        v7 = v7
                    case 32768:
                        fmt.Println(v10, v7)
                    default:
                        switch "hello" {
                        case "":
                            switch 383 {
                            case 507:
                                switch -342 {
                                case -2147483648:
                                    v7 -= v7
                                case 1:
                                    if true {
                                        v7 = v7
                                        if v10 {
                                            var v11 int16 = 1
                                            var v12 uint64 = 937
                                            _ = v11
                                            _ = v12
                                        }
                                        if (-9223372036854775808 != 0) {
                                            v13 := GenS1{F0: 42}
                                            _ = v13
                                        }
                                    } else {
                                        fmt.Println(v7)
                                    }
                                default:
                                    var v14 bool = ((true || v10) && false)
                                    _ = v14
                                }
                            case 574:
                                var v15 string = ("" + "\n\t")
                                _ = v15
                            default:
                                fmt.Println(v10)
                            }
                        case "a":
                            switch "hello" {
                            case "\n\t":
                                v7 -= (-2.25 / v7)
                            case "":
                                fmt.Println(v7)
                            default:
                                if v10 {
                                    var v16 uint = 0
                                    var v17 uint32 = 128
                                    _ = v16
                                    _ = v17
                                } else {
                                    if v10 {
                                        v18 := []int{42, 32767}
                                        v7 -= (v7 - float32(((-2.25 - 3.14159) + 1e10)))
                                        if v10 {
                                            for fi19 := 0; fi19 < 11; fi19++ {
                                                var v20 uint64 = 525
                                                v21 := GenS2{F0: -2147483648, F1: 42, F2: 255}
                                                var v22 int16 = 42
                                                _ = fi19
                                                _ = v20
                                                _ = v21
                                                _ = v22
                                            }
                                            v23 := GenS2{F0: -686, F1: 65535, F2: 2}
                                            v24 := GenS3{F0: -1, F1: 42}
                                            _ = v23
                                            _ = v24
                                        }
                                        _ = v18
                                    }
                                    v10 = v10
                                    v25 := GenS2{F0: -1, F1: 18446744073709551615, F2: 1}
                                    _ = v25
                                }
                            }
                        }
                    }
                } else {
                    var v26 uint16 = 442
                    _ = v26
                }
                if (1e10 <= (3.14159 / ((float32(1e10) * ((((0.0 * (((((v7 / (0.1 - v7)) - (float32(0.0) / v7)) - v7) - v7) + (1.5 - 0.0))) / ((-2.25 * float32(3.14159)) - 1e10)) + ((v7 - v7) - float32(0.0))) / -2.25)) - v7))) {
                    v27 := []int{4294967295, -482, 1}
                    for fi28 := 0; fi28 < 3; fi28++ {
                        for fi29 := 0; fi29 < 6; fi29++ {
                            v10 = (!v10)
                            v30 := GenS4{F0: 4294967295, F1: (fi29 + fi29), F2: 741}
                            _ = fi29
                            _ = v30
                        }
                        _ = fi28
                    }
                    for fi31 := 0; fi31 < 15; fi31++ {
                        for fi32 := 0; fi32 < 5; fi32++ {
                            v33 := GenS2{F0: 2147483647, F1: 18446744073709551615, F2: 255}
                            v27 = append(v27, (fi32 + (fi31 >> ((255) & 63))))
                            _ = fi32
                            _ = v33
                        }
                        var v34 float64 = (0.0 - 0.0)
                        var v35 string = "foo bar"
                        _ = fi31
                        _ = v34
                        _ = v35
                    }
                    _ = v27
                }
            } else {
                var v36 float64 = float64(((1.5 + v7) + v7))
                v37 := GenS1{F0: 0}
                _ = v36
                _ = v37
            }
            var v38 float32 = v7
            _ = v38
        } else {
            v39 := GenS5{F0: -947, F1: 2, F2: 42}
            v7 += v7
            var v40 int32 = 866
            _ = v39
            _ = v40
        }
        _ = v10
        _ = v9
    }
    _ = v7
    return 2
}
func genFn1(p0 uint8, p1 int) uint8 {
    v41 := GenS5{F0: (p1 - p1), F1: 834, F2: 127}
    p0 = (p0 >> ((255) & 63))
    _ = p0
    _ = p1
    _ = v41
    return 255
}

func main() {
    var v1 uint8 = 255
    var v2 int = -9223372036854775808
    v3 := make(map[string]bool)
    var v4 uint = 259
    var v5 int8 = -128
    v6 := make(map[uint32]string)
    if true {
        v5 = (v5 + (-(v5)))
    }
    var v42 int16 = 0
    v43 := GenS5{F0: (v2 << ((v4) & 63)), F1: (v4 ^ v4), F2: v1}
    switch ("hello" + ("a" + "hello")) {
    case "":
        for fi44 := 0; fi44 < 16; fi44++ {
            for fi45 := 0; fi45 < 16; fi45++ {
                var v46 int32 = -438
                _ = fi45
                _ = v46
            }
            v47 := GenS2{F0: 283, F1: (^(v4)), F2: (v1 - uint8(v4))}
            v48 := GenS6{F0: int64(v42), F1: 667, F2: 56}
            _ = fi44
            _ = v47
            _ = v48
        }
    case "hello":
        v49 := []int{}
        _ = v49
    default:
        if true {
            if ((false || (!(42 == 58))) && (v5 <= (v5 & v5))) {
                v50 := []int{-2147483648, (v2 >> (((v4 % (((v4 + v4)) | 1))) & 63))}
                _ = v50
            }
            if ("\n\t" == "\n\t") {
                for fi51 := 0; fi51 < 3; fi51++ {
                    v1 = v1
                    v4 -= 197
                    _ = fi51
                }
                v52 := GenS7{F0: v1, F1: ("hello" + "a")}
                _ = v52
            }
            v53 := GenS6{F0: 9223372036854775807, F1: 2147483647, F2: 483}
            _ = v53
        } else {
            v1 = (v1 ^ (v1 + (^(v1))))
            for fi54 := 0; fi54 < 7; fi54++ {
                v5 += v5
                if (-129 <= (v2 >> ((18446744073709551615) & 63))) {
                    v1 = 127
                    v42 = v42
                } else {
                    v2 = fi54
                    v55 := []int{int(v4)}
                    _ = v55
                }
                v56 := GenS8{F0: 9223372036854775807, F1: 0, F2: (v5 << (((v4 % (((v4 - (v4 % (((^(v4))) | 1)))) | 1))) & 63))}
                _ = fi54
                _ = v56
            }
        }
    }
    var v57 int = (-(v2))
    var v58 string = "a"
    switch (v57 - 244) {
    case -128:
        var v59 int8 = v5
        _ = v59
    case 2147483647:
        var v60 uint64 = 18446744073709551615
        _ = v60
    default:
        var v61 uint16 = 32767
        _ = v61
    }
    switch (("foo bar" + ((v58 + v58) + "a")) + (v58 + "\n\t")) {
    case "\n\t":
        var v62 float64 = (1.5 + ((((-2.25 * (-2.25 + -2.25)) / 1.5) - 0.0) * 3.14159))
        _ = v62
    }
    for fi63 := 0; fi63 < 6; fi63++ {
        var v64 uint8 = v1
        if ("\n\t" < "\n\t") {
            if (2 > 228) {
                if true {
                    v57 -= v57
                }
                switch (v57 ^ (^(v2))) {
                case 638:
                    if false {
                        if true {
                            v65 := GenS7{F0: (v1 + (v1 / ((168) | 1))), F1: "a"}
                            for fi66 := 0; fi66 < 16; fi66++ {
                                if (9223372036854775807 == 42) {
                                    if (genFn1(255, (v57 / (((-(v57))) | 1))) <= 52) {
                                        var v67 int16 = v42
                                        v5 -= v5
                                        switch v58 {
                                        case "":
                                            if (true && true) {
                                                v68 := []int{(v57 - -947), (fi63 << (((v4 / ((601) | 1))) & 63)), v57, fi66}
                                                _ = v68
                                            } else {
                                                var v69 float32 = 3.14159
                                                _ = v69
                                            }
                                            var v70 float32 = 3.14159
                                            var v71 bool = true
                                            v4 = v4
                                            _ = v70
                                            _ = v71
                                        }
                                        var v72 int64 = -761
                                        _ = v67
                                        _ = v72
                                    }
                                    var v73 uint = v4
                                    var v74 bool = (true || false)
                                    v2 -= 224
                                    _ = v73
                                    _ = v74
                                }
                                if true {
                                    if ((v5 >> (((v4 + (v4 % (((v4 + (v4 / ((255) | 1)))) | 1)))) & 63)) == (v5 | 10)) {
                                        v5 = (^(v5))
                                        delete(v3, ("hello" + "hello"))
                                    } else {
                                        v1 = genFn1((v64 / ((2) | 1)), fi63)
                                        switch "\n\t" {
                                        case "\n\t":
                                            fmt.Println(v2)
                                        }
                                        delete(v6, 1)
                                    }
                                    switch v2 {
                                    case -952:
                                        var v75 int16 = (v42 << (((v4 << ((v4) & 63))) & 63))
                                        _ = v75
                                    case -1:
                                        fmt.Println(v2, v5, v5)
                                    case 128:
                                        var v76 uint64 = 0
                                        var v77 uint16 = 42
                                        _ = v76
                                        _ = v77
                                    }
                                } else {
                                    v1 = (v64 / ((255) | 1))
                                    if (v5 != (-(v5))) {
                                        v78 := GenS9{F0: 2, F1: ((("a" + (("" + v58) + v58)) + "a") + ((("foo bar" + "hello") + v58) + (v58 + v58)))}
                                        var v79 int = 2147483647
                                        v80 := v3["foo bar"]
                                        var v81 int8 = -128
                                        _ = v78
                                        _ = v79
                                        _ = v80
                                        _ = v81
                                    } else {
                                        v42 = v42
                                        v2 = (fi66 % ((fi66) | 1))
                                    }
                                    v82 := GenS10{F0: 0, F1: (-(v42))}
                                    _ = v82
                                }
                                fmt.Println(fi66)
                                _ = fi66
                            }
                            _ = v65
                        } else {
                            var v83 int32 = 0
                            _ = v83
                        }
                    }
                case 128:
                    v84 := GenS9{F0: 42, F1: v58}
                    _ = v84
                case -1:
                    v57 -= v57
                }
                v85 := []int{32767, (v57 << ((1) & 63))}
                _ = v85
            } else {
                v86 := GenS2{F0: 454, F1: (v4 | v4), F2: 1}
                _ = len(v6)
                _ = v86
            }
        } else {
            var v87 int32 = 1
            v87 += v87
            _ = v87
        }
        _ = fi63
        _ = v64
    }
    if (!((v4 * (^(v4))) <= v4)) {
        v88 := v3["hello"]
        v89 := []int{v2, v2, v2, v57}
        _ = v88
        _ = v89
    } else {
        fmt.Println(v2, v2)
    }
    fmt.Println(v5, v4)
    switch v57 {
    case 127:
        switch (v2 ^ v2) {
        case 157:
            v58 = "a"
        default:
            v57 = (v57 * v2)
        }
    case -2147483648:
        switch ("a" + v58) {
        case "\n\t":
            var v90 int = v2
            _ = v90
        case "a":
            var v91 uint64 = 18446744073709551615
            _ = v91
        case "hello":
            var v92 uint64 = 333
            _ = v92
        default:
            if false {
                var v93 uint64 = 0
                _ = v93
            } else {
                v94 := GenS9{F0: 42, F1: v58}
                for fi95 := 0; fi95 < 17; fi95++ {
                    if (v58 >= v58) {
                        var v96 int64 = -9223372036854775808
                        for fi97 := 0; fi97 < 6; fi97++ {
                            switch (v2 % ((int(v42)) | 1)) {
                            case -899:
                                v98 := GenS11{F0: (v5 + v5), F1: uint8(v4), F2: v96}
                                var v99 uint16 = 42
                                _ = v98
                                _ = v99
                            }
                            v57 = int(v96)
                            v100 := []int{128, 2147483647}
                            v101 := GenS12{F0: 0, F1: 65535, F2: v5}
                            var v102 uint32 = 920
                            _ = fi97
                            _ = v100
                            _ = v101
                            _ = v102
                        }
                        if false {
                            var v103 uint8 = 201
                            fmt.Println(v5, fi95, v1)
                            _ = v103
                        }
                        _ = v96
                    }
                    var v104 uint8 = 128
                    if (-1 == v5) {
                        for fi105 := 0; fi105 < 9; fi105++ {
                            var v106 string = "a"
                            _ = fi105
                            _ = v106
                        }
                        v5 = 100
                        switch (v58 + v58) {
                        case "hello":
                            v57 += v2
                        case "\n\t":
                            v107 := GenS12{F0: 1, F1: 2147483647, F2: (v5 | (^(v5)))}
                            _ = v107
                        default:
                            var v108 float64 = -2.25
                            _ = v108
                        }
                    }
                    _ = fi95
                    _ = v104
                }
                _ = v94
            }
        }
    case 32768:
        v109 := GenS1{F0: 634}
        _ = v109
    }
    v110 := GenS13{F0: v42, F1: 2}
    v1 += (v1 % (((v1 - v1)) | 1))
    v4 = (v4 - 0)
    if false {
        v111 := GenS2{F0: -2147483648, F1: (v4 % ((255) | 1)), F2: v1}
        if (-2.25 >= (1.5 - ((((float32(0.0) + 0.0) * 0.1) + 0.1) + -2.25))) {
            _ = len(v6)
            if false {
                v4 += (v4 / ((v4) | 1))
                fmt.Println(v2, v57)
            }
        } else {
            for fi112 := 0; fi112 < 12; fi112++ {
                v113 := GenS9{F0: 526, F1: v58}
                _ = fi112
                _ = v113
            }
        }
        v114 := GenS9{F0: genFn0(), F1: ((v58 + v58) + "foo bar")}
        _ = v111
        _ = v114
    } else {
        var v115 string = "a"
        _ = v115
    }
    fmt.Println(v1, v5)
    switch v57 {
    case 42:
        if false {
            if (!false) {
                v57 += v2
                if (true && (!true)) {
                    for fi116 := 0; fi116 < 7; fi116++ {
                        v117 := GenS14{F0: 969, F1: (v1 % ((uint8(v4)) | 1)), F2: 42}
                        _ = fi116
                        _ = v117
                    }
                }
            } else {
                v118 := []int{v57, (v57 >> (((v4 / (((v4 - v4)) | 1))) & 63)), (v2 << (((v4 & v4)) & 63)), v57}
                if ("\n\t" != v58) {
                    if (!(!true)) {
                        var v119 uint8 = 255
                        if ((0 == 749) && (42 != -806)) {
                            v120 := GenS15{F0: 0, F1: v57, F2: (-(v5))}
                            if len(v118) >= 2 {
                                v118 = v118[1:len(v118)]
                            }
                            v42 = v42
                            _ = v120
                        } else {
                            switch v2 {
                            case 2147483647:
                                for fi121 := 0; fi121 < 5; fi121++ {
                                    if false {
                                        var v122 float64 = ((-2.25 * (0.1 - 1e10)) + (float64(0.1) - -2.25))
                                        v42 = int16(v5)
                                        _ = v122
                                    }
                                    _ = fi121
                                }
                            }
                        }
                        v2 = (v57 & (v2 * v57))
                        _ = v119
                    } else {
                        if true {
                            var v123 int64 = 0
                            _ = v123
                        }
                        var v124 uint16 = 32767
                        v5 -= v5
                        _ = v124
                    }
                    for fi125 := 0; fi125 < 8; fi125++ {
                        switch ((("foo bar" + (v58 + "foo bar")) + "") + (v58 + (v58 + (("\n\t" + v58) + "\n\t")))) {
                        case "a":
                            for fi126 := 0; fi126 < 11; fi126++ {
                                v127 := GenS16{F0: 2147483648, F1: 42, F2: (v1 * (v1 ^ (v1 * (v1 >> ((v4) & 63)))))}
                                if false {
                                    v128 := GenS17{F0: 264, F1: (v4 & v4), F2: 0}
                                    _ = v128
                                } else {
                                    v57 += (-(v57))
                                    v4 += (v4 % ((v4) | 1))
                                    fmt.Println(v5, v1)
                                    switch (("\n\t" + ("hello" + v58)) + ((v58 + "foo bar") + "hello")) {
                                    case "\n\t":
                                        fmt.Println(v42, fi126, v4)
                                    case "hello":
                                        if false {
                                            if false {
                                                var v129 uint8 = (^(v1))
                                                v130 := GenS15{F0: 255, F1: (fi126 ^ (^(v2))), F2: v5}
                                                _ = len(v3)
                                                if false {
                                                    delete(v6, 65535)
                                                    if (0 >= 838) {
                                                        v129 -= genFn1(2, (^(fi125)))
                                                        fmt.Println(fi125)
                                                        v57 -= int(v1)
                                                    } else {
                                                        v4 = v4
                                                        v129 = v129
                                                    }
                                                    v131 := GenS18{F0: 1}
                                                    _ = v131
                                                }
                                                v1 = v129
                                                v58 = v58
                                                _ = v129
                                                _ = v130
                                            }
                                            v132 := GenS11{F0: -44, F1: (v1 % ((0) | 1)), F2: -9223372036854775808}
                                            v4 = (v4 << (((v4 / ((v4) | 1))) & 63))
                                            _ = v132
                                        } else {
                                            v57 = 127
                                            switch v58 {
                                            case "\n\t":
                                                fmt.Println(fi125, v5)
                                            default:
                                                v2 += fi125
                                            }
                                        }
                                    default:
                                        switch v58 {
                                        case "":
                                            var v133 float32 = (float32((3.14159 + (0.1 + -2.25))) + (1.5 / (0.1 + ((3.14159 * -2.25) - 1e10))))
                                            _ = v133
                                        default:
                                            v2 += 0
                                        }
                                    }
                                }
                                fmt.Println(v42)
                                _ = fi126
                                _ = v127
                            }
                        case "hello":
                            if (235 <= 12) {
                                var v134 float32 = 1e10
                                if true {
                                    for fi135 := 0; fi135 < 13; fi135++ {
                                        if true {
                                            switch -904 {
                                            case -234:
                                                fmt.Println(v42, v134, fi135)
                                            case -1:
                                                _ = len(v118)
                                            case -604:
                                                if false {
                                                    v58 = (v58 + v58)
                                                }
                                            default:
                                                if true {
                                                    v118 = append(v118, 32768)
                                                    switch (v2 ^ (v57 | fi125)) {
                                                    case 1:
                                                        v134 += (v134 + 0.0)
                                                    default:
                                                        var v136 uint64 = 859
                                                        var v137 int16 = -32768
                                                        if (v57 >= (v57 + (v57 % ((fi135) | 1)))) {
                                                            v134 = v134
                                                            v138 := GenS3{F0: 100, F1: 42}
                                                            if (false || ((v4 + 698) <= v4)) {
                                                                switch fi135 {
                                                                case 4294967295:
                                                                    if false {
                                                                        var v139 uint16 = 32767
                                                                        v57 = v57
                                                                        _ = v139
                                                                    }
                                                                    v136 += v136
                                                                    v140 := GenS10{F0: -736, F1: v137}
                                                                    _ = v140
                                                                case 2147483647:
                                                                    delete(v3, ("hello" + v58))
                                                                default:
                                                                    if false {
                                                                        if (!true) {
                                                                            v141 := GenS10{F0: v42, F1: (v42 - v42)}
                                                                            _ = v141
                                                                        } else {
                                                                            fmt.Println(v42)
                                                                            var v142 bool = true
                                                                            _ = v142
                                                                        }
                                                                    } else {
                                                                        v143 := v6[2147483647]
                                                                        _ = v143
                                                                    }
                                                                }
                                                                _ = len(v6)
                                                                _ = len(v6)
                                                            }
                                                            _ = v138
                                                        }
                                                        _ = v136
                                                        _ = v137
                                                    }
                                                } else {
                                                    v4 = (v4 << ((v4) & 63))
                                                    var v144 float32 = 1.5
                                                    _ = v144
                                                }
                                            }
                                        }
                                        _ = fi135
                                    }
                                    for fi145 := 0; fi145 < 14; fi145++ {
                                        var v146 float64 = float64(float32(1e10))
                                        var v147 int64 = 0
                                        _ = fi145
                                        _ = v146
                                        _ = v147
                                    }
                                    for fi148 := 0; fi148 < 18; fi148++ {
                                        fmt.Println(v5, v4, v2)
                                        fmt.Println(fi148, v5, v2)
                                        v58 = "\n\t"
                                        _ = fi148
                                    }
                                } else {
                                    var v149 uint64 = 18446744073709551615
                                    var v150 float64 = 0.1
                                    _ = v149
                                    _ = v150
                                }
                                if len(v118) >= 2 {
                                    v118 = v118[1:len(v118)]
                                }
                                _ = v134
                            } else {
                                for fi151 := 0; fi151 < 12; fi151++ {
                                    v5 -= v5
                                    if (!false) {
                                        var v152 int16 = v42
                                        v4 = (v4 - (v4 - v4))
                                        var v153 uint = 919
                                        _ = v152
                                        _ = v153
                                    } else {
                                        if (((1.5 / -2.25) / 3.14159) != 1.5) {
                                            v42 = (v42 & 366)
                                        } else {
                                            var v154 int64 = 9223372036854775807
                                            var v155 uint16 = 1
                                            var v156 uint64 = 18446744073709551615
                                            _ = v154
                                            _ = v155
                                            _ = v156
                                        }
                                        fmt.Println(v4, v4)
                                        if (v57 < (v57 ^ v57)) {
                                            var v157 int32 = -186
                                            switch -566 {
                                            case -2147483648:
                                                fmt.Println(v2)
                                            case 757:
                                                v118 = append(v118, int(v42))
                                            }
                                            v2 = fi151
                                            _ = v157
                                        } else {
                                            var v158 float32 = 1.5
                                            var v159 int16 = v42
                                            v5 -= (v5 | v5)
                                            _ = v158
                                            _ = v159
                                        }
                                    }
                                    _ = fi151
                                }
                                switch ("a" + v58) {
                                case "hello":
                                    if true {
                                        var v160 uint32 = 2147483648
                                        var v161 int = fi125
                                        v161 -= (v2 + v2)
                                        _ = v160
                                        _ = v161
                                    }
                                case "":
                                    for fi162 := 0; fi162 < 14; fi162++ {
                                        var v163 uint16 = 739
                                        var v164 uint = (v4 | uint(v163))
                                        _ = fi162
                                        _ = v163
                                        _ = v164
                                    }
                                default:
                                    if len(v118) >= 2 {
                                        v118 = v118[1:len(v118)]
                                    }
                                }
                                _ = len(v6)
                            }
                        case "\n\t":
                            v57 = v2
                        }
                        _ = fi125
                    }
                } else {
                    var v165 int8 = v5
                    switch (v2 >> (((^(v4))) & 63)) {
                    case 32768:
                        switch v2 {
                        case 127:
                            fmt.Println(v1)
                        case -9223372036854775808:
                            if false {
                                var v166 uint8 = v1
                                v57 -= 32767
                                if true {
                                    var v167 int = v2
                                    var v168 int32 = 42
                                    var v169 int32 = (v168 << ((1) & 63))
                                    _ = v167
                                    _ = v168
                                    _ = v169
                                } else {
                                    if false {
                                        if ((v5 | v165) < (v5 % (((v165 * (v165 - (-(v165))))) | 1))) {
                                            var v170 string = v58
                                            var v171 uint32 = 42
                                            v6[v171] = v58
                                            _ = v170
                                            _ = v171
                                        }
                                        var v172 uint32 = uint32(v4)
                                        _ = len(v118)
                                        _ = v172
                                    } else {
                                        var v173 string = v58
                                        _ = v173
                                    }
                                    if (v58 == v58) {
                                        if false {
                                            v42 -= (v42 + 1)
                                            var v174 uint32 = uint32(v4)
                                            _ = v174
                                        }
                                    } else {
                                        var v175 uint8 = genFn1(v166, int(v5))
                                        var v176 int8 = -128
                                        v177 := GenS16{F0: 65535, F1: 1, F2: (v1 / ((42) | 1))}
                                        _ = v175
                                        _ = v176
                                        _ = v177
                                    }
                                    v178 := GenS10{F0: v42, F1: v42}
                                    _ = v178
                                }
                                _ = v166
                            }
                        case 4294967295:
                            v2 = (-(v57))
                        }
                    case 99:
                        switch (v58 + v58) {
                        case "\n\t":
                            v1 = 0
                        case "foo bar":
                            delete(v6, 1)
                        }
                    case 36:
                        v5 = (v5 + v5)
                    }
                    if false {
                        v179 := GenS9{F0: 122, F1: v58}
                        if true {
                            v57 = (v57 >> (((v4 | v4)) & 63))
                            if (false || true) {
                                var v180 int8 = v5
                                for fi181 := 0; fi181 < 16; fi181++ {
                                    if (((float32(3.14159) * ((float32(float64(1e10)) + -2.25) / float32(-2.25))) + (((1.5 + 1e10) + 0.1) * ((0.0 + (3.14159 - -2.25)) / (((-2.25 - float32(0.0)) + ((0.0 / 1.0) - 3.14159)) * 3.14159)))) < -2.25) {
                                        for fi182 := 0; fi182 < 3; fi182++ {
                                            if (false && ((!((2 <= 2) || false)) && (-1 != 1))) {
                                                v42 = v42
                                            }
                                            v4 -= (v4 / ((v4) | 1))
                                            if false {
                                                fmt.Println(v1)
                                                var v183 int16 = (v42 & -1)
                                                v184 := GenS19{F0: 784, F1: fi182, F2: 1}
                                                var v185 bool = false
                                                if (1e10 != 0.1) {
                                                    v186 := GenS20{F0: v2, F1: 4294967295}
                                                    v185 = ((v1 + 0) <= (^(v1)))
                                                    v187 := GenS21{F0: 310, F1: v1, F2: 1}
                                                    _ = v186
                                                    _ = v187
                                                }
                                                _ = v183
                                                _ = v184
                                                _ = v185
                                            }
                                            _ = fi182
                                        }
                                        for fi188 := 0; fi188 < 17; fi188++ {
                                            switch v57 {
                                            case -2147483648:
                                                v189 := GenS22{F0: 24, F1: v57, F2: v58}
                                                _ = v189
                                            case 2147483647:
                                                if (true || (!((v4 % (((v4 | 42)) | 1)) < (v4 << ((v4) & 63))))) {
                                                    var v190 float64 = 1.5
                                                    var v191 int32 = 42
                                                    _ = v190
                                                    _ = v191
                                                }
                                            case 0:
                                                v192 := GenS12{F0: 181, F1: 1, F2: (v180 ^ -37)}
                                                _ = v192
                                            }
                                            if len(v118) > 0 {
                                                v118[0] = 32768
                                            }
                                            _ = fi188
                                        }
                                    } else {
                                        fmt.Println(v5, v2)
                                        for fi193 := 0; fi193 < 20; fi193++ {
                                            var v194 uint64 = 745
                                            if true {
                                                switch (v58 + v58) {
                                                case "\n\t":
                                                    v165 = (v180 >> ((v4) & 63))
                                                case "foo bar":
                                                    v58 = v58
                                                case "hello":
                                                    switch v2 {
                                                    case 408:
                                                        v5 = -36
                                                    }
                                                }
                                                switch "foo bar" {
                                                case "hello":
                                                    var v195 uint16 = 65535
                                                    _ = v195
                                                default:
                                                    v118 = append(v118, v2)
                                                }
                                            } else {
                                                var v196 float64 = 0.0
                                                _ = v196
                                            }
                                            fmt.Println(v165)
                                            _ = fi193
                                            _ = v194
                                        }
                                    }
                                    _ = fi181
                                }
                                _ = v180
                            } else {
                                if (true && true) {
                                    var v197 int = v2
                                    if (!true) {
                                        for fi198 := 0; fi198 < 11; fi198++ {
                                            v199 := GenS22{F0: 848, F1: (^(fi198)), F2: (("" + v58) + (v58 + v58))}
                                            v5 = (v165 << ((v4) & 63))
                                            v1 = genFn1((^(v1)), (^(fi198)))
                                            _ = fi198
                                            _ = v199
                                        }
                                        var v200 float64 = 0.1
                                        fmt.Println(v2, v5, v5)
                                        _ = v200
                                    }
                                    v201 := GenS20{F0: (-(v197)), F1: 2}
                                    _ = v197
                                    _ = v201
                                }
                                v202 := GenS16{F0: 0, F1: v1, F2: (^(v1))}
                                _ = v202
                            }
                        }
                        _ = v179
                    } else {
                        for fi203 := 0; fi203 < 12; fi203++ {
                            if (((-2.25 - 0.0) - 0.0) >= (((0.0 / (0.1 + float32(1.5))) - ((0.0 * 0.0) - 0.0)) - (3.14159 * 3.14159))) {
                                switch (v2 | 32767) {
                                case 324:
                                    _ = len(v6)
                                case 32767:
                                    _ = len(v118)
                                }
                                var v204 uint64 = 2
                                _ = v204
                            }
                            _ = fi203
                        }
                        v165 = (v165 << ((405) & 63))
                    }
                    _ = v165
                }
                _ = v118
            }
        } else {
            v2 = v57
        }
    case -25:
        if false {
            v205 := GenS9{F0: 42, F1: v58}
            for fi206 := 0; fi206 < 14; fi206++ {
                v58 = v58
                var v207 float64 = 1e10
                _ = fi206
                _ = v207
            }
            _ = v205
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v110
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v57
    _ = v58
    _ = v6
}
