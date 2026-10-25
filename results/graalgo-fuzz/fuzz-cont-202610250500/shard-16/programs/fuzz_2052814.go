package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 int64; F1 uint16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint; F1 uint16 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int32 }
func genFn0() uint {
    var v5 string = "a"
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 int16 = 0
        v8 := GenS1{F0: 1, F1: 44, F2: 255}
        _ = fi6
        _ = v7
        _ = v8
    }
    if ((!((!(!(((v5 + v5) == (v5 + "a")) && (!false)))) || true)) && true) {
        var v9 uint = 928
        v10 := []int{-2147483648, 2147483647, 127}
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 int8 = 1
        fmt.Println(v12, v12)
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 int32 = -487
            fmt.Println(v12, v14, v14)
            if ((!(18446744073709551615 == 0)) || true) {
                if (-2.25 < 0.0) {
                    var v15 int = -2147483648
                    var v16 float64 = (float64(1.5) - 1.5)
                    _ = v15
                    _ = v16
                } else {
                    v5 = "\n\t"
                    if (false || true) {
                        v17 := GenS1{F0: (fi11 * fi11), F1: 1, F2: 4294967295}
                        v14 = v14
                        _ = v17
                    } else {
                        if ((!(42 > 127)) || ((-(fi11)) == 32767)) {
                            v14 -= v14
                        }
                        fmt.Println(fi11, fi11)
                    }
                }
                if (v14 == (v14 >> ((830) & 63))) {
                    v12 += 42
                    fmt.Println(v5)
                    fmt.Println(fi13, fi11)
                    switch 0 {
                    case -2147483648:
                        if (2 != 0) {
                            v5 = v5
                            v14 = 42
                        }
                    }
                    v5 = v5
                }
                var v18 int32 = v14
                _ = v18
            }
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    var v19 uint64 = 466
    _ = v19
    _ = v5
    return 863
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 0
    var v3 uint64 = 1
    var v4 int32 = 856
    switch -2147483648 {
    case -618:
        if ((1.5 / 1e10) > 0.1) {
            for fi20 := 0; fi20 < 6; fi20++ {
                v3 = (v3 >> ((genFn0()) & 63))
                if true {
                    var v21 bool = false
                    v2 += -32768
                    _ = v21
                } else {
                    var v22 uint64 = (^(v3))
                    _ = v22
                }
                if true {
                    v2 += (v2 - v2)
                }
                _ = fi20
            }
            v4 += (v4 << ((868) & 63))
        } else {
            for fi23 := 0; fi23 < 5; fi23++ {
                v24 := GenS2{F0: "hello", F1: 2147483647}
                v1[268] = "foo bar"
                var v25 float32 = (((1e10 - (1.5 - 1.5)) * (3.14159 - 1e10)) / float32(3.14159))
                _ = fi23
                _ = v24
                _ = v25
            }
            v26 := GenS2{F0: ("a" + ""), F1: 222}
            v27 := GenS3{F0: "foo bar", F1: 0}
            _ = v26
            _ = v27
        }
    case 0:
        v28 := GenS3{F0: "\n\t", F1: 99}
        _ = v28
    case -928:
        if false {
            v29 := GenS4{F0: v2}
            if false {
                for fi30 := 0; fi30 < 11; fi30++ {
                    v31 := GenS1{F0: fi30, F1: 316, F2: 42}
                    v32 := []int{fi30}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                switch "hello" {
                case "":
                    var v33 uint = 488
                    _ = v33
                }
                if true {
                    switch -435 {
                    case 42:
                        for fi34 := 0; fi34 < 11; fi34++ {
                            if (false || true) {
                                v35 := GenS4{F0: (v2 - (v2 & (v2 * (v2 / (((v2 & (v2 >> ((550) & 63)))) | 1)))))}
                                v2 = v2
                                _ = v35
                            }
                            _ = fi34
                        }
                    case -995:
                        if (2 > (v3 % ((0) | 1))) {
                            var v36 int32 = v4
                            v37 := GenS3{F0: "foo bar", F1: 0}
                            var v38 bool = ((32767 < 65535) && (65535 != 32767))
                            _ = v36
                            _ = v37
                            _ = v38
                        } else {
                            v39 := GenS3{F0: ("foo bar" + (("hello" + "\n\t") + "hello")), F1: 42}
                            for fi40 := 0; fi40 < 4; fi40++ {
                                if (32767 != 1) {
                                    v3 -= 2
                                    for fi41 := 0; fi41 < 10; fi41++ {
                                        switch fi41 {
                                        case -1:
                                            v2 -= 1
                                        case -433:
                                            v4 = v4
                                        }
                                        v42 := GenS5{F0: 2, F1: (^(v3))}
                                        _ = fi41
                                        _ = v42
                                    }
                                    var v43 int32 = v4
                                    _ = v43
                                }
                                var v44 uint64 = 4294967295
                                fmt.Println(v2, v3, fi40)
                                _ = fi40
                                _ = v44
                            }
                            var v45 int = 9223372036854775807
                            _ = v39
                            _ = v45
                        }
                    default:
                        if (189 != 392) {
                            v2 = (-(v2))
                            var v46 bool = (!true)
                            var v47 string = ""
                            _ = v46
                            _ = v47
                        }
                    }
                } else {
                    if ("" != "a") {
                        _ = len(v1)
                        for fi48 := 0; fi48 < 4; fi48++ {
                            fmt.Println(fi48)
                            for fi49 := 0; fi49 < 9; fi49++ {
                                if (true || true) {
                                    if true {
                                        v4 = (v4 & -474)
                                        var v50 uint32 = 65535
                                        var v51 uint = genFn0()
                                        v52 := GenS2{F0: "foo bar", F1: 65535}
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                    }
                                    v2 += 874
                                } else {
                                    var v53 uint16 = 178
                                    fmt.Println(fi48)
                                    var v54 uint64 = v3
                                    _ = v53
                                    _ = v54
                                }
                                v2 = v2
                                switch (fi49 << ((4294967295) & 63)) {
                                case -9223372036854775808:
                                    if false {
                                        v2 += (v2 + v2)
                                        switch 127 {
                                        case -2147483648:
                                            var v55 int32 = v4
                                            _ = v55
                                        case 32767:
                                            fmt.Println(fi48, v2, fi49)
                                        case 9223372036854775807:
                                            v2 -= v2
                                            v3 = (v3 >> ((255) & 63))
                                        default:
                                            var v56 uint16 = 0
                                            _ = v56
                                        }
                                    } else {
                                        switch fi49 {
                                        case 42:
                                            fmt.Println(fi48, v3, fi48)
                                        case 65:
                                            var v57 uint16 = 65535
                                            var v58 uint32 = uint32(v3)
                                            if true {
                                                v4 -= 934
                                            }
                                            var v59 int8 = 54
                                            v1[2147483647] = "\n\t"
                                            _ = v57
                                            _ = v58
                                            _ = v59
                                        default:
                                            v60 := GenS3{F0: ("" + (("\n\t" + ("foo bar" + ("hello" + "foo bar"))) + "")), F1: 128}
                                            _ = v60
                                        }
                                        switch (((("\n\t" + "hello") + "\n\t") + "hello") + ("hello" + "\n\t")) {
                                        case "a":
                                            v2 = (v2 * (-(v2)))
                                        }
                                    }
                                case -152:
                                    v61 := GenS6{F0: -468, F1: 42}
                                    _ = v61
                                case 2147483647:
                                    v62 := []int{}
                                    _ = v62
                                default:
                                    v63 := GenS7{F0: 42}
                                    _ = v63
                                }
                                _ = fi49
                            }
                            fmt.Println(v4)
                            _ = fi48
                        }
                        if true {
                            var v64 int = 2147483647
                            _ = v64
                        }
                    }
                    var v65 float64 = 0.0
                    if false {
                        for fi66 := 0; fi66 < 20; fi66++ {
                            var v67 int32 = (v4 * (v4 - (-(v4))))
                            v68 := GenS8{F0: 206, F1: 255}
                            for fi69 := 0; fi69 < 8; fi69++ {
                                switch "a" {
                                case "":
                                    v70 := GenS6{F0: -1, F1: 1}
                                    _ = v70
                                }
                                var v71 uint16 = 255
                                var v72 uint16 = v71
                                _ = fi69
                                _ = v71
                                _ = v72
                            }
                            _ = fi66
                            _ = v67
                            _ = v68
                        }
                        var v73 int32 = v4
                        var v74 float64 = (v65 - v65)
                        _ = v73
                        _ = v74
                    } else {
                        v75 := []int{4294967295, 128, 0}
                        v75 = append(v75, -735)
                        _ = v75
                    }
                    _ = v65
                }
            }
            _ = v29
        } else {
            var v76 float64 = (1.5 - 0.1)
            fmt.Println(v2)
            _ = v76
        }
    default:
        v77 := []int{}
        _ = v77
    }
    v3 = (v3 / ((v3) | 1))
    if false {
        for fi78 := 0; fi78 < 3; fi78++ {
            v4 -= -296
            switch (("foo bar" + ("a" + ("\n\t" + ""))) + "hello") {
            case "foo bar":
                switch (fi78 - (fi78 ^ (fi78 % ((fi78) | 1)))) {
                case -2147483648:
                    var v79 uint32 = 0
                    _ = v79
                case 128:
                    v80 := GenS8{F0: 18446744073709551615, F1: 1}
                    _ = v80
                case -929:
                    if ((0 == 255) && true) {
                        if (genFn0() <= 255) {
                            var v81 uint8 = 236
                            for fi82 := 0; fi82 < 6; fi82++ {
                                v81 += 119
                                _ = fi82
                            }
                            var v83 int64 = 1
                            _ = v81
                            _ = v83
                        }
                    }
                }
            case "\n\t":
                v4 -= 0
            }
            _ = fi78
        }
        if true {
            if (true && (1.5 >= ((float32(0.1) - (0.1 + (float32(1e10) / (float32((((-2.25 + ((1e10 - float64(((1e10 / (0.0 + ((float32(float64(1.5)) * 1.5) - (0.1 + 0.0)))) + 1.5))) / 3.14159)) + (-2.25 * (-2.25 / 0.1))) - -2.25)) - 1e10)))) / float32((1e10 * 0.1))))) {
                var v84 uint32 = 197
                if ((true || true) || true) {
                    if (32768 <= 381) {
                        for fi85 := 0; fi85 < 20; fi85++ {
                            var v86 uint16 = 32767
                            v87 := []int{(fi85 & fi85), 4294967295}
                            _ = fi85
                            _ = v86
                            _ = v87
                        }
                        v88 := []int{384, 2147483647, 0, 0}
                        _ = v88
                    }
                }
                _ = v84
            }
            var v89 uint64 = 42
            _ = v89
        }
    }
    fmt.Println(v3)
    v90 := GenS8{F0: 18446744073709551615, F1: 65535}
    for fi91 := 0; fi91 < 19; fi91++ {
        v92 := []int{}
        fmt.Println(v4, v3, v3)
        _ = fi91
        _ = v92
    }
    v93 := []int{-939, 830}
    v4 += (v4 * v4)
    for fi94 := 0; fi94 < 1; fi94++ {
        if (1 < 127) {
            var v95 float32 = -2.25
            _ = v95
        } else {
            switch ("a" + "a") {
            case "foo bar":
                if (v4 > v4) {
                    for fi96 := 0; fi96 < 15; fi96++ {
                        v2 = 1
                        v97 := GenS7{F0: 1}
                        var v98 bool = true
                        var v99 float32 = 1e10
                        v99 += v99
                        _ = fi96
                        _ = v97
                        _ = v98
                        _ = v99
                    }
                    v2 -= v2
                    if false {
                        var v100 int64 = 42
                        _ = v100
                    }
                }
            default:
                fmt.Println(fi94, fi94, fi94)
            }
        }
        _ = fi94
    }
    var v101 int = 1
    switch "\n\t" {
    case "hello":
        switch v101 {
        case 32767:
            switch ("a" + ("foo bar" + ((("hello" + "\n\t") + "a") + ""))) {
            case "hello":
                v102 := GenS1{F0: v101, F1: 42, F2: 2}
                _ = v102
            }
        case 655:
            var v103 uint64 = 840
            _ = v103
        case -9223372036854775808:
            var v104 int8 = 1
            _ = v104
        }
    case "foo bar":
        for fi105 := 0; fi105 < 14; fi105++ {
            fmt.Println(v4, fi105, v4)
            for fi106 := 0; fi106 < 3; fi106++ {
                if (!false) {
                    v101 = (v101 % ((-821) | 1))
                }
                v107 := GenS8{F0: 65535, F1: 799}
                if ((v2 & (v2 % ((-32768) | 1))) > -868) {
                    v3 = v3
                    if false {
                        if (false && (150 >= 0)) {
                            v4 += 2147483647
                            v2 -= v2
                            switch "\n\t" {
                            case "\n\t":
                                var v108 int8 = -1
                                _ = v108
                            case "a":
                                v1[2147483647] = ("\n\t" + "\n\t")
                            default:
                                var v109 float32 = ((1.5 + -2.25) * 3.14159)
                                v93 = append(v93, (fi106 + 4294967295))
                                var v110 uint = 0
                                _ = v109
                                _ = v110
                            }
                        } else {
                            v111 := GenS9{F0: 255}
                            _ = v111
                        }
                        var v112 bool = (!true)
                        var v113 uint64 = (v3 & (v3 & v3))
                        _ = v112
                        _ = v113
                    }
                } else {
                    v101 = fi106
                    if true {
                        switch (v101 ^ (v101 + 746)) {
                        case 9223372036854775807:
                            var v114 int32 = (v4 % (((^(v4))) | 1))
                            _ = v114
                        case 2147483647:
                            fmt.Println(v4)
                        case -2147483648:
                            if len(v93) >= 2 {
                                v93 = v93[1:len(v93)]
                            }
                        }
                        v4 += 42
                        _ = len(v93)
                        v101 = 128
                        if (7 <= -76) {
                            if (true && true) {
                                switch ("foo bar" + ("hello" + "hello")) {
                                case "foo bar":
                                    v115 := GenS1{F0: (fi105 & (fi105 ^ (fi105 & (fi105 / (((^(v101))) | 1))))), F1: v2, F2: 18}
                                    _ = v115
                                }
                                var v116 uint16 = 32767
                                _ = v116
                            }
                            if len(v93) > 0 {
                                v93[0] = -9223372036854775808
                            }
                        } else {
                            v2 = (-(v2))
                            fmt.Println(fi106, v4, v3)
                            var v117 uint16 = 32767
                            if true {
                                var v118 uint = genFn0()
                                var v119 uint16 = v117
                                fmt.Println(v119)
                                _ = v118
                                _ = v119
                            }
                            switch "\n\t" {
                            case "\n\t":
                                if true {
                                    if false {
                                        fmt.Println(fi105)
                                        var v120 uint64 = v3
                                        var v121 uint16 = (v117 >> ((489) & 63))
                                        var v122 uint8 = 128
                                        _ = v120
                                        _ = v121
                                        _ = v122
                                    }
                                    if (914 != 42) {
                                        var v123 float32 = 1e10
                                        switch ("\n\t" + "foo bar") {
                                        case "a":
                                            if true {
                                                v124 := GenS9{F0: 218}
                                                var v125 uint8 = 0
                                                v126 := GenS9{F0: 255}
                                                _ = v124
                                                _ = v125
                                                _ = v126
                                            }
                                        case "\n\t":
                                            v117 -= uint16(v3)
                                        }
                                        if len(v93) >= 2 {
                                            v93 = v93[1:len(v93)]
                                        }
                                        _ = v123
                                    }
                                } else {
                                    _ = len(v1)
                                    var v127 string = "a"
                                    v2 = (v2 ^ v2)
                                    _ = v127
                                }
                            case "hello":
                                if false {
                                    fmt.Println(fi105)
                                    v128 := GenS10{F0: (^(v4))}
                                    _ = v128
                                } else {
                                    fmt.Println(v3, fi106)
                                    v4 = (v4 >> ((1) & 63))
                                    var v129 int32 = (v4 / (((v4 >> ((uint(v117)) & 63))) | 1))
                                    _ = v129
                                }
                            }
                            _ = v117
                        }
                        var v130 int8 = 0
                        _ = v130
                    } else {
                        var v131 uint8 = 223
                        v1[329] = ((("" + ("hello" + "a")) + (("foo bar" + "") + (("foo bar" + "foo bar") + "foo bar"))) + "hello")
                        var v132 uint64 = v3
                        _ = v131
                        _ = v132
                    }
                    v93 = append(v93, (-(v101)))
                }
                _ = fi106
                _ = v107
            }
            _ = fi105
        }
    case "":
        var v133 uint = 1
        _ = v133
    }
    if ((1.5 + 1.5) < 0.0) {
        if true {
            switch (-(v101)) {
            case -1:
                v101 = 2147483647
            case -54:
                var v134 uint8 = 1
                _ = v134
            }
            if len(v93) > 0 {
                v93[0] = (v101 << ((18446744073709551615) & 63))
            }
        } else {
            switch "foo bar" {
            case "foo bar":
                v2 -= v2
            default:
                _ = len(v1)
            }
            v135 := GenS9{F0: 18446744073709551615}
            v101 -= v101
            _ = v135
        }
    }
    v2 = int16(v3)
    v101 -= (v101 / ((v101) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v93:", len(v93), v93)
    _ = v1
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v90
    _ = v93
}
