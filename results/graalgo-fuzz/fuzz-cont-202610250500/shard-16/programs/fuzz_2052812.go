package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 bool }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int; F1 int; F2 uint32 }
type GenS4 struct { F0 int16; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint8; F1 uint; F2 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64; F1 uint }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 string; F1 int32; F2 bool }
type GenS11 struct { F0 int8; F1 uint32 }
type GenS12 struct { F0 uint64 }

func main() {
    var v1 uint16 = 576
    v2 := make(map[int]int)
    v3 := []int{-1, 2147483647}
    var v4 uint64 = 4294967295
    if false {
        switch 2147483647 {
        case -829:
            for fi5 := 0; fi5 < 19; fi5++ {
                v6 := GenS1{F0: 1, F1: 18446744073709551615, F2: (!false)}
                v7 := GenS1{F0: fi5, F1: 599, F2: ((!true) && (2147483648 > 2))}
                _ = fi5
                _ = v6
                _ = v7
            }
        case 4294967295:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            v8 := GenS1{F0: 127, F1: 65535, F2: false}
            _ = v8
        }
        for fi9 := 0; fi9 < 18; fi9++ {
            v4 = v4
            fmt.Println(v4, fi9)
            _ = fi9
        }
        if false {
            for fi10 := 0; fi10 < 2; fi10++ {
                v1 = v1
                var v11 int16 = 220
                _ = fi10
                _ = v11
            }
        }
    } else {
        v4 += v4
    }
    if ((float64(((-2.25 + 1e10) * (float32(float64((0.1 / -2.25))) + -2.25))) - 1.5) <= (1e10 / 3.14159)) {
        var v12 uint = uint(v1)
        _ = v12
    } else {
        switch "\n\t" {
        case "foo bar":
            v4 += 192
        case "a":
            var v13 float32 = 0.0
            _ = v13
        default:
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 uint8 = 83
                fmt.Println(v4)
                _ = fi14
                _ = v15
            }
        }
    }
    _ = len(v3)
    v4 = v4
    var v16 uint16 = v1
    v2[4294967295] = -9223372036854775808
    if (!true) {
        v17 := GenS2{F0: 255, F1: false}
        _ = v17
    }
    var v18 int8 = 61
    var v19 bool = (!false)
    if (-70 <= 1) {
        v20 := GenS1{F0: 1, F1: 1, F2: (!(!(0 < -1)))}
        switch 32767 {
        case -183:
            switch 2147483647 {
            case 32768:
                var v21 int16 = 0
                _ = v21
            case 4294967295:
                for fi22 := 0; fi22 < 14; fi22++ {
                    _ = len(v3)
                    for fi23 := 0; fi23 < 10; fi23++ {
                        var v24 int16 = 32767
                        v25 := GenS1{F0: 553, F1: 18446744073709551615, F2: (v19 && v19)}
                        var v26 int = (fi23 ^ 0)
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    switch (^(fi22)) {
                    case 0:
                        var v27 int16 = 42
                        _ = v27
                    case -2147483648:
                        switch (("" + "hello") + "") {
                        case "\n\t":
                            v28 := GenS2{F0: 1, F1: (-413 > 1000)}
                            _ = v28
                        case "":
                            if ((((v19 || (-343 == 301)) && (858 < 0)) && v19) || (769 >= 224)) {
                                v3 = append(v3, (fi22 * fi22))
                            } else {
                                if true {
                                    v29 := GenS3{F0: (fi22 << ((65535) & 63)), F1: 42, F2: 2}
                                    v1 = v16
                                    v30 := GenS4{F0: 1, F1: 0, F2: v16}
                                    _ = v29
                                    _ = v30
                                }
                                if true {
                                    for fi31 := 0; fi31 < 18; fi31++ {
                                        var v32 uint = 65535
                                        var v33 uint = 126
                                        var v34 string = "foo bar"
                                        _ = fi31
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    }
                                    fmt.Println(v4)
                                    v35 := GenS1{F0: (fi22 % ((0) | 1)), F1: 255, F2: v19}
                                    _ = v35
                                }
                                switch (("foo bar" + "") + ("a" + ("foo bar" + (((("\n\t" + "hello") + "hello") + "foo bar") + "foo bar")))) {
                                case "hello":
                                    switch fi22 {
                                    case -1:
                                        for fi36 := 0; fi36 < 5; fi36++ {
                                            if (!(!(!v19))) {
                                                var v37 int8 = 127
                                                var v38 uint64 = v4
                                                _ = v37
                                                _ = v38
                                            }
                                            var v39 uint = 485
                                            if false {
                                                switch -1 {
                                                case -1:
                                                    var v40 uint = (v39 >> (((v39 / ((v39) | 1))) & 63))
                                                    _ = v40
                                                case 963:
                                                    v3 = append(v3, fi22)
                                                case 118:
                                                    if (255 <= 147) {
                                                        v41 := GenS5{F0: 192, F1: 255, F2: 1}
                                                        v42 := GenS1{F0: -804, F1: (v39 & (v39 * v39)), F2: true}
                                                        var v43 uint16 = (v16 * uint16(v39))
                                                        _ = v41
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                }
                                                var v44 uint = v39
                                                _ = v44
                                            } else {
                                                switch "hello" {
                                                case "foo bar":
                                                    if (!(128 >= 0)) {
                                                        var v45 uint = v39
                                                        if true {
                                                            v45 = (v45 * v45)
                                                            var v46 uint32 = 2
                                                            _ = v46
                                                        } else {
                                                            if v19 {
                                                                var v47 int64 = 1
                                                                var v48 uint8 = 86
                                                                fmt.Println(fi22, v48)
                                                                if v19 {
                                                                    fmt.Println(v45, v19, v19)
                                                                }
                                                                v3 = append(v3, (fi22 << ((v45) & 63)))
                                                                _ = v47
                                                                _ = v48
                                                            } else {
                                                                if (v19 && (0.1 == (float32(0.1) - (((1.5 + -2.25) / 3.14159) / 0.1)))) {
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        if false {
                                                                            v1 -= v16
                                                                            var v49 string = ("foo bar" + "")
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            v45 = v39
                                                                            v1 = v16
                                                                            _ = v49
                                                                        }
                                                                    case "":
                                                                        if true {
                                                                            if v19 {
                                                                                var v50 float32 = float32(-2.25)
                                                                                var v51 bool = false
                                                                                var v52 int32 = -535
                                                                                _ = v50
                                                                                _ = v51
                                                                                _ = v52
                                                                            }
                                                                            if len(v3) > 0 {
                                                                                v3[0] = fi36
                                                                            }
                                                                        }
                                                                    case "foo bar":
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            v53 := GenS1{F0: fi22, F1: v45, F2: (!((!v19) && v19))}
                                                                            _ = v53
                                                                        }
                                                                    }
                                                                }
                                                                var v54 int64 = -9223372036854775808
                                                                _ = v54
                                                            }
                                                            var v55 bool = (-945 >= -826)
                                                            _ = v55
                                                        }
                                                        delete(v2, (fi36 % (((-(fi36))) | 1)))
                                                        _ = v45
                                                    } else {
                                                        fmt.Println(v39, v18, v16)
                                                        if ("foo bar" <= ("" + "")) {
                                                            if v19 {
                                                                if true {
                                                                    v18 -= 29
                                                                }
                                                            } else {
                                                                if true {
                                                                    if (v18 == -128) {
                                                                        v39 = v39
                                                                    }
                                                                    fmt.Println(v19, v4)
                                                                    _ = len(v3)
                                                                    v4 = (v4 / (((v4 & v4)) | 1))
                                                                    var v56 int32 = -2147483648
                                                                    _ = v56
                                                                } else {
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                                var v57 uint8 = 0
                                                                v16 = v1
                                                                _ = v57
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = (fi36 / ((fi36) | 1))
                                                            }
                                                            var v58 bool = v19
                                                            v1 += v16
                                                            switch "foo bar" {
                                                            case "hello":
                                                                switch (fi36 >> (((v39 / ((v39) | 1))) & 63)) {
                                                                case 42:
                                                                    switch (^(fi36)) {
                                                                    case -663:
                                                                        if ((!(173 != 1)) && (580 < uint32(v16))) {
                                                                            var v59 float32 = 0.0
                                                                            _ = v59
                                                                        } else {
                                                                            var v60 bool = false
                                                                            _ = v60
                                                                        }
                                                                    case 2147483647:
                                                                        fmt.Println(v58)
                                                                    default:
                                                                        var v61 float64 = ((((0.1 + 1e10) / 1.5) * float64((1e10 - 0.1))) / -2.25)
                                                                        _ = v61
                                                                    }
                                                                case -1:
                                                                    switch ("\n\t" + "foo bar") {
                                                                    case "a":
                                                                        v62 := GenS6{F0: 0}
                                                                        _ = v62
                                                                    }
                                                                }
                                                            case "a":
                                                                var v63 int16 = -1
                                                                _ = v63
                                                            default:
                                                                fmt.Println(fi36, fi22)
                                                            }
                                                            _ = v58
                                                        } else {
                                                            delete(v2, fi22)
                                                        }
                                                    }
                                                case "a":
                                                    var v64 float64 = (0.0 / -2.25)
                                                    _ = v64
                                                default:
                                                    var v65 float32 = 1.5
                                                    _ = v65
                                                }
                                                v66 := GenS4{F0: 0, F1: -2147483648, F2: 2}
                                                var v67 uint32 = 525
                                                _ = v66
                                                _ = v67
                                            }
                                            var v68 int8 = (v18 >> ((v39) & 63))
                                            _ = fi36
                                            _ = v39
                                            _ = v68
                                        }
                                    case -2147483648:
                                        var v69 int8 = v18
                                        _ = v69
                                    default:
                                        if true {
                                            if (-1 != -760) {
                                                _ = len(v2)
                                                v70 := GenS7{F0: 42, F1: 255}
                                                _ = v70
                                            }
                                        } else {
                                            var v71 int32 = 560
                                            v19 = (0 <= 0)
                                            if false {
                                                var v72 int16 = 32767
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v72
                                            }
                                            _ = v71
                                        }
                                    }
                                case "\n\t":
                                    switch "hello" {
                                    case "":
                                        v3 = append(v3, 32767)
                                    case "a":
                                        if (v19 || (!true)) {
                                            if (!v19) {
                                                v16 = (v1 % ((42) | 1))
                                                fmt.Println(v4, v18, v19)
                                            } else {
                                                for fi73 := 0; fi73 < 5; fi73++ {
                                                    var v74 uint8 = 1
                                                    if (v4 >= (v4 | 1)) {
                                                        var v75 float32 = 1.5
                                                        _ = v75
                                                    } else {
                                                        var v76 uint = 2
                                                        _ = v76
                                                    }
                                                    var v77 bool = ((-1 <= v18) || v19)
                                                    _ = fi73
                                                    _ = v74
                                                    _ = v77
                                                }
                                                v2[710] = (-(fi22))
                                            }
                                        }
                                    }
                                case "a":
                                    switch fi22 {
                                    case 1:
                                        switch ("" + "\n\t") {
                                        case "a":
                                            if v19 {
                                                if (!false) {
                                                    v4 = 4294967295
                                                    var v78 float64 = 1.5
                                                    if (!v19) {
                                                        for fi79 := 0; fi79 < 10; fi79++ {
                                                            v16 = 42
                                                            var v80 int = fi22
                                                            if v19 {
                                                                if v19 {
                                                                    v18 += (-(v18))
                                                                }
                                                                if v19 {
                                                                    v78 -= float64(3.14159)
                                                                    v2[fi22] = -2147483648
                                                                    v3 = append(v3, 127)
                                                                    v1 = v16
                                                                    v18 = v18
                                                                } else {
                                                                    v19 = (v19 && v19)
                                                                    v81 := GenS3{F0: v80, F1: 42, F2: 422}
                                                                    _ = v81
                                                                }
                                                                fmt.Println(v19, v16, v16)
                                                                if v19 {
                                                                    if true {
                                                                        v80 = (fi79 ^ v80)
                                                                        v18 += -128
                                                                    }
                                                                    if (int16(fi79) < 0) {
                                                                        var v82 int16 = -32768
                                                                        if (v19 || false) {
                                                                            var v83 int64 = 138
                                                                            if v19 {
                                                                                v84 := GenS2{F0: 928, F1: (1 != 808)}
                                                                                _ = len(v3)
                                                                                _ = v84
                                                                            } else {
                                                                                if ((100 > (v18 >> ((42) & 63))) || (!("" > "\n\t"))) {
                                                                                    var v85 uint32 = 834
                                                                                    v85 += (v85 >> ((42) & 63))
                                                                                    var v86 int32 = -1
                                                                                    _ = v85
                                                                                    _ = v86
                                                                                }
                                                                                v78 = v78
                                                                                var v87 int = (fi79 % ((v80) | 1))
                                                                                _ = v87
                                                                            }
                                                                            _ = v83
                                                                        }
                                                                        fmt.Println(v82, v80, v78)
                                                                        _ = v82
                                                                    }
                                                                }
                                                                v16 += 1
                                                                switch ("foo bar" + "a") {
                                                                case "\n\t":
                                                                    fmt.Println(v18, v78, v80)
                                                                default:
                                                                    v1 = 490
                                                                }
                                                                v88 := GenS8{F0: -1, F1: v18}
                                                                _ = v88
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = fi79
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = v80
                                                            }
                                                            if true {
                                                                switch -9223372036854775808 {
                                                                case -755:
                                                                    v19 = (false && (v19 || v19))
                                                                default:
                                                                    v16 = v1
                                                                }
                                                                v80 += v80
                                                                _ = len(v3)
                                                            } else {
                                                                if v19 {
                                                                    switch fi22 {
                                                                    case 4294967295:
                                                                        switch 0 {
                                                                        case -1:
                                                                            v4 = 399
                                                                        default:
                                                                            var v89 uint8 = 183
                                                                            _ = v89
                                                                        }
                                                                    case -880:
                                                                        if v19 {
                                                                            switch fi22 {
                                                                            case 0:
                                                                                v4 += v4
                                                                            default:
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    if v19 {
                                                                                        fmt.Println(fi79, v1)
                                                                                        var v90 float64 = v78
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                        _ = v90
                                                                                    }
                                                                                case "hello":
                                                                                    v91 := GenS3{F0: fi22, F1: (v80 >> ((65535) & 63)), F2: 2147483647}
                                                                                    _ = v91
                                                                                case "":
                                                                                    var v92 uint8 = uint8(v4)
                                                                                    _ = v92
                                                                                default:
                                                                                    if (243 == 711) {
                                                                                        v4 -= (^(v4))
                                                                                    } else {
                                                                                        var v93 float32 = 1e10
                                                                                        v94 := GenS9{F0: v80}
                                                                                        var v95 float32 = v93
                                                                                        _ = v93
                                                                                        _ = v94
                                                                                        _ = v95
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v96 int16 = -279
                                                                            var v97 uint64 = 1
                                                                            _ = v96
                                                                            _ = v97
                                                                        } else {
                                                                            var v98 int64 = -930
                                                                            var v99 uint64 = (v4 * 722)
                                                                            if false {
                                                                                v1 = 689
                                                                                v100 := GenS5{F0: 255, F1: uint(v1), F2: -32768}
                                                                                _ = v100
                                                                            } else {
                                                                                switch "a" {
                                                                                case "":
                                                                                    var v101 int64 = (v98 / ((v98) | 1))
                                                                                    v102 := GenS5{F0: 185, F1: uint(v16), F2: 1000}
                                                                                    var v103 int32 = 868
                                                                                    _ = v101
                                                                                    _ = v102
                                                                                    _ = v103
                                                                                case "a":
                                                                                    var v104 int16 = 1
                                                                                    _ = v104
                                                                                default:
                                                                                    fmt.Println(v98)
                                                                                }
                                                                                var v105 uint64 = v4
                                                                                var v106 int64 = (v98 / ((v98) | 1))
                                                                                _ = v105
                                                                                _ = v106
                                                                            }
                                                                            _ = v98
                                                                            _ = v99
                                                                        }
                                                                    default:
                                                                        var v107 float64 = v78
                                                                        _ = v107
                                                                    }
                                                                    v3 = append(v3, fi22)
                                                                } else {
                                                                    switch fi22 {
                                                                    case -9223372036854775808:
                                                                        v108 := GenS1{F0: 127, F1: 433, F2: v19}
                                                                        fmt.Println(fi79, fi22, fi79)
                                                                        _ = v108
                                                                    default:
                                                                        fmt.Println(v80, v18, fi79)
                                                                    }
                                                                    var v109 int8 = v18
                                                                    _ = v109
                                                                }
                                                                v80 = (fi22 - (v80 | fi79))
                                                                delete(v2, fi22)
                                                            }
                                                            var v110 float32 = 3.14159
                                                            _ = fi79
                                                            _ = v110
                                                            _ = v80
                                                        }
                                                    } else {
                                                        v2[(^(fi22))] = fi22
                                                    }
                                                    _ = v78
                                                } else {
                                                    v19 = v19
                                                    v4 -= (v4 + (v4 >> ((1) & 63)))
                                                }
                                            } else {
                                                var v111 bool = v19
                                                v111 = (v19 && v111)
                                                _ = v111
                                            }
                                        case "\n\t":
                                            if (v19 && (65535 <= 446)) {
                                                _ = len(v2)
                                                v16 -= 1
                                                if v19 {
                                                    _ = len(v2)
                                                } else {
                                                    var v112 int16 = -1
                                                    for fi113 := 0; fi113 < 16; fi113++ {
                                                        var v114 float64 = 1e10
                                                        fmt.Println(v19)
                                                        _ = fi113
                                                        _ = v114
                                                    }
                                                    _ = v112
                                                }
                                            }
                                        case "hello":
                                            for fi115 := 0; fi115 < 13; fi115++ {
                                                if false {
                                                    v116 := v2[int(v4)]
                                                    v117 := GenS10{F0: "hello", F1: 0, F2: v19}
                                                    fmt.Println(v16, v16)
                                                    _ = v116
                                                    _ = v117
                                                } else {
                                                    v16 = uint16(v4)
                                                    v1 = v16
                                                }
                                                _ = fi115
                                            }
                                        default:
                                            for fi118 := 0; fi118 < 12; fi118++ {
                                                v119 := GenS8{F0: -117, F1: (v18 << ((628) & 63))}
                                                _ = fi118
                                                _ = v119
                                            }
                                        }
                                    }
                                default:
                                    var v120 uint32 = 493
                                    _ = v120
                                }
                            }
                        case "foo bar":
                            v18 -= v18
                        }
                    }
                    _ = fi22
                }
            default:
                var v121 uint32 = 658
                _ = v121
            }
        case -1:
            for fi122 := 0; fi122 < 20; fi122++ {
                v123 := GenS8{F0: 379, F1: (v18 | (^(v18)))}
                _ = fi122
                _ = v123
            }
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = len(v2)
        _ = v20
    }
    var v124 uint = 4294967295
    var v125 int32 = -1
    if (201 > 128) {
        v126 := GenS3{F0: 2147483647, F1: -1, F2: 847}
        _ = v126
    } else {
        v4 = 4294967295
        var v127 uint64 = (v4 / ((v4) | 1))
        _ = v127
    }
    if (1000 == -868) {
        var v128 string = "hello"
        delete(v2, 2147483647)
        _ = v128
    } else {
        v19 = v19
        if (((!(1 <= (^(v124)))) || (float64((0.0 / 1e10)) == (1.5 + 0.1))) && v19) {
            var v129 float32 = float32((0.1 + (float64(1e10) + ((1.5 / ((1e10 + 1e10) + (((float64((0.1 / 1.0)) - (0.0 - 1.5)) - 1e10) + (0.1 - 1.5)))) * 0.1))))
            v130 := GenS5{F0: 1, F1: 42, F2: 1}
            var v131 float64 = 0.0
            _ = v129
            _ = v130
            _ = v131
        } else {
            for fi132 := 0; fi132 < 2; fi132++ {
                var v133 int16 = 32767
                if v19 {
                    v134 := GenS11{F0: 0, F1: 0}
                    _ = v134
                } else {
                    fmt.Println(v1, v19, v125)
                    fmt.Println(v1, v1, v124)
                    v19 = (!(((255 == 51) || v19) || v19))
                }
                v18 = (v18 / (((v18 >> ((v124) & 63))) | 1))
                _ = fi132
                _ = v133
            }
            if true {
                for fi135 := 0; fi135 < 3; fi135++ {
                    v19 = true
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi135
                }
                v16 += 255
                for fi136 := 0; fi136 < 4; fi136++ {
                    v18 -= v18
                    _ = fi136
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = -803
                }
                v137 := GenS12{F0: v4}
                v4 = 4294967295
                _ = v137
            }
            for fi138 := 0; fi138 < 5; fi138++ {
                if v19 {
                    var v139 float32 = 3.14159
                    var v140 uint16 = (v1 + (v1 / (((v1 * v16)) | 1)))
                    for fi141 := 0; fi141 < 5; fi141++ {
                        if v19 {
                            if v19 {
                                fmt.Println(fi138, v18)
                            } else {
                                v2[860] = fi138
                            }
                            v4 += (^(v4))
                            v139 = v139
                        }
                        var v142 float32 = (v139 * float32((3.14159 * 3.14159)))
                        _ = fi141
                        _ = v142
                    }
                    _ = v139
                    _ = v140
                } else {
                    for fi143 := 0; fi143 < 12; fi143++ {
                        v144 := GenS4{F0: 1, F1: (v125 / (((v125 >> ((v124) & 63))) | 1)), F2: 32767}
                        _ = fi143
                        _ = v144
                    }
                }
                for fi145 := 0; fi145 < 17; fi145++ {
                    switch (((("a" + "hello") + "foo bar") + "a") + "foo bar") {
                    case "":
                        var v146 int32 = 1
                        _ = v146
                    default:
                        v147 := v2[fi145]
                        _ = v147
                    }
                    fmt.Println(v4, v19)
                    _ = fi145
                }
                _ = fi138
            }
        }
    }
    if len(v3) > 0 {
        v3[0] = 9223372036854775807
    }
    v124 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v124:", v124)
    fmt.Println("v125:", v125)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v124
    _ = v125
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
