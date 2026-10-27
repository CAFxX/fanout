package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int32 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int64; F1 uint16; F2 int64 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 int16; F1 string; F2 int64 }
type GenS7 struct { F0 int; F1 int16; F2 int64 }
type GenS8 struct { F0 uint16; F1 uint; F2 int8 }
type GenS9 struct { F0 uint64; F1 uint16; F2 string }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint64; F1 string; F2 int }
type GenS12 struct { F0 int8; F1 string; F2 int64 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int; F1 uint64; F2 uint64 }
type GenS15 struct { F0 int64; F1 int; F2 uint8 }
type GenS16 struct { F0 int16; F1 uint64 }
func genFn0() uint {
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 float32 = (3.14159 - 0.1)
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(fi6)
        fmt.Println(fi6)
        switch (("\n\t" + "\n\t") + "foo bar") {
        case "hello":
            var v7 float64 = 1e10
            _ = v7
        default:
            v8 := []int{(fi6 & fi6), -9223372036854775808, fi6}
            _ = v8
        }
        _ = fi6
    }
    var v9 string = ""
    for fi10 := 0; fi10 < 5; fi10++ {
        v9 = v9
        v11 := GenS1{F0: -128, F1: 290, F2: 1}
        _ = fi10
        _ = v11
    }
    var v12 uint16 = 635
    _ = v12
    _ = v9
    return 0
}
func genFn1(p0 bool, p1 int64, p2 int8) uint {
    p1 -= 9223372036854775807
    if (!(!p0)) {
        var v13 int64 = (p1 >> ((2) & 63))
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 428
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 42
    var v3 uint32 = 42
    var v14 uint = (^(v2))
    if (!(float64(1e10) > ((3.14159 + 1.5) + 0.0))) {
        if (("foo bar" + "a") < ("" + "a")) {
            switch 9223372036854775807 {
            case 128:
                v15 := GenS1{F0: 0, F1: v3, F2: -1}
                _ = v15
            case 2147483647:
                if true {
                    fmt.Println(v14)
                } else {
                    if (true || true) {
                        v16 := GenS1{F0: -1, F1: v3, F2: 42}
                        var v17 uint16 = 65535
                        _ = v16
                        _ = v17
                    } else {
                        v18 := []int{}
                        for fi19 := 0; fi19 < 15; fi19++ {
                            var v20 uint = v14
                            _ = fi19
                            _ = v20
                        }
                        if true {
                            var v21 int8 = -1
                            _ = v21
                        } else {
                            v2 = 874
                        }
                        _ = v18
                    }
                }
            }
            switch 42 {
            case -9223372036854775808:
                v2 = (v2 >> (((v2 - (v2 - v2))) & 63))
            }
        } else {
            for fi22 := 0; fi22 < 11; fi22++ {
                switch "hello" {
                case "":
                    var v23 uint = (v2 | v2)
                    _ = v23
                case "\n\t":
                    for fi24 := 0; fi24 < 5; fi24++ {
                        v25 := []int{}
                        if (true && (false || false)) {
                            if true {
                                var v26 uint16 = 578
                                var v27 int64 = -1
                                switch (fi22 >> (((v2 | uint(v3))) & 63)) {
                                case 767:
                                    var v28 int16 = 1
                                    _ = v28
                                }
                                v29 := GenS2{F0: (fi24 * fi24), F1: 42}
                                if (!false) {
                                    v3 += v3
                                    v30 := GenS1{F0: 1, F1: (v3 << ((uint(v26)) & 63)), F2: 417}
                                    v14 = (v2 * (v14 / (((v2 & (^(v2)))) | 1)))
                                    v25 = append(v25, fi24)
                                    if (-32768 >= 293) {
                                        v26 = (^(v26))
                                        if false {
                                            var v31 string = "foo bar"
                                            v32 := v1[fi22]
                                            _ = v31
                                            _ = v32
                                        } else {
                                            v1[fi24] = false
                                            if true {
                                                _ = len(v1)
                                            } else {
                                                if (128 <= 255) {
                                                    if len(v25) > 0 {
                                                        v25[0] = fi22
                                                    }
                                                    if (((v3 / ((v3) | 1)) <= v3) || false) {
                                                        v14 += (v2 & v2)
                                                        if len(v25) > 0 {
                                                            v25[0] = int(v2)
                                                        }
                                                    } else {
                                                        var v33 uint64 = 838
                                                        if true {
                                                            fmt.Println(v26, v27)
                                                            if false {
                                                                var v34 float32 = ((1e10 + ((0.0 + (3.14159 / 1.5)) / float32(-2.25))) * ((3.14159 / (-2.25 + 0.0)) + 0.0))
                                                                v27 = (v27 | -389)
                                                                var v35 float32 = v34
                                                                _ = v34
                                                                _ = v35
                                                            } else {
                                                                var v36 int32 = 0
                                                                _ = v36
                                                            }
                                                            if len(v25) > 0 {
                                                                v25[0] = (fi24 << ((v2) & 63))
                                                            }
                                                        } else {
                                                            delete(v1, (fi22 % ((9223372036854775807) | 1)))
                                                        }
                                                        _ = v33
                                                    }
                                                    fmt.Println(v14)
                                                    var v37 int32 = -46
                                                    _ = v37
                                                }
                                                v38 := GenS2{F0: 32767, F1: 2147483647}
                                                _ = v38
                                            }
                                        }
                                        switch (fi24 >> (((v2 % ((v2) | 1))) & 63)) {
                                        case 1:
                                            switch "hello" {
                                            case "":
                                                fmt.Println(v14)
                                                v3 = 4294967295
                                                if len(v25) > 0 {
                                                    v25[0] = (fi24 ^ -506)
                                                }
                                            case "\n\t":
                                                v39 := GenS1{F0: -1, F1: 42, F2: 0}
                                                _ = v39
                                            }
                                        default:
                                            v27 = (v27 | (v27 ^ 42))
                                        }
                                    }
                                    _ = v30
                                }
                                _ = v26
                                _ = v27
                                _ = v29
                            }
                        } else {
                            fmt.Println(fi22)
                            if len(v25) >= 2 {
                                v25 = v25[1:len(v25)]
                            }
                            v2 += (^(v14))
                            var v40 uint8 = 141
                            _ = v40
                        }
                        v41 := GenS2{F0: fi24, F1: 50}
                        _ = fi24
                        _ = v25
                        _ = v41
                    }
                case "hello":
                    for fi42 := 0; fi42 < 3; fi42++ {
                        v3 = uint32(v2)
                        _ = fi42
                    }
                default:
                    v14 = 2
                }
                v43 := GenS2{F0: fi22, F1: -595}
                v44 := GenS2{F0: 0, F1: 2147483647}
                _ = fi22
                _ = v43
                _ = v44
            }
            v3 -= v3
            var v45 int = 9223372036854775807
            _ = v45
        }
        if (1 == 0) {
            var v46 string = ""
            v47 := GenS2{F0: 9223372036854775807, F1: 42}
            if true {
                v48 := []int{286, -1, 654, 42}
                v46 = "foo bar"
                _ = v48
            }
            _ = v46
            _ = v47
        } else {
            if true {
                for fi49 := 0; fi49 < 9; fi49++ {
                    for fi50 := 0; fi50 < 2; fi50++ {
                        var v51 bool = false
                        _ = fi50
                        _ = v51
                    }
                    _ = fi49
                }
            } else {
                switch ("foo bar" + (("\n\t" + "\n\t") + "foo bar")) {
                case "a":
                    var v52 uint = 4294967295
                    _ = v52
                case "":
                    var v53 int8 = -23
                    _ = v53
                default:
                    var v54 uint = v14
                    _ = v54
                }
                for fi55 := 0; fi55 < 13; fi55++ {
                    for fi56 := 0; fi56 < 4; fi56++ {
                        var v57 uint16 = 42
                        v58 := v1[(fi56 << (((v14 + (v2 << ((v14) & 63)))) & 63))]
                        v59 := GenS1{F0: 1, F1: v3, F2: int32(v3)}
                        v1[fi56] = v58
                        if v58 {
                            v58 = true
                        } else {
                            var v60 float64 = float64(0.0)
                            if false {
                                v58 = (false && false)
                                v61 := []int{fi55, (fi55 / ((128) | 1))}
                                v62 := GenS3{F0: 9223372036854775807, F1: genFn1(true, 0, 100)}
                                _ = v61
                                _ = v62
                            } else {
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(v3, v14)
                                }
                                var v63 uint8 = 128
                                _ = v63
                            }
                            v64 := GenS1{F0: -128, F1: 160, F2: 283}
                            _ = v60
                            _ = v64
                        }
                        _ = fi56
                        _ = v57
                        _ = v58
                        _ = v59
                    }
                    v3 = 0
                    _ = fi55
                }
            }
            fmt.Println(v2, v14, v14)
        }
    } else {
        v3 = (v3 - v3)
        for fi65 := 0; fi65 < 17; fi65++ {
            var v66 uint16 = 0
            for fi67 := 0; fi67 < 2; fi67++ {
                var v68 int64 = int64(v3)
                v1[fi67] = (!false)
                _ = len(v1)
                v69 := GenS4{F0: 0, F1: v66, F2: v68}
                _ = fi67
                _ = v68
                _ = v69
            }
            v70 := GenS3{F0: -1, F1: uint(v66)}
            _ = fi65
            _ = v66
            _ = v70
        }
    }
    var v71 int16 = 195
    for fi72 := 0; fi72 < 5; fi72++ {
        v73 := GenS1{F0: 0, F1: v3, F2: 2147483647}
        var v74 uint = (v2 | (v2 + uint(v3)))
        v75 := []int{fi72, (fi72 >> (((v2 << ((0) & 63))) & 63)), 32767}
        _ = fi72
        _ = v73
        _ = v74
        _ = v75
    }
    if ((v71 % ((v71) | 1)) >= 42) {
        for fi76 := 0; fi76 < 20; fi76++ {
            v77 := []int{}
            _ = fi76
            _ = v77
        }
    }
    var v78 float32 = 1.5
    switch "foo bar" {
    case "":
        switch -1 {
        case -1:
            v71 = (v71 % (((v71 & -32768)) | 1))
        case 356:
            v78 = v78
        }
    case "foo bar":
        v79 := GenS5{F0: v3, F1: 1}
        _ = v79
    case "a":
        switch "foo bar" {
        case "hello":
            if ("" != "hello") {
                if false {
                    var v80 int64 = 9223372036854775807
                    if false {
                        switch 4294967295 {
                        case -964:
                            if false {
                                if (0.0 < 1.5) {
                                    if ((42 > 1) || true) {
                                        if (uint8(v14) < 2) {
                                            for fi81 := 0; fi81 < 18; fi81++ {
                                                for fi82 := 0; fi82 < 13; fi82++ {
                                                    fmt.Println(v3)
                                                    _ = fi82
                                                }
                                                _ = fi81
                                            }
                                            v80 = v80
                                        } else {
                                            if (((float32(1.5) * 1.5) / (float32(float64(float32(1e10))) / 0.1)) < (float32(1e10) / 0.1)) {
                                                for fi83 := 0; fi83 < 20; fi83++ {
                                                    v71 -= (v71 >> ((643) & 63))
                                                    v3 = (^(v3))
                                                    v84 := GenS1{F0: int8(v80), F1: (v3 % (((v3 & (v3 / (((v3 / (((^(v3))) | 1))) | 1)))) | 1)), F2: -2147483648}
                                                    _ = fi83
                                                    _ = v84
                                                }
                                            }
                                        }
                                        for fi85 := 0; fi85 < 20; fi85++ {
                                            v3 -= 2147483647
                                            var v86 int32 = 42
                                            v87 := []int{fi85, (^(fi85))}
                                            _ = fi85
                                            _ = v86
                                            _ = v87
                                        }
                                        v78 += v78
                                    }
                                    for fi88 := 0; fi88 < 17; fi88++ {
                                        switch (^(fi88)) {
                                        case 1:
                                            for fi89 := 0; fi89 < 11; fi89++ {
                                                var v90 float64 = 0.0
                                                v1[(fi89 % (((-(fi88))) | 1))] = true
                                                if (false && (false && false)) {
                                                    v91 := GenS6{F0: (^(v71)), F1: "", F2: (v80 * int64(fi89))}
                                                    v92 := []int{fi89, (-(fi89))}
                                                    _ = v91
                                                    _ = v92
                                                } else {
                                                    v93 := GenS4{F0: (v80 & v80), F1: 928, F2: -988}
                                                    var v94 int64 = (v80 & (-(v80)))
                                                    _ = v93
                                                    _ = v94
                                                }
                                                v3 = v3
                                                v95 := GenS1{F0: 100, F1: 0, F2: 1}
                                                var v96 uint = 297
                                                var v97 uint32 = v3
                                                _ = fi89
                                                _ = v90
                                                _ = v95
                                                _ = v96
                                                _ = v97
                                            }
                                        case -1:
                                            switch fi88 {
                                            case 127:
                                                var v98 bool = false
                                                _ = v98
                                            case 707:
                                                if (true || false) {
                                                    v71 -= v71
                                                } else {
                                                    var v99 uint64 = 985
                                                    fmt.Println(v78)
                                                    if (!false) {
                                                        for fi100 := 0; fi100 < 11; fi100++ {
                                                            fmt.Println(v14, v99, v3)
                                                            switch 32767 {
                                                            case 412:
                                                                v101 := GenS7{F0: fi100, F1: v71, F2: v80}
                                                                _ = v101
                                                            default:
                                                                switch fi88 {
                                                                case 2147483647:
                                                                    switch fi88 {
                                                                    case 2147483647:
                                                                        fmt.Println(fi100, fi88)
                                                                        v102 := GenS8{F0: 0, F1: 18446744073709551615, F2: 100}
                                                                        _ = v102
                                                                    case 32768:
                                                                        if true {
                                                                            if false {
                                                                                switch "\n\t" {
                                                                                case "\n\t":
                                                                                    if false {
                                                                                        v71 = 1000
                                                                                    }
                                                                                default:
                                                                                    v14 = v2
                                                                                }
                                                                            } else {
                                                                                var v103 uint64 = (v99 - (^(v99)))
                                                                                switch 32767 {
                                                                                case 9223372036854775807:
                                                                                    v80 -= v80
                                                                                case 128:
                                                                                    v104 := GenS9{F0: v99, F1: 945, F2: ("foo bar" + "\n\t")}
                                                                                    _ = v104
                                                                                default:
                                                                                    v105 := GenS4{F0: v80, F1: 18, F2: (v80 >> (((v14 << ((v14) & 63))) & 63))}
                                                                                    _ = v105
                                                                                }
                                                                                _ = v103
                                                                            }
                                                                        }
                                                                    case -260:
                                                                        v106 := GenS1{F0: int8(v80), F1: (^(v3)), F2: -711}
                                                                        _ = v106
                                                                    }
                                                                case 406:
                                                                    v1[(fi88 % ((-9223372036854775808) | 1))] = false
                                                                }
                                                            }
                                                            v99 = (v99 & (v99 % ((18446744073709551615) | 1)))
                                                            _ = fi100
                                                        }
                                                    } else {
                                                        for fi107 := 0; fi107 < 11; fi107++ {
                                                            _ = len(v1)
                                                            _ = fi107
                                                        }
                                                    }
                                                    _ = v99
                                                }
                                            default:
                                                fmt.Println(v3)
                                            }
                                        }
                                        for fi108 := 0; fi108 < 18; fi108++ {
                                            var v109 uint32 = (^(v3))
                                            _ = fi108
                                            _ = v109
                                        }
                                        v110 := GenS10{F0: 42}
                                        _ = fi88
                                        _ = v110
                                    }
                                    var v111 int64 = 0
                                    _ = v111
                                }
                                var v112 int8 = 0
                                _ = v112
                            }
                        default:
                            switch "" {
                            case "":
                                if (v78 != (((0.0 * (1.5 - 0.1)) / v78) * v78)) {
                                    var v113 float64 = (((0.0 - (float64(3.14159) - 1.5)) * (0.0 + 1e10)) + -2.25)
                                    v80 -= (v80 | (v80 - v80))
                                    switch -1 {
                                    case -2147483648:
                                        fmt.Println(v71, v3, v71)
                                    case 4294967295:
                                        if (!(v113 >= ((v113 / -2.25) / v113))) {
                                            v1[4294967295] = true
                                            if (genFn1(((v80 % ((0) | 1)) >= (v80 - v80)), v80, 1) >= (^(v2))) {
                                                var v114 uint16 = 2
                                                _ = v114
                                            } else {
                                                fmt.Println(v14, v80)
                                                v113 = v113
                                                var v115 bool = ((1 == -5) || ((v80 ^ (v80 * (v80 & 0))) < (v80 >> ((819) & 63))))
                                                _ = v115
                                            }
                                        }
                                    case -270:
                                        fmt.Println(v78, v113)
                                    }
                                    _ = v113
                                }
                            case "\n\t":
                                for fi116 := 0; fi116 < 14; fi116++ {
                                    switch fi116 {
                                    case -1:
                                        var v117 int16 = v71
                                        _ = v117
                                    case 128:
                                        v118 := v1[(fi116 ^ (-(fi116)))]
                                        _ = v118
                                    case 0:
                                        v119 := GenS3{F0: (v80 + (v80 / ((v80) | 1))), F1: 2}
                                        _ = v119
                                    default:
                                        fmt.Println(v3)
                                    }
                                    var v120 int16 = v71
                                    if true {
                                        for fi121 := 0; fi121 < 9; fi121++ {
                                            fmt.Println(v78)
                                            fmt.Println(fi121, v120)
                                            fmt.Println(v3)
                                            _ = fi121
                                        }
                                        var v122 int = fi116
                                        fmt.Println(v3)
                                        _ = v122
                                    } else {
                                        var v123 string = ""
                                        v14 = v2
                                        _ = v123
                                    }
                                    _ = fi116
                                    _ = v120
                                }
                            }
                        }
                        v124 := []int{2147483647, 0}
                        var v125 bool = true
                        _ = v124
                        _ = v125
                    }
                    _ = v80
                }
                fmt.Println(v78)
                switch 9223372036854775807 {
                case -9223372036854775808:
                    v3 = 2
                default:
                    v3 += v3
                }
            } else {
                v71 = (v71 | 355)
            }
        case "\n\t":
            fmt.Println(v3)
        case "":
            for fi126 := 0; fi126 < 12; fi126++ {
                delete(v1, fi126)
                for fi127 := 0; fi127 < 10; fi127++ {
                    v3 = (v3 % ((65535) | 1))
                    _ = fi127
                }
                _ = fi126
            }
        default:
            var v128 uint64 = 291
            _ = v128
        }
    }
    var v129 uint8 = 42
    v130 := GenS10{F0: uint8(v3)}
    if (false && (!(!(!true)))) {
        delete(v1, 32767)
        if false {
            var v131 uint = genFn0()
            var v132 uint32 = v3
            v133 := []int{}
            _ = v131
            _ = v132
            _ = v133
        } else {
            v3 += (v3 * (^(v3)))
        }
    } else {
        _ = len(v1)
        if true {
            v134 := v1[-2147483648]
            v134 = v134
            _ = v134
        } else {
            v14 += (v2 - (v14 * 2))
            v135 := []int{4294967295, 0, 9223372036854775807}
            if len(v135) >= 2 {
                v135 = v135[1:len(v135)]
            }
            _ = v135
        }
        var v136 uint = 18446744073709551615
        _ = v136
    }
    v137 := GenS11{F0: 0, F1: ("foo bar" + "a"), F2: -328}
    if (!(("a" <= "foo bar") && ((180 < -9223372036854775808) && (9223372036854775807 < 934)))) {
        v138 := []int{127, -2147483648, 32767}
        _ = v138
    } else {
        if true {
            v129 -= (v129 >> (((^(v2))) & 63))
            if (!(!true)) {
                if false {
                    var v139 int8 = 100
                    v140 := []int{4294967295, -218, -1}
                    for fi141 := 0; fi141 < 4; fi141++ {
                        var v142 uint = v2
                        var v143 uint8 = 1
                        _ = fi141
                        _ = v142
                        _ = v143
                    }
                    _ = v139
                    _ = v140
                }
                var v144 int8 = 99
                switch -9223372036854775808 {
                case 127:
                    switch 42 {
                    case -2147483648:
                        v145 := []int{}
                        _ = v145
                    }
                case 0:
                    if true {
                        for fi146 := 0; fi146 < 4; fi146++ {
                            for fi147 := 0; fi147 < 11; fi147++ {
                                if false {
                                    var v148 float32 = v78
                                    var v149 bool = (!false)
                                    var v150 uint64 = 172
                                    v2 += 65535
                                    _ = v148
                                    _ = v149
                                    _ = v150
                                }
                                fmt.Println(v14)
                                _ = fi147
                            }
                            if false {
                                for fi151 := 0; fi151 < 18; fi151++ {
                                    v152 := []int{fi151, (fi146 ^ -828), -1}
                                    fmt.Println(fi146, v78)
                                    v129 = v129
                                    _ = len(v1)
                                    v129 += uint8(v3)
                                    _ = fi151
                                    _ = v152
                                }
                            }
                            _ = fi146
                        }
                        var v153 uint16 = 244
                        _ = v153
                    }
                default:
                    v154 := GenS4{F0: 42, F1: 953, F2: 42}
                    _ = v154
                }
                _ = v144
            }
            v155 := GenS12{F0: -114, F1: "hello", F2: 0}
            _ = v155
        } else {
            if false {
                if ((^(v2)) > v2) {
                    for fi156 := 0; fi156 < 3; fi156++ {
                        v14 = uint(v3)
                        _ = fi156
                    }
                } else {
                    var v157 uint8 = 107
                    if (18446744073709551615 == 724) {
                        for fi158 := 0; fi158 < 2; fi158++ {
                            v159 := GenS13{F0: -9223372036854775808}
                            _ = fi158
                            _ = v159
                        }
                        for fi160 := 0; fi160 < 4; fi160++ {
                            for fi161 := 0; fi161 < 3; fi161++ {
                                v162 := GenS10{F0: v129}
                                var v163 float64 = float64((1.5 - v78))
                                _ = fi161
                                _ = v162
                                _ = v163
                            }
                            _ = fi160
                        }
                        for fi164 := 0; fi164 < 19; fi164++ {
                            v3 = 0
                            var v165 uint8 = (^(v129))
                            if false {
                                switch fi164 {
                                case 2147483647:
                                    for fi166 := 0; fi166 < 20; fi166++ {
                                        var v167 float32 = v78
                                        _ = fi166
                                        _ = v167
                                    }
                                }
                                fmt.Println(fi164, v3)
                                switch (fi164 ^ (fi164 | fi164)) {
                                case -1:
                                    switch "" {
                                    case "a":
                                        if (255 > 322) {
                                            if false {
                                                v168 := GenS9{F0: 0, F1: 65535, F2: "hello"}
                                                v14 = (^(v14))
                                                _ = v168
                                            } else {
                                                v2 = 2
                                                v2 = v14
                                                fmt.Println(v165, v2, v129)
                                            }
                                            var v169 uint64 = 0
                                            if (v71 > v71) {
                                                if (!(65535 <= (v3 | (v3 / (((^(v3))) | 1))))) {
                                                    var v170 uint64 = 18446744073709551615
                                                    v165 -= v165
                                                    v169 = (v170 | v169)
                                                    _ = v170
                                                }
                                                switch (^(fi164)) {
                                                case -1:
                                                    v171 := GenS14{F0: (fi164 / ((fi164) | 1)), F1: uint64(v165), F2: uint64(v129)}
                                                    _ = v171
                                                case -2147483648:
                                                    var v172 uint16 = uint16(v2)
                                                    _ = v172
                                                case 32768:
                                                    v157 = 1
                                                }
                                            } else {
                                                v3 += v3
                                                var v173 uint16 = uint16(v3)
                                                _ = v173
                                            }
                                            _ = v169
                                        } else {
                                            v174 := GenS15{F0: 0, F1: fi164, F2: uint8(v3)}
                                            _ = v174
                                        }
                                    case "foo bar":
                                        if (("\n\t" + "\n\t") == ("hello" + "")) {
                                            v157 = v157
                                            switch (fi164 | (fi164 ^ fi164)) {
                                            case -1:
                                                if false {
                                                    if false {
                                                        v14 = v14
                                                    } else {
                                                        v71 = (^(v71))
                                                    }
                                                    switch "foo bar" {
                                                    case "hello":
                                                        if true {
                                                            switch "a" {
                                                            case "foo bar":
                                                                switch 9223372036854775807 {
                                                                case 584:
                                                                    switch ("a" + "\n\t") {
                                                                    case "foo bar":
                                                                        switch fi164 {
                                                                        case -431:
                                                                            var v175 uint32 = v3
                                                                            _ = v175
                                                                        }
                                                                    case "a":
                                                                        fmt.Println(v157, v3)
                                                                    }
                                                                }
                                                            case "hello":
                                                                if (!false) {
                                                                    v165 = v157
                                                                }
                                                            case "a":
                                                                switch -2147483648 {
                                                                case 2147483647:
                                                                    v2 = genFn1(true, 1, -1)
                                                                case 283:
                                                                    var v176 uint64 = 42
                                                                    _ = v176
                                                                default:
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        v177 := GenS7{F0: (fi164 | 283), F1: v71, F2: -9223372036854775808}
                                                                        _ = v177
                                                                    case "foo bar":
                                                                        fmt.Println(v165)
                                                                    case "hello":
                                                                        v2 = v14
                                                                    default:
                                                                        var v178 int8 = -51
                                                                        _ = v178
                                                                    }
                                                                }
                                                            }
                                                            v2 += (v2 ^ v2)
                                                            v179 := []int{(^(fi164)), 9223372036854775807}
                                                            _ = v179
                                                        } else {
                                                            var v180 uint16 = 0
                                                            var v181 uint16 = v180
                                                            _ = v180
                                                            _ = v181
                                                        }
                                                    case "a":
                                                        fmt.Println(fi164, v2)
                                                    case "\n\t":
                                                        if false {
                                                            v78 = 1e10
                                                            v78 = (v78 * v78)
                                                        }
                                                    }
                                                } else {
                                                    switch 757 {
                                                    case 211:
                                                        v2 = v2
                                                    }
                                                }
                                            case 9223372036854775807:
                                                v3 = 1
                                            case 225:
                                                var v182 int8 = 1
                                                _ = v182
                                            default:
                                                v183 := GenS16{F0: 1, F1: 18446744073709551615}
                                                _ = v183
                                            }
                                            v184 := []int{-877, fi164, -691}
                                            _ = v184
                                        }
                                    case "":
                                        v129 = v165
                                    default:
                                        v165 -= v129
                                    }
                                case -934:
                                    v185 := GenS5{F0: (v3 | v3), F1: 0}
                                    _ = v185
                                default:
                                    for fi186 := 0; fi186 < 1; fi186++ {
                                        v187 := GenS10{F0: v157}
                                        _ = fi186
                                        _ = v187
                                    }
                                }
                            }
                            _ = fi164
                            _ = v165
                        }
                    } else {
                        v188 := GenS7{F0: -35, F1: v71, F2: 1}
                        v78 += 0.0
                        v189 := GenS2{F0: -687, F1: -683}
                        _ = v188
                        _ = v189
                    }
                    v190 := GenS1{F0: -128, F1: 742, F2: -254}
                    _ = v157
                    _ = v190
                }
                if false {
                    var v191 float32 = ((1.5 + v78) / v78)
                    var v192 uint = v14
                    var v193 int = 42
                    _ = v191
                    _ = v192
                    _ = v193
                } else {
                    fmt.Println(v3, v14)
                    v194 := GenS12{F0: -74, F1: "foo bar", F2: 42}
                    if (42 >= uint(v129)) {
                        delete(v1, -1)
                        var v195 uint = v14
                        _ = v195
                    }
                    _ = v194
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v129:", v129)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v71:", v71)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v129
    _ = v130
    _ = v137
    _ = v14
    _ = v2
    _ = v3
    _ = v71
    _ = v78
}
