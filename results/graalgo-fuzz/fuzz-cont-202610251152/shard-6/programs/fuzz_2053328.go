package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint16; F2 int32 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 uint32; F2 string }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 uint16; F1 int8 }
type GenS8 struct { F0 uint16; F1 uint32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint8; F1 uint }
type GenS12 struct { F0 int64; F1 int16; F2 uint8 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 int8; F1 uint16; F2 uint }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint8; F1 string }
type GenS18 struct { F0 uint; F1 int; F2 uint8 }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 uint8; F1 uint32 }
type GenS21 struct { F0 int32 }
type GenS22 struct { F0 uint }
type GenS23 struct { F0 uint32; F1 uint }
type GenS24 struct { F0 uint8 }
type GenS25 struct { F0 int; F1 int64 }

func main() {
    var v1 int64 = 1
    var v2 float32 = -2.25
    var v3 uint16 = 42
    var v4 uint32 = 1
    switch 598 {
    case 436:
        v5 := GenS1{F0: 2}
        _ = v5
    }
    v3 -= (^(v3))
    for fi6 := 0; fi6 < 8; fi6++ {
        v4 = 232
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 uint16 = 581
            if (("foo bar" > "foo bar") || (2147483647 < -2147483648)) {
                v9 := []int{128}
                _ = v9
            } else {
                switch "a" {
                case "":
                    switch "\n\t" {
                    case "":
                        v10 := GenS1{F0: 65535}
                        _ = v10
                    case "foo bar":
                        if false {
                            v8 += (v3 * 970)
                            var v11 uint64 = 0
                            v1 -= v1
                            _ = v11
                        }
                    }
                }
                v4 += v4
            }
            switch ("hello" + "hello") {
            case "a":
                v12 := []int{(-(fi6)), (fi7 << ((0) & 63)), (fi6 | fi7)}
                _ = v12
            case "hello":
                var v13 int64 = (v1 * v1)
                _ = v13
            case "":
                v14 := []int{(fi6 + (fi6 % ((fi6) | 1))), (fi6 % (((-(fi6))) | 1)), 32767}
                _ = v14
            }
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    switch 32768 {
    case 4294967295:
        v2 = v2
    case 42:
        var v15 uint8 = 116
        _ = v15
    }
    v16 := GenS1{F0: 18}
    var v17 int8 = -64
    var v18 string = ("a" + "a")
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 string = v18
        _ = fi19
        _ = v20
    }
    var v21 float64 = 3.14159
    fmt.Println(v4, v21)
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := GenS2{F0: 1, F1: v3, F2: 884}
        v18 = ((v18 + v18) + (("foo bar" + (("" + v18) + v18)) + v18))
        var v24 int16 = -32768
        _ = fi22
        _ = v23
        _ = v24
    }
    if ("hello" == v18) {
        switch v18 {
        case "":
            var v25 uint16 = 42
            _ = v25
        case "\n\t":
            var v26 float64 = v21
            _ = v26
        case "hello":
            var v27 uint16 = uint16(v4)
            _ = v27
        }
        fmt.Println(v1, v1)
        for fi28 := 0; fi28 < 18; fi28++ {
            v29 := []int{2147483647}
            v18 = ("foo bar" + "\n\t")
            switch (fi28 % (((fi28 / ((983) | 1))) | 1)) {
            case 142:
                var v30 uint8 = 127
                _ = v30
            case -2147483648:
                v3 -= 65535
            default:
                for fi31 := 0; fi31 < 19; fi31++ {
                    var v32 int8 = (v17 * (^(v17)))
                    _ = fi31
                    _ = v32
                }
            }
            _ = fi28
            _ = v29
        }
    }
    switch v18 {
    case "a":
        v17 = (v17 * (v17 | (v17 % (((^(v17))) | 1))))
    case "":
        var v33 uint8 = 255
        _ = v33
    case "hello":
        if false {
            switch 9223372036854775807 {
            case 42:
                v34 := GenS3{F0: (v18 + ""), F1: (v1 % ((977) | 1))}
                _ = v34
            case 32767:
                if true {
                    if true {
                        if ((true || true) || false) {
                            var v35 float64 = v21
                            _ = v35
                        } else {
                            if (v17 <= (v17 << ((65535) & 63))) {
                                if false {
                                    for fi36 := 0; fi36 < 1; fi36++ {
                                        for fi37 := 0; fi37 < 19; fi37++ {
                                            fmt.Println(fi37, v17)
                                            switch (fi37 - fi36) {
                                            case 2147483647:
                                                if (289 >= 1) {
                                                    fmt.Println(v1, fi37)
                                                } else {
                                                    var v38 uint16 = v3
                                                    v39 := GenS4{F0: (!((^(v4)) < (v4 << ((65535) & 63))))}
                                                    _ = v38
                                                    _ = v39
                                                }
                                            }
                                            fmt.Println(v2, v2)
                                            v40 := GenS5{F0: 9223372036854775807, F1: (v4 % ((v4) | 1)), F2: v18}
                                            _ = fi37
                                            _ = v40
                                        }
                                        v1 = (v1 + 1)
                                        v41 := GenS6{F0: 255, F1: v18}
                                        _ = fi36
                                        _ = v41
                                    }
                                    switch v18 {
                                    case "hello":
                                        v18 = v18
                                    case "\n\t":
                                        for fi42 := 0; fi42 < 6; fi42++ {
                                            for fi43 := 0; fi43 < 20; fi43++ {
                                                v17 = (v17 ^ 1)
                                                _ = fi43
                                            }
                                            for fi44 := 0; fi44 < 4; fi44++ {
                                                fmt.Println(v1)
                                                _ = fi44
                                            }
                                            _ = fi42
                                        }
                                    }
                                } else {
                                    if false {
                                        v18 = "\n\t"
                                        v45 := GenS2{F0: 176, F1: v3, F2: -334}
                                        v2 = (v2 * 1e10)
                                        _ = v45
                                    } else {
                                        var v46 int32 = -2147483648
                                        v47 := GenS7{F0: (v3 / (((v3 >> ((0) & 63))) | 1)), F1: (v17 % (((v17 << ((18446744073709551615) & 63))) | 1))}
                                        _ = v46
                                        _ = v47
                                    }
                                    v1 = (v1 + -1)
                                }
                                v48 := []int{-9223372036854775808, 4294967295, 2147483647}
                                for fi49 := 0; fi49 < 3; fi49++ {
                                    v48 = append(v48, (fi49 / ((int(v1)) | 1)))
                                    v4 -= (v4 & v4)
                                    _ = fi49
                                }
                                _ = v48
                            }
                            v50 := GenS6{F0: 367, F1: "\n\t"}
                            _ = v50
                        }
                    } else {
                        if (!true) {
                            if false {
                                var v51 int32 = 1
                                _ = v51
                            } else {
                                if false {
                                    if false {
                                        if ((int64(v17) > (v1 & v1)) && true) {
                                            switch -9223372036854775808 {
                                            case -688:
                                                var v52 int64 = 1
                                                _ = v52
                                            case -2147483648:
                                                v53 := []int{988, -9223372036854775808}
                                                _ = v53
                                            case 4294967295:
                                                v17 += (v17 * (v17 % (((v17 - v17)) | 1)))
                                            default:
                                                v54 := []int{4294967295, 42, 2147483647, -644}
                                                _ = v54
                                            }
                                            v55 := GenS8{F0: (v3 & 65535), F1: 42}
                                            v2 = 3.14159
                                            _ = v55
                                        }
                                        v56 := GenS9{F0: -2147483648}
                                        _ = v56
                                    }
                                    v3 = (v3 % ((v3) | 1))
                                } else {
                                    for fi57 := 0; fi57 < 6; fi57++ {
                                        var v58 int16 = 32767
                                        for fi59 := 0; fi59 < 19; fi59++ {
                                            if true {
                                                v60 := GenS6{F0: uint(v3), F1: ("\n\t" + v18)}
                                                _ = v60
                                            }
                                            var v61 float32 = ((((0.1 * float32(1.5)) * (v2 - 0.0)) * v2) + 3.14159)
                                            _ = fi59
                                            _ = v61
                                        }
                                        if (!((true || false) && false)) {
                                            v18 = "foo bar"
                                            var v62 int64 = 42
                                            fmt.Println(v21)
                                            _ = v62
                                        }
                                        _ = fi57
                                        _ = v58
                                    }
                                    switch ("hello" + v18) {
                                    case "":
                                        if ((v17 ^ v17) > 100) {
                                            v17 = v17
                                            for fi63 := 0; fi63 < 9; fi63++ {
                                                v64 := []int{(fi63 & (fi63 / ((686) | 1))), int(v3), 0}
                                                if false {
                                                    if true {
                                                        if true {
                                                            var v65 bool = (9223372036854775807 <= v1)
                                                            v4 -= (v4 * (v4 & uint32(v3)))
                                                            fmt.Println(v1, v1)
                                                            _ = v65
                                                        }
                                                        v64 = append(v64, -9223372036854775808)
                                                    } else {
                                                        var v66 float32 = v2
                                                        v64 = append(v64, 406)
                                                        _ = v66
                                                    }
                                                    if true {
                                                        var v67 uint32 = (v4 - 65535)
                                                        _ = v67
                                                    } else {
                                                        var v68 bool = true
                                                        _ = v68
                                                    }
                                                    var v69 uint32 = (v4 * v4)
                                                    _ = v69
                                                } else {
                                                    var v70 uint64 = 693
                                                    v17 += 1
                                                    if false {
                                                        v21 += 0.1
                                                        switch ((v18 + v18) + "\n\t") {
                                                        case "foo bar":
                                                            for fi71 := 0; fi71 < 6; fi71++ {
                                                                var v72 uint = 999
                                                                var v73 int32 = 0
                                                                _ = fi71
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                        case "a":
                                                            switch fi63 {
                                                            case 9223372036854775807:
                                                                v4 = v4
                                                            case -9223372036854775808:
                                                                switch fi63 {
                                                                case -182:
                                                                    v3 -= 42
                                                                case 2147483647:
                                                                    if ((v17 % (((v17 & v17)) | 1)) >= (v17 & v17)) {
                                                                        v74 := GenS1{F0: 2}
                                                                        _ = v74
                                                                    }
                                                                default:
                                                                    v75 := GenS10{F0: v70}
                                                                    _ = v75
                                                                }
                                                            default:
                                                                v76 := GenS3{F0: v18, F1: (v1 % (((v1 - int64(v70))) | 1))}
                                                                _ = v76
                                                            }
                                                        case "":
                                                            v21 = v21
                                                        default:
                                                            switch "a" {
                                                            case "":
                                                                v21 += 0.0
                                                            }
                                                        }
                                                        if true {
                                                            v4 = (v4 & v4)
                                                            var v77 uint16 = 42
                                                            _ = v77
                                                        }
                                                    }
                                                    _ = v70
                                                }
                                                v18 = (v18 + (v18 + "\n\t"))
                                                _ = fi63
                                                _ = v64
                                            }
                                            v3 = (v3 * 42)
                                        } else {
                                            fmt.Println(v21)
                                            if ((((^(v3)) >= v3) || (88 != 1)) || (2 != 1)) {
                                                var v78 float32 = v2
                                                v78 -= (v78 / 1.0)
                                                _ = v78
                                            } else {
                                                fmt.Println(v18, v21, v1)
                                                if false {
                                                    v18 = "hello"
                                                    if false {
                                                        if (!(4294967295 <= 2)) {
                                                            for fi79 := 0; fi79 < 1; fi79++ {
                                                                for fi80 := 0; fi80 < 14; fi80++ {
                                                                    v81 := []int{}
                                                                    v2 += v2
                                                                    _ = fi80
                                                                    _ = v81
                                                                }
                                                                var v82 string = "foo bar"
                                                                _ = fi79
                                                                _ = v82
                                                            }
                                                            v83 := []int{1, 32768}
                                                            v84 := GenS8{F0: 733, F1: 65535}
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                        if (v21 == v21) {
                                                            var v85 int16 = 32767
                                                            _ = v85
                                                        } else {
                                                            for fi86 := 0; fi86 < 10; fi86++ {
                                                                v87 := GenS3{F0: v18, F1: 0}
                                                                _ = fi86
                                                                _ = v87
                                                            }
                                                            for fi88 := 0; fi88 < 1; fi88++ {
                                                                var v89 float64 = (0.1 / (v21 - v21))
                                                                v18 = ("" + ((v18 + (v18 + ((v18 + v18) + v18))) + (v18 + ("" + ("\n\t" + "\n\t")))))
                                                                v89 += v89
                                                                _ = fi88
                                                                _ = v89
                                                            }
                                                            v1 += -1
                                                        }
                                                    } else {
                                                        fmt.Println(v4)
                                                        v90 := []int{-815, -942}
                                                        if false {
                                                            v91 := GenS11{F0: 128, F1: 18446744073709551615}
                                                            if len(v90) >= 2 {
                                                                v90 = v90[1:len(v90)]
                                                            }
                                                            v4 -= (v4 - v4)
                                                            _ = v91
                                                        }
                                                        _ = v90
                                                    }
                                                }
                                                v1 -= 558
                                            }
                                        }
                                    case "a":
                                        switch ("a" + (v18 + ((v18 + "a") + "hello"))) {
                                        case "a":
                                            if false {
                                                v4 = (v4 & (v4 + (v4 | v4)))
                                            }
                                        case "hello":
                                            fmt.Println(v3, v17, v17)
                                        }
                                    case "foo bar":
                                        v3 += (v3 | 32767)
                                    }
                                    for fi92 := 0; fi92 < 12; fi92++ {
                                        var v93 uint32 = (^(v4))
                                        if false {
                                            var v94 uint8 = 1
                                            var v95 uint16 = 42
                                            var v96 int32 = 207
                                            _ = v94
                                            _ = v95
                                            _ = v96
                                        }
                                        _ = fi92
                                        _ = v93
                                    }
                                }
                                if (659 >= (^(v1))) {
                                    if true {
                                        v3 = (v3 % (((v3 / (((^(v3))) | 1))) | 1))
                                        var v97 int = -9223372036854775808
                                        v98 := []int{int(v17), v97, 9223372036854775807}
                                        _ = v97
                                        _ = v98
                                    } else {
                                        v99 := []int{-954}
                                        _ = v99
                                    }
                                }
                                if false {
                                    v100 := []int{}
                                    _ = v100
                                } else {
                                    v101 := GenS12{F0: v1, F1: int16(v1), F2: 128}
                                    var v102 uint = 2
                                    v103 := GenS13{F0: uint16(v4)}
                                    _ = v101
                                    _ = v102
                                    _ = v103
                                }
                            }
                            var v104 uint8 = 255
                            _ = v104
                        }
                    }
                    v105 := []int{-1, 9223372036854775807}
                    switch (v18 + v18) {
                    case "a":
                        for fi106 := 0; fi106 < 15; fi106++ {
                            if true {
                                for fi107 := 0; fi107 < 15; fi107++ {
                                    if true {
                                        if ((v3 + (v3 % ((v3) | 1))) <= v3) {
                                            if false {
                                                switch v18 {
                                                case "\n\t":
                                                    v3 = (v3 ^ v3)
                                                case "":
                                                    v108 := GenS2{F0: 2147483647, F1: v3, F2: -220}
                                                    var v109 bool = (true || false)
                                                    if (true || ((fi106 - fi107) <= (-(fi107)))) {
                                                        var v110 uint8 = 42
                                                        _ = v110
                                                    }
                                                    _ = v108
                                                    _ = v109
                                                default:
                                                    if false {
                                                        if false {
                                                            v3 -= v3
                                                        } else {
                                                            var v111 uint8 = 1
                                                            if (1 <= 896) {
                                                                if true {
                                                                    if true {
                                                                        switch 232 {
                                                                        case -2147483648:
                                                                            if true {
                                                                                var v112 int8 = (v17 | v17)
                                                                                _ = v112
                                                                            } else {
                                                                                fmt.Println(fi107, v21)
                                                                                v3 = (v3 + 65535)
                                                                                if (v21 == float64((v2 + 0.0))) {
                                                                                    var v113 bool = false
                                                                                    var v114 string = ((("\n\t" + ((((v18 + (((v18 + "a") + (v18 + "a")) + (("a" + v18) + ((((v18 + v18) + v18) + (v18 + (v18 + "hello"))) + ("" + v18))))) + "a") + (v18 + ((("\n\t" + ((v18 + v18) + v18)) + v18) + ""))) + v18)) + v18) + "")
                                                                                    if v113 {
                                                                                        switch v18 {
                                                                                        case "\n\t":
                                                                                            switch ("" + ("" + (("a" + v18) + v114))) {
                                                                                            case "a":
                                                                                                v115 := GenS6{F0: 565, F1: v18}
                                                                                                _ = v115
                                                                                            case "hello":
                                                                                                var v116 uint64 = 1
                                                                                                _ = v116
                                                                                            }
                                                                                        default:
                                                                                            var v117 int64 = v1
                                                                                            _ = v117
                                                                                        }
                                                                                    } else {
                                                                                        if len(v105) >= 2 {
                                                                                            v105 = v105[1:len(v105)]
                                                                                        }
                                                                                    }
                                                                                    if (true || (!true)) {
                                                                                        fmt.Println(v2, v3)
                                                                                        fmt.Println(v1)
                                                                                    } else {
                                                                                        if v113 {
                                                                                            var v118 float64 = v21
                                                                                            v119 := GenS14{F0: v17}
                                                                                            _ = v118
                                                                                            _ = v119
                                                                                        } else {
                                                                                            fmt.Println(v1, fi107, v114)
                                                                                            var v120 int = (fi107 - fi107)
                                                                                            switch v18 {
                                                                                            case "\n\t":
                                                                                                v105 = append(v105, v120)
                                                                                            case "":
                                                                                                fmt.Println(v111, v21)
                                                                                                if len(v105) >= 2 {
                                                                                                    v105 = v105[1:len(v105)]
                                                                                                }
                                                                                                v17 += (v17 / (((v17 - (v17 % ((v17) | 1)))) | 1))
                                                                                            case "foo bar":
                                                                                                if len(v105) >= 2 {
                                                                                                    v105 = v105[1:len(v105)]
                                                                                                }
                                                                                            }
                                                                                            _ = v120
                                                                                        }
                                                                                    }
                                                                                    _ = v113
                                                                                    _ = v114
                                                                                } else {
                                                                                    v21 = v21
                                                                                    v3 = 32767
                                                                                    v121 := GenS7{F0: (v3 | v3), F1: (v17 ^ 100)}
                                                                                    _ = v121
                                                                                }
                                                                            }
                                                                        case 32767:
                                                                            if false {
                                                                                var v122 uint64 = 577
                                                                                if (true || true) {
                                                                                    var v123 string = v18
                                                                                    _ = v123
                                                                                } else {
                                                                                    _ = len(v105)
                                                                                }
                                                                                fmt.Println(v18, v17)
                                                                                _ = v122
                                                                            } else {
                                                                                if (!(-1 >= -32768)) {
                                                                                    v124 := GenS9{F0: 822}
                                                                                    _ = v124
                                                                                }
                                                                            }
                                                                            if true {
                                                                                if false {
                                                                                    v125 := GenS15{F0: (v17 * v17), F1: 1, F2: 0}
                                                                                    v126 := GenS6{F0: 65535, F1: (("a" + ("" + v18)) + v18)}
                                                                                    var v127 float32 = ((float32(v21) * (float32(v21) / (v2 / v2))) * (v2 - (-2.25 * float32(((v21 / 1e10) / v21)))))
                                                                                    v21 = float64((0.0 + float32(v21)))
                                                                                    _ = v125
                                                                                    _ = v126
                                                                                    _ = v127
                                                                                } else {
                                                                                    if (0 != 47) {
                                                                                        v111 = (v111 | v111)
                                                                                        if false {
                                                                                            v3 = (v3 % ((1) | 1))
                                                                                        } else {
                                                                                            fmt.Println(fi106)
                                                                                        }
                                                                                        v1 -= (v1 | int64(v4))
                                                                                    } else {
                                                                                        v17 = 98
                                                                                    }
                                                                                }
                                                                                var v128 string = ((v18 + v18) + v18)
                                                                                fmt.Println(v17)
                                                                                _ = v128
                                                                            }
                                                                        case -1:
                                                                            v3 = 2
                                                                        default:
                                                                            v18 = "\n\t"
                                                                        }
                                                                        v1 = v1
                                                                    } else {
                                                                        v21 += (((v21 * 1.5) - v21) - v21)
                                                                    }
                                                                } else {
                                                                    v1 += 9223372036854775807
                                                                }
                                                            }
                                                            v4 = (v4 ^ v4)
                                                            v105 = append(v105, fi107)
                                                            fmt.Println(v17, fi107)
                                                            fmt.Println(v18, v111)
                                                            _ = v111
                                                        }
                                                        if (2 != 0) {
                                                            v18 = "\n\t"
                                                            if ((false || ((((float64(((((1.5 * -2.25) + 0.1) * float32(1.5)) * (1.5 - float32((v21 - v21))))) / v21) / v21) + v21) <= (float64((v2 + v2)) - (((float64(v2) / 1e10) + float64(3.14159)) + v21)))) && true) {
                                                                var v129 uint32 = v4
                                                                v130 := GenS15{F0: v17, F1: 2, F2: 0}
                                                                v131 := GenS16{F0: 1}
                                                                _ = v129
                                                                _ = v130
                                                                _ = v131
                                                            }
                                                        }
                                                        if (0 <= uint(v3)) {
                                                            v132 := GenS17{F0: 255, F1: v18}
                                                            var v133 float64 = (v21 * 1.5)
                                                            v134 := GenS18{F0: 4294967295, F1: (fi106 ^ fi107), F2: 255}
                                                            var v135 float32 = (1e10 - v2)
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                if false {
                                                                    if len(v105) >= 2 {
                                                                        v105 = v105[1:len(v105)]
                                                                    }
                                                                    fmt.Println(v18, v21)
                                                                    if (v1 < (v1 * (v1 << ((42) & 63)))) {
                                                                        v136 := GenS1{F0: 0}
                                                                        if len(v105) >= 2 {
                                                                            v105 = v105[1:len(v105)]
                                                                        }
                                                                        _ = v136
                                                                    } else {
                                                                        v137 := GenS7{F0: v3, F1: (^(v17))}
                                                                        if false {
                                                                            v21 = v21
                                                                            v18 = v18
                                                                        } else {
                                                                            fmt.Println(v21, fi107)
                                                                        }
                                                                        var v138 int = fi107
                                                                        v1 = 9223372036854775807
                                                                        v135 = (3.14159 / float32(3.14159))
                                                                        _ = v137
                                                                        _ = v138
                                                                    }
                                                                } else {
                                                                    var v139 bool = false
                                                                    _ = len(v105)
                                                                    _ = v139
                                                                }
                                                            case "foo bar":
                                                                var v140 float32 = float32(v21)
                                                                _ = v140
                                                            default:
                                                                var v141 float64 = (1e10 - v21)
                                                                _ = v141
                                                            }
                                                            _ = v132
                                                            _ = v133
                                                            _ = v134
                                                            _ = v135
                                                        } else {
                                                            var v142 uint64 = 2
                                                            switch fi106 {
                                                            case 875:
                                                                fmt.Println(v4, fi106, v3)
                                                            default:
                                                                v143 := GenS12{F0: v1, F1: -32768, F2: 1}
                                                                _ = v143
                                                            }
                                                            v1 -= -811
                                                            _ = v142
                                                        }
                                                    } else {
                                                        if (v18 != ("\n\t" + "")) {
                                                            v144 := GenS9{F0: (fi107 / ((fi107) | 1))}
                                                            fmt.Println(fi106, v17)
                                                            v4 += (v4 | 462)
                                                            _ = v144
                                                        } else {
                                                            var v145 bool = (1 == 76)
                                                            var v146 uint16 = (^(v3))
                                                            var v147 int = 32767
                                                            _ = v145
                                                            _ = v146
                                                            _ = v147
                                                        }
                                                        _ = len(v105)
                                                        var v148 int8 = v17
                                                        _ = v148
                                                    }
                                                }
                                                fmt.Println(v21, v18, v1)
                                                v149 := GenS19{F0: v17}
                                                _ = v149
                                            }
                                            switch (v18 + v18) {
                                            case "a":
                                                v17 += (-(v17))
                                            case "hello":
                                                switch v18 {
                                                case "":
                                                    if (2 <= 255) {
                                                        switch fi106 {
                                                        case -1:
                                                            v1 += int64(v17)
                                                        case 1:
                                                            if false {
                                                                var v150 uint64 = 398
                                                                _ = v150
                                                            } else {
                                                                fmt.Println(v17, v21)
                                                            }
                                                        default:
                                                            var v151 uint16 = (v3 / ((uint16(v4)) | 1))
                                                            _ = v151
                                                        }
                                                        var v152 float32 = v2
                                                        v3 += v3
                                                        if (1 == -1) {
                                                            var v153 uint = 777
                                                            var v154 uint32 = 889
                                                            _ = v153
                                                            _ = v154
                                                        } else {
                                                            v4 = (v4 << ((42) & 63))
                                                            v155 := GenS11{F0: 0, F1: uint(v4)}
                                                            if len(v105) > 0 {
                                                                v105[0] = (^(fi106))
                                                            }
                                                            _ = v155
                                                        }
                                                        v156 := GenS20{F0: 98, F1: v4}
                                                        _ = v152
                                                        _ = v156
                                                    }
                                                case "foo bar":
                                                    v3 = 255
                                                    if false {
                                                        fmt.Println(v2)
                                                    } else {
                                                        v2 = -2.25
                                                    }
                                                    var v157 int64 = (v1 * (v1 | -9223372036854775808))
                                                    _ = v157
                                                default:
                                                    var v158 uint16 = (v3 / (((v3 % ((v3) | 1))) | 1))
                                                    var v159 int32 = int32(v4)
                                                    v160 := GenS3{F0: "a", F1: v1}
                                                    fmt.Println(v21, v1)
                                                    if (0 > 0) {
                                                        if (int64(v158) == int64(v159)) {
                                                            v105 = append(v105, (fi106 % ((-233) | 1)))
                                                            if (v17 >= v17) {
                                                                var v161 float32 = (float32(v21) / 1e10)
                                                                _ = v161
                                                            } else {
                                                                fmt.Println(v159)
                                                                fmt.Println(v18)
                                                            }
                                                            if len(v105) > 0 {
                                                                v105[0] = (fi107 % (((^(fi107))) | 1))
                                                            }
                                                            v17 += (v17 ^ (v17 >> ((255) & 63)))
                                                            var v162 float64 = 1.5
                                                            v158 = v3
                                                            v163 := GenS15{F0: v17, F1: v158, F2: 65535}
                                                            v17 += v17
                                                            _ = v162
                                                            _ = v163
                                                        } else {
                                                            v164 := GenS21{F0: (v159 << ((uint(v158)) & 63))}
                                                            v165 := GenS21{F0: -433}
                                                            switch int(v17) {
                                                            case 498:
                                                                _ = len(v105)
                                                            case 127:
                                                                var v166 int64 = (v1 - v1)
                                                                _ = v166
                                                            }
                                                            _ = v164
                                                            _ = v165
                                                        }
                                                        var v167 int64 = 1
                                                        _ = v167
                                                    } else {
                                                        v168 := GenS17{F0: 127, F1: ""}
                                                        _ = v168
                                                    }
                                                    _ = v158
                                                    _ = v159
                                                    _ = v160
                                                }
                                            case "\n\t":
                                                var v169 string = v18
                                                _ = v169
                                            default:
                                                v170 := GenS5{F0: (v1 * (v1 - (v1 % ((int64(v3)) | 1)))), F1: 4294967295, F2: ("" + v18)}
                                                _ = v170
                                            }
                                            v4 = (v4 | (v4 - 65535))
                                            var v171 int16 = 32767
                                            v172 := GenS7{F0: v3, F1: (v17 >> ((4294967295) & 63))}
                                            _ = v171
                                            _ = v172
                                        }
                                        if false {
                                            var v173 float64 = 0.1
                                            if (!false) {
                                                var v174 int16 = -769
                                                fmt.Println(v18, fi107, v18)
                                                _ = v174
                                            }
                                            _ = v173
                                        }
                                        v175 := GenS19{F0: 127}
                                        _ = v175
                                    }
                                    fmt.Println(v1, fi106)
                                    v4 -= (v4 | 42)
                                    _ = fi107
                                }
                                if false {
                                    switch fi106 {
                                    case 311:
                                        v1 += (v1 % (((-(v1))) | 1))
                                    default:
                                        switch 128 {
                                        case 2147483647:
                                            v176 := GenS22{F0: 18446744073709551615}
                                            _ = v176
                                        }
                                    }
                                    v177 := GenS1{F0: 255}
                                    _ = v177
                                }
                            } else {
                                if true {
                                    fmt.Println(v2)
                                    v178 := GenS17{F0: 127, F1: v18}
                                    for fi179 := 0; fi179 < 3; fi179++ {
                                        var v180 uint16 = (^(v3))
                                        v21 += v21
                                        v4 = v4
                                        fmt.Println(v180, v180, v21)
                                        _ = fi179
                                        _ = v180
                                    }
                                    _ = v178
                                } else {
                                    v1 = (v1 / ((v1) | 1))
                                    var v181 bool = ((v3 & v3) < v3)
                                    _ = v181
                                }
                                if true {
                                    if len(v105) >= 2 {
                                        v105 = v105[1:len(v105)]
                                    }
                                } else {
                                    switch fi106 {
                                    case 127:
                                        switch v18 {
                                        case "foo bar":
                                            v17 = v17
                                        case "":
                                            if ((!false) || true) {
                                                if (true && (!((false || true) && false))) {
                                                    if false {
                                                        var v182 uint16 = (v3 / ((v3) | 1))
                                                        v183 := GenS20{F0: 42, F1: (^(v4))}
                                                        _ = v182
                                                        _ = v183
                                                    } else {
                                                        v17 -= v17
                                                        v184 := GenS17{F0: 43, F1: v18}
                                                        _ = v184
                                                    }
                                                }
                                                var v185 uint32 = 856
                                                fmt.Println(v1)
                                                _ = v185
                                            }
                                        case "hello":
                                            v186 := GenS23{F0: v4, F1: 42}
                                            _ = v186
                                        }
                                    case -536:
                                        if len(v105) >= 2 {
                                            v105 = v105[1:len(v105)]
                                        }
                                    case 1:
                                        var v187 uint8 = 100
                                        _ = v187
                                    }
                                    switch ("a" + ("\n\t" + v18)) {
                                    case "\n\t":
                                        if false {
                                            if false {
                                                var v188 int = (fi106 - (fi106 % ((fi106) | 1)))
                                                v21 = (v21 + (((3.14159 / v21) / ((((float64(0.1) - (0.1 - (float64((1e10 + ((1e10 + (v2 / (3.14159 - 1.5))) - v2))) / 3.14159))) / (((v21 + (v21 - float64((0.1 + (float32(((v21 / ((v21 + ((float64(((v2 + v2) * v2)) * (v21 + float64(0.0))) * ((v21 - 3.14159) + 3.14159))) / float64(v2))) - v21)) / (1.5 * (v2 / 0.1))))))) * (v21 - ((v21 + v21) * 0.0))) * 1.5)) / 1.5) * 1.5)) * v21))
                                                _ = v188
                                            }
                                        } else {
                                            var v189 int64 = 9223372036854775807
                                            var v190 uint8 = 2
                                            v1 = -529
                                            _ = v189
                                            _ = v190
                                        }
                                    default:
                                        var v191 string = ("\n\t" + ((v18 + v18) + ((v18 + "a") + ((v18 + v18) + v18))))
                                        _ = v191
                                    }
                                }
                            }
                            _ = fi106
                        }
                    case "hello":
                        v21 += (0.1 / v21)
                    }
                    _ = v105
                }
            }
            v3 -= v3
        }
    default:
        v192 := GenS24{F0: 127}
        _ = v192
    }
    if (!(v3 == uint16(v4))) {
        v3 -= v3
    }
    if (!false) {
        v193 := GenS12{F0: v1, F1: -1, F2: 0}
        var v194 int64 = v1
        _ = v193
        _ = v194
    } else {
        var v195 uint16 = v3
        v196 := GenS25{F0: 9223372036854775807, F1: 42}
        _ = v195
        _ = v196
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
