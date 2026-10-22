package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint32; F1 bool; F2 uint }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 int32; F1 int16; F2 bool }
type GenS6 struct { F0 int16; F1 int; F2 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint64; F1 uint64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int; F1 int; F2 int16 }
type GenS11 struct { F0 int64; F1 int16 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint8) bool {
    if (!true) {
        v7 := []int{105, 0}
        if true {
            v7 = append(v7, 128)
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 float32 = ((0.1 * -2.25) - 1.5)
                fmt.Println(v9, fi8, v9)
                v10 := GenS1{F0: 0}
                _ = fi8
                _ = v10
                _ = v9
            }
            v11 := GenS1{F0: 2}
            _ = v11
        } else {
            for fi12 := 0; fi12 < 19; fi12++ {
                switch (("hello" + ((("a" + ("hello" + "")) + "") + "a")) + "") {
                case "":
                    if (42 <= 1) {
                        v13 := GenS1{F0: 1}
                        var v14 uint = 23
                        _ = v13
                        _ = v14
                    } else {
                        v7 = append(v7, (fi12 * fi12))
                        fmt.Println(p0, fi12)
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                case "a":
                    v15 := GenS2{F0: 54, F1: ("" + ("\n\t" + ("foo bar" + ("" + "foo bar"))))}
                    _ = v15
                default:
                    var v16 uint8 = p0
                    _ = v16
                }
                var v17 int = (fi12 >> ((18446744073709551615) & 63))
                var v18 int = (fi12 << ((0) & 63))
                _ = fi12
                _ = v17
                _ = v18
            }
            var v19 float64 = ((((((-2.25 * (1e10 / 0.1)) + ((1e10 - 1e10) + 0.1)) + 1.5) - float64(float32(3.14159))) + float64(0.1)) * float64(-2.25))
            _ = v19
        }
        v20 := GenS2{F0: 2, F1: ("" + ("hello" + "foo bar"))}
        _ = v20
        _ = v7
    } else {
        for fi21 := 0; fi21 < 15; fi21++ {
            if false {
                p0 -= p0
                v22 := GenS3{F0: 1, F1: false, F2: 65535}
                _ = v22
            } else {
                var v23 float64 = 1e10
                var v24 int64 = -9223372036854775808
                var v25 int64 = (v24 % ((-1) | 1))
                _ = v23
                _ = v24
                _ = v25
            }
            var v26 int16 = 1000
            _ = fi21
            _ = v26
        }
    }
    var v27 int64 = 9223372036854775807
    if (!false) {
        p0 = p0
    }
    var v28 int32 = 1
    _ = p0
    _ = v27
    _ = v28
    return (407 == 78)
}
func genFn1(p0 int, p1 int) uint {
    if ((!true) && false) {
        fmt.Println(p1, p1)
    } else {
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 uint = 119
            _ = fi29
            _ = v30
        }
    }
    v31 := []int{(p1 << ((65535) & 63)), 32768, p0, -1}
    _ = p0
    _ = p1
    _ = v31
    return 383
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]bool)
    var v3 int32 = 0
    v4 := make(map[int]string)
    var v5 float32 = (3.14159 + 3.14159)
    var v6 int64 = 42
    v32 := v4[1]
    var v33 uint32 = 1
    for fi34 := 0; fi34 < 14; fi34++ {
        v35 := GenS4{F0: (v33 / (((^(v33))) | 1)), F1: 288}
        v3 -= 1
        var v36 uint16 = 2
        _ = fi34
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 4; fi37++ {
        v3 = v3
        v32 = ""
        if false {
            for fi38 := 0; fi38 < 2; fi38++ {
                v39 := GenS3{F0: (v33 + (^(v33))), F1: true, F2: 42}
                delete(v4, 9223372036854775807)
                _ = fi38
                _ = v39
            }
        } else {
            var v40 uint32 = v33
            _ = v40
        }
        _ = fi37
    }
    fmt.Println(v3)
    v6 -= (v6 * v6)
    if true {
        if false {
            var v41 string = v32
            v42 := GenS1{F0: v33}
            if ((89 <= genFn1(-1, int(v3))) || true) {
                v33 += v33
                v43 := GenS5{F0: (-(v3)), F1: 416, F2: false}
                if true {
                    if true {
                        switch -682 {
                        case -1:
                            if ((v6 ^ 0) != (v6 | (v6 / ((854) | 1)))) {
                                for fi44 := 0; fi44 < 15; fi44++ {
                                    var v45 uint32 = v33
                                    v41 = v32
                                    for fi46 := 0; fi46 < 7; fi46++ {
                                        v47 := GenS2{F0: 2, F1: "hello"}
                                        _ = len(v4)
                                        v48 := GenS1{F0: (v45 & (v33 + v45))}
                                        _ = fi46
                                        _ = v47
                                        _ = v48
                                    }
                                    _ = fi44
                                    _ = v45
                                }
                                delete(v2, (v33 << ((4294967295) & 63)))
                            } else {
                                v49 := GenS6{F0: -32768, F1: -325, F2: 127}
                                _ = v49
                            }
                        case 32767:
                            var v50 float32 = ((3.14159 * 0.0) + v5)
                            _ = v50
                        case 2147483647:
                            if true {
                                var v51 uint8 = uint8(v33)
                                var v52 uint64 = 18446744073709551615
                                v53 := GenS7{F0: 32767}
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                        default:
                            _ = len(v1)
                        }
                        v33 = v33
                        for fi54 := 0; fi54 < 20; fi54++ {
                            var v55 int64 = (v6 * (v6 >> ((genFn1(fi54, fi54)) & 63)))
                            _ = fi54
                            _ = v55
                        }
                    } else {
                        switch 128 {
                        case -9223372036854775808:
                            v56 := GenS1{F0: 4294967295}
                            _ = v56
                        case 470:
                            for fi57 := 0; fi57 < 14; fi57++ {
                                if (false || true) {
                                    if true {
                                        v58 := GenS4{F0: 282, F1: 247}
                                        fmt.Println(v6, v5, fi57)
                                        _ = v58
                                    } else {
                                        v59 := GenS2{F0: 255, F1: ((("a" + ("hello" + "hello")) + v41) + v41)}
                                        for fi60 := 0; fi60 < 9; fi60++ {
                                            fmt.Println(v3, v41)
                                            v1 = append(v1, "\n\t")
                                            _ = fi60
                                        }
                                        _ = v59
                                    }
                                    v61 := v2[42]
                                    _ = v61
                                }
                                _ = fi57
                            }
                        case 127:
                            for fi62 := 0; fi62 < 14; fi62++ {
                                for fi63 := 0; fi63 < 8; fi63++ {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v64 int = fi63
                                    _ = fi63
                                    _ = v64
                                }
                                for fi65 := 0; fi65 < 8; fi65++ {
                                    if false {
                                        var v66 uint16 = 10
                                        switch (-(fi65)) {
                                        case -9223372036854775808:
                                            v1 = append(v1, "\n\t")
                                            v1 = append(v1, v41)
                                            fmt.Println(v32, v6, v41)
                                        case 332:
                                            v5 = v5
                                            v67 := GenS5{F0: 1, F1: 461, F2: true}
                                            switch (fi65 & -776) {
                                            case -1:
                                                v5 = (v5 - v5)
                                            case 9223372036854775807:
                                                v6 += (v6 * v6)
                                            }
                                            _ = v67
                                        case 42:
                                            if true {
                                                fmt.Println(v3, fi62)
                                                v33 += (v33 - v33)
                                            } else {
                                                v3 = 0
                                                _ = len(v2)
                                            }
                                        }
                                        _ = v66
                                    } else {
                                        var v68 float64 = 1.5
                                        if ((fi62 ^ (fi62 * fi62)) == fi62) {
                                            v6 = (v6 * (^(v6)))
                                            fmt.Println(v33, v6, v3)
                                            v69 := GenS2{F0: 14, F1: "foo bar"}
                                            if (v32 >= (v32 + v41)) {
                                                switch v41 {
                                                case "\n\t":
                                                    var v70 float32 = v5
                                                    _ = v70
                                                case "":
                                                    fmt.Println(v3)
                                                case "a":
                                                    fmt.Println(v68)
                                                }
                                                _ = len(v4)
                                            }
                                            _ = v69
                                        } else {
                                            v71 := GenS6{F0: 1, F1: (fi65 & -387), F2: 8}
                                            var v72 int = (fi62 >> ((4294967295) & 63))
                                            if (v33 == 1) {
                                                var v73 int8 = -128
                                                _ = v73
                                            }
                                            _ = v71
                                            _ = v72
                                        }
                                        _ = v68
                                    }
                                    if false {
                                        v74 := v4[(^(fi65))]
                                        var v75 int64 = v6
                                        if (fi65 > (fi65 * fi62)) {
                                            var v76 uint8 = 170
                                            _ = v76
                                        } else {
                                            fmt.Println(v6)
                                            v77 := v4[(fi65 << ((599) & 63))]
                                            if true {
                                                var v78 int32 = (v3 | v3)
                                                _ = v78
                                            } else {
                                                v3 -= (v3 >> ((18446744073709551615) & 63))
                                            }
                                            _ = v77
                                        }
                                        _ = v74
                                        _ = v75
                                    }
                                    v6 -= int64(v3)
                                    _ = fi65
                                }
                                _ = fi62
                            }
                        default:
                            switch "\n\t" {
                            case "a":
                                for fi79 := 0; fi79 < 4; fi79++ {
                                    v3 += v3
                                    for fi80 := 0; fi80 < 8; fi80++ {
                                        _ = len(v1)
                                        switch (fi79 << ((genFn1(fi79, fi80)) & 63)) {
                                        case 9223372036854775807:
                                            fmt.Println(fi80)
                                        }
                                        fmt.Println(v32)
                                        var v81 float64 = 1.5
                                        v5 = float32(v81)
                                        var v82 uint16 = 255
                                        var v83 uint32 = (v33 / ((170) | 1))
                                        var v84 int32 = (-(v3))
                                        v4[0] = v41
                                        _ = fi80
                                        _ = v81
                                        _ = v82
                                        _ = v83
                                        _ = v84
                                    }
                                    _ = fi79
                                }
                            default:
                                for fi85 := 0; fi85 < 16; fi85++ {
                                    var v86 string = v32
                                    _ = fi85
                                    _ = v86
                                }
                            }
                        }
                    }
                }
                _ = v43
            }
            _ = v41
            _ = v42
        } else {
            if (((1 >= 65535) || false) || (!(!(-619 > 2147483647)))) {
                if true {
                    fmt.Println(v5, v32, v5)
                    v33 += v33
                } else {
                    v87 := v4[42]
                    for fi88 := 0; fi88 < 11; fi88++ {
                        fmt.Println(fi88)
                        switch (-(fi88)) {
                        case -904:
                            if true {
                                for fi89 := 0; fi89 < 17; fi89++ {
                                    switch v87 {
                                    case "foo bar":
                                        fmt.Println(fi88, v6)
                                    case "":
                                        v90 := GenS6{F0: -1, F1: fi89, F2: -88}
                                        _ = v90
                                    case "hello":
                                        if false {
                                            var v91 bool = (float32(0.0) < v5)
                                            v91 = (!(v91 && (!(true && ((9223372036854775807 >= 0) && (0 > 0))))))
                                            _ = v91
                                        }
                                    default:
                                        var v92 uint64 = 854
                                        if true {
                                            fmt.Println(v32, fi88, v33)
                                            if false {
                                                v93 := GenS1{F0: v33}
                                                if false {
                                                    v1 = append(v1, (v32 + (v87 + v87)))
                                                    _ = len(v1)
                                                    fmt.Println(v87, v92)
                                                    if (42 == 65535) {
                                                        v3 = v3
                                                        v32 = (v32 + v87)
                                                    } else {
                                                        var v94 string = "a"
                                                        _ = v94
                                                    }
                                                    if false {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            var v95 string = ("a" + v87)
                                                            _ = v95
                                                        }
                                                    }
                                                    var v96 float32 = v5
                                                    _ = v96
                                                } else {
                                                    var v97 int32 = v3
                                                    _ = v97
                                                }
                                                v98 := GenS8{F0: (v92 * 727), F1: (v92 - v92)}
                                                _ = v93
                                                _ = v98
                                            } else {
                                                if ((^(v3)) != (v3 << ((378) & 63))) {
                                                    var v99 bool = (true && false)
                                                    _ = v99
                                                }
                                            }
                                        }
                                        v100 := GenS2{F0: 2, F1: v87}
                                        _ = v100
                                        _ = v92
                                    }
                                    v101 := v4[127]
                                    _ = fi89
                                    _ = v101
                                }
                            } else {
                                var v102 uint32 = (v33 - (^(v33)))
                                v102 = v102
                                _ = v102
                            }
                        default:
                            if false {
                                var v103 float64 = -2.25
                                var v104 uint32 = 42
                                var v105 uint8 = 1
                                _ = v103
                                _ = v104
                                _ = v105
                            } else {
                                v87 = v87
                            }
                        }
                        for fi106 := 0; fi106 < 8; fi106++ {
                            if (!true) {
                                var v107 uint64 = 2
                                v108 := GenS9{F0: v6}
                                var v109 int64 = (v6 << ((genFn1(fi88, fi106)) & 63))
                                if len(v1) > 0 {
                                    v1[0] = ((("" + "") + ((v32 + v87) + "a")) + "")
                                }
                                _ = v107
                                _ = v108
                                _ = v109
                            } else {
                                if (!(!(true || false))) {
                                    fmt.Println(v87, fi88)
                                    fmt.Println(v5)
                                    var v110 uint16 = uint16(v33)
                                    var v111 uint16 = (v110 + v110)
                                    switch (v87 + "foo bar") {
                                    case "":
                                        v112 := GenS9{F0: (v6 | -573)}
                                        _ = v112
                                    case "foo bar":
                                        var v113 uint16 = v110
                                        _ = v113
                                    case "\n\t":
                                        var v114 int16 = 0
                                        _ = v114
                                    default:
                                        if true {
                                            v111 = v111
                                        } else {
                                            v110 = (v110 | v110)
                                            var v115 float64 = 0.0
                                            fmt.Println(v5, v5, fi88)
                                            _ = v115
                                        }
                                    }
                                    v32 = "a"
                                    _ = v110
                                    _ = v111
                                }
                                v1 = append(v1, v32)
                            }
                            _ = fi106
                        }
                        _ = fi88
                    }
                    switch -177 {
                    case 687:
                        if (432 >= 1) {
                            var v116 uint32 = v33
                            v32 = v87
                            var v117 string = v87
                            _ = v116
                            _ = v117
                        } else {
                            var v118 string = (v32 + v32)
                            if true {
                                v33 = 4294967295
                                switch v32 {
                                case "\n\t":
                                    if (true || (!true)) {
                                        for fi119 := 0; fi119 < 12; fi119++ {
                                            switch ((v118 + v32) + (v87 + v87)) {
                                            case "a":
                                                v120 := GenS10{F0: (fi119 & (fi119 / ((fi119) | 1))), F1: (fi119 + (fi119 - -514)), F2: -32768}
                                                _ = v120
                                            case "":
                                                if false {
                                                    var v121 int = (fi119 | 2147483647)
                                                    for fi122 := 0; fi122 < 13; fi122++ {
                                                        if false {
                                                            v118 = (v32 + "")
                                                            if false {
                                                                delete(v2, 58)
                                                                v123 := GenS11{F0: v6, F1: -445}
                                                                _ = v123
                                                            } else {
                                                                if true {
                                                                    v1 = append(v1, (v32 + v118))
                                                                    if true {
                                                                        v118 = v32
                                                                    } else {
                                                                        var v124 float32 = (v5 + float32(3.14159))
                                                                        var v125 uint32 = v33
                                                                        var v126 string = v87
                                                                        v87 = "\n\t"
                                                                        _ = v124
                                                                        _ = v125
                                                                        _ = v126
                                                                    }
                                                                    var v127 int32 = v3
                                                                    _ = v127
                                                                } else {
                                                                    _ = len(v4)
                                                                    if true {
                                                                        fmt.Println(v32, v33)
                                                                        var v128 int = -1
                                                                        _ = v128
                                                                    }
                                                                    v121 -= (fi119 | (v121 ^ v121))
                                                                }
                                                                delete(v4, (fi119 << ((365) & 63)))
                                                            }
                                                        }
                                                        _ = fi122
                                                    }
                                                    _ = v121
                                                }
                                            }
                                            _ = fi119
                                        }
                                        fmt.Println(v87)
                                    }
                                case "a":
                                    for fi129 := 0; fi129 < 16; fi129++ {
                                        switch v118 {
                                        case "\n\t":
                                            var v130 float32 = v5
                                            _ = v130
                                        case "a":
                                            switch fi129 {
                                            case -1:
                                                for fi131 := 0; fi131 < 17; fi131++ {
                                                    fmt.Println(v118, v118, fi129)
                                                    _ = fi131
                                                }
                                            default:
                                                if (!(int(v33) >= fi129)) {
                                                    var v132 bool = false
                                                    switch v87 {
                                                    case "\n\t":
                                                        var v133 float32 = ((v5 * (((3.14159 / ((1e10 * v5) + (1e10 / -2.25))) / (v5 - v5)) * v5)) / v5)
                                                        _ = v133
                                                    case "hello":
                                                        v132 = (v132 || (473 > 4294967295))
                                                    case "":
                                                        if (0 == (v6 & v6)) {
                                                            if (!true) {
                                                                _ = len(v4)
                                                                for fi134 := 0; fi134 < 11; fi134++ {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = (v32 + ("hello" + "\n\t"))
                                                                    }
                                                                    v135 := GenS10{F0: -288, F1: (fi129 + (-(fi129))), F2: 834}
                                                                    var v136 bool = (2147483647 != v3)
                                                                    if (32767 == 1000) {
                                                                        v32 = ""
                                                                    }
                                                                    var v137 float32 = (v5 - (v5 + v5))
                                                                    _ = fi134
                                                                    _ = v135
                                                                    _ = v136
                                                                    _ = v137
                                                                }
                                                            } else {
                                                                for fi138 := 0; fi138 < 17; fi138++ {
                                                                    if ((v132 && (2147483647 != int(v33))) && false) {
                                                                        v5 = 0.0
                                                                        v139 := GenS8{F0: 18446744073709551615, F1: 4294967295}
                                                                        _ = v139
                                                                    } else {
                                                                        var v140 int16 = int16(fi138)
                                                                        fmt.Println(v5, v3, v140)
                                                                        fmt.Println(fi138, fi129, fi138)
                                                                        _ = v140
                                                                    }
                                                                    _ = fi138
                                                                }
                                                            }
                                                        }
                                                    }
                                                    _ = v132
                                                }
                                            }
                                        case "":
                                            switch fi129 {
                                            case -9223372036854775808:
                                                switch (((v118 + v118) + (v32 + (("foo bar" + "a") + (v118 + "hello")))) + (v118 + (((v118 + v87) + v87) + v32))) {
                                                case "a":
                                                    switch v118 {
                                                    case "":
                                                        for fi141 := 0; fi141 < 5; fi141++ {
                                                            v142 := GenS1{F0: 353}
                                                            _ = fi141
                                                            _ = v142
                                                        }
                                                    case "foo bar":
                                                        v33 = (v33 * (^(v33)))
                                                    default:
                                                        v143 := GenS11{F0: 1, F1: -845}
                                                        _ = v143
                                                    }
                                                default:
                                                    if len(v1) > 0 {
                                                        v1[0] = v32
                                                    }
                                                }
                                            }
                                        default:
                                            var v144 string = v32
                                            _ = v144
                                        }
                                        fmt.Println(v5, v33)
                                        _ = fi129
                                    }
                                case "":
                                    switch -1 {
                                    case 42:
                                        for fi145 := 0; fi145 < 20; fi145++ {
                                            v33 = (v33 | v33)
                                            v5 += v5
                                            _ = fi145
                                        }
                                    case 9223372036854775807:
                                        for fi146 := 0; fi146 < 9; fi146++ {
                                            v5 = (v5 - 3.14159)
                                            v3 = v3
                                            _ = fi146
                                        }
                                    case -595:
                                        v147 := GenS12{F0: 916}
                                        _ = v147
                                    }
                                }
                            } else {
                                v148 := GenS9{F0: v6}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v148
                            }
                            _ = v118
                        }
                    case -71:
                        if false {
                            var v149 int64 = 602
                            v32 = "\n\t"
                            for fi150 := 0; fi150 < 1; fi150++ {
                                if false {
                                    switch -708 {
                                    case 18:
                                        var v151 int64 = (v149 - v149)
                                        _ = v151
                                    case -763:
                                        if true {
                                            switch (fi150 >> ((1) & 63)) {
                                            case -9223372036854775808:
                                                fmt.Println(v33, v33, v5)
                                            case 42:
                                                if len(v1) > 0 {
                                                    v1[0] = "a"
                                                }
                                            case 2147483647:
                                                v32 = v87
                                            default:
                                                for fi152 := 0; fi152 < 11; fi152++ {
                                                    var v153 uint32 = (^(v33))
                                                    v154 := GenS11{F0: v149, F1: 32767}
                                                    if ((false && true) && false) {
                                                        var v155 int32 = v3
                                                        _ = v155
                                                    } else {
                                                        var v156 uint32 = v153
                                                        if false {
                                                            if false {
                                                                if (-32768 < -1) {
                                                                    fmt.Println(fi150)
                                                                    var v157 int32 = (v3 & v3)
                                                                    _ = v157
                                                                }
                                                                v87 = v32
                                                            }
                                                        } else {
                                                            var v158 uint32 = v153
                                                            switch (^(fi152)) {
                                                            case 32768:
                                                                var v159 uint64 = 382
                                                                _ = v159
                                                            case -9223372036854775808:
                                                                fmt.Println(v87, v33, v3)
                                                                if ((float64(((v5 / (1e10 * v5)) - 0.0)) >= 1.5) || true) {
                                                                    _ = len(v4)
                                                                    var v160 float64 = 1.5
                                                                    _ = v160
                                                                }
                                                                var v161 uint32 = (v153 >> ((920) & 63))
                                                                _ = v161
                                                            case 863:
                                                                v2[v156] = false
                                                            default:
                                                                v87 = ""
                                                            }
                                                            delete(v2, v153)
                                                            v149 = 1
                                                            _ = v158
                                                        }
                                                        v162 := GenS1{F0: (v153 + (v156 % ((65535) | 1)))}
                                                        _ = v156
                                                        _ = v162
                                                    }
                                                    v33 = v153
                                                    _ = fi152
                                                    _ = v153
                                                    _ = v154
                                                }
                                            }
                                        }
                                    case -282:
                                        var v163 float32 = v5
                                        _ = v163
                                    }
                                    switch (fi150 + (fi150 ^ (^(fi150)))) {
                                    case 1:
                                        v149 = (v149 % ((v149) | 1))
                                    }
                                } else {
                                    fmt.Println(v32, v33)
                                }
                                _ = fi150
                            }
                            _ = v149
                        }
                    case -1:
                        v164 := GenS8{F0: 543, F1: 890}
                        _ = v164
                    default:
                        for fi165 := 0; fi165 < 10; fi165++ {
                            v166 := GenS2{F0: 65535, F1: (v32 + v32)}
                            _ = fi165
                            _ = v166
                        }
                    }
                    _ = v87
                }
                switch -1 {
                case -437:
                    var v167 int8 = 127
                    _ = v167
                default:
                    v168 := v2[v33]
                    _ = v168
                }
                _ = len(v4)
            } else {
                var v169 float64 = 1e10
                if false {
                    if (!true) {
                        for fi170 := 0; fi170 < 6; fi170++ {
                            var v171 uint = 18446744073709551615
                            if true {
                                v32 = v32
                                var v172 int = fi170
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v172
                            }
                            _ = fi170
                            _ = v171
                        }
                        var v173 uint8 = 255
                        _ = v173
                    }
                    v174 := GenS9{F0: (v6 - v6)}
                    _ = v174
                } else {
                    _ = len(v2)
                    if false {
                        _ = len(v1)
                        if (v32 == "hello") {
                            var v175 uint16 = 941
                            var v176 float64 = (float64((v5 / (v5 - float32((1e10 * 1.5))))) * (v169 * 1e10))
                            _ = v175
                            _ = v176
                        } else {
                            var v177 uint64 = 278
                            for fi178 := 0; fi178 < 10; fi178++ {
                                var v179 int = fi178
                                _ = fi178
                                _ = v179
                            }
                            for fi180 := 0; fi180 < 3; fi180++ {
                                v6 -= 42
                                _ = fi180
                            }
                            _ = v177
                        }
                    } else {
                        v33 = v33
                        switch 817 {
                        case 42:
                            if (!true) {
                                v32 = "\n\t"
                                v32 = ((v32 + v32) + "foo bar")
                            }
                        case 32767:
                            fmt.Println(v32, v32, v6)
                        }
                    }
                    v169 -= (v169 - float64(-2.25))
                }
                _ = v169
            }
        }
        switch (("\n\t" + v32) + ("" + (v32 + (v32 + "a")))) {
        case "\n\t":
            fmt.Println(v3, v5)
        case "a":
            if ((true || (4294967295 == 1)) && false) {
                if ("a" >= "") {
                    v181 := GenS12{F0: 2}
                    _ = v181
                } else {
                    if false {
                        v1 = append(v1, v32)
                        for fi182 := 0; fi182 < 5; fi182++ {
                            if true {
                                if true {
                                    v33 -= (v33 >> ((65535) & 63))
                                    v6 = (v6 ^ -9223372036854775808)
                                    v5 += v5
                                }
                            } else {
                                switch "" {
                                case "a":
                                    if true {
                                        v6 += v6
                                        v33 = (v33 + (^(v33)))
                                    }
                                case "":
                                    var v183 uint32 = (v33 >> ((1) & 63))
                                    _ = v183
                                default:
                                    var v184 int64 = (v6 | -1)
                                    _ = v184
                                }
                            }
                            var v185 int = fi182
                            for fi186 := 0; fi186 < 3; fi186++ {
                                if true {
                                    if (1 > (v3 / ((42) | 1))) {
                                        var v187 int64 = -1
                                        fmt.Println(v5, v32, v5)
                                        var v188 uint = genFn1(768, int(v3))
                                        v189 := GenS13{F0: 4294967295, F1: 42}
                                        _ = len(v2)
                                        _ = v187
                                        _ = v188
                                        _ = v189
                                    }
                                    var v190 int16 = -331
                                    if true {
                                        var v191 int64 = (-(v6))
                                        _ = len(v1)
                                        v190 = -32768
                                        _ = v191
                                    }
                                    _ = v190
                                } else {
                                    var v192 uint = 901
                                    _ = v192
                                }
                                _ = fi186
                            }
                            _ = fi182
                            _ = v185
                        }
                        v32 = ((v32 + v32) + ("hello" + "hello"))
                    } else {
                        v2[v33] = (328 != 1)
                    }
                }
            } else {
                fmt.Println(v6, v33, v3)
                v193 := GenS7{F0: 42}
                _ = v193
            }
        case "hello":
            if true {
                _ = len(v4)
                for fi194 := 0; fi194 < 17; fi194++ {
                    for fi195 := 0; fi195 < 18; fi195++ {
                        v32 = (v32 + "hello")
                        _ = fi195
                    }
                    v33 = v33
                    _ = fi194
                }
                var v196 int64 = v6
                _ = v196
            }
        default:
            v197 := GenS4{F0: v33, F1: 0}
            _ = v197
        }
        if (false && true) {
            v33 += 2147483648
        }
    }
    v32 = (("foo bar" + "hello") + "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
