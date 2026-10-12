package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint16 }
type GenS2 struct { F0 bool; F1 int8; F2 uint8 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8; F1 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32; F1 int8; F2 string }
type GenS8 struct { F0 uint8; F1 uint16 }
type GenS9 struct { F0 int16; F1 uint32 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int; F1 int64 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int8; F1 uint32 }
type GenS17 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 uint; F1 int64 }
type GenS21 struct { F0 int32; F1 int32; F2 int32 }
type GenS22 struct { F0 uint32; F1 uint }
type GenS23 struct { F0 uint64; F1 uint; F2 uint }
type GenS24 struct { F0 uint32; F1 int16; F2 string }
type GenS25 struct { F0 string; F1 uint }
type GenS26 struct { F0 uint16 }
type GenS27 struct { F0 uint64; F1 uint; F2 int16 }
type GenS28 struct { F0 uint }
func genFn0() int {
    for fi4 := 0; fi4 < 16; fi4++ {
        switch (("" + "") + "foo bar") {
        case "a":
            var v5 int = 9223372036854775807
            _ = v5
        case "\n\t":
            v6 := []int{fi4, fi4, fi4, (^(fi4))}
            _ = v6
        }
        var v7 int8 = -128
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 string = ("a" + "")
            v9 = ("foo bar" + (("a" + (((v9 + "a") + "foo bar") + v9)) + v9))
            v10 := GenS1{F0: (fi4 >> ((18446744073709551615) & 63)), F1: 65535, F2: 42}
            var v11 int16 = 1
            v7 -= (v7 + 100)
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi4
        _ = v7
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        if true {
            var v13 uint8 = 128
            v13 = v13
            _ = v13
        }
        var v14 bool = true
        _ = fi12
        _ = v14
    }
    var v15 uint = 4294967295
    for fi16 := 0; fi16 < 11; fi16++ {
        switch fi16 {
        case 127:
            v17 := GenS2{F0: (42 <= 93), F1: 1, F2: 1}
            _ = v17
        case 42:
            var v18 int8 = -19
            _ = v18
        default:
            if false {
                if (false && (("\n\t" + "") < "hello")) {
                    v19 := GenS3{F0: -207, F1: (fi16 - (fi16 % (((fi16 + (fi16 << (((v15 % (((v15 >> (((v15 * (^(v15)))) & 63))) | 1))) & 63)))) | 1)))}
                    var v20 int = -2147483648
                    v21 := GenS4{F0: true}
                    _ = v19
                    _ = v20
                    _ = v21
                } else {
                    switch "a" {
                    case "\n\t":
                        for fi22 := 0; fi22 < 17; fi22++ {
                            v15 += v15
                            _ = fi22
                        }
                    default:
                        v15 = v15
                    }
                    v15 -= v15
                }
                v23 := GenS5{F0: -128, F1: 0}
                _ = v23
            }
        }
        if (!true) {
            switch "" {
            case "":
                switch (^(fi16)) {
                case -2147483648:
                    fmt.Println(v15)
                default:
                    var v24 uint8 = 69
                    _ = v24
                }
            default:
                v15 += (v15 + 946)
            }
        }
        _ = fi16
    }
    v15 = 255
    _ = v15
    return -456
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 694
    v3 := []uint32{}
    var v25 string = "foo bar"
    v26 := GenS4{F0: (-128 < -1)}
    fmt.Println(v2)
    for fi27 := 0; fi27 < 1; fi27++ {
        if (1 > 0) {
            for fi28 := 0; fi28 < 10; fi28++ {
                if ((^(v2)) != 42) {
                    var v29 uint32 = (v2 & v2)
                    v30 := v1[2147483648]
                    fmt.Println(fi28, v25, fi27)
                    _ = v29
                    _ = v30
                } else {
                    v31 := GenS5{F0: 100, F1: 1}
                    var v32 uint8 = 127
                    var v33 int32 = -169
                    _ = v31
                    _ = v32
                    _ = v33
                }
                v25 = "\n\t"
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi28
            }
            v34 := GenS3{F0: -943, F1: (fi27 * (-(fi27)))}
            _ = v34
        } else {
            v35 := GenS5{F0: -128, F1: 0}
            for fi36 := 0; fi36 < 1; fi36++ {
                fmt.Println(fi36)
                var v37 float32 = 1.5
                _ = fi36
                _ = v37
            }
            if true {
                fmt.Println(v2, v25)
            }
            _ = v35
        }
        _ = fi27
    }
    if false {
        fmt.Println(v25, v2)
        var v38 uint32 = v2
        _ = v38
    } else {
        v39 := GenS2{F0: true, F1: 1, F2: 42}
        _ = v39
    }
    var v40 uint32 = (v2 % ((v2) | 1))
    switch -2147483648 {
    case -401:
        switch -1 {
        case 1:
            if false {
                v25 = v25
                if (!true) {
                    var v41 int16 = -1
                    delete(v1, 4294967295)
                    _ = v41
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            } else {
                fmt.Println(v2, v40)
                var v42 int64 = 9223372036854775807
                if len(v3) > 0 {
                    v3[0] = v2
                }
                _ = v42
            }
        case 9223372036854775807:
            v2 += (v40 | 305)
        case -1:
            if ((0.0 >= ((-2.25 / 1.0) + -2.25)) || true) {
                var v43 int = 0
                for fi44 := 0; fi44 < 15; fi44++ {
                    v2 = (v40 - (^(v2)))
                    for fi45 := 0; fi45 < 2; fi45++ {
                        var v46 float32 = (0.0 + 0.0)
                        if false {
                            var v47 string = (v25 + ((v25 + (v25 + v25)) + v25))
                            v48 := GenS5{F0: 100, F1: 985}
                            v49 := v1[v40]
                            _ = v47
                            _ = v48
                            _ = v49
                        } else {
                            if (false && (4294967295 > 18446744073709551615)) {
                                v25 = ("foo bar" + "\n\t")
                            } else {
                                v50 := GenS6{F0: (374 <= 394)}
                                if false {
                                    v51 := GenS7{F0: (v2 ^ (v2 * (^(v2)))), F1: 0, F2: (((v25 + "foo bar") + v25) + v25)}
                                    var v52 int = (v43 % (((fi44 & v43)) | 1))
                                    v53 := GenS4{F0: false}
                                    var v54 uint64 = 4294967295
                                    fmt.Println(v25, v46)
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                } else {
                                    v55 := GenS3{F0: 0, F1: genFn0()}
                                    fmt.Println(v43)
                                    _ = v55
                                }
                                _ = v50
                            }
                        }
                        v2 = 2147483648
                        var v56 int16 = 952
                        if (2 <= 0) {
                            if (false || (3.14159 >= (((3.14159 * 3.14159) + -2.25) * 1e10))) {
                                if ("foo bar" == v25) {
                                    var v57 int8 = -39
                                    _ = v57
                                } else {
                                    v2 -= 300
                                    switch fi44 {
                                    case 551:
                                        v56 -= v56
                                    }
                                }
                                switch (v25 + "\n\t") {
                                case "foo bar":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case "a":
                                    var v58 uint64 = 2
                                    _ = v58
                                case "":
                                    _ = len(v1)
                                }
                            } else {
                                var v59 uint = 65535
                                v60 := GenS2{F0: true, F1: -123, F2: 34}
                                var v61 int32 = 0
                                _ = v59
                                _ = v60
                                _ = v61
                            }
                            v1[(^(v40))] = (v40 + 1)
                            v62 := GenS2{F0: false, F1: -128, F2: 42}
                            _ = v62
                        }
                        v46 += (float32(0.1) + float32(1.5))
                        v40 -= 1
                        v43 = genFn0()
                        _ = fi45
                        _ = v46
                        _ = v56
                    }
                    var v63 int64 = 1
                    _ = fi44
                    _ = v63
                }
                _ = v43
            }
        }
    case 329:
        for fi64 := 0; fi64 < 8; fi64++ {
            var v65 string = ((v25 + v25) + "")
            if false {
                for fi66 := 0; fi66 < 12; fi66++ {
                    v2 += (v2 ^ 658)
                    v25 = v25
                    _ = len(v3)
                    _ = fi66
                }
                var v67 uint16 = 0
                _ = v67
            }
            _ = fi64
            _ = v65
        }
    default:
        for fi68 := 0; fi68 < 17; fi68++ {
            var v69 uint64 = 42
            if true {
                switch (^(fi68)) {
                case -343:
                    if (85 < 42) {
                        fmt.Println(v69, v40)
                    } else {
                        v3 = append(v3, 2)
                    }
                case 888:
                    var v70 int8 = 42
                    _ = v70
                default:
                    v69 += v69
                }
                var v71 uint16 = 1
                var v72 uint = 255
                _ = v71
                _ = v72
            } else {
                var v73 int = 127
                v3 = append(v3, v2)
                switch 42 {
                case 32768:
                    v25 = "hello"
                case 42:
                    fmt.Println(v73)
                default:
                    v25 = v25
                }
                _ = v73
            }
            _ = fi68
            _ = v69
        }
    }
    if true {
        _ = len(v1)
        switch 1 {
        case -9223372036854775808:
            for fi74 := 0; fi74 < 17; fi74++ {
                var v75 uint8 = 128
                v25 = "a"
                v3 = append(v3, 2)
                _ = fi74
                _ = v75
            }
        default:
            if false {
                v76 := GenS2{F0: true, F1: -1, F2: 214}
                if (32767 != -1) {
                    v2 += v2
                }
                fmt.Println(v2)
                _ = v76
            }
        }
    } else {
        for fi77 := 0; fi77 < 12; fi77++ {
            v78 := GenS5{F0: -76, F1: -217}
            v79 := GenS6{F0: (true || false)}
            if ((v25 + v25) < v25) {
                if (v25 <= (v25 + ((((v25 + ("hello" + v25)) + v25) + (v25 + "foo bar")) + ((v25 + v25) + v25)))) {
                    if (1.5 > 0.0) {
                        var v80 int8 = 1
                        v2 += 792
                        var v81 int32 = int32(fi77)
                        _ = v80
                        _ = v81
                    } else {
                        if true {
                            v82 := v1[2]
                            if (1 <= -1) {
                                if true {
                                    v83 := GenS8{F0: 0, F1: 405}
                                    _ = v83
                                } else {
                                    fmt.Println(v40, v82, v82)
                                    v84 := GenS4{F0: (!(0.1 == 1.5))}
                                    if (!(!(216 < 2))) {
                                        for fi85 := 0; fi85 < 6; fi85++ {
                                            v86 := GenS3{F0: 540, F1: fi77}
                                            var v87 int8 = -46
                                            _ = fi85
                                            _ = v86
                                            _ = v87
                                        }
                                        var v88 float64 = ((-2.25 * 1.5) * 0.0)
                                        _ = v88
                                    }
                                    _ = v84
                                }
                                if ((^(fi77)) >= (fi77 ^ (fi77 << ((796) & 63)))) {
                                    switch "" {
                                    case "hello":
                                        var v89 uint16 = 255
                                        _ = v89
                                    case "":
                                        v90 := GenS1{F0: (-(fi77)), F1: uint(v2), F2: 663}
                                        _ = v90
                                    }
                                    fmt.Println(v2, v2, v82)
                                }
                                for fi91 := 0; fi91 < 13; fi91++ {
                                    if len(v3) > 0 {
                                        v3[0] = 42
                                    }
                                    v40 -= (v40 - (v40 - 4294967295))
                                    v40 = 2147483648
                                    _ = fi91
                                }
                            } else {
                                var v92 string = "hello"
                                for fi93 := 0; fi93 < 15; fi93++ {
                                    v25 = ""
                                    var v94 bool = false
                                    var v95 uint = 42
                                    _ = fi93
                                    _ = v94
                                    _ = v95
                                }
                                switch "hello" {
                                case "hello":
                                    for fi96 := 0; fi96 < 10; fi96++ {
                                        v40 = v82
                                        if (!false) {
                                            var v97 uint = 2
                                            var v98 uint64 = 218
                                            v82 += v82
                                            switch v92 {
                                            case "":
                                                var v99 bool = ((v40 != 264) && ((((v2 ^ (v82 - (v2 | v2))) <= v2) && (65535 > 65535)) && (0 == 0)))
                                                _ = v99
                                            case "a":
                                                var v100 float32 = (1e10 / -2.25)
                                                _ = v100
                                            default:
                                                v40 -= 2
                                            }
                                            v101 := GenS9{F0: 32767, F1: v40}
                                            _ = v101
                                            _ = v97
                                            _ = v98
                                        }
                                        _ = fi96
                                    }
                                case "a":
                                    fmt.Println(fi77)
                                default:
                                    v1[v2] = 65535
                                }
                                _ = v92
                            }
                            v25 = "foo bar"
                            _ = v82
                        } else {
                            if (42 > 89) {
                                var v102 int64 = 1
                                _ = v102
                            }
                        }
                        var v103 uint16 = 255
                        _ = v103
                    }
                    if false {
                        var v104 uint32 = v40
                        _ = v104
                    } else {
                        v105 := GenS10{F0: 972}
                        for fi106 := 0; fi106 < 17; fi106++ {
                            switch (fi77 * -438) {
                            case 655:
                                v107 := GenS1{F0: fi77, F1: 42, F2: 2}
                                v40 += v2
                                delete(v1, (v40 >> ((32) & 63)))
                                _ = v107
                            case -9223372036854775808:
                                fmt.Println(v25)
                            case -781:
                                if (v2 > (v2 - v2)) {
                                    if (!false) {
                                        if (!((1 == 1000) && (v25 < "hello"))) {
                                            v108 := GenS11{F0: (fi106 - genFn0())}
                                            var v109 uint16 = 65535
                                            v3 = append(v3, (^(v2)))
                                            if (!false) {
                                                _ = len(v3)
                                                var v110 int32 = 0
                                                var v111 int8 = -128
                                                _ = v110
                                                _ = v111
                                            }
                                            v109 += (v109 + v109)
                                            _ = v108
                                            _ = v109
                                        } else {
                                            v112 := GenS9{F0: 32767, F1: v40}
                                            _ = v112
                                        }
                                    }
                                    fmt.Println(fi106)
                                } else {
                                    switch v25 {
                                    case "a":
                                        switch 196 {
                                        case 0:
                                            v113 := GenS1{F0: 1, F1: 4294967295, F2: 65535}
                                            _ = v113
                                        default:
                                            var v114 uint8 = 129
                                            _ = v114
                                        }
                                    case "\n\t":
                                        if (1 == -981) {
                                            var v115 uint8 = 121
                                            if ((!true) && true) {
                                                switch v25 {
                                                case "hello":
                                                    fmt.Println(v2)
                                                }
                                                v2 = (^(v40))
                                            } else {
                                                if len(v3) > 0 {
                                                    v3[0] = (v2 * 1)
                                                }
                                                if ("" >= "hello") {
                                                    fmt.Println(fi77, v115, fi106)
                                                    if true {
                                                        var v116 int64 = 42
                                                        var v117 int = (fi106 + (^(fi106)))
                                                        _ = v116
                                                        _ = v117
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = (^(v2))
                                                    }
                                                    var v118 int64 = 9223372036854775807
                                                    _ = v118
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = (v40 & v40)
                                                    }
                                                }
                                                v3 = append(v3, v40)
                                            }
                                            _ = v115
                                        } else {
                                            if (false || (v25 != "")) {
                                                v2 = v2
                                                switch "" {
                                                case "foo bar":
                                                    fmt.Println(v40, v40)
                                                case "":
                                                    switch 42 {
                                                    case -151:
                                                        v119 := GenS1{F0: genFn0(), F1: 18446744073709551615, F2: 32767}
                                                        _ = v119
                                                    case -1:
                                                        _ = len(v3)
                                                    case -2147483648:
                                                        switch v25 {
                                                        case "foo bar":
                                                            if (0.1 >= 0.0) {
                                                                var v120 int8 = 100
                                                                v40 = 42
                                                                var v121 float32 = 0.0
                                                                v121 = ((v121 / v121) - 0.0)
                                                                _ = v120
                                                                _ = v121
                                                            } else {
                                                                if ((-(fi77)) > (fi106 / (((fi106 | genFn0())) | 1))) {
                                                                    var v122 uint32 = v2
                                                                    _ = v122
                                                                } else {
                                                                    v3 = append(v3, v2)
                                                                    var v123 uint8 = 42
                                                                    v2 -= (v40 >> ((429) & 63))
                                                                    _ = v123
                                                                }
                                                                _ = len(v3)
                                                            }
                                                        case "":
                                                            v40 = v2
                                                            if false {
                                                                var v124 float64 = -2.25
                                                                _ = v124
                                                            }
                                                            if false {
                                                                v125 := GenS1{F0: (fi106 & (^(fi106))), F1: 18446744073709551615, F2: 32767}
                                                                _ = v125
                                                            }
                                                        case "a":
                                                            v126 := GenS7{F0: (v40 * (v2 & 468)), F1: 0, F2: ("a" + "a")}
                                                            _ = v126
                                                        }
                                                    }
                                                }
                                                v2 = 695
                                                if (-731 < -9223372036854775808) {
                                                    if false {
                                                        if false {
                                                            v25 = v25
                                                            fmt.Println(fi77, v25)
                                                            v40 = 1
                                                            var v127 uint32 = v40
                                                            _ = v127
                                                        }
                                                        v128 := v1[(v2 + (v2 / ((v40) | 1)))]
                                                        _ = v128
                                                    }
                                                    v129 := GenS6{F0: false}
                                                    var v130 string = v25
                                                    if true {
                                                        if true {
                                                            fmt.Println(fi77, fi106)
                                                        }
                                                        v130 = "foo bar"
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = (v2 & (v2 << ((65535) & 63)))
                                                    }
                                                    switch "hello" {
                                                    case "a":
                                                        if false {
                                                            if (0 >= 234) {
                                                                if true {
                                                                    switch -1 {
                                                                    case 2147483647:
                                                                        v40 += 517
                                                                    case 42:
                                                                        v40 -= 1
                                                                    case 32768:
                                                                        if (65535 < 65535) {
                                                                            v25 = v130
                                                                            if true {
                                                                                var v131 uint32 = 690
                                                                                _ = v131
                                                                            } else {
                                                                                v2 = v2
                                                                                v3 = append(v3, v2)
                                                                                if false {
                                                                                    fmt.Println(v40, fi106, v2)
                                                                                    var v132 int16 = 1000
                                                                                    var v133 uint8 = 255
                                                                                    _ = v132
                                                                                    _ = v133
                                                                                }
                                                                                var v134 uint32 = (^(v2))
                                                                                _ = v134
                                                                            }
                                                                        }
                                                                    }
                                                                    var v135 int16 = 32767
                                                                    _ = v135
                                                                } else {
                                                                    var v136 bool = (false || (!true))
                                                                    v137 := GenS2{F0: (v136 && (1 == -1)), F1: 20, F2: 128}
                                                                    if (1 > 1) {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = (^(v2))
                                                                        }
                                                                        v138 := GenS12{F0: -2147483648}
                                                                        v3 = append(v3, (^(v40)))
                                                                        _ = v138
                                                                    }
                                                                    _ = v136
                                                                    _ = v137
                                                                }
                                                                if (787 > 431) {
                                                                    v2 += v2
                                                                }
                                                            } else {
                                                                switch (fi77 + 791) {
                                                                case 393:
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                case 465:
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                            }
                                                            v139 := GenS5{F0: -66, F1: 0}
                                                            if (42 >= 808) {
                                                                v140 := GenS3{F0: -2147483648, F1: (fi77 ^ fi77)}
                                                                if (2 >= 127) {
                                                                    fmt.Println(fi77, v2, v130)
                                                                    fmt.Println(v2, v2)
                                                                }
                                                                _ = v140
                                                            }
                                                            _ = v139
                                                        } else {
                                                            v3 = append(v3, (v40 | v2))
                                                            if false {
                                                                if len(v3) > 0 {
                                                                    v3[0] = v2
                                                                }
                                                                v141 := GenS13{F0: fi77, F1: -715}
                                                                if (-61 == 42) {
                                                                    v40 = v40
                                                                    v142 := GenS2{F0: (18446744073709551615 <= 4294967295), F1: 100, F2: 128}
                                                                    if true {
                                                                        if (4294967295 == 18446744073709551615) {
                                                                            switch "a" {
                                                                            case "foo bar":
                                                                                _ = len(v3)
                                                                            case "\n\t":
                                                                                v130 = ""
                                                                            default:
                                                                                v143 := GenS14{F0: (225 < 32767)}
                                                                                _ = v143
                                                                            }
                                                                            v144 := GenS15{F0: 255}
                                                                            _ = v144
                                                                        } else {
                                                                            if (2 < 487) {
                                                                                var v145 float64 = float64(0.1)
                                                                                var v146 int16 = 32767
                                                                                _ = v145
                                                                                _ = v146
                                                                            } else {
                                                                                v40 = v40
                                                                                if true {
                                                                                    var v147 int64 = 744
                                                                                    _ = v147
                                                                                }
                                                                            }
                                                                            _ = len(v1)
                                                                            if len(v3) > 0 {
                                                                                v3[0] = (^(v2))
                                                                            }
                                                                        }
                                                                        v25 = "foo bar"
                                                                        v2 -= 626
                                                                    }
                                                                    _ = v142
                                                                }
                                                                _ = v141
                                                            } else {
                                                                if false {
                                                                    v148 := GenS9{F0: 32767, F1: (^(v40))}
                                                                    var v149 uint8 = 202
                                                                    if false {
                                                                        if false {
                                                                            var v150 int64 = -704
                                                                            if (42 > 732) {
                                                                                if true {
                                                                                    var v151 bool = true
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                    if ((0.1 * (float64(0.0) - (3.14159 / -2.25))) != ((0.0 - 3.14159) * 0.0)) {
                                                                                        v152 := GenS1{F0: (fi77 >> ((uint(v40)) & 63)), F1: 266, F2: 215}
                                                                                        _ = v152
                                                                                    }
                                                                                    var v153 string = (v130 + (v25 + ((v130 + ((v25 + v130) + (v130 + v25))) + "\n\t")))
                                                                                    _ = v151
                                                                                    _ = v153
                                                                                }
                                                                                if (225 >= v149) {
                                                                                    switch "a" {
                                                                                    case "":
                                                                                        fmt.Println(fi106, v150)
                                                                                    case "\n\t":
                                                                                        v149 -= v149
                                                                                        v149 -= (v149 % ((1) | 1))
                                                                                        var v154 bool = (!true)
                                                                                        _ = v154
                                                                                    case "foo bar":
                                                                                        if (-1 < 0) {
                                                                                            var v155 int32 = 42
                                                                                            switch v25 {
                                                                                            case "foo bar":
                                                                                                delete(v1, (v40 % ((v2) | 1)))
                                                                                            case "\n\t":
                                                                                                if len(v3) >= 2 {
                                                                                                    v3 = v3[1:len(v3)]
                                                                                                }
                                                                                            }
                                                                                            _ = v155
                                                                                        }
                                                                                    }
                                                                                    v156 := GenS3{F0: 1, F1: -778}
                                                                                    v149 -= (^(v149))
                                                                                    _ = v156
                                                                                } else {
                                                                                    if (int8(v2) == 127) {
                                                                                        var v157 string = v130
                                                                                        _ = v157
                                                                                    }
                                                                                }
                                                                                if (0 != -1) {
                                                                                    v158 := GenS16{F0: 0, F1: v2}
                                                                                    var v159 uint = 255
                                                                                    _ = v158
                                                                                    _ = v159
                                                                                } else {
                                                                                    v160 := GenS17{F0: 32767, F1: (v40 * 341), F2: 127}
                                                                                    switch v130 {
                                                                                    case "":
                                                                                        fmt.Println(v25, v2, v40)
                                                                                    }
                                                                                    var v161 uint32 = 1
                                                                                    _ = v160
                                                                                    _ = v161
                                                                                }
                                                                                var v162 float64 = ((-2.25 - ((1.5 - -2.25) / (0.1 * 3.14159))) * 1e10)
                                                                                _ = v162
                                                                            } else {
                                                                                delete(v1, (v40 / ((v2) | 1)))
                                                                                v1[2147483647] = 178
                                                                                if false {
                                                                                    var v163 float32 = 1e10
                                                                                    v164 := GenS18{F0: 255}
                                                                                    var v165 bool = ((v150 * (v150 ^ -9223372036854775808)) != v150)
                                                                                    var v166 uint8 = 0
                                                                                    _ = v163
                                                                                    _ = v164
                                                                                    _ = v165
                                                                                    _ = v166
                                                                                } else {
                                                                                    v167 := GenS14{F0: false}
                                                                                    _ = v167
                                                                                }
                                                                            }
                                                                            fmt.Println(v2, v2, fi106)
                                                                            _ = v150
                                                                        } else {
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            var v168 uint64 = 0
                                                                            _ = v168
                                                                        }
                                                                    } else {
                                                                        switch fi106 {
                                                                        case 4294967295:
                                                                            v2 += (v40 - v40)
                                                                        }
                                                                        if ((v149 % (((v149 << ((uint(v2)) & 63))) | 1)) >= v149) {
                                                                            v169 := GenS5{F0: 0, F1: 0}
                                                                            v149 = v149
                                                                            _ = v169
                                                                        }
                                                                        var v170 uint32 = (v40 & (v40 / ((v2) | 1)))
                                                                        _ = v170
                                                                    }
                                                                    var v171 int16 = 1000
                                                                    v25 = "foo bar"
                                                                    _ = v148
                                                                    _ = v149
                                                                    _ = v171
                                                                }
                                                                switch (fi106 << ((42) & 63)) {
                                                                case -641:
                                                                    v130 = "foo bar"
                                                                    _ = len(v3)
                                                                    var v172 float64 = 0.1
                                                                    _ = v172
                                                                default:
                                                                    v173 := GenS10{F0: 1}
                                                                    _ = v173
                                                                }
                                                                v40 += v40
                                                            }
                                                        }
                                                    case "":
                                                        var v174 uint32 = 197
                                                        _ = v174
                                                    }
                                                    if false {
                                                        var v175 uint = 18446744073709551615
                                                        v2 -= v2
                                                        _ = v175
                                                    } else {
                                                        var v176 int8 = -128
                                                        _ = v176
                                                    }
                                                    v130 = "hello"
                                                    v130 = v130
                                                    _ = v129
                                                    _ = v130
                                                } else {
                                                    switch (fi77 & (fi106 | fi106)) {
                                                    case 1:
                                                        switch "foo bar" {
                                                        case "a":
                                                            switch fi106 {
                                                            case 320:
                                                                if len(v3) > 0 {
                                                                    v3[0] = (v40 & (v40 * v2))
                                                                }
                                                            }
                                                        case "":
                                                            v177 := GenS4{F0: (!false)}
                                                            _ = v177
                                                        case "\n\t":
                                                            v178 := v1[v2]
                                                            _ = v178
                                                        }
                                                        if (0 != 74) {
                                                            v179 := GenS11{F0: (-(fi77))}
                                                            if len(v3) > 0 {
                                                                v3[0] = 418
                                                            }
                                                            _ = v179
                                                        } else {
                                                            var v180 uint = 42
                                                            v25 = ""
                                                            _ = v180
                                                        }
                                                        if false {
                                                            var v181 string = v25
                                                            var v182 int8 = 40
                                                            v1[0] = (v2 | (^(v40)))
                                                            _ = v181
                                                            _ = v182
                                                        }
                                                        if ((-1 <= 0) || false) {
                                                            if true {
                                                                if (v2 > 2147483648) {
                                                                    v40 = 443
                                                                    fmt.Println(v40)
                                                                } else {
                                                                    if (-572 == -32768) {
                                                                        v40 = (v2 * 2147483648)
                                                                    }
                                                                    v183 := GenS17{F0: 113, F1: (^(v2)), F2: 48}
                                                                    _ = v183
                                                                }
                                                            }
                                                        } else {
                                                            v40 += (v40 % (((v2 / ((0) | 1))) | 1))
                                                        }
                                                        v25 = "a"
                                                    case 0:
                                                        fmt.Println(v40, v40)
                                                    case 127:
                                                        v2 = v2
                                                    }
                                                    v184 := GenS19{F0: 0}
                                                    if true {
                                                        switch v25 {
                                                        case "":
                                                            v185 := GenS20{F0: 653, F1: -9223372036854775808}
                                                            _ = v185
                                                        case "foo bar":
                                                            fmt.Println(v25)
                                                        }
                                                        v186 := GenS5{F0: -128, F1: 120}
                                                        _ = v186
                                                    }
                                                    _ = v184
                                                }
                                            } else {
                                                var v187 int64 = -813
                                                var v188 uint8 = 0
                                                var v189 uint = 0
                                                _ = v187
                                                _ = v188
                                                _ = v189
                                            }
                                        }
                                    }
                                }
                            default:
                                fmt.Println(v25, v40, fi106)
                            }
                            if true {
                                if (!true) {
                                    if len(v3) > 0 {
                                        v3[0] = v2
                                    }
                                    v190 := GenS10{F0: 0}
                                    v191 := GenS21{F0: 1, F1: -1, F2: 824}
                                    _ = v190
                                    _ = v191
                                } else {
                                    switch 9223372036854775807 {
                                    case 992:
                                        v192 := GenS22{F0: 77, F1: 18446744073709551615}
                                        _ = v192
                                    }
                                    var v193 uint16 = 65535
                                    var v194 uint = 255
                                    _ = v193
                                    _ = v194
                                }
                            }
                            _ = fi106
                        }
                        if len(v3) > 0 {
                            v3[0] = (^(v2))
                        }
                        _ = v105
                    }
                    switch ("a" + ((v25 + v25) + ("foo bar" + (v25 + v25)))) {
                    case "\n\t":
                        switch genFn0() {
                        case -2147483648:
                            var v195 float32 = (((3.14159 + float32((float64(0.1) * 3.14159))) * 3.14159) / 0.1)
                            _ = v195
                        case 182:
                            v196 := GenS19{F0: 100}
                            _ = v196
                        case 580:
                            switch int(v40) {
                            case 128:
                                if true {
                                    v25 = ("foo bar" + v25)
                                } else {
                                    if false {
                                        v40 = v40
                                        for fi197 := 0; fi197 < 7; fi197++ {
                                            var v198 int16 = -32768
                                            _ = fi197
                                            _ = v198
                                        }
                                    }
                                }
                            default:
                                if (32767 >= -175) {
                                    var v199 uint8 = 0
                                    _ = v199
                                }
                            }
                        }
                    }
                }
            }
            _ = fi77
            _ = v78
            _ = v79
        }
        v2 = v2
        _ = len(v1)
    }
    if (!(false && (400 < 4294967295))) {
        if (9223372036854775807 == 128) {
            v40 = (v40 * (^(v2)))
            v2 = v2
            _ = len(v3)
        } else {
            fmt.Println(v25, v40)
            if (uint8(v2) <= 255) {
                v40 += 750
                v200 := GenS23{F0: 18446744073709551615, F1: 1, F2: 0}
                if false {
                    fmt.Println(v40, v2, v25)
                    fmt.Println(v40, v2, v2)
                }
                _ = v200
            } else {
                for fi201 := 0; fi201 < 2; fi201++ {
                    if (514 >= 849) {
                        v202 := GenS17{F0: 299, F1: (^(v40)), F2: 100}
                        _ = v202
                    } else {
                        var v203 uint64 = 2
                        for fi204 := 0; fi204 < 14; fi204++ {
                            if (2147483647 == -880) {
                                if true {
                                    v205 := GenS18{F0: 1}
                                    _ = v205
                                } else {
                                    v40 = 42
                                    v40 += v40
                                    _ = len(v3)
                                }
                                delete(v1, 4294967295)
                            }
                            v203 = (^(v203))
                            _ = fi204
                        }
                        v1[v2] = 1
                        _ = v203
                    }
                    _ = len(v3)
                    if false {
                        for fi206 := 0; fi206 < 3; fi206++ {
                            var v207 uint32 = v40
                            _ = fi206
                            _ = v207
                        }
                        fmt.Println(v2, v25)
                        if false {
                            v25 = v25
                            var v208 int64 = -9223372036854775808
                            v25 = ((v25 + (v25 + v25)) + "hello")
                            _ = v208
                        } else {
                            v209 := GenS16{F0: 52, F1: v2}
                            switch (fi201 << ((264) & 63)) {
                            case -238:
                                fmt.Println(v40)
                            case 585:
                                var v210 uint16 = 2
                                _ = v210
                            default:
                                switch ("" + "\n\t") {
                                case "foo bar":
                                    var v211 string = v25
                                    _ = v211
                                case "":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case "a":
                                    var v212 uint16 = 65535
                                    _ = v212
                                }
                            }
                            _ = v209
                        }
                    } else {
                        for fi213 := 0; fi213 < 8; fi213++ {
                            fmt.Println(fi201, v40)
                            _ = len(v3)
                            fmt.Println(v2)
                            _ = fi213
                        }
                    }
                    _ = fi201
                }
                if ((float32((0.0 * (-2.25 - -2.25))) + (1.5 + float32(0.0))) == -2.25) {
                    fmt.Println(v25, v25)
                } else {
                    for fi214 := 0; fi214 < 5; fi214++ {
                        v2 += v40
                        var v215 int = fi214
                        _ = fi214
                        _ = v215
                    }
                }
            }
            v216 := GenS13{F0: 9223372036854775807, F1: 1}
            _ = v216
        }
        if true {
            fmt.Println(v25, v25)
            for fi217 := 0; fi217 < 6; fi217++ {
                v2 = v2
                _ = fi217
            }
            if (128 <= 128) {
                switch 95 {
                case -229:
                    if true {
                        v25 = (((v25 + ("hello" + ("" + v25))) + v25) + v25)
                    } else {
                        if false {
                            switch ("hello" + "\n\t") {
                            case "hello":
                                if (2147483647 > 126) {
                                    fmt.Println(v2, v40, v40)
                                    if false {
                                        var v218 int32 = 478
                                        _ = v218
                                    }
                                } else {
                                    for fi219 := 0; fi219 < 1; fi219++ {
                                        if true {
                                            if false {
                                                var v220 uint32 = (v40 * (v2 * (^(v40))))
                                                for fi221 := 0; fi221 < 10; fi221++ {
                                                    if true {
                                                        v25 = "foo bar"
                                                    }
                                                    _ = fi221
                                                }
                                                _ = v220
                                            } else {
                                                var v222 float32 = ((0.0 * (-2.25 + ((0.0 - 1.5) - float32((0.0 + ((float64(float32(1e10)) + 1.5) + 0.0)))))) + 1.5)
                                                v40 = v40
                                                v25 = "hello"
                                                _ = v222
                                            }
                                            v223 := GenS9{F0: 32767, F1: 778}
                                            _ = v223
                                        }
                                        v25 = "\n\t"
                                        _ = fi219
                                    }
                                    switch genFn0() {
                                    case 2147483647:
                                        var v224 uint = 18446744073709551615
                                        _ = v224
                                    case 417:
                                        switch 23 {
                                        case -621:
                                            switch v25 {
                                            case "foo bar":
                                                if (!("hello" <= v25)) {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    for fi225 := 0; fi225 < 1; fi225++ {
                                                        v226 := GenS24{F0: v2, F1: -32768, F2: "hello"}
                                                        v40 = v2
                                                        fmt.Println(fi225, v25, v40)
                                                        _ = fi225
                                                        _ = v226
                                                    }
                                                } else {
                                                    var v227 uint32 = 2
                                                    if true {
                                                        fmt.Println(v227, v2)
                                                        for fi228 := 0; fi228 < 17; fi228++ {
                                                            switch 2147483647 {
                                                            case 71:
                                                                var v229 uint16 = 1
                                                                _ = v229
                                                            case 42:
                                                                v227 = v40
                                                            default:
                                                                switch ((v25 + ((v25 + v25) + "hello")) + ((v25 + (v25 + v25)) + "")) {
                                                                case "foo bar":
                                                                    v230 := GenS25{F0: v25, F1: 1}
                                                                    _ = v230
                                                                }
                                                            }
                                                            _ = fi228
                                                        }
                                                    }
                                                    v25 = (v25 + (((v25 + v25) + v25) + v25))
                                                    _ = v227
                                                }
                                            case "a":
                                                v2 = (v2 ^ (v2 % ((v2) | 1)))
                                            case "":
                                                if false {
                                                    for fi231 := 0; fi231 < 5; fi231++ {
                                                        if (112 != 255) {
                                                            if (32767 == 0) {
                                                                v40 = v2
                                                            }
                                                        } else {
                                                            if true {
                                                                var v232 uint32 = 2147483647
                                                                v2 += v2
                                                                _ = v232
                                                            }
                                                            v25 = v25
                                                            for fi233 := 0; fi233 < 9; fi233++ {
                                                                var v234 bool = false
                                                                _ = fi233
                                                                _ = v234
                                                            }
                                                        }
                                                        _ = len(v1)
                                                        var v235 uint32 = v40
                                                        _ = fi231
                                                        _ = v235
                                                    }
                                                    if false {
                                                        v40 = v40
                                                    } else {
                                                        v236 := v1[(v2 | v2)]
                                                        _ = v236
                                                    }
                                                }
                                            }
                                        }
                                    case 843:
                                        var v237 int64 = -9223372036854775808
                                        _ = v237
                                    }
                                    v2 = 2
                                }
                            case "\n\t":
                                switch "" {
                                case "foo bar":
                                    v238 := GenS26{F0: 65535}
                                    _ = v238
                                case "a":
                                    fmt.Println(v2, v25, v2)
                                default:
                                    for fi239 := 0; fi239 < 1; fi239++ {
                                        v2 += v2
                                        var v240 int32 = -2147483648
                                        _ = fi239
                                        _ = v240
                                    }
                                }
                            }
                        }
                        var v241 uint8 = 0
                        delete(v1, v40)
                        _ = v241
                    }
                case 4294967295:
                    fmt.Println(v40, v40, v25)
                }
            }
        }
    } else {
        if (!false) {
            var v242 string = v25
            v25 = v242
            _ = v242
        }
        v25 = v25
    }
    v243 := GenS27{F0: 18446744073709551615, F1: 255, F2: -1}
    var v244 int32 = 2147483647
    var v245 string = ""
    var v246 int64 = int64(v244)
    v40 = (v40 << ((1) & 63))
    if false {
        var v247 uint64 = 276
        for fi248 := 0; fi248 < 5; fi248++ {
            v247 -= 703
            _ = fi248
        }
        _ = v247
    }
    v249 := GenS28{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v244:", v244)
    fmt.Println("v245:", v245)
    fmt.Println("v246:", v246)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v243
    _ = v244
    _ = v245
    _ = v246
    _ = v249
    _ = v25
    _ = v26
    _ = v3
    _ = v40
}
