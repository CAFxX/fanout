package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int32; F1 int8; F2 int }
type GenS4 struct { F0 int32; F1 uint16; F2 bool }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 string; F1 int8 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int32; F1 int8 }
func genFn0() int32 {
    var v5 uint8 = 197
    var v6 uint8 = (^(v5))
    if true {
        var v7 int16 = 0
        _ = v7
    } else {
        v8 := GenS1{F0: 4294967295, F1: 32767}
        if (-1 != 9223372036854775807) {
            v6 = (v5 & v6)
        } else {
            var v9 uint32 = 311
            var v10 uint = 202
            _ = v10
            _ = v9
        }
        if false {
            v11 := []int{720, 42}
            _ = v11
        } else {
            v12 := []int{2147483647, 2147483647, 2147483647}
            if (!false) {
                var v13 uint8 = v5
                _ = v13
            }
            _ = v12
        }
        _ = v8
    }
    fmt.Println(v5, v6)
    for fi14 := 0; fi14 < 12; fi14++ {
        switch "" {
        case "foo bar":
            if false {
                switch "\n\t" {
                case "a":
                    v15 := []int{fi14, fi14, fi14}
                    _ = v15
                case "\n\t":
                    v5 = 2
                }
                var v16 string = "\n\t"
                _ = v16
            }
        }
        v6 += (v6 ^ v6)
        _ = fi14
    }
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 uint8, p1 uint32) int32 {
    v17 := []int{-2147483648, 127, -9223372036854775808, 432}
    for fi18 := 0; fi18 < 2; fi18++ {
        for fi19 := 0; fi19 < 20; fi19++ {
            if ((p1 != 959) || true) {
                p0 = p0
                switch fi19 {
                case 2147483647:
                    switch (("a" + ("" + ((("a" + ("a" + "hello")) + ("\n\t" + "hello")) + "hello"))) + "a") {
                    case "a":
                        if (1 < 1000) {
                            switch "hello" {
                            case "hello":
                                v20 := GenS1{F0: 1, F1: 2}
                                _ = v20
                            }
                        }
                    }
                case 128:
                    var v21 int16 = -836
                    _ = v21
                default:
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                }
                fmt.Println(fi18, p1)
            }
            fmt.Println(fi19, p1)
            fmt.Println(p0, fi19, fi18)
            var v22 int8 = 0
            _ = fi19
            _ = v22
        }
        for fi23 := 0; fi23 < 9; fi23++ {
            if true {
                var v24 bool = true
                _ = v24
            } else {
                v25 := GenS2{F0: -816, F1: fi23}
                p1 = (p1 | (p1 * p1))
                p0 = 69
                _ = v25
            }
            _ = fi23
        }
        if true {
            v26 := GenS2{F0: genFn0(), F1: int(p1)}
            if true {
                switch "a" {
                case "a":
                    p1 += (p1 >> ((42) & 63))
                case "foo bar":
                    var v27 uint64 = 2
                    _ = v27
                }
                if false {
                    if len(v17) > 0 {
                        v17[0] = -241
                    }
                } else {
                    fmt.Println(p1)
                    v17 = append(v17, fi18)
                }
            }
            _ = v26
        }
        _ = fi18
    }
    v28 := GenS1{F0: 2, F1: 2}
    _ = p0
    _ = p1
    _ = v17
    _ = v28
    return -2147483648
}

func main() {
    var v1 bool = (false && (0.0 < ((-2.25 / (0.1 + (1e10 + (0.0 + 0.1)))) - 3.14159)))
    v2 := []string{}
    v3 := []uint32{}
    var v4 int32 = 1
    if false {
        var v29 string = ""
        _ = v29
    } else {
        var v30 int = 32767
        fmt.Println(v1)
        _ = v30
    }
    switch 107 {
    case 653:
        if (!v1) {
            var v31 float32 = 0.0
            var v32 int32 = genFn1(127, 42)
            _ = v31
            _ = v32
        } else {
            switch "hello" {
            case "\n\t":
                for fi33 := 0; fi33 < 16; fi33++ {
                    var v34 int64 = -9223372036854775808
                    v1 = true
                    switch ("foo bar" + "") {
                    case "foo bar":
                        v34 -= v34
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = fi33
                    _ = v34
                }
            default:
                var v35 float64 = 3.14159
                _ = v35
            }
        }
    case 9223372036854775807:
        if true {
            for fi36 := 0; fi36 < 1; fi36++ {
                if v1 {
                    v1 = v1
                } else {
                    if len(v2) > 0 {
                        v2[0] = (("a" + "\n\t") + "hello")
                    }
                    if (fi36 >= fi36) {
                        for fi37 := 0; fi37 < 4; fi37++ {
                            v1 = v1
                            switch fi37 {
                            case -487:
                                switch -1 {
                                case 710:
                                    v1 = (!((v4 / ((1) | 1)) == v4))
                                    if v1 {
                                        var v38 int32 = (v4 - (v4 % ((v4) | 1)))
                                        _ = v38
                                    } else {
                                        _ = len(v2)
                                        v4 += v4
                                    }
                                case 42:
                                    if false {
                                        if false {
                                            v4 = genFn0()
                                            v4 = v4
                                            var v39 float64 = (-2.25 / ((3.14159 / 3.14159) - 1e10))
                                            _ = v39
                                        }
                                        fmt.Println(fi36, fi37, fi37)
                                        var v40 uint8 = 1
                                        _ = v40
                                    } else {
                                        v4 = genFn0()
                                    }
                                case 9223372036854775807:
                                    v41 := GenS3{F0: (v4 >> ((4294967295) & 63)), F1: 0, F2: (fi36 / ((4294967295) | 1))}
                                    _ = v41
                                }
                            case 128:
                                switch fi36 {
                                case 2147483647:
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(v1)
                                    case "":
                                        fmt.Println(fi37, v1)
                                    case "hello":
                                        if (!false) {
                                            v1 = v1
                                            if ((!v1) || (1 != 42)) {
                                                _ = len(v2)
                                            }
                                            v4 += (v4 % ((genFn1(42, 990)) | 1))
                                        } else {
                                            v1 = true
                                            var v42 bool = ((v4 != -1) || (140 == 1))
                                            v43 := GenS2{F0: (v4 % (((-(v4))) | 1)), F1: 42}
                                            v44 := GenS1{F0: 4294967295, F1: 2}
                                            v4 += 2147483647
                                            v42 = false
                                            _ = v42
                                            _ = v43
                                            _ = v44
                                        }
                                    default:
                                        if ("" < "hello") {
                                            v45 := GenS2{F0: 0, F1: (fi36 | -940)}
                                            fmt.Println(v4, v4, fi36)
                                            v46 := GenS2{F0: v4, F1: 555}
                                            _ = v45
                                            _ = v46
                                        }
                                    }
                                case -2147483648:
                                    v4 = 1
                                }
                            case 0:
                                v4 = -1
                            default:
                                if (genFn1(60, 4294967295) != (^(v4))) {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = len(v3)
                                } else {
                                    var v47 float32 = 0.1
                                    var v48 uint32 = 65535
                                    _ = v47
                                    _ = v48
                                }
                            }
                            if len(v3) > 0 {
                                v3[0] = 0
                            }
                            var v49 bool = (v1 && ((!(9223372036854775807 >= 0)) && (("" + "\n\t") <= "")))
                            var v50 uint8 = 128
                            if v49 {
                                var v51 uint64 = 0
                                _ = v51
                            } else {
                                v4 = v4
                                var v52 uint = 18446744073709551615
                                fmt.Println(fi36, v52)
                                _ = v52
                            }
                            v1 = v1
                            _ = fi37
                            _ = v49
                            _ = v50
                        }
                        if v1 {
                            v4 = (v4 % ((v4) | 1))
                            v2 = append(v2, ("" + ""))
                        } else {
                            for fi53 := 0; fi53 < 4; fi53++ {
                                switch "" {
                                case "foo bar":
                                    fmt.Println(v4)
                                    if v1 {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v1 = (!v1)
                                        switch "a" {
                                        case "a":
                                            var v54 string = "a"
                                            _ = v54
                                        case "hello":
                                            v4 += v4
                                        case "\n\t":
                                            var v55 uint32 = 992
                                            _ = v55
                                        default:
                                            var v56 int64 = -9223372036854775808
                                            _ = v56
                                        }
                                        if v1 {
                                            var v57 int8 = 21
                                            _ = v57
                                        } else {
                                            if (!("\n\t" >= "a")) {
                                                v1 = v1
                                                if true {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    fmt.Println(v4, v4, fi36)
                                                } else {
                                                    var v58 uint32 = 0
                                                    if len(v2) > 0 {
                                                        v2[0] = "\n\t"
                                                    }
                                                    v4 += v4
                                                    _ = v58
                                                }
                                                v59 := GenS1{F0: 255, F1: 0}
                                                v4 = v4
                                                _ = v59
                                            }
                                            if (-903 <= 242) {
                                                var v60 uint32 = 1
                                                _ = v60
                                            }
                                            v1 = v1
                                        }
                                    }
                                case "a":
                                    v61 := GenS3{F0: v4, F1: 100, F2: fi53}
                                    _ = v61
                                case "":
                                    fmt.Println(v1, fi36, fi36)
                                default:
                                    v4 -= (-(v4))
                                    v1 = v1
                                    v1 = (!v1)
                                }
                                var v62 float64 = ((0.0 - 1.5) - 0.1)
                                var v63 string = ""
                                _ = fi53
                                _ = v62
                                _ = v63
                            }
                            v1 = (42 >= 229)
                            v4 += (v4 | v4)
                        }
                    }
                    fmt.Println(v1)
                }
                _ = fi36
            }
            if ((-9223372036854775808 <= 880) || v1) {
                v1 = false
            }
        }
    }
    v4 += (^(v4))
    for fi64 := 0; fi64 < 3; fi64++ {
        var v65 uint = 1
        v1 = v1
        _ = fi64
        _ = v65
    }
    v4 = (v4 / ((v4) | 1))
    v1 = v1
    if len(v3) > 0 {
        v3[0] = 1
    }
    if (!v1) {
        for fi66 := 0; fi66 < 12; fi66++ {
            for fi67 := 0; fi67 < 14; fi67++ {
                v68 := GenS4{F0: genFn0(), F1: 2, F2: v1}
                var v69 uint16 = 42
                v1 = ((0.1 - float64(0.1)) > (((1e10 / ((((1.5 - 0.1) / 3.14159) - (1.5 / float64((0.1 + 3.14159)))) + ((((0.0 + (((0.1 + float64(1.5)) + (float64(0.1) - 0.1)) - ((float64((0.0 - ((float32(-2.25) * 1e10) * 0.0))) - -2.25) / ((((float64(1e10) / ((0.1 + 0.0) * float64(0.1))) + (1e10 + 0.1)) / float64((0.1 - -2.25))) - 0.0)))) * 0.0) * 1.5) - 0.1))) / ((-2.25 - 0.1) + float64(((1.5 / -2.25) + 1.5)))) - 0.0))
                switch "" {
                case "\n\t":
                    if (false || v1) {
                        v70 := GenS1{F0: 255, F1: 42}
                        _ = v70
                    } else {
                        v2 = append(v2, ((("\n\t" + "") + "foo bar") + "foo bar"))
                        v71 := GenS4{F0: (v4 >> ((1) & 63)), F1: v69, F2: ((v1 || false) && (!v1))}
                        _ = v71
                    }
                case "foo bar":
                    v69 += (v69 * (^(v69)))
                    v3 = append(v3, 712)
                }
                fmt.Println(fi66, fi66)
                _ = fi67
                _ = v68
                _ = v69
            }
            _ = fi66
        }
        fmt.Println(v4)
    } else {
        v4 += v4
        if (!v1) {
            if ((65535 > 1) && false) {
                v3 = append(v3, 551)
                v72 := GenS2{F0: 1, F1: int(v4)}
                _ = v72
            } else {
                fmt.Println(v1, v1)
                for fi73 := 0; fi73 < 3; fi73++ {
                    for fi74 := 0; fi74 < 12; fi74++ {
                        v1 = ((false && ((float64(0.0) * 1e10) > (0.0 - float64(3.14159)))) || (42 < -1))
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                        v4 += v4
                        _ = fi74
                    }
                    _ = fi73
                }
                if v1 {
                    v3 = append(v3, 0)
                } else {
                    v2 = append(v2, ("a" + ""))
                    switch 0 {
                    case -1:
                        if (405 > 391) {
                            for fi75 := 0; fi75 < 13; fi75++ {
                                fmt.Println(v1, fi75)
                                _ = fi75
                            }
                            if v1 {
                                if v1 {
                                    v1 = (("a" + "hello") >= ("foo bar" + "\n\t"))
                                }
                                switch "" {
                                case "\n\t":
                                    v4 -= (v4 + (v4 + genFn1(2, 516)))
                                case "hello":
                                    for fi76 := 0; fi76 < 1; fi76++ {
                                        var v77 int8 = 31
                                        _ = fi76
                                        _ = v77
                                    }
                                case "":
                                    if len(v3) > 0 {
                                        v3[0] = 897
                                    }
                                default:
                                    for fi78 := 0; fi78 < 19; fi78++ {
                                        v79 := GenS5{F0: 2147483647, F1: (!v1)}
                                        v1 = false
                                        fmt.Println(fi78)
                                        _ = fi78
                                        _ = v79
                                    }
                                }
                            }
                        }
                    case 2147483647:
                        v1 = (47 != 0)
                    default:
                        var v80 float32 = 3.14159
                        _ = v80
                    }
                }
            }
            var v81 int16 = 175
            _ = v81
        } else {
            for fi82 := 0; fi82 < 18; fi82++ {
                switch "a" {
                case "foo bar":
                    var v83 int8 = 46
                    _ = v83
                case "\n\t":
                    var v84 int8 = -128
                    _ = v84
                default:
                    fmt.Println(fi82, v1)
                }
                _ = fi82
            }
            if (false || v1) {
                for fi85 := 0; fi85 < 12; fi85++ {
                    if true {
                        for fi86 := 0; fi86 < 1; fi86++ {
                            v4 = v4
                            _ = fi86
                        }
                        if (!(100 != 57)) {
                            fmt.Println(v4, v1, v1)
                        }
                    } else {
                        v4 -= v4
                    }
                    if false {
                        switch (("a" + "\n\t") + "a") {
                        case "foo bar":
                            var v87 bool = ((1 != 2) || (((v1 && ((true && ((-2.25 < (3.14159 / 1.5)) && ((!(true && (103 != 0))) || ((-32768 != 1) || v1)))) || (fi85 > (fi85 - fi85)))) && true) || v1))
                            _ = v87
                        case "\n\t":
                            v88 := GenS6{F0: 65535}
                            _ = v88
                        }
                        switch ("a" + "") {
                        case "":
                            var v89 uint32 = 2147483647
                            _ = v89
                        case "foo bar":
                            v2 = append(v2, "hello")
                        case "hello":
                            switch -774 {
                            case 128:
                                if len(v3) > 0 {
                                    v3[0] = 246
                                }
                            case 4294967295:
                                for fi90 := 0; fi90 < 12; fi90++ {
                                    v1 = v1
                                    _ = fi90
                                }
                            case -1:
                                fmt.Println(v1)
                            }
                        default:
                            if len(v2) > 0 {
                                v2[0] = "a"
                            }
                        }
                        switch fi85 {
                        case -9223372036854775808:
                            v91 := GenS7{F0: -32768}
                            _ = v91
                        case 336:
                            var v92 bool = true
                            _ = v92
                        }
                    }
                    for fi93 := 0; fi93 < 3; fi93++ {
                        if (((1 <= 2) && v1) && (0.1 >= ((0.1 - 0.1) / (1.5 + 1.5)))) {
                            if (255 < 575) {
                                v4 += v4
                                if v1 {
                                    var v94 uint = 0
                                    if v1 {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v2 = append(v2, ("" + "a"))
                                    }
                                    _ = v94
                                } else {
                                    var v95 string = ""
                                    switch (fi93 * fi85) {
                                    case 1:
                                        var v96 int64 = 0
                                        _ = v96
                                    case 4294967295:
                                        _ = len(v2)
                                    case 580:
                                        v4 = (v4 * v4)
                                        var v97 uint = 18446744073709551615
                                        fmt.Println(v4, fi85, fi85)
                                        _ = v97
                                    default:
                                        switch v95 {
                                        case "foo bar":
                                            v2 = append(v2, "")
                                        case "":
                                            v95 = v95
                                        case "a":
                                            v1 = true
                                        }
                                        if v1 {
                                            v98 := GenS1{F0: 18446744073709551615, F1: 42}
                                            var v99 int16 = 42
                                            var v100 float64 = (-2.25 - ((0.1 * 0.1) - ((((3.14159 / 0.1) * 0.1) * float64((0.1 / -2.25))) * (0.0 + 3.14159))))
                                            switch "foo bar" {
                                            case "\n\t":
                                                v99 += (-(v99))
                                            }
                                            switch (-(fi93)) {
                                            case -55:
                                                var v101 float64 = float64((0.0 - float32(v100)))
                                                _ = v101
                                            default:
                                                var v102 int = (fi93 + 2147483647)
                                                _ = v102
                                            }
                                            v103 := GenS8{F0: (((v95 + (v95 + "\n\t")) + v95) + (v95 + "hello")), F1: 100}
                                            v1 = false
                                            _ = v100
                                            _ = v103
                                            _ = v98
                                            _ = v99
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            if (true || (-83 == -1)) {
                                                _ = len(v2)
                                                var v104 uint32 = 2147483648
                                                _ = v104
                                            }
                                            var v105 bool = v1
                                            _ = v105
                                        }
                                        if (float32((1.5 * 1e10)) <= (-2.25 * (0.0 / ((3.14159 + 0.0) - -2.25)))) {
                                            var v106 int32 = v4
                                            _ = len(v2)
                                            fmt.Println(v1, v4)
                                            _ = v106
                                        } else {
                                            v4 = v4
                                            v4 = v4
                                            v95 = "a"
                                        }
                                    }
                                    v107 := GenS9{F0: (("foo bar" + v95) + (v95 + ("a" + v95)))}
                                    _ = v107
                                    _ = v95
                                }
                                v4 += (v4 * genFn0())
                            }
                            v108 := GenS1{F0: 662, F1: 0}
                            var v109 uint = 0
                            _ = v108
                            _ = v109
                        } else {
                            if v1 {
                                if v1 {
                                    v1 = (149 == 2)
                                } else {
                                    if (!true) {
                                        v1 = false
                                    } else {
                                        fmt.Println(fi93)
                                        _ = len(v2)
                                        _ = len(v2)
                                    }
                                    if true {
                                        var v110 int32 = v4
                                        switch "" {
                                        case "\n\t":
                                            if ((1e10 - ((0.1 - float32(0.0)) + -2.25)) > 1e10) {
                                                switch "a" {
                                                case "":
                                                    var v111 int64 = -545
                                                    _ = v111
                                                }
                                                v110 -= (v110 + (v110 << ((331) & 63)))
                                            } else {
                                                v4 = 1
                                            }
                                        }
                                        switch "a" {
                                        case "\n\t":
                                            var v112 bool = v1
                                            _ = v112
                                        default:
                                            v113 := GenS8{F0: "", F1: 1}
                                            _ = v113
                                        }
                                        var v114 float32 = (1e10 + (3.14159 * 3.14159))
                                        _ = v110
                                        _ = v114
                                    } else {
                                        v4 = v4
                                        var v115 float64 = 1.5
                                        _ = v115
                                    }
                                }
                                v1 = true
                            } else {
                                if true {
                                    v4 = (v4 + int32(fi93))
                                    var v116 uint16 = 181
                                    _ = v116
                                } else {
                                    v117 := GenS4{F0: genFn0(), F1: 32767, F2: v1}
                                    var v118 float32 = ((float32(-2.25) + ((3.14159 - (float32((-2.25 - float64(3.14159))) - 0.0)) + 1.5)) / (3.14159 * (-2.25 - (1e10 - 1e10))))
                                    switch 704 {
                                    case 128:
                                        v119 := GenS8{F0: "a", F1: 0}
                                        _ = v119
                                    }
                                    _ = v117
                                    _ = v118
                                }
                                var v120 bool = (1 == 2)
                                _ = v120
                            }
                        }
                        v1 = ((9223372036854775807 != 0) || (true && (760 == 32767)))
                        v1 = false
                        _ = fi93
                    }
                    _ = fi85
                }
                for fi121 := 0; fi121 < 19; fi121++ {
                    if ("\n\t" >= "hello") {
                        var v122 int16 = -32768
                        v2 = append(v2, "")
                        _ = v122
                    }
                    _ = fi121
                }
                for fi123 := 0; fi123 < 14; fi123++ {
                    if v1 {
                        if ((0 == 2) || ((^(fi123)) >= (fi123 + -375))) {
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = (("hello" + "") + "\n\t")
                                }
                            }
                        }
                        if (false || ((-2.25 / ((0.1 * 1e10) - -2.25)) == (0.0 - float32(((0.0 - float64(float32(0.0))) * (float64(1.5) + (1e10 + 3.14159))))))) {
                            _ = len(v2)
                            var v124 int16 = -32768
                            switch "a" {
                            case "\n\t":
                                v124 = v124
                            case "foo bar":
                                var v125 int16 = 163
                                _ = v125
                            case "hello":
                                var v126 uint8 = 182
                                _ = v126
                            }
                            _ = v124
                        }
                        for fi127 := 0; fi127 < 20; fi127++ {
                            v128 := GenS10{F0: (v4 << ((18446744073709551615) & 63)), F1: -13}
                            _ = fi127
                            _ = v128
                        }
                    }
                    v129 := GenS4{F0: genFn0(), F1: 266, F2: (v1 || v1)}
                    _ = fi123
                    _ = v129
                }
            }
        }
        var v130 int32 = v4
        _ = v130
    }
    v4 -= genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
