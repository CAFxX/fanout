package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint; F1 int64; F2 bool }
type GenS8 struct { F0 bool; F1 uint; F2 int }
type GenS9 struct { F0 uint16; F1 uint }
type GenS10 struct { F0 int; F1 uint64 }
type GenS11 struct { F0 string; F1 bool }
type GenS12 struct { F0 int16; F1 bool }
type GenS13 struct { F0 int; F1 uint }
type GenS14 struct { F0 bool; F1 uint32; F2 int32 }
type GenS15 struct { F0 int64; F1 int }
type GenS16 struct { F0 uint8; F1 int }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 uint32 }
type GenS19 struct { F0 uint8 }
type GenS20 struct { F0 int; F1 string }
type GenS21 struct { F0 int8; F1 uint }
type GenS22 struct { F0 uint8 }
type GenS23 struct { F0 uint8; F1 string; F2 string }
type GenS24 struct { F0 int8; F1 int }
type GenS25 struct { F0 int64; F1 string }
type GenS26 struct { F0 int; F1 bool }
type GenS27 struct { F0 uint16 }
type GenS28 struct { F0 int; F1 int16 }
type GenS29 struct { F0 uint32; F1 uint32; F2 bool }
type GenS30 struct { F0 uint16; F1 uint32 }
type GenS31 struct { F0 uint16 }
type GenS32 struct { F0 int8; F1 uint8; F2 int16 }
type GenS33 struct { F0 bool }
type GenS34 struct { F0 bool; F1 bool; F2 uint32 }
type GenS35 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 int8 = 42
    var v2 int = -9223372036854775808
    var v3 bool = ("hello" == ("\n\t" + "hello"))
    v4 := make(map[int]bool)
    var v5 uint8 = 127
    if (481 < 861) {
        v2 = 127
    } else {
        fmt.Println(v5)
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        if false {
            for fi7 := 0; fi7 < 13; fi7++ {
                fmt.Println(fi7, v1)
                _ = fi7
            }
            if (!v3) {
                v8 := []int{fi6, -9223372036854775808, (fi6 & (fi6 ^ fi6))}
                for fi9 := 0; fi9 < 11; fi9++ {
                    var v10 string = "hello"
                    switch v2 {
                    case -944:
                        var v11 uint64 = 878
                        _ = v11
                    default:
                        v10 = ("\n\t" + v10)
                    }
                    _ = fi9
                    _ = v10
                }
                _ = v8
            } else {
                delete(v4, -9223372036854775808)
                for fi12 := 0; fi12 < 4; fi12++ {
                    var v13 uint32 = 745
                    _ = fi12
                    _ = v13
                }
                for fi14 := 0; fi14 < 7; fi14++ {
                    v15 := GenS1{F0: 193}
                    switch "a" {
                    case "foo bar":
                        v16 := []int{}
                        _ = v16
                    case "a":
                        fmt.Println(fi6)
                    }
                    var v17 bool = (((!(v1 != 100)) || (!((80 == 498) && (!(v3 && true))))) && v3)
                    _ = fi14
                    _ = v15
                    _ = v17
                }
            }
            fmt.Println(v2)
        } else {
            v18 := []int{v2}
            if ("hello" <= "hello") {
                for fi19 := 0; fi19 < 3; fi19++ {
                    v5 = v5
                    fmt.Println(v5, v2)
                    _ = fi19
                }
                if v3 {
                    v20 := GenS1{F0: v5}
                    var v21 float64 = 3.14159
                    if v3 {
                        var v22 uint16 = 32767
                        _ = v22
                    }
                    _ = v20
                    _ = v21
                } else {
                    v23 := GenS1{F0: (v5 - 127)}
                    _ = v23
                }
                v5 -= v5
            }
            var v24 int = -1
            _ = v18
            _ = v24
        }
        fmt.Println(v3, fi6)
        _ = fi6
    }
    var v25 uint64 = 4294967295
    var v26 int = v2
    fmt.Println(v25)
    _ = len(v4)
    fmt.Println(v2, v2)
    fmt.Println(v1)
    switch "\n\t" {
    case "foo bar":
        v27 := []int{}
        _ = v27
    case "\n\t":
        for fi28 := 0; fi28 < 6; fi28++ {
            v29 := []int{fi28}
            var v30 int32 = 42
            v31 := GenS1{F0: 42}
            _ = fi28
            _ = v29
            _ = v30
            _ = v31
        }
    default:
        if false {
            for fi32 := 0; fi32 < 8; fi32++ {
                if (0.0 < -2.25) {
                    if false {
                        var v33 uint32 = 2147483647
                        _ = v33
                    }
                } else {
                    delete(v4, -933)
                }
                _ = fi32
            }
            var v34 int16 = 42
            _ = v34
        }
    }
    v5 -= 91
    for fi35 := 0; fi35 < 9; fi35++ {
        if v3 {
            if v3 {
                switch (v2 << ((0) & 63)) {
                case 609:
                    if (((v5 << ((4294967295) & 63)) >= v5) || (true && true)) {
                        var v36 int16 = 604
                        v26 = (-(fi35))
                        var v37 uint16 = 0
                        _ = v36
                        _ = v37
                    } else {
                        for fi38 := 0; fi38 < 9; fi38++ {
                            var v39 uint = 255
                            v40 := v4[(^(fi35))]
                            if true {
                                delete(v4, fi38)
                            } else {
                                var v41 int32 = -301
                                var v42 bool = v3
                                if (!false) {
                                    v4[fi35] = (!(v39 <= (^(v39))))
                                    v3 = ((v26 >> ((v39) & 63)) != v2)
                                    switch ("a" + "a") {
                                    case "a":
                                        v43 := []int{-378, fi38}
                                        _ = v43
                                    case "hello":
                                        var v44 int16 = 1000
                                        _ = v44
                                    case "":
                                        var v45 uint32 = 2
                                        _ = v45
                                    }
                                }
                                _ = v41
                                _ = v42
                            }
                            if (!false) {
                                v46 := []int{(-(fi35)), 127, (v2 + fi38), (fi35 + (fi35 ^ (v26 % ((794) | 1))))}
                                if len(v46) >= 2 {
                                    v46 = v46[1:len(v46)]
                                }
                                _ = v46
                            } else {
                                v47 := GenS2{F0: 0}
                                v25 = v25
                                _ = v47
                            }
                            var v48 uint16 = 2
                            var v49 uint8 = (v5 << (((^(v39))) & 63))
                            if (0 < 0) {
                                v48 = v48
                                fmt.Println(fi35, v26, v48)
                                var v50 float32 = 0.0
                                _ = v50
                            } else {
                                if false {
                                    var v51 int16 = 1
                                    var v52 int8 = v1
                                    if (!v3) {
                                        switch (^(v2)) {
                                        case 430:
                                            v53 := []int{128, (v26 << ((v39) & 63))}
                                            v40 = (v3 && v3)
                                            _ = v53
                                        }
                                        if (!((0.1 < 0.0) || v3)) {
                                            if ("a" != "foo bar") {
                                                switch "hello" {
                                                case "a":
                                                    if false {
                                                        switch (fi35 * (fi38 / (((fi38 >> (((v39 & (^(v39)))) & 63))) | 1))) {
                                                        case 415:
                                                            var v54 int8 = (v52 << (((v39 / ((uint(v25)) | 1))) & 63))
                                                            _ = v54
                                                        case 306:
                                                            var v55 uint16 = (v48 - (^(v48)))
                                                            _ = v55
                                                        case 9223372036854775807:
                                                            if ((-1 > 42) || (258 <= 9223372036854775807)) {
                                                                switch ("foo bar" + "hello") {
                                                                case "":
                                                                    fmt.Println(v26, v52)
                                                                case "hello":
                                                                    var v56 int = 196
                                                                    _ = v56
                                                                default:
                                                                    var v57 float64 = (1e10 / 1.0)
                                                                    _ = v57
                                                                }
                                                                v25 += uint64(v48)
                                                            }
                                                        }
                                                        v40 = (-2147483648 >= 0)
                                                        var v58 string = "\n\t"
                                                        switch (((v58 + v58) + v58) + v58) {
                                                        case "foo bar":
                                                            v40 = (0 == 42)
                                                        case "":
                                                            v48 -= v48
                                                        default:
                                                            if v3 {
                                                                v4[v2] = (-9223372036854775808 <= -1)
                                                                v59 := GenS2{F0: (v48 >> ((v39) & 63))}
                                                                switch v58 {
                                                                case "\n\t":
                                                                    var v60 uint16 = (v48 ^ (v48 | v48))
                                                                    _ = v60
                                                                case "":
                                                                    if ((v49 * (v49 % (((v49 ^ (v49 / (((v49 >> ((uint(v48)) & 63))) | 1)))) | 1))) > 101) {
                                                                        var v61 uint64 = (^(v25))
                                                                        var v62 uint64 = (v61 % ((v25) | 1))
                                                                        v49 = v5
                                                                        if (((v49 >= (v49 - 1)) || v40) || v40) {
                                                                            v48 += (v48 << (((v39 & v39)) & 63))
                                                                            if v3 {
                                                                                v63 := []int{4294967295, -9223372036854775808, 2147483647}
                                                                                if len(v63) >= 2 {
                                                                                    v63 = v63[1:len(v63)]
                                                                                }
                                                                                if (-2147483648 == 42) {
                                                                                    var v64 uint16 = v48
                                                                                    fmt.Println(v5, fi35, v40)
                                                                                    _ = v64
                                                                                }
                                                                                _ = v63
                                                                            }
                                                                        }
                                                                        var v65 int32 = 42
                                                                        _ = v61
                                                                        _ = v62
                                                                        _ = v65
                                                                    } else {
                                                                        var v66 uint8 = 2
                                                                        var v67 float32 = (-2.25 + float32((0.0 - 1.5)))
                                                                        var v68 uint8 = v5
                                                                        _ = len(v4)
                                                                        var v69 uint = v39
                                                                        _ = v66
                                                                        _ = v67
                                                                        _ = v68
                                                                        _ = v69
                                                                    }
                                                                case "foo bar":
                                                                    _ = len(v4)
                                                                default:
                                                                    var v70 uint = v39
                                                                    _ = v70
                                                                }
                                                                fmt.Println(v52, v25, v58)
                                                                v71 := GenS2{F0: uint16(v25)}
                                                                v72 := GenS1{F0: (^(v5))}
                                                                var v73 float64 = 0.0
                                                                var v74 string = (("a" + "hello") + v58)
                                                                fmt.Println(v48, v58)
                                                                switch (fi35 & int(v48)) {
                                                                case -416:
                                                                    switch fi35 {
                                                                    case -527:
                                                                        if v40 {
                                                                            var v75 int8 = (^(v52))
                                                                            _ = len(v4)
                                                                            v49 -= v5
                                                                            switch (fi38 * (v2 | (v26 - (fi35 << ((v39) & 63))))) {
                                                                            case 0:
                                                                                switch (v26 / ((v26) | 1)) {
                                                                                case 960:
                                                                                    if (v39 >= (v39 >> (((v39 << ((v39) & 63))) & 63))) {
                                                                                        fmt.Println(v5)
                                                                                        if v3 {
                                                                                            v76 := GenS3{F0: 255}
                                                                                            var v77 int8 = (v1 + (v1 >> (((v39 << (((v39 | 255)) & 63))) & 63)))
                                                                                            v77 = 82
                                                                                            _ = v76
                                                                                            _ = v77
                                                                                        }
                                                                                        _ = len(v4)
                                                                                    }
                                                                                case 32768:
                                                                                    v25 += v25
                                                                                }
                                                                            case -9223372036854775808:
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    if v3 {
                                                                                        if true {
                                                                                            v2 = (^(v26))
                                                                                        } else {
                                                                                            v73 -= (-2.25 * ((v73 / v73) * (0.0 - (v73 - (((v73 * (v73 * v73)) - (1.5 + (float64((float32((float64(1.5) + ((1e10 + 3.14159) - v73))) - ((((float32(-2.25) * (1.5 - float32(v73))) + -2.25) / 1.0) + float32(v73)))) + float64(1e10)))) / v73)))))
                                                                                        }
                                                                                        var v78 uint16 = v48
                                                                                        if (false && v3) {
                                                                                            var v79 uint = 484
                                                                                            _ = v79
                                                                                        } else {
                                                                                            switch -691 {
                                                                                            case -1:
                                                                                                v80 := v4[fi38]
                                                                                                _ = v80
                                                                                            case 42:
                                                                                                v81 := GenS4{F0: int16(v26)}
                                                                                                _ = v81
                                                                                            case 690:
                                                                                                if (v40 && v40) {
                                                                                                    var v82 uint64 = v25
                                                                                                    if (((v74 + "a") == v74) && false) {
                                                                                                        var v83 uint = v39
                                                                                                        var v84 uint32 = uint32(v25)
                                                                                                        _ = v83
                                                                                                        _ = v84
                                                                                                    } else {
                                                                                                        v85 := []int{v26}
                                                                                                        if v40 {
                                                                                                            var v86 int16 = 330
                                                                                                            switch (v58 + "hello") {
                                                                                                            case "hello":
                                                                                                                fmt.Println(v58, fi35)
                                                                                                            case "":
                                                                                                                var v87 int = (fi35 >> ((uint(v25)) & 63))
                                                                                                                _ = v87
                                                                                                            default:
                                                                                                                v86 += (v86 | (v86 + (v86 % ((32767) | 1))))
                                                                                                            }
                                                                                                            fmt.Println(v58, v58)
                                                                                                            _ = v86
                                                                                                        }
                                                                                                        _ = v85
                                                                                                    }
                                                                                                    _ = v82
                                                                                                }
                                                                                            default:
                                                                                                v88 := GenS3{F0: v39}
                                                                                                _ = v88
                                                                                            }
                                                                                            if v3 {
                                                                                                if ((v78 + (v78 + v78)) <= 32767) {
                                                                                                    v89 := GenS5{F0: 42}
                                                                                                    if v40 {
                                                                                                        var v90 float64 = 0.1
                                                                                                        var v91 uint32 = 42
                                                                                                        v92 := GenS3{F0: v39}
                                                                                                        var v93 uint8 = uint8(v78)
                                                                                                        fmt.Println(v51)
                                                                                                        if true {
                                                                                                            switch -487 {
                                                                                                            case 575:
                                                                                                                v2 = (v2 + (v26 - v26))
                                                                                                            }
                                                                                                            if (!((^(v25)) != 1)) {
                                                                                                                _ = len(v4)
                                                                                                                v40 = v3
                                                                                                            } else {
                                                                                                                v94 := GenS6{F0: (v51 % (((v51 | v51)) | 1))}
                                                                                                                fmt.Println(v78, v75)
                                                                                                                if (0.1 >= -2.25) {
                                                                                                                    var v95 int32 = -118
                                                                                                                    v96 := GenS7{F0: v39, F1: 9223372036854775807, F2: v40}
                                                                                                                    fmt.Println(v73)
                                                                                                                    _ = v95
                                                                                                                    _ = v96
                                                                                                                }
                                                                                                                _ = v94
                                                                                                            }
                                                                                                            v5 = v93
                                                                                                            fmt.Println(v48, v91, v51)
                                                                                                        }
                                                                                                        var v97 uint32 = v91
                                                                                                        switch (v2 / ((fi38) | 1)) {
                                                                                                        case 42:
                                                                                                            v98 := []int{4294967295, v26, 669}
                                                                                                            _ = v98
                                                                                                        }
                                                                                                        if (v48 >= (v78 + 32767)) {
                                                                                                            fmt.Println(v40)
                                                                                                            switch v26 {
                                                                                                            case 0:
                                                                                                                v99 := v4[4294967295]
                                                                                                                _ = v99
                                                                                                            case 4294967295:
                                                                                                                var v100 uint16 = v78
                                                                                                                _ = v100
                                                                                                            }
                                                                                                            fmt.Println(v25, v51)
                                                                                                        } else {
                                                                                                            v52 = v75
                                                                                                        }
                                                                                                        _ = v90
                                                                                                        _ = v91
                                                                                                        _ = v92
                                                                                                        _ = v93
                                                                                                        _ = v97
                                                                                                    }
                                                                                                    _ = v89
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v40 = true
                                                                                        _ = v78
                                                                                    } else {
                                                                                        v25 -= v25
                                                                                        delete(v4, v2)
                                                                                    }
                                                                                case "a":
                                                                                    v101 := v4[9223372036854775807]
                                                                                    _ = v101
                                                                                default:
                                                                                    v52 -= (v1 % ((-109) | 1))
                                                                                }
                                                                                if v40 {
                                                                                    var v102 int = (^(v2))
                                                                                    if v40 {
                                                                                        v103 := GenS4{F0: v51}
                                                                                        _ = v103
                                                                                    }
                                                                                    var v104 bool = (v26 <= 4294967295)
                                                                                    _ = v102
                                                                                    _ = v104
                                                                                }
                                                                                if v3 {
                                                                                    fmt.Println(v39, v40)
                                                                                    v25 = v25
                                                                                    var v105 int = -606
                                                                                    _ = v105
                                                                                } else {
                                                                                    var v106 uint8 = (v49 & (v5 * v5))
                                                                                    v107 := GenS7{F0: (v39 - 42), F1: 763, F2: (894 >= -1)}
                                                                                    v39 = 4294967295
                                                                                    _ = v106
                                                                                    _ = v107
                                                                                }
                                                                            case -399:
                                                                                v108 := GenS8{F0: ((v5 % (((v49 % ((2) | 1))) | 1)) == 2), F1: v39, F2: 0}
                                                                                _ = v108
                                                                            }
                                                                            _ = v75
                                                                        }
                                                                    default:
                                                                        if v3 {
                                                                            if v3 {
                                                                                if v40 {
                                                                                    fmt.Println(v40, v2, v48)
                                                                                    fmt.Println(v48, v5)
                                                                                    v109 := GenS9{F0: 65535, F1: (v39 << (((v39 % ((v39) | 1))) & 63))}
                                                                                    if (!((v52 & 1) > v52)) {
                                                                                        switch v58 {
                                                                                        case "\n\t":
                                                                                            v58 = ""
                                                                                        case "hello":
                                                                                            var v110 bool = (-63 > int64(v25))
                                                                                            _ = v110
                                                                                        }
                                                                                        if (v3 && v3) {
                                                                                            v111 := GenS10{F0: (v2 & fi35), F1: v25}
                                                                                            _ = v111
                                                                                        } else {
                                                                                            switch fi38 {
                                                                                            case -616:
                                                                                                _ = len(v4)
                                                                                            case 2147483647:
                                                                                                fmt.Println(fi35, v48)
                                                                                            case -842:
                                                                                                switch "foo bar" {
                                                                                                case "foo bar":
                                                                                                    var v112 uint = 18446744073709551615
                                                                                                    _ = v112
                                                                                                case "\n\t":
                                                                                                    if v3 {
                                                                                                        v113 := GenS11{F0: (v58 + v74), F1: (!v40)}
                                                                                                        _ = v113
                                                                                                    }
                                                                                                case "":
                                                                                                    var v114 int8 = v52
                                                                                                    _ = v114
                                                                                                default:
                                                                                                    if (1 < -543) {
                                                                                                        v115 := []int{(v2 + (fi38 ^ (^(fi38)))), (v2 << ((v39) & 63)), 272}
                                                                                                        if v40 {
                                                                                                            v2 = fi35
                                                                                                            var v116 int = (fi38 / ((-999) | 1))
                                                                                                            if v40 {
                                                                                                                var v117 uint = 983
                                                                                                                _ = v117
                                                                                                            }
                                                                                                            _ = v116
                                                                                                        } else {
                                                                                                            v1 = (v1 + (v1 & 0))
                                                                                                            v51 = 242
                                                                                                            var v118 int16 = 1
                                                                                                            _ = v118
                                                                                                        }
                                                                                                        if ((0.1 + 3.14159) != v73) {
                                                                                                            v119 := GenS12{F0: v51, F1: v3}
                                                                                                            if false {
                                                                                                                v39 += (v39 % (((v39 & v39)) | 1))
                                                                                                                _ = len(v115)
                                                                                                            }
                                                                                                            fmt.Println(fi38, v40)
                                                                                                            _ = v119
                                                                                                        } else {
                                                                                                            _ = len(v4)
                                                                                                            var v120 uint16 = 2
                                                                                                            v115 = append(v115, (fi35 + (v2 >> ((v39) & 63))))
                                                                                                            v73 = v73
                                                                                                            _ = v120
                                                                                                        }
                                                                                                        _ = v115
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                switch v2 {
                                                                                                case 1:
                                                                                                    switch (v74 + ("a" + ("" + (v74 + (v58 + (v58 + "a")))))) {
                                                                                                    case "\n\t":
                                                                                                        v121 := GenS13{F0: (fi38 * v26), F1: v39}
                                                                                                        v122 := v4[fi38]
                                                                                                        fmt.Println(v58, v3, v122)
                                                                                                        _ = v121
                                                                                                        _ = v122
                                                                                                    default:
                                                                                                        v39 = v39
                                                                                                    }
                                                                                                case 42:
                                                                                                    if (v51 == (v51 | v51)) {
                                                                                                        _ = len(v4)
                                                                                                        delete(v4, fi35)
                                                                                                        var v123 int64 = -9223372036854775808
                                                                                                        _ = v123
                                                                                                    }
                                                                                                default:
                                                                                                    var v124 int8 = v52
                                                                                                    _ = v124
                                                                                                }
                                                                                            }
                                                                                            var v125 bool = v3
                                                                                            _ = v125
                                                                                        }
                                                                                        fmt.Println(v1, v52, v3)
                                                                                    }
                                                                                    v126 := GenS13{F0: (fi38 ^ fi35), F1: v39}
                                                                                    v127 := GenS7{F0: 42, F1: -1, F2: (!(uint64(v49) == v25))}
                                                                                    _ = v109
                                                                                    _ = v126
                                                                                    _ = v127
                                                                                }
                                                                                if v3 {
                                                                                    if v40 {
                                                                                        switch v26 {
                                                                                        case -819:
                                                                                            v2 = v26
                                                                                        case 32768:
                                                                                            if v40 {
                                                                                                v128 := GenS14{F0: false, F1: uint32(v25), F2: 0}
                                                                                                if v3 {
                                                                                                    var v129 float64 = (v73 + ((float64((float32(-2.25) / 1e10)) * 0.0) + v73))
                                                                                                    v1 += -101
                                                                                                    switch (fi35 * (-(v2))) {
                                                                                                    case 1:
                                                                                                        var v130 uint = 1
                                                                                                        _ = v130
                                                                                                    case 4294967295:
                                                                                                        v131 := []int{(v2 % ((fi35) | 1)), (fi35 ^ v26), int(v39), int(v52)}
                                                                                                        _ = v131
                                                                                                    default:
                                                                                                        v132 := GenS15{F0: 680, F1: 128}
                                                                                                        _ = v132
                                                                                                    }
                                                                                                    switch (v26 / (((fi35 - fi35)) | 1)) {
                                                                                                    case 42:
                                                                                                        var v133 float32 = 3.14159
                                                                                                        _ = v133
                                                                                                    case 0:
                                                                                                        fmt.Println(v58, v3)
                                                                                                    case -938:
                                                                                                        fmt.Println(v1, v73, v25)
                                                                                                    default:
                                                                                                        var v134 int16 = -1
                                                                                                        _ = v134
                                                                                                    }
                                                                                                    if true {
                                                                                                        var v135 uint64 = (^(v25))
                                                                                                        var v136 int8 = (v1 - (v1 / ((v1) | 1)))
                                                                                                        _ = v135
                                                                                                        _ = v136
                                                                                                    } else {
                                                                                                        var v137 int8 = (v1 % (((v1 + 127)) | 1))
                                                                                                        var v138 bool = v40
                                                                                                        v139 := []int{fi35, fi38}
                                                                                                        _ = v137
                                                                                                        _ = v138
                                                                                                        _ = v139
                                                                                                    }
                                                                                                    _ = v129
                                                                                                }
                                                                                                v140 := GenS3{F0: (v39 >> ((v39) & 63))}
                                                                                                v40 = true
                                                                                                v1 = 1
                                                                                                _ = v128
                                                                                                _ = v140
                                                                                            } else {
                                                                                                v141 := GenS16{F0: (^(v49)), F1: -78}
                                                                                                if v3 {
                                                                                                    if v3 {
                                                                                                        if v40 {
                                                                                                            v142 := GenS17{F0: v1}
                                                                                                            if v3 {
                                                                                                                var v143 int8 = (v1 - 0)
                                                                                                                _ = v143
                                                                                                            }
                                                                                                            _ = v142
                                                                                                        } else {
                                                                                                            var v144 string = (v58 + "a")
                                                                                                            _ = v144
                                                                                                        }
                                                                                                        v145 := []int{fi38, int(v51), 128}
                                                                                                        var v146 uint = 195
                                                                                                        switch "foo bar" {
                                                                                                        case "":
                                                                                                            v147 := GenS3{F0: v146}
                                                                                                            _ = v147
                                                                                                        case "foo bar":
                                                                                                            v1 = 100
                                                                                                        }
                                                                                                        _ = v145
                                                                                                        _ = v146
                                                                                                    } else {
                                                                                                        switch v2 {
                                                                                                        case 42:
                                                                                                            var v148 uint = (v39 * (^(v39)))
                                                                                                            _ = v148
                                                                                                        }
                                                                                                        v49 -= v49
                                                                                                        var v149 int32 = int32(v51)
                                                                                                        _ = v149
                                                                                                    }
                                                                                                    v150 := []int{(v2 ^ (fi35 << ((18446744073709551615) & 63))), 42}
                                                                                                    _ = v150
                                                                                                } else {
                                                                                                    if ((v48 << (((v39 ^ (v39 * (v39 - v39)))) & 63)) == 42) {
                                                                                                        v52 -= (v52 + (v1 | v52))
                                                                                                    }
                                                                                                    switch -2147483648 {
                                                                                                    case -427:
                                                                                                        switch 128 {
                                                                                                        case 32768:
                                                                                                            if true {
                                                                                                                var v151 uint64 = 42
                                                                                                                var v152 int16 = v51
                                                                                                                if (!((702 == 572) || v40)) {
                                                                                                                    var v153 float32 = 0.1
                                                                                                                    _ = v153
                                                                                                                } else {
                                                                                                                    switch v58 {
                                                                                                                    case "a":
                                                                                                                        v154 := GenS6{F0: (v51 % ((634) | 1))}
                                                                                                                        _ = v154
                                                                                                                    case "":
                                                                                                                        var v155 uint64 = (v25 - v151)
                                                                                                                        var v156 uint = (^(v39))
                                                                                                                        fmt.Println(v52, v73)
                                                                                                                        _ = v155
                                                                                                                        _ = v156
                                                                                                                    }
                                                                                                                    v40 = v3
                                                                                                                }
                                                                                                                v157 := []int{(-(fi35)), (fi38 * 502), (v2 - -736)}
                                                                                                                _ = v151
                                                                                                                _ = v152
                                                                                                                _ = v157
                                                                                                            } else {
                                                                                                                if v3 {
                                                                                                                    fmt.Println(fi38)
                                                                                                                }
                                                                                                            }
                                                                                                        case -1:
                                                                                                            v158 := []int{(fi38 | (fi38 << ((v39) & 63)))}
                                                                                                            v159 := GenS18{F0: 1}
                                                                                                            v2 -= (v2 << (((v39 + (v39 % ((v39) | 1)))) & 63))
                                                                                                            _ = v158
                                                                                                            _ = v159
                                                                                                        }
                                                                                                    case 674:
                                                                                                        switch ("" + (v74 + ((v74 + v58) + (("a" + v74) + ((v58 + ("" + (v74 + ("a" + (v58 + v74))))) + ("\n\t" + v58)))))) {
                                                                                                        case "a":
                                                                                                            var v160 int = 32768
                                                                                                            _ = v160
                                                                                                        default:
                                                                                                            v58 = ((v74 + ((((v74 + (v74 + v74)) + "") + "\n\t") + v74)) + v58)
                                                                                                        }
                                                                                                    case 71:
                                                                                                        if v40 {
                                                                                                            if (int32(fi38) >= 0) {
                                                                                                                var v161 string = (v58 + (((v74 + v58) + (v58 + (v74 + ((((v58 + ("foo bar" + v58)) + "hello") + (v58 + v74)) + v74)))) + v58))
                                                                                                                if true {
                                                                                                                    v162 := GenS19{F0: (v49 & 2)}
                                                                                                                    _ = v162
                                                                                                                } else {
                                                                                                                    _ = len(v4)
                                                                                                                }
                                                                                                                v40 = (1 <= (v1 >> ((v39) & 63)))
                                                                                                                v40 = ((v25 & 18446744073709551615) != (v25 * v25))
                                                                                                                fmt.Println(v3, v74)
                                                                                                                _ = v161
                                                                                                            }
                                                                                                        }
                                                                                                    default:
                                                                                                        v74 = (v74 + v58)
                                                                                                    }
                                                                                                }
                                                                                                var v163 uint64 = 4294967295
                                                                                                _ = v141
                                                                                                _ = v163
                                                                                            }
                                                                                        default:
                                                                                            fmt.Println(v48, v73, v74)
                                                                                        }
                                                                                        var v164 int32 = 984
                                                                                        var v165 bool = v40
                                                                                        _ = v164
                                                                                        _ = v165
                                                                                    } else {
                                                                                        v166 := GenS4{F0: (v51 | 32767)}
                                                                                        if (42 < -619) {
                                                                                            if false {
                                                                                                v49 = v5
                                                                                                v167 := GenS2{F0: (^(v48))}
                                                                                                v168 := GenS4{F0: (v51 >> (((v39 ^ (v39 % ((v39) | 1)))) & 63))}
                                                                                                _ = v167
                                                                                                _ = v168
                                                                                            } else {
                                                                                                v52 -= 1
                                                                                            }
                                                                                            v169 := v4[(-(fi38))]
                                                                                            v49 += 116
                                                                                            _ = v169
                                                                                        } else {
                                                                                            fmt.Println(v25, v58)
                                                                                            var v170 string = (("a" + v58) + v58)
                                                                                            _ = v170
                                                                                        }
                                                                                        v171 := []int{640, fi38, (v26 + v26), int(v39)}
                                                                                        _ = v166
                                                                                        _ = v171
                                                                                    }
                                                                                    fmt.Println(v40, fi35)
                                                                                }
                                                                            } else {
                                                                                if ((v5 == (v49 << (((v39 + (^(v39)))) & 63))) && v3) {
                                                                                    var v172 uint32 = 7
                                                                                    _ = v172
                                                                                }
                                                                                v49 += 184
                                                                                v173 := GenS20{F0: 9223372036854775807, F1: "a"}
                                                                                v58 = v74
                                                                                v174 := GenS21{F0: (v52 | v1), F1: 18446744073709551615}
                                                                                _ = v173
                                                                                _ = v174
                                                                            }
                                                                            v51 = 42
                                                                            if v40 {
                                                                                var v175 uint = v39
                                                                                v176 := GenS11{F0: v58, F1: true}
                                                                                if (!((1.5 == 0.0) && (!v40))) {
                                                                                    var v177 int64 = int64(v26)
                                                                                    var v178 int8 = v52
                                                                                    _ = v177
                                                                                    _ = v178
                                                                                }
                                                                                _ = v175
                                                                                _ = v176
                                                                            }
                                                                            v39 -= uint(v25)
                                                                        } else {
                                                                            v52 = (v1 % ((100) | 1))
                                                                            var v179 uint64 = v25
                                                                            var v180 uint32 = 65535
                                                                            _ = v179
                                                                            _ = v180
                                                                        }
                                                                    }
                                                                    v5 = (v49 % ((v5) | 1))
                                                                    var v181 float64 = 0.0
                                                                    _ = v181
                                                                case 2147483647:
                                                                    var v182 uint16 = 1
                                                                    _ = v182
                                                                default:
                                                                    fmt.Println(v58, v52)
                                                                }
                                                                _ = v59
                                                                _ = v71
                                                                _ = v72
                                                                _ = v73
                                                                _ = v74
                                                            }
                                                        }
                                                        v183 := GenS19{F0: (^(v5))}
                                                        if false {
                                                            v48 += (v48 + (v48 + v48))
                                                        } else {
                                                            var v184 float64 = 0.1
                                                            v185 := []int{}
                                                            _ = v184
                                                            _ = v185
                                                        }
                                                        switch "hello" {
                                                        case "":
                                                            var v186 int32 = 2147483647
                                                            _ = v186
                                                        case "hello":
                                                            var v187 uint64 = (v25 + (v25 >> ((v39) & 63)))
                                                            _ = v187
                                                        default:
                                                            var v188 float64 = 0.1
                                                            _ = v188
                                                        }
                                                        _ = v183
                                                        _ = v58
                                                    }
                                                case "":
                                                    var v189 int64 = -9223372036854775808
                                                    _ = v189
                                                default:
                                                    var v190 int64 = 9223372036854775807
                                                    _ = v190
                                                }
                                                switch ("foo bar" + "\n\t") {
                                                case "hello":
                                                    switch v26 {
                                                    case 628:
                                                        var v191 int = (fi35 << ((v39) & 63))
                                                        _ = v191
                                                    case 32767:
                                                        var v192 uint16 = 429
                                                        _ = v192
                                                    }
                                                default:
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v193 := []int{(v2 ^ int(v51)), (^(fi38)), (v2 | (fi38 >> ((uint(v25)) & 63)))}
                                                        _ = v193
                                                    case "\n\t":
                                                        switch "hello" {
                                                        case "":
                                                            var v194 int16 = 42
                                                            _ = v194
                                                        }
                                                    case "a":
                                                        if (42 == -966) {
                                                            switch v2 {
                                                            case 0:
                                                                var v195 float32 = (0.0 - (((((0.0 * 3.14159) * 3.14159) - 0.0) * 1.5) + (-2.25 / 0.1)))
                                                                _ = v195
                                                            case 127:
                                                                switch 4294967295 {
                                                                case 2147483647:
                                                                    v52 -= -1
                                                                case 4294967295:
                                                                    var v196 float64 = float64((((1.5 + 1.5) + -2.25) * (0.0 - -2.25)))
                                                                    _ = v196
                                                                case 845:
                                                                    v197 := GenS22{F0: (v49 | (^(v5)))}
                                                                    fmt.Println(fi35, v49)
                                                                    delete(v4, int(v39))
                                                                    _ = v197
                                                                default:
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        var v198 int = fi38
                                                                        _ = v198
                                                                    }
                                                                }
                                                            default:
                                                                if false {
                                                                    v2 = 42
                                                                } else {
                                                                    if (((0.0 / (float64(((0.1 / 1.0) - (0.0 * 1e10))) * float64(float32(3.14159)))) != 3.14159) && v40) {
                                                                        var v199 float64 = -2.25
                                                                        _ = len(v4)
                                                                        if (!((v25 >> ((v39) & 63)) <= v25)) {
                                                                            if (v40 || false) {
                                                                                v200 := v4[v26]
                                                                                var v201 uint = v39
                                                                                if v3 {
                                                                                    switch (v26 | (fi35 - v2)) {
                                                                                    case -9223372036854775808:
                                                                                        if (true && v40) {
                                                                                            if v40 {
                                                                                                v200 = (0 != 2)
                                                                                            }
                                                                                            var v202 uint16 = uint16(v25)
                                                                                            _ = v202
                                                                                        } else {
                                                                                            var v203 int8 = 127
                                                                                            v204 := GenS23{F0: v49, F1: ("" + ((((("a" + ("a" + "hello")) + "\n\t") + "\n\t") + "hello") + "a")), F2: "foo bar"}
                                                                                            delete(v4, int(v52))
                                                                                            _ = v203
                                                                                            _ = v204
                                                                                        }
                                                                                    case 32767:
                                                                                        var v205 uint16 = 306
                                                                                        _ = v205
                                                                                    case 128:
                                                                                        switch (v26 / ((fi38) | 1)) {
                                                                                        case 0:
                                                                                            fmt.Println(v26, v52)
                                                                                        case 42:
                                                                                            v48 = v48
                                                                                        case 127:
                                                                                            delete(v4, v2)
                                                                                        default:
                                                                                            fmt.Println(fi38)
                                                                                        }
                                                                                    default:
                                                                                        delete(v4, -2147483648)
                                                                                    }
                                                                                    if (719 <= v201) {
                                                                                        if true {
                                                                                            if true {
                                                                                                v48 -= (v48 * uint16(v5))
                                                                                                if (2 >= 121) {
                                                                                                    var v206 int16 = (-(v51))
                                                                                                    v207 := GenS8{F0: false, F1: 42, F2: (fi35 + 4294967295)}
                                                                                                    v208 := GenS6{F0: v51}
                                                                                                    v209 := GenS11{F0: ("\n\t" + "foo bar"), F1: (!true)}
                                                                                                    _ = v206
                                                                                                    _ = v207
                                                                                                    _ = v208
                                                                                                    _ = v209
                                                                                                } else {
                                                                                                    if v3 {
                                                                                                        v5 -= (^(v5))
                                                                                                    }
                                                                                                    var v210 int8 = (v52 | v52)
                                                                                                    _ = v210
                                                                                                }
                                                                                            } else {
                                                                                                var v211 int8 = 127
                                                                                                _ = v211
                                                                                            }
                                                                                            v4[v2] = ("a" != "\n\t")
                                                                                            v212 := GenS15{F0: -9223372036854775808, F1: 128}
                                                                                            _ = v212
                                                                                        }
                                                                                        v213 := []int{4294967295}
                                                                                        _ = v213
                                                                                    } else {
                                                                                        if (3.14159 > v199) {
                                                                                            v26 += (v2 << (((^(v201))) & 63))
                                                                                            v214 := GenS14{F0: (v3 && (v3 || (v25 > v25))), F1: 2, F2: 0}
                                                                                            _ = v214
                                                                                        } else {
                                                                                            if (((((("hello" + "hello") + ("\n\t" + "hello")) + "a") + "a") + "foo bar") <= ("foo bar" + "")) {
                                                                                                if false {
                                                                                                    v26 = -963
                                                                                                    v215 := GenS24{F0: v1, F1: (fi38 | int(v25))}
                                                                                                    v216 := []int{492, (-(fi35)), fi35}
                                                                                                    _ = v215
                                                                                                    _ = v216
                                                                                                }
                                                                                                if true {
                                                                                                    v217 := []int{(fi38 / ((fi35) | 1)), (fi38 % (((v2 << ((986) & 63))) | 1)), fi35}
                                                                                                    _ = v217
                                                                                                }
                                                                                            } else {
                                                                                                var v218 uint16 = v48
                                                                                                switch "hello" {
                                                                                                case "":
                                                                                                    _ = len(v4)
                                                                                                }
                                                                                                v201 = (v39 - v201)
                                                                                                _ = v218
                                                                                            }
                                                                                            var v219 int16 = int16(fi35)
                                                                                            var v220 uint16 = (v48 >> (((v39 ^ 18446744073709551615)) & 63))
                                                                                            _ = v219
                                                                                            _ = v220
                                                                                        }
                                                                                        if (2 != v48) {
                                                                                            switch "foo bar" {
                                                                                            case "hello":
                                                                                                v199 = ((v199 + v199) - v199)
                                                                                            case "\n\t":
                                                                                                var v221 float32 = float32(v199)
                                                                                                _ = v221
                                                                                            default:
                                                                                                if (!v3) {
                                                                                                    var v222 uint16 = 65535
                                                                                                    _ = v222
                                                                                                }
                                                                                            }
                                                                                            if ((int(v25) > fi38) && v40) {
                                                                                                switch int(v52) {
                                                                                                case -1:
                                                                                                    if v3 {
                                                                                                        var v223 uint16 = (^(v48))
                                                                                                        switch (("foo bar" + "a") + "foo bar") {
                                                                                                        case "foo bar":
                                                                                                            switch 4294967295 {
                                                                                                            case 2147483647:
                                                                                                                var v224 int8 = v1
                                                                                                                _ = v224
                                                                                                            default:
                                                                                                                var v225 uint8 = 127
                                                                                                                _ = v225
                                                                                                            }
                                                                                                        case "\n\t":
                                                                                                            switch (fi35 >> ((v39) & 63)) {
                                                                                                            case -369:
                                                                                                                var v226 uint8 = 0
                                                                                                                _ = v226
                                                                                                            case 32767:
                                                                                                                v4[fi35] = (v48 != (v48 | (v223 << ((4294967295) & 63))))
                                                                                                                v5 = uint8(v25)
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v223
                                                                                                    }
                                                                                                }
                                                                                                switch ("a" + "a") {
                                                                                                case "\n\t":
                                                                                                    v40 = (0.1 > 1.5)
                                                                                                case "foo bar":
                                                                                                    if v3 {
                                                                                                        var v227 uint = uint(v25)
                                                                                                        _ = v227
                                                                                                    } else {
                                                                                                        if false {
                                                                                                            v228 := GenS6{F0: v51}
                                                                                                            var v229 string = ("" + ("" + (("\n\t" + "\n\t") + "")))
                                                                                                            _ = v228
                                                                                                            _ = v229
                                                                                                        }
                                                                                                        switch ("" + ("a" + ("hello" + "a"))) {
                                                                                                        case "hello":
                                                                                                            v1 = (v1 << ((v39) & 63))
                                                                                                        case "\n\t":
                                                                                                            v230 := []int{(fi35 - 406), fi38, v26}
                                                                                                            _ = v230
                                                                                                        }
                                                                                                    }
                                                                                                    v231 := GenS18{F0: 65535}
                                                                                                    _ = v231
                                                                                                default:
                                                                                                    var v232 bool = true
                                                                                                    v4[(-(v26))] = (!(!(uint8(v25) < (v49 << ((v39) & 63)))))
                                                                                                    _ = v232
                                                                                                }
                                                                                                fmt.Println(v200, v48, fi38)
                                                                                                v233 := []int{fi35, fi35, v2}
                                                                                                var v234 uint8 = v49
                                                                                                _ = v233
                                                                                                _ = v234
                                                                                            }
                                                                                            v26 = fi38
                                                                                        } else {
                                                                                            if (int(v52) != 412) {
                                                                                                delete(v4, (fi38 ^ (v2 | v26)))
                                                                                                var v235 uint64 = v25
                                                                                                _ = v235
                                                                                            } else {
                                                                                                switch "" {
                                                                                                case "\n\t":
                                                                                                    var v236 uint8 = (^(v49))
                                                                                                    _ = v236
                                                                                                case "a":
                                                                                                    switch -9223372036854775808 {
                                                                                                    case 2147483647:
                                                                                                        var v237 int64 = 0
                                                                                                        _ = v237
                                                                                                    case 42:
                                                                                                        var v238 int16 = 32
                                                                                                        _ = v238
                                                                                                    case -964:
                                                                                                        if ((v25 & (v25 | v25)) <= 4294967295) {
                                                                                                            v239 := GenS25{F0: 9223372036854775807, F1: "\n\t"}
                                                                                                            v240 := []int{v2, (fi35 | fi38)}
                                                                                                            _ = v239
                                                                                                            _ = v240
                                                                                                        } else {
                                                                                                            var v241 uint16 = (v48 * v48)
                                                                                                            var v242 int32 = -2147483648
                                                                                                            v243 := []int{(fi38 * (fi35 + v2)), (fi35 | fi35), 32767}
                                                                                                            _ = v241
                                                                                                            _ = v242
                                                                                                            _ = v243
                                                                                                        }
                                                                                                    }
                                                                                                case "":
                                                                                                    v5 -= 0
                                                                                                default:
                                                                                                    var v244 float64 = v199
                                                                                                    _ = v244
                                                                                                }
                                                                                                fmt.Println(v40, v3, v39)
                                                                                            }
                                                                                            switch (fi38 << ((4294967295) & 63)) {
                                                                                            case -748:
                                                                                                var v245 uint8 = (^(v5))
                                                                                                _ = v245
                                                                                            case 42:
                                                                                                var v246 uint16 = 497
                                                                                                _ = v246
                                                                                            }
                                                                                            v49 = 0
                                                                                        }
                                                                                    }
                                                                                    if ((v5 | 0) >= 42) {
                                                                                        v3 = ((-733 > -654) || v3)
                                                                                    }
                                                                                    v201 = uint(v48)
                                                                                }
                                                                                v5 = (v49 - v49)
                                                                                _ = v200
                                                                                _ = v201
                                                                            }
                                                                        } else {
                                                                            if v3 {
                                                                                switch -105 {
                                                                                case -2147483648:
                                                                                    switch v26 {
                                                                                    case 198:
                                                                                        v247 := []int{-9223372036854775808, (fi35 >> (((v39 * (v39 / ((2) | 1)))) & 63)), 2147483647, v26}
                                                                                        _ = v247
                                                                                    case -57:
                                                                                        switch ("hello" + ("a" + ((("hello" + "") + "a") + ("hello" + "foo bar")))) {
                                                                                        case "\n\t":
                                                                                            var v248 int16 = 42
                                                                                            v249 := GenS10{F0: 0, F1: v25}
                                                                                            v250 := []int{}
                                                                                            if v3 {
                                                                                                var v251 int32 = -2147483648
                                                                                                _ = v251
                                                                                            } else {
                                                                                                v1 += v52
                                                                                            }
                                                                                            if ((!false) && ((v39 | v39) < v39)) {
                                                                                                v26 += v26
                                                                                                var v252 int32 = -1
                                                                                                _ = v252
                                                                                            } else {
                                                                                                switch "" {
                                                                                                case "\n\t":
                                                                                                    var v253 int64 = 167
                                                                                                    _ = v253
                                                                                                case "hello":
                                                                                                    if (0.0 < 0.0) {
                                                                                                        v254 := GenS25{F0: 0, F1: ("\n\t" + "foo bar")}
                                                                                                        v26 = 0
                                                                                                        _ = v254
                                                                                                    } else {
                                                                                                        var v255 int64 = -821
                                                                                                        _ = v255
                                                                                                    }
                                                                                                case "":
                                                                                                    v3 = v3
                                                                                                }
                                                                                                v4[fi38] = v40
                                                                                                var v256 int16 = v51
                                                                                                if (v49 > v49) {
                                                                                                    var v257 uint = (v39 * 1)
                                                                                                    var v258 float64 = (3.14159 / v199)
                                                                                                    v25 -= v25
                                                                                                    _ = v257
                                                                                                    _ = v258
                                                                                                }
                                                                                                _ = v256
                                                                                            }
                                                                                            _ = v248
                                                                                            _ = v249
                                                                                            _ = v250
                                                                                        case "hello":
                                                                                            v49 -= (v49 % ((v49) | 1))
                                                                                        }
                                                                                    case 4294967295:
                                                                                        switch (fi35 >> (((v39 >> (((v39 | v39)) & 63))) & 63)) {
                                                                                        case 128:
                                                                                            v259 := GenS24{F0: v52, F1: 42}
                                                                                            _ = v259
                                                                                        default:
                                                                                            v1 -= 127
                                                                                        }
                                                                                        var v260 bool = false
                                                                                        if (!v3) {
                                                                                            var v261 float64 = ((1e10 - 1.5) - v199)
                                                                                            v262 := GenS13{F0: int(v48), F1: (^(v39))}
                                                                                            _ = v261
                                                                                            _ = v262
                                                                                        }
                                                                                        _ = v260
                                                                                    default:
                                                                                        v48 -= (v48 + (^(v48)))
                                                                                    }
                                                                                }
                                                                                var v263 float32 = (1e10 - ((-2.25 + (3.14159 - 3.14159)) * 0.1))
                                                                                _ = v263
                                                                            }
                                                                            switch "a" {
                                                                            case "":
                                                                                v52 = v52
                                                                            case "foo bar":
                                                                                switch (v2 & 2147483647) {
                                                                                case 9223372036854775807:
                                                                                    var v264 int64 = 9223372036854775807
                                                                                    v265 := GenS26{F0: -2147483648, F1: ((!v40) && (("hello" + "a") == ("foo bar" + "hello")))}
                                                                                    v266 := []int{32768}
                                                                                    _ = v264
                                                                                    _ = v265
                                                                                    _ = v266
                                                                                case 57:
                                                                                    v267 := GenS27{F0: v48}
                                                                                    _ = v267
                                                                                default:
                                                                                    v268 := GenS5{F0: v52}
                                                                                    _ = v268
                                                                                }
                                                                            case "a":
                                                                                v269 := []int{0, 207, v2}
                                                                                _ = v269
                                                                            }
                                                                            var v270 int16 = v51
                                                                            v271 := GenS22{F0: v49}
                                                                            _ = v270
                                                                            _ = v271
                                                                        }
                                                                        _ = v199
                                                                    }
                                                                    v272 := GenS7{F0: v39, F1: 1, F2: v3}
                                                                    v48 = v48
                                                                    _ = v272
                                                                }
                                                            }
                                                            if (!v3) {
                                                                v273 := GenS25{F0: 1, F1: (("hello" + "a") + "foo bar")}
                                                                v274 := GenS9{F0: 2, F1: v39}
                                                                var v275 float64 = 1.5
                                                                _ = v273
                                                                _ = v274
                                                                _ = v275
                                                            } else {
                                                                v1 -= v52
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    var v276 float32 = (0.1 + ((-2.25 + 1.5) / float32(float64(3.14159))))
                                                                    _ = v276
                                                                case "a":
                                                                    if v40 {
                                                                        v277 := []int{int(v25)}
                                                                        _ = v277
                                                                    }
                                                                default:
                                                                    switch fi38 {
                                                                    case 32768:
                                                                        if v40 {
                                                                            var v278 string = "a"
                                                                            v26 = (v26 % ((664) | 1))
                                                                            v51 = -32768
                                                                            fmt.Println(v51, v49, v26)
                                                                            if (v278 != "hello") {
                                                                                if v3 {
                                                                                    var v279 uint16 = v48
                                                                                    var v280 uint8 = (v49 << (((v39 << ((v39) & 63))) & 63))
                                                                                    _ = v279
                                                                                    _ = v280
                                                                                } else {
                                                                                    v281 := []int{}
                                                                                    var v282 float32 = (0.0 + ((1.5 * ((3.14159 + -2.25) / (((1e10 / 1.0) * 3.14159) * -2.25))) - 1e10))
                                                                                    _ = v281
                                                                                    _ = v282
                                                                                }
                                                                                v48 = 65535
                                                                                var v283 uint8 = v49
                                                                                _ = v283
                                                                            }
                                                                            _ = v278
                                                                        } else {
                                                                            v284 := GenS28{F0: 9223372036854775807, F1: -524}
                                                                            v285 := GenS29{F0: 829, F1: 2, F2: v40}
                                                                            fmt.Println(v5, fi35)
                                                                            fmt.Println(fi35, v3, v26)
                                                                            v26 = fi35
                                                                            _ = v284
                                                                            _ = v285
                                                                        }
                                                                    case 2147483647:
                                                                        var v286 bool = true
                                                                        _ = v286
                                                                    case 557:
                                                                        switch ("a" + ((((("foo bar" + "foo bar") + "foo bar") + "\n\t") + ("a" + "")) + "foo bar")) {
                                                                        case "\n\t":
                                                                            var v287 uint8 = (v49 % (((v5 * (v49 / (((v49 & v5)) | 1)))) | 1))
                                                                            _ = v287
                                                                        case "foo bar":
                                                                            fmt.Println(v26)
                                                                            v1 = v1
                                                                        case "":
                                                                            switch (fi38 * (fi35 + fi35)) {
                                                                            case 127:
                                                                                v288 := []int{}
                                                                                _ = v288
                                                                            case 2147483647:
                                                                                v289 := GenS19{F0: v49}
                                                                                _ = v289
                                                                            case -1:
                                                                                var v290 float32 = float32(3.14159)
                                                                                _ = v290
                                                                            default:
                                                                                v48 = v48
                                                                            }
                                                                        default:
                                                                            v51 = 32767
                                                                        }
                                                                    }
                                                                }
                                                                v291 := []int{9223372036854775807, -9223372036854775808, (fi35 & int(v1))}
                                                                fmt.Println(v40, v2)
                                                                _ = v291
                                                            }
                                                        } else {
                                                            var v292 float32 = 0.0
                                                            _ = v292
                                                        }
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v3, v48)
                                            }
                                        } else {
                                            switch (fi35 / ((fi35) | 1)) {
                                            case 174:
                                                v293 := GenS7{F0: v39, F1: 9223372036854775807, F2: ((1.5 - 0.1) == 0.1)}
                                                _ = v293
                                            default:
                                                if v3 {
                                                    delete(v4, 32768)
                                                    var v294 int32 = 42
                                                    var v295 float32 = 3.14159
                                                    _ = v294
                                                    _ = v295
                                                }
                                            }
                                            switch (v2 << (((v39 & v39)) & 63)) {
                                            case -2147483648:
                                                v25 += v25
                                            case 384:
                                                var v296 uint64 = (v25 * uint64(v49))
                                                _ = v296
                                            case 32767:
                                                if (42 == 42) {
                                                    var v297 float64 = 3.14159
                                                    var v298 int64 = 0
                                                    v299 := GenS26{F0: fi38, F1: v3}
                                                    _ = v297
                                                    _ = v298
                                                    _ = v299
                                                } else {
                                                    v40 = ((v5 | v49) < (v5 / (((v49 >> (((v39 % (((v39 & (v39 + 18446744073709551615))) | 1))) & 63))) | 1)))
                                                    switch ("" + "hello") {
                                                    case "":
                                                        var v300 int32 = 0
                                                        _ = v300
                                                    case "foo bar":
                                                        fmt.Println(v49, fi35, v51)
                                                    case "a":
                                                        v48 += (v48 & 2)
                                                    default:
                                                        if false {
                                                            if v3 {
                                                                switch (-(v2)) {
                                                                case 29:
                                                                    fmt.Println(v26, v51, fi38)
                                                                case 852:
                                                                    v1 = -88
                                                                default:
                                                                    v301 := []int{(fi38 >> (((^(v39))) & 63)), fi38, fi35}
                                                                    _ = v301
                                                                }
                                                            }
                                                            if (true && (v39 != (^(v39)))) {
                                                                var v302 uint32 = 880
                                                                var v303 int64 = 0
                                                                fmt.Println(v51, v25, v1)
                                                                v304 := GenS30{F0: 415, F1: v302}
                                                                if true {
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        v2 -= fi38
                                                                    case "a":
                                                                        switch "a" {
                                                                        case "a":
                                                                            var v305 int64 = (-(v303))
                                                                            _ = v305
                                                                        case "foo bar":
                                                                            var v306 int8 = (-(v1))
                                                                            _ = v306
                                                                        default:
                                                                            if v40 {
                                                                                if (v40 && (int32(v25) == 755)) {
                                                                                    if v40 {
                                                                                        fmt.Println(v1, v52, v40)
                                                                                        v51 = v51
                                                                                        var v307 float32 = 0.0
                                                                                        _ = v307
                                                                                    }
                                                                                } else {
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        var v308 float64 = (-2.25 - ((0.1 / -2.25) / -2.25))
                                                                                        _ = v308
                                                                                    case "a":
                                                                                        v1 -= v1
                                                                                    case "hello":
                                                                                        var v309 int64 = (v303 * v303)
                                                                                        _ = v309
                                                                                    default:
                                                                                        var v310 uint = v39
                                                                                        _ = v310
                                                                                    }
                                                                                    switch (fi35 & (v2 << (((^(v39))) & 63))) {
                                                                                    case -942:
                                                                                        v311 := GenS6{F0: (v51 & int16(v48))}
                                                                                        _ = v311
                                                                                    case 32767:
                                                                                        var v312 uint = 2
                                                                                        _ = v312
                                                                                    }
                                                                                }
                                                                                var v313 int64 = (v303 & v303)
                                                                                _ = v313
                                                                            }
                                                                        }
                                                                    }
                                                                    v314 := GenS31{F0: v48}
                                                                    if (!v40) {
                                                                        var v315 uint64 = v25
                                                                        var v316 int8 = 100
                                                                        if v40 {
                                                                            if v40 {
                                                                                switch 2147483647 {
                                                                                case -912:
                                                                                    var v317 int64 = v303
                                                                                    _ = v317
                                                                                case 128:
                                                                                    switch -9223372036854775808 {
                                                                                    case 42:
                                                                                        fmt.Println(v2, v51)
                                                                                    }
                                                                                default:
                                                                                    switch "a" {
                                                                                    case "a":
                                                                                        if v40 {
                                                                                            v2 = v2
                                                                                            if true {
                                                                                                fmt.Println(v315, v5)
                                                                                                v40 = true
                                                                                            } else {
                                                                                                v51 = int16(v52)
                                                                                                v318 := []int{42}
                                                                                                v319 := GenS27{F0: uint16(v315)}
                                                                                                _ = v318
                                                                                                _ = v319
                                                                                            }
                                                                                            var v320 bool = ((v316 & v52) >= (v52 ^ (v52 >> ((v39) & 63))))
                                                                                            _ = v320
                                                                                        } else {
                                                                                            v303 += v303
                                                                                            delete(v4, -2147483648)
                                                                                        }
                                                                                    case "hello":
                                                                                        var v321 int16 = v51
                                                                                        _ = v321
                                                                                    case "":
                                                                                        fmt.Println(v26)
                                                                                        switch fi38 {
                                                                                        case 9223372036854775807:
                                                                                            v302 += 42
                                                                                        case 252:
                                                                                            if false {
                                                                                                v49 = v49
                                                                                            } else {
                                                                                                _ = len(v4)
                                                                                            }
                                                                                        }
                                                                                        var v322 uint64 = v315
                                                                                        _ = v322
                                                                                    default:
                                                                                        v323 := GenS32{F0: 42, F1: uint8(v302), F2: -649}
                                                                                        _ = v323
                                                                                    }
                                                                                }
                                                                                var v324 uint16 = 32767
                                                                                _ = v324
                                                                            } else {
                                                                                var v325 uint = (v39 & v39)
                                                                                var v326 uint16 = 65535
                                                                                v326 = 32767
                                                                                _ = v325
                                                                                _ = v326
                                                                            }
                                                                        } else {
                                                                            var v327 float64 = 1e10
                                                                            var v328 uint = 1
                                                                            _ = v327
                                                                            _ = v328
                                                                        }
                                                                        v329 := GenS3{F0: v39}
                                                                        v316 -= 0
                                                                        var v330 bool = true
                                                                        fmt.Println(v26, v315)
                                                                        _ = v315
                                                                        _ = v316
                                                                        _ = v329
                                                                        _ = v330
                                                                    } else {
                                                                        v3 = v40
                                                                        if (!v40) {
                                                                            v1 -= (-(v52))
                                                                            var v331 int16 = 0
                                                                            v332 := GenS33{F0: ((v39 - (v39 * v39)) == 65535)}
                                                                            _ = v331
                                                                            _ = v332
                                                                        } else {
                                                                            var v333 float32 = 1.5
                                                                            v5 = (v49 + v49)
                                                                            _ = v333
                                                                        }
                                                                        if false {
                                                                            v334 := GenS6{F0: v51}
                                                                            _ = v334
                                                                        } else {
                                                                            var v335 int64 = int64(v302)
                                                                            v336 := []int{fi38, fi38, 9223372036854775807, v26}
                                                                            _ = v335
                                                                            _ = v336
                                                                        }
                                                                    }
                                                                    _ = v314
                                                                }
                                                                var v337 float32 = (0.0 * 1e10)
                                                                _ = v302
                                                                _ = v303
                                                                _ = v304
                                                                _ = v337
                                                            } else {
                                                                var v338 uint32 = 2
                                                                switch v26 {
                                                                case 32768:
                                                                    v48 = uint16(v338)
                                                                case 127:
                                                                    var v339 uint16 = 42
                                                                    _ = v339
                                                                default:
                                                                    v340 := []int{fi35}
                                                                    _ = v340
                                                                }
                                                                if v3 {
                                                                    var v341 uint32 = (v338 ^ (v338 & v338))
                                                                    var v342 uint = 255
                                                                    delete(v4, 32768)
                                                                    _ = v341
                                                                    _ = v342
                                                                } else {
                                                                    var v343 int = v26
                                                                    fmt.Println(v1, v49)
                                                                    _ = v343
                                                                }
                                                                _ = v338
                                                            }
                                                            var v344 uint8 = v5
                                                            _ = v344
                                                        } else {
                                                            v345 := GenS15{F0: 1, F1: (fi35 >> (((v39 << (((v39 << (((v39 << ((v39) & 63))) & 63))) & 63))) & 63))}
                                                            _ = v345
                                                        }
                                                    }
                                                    var v346 int32 = -897
                                                    fmt.Println(v49)
                                                    _ = v346
                                                }
                                            }
                                            var v347 uint64 = v25
                                            _ = v347
                                        }
                                    }
                                    switch fi38 {
                                    case -699:
                                        var v348 int = fi38
                                        _ = v348
                                    case 318:
                                        v49 -= (^(v49))
                                    }
                                    _ = v51
                                    _ = v52
                                } else {
                                    var v349 bool = v3
                                    v49 = (v49 / (((v5 / (((v49 & v5)) | 1))) | 1))
                                    var v350 uint32 = 2147483647
                                    var v351 uint16 = (v48 >> ((v39) & 63))
                                    switch -409 {
                                    case 4294967295:
                                        switch (fi35 - v2) {
                                        case -9223372036854775808:
                                            v2 += fi38
                                        case 128:
                                            if ((v25 >= (v25 | 18446744073709551615)) && v40) {
                                                var v352 uint32 = v350
                                                _ = v352
                                            }
                                        case 4294967295:
                                            v353 := GenS34{F0: v349, F1: v40, F2: v350}
                                            _ = v353
                                        default:
                                            v354 := GenS27{F0: (v351 - (^(v48)))}
                                            _ = v354
                                        }
                                    case 491:
                                        fmt.Println(v40)
                                    }
                                    if (!(v349 && ((v25 - v25) != (v25 ^ (v25 / (((v25 ^ v25)) | 1)))))) {
                                        var v355 int16 = 1000
                                        switch "hello" {
                                        case "":
                                            v1 = (v1 + v1)
                                        }
                                        v1 = (v1 + int8(v355))
                                        _ = v355
                                    } else {
                                        v356 := GenS23{F0: 2, F1: ("hello" + "hello"), F2: ("hello" + "hello")}
                                        v4[fi35] = v349
                                        _ = v356
                                    }
                                    _ = v349
                                    _ = v350
                                    _ = v351
                                }
                            }
                            v25 -= v25
                            switch (fi35 ^ v2) {
                            case -2147483648:
                                switch (v2 << ((v39) & 63)) {
                                case 367:
                                    var v357 bool = (0 > 0)
                                    _ = v357
                                }
                            default:
                                v358 := GenS16{F0: 255, F1: fi35}
                                if (575 > 65535) {
                                    v26 = fi38
                                    switch fi38 {
                                    case 0:
                                        fmt.Println(fi35, v49)
                                    case -83:
                                        if v40 {
                                            if v3 {
                                                var v359 int16 = 1000
                                                _ = v359
                                            } else {
                                                if v3 {
                                                    v360 := []int{(fi35 - 128), 127}
                                                    v5 = uint8(v39)
                                                    var v361 uint8 = (v5 / (((v5 - v49)) | 1))
                                                    _ = v360
                                                    _ = v361
                                                }
                                                var v362 float32 = 0.0
                                                _ = v362
                                            }
                                            v25 = (v25 / (((v25 << ((v39) & 63))) | 1))
                                            var v363 uint64 = (v25 | v25)
                                            _ = v363
                                        } else {
                                            fmt.Println(v1)
                                            if true {
                                                v25 = v25
                                                fmt.Println(fi38, v5, v1)
                                                switch "hello" {
                                                case "foo bar":
                                                    v25 = (v25 & (v25 - (v25 << (((v39 >> ((68) & 63))) & 63))))
                                                }
                                            }
                                            switch (^(fi38)) {
                                            case -553:
                                                v364 := []int{(v26 + (fi35 >> (((^(v39))) & 63)))}
                                                _ = v364
                                            default:
                                                var v365 bool = (v39 < 255)
                                                _ = v365
                                            }
                                        }
                                    case 128:
                                        v48 = (^(v48))
                                    default:
                                        var v366 float32 = (3.14159 - -2.25)
                                        _ = v366
                                    }
                                }
                                _ = v358
                            }
                            _ = fi38
                            _ = v39
                            _ = v40
                            _ = v48
                            _ = v49
                        }
                    }
                case 9223372036854775807:
                    v367 := GenS35{F0: 127, F1: (v5 & v5)}
                    _ = v367
                }
                var v368 float32 = (float32(((1e10 + float64(0.0)) / 1.0)) * 3.14159)
                v369 := v4[(fi35 << ((916) & 63))]
                _ = v368
                _ = v369
            }
            v370 := GenS12{F0: 1000, F1: v3}
            fmt.Println(v25, v1, v1)
            _ = v370
        } else {
            v25 = v25
        }
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
