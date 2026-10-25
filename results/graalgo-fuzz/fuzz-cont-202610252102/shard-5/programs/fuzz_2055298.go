package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 int8; F2 int8 }
type GenS5 struct { F0 int; F1 uint32; F2 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int64; F1 int64; F2 uint8 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int64; F1 int16; F2 string }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 string; F1 bool; F2 uint }
type GenS16 struct { F0 bool; F1 int64 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 int; F1 string }
type GenS19 struct { F0 int16; F1 bool }
type GenS20 struct { F0 uint32; F1 string }
type GenS21 struct { F0 uint16 }
type GenS22 struct { F0 int16; F1 uint16 }
type GenS23 struct { F0 int32; F1 string }
type GenS24 struct { F0 int16; F1 uint }
type GenS25 struct { F0 int64 }
type GenS26 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS27 struct { F0 uint16; F1 uint8 }
type GenS28 struct { F0 int32 }
type GenS29 struct { F0 int32; F1 int64; F2 int }
type GenS30 struct { F0 string; F1 uint }
type GenS31 struct { F0 uint32; F1 bool; F2 bool }
func genFn0(p0 int8, p1 bool, p2 uint) string {
    for fi5 := 0; fi5 < 13; fi5++ {
        p2 = (p2 / ((255) | 1))
        var v6 bool = p1
        _ = fi5
        _ = v6
    }
    p1 = (65535 == 42)
    _ = p0
    _ = p1
    _ = p2
    return "a"
}

func main() {
    var v1 int32 = 727
    var v2 uint32 = 81
    v3 := make(map[uint32]bool)
    var v4 bool = true
    var v7 float32 = 3.14159
    v4 = v4
    delete(v3, 1)
    switch int(v2) {
    case 1:
        fmt.Println(v4)
    default:
        switch 32767 {
        case 1:
            v7 = 0.0
        case -9223372036854775808:
            if true {
                switch 468 {
                case 42:
                    if v4 {
                        switch (("" + ("foo bar" + ("\n\t" + ("foo bar" + (("\n\t" + "foo bar") + "\n\t"))))) + "hello") {
                        case "a":
                            v3[0] = ((42 > 42) && v4)
                        default:
                            if v4 {
                                if false {
                                    if true {
                                        v8 := v3[v2]
                                        for fi9 := 0; fi9 < 5; fi9++ {
                                            v10 := GenS1{F0: 4294967295}
                                            v2 = (^(v2))
                                            fmt.Println(v8, v1)
                                            _ = fi9
                                            _ = v10
                                        }
                                        _ = v8
                                    } else {
                                        var v11 int32 = -1
                                        v12 := []int{}
                                        _ = v11
                                        _ = v12
                                    }
                                    var v13 uint = 2
                                    _ = v13
                                } else {
                                    v14 := GenS2{F0: v1}
                                    _ = v14
                                }
                            }
                        }
                    }
                case 128:
                    var v15 uint8 = 2
                    _ = v15
                case 941:
                    v16 := []int{}
                    _ = v16
                default:
                    v1 += (v1 >> ((65535) & 63))
                }
            }
        case 9223372036854775807:
            var v17 int32 = v1
            _ = v17
        default:
            if v4 {
                v7 += (v7 - (3.14159 * v7))
                v18 := []int{-2147483648}
                _ = v18
            }
        }
    }
    if v4 {
        v1 = (v1 + v1)
        if (!v4) {
            if ((v1 % ((946) | 1)) > (v1 / ((v1) | 1))) {
                v2 = v2
                if (3.14159 <= -2.25) {
                    v7 = v7
                } else {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        v7 = (v7 / v7)
                        switch (^(fi19)) {
                        case 1:
                            _ = len(v3)
                        case -9223372036854775808:
                            for fi20 := 0; fi20 < 7; fi20++ {
                                v21 := GenS2{F0: 0}
                                var v22 float32 = (v7 * ((1e10 / v7) / (v7 / v7)))
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        case -368:
                            if v4 {
                                v2 = v2
                                var v23 bool = v4
                                _ = v23
                            } else {
                                v24 := []int{fi19, fi19, 1, fi19}
                                v2 = (^(v2))
                                var v25 uint = 4294967295
                                _ = v24
                                _ = v25
                            }
                        default:
                            fmt.Println(v1)
                        }
                        _ = fi19
                    }
                    if (-2.25 > v7) {
                        if v4 {
                            var v26 uint16 = 65535
                            _ = v26
                        } else {
                            var v27 int8 = 0
                            if (1 != -1) {
                                switch -476 {
                                case 4294967295:
                                    if v4 {
                                        var v28 bool = (65535 > 18446744073709551615)
                                        v29 := GenS3{F0: int32(v27)}
                                        _ = v28
                                        _ = v29
                                    } else {
                                        var v30 uint = 570
                                        _ = v30
                                    }
                                case -9223372036854775808:
                                    v31 := GenS4{F0: 42, F1: -33, F2: v27}
                                    _ = v31
                                case 9223372036854775807:
                                    if v4 {
                                        for fi32 := 0; fi32 < 15; fi32++ {
                                            if v4 {
                                                var v33 uint16 = 2
                                                fmt.Println(v33, v2)
                                                _ = v33
                                            }
                                            var v34 uint64 = 332
                                            v27 = (v27 | 0)
                                            _ = fi32
                                            _ = v34
                                        }
                                    } else {
                                        for fi35 := 0; fi35 < 13; fi35++ {
                                            v36 := GenS5{F0: (fi35 | fi35), F1: (v2 | (v2 | (^(v2)))), F2: -896}
                                            if (-32768 > 42) {
                                                v7 = 0.1
                                                v3[v2] = v4
                                            }
                                            for fi37 := 0; fi37 < 18; fi37++ {
                                                var v38 bool = (!((v4 && false) && v4))
                                                _ = fi37
                                                _ = v38
                                            }
                                            _ = fi35
                                            _ = v36
                                        }
                                        if ((3.14159 >= float32(-2.25)) || ((643 <= 0) || v4)) {
                                            var v39 uint8 = 61
                                            if v4 {
                                                if (36 != 231) {
                                                    switch 128 {
                                                    case -405:
                                                        var v40 float32 = 0.0
                                                        _ = v40
                                                    case 14:
                                                        v27 -= v27
                                                    }
                                                }
                                            } else {
                                                v4 = true
                                                v41 := GenS2{F0: v1}
                                                var v42 int8 = (v27 / ((v27) | 1))
                                                _ = v41
                                                _ = v42
                                            }
                                            if (("hello" > ("a" + "\n\t")) && false) {
                                                var v43 uint64 = 860
                                                if (0 <= -203) {
                                                    var v44 int = 9223372036854775807
                                                    if (v4 && ((v4 && (!(413 <= 42))) && (!v4))) {
                                                        for fi45 := 0; fi45 < 20; fi45++ {
                                                            v46 := GenS2{F0: 0}
                                                            _ = fi45
                                                            _ = v46
                                                        }
                                                    } else {
                                                        v47 := GenS5{F0: (v44 / ((v44) | 1)), F1: v2, F2: int64(v39)}
                                                        var v48 uint16 = 42
                                                        v27 -= v27
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                    for fi49 := 0; fi49 < 17; fi49++ {
                                                        fmt.Println(v1, v43)
                                                        _ = fi49
                                                    }
                                                    _ = v44
                                                } else {
                                                    var v50 int16 = 1000
                                                    fmt.Println(v1, v43, v39)
                                                    v51 := []int{-440, -567}
                                                    _ = v50
                                                    _ = v51
                                                }
                                                _ = v43
                                            } else {
                                                var v52 int8 = (v27 / (((v27 << ((2) & 63))) | 1))
                                                _ = v52
                                            }
                                            _ = v39
                                        }
                                        v27 -= 100
                                    }
                                }
                                var v53 uint64 = 18446744073709551615
                                v2 = 1
                                _ = v53
                            }
                            v54 := []int{879, 237}
                            _ = v27
                            _ = v54
                        }
                        for fi55 := 0; fi55 < 11; fi55++ {
                            var v56 float32 = (v7 - ((-2.25 / 1.0) + (v7 - v7)))
                            if true {
                                for fi57 := 0; fi57 < 18; fi57++ {
                                    v2 = v2
                                    switch "\n\t" {
                                    case "\n\t":
                                        v1 = v1
                                    case "foo bar":
                                        _ = len(v3)
                                    case "a":
                                        if (!((164 > 255) && (1000 <= -490))) {
                                            v56 += (v56 - 1.5)
                                            if v4 {
                                                var v58 uint64 = 667
                                                fmt.Println(fi55, v4, v1)
                                                v59 := []int{}
                                                if (744 >= 32767) {
                                                    var v60 uint = 65535
                                                    _ = v60
                                                } else {
                                                    if v4 {
                                                        delete(v3, v2)
                                                    }
                                                    v61 := v3[v2]
                                                    _ = v61
                                                }
                                                _ = v58
                                                _ = v59
                                            } else {
                                                if false {
                                                    v7 -= (1e10 * v56)
                                                } else {
                                                    v2 -= v2
                                                }
                                                switch "foo bar" {
                                                case "hello":
                                                    if (1 < -1) {
                                                        switch (("a" + ("" + "hello")) + "hello") {
                                                        case "\n\t":
                                                            v7 = v7
                                                        case "a":
                                                            fmt.Println(v56, v1)
                                                        case "":
                                                            switch fi57 {
                                                            case -730:
                                                                var v62 bool = (0.1 > 1e10)
                                                                if v4 {
                                                                    v4 = v62
                                                                    fmt.Println(v62)
                                                                    delete(v3, 2)
                                                                } else {
                                                                    var v63 uint64 = 2
                                                                    v64 := GenS4{F0: 1, F1: 100, F2: 127}
                                                                    v63 += v63
                                                                    _ = v63
                                                                    _ = v64
                                                                }
                                                                var v65 int8 = 100
                                                                _ = v62
                                                                _ = v65
                                                            case 32767:
                                                                fmt.Println(v7, v7)
                                                            default:
                                                                var v66 uint64 = 131
                                                                _ = v66
                                                            }
                                                        default:
                                                            if (4294967295 > 18446744073709551615) {
                                                                v7 += (v7 / (0.0 - v7))
                                                                v4 = v4
                                                                v4 = (true || v4)
                                                                fmt.Println(v7, fi55, v4)
                                                            }
                                                        }
                                                        var v67 int32 = (v1 | (^(v1)))
                                                        if (127 != 9223372036854775807) {
                                                            if true {
                                                                v68 := []int{-216, 32767}
                                                                fmt.Println(v2, v4, v1)
                                                                if ((("a" + "hello") + ("a" + "")) != ("" + "\n\t")) {
                                                                    var v69 float64 = 0.0
                                                                    v1 += 0
                                                                    _ = v69
                                                                } else {
                                                                    var v70 int64 = 9223372036854775807
                                                                    v71 := GenS2{F0: v67}
                                                                    var v72 uint64 = 1
                                                                    _ = v70
                                                                    _ = v71
                                                                    _ = v72
                                                                }
                                                                _ = v68
                                                            }
                                                            switch ("a" + ((((("a" + "\n\t") + "hello") + "a") + "foo bar") + "")) {
                                                            case "a":
                                                                var v73 float32 = v56
                                                                _ = v73
                                                            case "\n\t":
                                                                if v4 {
                                                                    delete(v3, v2)
                                                                    fmt.Println(v4, v56)
                                                                } else {
                                                                    var v74 uint32 = v2
                                                                    _ = v74
                                                                }
                                                            case "":
                                                                switch (^(fi57)) {
                                                                case 32767:
                                                                    if (826 > 262) {
                                                                        if (!(v2 < v2)) {
                                                                            v1 += 338
                                                                            var v75 int16 = 0
                                                                            delete(v3, v2)
                                                                            _ = v75
                                                                        }
                                                                        if ((v4 && v4) && (fi57 <= (fi55 & (^(fi55))))) {
                                                                            var v76 uint64 = 42
                                                                            fmt.Println(fi57, v4)
                                                                            _ = v76
                                                                        }
                                                                        if v4 {
                                                                            v3[2147483648] = v4
                                                                            switch (fi55 % ((0) | 1)) {
                                                                            case 0:
                                                                                switch "a" {
                                                                                case "":
                                                                                    if v4 {
                                                                                        fmt.Println(v67)
                                                                                        v77 := []int{0, 128}
                                                                                        var v78 int8 = 100
                                                                                        _ = v77
                                                                                        _ = v78
                                                                                    }
                                                                                case "a":
                                                                                    if (!v4) {
                                                                                        v79 := []int{}
                                                                                        var v80 int32 = -1
                                                                                        v7 = 0.1
                                                                                        _ = v79
                                                                                        _ = v80
                                                                                    }
                                                                                case "hello":
                                                                                    switch 356 {
                                                                                    case -505:
                                                                                        v81 := GenS1{F0: 0}
                                                                                        v82 := GenS6{F0: 1}
                                                                                        _ = v81
                                                                                        _ = v82
                                                                                    default:
                                                                                        var v83 uint64 = 0
                                                                                        _ = v83
                                                                                    }
                                                                                default:
                                                                                    fmt.Println(fi57)
                                                                                }
                                                                            case 42:
                                                                                _ = len(v3)
                                                                                if (!("\n\t" == "\n\t")) {
                                                                                    if false {
                                                                                        v84 := GenS7{F0: v4}
                                                                                        v3[v2] = v4
                                                                                        _ = v84
                                                                                    }
                                                                                    v85 := GenS8{F0: -1}
                                                                                    v2 -= (v2 / ((v2) | 1))
                                                                                    v67 += (v67 >> ((583) & 63))
                                                                                    _ = v85
                                                                                }
                                                                                v86 := GenS9{F0: v4, F1: 1, F2: 865}
                                                                                v4 = ((0 < -321) || v4)
                                                                                _ = v86
                                                                            }
                                                                            v1 -= (v1 - -2147483648)
                                                                        }
                                                                    }
                                                                case -9223372036854775808:
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        var v87 float64 = 0.1
                                                                        _ = v87
                                                                    case "foo bar":
                                                                        var v88 float64 = 0.1
                                                                        _ = v88
                                                                    case "hello":
                                                                        var v89 int8 = 89
                                                                        if true {
                                                                            var v90 uint16 = 2
                                                                            _ = v90
                                                                        }
                                                                        if v4 {
                                                                            _ = len(v3)
                                                                            if v4 {
                                                                                var v91 bool = v4
                                                                                v91 = false
                                                                                _ = v91
                                                                            } else {
                                                                                v92 := GenS9{F0: v4, F1: 122, F2: 0}
                                                                                fmt.Println(v2, v89)
                                                                                if true {
                                                                                    v67 = v1
                                                                                    fmt.Println(fi57)
                                                                                }
                                                                                _ = v92
                                                                            }
                                                                            var v93 bool = true
                                                                            _ = v93
                                                                        }
                                                                        _ = v89
                                                                    }
                                                                }
                                                            default:
                                                                var v94 uint8 = 167
                                                                if (false && ((v1 < v67) && (int16(v1) != -32768))) {
                                                                    var v95 string = "\n\t"
                                                                    var v96 int64 = -1
                                                                    if true {
                                                                        v3[(v2 << ((167) & 63))] = v4
                                                                        v97 := GenS8{F0: -615}
                                                                        if v4 {
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                v2 -= v2
                                                                            }
                                                                            v98 := GenS10{F0: 0}
                                                                            v99 := GenS7{F0: true}
                                                                            if (v4 && false) {
                                                                                v4 = true
                                                                            } else {
                                                                                if (65535 < 469) {
                                                                                    var v100 uint = 18446744073709551615
                                                                                    _ = v100
                                                                                } else {
                                                                                    v96 += (v96 % ((v96) | 1))
                                                                                    var v101 int64 = v96
                                                                                    _ = v101
                                                                                }
                                                                                var v102 int32 = v67
                                                                                _ = v102
                                                                            }
                                                                            _ = v98
                                                                            _ = v99
                                                                        }
                                                                        var v103 int = (fi57 * (fi57 ^ 9223372036854775807))
                                                                        _ = v103
                                                                        _ = v97
                                                                    }
                                                                    var v104 uint8 = (v94 - v94)
                                                                    _ = v104
                                                                    _ = v95
                                                                    _ = v96
                                                                }
                                                                var v105 int64 = 522
                                                                _ = v105
                                                                _ = v94
                                                            }
                                                            v4 = v4
                                                        } else {
                                                            if true {
                                                                if (v4 && v4) {
                                                                    _ = len(v3)
                                                                    if ((((3.14159 >= 1.5) || ((v4 && v4) && v4)) || (!(0 == (v67 >> ((2) & 63))))) && (("" + "hello") <= ("foo bar" + ""))) {
                                                                        var v106 float64 = 1.5
                                                                        v107 := GenS8{F0: -1}
                                                                        v108 := []int{fi57, -132, (fi57 / ((fi55) | 1))}
                                                                        _ = v106
                                                                        _ = v107
                                                                        _ = v108
                                                                    } else {
                                                                        var v109 string = "foo bar"
                                                                        delete(v3, (v2 << ((862) & 63)))
                                                                        _ = v109
                                                                    }
                                                                    fmt.Println(fi55, fi55)
                                                                }
                                                                v110 := []int{fi55, fi57}
                                                                if len(v110) > 0 {
                                                                    v110[0] = fi57
                                                                }
                                                                _ = v110
                                                            } else {
                                                                v1 = (^(v67))
                                                                switch (fi55 % (((fi57 << ((4294967295) & 63))) | 1)) {
                                                                case -1:
                                                                    v7 = (v56 - v56)
                                                                case -822:
                                                                    v111 := GenS9{F0: (0 >= 2), F1: 187, F2: 1}
                                                                    _ = v111
                                                                case 4294967295:
                                                                    switch -452 {
                                                                    case -2147483648:
                                                                        v56 = v56
                                                                    case 32767:
                                                                        fmt.Println(v56)
                                                                    }
                                                                default:
                                                                    var v112 int = (fi55 - (fi55 / ((-9223372036854775808) | 1)))
                                                                    _ = v112
                                                                }
                                                                switch (fi57 ^ (-(fi57))) {
                                                                case 32767:
                                                                    var v113 uint16 = 598
                                                                    _ = v113
                                                                case 0:
                                                                    var v114 int8 = int8(fi57)
                                                                    _ = v114
                                                                }
                                                            }
                                                        }
                                                        _ = v67
                                                    } else {
                                                        var v115 int16 = -1
                                                        _ = v115
                                                    }
                                                case "foo bar":
                                                    if (65535 >= 2147483648) {
                                                        v116 := []int{}
                                                        _ = v116
                                                    } else {
                                                        v56 = -2.25
                                                    }
                                                    v2 -= v2
                                                    v4 = false
                                                default:
                                                    fmt.Println(v7)
                                                    switch (fi55 % ((fi57) | 1)) {
                                                    case -2147483648:
                                                        delete(v3, (v2 | v2))
                                                    case -1:
                                                        v7 += (float32(float64(-2.25)) * ((v56 - (v7 + v56)) - v56))
                                                    default:
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            if v4 {
                                                                v3[v2] = (!v4)
                                                                var v117 uint = 65535
                                                                _ = v117
                                                            }
                                                        case "foo bar":
                                                            v118 := v3[473]
                                                            _ = v118
                                                        case "a":
                                                            var v119 uint64 = 579
                                                            v4 = ((v119 & v119) >= v119)
                                                            switch "a" {
                                                            case "":
                                                                switch int(v119) {
                                                                case 128:
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        _ = len(v3)
                                                                    case "":
                                                                        fmt.Println(v119, fi57, v119)
                                                                    case "a":
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            var v120 uint16 = 0
                                                                            _ = v120
                                                                        }
                                                                    default:
                                                                        var v121 int8 = 42
                                                                        if ((fi57 * (fi57 / ((fi55) | 1))) != 0) {
                                                                            fmt.Println(v56)
                                                                        }
                                                                        _ = v121
                                                                    }
                                                                case -685:
                                                                    if v4 {
                                                                        v2 += 2147483647
                                                                        var v122 uint64 = (v119 & v119)
                                                                        _ = v122
                                                                    } else {
                                                                        if v4 {
                                                                            var v123 uint32 = v2
                                                                            if (72 != 42) {
                                                                                v56 = v7
                                                                            }
                                                                            _ = v123
                                                                        } else {
                                                                            fmt.Println(v7)
                                                                            switch ((("hello" + "foo bar") + ("" + "hello")) + (("hello" + ("a" + "foo bar")) + ("" + ("" + "foo bar")))) {
                                                                            case "hello":
                                                                                var v124 int = fi55
                                                                                _ = v124
                                                                            }
                                                                        }
                                                                        _ = len(v3)
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            if (v119 >= v119) {
                                                                                var v125 int32 = -1
                                                                                v126 := GenS11{F0: 0, F1: -1, F2: 78}
                                                                                v127 := GenS1{F0: 255}
                                                                                _ = v125
                                                                                _ = v126
                                                                                _ = v127
                                                                            } else {
                                                                                if v4 {
                                                                                    var v128 int = fi55
                                                                                    _ = v128
                                                                                } else {
                                                                                    if (!true) {
                                                                                        var v129 uint16 = uint16(v2)
                                                                                        v119 -= v119
                                                                                        _ = v129
                                                                                    }
                                                                                    if false {
                                                                                        if true {
                                                                                            delete(v3, 2)
                                                                                            if (v7 != ((v56 - (float32(float64(3.14159)) + v56)) + 3.14159)) {
                                                                                                var v130 float32 = (0.1 - v7)
                                                                                                v119 = v119
                                                                                                _ = v130
                                                                                            } else {
                                                                                                var v131 string = ""
                                                                                                var v132 int32 = -1
                                                                                                _ = v131
                                                                                                _ = v132
                                                                                            }
                                                                                            if ((v2 / (((v2 - (v2 << ((2) & 63)))) | 1)) <= v2) {
                                                                                                fmt.Println(v2)
                                                                                            } else {
                                                                                                var v133 string = ("a" + "a")
                                                                                                v134 := GenS2{F0: v1}
                                                                                                v2 += v2
                                                                                                if (!true) {
                                                                                                    if v4 {
                                                                                                        v135 := GenS3{F0: v1}
                                                                                                        var v136 uint64 = uint64(v2)
                                                                                                        v137 := GenS12{F0: v133}
                                                                                                        v1 -= 0
                                                                                                        _ = v135
                                                                                                        _ = v136
                                                                                                        _ = v137
                                                                                                    }
                                                                                                    switch fi57 {
                                                                                                    case 449:
                                                                                                        if v4 {
                                                                                                            delete(v3, v2)
                                                                                                            var v138 int16 = -1
                                                                                                            _ = v138
                                                                                                        } else {
                                                                                                            var v139 int64 = -9223372036854775808
                                                                                                            _ = v139
                                                                                                        }
                                                                                                    case 4294967295:
                                                                                                        v140 := []int{(fi55 & (fi55 * 4294967295)), 128}
                                                                                                        _ = v140
                                                                                                    case -9223372036854775808:
                                                                                                        var v141 int8 = -1
                                                                                                        _ = v141
                                                                                                    }
                                                                                                    switch "a" {
                                                                                                    case "a":
                                                                                                        var v142 uint32 = (^(v2))
                                                                                                        _ = v142
                                                                                                    case "foo bar":
                                                                                                        v143 := GenS13{F0: -1, F1: -1, F2: "a"}
                                                                                                        _ = v143
                                                                                                    }
                                                                                                } else {
                                                                                                    if v4 {
                                                                                                        if v4 {
                                                                                                            v144 := GenS1{F0: 572}
                                                                                                            var v145 uint16 = 65535
                                                                                                            var v146 float32 = v7
                                                                                                            _ = v144
                                                                                                            _ = v145
                                                                                                            _ = v146
                                                                                                        } else {
                                                                                                            v147 := GenS13{F0: 1, F1: 642, F2: "hello"}
                                                                                                            _ = v147
                                                                                                        }
                                                                                                        v119 += (v119 - (v119 | (v119 >> ((42) & 63))))
                                                                                                    } else {
                                                                                                        if v4 {
                                                                                                            v148 := GenS10{F0: 1}
                                                                                                            var v149 int32 = 42
                                                                                                            v150 := []int{42, fi57, fi55}
                                                                                                            v149 = -2147483648
                                                                                                            _ = v148
                                                                                                            _ = v149
                                                                                                            _ = v150
                                                                                                        }
                                                                                                    }
                                                                                                    v151 := GenS14{F0: 42}
                                                                                                    if false {
                                                                                                        v119 -= 833
                                                                                                        v119 = v119
                                                                                                        v2 = 42
                                                                                                    }
                                                                                                    var v152 int32 = int32(v119)
                                                                                                    _ = v151
                                                                                                    _ = v152
                                                                                                }
                                                                                                if v4 {
                                                                                                    if true {
                                                                                                        if ((0 > -150) && true) {
                                                                                                            v153 := GenS15{F0: (("hello" + v133) + ""), F1: (true && (128 >= 128)), F2: 42}
                                                                                                            if v4 {
                                                                                                                v7 -= v56
                                                                                                                v2 = v2
                                                                                                            }
                                                                                                            if true {
                                                                                                                v154 := GenS5{F0: (fi55 + -9223372036854775808), F1: 224, F2: -842}
                                                                                                                _ = len(v3)
                                                                                                                _ = v154
                                                                                                            } else {
                                                                                                                var v155 float32 = v56
                                                                                                                _ = v155
                                                                                                            }
                                                                                                            if false {
                                                                                                                v156 := []int{32767, (-(fi57)), fi57, fi57}
                                                                                                                var v157 uint16 = 2
                                                                                                                _ = v156
                                                                                                                _ = v157
                                                                                                            }
                                                                                                            switch 4294967295 {
                                                                                                            case -765:
                                                                                                                var v158 uint16 = 517
                                                                                                                _ = v158
                                                                                                            case 639:
                                                                                                                if v4 {
                                                                                                                    v1 = v1
                                                                                                                    if ((-841 <= 0) || false) {
                                                                                                                        v159 := GenS16{F0: v4, F1: 666}
                                                                                                                        v160 := []int{fi57, fi55}
                                                                                                                        if (0.1 < (3.14159 + -2.25)) {
                                                                                                                            v7 -= v56
                                                                                                                        } else {
                                                                                                                            v161 := v3[(v2 | 301)]
                                                                                                                            switch "hello" {
                                                                                                                            case "a":
                                                                                                                                switch v133 {
                                                                                                                                case "\n\t":
                                                                                                                                    if v4 {
                                                                                                                                        if true {
                                                                                                                                            v2 = v2
                                                                                                                                        } else {
                                                                                                                                            _ = len(v3)
                                                                                                                                        }
                                                                                                                                        switch "\n\t" {
                                                                                                                                        case "hello":
                                                                                                                                            _ = len(v3)
                                                                                                                                        }
                                                                                                                                        var v162 int = (^(fi55))
                                                                                                                                        _ = v162
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    if len(v160) > 0 {
                                                                                                                                        v160[0] = int(v1)
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                if v4 {
                                                                                                                                    switch (v133 + v133) {
                                                                                                                                    case "foo bar":
                                                                                                                                        delete(v3, v2)
                                                                                                                                    case "":
                                                                                                                                        v163 := GenS1{F0: 483}
                                                                                                                                        _ = v163
                                                                                                                                    default:
                                                                                                                                        var v164 float64 = (((0.0 / 0.1) - (-2.25 + -2.25)) - ((-2.25 + float64(v56)) * 3.14159))
                                                                                                                                        _ = v164
                                                                                                                                    }
                                                                                                                                    var v165 string = ("hello" + ((v133 + "a") + "a"))
                                                                                                                                    _ = v165
                                                                                                                                } else {
                                                                                                                                    fmt.Println(v7, v1)
                                                                                                                                }
                                                                                                                            }
                                                                                                                            v1 = v1
                                                                                                                            _ = v161
                                                                                                                        }
                                                                                                                        v7 += -2.25
                                                                                                                        _ = v159
                                                                                                                        _ = v160
                                                                                                                    } else {
                                                                                                                        switch int(v2) {
                                                                                                                        case -30:
                                                                                                                            if (true || v4) {
                                                                                                                                switch v133 {
                                                                                                                                case "":
                                                                                                                                    v2 += v2
                                                                                                                                case "\n\t":
                                                                                                                                    fmt.Println(v119, v1, v1)
                                                                                                                                }
                                                                                                                                var v166 string = v133
                                                                                                                                v133 = ((("a" + v166) + ((v166 + v166) + v166)) + v166)
                                                                                                                                _ = v166
                                                                                                                            }
                                                                                                                        case -514:
                                                                                                                            v4 = (!v4)
                                                                                                                        case 9223372036854775807:
                                                                                                                            var v167 uint64 = 641
                                                                                                                            _ = v167
                                                                                                                        default:
                                                                                                                            if v4 {
                                                                                                                                v119 += v119
                                                                                                                                var v168 int8 = -60
                                                                                                                                _ = v168
                                                                                                                            } else {
                                                                                                                                if (0.1 >= 0.1) {
                                                                                                                                    var v169 float32 = (-2.25 / v7)
                                                                                                                                    _ = v169
                                                                                                                                } else {
                                                                                                                                    v170 := GenS17{F0: true}
                                                                                                                                    v171 := v3[1]
                                                                                                                                    if v4 {
                                                                                                                                        v4 = (v119 <= (v119 / ((0) | 1)))
                                                                                                                                        v172 := GenS9{F0: v171, F1: (v119 >> ((255) & 63)), F2: 32767}
                                                                                                                                        var v173 int64 = -9223372036854775808
                                                                                                                                        _ = v172
                                                                                                                                        _ = v173
                                                                                                                                    }
                                                                                                                                    _ = v170
                                                                                                                                    _ = v171
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v174 := GenS14{F0: -9223372036854775808}
                                                                                                                        var v175 uint16 = 42
                                                                                                                        _ = v174
                                                                                                                        _ = v175
                                                                                                                    }
                                                                                                                    var v176 bool = (!false)
                                                                                                                    _ = v176
                                                                                                                } else {
                                                                                                                    var v177 int32 = (v1 >> ((42) & 63))
                                                                                                                    fmt.Println(v119, fi57, v177)
                                                                                                                    v4 = ((fi57 << ((586) & 63)) != 127)
                                                                                                                    _ = v177
                                                                                                                }
                                                                                                            default:
                                                                                                                if ((float32(-2.25) - 0.0) == v56) {
                                                                                                                    fmt.Println(fi55, v119, v2)
                                                                                                                    switch (fi57 & (fi55 << ((1) & 63))) {
                                                                                                                    case 42:
                                                                                                                        fmt.Println(v56, v56)
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    if (42 <= 2) {
                                                                                                                        switch "foo bar" {
                                                                                                                        case "\n\t":
                                                                                                                            if false {
                                                                                                                                v119 = v119
                                                                                                                                v2 = (v2 % (((v2 - (v2 + v2))) | 1))
                                                                                                                            }
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        if ((v119 / ((v119) | 1)) > 767) {
                                                                                                                            var v178 uint64 = (v119 << ((4294967295) & 63))
                                                                                                                            var v179 uint32 = 763
                                                                                                                            _ = v178
                                                                                                                            _ = v179
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v153
                                                                                                        } else {
                                                                                                            v7 -= (v7 / 3.14159)
                                                                                                            v4 = ((v119 - (v119 % (((^(v119))) | 1))) == v119)
                                                                                                            var v180 float32 = (v7 - v56)
                                                                                                            _ = v180
                                                                                                        }
                                                                                                        v4 = (v119 == (v119 >> ((2) & 63)))
                                                                                                        v4 = v4
                                                                                                    }
                                                                                                    switch "foo bar" {
                                                                                                    case "":
                                                                                                        v119 = 908
                                                                                                    case "foo bar":
                                                                                                        var v181 uint32 = v2
                                                                                                        _ = v181
                                                                                                    default:
                                                                                                        fmt.Println(v7, v133)
                                                                                                    }
                                                                                                    v182 := []int{2147483647}
                                                                                                    _ = v182
                                                                                                }
                                                                                                _ = v133
                                                                                                _ = v134
                                                                                            }
                                                                                            v183 := GenS7{F0: (!((0.0 / 1.5) <= ((((float64(((float32(1e10) * v7) * v7)) * float64((v56 - (v56 * 1e10)))) / 3.14159) * (0.0 / ((0.0 - -2.25) + 0.1))) * 0.0)))}
                                                                                            _ = v183
                                                                                        } else {
                                                                                            fmt.Println(v119, fi55, v1)
                                                                                            v56 = (1.5 + v7)
                                                                                            fmt.Println(v4, fi57, v7)
                                                                                        }
                                                                                        var v184 int8 = 1
                                                                                        _ = v184
                                                                                    }
                                                                                }
                                                                                fmt.Println(v2)
                                                                                v185 := []int{}
                                                                                _ = v185
                                                                            }
                                                                        case "a":
                                                                            _ = len(v3)
                                                                        }
                                                                        fmt.Println(v1, fi55)
                                                                        if true {
                                                                            var v186 float64 = 0.0
                                                                            _ = v186
                                                                        } else {
                                                                            switch fi57 {
                                                                            case 1:
                                                                                switch (fi57 >> ((42) & 63)) {
                                                                                case 2147483647:
                                                                                    var v187 int16 = -32768
                                                                                    _ = v187
                                                                                case -209:
                                                                                    var v188 int32 = (v1 & v1)
                                                                                    v189 := GenS18{F0: 9223372036854775807, F1: ("foo bar" + "foo bar")}
                                                                                    v190 := GenS19{F0: -84, F1: (42 >= 201)}
                                                                                    v191 := GenS10{F0: 42}
                                                                                    if (v4 && (!(!true))) {
                                                                                        v4 = (fi57 > (^(fi55)))
                                                                                    }
                                                                                    _ = v188
                                                                                    _ = v189
                                                                                    _ = v190
                                                                                    _ = v191
                                                                                }
                                                                            default:
                                                                                var v192 uint64 = (^(v119))
                                                                                _ = v192
                                                                            }
                                                                            if (v7 != (((v7 - -2.25) + (0.1 / (1e10 + 0.0))) + v7)) {
                                                                                v56 = ((1e10 + v7) * float32(1.5))
                                                                                switch (fi55 % (((fi57 ^ fi57)) | 1)) {
                                                                                case 790:
                                                                                    v193 := []int{-541, (fi55 / ((fi57) | 1)), (^(fi57)), (fi57 ^ (fi57 >> ((255) & 63)))}
                                                                                    _ = v193
                                                                                case 1:
                                                                                    v194 := []int{(fi57 ^ fi57), fi55, fi57}
                                                                                    _ = v194
                                                                                case -99:
                                                                                    if (!true) {
                                                                                        switch 42 {
                                                                                        case -645:
                                                                                            switch "a" {
                                                                                            case "":
                                                                                                v1 += (v1 ^ v1)
                                                                                            }
                                                                                            switch (fi57 - fi57) {
                                                                                            case 351:
                                                                                                v195 := []int{(fi55 / ((fi55) | 1)), -121}
                                                                                                _ = v195
                                                                                            case -1:
                                                                                                v196 := []int{(fi57 / (((fi57 >> ((4294967295) & 63))) | 1)), (fi55 % ((127) | 1)), fi57}
                                                                                                _ = v196
                                                                                            case 42:
                                                                                                var v197 string = ("hello" + "foo bar")
                                                                                                _ = v197
                                                                                            default:
                                                                                                if ((v119 % ((v119) | 1)) <= 332) {
                                                                                                    if v4 {
                                                                                                        fmt.Println(fi55, v4, v2)
                                                                                                    } else {
                                                                                                        var v198 int16 = 42
                                                                                                        if v4 {
                                                                                                            v56 = -2.25
                                                                                                        } else {
                                                                                                            fmt.Println(v7)
                                                                                                            var v199 uint32 = v2
                                                                                                            _ = v199
                                                                                                        }
                                                                                                        v56 -= (v56 * 0.1)
                                                                                                        _ = v198
                                                                                                    }
                                                                                                } else {
                                                                                                    v200 := []int{fi55, (-(fi55)), 1, (-(fi57))}
                                                                                                    v201 := GenS1{F0: 42}
                                                                                                    delete(v3, 0)
                                                                                                    _ = v200
                                                                                                    _ = v201
                                                                                                }
                                                                                                v7 += float32((1e10 * (-2.25 * 1.5)))
                                                                                                if v4 {
                                                                                                    var v202 uint64 = (^(v119))
                                                                                                    if (1.5 <= 1e10) {
                                                                                                        var v203 int = (fi57 >> ((945) & 63))
                                                                                                        var v204 int8 = 100
                                                                                                        v203 = 32767
                                                                                                        var v205 uint16 = uint16(v2)
                                                                                                        var v206 uint = 359
                                                                                                        _ = v203
                                                                                                        _ = v204
                                                                                                        _ = v205
                                                                                                        _ = v206
                                                                                                    } else {
                                                                                                        v207 := []int{fi55, (fi55 % (((-(fi57))) | 1))}
                                                                                                        v202 -= 0
                                                                                                        v202 += (v202 >> ((42) & 63))
                                                                                                        _ = v207
                                                                                                    }
                                                                                                    var v208 int8 = int8(fi57)
                                                                                                    if v4 {
                                                                                                        v208 += 100
                                                                                                        var v209 int = -130
                                                                                                        v1 += (v1 | v1)
                                                                                                        if ((v1 + 1) != 560) {
                                                                                                            var v210 uint32 = 2147483647
                                                                                                            v7 -= v7
                                                                                                            fmt.Println(v2)
                                                                                                            _ = v210
                                                                                                        } else {
                                                                                                            var v211 uint = 65535
                                                                                                            if ((("foo bar" + "a") + "foo bar") >= ("foo bar" + ("hello" + (((("\n\t" + "") + ("hello" + "")) + ((((("\n\t" + ("" + (("\n\t" + "foo bar") + ("hello" + "hello")))) + "hello") + ("foo bar" + ("foo bar" + ("foo bar" + (("foo bar" + ("" + ("\n\t" + ("\n\t" + "\n\t")))) + "\n\t"))))) + ("foo bar" + "foo bar")) + "a")) + "")))) {
                                                                                                                var v212 bool = v4
                                                                                                                var v213 uint64 = (v119 * (^(v202)))
                                                                                                                _ = v212
                                                                                                                _ = v213
                                                                                                            } else {
                                                                                                                var v214 float32 = (v56 - float32(1.5))
                                                                                                                if (!(!true)) {
                                                                                                                    if v4 {
                                                                                                                        v208 -= v208
                                                                                                                        if v4 {
                                                                                                                            var v215 string = "\n\t"
                                                                                                                            _ = v215
                                                                                                                        } else {
                                                                                                                            v202 += v202
                                                                                                                            switch "a" {
                                                                                                                            case "\n\t":
                                                                                                                                v216 := []int{(fi55 >> ((2) & 63)), fi57}
                                                                                                                                _ = v216
                                                                                                                            default:
                                                                                                                                fmt.Println(v211, v2)
                                                                                                                            }
                                                                                                                            var v217 string = "\n\t"
                                                                                                                            _ = v217
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v214
                                                                                                            }
                                                                                                            fmt.Println(v56)
                                                                                                            _ = v211
                                                                                                        }
                                                                                                        _ = v209
                                                                                                    } else {
                                                                                                        var v218 uint = 42
                                                                                                        _ = v218
                                                                                                    }
                                                                                                    fmt.Println(v202, fi57)
                                                                                                    _ = v202
                                                                                                    _ = v208
                                                                                                } else {
                                                                                                    if true {
                                                                                                        var v219 bool = v4
                                                                                                        _ = v219
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            v3[(^(v2))] = v4
                                                                                            v220 := GenS20{F0: v2, F1: ("foo bar" + ("" + ("foo bar" + ("" + "a"))))}
                                                                                            if (float64(v7) >= (1e10 * -2.25)) {
                                                                                                delete(v3, 0)
                                                                                                var v221 int16 = -147
                                                                                                _ = v221
                                                                                            }
                                                                                            v222 := []int{(fi55 & fi55), fi57, fi57}
                                                                                            switch "\n\t" {
                                                                                            case "":
                                                                                                if len(v222) >= 2 {
                                                                                                    v222 = v222[1:len(v222)]
                                                                                                }
                                                                                            default:
                                                                                                v7 += float32(float64(float32(3.14159)))
                                                                                            }
                                                                                            _ = v220
                                                                                            _ = v222
                                                                                        case 32767:
                                                                                            v223 := GenS21{F0: 602}
                                                                                            _ = v223
                                                                                        }
                                                                                    } else {
                                                                                        if (!false) {
                                                                                            var v224 int8 = -128
                                                                                            v56 = 0.1
                                                                                            var v225 string = "foo bar"
                                                                                            _ = v224
                                                                                            _ = v225
                                                                                        }
                                                                                        fmt.Println(fi55)
                                                                                    }
                                                                                default:
                                                                                    if (("a" + "foo bar") >= "foo bar") {
                                                                                        switch (^(fi55)) {
                                                                                        case 42:
                                                                                            var v226 uint = 678
                                                                                            _ = v226
                                                                                        case 1:
                                                                                            fmt.Println(fi57, v1, v56)
                                                                                        case 32768:
                                                                                            var v227 int = fi55
                                                                                            fmt.Println(v119, v2, v1)
                                                                                            v56 = v56
                                                                                            v56 += (1e10 / float32(1e10))
                                                                                            fmt.Println(fi57)
                                                                                            if (true || v4) {
                                                                                                v228 := GenS22{F0: 1000, F1: 304}
                                                                                                _ = v228
                                                                                            } else {
                                                                                                var v229 uint64 = 4294967295
                                                                                                v119 = (^(v229))
                                                                                                _ = v229
                                                                                            }
                                                                                            v56 = 0.0
                                                                                            v230 := GenS21{F0: 1}
                                                                                            _ = v227
                                                                                            _ = v230
                                                                                        default:
                                                                                            v231 := []int{(fi57 * fi55)}
                                                                                            _ = v231
                                                                                        }
                                                                                        var v232 float64 = ((3.14159 + float64((3.14159 + -2.25))) - 1.5)
                                                                                        var v233 uint8 = 159
                                                                                        v234 := GenS23{F0: (v1 + (-(v1))), F1: ("\n\t" + "")}
                                                                                        v235 := []int{int(v233), fi55, fi57}
                                                                                        _ = v232
                                                                                        _ = v233
                                                                                        _ = v234
                                                                                        _ = v235
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                v236 := GenS15{F0: "a", F1: false, F2: 4294967295}
                                                                                v7 += v7
                                                                                var v237 int = (fi55 | -385)
                                                                                _ = v236
                                                                                _ = v237
                                                                            }
                                                                        }
                                                                        var v238 float64 = (0.1 * float64(-2.25))
                                                                        _ = v238
                                                                    }
                                                                default:
                                                                    v4 = v4
                                                                }
                                                            case "a":
                                                                var v239 int8 = -1
                                                                _ = v239
                                                            default:
                                                                var v240 uint = 142
                                                                _ = v240
                                                            }
                                                            _ = v119
                                                        default:
                                                            if true {
                                                                if (2 >= 338) {
                                                                    if v4 {
                                                                        var v241 uint16 = 42
                                                                        var v242 int8 = 88
                                                                        var v243 uint = 255
                                                                        _ = v241
                                                                        _ = v242
                                                                        _ = v243
                                                                    } else {
                                                                        fmt.Println(fi57)
                                                                        if v4 {
                                                                            var v244 bool = false
                                                                            v245 := []int{}
                                                                            v244 = (false || (100 < 127))
                                                                            _ = v244
                                                                            _ = v245
                                                                        } else {
                                                                            var v246 string = "\n\t"
                                                                            _ = v246
                                                                        }
                                                                    }
                                                                    if v4 {
                                                                        v247 := GenS24{F0: 42, F1: 2}
                                                                        v2 = 1
                                                                        var v248 int8 = 127
                                                                        _ = v247
                                                                        _ = v248
                                                                    }
                                                                }
                                                                if (!v4) {
                                                                    switch -9223372036854775808 {
                                                                    case -1:
                                                                        var v249 uint32 = 2
                                                                        _ = v249
                                                                    default:
                                                                        var v250 float64 = ((float64(3.14159) * 0.1) * (1.5 * 3.14159))
                                                                        _ = v250
                                                                    }
                                                                    fmt.Println(v4, v4)
                                                                } else {
                                                                    v251 := []int{(^(fi55)), (fi55 - 127)}
                                                                    _ = len(v251)
                                                                    _ = v251
                                                                }
                                                                var v252 int16 = -462
                                                                fmt.Println(v252)
                                                                _ = v252
                                                            }
                                                        }
                                                    }
                                                }
                                                fmt.Println(fi55, v7, v7)
                                            }
                                        } else {
                                            v253 := GenS21{F0: 65535}
                                            v7 = (v7 * (0.0 - -2.25))
                                            _ = v253
                                        }
                                    }
                                    _ = fi57
                                }
                                v1 = 1
                            } else {
                                if v4 {
                                    var v254 float64 = float64(v7)
                                    var v255 uint16 = 547
                                    v256 := GenS25{F0: -9223372036854775808}
                                    _ = v254
                                    _ = v255
                                    _ = v256
                                } else {
                                    if (v2 == v2) {
                                        for fi257 := 0; fi257 < 3; fi257++ {
                                            if false {
                                                switch "a" {
                                                case "":
                                                    v258 := GenS26{F0: 65535, F1: 505, F2: 1}
                                                    var v259 string = ("" + "\n\t")
                                                    if (!(759 <= 32767)) {
                                                        if true {
                                                            fmt.Println(v2)
                                                            v260 := v3[v2]
                                                            if (379 > 1000) {
                                                                delete(v3, v2)
                                                                var v261 int64 = 0
                                                                _ = v261
                                                            }
                                                            _ = v260
                                                        }
                                                        fmt.Println(v1, v4, fi55)
                                                        var v262 uint = 18446744073709551615
                                                        _ = v262
                                                    }
                                                    _ = v258
                                                    _ = v259
                                                case "\n\t":
                                                    var v263 int32 = v1
                                                    _ = v263
                                                case "hello":
                                                    if v4 {
                                                        v2 = v2
                                                        if false {
                                                            v2 -= (^(v2))
                                                        } else {
                                                            v264 := GenS10{F0: 720}
                                                            var v265 float64 = float64(1.5)
                                                            if true {
                                                                v56 -= v56
                                                            }
                                                            _ = v264
                                                            _ = v265
                                                        }
                                                    }
                                                default:
                                                    var v266 uint32 = (v2 >> ((154) & 63))
                                                    _ = v266
                                                }
                                                v267 := []int{fi55, (^(fi55))}
                                                if v4 {
                                                    v1 = (-(v1))
                                                    if (1 < 4294967295) {
                                                        var v268 bool = v4
                                                        _ = v268
                                                    } else {
                                                        if v4 {
                                                            var v269 int = fi257
                                                            v269 -= (fi55 % (((^(fi257))) | 1))
                                                            fmt.Println(v2, fi257, fi55)
                                                            _ = v269
                                                        } else {
                                                            var v270 uint = 983
                                                            if len(v267) > 0 {
                                                                v267[0] = (^(fi257))
                                                            }
                                                            _ = v270
                                                        }
                                                        var v271 int32 = (v1 - v1)
                                                        _ = len(v3)
                                                        _ = v271
                                                    }
                                                    v7 = v56
                                                    v272 := GenS18{F0: (fi257 % (((fi55 & 128)) | 1)), F1: "hello"}
                                                    _ = v272
                                                }
                                                _ = v267
                                            }
                                            v56 = -2.25
                                            v2 += (v2 / (((^(v2))) | 1))
                                            _ = fi257
                                        }
                                        var v273 int64 = 509
                                        _ = v273
                                    } else {
                                        v7 = v56
                                    }
                                    var v274 string = ("foo bar" + "foo bar")
                                    for fi275 := 0; fi275 < 9; fi275++ {
                                        v56 = (v56 * 3.14159)
                                        _ = fi275
                                    }
                                    _ = v274
                                }
                            }
                            if v4 {
                                v276 := GenS17{F0: v4}
                                v277 := []int{}
                                _ = v276
                                _ = v277
                            } else {
                                v278 := GenS27{F0: 1, F1: 255}
                                var v279 uint16 = 58
                                _ = v278
                                _ = v279
                            }
                            _ = fi55
                            _ = v56
                        }
                        fmt.Println(v4, v7, v1)
                    } else {
                        v3[(v2 + (v2 >> ((734) & 63)))] = (255 > 255)
                        if v4 {
                            var v280 int16 = 974
                            _ = v280
                        } else {
                            if true {
                                v4 = v4
                                for fi281 := 0; fi281 < 20; fi281++ {
                                    if (657 >= 306) {
                                        fmt.Println(v7, fi281)
                                    } else {
                                        v7 -= (float32(1.5) + v7)
                                        if v4 {
                                            var v282 int32 = (v1 >> ((990) & 63))
                                            _ = v282
                                        }
                                    }
                                    _ = fi281
                                }
                                if v4 {
                                    for fi283 := 0; fi283 < 6; fi283++ {
                                        v284 := GenS8{F0: 0}
                                        if (!(v2 >= v2)) {
                                            var v285 uint32 = v2
                                            fmt.Println(v4, v1)
                                            _ = v285
                                        }
                                        fmt.Println(fi283)
                                        _ = fi283
                                        _ = v284
                                    }
                                    var v286 int16 = 42
                                    _ = v286
                                }
                            }
                            v287 := []int{-761, -390}
                            _ = v287
                        }
                    }
                    v288 := []int{}
                    _ = v288
                }
            } else {
                var v289 int8 = -128
                switch ("a" + ("" + "")) {
                case "a":
                    v290 := v3[v2]
                    _ = v290
                case "":
                    _ = len(v3)
                case "foo bar":
                    v289 = -128
                default:
                    var v291 uint32 = v2
                    _ = v291
                }
                _ = v289
            }
        } else {
            var v292 int32 = v1
            fmt.Println(v292, v292)
            v293 := []int{717, 42, 4294967295}
            _ = v292
            _ = v293
        }
        v7 = v7
    }
    v294 := GenS13{F0: 9223372036854775807, F1: 42, F2: "a"}
    var v295 float32 = v7
    switch -597 {
    case -58:
        var v296 float32 = float32(3.14159)
        _ = v296
    default:
        switch 2147483647 {
        case 352:
            for fi297 := 0; fi297 < 4; fi297++ {
                v298 := GenS23{F0: v1, F1: "a"}
                _ = fi297
                _ = v298
            }
        case -9223372036854775808:
            var v299 float32 = v295
            _ = v299
        case 127:
            for fi300 := 0; fi300 < 18; fi300++ {
                v2 -= (v2 - 65535)
                delete(v3, (v2 % (((v2 % ((v2) | 1))) | 1)))
                switch (("hello" + "foo bar") + "\n\t") {
                case "a":
                    switch "\n\t" {
                    case "foo bar":
                        var v301 int8 = -1
                        _ = v301
                    case "\n\t":
                        v3[v2] = v4
                    }
                case "foo bar":
                    v1 = (v1 >> ((513) & 63))
                case "\n\t":
                    if (0 >= 784) {
                        v1 = (v1 * v1)
                        v302 := GenS28{F0: v1}
                        switch (fi300 - fi300) {
                        case 2147483647:
                            if v4 {
                                for fi303 := 0; fi303 < 15; fi303++ {
                                    if (2 == v2) {
                                        var v304 bool = (1 != -9223372036854775808)
                                        _ = v304
                                    }
                                    _ = fi303
                                }
                                v1 = (v1 - (v1 ^ (v1 % ((-2147483648) | 1))))
                            }
                        case 128:
                            if (!true) {
                                if (123 <= 255) {
                                    v305 := GenS29{F0: 0, F1: -73, F2: fi300}
                                    v4 = (9223372036854775807 == fi300)
                                    var v306 uint32 = v2
                                    _ = v305
                                    _ = v306
                                }
                                fmt.Println(fi300, v1, v2)
                            } else {
                                v1 -= v1
                                for fi307 := 0; fi307 < 5; fi307++ {
                                    if v4 {
                                        var v308 int8 = 0
                                        _ = v308
                                    } else {
                                        if (true || (false || true)) {
                                            v309 := GenS30{F0: "a", F1: 1}
                                            if (32767 <= 13) {
                                                var v310 bool = (!v4)
                                                v311 := GenS4{F0: 151, F1: 1, F2: 1}
                                                _ = v310
                                                _ = v311
                                            }
                                            _ = v309
                                        } else {
                                            var v312 bool = v4
                                            var v313 int64 = 1
                                            _ = v312
                                            _ = v313
                                        }
                                    }
                                    if (!(0.1 <= float64(float32((-2.25 * 0.0))))) {
                                        v295 -= (v295 + (v7 - v295))
                                        var v314 uint64 = 18446744073709551615
                                        if v4 {
                                            v295 = 1.5
                                            v7 = v295
                                            switch "\n\t" {
                                            case "a":
                                                _ = len(v3)
                                            }
                                        }
                                        _ = v314
                                    }
                                    var v315 int16 = 1000
                                    delete(v3, (v2 / (((^(v2))) | 1)))
                                    _ = fi307
                                    _ = v315
                                }
                            }
                        case -269:
                            switch int(v1) {
                            case 127:
                                v316 := GenS31{F0: v2, F1: false, F2: v4}
                                _ = v316
                            }
                        default:
                            fmt.Println(v2)
                        }
                        _ = v302
                    } else {
                        var v317 float64 = float64(1.5)
                        _ = v317
                    }
                default:
                    v318 := v3[420]
                    _ = v318
                }
                _ = fi300
            }
        }
    }
    if v4 {
        v319 := v3[(v2 * v2)]
        v320 := GenS31{F0: v2, F1: false, F2: true}
        _ = v319
        _ = v320
    } else {
        if (((178 > v1) && v4) && (-171 > -630)) {
            switch "foo bar" {
            case "\n\t":
                var v321 int64 = 9223372036854775807
                _ = v321
            case "hello":
                v7 = ((float32(-2.25) - v7) - 0.1)
            }
            v1 = (v1 * 1)
        }
        var v322 int8 = 1
        _ = v322
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v295:", v295)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v294
    _ = v295
    _ = v3
    _ = v4
    _ = v7
}
