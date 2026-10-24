package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 int32; F2 uint8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int64; F1 int32 }
type GenS11 struct { F0 bool; F1 int64 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint8; F1 uint32; F2 int32 }
func genFn0() bool {
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 uint16 = 0
        if false {
            v6 = (v6 << ((42) & 63))
            if (-1 > int32(fi5)) {
                if false {
                    fmt.Println(fi5)
                    v7 := GenS1{F0: (v6 % (((v6 ^ (v6 * 32767))) | 1)), F1: -2147483648}
                    if false {
                        v8 := []int{fi5, (-(fi5)), (fi5 >> ((0) & 63))}
                        if len(v8) > 0 {
                            v8[0] = 427
                        }
                        v9 := GenS2{F0: 331}
                        _ = v8
                        _ = v9
                    } else {
                        var v10 int8 = -128
                        _ = v10
                    }
                    _ = v7
                }
                var v11 uint32 = 4294967295
                if ((^(v11)) != (v11 / ((v11) | 1))) {
                    if true {
                        v12 := GenS3{F0: -858, F1: (v6 + (v6 + v6))}
                        _ = v12
                    } else {
                        var v13 int32 = -2147483648
                        v13 = (v13 << ((566) & 63))
                        _ = v13
                    }
                    var v14 int8 = 1
                    _ = v14
                }
                _ = v11
            }
            if (true && (!false)) {
                switch "" {
                case "foo bar":
                    var v15 int32 = -1
                    _ = v15
                }
                switch fi5 {
                case -1:
                    v16 := GenS1{F0: 255, F1: 0}
                    _ = v16
                case -105:
                    fmt.Println(fi5, fi5)
                default:
                    var v17 uint64 = 0
                    _ = v17
                }
                for fi18 := 0; fi18 < 7; fi18++ {
                    if (true && true) {
                        var v19 int16 = 32767
                        var v20 uint32 = 1
                        v21 := GenS1{F0: v6, F1: 42}
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    v6 = v6
                    var v22 int16 = 1000
                    switch "hello" {
                    case "hello":
                        v23 := GenS2{F0: -1}
                        _ = v23
                    }
                    var v24 int32 = -1
                    var v25 bool = false
                    var v26 float64 = (1.5 - (float64(1e10) * 0.0))
                    if (0.0 < (v26 - v26)) {
                        if (!true) {
                            v27 := GenS4{F0: v25}
                            v28 := []int{fi5, 127, 9223372036854775807, fi5}
                            v24 = v24
                            _ = v27
                            _ = v28
                        } else {
                            switch "foo bar" {
                            case "":
                                var v29 uint16 = 65535
                                v30 := GenS1{F0: 33, F1: (v24 >> ((65535) & 63))}
                                _ = v29
                                _ = v30
                            case "hello":
                                switch fi18 {
                                case 32768:
                                    v31 := GenS2{F0: (v24 + (v24 | 42))}
                                    _ = v31
                                case 0:
                                    v25 = (-9223372036854775808 >= -1)
                                case -9223372036854775808:
                                    v24 = (v24 + 1)
                                }
                            case "\n\t":
                                if v25 {
                                    switch ("hello" + "hello") {
                                    case "hello":
                                        v26 += (1e10 + 0.1)
                                    case "foo bar":
                                        var v32 uint64 = 320
                                        _ = v32
                                    default:
                                        if (v22 < (v22 - (v22 + (v22 % (((v22 << ((291) & 63))) | 1))))) {
                                            if (!v25) {
                                                v33 := GenS1{F0: (^(v6)), F1: v24}
                                                var v34 int32 = (v24 / ((-1) | 1))
                                                _ = v33
                                                _ = v34
                                            }
                                            if (!(!(917 >= 42))) {
                                                if v25 {
                                                    var v35 int8 = 127
                                                    _ = v35
                                                }
                                            } else {
                                                if v25 {
                                                    v22 = 788
                                                    v36 := []int{fi18, fi5, 42, (fi5 % ((4294967295) | 1))}
                                                    _ = v36
                                                }
                                                v37 := GenS3{F0: fi5, F1: 516}
                                                _ = v37
                                            }
                                            var v38 float32 = 0.1
                                            _ = v38
                                        }
                                    }
                                    if true {
                                        v39 := GenS4{F0: (42 <= 0)}
                                        _ = v39
                                    }
                                    fmt.Println(fi18, v6)
                                    var v40 int16 = (v22 & (-(v22)))
                                    _ = v40
                                }
                            }
                            v6 += (v6 / ((v6) | 1))
                            v41 := GenS1{F0: v6, F1: (v24 % ((2147483647) | 1))}
                            _ = v41
                        }
                        var v42 uint8 = 112
                        _ = v42
                    } else {
                        v43 := GenS4{F0: (((!(!v25)) && v25) || (2 > 0))}
                        _ = v43
                    }
                    v26 -= 0.0
                    fmt.Println(v6, v26)
                    v44 := []int{-792}
                    v26 = 0.1
                    _ = fi18
                    _ = v22
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v44
                }
            } else {
                v45 := GenS5{F0: 0, F1: -2147483648, F2: 127}
                if false {
                    fmt.Println(v6)
                    v46 := GenS6{F0: 828}
                    _ = v46
                }
                _ = v45
            }
        }
        v47 := GenS4{F0: false}
        _ = fi5
        _ = v47
        _ = v6
    }
    for fi48 := 0; fi48 < 2; fi48++ {
        switch "hello" {
        case "hello":
            v49 := []int{(fi48 - (fi48 >> ((662) & 63)))}
            _ = v49
        case "":
            if ((2147483648 < 639) || false) {
                fmt.Println(fi48)
                var v50 float32 = 0.0
                _ = v50
            } else {
                fmt.Println(fi48)
            }
        default:
            if ((!(127 <= 128)) && (4294967295 == 844)) {
                var v51 int64 = 42
                _ = v51
            }
        }
        _ = fi48
    }
    return (18446744073709551615 < 1)
}
func genFn1() uint8 {
    v52 := []int{32767}
    var v53 string = ("\n\t" + "")
    v53 = "hello"
    fmt.Println(v53)
    if false {
        switch -434 {
        case 2147483647:
            v53 = v53
        case -2147483648:
            var v54 uint64 = 0
            _ = v54
        default:
            if true {
                if false {
                    var v55 uint16 = 331
                    var v56 int = 1
                    for fi57 := 0; fi57 < 10; fi57++ {
                        fmt.Println(v56, v55, v56)
                        _ = fi57
                    }
                    _ = v55
                    _ = v56
                } else {
                    switch (("" + v53) + "") {
                    case "a":
                        if (1 > 197) {
                            if (float32(((3.14159 + float64(float32(-2.25))) - (1e10 - 0.0))) != float32(-2.25)) {
                                var v58 float32 = 0.0
                                for fi59 := 0; fi59 < 15; fi59++ {
                                    v52 = append(v52, (fi59 - (fi59 * (fi59 - fi59))))
                                    v58 -= (((((1e10 * 3.14159) + 0.1) - ((float32(float64((float32(3.14159) + 0.1))) * v58) / (-2.25 / 1.0))) + (0.1 + (1e10 / (float32(3.14159) / float32(3.14159))))) - v58)
                                    _ = fi59
                                }
                                fmt.Println(v58, v58)
                                _ = v58
                            }
                            v53 = ""
                            if ("foo bar" == "foo bar") {
                                fmt.Println(v53)
                                var v60 uint = 351
                                _ = v60
                            } else {
                                for fi61 := 0; fi61 < 18; fi61++ {
                                    if true {
                                        v53 = v53
                                    } else {
                                        if (2147483648 < 998) {
                                            if false {
                                                for fi62 := 0; fi62 < 20; fi62++ {
                                                    var v63 int = -462
                                                    v64 := GenS3{F0: v63, F1: 1}
                                                    _ = fi62
                                                    _ = v63
                                                    _ = v64
                                                }
                                                if false {
                                                    for fi65 := 0; fi65 < 6; fi65++ {
                                                        if true {
                                                            var v66 float64 = 0.0
                                                            if len(v52) > 0 {
                                                                v52[0] = 9223372036854775807
                                                            }
                                                            v67 := GenS2{F0: 0}
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                        v53 = v53
                                                        _ = fi65
                                                    }
                                                    for fi68 := 0; fi68 < 20; fi68++ {
                                                        switch (-(fi68)) {
                                                        case 2147483647:
                                                            var v69 string = ("hello" + (v53 + "\n\t"))
                                                            _ = v69
                                                        case -9223372036854775808:
                                                            fmt.Println(fi68)
                                                        default:
                                                            switch fi61 {
                                                            case -9223372036854775808:
                                                                v52 = append(v52, (fi61 & 2147483647))
                                                            case 32768:
                                                                var v70 int32 = -1
                                                                _ = v70
                                                            }
                                                        }
                                                        v71 := GenS7{F0: "foo bar", F1: fi61}
                                                        _ = fi68
                                                        _ = v71
                                                    }
                                                }
                                            }
                                            for fi72 := 0; fi72 < 13; fi72++ {
                                                fmt.Println(fi61)
                                                var v73 bool = (-934 <= 2147483647)
                                                v52 = append(v52, -305)
                                                if (!(65535 != 2)) {
                                                    fmt.Println(fi72, fi72)
                                                    if v73 {
                                                        v53 = v53
                                                        v52 = append(v52, -546)
                                                        var v74 uint16 = 316
                                                        _ = v74
                                                    }
                                                    fmt.Println(v53, fi72, fi61)
                                                }
                                                _ = fi72
                                                _ = v73
                                            }
                                        } else {
                                            var v75 bool = false
                                            v75 = false
                                            v53 = ""
                                            _ = v75
                                        }
                                    }
                                    if false {
                                        var v76 int32 = 42
                                        if len(v52) > 0 {
                                            v52[0] = fi61
                                        }
                                        if (347 >= 42) {
                                            var v77 uint32 = 676
                                            _ = v77
                                        } else {
                                            v78 := GenS8{F0: 42, F1: 899, F2: -2147483648}
                                            _ = v78
                                        }
                                        _ = v76
                                    }
                                    v79 := GenS3{F0: 32767, F1: 2}
                                    _ = fi61
                                    _ = v79
                                }
                                for fi80 := 0; fi80 < 19; fi80++ {
                                    for fi81 := 0; fi81 < 18; fi81++ {
                                        v52 = append(v52, fi81)
                                        v82 := GenS7{F0: (v53 + (v53 + ("a" + ""))), F1: (fi81 / (((fi81 - 32767)) | 1))}
                                        _ = fi81
                                        _ = v82
                                    }
                                    v83 := GenS6{F0: 4294967295}
                                    _ = fi80
                                    _ = v83
                                }
                            }
                        }
                    case "hello":
                        _ = len(v52)
                    case "foo bar":
                        var v84 uint64 = 18446744073709551615
                        _ = v84
                    default:
                        switch -996 {
                        case -9223372036854775808:
                            var v85 int32 = -1
                            _ = v85
                        case 715:
                            v53 = v53
                        default:
                            v86 := GenS9{F0: -294}
                            _ = v86
                        }
                    }
                    if true {
                        for fi87 := 0; fi87 < 20; fi87++ {
                            v53 = ("\n\t" + (v53 + v53))
                            v88 := GenS9{F0: fi87}
                            switch v53 {
                            case "foo bar":
                                if true {
                                    switch v53 {
                                    case "hello":
                                        for fi89 := 0; fi89 < 15; fi89++ {
                                            if true {
                                                v90 := GenS10{F0: 42, F1: -958}
                                                if false {
                                                    if (96 == 40) {
                                                        if false {
                                                            fmt.Println(v53, v53)
                                                        } else {
                                                            var v91 int16 = 308
                                                            if true {
                                                                var v92 int = (fi87 - (fi87 << ((965) & 63)))
                                                                _ = v92
                                                            } else {
                                                                switch (fi87 | (fi87 - (fi87 % ((fi89) | 1)))) {
                                                                case 545:
                                                                    fmt.Println(fi87)
                                                                case 1:
                                                                    var v93 bool = false
                                                                    _ = v93
                                                                case -53:
                                                                    switch 42 {
                                                                    case 575:
                                                                        var v94 uint = 18446744073709551615
                                                                        _ = v94
                                                                    }
                                                                }
                                                                var v95 uint16 = 0
                                                                if true {
                                                                    var v96 uint = 4294967295
                                                                    var v97 int8 = -39
                                                                    _ = v96
                                                                    _ = v97
                                                                }
                                                                _ = v95
                                                            }
                                                            if false {
                                                                fmt.Println(fi87, fi89)
                                                                if (false && true) {
                                                                    var v98 int = (fi89 << ((283) & 63))
                                                                    fmt.Println(v91, v53, fi87)
                                                                    _ = v98
                                                                }
                                                                var v99 bool = true
                                                                _ = v99
                                                            }
                                                            _ = v91
                                                        }
                                                    }
                                                    v53 = "a"
                                                    var v100 int64 = 9223372036854775807
                                                    _ = v100
                                                } else {
                                                    fmt.Println(v53, v53)
                                                    var v101 string = "a"
                                                    if len(v52) > 0 {
                                                        v52[0] = 660
                                                    }
                                                    if false {
                                                        v53 = v53
                                                        var v102 float64 = ((1e10 - 3.14159) - ((-2.25 + 3.14159) + (0.1 + (0.0 / (float64(((-2.25 / 1e10) - (1.5 + -2.25))) * 1e10)))))
                                                        _ = v102
                                                    }
                                                    _ = v101
                                                }
                                                _ = v90
                                            } else {
                                                var v103 uint = 255
                                                switch (fi89 - 1) {
                                                case 520:
                                                    v103 += (v103 ^ v103)
                                                case 900:
                                                    v53 = v53
                                                case 850:
                                                    var v104 int64 = -9223372036854775808
                                                    _ = v104
                                                }
                                                var v105 bool = (v53 >= ("" + v53))
                                                var v106 uint64 = 18446744073709551615
                                                _ = v103
                                                _ = v105
                                                _ = v106
                                            }
                                            v52 = append(v52, (fi87 + fi87))
                                            v107 := GenS4{F0: (false || (-832 > 1))}
                                            _ = fi89
                                            _ = v107
                                        }
                                    case "a":
                                        if (!((fi87 | (fi87 & fi87)) < fi87)) {
                                            v53 = v53
                                            var v108 uint = 4294967295
                                            _ = len(v52)
                                            _ = v108
                                        }
                                    default:
                                        v53 = "\n\t"
                                    }
                                    var v109 float32 = ((((1e10 - -2.25) - ((1.5 - -2.25) * ((-2.25 * (((-2.25 + -2.25) * -2.25) * (0.0 / ((0.0 - 0.1) + 0.0)))) - 1e10))) - 0.0) * -2.25)
                                    var v110 uint64 = 18446744073709551615
                                    _ = v109
                                    _ = v110
                                }
                            }
                            _ = fi87
                            _ = v88
                        }
                        for fi111 := 0; fi111 < 1; fi111++ {
                            switch (v53 + v53) {
                            case "hello":
                                for fi112 := 0; fi112 < 8; fi112++ {
                                    if (false || false) {
                                        if len(v52) > 0 {
                                            v52[0] = 127
                                        }
                                        if false {
                                            fmt.Println(v53)
                                            fmt.Println(v53, fi112)
                                            var v113 uint32 = 0
                                            _ = v113
                                        } else {
                                            v114 := GenS11{F0: false, F1: 42}
                                            _ = v114
                                        }
                                        var v115 int = fi112
                                        _ = v115
                                    } else {
                                        var v116 int = fi112
                                        if (0 < 255) {
                                            if true {
                                                fmt.Println(v116, fi112, v53)
                                            } else {
                                                _ = len(v52)
                                                var v117 uint64 = 121
                                                fmt.Println(v117, fi112, fi112)
                                                _ = v117
                                            }
                                            v118 := GenS3{F0: (fi111 - (fi111 ^ (v116 ^ v116))), F1: 32767}
                                            var v119 int16 = 351
                                            _ = v118
                                            _ = v119
                                        } else {
                                            fmt.Println(fi112, v53, fi112)
                                            if (1 >= 32767) {
                                                var v120 int32 = 2147483647
                                                _ = v120
                                            }
                                        }
                                        _ = v116
                                    }
                                    _ = fi112
                                }
                            }
                            _ = fi111
                        }
                        v53 = v53
                    }
                }
                var v121 int8 = 100
                _ = len(v52)
                _ = v121
            } else {
                v122 := GenS7{F0: (v53 + ("\n\t" + v53)), F1: 32767}
                _ = v122
            }
        }
        var v123 uint = 1
        _ = v123
    }
    _ = v52
    _ = v53
    return 128
}

func main() {
    var v1 int8 = 67
    var v2 float32 = 3.14159
    v3 := make(map[string]int)
    var v4 uint64 = 1
    v1 -= -65
    var v124 float32 = v2
    var v125 int16 = -723
    for fi126 := 0; fi126 < 17; fi126++ {
        var v127 int32 = -1
        _ = fi126
        _ = v127
    }
    v124 = (float32(1e10) - float32(-2.25))
    var v128 int64 = -376
    if true {
        v129 := GenS2{F0: 0}
        if false {
            if (0 > 255) {
                v130 := GenS8{F0: 255, F1: 132, F2: -2147483648}
                v131 := GenS12{F0: 255}
                _ = v130
                _ = v131
            }
        } else {
            for fi132 := 0; fi132 < 10; fi132++ {
                fmt.Println(fi132)
                var v133 float64 = ((1e10 * 3.14159) - 1e10)
                for fi134 := 0; fi134 < 12; fi134++ {
                    v135 := GenS9{F0: -2147483648}
                    var v136 uint = 255
                    _ = fi134
                    _ = v135
                    _ = v136
                }
                _ = fi132
                _ = v133
            }
        }
        _ = v129
    } else {
        var v137 uint8 = 42
        _ = v137
    }
    v138 := GenS13{F0: 255, F1: 76, F2: 42}
    if (!(0 != 42)) {
        fmt.Println(v2, v1, v4)
    } else {
        if (!true) {
            var v139 uint = 65535
            _ = v139
        } else {
            var v140 int8 = 1
            _ = v140
        }
        v3[""] = 42
        v141 := []int{-1, -503, 1}
        _ = v141
    }
    var v142 int8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v124:", v124)
    fmt.Println("v125:", v125)
    fmt.Println("v128:", v128)
    fmt.Println("v142:", v142)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v124
    _ = v125
    _ = v128
    _ = v138
    _ = v142
    _ = v2
    _ = v3
    _ = v4
}
