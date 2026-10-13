package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 int64; F1 int16; F2 int }
type GenS7 struct { F0 int64; F1 uint16; F2 int64 }
func genFn0(p0 bool) uint8 {
    var v6 int = 128
    p0 = true
    v7 := []int{311, (v6 * v6), v6}
    _ = p0
    _ = v6
    _ = v7
    return 190
}
func genFn1(p0 int32, p1 int32, p2 int8) uint {
    v8 := []int{127, -837}
    switch 0 {
    case 32767:
        fmt.Println(p1, p1)
    case 42:
        if false {
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := GenS1{F0: genFn0(true)}
                for fi11 := 0; fi11 < 7; fi11++ {
                    var v12 uint16 = 485
                    _ = fi11
                    _ = v12
                }
                var v13 uint64 = 945
                _ = fi9
                _ = v10
                _ = v13
            }
            v8 = append(v8, -644)
        } else {
            if false {
                fmt.Println(p1, p1)
            }
            v14 := GenS1{F0: 128}
            v15 := GenS2{F0: 2}
            _ = v14
            _ = v15
        }
    default:
        fmt.Println(p2, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 4294967295
}

func main() {
    var v1 float64 = 1.5
    var v2 uint = 11
    v3 := []int{}
    v4 := make(map[uint32]uint32)
    var v5 string = (("foo bar" + (("a" + "a") + "hello")) + "foo bar")
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2, v1, v2)
    var v16 float64 = 1.5
    var v17 uint32 = 65535
    var v18 uint32 = 1
    v18 = 0
    var v19 int8 = -128
    if (90 <= 2) {
        for fi20 := 0; fi20 < 16; fi20++ {
            v5 = ""
            var v21 uint8 = 0
            v17 = (v17 ^ v18)
            _ = fi20
            _ = v21
        }
        if false {
            fmt.Println(v5)
            if true {
                switch v5 {
                case "\n\t":
                    fmt.Println(v18)
                default:
                    for fi22 := 0; fi22 < 5; fi22++ {
                        if false {
                            var v23 uint32 = (v17 & 65535)
                            for fi24 := 0; fi24 < 6; fi24++ {
                                _ = len(v4)
                                if len(v3) > 0 {
                                    v3[0] = (fi24 * (fi24 ^ (fi22 ^ (fi24 + fi22))))
                                }
                                switch v5 {
                                case "a":
                                    switch ((((v5 + "") + v5) + "") + "") {
                                    case "":
                                        fmt.Println(v18, v16)
                                    case "hello":
                                        var v25 float64 = 0.1
                                        _ = v25
                                    case "a":
                                        v4[(v18 ^ (v23 | v17))] = v17
                                    default:
                                        v16 = 1.5
                                    }
                                default:
                                    if (((false || true) && (!false)) && (243 == -9223372036854775808)) {
                                        fmt.Println(v17, v17, fi22)
                                    } else {
                                        v4[v18] = 2147483647
                                        v17 += (^(v17))
                                        fmt.Println(v19)
                                        delete(v4, (^(v18)))
                                        var v26 bool = (2 < 65535)
                                        _ = v26
                                    }
                                }
                                v19 += v19
                                if false {
                                    v27 := GenS2{F0: v18}
                                    v5 = "\n\t"
                                    delete(v4, (v18 - v18))
                                    _ = v27
                                } else {
                                    v23 = (v18 - v17)
                                    var v28 int = fi22
                                    switch (-(v28)) {
                                    case 9223372036854775807:
                                        v29 := GenS1{F0: 0}
                                        _ = v29
                                    }
                                    _ = len(v4)
                                    var v30 int8 = -48
                                    _ = v28
                                    _ = v30
                                }
                                _ = fi24
                            }
                            _ = v23
                        } else {
                            for fi31 := 0; fi31 < 11; fi31++ {
                                if len(v3) > 0 {
                                    v3[0] = 9223372036854775807
                                }
                                fmt.Println(fi22, v18, fi31)
                                v32 := v4[(v18 >> ((812) & 63))]
                                _ = fi31
                                _ = v32
                            }
                        }
                        v2 = genFn1(2147483647, 652, (v19 / ((v19) | 1)))
                        _ = fi22
                    }
                }
                v5 = ""
                switch -2147483648 {
                case 128:
                    var v33 int64 = -1
                    _ = v33
                case -2147483648:
                    var v34 uint = (^(v2))
                    _ = v34
                default:
                    if (1 == 9223372036854775807) {
                        var v35 int64 = 0
                        switch 398 {
                        case -2147483648:
                            var v36 uint = (v2 + genFn1(-348, int32(v18), v19))
                            _ = v36
                        case 32768:
                            if false {
                                var v37 int = -9223372036854775808
                                var v38 float32 = 3.14159
                                v3 = append(v3, (-(v37)))
                                _ = v37
                                _ = v38
                            }
                        case 438:
                            for fi39 := 0; fi39 < 18; fi39++ {
                                if false {
                                    var v40 float32 = 0.0
                                    _ = len(v3)
                                    fmt.Println(v1, v2)
                                    _ = v40
                                }
                                v41 := GenS1{F0: 164}
                                _ = fi39
                                _ = v41
                            }
                        default:
                            var v42 float64 = ((v1 + v1) + (((1.5 + v16) / v16) * (v1 + ((3.14159 / v16) + 0.1))))
                            _ = v42
                        }
                        v43 := GenS1{F0: uint8(v2)}
                        _ = v35
                        _ = v43
                    } else {
                        switch 4294967295 {
                        case 572:
                            for fi44 := 0; fi44 < 11; fi44++ {
                                if (!true) {
                                    var v45 int16 = -1
                                    _ = v45
                                }
                                for fi46 := 0; fi46 < 6; fi46++ {
                                    if true {
                                        fmt.Println(v16)
                                    } else {
                                        v47 := GenS2{F0: (v17 / (((v18 << ((2) & 63))) | 1))}
                                        v48 := GenS3{F0: 2, F1: v17}
                                        var v49 int32 = -193
                                        _ = v47
                                        _ = v48
                                        _ = v49
                                    }
                                    if (!(uint8(v2) > 0)) {
                                        v50 := v4[(v17 - v18)]
                                        v16 = v1
                                        _ = v50
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = -950
                                    }
                                    _ = fi46
                                }
                                v5 = ("\n\t" + "a")
                                _ = fi44
                            }
                        case -824:
                            fmt.Println(v5)
                        default:
                            for fi51 := 0; fi51 < 3; fi51++ {
                                var v52 string = ("hello" + (v5 + v5))
                                _ = fi51
                                _ = v52
                            }
                        }
                        var v53 int64 = -934
                        for fi54 := 0; fi54 < 2; fi54++ {
                            delete(v4, (^(v17)))
                            if true {
                                for fi55 := 0; fi55 < 2; fi55++ {
                                    var v56 float64 = v1
                                    v2 += (v2 + (v2 ^ v2))
                                    _ = fi55
                                    _ = v56
                                }
                                v3 = append(v3, fi54)
                                v57 := GenS4{F0: fi54}
                                _ = v57
                            } else {
                                var v58 uint = genFn1(-1, 0, (v19 * 100))
                                v59 := GenS4{F0: (^(fi54))}
                                var v60 uint8 = 255
                                _ = v58
                                _ = v59
                                _ = v60
                            }
                            _ = fi54
                        }
                        _ = v53
                    }
                }
            }
            if false {
                if (false || (-2147483648 == -686)) {
                    var v61 float64 = (((((v16 / 3.14159) + float64(1.5)) * float64(-2.25)) + 1e10) / (v16 * 0.0))
                    _ = v61
                }
                if false {
                    v4[(v18 | 1)] = v18
                    for fi62 := 0; fi62 < 9; fi62++ {
                        fmt.Println(fi62, v2, fi62)
                        if false {
                            var v63 float64 = -2.25
                            v64 := GenS4{F0: fi62}
                            switch (v5 + "") {
                            case "hello":
                                v2 += (v2 * (^(v2)))
                            }
                            _ = v63
                            _ = v64
                        }
                        var v65 int8 = (v19 & (v19 ^ (v19 * v19)))
                        _ = fi62
                        _ = v65
                    }
                    v3 = append(v3, 32768)
                } else {
                    if (true && (42 > 255)) {
                        for fi66 := 0; fi66 < 18; fi66++ {
                            var v67 uint64 = 4294967295
                            for fi68 := 0; fi68 < 13; fi68++ {
                                var v69 int = 32768
                                v70 := GenS2{F0: (v18 / ((v17) | 1))}
                                _ = fi68
                                _ = v69
                                _ = v70
                            }
                            _ = fi66
                            _ = v67
                        }
                        v71 := GenS4{F0: 32768}
                        switch ((v5 + v5) + v5) {
                        case "a":
                            switch v5 {
                            case "hello":
                                if true {
                                    switch "" {
                                    case "foo bar":
                                        if true {
                                            if false {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v72 bool = (!true)
                                                    _ = v72
                                                default:
                                                    v17 = v17
                                                }
                                                if (-1 != 0) {
                                                    fmt.Println(v5)
                                                    var v73 uint = v2
                                                    var v74 uint64 = 716
                                                    _ = v73
                                                    _ = v74
                                                }
                                                if false {
                                                    v75 := v4[(v17 | v17)]
                                                    _ = v75
                                                }
                                            }
                                            _ = len(v3)
                                        } else {
                                            fmt.Println(v18, v19, v16)
                                            var v76 int32 = -681
                                            _ = v76
                                        }
                                    }
                                }
                            case "a":
                                for fi77 := 0; fi77 < 20; fi77++ {
                                    var v78 bool = (v2 > (v2 + 0))
                                    v79 := GenS4{F0: (fi77 + (fi77 + (fi77 | fi77)))}
                                    switch v5 {
                                    case "":
                                        fmt.Println(v5, v2, v17)
                                    case "foo bar":
                                        if v78 {
                                            var v80 uint32 = v17
                                            _ = v80
                                        } else {
                                            var v81 float32 = -2.25
                                            _ = v81
                                        }
                                    }
                                    _ = fi77
                                    _ = v78
                                    _ = v79
                                }
                            }
                        case "hello":
                            v82 := GenS5{F0: (2 == 255), F1: 32767}
                            _ = v82
                        case "\n\t":
                            if true {
                                var v83 float32 = float32(1.5)
                                if true {
                                    if true {
                                        _ = len(v4)
                                        if len(v3) > 0 {
                                            v3[0] = 1
                                        }
                                        v84 := GenS1{F0: 1}
                                        _ = v84
                                    }
                                    fmt.Println(v83, v83)
                                }
                                _ = v83
                            } else {
                                fmt.Println(v17)
                            }
                        default:
                            for fi85 := 0; fi85 < 11; fi85++ {
                                if ((0.0 >= 1e10) || true) {
                                    switch "a" {
                                    case "a":
                                        fmt.Println(v2)
                                    case "foo bar":
                                        var v86 int32 = -606
                                        _ = v86
                                    case "hello":
                                        if false {
                                            fmt.Println(v2, v16, v5)
                                        }
                                    }
                                    for fi87 := 0; fi87 < 3; fi87++ {
                                        switch "foo bar" {
                                        case "a":
                                            if (-289 == 9223372036854775807) {
                                                v19 = 0
                                            } else {
                                                var v88 int8 = v19
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v88
                                            }
                                        }
                                        _ = fi87
                                    }
                                    if false {
                                        switch (fi85 * fi85) {
                                        case 0:
                                            v89 := GenS3{F0: 127, F1: (^(v17))}
                                            _ = v89
                                        case -9223372036854775808:
                                            v5 = v5
                                        case 2147483647:
                                            switch v5 {
                                            case "foo bar":
                                                var v90 int16 = -595
                                                _ = v90
                                            }
                                        }
                                        fmt.Println(v16, v1, v17)
                                        for fi91 := 0; fi91 < 7; fi91++ {
                                            if false {
                                                switch v5 {
                                                case "":
                                                    if false {
                                                        if false {
                                                            v19 = (-(v19))
                                                            _ = len(v3)
                                                            var v92 int64 = -9223372036854775808
                                                            if true {
                                                                if true {
                                                                    if true {
                                                                        if (false || true) {
                                                                            var v93 uint = (^(v2))
                                                                            v94 := GenS2{F0: (^(v17))}
                                                                            _ = v93
                                                                            _ = v94
                                                                        } else {
                                                                            if false {
                                                                                v95 := GenS5{F0: true, F1: -780}
                                                                                fmt.Println(v19)
                                                                                v96 := GenS4{F0: fi85}
                                                                                _ = v95
                                                                                _ = v96
                                                                            } else {
                                                                                v92 += (v92 << ((v2) & 63))
                                                                                fmt.Println(v19, v2, v1)
                                                                                v3 = append(v3, 633)
                                                                                v19 += (v19 >> ((v2) & 63))
                                                                                v19 = (v19 * (v19 << (((^(v2))) & 63)))
                                                                            }
                                                                            v97 := GenS1{F0: 111}
                                                                            _ = v97
                                                                        }
                                                                        if (!false) {
                                                                            if (-32768 < -32768) {
                                                                                v2 -= v2
                                                                                v5 = v5
                                                                            } else {
                                                                                fmt.Println(v17, v1, v2)
                                                                                var v98 float64 = (v1 + v1)
                                                                                _ = v98
                                                                            }
                                                                            switch "" {
                                                                            case "a":
                                                                                v19 -= (v19 + v19)
                                                                            default:
                                                                                switch v5 {
                                                                                case "foo bar":
                                                                                    v4[(v18 / (((^(v18))) | 1))] = v18
                                                                                case "hello":
                                                                                    if true {
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = 32767
                                                                                        }
                                                                                        var v99 uint16 = 922
                                                                                        switch "a" {
                                                                                        case "\n\t":
                                                                                            var v100 uint32 = (v18 ^ (v18 % (((^(v18))) | 1)))
                                                                                            _ = v100
                                                                                        case "hello":
                                                                                            _ = len(v4)
                                                                                        }
                                                                                        _ = v99
                                                                                    }
                                                                                    switch int(v2) {
                                                                                    case -954:
                                                                                        var v101 int64 = v92
                                                                                        _ = v101
                                                                                    case -2147483648:
                                                                                        var v102 int8 = -97
                                                                                        _ = v102
                                                                                    default:
                                                                                        v16 = (((v1 - v16) * (float64((float32(float64((3.14159 - -2.25))) + ((1e10 / (1.5 - (((3.14159 - float32(v16)) / 1e10) + (3.14159 - (1.5 / 1.0))))) * 1.5))) + ((float64(float32((1.5 * (((v16 + (((1e10 * v16) / float64(1e10)) / (((v16 + v16) + (v16 * ((v16 - v1) - ((0.0 * ((float64((3.14159 / float32(-2.25))) - v16) / v1)) * float64(float32((v16 * v1))))))) * v1))) / v1) + -2.25)))) - 1e10) * 1e10))) + (v1 / 1.0))
                                                                                    }
                                                                                }
                                                                            }
                                                                            v103 := GenS6{F0: v92, F1: -213, F2: -9223372036854775808}
                                                                            _ = v103
                                                                        }
                                                                    }
                                                                } else {
                                                                    v92 -= v92
                                                                }
                                                                if (128 > (fi91 ^ (fi91 % ((-1) | 1)))) {
                                                                    if (450 != v17) {
                                                                        fmt.Println(v5, v19, v16)
                                                                    }
                                                                }
                                                                v104 := GenS2{F0: (v17 % (((^(v18))) | 1))}
                                                                _ = v104
                                                            } else {
                                                                var v105 float32 = ((-2.25 + 3.14159) * 1.5)
                                                                _ = v105
                                                            }
                                                            _ = len(v3)
                                                            _ = v92
                                                        } else {
                                                            fmt.Println(v17, v16)
                                                        }
                                                        _ = len(v3)
                                                        v18 += v17
                                                    }
                                                case "a":
                                                    v4[452] = (^(v17))
                                                }
                                                switch (fi85 * (^(fi85))) {
                                                case 2147483647:
                                                    switch v5 {
                                                    case "\n\t":
                                                        if true {
                                                            v4[uint32(v2)] = v17
                                                            v18 += 2147483647
                                                        } else {
                                                            v106 := GenS5{F0: false, F1: fi91}
                                                            _ = v106
                                                        }
                                                    case "a":
                                                        var v107 string = ("" + v5)
                                                        _ = v107
                                                    default:
                                                        v108 := GenS1{F0: 1}
                                                        _ = v108
                                                    }
                                                    v4[(^(v17))] = v18
                                                    fmt.Println(fi91, v16, v5)
                                                }
                                            } else {
                                                if true {
                                                    var v109 uint = 166
                                                    var v110 uint16 = 946
                                                    _ = v109
                                                    _ = v110
                                                }
                                                v18 += v17
                                                v17 -= (v18 >> ((18446744073709551615) & 63))
                                                fmt.Println(fi91, v19, v18)
                                                fmt.Println(v1, v5)
                                            }
                                            _ = fi91
                                        }
                                    } else {
                                        if true {
                                            for fi111 := 0; fi111 < 19; fi111++ {
                                                var v112 uint64 = 0
                                                v113 := v4[v18]
                                                _ = fi111
                                                _ = v112
                                                _ = v113
                                            }
                                            var v114 uint8 = 239
                                            _ = v114
                                        } else {
                                            fmt.Println(v17, v1)
                                            var v115 uint16 = 32767
                                            for fi116 := 0; fi116 < 16; fi116++ {
                                                switch v5 {
                                                case "a":
                                                    var v117 uint = 42
                                                    _ = v117
                                                }
                                                fmt.Println(v17, v18, fi85)
                                                _ = fi116
                                            }
                                            _ = v115
                                        }
                                    }
                                } else {
                                    if false {
                                        v19 = v19
                                        for fi118 := 0; fi118 < 18; fi118++ {
                                            var v119 uint8 = genFn0(true)
                                            var v120 int32 = 0
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = fi118
                                            _ = v119
                                            _ = v120
                                        }
                                        var v121 int8 = v19
                                        _ = v121
                                    }
                                    v5 = v5
                                    for fi122 := 0; fi122 < 11; fi122++ {
                                        var v123 bool = false
                                        var v124 uint32 = (v18 + v18)
                                        if (v1 >= v16) {
                                            switch (^(fi85)) {
                                            case 768:
                                                if false {
                                                    var v125 string = v5
                                                    _ = v125
                                                }
                                            case 2147483647:
                                                v19 -= v19
                                            case 127:
                                                v126 := GenS5{F0: false, F1: (fi122 ^ fi122)}
                                                v16 = -2.25
                                                var v127 string = "hello"
                                                _ = v126
                                                _ = v127
                                            default:
                                                v123 = (v123 || (!(v18 < v18)))
                                            }
                                        }
                                        _ = fi122
                                        _ = v123
                                        _ = v124
                                    }
                                }
                                _ = fi85
                            }
                        }
                        _ = v71
                    } else {
                        v19 += (^(v19))
                        v128 := v4[(v17 + (v18 & 42))]
                        for fi129 := 0; fi129 < 20; fi129++ {
                            var v130 uint8 = 224
                            if true {
                                delete(v4, v17)
                                fmt.Println(v2, v17, v17)
                                if true {
                                    var v131 uint8 = genFn0(true)
                                    _ = v131
                                } else {
                                    if true {
                                        var v132 int = -214
                                        _ = v132
                                    }
                                    v130 = (v130 * 1)
                                }
                            } else {
                                v16 += v16
                            }
                            _ = fi129
                            _ = v130
                        }
                        _ = v128
                    }
                    if (!false) {
                        v18 += v17
                    }
                    if (true || true) {
                        var v133 uint16 = 255
                        switch v5 {
                        case "\n\t":
                            v2 = uint(v133)
                        default:
                            v133 = v133
                        }
                        if false {
                            if ((^(v133)) == v133) {
                                v2 = (v2 % ((255) | 1))
                            } else {
                                var v134 uint64 = 4294967295
                                v135 := GenS7{F0: -902, F1: 255, F2: 9223372036854775807}
                                _ = v134
                                _ = v135
                            }
                            var v136 uint16 = (v133 >> ((v2) & 63))
                            _ = v136
                        }
                        _ = v133
                    }
                }
                for fi137 := 0; fi137 < 20; fi137++ {
                    _ = len(v4)
                    _ = fi137
                }
            } else {
                v5 = ""
                for fi138 := 0; fi138 < 6; fi138++ {
                    var v139 float32 = (0.1 * 3.14159)
                    _ = fi138
                    _ = v139
                }
            }
        }
    }
    if true {
        var v140 bool = (-369 > 9223372036854775807)
        v17 -= (v18 & (v17 - 2147483647))
        v141 := v4[(v18 * (v17 * (v18 >> ((v2) & 63))))]
        _ = v140
        _ = v141
    } else {
        for fi142 := 0; fi142 < 10; fi142++ {
            v17 += v18
            _ = fi142
        }
        v19 = v19
    }
    if false {
        v143 := GenS2{F0: v18}
        _ = v143
    }
    switch "hello" {
    case "hello":
        v18 += v18
    default:
        _ = len(v4)
    }
    var v144 int64 = -9223372036854775808
    v145 := GenS2{F0: (v18 | v18)}
    switch 0 {
    case 2147483647:
        fmt.Println(v16, v18)
    case 9223372036854775807:
        if ((v17 & v18) <= v18) {
            v2 += genFn1(42, 0, v19)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch (((((v5 + "\n\t") + (v5 + v5)) + v5) + ("\n\t" + v5)) + ("hello" + "")) {
            case "a":
                var v146 uint32 = 2147483647
                _ = v146
            default:
                v18 -= v17
            }
        }
    case -687:
        if (!(-9223372036854775808 >= v144)) {
            fmt.Println(v2, v5, v17)
            switch -9223372036854775808 {
            case -655:
                var v147 int8 = 100
                _ = v147
            case -1:
                for fi148 := 0; fi148 < 13; fi148++ {
                    var v149 uint8 = 1
                    var v150 int64 = (v144 % ((v144) | 1))
                    fmt.Println(v1, v1, v1)
                    _ = fi148
                    _ = v149
                    _ = v150
                }
            case 128:
                if (((^(v18)) == v18) && false) {
                    if true {
                        v16 -= (v16 - v16)
                    } else {
                        fmt.Println(v2)
                    }
                    v5 = ((((v5 + "a") + ("" + v5)) + (v5 + v5)) + "a")
                    v18 -= v18
                }
            }
        }
    default:
        switch 32768 {
        case -2147483648:
            var v151 uint8 = uint8(v2)
            _ = v151
        case 625:
            v17 = v17
        case 9:
            v16 += v1
        default:
            var v152 uint8 = 255
            _ = v152
        }
    }
    var v153 int32 = 552
    fmt.Println("v1:", v1)
    fmt.Println("v144:", v144)
    fmt.Println("v153:", v153)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v144
    _ = v145
    _ = v153
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
