package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int16; F1 bool; F2 int32 }
type GenS3 struct { F0 uint16; F1 int16; F2 string }
type GenS4 struct { F0 int8; F1 string; F2 int8 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS8 struct { F0 uint; F1 uint32; F2 int32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int32; F1 uint64 }
type GenS11 struct { F0 int; F1 int16; F2 uint16 }
type GenS12 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS13 struct { F0 int; F1 int32 }
type GenS14 struct { F0 int8; F1 int32; F2 bool }
type GenS15 struct { F0 int16; F1 uint64 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint; F1 uint16 }
type GenS18 struct { F0 int64; F1 int8; F2 uint64 }
func genFn0(p0 bool) int16 {
    switch 128 {
    case 311:
        var v4 uint64 = 18446744073709551615
        _ = v4
    case 948:
        p0 = true
    }
    var v5 uint64 = 981
    v6 := GenS1{F0: -128, F1: 792}
    _ = p0
    _ = v5
    _ = v6
    return 934
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    var v7 int32 = 730
    for fi8 := 0; fi8 < 18; fi8++ {
        v1 = -2.25
        fmt.Println(v1, fi8)
        _ = fi8
    }
    switch 544 {
    case 9223372036854775807:
        v1 -= v1
    case 127:
        if (false || false) {
            fmt.Println(v1)
            var v9 float32 = v1
            _ = v9
        } else {
            var v10 bool = false
            switch 254 {
            case 32767:
                for fi11 := 0; fi11 < 11; fi11++ {
                    v12 := GenS1{F0: -32, F1: 32767}
                    for fi13 := 0; fi13 < 13; fi13++ {
                        v14 := GenS1{F0: 127, F1: 0}
                        v1 = v1
                        fmt.Println(fi11)
                        _ = fi13
                        _ = v14
                    }
                    _ = fi11
                    _ = v12
                }
            case -1:
                switch "" {
                case "\n\t":
                    var v15 int16 = -810
                    _ = v15
                case "foo bar":
                    v16 := []int{-34}
                    _ = v16
                case "hello":
                    v17 := GenS1{F0: 100, F1: 65535}
                    _ = v17
                }
            default:
                var v18 uint64 = 852
                _ = v18
            }
            for fi19 := 0; fi19 < 4; fi19++ {
                if (-1 > -495) {
                    for fi20 := 0; fi20 < 1; fi20++ {
                        fmt.Println(fi19)
                        _ = fi20
                    }
                    fmt.Println(v7, v1, v7)
                    v21 := []int{fi19, (^(fi19)), 128}
                    _ = v21
                }
                v10 = (0.0 != 0.1)
                var v22 int8 = 0
                _ = fi19
                _ = v22
            }
            _ = v10
        }
    default:
        if false {
            var v23 int32 = 1
            _ = v23
        } else {
            var v24 bool = ("a" < "hello")
            if (406 >= 1) {
                fmt.Println(v1, v24)
            }
            _ = v24
        }
    }
    fmt.Println(v7)
    for fi25 := 0; fi25 < 12; fi25++ {
        v26 := []int{fi25, (fi25 ^ fi25)}
        _ = fi25
        _ = v26
    }
    v1 = ((((v1 * v1) * v1) + 1e10) + (1e10 / v1))
    if (true && (127 > 255)) {
        if ((4294967295 == 1) && (!false)) {
            v7 += 2147483647
            if (2147483648 >= 34) {
                _ = len(v3)
                v1 = v1
                if false {
                    fmt.Println(v7)
                    v27 := GenS2{F0: genFn0(false), F1: true, F2: v7}
                    var v28 int8 = 127
                    _ = v27
                    _ = v28
                } else {
                    fmt.Println(v1, v1)
                }
            }
        } else {
            v29 := GenS3{F0: 255, F1: 157, F2: "\n\t"}
            switch -173 {
            case -9223372036854775808:
                for fi30 := 0; fi30 < 1; fi30++ {
                    for fi31 := 0; fi31 < 17; fi31++ {
                        switch (fi31 << ((42) & 63)) {
                        case -800:
                            var v32 bool = true
                            _ = v32
                        case 32768:
                            switch ("hello" + ("foo bar" + "\n\t")) {
                            case "hello":
                                v33 := GenS4{F0: -25, F1: (("" + "") + "foo bar"), F2: 98}
                                _ = v33
                            case "":
                                v34 := GenS2{F0: 457, F1: true, F2: v7}
                                _ = v34
                            }
                        default:
                            fmt.Println(fi30, v7, v1)
                        }
                        _ = fi31
                    }
                    if true {
                        fmt.Println(fi30, v1, fi30)
                        for fi35 := 0; fi35 < 11; fi35++ {
                            v7 -= v7
                            _ = fi35
                        }
                    } else {
                        if true {
                            var v36 uint64 = 18446744073709551615
                            _ = v36
                        }
                        v37 := []int{}
                        _ = v37
                    }
                    v7 = (v7 >> ((18446744073709551615) & 63))
                    _ = fi30
                }
            default:
                v1 += v1
            }
            if true {
                fmt.Println(v7, v1)
                for fi38 := 0; fi38 < 5; fi38++ {
                    var v39 uint64 = 0
                    for fi40 := 0; fi40 < 16; fi40++ {
                        var v41 int64 = -462
                        _ = fi40
                        _ = v41
                    }
                    v7 = (v7 >> ((4294967295) & 63))
                    _ = fi38
                    _ = v39
                }
                switch "foo bar" {
                case "hello":
                    delete(v2, 42)
                case "":
                    if false {
                        v7 = (v7 << ((4294967295) & 63))
                    } else {
                        if (!true) {
                            var v42 uint = 0
                            _ = v42
                        } else {
                            v43 := []int{42, -263, -347}
                            v1 -= 0.1
                            _ = v43
                        }
                        var v44 int32 = v7
                        for fi45 := 0; fi45 < 6; fi45++ {
                            if (!(!false)) {
                                v46 := GenS3{F0: 42, F1: 32767, F2: "\n\t"}
                                _ = v46
                            } else {
                                v7 += (v44 >> ((144) & 63))
                                for fi47 := 0; fi47 < 15; fi47++ {
                                    switch "" {
                                    case "a":
                                        if (-1 == -32768) {
                                            v3[2147483647] = 32768
                                        } else {
                                            v44 = (v7 + v7)
                                            var v48 int8 = 127
                                            delete(v3, (fi45 + fi47))
                                            v44 = -2147483648
                                            _ = v48
                                        }
                                    case "\n\t":
                                        var v49 string = ("" + "hello")
                                        _ = v49
                                    }
                                    fmt.Println(v44)
                                    var v50 int = 32767
                                    _ = fi47
                                    _ = v50
                                }
                                if false {
                                    if false {
                                        for fi51 := 0; fi51 < 10; fi51++ {
                                            if (23 == -32768) {
                                                var v52 int16 = -102
                                                _ = v52
                                            } else {
                                                if (true && (0 <= 910)) {
                                                    switch "\n\t" {
                                                    case "":
                                                        var v53 uint8 = 1
                                                        _ = v53
                                                    case "foo bar":
                                                        fmt.Println(v44)
                                                    default:
                                                        var v54 bool = (!(true && true))
                                                        _ = v54
                                                    }
                                                    var v55 int64 = -884
                                                    fmt.Println(fi51)
                                                    _ = v55
                                                } else {
                                                    var v56 int8 = -8
                                                    var v57 float32 = 1.5
                                                    _ = v56
                                                    _ = v57
                                                }
                                            }
                                            _ = fi51
                                        }
                                    }
                                    var v58 uint32 = 4294967295
                                    switch (("" + "hello") + "") {
                                    case "hello":
                                        var v59 uint8 = 237
                                        _ = v59
                                    default:
                                        if (127 < 127) {
                                            v60 := GenS2{F0: 1000, F1: true, F2: (v7 % ((int32(fi45)) | 1))}
                                            if false {
                                                for fi61 := 0; fi61 < 11; fi61++ {
                                                    v44 = -453
                                                    if false {
                                                        if (((0.0 * v1) == (float32(3.14159) + 1e10)) || true) {
                                                            delete(v2, v58)
                                                            switch fi61 {
                                                            case 32767:
                                                                var v62 int32 = 68
                                                                v63 := []int{(fi45 << ((190) & 63))}
                                                                if (false || (2 >= 42)) {
                                                                    _ = len(v63)
                                                                }
                                                                _ = v62
                                                                _ = v63
                                                            default:
                                                                v64 := GenS5{F0: (v7 / ((v7) | 1)), F1: (v7 - (v7 & v44))}
                                                                _ = v64
                                                            }
                                                        } else {
                                                            fmt.Println(v1, v58)
                                                            var v65 string = ""
                                                            if (true || (false || (!((!(fi45 > (fi61 * fi61))) && true)))) {
                                                                if (v58 < v58) {
                                                                    v66 := GenS4{F0: -128, F1: v65, F2: int8(v44)}
                                                                    if false {
                                                                        var v67 bool = true
                                                                        v1 = v1
                                                                        switch -136 {
                                                                        case 32768:
                                                                            if true {
                                                                                v44 = v44
                                                                                var v68 int = fi61
                                                                                v69 := []int{}
                                                                                _ = v68
                                                                                _ = v69
                                                                            } else {
                                                                                fmt.Println(fi61, v1, v7)
                                                                                v67 = v67
                                                                                if v67 {
                                                                                    var v70 int = (fi45 << ((4294967295) & 63))
                                                                                    var v71 float64 = (1e10 * 1e10)
                                                                                    var v72 int8 = -13
                                                                                    _ = v70
                                                                                    _ = v71
                                                                                    _ = v72
                                                                                }
                                                                            }
                                                                        case 0:
                                                                            v73 := GenS2{F0: 42, F1: (!false), F2: v44}
                                                                            _ = v73
                                                                        default:
                                                                            v74 := GenS6{F0: 876, F1: ("" + v65)}
                                                                            _ = v74
                                                                        }
                                                                        var v75 float64 = 1.5
                                                                        _ = v67
                                                                        _ = v75
                                                                    }
                                                                    _ = v66
                                                                }
                                                            } else {
                                                                var v76 uint32 = v58
                                                                fmt.Println(v7)
                                                                if false {
                                                                    if (1 != 446) {
                                                                        if (false && (928 < 4294967295)) {
                                                                            var v77 float64 = (1e10 / 1.5)
                                                                            if true {
                                                                                fmt.Println(v65)
                                                                                if (-1 != (v44 << ((18446744073709551615) & 63))) {
                                                                                    _ = len(v2)
                                                                                    v76 = 65535
                                                                                    v78 := []int{}
                                                                                    _ = v78
                                                                                }
                                                                                v79 := GenS5{F0: v44, F1: 2147483647}
                                                                                _ = v79
                                                                            } else {
                                                                                v58 += v76
                                                                            }
                                                                            if ((42 <= genFn0(((v58 >> ((uint(v58)) & 63)) < 42))) && false) {
                                                                                v7 = v7
                                                                                if (1 >= 0) {
                                                                                    v80 := []int{-9223372036854775808, (fi61 | fi61), (fi61 << ((0) & 63)), fi45}
                                                                                    _ = v80
                                                                                }
                                                                                fmt.Println(fi61, v44, v7)
                                                                                fmt.Println(fi61, v77, v58)
                                                                                switch ((v65 + v65) + (v65 + (v65 + (v65 + (v65 + v65))))) {
                                                                                case "\n\t":
                                                                                    v81 := GenS7{F0: 81, F1: 320, F2: v76}
                                                                                    _ = v81
                                                                                default:
                                                                                    var v82 int = 900
                                                                                    _ = v82
                                                                                }
                                                                                v7 -= (v7 << ((658) & 63))
                                                                            } else {
                                                                                switch -9223372036854775808 {
                                                                                case 32767:
                                                                                    if false {
                                                                                        var v83 float32 = (0.0 / v1)
                                                                                        _ = v83
                                                                                    }
                                                                                default:
                                                                                    fmt.Println(v65, fi61, v7)
                                                                                }
                                                                                v77 = 1e10
                                                                                v1 -= 1e10
                                                                            }
                                                                            v77 -= 0.0
                                                                            _ = v77
                                                                        } else {
                                                                            fmt.Println(v65, v7)
                                                                        }
                                                                        v58 += (^(v58))
                                                                    }
                                                                }
                                                                _ = v76
                                                            }
                                                            v2[v58] = fi61
                                                            var v84 uint32 = (v58 % ((v58) | 1))
                                                            _ = v65
                                                            _ = v84
                                                        }
                                                        if true {
                                                            var v85 int = 1
                                                            if true {
                                                                if (!(42 != -132)) {
                                                                    if false {
                                                                        if (-1 < 0) {
                                                                            var v86 string = ""
                                                                            v87 := GenS8{F0: 197, F1: v58, F2: -650}
                                                                            var v88 int8 = -72
                                                                            _ = v86
                                                                            _ = v87
                                                                            _ = v88
                                                                        }
                                                                        switch (fi61 | v85) {
                                                                        case -906:
                                                                            switch "hello" {
                                                                            case "\n\t":
                                                                                v89 := GenS9{F0: 2}
                                                                                _ = v89
                                                                            case "":
                                                                                switch ((("\n\t" + "\n\t") + ("a" + "\n\t")) + ((("foo bar" + "\n\t") + "\n\t") + (("a" + "\n\t") + "\n\t"))) {
                                                                                case "":
                                                                                    v90 := []int{32767, -327, -2147483648, (^(fi61))}
                                                                                    _ = v90
                                                                                default:
                                                                                    var v91 int16 = -121
                                                                                    _ = v91
                                                                                }
                                                                            case "a":
                                                                                v92 := GenS10{F0: -2147483648, F1: 4294967295}
                                                                                _ = v92
                                                                            }
                                                                        case -1:
                                                                            v93 := v3[32768]
                                                                            _ = v93
                                                                        default:
                                                                            v7 = -1
                                                                        }
                                                                        var v94 int = 42
                                                                        if (682 < 400) {
                                                                            v95 := GenS11{F0: (fi61 & fi61), F1: -1, F2: 65535}
                                                                            v96 := GenS3{F0: 1, F1: int16(v44), F2: "foo bar"}
                                                                            _ = v95
                                                                            _ = v96
                                                                        }
                                                                        _ = v94
                                                                    }
                                                                } else {
                                                                    v58 = 1
                                                                    if true {
                                                                        v97 := []int{}
                                                                        if len(v97) > 0 {
                                                                            v97[0] = fi45
                                                                        }
                                                                        if len(v97) >= 2 {
                                                                            v97 = v97[1:len(v97)]
                                                                        }
                                                                        switch fi61 {
                                                                        case -9223372036854775808:
                                                                            var v98 int = fi61
                                                                            _ = v98
                                                                        case -1:
                                                                            var v99 int32 = (v44 + v7)
                                                                            _ = v99
                                                                        case 4294967295:
                                                                            if len(v97) >= 2 {
                                                                                v97 = v97[1:len(v97)]
                                                                            }
                                                                        }
                                                                        v44 -= (v7 * int32(v58))
                                                                        v85 -= (fi45 ^ fi45)
                                                                        _ = v97
                                                                    } else {
                                                                        v44 = v44
                                                                        var v100 uint = 255
                                                                        switch "" {
                                                                        case "hello":
                                                                            var v101 int64 = 0
                                                                            _ = v101
                                                                        case "a":
                                                                            if (!((v44 - (v44 << ((18446744073709551615) & 63))) == (v44 - 2147483647))) {
                                                                                var v102 string = ""
                                                                                v7 += -2147483648
                                                                                var v103 string = ""
                                                                                _ = v102
                                                                                _ = v103
                                                                            } else {
                                                                                v100 -= 42
                                                                                var v104 int16 = genFn0((!false))
                                                                                v105 := GenS6{F0: v100, F1: "hello"}
                                                                                var v106 float32 = (v1 * ((1.5 / 0.1) / -2.25))
                                                                                _ = v104
                                                                                _ = v105
                                                                                _ = v106
                                                                            }
                                                                        case "foo bar":
                                                                            switch (fi61 % ((4294967295) | 1)) {
                                                                            case -2147483648:
                                                                                v58 += v58
                                                                            default:
                                                                                switch fi45 {
                                                                                case 32767:
                                                                                    v85 -= (v85 ^ v85)
                                                                                case -1:
                                                                                    var v107 float64 = -2.25
                                                                                    _ = v107
                                                                                case 127:
                                                                                    if true {
                                                                                        v1 = (-2.25 / (v1 * (1.5 * 0.0)))
                                                                                        v108 := []int{fi45, fi61}
                                                                                        var v109 uint8 = 127
                                                                                        v7 += v44
                                                                                        _ = v108
                                                                                        _ = v109
                                                                                    } else {
                                                                                        var v110 uint8 = 255
                                                                                        var v111 uint64 = 1
                                                                                        fmt.Println(v44, v1, v58)
                                                                                        _ = v110
                                                                                        _ = v111
                                                                                    }
                                                                                default:
                                                                                    var v112 string = ""
                                                                                    _ = v112
                                                                                }
                                                                            }
                                                                        default:
                                                                            v1 = (((((v1 - v1) - v1) + 3.14159) + 1.5) / (v1 - (v1 + 1e10)))
                                                                        }
                                                                        _ = v100
                                                                    }
                                                                    var v113 int32 = v44
                                                                    _ = v113
                                                                }
                                                            }
                                                            v114 := GenS1{F0: 0, F1: 32767}
                                                            switch (("" + "") + (("a" + "foo bar") + "a")) {
                                                            case "":
                                                                v115 := []int{2147483647}
                                                                _ = v115
                                                            case "foo bar":
                                                                fmt.Println(v44, v85, v1)
                                                            case "a":
                                                                var v116 uint64 = 0
                                                                _ = v116
                                                            default:
                                                                v117 := GenS11{F0: -2147483648, F1: -1, F2: 665}
                                                                _ = v117
                                                            }
                                                            _ = v114
                                                            _ = v85
                                                        }
                                                    }
                                                    v118 := []int{(fi45 | 2147483647), -53, (fi61 << ((851) & 63))}
                                                    _ = fi61
                                                    _ = v118
                                                }
                                                if (!true) {
                                                    if (false || (32767 >= 704)) {
                                                        for fi119 := 0; fi119 < 11; fi119++ {
                                                            v44 = 2147483647
                                                            var v120 float32 = -2.25
                                                            _ = fi119
                                                            _ = v120
                                                        }
                                                        var v121 float64 = (-2.25 - 0.0)
                                                        _ = v121
                                                    }
                                                    switch 42 {
                                                    case 9223372036854775807:
                                                        v44 += v44
                                                    case 0:
                                                        if ((v58 << ((0) & 63)) >= v58) {
                                                            var v122 int64 = 927
                                                            for fi123 := 0; fi123 < 7; fi123++ {
                                                                v124 := []int{1, fi123}
                                                                fmt.Println(fi45)
                                                                _ = fi123
                                                                _ = v124
                                                            }
                                                            _ = v122
                                                        } else {
                                                            var v125 int16 = 1
                                                            for fi126 := 0; fi126 < 16; fi126++ {
                                                                var v127 uint16 = 65535
                                                                _ = fi126
                                                                _ = v127
                                                            }
                                                            v128 := []int{(fi45 << ((0) & 63)), 4294967295, fi45}
                                                            _ = v125
                                                            _ = v128
                                                        }
                                                    case 32768:
                                                        v129 := []int{fi45, fi45, 42}
                                                        _ = v129
                                                    default:
                                                        if (v1 != (0.0 - 0.1)) {
                                                            fmt.Println(v1, fi45)
                                                            v7 -= (v44 * (v44 >> ((0) & 63)))
                                                            v58 = (v58 >> ((4294967295) & 63))
                                                        }
                                                    }
                                                } else {
                                                    v130 := GenS12{F0: (^(v58)), F1: 0, F2: 1}
                                                    _ = v130
                                                }
                                            } else {
                                                delete(v3, (fi45 & (fi45 << ((430) & 63))))
                                                var v131 uint64 = 571
                                                if ((v131 % ((429) | 1)) <= (v131 ^ (v131 - v131))) {
                                                    v1 -= float32(3.14159)
                                                    v132 := []int{}
                                                    switch (("foo bar" + (("hello" + ("hello" + ("a" + "\n\t"))) + "")) + "hello") {
                                                    case "hello":
                                                        v44 = 1
                                                    }
                                                    _ = v132
                                                } else {
                                                    for fi133 := 0; fi133 < 19; fi133++ {
                                                        v131 -= 5
                                                        _ = fi133
                                                    }
                                                    if false {
                                                        var v134 uint16 = 255
                                                        v58 += v58
                                                        _ = v134
                                                    }
                                                }
                                                _ = v131
                                            }
                                            _ = v60
                                        }
                                    }
                                    _ = v58
                                } else {
                                    var v135 float64 = (-2.25 / -2.25)
                                    switch "a" {
                                    case "foo bar":
                                        v136 := []int{fi45, (fi45 - (fi45 * (fi45 + fi45)))}
                                        _ = v136
                                    case "a":
                                        fmt.Println(v135, v1)
                                    default:
                                        var v137 int = (fi45 % ((fi45) | 1))
                                        _ = v137
                                    }
                                    _ = v135
                                }
                            }
                            var v138 uint64 = 356
                            v7 = v44
                            _ = fi45
                            _ = v138
                        }
                        _ = v44
                    }
                }
            } else {
                for fi139 := 0; fi139 < 13; fi139++ {
                    if false {
                        v140 := []int{-476, 2147483647}
                        if true {
                            for fi141 := 0; fi141 < 10; fi141++ {
                                var v142 int8 = -1
                                if true {
                                    v1 = v1
                                    v1 += 0.0
                                } else {
                                    var v143 int32 = (v7 << ((438) & 63))
                                    v1 += float32(-2.25)
                                    fmt.Println(v1, v143)
                                    v144 := GenS8{F0: 2, F1: 4294967295, F2: 106}
                                    v142 = v142
                                    _ = v143
                                    _ = v144
                                }
                                _ = fi141
                                _ = v142
                            }
                            switch "\n\t" {
                            case "a":
                                for fi145 := 0; fi145 < 17; fi145++ {
                                    delete(v2, 65535)
                                    _ = fi145
                                }
                            case "foo bar":
                                fmt.Println(v1, fi139, v7)
                            default:
                                for fi146 := 0; fi146 < 6; fi146++ {
                                    v147 := v3[(fi139 / (((fi146 * fi139)) | 1))]
                                    var v148 uint8 = 128
                                    switch ("\n\t" + "") {
                                    case "\n\t":
                                        v149 := v2[1]
                                        _ = v149
                                    }
                                    _ = fi146
                                    _ = v147
                                    _ = v148
                                }
                            }
                        } else {
                            var v150 string = "a"
                            var v151 float64 = (-2.25 / 1.5)
                            v7 += (v7 & (v7 >> ((65535) & 63)))
                            _ = v150
                            _ = v151
                        }
                        _ = v140
                    } else {
                        for fi152 := 0; fi152 < 4; fi152++ {
                            var v153 uint16 = 1
                            _ = fi152
                            _ = v153
                        }
                        v154 := GenS12{F0: 998, F1: 423, F2: 32767}
                        _ = v154
                    }
                    delete(v2, 755)
                    _ = fi139
                }
            }
            _ = v29
        }
        v155 := GenS13{F0: -704, F1: (v7 << ((42) & 63))}
        v7 = 1
        _ = v155
    } else {
        var v156 uint32 = 4294967295
        v1 += ((v1 / v1) / v1)
        _ = v156
    }
    v157 := GenS14{F0: -6, F1: (v7 << ((42) & 63)), F2: ("hello" < (("foo bar" + "foo bar") + "foo bar"))}
    v158 := GenS6{F0: 4294967295, F1: ("\n\t" + "foo bar")}
    if false {
        if true {
            v7 += 42
            v159 := GenS15{F0: 42, F1: 0}
            _ = v159
        }
        var v160 int = 0
        if (false && ((("hello" + ("\n\t" + "a")) + ("foo bar" + "\n\t")) == "a")) {
            v161 := GenS3{F0: 32767, F1: 0, F2: ("foo bar" + ((("a" + ("" + ("foo bar" + "a"))) + ("foo bar" + "foo bar")) + "a"))}
            fmt.Println(v1)
            _ = v161
        }
        _ = v160
    } else {
        for fi162 := 0; fi162 < 16; fi162++ {
            if (true || true) {
                v1 -= v1
                v7 = v7
            }
            if false {
                var v163 int16 = -666
                var v164 int32 = (^(v7))
                _ = v163
                _ = v164
            } else {
                var v165 uint16 = 0
                var v166 bool = false
                if (!(!false)) {
                    if (!v166) {
                        v167 := []int{fi162, (fi162 * fi162)}
                        for fi168 := 0; fi168 < 9; fi168++ {
                            var v169 int = -2147483648
                            var v170 int = fi162
                            if true {
                                switch "foo bar" {
                                case "\n\t":
                                    v171 := GenS16{F0: (v170 - fi162)}
                                    _ = v171
                                case "a":
                                    if len(v167) > 0 {
                                        v167[0] = v169
                                    }
                                default:
                                    v166 = (18446744073709551615 <= 2)
                                }
                                fmt.Println(v1, v166)
                                switch "a" {
                                case "foo bar":
                                    v172 := GenS7{F0: 0, F1: 4294967295, F2: 678}
                                    _ = v172
                                case "\n\t":
                                    if len(v167) >= 2 {
                                        v167 = v167[1:len(v167)]
                                    }
                                case "hello":
                                    var v173 string = ""
                                    _ = v173
                                default:
                                    var v174 uint64 = 42
                                    _ = v174
                                }
                            }
                            switch -791 {
                            case -2147483648:
                                v175 := GenS2{F0: 1000, F1: v166, F2: 161}
                                _ = v175
                            default:
                                if true {
                                    v170 = fi168
                                    if (15 == 6) {
                                        if (18446744073709551615 == 0) {
                                            var v176 uint = 4294967295
                                            _ = v176
                                        }
                                        if len(v167) >= 2 {
                                            v167 = v167[1:len(v167)]
                                        }
                                    } else {
                                        if v166 {
                                            v170 += fi168
                                        }
                                    }
                                    fmt.Println(v169, v7)
                                } else {
                                    if true {
                                        _ = len(v3)
                                    } else {
                                        v177 := GenS17{F0: 391, F1: v165}
                                        var v178 int16 = -32768
                                        var v179 string = "a"
                                        v179 = ((((v179 + v179) + ((v179 + "") + ("" + ("a" + "")))) + "hello") + ((("\n\t" + v179) + v179) + (v179 + ((((((v179 + v179) + "") + (v179 + v179)) + "foo bar") + (v179 + ("\n\t" + v179))) + (v179 + v179)))))
                                        v167 = append(v167, 32767)
                                        _ = v177
                                        _ = v178
                                        _ = v179
                                    }
                                    var v180 uint16 = (v165 - (v165 - v165))
                                    _ = v180
                                }
                            }
                            _ = fi168
                            _ = v169
                            _ = v170
                        }
                        _ = v167
                    }
                    if v166 {
                        var v181 int16 = 32767
                        fmt.Println(v165, v181, v165)
                        _ = v181
                    } else {
                        var v182 uint64 = 0
                        var v183 int64 = 42
                        _ = v182
                        _ = v183
                    }
                } else {
                    v7 -= 913
                    var v184 int8 = 127
                    _ = v184
                }
                _ = v165
                _ = v166
            }
            switch (("" + "hello") + ("foo bar" + ("a" + ("\n\t" + ("foo bar" + "a"))))) {
            case "":
                v1 -= (v1 * float32(0.1))
            case "\n\t":
                v185 := []int{}
                _ = v185
            }
            _ = fi162
        }
    }
    if false {
        if (0 >= 42) {
            var v186 float32 = float32(float64(v1))
            _ = v186
        } else {
            if true {
                v187 := GenS18{F0: 9223372036854775807, F1: 42, F2: 2}
                var v188 int8 = 100
                _ = v187
                _ = v188
            } else {
                v1 = float32(float64(v1))
                v7 -= v7
            }
        }
    }
    var v189 uint16 = 255
    if false {
        if true {
            if (924 == 2) {
                v189 += (v189 << ((4294967295) & 63))
                _ = len(v2)
            }
        } else {
            fmt.Println(v7)
            if false {
                var v190 int64 = -10
                v191 := GenS18{F0: 9223372036854775807, F1: 0, F2: 0}
                _ = v190
                _ = v191
            } else {
                v192 := []int{4294967295}
                _ = v192
            }
        }
    } else {
        _ = len(v3)
    }
    for fi193 := 0; fi193 < 14; fi193++ {
        _ = len(v2)
        switch 171 {
        case 127:
            var v194 int8 = -128
            _ = v194
        case -843:
            switch "a" {
            case "foo bar":
                var v195 uint32 = 65535
                _ = v195
            case "hello":
                var v196 uint64 = 2
                _ = v196
            default:
                fmt.Println(fi193, v1, v189)
            }
        case -57:
            v189 -= (^(v189))
        default:
            v189 = 0
        }
        _ = fi193
    }
    v1 += 0.0
    v7 -= 2147483647
    var v197 uint16 = (^(v189))
    for fi198 := 0; fi198 < 7; fi198++ {
        for fi199 := 0; fi199 < 19; fi199++ {
            v200 := []int{9223372036854775807, (-(fi199)), (fi199 >> ((255) & 63))}
            fmt.Println(v189, v197, v197)
            switch ("\n\t" + "a") {
            case "hello":
                switch fi199 {
                case 898:
                    v200 = append(v200, fi198)
                case 9223372036854775807:
                    v7 += int32(fi198)
                default:
                    var v201 uint64 = 18446744073709551615
                    _ = v201
                }
            default:
                var v202 int32 = -1
                _ = v202
            }
            v7 = int32(v189)
            _ = fi199
            _ = v200
        }
        _ = fi198
    }
    fmt.Println("v1:", v1)
    fmt.Println("v189:", v189)
    fmt.Println("v197:", v197)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v157
    _ = v158
    _ = v189
    _ = v197
    _ = v2
    _ = v3
    _ = v7
}
