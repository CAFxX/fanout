package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64; F1 uint32 }
type GenS8 struct { F0 int8; F1 bool; F2 int }
type GenS9 struct { F0 uint16; F1 int8; F2 int64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int8; F1 int64; F2 uint32 }
type GenS12 struct { F0 int; F1 uint8 }
type GenS13 struct { F0 int16; F1 int64; F2 bool }
type GenS14 struct { F0 string; F1 int; F2 uint32 }
func genFn0(p0 int64, p1 int64, p2 int8) int64 {
    v4 := GenS1{F0: 127}
    fmt.Println(p1)
    for fi5 := 0; fi5 < 7; fi5++ {
        p0 = p0
        var v6 uint8 = 255
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p0 - (-(p1)))
}

func main() {
    v1 := []int{-257, 979, 32768}
    var v2 int32 = 488
    v3 := make(map[int]bool)
    v3[641] = ((42 > 71) && (true && (0 != 2147483647)))
    if len(v1) > 0 {
        v1[0] = 44
    }
    if true {
        var v7 uint64 = 1
        _ = v7
    } else {
        for fi8 := 0; fi8 < 5; fi8++ {
            if (93 <= 127) {
                for fi9 := 0; fi9 < 15; fi9++ {
                    v2 += -327
                    var v10 float32 = 0.1
                    var v11 int = fi9
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            } else {
                if (657 <= 2147483647) {
                    if true {
                        fmt.Println(fi8, fi8)
                        v2 -= v2
                    } else {
                        for fi12 := 0; fi12 < 16; fi12++ {
                            if false {
                                v13 := GenS1{F0: 188}
                                fmt.Println(fi8, fi8, v2)
                                var v14 uint64 = 2
                                _ = v13
                                _ = v14
                            } else {
                                v3[127] = true
                                var v15 uint64 = 205
                                _ = v15
                            }
                            fmt.Println(v2, fi8, fi8)
                            switch fi8 {
                            case 0:
                                v2 += (v2 % (((v2 + v2)) | 1))
                            case 4294967295:
                                v16 := GenS2{F0: (("foo bar" + ("\n\t" + ("a" + "\n\t"))) + (("a" + "a") + "hello")), F1: 2}
                                _ = v16
                            }
                            _ = fi12
                        }
                        for fi17 := 0; fi17 < 10; fi17++ {
                            v2 -= (v2 % ((v2) | 1))
                            v2 -= (v2 >> ((10) & 63))
                            _ = fi17
                        }
                    }
                    if (!true) {
                        v2 = 0
                        v2 = (^(v2))
                    } else {
                        if true {
                            for fi18 := 0; fi18 < 9; fi18++ {
                                if (!false) {
                                    v2 = 227
                                    if (!true) {
                                        v2 -= v2
                                        v3[fi8] = false
                                    }
                                    v19 := GenS1{F0: 0}
                                    _ = v19
                                }
                                _ = fi18
                            }
                            if true {
                                switch fi8 {
                                case 4294967295:
                                    if (!(false && false)) {
                                        if false {
                                            v3[(fi8 * (fi8 % ((-9223372036854775808) | 1)))] = (782 == 767)
                                            var v20 int16 = 0
                                            _ = v20
                                        }
                                    }
                                default:
                                    v21 := GenS1{F0: 33}
                                    _ = v21
                                }
                            }
                            var v22 int16 = -32768
                            _ = v22
                        }
                        var v23 bool = true
                        _ = v23
                    }
                }
            }
            for fi24 := 0; fi24 < 5; fi24++ {
                if len(v1) > 0 {
                    v1[0] = fi8
                }
                var v25 uint = 4294967295
                v25 += (v25 + v25)
                _ = fi24
                _ = v25
            }
            v2 = v2
            _ = fi8
        }
        v2 += (v2 & -2147483648)
    }
    var v26 int8 = 42
    var v27 uint32 = 892
    for fi28 := 0; fi28 < 8; fi28++ {
        switch (fi28 ^ fi28) {
        case 4294967295:
            if (true && true) {
                v2 = (^(v2))
                v29 := v3[fi28]
                _ = v29
            }
        case 32767:
            var v30 bool = false
            _ = v30
        case -1:
            for fi31 := 0; fi31 < 5; fi31++ {
                fmt.Println(v26)
                _ = fi31
            }
        default:
            v26 -= v26
        }
        if (v27 > (v27 >> ((42) & 63))) {
            v32 := GenS3{F0: 122}
            var v33 int8 = (v26 | v26)
            _ = v32
            _ = v33
        } else {
            if (38 <= 2) {
                if (-32768 != -1) {
                    delete(v3, (fi28 >> ((457) & 63)))
                    v34 := GenS4{F0: int64(v26), F1: (v27 & (v27 | (v27 | 42)))}
                    v2 -= (v2 * 2147483647)
                    _ = v34
                }
            }
        }
        if (139 > 914) {
            v35 := GenS2{F0: ("a" + "\n\t"), F1: 1}
            if true {
                v27 = 1
            }
            var v36 int8 = (v26 | (v26 + v26))
            _ = v35
            _ = v36
        }
        _ = fi28
    }
    if false {
        var v37 int8 = (v26 & -67)
        for fi38 := 0; fi38 < 7; fi38++ {
            v26 += v26
            var v39 int = (fi38 << ((758) & 63))
            if true {
                var v40 int = (fi38 | (v39 % ((fi38) | 1)))
                _ = v40
            } else {
                v41 := GenS3{F0: (v37 % (((-(v37))) | 1))}
                switch (fi38 % ((int(v2)) | 1)) {
                case -1:
                    v37 = v26
                case 27:
                    var v42 uint64 = 1
                    _ = v42
                default:
                    v1 = append(v1, (v39 | 9223372036854775807))
                }
                v39 = v39
                _ = v41
            }
            _ = fi38
            _ = v39
        }
        for fi43 := 0; fi43 < 12; fi43++ {
            v26 = (-(v26))
            if (0 <= 1) {
                var v44 int16 = -32768
                var v45 uint32 = 2
                var v46 uint = 186
                _ = v44
                _ = v45
                _ = v46
            } else {
                for fi47 := 0; fi47 < 7; fi47++ {
                    if true {
                        var v48 float32 = 0.0
                        _ = v48
                    } else {
                        fmt.Println(fi43, fi43, v37)
                        switch (-(fi47)) {
                        case 9223372036854775807:
                            v27 = (v27 / (((v27 << ((0) & 63))) | 1))
                        default:
                            switch fi43 {
                            case 4294967295:
                                switch "\n\t" {
                                case "a":
                                    if false {
                                        if (fi47 < (fi47 << ((65535) & 63))) {
                                            var v49 int16 = -214
                                            var v50 int32 = -103
                                            _ = v49
                                            _ = v50
                                        }
                                        v51 := GenS5{F0: 42, F1: 0}
                                        fmt.Println(v2, fi43)
                                        _ = v51
                                    } else {
                                        v1 = append(v1, (fi43 + 733))
                                        v52 := GenS6{F0: -1}
                                        _ = v52
                                    }
                                case "":
                                    _ = len(v3)
                                case "\n\t":
                                    v53 := GenS3{F0: v37}
                                    _ = v53
                                }
                            case -182:
                                var v54 uint8 = 2
                                _ = v54
                            }
                        }
                    }
                    if (3.14159 != -2.25) {
                        if false {
                            v3[42] = (!false)
                            v27 -= (^(v27))
                        } else {
                            v37 -= (v37 * 99)
                            fmt.Println(v26, fi43, v37)
                        }
                    }
                    switch "hello" {
                    case "a":
                        v27 += 623
                    case "foo bar":
                        switch "a" {
                        case "":
                            var v55 uint32 = (v27 % (((v27 ^ (v27 - v27))) | 1))
                            _ = v55
                        case "foo bar":
                            if (65535 >= 255) {
                                switch ("a" + "foo bar") {
                                case "":
                                    if (true || (false && false)) {
                                        v37 -= -1
                                        var v56 int32 = -410
                                        if false {
                                            if false {
                                                v37 = (-(v37))
                                                fmt.Println(v56)
                                                v37 -= 127
                                                var v57 int32 = (v56 << ((0) & 63))
                                                _ = v57
                                            }
                                            if ("foo bar" != "") {
                                                var v58 bool = ((true && (false && (!true))) && false)
                                                if false {
                                                    v58 = (!(v56 < v56))
                                                } else {
                                                    switch int(v37) {
                                                    case 32768:
                                                        v59 := GenS7{F0: 18446744073709551615, F1: v27}
                                                        _ = v59
                                                    case 421:
                                                        var v60 int = -9223372036854775808
                                                        _ = v60
                                                    case -1:
                                                        fmt.Println(v26)
                                                    default:
                                                        v61 := GenS8{F0: v26, F1: v58, F2: 0}
                                                        _ = v61
                                                    }
                                                    v37 -= (v26 >> ((4294967295) & 63))
                                                }
                                                _ = v58
                                            } else {
                                                switch "foo bar" {
                                                case "a":
                                                    var v62 uint8 = 114
                                                    _ = v62
                                                case "foo bar":
                                                    if true {
                                                        v37 = (v26 * v37)
                                                        v56 -= v56
                                                        var v63 uint32 = (v27 << ((0) & 63))
                                                        _ = v63
                                                    } else {
                                                        if ((42 <= 42) && ((float32((0.0 - (1.5 / 1.5))) * 1e10) > ((3.14159 / 0.1) - 3.14159))) {
                                                            var v64 int8 = 127
                                                            switch fi43 {
                                                            case 743:
                                                                v56 -= v2
                                                            case 1:
                                                                var v65 int32 = (-(v56))
                                                                _ = v65
                                                            default:
                                                                if (1 < genFn0(9223372036854775807, 645, v64)) {
                                                                    v66 := GenS5{F0: (v2 ^ -1), F1: 1}
                                                                    v64 = 0
                                                                    _ = v66
                                                                } else {
                                                                    v2 = (v56 * v56)
                                                                    if true {
                                                                        var v67 float64 = 0.1
                                                                        _ = v67
                                                                    } else {
                                                                        v68 := GenS6{F0: -1}
                                                                        v56 += v2
                                                                        fmt.Println(fi43)
                                                                        _ = v68
                                                                    }
                                                                    var v69 int64 = 580
                                                                    _ = v69
                                                                }
                                                            }
                                                            fmt.Println(v2)
                                                            _ = v64
                                                        } else {
                                                            var v70 int8 = (v26 | (v26 % ((v26) | 1)))
                                                            _ = v70
                                                        }
                                                        switch 32767 {
                                                        case 1:
                                                            fmt.Println(v37)
                                                            switch "\n\t" {
                                                            case "":
                                                                if false {
                                                                    v27 = v27
                                                                    var v71 string = ""
                                                                    v56 = 764
                                                                    _ = v71
                                                                } else {
                                                                    var v72 bool = (-9223372036854775808 < -164)
                                                                    v73 := GenS9{F0: 1, F1: (v37 | (v26 & v37)), F2: 9223372036854775807}
                                                                    _ = v72
                                                                    _ = v73
                                                                }
                                                            case "a":
                                                                var v74 int64 = -9223372036854775808
                                                                _ = v74
                                                            case "\n\t":
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                        case 9223372036854775807:
                                                            v75 := GenS2{F0: "", F1: 784}
                                                            _ = v75
                                                        }
                                                        fmt.Println(v26)
                                                        fmt.Println(fi43, fi47, v2)
                                                    }
                                                case "\n\t":
                                                    fmt.Println(fi43)
                                                default:
                                                    _ = len(v3)
                                                }
                                                if true {
                                                    _ = len(v1)
                                                    var v76 string = "hello"
                                                    _ = v76
                                                }
                                                switch fi43 {
                                                case 2147483647:
                                                    var v77 uint8 = 127
                                                    _ = v77
                                                case -9223372036854775808:
                                                    v78 := GenS6{F0: 0}
                                                    _ = v78
                                                }
                                            }
                                            var v79 float32 = (1.5 * 0.1)
                                            _ = v79
                                        } else {
                                            v80 := v3[127]
                                            _ = v80
                                        }
                                        _ = v56
                                    }
                                case "a":
                                    v37 = -1
                                }
                                v27 = 42
                                var v81 uint = 2
                                var v82 uint16 = 827
                                if false {
                                    v83 := GenS3{F0: -128}
                                    _ = v83
                                } else {
                                    if ("a" == ((("\n\t" + "\n\t") + "a") + ((("\n\t" + "\n\t") + "foo bar") + "foo bar"))) {
                                        v82 = (v82 | v82)
                                        var v84 bool = true
                                        _ = v84
                                    } else {
                                        v37 = v26
                                    }
                                }
                                v81 -= (v81 << ((627) & 63))
                                _ = v81
                                _ = v82
                            }
                        case "hello":
                            if false {
                                if (3.14159 >= 0.1) {
                                    delete(v3, (fi43 << ((18446744073709551615) & 63)))
                                    var v85 uint32 = (v27 >> ((255) & 63))
                                    _ = v85
                                }
                            } else {
                                var v86 int16 = 1000
                                _ = v86
                            }
                            v87 := GenS8{F0: -1, F1: ((!false) && (2147483647 != (v2 & (v2 % (((v2 | (v2 / (((^(v2))) | 1)))) | 1))))), F2: fi43}
                            v3[(^(fi47))] = (false && false)
                            v2 -= -864
                            _ = v87
                        default:
                            var v88 int32 = -2147483648
                            _ = v88
                        }
                    default:
                        switch ("\n\t" + "\n\t") {
                        case "":
                            if len(v1) > 0 {
                                v1[0] = 1
                            }
                        case "\n\t":
                            switch (fi43 ^ fi43) {
                            case 9223372036854775807:
                                fmt.Println(fi47)
                            default:
                                fmt.Println(fi43, v2, v2)
                                var v89 uint64 = 42
                                v90 := v3[fi47]
                                switch (("a" + "hello") + "\n\t") {
                                case "":
                                    switch (fi43 << ((255) & 63)) {
                                    case -1:
                                        if v90 {
                                            v37 += v26
                                            var v91 uint32 = (^(v27))
                                            _ = v91
                                        }
                                    default:
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v92 bool = v90
                                            _ = v92
                                        case "":
                                            if true {
                                                _ = len(v1)
                                                var v93 int16 = 282
                                                var v94 uint8 = 128
                                                if v90 {
                                                    var v95 uint32 = 1
                                                    _ = v95
                                                }
                                                _ = v93
                                                _ = v94
                                            } else {
                                                v1 = append(v1, fi43)
                                                var v96 int = fi43
                                                var v97 int16 = 0
                                                var v98 uint = 18446744073709551615
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v96
                                                _ = v97
                                                _ = v98
                                            }
                                        }
                                    }
                                case "a":
                                    v90 = v90
                                default:
                                    v99 := GenS6{F0: 0}
                                    _ = v99
                                }
                                v90 = ("\n\t" == "\n\t")
                                _ = v89
                                _ = v90
                            }
                        default:
                            var v100 float32 = 1.5
                            _ = v100
                        }
                    }
                    _ = fi47
                }
                var v101 int16 = 32767
                _ = v101
            }
            var v102 int8 = (v26 / ((-20) | 1))
            _ = fi43
            _ = v102
        }
        _ = v37
    } else {
        fmt.Println(v2, v26)
    }
    var v103 int = 1
    for fi104 := 0; fi104 < 5; fi104++ {
        fmt.Println(v2)
        switch "\n\t" {
        case "foo bar":
            if ((819 > 2) && true) {
                v26 -= -1
            } else {
                v105 := GenS5{F0: (v2 ^ -1), F1: (v27 / ((4294967295) | 1))}
                if (-185 >= 0) {
                    var v106 uint16 = 299
                    for fi107 := 0; fi107 < 8; fi107++ {
                        v108 := GenS10{F0: 1}
                        fmt.Println(v27)
                        fmt.Println(fi107)
                        _ = fi107
                        _ = v108
                    }
                    _ = v106
                }
                _ = v105
            }
        }
        _ = fi104
    }
    for fi109 := 0; fi109 < 4; fi109++ {
        v110 := GenS10{F0: (v27 << ((65) & 63))}
        _ = fi109
        _ = v110
    }
    v111 := GenS3{F0: v26}
    for fi112 := 0; fi112 < 4; fi112++ {
        for fi113 := 0; fi113 < 1; fi113++ {
            fmt.Println(v2, fi113)
            _ = fi113
        }
        v114 := GenS7{F0: 50, F1: v27}
        for fi115 := 0; fi115 < 11; fi115++ {
            switch v103 {
            case -214:
                switch -1 {
                case 42:
                    var v116 bool = (797 == -9223372036854775808)
                    _ = v116
                case -445:
                    var v117 int = (^(fi112))
                    _ = v117
                case 942:
                    _ = len(v3)
                default:
                    if (2 == 32767) {
                        switch v103 {
                        case 2147483647:
                            var v118 uint8 = 77
                            _ = v118
                        default:
                            if (1 <= 375) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                var v119 uint = 884
                                _ = v119
                            }
                        }
                        var v120 uint = 611
                        v103 = (fi112 << ((v120) & 63))
                        _ = v120
                    }
                }
            case -379:
                v121 := v3[(-(fi112))]
                if false {
                    v122 := GenS7{F0: 878, F1: v27}
                    _ = v122
                } else {
                    v123 := GenS9{F0: 147, F1: (v26 ^ (v26 >> ((2) & 63))), F2: int64(v26)}
                    _ = v123
                }
                delete(v3, (v103 + v103))
                v124 := GenS11{F0: (v26 | (^(v26))), F1: 9223372036854775807, F2: v27}
                var v125 int32 = v2
                _ = v121
                _ = v124
                _ = v125
            case -795:
                var v126 float64 = float64((1e10 + ((0.1 / 1.0) + 1e10)))
                _ = v126
            default:
                v26 += v26
            }
            _ = fi115
        }
        _ = fi112
        _ = v114
    }
    for fi127 := 0; fi127 < 14; fi127++ {
        switch "a" {
        case "a":
            v128 := GenS12{F0: 2147483647, F1: 120}
            _ = v128
        case "":
            if ((^(v103)) >= v103) {
                for fi129 := 0; fi129 < 8; fi129++ {
                    v130 := v3[fi129]
                    if ((v26 << ((42) & 63)) != 127) {
                        v131 := GenS13{F0: -297, F1: int64(v27), F2: v130}
                        v132 := GenS1{F0: 64}
                        fmt.Println(v27)
                        _ = v131
                        _ = v132
                    }
                    if (0 > 2) {
                        v27 -= (^(v27))
                        if v130 {
                            var v133 bool = false
                            v134 := GenS13{F0: int16(v2), F1: genFn0(0, genFn0(9223372036854775807, -9223372036854775808, (v26 & v26)), v26), F2: (0 >= 4294967295)}
                            _ = v133
                            _ = v134
                        }
                        delete(v3, v103)
                    } else {
                        if false {
                            v2 = 42
                            switch (v103 >> ((255) & 63)) {
                            case 42:
                                var v135 int16 = 0
                                _ = v135
                            default:
                                var v136 uint8 = 42
                                _ = v136
                            }
                            v137 := GenS4{F0: -267, F1: v27}
                            v27 = v27
                            _ = v137
                        } else {
                            v103 = fi129
                        }
                        if v130 {
                            v2 = 575
                            switch "\n\t" {
                            case "foo bar":
                                v27 = 769
                                var v138 int32 = v2
                                var v139 int = fi129
                                _ = v138
                                _ = v139
                            default:
                                v140 := GenS14{F0: ("" + ""), F1: (fi129 >> ((2) & 63)), F2: (v27 + (v27 * (^(v27))))}
                                _ = v140
                            }
                        } else {
                            v141 := GenS11{F0: (v26 * v26), F1: -1, F2: (v27 ^ (v27 * v27))}
                            _ = v141
                        }
                    }
                    if (1e10 <= 0.1) {
                        if (-1 == int8(v27)) {
                            v2 += (v2 & v2)
                        }
                    }
                    _ = len(v1)
                    _ = fi129
                    _ = v130
                }
            } else {
                var v142 int64 = 64
                v26 += 1
                for fi143 := 0; fi143 < 14; fi143++ {
                    var v144 int16 = 42
                    v2 = (v2 / (((v2 / ((v2) | 1))) | 1))
                    switch 0 {
                    case -856:
                        if true {
                            switch "a" {
                            case "foo bar":
                                var v145 int = fi127
                                _ = v145
                            default:
                                var v146 uint16 = 255
                                _ = v146
                            }
                        } else {
                            v144 = v144
                            var v147 int32 = int32(fi127)
                            _ = v147
                        }
                    case -1:
                        delete(v3, -1)
                    case -9223372036854775808:
                        v1 = append(v1, fi127)
                    default:
                        switch "foo bar" {
                        case "foo bar":
                            v148 := GenS1{F0: 2}
                            _ = v148
                        case "a":
                            var v149 uint64 = 50
                            _ = v149
                        default:
                            v150 := GenS12{F0: (fi143 % (((v103 % (((fi127 >> ((18446744073709551615) & 63))) | 1))) | 1)), F1: 127}
                            _ = v150
                        }
                    }
                    _ = fi143
                    _ = v144
                }
                _ = v142
            }
        default:
            v2 = (^(v2))
        }
        _ = fi127
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v103
    _ = v111
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
