package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int16; F1 int32; F2 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 uint; F1 bool }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int8; F1 int32; F2 uint32 }
type GenS11 struct { F0 bool; F1 string }
type GenS12 struct { F0 int64; F1 uint64 }
func genFn0() uint {
    v5 := GenS1{F0: 2, F1: 1}
    switch 32767 {
    case 1:
        var v6 int16 = 476
        _ = v6
    default:
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 bool = false
            fmt.Println(v8, fi7)
            _ = fi7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 bool = true
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        fmt.Println(fi11)
        if true {
            v12 := []int{(fi11 << ((65535) & 63)), fi11}
            switch (("hello" + ("a" + "hello")) + "\n\t") {
            case "foo bar":
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 int16 = 1
                    var v15 uint = 117
                    v14 -= (v14 * (v14 << ((v15) & 63)))
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            }
            var v16 bool = false
            _ = v12
            _ = v16
        }
        _ = fi11
    }
    _ = v5
    return 0
}

func main() {
    var v1 bool = ((true && false) || true)
    var v2 float32 = 0.1
    v3 := make(map[int]uint32)
    var v4 uint = 4294967295
    for fi17 := 0; fi17 < 20; fi17++ {
        switch "\n\t" {
        case "\n\t":
            for fi18 := 0; fi18 < 20; fi18++ {
                var v19 int64 = -899
                _ = fi18
                _ = v19
            }
        case "hello":
            v20 := GenS1{F0: 786, F1: 0}
            _ = v20
        case "":
            var v21 int64 = -9223372036854775808
            _ = v21
        default:
            v22 := []int{}
            _ = v22
        }
        _ = fi17
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        if true {
            v2 = (v2 * v2)
        } else {
            switch (("" + ("foo bar" + "foo bar")) + "a") {
            case "hello":
                for fi24 := 0; fi24 < 11; fi24++ {
                    var v25 uint32 = 42
                    _ = fi24
                    _ = v25
                }
            case "\n\t":
                v2 = (v2 - (v2 / (((0.1 * v2) * ((1e10 / (v2 * (((v2 / ((float32((1e10 - float64(v2))) - -2.25) + (0.0 / ((v2 * 0.1) + ((v2 * (v2 + v2)) / ((1.5 * 0.0) - 1.5)))))) * -2.25) + 1.5))) + (((((v2 + v2) - ((-2.25 / float32(1e10)) * 3.14159)) - float32((0.0 * 3.14159))) + ((3.14159 * v2) * ((1.5 + v2) * 0.1))) / float32(-2.25)))) * v2)))
            case "a":
                v4 = v4
            default:
                for fi26 := 0; fi26 < 6; fi26++ {
                    fmt.Println(fi26, v4, fi26)
                    _ = fi26
                }
            }
            switch ("a" + "foo bar") {
            case "hello":
                if (262 <= 1) {
                    switch fi23 {
                    case 0:
                        var v27 uint = v4
                        _ = v27
                    case -1:
                        for fi28 := 0; fi28 < 4; fi28++ {
                            v1 = (!v1)
                            switch fi28 {
                            case -1:
                                if (716 < 42) {
                                    if false {
                                        _ = len(v3)
                                    } else {
                                        var v29 uint = (v4 & (v4 >> (((^(v4))) & 63)))
                                        v30 := []int{fi23, fi23, fi23, (^(fi23))}
                                        _ = v29
                                        _ = v30
                                    }
                                }
                            case 32768:
                                if v1 {
                                    switch (("a" + (("foo bar" + ("" + "foo bar")) + "a")) + (("" + "a") + ("a" + ""))) {
                                    case "hello":
                                        var v31 uint8 = 42
                                        _ = v31
                                    case "foo bar":
                                        switch ("foo bar" + "hello") {
                                        case "foo bar":
                                            v1 = v1
                                        }
                                    case "a":
                                        var v32 float64 = ((-2.25 * (float64(v2) + -2.25)) - 1.5)
                                        _ = v32
                                    }
                                    var v33 bool = v1
                                    v34 := GenS2{F0: 1, F1: 1, F2: 2}
                                    _ = v33
                                    _ = v34
                                } else {
                                    var v35 int8 = -128
                                    _ = v35
                                }
                            case 2147483647:
                                v4 = (^(v4))
                            default:
                                fmt.Println(fi28)
                            }
                            _ = fi28
                        }
                    case 32768:
                        if (true && v1) {
                            v2 = 0.0
                            var v36 int64 = 42
                            switch "a" {
                            case "a":
                                var v37 string = "hello"
                                _ = v37
                            case "\n\t":
                                _ = len(v3)
                            default:
                                if v1 {
                                    v38 := GenS2{F0: -817, F1: -867, F2: 127}
                                    v2 = 0.1
                                    _ = v38
                                }
                            }
                            _ = v36
                        } else {
                            v39 := GenS2{F0: 32767, F1: 655, F2: 238}
                            for fi40 := 0; fi40 < 7; fi40++ {
                                var v41 uint8 = 127
                                switch fi40 {
                                case 312:
                                    delete(v3, fi40)
                                case 0:
                                    v42 := GenS2{F0: 367, F1: 42, F2: v41}
                                    _ = v42
                                default:
                                    v43 := GenS3{F0: 127}
                                    _ = v43
                                }
                                _ = fi40
                                _ = v41
                            }
                            var v44 float32 = (1e10 / (v2 - 1e10))
                            _ = v39
                            _ = v44
                        }
                    }
                } else {
                    fmt.Println(v4, v4, v4)
                }
            case "":
                var v45 int16 = 42
                _ = v45
            case "\n\t":
                v2 = float32(-2.25)
            }
            for fi46 := 0; fi46 < 10; fi46++ {
                v1 = false
                if true {
                    switch ("" + "foo bar") {
                    case "foo bar":
                        v3[fi46] = 188
                    default:
                        if ("" == ("" + ("hello" + "hello"))) {
                            var v47 string = ("\n\t" + ("hello" + "\n\t"))
                            if true {
                                v48 := GenS3{F0: -42}
                                if (-2147483648 <= 1) {
                                    fmt.Println(fi23, v2, fi23)
                                    v1 = v1
                                } else {
                                    if v1 {
                                        v2 = v2
                                        v2 -= v2
                                        v2 += 0.0
                                    }
                                    switch v47 {
                                    case "a":
                                        v1 = ((v4 << (((v4 - 18446744073709551615)) & 63)) > v4)
                                    }
                                    var v49 int16 = -912
                                    _ = v49
                                }
                                var v50 string = "a"
                                _ = v48
                                _ = v50
                            } else {
                                v1 = v1
                                v47 = v47
                                v51 := []int{32767, fi46}
                                switch fi46 {
                                case 155:
                                    var v52 int16 = 32767
                                    if false {
                                        v4 -= genFn0()
                                        v53 := GenS3{F0: 0}
                                        _ = v53
                                    }
                                    _ = v52
                                }
                                if v1 {
                                    if false {
                                        switch (^(fi46)) {
                                        case 1:
                                            if v1 {
                                                v2 -= v2
                                                fmt.Println(v4)
                                            } else {
                                                v1 = (!(v47 < v47))
                                                v54 := GenS1{F0: 225, F1: (^(v4))}
                                                _ = v54
                                            }
                                        case 32767:
                                            v47 = ("hello" + "")
                                        case 2147483647:
                                            if (2 != 0) {
                                                if true {
                                                    if v1 {
                                                        v55 := GenS4{F0: 65535, F1: 42}
                                                        _ = v55
                                                    } else {
                                                        v56 := v3[fi46]
                                                        switch 4294967295 {
                                                        case 127:
                                                            v47 = ""
                                                        case 42:
                                                            v57 := GenS2{F0: -32768, F1: 810, F2: 231}
                                                            _ = v57
                                                        case 128:
                                                            switch fi23 {
                                                            case -1:
                                                                if true {
                                                                    v2 = (v2 - (((v2 + v2) / float32(0.1)) / v2))
                                                                    v58 := v3[(fi23 & fi23)]
                                                                    if (!(!false)) {
                                                                        var v59 float32 = v2
                                                                        var v60 int = (fi23 % (((fi23 ^ 9223372036854775807)) | 1))
                                                                        v58 -= (v56 - 65535)
                                                                        if (v47 <= v47) {
                                                                            var v61 uint32 = 891
                                                                            v4 = genFn0()
                                                                            v2 = v2
                                                                            var v62 int = 234
                                                                            _ = v61
                                                                            _ = v62
                                                                        }
                                                                        var v63 uint16 = uint16(v58)
                                                                        _ = v59
                                                                        _ = v60
                                                                        _ = v63
                                                                    } else {
                                                                        fmt.Println(v56, fi46)
                                                                        v56 = v56
                                                                    }
                                                                    _ = v58
                                                                } else {
                                                                    fmt.Println(fi23, fi23, fi23)
                                                                }
                                                            case 2147483647:
                                                                v2 += (v2 / v2)
                                                            default:
                                                                v64 := GenS1{F0: 255, F1: 37}
                                                                _ = v64
                                                            }
                                                        default:
                                                            _ = len(v51)
                                                        }
                                                        var v65 uint32 = v56
                                                        _ = v56
                                                        _ = v65
                                                    }
                                                    var v66 uint16 = 21
                                                    _ = v66
                                                }
                                                var v67 int8 = 127
                                                if false {
                                                    v67 += -1
                                                    switch (v47 + v47) {
                                                    case "hello":
                                                        v68 := GenS5{F0: 2, F1: -651, F2: 883}
                                                        _ = v68
                                                    }
                                                    if v1 {
                                                        if len(v51) >= 2 {
                                                            v51 = v51[1:len(v51)]
                                                        }
                                                        v67 = 42
                                                        v4 = v4
                                                    }
                                                }
                                                _ = v67
                                            } else {
                                                var v69 int32 = 1
                                                _ = v69
                                            }
                                        default:
                                            v4 += (^(v4))
                                        }
                                        if v1 {
                                            v4 = (v4 - (^(v4)))
                                        } else {
                                            var v70 int16 = -483
                                            _ = v70
                                        }
                                        v4 = (^(v4))
                                    } else {
                                        var v71 float64 = (1e10 / float64(1.5))
                                        var v72 bool = v1
                                        v71 += -2.25
                                        _ = v71
                                        _ = v72
                                    }
                                    var v73 string = "hello"
                                    v73 = (("" + ("\n\t" + "")) + v73)
                                    v3[fi23] = 2147483648
                                    if v1 {
                                        var v74 float32 = float32(1.5)
                                        _ = len(v3)
                                        _ = v74
                                    } else {
                                        switch "" {
                                        case "":
                                            v4 = 421
                                        case "\n\t":
                                            v75 := GenS6{F0: 924, F1: v47}
                                            if (!((false || ((v1 && ((v1 || (0 > -1)) && v1)) && v1)) && ((v4 % (((v4 - (^(v4)))) | 1)) > v4))) {
                                                v76 := GenS7{F0: v4, F1: v1}
                                                _ = v76
                                            }
                                            _ = v75
                                        case "foo bar":
                                            var v77 uint64 = 0
                                            _ = v77
                                        }
                                        var v78 int64 = 42
                                        _ = v78
                                    }
                                    v79 := GenS1{F0: 1, F1: (^(v4))}
                                    _ = v73
                                    _ = v79
                                }
                                var v80 uint32 = 0
                                _ = v51
                                _ = v80
                            }
                            _ = v47
                        } else {
                            var v81 string = ("" + "foo bar")
                            _ = v81
                        }
                    }
                    v2 = v2
                } else {
                    v4 = (v4 * (^(v4)))
                    var v82 int32 = -338
                    v83 := []int{(fi46 & (fi23 % ((fi46) | 1)))}
                    if (true && ((0 > 654) && true)) {
                        if false {
                            var v84 int16 = 32767
                            if (true && false) {
                                fmt.Println(fi23, v82)
                                var v85 int8 = -128
                                switch "a" {
                                case "a":
                                    if true {
                                        v1 = v1
                                    } else {
                                        fmt.Println(v1)
                                        _ = len(v3)
                                    }
                                case "hello":
                                    var v86 int32 = v82
                                    var v87 int64 = 0
                                    v87 += 9223372036854775807
                                    _ = v86
                                    _ = v87
                                }
                                _ = v85
                            } else {
                                fmt.Println(fi46, fi46, v4)
                                if false {
                                    var v88 uint = 54
                                    if (0 >= 185) {
                                        var v89 int64 = 753
                                        _ = v89
                                    } else {
                                        var v90 uint64 = 753
                                        fmt.Println(v90)
                                        _ = v90
                                    }
                                    _ = v88
                                } else {
                                    var v91 int16 = (v84 | (v84 & -1))
                                    fmt.Println(v84, fi46, fi46)
                                    v82 = v82
                                    fmt.Println(v82, v91)
                                    var v92 uint = (v4 / ((v4) | 1))
                                    _ = v91
                                    _ = v92
                                }
                                v1 = ((true && v1) && v1)
                                v3[fi23] = 1
                                fmt.Println(v1)
                            }
                            _ = v84
                        }
                        v1 = v1
                        var v93 uint8 = 128
                        _ = v93
                    } else {
                        if (4294967295 <= 2147483648) {
                            var v94 float64 = (1e10 * (-2.25 - 3.14159))
                            v94 = 3.14159
                            v4 = 814
                            _ = v94
                        } else {
                            if true {
                                var v95 int = (fi46 * fi46)
                                v1 = ((float64(v2) < float64(v2)) || false)
                                v96 := GenS8{F0: 576}
                                _ = v95
                                _ = v96
                            }
                        }
                        switch "hello" {
                        case "hello":
                            if (!false) {
                                switch ("foo bar" + ("foo bar" + "hello")) {
                                case "foo bar":
                                    if len(v83) > 0 {
                                        v83[0] = (fi46 % ((fi23) | 1))
                                    }
                                    var v97 uint8 = 141
                                    var v98 uint64 = 4294967295
                                    _ = v97
                                    _ = v98
                                case "":
                                    v99 := GenS7{F0: v4, F1: true}
                                    _ = v99
                                case "\n\t":
                                    fmt.Println(v2, v2, v2)
                                default:
                                    switch "a" {
                                    case "foo bar":
                                        v4 += 255
                                        switch ("hello" + "a") {
                                        case "hello":
                                            if false {
                                                if ((500 > 9223372036854775807) || (!v1)) {
                                                    if false {
                                                        var v100 uint64 = 0
                                                        if (42 == 1) {
                                                            _ = len(v83)
                                                            v1 = v1
                                                        } else {
                                                            var v101 float64 = (float64((v2 + (3.14159 - v2))) - (0.0 * -2.25))
                                                            v102 := v3[fi23]
                                                            var v103 int32 = v82
                                                            _ = v101
                                                            _ = v102
                                                            _ = v103
                                                        }
                                                        _ = v100
                                                    } else {
                                                        var v104 uint8 = 128
                                                        _ = v104
                                                    }
                                                    delete(v3, (fi23 ^ -1))
                                                } else {
                                                    v1 = false
                                                    fmt.Println(v82)
                                                    fmt.Println(v82, v1, v1)
                                                    v4 -= (v4 + 0)
                                                }
                                                v1 = v1
                                            } else {
                                                var v105 uint8 = uint8(v4)
                                                _ = v105
                                            }
                                            var v106 string = ("hello" + "a")
                                            var v107 uint = 1
                                            _ = v106
                                            _ = v107
                                        case "a":
                                            if (("a" == "foo bar") || ((0 > 1) && v1)) {
                                                if (v82 > v82) {
                                                    if v1 {
                                                        v108 := v3[(^(fi46))]
                                                        switch (fi46 - fi46) {
                                                        case 297:
                                                            switch 0 {
                                                            case 2147483647:
                                                                var v109 int = 127
                                                                _ = v109
                                                            case 335:
                                                                var v110 int8 = int8(v4)
                                                                _ = v110
                                                            }
                                                        case 252:
                                                            var v111 uint = (v4 & uint(v108))
                                                            if (!v1) {
                                                                if v1 {
                                                                    var v112 uint32 = 719
                                                                    fmt.Println(v1)
                                                                    v1 = (!v1)
                                                                    _ = v112
                                                                } else {
                                                                    if (false && false) {
                                                                        v113 := GenS3{F0: 1}
                                                                        var v114 int16 = -32768
                                                                        if v1 {
                                                                            v115 := GenS7{F0: 472, F1: ((1e10 / (-2.25 / (((float64(((1e10 / v2) * ((((v2 / (v2 - v2)) - 1.5) - (float32((0.1 + ((1e10 * 3.14159) / 1e10))) - v2)) + (v2 * v2)))) / (-2.25 - 1e10)) * 0.1) + 1.5))) > (0.0 - ((1.5 * -2.25) / 1.0)))}
                                                                            var v116 uint8 = 204
                                                                            var v117 uint16 = 543
                                                                            _ = len(v3)
                                                                            v118 := GenS9{F0: int8(v111)}
                                                                            _ = v115
                                                                            _ = v116
                                                                            _ = v117
                                                                            _ = v118
                                                                        } else {
                                                                            var v119 uint64 = 304
                                                                            if (v4 != (^(v111))) {
                                                                                fmt.Println(fi46, v111, v2)
                                                                                _ = len(v83)
                                                                            } else {
                                                                                v82 = (v82 >> (((v4 ^ v111)) & 63))
                                                                                v1 = v1
                                                                                var v120 int64 = -860
                                                                                v108 = v108
                                                                                var v121 uint8 = 128
                                                                                if v1 {
                                                                                    var v122 uint8 = v121
                                                                                    _ = v122
                                                                                }
                                                                                _ = v120
                                                                                _ = v121
                                                                            }
                                                                            var v123 bool = v1
                                                                            _ = v119
                                                                            _ = v123
                                                                        }
                                                                        _ = v113
                                                                        _ = v114
                                                                    } else {
                                                                        if v1 {
                                                                            switch ("a" + "") {
                                                                            case "\n\t":
                                                                                fmt.Println(fi23, fi46)
                                                                            }
                                                                            if false {
                                                                                var v124 uint8 = 1
                                                                                switch fi46 {
                                                                                case 32768:
                                                                                    var v125 int32 = v82
                                                                                    _ = v125
                                                                                default:
                                                                                    var v126 int8 = 100
                                                                                    _ = v126
                                                                                }
                                                                                fmt.Println(v2)
                                                                                _ = v124
                                                                            }
                                                                            switch "hello" {
                                                                            case "":
                                                                                var v127 int16 = 42
                                                                                _ = v127
                                                                            case "\n\t":
                                                                                var v128 uint = 1
                                                                                _ = v128
                                                                            case "a":
                                                                                var v129 string = "\n\t"
                                                                                _ = v129
                                                                            default:
                                                                                v130 := GenS3{F0: 5}
                                                                                _ = v130
                                                                            }
                                                                        }
                                                                        fmt.Println(fi23)
                                                                        v1 = (!(false && true))
                                                                    }
                                                                    _ = len(v3)
                                                                    v3[fi23] = 0
                                                                }
                                                                v131 := GenS2{F0: -32768, F1: (v82 + (-(v82))), F2: 255}
                                                                var v132 uint8 = 252
                                                                _ = v131
                                                                _ = v132
                                                            }
                                                            _ = v111
                                                        default:
                                                            var v133 int8 = 7
                                                            _ = v133
                                                        }
                                                        _ = v108
                                                    }
                                                    switch fi23 {
                                                    case 1:
                                                        var v134 uint64 = uint64(v4)
                                                        _ = v134
                                                    case 32768:
                                                        if len(v83) >= 2 {
                                                            v83 = v83[1:len(v83)]
                                                        }
                                                    default:
                                                        v1 = (fi23 >= fi23)
                                                    }
                                                    var v135 uint32 = 2147483647
                                                    v1 = v1
                                                    if ((1.5 + 0.0) < ((3.14159 - (-2.25 / 0.1)) + -2.25)) {
                                                        v82 = -1
                                                    } else {
                                                        if ((fi23 * fi46) == (fi46 % ((fi23) | 1))) {
                                                            if len(v83) > 0 {
                                                                v83[0] = int(v135)
                                                            }
                                                        }
                                                    }
                                                    v82 += (v82 % ((v82) | 1))
                                                    _ = v135
                                                }
                                                _ = len(v3)
                                            }
                                        default:
                                            switch fi23 {
                                            case 2147483647:
                                                v136 := GenS5{F0: 75, F1: -32768, F2: 1}
                                                _ = v136
                                            default:
                                                v3[2147483647] = 2147483648
                                            }
                                        }
                                        v4 += genFn0()
                                    default:
                                        fmt.Println(fi46, fi46)
                                    }
                                }
                            }
                        case "a":
                            switch (("\n\t" + ("\n\t" + ("" + ("foo bar" + "foo bar")))) + ("\n\t" + "")) {
                            case "hello":
                                var v137 int = fi23
                                _ = v137
                            }
                        default:
                            switch -858 {
                            case 925:
                                switch int(v82) {
                                case 1:
                                    if v1 {
                                        if (("\n\t" + "\n\t") > "\n\t") {
                                            delete(v3, (fi46 * (fi23 | 32768)))
                                            var v138 bool = (-32768 != 315)
                                            _ = v138
                                        } else {
                                            v82 -= 42
                                        }
                                    } else {
                                        v139 := GenS5{F0: 127, F1: 1, F2: 922}
                                        var v140 int = 32767
                                        _ = v139
                                        _ = v140
                                    }
                                case 9223372036854775807:
                                    switch ("hello" + "hello") {
                                    case "":
                                        if len(v83) > 0 {
                                            v83[0] = (fi46 + (fi23 + (fi46 - fi46)))
                                        }
                                    }
                                case 454:
                                    v1 = v1
                                default:
                                    v141 := v3[(fi23 ^ fi23)]
                                    _ = v141
                                }
                            case 128:
                                if v1 {
                                    if true {
                                        fmt.Println(fi23, fi23)
                                        v1 = ((255 <= 1) && (32767 != 42))
                                    } else {
                                        if v1 {
                                            fmt.Println(v4, fi23, v2)
                                            v1 = (!(v1 && false))
                                        } else {
                                            var v142 uint = (v4 % ((977) | 1))
                                            v4 -= 726
                                            _ = v142
                                        }
                                        var v143 uint16 = 32767
                                        _ = v143
                                    }
                                    _ = len(v3)
                                } else {
                                    var v144 int16 = 864
                                    _ = v144
                                }
                            case 908:
                                if v1 {
                                    switch fi23 {
                                    case -1:
                                        var v145 uint16 = 0
                                        _ = v145
                                    case 32767:
                                        v82 -= (v82 * (v82 * int32(v4)))
                                    }
                                    switch fi23 {
                                    case 63:
                                        if len(v83) > 0 {
                                            v83[0] = fi46
                                        }
                                    }
                                } else {
                                    var v146 uint = (v4 & (v4 * v4))
                                    _ = v146
                                }
                            }
                            if len(v83) > 0 {
                                v83[0] = (fi23 / ((4294967295) | 1))
                            }
                        }
                        v2 += (1.5 / v2)
                    }
                    _ = v82
                    _ = v83
                }
                v1 = (!((2 >= 32767) || ((false || (false && true)) || true)))
                var v147 int32 = 0
                _ = fi46
                _ = v147
            }
        }
        _ = fi23
    }
    v2 = (float32(0.0) * v2)
    switch "foo bar" {
    case "hello":
        v2 = 3.14159
    case "foo bar":
        v148 := GenS10{F0: 42, F1: 1, F2: 423}
        _ = v148
    }
    for fi149 := 0; fi149 < 19; fi149++ {
        if v1 {
            if v1 {
                v2 -= v2
            }
            v150 := GenS2{F0: 1000, F1: 0, F2: 128}
            v151 := []int{fi149}
            _ = v150
            _ = v151
        }
        v3[-9223372036854775808] = 2147483647
        _ = fi149
    }
    if (!(v1 || (-2.25 > v2))) {
        v152 := GenS11{F0: (!false), F1: ("" + "\n\t")}
        _ = v152
    } else {
        v4 += v4
    }
    delete(v3, 1)
    v153 := GenS12{F0: -9223372036854775808, F1: 2}
    for fi154 := 0; fi154 < 15; fi154++ {
        switch fi154 {
        case 1:
            v155 := GenS4{F0: 747, F1: 1}
            _ = v155
        case 2147483647:
            v2 = 3.14159
        }
        switch "foo bar" {
        case "foo bar":
            _ = len(v3)
        case "a":
            v4 -= (v4 << ((v4) & 63))
        case "hello":
            v156 := []int{(^(fi154)), fi154, -1, fi154}
            _ = v156
        default:
            for fi157 := 0; fi157 < 6; fi157++ {
                var v158 int = fi154
                var v159 float64 = 3.14159
                _ = fi157
                _ = v158
                _ = v159
            }
        }
        v160 := []int{(fi154 | -2147483648), (fi154 & fi154), fi154, (fi154 + (fi154 << ((0) & 63)))}
        _ = fi154
        _ = v160
    }
    var v161 int32 = 1
    var v162 uint32 = 2147483648
    delete(v3, 375)
    for fi163 := 0; fi163 < 20; fi163++ {
        v1 = false
        var v164 float32 = (0.1 / (-2.25 + 3.14159))
        v164 -= (1e10 + float32((3.14159 / 1.5)))
        _ = fi163
        _ = v164
    }
    v3[0] = 0
    fmt.Println("v1:", v1)
    fmt.Println("v161:", v161)
    fmt.Println("v162:", v162)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v153
    _ = v161
    _ = v162
    _ = v2
    _ = v3
    _ = v4
}
